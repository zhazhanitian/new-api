# 15. mai-image-2.6 1K参考图

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:08  |  模型：`mai-image-2.6`  |  档位：1K，参考图不加价

### 调用参数
```json
{
  "model": "mai-image-2.6",
  "prompt": "按参考图风格画一只橘猫",
  "resolution": "1K",
  "ratio": "1:1",
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
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 1 张 | 参考图 |
| extra_fields.width | 未传 | 成对宽，优先于 size |
| extra_fields.height | 未传 | 成对高，优先于 size |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **14,926** |
| 预期 USD | $0.0299 |
| 预期 RMB | ¥0.217920 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.35s |
| task_id | `task_900QeKImaTU8UTAwFQuJy3Z8h42bJ8Pu` |
| request_id | 202611100644086218750008268d9d6B0vrJmXD |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_900QeKImaTU8UTAwFQuJy3Z8h42bJ8Pu",
    "status": "NOT_START",
    "progress": "0%",
    "model": "mai-image-2.6",
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
| 任务首次查询 amount | 14,926 | 0.217920 |
| 消费日志 Quota | - | - |
| **预扣金额** | **14,926** | **¥0.217920** |
| 预期扣减 | 14,926 | ¥0.217920 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 11 |
| 完成耗时 | 80.3s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706923641-a84aeed7-1073-41a7-973b-626b9a8c4d74-image_task_01M72374C2BC3ZYXAKWFVJ1THP_0.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 14,926 | ¥0.217920 |
| 平台记录最终消费 | 14,926 | **¥0.217920** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 28,
    "billing_status": "settled",
    "completed": 1794293076,
    "cost": 0.032768,
    "created": 1794293048,
    "credits_cost": 0.32767999999999997,
    "estimated_time": 60,
    "id": "task_01M72369K7A49PPJWFPFP0PMKP",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379476,
          "url": [
            "https://getapib.org/image/9998205706923641-a84aeed7-1073-41a7-973b-626b9a8c4d74-image_task_01M72374C2BC3ZYXAKWFVJ1THP_0.png"
          ]
        }
      ]
    },
    "status": "completed",
    "usage": {
      "input_tokens": 256,
      "input_tokens_details": {
        "cached_tokens": 0,
        "image_tokens": 256,
        "text_tokens": 0
      },
      "output_tokens": 1024,
      "output_tokens_details": {
        "image_tokens": 1024,
        "text_tokens": 0
      },
      "total_tokens": 1280
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
    "task_id": "task_900QeKImaTU8UTAwFQuJy3Z8h42bJ8Pu",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "mai-image-2.6",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706923641-a84aeed7-1073-41a7-973b-626b9a8c4d74-image_task_01M72374C2BC3ZYXAKWFVJ1THP_0.png"
      }
    ],
    "amount": "0.217920",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293048,
    "submit_time": 1794293048,
    "finish_time": 1794293126
  }
}
```

### 消费日志
```json
[]
```
