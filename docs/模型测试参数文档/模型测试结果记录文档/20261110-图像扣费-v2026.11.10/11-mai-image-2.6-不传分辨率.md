# 11. mai-image-2.6 不传分辨率

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:08  |  模型：`mai-image-2.6`  |  档位：1K

### 调用参数
```json
{
  "model": "mai-image-2.6",
  "prompt": "一只在窗台上的橘猫",
  "ratio": "1:1",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 未传 | 分辨率档位 |
| ratio | 1:1 | 宽高比 |
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
| 预期 Quota | **14,926** |
| 预期 USD | $0.0299 |
| 预期 RMB | ¥0.217920 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.28s |
| task_id | `task_tthSIc7D46aJJ0WTFnxMhbvFu4RaIVcw` |
| request_id | 202611100644084213390008268d9d6fcW0svmq |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_tthSIc7D46aJJ0WTFnxMhbvFu4RaIVcw",
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
| 任务首次查询 amount | 14,926 | 0.217920 |
| 消费日志 Quota | - | - |
| **预扣金额** | **14,926** | **¥0.217920** |
| 预期扣减 | 14,926 | ¥0.217920 |
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

**图片 URL**: https://getapib.org/image/9998205706922312-150c302a-bd10-459a-af99-75568f2492d6-image_task_01M72375NMRC9SEM2X5ZA96J0H_0.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 14,926 | ¥0.217920 |
| 平台记录最终消费 | 14,926 | **¥0.217920** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 29,
    "billing_status": "settled",
    "completed": 1794293077,
    "cost": 0.031162,
    "created": 1794293048,
    "credits_cost": 0.31162,
    "estimated_time": 60,
    "id": "task_01M72369ASH6YCABE15345SRPH",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379477,
          "url": [
            "https://getapib.org/image/9998205706922312-150c302a-bd10-459a-af99-75568f2492d6-image_task_01M72375NMRC9SEM2X5ZA96J0H_0.png"
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
      "output_tokens": 1024,
      "output_tokens_details": {
        "image_tokens": 1024,
        "text_tokens": 0
      },
      "total_tokens": 1032
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
    "task_id": "task_tthSIc7D46aJJ0WTFnxMhbvFu4RaIVcw",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "mai-image-2.6",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706922312-150c302a-bd10-459a-af99-75568f2492d6-image_task_01M72375NMRC9SEM2X5ZA96J0H_0.png"
      }
    ],
    "amount": "0.217920",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293048,
    "submit_time": 1794293048,
    "finish_time": 1794293122
  }
}
```

### 消费日志
```json
[]
```
