# pixverse-video-c1 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-46（seq=46）: 无声 — 720p × 5s
> 执行时间：2026/8/5 11:25:38  |  模型：`pixverse-video-c1`  |  标签：standard

> 💡 silent-720p 40137×5=200685

### 调用参数
```json
{
  "model": "pixverse-video-c1",
  "prompt": "迷雾森林中的精灵，荧光粒子飘落，梦幻奇幻风格"
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
| 预期 Quota | **100,343** |
| 预期 USD | $0.2007 |
| 预期 RMB | ¥1.4650 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.14s |
| task_id | `task_wUrcVsnfubKWsTgkrjPD5mw0UJpZlQp7` |

```json
{
  "id": "task_wUrcVsnfubKWsTgkrjPD5mw0UJpZlQp7",
  "task_id": "task_wUrcVsnfubKWsTgkrjPD5mw0UJpZlQp7",
  "object": "video",
  "model": "pixverse-video-c1",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900262
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 4,259,553 | — |
| 提交后可用 Quota | 4,179,073 | — |
| **预扣金额** | **80,480** | **¥1.1750** |
| 预期扣减 | 100,343 | ¥1.4650 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | 80,480 | — |
| request_id | 20260805032422181155000f61X4gQRWCDodf52 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 6 |
| 完成耗时 | 75.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/8762daa55001834814640682893/aigcVideoGenFile.mp4

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
| 预扣金额 | 80,480 | ¥1.1750 |
| 平台记录最终消费 | 81,156 | **¥1.184878** |
| 差额 | — | +¥0.0099（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 147,
    "created_at": 1785900262,
    "updated_at": 1785900335,
    "task_id": "task_wUrcVsnfubKWsTgkrjPD5mw0UJpZlQp7",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.184878",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/8762daa55001834814640682893/aigcVideoGenFile.mp4",
    "submit_time": 1785900262,
    "start_time": 1785900271,
    "finish_time": 1785900335,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "pixverse-video-c1",
      "origin_model_name": "pixverse-video-c1"
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
            "ModelVersion": "c1",
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
            "Prompt": "迷雾森林中的精灵，荧光粒子飘落，梦幻奇幻风格",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:25:35Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/8762daa55001834814640682893/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 4484039,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 2826066,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 4478623,
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
          "TaskId": "1500065731-AigcVideoTask-30a79a1bc08007e3e100eb12cd285182t"
        },
        "BeginProcessTime": "2026-08-05T03:24:22Z",
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
        "CreateTime": "2026-08-05T03:24:22Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:25:21Z",
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
        "RequestId": "60333faf-f549-4d12-9f31-d202e20552ea",
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

## TC-47（seq=47）: 有声 — 1080p × 8s
> 执行时间：2026/8/5 11:27:40  |  模型：`pixverse-video-c1`  |  标签：standard

> 💡 audio-1080p 96438×8=771504

### 调用参数
```json
{
  "model": "pixverse-video-c1",
  "prompt": "迷雾森林中的精灵，荧光粒子飘落，梦幻奇幻风格",
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
| 预期 Quota | **385,752** |
| 预期 USD | $0.7715 |
| 预期 RMB | ¥5.6320 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_gr4hqlwD90FExmg63U8Cy6JZzzYwMg4I` |

```json
{
  "id": "task_gr4hqlwD90FExmg63U8Cy6JZzzYwMg4I",
  "task_id": "task_gr4hqlwD90FExmg63U8Cy6JZzzYwMg4I",
  "object": "video",
  "model": "pixverse-video-c1",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900339
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 4,177,132 | — |
| 提交后可用 Quota | 3,791,380 | — |
| **预扣金额** | **385,752** | **¥5.6320** |
| 预期扣减 | 385,752 | ¥5.6320 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 9 |
| 完成耗时 | 120.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/4e30375b5001834814646375050/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8.042s** |
| 计费参考时长 | 8s |
| 时长差值 | 0.042s |
| 输出分辨率 | 1920×1080 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 385,752 | ¥5.6320 |
| 平台记录最终消费 | 387,777 | **¥5.661544** |
| 差额 | — | +¥0.0296（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 148,
    "created_at": 1785900339,
    "updated_at": 1785900448,
    "task_id": "task_gr4hqlwD90FExmg63U8Cy6JZzzYwMg4I",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "5.661544",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/4e30375b5001834814646375050/aigcVideoGenFile.mp4",
    "submit_time": 1785900339,
    "start_time": 1785900351,
    "finish_time": 1785900448,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "pixverse-video-c1",
      "origin_model_name": "pixverse-video-c1"
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
            "ModelVersion": "c1",
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
            "Prompt": "迷雾森林中的精灵，荧光粒子飘落，梦幻奇幻风格",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:27:28Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/4e30375b5001834814646375050/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 7849918,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8.042,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 7891131,
                  "VideoDuration": 8.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 7845755,
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
          "TaskId": "1500065731-AigcVideoTask-b908df823e99e19bbb2494e6de532cdbt"
        },
        "BeginProcessTime": "2026-08-05T03:25:39Z",
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
        "CreateTime": "2026-08-05T03:25:39Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:27:23Z",
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
        "RequestId": "624c92f6-cc4f-4d14-a04b-c3d282f9542a",
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

