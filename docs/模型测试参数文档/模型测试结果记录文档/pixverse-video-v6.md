# pixverse-video-v6 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-44（seq=44）: 无声 — 720p × 5s
> 执行时间：2026/8/5 11:22:35  |  模型：`pixverse-video-v6`  |  标签：standard

> 💡 silent-720p 36164×5=180820

### 调用参数
```json
{
  "model": "pixverse-video-v6",
  "prompt": "冰川融化，蔚蓝海水与白色浮冰交织，鸟瞰航拍"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | **5** | 未传，使用表达式默认值 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **90,410** |
| 预期 USD | $0.1808 |
| 预期 RMB | ¥1.3200 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.14s |
| task_id | `task_MiQ0moEYvZwWMmcRvFZv7Teoq9kutHhK` |

```json
{
  "id": "task_MiQ0moEYvZwWMmcRvFZv7Teoq9kutHhK",
  "task_id": "task_MiQ0moEYvZwWMmcRvFZv7Teoq9kutHhK",
  "object": "video",
  "model": "pixverse-video-v6",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900078
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 4,340,621 | — |
| 提交后可用 Quota | 4,270,416 | — |
| **预扣金额** | **70,205** | **¥1.0250** |
| 预期扣减 | 90,410 | ¥1.3200 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | 70,205 | — |
| request_id | 20260805032118704415000f61X4gQR58g0Nmi0 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 6 |
| 完成耗时 | 75.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/0147a0835001834814645384671/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.042s** |
| 计费参考时长 | 5s |
| 时长差值 | 0.042s |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 70,205 | ¥1.0250 |
| 平台记录最终消费 | 70,794 | **¥1.033592** |
| 差额 | — | +¥0.0086（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 145,
    "created_at": 1785900078,
    "updated_at": 1785900142,
    "task_id": "task_MiQ0moEYvZwWMmcRvFZv7Teoq9kutHhK",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.033592",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/0147a0835001834814645384671/aigcVideoGenFile.mp4",
    "submit_time": 1785900078,
    "start_time": 1785900093,
    "finish_time": 1785900142,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "pixverse-video-v6",
      "origin_model_name": "pixverse-video-v6"
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
            "ModelName": "PixVerse",
            "ModelVersion": "v6",
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
              "Resolution": "720p",
              "StorageMode": "Temporary"
            },
            "Prompt": "冰川融化，蔚蓝海水与白色浮冰交织，鸟瞰航拍",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:22:22Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/0147a0835001834814645384671/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 3651358,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 2301269,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 3645913,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "SDR"
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
          "TaskId": "1500065731-AigcVideoTask-16188c195b9ab5db834c304174ab0135t"
        },
        "BeginProcessTime": "2026-08-05T03:21:18Z",
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
        "CreateTime": "2026-08-05T03:21:18Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:22:19Z",
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
        "RequestId": "a0e5423f-f62f-42c2-a715-8f9b3ca32e59",
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

## TC-45（seq=45）: 有声 — 1080p × 8s（audio_generation=Enabled）
> 执行时间：2026/8/5 11:24:22  |  模型：`pixverse-video-v6`  |  标签：standard

> 💡 audio-1080p 92466×8=739728

### 调用参数
```json
{
  "model": "pixverse-video-v6",
  "prompt": "冰川融化，蔚蓝海水与白色浮冰交织，鸟瞰航拍",
  "duration": 8,
  "size": "1080p",
  "metadata": {
    "audio_generation": "Enabled"
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 8 | **8** | 已传，直接使用 |
| size（分辨率） | 1080p | — | 未传时适配器默认 720P |
| audio_generation | Enabled | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **369,864** |
| 预期 USD | $0.7397 |
| 预期 RMB | ¥5.4000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.17s |
| task_id | `task_R0bAJJIMp5s0FBkelJ7jwFGTTMj0SMvX` |

```json
{
  "id": "task_R0bAJJIMp5s0FBkelJ7jwFGTTMj0SMvX",
  "task_id": "task_R0bAJJIMp5s0FBkelJ7jwFGTTMj0SMvX",
  "object": "video",
  "model": "pixverse-video-v6",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900155
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 369,864 | ¥5.4000 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 8 |
| 完成耗时 | 105.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/705feaa25001834814643745419/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8.042s** |
| 计费参考时长 | 8s |
| 时长差值 | - |
| 输出分辨率 | 1920×1080 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 371,805 | **¥5.428353** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 146,
    "created_at": 1785900155,
    "updated_at": 1785900254,
    "task_id": "task_R0bAJJIMp5s0FBkelJ7jwFGTTMj0SMvX",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "5.428353",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/705feaa25001834814643745419/aigcVideoGenFile.mp4",
    "submit_time": 1785900155,
    "start_time": 1785900158,
    "finish_time": 1785900254,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "pixverse-video-v6",
      "origin_model_name": "pixverse-video-v6"
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
            "ModelName": "PixVerse",
            "ModelVersion": "v6",
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
              "Resolution": "1080p",
              "StorageMode": "Temporary"
            },
            "Prompt": "冰川融化，蔚蓝海水与白色浮冰交织，鸟瞰航拍",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:24:15Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/705feaa25001834814643745419/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 7111267,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8.042,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 7148602,
                  "VideoDuration": 8.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 7107101,
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
              "InputTokens": 0,
              "ThoughtTokens": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-1f1e20652d009ba3210825096b66f82bt"
        },
        "BeginProcessTime": "2026-08-05T03:22:35Z",
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
        "CreateTime": "2026-08-05T03:22:35Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:24:11Z",
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
        "RequestId": "ce76c18f-b4e1-4b8d-9e62-cdbf17768646",
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

