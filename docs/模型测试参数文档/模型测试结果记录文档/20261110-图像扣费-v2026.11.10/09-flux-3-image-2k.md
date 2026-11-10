# 9. flux-3-image 2k

> 本文件由 `image_billing_batch_20261110.js` 生成。价格文档 `v2026.11.10.2`。

> 执行时间：2026/11/10 14:44:08  |  模型：`flux-3-image`  |  档位：2k

### 调用参数
```json
{
  "model": "flux-3-image",
  "prompt": "浩瀚星空下的沙漠",
  "resolution": "2k",
  "ratio": "16:9",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 2k | 分辨率档位 |
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
| 预期 Quota | **38,356** |
| 预期 USD | $0.0767 |
| 预期 RMB | ¥0.559998 |
| 结算方式 | 按张预扣，成功后不按 token 重算 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.37s |
| task_id | `task_jLS3tprT99pFwRrXsebsG8EzL0SoCAh3` |
| request_id | 202611100644083218410008268d9d6Ja42Z2W5 |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_jLS3tprT99pFwRrXsebsG8EzL0SoCAh3",
    "status": "NOT_START",
    "progress": "0%",
    "model": "flux-3-image",
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
| 任务首次查询 amount | 38,356 | 0.559998 |
| 消费日志 Quota | - | - |
| **预扣金额** | **38,356** | **¥0.559998** |
| 预期扣减 | 38,356 | ¥0.559998 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 10 |
| 完成耗时 | 72.3s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://getapib.org/image/9998205706904233-62b9e3ce-4b06-427d-af86-291e05114840-image_1794293092417328682.png

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 38,356 | ¥0.559998 |
| 平台记录最终消费 | 38,356 | **¥0.559998** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "code": 200,
  "data": {
    "actual_time": 47,
    "billing_status": "settled",
    "completed": 1794293095,
    "cost": 0.08,
    "created": 1794293048,
    "credits_cost": 0.8,
    "estimated_time": 100,
    "id": "task_01M723697BAJ1P7WPZ55D3V0F6",
    "progress": 100,
    "result": {
      "images": [
        {
          "expires_at": 1794379495,
          "url": [
            "https://getapib.org/image/9998205706904233-62b9e3ce-4b06-427d-af86-291e05114840-image_1794293092417328682.png"
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
    "task_id": "task_jLS3tprT99pFwRrXsebsG8EzL0SoCAh3",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "flux-3-image",
    "images": [
      {
        "url": "https://getapib.org/image/9998205706904233-62b9e3ce-4b06-427d-af86-291e05114840-image_1794293092417328682.png"
      }
    ],
    "amount": "0.559998",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794293048,
    "submit_time": 1794293048,
    "finish_time": 1794293120
  }
}
```

### 消费日志
```json
[]
```
