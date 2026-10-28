# kling-video-3.0-omni 测试报告

## 14. kling-video-3.0-omni 参考视频1080P x 5秒
> 执行时间：2026-10-28  |  模型：`kling-video-3.0-omni`

> 表达式值 821920，quota 410960

### 调用参数
```json
{
  "model": "kling-video-3.0-omni",
  "prompt": "沿用参考视频的主体动作",
  "duration": 5,
  "resolution": "1080P",
  "metadata": {
    "audio_generation": "Enabled",
    "content": [
      {
        "type": "video_url",
        "role": "reference_video",
        "video_url": {
          "url": "https://aigc-bucket.su.bcebos.com/v1/objects/0d/0dca41d01e3579e43d30ae64fd30126aa056963fe85c5f254b81a3868a37298d?authorization=bce-auth-v1%2FALTAKMD70FiEhTeFimgJH1yjAb%2F2026-10-28T07%3A26%3A04Z%2F600%2F%2F03d1d6955391e462eadcb73c5ef846a63b2ba1543d2d63bdbae041429f5767f2"
        }
      }
    ]
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| duration | 5 | 表达式按请求秒数预扣，缺省走模型默认 |
| resolution | 1080P | 外层分辨率 |
| ratio | 未传 | 不参与本轮计价 |
| audio_generation | Enabled | 仅 Enabled 走有声价 |
| generate_audio | 未传 | 豆包 1.5 显式 false 才无声 |
| content | video_url/reference_video | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **410,960** |
| 预期 RMB | ¥6.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 500（预期 200）❌ |
| task_id | `-` |

```json
{
  "code": "InvalidParameterValue",
  "message": "InvalidParameterValue: AudioGeneration Enabled is not supported when video reference is provided for model Kling 3.0-Omni",
  "data": null
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 410,960 | ¥6.0000 |
| 日志预扣 | - | - |
| 平台记录最终消费 | - | - |
| 是否符合 | ❌ 未生成任务 | |
| request_id | - | |

> InvalidParameterValue: AudioGeneration Enabled is not supported when video reference is provided for model Kling 3.0-Omni

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
