# kling-video-2.5 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-03（seq=3）: 不传 duration/size（默认 720P × 5s）
> 执行时间：2026/8/5 10:10:22  |  模型：`kling-video-2.5`  |  标签：standard

> 💡 41096×5=205480

### 调用参数
```json
{
  "model": "kling-video-2.5",
  "prompt": "秋天的银杏大道，金黄叶片漫天飞舞，阳光透过树梢，电影感"
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
| 预期 Quota | **102,740** |
| 预期 USD | $0.2055 |
| 预期 RMB | ¥1.5000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.21s |
| task_id | `task_Nw16zKcL9DecumyDfmaUwXzHvhpGGwDJ` |

```json
{
  "id": "task_Nw16zKcL9DecumyDfmaUwXzHvhpGGwDJ",
  "task_id": "task_Nw16zKcL9DecumyDfmaUwXzHvhpGGwDJ",
  "object": "video",
  "model": "kling-video-2.5",
  "status": "queued",
  "progress": 0,
  "created_at": 1785895746
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 18,005,065 | — |
| 提交后可用 Quota | 17,902,325 | — |
| **预扣金额** | **102,740** | **¥1.5000** |
| 预期扣减 | 102,740 | ¥1.5000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 102,740 | — |
| request_id | 20260805020906122696000f61X4gQRI5PAiwmp | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 6 |
| 完成耗时 | 75.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/3f7dcbef5001834814622280004/aigcVideoGenFile.mp4

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
| 预扣金额 | 102,740 | ¥1.5000 |
| 平台记录最终消费 | 103,603 | **¥1.512604** |
| 差额 | — | +¥0.0126（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 104,
    "created_at": 1785895746,
    "updated_at": 1785895816,
    "task_id": "task_Nw16zKcL9DecumyDfmaUwXzHvhpGGwDJ",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.512604",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/3f7dcbef5001834814622280004/aigcVideoGenFile.mp4",
    "submit_time": 1785895746,
    "start_time": 1785895752,
    "finish_time": 1785895816,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-2.5",
      "origin_model_name": "kling-video-2.5"
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
            "ModelVersion": "2.5",
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
            "Prompt": "秋天的银杏大道，金黄叶片漫天飞舞，阳光透过树梢，电影感",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:10:16Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/3f7dcbef5001834814622280004/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 22461704,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 14156489,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 22458704,
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
          "TaskId": "1500065731-AigcVideoTask-07e5da908609b31f0bd775a041e9b9f1t"
        },
        "BeginProcessTime": "2026-08-05T02:09:06Z",
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
        "CreateTime": "2026-08-05T02:09:06Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:10:09Z",
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
        "RequestId": "fc8492c0-8d93-434c-9563-6304291dfcfb",
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

## TC-04（seq=4）: duration=10 + size=1080P
> 执行时间：2026/8/5 10:13:09  |  模型：`kling-video-2.5`  |  标签：standard

> 💡 68493×10=684930

### 调用参数
```json
{
  "model": "kling-video-2.5",
  "prompt": "秋天的银杏大道，金黄叶片漫天飞舞，阳光透过树梢，电影感",
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
| 预期 Quota | **342,465** |
| 预期 USD | $0.6849 |
| 预期 RMB | ¥5.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.18s |
| task_id | `task_DSAh0v05WMWjPFBlKzWhgqvrZn3tadnm` |

```json
{
  "id": "task_DSAh0v05WMWjPFBlKzWhgqvrZn3tadnm",
  "task_id": "task_DSAh0v05WMWjPFBlKzWhgqvrZn3tadnm",
  "object": "video",
  "model": "kling-video-2.5",
  "status": "queued",
  "progress": 0,
  "created_at": 1785895823
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 17,902,325 | — |
| 提交后可用 Quota | 17,559,860 | — |
| **预扣金额** | **342,465** | **¥5.0000** |
| 预期扣减 | 342,465 | ¥5.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 342,465 | — |
| request_id | 20260805021022935164000f61X4gQRud20WBad | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 12 |
| 完成耗时 | 165.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f919de9d5001834814621526745/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **10.042s** |
| 计费参考时长 | 10s |
| 时长差值 | 0.042s |
| 输出分辨率 | 1920×1080 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 342,465 | ¥5.0000 |
| 平台记录最终消费 | 343,903 | **¥5.020984** |
| 差额 | — | +¥0.0210（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 105,
    "created_at": 1785895823,
    "updated_at": 1785895977,
    "task_id": "task_DSAh0v05WMWjPFBlKzWhgqvrZn3tadnm",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "5.020984",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f919de9d5001834814621526745/aigcVideoGenFile.mp4",
    "submit_time": 1785895823,
    "start_time": 1785895832,
    "finish_time": 1785895977,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "kling-video-2.5",
      "origin_model_name": "kling-video-2.5"
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
            "ModelVersion": "2.5",
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
            "Prompt": "秋天的银杏大道，金黄叶片漫天飞舞，阳光透过树梢，电影感",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:12:58Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/f919de9d5001834814621526745/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 31437257,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 10.042,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 39461617,
                  "VideoDuration": 10.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 31435555,
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
          "TaskId": "1500065731-AigcVideoTask-3b1c107fd0def5603ab6299136804be3t"
        },
        "BeginProcessTime": "2026-08-05T02:10:23Z",
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
        "CreateTime": "2026-08-05T02:10:23Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:12:56Z",
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
        "RequestId": "adabe2d7-c443-4212-aa0d-58463ed60de2",
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

