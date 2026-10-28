# doubao-seedance-1-5-pro-251215 测试报告

## 60. doubao-seedance-1-5-pro-251215 无声
> 执行时间：2026-10-28  |  模型：`doubao-seedance-1-5-pro-251215`

> generate_audio=false，silent_video=0.5，¥8.00/1M

### 调用参数
```json
{
  "model": "doubao-seedance-1-5-pro-251215",
  "prompt": "海边栈道，镜头跟着人物向前",
  "duration": 5,
  "resolution": "720p",
  "ratio": "16:9",
  "metadata": {
    "generate_audio": false
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
| generate_audio | false | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 | generate_audio=false，silent_video=0.5，¥8.00/1M |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 404（预期 200）❌ |
| task_id | `-` |

```json
{
  "code": "fail_to_fetch_task",
  "message": "{\"error\":{\"code\":\"InvalidEndpointOrModel.NotFound\",\"message\":\"The model or endpoint doubao-seedance-1-5-pro-251215 does not exist or you do not have access to it. Request id: 0217931728182258df8e1a1401bb5aa62b5b494646ada581b2d98\",\"param\":\"\",\"type\":\"Not Found\"}}",
  "data": null
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | - | - |
| 日志预扣 | - | - |
| 平台记录最终消费 | - | - |
| 是否符合 | ❌ 未生成任务 | |
| request_id | - | |

> {"error":{"code":"InvalidEndpointOrModel.NotFound","message":"The model or endpoint doubao-seedance-1-5-pro-251215 does not exist or you do not have access to it. Request id: 0217931728182258df8e1a1401bb5aa62b5b494646ada

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **-** |
| 轮询次数 | 0 |
| progress | - |

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | - |
| 输出分辨率 | - |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/`

#### 最终查询响应（全量 JSON）
```json
null
```
