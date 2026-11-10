# 6. gemini-3.1-flash-image 2K

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:21:20  |  模型：`gemini-3.1-flash-image`  |  档位：2K

### 调用参数
```json
{
  "model": "gemini-3.1-flash-image",
  "prompt": "未来城市夜景",
  "resolution": "2K",
  "ratio": "1:1",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 2K | 分辨率档位 |
| ratio | 1:1 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 无 | 参考图 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **51,370** |
| 预期 USD | $0.1027 |
| 预期 RMB | ¥0.750002 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.32s |
| task_id | `task_FaCm1rOAAnDvlWqdjXPAqeaJRLZSFVGI` |
| request_id | 202611071021202229650008268d9d6p3mYyAcN |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_FaCm1rOAAnDvlWqdjXPAqeaJRLZSFVGI",
    "status": "NOT_START",
    "progress": "0%",
    "model": "gemini-3.1-flash-image",
    "images": [],
    "amount": "",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 0,
    "submit_time": 0,
    "finish_time": 0
  }
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 任务首次查询 amount | 51,370 | 0.750002 |
| 消费日志 Quota | 51,370 | 0.750002 |
| **预扣金额** | **51,370** | **¥0.750002** |
| 预期扣减 | 51,370 | ¥0.750002 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 8 |
| 完成耗时 | 56.2s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/52e5ea2c5001834827298966180/aigcImageGenFile.jpg


#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 51,370 | ¥0.750002 |
| 平台记录最终消费 | 51,370 | **¥0.750002** |
| 差额 | — | ¥0.000000（无差额） |
| 最终扣费是否符合 | ✅ 最终扣费符合预期 | — |

#### 原始上游响应
```json
{
  "Response": {
    "AigcAudioTask": null,
    "AigcHunyuan3DTask": null,
    "AigcImageTask": {
      "ErrCode": 0,
      "ErrCodeExt": "",
      "Input": {
        "EnhancePrompt": "",
        "FileInfos": [],
        "GenerationMode": "",
        "ModelName": "GG",
        "ModelVersion": "3.1",
        "NegativePrompt": "",
        "OutputConfig": {
          "AspectRatio": "1:1",
          "ClassId": 0,
          "ExpireTime": "0000-00-00T00:00:00Z",
          "InputComplianceCheck": "",
          "LogoAdd": "",
          "MediaName": "",
          "OutputComplianceCheck": "",
          "OutputFormat": "",
          "PersonGeneration": "",
          "Resolution": "2K",
          "StorageMode": "Temporary"
        },
        "Prompt": "未来城市夜景",
        "SceneType": ""
      },
      "Message": "",
      "Output": {
        "FileInfos": [
          {
            "ClassId": 0,
            "ExpireTime": "2026-11-14T10:22:10Z",
            "FileId": "",
            "FileType": "",
            "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/52e5ea2c5001834827298966180/aigcImageGenFile.jpg",
            "MediaName": "",
            "MetaData": {
              "AudioDuration": 0,
              "AudioStreamSet": [],
              "Bitrate": 842797000,
              "Container": "image2",
              "Duration": 0.04,
              "Height": 2048,
              "Md5": "",
              "Rotate": 0,
              "Size": 4213985,
              "VideoDuration": 0.04,
              "VideoStreamSet": [
                {
                  "Bitrate": 842797000,
                  "Codec": "mjpeg",
                  "CodecTag": "",
                  "Codecs": "",
                  "DynamicRangeInfo": {
                    "HDRType": "",
                    "Type": "Unknown"
                  },
                  "Fps": 25,
                  "Height": 2048,
                  "Width": 2048
                }
              ],
              "Width": 2048
            },
            "StorageMode": "Temporary"
          }
        ],
        "Usage": {
          "SpecificationUsages": [
            {
              "Category": "aigc_image_gg31_2k",
              "Specification": "Gem3.1_2K"
            }
          ]
        }
      },
      "Progress": 100,
      "SessionContext": "",
      "SessionId": "",
      "Status": "FINISH",
      "TaskId": "1500065731-AigcImageTask-20893a07bc9a6e4c89ca0ed549f369c4t"
    },
    "AigcVideoRedrawTask": null,
    "AigcVideoTask": null,
    "BeginProcessTime": "2026-11-07T10:21:20Z",
    "ClipTask": null,
    "CloneVoiceAsyncTask": null,
    "ComplexAdaptiveDynamicStreamingTask": null,
    "ComposeMediaTask": null,
    "ConcatTask": null,
    "CreateAigcAdvancedCustomElementTask": null,
    "CreateAigcAudioCloneTask": null,
    "CreateAigcCustomVoiceTask": null,
    "CreateAigcMaterialTask": null,
    "CreateAigcSubjectTask": null,
    "CreateImageSpriteTask": null,
    "CreateTime": "2026-11-07T10:21:20Z",
    "DNATask": null,
    "DescribeAigcFaceInfoAsyncTask": null,
    "DescribeFileAttributesTask": null,
    "DesignVoiceAsyncTask": null,
    "EditMediaTask": null,
    "ExtractBlindWatermarkTask": null,
    "ExtractCopyRightWatermarkTask": null,
    "ExtractTraceWatermarkTask": null,
    "FastClipMediaTask": null,
    "FinishTime": "2026-11-07T10:21:51Z",
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
    "RequestId": "b6404d28-a63c-438c-9d69-521ce82fe4bf",
    "ReviewAudioVideoTask": null,
    "SceneAigcImageTask": null,
    "SceneAigcVideoTask": null,
    "SearchMediaBySemanticsAsync": null,
    "SnapshotByTimeOffsetTask": null,
    "SplitMediaTask": null,
    "Status": "FINISH",
    "TaskType": "AigcImageTask",
    "TextToSpeechAsyncTask": null,
    "TranscodeTask": null,
    "VideoDepthEstimationTask": null,
    "VideoDubbingAsyncTask": null,
    "WandAsrTask": null,
    "WechatMiniProgramPublishTask": null,
    "WechatPublishTask": null
  }
}
```
#### 网关查询响应
```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_FaCm1rOAAnDvlWqdjXPAqeaJRLZSFVGI",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "gemini-3.1-flash-image",
    "images": [
      {
        "url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/52e5ea2c5001834827298966180/aigcImageGenFile.jpg"
      }
    ],
    "amount": "0.750002",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794046880,
    "submit_time": 1794046880,
    "finish_time": 1794046930
  }
}
```

### 消费日志
```json
[
  {
    "id": 25,
    "user_id": 1,
    "created_at": 1794046880,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "gemini-3.1-flash-image",
    "quota": 51370,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 13,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611071021202229650008268d9d6p3mYyAcN",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
