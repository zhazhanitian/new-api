# 3. flux-3-image-klein 1.5K

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:08  |  模型：`flux-3-image-klein`  |  档位：1.5k

### 调用参数
```json
{
  "model": "flux-3-image-klein",
  "prompt": "一只在樱花树下打盹的橘猫",
  "resolution": "1.5K",
  "ratio": "4:3",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 1.5K | 分辨率档位 |
| ratio | 4:3 | 宽高比 |
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
| 预期 Quota | **4,219** |
| 预期 USD | $0.0084 |
| 预期 RMB | ¥0.061597 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.67s |
| task_id | `task_U8rN0PXYwxb7wABPlhdhW4uilap4Y2KS` |
| request_id | 20261110064408213710008268d9d6CsPhBje5 |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_U8rN0PXYwxb7wABPlhdhW4uilap4Y2KS",
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
| 任务首次查询 amount | 4,219 | 0.061597 |
| 消费日志 Quota | 4,219 | 0.061597 |
| **预扣金额** | **4,219** | **¥0.061597** |
| 预期扣减 | 4,219 | ¥0.061597 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 3 |
| 完成耗时 | 16.1s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706941198-861e20ff-f111-472b-9fb2-23a747d96787-image_1794293056573638920.jpg

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 4,219 | ¥0.061597 |
| 平台记录最终消费 | 4,219 | **¥0.061597** |
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
    "cost": 0.0088,
    "created": 1794293048,
    "credits_cost": 0.08800000000000001,
    "estimated_time": 100,
    "id": "task_01M723695NGVDZ0MMVPJMCW5Q8",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379458,
          "url": [
            "https://getapib.org/image/9998205706941198-861e20ff-f111-472b-9fb2-23a747d96787-image_1794293056573638920.jpg"
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
    "task_id": "task_U8rN0PXYwxb7wABPlhdhW4uilap4Y2KS",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "flux-3-image-klein",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706941198-861e20ff-f111-472b-9fb2-23a747d96787-image_1794293056573638920.jpg"
      }
    ],
    "amount": "0.061597",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293048,
    "submit_time": 1794293048,
    "finish_time": 1794293060
  }
}
```

### 消费日志
```json
[
  {
    "id": 30,
    "user_id": 1,
    "created_at": 1794293048,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "flux-3-image-klein",
    "quota": 4219,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "20261110064408213710008268d9d6CsPhBje5",
    "upstream_request_id": "202611101444084644394093QUudvld",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
