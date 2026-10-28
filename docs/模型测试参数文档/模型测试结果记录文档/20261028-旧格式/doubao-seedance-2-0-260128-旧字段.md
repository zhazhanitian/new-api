# doubao-seedance-2-0-260128 旧格式测试报告

## doubao-seedance-2-0-260128 旧字段
> 执行时间：2026-10-28  |  模型：`doubao-seedance-2-0-260128`

> 覆盖旧字段：`size`、`images`、`seconds`、`duration`、`metadata.resolution`、`metadata.ratio`、`metadata.input_usage`。支持尾帧的模型另带 `metadata.last_frame`。

### 调用参数
```json
{
  "model": "doubao-seedance-2-0-260128",
  "prompt": "山间公路，云雾缓慢流动",
  "size": "720p",
  "seconds": "5",
  "duration": 5,
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"
  ],
  "metadata": {
    "resolution": "720p",
    "ratio": "16:9",
    "input_usage": "FirstFrame"
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| size | 720p | 旧分辨率。外层 resolution 为空时收成 resolution |
| seconds / duration | 5 | 旧时长。同时存在时上游以 seconds 为准，表达式读 duration |
| images | 1 张 | 旧图片字段。input_usage=FirstFrame 时第一张当首帧 |
| metadata.resolution | 720p | 与 size 同值，确认旧字段仍被接受 |
| metadata.ratio | 16:9 | 旧宽高比 |
| metadata.input_usage | FirstFrame | 旧首帧标记 |
| metadata.last_frame | 未传 | 旧尾帧 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 | 不含视频 720p。文档新价 ¥48.30/1M；若后台仍是旧 1.5 倍则为 ¥67.16/1M |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_VlVI2TeCWZrobFWatfLxoETYlzYSZfEE` |

```json
{
  "id": "task_VlVI2TeCWZrobFWatfLxoETYlzYSZfEE",
  "task_id": "task_VlVI2TeCWZrobFWatfLxoETYlzYSZfEE",
  "object": "video",
  "model": "doubao-seedance-2-0-260128",
  "status": "queued",
  "progress": 0,
  "created_at": 1793178802
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | - | - |
| 日志预扣 | 827,050 | ¥12.0749 |
| 平台记录最终消费 | 360,262 | ¥5.2598 |
| 是否符合 | ✅ 请求通过（豆包按 token 结算，见最终消费） | |
| request_id | 20261028091321500110008268d9d6iOczlJp2 | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 19 |
| progress | 100% |

**视频 URL**: https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317880212700000000000000000000ffffac1802a36ae757.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T091722Z&X-Tos-Expires=86400&X-Tos-Signature=e447a62e5d77e46e68c0f7a0f0c8e30311cf3884199a677d13d430ba4904179a&X-Tos-SignedHeaders=host

#### 查询调用参数

`GET /v1/video/generations/task_VlVI2TeCWZrobFWatfLxoETYlzYSZfEE`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 286,
    "created_at": 1793178802,
    "updated_at": 1793179067,
    "task_id": "task_VlVI2TeCWZrobFWatfLxoETYlzYSZfEE",
    "platform": "45",
    "user_id": 1,
    "group": "default",
    "channel_id": 16,
    "amount": "5.259825",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317880212700000000000000000000ffffac1802a36ae757.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T091722Z&X-Tos-Expires=86400&X-Tos-Signature=e447a62e5d77e46e68c0f7a0f0c8e30311cf3884199a677d13d430ba4904179a&X-Tos-SignedHeaders=host",
    "submit_time": 1793178802,
    "start_time": 1793178820,
    "finish_time": 1793179067,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-0-260128",
      "origin_model_name": "doubao-seedance-2-0-260128"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317880212700000000000000000000ffffac1802a36ae757.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T091722Z&X-Tos-Expires=86400&X-Tos-Signature=e447a62e5d77e46e68c0f7a0f0c8e30311cf3884199a677d13d430ba4904179a&X-Tos-SignedHeaders=host"
      },
      "created_at": 1793178802,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261028171321-lshx7",
      "model": "doubao-seedance-2-0-260128",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "720p",
      "seed": 55192,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793179050,
      "usage": {
        "completion_tokens": 108900,
        "total_tokens": 108900
      }
    }
  }
}
```
