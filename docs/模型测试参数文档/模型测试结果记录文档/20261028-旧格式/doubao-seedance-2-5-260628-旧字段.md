# doubao-seedance-2-5-260628 旧格式测试报告

## doubao-seedance-2-5-260628 旧字段
> 执行时间：2026-10-28  |  模型：`doubao-seedance-2-5-260628`

> 覆盖旧字段：`size`、`images`、`seconds`、`duration`、`metadata.resolution`、`metadata.ratio`、`metadata.input_usage`。支持尾帧的模型另带 `metadata.last_frame`。

### 调用参数
```json
{
  "model": "doubao-seedance-2-5-260628",
  "prompt": "清晨海岸公路，镜头缓慢推进",
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
| 预期 | 不含视频 720p，有效单价 ¥73.50/1M |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 400（预期 200）❌ |
| task_id | `-` |

```json
{
  "code": "fail_to_fetch_task",
  "message": "{\"error\":{\"code\":\"InvalidParameter.TaskTypeConstraint\",\"message\":\"The parameter ratio specified in the request is not valid. For first-frame or first-last-frame generation, the output ratio follows the first-frame image. Request id: 0217931788012463d0449b03f13da7ad7c0359f0f1d760c753a7a\",\"param\":\"ratio\",\"type\":\"BadRequest\"}}",
  "data": null
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | - | - |
| 日志预扣 | - | - |
| 平台记录最终消费 | - | - |
| 是否符合 | ❌ 请求未通过 | |
| request_id | - | |

> {"error":{"code":"InvalidParameter.TaskTypeConstraint","message":"The parameter ratio specified in the request is not valid. For first-frame or first-last-frame generation, the output ratio follows the first-frame image. Request id: 0217931

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **-** |
| 轮询次数 | 0 |
| progress | - |

#### 查询调用参数

`GET /v1/video/generations/`

#### 最终查询响应（全量 JSON）
```json
null
```
