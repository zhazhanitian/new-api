# sora-video-2.0 测试报告

## 39. sora-video-2.0 duration=12
> 执行时间：2026-10-28  |  模型：`sora-video-2.0`

> 表达式值 1232880，quota 616440

### 调用参数
```json
{
  "model": "sora-video-2.0",
  "prompt": "纸船在水面上漂过桥洞",
  "duration": 12,
  "resolution": "720P",
  "ratio": "16:9"
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| duration | 12 | 表达式按请求秒数预扣，缺省走模型默认 |
| resolution | 720P | 外层分辨率 |
| ratio | 16:9 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | 未传 | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **616,440** |
| 预期 RMB | ¥9.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_NvWCzubRycNuFYTH2htGzNMht3GuyUF3` |

```json
{
  "id": "task_NvWCzubRycNuFYTH2htGzNMht3GuyUF3",
  "task_id": "task_NvWCzubRycNuFYTH2htGzNMht3GuyUF3",
  "object": "video",
  "model": "sora-video-2.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172935
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 616,440 | ¥9.0000 |
| 日志预扣 | - | - |
| 平台记录最终消费 | 621,577 | ¥9.0750 |
| 是否符合 | ✅ 符合（实扣含成片秒数微调） | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 8 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/e80a26a95001834825867427653/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 12.1 |
| 输出分辨率 | 1280×720 |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_NvWCzubRycNuFYTH2htGzNMht3GuyUF3`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 263,
    "created_at": 1793172935,
    "updated_at": 1793173106,
    "task_id": "task_NvWCzubRycNuFYTH2htGzNMht3GuyUF3",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "9.075024",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/e80a26a95001834825867427653/aigcVideoGenFile.mp4",
    "submit_time": 1793172935,
    "start_time": 1793172982,
    "finish_time": 1793173106,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "sora-video-2.0",
      "origin_model_name": "sora-video-2.0"
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
            "ModelName": "OS",
            "ModelVersion": "2.0",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 12,
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
              "Resolution": "720P",
              "StorageMode": "Temporary"
            },
            "Prompt": "纸船在水面上漂过桥洞",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-11-04T07:38:26Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/e80a26a95001834825867427653/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 12.1,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 127689,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 96000
                    }
                  ],
                  "Bitrate": 5629276,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 12.1,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 8514281,
                  "VideoDuration": 12,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 5532731,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 30,
                      "Height": 720,
                      "Width": 1280
                    }
                  ],
                  "Width": 1280
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
              "OutputSeconds": 12,
              "ThoughtTokens": 0,
              "TotalSeconds": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-3a7c512ca01919212c0efd69b9257db2t"
        },
        "BeginProcessTime": "2026-10-28T07:35:35Z",
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
        "CreateTime": "2026-10-28T07:35:35Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:38:04Z",
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
        "RequestId": "d4246a8d-23ee-44c3-ae57-61a976aa07cb",
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
