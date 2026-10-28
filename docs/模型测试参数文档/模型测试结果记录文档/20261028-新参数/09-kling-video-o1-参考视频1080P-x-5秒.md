# kling-video-o1 测试报告

## 9. kling-video-o1 参考视频1080P x 5秒
> 执行时间：2026-10-28  |  模型：`kling-video-o1`

> 表达式值 821920，quota 410960

### 调用参数
```json
{
  "model": "kling-video-o1",
  "prompt": "沿用参考视频的运镜，色调偏暖",
  "duration": 5,
  "resolution": "1080P",
  "metadata": {
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
| audio_generation | 未传 | 仅 Enabled 走有声价 |
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
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_ALrupMjYA0uqqT93auZHmlOvRTcfItOV` |

```json
{
  "id": "task_ALrupMjYA0uqqT93auZHmlOvRTcfItOV",
  "task_id": "task_ALrupMjYA0uqqT93auZHmlOvRTcfItOV",
  "object": "video",
  "model": "kling-video-o1",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172817
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 410,960 | ¥6.0000 |
| 日志预扣 | - | - |
| 平台记录最终消费 | 414,412 | ¥6.0504 |
| 是否符合 | ✅ 符合（实扣含成片秒数微调） | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 1 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/4e6e71265001834825865474227/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 5.042 |
| 输出分辨率 | 1920×1080 |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_ALrupMjYA0uqqT93auZHmlOvRTcfItOV`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 231,
    "created_at": 1793172817,
    "updated_at": 1793173072,
    "task_id": "task_ALrupMjYA0uqqT93auZHmlOvRTcfItOV",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.050415",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/4e6e71265001834825865474227/aigcVideoGenFile.mp4",
    "submit_time": 1793172817,
    "start_time": 1793172835,
    "finish_time": 1793173072,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-o1",
      "origin_model_name": "kling-video-o1"
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
                "Category": "Video",
                "FileId": "",
                "KeepOriginalSound": "",
                "ObjectId": "",
                "ReferenceType": "",
                "Text": "",
                "Type": "Url",
                "Url": "https://aigc-bucket.su.bcebos.com/v1/objects/0d/0dca41d01e3579e43d30ae64fd30126aa056963fe85c5f254b81a3868a37298d?authorization=bce-auth-v1%2FALTAKMD70FiEhTeFimgJH1yjAb%2F2026-10-28T07%3A26%3A04Z%2F600%2F%2F03d1d6955391e462eadcb73c5ef846a63b2ba1543d2d63bdbae041429f5767f2",
                "Usage": "Reference",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Kling",
            "ModelVersion": "O1",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "",
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
            "Prompt": "沿用参考视频的运镜，色调偏暖",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-11-04T07:37:52Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/4e6e71265001834825865474227/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 5.039,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 128483,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 44100
                    }
                  ],
                  "Bitrate": 10723990,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 6758795,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 10586097,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
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
              "OutputSeconds": 5,
              "ThoughtTokens": 0,
              "TotalSeconds": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-86076da34d89b6dde0fcab91d12d34f5t"
        },
        "BeginProcessTime": "2026-10-28T07:33:38Z",
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
        "CreateTime": "2026-10-28T07:33:38Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:37:37Z",
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
        "RequestId": "3592e4a0-c232-4f7f-93af-bbb402faa29e",
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
