# 17. mai-image-2.6-flash 2K

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:08  |  模型：`mai-image-2.6-flash`  |  档位：2K

### 调用参数
```json
{
  "model": "mai-image-2.6-flash",
  "prompt": "一只在窗台上的橘猫",
  "resolution": "2K",
  "ratio": "16:9",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 2K | 分辨率档位 |
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
| 预期 Quota | **16,791** |
| 预期 USD | $0.0336 |
| 预期 RMB | ¥0.245149 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.27s |
| task_id | `task_Vd2SA9OTOqYzIs1vzaF1IEGA9RXusMZV` |
| request_id | 202611100644087203740008268d9d69Pqd9DLq |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_Vd2SA9OTOqYzIs1vzaF1IEGA9RXusMZV",
    "status": "NOT_START",
    "progress": "0%",
    "model": "mai-image-2.6-flash",
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
| 任务首次查询 amount | 16,791 | 0.245149 |
| 消费日志 Quota | 16,791 | 0.245149 |
| **预扣金额** | **16,791** | **¥0.245149** |
| 预期扣减 | 16,791 | ¥0.245149 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 5 |
| 完成耗时 | 32.2s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706930264-d13dd9f7-01ed-4843-bb97-c66f740a0476-image_task_01M7236XWNFTHZYJ2V4TGZ1GJJ_0.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 16,791 | ¥0.245149 |
| 平台记录最终消费 | 16,791 | **¥0.245149** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 21,
    "billing_status": "settled",
    "completed": 1794293069,
    "cost": 0.035032,
    "created": 1794293048,
    "credits_cost": 0.35032,
    "estimated_time": 60,
    "id": "task_01M72369KY3CAFK4D0DSW53WCZ",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379469,
          "url": [
            "https://getapib.org/image/9998205706930264-d13dd9f7-01ed-4843-bb97-c66f740a0476-image_task_01M7236XWNFTHZYJ2V4TGZ1GJJ_0.png"
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
    "task_id": "task_Vd2SA9OTOqYzIs1vzaF1IEGA9RXusMZV",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "mai-image-2.6-flash",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706930264-d13dd9f7-01ed-4843-bb97-c66f740a0476-image_task_01M7236XWNFTHZYJ2V4TGZ1GJJ_0.png"
      }
    ],
    "amount": "0.245149",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293048,
    "submit_time": 1794293048,
    "finish_time": 1794293077
  }
}
```

### 消费日志
```json
[
  {
    "id": 17,
    "user_id": 1,
    "created_at": 1794293048,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "mai-image-2.6-flash",
    "quota": 16791,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611100644087203740008268d9d69Pqd9DLq",
    "upstream_request_id": "20261110144408919016284N4YTZkBl",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
