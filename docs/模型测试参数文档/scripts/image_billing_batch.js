#!/usr/bin/env node
/**
 * 图像生成扣费测试。并发提交，再并发轮询。
 * API Key 从环境变量 IMAGE_TEST_API_KEY 读取，不写进文件。
 * 用例对应 图像生成扣费测试用例-v2026.11.07.md。
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
const OUT_DIR = path.join(__dirname, '..', '模型测试结果记录文档', '20261107-图像扣费-v2026.11.07');
const SUBMIT_CONCURRENCY = 8;
const POLL_CONCURRENCY = 12;
const POLL_INTERVAL_MS = 8000;
const POLL_TIMEOUT_MS = 15 * 60 * 1000;
const DONE = new Set(['SUCCESS', 'FAILURE', 'FAILED', 'FAIL', 'CANCELLED']);

function Q(expr) { return Math.round(expr * 0.5); }

const CASES = [
  { n: 1, model: 'gemini-2.5-flash-image', name: '1K', expectQuota: Q(41096), tier: '1K', body: { prompt: '一只在樱花树下打盹的橘猫', resolution: '1K', ratio: '1:1', n: 1 } },
  { n: 2, model: 'gemini-2.5-flash-image', name: '4K', expectQuota: Q(63014), tier: '4K', body: { prompt: '一只在樱花树下打盹的橘猫', resolution: '4K', ratio: '16:9', n: 1 } },
  { n: 3, model: 'gemini-3-pro-image', name: '1K', expectQuota: Q(136986), tier: '1K', body: { prompt: '浩瀚星空下的沙漠', resolution: '1K', ratio: '4:3', n: 1 } },
  { n: 4, model: 'gemini-3-pro-image', name: '4K', expectQuota: Q(246575), tier: '4K', body: { prompt: '浩瀚星空下的沙漠', resolution: '4K', ratio: '9:16', n: 1 } },
  { n: 5, model: 'gemini-3.1-flash-image', name: '720P', expectQuota: Q(45616), tier: '720P', body: { prompt: '未来城市夜景', resolution: '720P', ratio: '16:9', n: 1 } },
  { n: 6, model: 'gemini-3.1-flash-image', name: '2K', expectQuota: Q(102740), tier: '2K', body: { prompt: '未来城市夜景', resolution: '2K', ratio: '1:1', n: 1 } },
  { n: 7, model: 'kling-image-3.0', name: '1张', expectQuota: Q(27397), tier: '1k x 1', body: { prompt: '一杯咖啡的顶视角海报', resolution: '1k', ratio: '1:1', n: 1 } },
  { n: 8, model: 'kling-image-3.0', name: '2张', expectQuota: Q(27397 * 2), tier: '1k x 2', body: { prompt: '一杯咖啡的顶视角海报', resolution: '1k', ratio: '1:1', n: 2 } },
  { n: 9, model: 'kling-image-3.0-omni', name: '1k', expectQuota: Q(27397), tier: '1k', body: { prompt: '新鲜草莓塔', resolution: '1k', ratio: '1:1', n: 1 } },
  { n: 10, model: 'kling-image-3.0-omni', name: '4k', expectQuota: Q(54795), tier: '4k', body: { prompt: '新鲜草莓塔', resolution: '4k', ratio: '1:1', n: 1 } },
  { n: 11, model: 'kling-image-o1', name: '2k', expectQuota: Q(27397), tier: '2k', body: { prompt: '木星轨道上的太空站', resolution: '2k', ratio: '16:9', n: 1 } },
  { n: 12, model: 'kling-image-2.1', name: '文生图', expectQuota: Q(13699), tier: '1k 文生图', body: { prompt: '汉服女子站在古典庭院中', resolution: '1k', ratio: '3:4', n: 1 } },
  { n: 13, model: 'kling-image-2.1', name: '单图参考', expectQuota: Q(27397), tier: '2k 单图', body: { prompt: '保持人物姿势，改为雪山背景', resolution: '2k', ratio: '9:16', n: 1, image: [IMG] } },
  { n: 14, model: 'kling-image-scene', name: '扩图', expectQuota: Q(27397), tier: '固定扩图', body: { prompt: '以原图风格自然延展画面', n: 1, image: [IMG], extra_fields: { scene_type: 'image_expand', expansion: { up: 0.2, down: 0.2, left: 0.4, right: 0.4 } } } },
  { n: 15, model: 'vidu-image-q2', name: '文生图1080p', expectQuota: Q(25753), tier: '1080p 文生图', body: { prompt: '仙鹤在荷花池边展翅', resolution: '1080p', ratio: '9:16', n: 1 } },
  { n: 16, model: 'vidu-image-q2', name: '参考图2K', expectQuota: Q(51370), tier: '2K 1张参考图', body: { prompt: '按参考图风格画一张海边日落', resolution: '2K', ratio: '16:9', n: 1, image: [IMG] } },
  { n: 17, model: 'qwen-image-0925', name: '1024x1024', expectQuota: Q(41096), tier: '1K', body: { prompt: '赛博朋克夜市', size: '1024x1024', n: 1 } },
  { n: 18, model: 'qwen-image-0925', name: '1536x1536', expectQuota: Q(52055), tier: '2K', body: { prompt: '晨雾中的竹林', size: '1536x1536', n: 1 } },
  { n: 19, model: 'hunyuan-image-3.0', name: '固定价', expectQuota: Q(27397), tier: '固定 0.200', body: { prompt: '元宵节灯笼', size: '1024x1024', n: 1 } },
  { n: 20, model: 'doubao-seedream-5-0-260128', name: '1张', expectQuota: Q(30137), tier: '固定 0.220', body: { prompt: '一只白瓷茶杯放在木桌上', resolution: '2K', ratio: '1:1', n: 1 } },
  { n: 21, model: 'doubao-seedream-5-0-260128', name: '2张', expectQuota: Q(30137 * 2), tier: '固定 0.220 x 2', body: { prompt: '一只白瓷茶杯放在木桌上', resolution: '2K', ratio: '1:1', n: 2 } },
  { n: 22, model: 'doubao-seedream-4-5-251128', name: '固定价', expectQuota: Q(34247), tier: '固定 0.250', body: { prompt: '雨后的青石巷', resolution: '2K', ratio: '3:4', n: 1 } },
  { n: 23, model: 'doubao-seedream-4-0-250828', name: '4K同价', expectQuota: Q(27397), tier: '固定 0.200，传 4K', body: { prompt: '雨后的青石巷', resolution: '4K', ratio: '1:1', n: 1 } },
  { n: 24, model: 'gpt-image-2', name: 'low-1K', expectQuota: Q(6164), tier: 'low 1K', body: { prompt: '一只橘猫', resolution: '1K', ratio: '1:1', quality: 'low', n: 1 } },
  { n: 25, model: 'gpt-image-2', name: 'medium-2K-参考图', expectQuota: Q(109973 + 13699), tier: 'medium 2K + 1 张参考图', body: { prompt: '按参考图风格画一只橘猫', resolution: '2K', ratio: '1:1', quality: 'medium', n: 1, image: [IMG] } },
  { n: 26, model: 'OG-image-2', name: 'high-1K', expectQuota: Q(216849), tier: 'high 1K', body: { prompt: '一只橘猫', resolution: '1K', ratio: '1:1', quality: 'high', n: 1 } },
  { n: 27, model: 'OG-image-2.6-flare', name: '文生图', expectQuota: 500000, tier: '按量预扣，输入价一半', tokenSettle: true, body: { prompt: '一只橘猫坐在窗台上', resolution: '1K', ratio: '1:1', quality: 'low', n: 1 } },
  { n: 28, model: 'OG-image-2.6-sunburst', name: '参考图', expectQuota: 500000, tier: '按量预扣，输入价一半', tokenSettle: true, body: { prompt: '按参考图风格画一只橘猫', resolution: '1K', ratio: '1:1', quality: 'low', n: 1, image: [IMG] } },
  { n: 29, model: 'OG-image-2.5-flare', name: '文生图', expectQuota: 500000, tier: '按量预扣，输入价一半', tokenSettle: true, body: { prompt: '一只橘猫坐在窗台上', resolution: '1K', ratio: '1:1', quality: 'low', n: 1 } },
  { n: 30, model: 'OG-image-2.5-sunburst', name: '文生图', expectQuota: 500000, tier: '按量预扣，输入价一半', tokenSettle: true, body: { prompt: '一只橘猫坐在窗台上', resolution: '1K', ratio: '1:1', quality: 'low', n: 1 } },
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

function amountOfQuota(quota) {
  return (quota / 500000 * 7.3).toFixed(6);
}

function quotaOfAmount(amount) {
  const n = Number(amount);
  if (!Number.isFinite(n)) return null;
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

// 定价页展示的人民币单价。后台按量填的是美元，展示时乘 7.3。
const TOKEN_CNY_PER_MILLION = { text: 29.2, cachedText: 7.3, image: 46.72, cachedImage: 11.68, imageOut: 175.2 };

function findUsage(node) {
  if (!node || typeof node !== 'object') return null;
  if (node.usage && typeof node.usage === 'object' && (node.usage.input_tokens != null || node.usage.total_tokens != null)) {
    return node.usage;
  }
  if (node.data && typeof node.data === 'object' && !Array.isArray(node.data)) {
    const nested = findUsage(node.data);
    if (nested) return nested;
  }
  return null;
}

function splitUpstreamUsage(usage) {
  const input = usage.input_tokens_details || {};
  const output = usage.output_tokens_details || {};
  let text = Number(input.text_tokens || 0);
  let cachedText = Number(input.cached_tokens || 0);
  let image = Number(input.image_tokens || 0);
  let cachedImage = Number(input.cached_image_tokens || 0);
  let imageOut = Number(output.image_tokens || 0);
  const textOut = Number(output.text_tokens || 0);
  if (cachedText > text) cachedText = text;
  if (cachedImage > image) cachedImage = image;
  text -= cachedText;
  image -= cachedImage;
  if (imageOut === 0 && textOut === 0) imageOut = Number(usage.output_tokens || 0);
  else imageOut += textOut;
  return { text, cachedText, image, cachedImage, imageOut };
}

function tokenUsage(node) {
  const usage = findUsage(node);
  if (!usage || !usage.input_tokens_details) return null;
  return usage;
}

// 表达式先得到美元计价的 token 合计，额度 = round(合计 * 0.5)，展示金额再乘 7.3。
function settledFromUsage(parts) {
  const usd = parts.text * 4 + parts.cachedText * 1 + parts.image * 6.4
    + parts.cachedImage * 1.6 + parts.imageOut * 24;
  const quota = Math.round(usd * 0.5);
  return { usd, quota, amount: (quota / 500000 * 7.3).toFixed(6) };
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

function appendUpstream(lines, upstream, finalAmount) {
  lines.push('#### 原始上游响应');
  lines.push('```json');
  lines.push(dump(upstream));
  lines.push('```');
  lines.push('');
  const usage = tokenUsage(upstream);
  if (!usage) return;
  const parts = splitUpstreamUsage(usage);
  const settled = settledFromUsage(parts);
  const ok = finalAmount === settled.amount;
  const rate = TOKEN_CNY_PER_MILLION;
  const row = (name, tokens, price) => `| ${name} | ${tokens} | ${price.toFixed(4)} | ${(tokens * price / 1e6).toFixed(6)} |`;
  lines.push('#### 按上游用量核算');
  lines.push('| 项目 | token | 单价（元/1M） | 金额（元） |');
  lines.push('|---|---:|---:|---:|');
  lines.push(row('文本输入', parts.text, rate.text));
  lines.push(row('缓存文本输入', parts.cachedText, rate.cachedText));
  lines.push(row('图片输入', parts.image, rate.image));
  lines.push(row('缓存图片输入', parts.cachedImage, rate.cachedImage));
  lines.push(row('图片输出', parts.imageOut, rate.imageOut));
  lines.push(`| 额度取整后 | ${settled.quota} | — | ${settled.amount} |`);
  lines.push(`| 平台 amount | — | — | ${finalAmount || '-'} |`);
  lines.push(`| 是否一致 | — | — | ${ok ? '✅ 一致' : '❌ 不一致'} |`);
  lines.push('');
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
  lines.push(`> 本文件由 \`image_billing_batch.js\` 生成。价格文档 \`v2026.11.07\`。`);
  lines.push('');
  lines.push(`> 执行时间：${record.startedAt}  |  模型：\`${c.model}\`  |  档位：${c.tier}`);
  lines.push('');
  lines.push('### 调用参数');
  lines.push('```json');
  lines.push(dump({ model: c.model, ...c.body }));
  lines.push('```');
  lines.push('');

  const body = c.body;
  lines.push('### 价格变量核对');
  lines.push('| 变量 | 请求值 | 说明 |');
  lines.push('|---|---|---|');
  lines.push(`| resolution | ${body.resolution || '未传'} | 分辨率档位 |`);
  lines.push(`| ratio | ${body.ratio || '未传'} | 宽高比 |`);
  lines.push(`| size | ${body.size || '未传'} | 宽x高 |`);
  lines.push(`| quality | ${body.quality || '未传'} | 渲染精度 |`);
  lines.push(`| n | ${body.n != null ? body.n : '未传'} | 张数 |`);
  lines.push(`| image | ${Array.isArray(body.image) ? `${body.image.length} 张` : '无'} | 参考图 |`);
  lines.push('');

  lines.push('### 预期扣费');
  lines.push('| 项目 | 值 |');
  lines.push('|---|---|');
  lines.push('| 预期 HTTP 状态 | 200 |');
  lines.push(`| 预期 Quota | **${fmt(c.expectQuota)}** |`);
  lines.push(`| 预期 USD | $${(c.expectQuota / 500000).toFixed(4)} |`);
  lines.push(`| 预期 RMB | ¥${amountOfQuota(c.expectQuota)} |`);
  if (c.tokenSettle) lines.push('| 结算方式 | 提交预扣 $1，成功后按 token 重算 |');
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
    const upstreamUsage = tokenUsage(record.upstream);
    const upstreamSettled = upstreamUsage ? settledFromUsage(splitUpstreamUsage(upstreamUsage)) : null;
    let finalNote = '';
    if (upstreamSettled && record.finalStatus === 'SUCCESS') {
      finalNote = finalAmount === upstreamSettled.amount
        ? '✅ 与上游用量按定价页核算一致'
        : `❌ 与上游用量核算 ¥${upstreamSettled.amount} 不一致`;
    } else if (c.tokenSettle && record.finalStatus === 'SUCCESS') {
      finalNote = finalAmount === '7.300000'
        ? '⚠️ 成功后金额仍是预扣，没有按 token 重算'
        : '成功后已离开预扣金额，按 token 结算';
    } else if (record.finalStatus && record.finalStatus !== 'SUCCESS') {
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

    if (record.upstream) appendUpstream(lines, record.upstream, finalAmount);

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

function mergeSummary(submitted) {
  const indexPath = path.join(OUT_DIR, '_汇总.md');
  const fresh = submitted.map((r) => {
    const pre = r.preQuota === r.c.expectQuota ? '符合' : '不符';
    const finalAmount = r.poll?.json?.data?.amount || '-';
    const finalQuota = quotaOfAmount(finalAmount);
    let finalMark = '-';
    if (!r.finalStatus || r.finalStatus === '未提交') finalMark = '未完成';
    else if (r.finalStatus !== 'SUCCESS') finalMark = '失败退款';
    else if (r.c.tokenSettle) finalMark = finalAmount === '7.300000' ? '未按 token 重算' : '按 token 结算';
    else if (finalQuota === r.c.expectQuota) finalMark = '符合';
    else finalMark = '不符';
    return `| ${r.c.n} | \`${r.c.model}\` | ${r.c.name} | ${r.finalStatus || '未提交'} | ${fmt(r.preQuota)} | ${fmt(r.c.expectQuota)} | ${pre} | ${finalAmount} | ${finalMark} |`;
  });
  if (!process.env.IMAGE_TEST_CASES || !fs.existsSync(indexPath)) {
    const rows = fresh;
    const index = [
      '# 图像扣费测试汇总 v2026.11.07',
      '',
      `执行时间：${now()}`,
      '',
      '预期金额按价格文档 `v2026.11.07` 的表达式、分组倍率 1。最终符合指成功任务的 `amount` 与预期 Quota 换算一致。',
      '',
      '| # | 模型 | 用例 | 状态 | 预扣 Quota | 预期 Quota | 预扣 | 最终 amount | 最终 |',
      '|---|---|---|---|---:|---:|---|---:|---|',
      ...rows,
      '',
    ].join('\n');
    fs.writeFileSync(indexPath, index);
    return;
  }
  let text = fs.readFileSync(indexPath, 'utf8');
  if (!text.includes('| 最终 |')) {
    text = text.replace(
      '| # | 模型 | 用例 | 状态 | 预扣 Quota | 预期 Quota | 预扣 | 最终 amount |',
      '| # | 模型 | 用例 | 状态 | 预扣 Quota | 预期 Quota | 预扣 | 最终 amount | 最终 |',
    );
    text = text.replace(
      '|---|---|---|---|---:|---:|---|---:|',
      '|---|---|---|---|---:|---:|---|---:|---|',
    );
  }
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
  const cases = selectedCases();
  const submitConcurrency = process.env.IMAGE_TEST_CASES ? 1 : SUBMIT_CONCURRENCY;
  console.log(`输出目录 ${OUT_DIR}`);
  console.log(`用例 ${cases.length}，提交并发 ${submitConcurrency}`);

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
    if (process.env.IMAGE_TEST_CASES) await sleep(2000);
    return record;
  }

  const submitted = await pool(cases, submitConcurrency, submitOne);
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
    console.log(`完成 ${record.c.n} ${record.finalStatus} ${mark} pre=${record.preQuota} expect=${record.c.expectQuota}`);
    return record;
  }

  await pool(runnable, POLL_CONCURRENCY, finish);
  for (const record of submitted.filter((r) => !r.taskId)) {
    await attachLogs(record);
    writeCase(record.c, record);
  }

  mergeSummary(submitted);
  const ok = submitted.filter((r) => r.preQuota === r.c.expectQuota).length;
  console.log(`预扣符合 ${ok}/${submitted.length}`);
  console.log('全部用例已写完');
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
