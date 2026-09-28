# 豆包 Seedance 2.5 价格验证测试参数

> **模型正式 ID**：`doubao-seedance-2-5-260628`（结算必须用此 ID；短名 `doubao-seedance-2.5` 仅影响广场展示）  
> **接口**：`POST /v1/video/generations`，轮询 `GET /v1/video/generations/{task_id}`  
> **计费方式**：按上游返回的 **Token 用量** × 有效输入单价（$/1M tokens）结算；后台只配一档基准价，请求时自动叠 `OtherRatios`。

---

## 价格基准（调价对照）

| 项 | 值 |
| --- | --- |
| 后台应填输入价 | **10.0685** `$/1M tokens`（= 官方不含视频·480/720 ×1.05 ÷7.3） |
| 基准人民币摘要 | **¥73.50** / 1M tokens |
| 汇率 | 1 USD = 7.3 CNY |
| `video_input` | `42/70 = 0.6`（metadata.content 含 `video_url` 时） |
| `resolution`（仅 1080p） | `46/42 ≈ 1.095238` |
| 有声/无声 | **不影响** 2.5 单价（无 `silent_video`） |
| 4K | **不支持** |

### 四档有效单价（经销后）

| 档位 | OtherRatios | 有效单价（元 / 1M tokens） |
| --- | --- | --- |
| 不含视频 · 480P/720P | （无 / 仅基准） | **73.50** |
| 含视频输入 · 480P/720P | `video_input=0.6` | **44.10** |
| 不含视频 · 1080P | `resolution≈1.095238` | **≈80.50** |
| 含视频输入 · 1080P | `0.6 × 1.095238` | **48.30** |

### 实扣验算

```
实扣（元）≈ (completion/usage tokens ÷ 1_000_000) × 有效单价（元）
```

核对时看任务日志 / 消费记录里的 **用量 tokens** 与是否打上 `video_input` / `resolution` 即可；绝对金额会随上游 token 波动，**重点验倍率是否进对档**。

---

## 素材占位（自行回填）

| 占位符 | 用途 | 说明 |
| --- | --- | --- |
| `{{REF_VIDEO_URL}}` | 参考视频 / 编辑 / 延长 | 公网可访问 mp4；含人脸须先过人像库，改传 `asset://{asset_id}` |
| `{{REF_IMAGE_URL}}` | 参考图 / 首帧 | 可先用下方风景图；含人脸同理用 `asset://` |
| `{{REF_AUDIO_URL}}` | 参考音频（可选） | 公网 mp3/wav 等 |

**可直接用的图片（无脸场景）：**

| 用途 | URL |
| --- | --- |
| 风景首帧 / 参考图 | `https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg` |

---

## 用例一览

| # | 场景 | 预期档位 | 预期有效单价 |
| --- | --- | ---: | ---: |
| 1 | 文生 · 720P | 不含视频 · 480/720 | ¥73.50 |
| 2 | 文生 · 1080P | 不含视频 · 1080P | ¥80.50 |
| 3 | 文生 · 480P | 不含视频 · 480/720（与 720 同价） | ¥73.50 |
| 4 | 图生首帧 · 720P | 不含视频（仅图不算视频输入） | ¥73.50 |
| 5 | 参考视频 · 720P | 含视频 · 480/720 | ¥44.10 |
| 6 | 参考视频 · 1080P | 含视频 · 1080P | ¥48.30 |
| 7 | 视频编辑 edit · 720P | 含视频 · 480/720 | ¥44.10 |
| 8 | 视频延长 extend · 1080P | 含视频 · 1080P | ¥48.30 |

---

## 1. 文生 · 720P（基准档）

- **计费**：无 `video_url`，`size=720p` → 无 OtherRatios  
- **预期有效单价**：¥73.50 / 1M

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "清晨海岸公路，镜头缓慢推进，海风吹过路旁草木，写实风格",
    "size": "720p",
    "duration": 5,
    "metadata": {
      "ratio": "16:9",
      "generate_audio": true,
      "watermark": false
    }
  }'
```

---

## 2. 文生 · 1080P（分辨率溢价）

- **计费**：`resolution≈1.095238`  
- **预期有效单价**：¥80.50 / 1M（73.50 × 46/42）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "霓虹雨夜，复古摩托车驶过积水路面，赛博朋克风",
    "size": "1080p",
    "duration": 5,
    "metadata": {
      "ratio": "16:9",
      "generate_audio": true
    }
  }'
```

---

## 3. 文生 · 480P（与 720 同价）

- **计费**：480P 无单独 resolution 系数，应与用例 1 同档  
- **预期有效单价**：¥73.50 / 1M

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "一只橘猫趴在阳光窗台上，懒洋洋眯眼",
    "size": "480p",
    "duration": 5,
    "metadata": {
      "ratio": "16:9",
      "generate_audio": true
    }
  }'
```

---

## 4. 图生首帧 · 720P（确认「仅图」不触发视频折扣）

- **计费**：`images` / `image_url` **不算** `video_input`  
- **预期有效单价**：¥73.50 / 1M（若误打成 44.10 则检测逻辑有 bug）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "镜头从远景缓缓推近，画面轻微呼吸感",
    "images": [
      "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"
    ],
    "size": "720p",
    "duration": 5,
    "metadata": {
      "ratio": "adaptive",
      "generate_audio": true
    }
  }'
```

---

## 5. 参考视频 · 720P（含视频折扣）

- **计费**：`video_input=0.6`  
- **预期有效单价**：¥44.10 / 1M  
- **回填**：把 `{{REF_VIDEO_URL}}` 换成你的视频（建议 2–15s 无脸场景片；有脸用 `asset://...`）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "保持视频 1 的运镜节奏，整体色调偏暖，写实风格",
    "size": "720p",
    "duration": 5,
    "metadata": {
      "omni_reference_task_type": "reference",
      "generate_audio": true,
      "content": [
        {
          "type": "video_url",
          "role": "reference_video",
          "video_url": {
            "url": "{{REF_VIDEO_URL}}"
          }
        }
      ]
    }
  }'
```

---

## 6. 参考视频 · 1080P（视频折扣 × 1080 溢价）

- **计费**：`video_input=0.6` × `resolution≈1.095238`  
- **预期有效单价**：¥48.30 / 1M

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "沿用视频 1 的主体运动，补充细腻光影与景深",
    "size": "1080p",
    "duration": 5,
    "metadata": {
      "omni_reference_task_type": "reference",
      "output_format": "mp4",
      "generate_audio": true,
      "content": [
        {
          "type": "image_url",
          "role": "reference_image",
          "image_url": {
            "url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"
          }
        },
        {
          "type": "video_url",
          "role": "reference_video",
          "video_url": {
            "url": "{{REF_VIDEO_URL}}"
          }
        }
      ]
    }
  }'
```

---

## 7. 视频编辑 edit · 720P（含视频）

- **约束**：`omni_reference_task_type=edit`，`duration=-1`，`ratio=adaptive`，须有 `reference_video`  
- **预期有效单价**：¥44.10 / 1M

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "把视频 1 的天空改成晚霞，其余构图与运镜保持不变",
    "size": "720p",
    "duration": -1,
    "metadata": {
      "omni_reference_task_type": "edit",
      "ratio": "adaptive",
      "generate_audio": true,
      "content": [
        {
          "type": "video_url",
          "role": "reference_video",
          "video_url": {
            "url": "{{REF_VIDEO_URL}}"
          }
        }
      ]
    }
  }'
```

---

## 8. 视频延长 extend · 1080P（含视频 × 1080）

- **约束**：`omni_reference_task_type=extend`，`ratio=adaptive`，须有 `reference_video`  
- **预期有效单价**：¥48.30 / 1M

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "在视频 1 结尾自然续写 3 秒，镜头继续向前推进",
    "size": "1080p",
    "duration": 5,
    "metadata": {
      "omni_reference_task_type": "extend",
      "ratio": "adaptive",
      "generate_audio": true,
      "content": [
        {
          "type": "video_url",
          "role": "reference_video",
          "video_url": {
            "url": "{{REF_VIDEO_URL}}"
          }
        }
      ]
    }
  }'
```

---

## 核对清单（跑完后填）

| # | task_id | usage tokens | 是否含 video_input | 是否含 resolution | 有效单价是否对 | 实扣（元） | 备注 |
| --- | --- | ---: | --- | --- | --- | ---: | --- |
| 1 |  |  | 否 | 否 | 73.50? |  |  |
| 2 |  |  | 否 | 是(1080p) | 80.50? |  |  |
| 3 |  |  | 否 | 否 | 73.50? |  |  |
| 4 |  |  | **否**（关键） | 否 | 73.50? |  | 仅图不得打折 |
| 5 |  |  | 是(0.6) | 否 | 44.10? |  | 回填视频 |
| 6 |  |  | 是(0.6) | 是 | 48.30? |  | 回填视频 |
| 7 |  |  | 是(0.6) | 否 | 44.10? |  | edit |
| 8 |  |  | 是(0.6) | 是 | 48.30? |  | extend |

---

## 备注

1. 视频占位统一搜 `{{REF_VIDEO_URL}}` 替换即可；含人脸素材请改 `asset://{asset_id}`。  
2. 用例 4 专门防回归：历史若把「有图」误判成视频输入，会错误打到 ¥44.10。  
3. `generate_audio` 在 2.5 **不改价**；若与有声/无声用例对比，单价应一致。  
4. 价格配置说明见 `docs/渠道模型价格配置说明/价格配置说明.md` §9.2 / §9.5。
