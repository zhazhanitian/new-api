# hailuo-video-02 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-25（seq=25）: 默认（768P × 6s）
> 执行时间：2026/8/5 10:49:46  |  模型：`hailuo-video-02`  |  标签：standard

> 💡 768P 45205×6=271230

### 调用参数
```json
{
  "model": "hailuo-video-02",
  "prompt": "清晨薄雾中的古镇水乡，乌篷船缓缓划过，倒影在水面荡漾"
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
| 耗时 | 0.19s |
| task_id | `task_XmiRbO3cjZGiPuHMdKaTJ6NCOdck4xvD` |

```json
{
  "id": "task_XmiRbO3cjZGiPuHMdKaTJ6NCOdck4xvD",
  "task_id": "task_XmiRbO3cjZGiPuHMdKaTJ6NCOdck4xvD",
  "object": "video",
  "model": "hailuo-video-02",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898110
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
| 轮询次数 | 6 |
| 完成耗时 | 75.0s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/61fbd6f75001834814721031576/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.875s** |
| 计费参考时长 | 6s |
| 时长差值 | - |
| 输出分辨率 | 1366×768 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 132,789 | **¥1.938719** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 126,
    "created_at": 1785898110,
    "updated_at": 1785898173,
    "task_id": "task_XmiRbO3cjZGiPuHMdKaTJ6NCOdck4xvD",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.938719",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/61fbd6f75001834814721031576/aigcVideoGenFile.mp4",
    "submit_time": 1785898110,
    "start_time": 1785898124,
    "finish_time": 1785898173,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-02",
      "origin_model_name": "hailuo-video-02"
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
            "ModelVersion": "02",
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
            "Prompt": "清晨薄雾中的古镇水乡，乌篷船缓缓划过，倒影在水面荡漾",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:49:33Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/61fbd6f75001834814721031576/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 1256476,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.875,
                  "Height": 768,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 922725,
                  "VideoDuration": 5.875,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 1251174,
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
          "TaskId": "1500065731-AigcVideoTask-54faaf97855204e783eb253a76ec149dt"
        },
        "BeginProcessTime": "2026-08-05T02:48:30Z",
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
        "CreateTime": "2026-08-05T02:48:30Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:49:31Z",
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
        "RequestId": "9e7b107a-7473-4eda-a740-8e72e3201f21",
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

## TC-26（seq=26）: duration=10 + size=1080P
> 执行时间：2026/8/5 10:53:03  |  模型：`hailuo-video-02`  |  标签：standard

> 💡 1080P 79452×10=794520

### 调用参数
```json
{
  "model": "hailuo-video-02",
  "prompt": "清晨薄雾中的古镇水乡，乌篷船缓缓划过，倒影在水面荡漾",
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
| 耗时 | 0.16s |
| task_id | `task_gYRwF9HJqGlqTna8Aaz0RaQot4lO9hYU` |

```json
{
  "id": "task_gYRwF9HJqGlqTna8Aaz0RaQot4lO9hYU",
  "task_id": "task_gYRwF9HJqGlqTna8Aaz0RaQot4lO9hYU",
  "object": "video",
  "model": "hailuo-video-02",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898186
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 397,260 | ¥5.8000 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 14 |
| 完成耗时 | 195.2s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/cc39e6025001834814634519393/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **10.125s** |
| 计费参考时长 | 10s |
| 时长差值 | - |
| 输出分辨率 | 1920×1080 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 402,225 | **¥5.872485** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 127,
    "created_at": 1785898186,
    "updated_at": 1785898383,
    "task_id": "task_gYRwF9HJqGlqTna8Aaz0RaQot4lO9hYU",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "5.872485",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/cc39e6025001834814634519393/aigcVideoGenFile.mp4",
    "submit_time": 1785898186,
    "start_time": 1785898189,
    "finish_time": 1785898382,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-02",
      "origin_model_name": "hailuo-video-02"
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
            "ModelVersion": "02",
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
            "Prompt": "清晨薄雾中的古镇水乡，乌篷船缓缓划过，倒影在水面荡漾",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:53:03Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/cc39e6025001834814634519393/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 3648758,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 10.125,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 4617960,
                  "VideoDuration": 10.125,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 3644670,
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
          "TaskId": "1500065731-AigcVideoTask-043c7bb62c46762247e1bec70a915089t"
        },
        "BeginProcessTime": "2026-08-05T02:49:47Z",
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
        "CreateTime": "2026-08-05T02:49:47Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:52:52Z",
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
        "RequestId": "4c78466e-0ec0-4346-9a0b-28dba7b4ed9a",
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

