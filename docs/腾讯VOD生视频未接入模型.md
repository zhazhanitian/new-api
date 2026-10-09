# 腾讯 VOD 生视频未接入模型

对照来源：

- 官方接口：[创建 AIGC 生视频任务](https://cloud.tencent.com/document/product/266/126239)
- 当前代码：`relay/channel/task/tencentvod/video_constants.go`（已接入 25 个版本）
- 官方价格：[云点播按量计费 · AIGC 生视频模型价格](https://cloud.tencent.com/document/product/266/95125)
- 核对日期：2026-09-29

OS、Mingmou、PixVerse 没有新版本。下面 14 个在官方 `ModelVersion` 里，代码还没有。

接入时改 `video_constants.go` 三处：`VideoModelList`、`videoModelNameMap`、`videoModelVersionMap`。`ModelVersion` 按官方原样，区分大小写。内部模型名沿用 `{系列}-video-{版本}`；GV 继续用已有前缀 `veo-video`。

---

## 未接入清单

| 系列 | 上游 ModelName | 上游 ModelVersion | 建议内部模型名 | 官方价格 |
| --- | --- | --- | --- | --- |
| Kling | `Kling` | `3.0-turbo` | `kling-video-3.0-turbo` | 有 |
| Kling | `Kling` | `scene` | `kling-video-scene` | 未单列 |
| Vidu | `Vidu` | `q3-mix` | `vidu-video-q3-mix` | 有 |
| Vidu | `Vidu` | `q3-drama` | `vidu-video-q3-drama` | 有 |
| Vidu | `Vidu` | `q3-ad` | `vidu-video-q3-ad` | 未单列 |
| Vidu | `Vidu` | `scene` | `vidu-video-scene` | 未单列 |
| Vidu | `Vidu` | `avatar-q2-pro` | `vidu-video-avatar-q2-pro` | 未单列 |
| Vidu | `Vidu` | `avatar-q2-turbo` | `vidu-video-avatar-q2-turbo` | 未单列 |
| Vidu | `Vidu` | `lip-sync` | `vidu-video-lip-sync` | 未单列 |
| Hailuo | `Hailuo` | `H3_regen` | `hailuo-video-h3-regen` | 未单列 |
| Hailuo | `Hailuo` | `H3-Max` | `hailuo-video-h3-max` | 未单列 |
| GV | `GV` | `3.1-lite` | `veo-video-3.1-lite` | 有 |
| GV | `GV` | `omni` | `veo-video-omni` | 有 |
| Hunyuan | `Hunyuan` | `3d_2.0` | `hunyuan-video-3d-2.0` | 未单列 |

`Hunyuan` / `3d_2.0` 挂在生视频接口上，和已接入的混元 AI3D（`hunyuan-3d-*`）不是同一套接口。

---

## 已有官方标价

单位：人民币元/秒。空档表示价格页该分辨率没有标价。

### Kling `3.0-turbo`

| 属性 | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: |
| - | 0.800 | 1.000 | 1.200 | 1.440 |

### Vidu `q3-mix`

| 属性 | 480P | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: | ---: |
| 参考生 | - | 0.782 | 0.938 | 1.126 | 1.352 |

### Vidu `q3-drama`

| 属性 | 480P | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: | ---: |
| 参考生 | - | - | 0.875 | 1.050 | 1.260 |

### GV `3.1-lite`

| 属性 | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: |
| 有声 | 0.375 | 0.600 | 0.900 | 1.125 |
| 无声 | 0.225 | 0.375 | 0.600 | 0.750 |

### GV `omni`

输出视频按秒。输入、输出文本另按百万 token 计费。

| 属性 | 720P | 1080P | 2K | 4K |
| --- | ---: | ---: | ---: | ---: |
| 输出视频 | 0.750 | 1.125 | 1.350 | 1.620 |

| 属性 | 单价 |
| --- | --- |
| 输入文本 | 11.25 元/百万 token |
| 输出文本 | 67.5 元/百万 token |
