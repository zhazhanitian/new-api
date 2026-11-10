# 30. muse-image-ext 参考图

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:09  |  模型：`muse-image-ext`  |  档位：固定，参考图不加价

### 调用参数
```json
{
  "model": "muse-image-ext",
  "prompt": "按参考图风格画海边的白色小屋",
  "ratio": "1:1",
  "n": 1,
  "image": [
    "https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk"
  ]
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
| image | 1 张 | 参考图 |
| extra_fields.width | 未传 | 成对宽，优先于 size |
| extra_fields.height | 未传 | 成对高，优先于 size |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **2,302** |
| 预期 USD | $0.0046 |
| 预期 RMB | ¥0.033609 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.29s |
| task_id | `task_3qqm18iuyxvwgnnNbYmuG8i3vRBstXxd` |
| request_id | 202611100644093711380008268d9d6xwNdDhR0 |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_3qqm18iuyxvwgnnNbYmuG8i3vRBstXxd",
    "status": "NOT_START",
    "progress": "0%",
    "model": "muse-image-ext",
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
| 任务首次查询 amount | 2,302 | 0.033609 |
| 消费日志 Quota | - | - |
| **预扣金额** | **2,302** | **¥0.033609** |
| 预期扣减 | 2,302 | ¥0.033609 |
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

**图片 URL**: https://getapib.org/f/images/muimg_7f1096d2e1300c0ce7400c14b7dc7154-0.webp

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 2,302 | ¥0.033609 |
| 平台记录最终消费 | 2,302 | **¥0.033609** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 61,
    "billing_status": "settled",
    "completed": 1794293110,
    "cost": 0.0048,
    "created": 1794293049,
    "credits_cost": 0.047999999999999994,
    "estimated_time": 50,
    "id": "task_01M7236A8H7MRDBCCASACQ60EA",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379510,
          "url": [
            "https://getapib.org/f/images/muimg_7f1096d2e1300c0ce7400c14b7dc7154-0.webp"
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
    "task_id": "task_3qqm18iuyxvwgnnNbYmuG8i3vRBstXxd",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "muse-image-ext",
    "images": [
      {
        "url": "https://getapib.org/f/images/muimg_7f1096d2e1300c0ce7400c14b7dc7154-0.webp"
      }
    ],
    "amount": "0.033609",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293049,
    "submit_time": 1794293049,
    "finish_time": 1794293129
  }
}
```

### 消费日志
```json
[]
```
