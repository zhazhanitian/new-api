# 10. flux-3-image 4k

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:08  |  模型：`flux-3-image`  |  档位：4k

### 调用参数
```json
{
  "model": "flux-3-image",
  "prompt": "浩瀚星空下的沙漠",
  "resolution": "4k",
  "ratio": "16:9",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 4k | 分辨率档位 |
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
| 预期 Quota | **232,822** |
| 预期 USD | $0.4656 |
| 预期 RMB | ¥3.399201 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.32s |
| task_id | `task_2Fy3BPJZSRwMJy1Ef2r3SKN2sQ8Zjdht` |
| request_id | 202611100644083722100008268d9d6iotdHHGf |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_2Fy3BPJZSRwMJy1Ef2r3SKN2sQ8Zjdht",
    "status": "NOT_START",
    "progress": "0%",
    "model": "flux-3-image",
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
| 任务首次查询 amount | 232,822 | 3.399201 |
| 消费日志 Quota | - | - |
| **预扣金额** | **232,822** | **¥3.399201** |
| 预期扣减 | 232,822 | ¥3.399201 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 14 |
| 完成耗时 | 104.3s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706867119-d66efe88-67cb-4b2f-875b-a3276120296b-image_1794293123020798166.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 232,822 | ¥3.399201 |
| 平台记录最终消费 | 232,822 | **¥3.399201** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 85,
    "billing_status": "settled",
    "completed": 1794293133,
    "cost": 0.4856,
    "created": 1794293048,
    "credits_cost": 4.856,
    "estimated_time": 100,
    "id": "task_01M723698HDZWH11C5AN1Q07KP",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379533,
          "url": [
            "https://getapib.org/image/9998205706867119-d66efe88-67cb-4b2f-875b-a3276120296b-image_1794293123020798166.png"
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
    "task_id": "task_2Fy3BPJZSRwMJy1Ef2r3SKN2sQ8Zjdht",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "flux-3-image",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706867119-d66efe88-67cb-4b2f-875b-a3276120296b-image_1794293123020798166.png"
      }
    ],
    "amount": "3.399201",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293048,
    "submit_time": 1794293048,
    "finish_time": 1794293146
  }
}
```

### 消费日志
```json
[]
```
