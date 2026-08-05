# veo-video-3.1 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-31（seq=31）: 无声（默认）— d=8 固定
> 执行时间：2026/8/5 11:01:12  |  模型：`veo-video-3.1`  |  标签：standard

> 💡 silent-720P 205479×8=1643832（GV duration 固定 8s）

### 调用参数
```json
{
  "model": "veo-video-3.1",
  "prompt": "极光在冰岛上空绽放，湖面倒映，绿色光带舞动"
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
| 预期 Quota | **821,916** |
| 预期 USD | $1.6438 |
| 预期 RMB | ¥12.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.14s |
| task_id | `task_dVGCmisA6u36VBGt5Agur81IAimqfBDl` |

```json
{
  "id": "task_dVGCmisA6u36VBGt5Agur81IAimqfBDl",
  "task_id": "task_dVGCmisA6u36VBGt5Agur81IAimqfBDl",
  "object": "video",
  "model": "veo-video-3.1",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898781
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 10,416,500 | — |
| 提交后可用 Quota | 9,594,584 | — |
| **预扣金额** | **821,916** | **¥12.0000** |
| 预期扣减 | 821,916 | ¥12.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 821,916 | — |
| request_id | 20260805025940902464000f61X4gQRpbWSwJk0 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 7 |
| 完成耗时 | 90.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/1518ec2b5001834814635396299/aigcVideoGenFile.mp4

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
| 预扣金额 | 821,916 | ¥12.0000 |
| 平台记录最终消费 | 821,916 | **¥11.999974** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 132,
    "created_at": 1785898781,
    "updated_at": 1785898867,
    "task_id": "task_dVGCmisA6u36VBGt5Agur81IAimqfBDl",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "11.999974",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/1518ec2b5001834814635396299/aigcVideoGenFile.mp4",
    "submit_time": 1785898781,
    "start_time": 1785898786,
    "finish_time": 1785898867,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "veo-video-3.1",
      "origin_model_name": "veo-video-3.1"
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
            "ModelName": "GV",
            "ModelVersion": "3.1",
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
            "Prompt": "极光在冰岛上空绽放，湖面倒映，绿色光带舞动",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:01:07Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/1518ec2b5001834814635396299/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 3748748,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 3748748,
                  "VideoDuration": 8,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 3736605,
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
          "TaskId": "1500065731-AigcVideoTask-c0bb19bde45cc273e8078850c0825bf5t"
        },
        "BeginProcessTime": "2026-08-05T02:59:41Z",
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
        "CreateTime": "2026-08-05T02:59:41Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:00:54Z",
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
        "RequestId": "7cf744ca-5358-49ae-9fb4-609640cf4e8d",
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

## TC-32（seq=32）: 有声 — audio_generation=Enabled
> 执行时间：2026/8/5 11:02:44  |  模型：`veo-video-3.1`  |  标签：standard

> 💡 audio-720P 410959×8=3287672

### 调用参数
```json
{
  "model": "veo-video-3.1",
  "prompt": "极光在冰岛上空绽放，湖面倒映，绿色光带舞动",
  "metadata": {
    "audio_generation": "Enabled"
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | **8** | 未传，使用表达式默认值 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | Enabled | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **1,643,836** |
| 预期 USD | $3.2877 |
| 预期 RMB | ¥24.0000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_cZj3laJC15iUKNsQr8wJmQBMhc8mLAon` |

```json
{
  "id": "task_cZj3laJC15iUKNsQr8wJmQBMhc8mLAon",
  "task_id": "task_cZj3laJC15iUKNsQr8wJmQBMhc8mLAon",
  "object": "video",
  "model": "veo-video-3.1",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898872
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 9,589,619 | — |
| 提交后可用 Quota | 7,945,783 | — |
| **预扣金额** | **1,643,836** | **¥24.0000** |
| 预期扣减 | 1,643,836 | ¥24.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 1,643,836 | — |
| request_id | 20260805030112667904000f61X4gQRWNGQ9H11 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 7 |
| 完成耗时 | 90.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/5d0d3fbe5001834814636144168/aigcVideoGenFile.mp4

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
| 预扣金额 | 1,643,836 | ¥24.0000 |
| 平台记录最终消费 | 1,643,836 | **¥24.000006** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 133,
    "created_at": 1785898872,
    "updated_at": 1785898963,
    "task_id": "task_cZj3laJC15iUKNsQr8wJmQBMhc8mLAon",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "24.000006",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/5d0d3fbe5001834814636144168/aigcVideoGenFile.mp4",
    "submit_time": 1785898872,
    "start_time": 1785898883,
    "finish_time": 1785898963,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "veo-video-3.1",
      "origin_model_name": "veo-video-3.1"
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
            "ModelName": "GV",
            "ModelVersion": "3.1",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "16:9",
              "AudioGeneration": "Enabled",
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
            "Prompt": "极光在冰岛上空绽放，湖面倒映，绿色光带舞动",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:02:43Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/5d0d3fbe5001834814636144168/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 8,
                  "AudioStreamSet": [
                    {
                      "Bitrate": 256576,
                      "Channel": 0,
                      "Codec": "aac",
                      "Codecs": "",
                      "Loudness": 0,
                      "SamplingRate": 48000
                    }
                  ],
                  "Bitrate": 4407260,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 4407260,
                  "VideoDuration": 8,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 4133651,
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
          "TaskId": "1500065731-AigcVideoTask-1c68fff3936dd925d8f193d87021b325t"
        },
        "BeginProcessTime": "2026-08-05T03:01:12Z",
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
        "CreateTime": "2026-08-05T03:01:12Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:02:30Z",
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
        "RequestId": "2d1caa1a-b6d5-4799-b050-054b6ed0a363",
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

