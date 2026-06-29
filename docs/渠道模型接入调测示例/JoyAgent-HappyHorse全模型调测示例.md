# JoyAgent · HappyHorse 全模型调测示例

渠道类型：`JoyAgent`（ChannelType=59）  
底层平台：京东 JoyAgent → 阿里云百炼 HappyHorse 系列  
接入端点：`POST /v1/video/generations`

> 替换说明：
> - `YOUR_API_KEY` → 系统分配的用户 token
> - `BASE_URL` → 服务地址（如 `http://localhost:3000`）
> - 所有 `https://placeholder.example.com/...` 链接替换为真实公网可访问 URL

---

## 一、文生视频（T2V）

### 1-1. happyhorse-1.0-t2v — 最简调用（默认参数）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-t2v",
    "prompt": "一座由硬纸板和瓶盖搭建的微型城市，在夜晚焕发出生机，小灯点缀其间，照亮前路"
  }'
```

### 1-2. happyhorse-1.0-t2v — 指定分辨率 + 时长 + 宽高比

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-t2v",
    "prompt": "一只橘猫慵懒地躺在阳光里打盹，微风轻轻吹动它的胡须",
    "size": "1080p",
    "duration": 8,
    "metadata": {
      "ratio": "16:9"
    }
  }'
```

### 1-3. happyhorse-1.0-t2v — 竖版短视频 + 固定随机种子

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-t2v",
    "prompt": "霓虹灯下的夜市街景，人群熙攘，烟火气十足",
    "size": "720p",
    "duration": 5,
    "metadata": {
      "ratio": "9:16",
      "seed": 42
    }
  }'
```

### 1-4. happyhorse-1.0-t2v — 方形比例（适合社交媒体）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-t2v",
    "prompt": "水墨画风格，山间云雾缭绕，远处有一叶扁舟",
    "size": "1080p",
    "duration": 10,
    "metadata": {
      "ratio": "1:1"
    }
  }'
```

### 1-5. happyhorse-1.0-t2v — 旧版兼容性验证

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-t2v",
    "prompt": "草原上奔跑的马群，远山如黛，蓝天白云",
    "size": "720p",
    "duration": 5
  }'
```

---

## 二、图生视频（I2V，基于首帧）

> 使用 `input_reference` 或 `images[0]` 传入首帧图片 URL；宽高比自动跟随首帧，**不支持 ratio 参数**。

### 2-1. happyhorse-1.0-i2v — input_reference 传图（推荐方式）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-i2v",
    "prompt": "猫咪慢慢抬起头，好奇地望向镜头",
    "input_reference": "https://placeholder.example.com/cat-first-frame.jpg"
  }'
```

### 2-2. happyhorse-1.0-i2v — images 数组传图

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-i2v",
    "prompt": "花瓣随风飘落，画面唯美流动",
    "images": ["https://placeholder.example.com/flower-scene.jpg"]
  }'
```

### 2-3. happyhorse-1.0-i2v — 指定分辨率 + 时长

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-i2v",
    "prompt": "城市夜景中的高楼，灯光逐渐亮起，车流穿梭",
    "input_reference": "https://placeholder.example.com/city-night.jpg",
    "size": "1080p",
    "duration": 8
  }'
```

### 2-4. happyhorse-1.0-i2v — prompt 为空（纯动态效果）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-i2v",
    "prompt": "",
    "input_reference": "https://placeholder.example.com/landscape.jpg",
    "size": "720p",
    "duration": 5
  }'
```

---

## 三、参考生视频（R2V）

> `images` 数组传入 1-9 张参考图，均以 `reference_image` 类型传给阿里云。
> prompt 中可用 `[Image 1]`、`[Image 2]` 等指代对应图片。

### 3-1. happyhorse-1.0-r2v — 单张参考图（最简）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-r2v",
    "prompt": "[Image 1]中的女孩缓缓回头，眼神温柔，微风吹动发丝",
    "images": [
      "https://placeholder.example.com/girl-portrait.jpg"
    ]
  }'
```

### 3-2. happyhorse-1.0-r2v — 多张参考图（角色一致性）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-r2v",
    "prompt": "[Image 1]中的女性，手持[Image 2]中的折扇，缓缓展开，[Image 3]中的流苏耳坠随之摆动",
    "images": [
      "https://placeholder.example.com/girl-portrait.jpg",
      "https://placeholder.example.com/folding-fan.jpg",
      "https://placeholder.example.com/earrings.jpg"
    ],
    "size": "1080p",
    "duration": 8,
    "metadata": {
      "ratio": "16:9"
    }
  }'
```

### 3-3. happyhorse-1.0-r2v — 指定竖版 + 固定 seed

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-r2v",
    "prompt": "[Image 1]中的模特，自信地走上舞台，展示身上的服装",
    "images": [
      "https://placeholder.example.com/model-photo.jpg"
    ],
    "size": "720p",
    "duration": 5,
    "metadata": {
      "ratio": "9:16",
      "seed": 12345
    }
  }'
```

### 3-4. happyhorse-1.0-r2v — 旧版兼容性验证

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-r2v",
    "prompt": "[Image 1]中的人物向前走去，背影渐渐远去",
    "images": [
      "https://placeholder.example.com/person-back.jpg"
    ],
    "size": "720p",
    "duration": 5
  }'
```

---

## 四、视频编辑（Video-Edit）

> `images[0]` 为**待编辑视频 URL**（mp4/mov），`images[1]` 及以后为可选参考图。
> 注意：video-edit **不支持 duration 和 ratio 参数**；输出时长跟随输入视频。

### 4-1. happyhorse-1.0-video-edit — 纯指令编辑（无参考图）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-video-edit",
    "prompt": "将视频风格转换为水墨画风格，保留原有动作",
    "images": [
      "https://placeholder.example.com/original-video.mp4"
    ]
  }'
```

### 4-2. happyhorse-1.0-video-edit — 指令 + 参考图（局部替换）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-video-edit",
    "prompt": "让视频中的人物穿上图片中的条纹毛衣",
    "images": [
      "https://placeholder.example.com/original-video.mp4",
      "https://placeholder.example.com/striped-sweater.jpg"
    ],
    "size": "720p"
  }'
```

### 4-3. happyhorse-1.0-video-edit — 指定分辨率 + 保留原声

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-video-edit",
    "prompt": "将背景替换为雪景，主体人物保持不变",
    "images": [
      "https://placeholder.example.com/original-video.mp4"
    ],
    "size": "1080p",
    "metadata": {
      "audio_setting": "origin"
    }
  }'
```

### 4-4. happyhorse-1.0-video-edit — 多张参考图（最多可传 5 张）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-video-edit",
    "prompt": "参考图片1中的服装样式和图片2中的场景，对视频进行风格化改造",
    "images": [
      "https://placeholder.example.com/original-video.mp4",
      "https://placeholder.example.com/ref-style1.jpg",
      "https://placeholder.example.com/ref-style2.jpg"
    ],
    "size": "720p",
    "metadata": {
      "audio_setting": "auto"
    }
  }'
```

---

## 五、异常/边界用例

### 5-1. 缺少 prompt（video-edit 中 prompt 是必填的，应报错）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-video-edit",
    "prompt": "",
    "images": [
      "https://placeholder.example.com/original-video.mp4"
    ]
  }'
```

### 5-2. video-edit 不传 images（应报错或静默忽略）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-video-edit",
    "prompt": "将视频转为黑白风格"
  }'
```

### 5-3. i2v 不传图片（应报错或静默忽略）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-i2v",
    "prompt": "镜头缓缓推进"
  }'
```

### 5-4. 不合法的 size 值（应被忽略或映射为默认值）

```bash
curl -s -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer $YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "happyhorse-1.0-t2v",
    "prompt": "测试不合法分辨率",
    "size": "2k",
    "duration": 5
  }'
```

---

## 六、任务状态查询

提交成功后，响应中会返回 `id`（公开任务 ID）。使用该 ID 轮询状态：

```bash
curl -s "$BASE_URL/v1/videos/TASK_ID" \
  -H "Authorization: Bearer $YOUR_API_KEY"
```

期望返回结构：

```json
{
  "id": "TASK_ID",
  "status": "processing",   // pending / processing / succeeded / failed
  "model": "happyhorse-1.0-t2v",
  "created_at": 1719000000,
  "metadata": {
    "url": "https://dashscope-result.oss-cn-beijing.aliyuncs.com/xxx.mp4"
  }
}
```

---

## 七、参数速查表

| 模型 | prompt | images / input_reference | size | duration | ratio | audio_setting |
|------|--------|--------------------------|------|----------|-------|---------------|
| 1.1-t2v | 必填 | 不需要 | 可选 | 可选(3-15) | 可选 | ❌ |
| 1.0-t2v | 必填 | 不需要 | 可选 | 可选(3-15) | 可选 | ❌ |
| 1.1-i2v | 可选 | 必填（1张，首帧） | 可选 | 可选(3-15) | ❌ | ❌ |
| 1.1-r2v | 必填 | 必填（1-9张参考图） | 可选 | 可选(3-15) | 可选 | ❌ |
| 1.0-r2v | 必填 | 必填（1-9张参考图） | 可选 | 可选(3-15) | 可选 | ❌ |
| 1.0-video-edit | 必填 | 必填（`[0]`=视频，`[1-5]`=参考图） | 可选 | ❌ | ❌ | 可选 |

**size 可选值：** `480p` / `720p` / `1080p`（默认 1080P）  
**ratio 可选值：** `16:9`（默认） / `9:16` / `1:1` / `4:3` / `3:4` / `4:5` / `5:4` / `9:21` / `21:9`  
**audio_setting 可选值：** `auto`（默认） / `origin`（保留原声）
