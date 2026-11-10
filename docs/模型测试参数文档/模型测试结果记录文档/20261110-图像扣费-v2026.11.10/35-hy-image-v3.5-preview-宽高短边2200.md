# 35. hy-image-v3.5-preview 宽高短边2200

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:53:16  |  模型：`hy-image-v3.5-preview`  |  档位：width/height 3008x2208 优先，短边 2208 按 4K

### 调用参数
```json
{
  "model": "hy-image-v3.5-preview",
  "prompt": "雨后的江南小镇",
  "size": "1024x1024",
  "n": 1,
  "extra_fields": {
    "width": 3008,
    "height": 2208
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
| extra_fields.width | 3008 | 成对宽，优先于 size |
| extra_fields.height | 2208 | 成对高，优先于 size |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **12,274** |
| 预期 USD | $0.0245 |
| 预期 RMB | ¥0.179200 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.46s |
| task_id | `task_u96wPb1ok0jLvASnMKaU8yVX4Ti9LFLR` |
| request_id | 202611100653164321710008268d9d67ToxZdvi |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_u96wPb1ok0jLvASnMKaU8yVX4Ti9LFLR",
    "status": "NOT_START",
    "progress": "0%",
    "model": "hy-image-v3.5-preview",
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
| 任务首次查询 amount | 12,274 | 0.179200 |
| 消费日志 Quota | - | - |
| **预扣金额** | **12,274** | **¥0.179200** |
| 预期扣减 | 12,274 | ¥0.179200 |
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

**图片 URL**: https://getapib.org/image/9998205706371149-8cebc68a-987a-4737-9b85-188c18e1bc09-image_1794293627784594642_0.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 12,274 | ¥0.179200 |
| 平台记录最终消费 | 12,274 | **¥0.179200** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 32,
    "billing_status": "settled",
    "completed": 1794293628,
    "cost": 0.0256,
    "created": 1794293596,
    "credits_cost": 0.256,
    "estimated_time": 60,
    "id": "task_01M723Q0MJJXDSV1MYG5M95K4K",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794380028,
          "url": [
            "https://getapib.org/image/9998205706371149-8cebc68a-987a-4737-9b85-188c18e1bc09-image_1794293627784594642_0.png"
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
    "task_id": "task_u96wPb1ok0jLvASnMKaU8yVX4Ti9LFLR",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "hy-image-v3.5-preview",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706371149-8cebc68a-987a-4737-9b85-188c18e1bc09-image_1794293627784594642_0.png"
      }
    ],
    "amount": "0.179200",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293596,
    "submit_time": 1794293596,
    "finish_time": 1794293641
  }
}
```

### 消费日志
```json
[]
```
