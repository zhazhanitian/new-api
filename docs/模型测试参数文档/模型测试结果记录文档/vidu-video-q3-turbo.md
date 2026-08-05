# vidu-video-q3-turbo 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-23（seq=23）: 720P × 5s
> 执行时间：2026/8/5 10:45:58  |  模型：`vidu-video-q3-turbo`  |  标签：standard

> 💡 720P 51370×5=256850

### 调用参数
```json
{
  "model": "vidu-video-q3-turbo",
  "prompt": "古堡在暮色中轮廓渐现，蝙蝠掠过天际，神秘哥特风",
  "duration": 5,
  "size": "720P",
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 5 | **5** | 已传，直接使用 |
| size（分辨率） | 720P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 1 张 | — | 参考图生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **128,425** |
| 预期 USD | $0.2569 |
| 预期 RMB | ¥1.8750 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.18s |
| task_id | `task_YUDgcH9DS12bJL00BLPPA1ebsn4vCZlC` |

```json
{
  "id": "task_YUDgcH9DS12bJL00BLPPA1ebsn4vCZlC",
  "task_id": "task_YUDgcH9DS12bJL00BLPPA1ebsn4vCZlC",
  "object": "video",
  "model": "vidu-video-q3-turbo",
  "status": "queued",
  "progress": 0,
  "created_at": 1785897881
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 128,425 | ¥1.8750 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 6 |
| 完成耗时 | 75.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/c08ddfd45001834814634015118/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.042s** |
| 计费参考时长 | 5s |
| 时长差值 | - |
| 输出分辨率 | 720×1282 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 129,503 | **¥1.890744** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 124,
    "created_at": 1785897881,
    "updated_at": 1785897947,
    "task_id": "task_YUDgcH9DS12bJL00BLPPA1ebsn4vCZlC",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.890744",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/c08ddfd45001834814634015118/aigcVideoGenFile.mp4",
    "submit_time": 1785897881,
    "start_time": 1785897882,
    "finish_time": 1785897947,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q3-turbo",
      "origin_model_name": "vidu-video-q3-turbo"
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
              "AspectRatio": "16:9",
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
              "Resolution": "720P",
              "StorageMode": "Temporary"
            },
            "Prompt": "古堡在暮色中轮廓渐现，蝙蝠掠过天际，神秘哥特风",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:45:47Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/c08ddfd45001834814634015118/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 5,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 129422,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 48000
                    }
                  ],
                  "Bitrate": 4948238,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 1282,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 3118627,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 4810656,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "SDR"
                      },
                      "Fps": 24,
                      "Height": 1282,
                      "Width": 720
                    }
                  ],
                  "Width": 720
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
          "TaskId": "1500065731-AigcVideoTask-09856699a79c70b16734205e279298a1t"
        },
        "BeginProcessTime": "2026-08-05T02:44:41Z",
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
        "CreateTime": "2026-08-05T02:44:41Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:45:31Z",
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
        "RequestId": "cfef63b7-53bc-4e3e-aeb2-91e634ab8047",
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

## TC-24（seq=24）: 1080P × 8s
> 执行时间：2026/8/5 10:48:30  |  模型：`vidu-video-q3-turbo`  |  标签：standard

> 💡 1080P 60000×8=480000

### 调用参数
```json
{
  "model": "vidu-video-q3-turbo",
  "prompt": "古堡在暮色中轮廓渐现，蝙蝠掠过天际，神秘哥特风",
  "duration": 8,
  "size": "1080P",
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809693246-v8bmrsse.jpeg"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 8 | **8** | 已传，直接使用 |
| size（分辨率） | 1080P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 1 张 | — | 参考图生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **240,000** |
| 预期 USD | $0.4800 |
| 预期 RMB | ¥3.5040 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_FdiUUaUPayEsi1qEHVkKzkwa0asF7GW2` |

```json
{
  "id": "task_FdiUUaUPayEsi1qEHVkKzkwa0asF7GW2",
  "task_id": "task_FdiUUaUPayEsi1qEHVkKzkwa0asF7GW2",
  "object": "video",
  "model": "vidu-video-q3-turbo",
  "status": "queued",
  "progress": 0,
  "created_at": 1785897958
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 240,000 | ¥3.5040 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 11 |
| 完成耗时 | 150.2s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/176f23e15001834814635498007/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8.042s** |
| 计费参考时长 | 8s |
| 时长差值 | - |
| 输出分辨率 | 1080×1922 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 241,260 | **¥3.522396** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 125,
    "created_at": 1785897958,
    "updated_at": 1785898108,
    "task_id": "task_FdiUUaUPayEsi1qEHVkKzkwa0asF7GW2",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.522396",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/176f23e15001834814635498007/aigcVideoGenFile.mp4",
    "submit_time": 1785897958,
    "start_time": 1785897963,
    "finish_time": 1785898108,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q3-turbo",
      "origin_model_name": "vidu-video-q3-turbo"
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
              "Resolution": "1080P",
              "StorageMode": "Temporary"
            },
            "Prompt": "古堡在暮色中轮廓渐现，蝙蝠掠过天际，神秘哥特风",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:48:28Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/176f23e15001834814635498007/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 8,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 97321,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 44100
                    }
                  ],
                  "Bitrate": 13669478,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8.042,
                  "Height": 1922,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 13741243,
                  "VideoDuration": 8.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 13564345,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "SDR"
                      },
                      "Fps": 24,
                      "Height": 1922,
                      "Width": 1080
                    }
                  ],
                  "Width": 1080
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
          "TaskId": "1500065731-AigcVideoTask-1c60043d8328c4bc214f80e8d10372fct"
        },
        "BeginProcessTime": "2026-08-05T02:45:58Z",
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
        "CreateTime": "2026-08-05T02:45:58Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:48:17Z",
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
        "RequestId": "89189494-07e9-4514-93e5-e97c93911060",
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

