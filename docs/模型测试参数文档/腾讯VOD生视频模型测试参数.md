# 腾讯 VOD 生视频模型测试参数（新请求结构）

> 旧版用 `size`、`images` 的用例已作废。本文件只覆盖当前对外字段：外层 `resolution`、`ratio`、`duration`，素材只走 `metadata.content`。
>
> 脚本 `scripts/tencent_vod_video_test.js` 里的 `CASES` 还是旧请求。本文确认后再改脚本，然后才能 live 跑。
>
> 预期按 `腾讯VOD生视频模型价格配置.md` 的表达式。表达式值是计价结果；quota = `round(表达式值 / 2)`。金额 = 表达式值 ÷ 1,000,000 美元。

**素材**

| 用途 | URL |
| --- | --- |
| 风景图 | `https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg` |
| 参考视频 | 跑之前把 `{{REF_VIDEO_URL}}` 换成可访问的 mp4 |

同价模型不重复开用例：`kling-video-2.0`、`kling-video-2.1` 与 1.6 同价；`hailuo-video-2.3` 与 02 同价，只留 1 条确认模型已配上。

---

## 1. `kling-video-1.6` 默认

不传 `duration`、`resolution`。计 5 秒、720P。

表达式值 **273975**，quota **136988**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-1.6",
    "prompt": "夏日海滩，海浪轻拍礁石，写实风格",
    "ratio": "16:9"
  }'
```

## 2. `kling-video-1.6` 1080P × 10 秒

表达式值 **958900**，quota **479450**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-1.6",
    "prompt": "夏日海滩，海浪轻拍礁石，写实风格",
    "duration": 10,
    "resolution": "1080P",
    "ratio": "16:9"
  }'
```

## 3. `kling-video-2.5` 默认

5 秒、720P。表达式值 **205480**，quota **102740**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-2.5",
    "prompt": "霓虹雨夜，摩托车驶过积水路面",
    "ratio": "16:9"
  }'
```

## 4. `kling-video-2.5` 1080P × 5 秒

表达式值 **342465**，quota **171233**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-2.5",
    "prompt": "霓虹雨夜，摩托车驶过积水路面",
    "duration": 5,
    "resolution": "1080P",
    "ratio": "16:9"
  }'
```

## 5. `kling-video-2.6` 无声默认

不传 `audio_generation`。5 秒、720P。表达式值 **205480**，quota **102740**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-2.6",
    "prompt": "竹林微风，阳光穿过竹叶",
    "ratio": "16:9"
  }'
```

## 6. `kling-video-2.6` 有声 1080P × 5 秒

表达式值 **684930**，quota **342465**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-2.6",
    "prompt": "竹林微风，阳光穿过竹叶",
    "duration": 5,
    "resolution": "1080P",
    "ratio": "16:9",
    "metadata": { "audio_generation": "Enabled" }
  }'
```

## 7. `kling-video-o1` 文生

无参考视频。5 秒、720P。表达式值 **410960**，quota **205480**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-o1",
    "prompt": "城市天际线从白天过渡到黄昏",
    "ratio": "16:9"
  }'
```

## 8. `kling-video-o1` 只有参考图

参考图不能算有参考视频，应与用例 7 同价。表达式值 **410960**，quota **205480**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-o1",
    "prompt": "保持画面风格，镜头缓慢推进",
    "metadata": {
      "content": [{
        "type": "image_url",
        "role": "reference_image",
        "image_url": { "url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg" }
      }]
    }
  }'
```

## 9. `kling-video-o1` 参考视频 1080P × 5 秒

表达式值 **821920**，quota **410960**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-o1",
    "prompt": "沿用参考视频的运镜，色调偏暖",
    "duration": 5,
    "resolution": "1080P",
    "metadata": {
      "content": [{
        "type": "video_url",
        "role": "reference_video",
        "video_url": { "url": "{{REF_VIDEO_URL}}" }
      }]
    }
  }'
```

## 10. `kling-video-3.0` 无声默认

表达式值 **410960**，quota **205480**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0",
    "prompt": "雨后街道，霓虹倒映在地面积水上",
    "ratio": "16:9"
  }'
```

## 11. `kling-video-3.0` 有声 1080P × 5 秒

表达式值 **821920**，quota **410960**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0",
    "prompt": "雨后街道，霓虹倒映在地面积水上",
    "duration": 5,
    "resolution": "1080P",
    "ratio": "16:9",
    "metadata": { "audio_generation": "Enabled" }
  }'
```

## 12. `kling-video-3.0-omni` 无参考视频、无声

表达式值 **410960**，quota **205480**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "一只猫从窗台跳下，慢动作",
    "ratio": "16:9"
  }'
```

## 13. `kling-video-3.0-omni` 无参考视频、有声

720P × 5 秒。表达式值 **547945**，quota **273973**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "一只猫从窗台跳下，慢动作",
    "duration": 5,
    "resolution": "720P",
    "ratio": "16:9",
    "metadata": { "audio_generation": "Enabled" }
  }'
```

## 14. `kling-video-3.0-omni` 参考视频 1080P × 5 秒

有参考视频时不再分有声无声。表达式值 **821920**，quota **410960**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "沿用参考视频的主体动作",
    "duration": 5,
    "resolution": "1080P",
    "metadata": {
      "audio_generation": "Enabled",
      "content": [{
        "type": "video_url",
        "role": "reference_video",
        "video_url": { "url": "{{REF_VIDEO_URL}}" }
      }]
    }
  }'
```

## 15. `vidu-video-q2` 文生

表达式值 **219180**，quota **109590**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2",
    "prompt": "山间公路，云雾缓慢流动",
    "ratio": "16:9"
  }'
```

## 16. `vidu-video-q2` 不写 role 的参考图，1080P × 8 秒

不写 `role` 应补成参考图，走参考生。表达式值 **898632**，quota **449316**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2",
    "prompt": "人物转头微笑，背景轻微虚化",
    "duration": 8,
    "resolution": "1080P",
    "metadata": {
      "content": [{
        "type": "image_url",
        "image_url": { "url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg" }
      }]
    }
  }'
```

## 17. `vidu-video-q2-pro` 首帧

只有首帧，走图生价，不是参考生。5 秒、720P。表达式值 **239725**，quota **119863**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-pro",
    "prompt": "人物从座位上站起来",
    "metadata": {
      "content": [{
        "type": "image_url",
        "role": "first_frame",
        "image_url": { "url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg" }
      }]
    }
  }'
```

## 18. `vidu-video-q2-pro` 参考图 1080P × 8 秒

表达式值 **986304**，quota **493152**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-pro",
    "prompt": "保持人物外观，镜头缓慢环绕",
    "duration": 8,
    "resolution": "1080P",
    "metadata": {
      "content": [{
        "type": "image_url",
        "role": "reference_image",
        "image_url": { "url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg" }
      }]
    }
  }'
```

## 19. `vidu-video-q2-turbo` 默认

表达式值 **171235**，quota **85618**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-turbo",
    "prompt": "海面日落，镜头平稳横移",
    "ratio": "16:9"
  }'
```

## 20. `vidu-video-q2-turbo` 1080P × 8 秒

表达式值 **515072**，quota **257536**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-turbo",
    "prompt": "海面日落，镜头平稳横移",
    "duration": 8,
    "resolution": "1080P",
    "ratio": "16:9"
  }'
```

## 21. `vidu-video-q3` 默认

表达式值 **428080**，quota **214040**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3",
    "prompt": "森林小径，光斑随风移动",
    "ratio": "16:9"
  }'
```

## 22. `vidu-video-q3` 1080P × 8 秒

表达式值 **856984**，quota **428492**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3",
    "prompt": "森林小径，光斑随风移动",
    "duration": 8,
    "resolution": "1080P",
    "ratio": "16:9"
  }'
```

## 23. `vidu-video-q3-pro` 默认

表达式值 **535615**，quota **267808**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3-pro",
    "prompt": "雪山日出，云层缓慢流动",
    "ratio": "16:9"
  }'
```

## 24. `vidu-video-q3-pro` 1080P × 8 秒

表达式值 **1027944**，quota **513972**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3-pro",
    "prompt": "雪山日出，云层缓慢流动",
    "duration": 8,
    "resolution": "1080P",
    "ratio": "16:9"
  }'
```

## 25. `vidu-video-q3-turbo` 文生默认

表达式值 **256850**，quota **128425**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3-turbo",
    "prompt": "城市夜景，车灯拉出光轨",
    "ratio": "16:9"
  }'
```

## 26. `vidu-video-q3-turbo` 参考图 1080P × 8 秒

表达式值 **444928**，quota **222464**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3-turbo",
    "prompt": "保持参考图风格，镜头缓慢拉远",
    "duration": 8,
    "resolution": "1080P",
    "metadata": {
      "content": [{
        "type": "image_url",
        "role": "reference_image",
        "image_url": { "url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg" }
      }]
    }
  }'
```

## 27. `hailuo-video-02` 默认

不传时长按 6 秒，不传分辨率按 768P。表达式值 **271230**，quota **135615**。`hailuo-video-2.3` 同价。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-02",
    "prompt": "雨中的咖啡馆橱窗，行人撑伞经过"
  }'
```

## 28. `hailuo-video-02` 1080P × 10 秒

表达式值 **794520**，quota **397260**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-02",
    "prompt": "雨中的咖啡馆橱窗，行人撑伞经过",
    "duration": 10,
    "resolution": "1080P"
  }'
```

## 29. `hailuo-video-2.3` 默认

与 02 同表达式。表达式值 **271230**，quota **135615**。只确认这个模型 ID 已配上。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-2.3",
    "prompt": "雨中的咖啡馆橱窗，行人撑伞经过"
  }'
```

## 30. `hailuo-video-2.3-fast` 默认

6 秒、768P。表达式值 **184932**，quota **92466**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-2.3-fast",
    "prompt": "桌面上的咖啡杯，热气缓缓升起"
  }'
```

## 31. `hailuo-video-2.3-fast` 1080P × 10 秒

表达式值 **527400**，quota **263700**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-2.3-fast",
    "prompt": "桌面上的咖啡杯，热气缓缓升起",
    "duration": 10,
    "resolution": "1080P"
  }'
```

## 32. `hailuo-video-h3` 默认

6 秒、768P。表达式值 **410958**，quota **205479**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-h3",
    "prompt": "海边栈道，镜头跟着人物向前走",
    "ratio": "16:9"
  }'
```

## 33. `hailuo-video-h3` 1080P × 10 秒

表达式值 **890410**，quota **445205**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-h3",
    "prompt": "海边栈道，镜头跟着人物向前走",
    "duration": 10,
    "resolution": "1080P",
    "ratio": "16:9"
  }'
```

## 34. `veo-video-3.1` 无声

时长固定 8 秒，720P 与 1080P 同价。表达式值 **1643832**，quota **821916**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-video-3.1",
    "prompt": "沙漠公路上的一辆旧汽车，热浪扭曲远景",
    "resolution": "720P",
    "ratio": "16:9"
  }'
```

## 35. `veo-video-3.1` 有声

表达式值 **3287672**，quota **1643836**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-video-3.1",
    "prompt": "沙漠公路上的一辆旧汽车，热浪扭曲远景",
    "resolution": "1080P",
    "ratio": "16:9",
    "metadata": { "audio_generation": "Enabled" }
  }'
```

## 36. `veo-video-3.1-fast` 无声

表达式值 **821920**，quota **410960**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-video-3.1-fast",
    "prompt": "厨房里的面团被慢慢拉开",
    "resolution": "720P",
    "ratio": "16:9"
  }'
```

## 37. `veo-video-3.1-fast` 有声

表达式值 **1232880**，quota **616440**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-video-3.1-fast",
    "prompt": "厨房里的面团被慢慢拉开",
    "resolution": "720P",
    "ratio": "16:9",
    "metadata": { "audio_generation": "Enabled" }
  }'
```

## 38. `sora-video-2.0` 默认

不传时长按 8 秒、720P。表达式值 **821920**，quota **410960**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "纸船在水面上漂过桥洞",
    "ratio": "16:9"
  }'
```

## 39. `sora-video-2.0` duration=12

12 秒落在最长档。表达式值 **1232880**，quota **616440**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "纸船在水面上漂过桥洞",
    "duration": 12,
    "resolution": "720P",
    "ratio": "16:9"
  }'
```

## 40. `hunyuan-video-1.5` 默认

不向腾讯发时长和分辨率，预扣仍按 5 秒、720P。表达式值 **205480**，quota **102740**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hunyuan-video-1.5",
    "prompt": "江南水乡，乌篷船慢慢划过"
  }'
```

## 41. `hunyuan-video-1.5` 1080P × 10 秒

表达式值 **684930**，quota **342465**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hunyuan-video-1.5",
    "prompt": "江南水乡，乌篷船慢慢划过",
    "duration": 10,
    "resolution": "1080P"
  }'
```

## 42. `mingmou-video-1.0` 默认

与混元 1.5 默认同价。表达式值 **205480**，quota **102740**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "mingmou-video-1.0",
    "prompt": "雪地里的一盏路灯，雪花落下"
  }'
```

## 43. `pixverse-video-v5.6` 默认

无声，5 秒、720P。表达式值 **215755**，quota **107878**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-v5.6",
    "prompt": "一只鸟从枝头飞起",
    "ratio": "16:9"
  }'
```

## 44. `pixverse-video-v5.6` 1080P × 10 秒

表达式值 **719180**，quota **359590**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-v5.6",
    "prompt": "一只鸟从枝头飞起",
    "duration": 10,
    "resolution": "1080P",
    "ratio": "16:9"
  }'
```

## 45. `pixverse-video-v6` 无声默认

表达式值 **180820**，quota **90410**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-v6",
    "prompt": "风吹过麦田，麦浪起伏",
    "ratio": "16:9"
  }'
```

## 46. `pixverse-video-v6` 有声 1080P × 8 秒

表达式值 **739728**，quota **369864**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-v6",
    "prompt": "风吹过麦田，麦浪起伏",
    "duration": 8,
    "resolution": "1080P",
    "ratio": "16:9",
    "metadata": { "audio_generation": "Enabled" }
  }'
```

## 47. `pixverse-video-c1` 无声默认

表达式值 **200685**，quota **100343**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-c1",
    "prompt": "夜晚的码头，渔灯轻轻晃动",
    "ratio": "16:9"
  }'
```

## 48. `pixverse-video-c1` 有声 1080P × 8 秒

表达式值 **771504**，quota **385752**。

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-c1",
    "prompt": "夜晚的码头，渔灯轻轻晃动",
    "duration": 8,
    "resolution": "1080P",
    "ratio": "16:9",
    "metadata": { "audio_generation": "Enabled" }
  }'
```
