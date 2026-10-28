#!/usr/bin/env node
/**
 * 新参数结构视频计费测试。
 * 并发提交，每条用例一个 markdown。API Key 从环境变量 VIDEO_TEST_API_KEY 读取。
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

const IMG = 'https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906?authorization=bce-auth-v1%2FALTAKMD70FiEhTeFimgJH1yjAb%2F2026-10-28T07%3A25%3A39Z%2F600%2F%2F83ecd58fe0f9818278fae698dd920bdc4b1ee0fe7012fc23f42fa5df0a7bb5e7';
const VID = 'https://aigc-bucket.su.bcebos.com/v1/objects/0d/0dca41d01e3579e43d30ae64fd30126aa056963fe85c5f254b81a3868a37298d?authorization=bce-auth-v1%2FALTAKMD70FiEhTeFimgJH1yjAb%2F2026-10-28T07%3A26%3A04Z%2F600%2F%2F03d1d6955391e462eadcb73c5ef846a63b2ba1543d2d63bdbae041429f5767f2';

const OUT_DIR = path.join(__dirname, '..', '模型测试结果记录文档', '20261028-新参数');
const SUBMIT_CONCURRENCY = 8;
const POLL_CONCURRENCY = 12;
const POLL_INTERVAL_MS = 15000;
const POLL_TIMEOUT_MS = 12 * 60 * 1000;
const DONE = new Set(['SUCCESS', 'FAILURE', 'FAILED', 'CANCELLED', 'SUCCEEDED']);

const image = (role) => ({
  type: 'image_url',
  ...(role ? { role } : {}),
  image_url: { url: IMG },
});
const video = (role) => ({
  type: 'video_url',
  role: role || 'reference_video',
  video_url: { url: VID },
});

function Q(expr) { return Math.round(expr / 2); }

const CASES = [
  { n: 1, model: 'kling-video-1.6', name: '默认', expectExpr: 273975, body: { prompt: '夏日海滩，海浪轻拍礁石，写实风格', ratio: '16:9' } },
  { n: 2, model: 'kling-video-1.6', name: '1080P x 10秒', expectExpr: 958900, body: { prompt: '夏日海滩，海浪轻拍礁石，写实风格', duration: 10, resolution: '1080P', ratio: '16:9' } },
  { n: 3, model: 'kling-video-2.5', name: '默认', expectExpr: 205480, body: { prompt: '霓虹雨夜，摩托车驶过积水路面', ratio: '16:9' } },
  { n: 4, model: 'kling-video-2.5', name: '1080P x 5秒', expectExpr: 342465, body: { prompt: '霓虹雨夜，摩托车驶过积水路面', duration: 5, resolution: '1080P', ratio: '16:9' } },
  { n: 5, model: 'kling-video-2.6', name: '无声默认', expectExpr: 205480, body: { prompt: '竹林微风，阳光穿过竹叶', ratio: '16:9' } },
  { n: 6, model: 'kling-video-2.6', name: '有声1080P x 5秒', expectExpr: 684930, body: { prompt: '竹林微风，阳光穿过竹叶', duration: 5, resolution: '1080P', ratio: '16:9', metadata: { audio_generation: 'Enabled' } } },
  { n: 7, model: 'kling-video-o1', name: '文生', expectExpr: 410960, body: { prompt: '城市天际线从白天过渡到黄昏', ratio: '16:9' } },
  { n: 8, model: 'kling-video-o1', name: '只有参考图', expectExpr: 410960, body: { prompt: '保持画面风格，镜头缓慢推进', metadata: { content: [image('reference_image')] } } },
  { n: 9, model: 'kling-video-o1', name: '参考视频1080P x 5秒', expectExpr: 821920, body: { prompt: '沿用参考视频的运镜，色调偏暖', duration: 5, resolution: '1080P', metadata: { content: [video()] } } },
  { n: 10, model: 'kling-video-3.0', name: '无声默认', expectExpr: 410960, body: { prompt: '雨后街道，霓虹倒映在地面积水上', ratio: '16:9' } },
  { n: 11, model: 'kling-video-3.0', name: '有声1080P x 5秒', expectExpr: 821920, body: { prompt: '雨后街道，霓虹倒映在地面积水上', duration: 5, resolution: '1080P', ratio: '16:9', metadata: { audio_generation: 'Enabled' } } },
  { n: 12, model: 'kling-video-3.0-omni', name: '无参考视频无声', expectExpr: 410960, body: { prompt: '一只猫从窗台跳下，慢动作', ratio: '16:9' } },
  { n: 13, model: 'kling-video-3.0-omni', name: '无参考视频有声', expectExpr: 547945, body: { prompt: '一只猫从窗台跳下，慢动作', duration: 5, resolution: '720P', ratio: '16:9', metadata: { audio_generation: 'Enabled' } } },
  { n: 14, model: 'kling-video-3.0-omni', name: '参考视频1080P x 5秒', expectExpr: 821920, body: { prompt: '沿用参考视频的主体动作', duration: 5, resolution: '1080P', metadata: { audio_generation: 'Enabled', content: [video()] } } },
  { n: 15, model: 'vidu-video-q2', name: '文生', expectExpr: 219180, body: { prompt: '山间公路，云雾缓慢流动', ratio: '16:9' } },
  { n: 16, model: 'vidu-video-q2', name: '不写role的参考图1080P x 8秒', expectExpr: 898632, body: { prompt: '人物转头微笑，背景轻微虚化', duration: 8, resolution: '1080P', metadata: { content: [image()] } } },
  { n: 17, model: 'vidu-video-q2-pro', name: '首帧', expectExpr: 239725, body: { prompt: '人物从座位上站起来', metadata: { content: [image('first_frame')] } } },
  { n: 18, model: 'vidu-video-q2-pro', name: '参考图1080P x 8秒', expectExpr: 986304, body: { prompt: '保持人物外观，镜头缓慢环绕', duration: 8, resolution: '1080P', metadata: { content: [image('reference_image')] } } },
  { n: 19, model: 'vidu-video-q2-turbo', name: '默认', expectExpr: 171235, body: { prompt: '海面日落，镜头平稳横移', ratio: '16:9' } },
  { n: 20, model: 'vidu-video-q2-turbo', name: '1080P x 8秒', expectExpr: 515072, body: { prompt: '海面日落，镜头平稳横移', duration: 8, resolution: '1080P', ratio: '16:9' } },
  { n: 21, model: 'vidu-video-q3', name: '默认', expectExpr: 428080, body: { prompt: '森林小径，光斑随风移动', ratio: '16:9' } },
  { n: 22, model: 'vidu-video-q3', name: '1080P x 8秒', expectExpr: 856984, body: { prompt: '森林小径，光斑随风移动', duration: 8, resolution: '1080P', ratio: '16:9' } },
  { n: 23, model: 'vidu-video-q3-pro', name: '默认', expectExpr: 535615, body: { prompt: '雪山日出，云层缓慢流动', ratio: '16:9' } },
  { n: 24, model: 'vidu-video-q3-pro', name: '1080P x 8秒', expectExpr: 1027944, body: { prompt: '雪山日出，云层缓慢流动', duration: 8, resolution: '1080P', ratio: '16:9' } },
  { n: 25, model: 'vidu-video-q3-turbo', name: '文生默认', expectExpr: 256850, body: { prompt: '城市夜景，车灯拉出光轨', ratio: '16:9' } },
  { n: 26, model: 'vidu-video-q3-turbo', name: '参考图1080P x 8秒', expectExpr: 444928, body: { prompt: '保持参考图风格，镜头缓慢拉远', duration: 8, resolution: '1080P', metadata: { content: [image('reference_image')] } } },
  { n: 27, model: 'hailuo-video-02', name: '默认', expectExpr: 271230, body: { prompt: '雨中的咖啡馆橱窗，行人撑伞经过' } },
  { n: 28, model: 'hailuo-video-02', name: '1080P x 10秒', expectExpr: 794520, body: { prompt: '雨中的咖啡馆橱窗，行人撑伞经过', duration: 10, resolution: '1080P' } },
  { n: 29, model: 'hailuo-video-2.3', name: '默认', expectExpr: 271230, body: { prompt: '雨中的咖啡馆橱窗，行人撑伞经过' } },
  { n: 30, model: 'hailuo-video-2.3-fast', name: '默认', expectExpr: 184932, body: { prompt: '桌面上的咖啡杯，热气缓缓升起' } },
  { n: 31, model: 'hailuo-video-2.3-fast', name: '1080P x 10秒', expectExpr: 527400, body: { prompt: '桌面上的咖啡杯，热气缓缓升起', duration: 10, resolution: '1080P' } },
  { n: 32, model: 'hailuo-video-h3', name: '默认', expectExpr: 410958, body: { prompt: '海边栈道，镜头跟着人物向前走', ratio: '16:9' } },
  { n: 33, model: 'hailuo-video-h3', name: '1080P x 10秒', expectExpr: 890410, body: { prompt: '海边栈道，镜头跟着人物向前走', duration: 10, resolution: '1080P', ratio: '16:9' } },
  { n: 34, model: 'veo-video-3.1', name: '无声', expectExpr: 1643832, body: { prompt: '沙漠公路上的一辆旧汽车，热浪扭曲远景', resolution: '720P', ratio: '16:9' } },
  { n: 35, model: 'veo-video-3.1', name: '有声', expectExpr: 3287672, body: { prompt: '沙漠公路上的一辆旧汽车，热浪扭曲远景', resolution: '1080P', ratio: '16:9', metadata: { audio_generation: 'Enabled' } } },
  { n: 36, model: 'veo-video-3.1-fast', name: '无声', expectExpr: 821920, body: { prompt: '厨房里的面团被慢慢拉开', resolution: '720P', ratio: '16:9' } },
  { n: 37, model: 'veo-video-3.1-fast', name: '有声', expectExpr: 1232880, body: { prompt: '厨房里的面团被慢慢拉开', resolution: '720P', ratio: '16:9', metadata: { audio_generation: 'Enabled' } } },
  { n: 38, model: 'sora-video-2.0', name: '默认', expectExpr: 821920, body: { prompt: '纸船在水面上漂过桥洞', ratio: '16:9' } },
  { n: 39, model: 'sora-video-2.0', name: 'duration=12', expectExpr: 1232880, body: { prompt: '纸船在水面上漂过桥洞', duration: 12, resolution: '720P', ratio: '16:9' } },
  { n: 40, model: 'hunyuan-video-1.5', name: '默认', expectExpr: 205480, body: { prompt: '江南水乡，乌篷船慢慢划过' } },
  { n: 41, model: 'hunyuan-video-1.5', name: '1080P x 10秒', expectExpr: 684930, body: { prompt: '江南水乡，乌篷船慢慢划过', duration: 10, resolution: '1080P' } },
  { n: 42, model: 'mingmou-video-1.0', name: '默认', expectExpr: 205480, body: { prompt: '雪地里的一盏路灯，雪花落下' } },
  { n: 43, model: 'pixverse-video-v5.6', name: '默认', expectExpr: 215755, body: { prompt: '一只鸟从枝头飞起', ratio: '16:9' } },
  { n: 44, model: 'pixverse-video-v5.6', name: '1080P x 10秒', expectExpr: 719180, body: { prompt: '一只鸟从枝头飞起', duration: 10, resolution: '1080P', ratio: '16:9' } },
  { n: 45, model: 'pixverse-video-v6', name: '无声默认', expectExpr: 180820, body: { prompt: '风吹过麦田，麦浪起伏', ratio: '16:9' } },
  { n: 46, model: 'pixverse-video-v6', name: '有声1080P x 8秒', expectExpr: 739728, body: { prompt: '风吹过麦田，麦浪起伏', duration: 8, resolution: '1080P', ratio: '16:9', metadata: { audio_generation: 'Enabled' } } },
  { n: 47, model: 'pixverse-video-c1', name: '无声默认', expectExpr: 200685, body: { prompt: '夜晚的码头，渔灯轻轻晃动', ratio: '16:9' } },
  { n: 48, model: 'pixverse-video-c1', name: '有声1080P x 8秒', expectExpr: 771504, body: { prompt: '夜晚的码头，渔灯轻轻晃动', duration: 8, resolution: '1080P', ratio: '16:9', metadata: { audio_generation: 'Enabled' } } },
  { n: 49, model: 'doubao-seedance-2-5-260628', name: '文生720p', expectNote: '无 OtherRatios，有效单价 ¥73.50/1M', body: { prompt: '清晨海岸公路，镜头缓慢推进', duration: 5, resolution: '720p', ratio: '16:9', metadata: { generate_audio: true, watermark: false } } },
  { n: 50, model: 'doubao-seedance-2-5-260628', name: '文生1080p', expectNote: 'resolution≈1.09524，有效单价 ¥80.50/1M', body: { prompt: '霓虹雨夜，摩托车驶过积水路面', duration: 5, resolution: '1080p', ratio: '16:9', metadata: { generate_audio: true } } },
  { n: 51, model: 'doubao-seedance-2-5-260628', name: '参考视频720p', expectNote: 'video_input=0.6，有效单价 ¥44.10/1M', needsVideoAsset: true, body: { prompt: '保持参考视频的运镜，色调偏暖', duration: 5, resolution: '720p', ratio: 'adaptive', metadata: { generate_audio: true, content: [video()] } } },
  { n: 52, model: 'doubao-seedance-2-0-260128', name: '文生720p', expectNote: '无系数，有效单价 ¥48.30/1M', body: { prompt: '山间公路，云雾缓慢流动', duration: 5, resolution: '720p', ratio: '16:9', metadata: { generate_audio: true } } },
  { n: 53, model: 'doubao-seedance-2-0-260128', name: '文生4k', expectNote: 'resolution≈0.5652，有效单价 ¥27.30/1M', body: { prompt: '山间公路，云雾缓慢流动', duration: 5, resolution: '4k', ratio: '16:9', metadata: { generate_audio: true } } },
  { n: 54, model: 'doubao-seedance-2-0-260128', name: '参考视频720p', expectNote: 'video_input≈0.6087，有效单价 ¥29.40/1M', needsVideoAsset: true, body: { prompt: '保持参考视频的运镜，色调偏暖', duration: 5, resolution: '720p', ratio: 'adaptive', metadata: { generate_audio: true, content: [video()] } } },
  { n: 55, model: 'doubao-seedance-2-0-fast-260128', name: '文生720p', expectNote: '无系数，有效单价 ¥38.85/1M', body: { prompt: '一只猫趴在窗台上', duration: 5, resolution: '720p', ratio: '16:9', metadata: { generate_audio: true } } },
  { n: 56, model: 'doubao-seedance-2-0-fast-260128', name: '参考视频720p', expectNote: 'video_input≈0.5946，有效单价 ¥23.10/1M', needsVideoAsset: true, body: { prompt: '保持参考视频的运镜', duration: 5, resolution: '720p', ratio: 'adaptive', metadata: { generate_audio: true, content: [video()] } } },
  { n: 57, model: 'doubao-seedance-2-0-mini-260615', name: '文生720p', expectNote: '无系数，有效单价 ¥24.15/1M', body: { prompt: '桌面上的咖啡杯，热气升起', duration: 5, resolution: '720p', ratio: '1:1', metadata: { generate_audio: true } } },
  { n: 58, model: 'doubao-seedance-2-0-mini-260615', name: '参考视频720p', expectNote: 'video_input≈0.6087，有效单价 ¥14.70/1M', needsVideoAsset: true, body: { prompt: '保持参考视频的运镜', duration: 5, resolution: '720p', ratio: 'adaptive', metadata: { generate_audio: true, content: [video()] } } },
  { n: 59, model: 'doubao-seedance-1-5-pro-251215', name: '有声', expectNote: '不传 generate_audio，有声 ¥16.00/1M', body: { prompt: '海边栈道，镜头跟着人物向前', duration: 5, resolution: '720p', ratio: '16:9' } },
  { n: 60, model: 'doubao-seedance-1-5-pro-251215', name: '无声', expectNote: 'generate_audio=false，silent_video=0.5，¥8.00/1M', body: { prompt: '海边栈道，镜头跟着人物向前', duration: 5, resolution: '720p', ratio: '16:9', metadata: { generate_audio: false } } },
];

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
        resolve({ status: res.statusCode, json, text });
      });
    });
    req.on('error', reject);
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

function fileFor(c) {
  const safe = `${String(c.n).padStart(2, '0')}-${c.model}-${c.name}`.replace(/[^\w.\-\u4e00-\u9fa5]+/g, '-');
  return path.join(OUT_DIR, `${safe}.md`);
}

function statusOf(json) {
  return json?.status || json?.data?.status || json?.data?.data?.status || '';
}

function taskIdOf(json) {
  return json?.task_id || json?.id || json?.data?.task_id || '';
}

async function fetchLogs() {
  const r = await request('GET', '/api/log/token?p=1&page_size=200');
  return Array.isArray(r.json?.data) ? r.json.data : [];
}

function pickLogs(logs, model, fromSec, toSec) {
  return logs.filter((l) => l.model_name === model && l.created_at >= fromSec - 2 && l.created_at <= toSec + 5);
}

function writeCase(c, extra) {
  const expectQuota = c.expectExpr != null ? Q(c.expectExpr) : null;
  const lines = [];
  lines.push(`# ${c.n}. ${c.model} ${c.name}`);
  lines.push('');
  lines.push(`- 预期：${c.expectNote || `表达式值 ${c.expectExpr}，quota ${expectQuota}`}`);
  lines.push('');
  lines.push('## 输入 JSON');
  lines.push('');
  lines.push('```json');
  lines.push(JSON.stringify({ model: c.model, ...c.body }, null, 2));
  lines.push('```');
  lines.push('');
  lines.push('## 提交');
  lines.push('');
  lines.push('- 方法：`POST /v1/video/generations`');
  lines.push(`- HTTP ${extra.submit.status}`);
  lines.push('');
  lines.push('```json');
  lines.push(JSON.stringify(extra.submit.json, null, 2));
  lines.push('```');
  lines.push('');
  lines.push('## 查询');
  lines.push('');
  lines.push(`- 方法：\`GET /v1/video/generations/${extra.taskId || ''}\``);
  lines.push(`- 轮询次数：${extra.polls || 0}`);
  lines.push(`- 最终状态：${extra.finalStatus || ''}`);
  lines.push('');
  lines.push('```json');
  lines.push(JSON.stringify(extra.pollJson || null, null, 2));
  lines.push('```');
  lines.push('');
  lines.push('## 消费日志');
  lines.push('');
  lines.push('```json');
  lines.push(JSON.stringify(extra.logs || [], null, 2));
  lines.push('```');
  lines.push('');
  fs.writeFileSync(fileFor(c), lines.join('\n'));
}

async function registerPortrait() {
  const note = { imageAsset: '', videoAsset: '', error: '' };
  try {
    const group = await request('POST', '/v1/portrait/groups', {
      name: `billing-test-${Date.now()}`,
      description: '视频计费测试素材',
      group_type: 'AIGC',
    });
    const groupId = group.json?.data?.group_id;
    if (!groupId) {
      note.error = `创建素材组失败 HTTP ${group.status} ${JSON.stringify(group.json)}`;
      return note;
    }
    const image = await request('POST', '/v1/portrait/assets', {
      group_id: groupId, url: IMG, asset_type: 'Image', name: 'billing-image',
    });
    const video = await request('POST', '/v1/portrait/assets', {
      group_id: groupId, url: VID, asset_type: 'Video', name: 'billing-video',
    });
    note.imageAsset = image.json?.data?.asset_id || '';
    note.videoAsset = video.json?.data?.asset_id || '';
    const deadline = Date.now() + 90_000;
    const pending = [note.imageAsset, note.videoAsset].filter(Boolean);
    while (Date.now() < deadline && pending.length) {
      await sleep(5000);
      for (const id of [...pending]) {
        const detail = await request('GET', `/v1/portrait/assets/${id}`);
        const status = detail.json?.data?.status;
        if (status === 'Active' || status === 'Failed') {
          pending.splice(pending.indexOf(id), 1);
          if (status !== 'Active') note.error += `${id}=${status};`;
        }
      }
    }
    if (pending.length) note.error += `审核未完成 ${pending.join(',')}`;
  } catch (err) {
    note.error = String(err);
  }
  return note;
}

function useAsset(c, portrait) {
  if (!c.needsVideoAsset || !portrait.videoAsset || portrait.error.includes(portrait.videoAsset)) return c;
  const copy = JSON.parse(JSON.stringify(c));
  const item = copy.body.metadata.content[0];
  item.video_url.url = `asset://${portrait.videoAsset}`;
  copy.assetNote = portrait.videoAsset;
  return copy;
}

async function main() {
  fs.mkdirSync(OUT_DIR, { recursive: true });
  console.log(`输出目录 ${OUT_DIR}`);
  console.log('登记人像库素材（豆包参考视频用，不挡住腾讯提交）');
  const portraitPromise = registerPortrait();

  const tencent = CASES.filter((c) => c.n <= 48);
  const doubaoText = CASES.filter((c) => c.n > 48 && !c.needsVideoAsset);
  const doubaoVideo = CASES.filter((c) => c.needsVideoAsset);

  async function runOne(c) {
    const started = Math.floor(Date.now() / 1000);
    const payload = { model: c.model, ...c.body };
    let submit;
    try {
      submit = await request('POST', '/v1/video/generations', payload);
    } catch (err) {
      submit = { status: 0, json: { error: String(err) } };
    }
    const taskId = taskIdOf(submit.json);
    console.log(`提交 ${c.n} ${c.model} HTTP ${submit.status} task=${taskId || '-'}`);
    const record = { c, submit, taskId, started, polls: 0, finalStatus: '', pollJson: null, logs: [] };
    writeCase(c, record);
    return record;
  }

  const firstWave = await pool([...tencent, ...doubaoText], SUBMIT_CONCURRENCY, runOne);
  const portrait = await portraitPromise;
  fs.writeFileSync(path.join(OUT_DIR, '_portrait.json'), JSON.stringify(portrait, null, 2));
  console.log('人像库', JSON.stringify(portrait));
  const videoCases = doubaoVideo.map((c) => useAsset(c, portrait));
  const secondWave = await pool(videoCases, SUBMIT_CONCURRENCY, runOne);
  const records = [...firstWave, ...secondWave].filter((r) => r && r.taskId);

  async function pollOne(record) {
    const deadline = Date.now() + POLL_TIMEOUT_MS;
    let last = null;
    while (Date.now() < deadline) {
      record.polls++;
      const r = await request('GET', `/v1/video/generations/${encodeURIComponent(record.taskId)}`);
      last = r.json;
      const status = String(statusOf(last) || '').toUpperCase();
      if (DONE.has(status)) {
        record.finalStatus = status;
        record.pollJson = last;
        break;
      }
      await sleep(POLL_INTERVAL_MS);
    }
    if (!record.finalStatus) {
      record.finalStatus = 'TIMEOUT';
      record.pollJson = last;
    }
    const logs = await fetchLogs();
    record.logs = pickLogs(logs, record.c.model, record.started, Math.floor(Date.now() / 1000));
    writeCase(record.c, record);
    console.log(`完成 ${record.c.n} ${record.finalStatus} polls=${record.polls}`);
    return record;
  }

  await pool(records, POLL_CONCURRENCY, pollOne);
  const failedSubmit = [...firstWave, ...secondWave].filter((r) => r && !r.taskId);
  for (const record of failedSubmit) {
    const logs = await fetchLogs();
    record.logs = pickLogs(logs, record.c.model, record.started, record.started + 30);
    writeCase(record.c, record);
  }
  console.log('全部用例已写完');
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
