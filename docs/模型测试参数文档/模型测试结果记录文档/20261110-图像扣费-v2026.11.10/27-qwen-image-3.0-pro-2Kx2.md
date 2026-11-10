# 27. qwen-image-3.0-pro 2Kx2

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:09  |  模型：`qwen-image-3.0-pro`  |  档位：2K x 2

### 调用参数
```json
{
  "model": "qwen-image-3.0-pro",
  "prompt": "一张密集排版的菜单",
  "resolution": "2K",
  "ratio": "1:1",
  "n": 2
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 2K | 分辨率档位 |
| ratio | 1:1 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 2 | 张数 |
| image | 无 | 参考图 |
| extra_fields.width | 未传 | 成对宽，优先于 size |
| extra_fields.height | 未传 | 成对高，优先于 size |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **54,792** |
| 预期 USD | $0.1096 |
| 预期 RMB | ¥0.799963 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.27s |
| task_id | `task_G6q0UaUI7X6NzEWTuHvDyklfsJ3qVovo` |
| request_id | 202611100644092212820008268d9d6DfyIbhyY |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_G6q0UaUI7X6NzEWTuHvDyklfsJ3qVovo",
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
| 任务首次查询 amount | 54,792 | 0.799963 |
| 消费日志 Quota | 54,792 | 0.799963 |
| **预扣金额** | **54,792** | **¥0.799963** |
| 预期扣减 | 54,792 | ¥0.799963 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 2 |
| 完成耗时 | 8.3s |
| progress | 100% |
| 输出张数 | 2 |

**图片 URL**: https://getapib.org/image/9998205706927255-fa665d34-59bb-4673-8e68-8d379dfb99bf-image_task_01M723709DV0PYVBBX1QWR8DN1_0.png  
https://getapib.org/image/9998205706926898-cacd1481-4c0c-4d70-8954-f374357027ab-image_task_01M723709DV0PYVBBX1QWR8DN1_1.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 54,792 | ¥0.799963 |
| 平台记录最终消费 | 54,792 | **¥0.799963** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 24,
    "billing_status": "settled",
    "completed": 1794293073,
    "cost": 0.114288,
    "created": 1794293049,
    "credits_cost": 1.14288,
    "estimated_time": 100,
    "id": "task_01M7236A3AZ2W7KWNN71D156N4",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379473,
          "url": [
            "https://getapib.org/image/9998205706927255-fa665d34-59bb-4673-8e68-8d379dfb99bf-image_task_01M723709DV0PYVBBX1QWR8DN1_0.png",
            "https://getapib.org/image/9998205706926898-cacd1481-4c0c-4d70-8954-f374357027ab-image_task_01M723709DV0PYVBBX1QWR8DN1_1.png"
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
    "task_id": "task_G6q0UaUI7X6NzEWTuHvDyklfsJ3qVovo",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "qwen-image-3.0-pro",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706927255-fa665d34-59bb-4673-8e68-8d379dfb99bf-image_task_01M723709DV0PYVBBX1QWR8DN1_0.png"
      },
      {
        "url": "https://getapib.org/image/9998205706926898-cacd1481-4c0c-4d70-8954-f374357027ab-image_task_01M723709DV0PYVBBX1QWR8DN1_1.png"
      }
    ],
    "amount": "0.799963",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293049,
    "submit_time": 1794293049,
    "finish_time": 1794293088
  }
}
```

### 消费日志
```json
[
  {
    "id": 10,
    "user_id": 1,
    "created_at": 1794293049,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "qwen-image-3.0-pro",
    "quota": 54792,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611100644092212820008268d9d6DfyIbhyY",
    "upstream_request_id": "20261110144409419409093PdL7KXVv",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
