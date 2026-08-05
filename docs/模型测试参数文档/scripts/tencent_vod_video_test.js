#!/usr/bin/env node
/**
 * ============================================================
 * 腾讯 VOD 生视频模型 — 自动化测试脚本
 * ============================================================
 * 运行环境：Node.js >= 18（直接使用内置 http 模块，无需 npm install）
 *
 * ⚠️ 与测试参数文档强关联（必读）：
 *   文档：../腾讯VOD生视频模型测试参数.md
 *   本脚本的 CASES 必须与文档用例保持一一对应。
 *   修改文档中的用例（序号 / 模型 / 请求体 / 预期扣费 / 边界用例）时，
 *   必须同步修改本文件 CASES；禁止只改一侧。
 *   改完后先 --list / --dry-run 核对，再 live 执行。
 *
 * ┌──────────────────── 使用方式 ─────────────────────────────┐
 * │                                                           │
 * │  node tencent_vod_video_test.js --list                    │
 * │    → 列出全部用例（序号、模型、预期 Quota、标签）             │
 * │                                                           │
 * │  node tencent_vod_video_test.js --case 2                  │
 * │    → 只跑第 2 条用例                                       │
 * │                                                           │
 * │  node tencent_vod_video_test.js --cases 1,5,10            │
 * │    → 跑第 1、5、10 条（逗号分隔）                           │
 * │                                                           │
 * │  node tencent_vod_video_test.js --from 5                  │
 * │    → 从第 5 条跑到最后                                     │
 * │                                                           │
 * │  node tencent_vod_video_test.js --from 5 --to 10          │
 * │    → 跑第 5～10 条（含两端）                               │
 * │                                                           │
 * │  node tencent_vod_video_test.js --model kling-video-1.6   │
 * │    → 跑该模型的全部用例                                    │
 * │                                                           │
 * │  node tencent_vod_video_test.js --tag edge                │
 * │    → 只跑边界/异常参数用例（EC-01~EC-13）                   │
 * │                                                           │
 * │  node tencent_vod_video_test.js --all                     │
 * │    → 跑全部用例（逐条顺序执行，共 60 条）                    │
 * │                                                           │
 * │  node tencent_vod_video_test.js --dry-run --case 2        │
 * │    → 只推演预期扣费，不发任何 HTTP 请求                      │
 * │                                                           │
 * │  node tencent_vod_video_test.js --force --case 2          │
 * │    → 强制重跑（绕过"已执行"保护，会重新扣费）                │
 * │                                                           │
 * │  node tencent_vod_video_test.js --resume --case 1         │
 * │    → 不重新提交，只轮询已记录的 task_id 并补写报告           │
 * │                                                           │
 * │  node tencent_vod_video_test.js --resume --case 1 \       │
 * │       --task-id task_xxx                                  │
 * │    → 指定 task_id 补轮询（状态文件丢失时用）                 │
 * │                                                           │
 * └────────────────────────────────────────────────────────────┘
 *
 * 成本控制原则：
 *   1. 每条用例只提交 1 次（除非加 --force）
 *   2. 提交成功后立刻落盘状态，轮询中断也不会重复扣费
 *   3. 用例严格按序执行，前一条轮询完成后才开始下一条
 *   4. 未指定范围时拒绝执行（必须给出 --all / --case / 等参数）
 *
 * 结果输出目录：
 *   ../模型测试结果记录文档/{model}.md（按模型追加写入）
 *
 * 进度状态文件：
 *   .tencent_vod_test_state.json（记录已执行用例，防止重复提交）
 * ============================================================
 */

'use strict';

const http = require('http');
const fs   = require('fs');
const path = require('path');

// ── 基础配置 ──────────────────────────────────────────────────
const BASE_HOST = 'localhost';
const BASE_PORT = 9006;
const API_KEY   = 'sk-bekfGkjueBjiia7A8JiKlfSJX3HKI7fMe8LupUr4Xn0P5xda';

const POLL_INTERVAL_MS = 15_000;  // 轮询间隔 15s
const POLL_TIMEOUT_MS  = 600_000; // 最长等待 10 分钟

const SCRIPT_DIR  = __dirname;
const RESULTS_DIR = path.join(SCRIPT_DIR, '..', '模型测试结果记录文档');
const STATE_FILE  = path.join(SCRIPT_DIR, '.tencent_vod_test_state.json');

// 素材图片（OSS 直链）
const IMG_LANDSCAPE = 'https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg';
const IMG_PERSON    = 'https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg';
const IMG_PORTRAIT  = 'https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809801856-14qkezq2.jpeg';

// ── 工具函数 ──────────────────────────────────────────────────
const Q = (exprVal) => Math.round(exprVal * 0.5); // exprVal → quota (÷1e6 × 500000)
const fmt = (n) => typeof n === 'number' ? n.toLocaleString('en-US') : String(n);
const now  = () => new Date().toLocaleString('zh-CN', { hour12: false });

// ── 用例定义 ──────────────────────────────────────────────────
// seq: 用例序号（1-based）
// id:  便于引用的字符串标识
// model: 调用的模型名
// name: 用例说明
// body: 请求体（不含 model 字段，由 runner 自动注入）
// expectedQuota: 预期扣减 quota（Math.round(exprVal × 0.5)）
// expectHttp: 预期 HTTP 状态码（默认 200）
// tag: 'standard' | 'edge'
// notes: 补充说明
const CASES = [
  // ─── 一、Kling 系列 ────────────────────────────────────────────────────────
  {
    seq: 1, id: 'TC-01', model: 'kling-video-1.6', tag: 'standard',
    name: '不传 duration/size（默认 720P × 5s）',
    body: { prompt: '夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格' },
    expectedQuota: Q(273_975), expectHttp: 200,
    notes: '适配器默认 720P；API 默认时长 5s。表达式 54795×5=273975',
  },
  {
    seq: 2, id: 'TC-02', model: 'kling-video-1.6', tag: 'standard',
    name: 'duration=10 + size=1080P',
    body: { prompt: '夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格', duration: 10, size: '1080P' },
    expectedQuota: Q(958_900), expectHttp: 200,
    notes: '95890×10=958900',
  },
  {
    seq: 3, id: 'TC-03', model: 'kling-video-2.5', tag: 'standard',
    name: '不传 duration/size（默认 720P × 5s）',
    body: { prompt: '秋天的银杏大道，金黄叶片漫天飞舞，阳光透过树梢，电影感' },
    expectedQuota: Q(205_480), expectHttp: 200,
    notes: '41096×5=205480',
  },
  {
    seq: 4, id: 'TC-04', model: 'kling-video-2.5', tag: 'standard',
    name: 'duration=10 + size=1080P',
    body: { prompt: '秋天的银杏大道，金黄叶片漫天飞舞，阳光透过树梢，电影感', duration: 10, size: '1080P' },
    expectedQuota: Q(684_930), expectHttp: 200,
    notes: '68493×10=684930',
  },
  {
    seq: 5, id: 'TC-05', model: 'kling-video-2.6', tag: 'standard',
    name: '无声 — 720P × 5s',
    body: { prompt: '城市夜景俯瞰，霓虹灯流光，车流如河，航拍视角' },
    expectedQuota: Q(205_480), expectHttp: 200,
    notes: 'silent-720P 41096×5=205480',
  },
  {
    seq: 6, id: 'TC-06', model: 'kling-video-2.6', tag: 'standard',
    name: '有声 — 1080P × 5s（audio_generation=Enabled）',
    body: {
      prompt: '城市夜景俯瞰，霓虹灯流光，车流如河，航拍视角',
      duration: 5, size: '1080P',
      metadata: { audio_generation: 'Enabled' },
    },
    expectedQuota: Q(684_930), expectHttp: 200,
    notes: 'audio-1080P 136986×5=684930',
  },
  {
    seq: 7, id: 'TC-07', model: 'kling-video-o1', tag: 'standard',
    name: '无参考图 — 720P × 5s',
    body: { prompt: '水墨山水画风格，群山叠翠，云雾飘渺，鸟儿掠过' },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: 't2v-720P 82192×5=410960',
  },
  {
    seq: 8, id: 'TC-08', model: 'kling-video-o1', tag: 'standard',
    name: '有参考图 — 1080P × 10s（2 张图）',
    body: {
      prompt: '参考图中的风景延伸为动态视频，轻风拂过水面，波光粼粼',
      duration: 10, size: '1080P',
      images: [IMG_LANDSCAPE, IMG_PORTRAIT],
    },
    expectedQuota: Q(1_643_840), expectHttp: 200,
    notes: 'ref-1080P 164384×10=1643840',
  },
  {
    seq: 9, id: 'TC-09', model: 'kling-video-3.0', tag: 'standard',
    name: '无声 — 720P × 5s',
    body: { prompt: '雨后竹林，翠绿水珠悬挂叶尖，鸟鸣声声，空气清新' },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: 'silent-720P 82192×5=410960',
  },
  {
    seq: 10, id: 'TC-10', model: 'kling-video-3.0', tag: 'standard',
    name: '有声 — 1080P × 5s',
    body: {
      prompt: '雨后竹林，翠绿水珠悬挂叶尖，鸟鸣声声，空气清新',
      duration: 5, size: '1080P',
      metadata: { audio_generation: 'Enabled' },
    },
    expectedQuota: Q(821_920), expectHttp: 200,
    notes: 'audio-1080P 164384×5=821920',
  },
  {
    seq: 11, id: 'TC-11', model: 'kling-video-3.0-omni', tag: 'standard',
    name: '无声无参考 — 720P × 5s',
    body: { prompt: '黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感' },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: 'silent-720P 82192×5=410960',
  },
  {
    seq: 12, id: 'TC-12', model: 'kling-video-3.0-omni', tag: 'standard',
    name: '有参考图 — 1080P × 5s',
    body: {
      prompt: '参考图中的人物微笑转身，背景为温馨的咖啡馆，自然光',
      duration: 5, size: '1080P',
      images: [IMG_PERSON],
      metadata: { input_usage: 'Reference' },
    },
    expectedQuota: Q(821_920), expectHttp: 200,
    notes: 'ref-1080P 164384×5=821920',
  },
  // ─── 二、Vidu 系列 ─────────────────────────────────────────────────────────
  {
    seq: 13, id: 'TC-13', model: 'vidu-video-q2', tag: 'standard',
    name: '文生视频 — 720P × 5s',
    body: { prompt: '热带雨林中，蝴蝶在花丛间飞舞，色彩缤纷，近景微距' },
    expectedQuota: Q(219_180), expectHttp: 200,
    notes: 't2v-720P 43836×5=219180',
  },
  {
    seq: 14, id: 'TC-14', model: 'vidu-video-q2', tag: 'standard',
    name: '参考图生视频 — 1080P × 8s',
    body: {
      prompt: '参考图中的风景在晚霞中渐渐变暗，远处炊烟袅袅升起',
      duration: 8, size: '1080P',
      images: [IMG_LANDSCAPE],
      metadata: { input_usage: 'Reference' },
    },
    expectedQuota: Q(898_632), expectHttp: 200,
    notes: 'ref-1080P 112329×8=898632',
  },
  {
    seq: 15, id: 'TC-15', model: 'vidu-video-q2-pro', tag: 'standard',
    name: '首帧图生 — 720P × 5s（不传 input_usage）',
    body: {
      prompt: '人物缓缓转身，露出微笑，阳光洒在发梢',
      duration: 5, size: '720P',
      images: [IMG_PERSON],
    },
    expectedQuota: Q(239_725), expectHttp: 200,
    notes: '无 input_usage → i2v-720P 47945×5=239725',
  },
  {
    seq: 16, id: 'TC-16', model: 'vidu-video-q2-pro', tag: 'standard',
    name: '参考生 — 1080P × 8s（input_usage=Reference）',
    body: {
      prompt: '人物向镜头走来，背景为樱花盛开的公园，写实风格',
      duration: 8, size: '1080P',
      images: [IMG_PERSON],
      metadata: { input_usage: 'Reference' },
    },
    expectedQuota: Q(986_304), expectHttp: 200,
    notes: 'ref-1080P 123288×8=986304',
  },
  {
    seq: 17, id: 'TC-17', model: 'vidu-video-q2-turbo', tag: 'standard',
    name: '720P × 5s',
    body: {
      prompt: '快节奏城市街景，行人穿梭，霓虹倒影，延时摄影感',
      duration: 5, size: '720P',
      images: [IMG_LANDSCAPE],
    },
    expectedQuota: Q(171_235), expectHttp: 200,
    notes: '720P 34247×5=171235',
  },
  {
    seq: 18, id: 'TC-18', model: 'vidu-video-q2-turbo', tag: 'standard',
    name: '1080P × 8s',
    body: {
      prompt: '快节奏城市街景，行人穿梭，霓虹倒影，延时摄影感',
      duration: 8, size: '1080P',
      images: [IMG_LANDSCAPE],
    },
    expectedQuota: Q(515_072), expectHttp: 200,
    notes: '1080P 64384×8=515072',
  },
  {
    seq: 19, id: 'TC-19', model: 'vidu-video-q3', tag: 'standard',
    name: '720P × 5s',
    body: {
      prompt: '冬日雪原，驯鹿群奔跑，雪花飞舞，北极光远闪',
      duration: 5, size: '720P',
      images: [IMG_LANDSCAPE],
    },
    expectedQuota: Q(428_080), expectHttp: 200,
    notes: '720P 85616×5=428080',
  },
  {
    seq: 20, id: 'TC-20', model: 'vidu-video-q3', tag: 'standard',
    name: '1080P × 8s',
    body: {
      prompt: '冬日雪原，驯鹿群奔跑，雪花飞舞，北极光远闪',
      duration: 8, size: '1080P',
      images: [IMG_LANDSCAPE],
    },
    expectedQuota: Q(856_984), expectHttp: 200,
    notes: '1080P 107123×8=856984',
  },
  {
    seq: 21, id: 'TC-21', model: 'vidu-video-q3-pro', tag: 'standard',
    name: '720P × 5s',
    body: {
      prompt: '女子站在落日余晖中，微风拂动发丝，暖色调电影感',
      duration: 5, size: '720P',
      images: [IMG_PORTRAIT],
    },
    expectedQuota: Q(535_615), expectHttp: 200,
    notes: '720P 107123×5=535615',
  },
  {
    seq: 22, id: 'TC-22', model: 'vidu-video-q3-pro', tag: 'standard',
    name: '1080P × 8s',
    body: {
      prompt: '女子站在落日余晖中，微风拂动发丝，暖色调电影感',
      duration: 8, size: '1080P',
      images: [IMG_PORTRAIT],
    },
    expectedQuota: Q(1_027_944), expectHttp: 200,
    notes: '1080P 128493×8=1027944',
  },
  {
    seq: 23, id: 'TC-23', model: 'vidu-video-q3-turbo', tag: 'standard',
    name: '720P × 5s',
    body: {
      prompt: '古堡在暮色中轮廓渐现，蝙蝠掠过天际，神秘哥特风',
      duration: 5, size: '720P',
      images: [IMG_LANDSCAPE],
    },
    expectedQuota: Q(256_850), expectHttp: 200,
    notes: '720P 51370×5=256850',
  },
  {
    seq: 24, id: 'TC-24', model: 'vidu-video-q3-turbo', tag: 'standard',
    name: '1080P × 8s',
    body: {
      prompt: '古堡在暮色中轮廓渐现，蝙蝠掠过天际，神秘哥特风',
      duration: 8, size: '1080P',
      images: [IMG_LANDSCAPE],
    },
    expectedQuota: Q(480_000), expectHttp: 200,
    notes: '1080P 60000×8=480000',
  },
  // ─── 三、Hailuo 系列 ───────────────────────────────────────────────────────
  {
    seq: 25, id: 'TC-25', model: 'hailuo-video-02', tag: 'standard',
    name: '默认（768P × 6s）',
    body: { prompt: '清晨薄雾中的古镇水乡，乌篷船缓缓划过，倒影在水面荡漾' },
    expectedQuota: Q(271_230), expectHttp: 200,
    notes: '768P 45205×6=271230',
  },
  {
    seq: 26, id: 'TC-26', model: 'hailuo-video-02', tag: 'standard',
    name: 'duration=10 + size=1080P',
    body: { prompt: '清晨薄雾中的古镇水乡，乌篷船缓缓划过，倒影在水面荡漾', duration: 10, size: '1080P' },
    expectedQuota: Q(794_520), expectHttp: 200,
    notes: '1080P 79452×10=794520',
  },
  {
    seq: 27, id: 'TC-27', model: 'hailuo-video-2.3-fast', tag: 'standard',
    name: '默认（768P × 6s）',
    body: { prompt: '樱花满开的公园小径，花瓣随风飘落，慢动作，浪漫唯美' },
    expectedQuota: Q(184_932), expectHttp: 200,
    notes: '768P 30822×6=184932',
  },
  {
    seq: 28, id: 'TC-28', model: 'hailuo-video-2.3-fast', tag: 'standard',
    name: 'duration=10 + size=1080P',
    body: { prompt: '樱花满开的公园小径，花瓣随风飘落，慢动作，浪漫唯美', duration: 10, size: '1080P' },
    expectedQuota: Q(527_400), expectHttp: 200,
    notes: '1080P 52740×10=527400',
  },
  {
    seq: 29, id: 'TC-29', model: 'hailuo-video-h3', tag: 'standard',
    name: '默认（768P × 6s，占位价格）',
    body: { prompt: '深秋枫林，一片片红叶在晨风中轻轻飘落，光线穿透树冠，静谧' },
    expectedQuota: Q(271_230), expectHttp: 200,
    notes: '⚠️ H3 价格未公布，使用 02 同档占位。45205×6=271230',
  },
  {
    seq: 30, id: 'TC-30', model: 'hailuo-video-h3', tag: 'standard',
    name: 'duration=10 + size=1080P（占位价格）',
    body: { prompt: '深秋枫林，一片片红叶在晨风中轻轻飘落，光线穿透树冠，静谧', duration: 10, size: '1080P' },
    expectedQuota: Q(794_520), expectHttp: 200,
    notes: '⚠️ H3 价格未公布，使用 02 同档占位。79452×10=794520',
  },
  // ─── 四、Google Veo 系列 ───────────────────────────────────────────────────
  {
    seq: 31, id: 'TC-31', model: 'veo-video-3.1', tag: 'standard',
    name: '无声（默认）— d=8 固定',
    body: { prompt: '极光在冰岛上空绽放，湖面倒映，绿色光带舞动' },
    expectedQuota: Q(1_643_832), expectHttp: 200,
    notes: 'silent-720P 205479×8=1643832（GV duration 固定 8s）',
  },
  {
    seq: 32, id: 'TC-32', model: 'veo-video-3.1', tag: 'standard',
    name: '有声 — audio_generation=Enabled',
    body: {
      prompt: '极光在冰岛上空绽放，湖面倒映，绿色光带舞动',
      metadata: { audio_generation: 'Enabled' },
    },
    expectedQuota: Q(3_287_672), expectHttp: 200,
    notes: 'audio-720P 410959×8=3287672',
  },
  {
    seq: 33, id: 'TC-33', model: 'veo-video-3.1-fast', tag: 'standard',
    name: '无声（默认）— d=8 固定',
    body: { prompt: '热带海底，珊瑚礁旁鱼群穿梭，阳光折射，蓝色清澈' },
    expectedQuota: Q(821_920), expectHttp: 200,
    notes: 'silent-720P 102740×8=821920',
  },
  {
    seq: 34, id: 'TC-34', model: 'veo-video-3.1-fast', tag: 'standard',
    name: '有声 — audio_generation=Enabled',
    body: {
      prompt: '热带海底，珊瑚礁旁鱼群穿梭，阳光折射，蓝色清澈',
      metadata: { audio_generation: 'Enabled' },
    },
    expectedQuota: Q(1_232_880), expectHttp: 200,
    notes: 'audio-720P 154110×8=1232880',
  },
  // ─── 五、Sora（OS）系列 ────────────────────────────────────────────────────
  {
    seq: 35, id: 'TC-35', model: 'sora-video-2.0', tag: 'standard',
    name: '不传 duration（API 默认 8s，snap→8）',
    body: { prompt: '未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁' },
    expectedQuota: Q(821_920), expectHttp: 200,
    notes: 'd0=0→d=8；720P 102740×8=821920',
  },
  {
    seq: 36, id: 'TC-36', model: 'sora-video-2.0', tag: 'standard',
    name: 'duration=5（≤6 → snap 到 4s）',
    body: { prompt: '未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁', duration: 5 },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: 'd0=5≤6→d=4；720P 102740×4=410960',
  },
  {
    seq: 37, id: 'TC-37', model: 'sora-video-2.0', tag: 'standard',
    name: 'duration=12（>10 → snap 到 12s）',
    body: { prompt: '未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁', duration: 12 },
    expectedQuota: Q(1_232_880), expectHttp: 200,
    notes: 'd0=12>10→d=12；720P 102740×12=1232880',
  },
  // ─── 六、Hunyuan / Mingmou ────────────────────────────────────────────────
  {
    seq: 38, id: 'TC-38', model: 'hunyuan-video-1.5', tag: 'standard',
    name: '720P × 5s',
    body: { prompt: '禅意庭院，枯山水中落叶旋转，晨光斑驳' },
    expectedQuota: Q(205_480), expectHttp: 200,
    notes: '720P 41096×5=205480',
  },
  {
    seq: 39, id: 'TC-39', model: 'hunyuan-video-1.5', tag: 'standard',
    name: '1080P × 10s',
    body: { prompt: '禅意庭院，枯山水中落叶旋转，晨光斑驳', duration: 10, size: '1080P' },
    expectedQuota: Q(684_930), expectHttp: 200,
    notes: '1080P 68493×10=684930',
  },
  {
    seq: 40, id: 'TC-40', model: 'mingmou-video-1.0', tag: 'standard',
    name: '720P × 5s',
    body: { prompt: '微缩城市模型，汽车穿行在精致街道，移轴摄影效果' },
    expectedQuota: Q(205_480), expectHttp: 200,
    notes: '720P 41096×5=205480（同 Hunyuan 定价）',
  },
  {
    seq: 41, id: 'TC-41', model: 'mingmou-video-1.0', tag: 'standard',
    name: '1080P × 10s',
    body: { prompt: '微缩城市模型，汽车穿行在精致街道，移轴摄影效果', duration: 10, size: '1080P' },
    expectedQuota: Q(684_930), expectHttp: 200,
    notes: '1080P 68493×10=684930',
  },
  // ─── 七、PixVerse 系列 ────────────────────────────────────────────────────
  {
    seq: 42, id: 'TC-42', model: 'pixverse-video-v5.6', tag: 'standard',
    name: '720p × 5s（无声）',
    body: { prompt: '节日烟花在夜空中绽放，五彩斑斓，倒映在湖面' },
    expectedQuota: Q(215_755), expectHttp: 200,
    notes: '720p 43151×5=215755',
  },
  {
    seq: 43, id: 'TC-43', model: 'pixverse-video-v5.6', tag: 'standard',
    name: '1080p × 10s（无声）',
    body: { prompt: '节日烟花在夜空中绽放，五彩斑斓，倒映在湖面', duration: 10, size: '1080p' },
    expectedQuota: Q(719_180), expectHttp: 200,
    notes: '1080p 71918×10=719180',
  },
  {
    seq: 44, id: 'TC-44', model: 'pixverse-video-v6', tag: 'standard',
    name: '无声 — 720p × 5s',
    body: { prompt: '冰川融化，蔚蓝海水与白色浮冰交织，鸟瞰航拍' },
    expectedQuota: Q(180_820), expectHttp: 200,
    notes: 'silent-720p 36164×5=180820',
  },
  {
    seq: 45, id: 'TC-45', model: 'pixverse-video-v6', tag: 'standard',
    name: '有声 — 1080p × 8s（audio_generation=Enabled）',
    body: {
      prompt: '冰川融化，蔚蓝海水与白色浮冰交织，鸟瞰航拍',
      duration: 8, size: '1080p',
      metadata: { audio_generation: 'Enabled' },
    },
    expectedQuota: Q(739_728), expectHttp: 200,
    notes: 'audio-1080p 92466×8=739728',
  },
  {
    seq: 46, id: 'TC-46', model: 'pixverse-video-c1', tag: 'standard',
    name: '无声 — 720p × 5s',
    body: { prompt: '迷雾森林中的精灵，荧光粒子飘落，梦幻奇幻风格' },
    expectedQuota: Q(200_685), expectHttp: 200,
    notes: 'silent-720p 40137×5=200685',
  },
  {
    seq: 47, id: 'TC-47', model: 'pixverse-video-c1', tag: 'standard',
    name: '有声 — 1080p × 8s',
    body: {
      prompt: '迷雾森林中的精灵，荧光粒子飘落，梦幻奇幻风格',
      duration: 8, size: '1080p',
      metadata: { audio_generation: 'Enabled' },
    },
    expectedQuota: Q(771_504), expectHttp: 200,
    notes: 'audio-1080p 96438×8=771504',
  },
  // ─── 八、边界与异常参数覆盖（3 个模型） ───────────────────────────────────
  // ── kling-video-3.0-omni ──────────────────────────────────────────────────
  {
    seq: 48, id: 'EC-01', model: 'kling-video-3.0-omni', tag: 'edge',
    name: 'duration=-1（负数 → clamp 到 5）',
    body: { prompt: '黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感', duration: -1, size: '720P' },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: '表达式 d0=-1≤0→d=5；silent-720P 82192×5=410960。验证 clamp 后仍按 5s 计费。',
  },
  {
    seq: 49, id: 'EC-02', model: 'kling-video-3.0-omni', tag: 'edge',
    name: 'size="random_xyz"（无效分辨率 → fallback 720P）',
    body: { prompt: '黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感', duration: 5, size: 'random_xyz' },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: '无效 size 命中不了任何 tier，fallback silent-720P 82192×5=410960。',
  },
  {
    seq: 50, id: 'EC-03', model: 'kling-video-3.0-omni', tag: 'edge',
    name: 'audio_generation="yes"（错误值，非 "Enabled" → 按无声计）',
    body: {
      prompt: '黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感',
      duration: 5, size: '720P',
      metadata: { audio_generation: 'yes' },
    },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: '表达式检查 == "Enabled"，"yes" 不匹配 → audio=false → silent-720P 82192×5=410960。',
  },
  {
    seq: 51, id: 'EC-04', model: 'kling-video-3.0-omni', tag: 'edge',
    name: 'images=[]（空数组 → 等同无参考图）',
    body: {
      prompt: '黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感',
      duration: 5, size: '720P',
      images: [],
    },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: '表达式 images.# = 0 → ref=false → silent-720P 82192×5=410960。',
  },
  {
    seq: 52, id: 'EC-05', model: 'kling-video-3.0-omni', tag: 'edge',
    name: 'prompt 为空字符串（应被服务端拒绝）',
    body: { prompt: '', duration: 5, size: '720P' },
    expectedQuota: 0, expectHttp: 400,
    notes: '预期 400。不应扣费；验证空 prompt 校验是否生效。',
  },
  // ── vidu-video-q2-pro ─────────────────────────────────────────────────────
  {
    seq: 53, id: 'EC-06', model: 'vidu-video-q2-pro', tag: 'edge',
    name: '有图但不传 input_usage → 走首帧(i2v)价格',
    body: {
      prompt: '人物缓缓转身，露出微笑，阳光洒在发梢',
      duration: 5, size: '720P',
      images: [IMG_PERSON],
    },
    expectedQuota: Q(239_725), expectHttp: 200,
    notes: '不传 input_usage → ref=false → i2v-720P 47945×5=239725。与 TC-15 相同预期，验证默认行为。',
  },
  {
    seq: 54, id: 'EC-07', model: 'vidu-video-q2-pro', tag: 'edge',
    name: 'input_usage=Reference 但不传 images → 参考生价（表达式仍按 ref 计）',
    body: {
      prompt: '人物缓缓转身，露出微笑，阳光洒在发梢',
      duration: 5, size: '720P',
      metadata: { input_usage: 'Reference' },
    },
    expectedQuota: Q(239_725), expectHttp: 200,
    notes: '表达式 ref=true(input_usage==Reference) → ref-720P 47945×5=239725。注意：无图时腾讯 API 实际行为待观察。',
  },
  {
    seq: 55, id: 'EC-08', model: 'vidu-video-q2-pro', tag: 'edge',
    name: 'duration=0（→ 按默认 5s 计费）',
    body: {
      prompt: '人物缓缓转身，露出微笑，阳光洒在发梢',
      duration: 0, size: '720P',
      images: [IMG_PERSON],
    },
    expectedQuota: Q(239_725), expectHttp: 200,
    notes: 'd0=0≤0→d=5；i2v-720P 47945×5=239725。验证 duration=0 的 clamp 行为。',
  },
  {
    seq: 56, id: 'EC-09', model: 'vidu-video-q2-pro', tag: 'edge',
    name: 'size="2K"（i2v-2K 档）',
    body: {
      prompt: '人物缓缓转身，露出微笑，阳光洒在发梢',
      duration: 5, size: '2K',
      images: [IMG_PERSON],
    },
    expectedQuota: Q(684_930), expectHttp: 200,
    notes: 'i2v-2K 136986×5=684930。验证高分辨率档计费。',
  },
  // ── sora-video-2.0（duration snap {4, 8, 12} 边界） ───────────────────────
  {
    seq: 57, id: 'EC-10', model: 'sora-video-2.0', tag: 'edge',
    name: 'duration=6（≤6 → snap 到 4s）',
    body: { prompt: '未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁', duration: 6 },
    expectedQuota: Q(410_960), expectHttp: 200,
    notes: 'd0=6≤6→d=4；720P 102740×4=410960。snap 下界验证。',
  },
  {
    seq: 58, id: 'EC-11', model: 'sora-video-2.0', tag: 'edge',
    name: 'duration=7（6< d ≤10 → snap 到 8s）',
    body: { prompt: '未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁', duration: 7 },
    expectedQuota: Q(821_920), expectHttp: 200,
    notes: 'd0=7→d=8；720P 102740×8=821920。snap 中间档验证。',
  },
  {
    seq: 59, id: 'EC-12', model: 'sora-video-2.0', tag: 'edge',
    name: 'duration=11（>10 → snap 到 12s）',
    body: { prompt: '未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁', duration: 11 },
    expectedQuota: Q(1_232_880), expectHttp: 200,
    notes: 'd0=11>10→d=12；720P 102740×12=1232880。snap 上界验证。',
  },
  {
    seq: 60, id: 'EC-13', model: 'sora-video-2.0', tag: 'edge',
    name: 'size="1080P"（表达式按 1080P 计，适配器可能实际仍发 720P）',
    body: { prompt: '未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁', size: '1080P' },
    expectedQuota: Q(1_232_880), expectHttp: 200,
    notes: 'd=8(默认)；1080P 154110×8=1232880。注意：适配器 OS 仅支持 720P，实际请求分辨率可能被降级，但计费按表达式走 1080P 档。观察实际扣费与此预期是否一致。',
  },
];

// ── 状态管理（防止重复提交） ────────────────────────────────────
function loadState() {
  try {
    return JSON.parse(fs.readFileSync(STATE_FILE, 'utf8'));
  } catch {
    return { executed: {} };
  }
}
function saveState(state) {
  fs.writeFileSync(STATE_FILE, JSON.stringify(state, null, 2), 'utf8');
}

// ── HTTP 工具（直连 localhost，不走代理） ────────────────────────
function httpRequest(method, urlPath, body, timeoutMs = 120_000) {
  return new Promise((resolve, reject) => {
    const payload = body ? JSON.stringify(body) : null;
    const opts = {
      hostname: BASE_HOST,
      port: BASE_PORT,
      path: urlPath,
      method,
      headers: {
        Authorization: `Bearer ${API_KEY}`,
        'Content-Type': 'application/json',
        ...(payload ? { 'Content-Length': Buffer.byteLength(payload) } : {}),
      },
      timeout: timeoutMs,
    };
    const t0 = Date.now();
    const req = http.request(opts, (res) => {
      let raw = '';
      res.on('data', (c) => { raw += c; });
      res.on('end', () => {
        let json = {};
        try { json = JSON.parse(raw || '{}'); } catch { json = { _raw: raw }; }
        resolve({ status: res.statusCode, json, elapsedMs: Date.now() - t0 });
      });
    });
    req.on('timeout', () => { req.destroy(new Error(`HTTP timeout after ${timeoutMs}ms`)); });
    req.on('error', reject);
    if (payload) req.write(payload);
    req.end();
  });
}

// ── 从查询响应中提取任务数据 ─────────────────────────────────────
// 实际接口格式：{ code: "success", data: { status: "SUCCESS", result_url, ... } }
// status 取值：QUEUED / SUBMITTED / IN_PROGRESS / SUCCESS / FAILURE
function extractTaskPayload(json) {
  if (!json || typeof json !== 'object') return {};
  // 兼容：有 data 包裹时取 data；否则退回顶层（兼容其他渠道格式）
  return (json.data && typeof json.data === 'object') ? json.data : json;
}

function extractTaskStatus(json) {
  const payload = extractTaskPayload(json);
  return String(payload.status || payload.Status || '').toUpperCase();
}

const DONE_STATUSES = new Set(['SUCCESS', 'FAILURE', 'FAILED', 'CANCELLED', 'SUCCEEDED']);

// ── 轮询任务状态 ────────────────────────────────────────────────
async function pollTask(taskId) {
  const deadline = Date.now() + POLL_TIMEOUT_MS;
  let polls = 0;
  let last = {};
  const t0 = Date.now();
  while (Date.now() < deadline) {
    polls++;
    const r = await httpRequest('GET', `/v1/video/generations/${taskId}`);
    last = r.json;
    const status = extractTaskStatus(last);
    const progress = extractTaskPayload(last).progress || '';
    process.stdout.write(`\r  轮询 #${polls}：${status || '(空)'}${progress ? ` ${progress}` : ''}  `);
    if (DONE_STATUSES.has(status)) {
      process.stdout.write('\n');
      return { finalStatus: status, polls, elapsedMs: Date.now() - t0, last };
    }
    await sleep(POLL_INTERVAL_MS);
  }
  process.stdout.write('\n');
  return { finalStatus: 'timeout', polls, elapsedMs: Date.now() - t0, last };
}

function sleep(ms) {
  return new Promise((r) => setTimeout(r, ms));
}

// ── 查询 Token 额度 ─────────────────────────────────────────────
async function getAvailableQuota() {
  try {
    const r = await httpRequest('GET', '/api/usage/token/');
    return (r.json.data || {}).total_available ?? null;
  } catch {
    return null;
  }
}

// ── 查询最近一条消费日志 ─────────────────────────────────────────
async function getLatestLog() {
  try {
    const r = await httpRequest('GET', '/api/log/token?p=1&page_size=1');
    const list = r.json.data || [];
    return list[0] || null;
  } catch {
    return null;
  }
}

// ── 写入结果文件（按模型追加） ────────────────────────────────────
function writeResult(model, text) {
  fs.mkdirSync(RESULTS_DIR, { recursive: true });
  const file = path.join(RESULTS_DIR, `${model}.md`);
  const isNew = !fs.existsSync(file);
  const fd = fs.openSync(file, 'a');
  if (isNew) {
    fs.writeSync(fd, `# ${model} 测试报告\n\n`);
    fs.writeSync(fd, `> 本文件由 \`tencent_vod_video_test.js\` 自动追加生成，禁止手动修改序号。\n\n`);
  }
  fs.writeSync(fd, text);
  if (!text.endsWith('\n')) fs.writeSync(fd, '\n');
  fs.writeSync(fd, '\n---\n\n');
  fs.closeSync(fd);
  return file;
}

// ── 价格变量提取辅助 ─────────────────────────────────────────────

// 根据模型名和请求 body 提取影响价格的变量，用于报告中的核对表
function extractBillingVars(model, body) {
  const meta         = (body.metadata && typeof body.metadata === 'object') ? body.metadata : {};
  const durationSent = body.duration != null ? body.duration : null;
  const sizeSent     = body.size || meta.resolution || null;
  const audio        = meta.audio_generation || null;
  const imgCount     = Array.isArray(body.images) ? body.images.length : 0;
  const inputUsage   = meta.input_usage || null;

  // 推导有效计费时长（与 Go getBillingDuration 对应）
  let billingDur;
  const family = model.split('-')[0].toLowerCase(); // kling / vidu / hailuo / veo / sora / hunyuan / mingmou / pixverse
  if (family === 'hunyuan' || family === 'mingmou') {
    billingDur = null; // 不发 Duration
  } else if (family === 'hailuo') {
    // snap {6,10}
    billingDur = (durationSent == null || durationSent <= 8) ? 6 : 10;
  } else if (family === 'veo') {
    billingDur = 8; // always 8
  } else if (family === 'sora') {
    // snap {4,8,12}
    if (durationSent == null || durationSent <= 0) {
      billingDur = 8;
    } else {
      const snaps = [4, 8, 12];
      billingDur = snaps.reduce((a, b) => Math.abs(b - durationSent) < Math.abs(a - durationSent) ? b : a);
    }
  } else {
    // Kling / Vidu / PixVerse — default 5
    billingDur = (durationSent == null || durationSent <= 0) ? 5 : durationSent;
  }

  return { durationSent, billingDur, sizeSent, audio, imgCount, inputUsage };
}

// 从最终轮询响应中提取实际视频元数据（时长、分辨率）
function extractVideoMeta(pollLastJson) {
  try {
    const resp = pollLastJson?.data?.data?.Response;
    if (!resp) return null;
    const videoTask = resp.AigcVideoTask;
    if (!videoTask?.Output?.FileInfos?.length) return null;
    // 取 UsageType 为空（视频正文）的条目
    const fi = videoTask.Output.FileInfos.find(f => !f.UsageType) || videoTask.Output.FileInfos[0];
    const meta = fi?.MetaData;
    if (!meta) return null;
    return {
      duration: meta.Duration ?? null,       // 实际时长（秒，8位小数精度）
      width:    meta.Width   ?? null,
      height:   meta.Height  ?? null,
    };
  } catch {
    return null;
  }
}

// ── 构建报告文本 ─────────────────────────────────────────────────
function buildReport(tc, { dryRun, submitResult, taskId, pollResult, billing } = {}) {
  const ts = now();
  const lines = [
    `## ${tc.id}（seq=${tc.seq}）: ${tc.name}`,
    `> 执行时间：${ts}  |  模型：\`${tc.model}\`  |  标签：${tc.tag}`,
    '',
  ];
  if (tc.notes) lines.push(`> 💡 ${tc.notes}`, '');

  const bodyWithModel = { model: tc.model, ...tc.body };
  lines.push('### 调用参数', '```json', JSON.stringify(bodyWithModel, null, 2), '```', '');

  // 价格变量核对
  const bv = extractBillingVars(tc.model, tc.body);
  lines.push(
    '### 价格变量核对',
    '| 变量 | 请求值 | 有效计费值 | 说明 |',
    '|---|---|---|---|',
    `| duration（时长/s） | ${bv.durationSent != null ? bv.durationSent : '未传'} | ${bv.billingDur != null ? `**${bv.billingDur}**` : 'N/A（不参与计费）'} | ${bv.billingDur == null ? '模型不发送 Duration' : bv.durationSent == null ? '未传，使用表达式默认值' : '已传，直接使用'} |`,
    `| size（分辨率） | ${bv.sizeSent || '未传'} | — | 未传时适配器默认 720P |`,
    `| audio_generation | ${bv.audio || '未传'} | — | 未传视为无声版 |`,
    `| images（参考图数） | ${bv.imgCount > 0 ? `${bv.imgCount} 张` : '无'} | — | ${bv.imgCount > 0 && bv.inputUsage ? `input_usage=${bv.inputUsage}` : bv.imgCount > 0 ? '参考图生视频' : '文生视频'} |`,
    '',
  );

  lines.push(
    '### 预期扣费',
    '| 项目 | 值 |',
    '|---|---|',
    `| 预期 HTTP 状态 | ${tc.expectHttp} |`,
    `| 预期 Quota | **${fmt(tc.expectedQuota)}** |`,
    `| 预期 USD | $${(tc.expectedQuota / 500_000).toFixed(4)} |`,
    `| 预期 RMB | ¥${(tc.expectedQuota / 500_000 * 7.3).toFixed(4)} |`,
    '',
  );

  if (dryRun) {
    lines.push('> ⚙️ dry-run 模式：未发送任何 HTTP 请求', '');
    return lines.join('\n');
  }

  // 提交结果
  const httpOk = submitResult.status === tc.expectHttp ? '✅' : '❌';
  lines.push(
    '### 提交结果',
    '| 项目 | 值 |',
    '|---|---|',
    `| HTTP 状态 | ${submitResult.status}（预期 ${tc.expectHttp}）${httpOk} |`,
    `| 耗时 | ${(submitResult.elapsedMs / 1000).toFixed(2)}s |`,
    `| task_id | \`${taskId || '-'}\` |`,
    '',
    '```json',
    JSON.stringify(submitResult.json, null, 2),
    '```',
    '',
  );

  // 扣费分析
  const match = billing.actualDeduct === tc.expectedQuota
    ? (tc.expectHttp === 200 ? '✅ 符合' : '✅（预期不扣费）')
    : '❌ 不符';
  const preRMB  = typeof billing.actualDeduct === 'number'
    ? `¥${(billing.actualDeduct / 500_000 * 7.3).toFixed(4)}` : '-';
  lines.push(
    '### 扣费分析（提交时）',
    '| 项目 | Quota | RMB |',
    '|---|---|---|',
    `| 提交前可用 Quota | ${fmt(billing.quotaBefore)} | — |`,
    `| 提交后可用 Quota | ${fmt(billing.quotaAfter)} | — |`,
    `| **预扣金额** | **${fmt(billing.actualDeduct)}** | **${preRMB}** |`,
    `| 预期扣减 | ${fmt(tc.expectedQuota)} | ¥${(tc.expectedQuota / 500_000 * 7.3).toFixed(4)} |`,
    `| 预扣是否符合 | ${match} | — |`,
    `| 消费日志 Quota | ${fmt(billing.logQuota)} | — |`,
    `| request_id | ${billing.requestId || '-'} | — |`,
    '',
  );

  // 轮询结果
  if (pollResult) {
    const payload    = extractTaskPayload(pollResult.last);
    const videoUrl   = payload.result_url || payload.url || (payload.metadata || {}).url || '';
    const failReason = payload.fail_reason || '';
    const videoMeta  = extractVideoMeta(pollResult.last);
    const finalAmt   = pollResult.last?.data?.amount;   // 平台记录的最终消费金额（RMB）

    lines.push(
      '### 任务轮询',
      '| 项目 | 值 |',
      '|---|---|',
      `| 最终状态 | **${pollResult.finalStatus}** |`,
      `| 轮询次数 | ${pollResult.polls} |`,
      `| 完成耗时 | ${(pollResult.elapsedMs / 1000).toFixed(1)}s |`,
      `| progress | ${payload.progress || '-'} |`,
      '',
    );
    if (videoUrl) lines.push(`**视频 URL**: ${videoUrl}`, '');
    if (failReason) lines.push(`**失败原因**: ${failReason}`, '');

    // 实际输出元数据（时长 + 分辨率）
    if (videoMeta) {
      const actualDur    = videoMeta.duration != null ? `${videoMeta.duration}s` : '-';
      const resolution   = (videoMeta.width && videoMeta.height) ? `${videoMeta.width}×${videoMeta.height}` : '-';
      const durDiff      = (typeof billing.actualDeduct === 'number' && bv.billingDur != null && videoMeta.duration != null)
        ? `${((videoMeta.duration - bv.billingDur)).toFixed(3)}s`
        : '-';
      lines.push(
        '#### 实际输出元数据',
        '| 项目 | 值 |',
        '|---|---|',
        `| 实际视频时长 | **${actualDur}** |`,
        `| 计费参考时长 | ${bv.billingDur != null ? `${bv.billingDur}s` : 'N/A'} |`,
        `| 时长差值 | ${durDiff} |`,
        `| 输出分辨率 | ${resolution} |`,
        '',
      );
    }

    // 最终结算对比
    if (typeof billing.actualDeduct === 'number' || finalAmt != null) {
      const preQuota  = billing.actualDeduct;
      const preRMBNum = typeof preQuota === 'number' ? (preQuota / 500_000 * 7.3) : null;
      const finalAmtNum = finalAmt != null ? parseFloat(finalAmt) : null;
      // 推算最终 Quota（若 amount 为 RMB）
      const finalQuota = finalAmtNum != null ? Math.round(finalAmtNum / 7.3 * 500_000) : null;
      const diff = (preRMBNum != null && finalAmtNum != null)
        ? (finalAmtNum - preRMBNum >= 0 ? `+¥${(finalAmtNum - preRMBNum).toFixed(4)}（补扣）` : `-¥${(preRMBNum - finalAmtNum).toFixed(4)}（退还）`)
        : '-';

      lines.push(
        '#### 最终结算（多退少补）',
        '| 项目 | Quota | RMB |',
        '|---|---|---|',
        `| 预扣金额 | ${fmt(preQuota)} | ${preRMBNum != null ? `¥${preRMBNum.toFixed(4)}` : '-'} |`,
        `| 平台记录最终消费 | ${finalQuota != null ? fmt(finalQuota) : '-'} | ${finalAmtNum != null ? `**¥${finalAmtNum.toFixed(6)}**` : '-'} |`,
        `| 差额 | — | ${diff} |`,
        '',
      );
    }

    // 最终一次查询接口的完整响应（便于核对上游字段）
    lines.push(
      '#### 最终查询响应（全量 JSON）',
      '```json',
      JSON.stringify(pollResult.last ?? {}, null, 2),
      '```',
      '',
    );
  }

  return lines.join('\n');
}

function pickTaskId(json) {
  if (!json || typeof json !== 'object') return null;
  return json.task_id || json.id
    || (json.data && (json.data.task_id || json.data.id))
    || null;
}

// ── 仅轮询已有任务并写报告（不重新提交，不扣费） ─────────────────
async function resumeCase(tc, taskId, savedBilling) {
  console.log(`\n${'='.repeat(60)}`);
  console.log(`▶ [resume] ${tc.id}: ${tc.name}`);
  console.log(`  模型: ${tc.model}  |  task_id: ${taskId}`);

  const billing = savedBilling || {
    quotaBefore: '-',
    quotaAfter: '-',
    actualDeduct: null,
    logQuota: null,
    requestId: null,
  };

  console.log(`  轮询任务 ${taskId}（每 ${POLL_INTERVAL_MS / 1000}s 一次，最长 ${POLL_TIMEOUT_MS / 60000} 分钟）`);
  const pollResult = await pollTask(taskId);
  console.log(`  任务完成：${pollResult.finalStatus}，耗时 ${(pollResult.elapsedMs / 1000).toFixed(1)}s`);

  // 用轮询响应构造一个伪 submitResult（resume 时不再有真实提交响应）
  const submitResult = {
    status: 200,
    elapsedMs: 0,
    json: { task_id: taskId, _note: 'resume：未重新提交，仅轮询已有任务' },
  };

  const report = buildReport(tc, { dryRun: false, submitResult, taskId, pollResult, billing });
  const file = writeResult(tc.model, report);
  console.log(`  报告 → ${path.relative(process.cwd(), file)}`);

  const state = loadState();
  state.executed = state.executed || {};
  state.executed[tc.id] = {
    ...(state.executed[tc.id] || {}),
    at: now(),
    taskId,
    expectedQuota: tc.expectedQuota,
    actualDeduct: billing.actualDeduct,
    pollStatus: pollResult.finalStatus,
  };
  saveState(state);
}

// ── 执行单条用例 ─────────────────────────────────────────────────
async function runCase(tc, { dryRun, force, resume, taskIdOverride }) {
  const state = loadState();
  const executed = state.executed || {};

  console.log(`\n${'='.repeat(60)}`);
  console.log(`▶ [${tc.seq}/${CASES.length}] ${tc.id}: ${tc.name}`);
  console.log(`  模型: ${tc.model}  |  预期 Quota: ${fmt(tc.expectedQuota)}  |  预期 HTTP: ${tc.expectHttp}`);

  // dry-run
  if (dryRun) {
    const report = buildReport(tc, { dryRun: true });
    const file = writeResult(tc.model, report);
    console.log(`  [dry-run] 推演完成 → ${path.relative(process.cwd(), file)}`);
    return;
  }

  // resume：只轮询已有任务，绝不重新 POST
  if (resume) {
    const taskId = taskIdOverride || executed[tc.id]?.taskId;
    if (!taskId) {
      console.error(`  ❌ resume 失败：找不到 ${tc.id} 的 task_id。请加 --task-id xxx`);
      return;
    }
    await resumeCase(tc, taskId, executed[tc.id] ? {
      quotaBefore: executed[tc.id].quotaBefore ?? '-',
      quotaAfter: executed[tc.id].quotaAfter ?? '-',
      actualDeduct: executed[tc.id].actualDeduct ?? null,
      logQuota: executed[tc.id].logQuota ?? null,
      requestId: executed[tc.id].requestId ?? null,
    } : null);
    return;
  }

  // 检查是否已执行
  if (executed[tc.id] && !force) {
    console.log(`  ⚠️  已执行过（${executed[tc.id].at}），跳过。`);
    console.log(`     task_id=${executed[tc.id].taskId || '-'}`);
    console.log(`     若只需补写轮询结果：加 --resume`);
    console.log(`     若需重新提交扣费：加 --force`);
    return;
  }

  // 提交前查额度
  process.stdout.write('  查询提交前额度...');
  const quotaBefore = await getAvailableQuota();
  console.log(` ${fmt(quotaBefore)}`);

  // 提交任务
  process.stdout.write('  提交任务...');
  const bodyWithModel = { model: tc.model, ...tc.body };
  const submitResult = await httpRequest('POST', '/v1/video/generations', bodyWithModel);
  console.log(` HTTP ${submitResult.status}，耗时 ${(submitResult.elapsedMs / 1000).toFixed(2)}s`);

  const taskId = pickTaskId(submitResult.json);

  // 提交后查额度
  await sleep(1500); // 等 1.5s 让扣费落账
  process.stdout.write('  查询提交后额度...');
  const quotaAfter = await getAvailableQuota();
  console.log(` ${fmt(quotaAfter)}`);

  const actualDeduct = quotaBefore != null && quotaAfter != null
    ? quotaBefore - quotaAfter : null;
  console.log(`  实际扣减: ${fmt(actualDeduct)}  |  预期: ${fmt(tc.expectedQuota)}`);

  // 查消费日志
  const logEntry = await getLatestLog();
  const logQuota = logEntry?.quota ?? null;
  const requestId = logEntry?.request_id ?? null;

  const billing = { quotaBefore, quotaAfter, actualDeduct, logQuota, requestId };

  // ★ 提交成功后立刻落盘，避免轮询中 Ctrl+C 导致下次重复扣费
  if (submitResult.status === 200 && taskId) {
    state.executed = state.executed || {};
    state.executed[tc.id] = {
      at: now(),
      taskId,
      expectedQuota: tc.expectedQuota,
      actualDeduct,
      quotaBefore,
      quotaAfter,
      logQuota,
      requestId,
      pollStatus: 'submitted',
    };
    saveState(state);
    console.log(`  已记录提交状态（防重复扣费）`);
  }

  // 轮询（仅在提交成功且有 task_id 时）
  let pollResult = null;
  if (submitResult.status === 200 && taskId) {
    console.log(`  轮询任务 ${taskId}（每 ${POLL_INTERVAL_MS / 1000}s 一次，最长 ${POLL_TIMEOUT_MS / 60000} 分钟）`);
    pollResult = await pollTask(taskId);
    console.log(`  任务完成：${pollResult.finalStatus}，耗时 ${(pollResult.elapsedMs / 1000).toFixed(1)}s`);
  }

  // 写报告
  const report = buildReport(tc, { dryRun: false, submitResult, taskId, pollResult, billing });
  const file = writeResult(tc.model, report);
  console.log(`  报告 → ${path.relative(process.cwd(), file)}`);

  // 更新轮询结果到状态
  if (submitResult.status === 200 && taskId) {
    state.executed[tc.id] = {
      ...state.executed[tc.id],
      pollStatus: pollResult?.finalStatus ?? 'n/a',
    };
    saveState(state);
  }
}

// ── 列出用例 ─────────────────────────────────────────────────────
function listCases(cases) {
  const state = loadState();
  const done = state.executed || {};
  console.log('\n序号  ID        模型                        预期Quota    标签      状态');
  console.log('─'.repeat(80));
  for (const tc of cases) {
    const isDone = done[tc.id] ? '✅已执行' : '⬜待执行';
    const seq    = String(tc.seq).padStart(3);
    const id     = tc.id.padEnd(9);
    const model  = tc.model.padEnd(28);
    const quota  = String(tc.expectedQuota).padStart(10);
    const tag    = tc.tag.padEnd(9);
    console.log(`${seq}  ${id} ${model} ${quota}    ${tag} ${isDone}`);
  }
  console.log(`\n共 ${cases.length} 条用例（standard: ${cases.filter(c=>c.tag==='standard').length}，edge: ${cases.filter(c=>c.tag==='edge').length}）`);
}

// ── CLI 解析 & 主函数 ─────────────────────────────────────────────
async function main() {
  const args = process.argv.slice(2);
  const has  = (flag) => args.includes(flag);
  const get  = (flag) => { const i = args.indexOf(flag); return i !== -1 ? args[i + 1] : null; };

  const dryRun = has('--dry-run');
  const force  = has('--force');
  const resume = has('--resume');
  const taskIdOverride = get('--task-id');

  // --list
  if (has('--list')) {
    listCases(CASES);
    return;
  }

  // 筛选用例
  let selected = [];

  if (has('--all')) {
    selected = CASES.slice();
  } else if (get('--case') !== null) {
    const n = parseInt(get('--case'), 10);
    if (isNaN(n) || n < 1 || n > CASES.length) {
      console.error(`❌ --case 需要 1~${CASES.length} 之间的整数`);
      process.exit(1);
    }
    selected = [CASES[n - 1]];
  } else if (get('--cases') !== null) {
    const nums = get('--cases').split(',').map(s => parseInt(s.trim(), 10));
    selected = nums.map((n) => {
      const tc = CASES[n - 1];
      if (!tc) { console.error(`❌ 用例序号 ${n} 不存在`); process.exit(1); }
      return tc;
    });
  } else if (get('--from') !== null || get('--to') !== null) {
    const from = parseInt(get('--from') || '1', 10);
    const to   = parseInt(get('--to') || String(CASES.length), 10);
    selected = CASES.slice(from - 1, to);
  } else if (get('--model') !== null) {
    const m = get('--model');
    selected = CASES.filter(tc => tc.model === m);
    if (!selected.length) {
      console.error(`❌ 未找到模型 "${m}" 的用例。可用模型：\n  ${[...new Set(CASES.map(c=>c.model))].join('\n  ')}`);
      process.exit(1);
    }
  } else if (get('--tag') !== null) {
    const t = get('--tag');
    selected = CASES.filter(tc => tc.tag === t);
    if (!selected.length) {
      console.error(`❌ 未找到标签 "${t}" 的用例。可用标签：standard | edge`);
      process.exit(1);
    }
  } else {
    console.log([
      '❌ 未指定执行范围。请使用以下参数之一：',
      '  --list                        列出所有用例',
      '  --case N                      只跑第 N 条',
      '  --cases N,M,K                 跑多条（逗号分隔序号）',
      '  --from N [--to M]             跑第 N～M 条',
      '  --model MODEL_NAME            跑指定模型全部用例',
      '  --tag standard|edge           按标签筛选',
      '  --all                         跑全部',
      '  附加：--dry-run（推演）, --force（强制重跑）, --resume [--task-id xxx]（只轮询补报告）',
    ].join('\n'));
    process.exit(1);
  }

  if (!selected.length) {
    console.log('未匹配到任何用例，退出。');
    return;
  }

  const modeLabel = dryRun ? '（dry-run 模式）' : resume ? '（resume：只轮询，不扣费）' : '';
  console.log(`\n📋 准备执行 ${selected.length} 条用例${modeLabel}：`);
  for (const tc of selected) {
    const label = dryRun || resume ? '' : (loadState().executed[tc.id] && !force ? ' [已执行，将跳过]' : '');
    console.log(`  [${tc.seq}] ${tc.id}: ${tc.name}${label}`);
  }

  // 只有真正会提交扣费时才倒计时警告
  if (!dryRun && !resume) {
    console.log('\n⚠️  将真实提交任务并扣费。5 秒后开始执行，Ctrl+C 取消...');
    await sleep(5000);
  }

  for (const tc of selected) {
    await runCase(tc, { dryRun, force, resume, taskIdOverride });
  }

  console.log('\n✅ 全部完成。');
  console.log(`   结果目录：${RESULTS_DIR}`);
}

main().catch((err) => {
  console.error('❌ 脚本异常：', err.message);
  process.exit(1);
});
