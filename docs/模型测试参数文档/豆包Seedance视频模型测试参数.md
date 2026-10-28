# 豆包 Seedance 视频模型测试参数（新请求结构）

按 token 结算，没有固定 quota。提交后核对消费记录里的 **OtherRatios**，不要拿生成耗时去对金额。

请求只用外层 `resolution`、`ratio`、`duration`。素材走 `metadata.content`，地址写在 `image_url.url` / `video_url.url`。不要再传 `size`、`images`、`metadata.ratio`。

旧文件 `豆包Seedance2.5模型测试参数.md` 仍是 `size`，不要用。

风景图：`https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg`  
参考视频：跑之前把 `{{REF_VIDEO_URL}}` 换成可访问、尽量不含脸的 mp4。含脸必须用 `asset://{asset_id}`。

后台基准价见 `价格配置说明.md` 第九节。下面的「有效单价」是经销后的人民币 / 1M tokens。

| 模型 | 后台输入价（$/1M） | 不含视频基准 |
| --- | ---: | --- |
| `doubao-seedance-2-5-260628` | 10.0685 | ¥73.50 |
| `doubao-seedance-2-0-260128`、`dreamina-seedance-2-0-260128` | 6.6164 | ¥48.30 |
| `doubao-seedance-2-0-fast-260128`、`dreamina-seedance-2-0-fast-260128` | 5.3219 | ¥38.85 |
| `doubao-seedance-2-0-mini-260615` | 3.3082 | ¥24.15 |
| `doubao-seedance-1-5-pro-251215` | 2.1918 | 有声 ¥16.00 |

---

## 1. Seedance 2.5 文生 720p

无视频、无 1080 系数。有效单价 **¥73.50**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "清晨海岸公路，镜头缓慢推进",
    "duration": 5,
    "resolution": "720p",
    "ratio": "16:9",
    "metadata": { "generate_audio": true, "watermark": false }
  }'
```

## 2. Seedance 2.5 文生 1080p

`resolution ≈ 1.09524`。有效单价 **¥80.50**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "霓虹雨夜，摩托车驶过积水路面",
    "duration": 5,
    "resolution": "1080p",
    "ratio": "16:9",
    "metadata": { "generate_audio": true }
  }'
```

## 3. Seedance 2.5 参考视频 720p

`video_input = 0.6`。只有参考图不能打这个系数。有效单价 **¥44.10**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-5-260628",
    "prompt": "保持参考视频的运镜，色调偏暖",
    "duration": 5,
    "resolution": "720p",
    "ratio": "adaptive",
    "metadata": {
      "generate_audio": true,
      "content": [{
        "type": "video_url",
        "role": "reference_video",
        "video_url": { "url": "{{REF_VIDEO_URL}}" }
      }]
    }
  }'
```

## 4. Seedance 2.0 文生 720p

无系数。有效单价 **¥48.30**。`dreamina-seedance-2-0-260128` 同价，后台若配了就改 model 再跑一条。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-0-260128",
    "prompt": "山间公路，云雾缓慢流动",
    "duration": 5,
    "resolution": "720p",
    "ratio": "16:9",
    "metadata": { "generate_audio": true }
  }'
```

## 5. Seedance 2.0 文生 4k

`resolution ≈ 0.5652`。有效单价 **¥27.30**（48.30 × 26/46）。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-0-260128",
    "prompt": "山间公路，云雾缓慢流动",
    "duration": 5,
    "resolution": "4k",
    "ratio": "16:9",
    "metadata": { "generate_audio": true }
  }'
```

## 6. Seedance 2.0 参考视频 720p

`video_input ≈ 0.6087`。有效单价 **¥29.40**（48.30 × 28/46）。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-0-260128",
    "prompt": "保持参考视频的运镜，色调偏暖",
    "duration": 5,
    "resolution": "720p",
    "ratio": "adaptive",
    "metadata": {
      "generate_audio": true,
      "content": [{
        "type": "video_url",
        "role": "reference_video",
        "video_url": { "url": "{{REF_VIDEO_URL}}" }
      }]
    }
  }'
```

## 7. Seedance 2.0-fast 文生 720p

无系数。有效单价 **¥38.85**。`dreamina-seedance-2-0-fast-260128` 同价。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-0-fast-260128",
    "prompt": "一只猫趴在窗台上",
    "duration": 5,
    "resolution": "720p",
    "ratio": "16:9",
    "metadata": { "generate_audio": true }
  }'
```

## 8. Seedance 2.0-fast 参考视频 720p

`video_input ≈ 0.5946`。有效单价 **¥23.10**（38.85 × 22/37）。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-0-fast-260128",
    "prompt": "保持参考视频的运镜",
    "duration": 5,
    "resolution": "720p",
    "ratio": "adaptive",
    "metadata": {
      "generate_audio": true,
      "content": [{
        "type": "video_url",
        "role": "reference_video",
        "video_url": { "url": "{{REF_VIDEO_URL}}" }
      }]
    }
  }'
```

## 9. Seedance 2.0-mini 文生 720p

有效单价 **¥24.15**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-0-mini-260615",
    "prompt": "桌面上的咖啡杯，热气升起",
    "duration": 5,
    "resolution": "720p",
    "ratio": "1:1",
    "metadata": { "generate_audio": true }
  }'
```

## 10. Seedance 2.0-mini 参考视频 720p

`video_input ≈ 0.6087`。有效单价 **¥14.70**（24.15 × 14/23）。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-2-0-mini-260615",
    "prompt": "保持参考视频的运镜",
    "duration": 5,
    "resolution": "720p",
    "ratio": "adaptive",
    "metadata": {
      "generate_audio": true,
      "content": [{
        "type": "video_url",
        "role": "reference_video",
        "video_url": { "url": "{{REF_VIDEO_URL}}" }
      }]
    }
  }'
```

## 11. Seedance 1.5 Pro 有声

不传 `generate_audio` 按有声，不打折。有效单价 **¥16.00**。分辨率不参与 1.5 调价。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-1-5-pro-251215",
    "prompt": "海边栈道，镜头跟着人物向前",
    "duration": 5,
    "resolution": "720p",
    "ratio": "16:9"
  }'
```

## 12. Seedance 1.5 Pro 无声

`generate_audio: false` → `silent_video = 0.5`。有效单价 **¥8.00**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "doubao-seedance-1-5-pro-251215",
    "prompt": "海边栈道，镜头跟着人物向前",
    "duration": 5,
    "resolution": "720p",
    "ratio": "16:9",
    "metadata": { "generate_audio": false }
  }'
```
