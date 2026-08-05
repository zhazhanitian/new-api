# mingmou-video-1.0 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-40（seq=40）: 720P × 5s
> 执行时间：2026/8/5 11:18:43  |  模型：`mingmou-video-1.0`  |  标签：standard

> 💡 720P 41096×5=205480（同 Hunyuan 定价）

### 调用参数
```json
{
  "model": "mingmou-video-1.0",
  "prompt": "微缩城市模型，汽车穿行在精致街道，移轴摄影效果"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | N/A（不参与计费） | 模型不发送 Duration |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **102,740** |
| 预期 USD | $0.2055 |
| 预期 RMB | ¥1.5000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.13s |
| task_id | `task_z97AntChyJyHoK1UJGXcrGLF22znt1SH` |

```json
{
  "id": "task_z97AntChyJyHoK1UJGXcrGLF22znt1SH",
  "task_id": "task_z97AntChyJyHoK1UJGXcrGLF22znt1SH",
  "object": "video",
  "model": "mingmou-video-1.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785899831
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 5,230,025 | — |
| 提交后可用 Quota | 5,127,285 | — |
| **预扣金额** | **102,740** | **¥1.5000** |
| 预期扣减 | 102,740 | ¥1.5000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 102,740 | — |
| request_id | 20260805031711806598000f61X4gQRw5cHxAeM | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 7 |
| 完成耗时 | 90.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/15c35ea15001834814642199290/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.875s** |
| 计费参考时长 | N/A |
| 时长差值 | - |
| 输出分辨率 | 1366×768 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 102,740 | ¥1.5000 |
| 平台记录最终消费 | 102,740 | **¥1.500004** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 141,
    "created_at": 1785899831,
    "updated_at": 1785899915,
    "task_id": "task_z97AntChyJyHoK1UJGXcrGLF22znt1SH",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.500004",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/15c35ea15001834814642199290/aigcVideoGenFile.mp4",
    "submit_time": 1785899831,
    "start_time": 1785899835,
    "finish_time": 1785899915,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "mingmou-video-1.0",
      "origin_model_name": "mingmou-video-1.0"
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
            "ModelName": "Mingmou",
            "ModelVersion": "1.0",
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
              "Resolution": "",
              "StorageMode": "Temporary"
            },
            "Prompt": "微缩城市模型，汽车穿行在精致街道，移轴摄影效果",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:18:36Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/15c35ea15001834814642199290/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 922225,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.875,
                  "Height": 768,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 677259,
                  "VideoDuration": 5.875,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 916922,
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
              "InputTokens": 0,
              "ThoughtTokens": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-317218231bc8a0fcb80ea38a2c947565t"
        },
        "BeginProcessTime": "2026-08-05T03:17:12Z",
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
        "CreateTime": "2026-08-05T03:17:12Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:18:23Z",
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
        "RequestId": "c6c2259e-5a34-4947-973e-28e5572f6acf",
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

## TC-41（seq=41）: 1080P × 10s
> 执行时间：2026/8/5 11:20:00  |  模型：`mingmou-video-1.0`  |  标签：standard

> 💡 1080P 68493×10=684930

### 调用参数
```json
{
  "model": "mingmou-video-1.0",
  "prompt": "微缩城市模型，汽车穿行在精致街道，移轴摄影效果",
  "duration": 10,
  "size": "1080P"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 10 | N/A（不参与计费） | 模型不发送 Duration |
| size（分辨率） | 1080P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **342,465** |
| 预期 USD | $0.6849 |
| 预期 RMB | ¥5.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.16s |
| task_id | `task_yeg8OJcWt314ndnt7r3gP5p7xuGibuaj` |

```json
{
  "id": "task_yeg8OJcWt314ndnt7r3gP5p7xuGibuaj",
  "task_id": "task_yeg8OJcWt314ndnt7r3gP5p7xuGibuaj",
  "object": "video",
  "model": "mingmou-video-1.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785899923
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 5,127,285 | — |
| 提交后可用 Quota | 4,784,820 | — |
| **预扣金额** | **342,465** | **¥5.0000** |
| 预期扣减 | 342,465 | ¥5.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 342,465 | — |
| request_id | 20260805031843546684000f61X4gQRGmfeEcet | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 6 |
| 完成耗时 | 75.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f98c34815001834814638272950/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.875s** |
| 计费参考时长 | N/A |
| 时长差值 | - |
| 输出分辨率 | 1366×768 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 342,465 | ¥5.0000 |
| 平台记录最终消费 | 342,465 | **¥4.999989** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 142,
    "created_at": 1785899923,
    "updated_at": 1785899996,
    "task_id": "task_yeg8OJcWt314ndnt7r3gP5p7xuGibuaj",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "4.999989",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f98c34815001834814638272950/aigcVideoGenFile.mp4",
    "submit_time": 1785899923,
    "start_time": 1785899932,
    "finish_time": 1785899996,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "mingmou-video-1.0",
      "origin_model_name": "mingmou-video-1.0"
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
            "ModelName": "Mingmou",
            "ModelVersion": "1.0",
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
              "Resolution": "",
              "StorageMode": "Temporary"
            },
            "Prompt": "微缩城市模型，汽车穿行在精致街道，移轴摄影效果",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:19:56Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f98c34815001834814638272950/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 1424501,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.875,
                  "Height": 768,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 1046118,
                  "VideoDuration": 5.875,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 1419195,
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
              "InputTokens": 0,
              "ThoughtTokens": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-5b46bd45a39ba571a86c07194e7e2664t"
        },
        "BeginProcessTime": "2026-08-05T03:18:43Z",
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
        "CreateTime": "2026-08-05T03:18:43Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:19:46Z",
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
        "RequestId": "105db439-dc87-4e52-a6f5-bb312cc048ae",
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

