# 腾讯 VOD 生视频模型价格配置说明

> 来源：腾讯云点播官方价格文档，AIGC 生视频计费。
> 官方价格为人民币元/秒；平台表达式值按固定汇率 `1 USD = 7.3 CNY` 换算：
>
> `表达式值 = 人民币单价 ÷ 7.3 × 1,000,000`
>
> 后台请使用「表达式/阶梯计费」。表达式返回值仍是平台内部的美元百万分比值，系统会按现有 quota 规则结算。

---

## 一、当前接入模型

来自 `relay/channel/task/tencentvod/video_constants.go` 的腾讯 VOD 视频模型：

| 系列 | 内部模型名 |
| --- | --- |
| Kling | `kling-video-1.6`、`kling-video-2.0`、`kling-video-2.1`、`kling-video-2.5`、`kling-video-2.6`、`kling-video-o1`、`kling-video-3.0`、`kling-video-3.0-omni` |
| Vidu | `vidu-video-q2`、`vidu-video-q2-pro`、`vidu-video-q2-turbo`、`vidu-video-q3`、`vidu-video-q3-pro`、`vidu-video-q3-turbo` |
| Hailuo | `hailuo-video-02`、`hailuo-video-2.3`、`hailuo-video-2.3-fast`、`hailuo-video-h3` |
| Google Veo | `veo-video-3.1`、`veo-video-3.1-fast` |
| OpenAI Sora | `sora-video-2.0` |
| Hunyuan | `hunyuan-video-1.5` |
| Mingmou | `mingmou-video-1.0` |
| PixVerse | `pixverse-video-v5.6`、`pixverse-video-v6`、`pixverse-video-c1` |

---

## 二、官方视频模型价格清单

单位：人民币元/秒。

| 厂商 | 模型/版本 | 属性 | 480/540P | 720/768P | 1080P | 2K | 4K |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Hunyuan | 1.5 | - | - | 0.300 | 0.500 | 0.750 | 1.120 |
| Vidu | q3 | 参考生 | 0.313 | 0.625 | 0.782 | 0.939 | 1.127 |
| Vidu | q3-pro | 图生、文生、首尾帧生 | 0.313 | 0.782 | 0.938 | 1.126 | 1.351 |
| Vidu | q3-turbo | 图生、文生、首尾帧生 | 0.250 | 0.375 | 0.438 | 0.526 | 0.631 |
| Vidu | q2 | 文生 | - | 0.320 | 0.470 | 0.700 | 1.050 |
| Vidu | q2 | 参考生 | 0.240 | 0.320 | 0.820 | 1.230 | 1.845 |
| Vidu | q2-pro | 图生、首尾帧生 | - | 0.350 | 0.700 | 1.000 | 1.500 |
| Vidu | q2-pro | 参考生 | 0.270 | 0.350 | 0.900 | 1.350 | 2.025 |
| Vidu | q2-turbo | 图生、首尾帧生 | - | 0.250 | 0.470 | 0.700 | 1.050 |
| Kling | 3.0-Omni | 无参考视频+无声 | - | 0.600 | 0.800 | 1.000 | 3.000 |
| Kling | 3.0-Omni | 无参考视频+有声 | - | 0.800 | 1.000 | 1.200 | 3.000 |
| Kling | 3.0-Omni | 有参考视频+无声 | - | 0.900 | 1.200 | 1.500 | 2.000 |
| Kling | 3.0 | 无声 | - | 0.600 | 0.800 | 1.000 | 3.000 |
| Kling | 3.0 | 有声+无音色 | - | 0.900 | 1.200 | 1.500 | 3.000 |
| Kling | O1 | 无参考视频 | - | 0.600 | 0.800 | 1.200 | 1.800 |
| Kling | O1 | 有参考视频 | - | 0.900 | 1.200 | 1.800 | 2.700 |
| Kling | 2.6 | 无声 | - | 0.300 | 0.500 | 0.750 | 1.120 |
| Kling | 2.6 | 有声 | - | - | 1.000 | 1.500 | 2.250 |
| Kling | 2.5-pro | - | - | 0.300 | 0.500 | 0.750 | 1.120 |
| Kling | 1.6、2.0、2.1 | - | - | 0.400 | 0.700 | 1.000 | 1.500 |
| Hailuo | 02、2.3 | - | - | 0.330 | 0.580 | 0.930 | 1.490 |
| Hailuo | 2.3-fast | - | - | 0.225 | 0.385 | 0.580 | 0.870 |
| Hailuo | H3 | - | - | **待确认** | **待确认** | **待确认** | **待确认** |
| GV | 3.1 | 有声 | - | 3.000 | 3.000 | 3.750 | 4.500 |
| GV | 3.1 | 无声 | - | 1.500 | 1.500 | 2.250 | 3.000 |
| GV | 3.1-fast | 有声 | - | 1.125 | 1.125 | 1.875 | 2.625 |
| GV | 3.1-fast | 无声 | - | 0.750 | 0.750 | 1.500 | 2.250 |
| OS | 2.0 | - | - | 0.750 | 1.125 | 1.688 | 2.531 |
| PixVerse | V5.6 | 无声 | 0.245 | 0.315 | 0.525 | 0.735 | 1.029 |
| PixVerse | V6.0 | 无声 | 0.205 | 0.264 | 0.528 | 0.634 | 0.760 |
| PixVerse | V6.0 | 有声 | 0.264 | 0.352 | 0.675 | 0.810 | 0.971 |
| PixVerse | C1 | 无声 | 0.235 | 0.293 | 0.557 | 0.669 | 0.803 |
| PixVerse | C1 | 有声 | 0.293 | 0.381 | 0.704 | 0.845 | 1.014 |
| Mingmou | 1.0 | - | - | 0.300 | 0.500 | 0.750 | 1.120 |

说明：

- 价格表达式默认按请求入参预扣。腾讯实际出账按输出视频秒数和实际落档结算，后台配置无法在预扣时读取腾讯回包中的最终落档。
- 当前适配器里 Kling、Hailuo、Vidu、GV、OS 只会向腾讯传 720/768P 或 1080P 分辨率档；PixVerse 会传 540p/720p/1080p/2k/4k。表达式保留官方 2K/4K 分支，便于后续适配器放开分辨率后直接生效。
- 视频时长读取 `duration`；缺省时长（不传 `duration`）各模型使用 API 官方默认值：Kling 为 5 秒，Hailuo 为 6 秒，Vidu / PixVerse 由 API 决定，GV 固定 8 秒，OS 为 8 秒，Hunyuan / Mingmou 不发送 Duration 字段。

---

## 三、表达式通用解读（先看这段）

所有视频模型表达式的结构都是同一套路：

```
最终预扣值 = 单价 u × 计费秒数 d
```

其中：

| 符号/函数 | 含义 |
| --- | --- |
| `param("x")` | 读取请求 body 字段 `x`；没有该字段时返回 `nil` |
| `param("metadata.xxx")` | 读取 `metadata` 对象里的子字段 |
| `param("images.#")` | `images` 数组长度（gjson `#`） |
| `has(s, "1080")` | 字符串 `s` 是否包含子串 `"1080"`（用于识别分辨率档） |
| `tier("名字", 数值)` | 标记当前命中的价格档，返回该数值（便于后台日志区分档位） |
| `let x = ...` | 定义中间变量 |

**换算关系（写死在配置说明里）：**

```
表达式单价 u = 官方人民币元/秒 ÷ 7.3 × 1,000,000
预扣表达式值 = u × d
实际扣减 Quota ≈ 预扣表达式值 × 0.5
约合人民币 ≈ 预扣表达式值 ÷ 1,000,000 × 7.3
```

**分辨率字段取值顺序：**

1. 优先读 `size`
2. `size` 为空时读 `metadata.resolution`
3. 两者都空 / 无法识别时，落到该模型的默认档（多数是 720P / 768P）

**有声判断统一规则：**

- 仅当 `metadata.audio_generation == "Enabled"` 时走有声价
- 不传、传错（如 `"yes"`）一律按无声价

---

## 四、逐模型价格表达式

### 任务 1：`kling-video-1.6`

官方价：720P `0.400`，1080P `0.700`，2K `1.000`，4K `1.500` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 3 ? 3 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 205479) : has(s, "2K") || has(s, "2k") ? tier("2K", 136986) : has(s, "1080") ? tier("1080P", 95890) : tier("720P", 54795);
u * d
```

**表达式解读：**

1. **读时长 `d0`**：有 `duration` 就转成整数，没有就当 `0`。
2. **规范化时长 `d`**（对齐 Kling API：合法 3–15 秒，默认 5）：
   - `d0 ≤ 0`（未传/0/负数）→ `d = 5`
   - `d0 < 3` → `d = 3`
   - `d0 > 15` → `d = 15`
   - 其余 → `d = d0`
3. **读分辨率字符串 `s`**：`size` 优先，否则 `metadata.resolution`。
4. **选单价 `u`**（按字符串包含关系匹配，优先级 4K > 2K > 1080 > 默认 720P）：

| 命中档 | 官方元/秒 | 表达式单价 u |
| --- | ---: | ---: |
| 4K | 1.500 | 205479 |
| 2K | 1.000 | 136986 |
| 1080P | 0.700 | 95890 |
| 720P（默认） | 0.400 | 54795 |

5. **结果**：`u * d`

**计算示例：**

| 请求 | d | u | 表达式值 | 约合 RMB |
| --- | ---: | ---: | ---: | ---: |
| 不传 duration/size | 5 | 54795 | 273975 | ¥2.00 |
| `duration=10, size=1080P` | 10 | 95890 | 958900 | ¥7.00 |

---

### 任务 2：`kling-video-2.0`

同 `kling-video-1.6`。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 3 ? 3 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 205479) : has(s, "2K") || has(s, "2k") ? tier("2K", 136986) : has(s, "1080") ? tier("1080P", 95890) : tier("720P", 54795);
u * d
```

**表达式解读：** 与任务 1 完全相同（时长 clamp、分辨率分档、单价表一致）。

---

### 任务 3：`kling-video-2.1`

同 `kling-video-1.6`。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 3 ? 3 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 205479) : has(s, "2K") || has(s, "2k") ? tier("2K", 136986) : has(s, "1080") ? tier("1080P", 95890) : tier("720P", 54795);
u * d
```

**表达式解读：** 与任务 1 完全相同。

---

### 任务 4：`kling-video-2.5`

按官方 `2.5-pro` 价格配置：720P `0.300`，1080P `0.500`，2K `0.750`，4K `1.120` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 3 ? 3 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 153425) : has(s, "2K") || has(s, "2k") ? tier("2K", 102740) : has(s, "1080") ? tier("1080P", 68493) : tier("720P", 41096);
u * d
```

**表达式解读：**

- **时长逻辑**：同 Kling 1.6（默认 5，clamp 到 3–15）。
- **单价表**（仅数字不同）：

| 命中档 | 官方元/秒 | 表达式单价 u |
| --- | ---: | ---: |
| 4K | 1.120 | 153425 |
| 2K | 0.750 | 102740 |
| 1080P | 0.500 | 68493 |
| 720P（默认） | 0.300 | 41096 |

- **结果**：`u * d`
- **示例**：不传参数 → `41096 × 5 = 205480` ≈ ¥1.50

---

### 任务 5：`kling-video-2.6`

无声价：720P `0.300`，1080P `0.500`，2K `0.750`，4K `1.120` 元/秒。有声价：1080P `1.000`，2K `1.500`，4K `2.250` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 3 ? 3 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let audio = param("metadata.audio_generation") == "Enabled";
let u = audio ? (has(s, "4K") || has(s, "4k") ? tier("audio-4K", 308219) : has(s, "2K") || has(s, "2k") ? tier("audio-2K", 205479) : tier("audio-1080P", 136986)) : (has(s, "4K") || has(s, "4k") ? tier("silent-4K", 153425) : has(s, "2K") || has(s, "2k") ? tier("silent-2K", 102740) : has(s, "1080") ? tier("silent-1080P", 68493) : tier("silent-720P", 41096));
u * d
```

**表达式解读：**

1. **时长 `d`**：同 Kling（默认 5，3–15）。
2. **分支 `audio`**：`metadata.audio_generation == "Enabled"` → 有声，否则无声。
3. **选单价 `u`**：先分有声/无声，再按分辨率：

| 分支 | 命中档 | 官方元/秒 | u |
| --- | --- | ---: | ---: |
| 无声 | 720P（默认） | 0.300 | 41096 |
| 无声 | 1080P | 0.500 | 68493 |
| 无声 | 2K | 0.750 | 102740 |
| 无声 | 4K | 1.120 | 153425 |
| 有声 | 1080P（默认，含未识别/720） | 1.000 | 136986 |
| 有声 | 2K | 1.500 | 205479 |
| 有声 | 4K | 2.250 | 308219 |

> 注意：有声分支官方没有单独列 720P；表达式在有声且未命中 2K/4K 时统一落到 `audio-1080P`。

4. **结果**：`u * d`
5. **示例**：
   - 无声默认 → `41096 × 5 = 205480`
   - 有声 + 1080P × 5s → `136986 × 5 = 684930`

---

### 任务 6：`kling-video-o1`

无参考视频：720P `0.600`，1080P `0.800`，2K `1.200`，4K `1.800` 元/秒。有参考视频：720P `0.900`，1080P `1.200`，2K `1.800`，4K `2.700` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 3 ? 3 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let ref = param("images.#") != nil && param("images.#") > 0;
let u = ref ? (has(s, "4K") || has(s, "4k") ? tier("ref-4K", 369863) : has(s, "2K") || has(s, "2k") ? tier("ref-2K", 246575) : has(s, "1080") ? tier("ref-1080P", 164384) : tier("ref-720P", 123288)) : (has(s, "4K") || has(s, "4k") ? tier("t2v-4K", 246575) : has(s, "2K") || has(s, "2k") ? tier("t2v-2K", 164384) : has(s, "1080") ? tier("t2v-1080P", 109589) : tier("t2v-720P", 82192));
u * d
```

**表达式解读：**

1. **时长 `d`**：同 Kling（默认 5，3–15）。
2. **分支 `ref`**：`images` 数组长度 > 0 → 有参考视频价，否则文生（t2v）价。
3. **单价表**：

| 分支 | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: |
| 无参考（t2v）官方元/秒 | 0.600 | 0.800 | 1.200 | 1.800 |
| 无参考 u | 82192 | 109589 | 164384 | 246575 |
| 有参考（ref）官方元/秒 | 0.900 | 1.200 | 1.800 | 2.700 |
| 有参考 u | 123288 | 164384 | 246575 | 369863 |

4. **结果**：`u * d`
5. **示例**：无图默认 → `82192 × 5 = 410960`；有图 + 1080P × 10s → `164384 × 10 = 1643840`

---

### 任务 7：`kling-video-3.0`

无声价：720P `0.600`，1080P `0.800`，2K `1.000`，4K `3.000` 元/秒。有声+无音色价：720P `0.900`，1080P `1.200`，2K `1.500`，4K `3.000` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 3 ? 3 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let audio = param("metadata.audio_generation") == "Enabled";
let u = audio ? (has(s, "4K") || has(s, "4k") ? tier("audio-4K", 410959) : has(s, "2K") || has(s, "2k") ? tier("audio-2K", 205479) : has(s, "1080") ? tier("audio-1080P", 164384) : tier("audio-720P", 123288)) : (has(s, "4K") || has(s, "4k") ? tier("silent-4K", 410959) : has(s, "2K") || has(s, "2k") ? tier("silent-2K", 136986) : has(s, "1080") ? tier("silent-1080P", 109589) : tier("silent-720P", 82192));
u * d
```

**表达式解读：**

1. **时长 `d`**：同 Kling（默认 5，3–15）。
2. **分支 `audio`**：`Enabled` → 有声价，否则无声价。
3. **单价表**：

| 分支 | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: |
| 无声 官方元/秒 | 0.600 | 0.800 | 1.000 | 3.000 |
| 无声 u | 82192 | 109589 | 136986 | 410959 |
| 有声 官方元/秒 | 0.900 | 1.200 | 1.500 | 3.000 |
| 有声 u | 123288 | 164384 | 205479 | 410959 |

4. **结果**：`u * d`
5. **示例**：无声默认 → `82192 × 5 = 410960`；有声 1080P × 5s → `164384 × 5 = 821920`

---

### 任务 8：`kling-video-3.0-omni`

无参考视频+无声：720P `0.600`，1080P `0.800`，2K `1.000`，4K `3.000` 元/秒。无参考视频+有声：720P `0.800`，1080P `1.000`，2K `1.200`，4K `3.000` 元/秒。有参考视频+无声：720P `0.900`，1080P `1.200`，2K `1.500`，4K `2.000` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 3 ? 3 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let audio = param("metadata.audio_generation") == "Enabled";
let ref = param("images.#") != nil && param("images.#") > 0;
let u = ref ? (has(s, "4K") || has(s, "4k") ? tier("ref-4K", 273973) : has(s, "2K") || has(s, "2k") ? tier("ref-2K", 205479) : has(s, "1080") ? tier("ref-1080P", 164384) : tier("ref-720P", 123288)) : audio ? (has(s, "4K") || has(s, "4k") ? tier("audio-4K", 410959) : has(s, "2K") || has(s, "2k") ? tier("audio-2K", 164384) : has(s, "1080") ? tier("audio-1080P", 136986) : tier("audio-720P", 109589)) : (has(s, "4K") || has(s, "4k") ? tier("silent-4K", 410959) : has(s, "2K") || has(s, "2k") ? tier("silent-2K", 136986) : has(s, "1080") ? tier("silent-1080P", 109589) : tier("silent-720P", 82192));
u * d
```

**表达式解读：**

这是计费变量最多的 Kling 模型，分支优先级：

```
有参考图(ref) ？ → 走「有参考+无声」价表
否则有声(audio) ？ → 走「无参考+有声」价表
否则 → 走「无参考+无声」价表
```

> 当前表达式在 `ref=true` 时**不再区分有声/无声**，统一走参考视频价（与官方表「有参考视频+无声」对齐）。

| 分支 | 720P u / 元 | 1080P u / 元 | 2K u / 元 | 4K u / 元 |
| --- | --- | --- | --- | --- |
| 无参考+无声 | 82192 / 0.600 | 109589 / 0.800 | 136986 / 1.000 | 410959 / 3.000 |
| 无参考+有声 | 109589 / 0.800 | 136986 / 1.000 | 164384 / 1.200 | 410959 / 3.000 |
| 有参考 | 123288 / 0.900 | 164384 / 1.200 | 205479 / 1.500 | 273973 / 2.000 |

- **时长**：同 Kling（默认 5，3–15）
- **结果**：`u * d`
- **示例**：无图无声默认 → `82192 × 5 = 410960`；有图 1080P × 5s → `164384 × 5 = 821920`

---

### 任务 9：`vidu-video-q2`

文生价：720P `0.320`，1080P `0.470`，2K `0.700`，4K `1.050` 元/秒。参考生价：540P `0.240`，720P `0.320`，1080P `0.820`，2K `1.230`，4K `1.845` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 10 ? 10 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let ref = param("images.#") != nil && param("images.#") > 0;
let u = ref ? (has(s, "4K") || has(s, "4k") ? tier("ref-4K", 252740) : has(s, "2K") || has(s, "2k") ? tier("ref-2K", 168493) : has(s, "1080") ? tier("ref-1080P", 112329) : has(s, "540") || has(s, "480") ? tier("ref-540P", 32877) : tier("ref-720P", 43836)) : (has(s, "4K") || has(s, "4k") ? tier("t2v-4K", 143836) : has(s, "2K") || has(s, "2k") ? tier("t2v-2K", 95890) : has(s, "1080") ? tier("t2v-1080P", 64384) : tier("t2v-720P", 43836));
u * d
```

**表达式解读：**

1. **时长 `d`**（Vidu：1–10 秒，默认按 5）：
   - `d0 ≤ 0` → 5
   - `d0 < 1` → 1
   - `d0 > 10` → 10
   - 其余 → `d0`
2. **分支 `ref`**：有 `images` → 参考生价，否则文生价。
3. **单价表**：

| 分支 | 540P | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: | ---: |
| 文生 官方元/秒 | - | 0.320 | 0.470 | 0.700 | 1.050 |
| 文生 u | - | 43836 | 64384 | 95890 | 143836 |
| 参考生 官方元/秒 | 0.240 | 0.320 | 0.820 | 1.230 | 1.845 |
| 参考生 u | 32877 | 43836 | 112329 | 168493 | 252740 |

4. **结果**：`u * d`
5. **示例**：文生默认 → `43836 × 5 = 219180`；参考生 1080P × 8s → `112329 × 8 = 898632`

---

### 任务 10：`vidu-video-q2-pro`

默认图生/首尾帧生价；当 `metadata.input_usage == "Reference"` 时按参考生价。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 10 ? 10 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let ref = param("metadata.input_usage") == "Reference";
let u = ref ? (has(s, "4K") || has(s, "4k") ? tier("ref-4K", 277397) : has(s, "2K") || has(s, "2k") ? tier("ref-2K", 184932) : has(s, "1080") ? tier("ref-1080P", 123288) : has(s, "540") || has(s, "480") ? tier("ref-540P", 36986) : tier("ref-720P", 47945)) : (has(s, "4K") || has(s, "4k") ? tier("i2v-4K", 205479) : has(s, "2K") || has(s, "2k") ? tier("i2v-2K", 136986) : has(s, "1080") ? tier("i2v-1080P", 95890) : tier("i2v-720P", 47945));
u * d
```

**表达式解读：**

与 q2 的关键差异：**不是看有没有图，而是看 `metadata.input_usage`**：

- `input_usage == "Reference"` → 参考生价（ref）
- 否则（含不传）→ 图生/首尾帧价（i2v）

| 分支 | 540P | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: | ---: |
| i2v 官方元/秒 | - | 0.350 | 0.700 | 1.000 | 1.500 |
| i2v u | - | 47945 | 95890 | 136986 | 205479 |
| ref 官方元/秒 | 0.270 | 0.350 | 0.900 | 1.350 | 2.025 |
| ref u | 36986 | 47945 | 123288 | 184932 | 277397 |

- **时长**：同 Vidu（默认 5，1–10）
- **示例**：有图但不传 input_usage → i2v-720P：`47945 × 5 = 239725`；`input_usage=Reference` + 1080P × 8s → `123288 × 8 = 986304`

---

### 任务 11：`vidu-video-q2-turbo`

图生/首尾帧生价：720P `0.250`，1080P `0.470`，2K `0.700`，4K `1.050` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 10 ? 10 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 143836) : has(s, "2K") || has(s, "2k") ? tier("2K", 95890) : has(s, "1080") ? tier("1080P", 64384) : tier("720P", 34247);
u * d
```

**表达式解读：**

- **无分支**：只按分辨率选单价。
- **时长**：Vidu（默认 5，1–10）。
- **单价**：720P=34247，1080P=64384，2K=95890，4K=143836。
- **示例**：720P × 5s → `34247 × 5 = 171235`

---

### 任务 12：`vidu-video-q3`

参考生价：540P `0.313`，720P `0.625`，1080P `0.782`，2K `0.939`，4K `1.127` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 10 ? 10 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 154384) : has(s, "2K") || has(s, "2k") ? tier("2K", 128630) : has(s, "1080") ? tier("1080P", 107123) : has(s, "540") || has(s, "480") ? tier("540P", 42877) : tier("720P", 85616);
u * d
```

**表达式解读：**

- 仅参考生一档价格；分辨率匹配顺序：4K > 2K > 1080 > 540/480 > 默认 720P。
- **单价 u**：540P=42877，720P=85616，1080P=107123，2K=128630，4K=154384。
- **示例**：720P × 5s → `85616 × 5 = 428080`

---

### 任务 13：`vidu-video-q3-pro`

图生、文生、首尾帧生价：540P `0.313`，720P `0.782`，1080P `0.938`，2K `1.126`，4K `1.351` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 10 ? 10 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 185068) : has(s, "2K") || has(s, "2k") ? tier("2K", 154247) : has(s, "1080") ? tier("1080P", 128493) : has(s, "540") || has(s, "480") ? tier("540P", 42877) : tier("720P", 107123);
u * d
```

**表达式解读：**

- 无额外分支，仅分辨率分档。
- **单价 u**：540P=42877，720P=107123，1080P=128493，2K=154247，4K=185068。
- **示例**：720P × 5s → `107123 × 5 = 535615`

---

### 任务 14：`vidu-video-q3-turbo`

图生、文生、首尾帧生价：540P `0.250`，720P `0.375`，1080P `0.438`，2K `0.526`，4K `0.631` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 10 ? 10 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 86438) : has(s, "2K") || has(s, "2k") ? tier("2K", 72055) : has(s, "1080") ? tier("1080P", 60000) : has(s, "540") || has(s, "480") ? tier("540P", 34247) : tier("720P", 51370);
u * d
```

**表达式解读：**

- 无额外分支，仅分辨率分档。
- **单价 u**：540P=34247，720P=51370，1080P=60000，2K=72055，4K=86438。
- **示例**：720P × 5s → `51370 × 5 = 256850`

---

### 任务 15：`hailuo-video-02`

官方价：768P `0.330`，1080P `0.580`，2K `0.930`，4K `1.490` 元/秒。海螺时长只支持 6 或 10 秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 8 ? 6 : 10;
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 204110) : has(s, "2K") || has(s, "2k") ? tier("2K", 127397) : has(s, "1080") ? tier("1080P", 79452) : tier("768P", 45205);
u * d
```

**表达式解读：**

1. **时长 snap**（对齐适配器：`≤8 → 6`，否则 `10`）：
   - 未传 `duration` 时 `d0=0 ≤ 8` → **默认 6 秒**
   - `duration=10` → 10 秒
2. **分辨率**：含 `1080` → 1080P；含 2K/4K 对应档；否则默认 **768P**（不是 720P）。
3. **单价**：768P=45205，1080P=79452，2K=127397，4K=204110。
4. **示例**：默认 → `45205 × 6 = 271230`；1080P × 10s → `79452 × 10 = 794520`

---

### 任务 16：`hailuo-video-2.3`

同 `hailuo-video-02`。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 8 ? 6 : 10;
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 204110) : has(s, "2K") || has(s, "2k") ? tier("2K", 127397) : has(s, "1080") ? tier("1080P", 79452) : tier("768P", 45205);
u * d
```

**表达式解读：** 与任务 15 完全相同。

---

### 任务 17：`hailuo-video-2.3-fast`

官方价：768P `0.225`，1080P `0.385`，2K `0.580`，4K `0.870` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 8 ? 6 : 10;
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 119178) : has(s, "2K") || has(s, "2k") ? tier("2K", 79452) : has(s, "1080") ? tier("1080P", 52740) : tier("768P", 30822);
u * d
```

**表达式解读：**

- **时长逻辑**：同 Hailuo（`≤8 → 6`，否则 10）。
- **单价**：768P=30822，1080P=52740，2K=79452，4K=119178。
- **示例**：默认 → `30822 × 6 = 184932`

---

### 任务 18：`veo-video-3.1`

GV 3.1 有声价：720P/1080P `3.000`，2K `3.750`，4K `4.500` 元/秒。无声价：720P/1080P `1.500`，2K `2.250`，4K `3.000` 元/秒。当前适配器固定传 8 秒。

```text
let d = 8;
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let audio = param("metadata.audio_generation") == "Enabled";
let u = audio ? (has(s, "4K") || has(s, "4k") ? tier("audio-4K", 616438) : has(s, "2K") || has(s, "2k") ? tier("audio-2K", 513699) : tier("audio-720_1080P", 410959)) : (has(s, "4K") || has(s, "4k") ? tier("silent-4K", 410959) : has(s, "2K") || has(s, "2k") ? tier("silent-2K", 308219) : tier("silent-720_1080P", 205479));
u * d
```

**表达式解读：**

1. **时长写死 `d = 8`**：不读请求里的 `duration`（适配器也会强制发 8）。
2. **分支 `audio`**：`Enabled` → 有声，否则无声。
3. **分辨率**：官方 720P 与 1080P 同价，所以表达式合并为 `720_1080P` 一档。

| 分支 | 720/1080P | 2K | 4K |
| --- | ---: | ---: | ---: |
| 无声 官方元/秒 | 1.500 | 2.250 | 3.000 |
| 无声 u | 205479 | 308219 | 410959 |
| 有声 官方元/秒 | 3.000 | 3.750 | 4.500 |
| 有声 u | 410959 | 513699 | 616438 |

4. **示例**：无声默认 → `205479 × 8 = 1643832`；有声默认 → `410959 × 8 = 3287672`

---

### 任务 19：`veo-video-3.1-fast`

GV 3.1-fast 有声价：720P/1080P `1.125`，2K `1.875`，4K `2.625` 元/秒。无声价：720P/1080P `0.750`，2K `1.500`，4K `2.250` 元/秒。当前适配器固定传 8 秒。

```text
let d = 8;
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let audio = param("metadata.audio_generation") == "Enabled";
let u = audio ? (has(s, "4K") || has(s, "4k") ? tier("audio-4K", 359589) : has(s, "2K") || has(s, "2k") ? tier("audio-2K", 256849) : tier("audio-720_1080P", 154110)) : (has(s, "4K") || has(s, "4k") ? tier("silent-4K", 308219) : has(s, "2K") || has(s, "2k") ? tier("silent-2K", 205479) : tier("silent-720_1080P", 102740));
u * d
```

**表达式解读：**

- 结构同任务 18（`d=8` + 有声/无声 + 720/1080 同价）。
- **单价**：

| 分支 | 720/1080P u | 2K u | 4K u |
| --- | ---: | ---: | ---: |
| 无声 | 102740 | 205479 | 308219 |
| 有声 | 154110 | 256849 | 359589 |

- **示例**：无声默认 → `102740 × 8 = 821920`

---

### 任务 20：`sora-video-2.0`

官方价：720P `0.750`，1080P `1.125`，2K `1.688`，4K `2.531` 元/秒。官方默认 8 秒；用户指定时 snap 到最近合法值 {4, 8, 12}。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 8 : (d0 <= 6 ? 4 : (d0 <= 10 ? 8 : 12));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 346712) : has(s, "2K") || has(s, "2k") ? tier("2K", 231233) : has(s, "1080") ? tier("1080P", 154110) : tier("720P", 102740);
u * d
```

**表达式解读：**

1. **时长 snap 规则**（对齐适配器对 OS 的处理）：

| 入参 duration | 计费秒数 d |
| --- | ---: |
| 未传 / ≤0 | 8（官方默认） |
| 1～6 | 4 |
| 7～10 | 8 |
| ≥11 | 12 |

2. **分辨率分档**：4K > 2K > 1080 > 默认 720P。
3. **单价**：720P=102740，1080P=154110，2K=231233，4K=346712。
4. **示例**：
   - 不传 duration → `102740 × 8 = 821920`
   - `duration=5` → snap 到 4 → `102740 × 4 = 410960`
   - `duration=12` → `102740 × 12 = 1232880`

---

### 任务 21：`hunyuan-video-1.5`

官方价：720P `0.300`，1080P `0.500`，2K `0.750`，4K `1.120` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : d0;
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 153425) : has(s, "2K") || has(s, "2k") ? tier("2K", 102740) : has(s, "1080") ? tier("1080P", 68493) : tier("720P", 41096);
u * d
```

**表达式解读：**

- **时长**：未传/≤0 → 默认按 5 秒计；有传则原样使用（适配器本身不向腾讯发送 Duration）。
- **单价**：720P=41096，1080P=68493，2K=102740，4K=153425（与 Kling 2.5 同档）。
- **示例**：默认 → `41096 × 5 = 205480`

---

### 任务 22：`mingmou-video-1.0`

官方价同 Hunyuan 1.5：720P `0.300`，1080P `0.500`，2K `0.750`，4K `1.120` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : d0;
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 153425) : has(s, "2K") || has(s, "2k") ? tier("2K", 102740) : has(s, "1080") ? tier("1080P", 68493) : tier("720P", 41096);
u * d
```

**表达式解读：** 与任务 21 完全相同。

---

### 任务 23：`pixverse-video-v5.6`

无声价：540P `0.245`，720P `0.315`，1080P `0.525`，2K `0.735`，4K `1.029` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 140959) : has(s, "2K") || has(s, "2k") ? tier("2K", 100685) : has(s, "1080") ? tier("1080P", 71918) : has(s, "540") || has(s, "480") ? tier("540P", 33562) : tier("720P", 43151);
u * d
```

**表达式解读：**

1. **时长**（PixVerse：1–15，默认 5）：`≤0→5`，`<1→1`，`>15→15`。
2. **分辨率匹配顺序**：4K > 2K > 1080 > **540/480** > 默认 **720P**（与适配器 `defaultVideoResolution["PixVerse"]="720p"` 对齐）。
3. **单价**：540P=33562，720P=43151，1080P=71918，2K=100685，4K=140959。
4. **示例**：不传 size → 720P × 5s = `43151 × 5 = 215755`；若传 `540p` → `33562 × 5 = 167810`

---

### 任务 24：`pixverse-video-v6`

无声价：540P `0.205`，720P `0.264`，1080P `0.528`，2K `0.634`，4K `0.760` 元/秒。有声价：540P `0.264`，720P `0.352`，1080P `0.675`，2K `0.810`，4K `0.971` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let audio = param("metadata.audio_generation") == "Enabled";
let u = audio ? (has(s, "4K") || has(s, "4k") ? tier("audio-4K", 133014) : has(s, "2K") || has(s, "2k") ? tier("audio-2K", 110959) : has(s, "1080") ? tier("audio-1080P", 92466) : has(s, "540") || has(s, "480") ? tier("audio-540P", 36164) : tier("audio-720P", 48219)) : (has(s, "4K") || has(s, "4k") ? tier("silent-4K", 104110) : has(s, "2K") || has(s, "2k") ? tier("silent-2K", 86849) : has(s, "1080") ? tier("silent-1080P", 72329) : has(s, "540") || has(s, "480") ? tier("silent-540P", 28082) : tier("silent-720P", 36164));
u * d
```

**表达式解读：**

1. **时长**：同 PixVerse（默认 5，1–15）。
2. **分支 `audio`**：`Enabled` → 有声，否则无声。
3. **分辨率**：4K > 2K > 1080 > **540/480** > 默认 **720P**（与适配器默认对齐）。
4. **单价要点**：无声 720P=`36164`；有声 1080P=`92466`。
5. **示例**：无声 720p × 5s → `36164 × 5 = 180820`；有声 1080p × 8s → `92466 × 8 = 739728`

---

### 任务 25：`pixverse-video-c1`

无声价：540P `0.235`，720P `0.293`，1080P `0.557`，2K `0.669`，4K `0.803` 元/秒。有声价：540P `0.293`，720P `0.381`，1080P `0.704`，2K `0.845`，4K `1.014` 元/秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 0 ? 5 : (d0 < 1 ? 1 : (d0 > 15 ? 15 : d0));
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let audio = param("metadata.audio_generation") == "Enabled";
let u = audio ? (has(s, "4K") || has(s, "4k") ? tier("audio-4K", 138904) : has(s, "2K") || has(s, "2k") ? tier("audio-2K", 115753) : has(s, "1080") ? tier("audio-1080P", 96438) : has(s, "540") || has(s, "480") ? tier("audio-540P", 40137) : tier("audio-720P", 52192)) : (has(s, "4K") || has(s, "4k") ? tier("silent-4K", 110000) : has(s, "2K") || has(s, "2k") ? tier("silent-2K", 91644) : has(s, "1080") ? tier("silent-1080P", 76301) : has(s, "540") || has(s, "480") ? tier("silent-540P", 32192) : tier("silent-720P", 40137));
u * d
```

**表达式解读：**

- 结构同任务 24（有声/无声 × 分辨率）。
- **分辨率**：4K > 2K > 1080 > **540/480** > 默认 **720P**（与适配器默认对齐）。
- **单价要点**：无声 720P=`40137`；有声 1080P=`96438`。
- **示例**：无声 720p × 5s → `40137 × 5 = 200685`；有声 1080p × 8s → `96438 × 8 = 771504`

---

### 任务 26：`hailuo-video-h3`

> ⚠️ **H3 官方价格尚未在腾讯文档中发布，以下表达式暂用 02/2.3 同档价格占位，配置前请先向腾讯确认 H3 正式定价，再更新对应数字。**
>
> 临时占位价：768P `0.330`，1080P `0.580`，2K `0.930`，4K `1.490` 元/秒（与 02/2.3 相同）。时长 snap 到 {6, 10} 秒。

```text
let d0 = param("duration") != nil ? int(param("duration")) : 0;
let d = d0 <= 8 ? 6 : 10;
let s = param("size") != nil ? param("size") : param("metadata.resolution");
let u = has(s, "4K") || has(s, "4k") ? tier("4K", 204110) : has(s, "2K") || has(s, "2k") ? tier("2K", 127397) : has(s, "1080") ? tier("1080P", 79452) : tier("768P", 45205);
u * d
```

**表达式解读：** 计算逻辑与任务 15（`hailuo-video-02`）完全相同；数字为占位，正式定价公布后需整体替换单价表。

---

## 五、配置时注意

- 平台后台选择「表达式/阶梯计费」，把对应模型的整段表达式整体复制进去。
- 如果业务侧不传 `metadata.audio_generation`，有声/无声模型默认按无声价计费。需要精确计费时，调用方应传 `metadata.audio_generation: "Enabled"`。
- 如果调用方传入 `size` 为像素值，例如 `1920x1080`，表达式只能通过字符串包含关系粗略识别 `1080`、`2K`、`4K`。当前后台表达式无法像 Go 代码一样计算短边落档。
- Hunyuan、Mingmou 适配器不向腾讯发送 `Duration`/`Resolution` 字段，表达式按请求意图预扣；任务完成后系统会按实际输出时长（MetaData.Duration）做多退少补，分辨率部分仍按请求档位计。
- 看不懂某条表达式时：先看「三、表达式通用解读」，再看该模型任务下的「表达式解读」——通常只要弄清 **d（秒数）怎么来、有没有 audio/ref 分支、u 命中哪一档**。
