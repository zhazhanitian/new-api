# 14. mai-image-2.6 宽高覆盖size

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:53:16  |  模型：`mai-image-2.6`  |  档位：width/height 1536x1536 优先，按 2K

### 调用参数
```json
{
  "model": "mai-image-2.6",
  "prompt": "一只在窗台上的橘猫",
  "size": "1024x1024",
  "n": 1,
  "extra_fields": {
    "width": 1536,
    "height": 1536
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 未传 | 分辨率档位 |
| ratio | 未传 | 宽高比 |
| size | 1024x1024 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 无 | 参考图 |
| extra_fields.width | 1536 | 成对宽，优先于 size |
| extra_fields.height | 1536 | 成对高，优先于 size |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **33,581** |
| 预期 USD | $0.0672 |
| 预期 RMB | ¥0.490283 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.66s |
| task_id | `task_g0TcRAus16ZlA4vqezsUTbHNCnsymAnw` |
| request_id | 202611100653162606220008268d9d60d7iIfPW |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_g0TcRAus16ZlA4vqezsUTbHNCnsymAnw",
    "status": "NOT_START",
    "progress": "0%",
    "model": "mai-image-2.6",
    "images": [],
    "amount": "",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 0,
    "submit_time": 0,
    "finish_time": 0
  }
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 任务首次查询 amount | 33,581 | 0.490283 |
| 消费日志 Quota | - | - |
| **预扣金额** | **33,581** | **¥0.490283** |
| 预期扣减 | 33,581 | ¥0.490283 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 7 |
| 完成耗时 | 48.1s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706358657-c9833bfc-035f-4622-83e5-9aded29d40ab-image_task_01M723RC3DNC9FP247ZDSFQ58Z_0.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 33,581 | ¥0.490283 |
| 平台记录最终消费 | 33,581 | **¥0.490283** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 45,
    "billing_status": "settled",
    "completed": 1794293641,
    "cost": 0.070074,
    "created": 1794293596,
    "credits_cost": 0.7007399999999999,
    "estimated_time": 60,
    "id": "task_01M723Q0MY77D1D1NTE2J09A5B",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794380041,
          "url": [
            "https://getapib.org/image/9998205706358657-c9833bfc-035f-4622-83e5-9aded29d40ab-image_task_01M723RC3DNC9FP247ZDSFQ58Z_0.png"
          ]
        }
      ]
    },
    "status": "completed",
    "usage": {
      "input_tokens": 8,
      "input_tokens_details": {
        "cached_tokens": 0,
        "image_tokens": 0,
        "text_tokens": 8
      },
      "output_tokens": 2304,
      "output_tokens_details": {
        "image_tokens": 2304,
        "text_tokens": 0
      },
      "total_tokens": 2312
    }
  }
}
```

#### 网关查询响应
```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_g0TcRAus16ZlA4vqezsUTbHNCnsymAnw",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "mai-image-2.6",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706358657-c9833bfc-035f-4622-83e5-9aded29d40ab-image_task_01M723RC3DNC9FP247ZDSFQ58Z_0.png"
      }
    ],
    "amount": "0.490283",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293596,
    "submit_time": 1794293596,
    "finish_time": 1794293642
  }
}
```

### 消费日志
```json
[]
```
