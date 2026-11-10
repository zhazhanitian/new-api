# 28. qwen-image-3.0-pro 1Kx2参考图

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:53:16  |  模型：`qwen-image-3.0-pro`  |  档位：1K x 2 + 1 张参考图

### 调用参数
```json
{
  "model": "qwen-image-3.0-pro",
  "prompt": "按参考图风格排一张菜单",
  "resolution": "1K",
  "ratio": "3:4",
  "n": 2,
  "image": [
    "https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 1K | 分辨率档位 |
| ratio | 3:4 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 2 | 张数 |
| image | 1 张 | 参考图 |
| extra_fields.width | 未传 | 成对宽，优先于 size |
| extra_fields.height | 未传 | 成对高，优先于 size |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **28,494** |
| 预期 USD | $0.0570 |
| 预期 RMB | ¥0.416012 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.53s |
| task_id | `task_NJy3aVijJvZOobZHeCR5JEguAKLmCNFp` |
| request_id | 202611100653163814960008268d9d6CVVaMJBM |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_NJy3aVijJvZOobZHeCR5JEguAKLmCNFp",
    "status": "NOT_START",
    "progress": "0%",
    "model": "qwen-image-3.0-pro",
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
| 任务首次查询 amount | 28,494 | 0.416012 |
| 消费日志 Quota | - | - |
| **预扣金额** | **28,494** | **¥0.416012** |
| 预期扣减 | 28,494 | ¥0.416012 |
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

**图片 URL**: https://getapib.org/image/9998205706387777-cf68ca45-b702-4b2f-8b6b-1a2815533efd-image_task_01M723QFD2E8VS7YPVG8RPXXT6_0.png  
https://getapib.org/image/9998205706387412-f3e2ca79-16a7-47f2-9fe8-e7b08da86727-image_task_01M723QFD2E8VS7YPVG8RPXXT6_1.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 28,494 | ¥0.416012 |
| 平台记录最终消费 | 28,494 | **¥0.416012** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 16,
    "billing_status": "settled",
    "completed": 1794293612,
    "cost": 0.05943,
    "created": 1794293596,
    "credits_cost": 0.5942999999999999,
    "estimated_time": 100,
    "id": "task_01M723Q0MWHGWBRRY876ZTA8AM",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794380012,
          "url": [
            "https://getapib.org/image/9998205706387777-cf68ca45-b702-4b2f-8b6b-1a2815533efd-image_task_01M723QFD2E8VS7YPVG8RPXXT6_0.png",
            "https://getapib.org/image/9998205706387412-f3e2ca79-16a7-47f2-9fe8-e7b08da86727-image_task_01M723QFD2E8VS7YPVG8RPXXT6_1.png"
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
    "task_id": "task_NJy3aVijJvZOobZHeCR5JEguAKLmCNFp",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "qwen-image-3.0-pro",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706387777-cf68ca45-b702-4b2f-8b6b-1a2815533efd-image_task_01M723QFD2E8VS7YPVG8RPXXT6_0.png"
      },
      {
        "url": "https://getapib.org/image/9998205706387412-f3e2ca79-16a7-47f2-9fe8-e7b08da86727-image_task_01M723QFD2E8VS7YPVG8RPXXT6_1.png"
      }
    ],
    "amount": "0.416012",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293596,
    "submit_time": 1794293596,
    "finish_time": 1794293623
  }
}
```

### 消费日志
```json
[]
```
