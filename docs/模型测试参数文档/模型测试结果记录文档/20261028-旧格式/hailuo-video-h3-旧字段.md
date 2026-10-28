# hailuo-video-h3 旧格式测试报告

## hailuo-video-h3 旧字段
> 执行时间：2026-10-28  |  模型：`hailuo-video-h3`

> 覆盖旧字段：`size`、`images`、`seconds`、`duration`、`metadata.resolution`、`metadata.ratio`、`metadata.input_usage`。支持尾帧的模型另带 `metadata.last_frame`。

### 调用参数
```json
{
  "model": "hailuo-video-h3",
  "prompt": "海边栈道，镜头跟着人物向前走",
  "size": "1080P",
  "seconds": "5",
  "duration": 5,
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"
  ],
  "metadata": {
    "resolution": "1080P",
    "ratio": "16:9",
    "input_usage": "FirstFrame",
    "last_frame": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"
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
| metadata.last_frame | 有 | 旧尾帧 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期表达式值 | 445,205 |
| 预期 Quota | **222,603** |
| 预期 RMB | ¥3.2500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_4Glmuz9AeZmBi6aEqrgrhaSpTrcIpi88` |

```json
{
  "id": "task_4Glmuz9AeZmBi6aEqrgrhaSpTrcIpi88",
  "task_id": "task_4Glmuz9AeZmBi6aEqrgrhaSpTrcIpi88",
  "object": "video",
  "model": "hailuo-video-h3",
  "status": "queued",
  "progress": 0,
  "created_at": 1793178801
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 222,603 | ¥3.2500 |
| 日志预扣 | 222,603 | ¥3.2500 |
| 平台记录最终消费 | 230,839 | ¥3.3702 |
| 是否符合 | ✅ 预扣符合 | |
| request_id | 20261028091321497030008268d9d6UaiB6LW9 | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 18 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/45eb4d635001834825875910282/aigcVideoGenFile.mp4

#### 查询调用参数

`GET /v1/video/generations/task_4Glmuz9AeZmBi6aEqrgrhaSpTrcIpi88`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 277,
    "created_at": 1793178801,
    "updated_at": 1793179045,
    "task_id": "task_4Glmuz9AeZmBi6aEqrgrhaSpTrcIpi88",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.370249",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/45eb4d635001834825875910282/aigcVideoGenFile.mp4",
    "submit_time": 1793178801,
    "start_time": 1793178810,
    "finish_time": 1793179045,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-h3",
      "origin_model_name": "hailuo-video-h3"
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
              },
              {
                "Category": "Image",
                "FileId": "",
                "KeepOriginalSound": "",
                "ObjectId": "",
                "ReferenceType": "",
                "Text": "",
                "Type": "Url",
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg",
                "Usage": "LastFrame",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Hailuo",
            "ModelVersion": "H3",
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
              "Resolution": "1080P",
              "StorageMode": "Temporary"
            },
            "Prompt": "海边栈道，镜头跟着人物向前走",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-11-04T09:17:25Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/45eb4d635001834825875910282/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 5.184,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 97436,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 44100
                    }
                  ],
                  "Bitrate": 4586564,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.185,
                  "Height": 1890,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 2972667,
                  "VideoDuration": 5.167,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 4494675,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 1890,
                      "Width": 1080
                    }
                  ],
                  "Width": 1080
                },
                "StorageMode": "Temporary",
                "UsageType": ""
              }
            ],
            "ProcedureTaskIds": [],
            "Usage": {
              "InputImageCount": 2,
              "InputSeconds": 0,
              "InputTokens": 0,
              "OutputSeconds": 5,
              "ThoughtTokens": 0,
              "TotalSeconds": 5
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-1cfcdb5145753e0edef69d284a7ed44dt"
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
        "FinishTime": "2026-10-28T09:17:08Z",
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
        "RequestId": "0c43983e-36e4-4796-b430-b1d518b5d525",
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
