# vidu-video-q2-pro 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-15（seq=15）: 首帧图生 — 720P × 5s（不传 input_usage）
> 执行时间：2026/8/5 10:33:59  |  模型：`vidu-video-q2-pro`  |  标签：standard

> 💡 无 input_usage → i2v-720P 47945×5=239725

### 调用参数
```json
{
  "model": "vidu-video-q2-pro",
  "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
  "duration": 5,
  "size": "720P",
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"
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
| 预期 Quota | **119,863** |
| 预期 USD | $0.2397 |
| 预期 RMB | ¥1.7500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_sHoUb9Cnod1YFh52Q53ZmU1kJVqAjxH8` |

```json
{
  "id": "task_sHoUb9Cnod1YFh52Q53ZmU1kJVqAjxH8",
  "task_id": "task_sHoUb9Cnod1YFh52Q53ZmU1kJVqAjxH8",
  "object": "video",
  "model": "vidu-video-q2-pro",
  "status": "queued",
  "progress": 0,
  "created_at": 1785897147
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 14,272,975 | — |
| 提交后可用 Quota | 14,153,112 | — |
| **预扣金额** | **119,863** | **¥1.7500** |
| 预期扣减 | 119,863 | ¥1.7500 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 119,863 | — |
| request_id | 20260805023227520227000f61X4gQRB76RyDEp | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 7 |
| 完成耗时 | 90.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/0ceebbdf5001834814628220596/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.084s** |
| 计费参考时长 | 5s |
| 时长差值 | 0.084s |
| 输出分辨率 | 960×960 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 119,863 | ¥1.7500 |
| 平台记录最终消费 | 121,876 | **¥1.779390** |
| 差额 | — | +¥0.0294（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 116,
    "created_at": 1785897147,
    "updated_at": 1785897237,
    "task_id": "task_sHoUb9Cnod1YFh52Q53ZmU1kJVqAjxH8",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.779390",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/0ceebbdf5001834814628220596/aigcVideoGenFile.mp4",
    "submit_time": 1785897147,
    "start_time": 1785897156,
    "finish_time": 1785897237,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q2-pro",
      "origin_model_name": "vidu-video-q2-pro"
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
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg",
                "Usage": "FirstFrame",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Vidu",
            "ModelVersion": "q2-pro",
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
            "Prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:33:57Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/0ceebbdf5001834814628220596/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 2605693,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.084,
                  "Height": 960,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 1655918,
                  "VideoDuration": 5.083,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 2600503,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "SDR"
                      },
                      "Fps": 24,
                      "Height": 960,
                      "Width": 960
                    }
                  ],
                  "Width": 960
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
          "TaskId": "1500065731-AigcVideoTask-5c293f1e6ecb9cda6d218720df18471ft"
        },
        "BeginProcessTime": "2026-08-05T02:32:27Z",
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
        "CreateTime": "2026-08-05T02:32:27Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:33:55Z",
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
        "RequestId": "697c8e9d-2dcb-4f0c-88a6-71162b9d79b4",
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

## TC-16（seq=16）: 参考生 — 1080P × 8s（input_usage=Reference）
> 执行时间：2026/8/5 10:35:46  |  模型：`vidu-video-q2-pro`  |  标签：standard

> 💡 ref-1080P 123288×8=986304

### 调用参数
```json
{
  "model": "vidu-video-q2-pro",
  "prompt": "人物向镜头走来，背景为樱花盛开的公园，写实风格",
  "duration": 8,
  "size": "1080P",
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"
  ],
  "metadata": {
    "input_usage": "Reference"
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 8 | **8** | 已传，直接使用 |
| size（分辨率） | 1080P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 1 张 | — | input_usage=Reference |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **493,152** |
| 预期 USD | $0.9863 |
| 预期 RMB | ¥7.2000 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.14s |
| task_id | `task_ruFWpKsYQZFZpk0btflcuD9XTW0CDHaW` |

```json
{
  "id": "task_ruFWpKsYQZFZpk0btflcuD9XTW0CDHaW",
  "task_id": "task_ruFWpKsYQZFZpk0btflcuD9XTW0CDHaW",
  "object": "video",
  "model": "vidu-video-q2-pro",
  "status": "queued",
  "progress": 0,
  "created_at": 1785897239
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 14,150,754 | — |
| 提交后可用 Quota | 13,657,602 | — |
| **预扣金额** | **493,152** | **¥7.2000** |
| 预期扣减 | 493,152 | ¥7.2000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 493,152 | — |
| request_id | 20260805023359266349000f61X4gQRQkoXtRiP | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 8 |
| 完成耗时 | 105.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/6ecf1ffd5001834814626091356/aigcVideoGenFile.mp4

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
| 预扣金额 | 493,152 | ¥7.2000 |
| 平台记录最终消费 | 495,741 | **¥7.237819** |
| 差额 | — | +¥0.0378（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 117,
    "created_at": 1785897239,
    "updated_at": 1785897333,
    "task_id": "task_ruFWpKsYQZFZpk0btflcuD9XTW0CDHaW",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "7.237819",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/6ecf1ffd5001834814626091356/aigcVideoGenFile.mp4",
    "submit_time": 1785897239,
    "start_time": 1785897253,
    "finish_time": 1785897333,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q2-pro",
      "origin_model_name": "vidu-video-q2-pro"
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
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg",
                "Usage": "Reference",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Vidu",
            "ModelVersion": "q2-pro",
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
            "Prompt": "人物向镜头走来，背景为樱花盛开的公园，写实风格",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:35:33Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/6ecf1ffd5001834814626091356/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 10219956,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8.042,
                  "Height": 1080,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 10273611,
                  "VideoDuration": 8.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 10215940,
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
          "TaskId": "1500065731-AigcVideoTask-fb0c5149be744744c13c8d868db044d9t"
        },
        "BeginProcessTime": "2026-08-05T02:33:59Z",
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
        "CreateTime": "2026-08-05T02:33:59Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:35:20Z",
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
        "RequestId": "c6f259fa-dc6f-4808-9fa6-e0b0db203283",
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

## EC-06（seq=53）: 有图但不传 input_usage → 走首帧(i2v)价格
> 执行时间：2026/8/5 11:33:35  |  模型：`vidu-video-q2-pro`  |  标签：edge

> 💡 不传 input_usage → ref=false → i2v-720P 47945×5=239725。与 TC-15 相同预期，验证默认行为。

### 调用参数
```json
{
  "model": "vidu-video-q2-pro",
  "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
  "duration": 5,
  "size": "720P",
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"
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
| 预期 Quota | **119,863** |
| 预期 USD | $0.2397 |
| 预期 RMB | ¥1.7500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_tY4k1tgT3BnRN0yY0pAXqO9oCB5JWjoE` |

```json
{
  "id": "task_tY4k1tgT3BnRN0yY0pAXqO9oCB5JWjoE",
  "task_id": "task_tY4k1tgT3BnRN0yY0pAXqO9oCB5JWjoE",
  "object": "video",
  "model": "vidu-video-q2-pro",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900709
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 119,863 | ¥1.7500 |
| 预扣是否符合 | ❌ 不符 | — |
| 消费日志 Quota | null | — |
| request_id | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 8 |
| 完成耗时 | 105.1s |
| progress | 100% |

**视频 URL**: http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/8df875fc5001834814640927621/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.084s** |
| 计费参考时长 | 5s |
| 时长差值 | - |
| 输出分辨率 | 960×960 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 121,876 | **¥1.779390** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 152,
    "created_at": 1785900709,
    "updated_at": 1785900803,
    "task_id": "task_tY4k1tgT3BnRN0yY0pAXqO9oCB5JWjoE",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.779390",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/8df875fc5001834814640927621/aigcVideoGenFile.mp4",
    "submit_time": 1785900709,
    "start_time": 1785900722,
    "finish_time": 1785900803,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q2-pro",
      "origin_model_name": "vidu-video-q2-pro"
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
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg",
                "Usage": "FirstFrame",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Vidu",
            "ModelVersion": "q2-pro",
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
            "Prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:33:23Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://251000800.vod2.myqcloud.com/b8c85f46vodnj251000800/8df875fc5001834814640927621/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 3195543,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.084,
                  "Height": 960,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 2030768,
                  "VideoDuration": 5.083,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 3190430,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "SDR"
                      },
                      "Fps": 24,
                      "Height": 960,
                      "Width": 960
                    }
                  ],
                  "Width": 960
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
          "TaskId": "1500065731-AigcVideoTask-53dc1e878d135216cdf21c1ccfcb133at"
        },
        "BeginProcessTime": "2026-08-05T03:31:49Z",
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
        "CreateTime": "2026-08-05T03:31:49Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:33:12Z",
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
        "RequestId": "a48305e0-a466-4c96-b23e-9e3392f041a6",
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

## EC-07（seq=54）: input_usage=Reference 但不传 images → 参考生价（表达式仍按 ref 计）
> 执行时间：2026/8/5 11:34:52  |  模型：`vidu-video-q2-pro`  |  标签：edge

> 💡 表达式 ref=true(input_usage==Reference) → ref-720P 47945×5=239725。注意：无图时腾讯 API 实际行为待观察。

### 调用参数
```json
{
  "model": "vidu-video-q2-pro",
  "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
  "duration": 5,
  "size": "720P",
  "metadata": {
    "input_usage": "Reference"
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 5 | **5** | 已传，直接使用 |
| size（分辨率） | 720P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **119,863** |
| 预期 USD | $0.2397 |
| 预期 RMB | ¥1.7500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.15s |
| task_id | `task_VDH92jpWo3DjcUpeFgM4hUbNwLrrseqZ` |

```json
{
  "id": "task_VDH92jpWo3DjcUpeFgM4hUbNwLrrseqZ",
  "task_id": "task_VDH92jpWo3DjcUpeFgM4hUbNwLrrseqZ",
  "object": "video",
  "model": "vidu-video-q2-pro",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900815
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 119,863 | ¥1.7500 |
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

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/ea66b1005001834814648463271/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.042s** |
| 计费参考时长 | 5s |
| 时长差值 | - |
| 输出分辨率 | 1280×720 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 120,869 | **¥1.764687** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 153,
    "created_at": 1785900815,
    "updated_at": 1785900884,
    "task_id": "task_VDH92jpWo3DjcUpeFgM4hUbNwLrrseqZ",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.764687",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/ea66b1005001834814648463271/aigcVideoGenFile.mp4",
    "submit_time": 1785900815,
    "start_time": 1785900819,
    "finish_time": 1785900884,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q2-pro",
      "origin_model_name": "vidu-video-q2-pro"
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
            "ModelName": "Vidu",
            "ModelVersion": "q2-pro",
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
            "Prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:34:44Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/ea66b1005001834814648463271/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 5081697,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 720,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 3202740,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 5076400,
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
          "TaskId": "1500065731-AigcVideoTask-c0b2661995c2b4182b68ea56b9d59b5ft"
        },
        "BeginProcessTime": "2026-08-05T03:33:36Z",
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
        "CreateTime": "2026-08-05T03:33:36Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:34:32Z",
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
        "RequestId": "b93be0b9-6928-4831-bc2a-de64b0a03646",
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

## EC-08（seq=55）: duration=0（→ 按默认 5s 计费）
> 执行时间：2026/8/5 11:36:09  |  模型：`vidu-video-q2-pro`  |  标签：edge

> 💡 d0=0≤0→d=5；i2v-720P 47945×5=239725。验证 duration=0 的 clamp 行为。

### 调用参数
```json
{
  "model": "vidu-video-q2-pro",
  "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
  "duration": 0,
  "size": "720P",
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 0 | **5** | 已传，直接使用 |
| size（分辨率） | 720P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 1 张 | — | 参考图生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **119,863** |
| 预期 USD | $0.2397 |
| 预期 RMB | ¥1.7500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.13s |
| task_id | `task_jiSr7mpXoaRoj6azIcQymmRYThpL7lri` |

```json
{
  "id": "task_jiSr7mpXoaRoj6azIcQymmRYThpL7lri",
  "task_id": "task_jiSr7mpXoaRoj6azIcQymmRYThpL7lri",
  "object": "video",
  "model": "vidu-video-q2-pro",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900892
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 119,863 | ¥1.7500 |
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

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/4b9a305e5001834814724181574/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.042s** |
| 计费参考时长 | 5s |
| 时长差值 | - |
| 输出分辨率 | 960×960 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 120,869 | **¥1.764687** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 154,
    "created_at": 1785900892,
    "updated_at": 1785900964,
    "task_id": "task_jiSr7mpXoaRoj6azIcQymmRYThpL7lri",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.764687",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/4b9a305e5001834814724181574/aigcVideoGenFile.mp4",
    "submit_time": 1785900892,
    "start_time": 1785900900,
    "finish_time": 1785900964,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q2-pro",
      "origin_model_name": "vidu-video-q2-pro"
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
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg",
                "Usage": "FirstFrame",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Vidu",
            "ModelVersion": "q2-pro",
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
            "Prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:36:04Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/4b9a305e5001834814724181574/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 1101856,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.042,
                  "Height": 960,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 694445,
                  "VideoDuration": 5.042,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 1096300,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 960,
                      "Width": 960
                    }
                  ],
                  "Width": 960
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
          "TaskId": "1500065731-AigcVideoTask-552b0c1b3e50390005c91044c2e90f80t"
        },
        "BeginProcessTime": "2026-08-05T03:34:52Z",
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
        "CreateTime": "2026-08-05T03:34:52Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:35:56Z",
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
        "RequestId": "30d88ac8-52c5-41d4-82df-b964a9fe8726",
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

## EC-09（seq=56）: size="2K"（i2v-2K 档）
> 执行时间：2026/8/5 11:37:56  |  模型：`vidu-video-q2-pro`  |  标签：edge

> 💡 i2v-2K 136986×5=684930。验证高分辨率档计费。

### 调用参数
```json
{
  "model": "vidu-video-q2-pro",
  "prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
  "duration": 5,
  "size": "2K",
  "images": [
    "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 5 | **5** | 已传，直接使用 |
| size（分辨率） | 2K | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 1 张 | — | 参考图生视频 |

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
| 耗时 | 0.15s |
| task_id | `task_TRBm7t5FHdliczspybkmTAlxbm53q3LA` |

```json
{
  "id": "task_TRBm7t5FHdliczspybkmTAlxbm53q3LA",
  "task_id": "task_TRBm7t5FHdliczspybkmTAlxbm53q3LA",
  "object": "video",
  "model": "vidu-video-q2-pro",
  "status": "queued",
  "progress": 0,
  "created_at": 1785900969
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 2,804,453 | — |
| 提交后可用 Quota | 2,461,988 | — |
| **预扣金额** | **342,465** | **¥5.0000** |
| 预期扣减 | 342,465 | ¥5.0000 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 342,465 | — |
| request_id | 20260805033609296660000f61X4gQRRo2UsrVe | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 8 |
| 完成耗时 | 105.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/bd1139e15001834814644704485/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.084s** |
| 计费参考时长 | 5s |
| 时长差值 | 0.084s |
| 输出分辨率 | 1440×1440 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 342,465 | ¥5.0000 |
| 平台记录最终消费 | 348,218 | **¥5.083983** |
| 差额 | — | +¥0.0840（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 155,
    "created_at": 1785900969,
    "updated_at": 1785901061,
    "task_id": "task_TRBm7t5FHdliczspybkmTAlxbm53q3LA",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "5.083983",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/bd1139e15001834814644704485/aigcVideoGenFile.mp4",
    "submit_time": 1785900969,
    "start_time": 1785900981,
    "finish_time": 1785901061,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q2-pro",
      "origin_model_name": "vidu-video-q2-pro"
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
                "Url": "https://seekingliren.oss-cn-hangzhou.aliyuncs.com/admin/uploads/1785809772583-jq2pzivz.jpeg",
                "Usage": "FirstFrame",
                "VoiceId": ""
              }
            ],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Vidu",
            "ModelVersion": "q2-pro",
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
              "Resolution": "1080P",
              "StorageMode": "Temporary"
            },
            "Prompt": "人物缓缓转身，露出微笑，阳光洒在发梢",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T03:37:41Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/bd1139e15001834814644704485/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 5777990,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.084,
                  "Height": 1440,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 3671913,
                  "VideoDuration": 5.083,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 5773166,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "SDR"
                      },
                      "Fps": 24,
                      "Height": 1440,
                      "Width": 1440
                    }
                  ],
                  "Width": 1440
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
          "TaskId": "1500065731-AigcVideoTask-0834d4cfb72b5d2a6b0a78711a37c825t"
        },
        "BeginProcessTime": "2026-08-05T03:36:09Z",
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
        "CreateTime": "2026-08-05T03:36:09Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T03:37:30Z",
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
        "RequestId": "cf52fd49-e90c-4d39-838c-60ad66839a86",
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

