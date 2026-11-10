#!/usr/bin/env node
/**
 * 图像生成扣费测试，v2026.11.10。
 * 用例对应 图像生成扣费测试用例-v2026.11.10.md。
 * 提交按每秒 20 条铺开，再并发轮询。
 * API Key 从环境变量 IMAGE_TEST_API_KEY 读取，不写进文件。
 */
'use strict';

const http = require('http');
const fs = require('fs');
const path = require('path');
const { execFileSync } = require('child_process');

const HOST = '127.0.0.1';
const PORT = 9006;
const API_KEY = process.env.IMAGE_TEST_API_KEY;
if (!API_KEY) {
  console.error('缺少 IMAGE_TEST_API_KEY');
  process.exit(1);
}

const IMG = 'https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk';
const OUT_DIR = path.join(__dirname, '..', '模型测试结果记录文档', '20261110-图像扣费-v2026.11.10');
const SUBMIT_PER_SECOND = 20;
const POLL_CONCURRENCY = 20;
const POLL_INTERVAL_MS = 8000;
const POLL_TIMEOUT_MS = 15 * 60 * 1000;
const DONE = new Set(['SUCCESS', 'FAILURE', 'FAILED', 'FAIL', 'CANCELLED']);

function Q(expr) { return Math.round(expr * 0.5); }

const CASES = [
  { n: 1, model: 'flux-3-image-klein', name: '不传分辨率', expectQuota: Q(6904), tier: '1k，参考图不加价', body: { prompt: '一只在樱花树下打盹的橘猫', ratio: '16:9', n: 1, image: [IMG] } },
  { n: 2, model: 'flux-3-image-klein', name: '768', expectQuota: Q(6137), tier: '768sq', body: { prompt: '一只在樱花树下打盹的橘猫', resolution: '768', ratio: '1:1', n: 1 } },
  { n: 3, model: 'flux-3-image-klein', name: '1.5K', expectQuota: Q(8438), tier: '1.5k', body: { prompt: '一只在樱花树下打盹的橘猫', resolution: '1.5K', ratio: '4:3', n: 1 } },
  { n: 4, model: 'flux-3-image-klein', name: '2k', expectQuota: Q(10740), tier: '2k', body: { prompt: '一只在樱花树下打盹的橘猫', resolution: '2k', ratio: '3:2', n: 1 } },
  { n: 5, model: 'flux-3-image-klein', name: '4K', expectQuota: Q(27616), tier: '4k', body: { prompt: '一只在樱花树下打盹的橘猫', resolution: '4K', ratio: '16:9', n: 1 } },
  { n: 6, model: 'flux-3-image', name: '1k', expectQuota: Q(36822), tier: '1k', body: { prompt: '浩瀚星空下的沙漠', resolution: '1k', ratio: '1:1', n: 1 } },
  { n: 7, model: 'flux-3-image', name: '768sq', expectQuota: Q(31452), tier: '768sq', body: { prompt: '浩瀚星空下的沙漠', resolution: '768sq', ratio: '1:1', n: 1 } },
  { n: 8, model: 'flux-3-image', name: '1.5k', expectQuota: Q(53699), tier: '1.5k', body: { prompt: '浩瀚星空下的沙漠', resolution: '1.5k', ratio: '4:3', n: 1 } },
  { n: 9, model: 'flux-3-image', name: '2k', expectQuota: Q(76712), tier: '2k', body: { prompt: '浩瀚星空下的沙漠', resolution: '2k', ratio: '16:9', n: 1 } },
  { n: 10, model: 'flux-3-image', name: '4k', expectQuota: Q(465644), tier: '4k', body: { prompt: '浩瀚星空下的沙漠', resolution: '4k', ratio: '16:9', n: 1 } },
  { n: 11, model: 'mai-image-2.6', name: '不传分辨率', expectQuota: Q(29851), tier: '1K', body: { prompt: '一只在窗台上的橘猫', ratio: '1:1', n: 1 } },
  { n: 12, model: 'mai-image-2.6', name: '2k', expectQuota: Q(67162), tier: '2K', body: { prompt: '一只在窗台上的橘猫', resolution: '2k', ratio: '16:9', n: 1 } },
  { n: 13, model: 'mai-image-2.6', name: '1536x1536', expectQuota: Q(67162), tier: '像素大于 1024x1024，按 2K', body: { prompt: '一只在窗台上的橘猫', size: '1536x1536', resolution: '1K', n: 1 } },
  { n: 14, model: 'mai-image-2.6', name: '宽高覆盖size', expectQuota: Q(67162), tier: 'width/height 1536x1536 优先，按 2K', body: { prompt: '一只在窗台上的橘猫', size: '1024x1024', n: 1, extra_fields: { width: 1536, height: 1536 } } },
  { n: 15, model: 'mai-image-2.6', name: '1K参考图', expectQuota: Q(29851), tier: '1K，参考图不加价', body: { prompt: '按参考图风格画一只橘猫', resolution: '1K', ratio: '1:1', n: 1, image: [IMG] } },
  { n: 16, model: 'mai-image-2.6-flash', name: '不传分辨率', expectQuota: Q(14921), tier: '1K', body: { prompt: '一只在窗台上的橘猫', ratio: '1:1', n: 1 } },
  { n: 17, model: 'mai-image-2.6-flash', name: '2K', expectQuota: Q(33581), tier: '2K', body: { prompt: '一只在窗台上的橘猫', resolution: '2K', ratio: '16:9', n: 1 } },
  { n: 18, model: 'qwen-image-3.0', name: '1K', expectQuota: Q(19725), tier: '1 张', body: { prompt: '一杯咖啡的顶视角海报', resolution: '1K', ratio: '1:1' } },
  { n: 19, model: 'qwen-image-3.0', name: '2K', expectQuota: Q(19725), tier: '1 张，与 1K 同价', body: { prompt: '一杯咖啡的顶视角海报', resolution: '2K', ratio: '16:9', n: 1 } },
  { n: 20, model: 'qwen-image-3.0', name: '2张', expectQuota: Q(19725 * 2), tier: '1K x 2', body: { prompt: '一杯咖啡的顶视角海报', resolution: '1K', ratio: '1:1', n: 2 } },
  { n: 21, model: 'qwen-image-3.0', name: '1张参考图', expectQuota: Q(19725 + 2196), tier: '1K + 1 张参考图（字符串）', body: { prompt: '按参考图风格画一杯咖啡', resolution: '1K', ratio: '1:1', n: 1, image: IMG } },
  { n: 22, model: 'qwen-image-3.0', name: '2张2参考图', expectQuota: Q(19725 * 2 + 2196 * 2), tier: '1K x 2 + 参考图 x 2', body: { prompt: '按参考图风格画一杯咖啡', resolution: '1K', ratio: '1:1', n: 2, image: [IMG, IMG] } },
  { n: 23, model: 'qwen-image-3.0-pro', name: '1K', expectQuota: Q(27396), tier: '1K', body: { prompt: '一张密集排版的菜单', resolution: '1K', ratio: '3:4', n: 1 } },
  { n: 24, model: 'qwen-image-3.0-pro', name: '2K', expectQuota: Q(54792), tier: '2K', body: { prompt: '一张密集排版的菜单', resolution: '2K', ratio: '3:4', n: 1 } },
  { n: 25, model: 'qwen-image-3.0-pro', name: '1600x900', expectQuota: Q(27396), tier: '像素未过线，按 1K', body: { prompt: '一张密集排版的菜单', size: '1600x900', n: 1 } },
  { n: 26, model: 'qwen-image-3.0-pro', name: '2048x2048', expectQuota: Q(54792), tier: '像素过线，按 2K', body: { prompt: '一张密集排版的菜单', size: '2048x2048', n: 1 } },
  { n: 27, model: 'qwen-image-3.0-pro', name: '2Kx2', expectQuota: Q(54792 * 2), tier: '2K x 2', body: { prompt: '一张密集排版的菜单', resolution: '2K', ratio: '1:1', n: 2 } },
  { n: 28, model: 'qwen-image-3.0-pro', name: '1Kx2参考图', expectQuota: Q(27396 * 2 + 2196), tier: '1K x 2 + 1 张参考图', body: { prompt: '按参考图风格排一张菜单', resolution: '1K', ratio: '3:4', n: 2, image: [IMG] } },
  { n: 29, model: 'muse-image-ext', name: '固定', expectQuota: Q(4603), tier: '固定', body: { prompt: '海边的白色小屋', ratio: '16:9', n: 1 } },
  { n: 30, model: 'muse-image-ext', name: '参考图', expectQuota: Q(4603), tier: '固定，参考图不加价', body: { prompt: '按参考图风格画海边的白色小屋', ratio: '1:1', n: 1, image: [IMG] } },
  { n: 31, model: 'hy-image-v3.5-preview', name: '不传分辨率', expectQuota: Q(18411), tier: '1K', body: { prompt: '雨后的江南小镇', ratio: '16:9', n: 1 } },
  { n: 32, model: 'hy-image-v3.5-preview', name: '2K', expectQuota: Q(18411), tier: '2K，与 1K 同价', body: { prompt: '雨后的江南小镇', resolution: '2K', ratio: '3:4', n: 1 } },
  { n: 33, model: 'hy-image-v3.5-preview', name: '4K', expectQuota: Q(24548), tier: '4K', body: { prompt: '雨后的江南小镇', resolution: '4K', ratio: '16:9', n: 1 } },
  { n: 34, model: 'hy-image-v3.5-preview', name: '2000x1104', expectQuota: Q(18411), tier: '短边 1104，按 2K 价', body: { prompt: '雨后的江南小镇', size: '2000x1104', n: 1 } },
  { n: 35, model: 'hy-image-v3.5-preview', name: '宽高短边2200', expectQuota: Q(24548), tier: 'width/height 3008x2208 优先，短边 2208 按 4K', body: { prompt: '雨后的江南小镇', size: '1024x1024', n: 1, extra_fields: { width: 3008, height: 2208 } } },
];

function request(method, urlPath, body) {
  return new Promise((resolve, reject) => {
    const payload = body ? JSON.stringify(body) : null;
    const started = Date.now();
    const req = http.request({
      host: HOST, port: PORT, method, path: urlPath, timeout: 120000,
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
        try { json = JSON.parse(text); } catch { json = { raw: text.slice(0, 4000) }; }
        resolve({
          status: res.statusCode,
          json,
          elapsedMs: Date.now() - started,
          requestId: res.headers['x-oneapi-request-id'] || '',
        });
      });
    });
    req.on('error', reject);
    req.on('timeout', () => req.destroy(new Error('request timeout')));
    if (payload) req.write(payload);
    req.end();
  });
}

function sleep(ms) { return new Promise((r) => setTimeout(r, ms)); }

async function pool(items, limit, worker) {
  const ret = new Array(items.length);
  let i = 0;
  async function run() {
    while (i < items.length) {
      const idx = i++;
      ret[idx] = await worker(items[idx], idx);
    }
  }
  await Promise.all(Array.from({ length: Math.min(limit, items.length) }, run));
  return ret;
}

async function paced(items, perSecond, worker) {
  const gap = 1000 / perSecond;
  const t0 = Date.now();
  const ret = new Array(items.length);
  await Promise.all(items.map(async (item, idx) => {
    const wait = t0 + idx * gap - Date.now();
    if (wait > 0) await sleep(wait);
    ret[idx] = await worker(item, idx);
  }));
  return ret;
}

function amountOfQuota(quota) {
  return (quota / 500000 * 7.3).toFixed(6);
}

function quotaOfAmount(amount) {
  const n = Number(amount);
  if (!Number.isFinite(n) || amount === '') return null;
  return Math.round(n / 7.3 * 500000);
}

function fmt(n) {
  if (n == null || n === '' || Number.isNaN(n)) return '-';
  return Number(n).toLocaleString('en-US');
}

function now() {
  return new Date().toLocaleString('zh-CN', { hour12: false, timeZone: 'Asia/Shanghai' });
}

function taskIdOf(json) {
  return json?.data?.task_id || json?.task_id || json?.data?.id || '';
}

function statusOf(json) {
  return String(json?.data?.status || json?.status || '').toUpperCase();
}

function fileFor(c) {
  const safe = `${String(c.n).padStart(2, '0')}-${c.model}-${c.name}`.replace(/[^\w.\-\u4e00-\u9fa5]+/g, '-');
  return path.join(OUT_DIR, `${safe}.md`);
}

function imageLabel(image) {
  if (Array.isArray(image)) return `${image.length} 张`;
  if (typeof image === 'string' && image) return '1 张（字符串）';
  return '无';
}

function shrink(value, depth = 0) {
  if (depth > 8) return value;
  if (typeof value === 'string') {
    if (value.length > 400 && !value.startsWith('http')) {
      return `${value.slice(0, 80)}…（已截断，原长 ${value.length}）`;
    }
    return value;
  }
  if (Array.isArray(value)) return value.map((v) => shrink(v, depth + 1));
  if (value && typeof value === 'object') {
    const out = {};
    for (const [k, v] of Object.entries(value)) out[k] = shrink(v, depth + 1);
    return out;
  }
  return value;
}

function dump(value) {
  return JSON.stringify(shrink(value), null, 2);
}

function fetchUpstream(taskId) {
  if (!taskId || !/^task_[A-Za-z0-9]+$/.test(taskId)) return null;
  try {
    const out = execFileSync('docker', [
      'exec', 'new-api-dev-pg', 'psql', '-U', 'root', '-d', 'new-api', '-t', '-A',
      '-c', `SELECT upstream_response::text FROM tasks WHERE task_id = '${taskId}'`,
    ], { encoding: 'utf8' }).trim();
    if (!out) return null;
    return JSON.parse(out);
  } catch {
    return null;
  }
}

async function fetchLogs() {
  const r = await request('GET', '/api/log/token?p=1&page_size=1000');
  return Array.isArray(r.json?.data) ? r.json.data : [];
}

function pickLogs(logs, record) {
  return logs.filter((l) => {
    if (record.requestId && l.request_id === record.requestId) return true;
    if (record.taskId && typeof l.other === 'string' && l.other.includes(record.taskId)) return true;
    return false;
  });
}

function prechargeLog(logs) {
  return logs.find((l) => l.type === 2 && l.request_id) || logs.find((l) => l.type === 2) || null;
}

function buildReport(c, record) {
  const lines = [];
  lines.push(`# ${c.n}. ${c.model} ${c.name}`);
  lines.push('');
  lines.push('> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。');
  lines.push('');
  lines.push(`> 执行时间：${record.startedAt}  |  模型：\`${c.model}\`  |  档位：${c.tier}`);
  lines.push('');
  lines.push('### 调用参数');
  lines.push('```json');
  lines.push(dump({ model: c.model, ...c.body }));
  lines.push('```');
  lines.push('');

  const body = c.body;
  const extra = body.extra_fields || {};
  lines.push('### 价格变量核对');
  lines.push('| 变量 | 请求值 | 说明 |');
  lines.push('|---|---|---|');
  lines.push(`| resolution | ${body.resolution || '未传'} | 分辨率档位 |`);
  lines.push(`| ratio | ${body.ratio || '未传'} | 宽高比 |`);
  lines.push(`| size | ${body.size || '未传'} | 宽x高 |`);
  lines.push(`| quality | ${body.quality || '未传'} | 渲染精度 |`);
  lines.push(`| n | ${body.n != null ? body.n : '未传'} | 张数 |`);
  lines.push(`| image | ${imageLabel(body.image)} | 参考图 |`);
  lines.push(`| extra_fields.width | ${extra.width != null ? extra.width : '未传'} | 成对宽，优先于 size |`);
  lines.push(`| extra_fields.height | ${extra.height != null ? extra.height : '未传'} | 成对高，优先于 size |`);
  lines.push('');

  lines.push('### 预期扣费');
  lines.push('| 项目 | 值 |');
  lines.push('|---|---|');
  lines.push('| 预期 HTTP 状态 | 200 |');
  lines.push(`| 预期 Quota | **${fmt(c.expectQuota)}** |`);
  lines.push(`| 预期 USD | $${(c.expectQuota / 500000).toFixed(4)} |`);
  lines.push(`| 预期 RMB | ¥${amountOfQuota(c.expectQuota)} |`);
  lines.push('| 结算方式 | 按张预扣，成功后不按 token 重算 |');
  lines.push('');

  const httpOk = record.submit?.status === 200 ? '✅' : '❌';
  lines.push('### 提交结果');
  lines.push('| 项目 | 值 |');
  lines.push('|---|---|');
  lines.push(`| HTTP 状态 | ${record.submit?.status ?? '-'}（预期 200）${httpOk} |`);
  lines.push(`| 耗时 | ${record.submit ? (record.submit.elapsedMs / 1000).toFixed(2) : '-'}s |`);
  lines.push(`| task_id | \`${record.taskId || '-'}\` |`);
  lines.push(`| request_id | ${record.requestId || '-'} |`);
  lines.push('');
  lines.push('```json');
  lines.push(dump(record.submit?.json ?? null));
  lines.push('```');
  lines.push('');

  const preQuota = record.preQuota;
  const preMatch = preQuota === c.expectQuota ? '✅ 符合' : '❌ 不符';
  lines.push('### 扣费分析（提交时）');
  lines.push('| 项目 | Quota | RMB |');
  lines.push('|---|---|---|');
  lines.push(`| 任务首次查询 amount | ${fmt(record.earlyQuota)} | ${record.earlyAmount || '-'} |`);
  lines.push(`| 消费日志 Quota | ${fmt(record.logQuota)} | ${record.logQuota != null ? amountOfQuota(record.logQuota) : '-'} |`);
  lines.push(`| **预扣金额** | **${fmt(preQuota)}** | **${preQuota != null ? `¥${amountOfQuota(preQuota)}` : '-'}** |`);
  lines.push(`| 预期扣减 | ${fmt(c.expectQuota)} | ¥${amountOfQuota(c.expectQuota)} |`);
  lines.push(`| 预扣是否符合 | ${record.submit?.status === 200 ? preMatch : '❌ 提交失败'} | — |`);
  lines.push(`| 分组倍率 | ${record.groupRatio ?? '-'} | — |`);
  lines.push('');

  if (record.poll) {
    const data = record.poll.json?.data || {};
    const images = Array.isArray(data.images) ? data.images : [];
    const urls = images.map((it) => it.url).filter(Boolean);
    lines.push('### 任务轮询');
    lines.push('| 项目 | 值 |');
    lines.push('|---|---|');
    lines.push(`| 最终状态 | **${record.finalStatus || '-'}** |`);
    lines.push(`| 轮询次数 | ${record.polls} |`);
    lines.push(`| 完成耗时 | ${(record.pollElapsedMs / 1000).toFixed(1)}s |`);
    lines.push(`| progress | ${data.progress || '-'} |`);
    lines.push(`| 输出张数 | ${images.length} |`);
    lines.push('');
    if (urls.length) lines.push(`**图片 URL**: ${urls.join('  \n')}`, '');
    if (data.fail_reason) lines.push(`**失败原因**: ${data.fail_reason}`, '');

    const finalAmount = data.amount || '';
    const finalQuota = quotaOfAmount(finalAmount);
    let diff = '-';
    if (preQuota != null && finalQuota != null) {
      const delta = finalQuota - preQuota;
      const deltaRmb = (delta / 500000 * 7.3);
      if (delta === 0) diff = '¥0.000000（无差额）';
      else if (delta > 0) diff = `+¥${deltaRmb.toFixed(6)}（补扣）`;
      else diff = `-¥${(-deltaRmb).toFixed(6)}（退还）`;
    }
    let finalNote = '';
    if (record.finalStatus && record.finalStatus !== 'SUCCESS') {
      finalNote = '任务未成功。失败会退款，最终金额不要当成成交价';
    } else if (finalQuota != null && finalQuota === c.expectQuota) {
      finalNote = '✅ 最终扣费符合预期';
    } else if (finalQuota != null) {
      finalNote = '❌ 最终扣费不符合预期';
    }

    lines.push('#### 最终结算（多退少补）');
    lines.push('| 项目 | Quota | RMB |');
    lines.push('|---|---|---|');
    lines.push(`| 预扣金额 | ${fmt(preQuota)} | ${preQuota != null ? `¥${amountOfQuota(preQuota)}` : '-'} |`);
    lines.push(`| 平台记录最终消费 | ${fmt(finalQuota)} | ${finalAmount ? `**¥${finalAmount}**` : '-'} |`);
    lines.push(`| 差额 | — | ${diff} |`);
    lines.push(`| 最终扣费是否符合 | ${finalNote || '-'} | — |`);
    lines.push('');

    if (record.upstream) {
      lines.push('#### 原始上游响应');
      lines.push('```json');
      lines.push(dump(record.upstream));
      lines.push('```');
      lines.push('');
    }

    lines.push('#### 网关查询响应');
    lines.push('```json');
    lines.push(dump(record.poll.json));
    lines.push('```');
    lines.push('');
  }

  lines.push('### 消费日志');
  lines.push('```json');
  lines.push(dump(record.logs || []));
  lines.push('```');
  lines.push('');
  return lines.join('\n');
}

function writeCase(c, record) {
  fs.writeFileSync(fileFor(c), buildReport(c, record));
}

async function pollOne(record) {
  const deadline = Date.now() + POLL_TIMEOUT_MS;
  const t0 = Date.now();
  let last = null;
  while (Date.now() < deadline) {
    record.polls++;
    const r = await request('GET', `/v2/image-tasks/${encodeURIComponent(record.taskId)}`);
    last = r;
    if (record.polls === 1) {
      record.earlyAmount = r.json?.data?.amount || '';
      record.earlyQuota = quotaOfAmount(record.earlyAmount);
    }
    const status = statusOf(r.json);
    if (DONE.has(status)) {
      record.finalStatus = status;
      record.poll = r;
      break;
    }
    await sleep(POLL_INTERVAL_MS);
  }
  record.pollElapsedMs = Date.now() - t0;
  if (!record.finalStatus) {
    record.finalStatus = 'TIMEOUT';
    record.poll = last;
  }
  return record;
}

async function attachLogs(record) {
  for (let i = 0; i < 4; i++) {
    const logs = await fetchLogs();
    record.logs = pickLogs(logs, record);
    if (record.logs.length) break;
    await sleep(400);
  }
  const pre = prechargeLog(record.logs || []);
  record.logQuota = pre ? pre.quota : null;
  if (pre && pre.other) {
    try { record.groupRatio = JSON.parse(pre.other).group_ratio; } catch { record.groupRatio = null; }
  }
  record.preQuota = record.logQuota != null ? record.logQuota : record.earlyQuota;
}

function selectedCases() {
  const raw = process.env.IMAGE_TEST_CASES;
  if (!raw) return CASES;
  const wanted = new Set(raw.split(',').map((s) => Number(s.trim())).filter((n) => n > 0));
  return CASES.filter((c) => wanted.has(c.n));
}

function finalMarkOf(r) {
  const finalAmount = r.poll?.json?.data?.amount || '';
  const finalQuota = quotaOfAmount(finalAmount);
  if (!r.finalStatus || r.finalStatus === '未提交') return '未完成';
  if (r.finalStatus !== 'SUCCESS') return '失败退款';
  if (finalQuota === r.c.expectQuota) return '符合';
  return '不符';
}

function summaryRow(r) {
  const pre = r.preQuota === r.c.expectQuota ? '符合' : '不符';
  const finalAmount = r.poll?.json?.data?.amount || '-';
  return `| ${r.c.n} | \`${r.c.model}\` | ${r.c.name} | ${r.finalStatus || '未提交'} | ${fmt(r.preQuota)} | ${fmt(r.c.expectQuota)} | ${pre} | ${finalAmount} | ${finalMarkOf(r)} |`;
}

function mergeSummary(submitted) {
  const indexPath = path.join(OUT_DIR, '_汇总.md');
  const fresh = submitted.map(summaryRow);
  if (!process.env.IMAGE_TEST_CASES || !fs.existsSync(indexPath)) {
    const index = [
      '# 图像扣费测试汇总 v2026.11.10',
      '',
      `执行时间：${now()}`,
      '',
      '预期金额按价格文档 `v2026.11.10.2` 的表达式、分组倍率 1。最终符合指成功任务的 `amount` 与预期 Quota 换算一致。这批按张预扣，成功后不按 token 重算。',
      '',
      '| # | 模型 | 用例 | 状态 | 预扣 Quota | 预期 Quota | 预扣 | 最终 amount | 最终 |',
      '|---|---|---|---|---:|---:|---|---:|---|',
      ...fresh,
      '',
    ].join('\n');
    fs.writeFileSync(indexPath, index);
    return;
  }
  let text = fs.readFileSync(indexPath, 'utf8');
  for (const row of fresh) {
    const n = row.match(/^\| (\d+) \|/)[1];
    const lineRe = new RegExp(`^\\| ${n} \\|.*$`, 'm');
    if (lineRe.test(text)) text = text.replace(lineRe, row);
    else text += `${row}\n`;
  }
  text += `\n补跑时间：${now()}。用例 ${submitted.map((r) => r.c.n).join(', ')}。\n`;
  fs.writeFileSync(indexPath, text);
}

async function main() {
  fs.mkdirSync(OUT_DIR, { recursive: true });
  const probe = await request('GET', '/api/status');
  if (probe.status !== 200) {
    console.error(`本地网关未就绪 HTTP ${probe.status}`);
    process.exit(1);
  }
  const cases = selectedCases();
  console.log(`输出目录 ${OUT_DIR}`);
  console.log(`用例 ${cases.length}，提交每秒 ${SUBMIT_PER_SECOND} 条`);

  async function submitOne(c) {
    const payload = { model: c.model, ...c.body };
    const record = {
      c, startedAt: now(), submit: null, taskId: '', requestId: '',
      polls: 0, finalStatus: '', poll: null, logs: [],
      earlyAmount: '', earlyQuota: null, logQuota: null, preQuota: null, groupRatio: null,
      pollElapsedMs: 0,
    };
    try {
      record.submit = await request('POST', '/v2/image-tasks', payload);
    } catch (err) {
      record.submit = { status: 0, json: { error: String(err) }, elapsedMs: 0, requestId: '' };
    }
    record.requestId = record.submit.requestId || '';
    record.taskId = taskIdOf(record.submit.json);
    console.log(`提交 ${c.n} ${c.model} ${c.name} HTTP ${record.submit.status} task=${record.taskId || '-'}`);
    writeCase(c, record);
    return record;
  }

  const submitted = await paced(cases, SUBMIT_PER_SECOND, submitOne);
  const runnable = submitted.filter((r) => r.taskId);

  async function finish(record) {
    await pollOne(record);
    await attachLogs(record);
    record.upstream = fetchUpstream(record.taskId);
    if (!record.upstream) {
      await sleep(1000);
      record.upstream = fetchUpstream(record.taskId);
    }
    writeCase(record.c, record);
    const mark = record.preQuota === record.c.expectQuota ? '预扣符合' : '预扣不符';
    const finalAmount = record.poll?.json?.data?.amount || '-';
    console.log(`完成 ${record.c.n} ${record.finalStatus} ${mark} pre=${record.preQuota} expect=${record.c.expectQuota} amount=${finalAmount}`);
    return record;
  }

  await pool(runnable, POLL_CONCURRENCY, finish);
  for (const record of submitted.filter((r) => !r.taskId)) {
    await attachLogs(record);
    writeCase(record.c, record);
  }

  mergeSummary(submitted);
  const preOk = submitted.filter((r) => r.preQuota === r.c.expectQuota).length;
  const finalOk = submitted.filter((r) => finalMarkOf(r) === '符合').length;
  console.log(`预扣符合 ${preOk}/${submitted.length}，最终符合 ${finalOk}/${submitted.length}`);
  console.log('全部用例已写完');
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
