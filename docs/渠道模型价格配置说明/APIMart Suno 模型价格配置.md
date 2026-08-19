# APIMart Suno 模型价格配置

> **渠道类型**：APIMart Suno（后台渠道类型编号 **60**）  
> **上游地址**：`https://apib.ai`  
> **固定汇率**：`1 USD = 7.3 CNY`（平台内部汇率，不随市场波动调整）；APIMart 平台自身汇率 `1 USD = 7 CNY`（固定，与我们无关）  
> **积分换算**：`1 Credit = $0.10 USD`（APIMart 固定，后台填值统一以 USD 为准）  
> **计费方式**：所有 Suno 工具均为**按次计费**（每次调用收固定费用，与版本无关）

---

## 一、配置结论（快速查表）

所有 31 个模型均选择 **按次计费**，后台填写 **ModelPrice**（单位：USD）。

| 模型 ID | 功能 | APIMart 成本（Credits） | 后台填值（ModelPrice / USD） | 对应人民币成本 |
|---------|------|----------------------|--------------------------|------------|
| `suno-music` | 生成音乐（支持全版本） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-lyrics` | 生成歌词 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-aligned-lyrics` | 歌词时间轴 | 0.008 | **0.0008** | ≈ 0.006 元 |
| `suno-bpm` | BPM 分析 | 0.008 | **0.0008** | ≈ 0.006 元 |
| `suno-concat` | 完整歌曲合成 | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-generate-video` | 生成音乐视频 | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-persona` | 创建 Persona | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-upload` | 上传音频 | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-upsample-tags` | 标签增强（同步接口） | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-vox` | 提取 Vox | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-wav` | 导出 WAV | 0.04 | **0.004** | ≈ 0.029 元 |
| `suno-crop` | 裁剪音频 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-fade-in` | 淡入 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-fade-out` | 淡出 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-remove-section` | 删除片段 | 0.08 | **0.008** | ≈ 0.058 元 |
| `suno-sounds` | 音效生成（v5/v5.5） | 0.096 | **0.0096** | ≈ 0.070 元 |
| `suno-create-voice` | 创建语音 | 0.16 | **0.016** | ≈ 0.117 元 |
| `suno-adjust-speed` | 调整速度 | 0.24 | **0.024** | ≈ 0.175 元 |
| `suno-add-instrumental` | 添加伴奏（v5/v5.5） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-add-stem` | 添加音轨（仅 v5.5） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-add-vocals` | 添加人声（v5/v5.5） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-cover` | 风格翻唱（支持全版本） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-extend` | 续写延长（支持全版本） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-mashup` | 生成混搭（支持全版本） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-midi` | 生成 MIDI | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-remaster` | 母带优化（v4.5+/v5/v5.5） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-replace-section` | 段落替换（v4/v4.5+/v5/v5.5） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-sample` | 样本转歌曲（支持全版本） | 0.5 | **0.05** | ≈ 0.365 元 |
| `suno-inspo` | 灵感生成（v4 及以上） | 0.68 | **0.068** | ≈ 0.496 元 |
| `suno-stems` | 分轨提取 | 1.0 | **0.1** | ≈ 0.730 元 |
| `suno-stems-all` | 全量分轨 | 2.4 | **0.24** | ≈ 1.752 元 |

> 上表的"对应人民币成本"为上游成本价（无加成），按 `USD × 7.3` 换算，供运营定价参考。  
> **后台只需填写 ModelPrice（USD列）**，平台自动按内部汇率换算扣费。

---

## 二、换算说明

```
APIMart Credits → USD：Credits × $0.10
USD → 人民币：USD × 7.3（平台固定汇率）
USD → 内部 quota：USD × 500,000（$1 = 500,000 quota）

示例：suno-music
  0.5 Credits × $0.10 = $0.05 USD
  $0.05 × 7.3 = 0.365 元（成本价）
  后台填 ModelPrice = 0.05
```

---

## 三、表达式值速查（如需使用表达式计费）

所有 Suno 工具价格固定（与 version 无关，与请求参数无关），表达式形式为：

```
tier("fixed", 表达式值)
```

其中：**表达式值 = USD单价 × 1,000,000 = Credits × 100,000**

| Credits | USD | 表达式值 | 涉及模型 |
|---------|-----|---------|---------|
| 0.008 | 0.0008 | **800** | `suno-aligned-lyrics`、`suno-bpm` |
| 0.04 | 0.004 | **4,000** | `suno-concat`、`suno-generate-video`、`suno-persona`、`suno-upload`、`suno-upsample-tags`、`suno-vox`、`suno-wav` |
| 0.08 | 0.008 | **8,000** | `suno-lyrics`、`suno-crop`、`suno-fade-in`、`suno-fade-out`、`suno-remove-section` |
| 0.096 | 0.0096 | **9,600** | `suno-sounds` |
| 0.16 | 0.016 | **16,000** | `suno-create-voice` |
| 0.24 | 0.024 | **24,000** | `suno-adjust-speed` |
| 0.5 | 0.05 | **50,000** | `suno-music`、`suno-add-instrumental`、`suno-add-stem`、`suno-add-vocals`、`suno-cover`、`suno-extend`、`suno-mashup`、`suno-midi`、`suno-remaster`、`suno-replace-section`、`suno-sample` |
| 0.68 | 0.068 | **68,000** | `suno-inspo` |
| 1.0 | 0.1 | **100,000** | `suno-stems` |
| 2.4 | 0.24 | **240,000** | `suno-stems-all` |

---

## 四、后台操作步骤

1. 进入 **系统设置 → 模型倍率**
2. 搜索模型 ID（如 `suno-music`）
3. 将计费方式切换为 **按次计费**
4. 在 **ModelPrice** 字段填入对应的 USD 值（见第一节表格"后台填值"列）
5. 保存并重复以上操作完成全部 31 个模型

> **注意**：同一个模型的所有 version（v3.5/v4/v4.5/v5/v5.5 等）共用同一个模型 ID，价格相同，不需要分 version 单独配置。如未来 APIMart 对特定 version 调整价格，届时可通过调整对应模型 ID 的 ModelPrice 统一变更，或联系开发在适配器层按 version 引入差价系数。

---

## 五、价格分组汇总（便于批量核对）

| 成本价格段 | 模型数 | 模型 ID 列表 |
|-----------|-------|------------|
| < 0.01 元 | 2 | `suno-aligned-lyrics`、`suno-bpm` |
| 0.01–0.05 元 | 7 | `suno-concat`、`suno-generate-video`、`suno-persona`、`suno-upload`、`suno-upsample-tags`、`suno-vox`、`suno-wav` |
| 0.05–0.10 元 | 5 | `suno-lyrics`、`suno-crop`、`suno-fade-in`、`suno-fade-out`、`suno-remove-section` |
| 0.10–0.20 元 | 2 | `suno-sounds`、`suno-create-voice` |
| 0.15–0.20 元 | 1 | `suno-adjust-speed` |
| 0.35–0.40 元 | 11 | `suno-music`、`suno-add-instrumental`、`suno-add-stem`、`suno-add-vocals`、`suno-cover`、`suno-extend`、`suno-mashup`、`suno-midi`、`suno-remaster`、`suno-replace-section`、`suno-sample` |
| 0.45–0.55 元 | 1 | `suno-inspo` |
| 0.70–0.75 元 | 1 | `suno-stems` |
| > 1.50 元 | 1 | `suno-stems-all` |
