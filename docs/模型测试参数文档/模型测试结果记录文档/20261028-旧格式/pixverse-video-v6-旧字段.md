# pixverse-video-v6 旧格式测试报告

## pixverse-video-v6 旧字段
> 执行时间：2026-10-28  |  模型：`pixverse-video-v6`

> 覆盖旧字段：`size`、`images`、`seconds`、`duration`、`metadata.resolution`、`metadata.ratio`、`metadata.input_usage`。支持尾帧的模型另带 `metadata.last_frame`。

### 调用参数
```json
{
  "model": "pixverse-video-v6",
  "prompt": "风吹过麦田，麦浪起伏",
  "size": "1080P",
  "seconds": "5",
  "duration": 5,
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"
  ],
  "metadata": {
    "resolution": "1080P",
    "ratio": "16:9",
    "input_usage": "FirstFrame"
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| size | 1080P | 旧分辨率。外层 resolution 为空时收成 resolution |
| seconds / duration | 5 | 旧时长。同时存在时上游以 seconds 为准，表达式读 duration |
| images | 1 张 | 旧图片字段。input_usage=FirstFrame 时第一张当首帧 |
| metadata.resolution | 1080P | 与 size 同值，确认旧字段仍被接受 |
| metadata.ratio | 16:9 | 旧宽高比 |
| metadata.input_usage | FirstFrame | 旧首帧标记 |
| metadata.last_frame | 未传 | 旧尾帧 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期表达式值 | 361,645 |
| 预期 Quota | **180,823** |
| 预期 RMB | ¥2.6400 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_H6UaGFNMRq0PskATs9OpicOu4OncOoqK` |

```json
{
  "id": "task_H6UaGFNMRq0PskATs9OpicOu4OncOoqK",
  "task_id": "task_H6UaGFNMRq0PskATs9OpicOu4OncOoqK",
  "object": "video",
  "model": "pixverse-video-v6",
  "status": "queued",
  "progress": 0,
  "created_at": 1793178801
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 180,823 | ¥2.6400 |
| 日志预扣 | 180,823 | ¥2.6400 |
| 平台记录最终消费 | 180,823 | ¥2.6400 |
| 是否符合 | ✅ 预扣符合 | |
| request_id | 20261028091321497970008268d9d6Iu9QWiaj | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **TIMEOUT** |
| 轮询次数 | 48 |
| progress | 0% |

#### 查询调用参数

`GET /v1/video/generations/task_H6UaGFNMRq0PskATs9OpicOu4OncOoqK`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 284,
    "created_at": 1793178801,
    "updated_at": 1793179496,
    "task_id": "task_H6UaGFNMRq0PskATs9OpicOu4OncOoqK",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "2.640016",
    "action": "generate",
    "status": "IN_PROGRESS",
    "fail_reason": "",
    "submit_time": 1793178801,
    "start_time": 1793178818,
    "finish_time": 0,
    "progress": "0%",
    "properties": {
      "input": "",
      "upstream_model_name": "pixverse-video-v6",
      "origin_model_name": "pixverse-video-v6"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcHunyuan3DTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ChannelErrCode": "",
          "ErrCode": 0,
          "ErrCodeExt": "",
          "Input": {
            "EnhancePrompt": "",
            "FileInfos": [
              {
                "Category": "Image",
                "FileId": "",
                "KeepOriginalSound": "",
                "ObjectId": "",
                "ReferenceType": "",
                "Text": "",
                "Type": "Url",
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg",
                "Usage": "FirstFrame",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "PixVerse",
            "ModelVersion": "v6",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 5,
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
            "Prompt": "风吹过麦田，麦浪起伏",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
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
          "Progress": 0,
          "SessionContext": "",
          "SessionId": "",
          "Status": "PROCESSING",
          "TaskId": "1500065731-AigcVideoTask-a0f38ae4785bab8224cf75f93aa1b755t"
        },
        "BeginProcessTime": "2026-10-28T09:13:21Z",
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
        "CreateTime": "2026-10-28T09:13:21Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "0001-01-01T00:00:00Z",
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
        "RequestId": "05620124-4125-4682-9a3d-a7c51dea1a8e",
        "ReviewAudioVideoTask": null,
        "SceneAigcImageTask": null,
        "SceneAigcVideoTask": null,
        "SearchMediaBySemanticsAsync": null,
        "SnapshotByTimeOffsetTask": null,
        "SplitMediaTask": null,
        "Status": "PROCESSING",
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
