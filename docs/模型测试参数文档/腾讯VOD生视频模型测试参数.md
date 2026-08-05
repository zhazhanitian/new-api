# 腾讯 VOD 生视频模型测试参数文档

> ⚠️ **与测试脚本强关联（必读）**
>
> 本文档与自动化测试脚本 **一一对应**，修改任一侧都必须同步另一侧：
>
> | 侧 | 路径 |
> | --- | --- |
> | 测试参数文档（本文件） | `new-api/docs/模型测试参数文档/腾讯VOD生视频模型测试参数.md` |
> | 测试脚本 | `new-api/docs/模型测试参数文档/scripts/tencent_vod_video_test.js` |
> | 测试结果输出 | `new-api/docs/模型测试参数文档/模型测试结果记录文档/{model}.md` |
>
> **同步规则：**
> 1. 新增 / 删除 / 调整用例（序号、模型、请求体、预期扣费、边界用例 EC-xx）时，**必须同步修改脚本内的 `CASES` 数组**
> 2. 脚本内 `seq` 与本文档用例序号应对齐（标准用例 1–47，边界用例 EC-01～EC-13 → seq 48–60）
> 3. 改完后建议先跑 `node tencent_vod_video_test.js --list` 与 `--dry-run --case N` 核对，再 live 执行
> 4. **禁止只改文档不改脚本**（反之亦然），否则实测覆盖与文档预期会漂移

> **说明**
> - `BASE_URL` 替换为实际服务地址，`YOUR_API_KEY` 替换为实际 key
> - 视频生成为异步长任务，接口为 `POST /v1/video/generations`；提交后用 `GET /v1/video/generations/{task_id}` 轮询结果
> - **预期扣费** 按价格表达式计算（汇率 1 USD = 7.3 CNY），最终以腾讯 VOD 实际出账为准
> - 预期金额 = 表达式值 ÷ 1,000,000 USD；quota = 表达式值 ÷ 2（1 USD = 500,000 quota）
> - **脚本测试成本控制**：通过 `scripts/tencent_vod_video_test.js` 执行 live 测试时，**禁止同一测试用例重复提交**（每个用例默认只 POST 一次）；AI 不得擅自反复执行脚本，除非用户明确要求重新执行该用例
>
> **测试用图片素材**（`images` 字段可直接使用）：
>
> | 用途 | URL |
> | --- | --- |
> | 风景图（首帧 / 风格参考） | `https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg` |
> | 人物女生上半身（人物参考 / 图生视频） | `https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg` |
> | 女生头像照（人物特写 / 参考生） | `https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809801856-14qkezq2.jpeg` |

---

## 一、Kling 系列

### 1. `kling-video-1.6` — 最简请求（不传计费变量）

> 适用同价模型：`kling-video-1.6`、`kling-video-2.0`、`kling-video-2.1`

- **计费参数**：不传 `duration`（适配器不发 Duration，API 默认 5 秒）、不传 `size`（适配器发 720P）
- **预期**：54,795 × 5 = **273,975** ≈ **$0.274**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-1.6",
    "prompt": "夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格"
  }'
```

---

### 2. `kling-video-1.6` — 指定时长 + 1080P

- **计费参数**：`duration=10`、`size="1080P"`
- **预期**：95,890 × 10 = **958,900** ≈ **$0.959**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-1.6",
    "prompt": "夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格",
    "duration": 10,
    "size": "1080P"
  }'
```

---

### 3. `kling-video-2.5` — 最简请求

> 对应腾讯官方 Kling 2.5-pro 价格档

- **计费参数**：不传（d=5，720P）
- **预期**：41,096 × 5 = **205,480** ≈ **$0.205**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-2.5",
    "prompt": "霓虹雨夜，一辆复古摩托车驶过倒映着霓虹灯的积水路面，赛博朋克风"
  }'
```

---

### 4. `kling-video-2.5` — 1080P × 10 秒

- **预期**：68,493 × 10 = **684,930** ≈ **$0.685**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-2.5",
    "prompt": "霓虹雨夜，一辆复古摩托车驶过倒映着霓虹灯的积水路面，赛博朋克风",
    "duration": 10,
    "size": "1080P"
  }'
```

---

### 5. `kling-video-2.6` — 无声（默认）

- **计费参数**：不传 `metadata.audio_generation`（无声档）、d=5、720P
- **预期**：41,096 × 5 = **205,480** ≈ **$0.205**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-2.6",
    "prompt": "初春的竹林，微风吹动竹叶，斑驳的阳光穿过竹叶间隙，治愈系"
  }'
```

---

### 6. `kling-video-2.6` — 有声 + 1080P × 5 秒

- **计费参数**：`audio_generation=Enabled`、`size="1080P"`、`duration=5`
- **预期**：136,986 × 5 = **684,930** ≈ **$0.685**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-2.6",
    "prompt": "初春的竹林，微风吹动竹叶，斑驳的阳光穿过竹叶间隙，治愈系",
    "duration": 5,
    "size": "1080P",
    "metadata": {
      "audio_generation": "Enabled"
    }
  }'
```

---

### 7. `kling-video-o1` — 纯文生视频（无参考图）

- **计费参数**：无 `images`（文生档）、d=5、720P
- **预期**：82,192 × 5 = **410,960** ≈ **$0.411**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-o1",
    "prompt": "一位宇航员在月球表面缓步行走，远处地球悬挂在漆黑星空中，史诗级画面"
  }'
```

---

### 8. `kling-video-o1` — 有参考图 + 1080P × 10 秒

- **计费参数**：`images` 有 1 张（参考图档）、`size="1080P"`、`duration=10`
- **预期**：164,384 × 10 = **1,643,840** ≈ **$1.644**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-o1",
    "prompt": "以参考图中人物为主体，生成其在城市广场漫步的视频，电影感镜头",
    "duration": 10,
    "size": "1080P",
    "images": ["https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"]
  }'
```

---

### 9. `kling-video-3.0` — 无声（默认）

- **预期**：82,192 × 5 = **410,960** ≈ **$0.411**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0",
    "prompt": "雪山之巅，一头雄鹰展翅翱翔，镜头随之俯冲，大自然震撼全景"
  }'
```

---

### 10. `kling-video-3.0` — 有声 + 1080P × 5 秒

- **计费参数**：`audio_generation=Enabled`、`size="1080P"`、`duration=5`
- **预期**：164,384 × 5 = **821,920** ≈ **$0.822**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0",
    "prompt": "雪山之巅，一头雄鹰展翅翱翔，镜头随之俯冲，大自然震撼全景",
    "duration": 5,
    "size": "1080P",
    "metadata": {
      "audio_generation": "Enabled"
    }
  }'
```

---

### 11. `kling-video-3.0-omni` — 纯文生视频（无声无参考）

- **计费参数**：无 `images`、无 `audio_generation`（文生无声档）、d=5、720P
- **预期**：82,192 × 5 = **410,960** ≈ **$0.411**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "一对老夫妻在秋日公园的长椅上相依而坐，落叶飘零，温馨治愈"
  }'
```

---

### 12. `kling-video-3.0-omni` — 参考图 + 1080P × 5 秒（参考档）

> 传入图片时表达式走参考档，不区分有声/无声

- **计费参数**：`images` 有 1 张、`size="1080P"`、`duration=5`
- **预期**：164,384 × 5 = **821,920** ≈ **$0.822**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "以参考图风格生成一段优雅的舞蹈视频，芭蕾风格，电影质感",
    "duration": 5,
    "size": "1080P",
    "images": ["https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"]
  }'
```

---

## 二、Vidu 系列

### 13. `vidu-video-q2` — 纯文生视频（无参考图）

- **计费参数**：无 `images`（文生档）、d=5、720P
- **预期**：43,836 × 5 = **219,180** ≈ **$0.219**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2",
    "prompt": "古镇清晨，炊烟袅袅，石板路上行人稀少，水墨江南意境"
  }'
```

---

### 14. `vidu-video-q2` — 参考图 + 1080P × 8 秒（参考生档）

> 传入 `images` 即触发参考生价格（1080P 价格差异较大，务必复核）

- **计费参数**：`images` 有 1 张、`size="1080P"`、`duration=8`
- **预期**：112,329 × 8 = **898,632** ≈ **$0.899**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2",
    "prompt": "以参考图风格生成城市俯瞰视角，从高处缓缓拉近，电影感镜头",
    "duration": 8,
    "size": "1080P",
    "images": ["https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"],
    "metadata": {
      "input_usage": "Reference"
    }
  }'
```

---

### 15. `vidu-video-q2-pro` — 图生视频（默认首帧模式）

> 不传 `input_usage`：单图走首帧（i2v）价格

- **计费参数**：`images` 有 1 张（首帧档）、d=5、720P
- **预期**：47,945 × 5 = **239,725** ≈ **$0.240**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-pro",
    "prompt": "图中场景继续延展，人物缓缓走向远方，光影渐渐变化",
    "images": ["https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"]
  }'
```

---

### 16. `vidu-video-q2-pro` — 参考图模式 + 1080P × 8 秒

> 传 `input_usage: "Reference"` 走参考生价格

- **计费参数**：`images` 有 1 张 + `input_usage="Reference"`、`size="1080P"`、`duration=8`
- **预期**：123,288 × 8 = **986,304** ≈ **$0.986**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-pro",
    "prompt": "以参考图中人物为主角，生成一段在咖啡馆窗边阅读的视频，自然光线",
    "duration": 8,
    "size": "1080P",
    "images": ["https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809801856-14qkezq2.jpeg"],
    "metadata": {
      "input_usage": "Reference"
    }
  }'
```

---

### 17. `vidu-video-q2-turbo` — 最简请求

- **计费参数**：d=5、720P
- **预期**：34,247 × 5 = **171,235** ≈ **$0.171**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-turbo",
    "prompt": "夜晚城市延时摄影，车流如河，星轨缓缓旋转，动感十足"
  }'
```

---

### 18. `vidu-video-q2-turbo` — 1080P × 8 秒

- **预期**：64,384 × 8 = **515,072** ≈ **$0.515**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-turbo",
    "prompt": "夜晚城市延时摄影，车流如河，星轨缓缓旋转，动感十足",
    "duration": 8,
    "size": "1080P"
  }'
```

---

### 19. `vidu-video-q3` — 最简请求

- **计费参数**：d=5、720P
- **预期**：85,616 × 5 = **428,080** ≈ **$0.428**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3",
    "prompt": "一只猫咪在阳光下伸懒腰，毛发细节丰富，超写实风格"
  }'
```

---

### 20. `vidu-video-q3` — 1080P × 8 秒

- **预期**：107,123 × 8 = **856,984** ≈ **$0.857**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3",
    "prompt": "一只猫咪在阳光下伸懒腰，毛发细节丰富，超写实风格",
    "duration": 8,
    "size": "1080P"
  }'
```

---

### 21. `vidu-video-q3-pro` — 最简请求

- **计费参数**：d=5、720P
- **预期**：107,123 × 5 = **535,615** ≈ **$0.536**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3-pro",
    "prompt": "冬日雪原，一只北极狐在雪地上轻盈跳跃，毛发蓬松，写实细腻"
  }'
```

---

### 22. `vidu-video-q3-pro` — 1080P × 8 秒

- **预期**：128,493 × 8 = **1,027,944** ≈ **$1.028**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3-pro",
    "prompt": "冬日雪原，一只北极狐在雪地上轻盈跳跃，毛发蓬松，写实细腻",
    "duration": 8,
    "size": "1080P"
  }'
```

---

### 23. `vidu-video-q3-turbo` — 最简请求

- **计费参数**：d=5、720P
- **预期**：51,370 × 5 = **256,850** ≈ **$0.257**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3-turbo",
    "prompt": "热带雨林深处，瀑布倾泻入碧绿潭水，水雾弥漫，绿意盎然"
  }'
```

---

### 24. `vidu-video-q3-turbo` — 1080P × 8 秒

- **预期**：60,000 × 8 = **480,000** ≈ **$0.480**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q3-turbo",
    "prompt": "热带雨林深处，瀑布倾泻入碧绿潭水，水雾弥漫，绿意盎然",
    "duration": 8,
    "size": "1080P"
  }'
```

---

## 三、Hailuo 系列

> **注意**：Hailuo 时长只支持 6 秒和 10 秒，适配器按"≤8 秒传 6，否则传 10"来 snap。`size` 只影响分辨率（768P 或 1080P），不传默认 768P。

### 25. `hailuo-video-02` — 不传 duration（默认 6 秒 + 768P）

> 适用同价模型：`hailuo-video-02`、`hailuo-video-2.3`

- **预期**：45,205 × 6 = **271,230** ≈ **$0.271**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-02",
    "prompt": "清晨薄雾中的古镇水乡，乌篷船缓缓划过，倒影在水面荡漾"
  }'
```

---

### 26. `hailuo-video-02` — duration=10 + 1080P

- **计费参数**：`duration=10`（≥9 → snap 到 10）、`size="1080P"`
- **预期**：79,452 × 10 = **794,520** ≈ **$0.795**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-02",
    "prompt": "清晨薄雾中的古镇水乡，乌篷船缓缓划过，倒影在水面荡漾",
    "duration": 10,
    "size": "1080P"
  }'
```

---

### 27. `hailuo-video-2.3-fast` — 不传 duration（6 秒 + 768P）

- **预期**：30,822 × 6 = **184,932** ≈ **$0.185**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-2.3-fast",
    "prompt": "樱花满开的公园小径，花瓣随风飘落，慢动作镜头，浪漫唯美"
  }'
```

---

### 28. `hailuo-video-2.3-fast` — duration=10 + 1080P

- **预期**：52,740 × 10 = **527,400** ≈ **$0.527**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-2.3-fast",
    "prompt": "樱花满开的公园小径，花瓣随风飘落，慢动作镜头，浪漫唯美",
    "duration": 10,
    "size": "1080P"
  }'
```

---

### 29. `hailuo-video-h3` — 不传 duration（6 秒 + 768P，占位价格）

> ⚠️ **H3 官方价格未公布，当前配置使用 02/2.3 同档价格占位。以下预期金额在价格确认后需重新核算。**

- **计费参数**：不传（d=6，768P）
- **预期（占位）**：45,205 × 6 = **271,230** ≈ **$0.271**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-h3",
    "prompt": "深秋枫林，一片片红叶在晨风中轻轻飘落，光线穿透树冠，氛围静谧"
  }'
```

---

### 30. `hailuo-video-h3` — duration=10 + 1080P（占位价格）

- **计费参数**：`duration=10`、`size="1080P"`
- **预期（占位）**：79,452 × 10 = **794,520** ≈ **$0.795**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hailuo-video-h3",
    "prompt": "深秋枫林，一片片红叶在晨风中轻轻飘落，光线穿透树冠，氛围静谧",
    "duration": 10,
    "size": "1080P"
  }'
```

---

## 四、Google Veo 系列

> **注意**：GV 系列时长固定 8 秒，适配器始终发 Duration=8，`duration` 字段填写无效。`size` 支持 "720P" / "1080P"，不传默认 720P。

### 31. `veo-video-3.1` — 无声（默认）

- **计费参数**：不传 `audio_generation`（无声档）、d=8 固定
- **预期**：205,479 × 8 = **1,643,832** ≈ **$1.644**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-video-3.1",
    "prompt": "深海中一群荧光水母缓缓漂浮，蓝色光芒在黑暗中流动，超写实"
  }'
```

---

### 32. `veo-video-3.1` — 有声

- **计费参数**：`audio_generation=Enabled`、d=8 固定
- **预期**：410,959 × 8 = **3,287,672** ≈ **$3.288**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-video-3.1",
    "prompt": "深海中一群荧光水母缓缓漂浮，蓝色光芒在黑暗中流动，超写实",
    "metadata": {
      "audio_generation": "Enabled"
    }
  }'
```

---

### 33. `veo-video-3.1-fast` — 无声（默认）

- **预期**：102,740 × 8 = **821,920** ≈ **$0.822**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-video-3.1-fast",
    "prompt": "沙漠中的骆驼商队，夕阳染红天边，长长的影子拉伸在金色沙丘上"
  }'
```

---

### 34. `veo-video-3.1-fast` — 有声

- **预期**：154,110 × 8 = **1,232,880** ≈ **$1.233**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "veo-video-3.1-fast",
    "prompt": "沙漠中的骆驼商队，夕阳染红天边，长长的影子拉伸在金色沙丘上",
    "metadata": {
      "audio_generation": "Enabled"
    }
  }'
```

---

## 五、OpenAI Sora

> **注意**：OS (Sora) 时长 snap 到 {4, 8, 12} 秒；适配器只支持 720P（发 2K/4K 也会被 snap 到 720P）。  
> 不传 `duration`：API 用官方默认 8 秒；传 1–6 → 4 秒；传 7–10 → 8 秒；传 11+ → 12 秒。

### 35. `sora-video-2.0` — 不传 duration（API 默认 8 秒）

- **计费参数**：d=8（默认）、720P
- **预期**：102,740 × 8 = **821,920** ≈ **$0.822**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "宇宙大爆炸瞬间，粒子与光芒向四周爆发，史诗级科幻视觉"
  }'
```

---

### 36. `sora-video-2.0` — duration=5（snap 到 4 秒）

- **计费参数**：d0=5 → d=4（snap），720P
- **预期**：102,740 × 4 = **410,960** ≈ **$0.411**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "宇宙大爆炸瞬间，粒子与光芒向四周爆发，史诗级科幻视觉",
    "duration": 5
  }'
```

---

### 37. `sora-video-2.0` — duration=12（最长档）

- **计费参数**：d=12（snap），720P
- **预期**：102,740 × 12 = **1,232,880** ≈ **$1.233**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "宇宙大爆炸瞬间，粒子与光芒向四周爆发，史诗级科幻视觉",
    "duration": 12
  }'
```

---

## 六、Hunyuan + Mingmou

> **注意**：适配器不向腾讯发送 Duration / Resolution 字段（模型无官方文档参数），表达式按请求意图预扣。上线后建议用腾讯账单抽样核对实际秒数和分辨率。

### 38. `hunyuan-video-1.5` — 不传 duration/size（按 5 秒 720P 预扣）

- **预期**：41,096 × 5 = **205,480** ≈ **$0.205**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hunyuan-video-1.5",
    "prompt": "长城在云雾中若隐若现，晨光初升，壮阔苍茫的历史感"
  }'
```

---

### 39. `hunyuan-video-1.5` — 显式指定 duration=10 + size=1080P

- **预期**：68,493 × 10 = **684,930** ≈ **$0.685**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "hunyuan-video-1.5",
    "prompt": "长城在云雾中若隐若现，晨光初升，壮阔苍茫的历史感",
    "duration": 10,
    "size": "1080P"
  }'
```

---

### 40. `mingmou-video-1.0` — 不传 duration/size（按 5 秒 720P 预扣）

- **预期**：41,096 × 5 = **205,480** ≈ **$0.205**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "mingmou-video-1.0",
    "prompt": "川西高原牧场，牦牛群在金黄草甸上漫步，远山皑皑白雪，藏区风情"
  }'
```

---

### 41. `mingmou-video-1.0` — duration=10 + size=1080P

- **预期**：68,493 × 10 = **684,930** ≈ **$0.685**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "mingmou-video-1.0",
    "prompt": "川西高原牧场，牦牛群在金黄草甸上漫步，远山皑皑白雪，藏区风情",
    "duration": 10,
    "size": "1080P"
  }'
```

---

## 七、PixVerse 系列

> **注意**：PixVerse 支持 540p/720p/1080p/2k/4k，`size` 传分辨率字符串（大小写均可，如 "720P" 或 "720p"）；不传默认 720p。时长范围 1-15 秒，不传默认 5 秒。

### 42. `pixverse-video-v5.6` — 最简请求（无声，5 秒 720p）

> V5.6 只有无声价格

- **预期**：43,151 × 5 = **215,755** ≈ **$0.216**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-v5.6",
    "prompt": "赛博朋克城市的俯瞰夜景，霓虹灯倒映在潮湿的街道，飞行器穿梭其间"
  }'
```

---

### 43. `pixverse-video-v5.6` — 1080p × 10 秒

- **预期**：71,918 × 10 = **719,180** ≈ **$0.719**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-v5.6",
    "prompt": "赛博朋克城市的俯瞰夜景，霓虹灯倒映在潮湿的街道，飞行器穿梭其间",
    "duration": 10,
    "size": "1080P"
  }'
```

---

### 44. `pixverse-video-v6` — 无声（默认，5 秒 720p）

- **预期**：36,164 × 5 = **180,820** ≈ **$0.181**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-v6",
    "prompt": "极简主义建筑内部，阳光穿过百叶窗投下平行光影，几何美感"
  }'
```

---

### 45. `pixverse-video-v6` — 有声 + 1080p × 8 秒

- **计费参数**：`audio_generation=Enabled`、`size="1080P"`、`duration=8`
- **预期**：92,466 × 8 = **739,728** ≈ **$0.740**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-v6",
    "prompt": "极简主义建筑内部，阳光穿过百叶窗投下平行光影，几何美感",
    "duration": 8,
    "size": "1080P",
    "metadata": {
      "audio_generation": "Enabled"
    }
  }'
```

---

### 46. `pixverse-video-c1` — 无声（默认，5 秒 720p）

- **预期**：40,137 × 5 = **200,685** ≈ **$0.201**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-c1",
    "prompt": "节日烟花在夜空中绽放，五彩斑斓，倒映在湖面上，庆典氛围"
  }'
```

---

### 47. `pixverse-video-c1` — 有声 + 1080p × 8 秒

- **计费参数**：`audio_generation=Enabled`、`size="1080P"`、`duration=8`
- **预期**：96,438 × 8 = **771,504** ≈ **$0.772**

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "pixverse-video-c1",
    "prompt": "节日烟花在夜空中绽放，五彩斑斓，倒映在湖面上，庆典氛围",
    "duration": 8,
    "size": "1080P",
    "metadata": {
      "audio_generation": "Enabled"
    }
  }'
```

---

## 八、边界与异常参数覆盖测试

> 选取 3 个计费变量最多的模型，验证非常规入参下的表达式计费行为。
>
> - `kling-video-3.0-omni`：5 个计费变量（duration / size / audio / ref / input_usage）
> - `vidu-video-q2-pro`：`input_usage` 的存在与否切换 ref/i2v 两种定价
> - `sora-video-2.0`：duration snap 到 {4, 8, 12} 的边界行为
>
> **脚本中对应序号：EC-01～EC-13（seq 48～60），标签 `edge`。**

---

### EC-01. `kling-video-3.0-omni` — duration=-1（负数 → clamp 到 5s）

- **测试目的**：表达式 `d0 ≤ 0 → d=5`，验证负数被正确 clamp，按 5s 计费。
- **预期**：silent-720P  82,192 × 5 = **410,960**（quota ≈ 205,480）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感",
    "duration": -1,
    "size": "720P"
  }'
```

---

### EC-02. `kling-video-3.0-omni` — size="random_xyz"（无效分辨率 → fallback 720P）

- **测试目的**：无效 size 无法命中任何 tier，表达式 fallback `silent-720P`。
- **预期**：silent-720P  82,192 × 5 = **410,960**（quota ≈ 205,480）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感",
    "duration": 5,
    "size": "random_xyz"
  }'
```

---

### EC-03. `kling-video-3.0-omni` — audio_generation="yes"（错误值，非 "Enabled"）

- **测试目的**：表达式检查 `== "Enabled"`，"yes" 不匹配 → audio=false → 按无声计费。
- **预期**：silent-720P  82,192 × 5 = **410,960**（quota ≈ 205,480）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感",
    "duration": 5,
    "size": "720P",
    "metadata": { "audio_generation": "yes" }
  }'
```

---

### EC-04. `kling-video-3.0-omni` — images=[]（空数组 → ref=false）

- **测试目的**：`images.#` = 0 → `ref=false` → 无参考图价格。
- **预期**：silent-720P  82,192 × 5 = **410,960**（quota ≈ 205,480）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "黄昏草原，骏马奔腾，夕阳将大地染成橙红，史诗感",
    "duration": 5,
    "size": "720P",
    "images": []
  }'
```

---

### EC-05. `kling-video-3.0-omni` — prompt 为空字符串（预期 400）

- **测试目的**：验证服务端对空 prompt 的校验；请求被拒绝时不应产生扣费。
- **预期 HTTP**：`400`；**预期扣费**：0 Quota

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "kling-video-3.0-omni",
    "prompt": "",
    "duration": 5,
    "size": "720P"
  }'
```

---

### EC-06. `vidu-video-q2-pro` — 有图但不传 input_usage（走首帧/i2v 价格）

- **测试目的**：仅有 images、不传 `input_usage` 时 `ref=false` → i2v 定价。
- **预期**：i2v-720P  47,945 × 5 = **239,725**（quota ≈ 119,863）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-pro",
    "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
    "duration": 5,
    "size": "720P",
    "images": ["https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"]
  }'
```

---

### EC-07. `vidu-video-q2-pro` — input_usage=Reference 但不传 images

- **测试目的**：表达式仍按 `input_usage=="Reference"` 判定 `ref=true` → ref 定价；腾讯 API 实际行为待观察（可能拒绝或按文生处理）。
- **预期（按表达式）**：ref-720P  47,945 × 5 = **239,725**（quota ≈ 119,863）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-pro",
    "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
    "duration": 5,
    "size": "720P",
    "metadata": { "input_usage": "Reference" }
  }'
```

---

### EC-08. `vidu-video-q2-pro` — duration=0（→ 按默认 5s 计费）

- **测试目的**：`d0=0 ≤ 0 → d=5`，验证 duration=0 与不传等价。
- **预期**：i2v-720P  47,945 × 5 = **239,725**（quota ≈ 119,863）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-pro",
    "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
    "duration": 0,
    "size": "720P",
    "images": ["https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"]
  }'
```

---

### EC-09. `vidu-video-q2-pro` — size="2K"（i2v-2K 档）

- **测试目的**：高分辨率档计费验证；i2v-2K 单价为 136,986。
- **预期**：i2v-2K  136,986 × 5 = **684,930**（quota ≈ 342,465）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "vidu-video-q2-pro",
    "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
    "duration": 5,
    "size": "2K",
    "images": ["https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"]
  }'
```

---

### EC-10. `sora-video-2.0` — duration=6（≤6 → snap 到 4s）

- **测试目的**：snap 下界，`d0=6≤6 → d=4`。
- **预期**：720P  102,740 × 4 = **410,960**（quota ≈ 205,480）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
    "duration": 6
  }'
```

---

### EC-11. `sora-video-2.0` — duration=7（6 < d ≤ 10 → snap 到 8s）

- **测试目的**：snap 中间档，`d0=7 → d=8`。
- **预期**：720P  102,740 × 8 = **821,920**（quota ≈ 410,960）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
    "duration": 7
  }'
```

---

### EC-12. `sora-video-2.0` — duration=11（>10 → snap 到 12s）

- **测试目的**：snap 上界，`d0=11>10 → d=12`。
- **预期**：720P  102,740 × 12 = **1,232,880**（quota ≈ 616,440）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
    "duration": 11
  }'
```

---

### EC-13. `sora-video-2.0` — size="1080P"（表达式按 1080P 计，适配器可能降级）

- **测试目的**：表达式走 1080P 档（154,110/s），但 OS(Sora) 适配器仅支持 720P，实际向腾讯发 720P 请求。观察实际扣费与表达式预期是否一致，以确认适配器与表达式是否匹配。
- **预期（按表达式）**：1080P  154,110 × 8 = **1,232,880**（quota ≈ 616,440）

```bash
curl -X POST "$BASE_URL/v1/video/generations" \
  -H "Authorization: Bearer YOUR_API_KEY" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "sora-video-2.0",
    "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
    "size": "1080P"
  }'
```

---

## 附：任务结果查询

```bash
curl -X GET "$BASE_URL/v1/video/generations/{task_id}" \
  -H "Authorization: Bearer YOUR_API_KEY"
```

返回 `status` 字段：

- `queued` — 排队中
- `inProgress` — 生成中（`progress` 字段显示百分比）
- `succeeded` — 完成，视频 URL 在 `metadata.url`
- `failed` — 失败，查看 `error` 字段

---

## 附：预期扣费汇总

> 换算：`预期 RMB = 预期表达式值 ÷ 1,000,000 × 7.3`（与价格配置文档汇率一致）

| 序号 | 模型 | 计费要素 | 预期表达式值 | 预期 USD | 预期 RMB |
|---|---|---|---|---|---|
| 1 | kling-video-1.6 | 720P × 5s | 273,975 | $0.274 | ¥2.00 |
| 2 | kling-video-1.6 | 1080P × 10s | 958,900 | $0.959 | ¥7.00 |
| 3 | kling-video-2.5 | 720P × 5s | 205,480 | $0.205 | ¥1.50 |
| 4 | kling-video-2.5 | 1080P × 10s | 684,930 | $0.685 | ¥5.00 |
| 5 | kling-video-2.6 | 无声 720P × 5s | 205,480 | $0.205 | ¥1.50 |
| 6 | kling-video-2.6 | 有声 1080P × 5s | 684,930 | $0.685 | ¥5.00 |
| 7 | kling-video-o1 | 无参考 720P × 5s | 410,960 | $0.411 | ¥3.00 |
| 8 | kling-video-o1 | 有参考 1080P × 10s | 1,643,840 | $1.644 | ¥12.00 |
| 9 | kling-video-3.0 | 无声 720P × 5s | 410,960 | $0.411 | ¥3.00 |
| 10 | kling-video-3.0 | 有声 1080P × 5s | 821,920 | $0.822 | ¥6.00 |
| 11 | kling-video-3.0-omni | 无声无参考 720P × 5s | 410,960 | $0.411 | ¥3.00 |
| 12 | kling-video-3.0-omni | 有参考 1080P × 5s | 821,920 | $0.822 | ¥6.00 |
| 13 | vidu-video-q2 | 文生 720P × 5s | 219,180 | $0.219 | ¥1.60 |
| 14 | vidu-video-q2 | 参考生 1080P × 8s | 898,632 | $0.899 | ¥6.56 |
| 15 | vidu-video-q2-pro | 首帧 720P × 5s | 239,725 | $0.240 | ¥1.75 |
| 16 | vidu-video-q2-pro | 参考生 1080P × 8s | 986,304 | $0.986 | ¥7.20 |
| 17 | vidu-video-q2-turbo | 720P × 5s | 171,235 | $0.171 | ¥1.25 |
| 18 | vidu-video-q2-turbo | 1080P × 8s | 515,072 | $0.515 | ¥3.76 |
| 19 | vidu-video-q3 | 720P × 5s | 428,080 | $0.428 | ¥3.12 |
| 20 | vidu-video-q3 | 1080P × 8s | 856,984 | $0.857 | ¥6.26 |
| 21 | vidu-video-q3-pro | 720P × 5s | 535,615 | $0.536 | ¥3.91 |
| 22 | vidu-video-q3-pro | 1080P × 8s | 1,027,944 | $1.028 | ¥7.50 |
| 23 | vidu-video-q3-turbo | 720P × 5s | 256,850 | $0.257 | ¥1.88 |
| 24 | vidu-video-q3-turbo | 1080P × 8s | 480,000 | $0.480 | ¥3.50 |
| 25 | hailuo-video-02 | 768P × 6s | 271,230 | $0.271 | ¥1.98 |
| 26 | hailuo-video-02 | 1080P × 10s | 794,520 | $0.795 | ¥5.80 |
| 27 | hailuo-video-2.3-fast | 768P × 6s | 184,932 | $0.185 | ¥1.35 |
| 28 | hailuo-video-2.3-fast | 1080P × 10s | 527,400 | $0.527 | ¥3.85 |
| 29 | hailuo-video-h3 ⚠️ | 768P × 6s（占位） | 271,230 | $0.271 | ¥1.98 |
| 30 | hailuo-video-h3 ⚠️ | 1080P × 10s（占位） | 794,520 | $0.795 | ¥5.80 |
| 31 | veo-video-3.1 | 无声 720P × 8s | 1,643,832 | $1.644 | ¥12.00 |
| 32 | veo-video-3.1 | 有声 720P × 8s | 3,287,672 | $3.288 | ¥24.00 |
| 33 | veo-video-3.1-fast | 无声 720P × 8s | 821,920 | $0.822 | ¥6.00 |
| 34 | veo-video-3.1-fast | 有声 720P × 8s | 1,232,880 | $1.233 | ¥9.00 |
| 35 | sora-video-2.0 | 720P × 8s (默认) | 821,920 | $0.822 | ¥6.00 |
| 36 | sora-video-2.0 | 720P × 4s (dur=5→snap4) | 410,960 | $0.411 | ¥3.00 |
| 37 | sora-video-2.0 | 720P × 12s | 1,232,880 | $1.233 | ¥9.00 |
| 38 | hunyuan-video-1.5 | 720P × 5s | 205,480 | $0.205 | ¥1.50 |
| 39 | hunyuan-video-1.5 | 1080P × 10s | 684,930 | $0.685 | ¥5.00 |
| 40 | mingmou-video-1.0 | 720P × 5s | 205,480 | $0.205 | ¥1.50 |
| 41 | mingmou-video-1.0 | 1080P × 10s | 684,930 | $0.685 | ¥5.00 |
| 42 | pixverse-video-v5.6 | 720p × 5s | 215,755 | $0.216 | ¥1.58 |
| 43 | pixverse-video-v5.6 | 1080p × 10s | 719,180 | $0.719 | ¥5.25 |
| 44 | pixverse-video-v6 | 无声 720p × 5s | 180,820 | $0.181 | ¥1.32 |
| 45 | pixverse-video-v6 | 有声 1080p × 8s | 739,728 | $0.740 | ¥5.40 |
| 46 | pixverse-video-c1 | 无声 720p × 5s | 200,685 | $0.201 | ¥1.47 |
| 47 | pixverse-video-c1 | 有声 1080p × 8s | 771,504 | $0.772 | ¥5.63 |

> ⚠️ 29-30 行（`hailuo-video-h3`）使用占位价格，H3 官方定价公布后需重新核算并更新价格配置文档。
