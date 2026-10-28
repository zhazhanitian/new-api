# hailuo-video-2.3 测试报告

## 29. hailuo-video-2.3 默认
> 执行时间：2026-10-28  |  模型：`hailuo-video-2.3`

> 表达式值 271230，quota 135615

### 调用参数
```json
{
  "model": "hailuo-video-2.3",
  "prompt": "雨中的咖啡馆橱窗，行人撑伞经过"
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| duration | 未传 | 表达式按请求秒数预扣，缺省走模型默认 |
| resolution | 未传 | 外层分辨率 |
| ratio | 未传 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | 未传 | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **135,615** |
| 预期 RMB | ¥1.9800 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_TSi5bi83unU0HwXuzlJKrwTqwoerREBq` |

```json
{
  "id": "task_TSi5bi83unU0HwXuzlJKrwTqwoerREBq",
  "task_id": "task_TSi5bi83unU0HwXuzlJKrwTqwoerREBq",
  "object": "video",
  "model": "hailuo-video-2.3",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172930
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 135,615 | ¥1.9800 |
| 日志预扣 | 135,615 | ¥1.9800 |
| 平台记录最终消费 | 132,789 | ¥1.9387 |
| 是否符合 | ✅ 符合 | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 3 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/948ce1285001834825866193428/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 5.875 |
| 输出分辨率 | 1366×768 |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_TSi5bi83unU0HwXuzlJKrwTqwoerREBq`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 253,
    "created_at": 1793172930,
    "updated_at": 1793173031,
    "task_id": "task_TSi5bi83unU0HwXuzlJKrwTqwoerREBq",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.938719",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/948ce1285001834825866193428/aigcVideoGenFile.mp4",
    "submit_time": 1793172930,
    "start_time": 1793172971,
    "finish_time": 1793173031,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-2.3",
      "origin_model_name": "hailuo-video-2.3"
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
            "FileInfos": [],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Hailuo",
            "ModelVersion": "2.3",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "",
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
              "Resolution": "768P",
              "StorageMode": "Temporary"
            },
            "Prompt": "雨中的咖啡馆橱窗，行人撑伞经过",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-11-04T07:37:11Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/948ce1285001834825866193428/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 2077992,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.875,
                  "Height": 768,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 1526026,
                  "VideoDuration": 5.875,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 2072910,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 768,
                      "Width": 1366
                    }
                  ],
                  "Width": 1366
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
          "TaskId": "1500065731-AigcVideoTask-141e6bc4ea8a16a466827d7654c5c77bt"
        },
        "BeginProcessTime": "2026-10-28T07:35:30Z",
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
        "CreateTime": "2026-10-28T07:35:30Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:37:06Z",
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
        "RequestId": "f8dfbe51-1b54-46d2-80dd-5a46d2808fe6",
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
