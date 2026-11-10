# 20. doubao-seedream-5-0-260128 1张

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:21:20  |  模型：`doubao-seedream-5-0-260128`  |  档位：固定 0.220

### 调用参数
```json
{
  "model": "doubao-seedream-5-0-260128",
  "prompt": "一只白瓷茶杯放在木桌上",
  "resolution": "2K",
  "ratio": "1:1",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 2K | 分辨率档位 |
| ratio | 1:1 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 无 | 参考图 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **15,069** |
| 预期 USD | $0.0301 |
| 预期 RMB | ¥0.220007 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.02s |
| task_id | `task_PI2LZm1VFJrY5iUG0tx80XzA9J7IhN9b` |
| request_id | 202611071021206466720008268d9d6lU2RrKsV |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_PI2LZm1VFJrY5iUG0tx80XzA9J7IhN9b",
    "status": "QUEUED",
    "progress": "0%",
    "model": "doubao-seedream-5-0-260128",
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
| 任务首次查询 amount | 15,069 | 0.220007 |
| 消费日志 Quota | 15,069 | 0.220007 |
| **预扣金额** | **15,069** | **¥0.220007** |
| 预期扣减 | 15,069 | ¥0.220007 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 1 |
| 完成耗时 | 0.0s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedream-5-0/021794046880811388e10ad0dcd28daa4414437ebb5a857924134_0.jpeg?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261107%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261107T102135Z&X-Tos-Expires=86400&X-Tos-Signature=b67cb6d6fda76e3100803d3b4c1df453a8e2e5144700faafc97283ae467ba6d1&X-Tos-SignedHeaders=host


#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 15,069 | ¥0.220007 |
| 平台记录最终消费 | 15,069 | **¥0.220007** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "created": 1794046895,
  "data": [
    {
      "size": "2048x2048",
      "url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedream-5-0/021794046880811388e10ad0dcd28daa4414437ebb5a857924134_0.jpeg?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261107%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261107T102135Z&X-Tos-Expires=86400&X-Tos-Signature=b67cb6d6fda76e3100803d3b4c1df453a8e2e5144700faafc97283ae467ba6d1&X-Tos-SignedHeaders=host"
    }
  ],
  "model": "doubao-seedream-5-0-260128",
  "usage": {
    "generated_images": 1,
    "output_tokens": 16384,
    "total_tokens": 16384
  }
}
```

#### 网关查询响应
```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_PI2LZm1VFJrY5iUG0tx80XzA9J7IhN9b",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "doubao-seedream-5-0-260128",
    "images": [
      {
        "url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedream-5-0/021794046880811388e10ad0dcd28daa4414437ebb5a857924134_0.jpeg?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261107%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261107T102135Z&X-Tos-Expires=86400&X-Tos-Signature=b67cb6d6fda76e3100803d3b4c1df453a8e2e5144700faafc97283ae467ba6d1&X-Tos-SignedHeaders=host"
      }
    ],
    "amount": "0.220007",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 16384,
      "total_tokens": 16384
    },
    "created_at": 1794046880,
    "submit_time": 1794046880,
    "finish_time": 1794046895
  }
}
```

### 消费日志
```json
[
  {
    "id": 9,
    "user_id": 1,
    "created_at": 1794046880,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "doubao-seedream-5-0-260128",
    "quota": 15069,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 15,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611071021206466720008268d9d6lU2RrKsV",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
