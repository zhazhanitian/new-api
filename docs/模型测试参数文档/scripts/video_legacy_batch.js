#!/usr/bin/env node
/**
 * 旧请求字段兼容测试。全部同时提交，不设并发上限。
 * 不改已有测试文档。API Key：VIDEO_TEST_API_KEY。
 */
'use strict';

const http = require('http');
const fs = require('fs');
const path = require('path');

const HOST = '127.0.0.1';
const PORT = 9006;
const API_KEY = process.env.VIDEO_TEST_API_KEY;
if (!API_KEY) {
  console.error('缺少 VIDEO_TEST_API_KEY');
  process.exit(1);
}

const IMG = 'https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg';
const IMG_LAST = 'https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg';
const OUT_DIR = path.join(__dirname, '..', '模型测试结果记录文档', '20261028-旧格式');
const POLL_INTERVAL_MS = 15000;
const POLL_TIMEOUT_MS = 12 * 60 * 1000;
const DONE = new Set(['SUCCESS', 'FAILURE', 'FAILED', 'CANCELLED', 'SUCCEEDED']);

function legacyBody(extra) {
  return {
    prompt: extra.prompt,
    size: extra.size,
    seconds: String(extra.seconds),
    duration: extra.seconds,
    images: [IMG],
    metadata: {
      resolution: extra.size,
      ratio: extra.ratio || '16:9',
      input_usage: 'FirstFrame',
      ...(extra.lastFrame ? { last_frame: IMG_LAST } : {}),
    },
  };
}

const CASES = [
  { model: 'kling-video-3.0-omni', name: '旧字段', expectExpr: 547945, body: legacyBody({ prompt: '沿用画面风格，镜头缓慢推进', size: '1080P', seconds: 5 }) },
  { model: 'kling-video-3.0', name: '旧字段', expectExpr: 547945, body: legacyBody({ prompt: '雨后街道，霓虹倒映在地面积水上', size: '1080P', seconds: 5 }) },
  { model: 'vidu-video-q3-pro', name: '旧字段', expectExpr: 642465, body: legacyBody({ prompt: '雪山日出，云层缓慢流动', size: '1080P', seconds: 5, lastFrame: true }) },
  { model: 'vidu-video-q3-turbo', name: '旧字段', expectExpr: 300000, body: legacyBody({ prompt: '城市夜景，车灯拉出光轨', size: '1080P', seconds: 5 }) },
  { model: 'hailuo-video-h3', name: '旧字段', expectExpr: 445205, body: legacyBody({ prompt: '海边栈道，镜头跟着人物向前走', size: '1080P', seconds: 5, lastFrame: true }) },
  { model: 'hailuo-video-2.3', name: '旧字段', expectExpr: 476712, body: legacyBody({ prompt: '雨中的咖啡馆橱窗，行人撑伞经过', size: '1080P', seconds: 6, ratio: '' }) },
  { model: 'veo-video-3.1', name: '旧字段', expectExpr: 1643832, body: legacyBody({ prompt: '沙漠公路上的一辆旧汽车', size: '1080P', seconds: 8, lastFrame: true }) },
  { model: 'veo-video-3.1-fast', name: '旧字段', expectExpr: 821920, body: legacyBody({ prompt: '厨房里的面团被慢慢拉开', size: '720P', seconds: 8 }) },
  { model: 'sora-video-2.0', name: '旧字段', expectExpr: 1232880, body: legacyBody({ prompt: '纸船在水面上漂过桥洞', size: '1080P', seconds: 8 }) },
  { model: 'hunyuan-video-1.5', name: '旧字段', expectExpr: 342465, body: legacyBody({ prompt: '江南水乡，乌篷船慢慢划过', size: '1080P', seconds: 5 }) },
  { model: 'mingmou-video-1.0', name: '旧字段', expectExpr: 342465, body: legacyBody({ prompt: '雪地里的一盏路灯，雪花落下', size: '1080P', seconds: 5 }) },
  { model: 'pixverse-video-v6', name: '旧字段', expectExpr: 361645, body: legacyBody({ prompt: '风吹过麦田，麦浪起伏', size: '1080P', seconds: 5 }) },
  { model: 'pixverse-video-c1', name: '旧字段', expectExpr: 381505, body: legacyBody({ prompt: '夜晚的码头，渔灯轻轻晃动', size: '1080P', seconds: 5 }) },
  { model: 'doubao-seedance-2-5-260628', name: '旧字段', expectNote: '不含视频 720p，有效单价 ¥73.50/1M', body: legacyBody({ prompt: '清晨海岸公路，镜头缓慢推进', size: '720p', seconds: 5 }) },
  { model: 'doubao-seedance-2-0-260128', name: '旧字段', expectNote: '不含视频 720p。文档新价 ¥48.30/1M；若后台仍是旧 1.5 倍则为 ¥67.16/1M', body: legacyBody({ prompt: '山间公路，云雾缓慢流动', size: '720p', seconds: 5 }) },
];

for (const c of CASES) {
  if (c.model === 'hailuo-video-2.3') delete c.body.metadata.ratio;
}

function request(method, urlPath, body) {
  return new Promise((resolve, reject) => {
    const payload = body ? JSON.stringify(body) : null;
    const req = http.request({
      host: HOST, port: PORT, method, path: urlPath,
      headers: {
        Authorization: `Bearer ${API_KEY}`,
        'Content-Type': 'application/json',
        ...(payload ? { 'Content-Length': Buffer.byteLength(payload) } : {}),
      },
    }, (res) => {
      const chunks = [];
      res.on('data', (c) => chunks.push(c));
      res.on('end', () => {
        const text = Buffer.concat(chunks).toString('utf8');
        let json = null;
        try { json = JSON.parse(text); } catch { json = { raw: text }; }
        resolve({ status: res.statusCode, json });
      });
    });
    req.on('error', reject);
    if (payload) req.write(payload);
    req.end();
  });
}

function sleep(ms) { return new Promise((r) => setTimeout(r, ms)); }
function statusOf(json) { return String(json?.status || json?.data?.status || '').toUpperCase(); }
function taskIdOf(json) { return json?.task_id || json?.id || json?.data?.task_id || ''; }

async function fetchLogs() {
  const r = await request('GET', '/api/log/token?p=1&page_size=100');
  return Array.isArray(r.json?.data) ? r.json.data : [];
}

function fileFor(c) {
  const safe = `${c.model}-旧字段`.replace(/[^\w.\-\u4e00-\u9fa5]+/g, '-');
  return path.join(OUT_DIR, `${safe}.md`);
}

function fmt(n) {
  if (n == null || Number.isNaN(Number(n))) return '-';
  return Math.round(Number(n)).toLocaleString('en-US');
}

function writeReport(c, rec) {
  const expr = c.expectExpr || 0;
  const expectQuota = expr ? Math.round(expr / 2) : null;
  const expectRmb = expr ? expr / 1e6 * 7.3 : null;
  const amount = rec.poll?.data?.amount != null ? Number(rec.poll.data.amount) : null;
  const settledQuota = amount != null ? Math.round(amount / 7.3 * 500000) : null;
  const taskId = rec.taskId || '';
  let pre = null;
  let actual = null;
  let requestId = '-';
  for (const l of rec.logs || []) {
    let other = {};
    try { other = JSON.parse(l.other || '{}'); } catch { other = {}; }
    if (taskId && (l.other || '').includes(taskId)) {
      if (other.pre_consumed_quota) pre = other.pre_consumed_quota;
      if (other.actual_quota) actual = other.actual_quota;
    }
    if (l.type === 2 && l.model_name === c.model && pre == null && l.quota) {
      pre = l.quota;
    }
    if (l.request_id) requestId = l.request_id;
  }
  let verdict = '—';
  const notes = [];
  if (rec.submit.status !== 200) {
    verdict = '❌ 请求未通过';
    notes.push(String(rec.submit.json?.message || rec.submit.json?.code || '').slice(0, 240));
  } else if (pre != null && expectQuota != null) {
    verdict = Number(pre) === expectQuota ? '✅ 预扣符合' : `❌ 预扣 ${pre}，预期 ${expectQuota}`;
  } else if (c.expectNote) {
    verdict = rec.finalStatus === 'SUCCESS' ? '✅ 请求通过（豆包按 token 结算，见最终消费）' : '❌ 请求未完成';
  } else if (expectRmb != null && amount != null && Math.abs(amount - expectRmb) / expectRmb < 0.03) {
    verdict = '✅ 预扣日志未单独匹配，最终消费与预期价档一致';
  } else if (expectRmb != null && amount != null) {
    verdict = '❌ 最终消费与预期价档差距超过 3%';
  }
  const lines = [];
  lines.push(`# ${c.model} 旧格式测试报告`);
  lines.push('');
  lines.push(`## ${c.model} ${c.name}`);
  lines.push(`> 执行时间：2026-10-28  |  模型：\`${c.model}\``);
  lines.push('');
  lines.push('> 覆盖旧字段：`size`、`images`、`seconds`、`duration`、`metadata.resolution`、`metadata.ratio`、`metadata.input_usage`。支持尾帧的模型另带 `metadata.last_frame`。');
  lines.push('');
  lines.push('### 调用参数');
  lines.push('```json');
  lines.push(JSON.stringify({ model: c.model, ...c.body }, null, 2));
  lines.push('```');
  lines.push('');
  lines.push('### 价格变量核对');
  lines.push('| 变量 | 请求值 | 说明 |');
  lines.push('|---|---|---|');
  lines.push(`| size | ${c.body.size} | 旧分辨率。外层 resolution 为空时收成 resolution |`);
  lines.push(`| seconds / duration | ${c.body.seconds} | 旧时长。同时存在时上游以 seconds 为准，表达式读 duration |`);
  lines.push(`| images | 1 张 | 旧图片字段。input_usage=FirstFrame 时第一张当首帧 |`);
  lines.push(`| metadata.resolution | ${c.body.metadata.resolution} | 与 size 同值，确认旧字段仍被接受 |`);
  lines.push(`| metadata.ratio | ${c.body.metadata.ratio || '未传'} | 旧宽高比 |`);
  lines.push(`| metadata.input_usage | FirstFrame | 旧首帧标记 |`);
  lines.push(`| metadata.last_frame | ${c.body.metadata.last_frame ? '有' : '未传'} | 旧尾帧 |`);
  lines.push('');
  lines.push('### 预期扣费');
  lines.push('| 项目 | 值 |');
  lines.push('|---|---|');
  lines.push('| 预期 HTTP 状态 | 200 |');
  if (expectQuota != null) {
    lines.push(`| 预期表达式值 | ${fmt(expr)} |`);
    lines.push(`| 预期 Quota | **${fmt(expectQuota)}** |`);
    lines.push(`| 预期 RMB | ¥${expectRmb.toFixed(4)} |`);
  } else {
    lines.push(`| 预期 | ${c.expectNote} |`);
  }
  lines.push('');
  lines.push('### 提交结果');
  lines.push('| 项目 | 值 |');
  lines.push('|---|---|');
  lines.push(`| HTTP 状态 | ${rec.submit.status}（预期 200）${rec.submit.status === 200 ? '✅' : '❌'} |`);
  lines.push(`| task_id | \`${taskId || '-'}\` |`);
  lines.push('');
  lines.push('```json');
  lines.push(JSON.stringify(rec.submit.json, null, 2));
  lines.push('```');
  lines.push('');
  lines.push('### 扣费分析（提交时）');
  lines.push('| 项目 | Quota | RMB |');
  lines.push('|---|---|---|');
  lines.push(`| 预期扣减 | ${fmt(expectQuota)} | ${expectRmb != null ? '¥' + expectRmb.toFixed(4) : '-'} |`);
  lines.push(`| 日志预扣 | ${fmt(pre)} | ${pre != null ? '¥' + (pre / 500000 * 7.3).toFixed(4) : '-'} |`);
  lines.push(`| 平台记录最终消费 | ${fmt(actual != null ? actual : settledQuota)} | ${amount != null ? '¥' + amount.toFixed(4) : '-'} |`);
  lines.push(`| 是否符合 | ${verdict} | |`);
  lines.push(`| request_id | ${requestId} | |`);
  if (notes.length) lines.push('', notes.map((x) => `> ${x}`).join('\n'));
  lines.push('');
  lines.push('### 任务轮询');
  lines.push('| 项目 | 值 |');
  lines.push('|---|---|');
  lines.push(`| 最终状态 | **${rec.finalStatus || '-'}** |`);
  lines.push(`| 轮询次数 | ${rec.polls || 0} |`);
  lines.push(`| progress | ${rec.poll?.data?.progress || '-'} |`);
  if (rec.poll?.data?.fail_reason) lines.push(`| 失败原因 | ${String(rec.poll.data.fail_reason).replace(/\|/g, '/').slice(0, 300)} |`);
  if (rec.poll?.data?.result_url) lines.push('', `**视频 URL**: ${rec.poll.data.result_url}`);
  lines.push('');
  lines.push('#### 查询调用参数');
  lines.push('');
  lines.push(`\`GET /v1/video/generations/${taskId || ''}\``);
  lines.push('');
  lines.push('#### 最终查询响应（全量 JSON）');
  lines.push('```json');
  lines.push(JSON.stringify(rec.poll, null, 2));
  lines.push('```');
  lines.push('');
  fs.writeFileSync(fileFor(c), lines.join('\n'));
}

async function pollOne(rec) {
  if (!rec.taskId) {
    rec.finalStatus = '';
    rec.polls = 0;
    return rec;
  }
  const deadline = Date.now() + POLL_TIMEOUT_MS;
  while (Date.now() < deadline) {
    rec.polls++;
    const r = await request('GET', `/v1/video/generations/${encodeURIComponent(rec.taskId)}`);
    rec.poll = r.json;
    const status = statusOf(r.json);
    if (DONE.has(status)) {
      rec.finalStatus = status;
      return rec;
    }
    await sleep(POLL_INTERVAL_MS);
  }
  rec.finalStatus = 'TIMEOUT';
  return rec;
}

async function main() {
  fs.mkdirSync(OUT_DIR, { recursive: true });
  console.log(`同时提交 ${CASES.length} 条`);
  const started = Math.floor(Date.now() / 1000);
  const records = await Promise.all(CASES.map(async (c) => {
    let submit;
    try {
      submit = await request('POST', '/v1/video/generations', { model: c.model, ...c.body });
    } catch (err) {
      submit = { status: 0, json: { error: String(err) } };
    }
    const taskId = taskIdOf(submit.json);
    console.log(`${c.model} HTTP ${submit.status} ${taskId || submit.json?.code || ''}`);
    return { c, submit, taskId, polls: 0, finalStatus: '', poll: null, logs: [] };
  }));

  await sleep(1500);
  let logs = [];
  try { logs = await fetchLogs(); } catch (err) { console.log('日志读取失败', err); }
  for (const rec of records) {
    rec.logs = logs.filter((l) => l.model_name === rec.c.model && l.created_at >= started - 2);
    writeReport(rec.c, rec);
  }

  await Promise.all(records.map(async (rec) => {
    await pollOne(rec);
    try {
      const later = await fetchLogs();
      const more = later.filter((l) => l.model_name === rec.c.model && l.created_at >= started - 2);
      const seen = new Set(rec.logs.map((l) => l.id));
      for (const l of more) if (!seen.has(l.id)) rec.logs.push(l);
    } catch { /* keep earlier logs */ }
    writeReport(rec.c, rec);
    console.log(`完成 ${rec.c.model} ${rec.finalStatus || rec.submit.status}`);
  }));
  console.log('旧格式测试结束');
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
