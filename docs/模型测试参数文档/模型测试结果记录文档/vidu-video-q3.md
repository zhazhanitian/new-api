# vidu-video-q3 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-19（seq=19）: 720P × 5s
> 执行时间：2026/8/5 10:38:51  |  模型：`vidu-video-q3`  |  标签：standard

> 💡 720P 85616×5=428080

### 调用参数
```json
{
  "model": "vidu-video-q3",
  "prompt": "冬日雪原，驯鹿群奔跑，雪花飞舞，北极光远闪",
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
| 预期 Quota | **214,040** |
| 预期 USD | $0.4281 |
| 预期 RMB | ¥3.1250 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.16s |
| task_id | `task_KOJAxP3jyKCL85bzaTMW733rrCbfNVEk` |

```json
{
  "id": "task_KOJAxP3jyKCL85bzaTMW733rrCbfNVEk",
  "task_id": "task_KOJAxP3jyKCL85bzaTMW733rrCbfNVEk",
  "object": "video",
  "model": "vidu-video-q3",
  "status": "queued",
  "progress": 0,
  "created_at": 1785897484
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 13,308,408 | — |
| 提交后可用 Quota | 13,094,368 | — |
| **预扣金额** | **214,040** | **¥3.1250** |
| 预期扣减 | 214,040 | ¥3.1250 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 214,040 | — |
| request_id | 20260805023804477841000f61X4gQRmV0grYQ9 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 4 |
| 完成耗时 | 45.0s |
| progress | 100% |

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/0e0876815001834814635088372/aigcVideoGenFile.mp4

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
| 预扣金额 | 214,040 | ¥3.1250 |
| 平台记录最终消费 | 217,635 | **¥3.177471** |
| 差额 | — | +¥0.0525（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 120,
    "created_at": 1785897484,
    "updated_at": 1785897527,
    "task_id": "task_KOJAxP3jyKCL85bzaTMW733rrCbfNVEk",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.177471",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/0e0876815001834814635088372/aigcVideoGenFile.mp4",
    "submit_time": 1785897484,
    "start_time": 1785897495,
    "finish_time": 1785897527,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q3",
      "origin_model_name": "vidu-video-q3"
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
            "ModelVersion": "q3",
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
            "Prompt": "冬日雪原，驯鹿群奔跑，雪花飞舞，北极光远闪",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:38:47Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/0e0876815001834814635088372/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 5399301,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.084,
                  "Height": 1282,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 3431256,
                  "VideoDuration": 5.083,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 5394425,
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
          "TaskId": "1500065731-AigcVideoTask-d4579869e3a05a12a77c7cc5fc7dbc89t"
        },
        "BeginProcessTime": "2026-08-05T02:38:04Z",
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
        "CreateTime": "2026-08-05T02:38:04Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:38:41Z",
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
        "RequestId": "564aae44-3f04-4330-856f-b1ecf0dd085b",
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

## TC-20（seq=20）: 1080P × 8s
> 执行时间：2026/8/5 10:40:22  |  模型：`vidu-video-q3`  |  标签：standard

> 💡 1080P 107123×8=856984

### 调用参数
```json
{
  "model": "vidu-video-q3",
  "prompt": "冬日雪原，驯鹿群奔跑，雪花飞舞，北极光远闪",
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
| 预期 Quota | **428,492** |
| 预期 USD | $0.8570 |
| 预期 RMB | ¥6.2560 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.14s |
| task_id | `task_SB1Rm860TR1rLLf36ydLZxoXQVN0imFr` |

```json
{
  "id": "task_SB1Rm860TR1rLLf36ydLZxoXQVN0imFr",
  "task_id": "task_SB1Rm860TR1rLLf36ydLZxoXQVN0imFr",
  "object": "video",
  "model": "vidu-video-q3",
  "status": "queued",
  "progress": 0,
  "created_at": 1785897531
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | null | — |
| 提交后可用 Quota | null | — |
| **预扣金额** | **null** | **-** |
| 预期扣减 | 428,492 | ¥6.2560 |
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

**视频 URL**: http://store.vod-qcloud.com/b8c85f46vodnj251000800/7e0df5495001834814619931939/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | **8.084s** |
| 计费参考时长 | 8s |
| 时长差值 | - |
| 输出分辨率 | 1080×1922 |

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | null | - |
| 平台记录最终消费 | 432,991 | **¥6.321669** |
| 差额 | — | - |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 121,
    "created_at": 1785897531,
    "updated_at": 1785897608,
    "task_id": "task_SB1Rm860TR1rLLf36ydLZxoXQVN0imFr",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "6.321669",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/7e0df5495001834814619931939/aigcVideoGenFile.mp4",
    "submit_time": 1785897531,
    "start_time": 1785897543,
    "finish_time": 1785897608,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "vidu-video-q3",
      "origin_model_name": "vidu-video-q3"
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
            "ModelVersion": "q3",
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
            "Prompt": "冬日雪原，驯鹿群奔跑，雪花飞舞，北极光远闪",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-08-12T02:40:08Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "http://store.vod-qcloud.com/b8c85f46vodnj251000800/7e0df5495001834814619931939/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 12625954,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 8.084,
                  "Height": 1922,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 12758527,
                  "VideoDuration": 8.083,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 12622571,
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
          "TaskId": "1500065731-AigcVideoTask-7d409fcd8b0f238acb523ff14d92ac96t"
        },
        "BeginProcessTime": "2026-08-05T02:38:51Z",
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
        "CreateTime": "2026-08-05T02:38:51Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:39:55Z",
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
        "RequestId": "7982d37d-b1e5-42a1-a6ac-58a11535aa2f",
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

