# kling-video-3.0 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-09（seq=9）: 无声 — 720P × 5s
> 执行时间：2026/8/5 10:22:18  |  模型：`kling-video-3.0`  |  标签：standard

> 💡 silent-720P 82192×5=410960

### 调用参数
```json
{
  "model": "kling-video-3.0",
  "prompt": "雨后竹林，翠绿水珠悬挂叶尖，鸟鸣声声，空气清新"
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
| 预期 Quota | **205,480** |
| 预期 USD | $0.4110 |
| 预期 RMB | ¥3.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.14s |
| task_id | `task_CKJ7yygbdR79JFiT03eyIYstoH8xDPs9` |

```json
{
  "id": "task_CKJ7yygbdR79JFiT03eyIYstoH8xDPs9",
  "task_id": "task_CKJ7yygbdR79JFiT03eyIYstoH8xDPs9",
  "object": "video",
  "model": "kling-video-3.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785896356
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 16,079,489 | — |
| 提交后可用 Quota | 15,874,009 | — |
| **预扣金额** | **205,480** | **¥3.0000** |
| 预期扣减 | 205,480 | ¥3.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 13 |
| 完成耗时 | 180.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/2fa053415001834814625666506/aigcVideoGenFile.mp4

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
| 预扣金额 | 205,480 | ¥3.0000 |
| 平台记录最终消费 | 207,206 | **¥3.025208** |
| 差额 | — | +¥0.0252（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 110,
    "created_at": 1785896356,
    "updated_at": 1785896526,
    "task_id": "task_CKJ7yygbdR79JFiT03eyIYstoH8xDPs9",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.025208",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/2fa053415001834814625666506/aigcVideoGenFile.mp4",
    "submit_time": 1785896356,
    "start_time": 1785896365,
    "finish_time": 1785896526,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-3.0",
      "origin_model_name": "kling-video-3.0"
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
            "ModelName": "Kling",
            "ModelVersion": "3.0",
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
              "Resolution": "720P",
              "StorageMode": "Temporary"
            },
            "Prompt": "雨后竹林，翠绿水珠悬挂叶尖，鸟鸣声声，空气清新",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:22:06Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/2fa053415001834814625666506/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 4362119,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 2749226,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 4358092,
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
          "TaskId": "1500065731-AigcVideoTask-fe000c8c83dab9d175fea9719fca1258t"
        },
        "BeginProcessTime": "2026-08-05T02:19:17Z",
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
        "CreateTime": "2026-08-05T02:19:16Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:22:06Z",
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
        "RequestId": "2d916480-fc22-4b6e-8411-d83efffeff04",
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

## TC-10（seq=10）: 有声 — 1080P × 5s
> 执行时间：2026/8/5 10:24:05  |  模型：`kling-video-3.0`  |  标签：standard

> 💡 audio-1080P 164384×5=821920

### 调用参数
```json
{
  "model": "kling-video-3.0",
  "prompt": "雨后竹林，翠绿水珠悬挂叶尖，鸟鸣声声，空气清新",
  "duration": 5,
  "size": "1080P",
  "metadata": {
    "audio_generation": "Enabled"
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 5 | **5** | 已传，直接使用 |
| size（分辨率） | 1080P | — | 未传时适配器默认 720P |
| audio_generation | Enabled | — | 未传视为无声版 |
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
| 耗时 | 0.16s |
| task_id | `task_gNQvmvXmZwsLR0UTXD3rhiBcb23N6aIF` |

```json
{
  "id": "task_gNQvmvXmZwsLR0UTXD3rhiBcb23N6aIF",
  "task_id": "task_gNQvmvXmZwsLR0UTXD3rhiBcb23N6aIF",
  "object": "video",
  "model": "kling-video-3.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785896538
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 410,960 | ¥6.0000 |
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

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/d91992135001834814624233800/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.042s** |
| 计费参考时长 | 5s |
| 时长差值 | - |
| 输出分辨率 | 1920×1080 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 414,412 | **¥6.050415** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 111,
    "created_at": 1785896538,
    "updated_at": 1785896639,
    "task_id": "task_gNQvmvXmZwsLR0UTXD3rhiBcb23N6aIF",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.050415",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/d91992135001834814624233800/aigcVideoGenFile.mp4",
    "submit_time": 1785896538,
    "start_time": 1785896542,
    "finish_time": 1785896639,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-3.0",
      "origin_model_name": "kling-video-3.0"
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
            "ModelName": "Kling",
            "ModelVersion": "3.0",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "Enabled",
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
            "Prompt": "雨后竹林，翠绿水珠悬挂叶尖，鸟鸣声声，空气清新",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:23:59Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/d91992135001834814624233800/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 5.039,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 128349,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 44100
                    }
                  ],
                  "Bitrate": 14493824,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 9134733,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 14356316,
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
              "InputTokens": 0,
              "ThoughtTokens": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-3b43e3c8bf76f331c7836db23a141e40t"
        },
        "BeginProcessTime": "2026-08-05T02:22:18Z",
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
        "CreateTime": "2026-08-05T02:22:18Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:23:46Z",
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
        "RequestId": "7b331cf9-f343-4d91-9d79-415af7b0ff71",
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

