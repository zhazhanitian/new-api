# sora-video-2.0 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-35（seq=35）: 不传 duration（API 默认 8s，snap→8）
> 执行时间：2026/8/5 11:08:04  |  模型：`sora-video-2.0`  |  标签：standard

> 💡 d0=0→d=8；720P 102740×8=821920

### 调用参数
```json
{
  "model": "sora-video-2.0",
  "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | **8** | 未传，使用表达式默认值 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
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
| task_id | `task_xqI7lRChgCVZeqvxMttsTiOLWa6QzXUT` |

```json
{
  "id": "task_xqI7lRChgCVZeqvxMttsTiOLWa6QzXUT",
  "task_id": "task_xqI7lRChgCVZeqvxMttsTiOLWa6QzXUT",
  "object": "video",
  "model": "sora-video-2.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785899163
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
| 轮询次数 | 9 |
| 完成耗时 | 120.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/749269f25001834814721820241/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8s** |
| 计费参考时长 | 8s |
| 时长差值 | - |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 410,960 | **¥6.000016** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 136,
    "created_at": 1785899163,
    "updated_at": 1785899270,
    "task_id": "task_xqI7lRChgCVZeqvxMttsTiOLWa6QzXUT",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.000016",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/749269f25001834814721820241/aigcVideoGenFile.mp4",
    "submit_time": 1785899163,
    "start_time": 1785899173,
    "finish_time": 1785899270,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "sora-video-2.0",
      "origin_model_name": "sora-video-2.0"
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
            "ModelName": "OS",
            "ModelVersion": "2.0",
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
            "Prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:07:50Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/749269f25001834814721820241/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 8,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 128173,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 96000
                    }
                  ],
                  "Bitrate": 8193330,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 8193330,
                  "VideoDuration": 8,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 8047443,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 30,
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
          "TaskId": "1500065731-AigcVideoTask-e27f5babb970e74c7b1c44bd3104b64bt"
        },
        "BeginProcessTime": "2026-08-05T03:06:03Z",
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
        "CreateTime": "2026-08-05T03:06:03Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:07:49Z",
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
        "RequestId": "6aa67be3-b4c0-414a-a845-c294dd2f91f7",
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

## TC-36（seq=36）: duration=5（≤6 → snap 到 4s）
> 执行时间：2026/8/5 11:09:36  |  模型：`sora-video-2.0`  |  标签：standard

> 💡 d0=5≤6→d=4；720P 102740×4=410960

### 调用参数
```json
{
  "model": "sora-video-2.0",
  "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
  "duration": 5
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 5 | **4** | 已传，直接使用 |
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
| 耗时 | 0.15s |
| task_id | `task_k3mNbfR5fwqYwfoQLSN1dNKwpWHVWUm4` |

```json
{
  "id": "task_k3mNbfR5fwqYwfoQLSN1dNKwpWHVWUm4",
  "task_id": "task_k3mNbfR5fwqYwfoQLSN1dNKwpWHVWUm4",
  "object": "video",
  "model": "sora-video-2.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785899284
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 205,480 | ¥3.0000 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 7 |
| 完成耗时 | 90.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/819323325001834814633616110/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **4.1s** |
| 计费参考时长 | 4s |
| 时长差值 | - |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 210,616 | **¥3.074994** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 137,
    "created_at": 1785899284,
    "updated_at": 1785899367,
    "task_id": "task_k3mNbfR5fwqYwfoQLSN1dNKwpWHVWUm4",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.074994",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/819323325001834814633616110/aigcVideoGenFile.mp4",
    "submit_time": 1785899284,
    "start_time": 1785899286,
    "finish_time": 1785899367,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "sora-video-2.0",
      "origin_model_name": "sora-video-2.0"
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
            "ModelName": "OS",
            "ModelVersion": "2.0",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 4,
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
            "Prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:09:27Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/819323325001834814633616110/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 4.1,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 124914,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 96000
                    }
                  ],
                  "Bitrate": 7050608,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 4.1,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 3613437,
                  "VideoDuration": 4,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 7071790,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 30,
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
          "TaskId": "1500065731-AigcVideoTask-cca158f25cc96a3a1f24187ed07b4a25t"
        },
        "BeginProcessTime": "2026-08-05T03:08:04Z",
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
        "CreateTime": "2026-08-05T03:08:04Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:09:18Z",
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
        "RequestId": "f96a4660-1baa-436b-bafd-f8f8c1d9f0a7",
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

## TC-37（seq=37）: duration=12（>10 → snap 到 12s）
> 执行时间：2026/8/5 11:12:38  |  模型：`sora-video-2.0`  |  标签：standard

> 💡 d0=12>10→d=12；720P 102740×12=1232880

### 调用参数
```json
{
  "model": "sora-video-2.0",
  "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
  "duration": 12
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 12 | **12** | 已传，直接使用 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **616,440** |
| 预期 USD | $1.2329 |
| 预期 RMB | ¥9.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.22s |
| task_id | `task_Q895GPnfH4krq1QaxUBt09Ig3LO7zXuN` |

```json
{
  "id": "task_Q895GPnfH4krq1QaxUBt09Ig3LO7zXuN",
  "task_id": "task_Q895GPnfH4krq1QaxUBt09Ig3LO7zXuN",
  "object": "video",
  "model": "sora-video-2.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785899376
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 616,440 | ¥9.0000 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 13 |
| 完成耗时 | 180.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/5d0bb89d5001834814636141881/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **12.1s** |
| 计费参考时长 | 12s |
| 时长差值 | - |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 621,577 | **¥9.075024** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 138,
    "created_at": 1785899376,
    "updated_at": 1785899544,
    "task_id": "task_Q895GPnfH4krq1QaxUBt09Ig3LO7zXuN",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "9.075024",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/5d0bb89d5001834814636141881/aigcVideoGenFile.mp4",
    "submit_time": 1785899376,
    "start_time": 1785899383,
    "finish_time": 1785899544,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "sora-video-2.0",
      "origin_model_name": "sora-video-2.0"
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
            "ModelName": "OS",
            "ModelVersion": "2.0",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 12,
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
            "Prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:12:24Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/5d0bb89d5001834814636141881/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 12.1,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 127392,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 96000
                    }
                  ],
                  "Bitrate": 7626602,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 12.1,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 11535237,
                  "VideoDuration": 12,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 7546995,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 30,
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
          "TaskId": "1500065731-AigcVideoTask-aacb524ec65e894ac7be239abbf8a23at"
        },
        "BeginProcessTime": "2026-08-05T03:09:36Z",
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
        "CreateTime": "2026-08-05T03:09:36Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:12:11Z",
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
        "RequestId": "46acf463-7121-49b2-9645-84a1d3ae6ec9",
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

## EC-10（seq=57）: duration=6（≤6 → snap 到 4s）
> 执行时间：2026/8/5 11:39:27  |  模型：`sora-video-2.0`  |  标签：edge

> 💡 d0=6≤6→d=4；720P 102740×4=410960。snap 下界验证。

### 调用参数
```json
{
  "model": "sora-video-2.0",
  "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
  "duration": 6
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 6 | **4** | 已传，直接使用 |
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
| 耗时 | 0.17s |
| task_id | `task_Lv3Q6SauOryXNQyWqUYdQbXcIRCSnJcq` |

```json
{
  "id": "task_Lv3Q6SauOryXNQyWqUYdQbXcIRCSnJcq",
  "task_id": "task_Lv3Q6SauOryXNQyWqUYdQbXcIRCSnJcq",
  "object": "video",
  "model": "sora-video-2.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785901076
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 2,460,982 | — |
| 提交后可用 Quota | 2,255,502 | — |
| **预扣金额** | **205,480** | **¥3.0000** |
| 预期扣减 | 205,480 | ¥3.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 205,480 | — |
| request_id | 2026080503375643649000f61X4gQRMr5cyfA2 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 7 |
| 完成耗时 | 90.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/4998a8a35001834814646186742/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **4.1s** |
| 计费参考时长 | 4s |
| 时长差值 | 0.100s |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 205,480 | ¥3.0000 |
| 平台记录最终消费 | 210,616 | **¥3.074994** |
| 差额 | — | +¥0.0750（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 156,
    "created_at": 1785901076,
    "updated_at": 1785901158,
    "task_id": "task_Lv3Q6SauOryXNQyWqUYdQbXcIRCSnJcq",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.074994",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/4998a8a35001834814646186742/aigcVideoGenFile.mp4",
    "submit_time": 1785901076,
    "start_time": 1785901077,
    "finish_time": 1785901158,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "sora-video-2.0",
      "origin_model_name": "sora-video-2.0"
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
            "ModelName": "OS",
            "ModelVersion": "2.0",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 4,
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
            "Prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:39:18Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/4998a8a35001834814646186742/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 4.1,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 125132,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 96000
                    }
                  ],
                  "Bitrate": 6966950,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 4.1,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 3570562,
                  "VideoDuration": 4,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 6985752,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 30,
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
          "TaskId": "1500065731-AigcVideoTask-0be14a6279b84138332d430f883996eft"
        },
        "BeginProcessTime": "2026-08-05T03:37:56Z",
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
        "CreateTime": "2026-08-05T03:37:56Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:39:04Z",
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
        "RequestId": "f93a73f7-bcd3-4c22-9e26-650986b9c4ab",
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

## EC-11（seq=58）: duration=7（6< d ≤10 → snap 到 8s）
> 执行时间：2026/8/5 11:41:29  |  模型：`sora-video-2.0`  |  标签：edge

> 💡 d0=7→d=8；720P 102740×8=821920。snap 中间档验证。

### 调用参数
```json
{
  "model": "sora-video-2.0",
  "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
  "duration": 7
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 7 | **8** | 已传，直接使用 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
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
| task_id | `task_7WcAdhW56Df4WWAv9SS0JSOqYRK5xYEb` |

```json
{
  "id": "task_7WcAdhW56Df4WWAv9SS0JSOqYRK5xYEb",
  "task_id": "task_7WcAdhW56Df4WWAv9SS0JSOqYRK5xYEb",
  "object": "video",
  "model": "sora-video-2.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785901167
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 2,249,749 | — |
| 提交后可用 Quota | 1,838,789 | — |
| **预扣金额** | **410,960** | **¥6.0000** |
| 预期扣减 | 410,960 | ¥6.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 410,960 | — |
| request_id | 20260805033927820167000f61X4gQRUMPBWqrC | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 9 |
| 完成耗时 | 120.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/499923e45001834814646187337/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8s** |
| 计费参考时长 | 8s |
| 时长差值 | 0.000s |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 410,960 | ¥6.0000 |
| 平台记录最终消费 | 410,960 | **¥6.000016** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 157,
    "created_at": 1785901167,
    "updated_at": 1785901287,
    "task_id": "task_7WcAdhW56Df4WWAv9SS0JSOqYRK5xYEb",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.000016",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/499923e45001834814646187337/aigcVideoGenFile.mp4",
    "submit_time": 1785901167,
    "start_time": 1785901174,
    "finish_time": 1785901287,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "sora-video-2.0",
      "origin_model_name": "sora-video-2.0"
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
            "ModelName": "OS",
            "ModelVersion": "2.0",
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
              "Resolution": "720P",
              "StorageMode": "Temporary"
            },
            "Prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:41:27Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/499923e45001834814646187337/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 8,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 128121,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 96000
                    }
                  ],
                  "Bitrate": 8617820,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 8617820,
                  "VideoDuration": 8,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 8472023,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 30,
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
          "TaskId": "1500065731-AigcVideoTask-d3daffe4b447b4eabb9d0286d8c89155t"
        },
        "BeginProcessTime": "2026-08-05T03:39:28Z",
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
        "CreateTime": "2026-08-05T03:39:28Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:41:27Z",
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
        "RequestId": "14b9f96d-58ee-46b5-aa07-e7a8da59d007",
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

## EC-12（seq=59）: duration=11（>10 → snap 到 12s）
> 执行时间：2026/8/5 11:44:16  |  模型：`sora-video-2.0`  |  标签：edge

> 💡 d0=11>10→d=12；720P 102740×12=1232880。snap 上界验证。

### 调用参数
```json
{
  "model": "sora-video-2.0",
  "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
  "duration": 11
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 11 | **12** | 已传，直接使用 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **616,440** |
| 预期 USD | $1.2329 |
| 预期 RMB | ¥9.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.16s |
| task_id | `task_S47T2H2wNDa3gvrEfOsiZijjhWVau7xx` |

```json
{
  "id": "task_S47T2H2wNDa3gvrEfOsiZijjhWVau7xx",
  "task_id": "task_S47T2H2wNDa3gvrEfOsiZijjhWVau7xx",
  "object": "video",
  "model": "sora-video-2.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785901289
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 1,833,653 | — |
| 提交后可用 Quota | 1,217,213 | — |
| **预扣金额** | **616,440** | **¥9.0000** |
| 预期扣减 | 616,440 | ¥9.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 616,440 | — |
| request_id | 20260805034129620776000f61X4gQRtG1nXTwF | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 12 |
| 完成耗时 | 165.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/bdcdb2fa5001834814651517019/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **12.1s** |
| 计费参考时长 | 12s |
| 时长差值 | 0.100s |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 616,440 | ¥9.0000 |
| 平台记录最终消费 | 621,577 | **¥9.075024** |
| 差额 | — | +¥0.0750（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 158,
    "created_at": 1785901289,
    "updated_at": 1785901448,
    "task_id": "task_S47T2H2wNDa3gvrEfOsiZijjhWVau7xx",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "9.075024",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/bdcdb2fa5001834814651517019/aigcVideoGenFile.mp4",
    "submit_time": 1785901289,
    "start_time": 1785901303,
    "finish_time": 1785901448,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "sora-video-2.0",
      "origin_model_name": "sora-video-2.0"
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
            "ModelName": "OS",
            "ModelVersion": "2.0",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "",
              "ClassId": 0,
              "Duration": 12,
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
            "Prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:44:08Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/bdcdb2fa5001834814651517019/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 12.1,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 127263,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 96000
                    }
                  ],
                  "Bitrate": 8265866,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 12.1,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 12502123,
                  "VideoDuration": 12,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 8191726,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 30,
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
          "TaskId": "1500065731-AigcVideoTask-9f992e224e8afa8bcca5fca28d4a0fabt"
        },
        "BeginProcessTime": "2026-08-05T03:41:29Z",
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
        "CreateTime": "2026-08-05T03:41:29Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:44:06Z",
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
        "RequestId": "cbb24d74-d95a-483e-be43-d0ebfd3354b1",
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

## EC-13（seq=60）: size="1080P"（表达式按 1080P 计，适配器可能实际仍发 720P）
> 执行时间：2026/8/5 11:46:18  |  模型：`sora-video-2.0`  |  标签：edge

> 💡 d=8(默认)；1080P 154110×8=1232880。注意：适配器 OS 仅支持 720P，实际请求分辨率可能被降级，但计费按表达式走 1080P 档。观察实际扣费与此预期是否一致。

### 调用参数
```json
{
  "model": "sora-video-2.0",
  "prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
  "size": "1080P"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | **8** | 未传，使用表达式默认值 |
| size（分辨率） | 1080P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **616,440** |
| 预期 USD | $1.2329 |
| 预期 RMB | ¥9.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.16s |
| task_id | `task_Sz5roXb746MniM9Us1cFUFJhRnEefVBM` |

```json
{
  "id": "task_Sz5roXb746MniM9Us1cFUFJhRnEefVBM",
  "task_id": "task_Sz5roXb746MniM9Us1cFUFJhRnEefVBM",
  "object": "video",
  "model": "sora-video-2.0",
  "status": "queued",
  "progress": 0,
  "created_at": 1785901456
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 1,217,213 | — |
| 提交后可用 Quota | 600,773 | — |
| **预扣金额** | **616,440** | **¥9.0000** |
| 预期扣减 | 616,440 | ¥9.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 616,440 | — |
| request_id | 20260805034416413350000f61X4gQR1bSOSyS3 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 9 |
| 完成耗时 | 120.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/48fea91e5001834814639390711/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8s** |
| 计费参考时长 | 8s |
| 时长差值 | 0.000s |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 616,440 | ¥9.0000 |
| 平台记录最终消费 | 616,440 | **¥9.000024** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 159,
    "created_at": 1785901456,
    "updated_at": 1785901577,
    "task_id": "task_Sz5roXb746MniM9Us1cFUFJhRnEefVBM",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "9.000024",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/48fea91e5001834814639390711/aigcVideoGenFile.mp4",
    "submit_time": 1785901456,
    "start_time": 1785901464,
    "finish_time": 1785901577,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "sora-video-2.0",
      "origin_model_name": "sora-video-2.0"
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
            "ModelName": "OS",
            "ModelVersion": "2.0",
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
            "Prompt": "未来城市，高架轨道飞行器穿梭，全息广告在空中闪烁",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:46:17Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/48fea91e5001834814639390711/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 8,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 128095,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 96000
                    }
                  ],
                  "Bitrate": 8716684,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 8716684,
                  "VideoDuration": 8,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 8570819,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 30,
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
          "TaskId": "1500065731-AigcVideoTask-f2401ad063b24166983e6c65cafaca86t"
        },
        "BeginProcessTime": "2026-08-05T03:44:16Z",
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
        "CreateTime": "2026-08-05T03:44:16Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:46:14Z",
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
        "RequestId": "568e2460-6fea-4886-bf43-45d11361963b",
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

