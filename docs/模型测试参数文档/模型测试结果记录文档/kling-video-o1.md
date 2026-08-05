# kling-video-o1 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-07（seq=7）: 无参考图 — 720P × 5s
> 执行时间：2026/8/5 10:16:44  |  模型：`kling-video-o1`  |  标签：standard

> 💡 t2v-720P 82192×5=410960

### 调用参数
```json
{
  "model": "kling-video-o1",
  "prompt": "水墨山水画风格，群山叠翠，云雾飘渺，鸟儿掠过"
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
| task_id | `task_6q5iEckzdcvZJHecZ1OR1ORmeC8sYrzG` |

```json
{
  "id": "task_6q5iEckzdcvZJHecZ1OR1ORmeC8sYrzG",
  "task_id": "task_6q5iEckzdcvZJHecZ1OR1ORmeC8sYrzG",
  "object": "video",
  "model": "kling-video-o1",
  "status": "queued",
  "progress": 0,
  "created_at": 1785896143
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 17,111,491 | — |
| 提交后可用 Quota | 16,906,011 | — |
| **预扣金额** | **205,480** | **¥3.0000** |
| 预期扣减 | 205,480 | ¥3.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 205,480 | — |
| request_id | 20260805021543240562000f61X4gQRZLpiu5x0 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 5 |
| 完成耗时 | 60.0s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f8f37c005001834814716133794/aigcVideoGenFile.mp4

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
    "id": 108,
    "created_at": 1785896143,
    "updated_at": 1785896203,
    "task_id": "task_6q5iEckzdcvZJHecZ1OR1ORmeC8sYrzG",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.025208",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f8f37c005001834814716133794/aigcVideoGenFile.mp4",
    "submit_time": 1785896143,
    "start_time": 1785896155,
    "finish_time": 1785896203,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-o1",
      "origin_model_name": "kling-video-o1"
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
            "ModelVersion": "O1",
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
            "Prompt": "水墨山水画风格，群山叠翠，云雾飘渺，鸟儿掠过",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:16:43Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f8f37c005001834814716133794/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 8926559,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 5625964,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 8922817,
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
          "TaskId": "1500065731-AigcVideoTask-03eb3f9cd41194942b94907f4f673f78t"
        },
        "BeginProcessTime": "2026-08-05T02:15:43Z",
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
        "CreateTime": "2026-08-05T02:15:43Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:16:28Z",
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
        "RequestId": "800c2e0c-b8d9-434e-aac2-7bf97f9478ea",
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

## TC-08（seq=8）: 有参考图 — 1080P × 10s（2 张图）
> 执行时间：2026/8/5 10:19:16  |  模型：`kling-video-o1`  |  标签：standard

> 💡 ref-1080P 164384×10=1643840

### 调用参数
```json
{
  "model": "kling-video-o1",
  "prompt": "参考图中的风景延伸为动态视频，轻风拂过水面，波光粼粼",
  "duration": 10,
  "size": "1080P",
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg",
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809801856-14qkezq2.jpeg"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 10 | **10** | 已传，直接使用 |
| size（分辨率） | 1080P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 2 张 | — | 参考图生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **821,920** |
| 预期 USD | $1.6438 |
| 预期 RMB | ¥12.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.24s |
| task_id | `task_yKHPmmJqbcTILGUBL73XSBPkFLTB2Gdh` |

```json
{
  "id": "task_yKHPmmJqbcTILGUBL73XSBPkFLTB2Gdh",
  "task_id": "task_yKHPmmJqbcTILGUBL73XSBPkFLTB2Gdh",
  "object": "video",
  "model": "kling-video-o1",
  "status": "queued",
  "progress": 0,
  "created_at": 1785896205
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 16,903,135 | — |
| 提交后可用 Quota | 16,081,215 | — |
| **预扣金额** | **821,920** | **¥12.0000** |
| 预期扣减 | 821,920 | ¥12.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 821,920 | — |
| request_id | 20260805021644945669000f61X4gQRsO0bxLE7 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 11 |
| 完成耗时 | 150.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/dd41eb615001834814617658207/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **10.042s** |
| 计费参考时长 | 10s |
| 时长差值 | 0.042s |
| 输出分辨率 | 1076×1920 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 821,920 | ¥12.0000 |
| 平台记录最终消费 | 825,372 | **¥12.050431** |
| 差额 | — | +¥0.0504（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 109,
    "created_at": 1785896205,
    "updated_at": 1785896349,
    "task_id": "task_yKHPmmJqbcTILGUBL73XSBPkFLTB2Gdh",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "12.050431",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/dd41eb615001834814617658207/aigcVideoGenFile.mp4",
    "submit_time": 1785896205,
    "start_time": 1785896220,
    "finish_time": 1785896349,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-o1",
      "origin_model_name": "kling-video-o1"
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
            "FileInfos": [
              {
                "Category": "Image",
                "FileId": "",
                "KeepOriginalSound": "",
                "ObjectId": "",
                "ReferenceType": "",
                "Text": "",
                "Type": "Url",
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg",
                "Usage": "FirstFrame",
                "VoiceId": ""
              },
              {
                "Category": "Image",
                "FileId": "",
                "KeepOriginalSound": "",
                "ObjectId": "",
                "ReferenceType": "",
                "Text": "",
                "Type": "Url",
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809801856-14qkezq2.jpeg",
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
            "Prompt": "参考图中的风景延伸为动态视频，轻风拂过水面，波光粼粼",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:19:09Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/dd41eb615001834814617658207/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 16739243,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 10.042,
                  "Height": 1920,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 21011935,
                  "VideoDuration": 10.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 16737134,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 1920,
                      "Width": 1076
                    }
                  ],
                  "Width": 1076
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
          "TaskId": "1500065731-AigcVideoTask-d44fa4d7dd30960473a5be4f0cd2b9dat"
        },
        "BeginProcessTime": "2026-08-05T02:16:45Z",
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
        "CreateTime": "2026-08-05T02:16:45Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:19:05Z",
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
        "RequestId": "a07fe061-ff6c-4ffa-bf21-5f47627473cc",
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

