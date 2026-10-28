# doubao-seedance-2-5-260628 测试报告

## 50. doubao-seedance-2-5-260628 文生1080p
> 执行时间：2026-10-28  |  模型：`doubao-seedance-2-5-260628`

> resolution≈1.09524，有效单价 ¥80.50/1M

### 调用参数
```json
{
  "model": "doubao-seedance-2-5-260628",
  "prompt": "霓虹雨夜，摩托车驶过积水路面",
  "duration": 5,
  "resolution": "1080p",
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
| resolution | 1080p | 外层分辨率 |
| ratio | 16:9 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | true | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 | resolution≈1.09524，有效单价 ¥80.50/1M |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_KTIfDokQ8knELFEg0jlKxERh8aXQdcDU` |

```json
{
  "id": "task_KTIfDokQ8knELFEg0jlKxERh8aXQdcDU",
  "task_id": "task_KTIfDokQ8knELFEg0jlKxERh8aXQdcDU",
  "object": "video",
  "model": "doubao-seedance-2-5-260628",
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
| 平台记录最终消费 | 1,350,994 | ¥19.7245 |
| 是否符合 | ✅ 符合 | |
| request_id | - | |

> 折算单价 ¥80.50 / 1M tokens，用量 245025。

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 1 |
| progress | 100% |

**视频 URL**: https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-5/2122754587/Video/unv1z-83335182016202-tgc_808d9e4ccd51caffff00000000000000000000336818271397120/cgt-20261028153338-z1vnu.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T073620Z&X-Tos-Expires=86400&X-Tos-Max-Requests=100&X-Tos-Signature=b1eaaaefd350d9df14343b80d7008ae97a8085ff6508bd61d871976ccf99feeb&X-Tos-SignedHeaders=host

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 5 |
| 输出分辨率 | - |
| tokens | 245025 |

#### 查询调用参数

`GET /v1/video/generations/task_KTIfDokQ8knELFEg0jlKxERh8aXQdcDU`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 240,
    "created_at": 1793172818,
    "updated_at": 1793172995,
    "task_id": "task_KTIfDokQ8knELFEg0jlKxERh8aXQdcDU",
    "platform": "45",
    "user_id": 1,
    "group": "default",
    "channel_id": 16,
    "amount": "19.724512",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-5/2122754587/Video/unv1z-83335182016202-tgc_808d9e4ccd51caffff00000000000000000000336818271397120/cgt-20261028153338-z1vnu.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T073620Z&X-Tos-Expires=86400&X-Tos-Max-Requests=100&X-Tos-Signature=b1eaaaefd350d9df14343b80d7008ae97a8085ff6508bd61d871976ccf99feeb&X-Tos-SignedHeaders=host",
    "submit_time": 1793172818,
    "start_time": 1793172845,
    "finish_time": 1793172994,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-5-260628",
      "origin_model_name": "doubao-seedance-2-5-260628"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-5/2122754587/Video/unv1z-83335182016202-tgc_808d9e4ccd51caffff00000000000000000000336818271397120/cgt-20261028153338-z1vnu.mp4?X-Tos-Algorithm=TOS4-HMAC-SHA256&X-Tos-Credential=AKLTYWJkZTExNjA1ZDUyNDc3YzhjNTM5OGIyNjBhNDcyOTQ%2F20261028%2Fcn-beijing%2Ftos%2Frequest&X-Tos-Date=20261028T073620Z&X-Tos-Expires=86400&X-Tos-Max-Requests=100&X-Tos-Signature=b1eaaaefd350d9df14343b80d7008ae97a8085ff6508bd61d871976ccf99feeb&X-Tos-SignedHeaders=host"
      },
      "created_at": 1793172818,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261028153338-z1vnu",
      "model": "doubao-seedance-2-5-260628",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "1080p",
      "seed": 68036,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793172982,
      "usage": {
        "completion_tokens": 245025,
        "total_tokens": 245025
      }
    }
  }
}
```
