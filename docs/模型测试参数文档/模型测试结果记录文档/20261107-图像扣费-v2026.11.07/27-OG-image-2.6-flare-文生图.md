# 27. OG-image-2.6-flare 文生图

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:41:08  |  模型：`OG-image-2.6-flare`  |  档位：按量预扣，输入价一半

### 调用参数
```json
{
  "model": "OG-image-2.6-flare",
  "prompt": "一只橘猫坐在窗台上",
  "resolution": "1K",
  "ratio": "1:1",
  "quality": "low",
  "n": 1
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
| image | 无 | 参考图 |

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
| 耗时 | 0.66s |
| task_id | `task_Pgo719UBMORTMqlm13exGfTPs15b1ZuO` |
| request_id | 20261107104108424370008268d9d6LPrAoMxs |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_Pgo719UBMORTMqlm13exGfTPs15b1ZuO",
    "status": "NOT_START",
    "progress": "0%",
    "model": "OG-image-2.6-flare",
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
| 轮询次数 | 2 |
| 完成耗时 | 8.0s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205951924010-a9c89c39-782a-4bed-9a3e-90cb3542c411-image_task_01M6TSJA6HBJSFN0QBJ9WTF5XV_0.png


#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 500,000 | ¥7.300000 |
| 平台记录最终消费 | 9,920 | **¥0.144832** |
| 差额 | — | -¥7.155168（退还） |
| 最终扣费是否符合 | ✅ 与上游用量按定价页核算一致 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 8,
    "completed": 1794048076,
    "cost": 0.01984,
    "created": 1794048068,
    "credits_cost": 0.1984,
    "estimated_time": 60,
    "id": "task_01M6TSJ32911NWVGK9XDRKTN24",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794134476,
          "url": [
            "https://getapib.org/image/9998205951924010-a9c89c39-782a-4bed-9a3e-90cb3542c411-image_task_01M6TSJA6HBJSFN0QBJ9WTF5XV_0.png"
          ]
        }
      ]
    },
    "status": "completed",
    "usage": {
      "input_tokens": 16,
      "input_tokens_details": {
        "cached_tokens": 0,
        "image_tokens": 0,
        "text_tokens": 16
      },
      "output_tokens": 824,
      "output_tokens_details": {
        "image_tokens": 824,
        "text_tokens": 0
      },
      "total_tokens": 840
    }
  }
}
```

#### 按上游用量核算
| 项目 | token | 单价（元/1M） | 金额（元） |
|---|---:|---:|---:|
| 文本输入 | 16 | 29.2000 | 0.000467 |
| 缓存文本输入 | 0 | 7.3000 | 0.000000 |
| 图片输入 | 0 | 46.7200 | 0.000000 |
| 缓存图片输入 | 0 | 11.6800 | 0.000000 |
| 图片输出 | 824 | 175.2000 | 0.144365 |
| 额度取整后 | 9,920 | — | 0.144832 |
| 平台 amount | — | — | 0.144832 |
| 是否一致 | — | — | ✅ 一致 |

#### 网关查询响应
```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_Pgo719UBMORTMqlm13exGfTPs15b1ZuO",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "OG-image-2.6-flare",
    "images": [
      {
        "url": "https://getapib.org/image/9998205951924010-a9c89c39-782a-4bed-9a3e-90cb3542c411-image_task_01M6TSJA6HBJSFN0QBJ9WTF5XV_0.png"
      }
    ],
    "amount": "0.144832",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794048068,
    "submit_time": 1794048068,
    "finish_time": 1794048081
  }
}
```

### 消费日志
```json
[
  {
    "id": 1,
    "user_id": 1,
    "created_at": 1794048081,
    "type": 6,
    "content": "adaptor计费调整",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "OG-image-2.6-flare",
    "quota": 490080,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 21,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "other": "{\"actual_quota\":9920,\"group_ratio\":1,\"is_model_mapped\":true,\"model_price\":0,\"pre_consumed_quota\":500000,\"task_id\":\"task_Pgo719UBMORTMqlm13exGfTPs15b1ZuO\",\"upstream_model_name\":\"gpt-image-2.6-flare\"}"
  },
  {
    "id": 5,
    "user_id": 1,
    "created_at": 1794048068,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "OG-image-2.6-flare",
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
    "request_id": "20261107104108424370008268d9d6LPrAoMxs",
    "upstream_request_id": "20261107184108640142769BC1GZ9hd",
    "other": "{\"group_ratio\":1,\"is_model_mapped\":true,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\",\"upstream_model_name\":\"gpt-image-2.6-flare\"}"
  }
]
```
