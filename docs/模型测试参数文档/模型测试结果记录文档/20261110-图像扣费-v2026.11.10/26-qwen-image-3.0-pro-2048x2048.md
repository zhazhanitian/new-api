# 26. qwen-image-3.0-pro 2048x2048

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:09  |  模型：`qwen-image-3.0-pro`  |  档位：像素过线，按 2K

### 调用参数
```json
{
  "model": "qwen-image-3.0-pro",
  "prompt": "一张密集排版的菜单",
  "size": "2048x2048",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 未传 | 分辨率档位 |
| ratio | 未传 | 宽高比 |
| size | 2048x2048 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 无 | 参考图 |
| extra_fields.width | 未传 | 成对宽，优先于 size |
| extra_fields.height | 未传 | 成对高，优先于 size |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **27,396** |
| 预期 USD | $0.0548 |
| 预期 RMB | ¥0.399982 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.26s |
| task_id | `task_VXKQGKhc4i35oA7ZC7ossXdJSK6dC26a` |
| request_id | 202611100644091708670008268d9d6m9ulkrmT |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_VXKQGKhc4i35oA7ZC7ossXdJSK6dC26a",
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
| 任务首次查询 amount | 27,396 | 0.399982 |
| 消费日志 Quota | 27,396 | 0.399982 |
| **预扣金额** | **27,396** | **¥0.399982** |
| 预期扣减 | 27,396 | ¥0.399982 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 2 |
| 完成耗时 | 8.0s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706929335-176cb3bb-803b-4500-9dbe-42beb85fa9c3-image_task_01M7236Y3JGAKS46GWVP664E64_0.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 27,396 | ¥0.399982 |
| 平台记录最终消费 | 27,396 | **¥0.399982** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 21,
    "billing_status": "settled",
    "completed": 1794293070,
    "cost": 0.057144,
    "created": 1794293049,
    "credits_cost": 0.57144,
    "estimated_time": 100,
    "id": "task_01M7236A1BDR075G7SJ3GB6285",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379470,
          "url": [
            "https://getapib.org/image/9998205706929335-176cb3bb-803b-4500-9dbe-42beb85fa9c3-image_task_01M7236Y3JGAKS46GWVP664E64_0.png"
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
    "task_id": "task_VXKQGKhc4i35oA7ZC7ossXdJSK6dC26a",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "qwen-image-3.0-pro",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706929335-176cb3bb-803b-4500-9dbe-42beb85fa9c3-image_task_01M7236Y3JGAKS46GWVP664E64_0.png"
      }
    ],
    "amount": "0.399982",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293049,
    "submit_time": 1794293049,
    "finish_time": 1794293087
  }
}
```

### 消费日志
```json
[
  {
    "id": 11,
    "user_id": 1,
    "created_at": 1794293049,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "qwen-image-3.0-pro",
    "quota": 27396,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611100644091708670008268d9d6m9ulkrmT",
    "upstream_request_id": "20261110144409359082737rl9sSwA4",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
