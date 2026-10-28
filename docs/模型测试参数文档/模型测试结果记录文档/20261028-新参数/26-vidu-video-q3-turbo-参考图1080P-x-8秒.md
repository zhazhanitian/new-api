# vidu-video-q3-turbo 测试报告

## 26. vidu-video-q3-turbo 参考图1080P x 8秒
> 执行时间：2026-10-28  |  模型：`vidu-video-q3-turbo`

> 表达式值 444928，quota 222464

### 调用参数
```json
{
  "model": "vidu-video-q3-turbo",
  "prompt": "保持参考图风格，镜头缓慢拉远",
  "duration": 8,
  "resolution": "1080P",
  "metadata": {
    "content": [
      {
        "type": "image_url",
        "role": "reference_image",
        "image_url": {
          "url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906?authorization=bce-auth-v1%2FALTAKMD70FiEhTeFimgJH1yjAb%2F2026-10-28T07%3A25%3A39Z%2F600%2F%2F83ecd58fe0f9818278fae698dd920bdc4b1ee0fe7012fc23f42fa5df0a7bb5e7"
        }
      }
    ]
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| duration | 8 | 表达式按请求秒数预扣，缺省走模型默认 |
| resolution | 1080P | 外层分辨率 |
| ratio | 未传 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | 未传 | 豆包 1.5 显式 false 才无声 |
| content | image_url/reference_image | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **222,464** |
| 预期 RMB | ¥3.2480 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_7MZk3lpxny1vHpL2nnQI67P8lDMFFBwh` |

```json
{
  "id": "task_7MZk3lpxny1vHpL2nnQI67P8lDMFFBwh",
  "task_id": "task_7MZk3lpxny1vHpL2nnQI67P8lDMFFBwh",
  "object": "video",
  "model": "vidu-video-q3-turbo",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172928
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 222,464 | ¥3.2480 |
| 日志预扣 | - | - |
| 平台记录最终消费 | 223,631 | ¥3.2650 |
| 是否符合 | ✅ 符合（实扣含成片秒数微调） | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 10 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/9b6a60705001834825866477575/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 8.042 |
| 输出分辨率 | 1920×1080 |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_7MZk3lpxny1vHpL2nnQI67P8lDMFFBwh`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 250,
    "created_at": 1793172928,
    "updated_at": 1793173126,
    "task_id": "task_7MZk3lpxny1vHpL2nnQI67P8lDMFFBwh",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.265013",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/9b6a60705001834825866477575/aigcVideoGenFile.mp4",
    "submit_time": 1793172928,
    "start_time": 1793172968,
    "finish_time": 1793173126,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q3-turbo",
      "origin_model_name": "vidu-video-q3-turbo"
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
                "Url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906?authorization=bce-auth-v1%2FALTAKMD70FiEhTeFimgJH1yjAb%2F2026-10-28T07%3A25%3A39Z%2F600%2F%2F83ecd58fe0f9818278fae698dd920bdc4b1ee0fe7012fc23f42fa5df0a7bb5e7",
                "Usage": "Reference",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Vidu",
            "ModelVersion": "q3-turbo",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "",
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
            "Prompt": "保持参考图风格，镜头缓慢拉远",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-11-04T07:38:46Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/9b6a60705001834825866477575/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 18154288,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8.042,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 18249599,
                  "VideoDuration": 8.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 18150273,
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
          "TaskId": "1500065731-AigcVideoTask-d671009793c4b93d868392cdd41c7499t"
        },
        "BeginProcessTime": "2026-10-28T07:35:28Z",
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
        "CreateTime": "2026-10-28T07:35:28Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:38:34Z",
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
        "RequestId": "ec817b28-bd1d-4fa0-8f92-7593f9f6496b",
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
