# vidu-video-q3 测试报告

## 22. vidu-video-q3 1080P x 8秒
> 执行时间：2026-10-28  |  模型：`vidu-video-q3`

> 表达式值 856984，quota 428492

### 调用参数
```json
{
  "model": "vidu-video-q3",
  "prompt": "森林小径，光斑随风移动",
  "duration": 8,
  "resolution": "1080P",
  "ratio": "16:9"
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| duration | 8 | 表达式按请求秒数预扣，缺省走模型默认 |
| resolution | 1080P | 外层分辨率 |
| ratio | 16:9 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | 未传 | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **428,492** |
| 预期 RMB | ¥6.2560 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_renJ2EC1QMXIxirtCKCXprdDFxPSu36u` |

```json
{
  "id": "task_renJ2EC1QMXIxirtCKCXprdDFxPSu36u",
  "task_id": "task_renJ2EC1QMXIxirtCKCXprdDFxPSu36u",
  "object": "video",
  "model": "vidu-video-q3",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172926
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 428,492 | ¥6.2560 |
| 日志预扣 | 428,492 | ¥6.2560 |
| 平台记录最终消费 | 430,741 | ¥6.2888 |
| 是否符合 | ✅ 符合 | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 3 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/1e1427815001834825932758213/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 8.042 |
| 输出分辨率 | 1920×1080 |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_renJ2EC1QMXIxirtCKCXprdDFxPSu36u`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 246,
    "created_at": 1793172926,
    "updated_at": 1793173023,
    "task_id": "task_renJ2EC1QMXIxirtCKCXprdDFxPSu36u",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.288819",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/1e1427815001834825932758213/aigcVideoGenFile.mp4",
    "submit_time": 1793172926,
    "start_time": 1793172964,
    "finish_time": 1793173023,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q3",
      "origin_model_name": "vidu-video-q3"
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
            "ModelName": "Vidu",
            "ModelVersion": "q3",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 8,
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
            "Prompt": "森林小径，光斑随风移动",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-11-04T07:37:03Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/1e1427815001834825932758213/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 34132686,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8.042,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 34311883,
                  "VideoDuration": 8.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 34129264,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "SDR"
                      },
                      "Fps": 24,
                      "Height": 1080,
                      "Width": 1920
                    }
                  ],
                  "Width": 1920
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
              "OutputSeconds": 8,
              "ThoughtTokens": 0,
              "TotalSeconds": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-422116d22b67dfcaef319014c464ec56t"
        },
        "BeginProcessTime": "2026-10-28T07:35:26Z",
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
        "CreateTime": "2026-10-28T07:35:26Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:36:37Z",
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
        "RequestId": "d3990e40-e022-47e3-a43e-1adaf5435c95",
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
