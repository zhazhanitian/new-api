# kling-video-1.6 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-01（seq=1）: 不传 duration/size（默认 720P × 5s）
> 执行时间：2026/8/5 09:51:15  |  模型：`kling-video-1.6`  |  标签：standard

> 💡 适配器默认 720P；API 默认时长 5s。表达式 54795×5=273975

### 调用参数
```json
{
  "model": "kling-video-1.6",
  "prompt": "夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格"
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
| 预期 Quota | **136,988** |
| 预期 USD | $0.2740 |
| 预期 RMB | ¥2.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.17s |
| task_id | `task_zP9XeNemCOLcAbmnMOckV4VtlaRr7aL1` |

```json
{
  "id": "task_zP9XeNemCOLcAbmnMOckV4VtlaRr7aL1",
  "task_id": "task_zP9XeNemCOLcAbmnMOckV4VtlaRr7aL1",
  "object": "video",
  "model": "kling-video-1.6",
  "status": "queued",
  "progress": 0,
  "created_at": 1785894388
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 18,624,666 | — |
| 提交后可用 Quota | 18,487,678 | — |
| **预扣金额** | **136,988** | **¥2.0000** |
| 预期扣减 | 136,988 | ¥2.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 136,988 | — |
| request_id | 20260805014628761334000f61X4gQRxobFOTEs | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 20 |
| 完成耗时 | 285.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/a9c72e3e5001834814715051749/aigcVideoGenFile.mp4

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
| 预扣金额 | 136,988 | ¥2.0000 |
| 平台记录最终消费 | 138,138 | **¥2.016815** |
| 差额 | — | +¥0.0168（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 102,
    "created_at": 1785894388,
    "updated_at": 1785894673,
    "task_id": "task_zP9XeNemCOLcAbmnMOckV4VtlaRr7aL1",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "2.016815",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/a9c72e3e5001834814715051749/aigcVideoGenFile.mp4",
    "submit_time": 1785894388,
    "start_time": 1785894398,
    "finish_time": 1785894673,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-1.6",
      "origin_model_name": "kling-video-1.6"
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
            "ModelVersion": "1.6",
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
            "Prompt": "夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T01:51:13Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/a9c72e3e5001834814715051749/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 7582035,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 4778578,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 7576810,
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
          "TaskId": "1500065731-AigcVideoTask-68a21dd23ac5253ef670ea2736e07700t"
        },
        "BeginProcessTime": "2026-08-05T01:46:28Z",
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
        "CreateTime": "2026-08-05T01:46:28Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T01:51:01Z",
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
        "RequestId": "b49f09e7-1f85-434b-bf05-a5a50999cc8a",
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

## TC-02（seq=2）: duration=10 + size=1080P
> 执行时间：2026/8/5 10:01:17  |  模型：`kling-video-1.6`  |  标签：standard

> 💡 95890×10=958900

### 调用参数
```json
{
  "model": "kling-video-1.6",
  "prompt": "夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格",
  "duration": 10,
  "size": "1080P"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 10 | **10** | 已传，直接使用 |
| size（分辨率） | 1080P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **479,450** |
| 预期 USD | $0.9589 |
| 预期 RMB | ¥7.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_J2SnugOgiiDvpnSAWhZHzgjKylmW8Pmj` |

```json
{
  "id": "task_J2SnugOgiiDvpnSAWhZHzgjKylmW8Pmj",
  "task_id": "task_J2SnugOgiiDvpnSAWhZHzgjKylmW8Pmj",
  "object": "video",
  "model": "kling-video-1.6",
  "status": "queued",
  "progress": 0,
  "created_at": 1785894675
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 18,487,678 | — |
| 提交后可用 Quota | 18,008,228 | — |
| **预扣金额** | **479,450** | **¥7.0000** |
| 预期扣减 | 479,450 | ¥7.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 479,450 | — |
| request_id | 20260805015115610516000f61X4gQRjfxLetbC | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **timeout** |
| 轮询次数 | 40 |
| 完成耗时 | 600.4s |
| progress | 0% |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 479,450 | ¥7.0000 |
| 平台记录最终消费 | 479,450 | **¥6.999970** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 103,
    "created_at": 1785894675,
    "updated_at": 1785895253,
    "task_id": "task_J2SnugOgiiDvpnSAWhZHzgjKylmW8Pmj",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.999970",
    "action": "generate",
    "status": "IN_PROGRESS",
    "fail_reason": "",
    "submit_time": 1785894675,
    "start_time": 1785894689,
    "finish_time": 0,
    "progress": "0%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-1.6",
      "origin_model_name": "kling-video-1.6"
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
            "ModelVersion": "1.6",
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
            "Prompt": "夏日海滩，海浪轻拍礁石，阳光折射出金色光芒，写实风格",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [],
            "ProcedureTaskIds": [],
            "Usage": {
              "InputTokens": 0,
              "ThoughtTokens": 0
            }
          },
          "Progress": 0,
          "SessionContext": "",
          "SessionId": "",
          "Status": "PROCESSING",
          "TaskId": "1500065731-AigcVideoTask-fc7f9e060b29a827a6abfb2477547e62t"
        },
        "BeginProcessTime": "2026-08-05T01:51:15Z",
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
        "CreateTime": "2026-08-05T01:51:15Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "0001-01-01T00:00:00Z",
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
        "RequestId": "0434dd86-9b7a-4692-94e4-8631efc32be4",
        "ReviewAudioVideoTask": null,
        "SceneAigcImageTask": null,
        "SceneAigcVideoTask": null,
        "SnapshotByTimeOffsetTask": null,
        "SplitMediaTask": null,
        "Status": "PROCESSING",
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

