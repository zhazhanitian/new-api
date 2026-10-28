# doubao-seedance-2-0-260128 测试报告

## 53. doubao-seedance-2-0-260128 文生4k
> 执行时间：2026-10-28  |  模型：`doubao-seedance-2-0-260128`

> resolution≈0.5652，有效单价 ¥27.30/1M

### 调用参数
```json
{
  "model": "doubao-seedance-2-0-260128",
  "prompt": "山间公路，云雾缓慢流动",
  "duration": 5,
  "resolution": "4k",
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
| resolution | 4k | 外层分辨率 |
| ratio | 16:9 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | true | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 | resolution≈0.5652，有效单价 ¥27.30/1M |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_q7ryAl8b7aZ3uMJlhe3caXjevCyQdPaD` |

```json
{
  "id": "task_q7ryAl8b7aZ3uMJlhe3caXjevCyQdPaD",
  "task_id": "task_q7ryAl8b7aZ3uMJlhe3caXjevCyQdPaD",
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
| 平台记录最终消费 | 2,548,260 | ¥37.2046 |
| 是否符合 | ❌ 不符 | |
| request_id | - | |

> 折算单价 ¥37.96 / 1M tokens，用量 980100。

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 6 |
| progress | 100% |

**视频 URL**: https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317281870200000000000000000000ffffac175b6f924f42.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T074051Z&X-Tos-Expires=86400&X-Tos-Signature=bda310d99d5f43e8e50f74dd9373a4a96dc45058cea7ab08e1a25c0220d7c858&X-Tos-SignedHeaders=host

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 5 |
| 输出分辨率 | - |
| tokens | 980100 |

#### 查询调用参数

`GET /v1/video/generations/task_q7ryAl8b7aZ3uMJlhe3caXjevCyQdPaD`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 243,
    "created_at": 1793172818,
    "updated_at": 1793173267,
    "task_id": "task_q7ryAl8b7aZ3uMJlhe3caXjevCyQdPaD",
    "platform": "45",
    "user_id": 1,
    "group": "default",
    "channel_id": 16,
    "amount": "37.204596",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317281870200000000000000000000ffffac175b6f924f42.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T074051Z&X-Tos-Expires=86400&X-Tos-Signature=bda310d99d5f43e8e50f74dd9373a4a96dc45058cea7ab08e1a25c0220d7c858&X-Tos-SignedHeaders=host",
    "submit_time": 1793172818,
    "start_time": 1793172848,
    "finish_time": 1793173267,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-0-260128",
      "origin_model_name": "doubao-seedance-2-0-260128"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0/02179317281870200000000000000000000ffffac175b6f924f42.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T074051Z&X-Tos-Expires=86400&X-Tos-Signature=bda310d99d5f43e8e50f74dd9373a4a96dc45058cea7ab08e1a25c0220d7c858&X-Tos-SignedHeaders=host"
      },
      "created_at": 1793172818,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261028153338-gwgcn",
      "model": "doubao-seedance-2-0-260128",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "4k",
      "seed": 71860,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793173261,
      "usage": {
        "completion_tokens": 980100,
        "total_tokens": 980100
      }
    }
  }
}
```
