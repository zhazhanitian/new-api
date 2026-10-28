# pixverse-video-v5.6 测试报告

## 44. pixverse-video-v5.6 1080P x 10秒
> 执行时间：2026-10-28  |  模型：`pixverse-video-v5.6`

> 表达式值 719180，quota 359590

### 调用参数
```json
{
  "model": "pixverse-video-v5.6",
  "prompt": "一只鸟从枝头飞起",
  "duration": 10,
  "resolution": "1080P",
  "ratio": "16:9"
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| duration | 10 | 表达式按请求秒数预扣，缺省走模型默认 |
| resolution | 1080P | 外层分辨率 |
| ratio | 16:9 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | 未传 | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **359,590** |
| 预期 RMB | ¥5.2500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_qVWguXDcGAboVhaHY6eooIKM3VYkcvFj` |

```json
{
  "id": "task_qVWguXDcGAboVhaHY6eooIKM3VYkcvFj",
  "task_id": "task_qVWguXDcGAboVhaHY6eooIKM3VYkcvFj",
  "object": "video",
  "model": "pixverse-video-v5.6",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172938
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 359,590 | ¥5.2500 |
| 日志预扣 | - | - |
| 平台记录最终消费 | 359,590 | ¥5.2500 |
| 是否符合 | ✅ 预扣符合，任务失败后退款 | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **FAILURE** |
| 轮询次数 | 4 |
| progress | 100% |
| 失败原因 | task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds. |

**视频 URL**: task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | - |
| 输出分辨率 | - |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_qVWguXDcGAboVhaHY6eooIKM3VYkcvFj`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 268,
    "created_at": 1793172938,
    "updated_at": 1793173047,
    "task_id": "task_qVWguXDcGAboVhaHY6eooIKM3VYkcvFj",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "5.250014",
    "action": "generate",
    "status": "FAILURE",
    "fail_reason": "task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.",
    "result_url": "task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.",
    "submit_time": 1793172938,
    "start_time": 1793172988,
    "finish_time": 1793173047,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "pixverse-video-v5.6",
      "origin_model_name": "pixverse-video-v5.6"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcHunyuan3DTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ChannelErrCode": "",
          "ErrCode": 70000,
          "ErrCodeExt": "InvalidParameter",
          "Input": {
            "EnhancePrompt": "",
            "FileInfos": [],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "PixVerse",
            "ModelVersion": "v5.6",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 10,
              "EnableBGM": "",
              "EnhanceSwitch": "",
              "ExpireTime": "0000-00-00T00:00:00Z",
              "FrameInterpolate": "",
              "InputComplianceCheck": "",
              "LogoAdd": "",
              "MediaName": "",
              "OffPeak": "",
              "OutputComplianceCheck": "",
              "PersonGeneration": "",
              "Resolution": "1080p",
              "StorageMode": "Temporary"
            },
            "Prompt": "一只鸟从枝头飞起",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.",
          "Output": {
            "FileInfos": [],
            "ProcedureTaskIds": [],
            "Usage": {
              "InputImageCount": 0,
              "InputSeconds": 0,
              "InputTokens": 0,
              "OutputSeconds": 0,
              "ThoughtTokens": 0,
              "TotalSeconds": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-cea82edf24d386b2c020a072307f61dat"
        },
        "BeginProcessTime": "2026-10-28T07:35:38Z",
        "ClipTask": null,
        "CloneVoiceAsyncTask": null,
        "ComplexAdaptiveDynamicStreamingTask": null,
        "ComposeMediaTask": null,
        "ConcatTask": null,
        "CreateAigcAdvancedCustomElementTask": null,
        "CreateAigcAudioCloneTask": null,
        "CreateAigcCustomVoiceTask": null,
        "CreateAigcMaterialTask": null,
        "CreateAigcSubjectTask": null,
        "CreateImageSpriteTask": null,
        "CreateTime": "2026-10-28T07:35:38Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:37:17Z",
        "ImportMediaKnowledge": null,
        "ProcedureTask": null,
        "ProcessImageAsyncTask": null,
        "ProcessMediaByMPSTask": null,
        "PullUploadTask": null,
        "QualityEnhanceTask": null,
        "QualityInspectTask": null,
        "RebuildMediaTask": null,
        "ReduceMediaBitrateTask": null,
        "RemoveWatermarkTask": null,
        "RequestId": "514ae595-3521-45cf-bc58-fabd6bb5ed8a",
        "ReviewAudioVideoTask": null,
        "SceneAigcImageTask": null,
        "SceneAigcVideoTask": null,
        "SearchMediaBySemanticsAsync": null,
        "SnapshotByTimeOffsetTask": null,
        "SplitMediaTask": null,
        "Status": "FINISH",
        "TaskType": "AigcVideoTask",
        "TextToSpeechAsyncTask": null,
        "TranscodeTask": null,
        "VideoDepthEstimationTask": null,
        "VideoDubbingAsyncTask": null,
        "WandAsrTask": null,
        "WechatMiniProgramPublishTask": null,
        "WechatPublishTask": null
      }
    }
  }
}
```
