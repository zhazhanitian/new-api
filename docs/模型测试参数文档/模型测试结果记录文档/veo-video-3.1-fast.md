# veo-video-3.1-fast 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-33（seq=33）: 无声（默认）— d=8 固定
> 执行时间：2026/8/5 11:04:16  |  模型：`veo-video-3.1-fast`  |  标签：standard

> 💡 silent-720P 102740×8=821920

### 调用参数
```json
{
  "model": "veo-video-3.1-fast",
  "prompt": "热带海底，珊瑚礁旁鱼群穿梭，阳光折射，蓝色清澈"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | **8** | 未传，使用表达式默认值 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **410,960** |
| 预期 USD | $0.8219 |
| 预期 RMB | ¥6.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.14s |
| task_id | `task_UxHS9LalO60gMWoqTp6NhEkpLDeDkk2T` |

```json
{
  "id": "task_UxHS9LalO60gMWoqTp6NhEkpLDeDkk2T",
  "task_id": "task_UxHS9LalO60gMWoqTp6NhEkpLDeDkk2T",
  "object": "video",
  "model": "veo-video-3.1-fast",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898964
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 7,945,783 | — |
| 提交后可用 Quota | 7,534,823 | — |
| **预扣金额** | **410,960** | **¥6.0000** |
| 预期扣减 | 410,960 | ¥6.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 410,960 | — |
| request_id | 20260805030244416401000f61X4gQRdMGAwvNq | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 7 |
| 完成耗时 | 90.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/1e5c89a25001834814635788114/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8s** |
| 计费参考时长 | 8s |
| 时长差值 | 0.000s |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 410,960 | ¥6.0000 |
| 平台记录最终消费 | 410,960 | **¥6.000016** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 134,
    "created_at": 1785898964,
    "updated_at": 1785899044,
    "task_id": "task_UxHS9LalO60gMWoqTp6NhEkpLDeDkk2T",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.000016",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/1e5c89a25001834814635788114/aigcVideoGenFile.mp4",
    "submit_time": 1785898964,
    "start_time": 1785898980,
    "finish_time": 1785899044,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "veo-video-3.1-fast",
      "origin_model_name": "veo-video-3.1-fast"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ErrCode": 0,
          "ErrCodeExt": "",
          "Input": {
            "EnhancePrompt": "",
            "FileInfos": [],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "GV",
            "ModelVersion": "3.1-fast",
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
              "Resolution": "720P",
              "StorageMode": "Temporary"
            },
            "Prompt": "热带海底，珊瑚礁旁鱼群穿梭，阳光折射，蓝色清澈",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:04:04Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/1e5c89a25001834814635788114/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 17445340,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 17445340,
                  "VideoDuration": 8,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 17433069,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
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
              "InputTokens": 0,
              "ThoughtTokens": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-8e7950b8ed4369288005a955925eac69t"
        },
        "BeginProcessTime": "2026-08-05T03:02:44Z",
        "ClipTask": null,
        "ComplexAdaptiveDynamicStreamingTask": null,
        "ComposeMediaTask": null,
        "ConcatTask": null,
        "CreateAigcAdvancedCustomElementTask": null,
        "CreateAigcAudioCloneTask": null,
        "CreateAigcCustomVoiceTask": null,
        "CreateAigcMaterialTask": null,
        "CreateAigcSubjectTask": null,
        "CreateImageSpriteTask": null,
        "CreateTime": "2026-08-05T03:02:44Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:03:59Z",
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
        "RequestId": "6bef8652-3bb0-45d4-8d87-344d89e7e9fb",
        "ReviewAudioVideoTask": null,
        "SceneAigcImageTask": null,
        "SceneAigcVideoTask": null,
        "SnapshotByTimeOffsetTask": null,
        "SplitMediaTask": null,
        "Status": "FINISH",
        "TaskType": "AigcVideoTask",
        "TranscodeTask": null,
        "WechatMiniProgramPublishTask": null,
        "WechatPublishTask": null
      }
    }
  }
}
```

---

## TC-34（seq=34）: 有声 — audio_generation=Enabled
> 执行时间：2026/8/5 11:06:02  |  模型：`veo-video-3.1-fast`  |  标签：standard

> 💡 audio-720P 154110×8=1232880

### 调用参数
```json
{
  "model": "veo-video-3.1-fast",
  "prompt": "热带海底，珊瑚礁旁鱼群穿梭，阳光折射，蓝色清澈",
  "metadata": {
    "audio_generation": "Enabled"
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | **8** | 未传，使用表达式默认值 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | Enabled | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **616,440** |
| 预期 USD | $1.2329 |
| 预期 RMB | ¥9.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_Kxhdpr8sfaGJfgyjcECA6TTMK1gNEoLj` |

```json
{
  "id": "task_Kxhdpr8sfaGJfgyjcECA6TTMK1gNEoLj",
  "task_id": "task_Kxhdpr8sfaGJfgyjcECA6TTMK1gNEoLj",
  "object": "video",
  "model": "veo-video-3.1-fast",
  "status": "queued",
  "progress": 0,
  "created_at": 1785899056
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 7,534,823 | — |
| 提交后可用 Quota | 6,918,383 | — |
| **预扣金额** | **616,440** | **¥9.0000** |
| 预期扣减 | 616,440 | ¥9.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 8 |
| 完成耗时 | 105.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/81a3b5205001834814633623873/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8s** |
| 计费参考时长 | 8s |
| 时长差值 | 0.000s |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 616,440 | ¥9.0000 |
| 平台记录最终消费 | 616,440 | **¥9.000024** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 135,
    "created_at": 1785899056,
    "updated_at": 1785899157,
    "task_id": "task_Kxhdpr8sfaGJfgyjcECA6TTMK1gNEoLj",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "9.000024",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/81a3b5205001834814633623873/aigcVideoGenFile.mp4",
    "submit_time": 1785899056,
    "start_time": 1785899060,
    "finish_time": 1785899157,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "veo-video-3.1-fast",
      "origin_model_name": "veo-video-3.1-fast"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ErrCode": 0,
          "ErrCodeExt": "",
          "Input": {
            "EnhancePrompt": "",
            "FileInfos": [],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "GV",
            "ModelVersion": "3.1-fast",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "Enabled",
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
              "Resolution": "720P",
              "StorageMode": "Temporary"
            },
            "Prompt": "热带海底，珊瑚礁旁鱼群穿梭，阳光折射，蓝色清澈",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:05:57Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/81a3b5205001834814633623873/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 8,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 256297,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 48000
                    }
                  ],
                  "Bitrate": 22226267,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 22226267,
                  "VideoDuration": 8,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 21952949,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
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
              "InputTokens": 0,
              "ThoughtTokens": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-535c3f76144c38a0611f76abdaf4d0b0t"
        },
        "BeginProcessTime": "2026-08-05T03:04:16Z",
        "ClipTask": null,
        "ComplexAdaptiveDynamicStreamingTask": null,
        "ComposeMediaTask": null,
        "ConcatTask": null,
        "CreateAigcAdvancedCustomElementTask": null,
        "CreateAigcAudioCloneTask": null,
        "CreateAigcCustomVoiceTask": null,
        "CreateAigcMaterialTask": null,
        "CreateAigcSubjectTask": null,
        "CreateImageSpriteTask": null,
        "CreateTime": "2026-08-05T03:04:16Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:05:43Z",
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
        "RequestId": "d27d9b6d-5b56-420c-9d89-ddc633f777cb",
        "ReviewAudioVideoTask": null,
        "SceneAigcImageTask": null,
        "SceneAigcVideoTask": null,
        "SnapshotByTimeOffsetTask": null,
        "SplitMediaTask": null,
        "Status": "FINISH",
        "TaskType": "AigcVideoTask",
        "TranscodeTask": null,
        "WechatMiniProgramPublishTask": null,
        "WechatPublishTask": null
      }
    }
  }
}
```

---

