# 2. flux-3-image-klein 768

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:07  |  模型：`flux-3-image-klein`  |  档位：768sq

### 调用参数
```json
{
  "model": "flux-3-image-klein",
  "prompt": "一只在樱花树下打盹的橘猫",
  "resolution": "768",
  "ratio": "1:1",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 768 | 分辨率档位 |
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
| 预期 Quota | **3,069** |
| 预期 USD | $0.0061 |
| 预期 RMB | ¥0.044807 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.71s |
| task_id | `task_33buVThjkYCkOfe5UjbEZ1oI8Y2NTbT7` |
| request_id | 202611100644079707350008268d9d6VvMVT8Z2 |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_33buVThjkYCkOfe5UjbEZ1oI8Y2NTbT7",
    "status": "NOT_START",
    "progress": "0%",
    "model": "flux-3-image-klein",
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
| 任务首次查询 amount | 3,069 | 0.044807 |
| 消费日志 Quota | - | - |
| **预扣金额** | **3,069** | **¥0.044807** |
| 预期扣减 | 3,069 | ¥0.044807 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 9 |
| 完成耗时 | 64.3s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706941565-b5cfa014-35ee-49d0-b7a0-070ca0e0c62e-image_1794293056603364295.jpg

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 3,069 | ¥0.044807 |
| 平台记录最终消费 | 3,069 | **¥0.044807** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 10,
    "billing_status": "settled",
    "completed": 1794293058,
    "cost": 0.0064,
    "created": 1794293048,
    "credits_cost": 0.064,
    "estimated_time": 100,
    "id": "task_01M723695MV0CPWWECC0ZPN0ND",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379458,
          "url": [
            "https://getapib.org/image/9998205706941565-b5cfa014-35ee-49d0-b7a0-070ca0e0c62e-image_1794293056603364295.jpg"
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
    "task_id": "task_33buVThjkYCkOfe5UjbEZ1oI8Y2NTbT7",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "flux-3-image-klein",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706941565-b5cfa014-35ee-49d0-b7a0-070ca0e0c62e-image_1794293056603364295.jpg"
      }
    ],
    "amount": "0.044807",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293048,
    "submit_time": 1794293048,
    "finish_time": 1794293113
  }
}
```

### 消费日志
```json
[]
```
