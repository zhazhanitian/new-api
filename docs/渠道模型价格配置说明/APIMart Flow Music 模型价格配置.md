# APIMart Flow Music 模型价格配置

> **渠道类型**：APIMart Flow Music（后台渠道类型编号 **63**）  
> **上游地址**：`https://api.apimart.ai`（代码默认 Base URL 为 `https://apib.ai`，渠道里填哪个就用哪个）  
> **汇率**：APIMart 按 `1 USD = 7 CNY` 标人民币；本平台按 `1 USD = 7.3 CNY` 入账。后台不能直接填 APIMart 美元  
> **积分换算**：`1 Credit = $0.10 USD`（APIMart 固定）  
> **计费方式**：全部工具**按次计费**。Lyria 3 Pro 与 Lyria 3.5 同价，不按 `version`、`length` 加价

---

## 一、模型 ID

一个工具一个模型 ID。上游 `model` 固定是 `flowmusic`，由适配器改写，客户端不要传 `lyria-3.5` 当模型名。

`version` 枚举为 `lyria-3-pro`（Lyria 3 Pro，默认）和 `lyria-3.5`。不传时按默认 `lyria-3-pro`。两个版本的业务字段相同，网关不校验取值，原样交给上游。

| 模型 ID | 功能 | 上游路径 |
|---------|------|----------|
| `flowmusic` | 生成音乐 | `POST /v1/music/generations` |
| `flowmusic-lyrics` | 生成歌词 | `POST /v1/music/generations/lyricsFlowMusic` |
| `flowmusic-extend` | 音乐延伸 | `POST /v1/music/generations/extendFlowMusic` |
| `flowmusic-replace` | 片段替换 | `POST /v1/music/generations/replaceFlowMusic` |
| `flowmusic-cover` | Cover 改编 | `POST /v1/music/generations/coverFlowMusic` |
| `flowmusic-stems` | 词曲分离 | `POST /v1/music/generations/stemsFlowMusic` |
| `flowmusic-upload` | 上传音频 | `POST /v1/music/generations/uploadAudioFlowMusic` |
| `flowmusic-download` | 下载音频 | `POST /v1/music/generations/downloadAudioFlowMusic` |
| `flowmusic-video` | 音乐视频渲染 | `POST /v1/music/generations/videoClipFlowMusic` |

查询走现有 `GET /v1/music/generations/:task_id`。后台轮询上游 `GET /v1/music/tasks/{task_id}`。

跟进操作使用结果里的 `clip_id`。多 Key 时可以在请求里额外带我们自己的 `task_id` 锁定原渠道，这个字段不会转发给上游。

---

## 二、参数

生成、延伸、替换、Cover 的 Lyria 3 Pro 和 Lyria 3.5 字段相同。唯一差别是 3.5 多一个 `version`。歌词、分离、上传、下载、视频没有 `version`。网关不白名单 `version`，客户端传了就原样转发。

取值（时长、强度、格式等）由上游返回 4xx。提交失败会退掉预扣。

| 模型 ID | 字段 |
|---------|------|
| `flowmusic` | `sound_prompt`、`lyrics`（至少有一个）、`title`、`bpm`、`length`（1–240 秒）、`seed`，可选 `version` |
| `flowmusic-lyrics` | `prompt`（≤ 3000） |
| `flowmusic-extend` | `clip_id`、`extend_from_s`、`extend_s`（最大 164）、`instruction`、`title`、`seed`，可选 `version` |
| `flowmusic-replace` | `clip_id`、`start_s`、`end_s`（必须大于 `start_s`）、`instruction`、`title`、`seed`，可选 `version` |
| `flowmusic-cover` | `clip_id`、`instruction`、`strength`（0–1）、`title`、`seed`，可选 `version` |
| `flowmusic-stems` | `clip_id`。结果是 zip（`result.music[0].file_url`） |
| `flowmusic-upload` | `audio_url`（公网 `.mp3` / `.wav`） |
| `flowmusic-download` | `clip_id`、`format`（`wav` 或 `mp3`） |
| `flowmusic-video` | `clip_id`、`preset`（`simple` / `modern` / `player`，默认 `simple`）。结果在 `video_url` |

Lyria 3.5 成曲示例：

```json
{
  "model": "flowmusic",
  "version": "lyria-3.5",
  "title": "坚持",
  "lyrics": "[Verse 1]\n黑夜再长也会天亮",
  "sound_prompt": "energetic rock with electric guitar",
  "length": 120
}
```

`version` 不传时默认 `lyria-3-pro`。也可以显式传 `lyria-3-pro` 或 `lyria-3.5`。

---

## 三、价格

填**黄金价（价目表「当前价格」）**。官方价只作对照。铂金、钻石是账号档位，不写进 ModelPrice。

APIMart 美元先乘 7 得到人民币，再除以 7.3 才是后台 ModelPrice。这样平台按 7.3 展示和扣费时，人民币与 APIMart 标价一致。直接填 0.06 会按 0.438 元入账，比 APIMart 的 0.42 元高。

| 模型 ID | 功能 | 黄金价 Credits | APIMart 美元 | APIMart 人民币（×7） | 后台填值（ModelPrice） | 官方价 Credits |
|---------|------|----------------|--------------|----------------------|------------------------|----------------|
| `flowmusic` | 生成音乐 | 0.6 | 0.06 | 0.42 元 | **0.057534** | 0.75 |
| `flowmusic-extend` | 音乐延伸 | 0.6 | 0.06 | 0.42 元 | **0.057534** | 0.75 |
| `flowmusic-replace` | 片段替换 | 0.6 | 0.06 | 0.42 元 | **0.057534** | 0.75 |
| `flowmusic-cover` | Cover 改编 | 0.6 | 0.06 | 0.42 元 | **0.057534** | 0.75 |
| `flowmusic-stems` | 词曲分离 | 0.6 | 0.06 | 0.42 元 | **0.057534** | 0.75 |
| `flowmusic-lyrics` | 生成歌词 | 0.2 | 0.02 | 0.14 元 | **0.019178** | 0.25 |
| `flowmusic-download` | 下载音频 | 0.2 | 0.02 | 0.14 元 | **0.019178** | 0.25 |
| `flowmusic-video` | 音乐视频 | 0.2 | 0.02 | 0.14 元 | **0.019178** | 0.25 |
| `flowmusic-upload` | 上传音频 | 0.1 | 0.01 | 0.07 元 | **0.009589** | 0.125 |

分组倍率仍由分组配置乘上去。`version` 不改变价格。

文档示例里出现过的 `credits_cost: 0.48` 是旧样本。查询接口示例已是 `0.6`，与当前黄金价一致。

---

## 四、换算

```
APIMart 美元 = Credits × $0.10
APIMart 人民币 = APIMart 美元 × 7
后台 ModelPrice = APIMart 人民币 ÷ 7.3
平台人民币 = ModelPrice × 7.3
内部 quota = ModelPrice × 500,000

示例：flowmusic
  0.6 Credits × $0.10 = $0.06
  $0.06 × 7 = 0.42 元
  0.42 ÷ 7.3 = 0.057534
  0.057534 × 7.3 = 0.42 元
  0.057534 × 500,000 = 28,767 quota
  后台填 ModelPrice = 0.057534
```

默认价已经写在 `setting/ratio_setting/model_ratio.go` 的 `defaultModelPrice`。后台如果已经保存过模型价格，以后台为准。

---

## 五、后台操作步骤

1. 新建渠道，类型选 **APIMart Flow Music**
2. Base URL 填 `https://api.apimart.ai` 或 `https://apib.ai`，Key 填 APIMart Bearer Token
3. 模型填上表 9 个 ID（经典后台切换渠道类型时会带出预设）
4. 进入 **系统设置 → 模型倍率**，确认这 9 个模型是**按次计费**，ModelPrice 与第三节「后台填值」一致
