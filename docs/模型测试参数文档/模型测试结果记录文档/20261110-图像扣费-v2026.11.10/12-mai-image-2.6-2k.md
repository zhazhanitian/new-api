# 12. mai-image-2.6 2k

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:08  |  模型：`mai-image-2.6`  |  档位：2K

### 调用参数
```json
{
  "model": "mai-image-2.6",
  "prompt": "一只在窗台上的橘猫",
  "resolution": "2k",
  "ratio": "16:9",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 2k | 分辨率档位 |
| ratio | 16:9 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 无 | 参考图 |
| extra_fields.width | 未传 | 成对宽，优先于 size |
| extra_fields.height | 未传 | 成对高，优先于 size |

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
| 耗时 | 0.25s |
| task_id | `task_PzCsD3mZ2DodfSpv8nE6QUvQbjcR3p43` |
| request_id | 202611100644084709180008268d9d6XV6kF2Sa |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_PzCsD3mZ2DodfSpv8nE6QUvQbjcR3p43",
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
| 轮询次数 | 11 |
| 完成耗时 | 80.3s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706907527-a3b70796-0f91-479f-8316-f210b19e52ca-image_task_01M7237M3JNC3ZR1FY1GAMXZVY_0.png

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
    "actual_time": 44,
    "billing_status": "settled",
    "completed": 1794293092,
    "cost": 0.070074,
    "created": 1794293048,
    "credits_cost": 0.7007399999999999,
    "estimated_time": 60,
    "id": "task_01M72369BP5QR1D8CHB9ZB0610",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379492,
          "url": [
            "https://getapib.org/image/9998205706907527-a3b70796-0f91-479f-8316-f210b19e52ca-image_task_01M7237M3JNC3ZR1FY1GAMXZVY_0.png"
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
    "task_id": "task_PzCsD3mZ2DodfSpv8nE6QUvQbjcR3p43",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "mai-image-2.6",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706907527-a3b70796-0f91-479f-8316-f210b19e52ca-image_task_01M7237M3JNC3ZR1FY1GAMXZVY_0.png"
      }
    ],
    "amount": "0.490283",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293048,
    "submit_time": 1794293048,
    "finish_time": 1794293123
  }
}
```

### 消费日志
```json
[]
```
