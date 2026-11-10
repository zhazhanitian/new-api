# 28. gpt-image-2.6-sunburst 参考图

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:31:50  |  模型：`gpt-image-2.6-sunburst`  |  档位：预扣 $1

### 调用参数
```json
{
  "model": "gpt-image-2.6-sunburst",
  "prompt": "按参考图风格画一只橘猫",
  "resolution": "1K",
  "ratio": "1:1",
  "quality": "low",
  "n": 1,
  "image": [
    "https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 1K | 分辨率档位 |
| ratio | 1:1 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | low | 渲染精度 |
| n | 1 | 张数 |
| image | 1 张 | 参考图 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **500,000** |
| 预期 USD | $1.0000 |
| 预期 RMB | ¥7.300000 |
| 结算方式 | 提交预扣 $1，成功后按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 503（预期 200）❌ |
| 耗时 | 0.01s |
| task_id | `-` |
| request_id | 202611071031509348520008268d9d6tFDxvSlt |

```json
{
  "error": {
    "code": "model_not_found",
    "message": "No available channel for model gpt-image-2.6-sunburst under group default (distributor) (request id: 202611071031509348520008268d9d6tFDxvSlt)",
    "type": "new_api_error"
  }
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 任务首次查询 amount | - | - |
| 消费日志 Quota | - | - |
| **预扣金额** | **-** | **-** |
| 预期扣减 | 500,000 | ¥7.300000 |
| 预扣是否符合 | ❌ 提交失败 | — |
| 分组倍率 | - | — |

### 消费日志
```json
[]
```
