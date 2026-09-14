# Suno V6 模型测试参数文档（生成音乐 / 生成歌词）

> 只覆盖 `suno-music`、`suno-lyrics`。旧版全模型清单见 [`Suno模型测试参数文档.md`](./Suno模型测试参数文档.md)（v5，勿再用于 V6 验收）。

**提交**：`POST /v1/music/generations`  
**查询**：`GET /v1/music/generations/{task_id}`  
**认证**：`Authorization: Bearer sk-xxxxxx`

```bash
# 提交后拿 data[0].task_id 轮询
curl -X GET "$BASE_URL/v1/music/generations/$TASK_ID" \
  -H "Authorization: Bearer YOUR_API_KEY"
```

- 平台状态：`SUBMITTED` / `QUEUED` / `IN_PROGRESS` / `SUCCESS` / `FAILURE` / `UNKNOWN`
- 音乐完成：读 `result.music[]` 的 `audio_url` / `title` / `lyrics` / `tags`
- 歌词完成：优先读 `result.lyrics[]` 的 `text` / `title` / `tags`
- `v6` / `v6-wild` / `v6-mini` / `custom_model_id` **同价**
- `suno-music` 普通档 **0.05 USD**；仅 `max_mode=true` 才 **×2 = 0.10 USD**。`variety: "max"` 不加价
- `suno-lyrics` **0.008 USD**，无 Max、无版本维度
- 依赖前置的占位符：`YOUR_PERSONA_ID`、`YOUR_CUSTOM_MODEL_ID`（分别来自 `suno-persona`、`suno-create-model` 成功结果）

---

## 字段速查（测前先看）

### `suno-music`

| 字段 | 要点 |
| --- | --- |
| `version` | `v6` / `v6-wild` / `v6-mini`；与 `custom_model_id` **至少其一、不能同时传** |
| `custom` | `false`（默认）灵感：`prompt` 是描述；`true` 自定义：`prompt` 是歌词 |
| `style` | **本模型用 `style`，不要传 `tags`**。仅自定义模式生效，≤1000 |
| `weirdness_constraint` | 创意度 0–1。**不要传 `weirdness`**。两种模式都会校验并生效 |
| `style_weight` / `audio_weight` | 0–1，两种模式都会校验并生效；`0` 合法 |
| `prompt` | 灵感必填 ≤3000；自定义且非纯音乐必填歌词 ≤5000；自定义纯音乐可省略 |
| `title` | ≤80，仅自定义生效 |
| `max_mode` | 文档要求 `custom=true`；开启后按 2 倍扣费 |
| `duration` | 整数 10–360，仅自定义；成品时长以结果为准 |
| `variety` | `off` / `normal` / `high` / `extra` / `max`；`max` **不**触发 Max 计费 |
| `audio_format` | `mp3` / `m4a` / `wav` |
| `vocal_gender` | `Male` / `Female` / `m` / `f` / `male` / `female` |
| `persona_id` | 仅自定义；与 `custom_model_id` 互斥 |

### `suno-lyrics`

| 字段 | 要点 |
| --- | --- |
| `prompt` | 必填，去首尾空白后不能为空 |
| `lyrics_model` | `classic` / `remi`；不传走服务默认 |
| `version` | 不要传；传入会被丢弃，不影响计费 |

---

## 一、`suno-music` 正向用例

### 1. 灵感模式 · v6 中文（基线）

目的：最小可用请求。  
预期：提交成功，普通档 0.05；完成后 `result.music[]` 有 `audio_url`。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "深夜城市中的 lo-fi 钢琴，伴随轻柔雨声，节奏缓慢，适合独处"
}
```

### 2. 灵感模式 · v6-wild

目的：覆盖 `v6-wild`，同价。

```json
{
  "model": "suno-music",
  "version": "v6-wild",
  "custom": false,
  "prompt": "赛博朋克夜店，失真电吉他和合成器对撞，能量很高"
}
```

### 3. 灵感模式 · v6-mini

目的：覆盖 `v6-mini`，同价。

```json
{
  "model": "suno-music",
  "version": "v6-mini",
  "custom": false,
  "prompt": "清晨公园里的轻快木吉他小品，明亮、短小、干净"
}
```

### 4. 灵感模式 · 省略 `custom`（默认 false）

目的：确认不传 `custom` 走灵感模式。

```json
{
  "model": "suno-music",
  "version": "v6",
  "prompt": "夏日海边日落，暖色合成器铺底，带一点热带打击乐"
}
```

### 5. 灵感模式 · 英文描述

目的：非中文 prompt。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "An uplifting indie folk song about driving across the desert at dawn, acoustic guitar and warm male vocal"
}
```

### 6. 灵感模式 · 纯音乐

目的：`instrumental=true` + 灵感描述，无人声。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "instrumental": true,
  "prompt": "空灵氛围电子，缓慢脉冲低音，适合冥想，不要人声"
}
```

### 7. 灵感模式 · 女声

目的：灵感模式也生效 `vocal_gender`。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "vocal_gender": "Female",
  "prompt": "日系流行，清亮女声，关于放学后的单车和晚霞"
}
```

### 8. 灵感模式 · 男声（简写 `m`）

目的：覆盖 `vocal_gender` 简写枚举。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "vocal_gender": "m",
  "prompt": "蓝调摇滚，沙哑男声，酒吧舞台，慢摇"
}
```

### 9. 灵感模式 · 权重中间值

目的：V6 权重在灵感模式也会校验并生效，不是静默忽略。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "电影感弦乐配器的中文抒情歌，雨夜城市",
  "style_weight": 0.65,
  "weirdness_constraint": 0.35,
  "audio_weight": 0.5
}
```

### 10. 灵感模式 · 权重下界 0

目的：`0` 合法，不要当成「没传」。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "极简钢琴夜曲，几乎没有打击乐",
  "style_weight": 0,
  "weirdness_constraint": 0,
  "audio_weight": 0
}
```

### 11. 灵感模式 · 权重上界 1

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "强烈实验电子，怪异音色层叠",
  "style_weight": 1,
  "weirdness_constraint": 1,
  "audio_weight": 1
}
```

### 12. 灵感模式 · `variety=off`

目的：覆盖 variety 枚举；不计 2 倍。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "稳定的电台流行，旋律工整，变化不要太大",
  "variety": "off"
}
```

### 13. 灵感模式 · `variety=normal`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "标准中文流行，适中编排",
  "variety": "normal"
}
```

### 14. 灵感模式 · `variety=high`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "编曲层次更丰富的都市流行",
  "variety": "high"
}
```

### 15. 灵感模式 · `variety=extra`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "多段式概念专辑开场曲，风格跳跃但仍能听完",
  "variety": "extra"
}
```

### 16. 灵感模式 · `variety=max`（确认不加价）

目的：**必须核对扣费仍是 0.05**，不是 0.10。`variety: "max"` 不会触发 Max。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "把同一主题做成尽可能多变的流行编曲",
  "variety": "max"
}
```

### 17. 灵感模式 · `audio_format=mp3`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "轻松咖啡馆爵士，人声轻哼",
  "audio_format": "mp3"
}
```

### 18. 灵感模式 · `audio_format=m4a`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "清爽清晨民谣",
  "audio_format": "m4a"
}
```

### 19. 灵感模式 · `audio_format=wav`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "需要后期的原声音色钢琴曲",
  "audio_format": "wav"
}
```

### 20. 自定义模式 · 中文歌词 + 风格（基线）

目的：`custom=true` 时 `prompt` 当歌词，`style` 当风格。  
预期：普通档 0.05。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "instrumental": false,
  "title": "雨窗慢拍",
  "style": "lo-fi, piano ballad, trip-hop, rainy night, chill",
  "prompt": "[Verse]\n我把灯关小一点\n让雨落进窗边\n你留下的那只杯子\n还在桌上转圈\n\n[Chorus]\n深夜的城，慢慢睡去\n我还醒着，等一场回忆"
}
```

### 21. 自定义模式 · 英文歌词

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "title": "Neon After Midnight",
  "style": "synthwave, cinematic, female vocal",
  "prompt": "[Verse]\nStreetlights paint the rain in gold\nI keep your name in the static of the radio\n\n[Chorus]\nAfter midnight we don't look back\nThe city hums like a heart attack"
}
```

### 22. 自定义模式 · 中英混歌词

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "title": "Midnight Taxi",
  "style": "r&b, city pop, male vocal",
  "prompt": "[Verse]\n凌晨两点的出租车上\nI replay the last goodbye\n霓虹一闪一闪\nbut I don't ask why\n\n[Chorus]\n带我回家，don't talk tonight"
}
```

### 23. 自定义纯音乐 · 省略歌词

目的：`custom=true` + `instrumental=true` 可不传 `prompt`。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "instrumental": true,
  "title": "雨巷钢琴",
  "style": "lofi, piano, rain, chill, no vocal"
}
```

### 24. 自定义纯音乐 · 仍带描述性 prompt

目的：纯音乐也可以带 prompt，看上游是否当编曲提示而不是当歌词唱出来。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "instrumental": true,
  "title": "远山",
  "style": "cinematic, strings, ambient",
  "prompt": "缓慢铺开的弦乐，中段加入低频脉冲，结尾渐弱"
}
```

### 25. 自定义 · `auto_lyrics=true`

目的：对输入歌词二次创作，成品歌词不必与输入逐字相同。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "auto_lyrics": true,
  "title": "海边重逢",
  "style": "c-pop, acoustic, warm",
  "prompt": "[Verse]\n夏夜的风把名字吹回来\n沙滩上还留着去年的脚印\n\n[Chorus]\n如果你还记得那首歌\n我们就再唱一次"
}
```

### 26. 自定义 · `auto_lyrics=false` + `negative_tags`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "auto_lyrics": false,
  "title": "安静房间",
  "style": "indie folk, soft vocal, acoustic guitar",
  "negative_tags": "metal, screaming, EDM drop, rap",
  "prompt": "[Verse]\n把窗打开一点点\n让风替我说完那句抱歉\n\n[Chorus]\n今晚不要热闹\n只要一把吉他和你的呼吸"
}
```

### 27. 自定义 · 女声 `Female` + 权重

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "title": "晚风邮局",
  "style": "j-pop, bright, female vocal",
  "vocal_gender": "Female",
  "style_weight": 0.7,
  "weirdness_constraint": 0.25,
  "audio_weight": 0.4,
  "prompt": "[Verse]\n明信片写到一半就停了\n墨水被晚风吹干\n\n[Chorus]\n如果有下一次夏天\n请把地址写得更清楚一点"
}
```

### 28. 自定义 · 男声 `male`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "title": "旧收音机",
  "style": "folk rock, male vocal, storytelling",
  "vocal_gender": "male",
  "prompt": "[Verse]\n父亲留下的收音机还在响\n沙沙的电流里藏着一首老歌\n\n[Chorus]\n我调不准那个年代\n却总在夜里把它打开"
}
```

### 29. 自定义 · `duration=10`（下界）

目的：目标时长最短。成品时长以结果为准，允许偏差。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "duration": 10,
  "title": "十秒钩子",
  "style": "electronic hook, short, punchy",
  "prompt": "[Chorus]\n只给你十秒\n记住我的名字"
}
```

### 30. 自定义 · `duration=120`

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "duration": 120,
  "title": "两分钟情歌",
  "style": "pop ballad, piano, female vocal",
  "prompt": "[Verse]\n路灯把影子拉得很长\n我数着回家的台阶\n\n[Chorus]\n如果时间可以再慢一点\n我想把这首歌唱完"
}
```

### 31. 自定义 · `duration=360`（上界）

目的：最长目标时长，任务会更久，轮询耐心一点。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "duration": 360,
  "title": "长途列车",
  "style": "cinematic folk, narrative, slow build",
  "prompt": "[Verse]\n列车穿过无名的站\n玻璃上全是别人的倒影\n\n[Chorus]\n我带着未寄出的信\n去一个没有坐标的夏天\n\n[Bridge]\n如果终点不是家\n那就把沿途当成答案"
}
```

### 32. 自定义 · `max_mode=true`（必须核对 2 倍扣费）

目的：唯一会变价的正向参数。  
预期：**0.10 USD**，不是 0.05。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "max_mode": true,
  "title": "星光铺路",
  "style": "anthemic pop, piano, wide chorus",
  "prompt": "[Verse]\n星光洒落在深夜的街\n我把脚步放得很轻\n\n[Chorus]\n这条路还很长\n但灯已经亮了"
}
```

### 33. 自定义 · Max + 时长 + wav（计费仍只看 max_mode）

目的：`duration` / `audio_format` 不加价；扣费仍是 0.10。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "max_mode": true,
  "duration": 180,
  "audio_format": "wav",
  "variety": "high",
  "title": "高规格样带",
  "style": "cinematic pop, orchestra, female vocal",
  "vocal_gender": "Female",
  "prompt": "[Verse]\n把所有未说的话\n放进这首还没写完的歌\n\n[Chorus]\n请用最大的声音回答我"
}
```

### 34. 自定义 · `max_mode=false` 显式关闭

目的：显式 false 应与不传相同，扣 0.05。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "max_mode": false,
  "title": "普通档对照",
  "style": "acoustic pop",
  "prompt": "[Verse]\n这是对照样例\n用来确认没有误扣 Max"
}
```

### 35. 自定义 · `v6-wild` + 高创意

```json
{
  "model": "suno-music",
  "version": "v6-wild",
  "custom": true,
  "title": "失控夜",
  "style": "experimental electronic, glitch, alt pop",
  "weirdness_constraint": 0.85,
  "variety": "extra",
  "prompt": "[Verse]\n信号断了又接上\n我在噪声里找你的旋律\n\n[Chorus]\n今晚允许一点失控"
}
```

### 36. 自定义 · `v6-mini` 短歌

```json
{
  "model": "suno-music",
  "version": "v6-mini",
  "custom": true,
  "duration": 45,
  "title": "便利店灯",
  "style": "bedroom pop, lo-fi, soft vocal",
  "prompt": "[Verse]\n便利店的灯亮到很晚\n我买一瓶汽水当作庆祝\n\n[Chorus]\n平凡的胜利也算胜利"
}
```

### 37. 自定义 · 完整字段（厨房水槽，不含互斥项）

目的：一次带齐常用可选字段。不含 `custom_model_id` / `persona_id`。  
预期：`max_mode=true` → **0.10 USD**。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "instrumental": false,
  "title": "全字段样例",
  "style": "synth pop, cinematic, female vocal, 80s drums",
  "negative_tags": "death metal, scream, kids choir",
  "auto_lyrics": false,
  "vocal_gender": "Female",
  "style_weight": 0.62,
  "weirdness_constraint": 0.28,
  "audio_weight": 0.45,
  "variety": "high",
  "max_mode": true,
  "audio_format": "mp3",
  "duration": 150,
  "prompt": "[Verse]\n我把今天折叠好\n放进外套内侧的口袋\n\n[Pre-Chorus]\n路灯一盏盏亮起来\n像有人替我点名\n\n[Chorus]\n如果世界再大声一点\n也请留一句给我"
}
```

### 38. 依赖前置 · `persona_id`（不要同时传 custom_model_id）

目的：自定义模式下绑定人声/风格 Persona。先跑通 `suno-persona` 再测。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "persona_id": "YOUR_PERSONA_ID",
  "title": "Persona 绑定",
  "style": "indie pop, warm vocal",
  "prompt": "[Verse]\n用你熟悉的声音\n把这首未完成的歌唱完\n\n[Chorus]\n我认得这个音色"
}
```

### 39. 依赖前置 · 只传 `custom_model_id`（不要传 version / persona_id）

目的：与公共版本同价 0.05；创建模型费用在 `suno-create-model`，不在本请求。

```json
{
  "model": "suno-music",
  "custom_model_id": "YOUR_CUSTOM_MODEL_ID",
  "custom": true,
  "title": "自定义模型出歌",
  "style": "按训练风格继续写一首新的",
  "prompt": "[Verse]\n沿用你学会的味道\n但歌词是全新的夜晚\n\n[Chorus]\n这是同一条路上的下一站"
}
```

### 40. 行为对照 · 误传 `tags`（应不生效风格）

目的：本模型风格字段是 `style`。此条用于确认误传 `tags` 时请求能否提交、成品风格是否明显未按 tags 走。  
预期：平台不因多传 `tags` 直接 400（字段会被忽略）；风格可能很泛。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "title": "错用 tags",
  "tags": "heavy metal, scream, djent",
  "prompt": "[Verse]\n如果 tags 生效我应该听到金属\n如果没生效就只是普通流行"
}
```

---

## 二、`suno-music` 负例（应 400，不产生成品）

这些用例用来验校验，不要当成功样例记账。

### N1. 既不传 `version` 也不传 `custom_model_id`

预期：400，`suno-music requires either version ... or custom_model_id`。

```json
{
  "model": "suno-music",
  "custom": false,
  "prompt": "缺少版本"
}
```

### N2. `version` 与 `custom_model_id` 同时传

预期：400，二者互斥。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom_model_id": "YOUR_CUSTOM_MODEL_ID",
  "custom": false,
  "prompt": "互斥字段"
}
```

### N3. 旧版本 `v5.5`

预期：400，`unsupported_version`。

```json
{
  "model": "suno-music",
  "version": "v5.5",
  "custom": false,
  "prompt": "旧版本应被拒绝"
}
```

### N4. 权重越界

预期：400，`style_weight must be between 0.00 and 1.00`。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "权重越界",
  "style_weight": 1.5
}
```

### N5. 灵感模式空 prompt

预期：上游或平台 400（灵感模式 prompt 必填）。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": false,
  "prompt": "   "
}
```

### N6. 自定义有人声但没有歌词

预期：400 或上游拒绝（`custom=true` 且 `instrumental=false` 时 prompt 必填）。

```json
{
  "model": "suno-music",
  "version": "v6",
  "custom": true,
  "instrumental": false,
  "title": "缺歌词",
  "style": "pop"
}
```

---

## 三、`suno-lyrics` 正向用例

完成后优先看 `result.lyrics[]`。扣费均为 **0.008 USD**。不要传 `version`。

### 41. 仅 prompt（默认 lyrics_model）

```json
{
  "model": "suno-lyrics",
  "prompt": "写一首关于夏夜海边重逢的中文流行歌词"
}
```

### 42. `lyrics_model=classic` · 中文流行

```json
{
  "model": "suno-lyrics",
  "prompt": "写一首华语流行情歌，主题是下雨天错过末班车，要有主歌、副歌，语言口语化",
  "lyrics_model": "classic"
}
```

### 43. `lyrics_model=remi` · 中文说唱

```json
{
  "model": "suno-lyrics",
  "prompt": "写一首中文说唱，主题是北漂租房和加班，押韵密集，带 Hook，不要脏话",
  "lyrics_model": "remi"
}
```

### 44. classic · 英文流行

```json
{
  "model": "suno-lyrics",
  "prompt": "Write English pop lyrics about leaving a small town at 18, with verse / pre-chorus / chorus / bridge",
  "lyrics_model": "classic"
}
```

### 45. remi · 英文 R&B

```json
{
  "model": "suno-lyrics",
  "prompt": "Write modern R&B lyrics about a late-night phone call you should not pick up, intimate, repetitive hook",
  "lyrics_model": "remi"
}
```

### 46. 日文歌词

```json
{
  "model": "suno-lyrics",
  "prompt": "放課後の自転車と夕焼けをテーマにした日本語のJ-Pop歌詞。Aメロ・サビを含めて。",
  "lyrics_model": "classic"
}
```

### 47. 指定段落结构

目的：看结果是否按 Section 标记组织。

```json
{
  "model": "suno-lyrics",
  "prompt": "写中文摇滚歌词，必须包含：[Verse 1] [Pre-Chorus] [Chorus] [Verse 2] [Bridge] [Final Chorus]。主题：放弃稳定工作去旅行。副歌要好记、适合合唱。",
  "lyrics_model": "classic"
}
```

### 48. 双语歌词

```json
{
  "model": "suno-lyrics",
  "prompt": "写中英双语流行歌词，主歌中文、副歌英文，主题是异地恋时差，副歌要短、好哼",
  "lyrics_model": "classic"
}
```

### 49. 极短主题

目的：最短可用 prompt。

```json
{
  "model": "suno-lyrics",
  "prompt": "冬天",
  "lyrics_model": "classic"
}
```

### 50. 详细创作要求

```json
{
  "model": "suno-lyrics",
  "prompt": "写一首适合女声演唱的都市民谣。人物是 28 岁女性，在便利店兼职，歌词避免直说“加班”“崩溃”。要有具体物象：荧光灯、过期便当、公交末班、手机没电。副歌重复“我还在路上”。不要英文，不要网络梗。",
  "lyrics_model": "classic"
}
```

### 51. 纯音乐向的歌词需求（仍走歌词模型）

目的：描述「不要人声」时，看返回的是空歌词、注释，还是仍写出词。

```json
{
  "model": "suno-lyrics",
  "prompt": "给一首无人声的电影感钢琴曲起一个标题和风格标签，如果必须写词就写极少的哼唱音节，不要完整故事",
  "lyrics_model": "classic"
}
```

### 52. remi · 儿童向

```json
{
  "model": "suno-lyrics",
  "prompt": "写一首给小朋友的中文歌曲歌词，主题是刷牙，欢快、押韵、无恐怖意象",
  "lyrics_model": "remi"
}
```

### 53. 行为对照 · 多传 `version`（应被丢弃）

目的：歌词模型无版本维度。传入 `v6` 应仍按 0.008 计费并成功，不能因为 version 报错。

```json
{
  "model": "suno-lyrics",
  "version": "v6",
  "prompt": "写一首关于春天窗台绿植的短歌词",
  "lyrics_model": "classic"
}
```

---

## 四、`suno-lyrics` 负例

### N7. 缺少 prompt

预期：400，`prompt is required`。

```json
{
  "model": "suno-lyrics",
  "lyrics_model": "classic"
}
```

### N8. prompt 全空白

预期：400，去空白后为空。

```json
{
  "model": "suno-lyrics",
  "prompt": "   "
}
```

---

## 五、建议测试顺序

1. **先跑 1、20、41**：确认提交、轮询、取结果链路通。
2. **再跑 2、3**：三个公共版本都出歌。
3. **计费对照**：16（variety=max，仍 0.05）vs 32（max_mode=true，0.10）vs 34（显式 false，0.05）。
4. **字段面**：7–11 权重、12–19 variety/格式、23–31 自定义/时长。
5. **依赖项**：38、39 等前置任务有了再测。
6. **负例 N1–N8**：确认 400 文案，不要当成失败任务去等音频。
7. **歌词面**：42 vs 43 看 classic / remi 差异；47 看段落结构。

查询时建议 3–5 秒轮询一次。音乐通常比歌词慢；`duration=360` 和 Max 更慢。
