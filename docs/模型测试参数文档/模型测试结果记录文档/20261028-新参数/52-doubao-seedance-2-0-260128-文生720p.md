# doubao-seedance-2-0-260128 测试报告

## 52. doubao-seedance-2-0-260128 文生720p
> 执行时间：2026-10-28  |  模型：`doubao-seedance-2-0-260128`

> 无系数，有效单价 ¥48.30/1M

### 调用参数
```json
{
  "model": "doubao-seedance-2-0-260128",
  "prompt": "山间公路，云雾缓慢流动",
  "duration": 5,
  "resolution": "720p",
  "ratio": "16:9",
  "metadata": {
    "generate_audio": true
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| duration | 5 | 表达式按请求秒数预扣，缺省走模型默认 |
| resolution | 720p | 外层分辨率 |
| ratio | 16:9 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | true | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 | 无系数，有效单价 ¥48.30/1M |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_RJsRht8EX1XeYXqg3Gthc9kVrSiHer2Z` |

```json
{
  "id": "task_RJsRht8EX1XeYXqg3Gthc9kVrSiHer2Z",
  "task_id": "task_RJsRht8EX1XeYXqg3Gthc9kVrSiHer2Z",
  "object": "video",
  "model": "doubao-seedance-2-0-260128",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172818
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | - | - |
| 日志预扣 | - | - |
| 平台记录最终消费 | 500,939 | ¥7.3137 |
| 是否符合 | ❌ 不符 | |
| request_id | - | |

> 折算单价 ¥67.16 / 1M tokens，用量 108900。

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 1 |
| progress | 100% |

**视频 URL**: https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317281874300000000000000000000ffffac141a6155b9b1.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T073650Z&X-Tos-Expires=86400&X-Tos-Signature=360e65ad733e282cff18f4977b12061ca587c1054726ec4cbadc74df2ef8c376&X-Tos-SignedHeaders=host

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 5 |
| 输出分辨率 | - |
| tokens | 108900 |

#### 查询调用参数

`GET /v1/video/generations/task_RJsRht8EX1XeYXqg3Gthc9kVrSiHer2Z`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 244,
    "created_at": 1793172818,
    "updated_at": 1793173069,
    "task_id": "task_RJsRht8EX1XeYXqg3Gthc9kVrSiHer2Z",
    "platform": "45",
    "user_id": 1,
    "group": "default",
    "channel_id": 16,
    "amount": "7.313709",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317281874300000000000000000000ffffac141a6155b9b1.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T073650Z&X-Tos-Expires=86400&X-Tos-Signature=360e65ad733e282cff18f4977b12061ca587c1054726ec4cbadc74df2ef8c376&X-Tos-SignedHeaders=host",
    "submit_time": 1793172818,
    "start_time": 1793172849,
    "finish_time": 1793173069,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-0-260128",
      "origin_model_name": "doubao-seedance-2-0-260128"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317281874300000000000000000000ffffac141a6155b9b1.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T073650Z&X-Tos-Expires=86400&X-Tos-Signature=360e65ad733e282cff18f4977b12061ca587c1054726ec4cbadc74df2ef8c376&X-Tos-SignedHeaders=host"
      },
      "created_at": 1793172818,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261028153338-9n5nq",
      "model": "doubao-seedance-2-0-260128",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "720p",
      "seed": 36391,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793173015,
      "usage": {
        "completion_tokens": 108900,
        "total_tokens": 108900
      }
    }
  }
}
```
