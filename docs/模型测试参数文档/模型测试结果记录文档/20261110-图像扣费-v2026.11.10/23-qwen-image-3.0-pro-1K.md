# 23. qwen-image-3.0-pro 1K

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:09  |  模型：`qwen-image-3.0-pro`  |  档位：1K

### 调用参数
```json
{
  "model": "qwen-image-3.0-pro",
  "prompt": "一张密集排版的菜单",
  "resolution": "1K",
  "ratio": "3:4",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 1K | 分辨率档位 |
| ratio | 3:4 | 宽高比 |
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
| 预期 Quota | **13,698** |
| 预期 USD | $0.0274 |
| 预期 RMB | ¥0.199991 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.29s |
| task_id | `task_5IQ7zBo5M0tHPmmLcykU58GF0CXSswJt` |
| request_id | 20261110064409206320008268d9d66SdxsPZS |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_5IQ7zBo5M0tHPmmLcykU58GF0CXSswJt",
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
| 任务首次查询 amount | 13,698 | 0.199991 |
| 消费日志 Quota | 13,698 | 0.199991 |
| **预扣金额** | **13,698** | **¥0.199991** |
| 预期扣减 | 13,698 | ¥0.199991 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 4 |
| 完成耗时 | 24.5s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706941960-46bb2511-4786-474b-b664-8c1903e91a9d-image_task_01M7236J9PNNXMFPTPAKS4JTWH_0.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 13,698 | ¥0.199991 |
| 平台记录最终消费 | 13,698 | **¥0.199991** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 9,
    "billing_status": "settled",
    "completed": 1794293058,
    "cost": 0.028572,
    "created": 1794293049,
    "credits_cost": 0.28572,
    "estimated_time": 100,
    "id": "task_01M72369XFGTTP6VVA4ZG3V60E",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379458,
          "url": [
            "https://getapib.org/image/9998205706941960-46bb2511-4786-474b-b664-8c1903e91a9d-image_task_01M7236J9PNNXMFPTPAKS4JTWH_0.png"
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
    "task_id": "task_5IQ7zBo5M0tHPmmLcykU58GF0CXSswJt",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "qwen-image-3.0-pro",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706941960-46bb2511-4786-474b-b664-8c1903e91a9d-image_task_01M7236J9PNNXMFPTPAKS4JTWH_0.png"
      }
    ],
    "amount": "0.199991",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293049,
    "submit_time": 1794293049,
    "finish_time": 1794293084
  }
}
```

### 消费日志
```json
[
  {
    "id": 13,
    "user_id": 1,
    "created_at": 1794293049,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "qwen-image-3.0-pro",
    "quota": 13698,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "20261110064409206320008268d9d66SdxsPZS",
    "upstream_request_id": "20261110144409233516302PTOE2e5p",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
