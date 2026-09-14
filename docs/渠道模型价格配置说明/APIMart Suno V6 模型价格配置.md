# APIMart Suno V6 模型价格配置

> **渠道类型**：APIMart Suno（后台渠道类型编号 **60**）  
> **上游地址**：`https://api.apimart.ai`  
> **固定汇率**：`1 USD = 7.3 CNY`（平台内部汇率，不随市场波动调整）；APIMart 平台自身汇率 `1 USD = 7 CNY`（固定，与我们无关）  
> **积分换算**：`1 Credit = $0.10 USD`（APIMart 固定，后台填值统一以 USD 为准）  
> **计费方式**：所有 Suno 工具均为**按次计费**；支持 Max 的模型在请求里带 `max_mode=true` 时，由适配器自动按 **2 倍**扣费

旧文档见同目录 [`APIMart Suno 模型价格配置.md`](./APIMart%20Suno%20模型价格配置.md)（v3.5–v5.5，已下架，勿再用于新配置）。

---

## 一、和旧版最大的差别：同一模型、不同参数

V6 价格页里，`v6` / `v6-wild` / `v6-mini` / `custom`（指 `custom_model_id`，不是 `custom: true` 歌词模式）**单价相同**，不需要按版本拆模型、也不需要按版本写表达式。

真正会变价的只有一个请求参数：

| 参数 | 是否影响价格 | 说明 |
|------|------------|------|
| `version`（v6 / v6-wild / v6-mini） | 否 | 同模型同价 |
| `custom_model_id` | 否 | 与公共版本同价；创建模型本身另收 `suno-create-model` |
| `custom: true`（自定义歌词） | 否 | 与计费无关 |
| `variety: "max"` | 否 | **不会**触发 Max 计费 |
| `max_mode: true` | **是，2 倍** | 仅下列 12 个模型支持 |

这和视频/生图模型不一样：Kling、Veo、gpt-image-2 是「分辨率 × 时长 × 有声」多维分档，后台必须用**表达式**读 `param("size")` 等字段。Suno V6 只有 **Max / 非 Max 两档，且刚好是 2 倍**，所以：

**推荐做法（已在代码落地）：后台仍然选「按次计费」，只填普通档 ModelPrice；`max_mode=true` 由适配器 `EstimateBilling` 乘 `2.0`。**

不要把视频模型那套「表达式里再写一遍 max」套过来——会和代码倍率叠乘，变成 4 倍。

---

## 二、配置结论（快速查表）

所有 33 个模型均选择 **按次计费**，后台填写 **ModelPrice**（单位：USD，填**普通档**）。

### 2.1 固定价（无 Max，后台填多少扣多少）

| 模型 ID | 功能 | APIMart 成本（Credits） | 后台填值（ModelPrice / USD） | 对应人民币成本 |
|---------|------|----------------------|--------------------------|------------|
| `suno-aligned-lyrics` | 歌词时间轴 | 0.008 | **0.0008** | ≈ 0.006 元 |
| `suno-bpm` | BPM 分析 | 0.008 | **0.0008** | ≈ 0.006 元 |
| `suno-download` | 下载音频（替代旧 wav） | 0.016 | **0.0016** | ≈ 0.012 元 |
| `suno-concat` | 完整歌曲合成 | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-generate-video` | 生成音乐视频 | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-persona` | 创建 Persona | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-upload` | 上传音频 | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-upsample-tags` | 标签增强 | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-crop` | 裁剪音频 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-fade-in` | 淡入 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-fade-out` | 淡出 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-lyrics` | 生成歌词 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-remove-section` | 删除片段 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-sounds` | 音效生成 | 0.096 | **0.0096** | ≈ 0.070 元 |
| `suno-create-voice` | 创建语音 | 0.16 | **0.016** | ≈ 0.117 元 |
| `suno-adjust-speed` | 调整速度 | 0.24 | **0.024** | ≈ 0.175 元 |
| `suno-midi` | 生成 MIDI | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-remaster` | 母带优化（V6 无版本参数） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-stems` | 分轨提取 | 1.0 | **0.1** | ≈ 0.730 元 |
| `suno-stems-all` | 全量分轨 | 2.4 | **0.24** | ≈ 1.752 元 |
| `suno-create-model` | 创建自定义模型 | 9.6 | **0.96** | ≈ 7.008 元 |

### 2.2 支持 Max 的模型（后台只填普通档，Max 自动 ×2）

| 模型 ID | 功能 | 普通档 Credits / USD | 后台填值 | Max 档实扣 USD | 普通档人民币 | Max 档人民币 |
|---------|------|---------------------|---------|---------------|-----------|-----------|
| `suno-music` | 生成音乐 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-cover` | 风格翻唱 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-extend` | 续写延长 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-mashup` | 生成混搭 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-sample` | 样本转歌曲 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-replace-section` | 段落替换 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-add-instrumental` | 添加伴奏 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-add-stem` | 添加音轨 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-add-vocals` | 添加人声 | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-upload-cover` | 上传并翻唱（新增） | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-upload-extend` | 上传并延伸（新增） | 0.5 / 0.05 | **0.05** | 0.10 | ≈ 0.365 元 | ≈ 0.730 元 |
| `suno-inspo` | 灵感生成 | 0.68 / 0.068 | **0.068** | 0.136 | ≈ 0.496 元 | ≈ 0.993 元 |

> 「对应人民币成本」为上游成本价（无加成），按 `USD × 7.3` 换算，供运营定价参考。  
> **后台只需填写 ModelPrice（USD 列的普通档）**，平台自动按内部汇率换算扣费；Max 请求由代码翻倍。

---

## 三、计费链路（为什么不用表达式）

```
用户请求（max_mode 可有可无）
  → 后台 ModelPrice = 普通档（如 suno-music = 0.05）
  → 适配器 EstimateBilling：仅当该模型在 Max 白名单且 max_mode=true 时，OtherRatios["max_mode"]=2.0
  → 实扣 Quota = 基础额度 × 2（或 × 1）
```

对照视频模型（腾讯 VOD）：同一模型要按 `duration`、`size`、有声/无声分很多档，后台必须写 `param("size")` 表达式。Suno V6 没有分辨率、没有按时长计价，**唯一变量是 2 倍开关**，用代码倍率更简单、也不容易配错。

**禁止同时使用：**

- 后台填「按次计费」普通档 **且** 再配表达式读 `max_mode`  
- 或后台填 Max 档价格（0.10）再指望代码 ×2  

这两种都会多扣。

---

## 四、换算说明

```
APIMart Credits → USD：Credits × $0.10
USD → 人民币：USD × 7.3（平台固定汇率）
USD → 内部 quota：USD × 500,000（$1 = 500,000 quota）

示例：suno-music 普通档
  0.5 Credits × $0.10 = $0.05 USD
  $0.05 × 7.3 = 0.365 元（成本价）
  后台填 ModelPrice = 0.05

示例：suno-music + max_mode=true
  后台仍填 0.05，代码 ×2 → 实扣 $0.10 ≈ 0.730 元
```

---

## 五、后台操作步骤

1. 进入 **系统设置 → 模型倍率**
2. 搜索模型 ID（如 `suno-music`）
3. 将计费方式切换为 **按次计费**
4. 在 **ModelPrice** 填入第二节表格「后台填值」列（一律普通档）
5. 保存并重复，完成全部 **33** 个模型

### 相对旧配置需要改动的项

| 动作 | 模型 ID | 说明 |
|------|---------|------|
| **下架 / 停用** | `suno-vox` | 提取 Vox 已删除 |
| **下架 / 停用** | `suno-wav` | 导出 WAV 已删除，改用 `suno-download` |
| **新增** | `suno-download` | ModelPrice **0.0016** |
| **新增** | `suno-upload-cover` | ModelPrice **0.05** |
| **新增** | `suno-upload-extend` | ModelPrice **0.05** |
| **新增** | `suno-create-model` | ModelPrice **0.96** |
| **价格不变** | 其余旧模型 | 普通档与旧文档相同，可保留原填值 |

旧模型里 `suno-music` 等已经是 `0.05`，**不用改数字**；上线 V6 代码后，带 `max_mode=true` 的请求会自动按 0.10 扣。

---

## 六、价格分组汇总（便于批量核对）

| 成本价格段（普通档人民币） | 模型数 | 模型 ID 列表 |
|-----------|-------|------------|
| < 0.01 元 | 2 | `suno-aligned-lyrics`、`suno-bpm` |
| 0.01–0.05 元 | 6 | `suno-download`、`suno-concat`、`suno-generate-video`、`suno-persona`、`suno-upload`、`suno-upsample-tags` |
| 0.05–0.10 元 | 5 | `suno-lyrics`、`suno-crop`、`suno-fade-in`、`suno-fade-out`、`suno-remove-section` |
| 0.10–0.20 元 | 2 | `suno-sounds`、`suno-create-voice` |
| 0.15–0.20 元 | 1 | `suno-adjust-speed` |
| 0.35–0.40 元 | 13 | `suno-music`、`suno-cover`、`suno-extend`、`suno-mashup`、`suno-sample`、`suno-replace-section`、`suno-add-instrumental`、`suno-add-stem`、`suno-add-vocals`、`suno-upload-cover`、`suno-upload-extend`、`suno-midi`、`suno-remaster` |
| 0.45–0.55 元 | 1 | `suno-inspo` |
| 0.70–0.75 元 | 1 | `suno-stems` |
| > 1.50 元 | 1 | `suno-stems-all` |
| > 7.00 元 | 1 | `suno-create-model` |

---

## 七、仅供对照：若坚持用表达式（不推荐）

仅在**关闭适配器 Max 倍率**之后才能用表达式，否则会 ×4。当前代码会乘 2，**默认不要配表达式**。

若将来改成纯后台表达式，普通档 / Max 档对应值为：

| Credits | USD | 表达式值 | 用途 |
|---------|-----|---------|------|
| 0.008 | 0.0008 | **800** | aligned-lyrics、bpm |
| 0.016 | 0.0016 | **1,600** | download |
| 0.04 | 0.004 | **4,000** | concat、generate-video、persona、upload、upsample-tags |
| 0.08 | 0.008 | **8,000** | lyrics、crop、fade-in、fade-out、remove-section |
| 0.096 | 0.0096 | **9,600** | sounds |
| 0.16 | 0.016 | **16,000** | create-voice |
| 0.24 | 0.024 | **24,000** | adjust-speed |
| 0.5 | 0.05 | **50,000** | music 等普通档 |
| 0.68 | 0.068 | **68,000** | inspo 普通档 |
| 1.0 | 0.10 | **100,000** | music 等 Max 档 |
| 1.36 | 0.136 | **136,000** | inspo Max 档 |
| 2.4 | 0.24 | **240,000** | stems-all |
| 9.6 | 0.96 | **960,000** | create-model |

支持 Max 的模型表达式示例（**当前切勿与按次计费+代码倍率同时启用**）：

```
param("max_mode") == true ? tier("max", 100000) : tier("normal", 50000)
```

`suno-inspo` 把数字换成 `136000` / `68000`。固定价模型继续用 `tier("fixed", 表达式值)`。
