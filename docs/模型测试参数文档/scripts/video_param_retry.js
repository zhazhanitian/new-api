#!/usr/bin/env node
'use strict';

const http = require('http');
const fs = require('fs');
const path = require('path');

const HOST = '127.0.0.1';
const PORT = 9006;
const API_KEY = process.env.VIDEO_TEST_API_KEY;
const DIR = path.join(__dirname, '..', '模型测试结果记录文档', '20261028-新参数');
const VID = 'https://aigc-bucket.su.bcebos.com/v1/objects/0d/0dca41d01e3579e43d30ae64fd30126aa056963fe85c5f254b81a3868a37298d?authorization=bce-auth-v1%2FALTAKMD70FiEhTeFimgJH1yjAb%2F2026-10-28T07%3A26%3A04Z%2F600%2F%2F03d1d6955391e462eadcb73c5ef846a63b2ba1543d2d63bdbae041429f5767f2';
const DONE = new Set(['SUCCESS', 'FAILURE', 'FAILED', 'CANCELLED', 'SUCCEEDED']);
const POLL_TIMEOUT_MS = 12 * 60 * 1000;
const POLL_INTERVAL_MS = 15000;

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
function taskIdOf(json) { return json?.task_id || json?.id || ''; }

function parseInput(text) {
  const m = text.match(/## 输入 JSON\s+```json\s+([\s\S]*?)```/);
  return m ? JSON.parse(m[1]) : null;
}

async function poll(taskId) {
  const deadline = Date.now() + POLL_TIMEOUT_MS;
  let polls = 0;
  let last = null;
  while (Date.now() < deadline) {
    polls++;
    const r = await request('GET', `/v1/video/generations/${encodeURIComponent(taskId)}`);
    last = r.json;
    if (DONE.has(statusOf(last))) return { polls, finalStatus: statusOf(last), pollJson: last };
    await sleep(POLL_INTERVAL_MS);
  }
  return { polls, finalStatus: 'TIMEOUT', pollJson: last };
}

async function logsFor(model, fromSec) {
  const r = await request('GET', '/api/log/token?p=1&page_size=200');
  const list = Array.isArray(r.json?.data) ? r.json.data : [];
  const now = Math.floor(Date.now() / 1000);
  return list.filter((l) => l.model_name === model && l.created_at >= fromSec - 2 && l.created_at <= now + 5);
}

function rewrite(file, input, submit, taskId, poll, logs, note) {
  const title = fs.readFileSync(file, 'utf8').split('\n')[0];
  const expect = (fs.readFileSync(file, 'utf8').match(/^- 预期：.*$/m) || [''])[0];
  const lines = [
    title, '', expect, '',
    note ? `> 补提交：${note}` : '', '',
    '## 输入 JSON', '', '```json', JSON.stringify(input, null, 2), '```', '',
    '## 提交', '', '- 方法：`POST /v1/video/generations`', `- HTTP ${submit.status}`, '',
    '```json', JSON.stringify(submit.json, null, 2), '```', '',
    '## 查询', '', `- 方法：\`GET /v1/video/generations/${taskId || ''}\``,
    `- 轮询次数：${poll.polls || 0}`, `- 最终状态：${poll.finalStatus || ''}`, '',
    '```json', JSON.stringify(poll.pollJson || null, null, 2), '```', '',
    '## 消费日志', '', '```json', JSON.stringify(logs, null, 2), '```', '',
  ];
  fs.writeFileSync(file, lines.filter((x, i, a) => !(x === '' && a[i - 1] === '')).join('\n'));
}

function existingTask(text) {
  const m = text.match(/"task_id": "([^"]+)"/);
  return m ? m[1] : '';
}

async function main() {
  const files = fs.readdirSync(DIR).filter((f) => f.endsWith('.md')).sort();
  const jobs = [];
  const already = [];
  for (const name of files) {
    const file = path.join(DIR, name);
    const text = fs.readFileSync(file, 'utf8');
    const http = (text.match(/- HTTP (\d+)/) || [])[1];
    const input = parseInput(text);
    if (!input) continue;
    if (http === '200') {
      const taskId = existingTask(text);
      if (taskId && !/最终状态：\*\*(SUCCESS|FAILURE|FAILED)/.test(text) && text.includes('最终状态：\n')) {
        already.push({ file, input, taskId, started: Math.floor(Date.now() / 1000) - 120, submit: { status: 200, json: { task_id: taskId } } });
      }
      continue;
    }
    if (http !== '500' && http !== '400') continue;
    if (input.metadata?.content) {
      for (const item of input.metadata.content) {
        if (item.video_url && String(item.video_url.url).startsWith('asset://')) item.video_url.url = VID;
      }
    }
    jobs.push({ file, input, http });
  }
  console.log(`补提交 ${jobs.length} 条，继续轮询 ${already.length} 条`);
  const pending = [];
  for (const job of jobs) {
    const started = Math.floor(Date.now() / 1000);
    const submit = await request('POST', '/v1/video/generations', job.input);
    const taskId = taskIdOf(submit.json);
    console.log(`${path.basename(job.file)} HTTP ${submit.status} ${taskId || submit.json?.code || ''}`);
    if (taskId) pending.push({ ...job, submit, taskId, started });
    else {
      const logs = await logsFor(job.input.model, started);
      rewrite(job.file, job.input, submit, '', { polls: 0, finalStatus: '', pollJson: null }, logs, '限流或上游拒绝后的补提交');
    }
    await sleep(400);
  }
  await Promise.all(pending.map(async (job) => {
    const polled = await poll(job.taskId);
    const logs = await logsFor(job.input.model, job.started);
    rewrite(job.file, job.input, job.submit, job.taskId, polled, logs, '限流后的补提交');
    console.log(`完成补提交 ${path.basename(job.file)} ${polled.finalStatus}`);
  }));
  await Promise.all(already.map(async (job) => {
    const polled = await poll(job.taskId);
    const logs = await logsFor(job.input.model, job.started);
    const text = fs.readFileSync(job.file, 'utf8');
    const head = text.split('## 查询')[0];
    const tail = [
      '## 查询', '',
      `- 方法：\`GET /v1/video/generations/${job.taskId}\``,
      `- 轮询次数：${polled.polls}`,
      `- 最终状态：${polled.finalStatus}`, '',
      '```json', JSON.stringify(polled.pollJson, null, 2), '```', '',
      '## 消费日志', '', '```json', JSON.stringify(logs, null, 2), '```', '',
    ].join('\n');
    fs.writeFileSync(job.file, head + tail);
    console.log(`完成轮询 ${path.basename(job.file)} ${polled.finalStatus}`);
  }));
  console.log('补提交结束');
}

main().catch((err) => { console.error(err); process.exit(1); });
