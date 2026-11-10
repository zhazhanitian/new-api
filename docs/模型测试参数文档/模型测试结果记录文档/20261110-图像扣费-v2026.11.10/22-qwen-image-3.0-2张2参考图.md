# 22. qwen-image-3.0 2张2参考图

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:53:16  |  模型：`qwen-image-3.0`  |  档位：1K x 2 + 参考图 x 2

### 调用参数
```json
{
  "model": "qwen-image-3.0",
  "prompt": "按参考图风格画一杯咖啡",
  "resolution": "1K",
  "ratio": "1:1",
  "n": 2,
  "image": [
    "https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk",
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
| quality | 未传 | 渲染精度 |
| n | 2 | 张数 |
| image | 2 张 | 参考图 |
| extra_fields.width | 未传 | 成对宽，优先于 size |
| extra_fields.height | 未传 | 成对高，优先于 size |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **21,921** |
| 预期 USD | $0.0438 |
| 预期 RMB | ¥0.320047 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.62s |
| task_id | `task_rUEirW2EPXXLqvAQ0DKDgAPfVOCj8ix8` |
| request_id | 202611100653162912880008268d9d6VFEFE81L |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_rUEirW2EPXXLqvAQ0DKDgAPfVOCj8ix8",
    "status": "NOT_START",
    "progress": "0%",
    "model": "qwen-image-3.0",
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
| 任务首次查询 amount | 21,921 | 0.320047 |
| 消费日志 Quota | - | - |
| **预扣金额** | **21,921** | **¥0.320047** |
| 预期扣减 | 21,921 | ¥0.320047 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 5 |
| 完成耗时 | 32.1s |
| progress | 100% |
| 输出张数 | 2 |

**图片 URL**: https://getapib.org/image/9998205706378774-f2a49c28-6657-497e-9c6e-fc3155ba57ea-image_task_01M723QR6DCB3F57FCHJ8ZQ9Z4_0.png  
https://getapib.org/image/9998205706378507-572b55f6-8e72-4ee7-a056-e455eee92431-image_task_01M723QR6DCB3F57FCHJ8ZQ9Z4_1.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 21,921 | ¥0.320047 |
| 平台记录最终消费 | 21,921 | **¥0.320047** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 25,
    "billing_status": "settled",
    "completed": 1794293621,
    "cost": 0.045716,
    "created": 1794293596,
    "credits_cost": 0.45716,
    "estimated_time": 100,
    "id": "task_01M723Q0N7SH2KXRRJVEFEVZC3",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794380021,
          "url": [
            "https://getapib.org/image/9998205706378774-f2a49c28-6657-497e-9c6e-fc3155ba57ea-image_task_01M723QR6DCB3F57FCHJ8ZQ9Z4_0.png",
            "https://getapib.org/image/9998205706378507-572b55f6-8e72-4ee7-a056-e455eee92431-image_task_01M723QR6DCB3F57FCHJ8ZQ9Z4_1.png"
          ]
        }
      ]
    },
    "status": "completed"
  }
}
```

#### 网关查询响应
```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_rUEirW2EPXXLqvAQ0DKDgAPfVOCj8ix8",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "qwen-image-3.0",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706378774-f2a49c28-6657-497e-9c6e-fc3155ba57ea-image_task_01M723QR6DCB3F57FCHJ8ZQ9Z4_0.png"
      },
      {
        "url": "https://getapib.org/image/9998205706378507-572b55f6-8e72-4ee7-a056-e455eee92431-image_task_01M723QR6DCB3F57FCHJ8ZQ9Z4_1.png"
      }
    ],
    "amount": "0.320047",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293596,
    "submit_time": 1794293596,
    "finish_time": 1794293625
  }
}
```

### 消费日志
```json
[]
```
