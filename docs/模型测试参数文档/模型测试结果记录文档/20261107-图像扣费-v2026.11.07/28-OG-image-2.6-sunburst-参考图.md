# 28. OG-image-2.6-sunburst 参考图

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:41:10  |  模型：`OG-image-2.6-sunburst`  |  档位：按量预扣，输入价一半

### 调用参数
```json
{
  "model": "OG-image-2.6-sunburst",
  "prompt": "按参考图风格画一只橘猫",
  "resolution": "1K",
  "ratio": "1:1",
  "quality": "low",
  "n": 1,
  "image": [
    "https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 1K | 分辨率档位 |
| ratio | 1:1 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | low | 渲染精度 |
| n | 1 | 张数 |
| image | 1 张 | 参考图 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **500,000** |
| 预期 USD | $1.0000 |
| 预期 RMB | ¥7.300000 |
| 结算方式 | 提交预扣 $1，成功后按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.32s |
| task_id | `task_WsvI68XZxwtzvYCkd9o2kdzGHnwEVrHD` |
| request_id | 202611071041107076600008268d9d6O9g7Nqvf |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_WsvI68XZxwtzvYCkd9o2kdzGHnwEVrHD",
    "status": "NOT_START",
    "progress": "0%",
    "model": "OG-image-2.6-sunburst",
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
| 任务首次查询 amount | 500,000 | 7.300000 |
| 消费日志 Quota | 500,000 | 7.300000 |
| **预扣金额** | **500,000** | **¥7.300000** |
| 预期扣减 | 500,000 | ¥7.300000 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 4 |
| 完成耗时 | 24.0s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205951911188-114b96a7-5836-4fcc-b8ce-9e5fe9fc0d62-image_task_01M6TSJPQ7Y8JVQP6QH101HZGD_0.png


#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 500,000 | ¥7.300000 |
| 平台记录最终消费 | 13,219 | **¥0.192997** |
| 差额 | — | -¥7.107003（退还） |
| 最终扣费是否符合 | ✅ 与上游用量按定价页核算一致 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 17,
    "completed": 1794048088,
    "cost": 0.026438,
    "created": 1794048071,
    "credits_cost": 0.26438,
    "estimated_time": 60,
    "id": "task_01M6TSJ5BN2HP9AS2RX8DZPTYT",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794134488,
          "url": [
            "https://getapib.org/image/9998205951911188-114b96a7-5836-4fcc-b8ce-9e5fe9fc0d62-image_task_01M6TSJPQ7Y8JVQP6QH101HZGD_0.png"
          ]
        }
      ]
    },
    "status": "completed",
    "usage": {
      "input_tokens": 1051,
      "input_tokens_details": {
        "cached_tokens": 0,
        "image_tokens": 1024,
        "text_tokens": 27
      },
      "output_tokens": 824,
      "output_tokens_details": {
        "image_tokens": 824,
        "text_tokens": 0
      },
      "total_tokens": 1875
    }
  }
}
```

#### 按上游用量核算
| 项目 | token | 单价（元/1M） | 金额（元） |
|---|---:|---:|---:|
| 文本输入 | 27 | 29.2000 | 0.000788 |
| 缓存文本输入 | 0 | 7.3000 | 0.000000 |
| 图片输入 | 1024 | 46.7200 | 0.047841 |
| 缓存图片输入 | 0 | 11.6800 | 0.000000 |
| 图片输出 | 824 | 175.2000 | 0.144365 |
| 额度取整后 | 13,219 | — | 0.192997 |
| 平台 amount | — | — | 0.192997 |
| 是否一致 | — | — | ✅ 一致 |

#### 网关查询响应
```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_WsvI68XZxwtzvYCkd9o2kdzGHnwEVrHD",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "OG-image-2.6-sunburst",
    "images": [
      {
        "url": "https://getapib.org/image/9998205951911188-114b96a7-5836-4fcc-b8ce-9e5fe9fc0d62-image_task_01M6TSJPQ7Y8JVQP6QH101HZGD_0.png"
      }
    ],
    "amount": "0.192997",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794048071,
    "submit_time": 1794048071,
    "finish_time": 1794048101
  }
}
```

### 消费日志
```json
[
  {
    "id": 1,
    "user_id": 1,
    "created_at": 1794048101,
    "type": 6,
    "content": "adaptor计费调整",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "OG-image-2.6-sunburst",
    "quota": 486781,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "other": "{\"actual_quota\":13219,\"group_ratio\":1,\"is_model_mapped\":true,\"model_price\":0,\"pre_consumed_quota\":500000,\"task_id\":\"task_WsvI68XZxwtzvYCkd9o2kdzGHnwEVrHD\",\"upstream_model_name\":\"gpt-image-2.6-sunburst\"}"
  },
  {
    "id": 5,
    "user_id": 1,
    "created_at": 1794048071,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "OG-image-2.6-sunburst",
    "quota": 500000,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611071041107076600008268d9d6O9g7Nqvf",
    "upstream_request_id": "20261107184110969713928e00nsZbR",
    "other": "{\"group_ratio\":1,\"is_model_mapped\":true,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\",\"upstream_model_name\":\"gpt-image-2.6-sunburst\"}"
  }
]
```
