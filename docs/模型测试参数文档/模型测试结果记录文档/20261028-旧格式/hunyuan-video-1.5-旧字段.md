# hunyuan-video-1.5 旧格式测试报告

## hunyuan-video-1.5 旧字段
> 执行时间：2026-10-28  |  模型：`hunyuan-video-1.5`

> 覆盖旧字段：`size`、`images`、`seconds`、`duration`、`metadata.resolution`、`metadata.ratio`、`metadata.input_usage`。支持尾帧的模型另带 `metadata.last_frame`。

### 调用参数
```json
{
  "model": "hunyuan-video-1.5",
  "prompt": "江南水乡，乌篷船慢慢划过",
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
| 预期表达式值 | 342,465 |
| 预期 Quota | **171,233** |
| 预期 RMB | ¥2.5000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_ojz3OA5Z3yFCxG2TjjLRSpfpAOe4Fkid` |

```json
{
  "id": "task_ojz3OA5Z3yFCxG2TjjLRSpfpAOe4Fkid",
  "task_id": "task_ojz3OA5Z3yFCxG2TjjLRSpfpAOe4Fkid",
  "object": "video",
  "model": "hunyuan-video-1.5",
  "status": "queued",
  "progress": 0,
  "created_at": 1793178801
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 171,233 | ¥2.5000 |
| 日志预扣 | 171,233 | ¥2.5000 |
| 平台记录最终消费 | 172,671 | ¥2.5210 |
| 是否符合 | ✅ 预扣符合 | |
| request_id | 20261028091321498160008268d9d6XnAuOTwz | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 12 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/87986a885001834825876464246/aigcVideoGenFile.mp4

#### 查询调用参数

`GET /v1/video/generations/task_ojz3OA5Z3yFCxG2TjjLRSpfpAOe4Fkid`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 274,
    "created_at": 1793178801,
    "updated_at": 1793178956,
    "task_id": "task_ojz3OA5Z3yFCxG2TjjLRSpfpAOe4Fkid",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "2.520997",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/87986a885001834825876464246/aigcVideoGenFile.mp4",
    "submit_time": 1793178801,
    "start_time": 1793178807,
    "finish_time": 1793178956,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hunyuan-video-1.5",
      "origin_model_name": "hunyuan-video-1.5"
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
            "ModelName": "Hunyuan",
            "ModelVersion": "1.5",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 0,
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
              "Resolution": "",
              "StorageMode": "Temporary"
            },
            "Prompt": "江南水乡，乌篷船慢慢划过",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-11-04T09:15:56Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/87986a885001834825876464246/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 785964,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 1280,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 495354,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 780219,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 1280,
                      "Width": 720
                    }
                  ],
                  "Width": 720
                },
                "StorageMode": "Temporary",
                "UsageType": ""
              }
            ],
            "ProcedureTaskIds": [],
            "Usage": {
              "InputImageCount": 0,
              "InputSeconds": 0,
              "InputTokens": 0,
              "OutputSeconds": 5,
              "ThoughtTokens": 0,
              "TotalSeconds": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-195228213a2eeb06af1be7eddfc8293ft"
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
        "FinishTime": "2026-10-28T09:15:56Z",
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
        "RequestId": "aaee8b17-506a-4a12-964b-219ea56d6734",
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
