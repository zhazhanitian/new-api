# hailuo-video-h3 测试报告

## 33. hailuo-video-h3 1080P x 10秒
> 执行时间：2026-10-28  |  模型：`hailuo-video-h3`

> 表达式值 890410，quota 445205

### 调用参数
```json
{
  "model": "hailuo-video-h3",
  "prompt": "海边栈道，镜头跟着人物向前走",
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
| 预期 Quota | **445,205** |
| 预期 RMB | ¥6.5000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_vDje6uOFYx7oKPgeDnvVuXRIGNMRYG6l` |

```json
{
  "id": "task_vDje6uOFYx7oKPgeDnvVuXRIGNMRYG6l",
  "task_id": "task_vDje6uOFYx7oKPgeDnvVuXRIGNMRYG6l",
  "object": "video",
  "model": "hailuo-video-h3",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172932
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 445,205 | ¥6.5000 |
| 日志预扣 | - | - |
| 平台记录最终消费 | 451,660 | ¥6.5942 |
| 是否符合 | ✅ 符合（实扣含成片秒数微调） | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 14 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/e372764a5001834825867238573/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 10.145 |
| 输出分辨率 | 1890×1080 |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_vDje6uOFYx7oKPgeDnvVuXRIGNMRYG6l`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 257,
    "created_at": 1793172932,
    "updated_at": 1793173186,
    "task_id": "task_vDje6uOFYx7oKPgeDnvVuXRIGNMRYG6l",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.594236",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/e372764a5001834825867238573/aigcVideoGenFile.mp4",
    "submit_time": 1793172932,
    "start_time": 1793172976,
    "finish_time": 1793173186,
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
            "FileInfos": [],
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
                "ExpireTime": "2026-11-04T07:39:46Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/e372764a5001834825867238573/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 10.144,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 97077,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 44100
                    }
                  ],
                  "Bitrate": 8075722,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 10.145,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 10241025,
                  "VideoDuration": 10.125,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 7986082,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 1080,
                      "Width": 1890
                    }
                  ],
                  "Width": 1890
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
              "OutputSeconds": 10,
              "ThoughtTokens": 0,
              "TotalSeconds": 10
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-35b48442fb707d5530387c47d315a30et"
        },
        "BeginProcessTime": "2026-10-28T07:35:32Z",
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
        "CreateTime": "2026-10-28T07:35:32Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:39:40Z",
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
        "RequestId": "57c028ed-05dc-4cc8-9701-cdfbbdecf073",
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
