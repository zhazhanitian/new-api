# vidu-video-q2-turbo 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-17（seq=17）: 720P × 5s
> 执行时间：2026/8/5 10:36:47  |  模型：`vidu-video-q2-turbo`  |  标签：standard

> 💡 720P 34247×5=171235

### 调用参数
```json
{
  "model": "vidu-video-q2-turbo",
  "prompt": "快节奏城市街景，行人穿梭，霓虹倒影，延时摄影感",
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
| 预期 Quota | **85,618** |
| 预期 USD | $0.1712 |
| 预期 RMB | ¥1.2500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.16s |
| task_id | `task_RmYYDu2Q1YsOugTz6k8yHCw0EGaHOC5Q` |

```json
{
  "id": "task_RmYYDu2Q1YsOugTz6k8yHCw0EGaHOC5Q",
  "task_id": "task_RmYYDu2Q1YsOugTz6k8yHCw0EGaHOC5Q",
  "object": "video",
  "model": "vidu-video-q2-turbo",
  "status": "queued",
  "progress": 0,
  "created_at": 1785897346
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 13,655,589 | — |
| 提交后可用 Quota | 13,569,971 | — |
| **预扣金额** | **85,618** | **¥1.2500** |
| 预期扣减 | 85,618 | ¥1.2500 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 85,618 | — |
| request_id | 2026080502354618491000f61X4gQRtoZ7b9um | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 5 |
| 完成耗时 | 60.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/0e06dddf5001834814635085652/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **5.084s** |
| 计费参考时长 | 5s |
| 时长差值 | 0.084s |
| 输出分辨率 | 720×1282 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 85,618 | ¥1.2500 |
| 平台记录最终消费 | 87,056 | **¥1.271018** |
| 差额 | — | +¥0.0210（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 118,
    "created_at": 1785897346,
    "updated_at": 1785897398,
    "task_id": "task_RmYYDu2Q1YsOugTz6k8yHCw0EGaHOC5Q",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.271018",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/0e06dddf5001834814635085652/aigcVideoGenFile.mp4",
    "submit_time": 1785897346,
    "start_time": 1785897349,
    "finish_time": 1785897398,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q2-turbo",
      "origin_model_name": "vidu-video-q2-turbo"
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
            "ModelVersion": "q2-turbo",
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
            "Prompt": "快节奏城市街景，行人穿梭，霓虹倒影，延时摄影感",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:36:38Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/0e06dddf5001834814635085652/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 7991743,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.084,
                  "Height": 1282,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 5078753,
                  "VideoDuration": 5.083,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 7987229,
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
          "TaskId": "1500065731-AigcVideoTask-b6d53c38339996b3f75405764d66fe34t"
        },
        "BeginProcessTime": "2026-08-05T02:35:46Z",
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
        "CreateTime": "2026-08-05T02:35:46Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:36:34Z",
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
        "RequestId": "350cc2e3-4a5a-42b6-96a0-fa36673cfcee",
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

## TC-18（seq=18）: 1080P × 8s
> 执行时间：2026/8/5 10:38:04  |  模型：`vidu-video-q2-turbo`  |  标签：standard

> 💡 1080P 64384×8=515072

### 调用参数
```json
{
  "model": "vidu-video-q2-turbo",
  "prompt": "快节奏城市街景，行人穿梭，霓虹倒影，延时摄影感",
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
| 预期 Quota | **257,536** |
| 预期 USD | $0.5151 |
| 预期 RMB | ¥3.7600 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.13s |
| task_id | `task_3bECINg7EltWkWvLAS7VzDbyQkA8Nx7D` |

```json
{
  "id": "task_3bECINg7EltWkWvLAS7VzDbyQkA8Nx7D",
  "task_id": "task_3bECINg7EltWkWvLAS7VzDbyQkA8Nx7D",
  "object": "video",
  "model": "vidu-video-q2-turbo",
  "status": "queued",
  "progress": 0,
  "created_at": 1785897407
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 13,567,382 | — |
| 提交后可用 Quota | 13,309,846 | — |
| **预扣金额** | **257,536** | **¥3.7600** |
| 预期扣减 | 257,536 | ¥3.7600 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 257,536 | — |
| request_id | 20260805023647763854000f61X4gQRrvFBdBAk | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 6 |
| 完成耗时 | 75.1s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/ddd6fc1a5001834814624445280/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8.084s** |
| 计费参考时长 | 8s |
| 时长差值 | 0.084s |
| 输出分辨率 | 1080×1922 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 257,536 | ¥3.7600 |
| 平台记录最终消费 | 260,240 | **¥3.799504** |
| 差额 | — | +¥0.0395（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 119,
    "created_at": 1785897407,
    "updated_at": 1785897479,
    "task_id": "task_3bECINg7EltWkWvLAS7VzDbyQkA8Nx7D",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.799504",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/ddd6fc1a5001834814624445280/aigcVideoGenFile.mp4",
    "submit_time": 1785897407,
    "start_time": 1785897414,
    "finish_time": 1785897479,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q2-turbo",
      "origin_model_name": "vidu-video-q2-turbo"
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
            "ModelVersion": "q2-turbo",
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
            "Prompt": "快节奏城市街景，行人穿梭，霓虹倒影，延时摄影感",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:37:59Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/ddd6fc1a5001834814624445280/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 12369804,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8.084,
                  "Height": 1922,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 12499687,
                  "VideoDuration": 8.083,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 12366382,
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
          "TaskId": "1500065731-AigcVideoTask-e4542007380060116181c4627205f136t"
        },
        "BeginProcessTime": "2026-08-05T02:36:47Z",
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
        "CreateTime": "2026-08-05T02:36:47Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:37:50Z",
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
        "RequestId": "d61116f0-0bca-4ea7-a2ef-ca4738b6b02a",
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

