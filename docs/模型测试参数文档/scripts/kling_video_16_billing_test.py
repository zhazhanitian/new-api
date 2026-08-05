#!/usr/bin/env python3
"""kling-video-1.6 扣费测试脚本

规则（严格成本控制）：
  - 一次命令只允许 1 个用例：必须 --case TC-XX，禁止批量/逗号分隔/多 ID
  - 禁止并行 live：进程锁保证同一时刻只有 1 个 --run 在跑
  - 禁止重复提交：同一用例 live 过则拒绝再次 POST，除非 --force
  - 允许轮询查任务、查日志/额度（用于拿结果和核对扣费）
  - 每次 --run --case：1 次 POST 提交 + 轮询至完成 + 查扣费

用法：
  python3 kling_video_16_billing_test.py --dry-run --case TC-02   # 只推演，不提交
  python3 kling_video_16_billing_test.py --run --case TC-02       # 提交+轮询+写报告
  python3 kling_video_16_billing_test.py --run --case TC-02 --force  # 强制重提交
"""

import argparse
import atexit
import json
import os
import subprocess
import sys
import time
from dataclasses import dataclass, field
from datetime import datetime
from typing import Optional

BASE_URL = "http://localhost:9006"
API_KEY = "sk-bekfGkjueBjiia7A8JiKlfSJX3HKI7fMe8LupUr4Xn0P5xda"
MODEL = "kling-video-1.6"
PROMPT = "夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格"
LANDSCAPE = "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"
FX = 7.3
QUOTA_PER_USD = 500_000
POLL_INTERVAL = 15
POLL_MAX_WAIT = 600

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
STATE_FILE = os.path.join(SCRIPT_DIR, ".kling_video_16_executed.json")
LOCK_FILE = os.path.join(SCRIPT_DIR, ".kling_video_16_run.lock")
REPORT_PATH = os.path.join(os.path.dirname(SCRIPT_DIR), "kling-video-1.6扣费测试报告.md")

EXPR_UNIT = {"720P": 54795, "1080P": 95890, "2K": 136986, "4K": 205479}


@dataclass
class Case:
    id: str
    name: str
    body: dict
    expect_http: int = 200
    expect_submit: bool = True
    notes: str = ""
    billing_overrides: dict = field(default_factory=dict)


CASES = [
    Case("TC-01", "最简请求（不传 duration/size）", {"model": MODEL, "prompt": PROMPT},
         notes="适配器默认 720P；duration 不传时 API 默认 5s"),
    Case("TC-02", "标准：duration=10 + size=1080P",
         {"model": MODEL, "prompt": PROMPT, "duration": 10, "size": "1080P"}),
    Case("TC-03", "边界：duration=3（下限）",
         {"model": MODEL, "prompt": PROMPT, "duration": 3, "size": "720P"}),
    Case("TC-04", "边界：duration=15（上限）",
         {"model": MODEL, "prompt": PROMPT, "duration": 15, "size": "720P"}),
    Case("TC-05", "非法 duration=1（应 clamp 到 3）",
         {"model": MODEL, "prompt": PROMPT, "duration": 1},
         notes="表达式按 clamp 后 d=3 计费"),
    Case("TC-06", "非法 duration=20（应 clamp 到 15）",
         {"model": MODEL, "prompt": PROMPT, "duration": 20}),
    Case("TC-07", "仅 size=1080P，不传 duration",
         {"model": MODEL, "prompt": PROMPT, "size": "1080P"}),
    Case("TC-08", "像素 size=1920x1080",
         {"model": MODEL, "prompt": PROMPT, "duration": 5, "size": "1920x1080"}),
    Case("TC-09", "metadata.resolution=1080P 替代 size",
         {"model": MODEL, "prompt": PROMPT, "duration": 5, "metadata": {"resolution": "1080P"}}),
    Case("TC-10", "错传 size=4K（表达式走 4K 档）",
         {"model": MODEL, "prompt": PROMPT, "duration": 5, "size": "4K"},
         notes="表达式按 4K 计费；适配器实际可能只发 720P/1080P"),
    Case("TC-11", "空 prompt（应拒绝）", {"model": MODEL, "prompt": "   "},
         expect_http=400, expect_submit=False, notes="校验失败不应提交上游"),
    Case("TC-12", "duration 字符串 \"10\"",
         {"model": MODEL, "prompt": PROMPT, "duration": "10", "size": "1080P"}),
    Case("TC-13", "多传无关 metadata + 参考图",
         {"model": MODEL, "prompt": PROMPT, "duration": 5, "size": "720P",
          "images": [LANDSCAPE], "metadata": {"audio_generation": "Enabled", "foo": "bar"}},
         notes="kling 1.6 表达式无 audio/ref 分支"),
    Case("TC-14", "显式 duration=0", {"model": MODEL, "prompt": PROMPT, "duration": 0},
         notes="d0=0 应按默认 5 秒计费"),
]


def fmt_num(v) -> str:
    if isinstance(v, (int, float)):
        return f"{int(v):,}"
    return str(v)


def clamp_duration(d: int) -> int:
    if d <= 0:
        return 5
    if d < 3:
        return 3
    if d > 15:
        return 15
    return d


def resolve_size_token(size: Optional[str], resolution: Optional[str]) -> str:
    s = size or resolution or ""
    if not s:
        return "720P"
    sl = s.lower()
    if "4k" in sl:
        return "4K"
    if "2k" in sl:
        return "2K"
    if "1080" in sl:
        return "1080P"
    if "x" in sl:
        parts = s.split("x")
        if len(parts) == 2:
            try:
                w, h = int(parts[0]), int(parts[1])
                if min(w, h) >= 1080:
                    return "1080P"
            except ValueError:
                pass
    return "720P"


def expected_billing(body: dict) -> dict:
    d0 = body.get("duration")
    if d0 is None:
        d0 = 0
    elif isinstance(d0, str):
        d0 = int(d0) if d0.isdigit() else 0
    d = clamp_duration(int(d0))
    meta = body.get("metadata") or {}
    tier = resolve_size_token(body.get("size"), meta.get("resolution"))
    unit = EXPR_UNIT[tier]
    expr_val = unit * d
    usd = expr_val / 1_000_000
    return {
        "duration_input": body.get("duration"),
        "duration_billed": d,
        "size_input": body.get("size"),
        "resolution_input": meta.get("resolution"),
        "tier": tier,
        "unit_per_sec": unit,
        "expression_value": expr_val,
        "expected_usd": round(usd, 6),
        "expected_rmb": round(usd * FX, 2),
        "expected_quota": round(expr_val / 1_000_000 * QUOTA_PER_USD),
    }


def curl(method: str, path: str, body: Optional[dict] = None, timeout: int = 120) -> dict:
    cmd = [
        "curl", "-s", "--noproxy", "*",
        "-w", "\n__HTTP_CODE__:%{http_code}\n__TIME_TOTAL__:%{time_total}",
        "-X", method, f"{BASE_URL}{path}",
        "-H", f"Authorization: Bearer {API_KEY}",
        "-H", "Content-Type: application/json",
    ]
    if body is not None:
        cmd += ["-d", json.dumps(body, ensure_ascii=False)]
    proc = subprocess.run(cmd, capture_output=True, text=True, timeout=timeout)
    raw = proc.stdout
    http_code, time_total = 0, 0.0
    if "__HTTP_CODE__:" in raw:
        main, meta = raw.rsplit("__HTTP_CODE__:", 1)
        code_part, time_part = meta.split("__TIME_TOTAL__:", 1)
        http_code = int(code_part.strip())
        time_total = float(time_part.strip())
        raw = main
    try:
        parsed = json.loads(raw.strip() or "{}")
    except json.JSONDecodeError:
        parsed = {"_raw": raw.strip()}
    return {"http_code": http_code, "time_total": time_total, "json": parsed}


def load_state() -> dict:
    if not os.path.exists(STATE_FILE):
        return {"executed": {}}
    with open(STATE_FILE, encoding="utf-8") as f:
        return json.load(f)


def save_state(state: dict) -> None:
    with open(STATE_FILE, "w", encoding="utf-8") as f:
        json.dump(state, f, ensure_ascii=False, indent=2)


def _pid_alive(pid: int) -> bool:
    try:
        os.kill(pid, 0)
        return True
    except OSError:
        return False


def acquire_run_lock() -> None:
    """保证同一时刻只有 1 个 --run 在执行，避免并行提交干扰结果。"""
    if os.path.exists(LOCK_FILE):
        try:
            with open(LOCK_FILE, encoding="utf-8") as f:
                meta = json.load(f)
            old_pid = int(meta.get("pid", 0))
            if old_pid and _pid_alive(old_pid):
                print(
                    f"拒绝并行执行：已有 live 测试在跑\n"
                    f"  pid={old_pid} case={meta.get('case')} since={meta.get('at')}\n"
                    f"  请等当前用例完成后再测下一个，禁止一次性提交多个用例",
                    file=sys.stderr,
                )
                sys.exit(2)
        except (json.JSONDecodeError, ValueError, OSError):
            pass  # 锁文件损坏，覆盖重建
    with open(LOCK_FILE, "w", encoding="utf-8") as f:
        json.dump(
            {
                "pid": os.getpid(),
                "at": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
                "case": os.environ.get("KLING_TEST_CASE", ""),
            },
            f,
        )
    atexit.register(release_run_lock)


def release_run_lock() -> None:
    try:
        if os.path.exists(LOCK_FILE):
            with open(LOCK_FILE, encoding="utf-8") as f:
                meta = json.load(f)
            if meta.get("pid") == os.getpid():
                os.remove(LOCK_FILE)
    except (json.JSONDecodeError, OSError, ValueError):
        try:
            os.remove(LOCK_FILE)
        except OSError:
            pass


def parse_single_case_id(raw: str) -> str:
    """强制只接受 1 个用例 ID，拒绝批量写法。"""
    s = (raw or "").strip()
    if not s:
        print("拒绝：必须指定 --case，且一次只能 1 个用例", file=sys.stderr)
        sys.exit(2)
    # 常见批量写法：逗号/空格/分号/斜杠分隔
    for sep in (",", ";", "|", "/", " ", "\t"):
        if sep in s:
            parts = [p for p in s.replace(",", " ").replace(";", " ").replace("|", " ").split() if p]
            if len(parts) > 1:
                print(
                    f"拒绝批量提交：检测到多个用例 {parts}\n"
                    f"  一次命令只能测 1 个用例，请改为：--case {parts[0]}\n"
                    f"  测完再单独执行下一个，避免干扰扣费结果",
                    file=sys.stderr,
                )
                sys.exit(2)
            s = parts[0] if parts else s
            break
    return s.upper() if s.upper().startswith("TC-") else s


def poll_task(task_id: str) -> dict:
    start = time.time()
    last = {}
    polls = 0
    while time.time() - start < POLL_MAX_WAIT:
        polls += 1
        r = curl("GET", f"/v1/video/generations/{task_id}", timeout=60)
        last = r["json"]
        status = last.get("status", "")
        if status in ("succeeded", "failed", "cancelled"):
            return {
                "final_status": status,
                "poll_count": polls,
                "poll_seconds": round(time.time() - start, 1),
                "last_response": last,
            }
        time.sleep(POLL_INTERVAL)
    return {
        "final_status": "timeout",
        "poll_count": polls,
        "poll_seconds": round(time.time() - start, 1),
        "last_response": last,
    }


def append_report(text: str, path: str) -> None:
    header = not os.path.exists(path)
    with open(path, "a", encoding="utf-8") as f:
        if header:
            f.write("# kling-video-1.6 扣费测试报告\n\n")
        f.write(text)
        if not text.endswith("\n"):
            f.write("\n")
        f.write("\n---\n\n")


def build_report(case: Case, exp: dict, *, dry_run: bool, live: Optional[dict] = None) -> str:
    now = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    lines = [
        f"## {case.id}: {case.name} @ {now}",
        "",
    ]
    if case.notes:
        lines += [f"> {case.notes}", ""]
    lines += [
        "### 调用参数",
        "```json",
        json.dumps(case.body, ensure_ascii=False, indent=2),
        "```",
        "",
        "### 测试价格变量",
        "| 字段 | 值 |",
        "| --- | --- |",
        f"| duration 入参 | {exp['duration_input']} |",
        f"| duration 计费用 | {exp['duration_billed']} 秒 |",
        f"| size 入参 | {exp['size_input'] or '(未传)'} |",
        f"| metadata.resolution | {exp['resolution_input'] or '(未传)'} |",
        f"| 命中分辨率档 | {exp['tier']} |",
        f"| 预期表达式值 | {exp['expression_value']:,} |",
        f"| 预期 USD | ${exp['expected_usd']:.3f} |",
        f"| 预期 RMB | ¥{exp['expected_rmb']:.2f} |",
        f"| **预期 Quota** | **{exp['expected_quota']:,}** |",
        "",
    ]
    if dry_run:
        lines += ["> dry-run：未提交任务", ""]
        return "\n".join(lines)

    assert live is not None
    submit = live["submit"]
    poll = live.get("poll")
    billing = live.get("billing", {})

    lines += [
        "### 提交结果",
        f"- HTTP: {submit['http_code']}（预期 {case.expect_http}）",
        f"- 提交耗时: {submit['time_total']:.3f}s",
        f"- task_id: `{live.get('task_id', '-')}`",
        "",
        "```json",
        json.dumps(submit["json"], ensure_ascii=False, indent=2),
        "```",
        "",
        "### 扣费分析",
        "| 指标 | 值 |",
        "| --- | --- |",
        f"| 提交前可用额度 | {fmt_num(billing.get('avail_before', '-'))} |",
        f"| 提交后可用额度 | {fmt_num(billing.get('avail_after', '-'))} |",
        f"| **实际扣减 Quota** | **{fmt_num(billing.get('actual_deduct', '-'))}** |",
        f"| 消费日志 quota | {fmt_num(billing.get('log_quota', '-'))} |",
        f"| request_id | {billing.get('request_id', '-')} |",
        f"| 扣费是否符合预期 | {'✅' if billing.get('match') else '❌'} |",
        "",
    ]
    if poll:
        lines += [
            "### 任务轮询",
            f"- 轮询次数: {poll['poll_count']}",
            f"- 轮询耗时: {poll['poll_seconds']}s",
            f"- 最终状态: **{poll['final_status']}**",
            "",
        ]
        meta = (poll.get("last_response") or {}).get("metadata") or {}
        if meta.get("url"):
            lines += [f"- 视频 URL: {meta['url']}", ""]
        err = (poll.get("last_response") or {}).get("error")
        if err:
            lines += ["```json", json.dumps(err, ensure_ascii=False, indent=2), "```", ""]
    return "\n".join(lines)


def main() -> int:
    p = argparse.ArgumentParser(
        description="kling-video-1.6 测试（一次只跑 1 个用例；禁止批量/并行提交）",
        epilog="示例：python3 kling_video_16_billing_test.py --run --case TC-02",
    )
    mode = p.add_mutually_exclusive_group(required=True)
    mode.add_argument("--dry-run", action="store_true", help="只推演，不提交")
    mode.add_argument("--run", action="store_true", help="提交 1 个用例 + 轮询 + 写报告")
    p.add_argument(
        "--case",
        required=True,
        metavar="TC-XX",
        help="用例 ID（必填，且只能 1 个，如 TC-02；禁止 TC-01,TC-02）",
    )
    p.add_argument("--force", action="store_true", help="强制重新提交已执行过的用例")
    p.add_argument("--report", default=REPORT_PATH)
    args = p.parse_args()

    case_id = parse_single_case_id(args.case)
    cases = [c for c in CASES if c.id == case_id]
    if not cases:
        known = ", ".join(c.id for c in CASES)
        print(f"未知用例: {case_id}\n可用: {known}", file=sys.stderr)
        return 2
    case = cases[0]
    exp = expected_billing(case.body)

    if args.dry_run:
        report = build_report(case, exp, dry_run=True)
        append_report(report, args.report)
        print(report)
        return 0

    # ---- live：强制单用例、单进程 ----
    os.environ["KLING_TEST_CASE"] = case.id
    acquire_run_lock()

    state = load_state()
    if case.id in state.get("executed", {}) and not args.force:
        rec = state["executed"][case.id]
        print(
            f"拒绝重复提交：{case.id} 已于 {rec.get('at')} 提交过\n"
            f"  task_id={rec.get('task_id')}\n"
            f"  如需重新提交请加 --force",
            file=sys.stderr,
        )
        return 2

    # 刷新锁文件中的 case 字段
    with open(LOCK_FILE, "w", encoding="utf-8") as f:
        json.dump(
            {
                "pid": os.getpid(),
                "at": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
                "case": case.id,
            },
            f,
        )

    # live: 1 次 POST + 轮询 + 查额度/日志
    usage_before = curl("GET", "/api/usage/token/")["json"].get("data") or {}
    avail_before = usage_before.get("total_available", 0)

    submit = curl("POST", "/v1/video/generations", case.body)
    task_id = submit["json"].get("task_id") or submit["json"].get("id")

    usage_after = curl("GET", "/api/usage/token/")["json"].get("data") or {}
    avail_after = usage_after.get("total_available", 0)
    actual_deduct = (avail_before or 0) - (avail_after or 0)

    latest_log = curl("GET", "/api/log/token?p=1&page_size=1")["json"].get("data") or []
    log_entry = latest_log[0] if latest_log else {}
    log_quota = log_entry.get("quota")
    match = (
        actual_deduct == exp["expected_quota"]
        if submit["http_code"] == 200 and case.expect_submit
        else actual_deduct == 0
    )

    poll = None
    if task_id and submit["http_code"] == 200:
        poll = poll_task(task_id)

    live = {
        "submit": submit,
        "task_id": task_id,
        "poll": poll,
        "billing": {
            "avail_before": avail_before,
            "avail_after": avail_after,
            "actual_deduct": actual_deduct,
            "log_quota": log_quota,
            "request_id": log_entry.get("request_id"),
            "match": match,
        },
    }
    report = build_report(case, exp, dry_run=False, live=live)
    append_report(report, args.report)
    print(report)

    if submit["http_code"] == 200 and case.expect_submit:
        state = load_state()
        state.setdefault("executed", {})[case.id] = {
            "at": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
            "task_id": task_id,
            "expected_quota": exp["expected_quota"],
            "actual_deduct": actual_deduct,
            "poll_status": poll["final_status"] if poll else "n/a",
        }
        save_state(state)

    print(f"\n报告: {args.report}", file=sys.stderr)
    return 0


if __name__ == "__main__":
    sys.exit(main())
