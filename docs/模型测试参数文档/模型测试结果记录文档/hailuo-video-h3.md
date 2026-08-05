# hailuo-video-h3 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-29（seq=29）: 默认（768P × 6s，占位价格）
> 执行时间：2026/8/5 10:56:38  |  模型：`hailuo-video-h3`  |  标签：standard

> 💡 ⚠️ H3 价格未公布，使用 02 同档占位。45205×6=271230

### 调用参数
```json
{
  "model": "hailuo-video-h3",
  "prompt": "深秋枫林，一片片红叶在晨风中轻轻飘落，光线穿透树冠，静谧"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | **6** | 未传，使用表达式默认值 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **135,615** |
| 预期 USD | $0.2712 |
| 预期 RMB | ¥1.9800 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.18s |
| task_id | `task_Sd66VTnF5pqPpWVz1CyixpaXNGha2Qst` |

```json
{
  "id": "task_Sd66VTnF5pqPpWVz1CyixpaXNGha2Qst",
  "task_id": "task_Sd66VTnF5pqPpWVz1CyixpaXNGha2Qst",
  "object": "video",
  "model": "hailuo-video-h3",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898417
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 135,615 | ¥1.9800 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 13 |
| 完成耗时 | 180.2s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/f98db8285001834814638275198/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **6.584s** |
| 计费参考时长 | 6s |
| 时长差值 | - |
| 输出分辨率 | 1344×768 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 148,814 | **¥2.172684** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 130,
    "created_at": 1785898417,
    "updated_at": 1785898592,
    "task_id": "task_Sd66VTnF5pqPpWVz1CyixpaXNGha2Qst",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "2.172684",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/f98db8285001834814638275198/aigcVideoGenFile.mp4",
    "submit_time": 1785898417,
    "start_time": 1785898431,
    "finish_time": 1785898592,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-h3",
      "origin_model_name": "hailuo-video-h3"
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
            "ModelName": "Hailuo",
            "ModelVersion": "H3",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 6,
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
              "Resolution": "768P",
              "StorageMode": "Temporary"
            },
            "Prompt": "深秋枫林，一片片红叶在晨风中轻轻飘落，光线穿透树冠，静谧",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:56:32Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/f98db8285001834814638275198/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 6.575,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 129611,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 32000
                    }
                  ],
                  "Bitrate": 4146516,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 6.584,
                  "Height": 768,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 3412583,
                  "VideoDuration": 6.583,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 4007270,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 768,
                      "Width": 1344
                    }
                  ],
                  "Width": 1344
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
          "TaskId": "1500065731-AigcVideoTask-a3c12b78dba6f483f9f9f2f0c79d1ac3t"
        },
        "BeginProcessTime": "2026-08-05T02:53:37Z",
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
        "CreateTime": "2026-08-05T02:53:37Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:56:18Z",
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
        "RequestId": "6b07add6-a3be-4c53-985a-792443f7f89f",
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

## TC-30（seq=30）: duration=10 + size=1080P（占位价格）
> 执行时间：2026/8/5 10:59:40  |  模型：`hailuo-video-h3`  |  标签：standard

> 💡 ⚠️ H3 价格未公布，使用 02 同档占位。79452×10=794520

### 调用参数
```json
{
  "model": "hailuo-video-h3",
  "prompt": "深秋枫林，一片片红叶在晨风中轻轻飘落，光线穿透树冠，静谧",
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
| 预期 Quota | **397,260** |
| 预期 USD | $0.7945 |
| 预期 RMB | ¥5.8000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.24s |
| task_id | `task_XQ8EKzDACEW1pXWLiDPaMpM8A11xAeUu` |

```json
{
  "id": "task_XQ8EKzDACEW1pXWLiDPaMpM8A11xAeUu",
  "task_id": "task_XQ8EKzDACEW1pXWLiDPaMpM8A11xAeUu",
  "object": "video",
  "model": "hailuo-video-h3",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898599
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 10,826,959 | — |
| 提交后可用 Quota | 10,429,699 | — |
| **预扣金额** | **397,260** | **¥5.8000** |
| 预期扣减 | 397,260 | ¥5.8000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 397,260 | — |
| request_id | 20260805025638993444000f61X4gQRqxU1Ekqp | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 13 |
| 完成耗时 | 180.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/dc08b5fe5001834814631129737/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **10.125s** |
| 计费参考时长 | 10s |
| 时长差值 | 0.125s |
| 输出分辨率 | 1344×768 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 397,260 | ¥5.8000 |
| 平台记录最终消费 | 402,225 | **¥5.872485** |
| 差额 | — | +¥0.0725（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 131,
    "created_at": 1785898599,
    "updated_at": 1785898770,
    "task_id": "task_XQ8EKzDACEW1pXWLiDPaMpM8A11xAeUu",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "5.872485",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/dc08b5fe5001834814631129737/aigcVideoGenFile.mp4",
    "submit_time": 1785898599,
    "start_time": 1785898609,
    "finish_time": 1785898770,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-h3",
      "origin_model_name": "hailuo-video-h3"
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
            "ModelName": "Hailuo",
            "ModelVersion": "H3",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "",
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
            "Prompt": "深秋枫林，一片片红叶在晨风中轻轻飘落，光线穿透树冠，静谧",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:59:30Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/dc08b5fe5001834814631129737/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 10.125,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 128574,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 32000
                    }
                  ],
                  "Bitrate": 4429300,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 10.125,
                  "Height": 768,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 5605834,
                  "VideoDuration": 10.125,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 4291902,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 768,
                      "Width": 1344
                    }
                  ],
                  "Width": 1344
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
          "TaskId": "1500065731-AigcVideoTask-9b9d10719a4759563445abd312568dd0t"
        },
        "BeginProcessTime": "2026-08-05T02:56:39Z",
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
        "CreateTime": "2026-08-05T02:56:39Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:59:28Z",
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
        "RequestId": "881bc996-ae52-4465-9d1e-4b32b5062bf8",
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

