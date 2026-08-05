# pixverse-video-v5.6 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-42（seq=42）: 720p × 5s（无声）
> 执行时间：2026/8/5 11:21:02  |  模型：`pixverse-video-v5.6`  |  标签：standard

> 💡 720p 43151×5=215755

### 调用参数
```json
{
  "model": "pixverse-video-v5.6",
  "prompt": "节日烟花在夜空中绽放，五彩斑斓，倒映在湖面"
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
| 预期 Quota | **107,878** |
| 预期 USD | $0.2158 |
| 预期 RMB | ¥1.5750 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_sDs9NllBbcE0jHwzUx8epSMlP8GuoVoR` |

```json
{
  "id": "task_sDs9NllBbcE0jHwzUx8epSMlP8GuoVoR",
  "task_id": "task_sDs9NllBbcE0jHwzUx8epSMlP8GuoVoR",
  "object": "video",
  "model": "pixverse-video-v5.6",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900000
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 4,784,820 | — |
| 提交后可用 Quota | 4,700,915 | — |
| **预扣金额** | **83,905** | **¥1.2250** |
| 预期扣减 | 107,878 | ¥1.5750 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | 83,905 | — |
| request_id | 20260805032000307279000f61X4gQRzHljtl1N | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 5 |
| 完成耗时 | 60.0s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/c64b64d65001834814641071264/aigcVideoGenFile.mp4

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
| 预扣金额 | 83,905 | ¥1.2250 |
| 平台记录最终消费 | 84,609 | **¥1.235291** |
| 差额 | — | +¥0.0103（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 143,
    "created_at": 1785900000,
    "updated_at": 1785900061,
    "task_id": "task_sDs9NllBbcE0jHwzUx8epSMlP8GuoVoR",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.235291",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/c64b64d65001834814641071264/aigcVideoGenFile.mp4",
    "submit_time": 1785900000,
    "start_time": 1785900012,
    "finish_time": 1785900061,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "pixverse-video-v5.6",
      "origin_model_name": "pixverse-video-v5.6"
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
            "ModelVersion": "v5.6",
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
            "Prompt": "节日烟花在夜空中绽放，五彩斑斓，倒映在湖面",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:21:01Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/c64b64d65001834814641071264/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 3823487,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 2409753,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 3819565,
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
          "TaskId": "1500065731-AigcVideoTask-db808f27b2658a79aa04a772682482b6t"
        },
        "BeginProcessTime": "2026-08-05T03:20:00Z",
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
        "CreateTime": "2026-08-05T03:20:00Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:20:46Z",
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
        "RequestId": "909d1df1-c223-47e4-a0e8-bdd30c03c038",
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

## TC-43（seq=43）: 1080p × 10s（无声）
> 执行时间：2026/8/5 11:21:18  |  模型：`pixverse-video-v5.6`  |  标签：standard

> 💡 1080p 71918×10=719180

### 调用参数
```json
{
  "model": "pixverse-video-v5.6",
  "prompt": "节日烟花在夜空中绽放，五彩斑斓，倒映在湖面",
  "duration": 10,
  "size": "1080p"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 10 | **10** | 已传，直接使用 |
| size（分辨率） | 1080p | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **359,590** |
| 预期 USD | $0.7192 |
| 预期 RMB | ¥5.2500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.16s |
| task_id | `task_vgywS1VplAV1iN3tR1LlOlNKhzIds6u3` |

```json
{
  "id": "task_vgywS1VplAV1iN3tR1LlOlNKhzIds6u3",
  "task_id": "task_vgywS1VplAV1iN3tR1LlOlNKhzIds6u3",
  "object": "video",
  "model": "pixverse-video-v5.6",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900062
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 359,590 | ¥5.2500 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **FAILURE** |
| 轮询次数 | 2 |
| 完成耗时 | 15.0s |
| progress | 100% |

**视频 URL**: task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.

**失败原因**: task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 359,590 | **¥5.250014** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 144,
    "created_at": 1785900062,
    "updated_at": 1785900077,
    "task_id": "task_vgywS1VplAV1iN3tR1LlOlNKhzIds6u3",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "5.250014",
    "action": "generate",
    "status": "FAILURE",
    "fail_reason": "task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.",
    "result_url": "task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.",
    "submit_time": 1785900062,
    "start_time": 0,
    "finish_time": 1785900077,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "pixverse-video-v5.6",
      "origin_model_name": "pixverse-video-v5.6"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ErrCode": 70000,
          "ErrCodeExt": "InternalError",
          "Input": {
            "EnhancePrompt": "",
            "FileInfos": [],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "PixVerse",
            "ModelVersion": "v5.6",
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
              "Resolution": "1080p",
              "StorageMode": "Temporary"
            },
            "Prompt": "节日烟花在夜空中绽放，五彩斑斓，倒映在湖面",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "task failed with status: FAIL, message: create pixverse t2v task failed. ret:-4,msg:pixverse api error, ErrCode:400017, ErrMsg:1080p does not support durations longer than 8 seconds.",
          "Output": {
            "FileInfos": [],
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
          "TaskId": "1500065731-AigcVideoTask-799c931b67414af7385a856ae9d6b6d2t"
        },
        "BeginProcessTime": "2026-08-05T03:21:02Z",
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
        "CreateTime": "2026-08-05T03:21:02Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:21:10Z",
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
        "RequestId": "cf630add-d355-4227-b567-1a5f36e74ba2",
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

