# 20. qwen-image-3.0 2张

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:08  |  模型：`qwen-image-3.0`  |  档位：1K x 2

### 调用参数
```json
{
  "model": "qwen-image-3.0",
  "prompt": "一杯咖啡的顶视角海报",
  "resolution": "1K",
  "ratio": "1:1",
  "n": 2
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 1K | 分辨率档位 |
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
| 预期 Quota | **19,725** |
| 预期 USD | $0.0394 |
| 预期 RMB | ¥0.287985 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.27s |
| task_id | `task_i0hv8cyuvwdBaLWcXAJsfNZrNP3JqogY` |
| request_id | 202611100644088704830008268d9d6wLHNA3La |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_i0hv8cyuvwdBaLWcXAJsfNZrNP3JqogY",
    "status": "NOT_START",
    "progress": "0%",
    "model": "qwen-image-3.0",
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
| 任务首次查询 amount | 19,725 | 0.287985 |
| 消费日志 Quota | 19,725 | 0.287985 |
| **预扣金额** | **19,725** | **¥0.287985** |
| 预期扣减 | 19,725 | ¥0.287985 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 5 |
| 完成耗时 | 32.2s |
| progress | 100% |
| 输出张数 | 2 |

**图片 URL**: https://getapib.org/image/9998205706945382-6b998110-6c3a-4198-8924-771cd88985cd-image_task_01M7236EWXX06W82MBMZ6ZGD9E_0.png  
https://getapib.org/image/9998205706945037-00d6ae7b-7468-4b70-b08c-5809b94007c1-image_task_01M7236EWXX06W82MBMZ6ZGD9E_1.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 19,725 | ¥0.287985 |
| 平台记录最终消费 | 19,725 | **¥0.287985** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 6,
    "billing_status": "settled",
    "completed": 1794293055,
    "cost": 0.041144,
    "created": 1794293049,
    "credits_cost": 0.41144000000000003,
    "estimated_time": 100,
    "id": "task_01M72369RBZGQWXZ8PJ3AG9DE7",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379455,
          "url": [
            "https://getapib.org/image/9998205706945382-6b998110-6c3a-4198-8924-771cd88985cd-image_task_01M7236EWXX06W82MBMZ6ZGD9E_0.png",
            "https://getapib.org/image/9998205706945037-00d6ae7b-7468-4b70-b08c-5809b94007c1-image_task_01M7236EWXX06W82MBMZ6ZGD9E_1.png"
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
    "task_id": "task_i0hv8cyuvwdBaLWcXAJsfNZrNP3JqogY",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "qwen-image-3.0",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706945382-6b998110-6c3a-4198-8924-771cd88985cd-image_task_01M7236EWXX06W82MBMZ6ZGD9E_0.png"
      },
      {
        "url": "https://getapib.org/image/9998205706945037-00d6ae7b-7468-4b70-b08c-5809b94007c1-image_task_01M7236EWXX06W82MBMZ6ZGD9E_1.png"
      }
    ],
    "amount": "0.287985",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293049,
    "submit_time": 1794293049,
    "finish_time": 1794293080
  }
}
```

### 消费日志
```json
[
  {
    "id": 14,
    "user_id": 1,
    "created_at": 1794293049,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "qwen-image-3.0",
    "quota": 19725,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611100644088704830008268d9d6wLHNA3La",
    "upstream_request_id": "2026111014440971025537MkcLEALW",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
