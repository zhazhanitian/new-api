# 25. gpt-image-2 medium-2K-参考图

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:21:20  |  模型：`gpt-image-2`  |  档位：medium 2K + 1 张参考图

### 调用参数
```json
{
  "model": "gpt-image-2",
  "prompt": "按参考图风格画一只橘猫",
  "resolution": "2K",
  "ratio": "1:1",
  "quality": "medium",
  "n": 1,
  "image": [
    "https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk"
  ]
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 2K | 分辨率档位 |
| ratio | 1:1 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | medium | 渲染精度 |
| n | 1 | 张数 |
| image | 1 张 | 参考图 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **61,836** |
| 预期 USD | $0.1237 |
| 预期 RMB | ¥0.902806 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.12s |
| task_id | `task_jJMxh20R7TWSppQTUHhfTmIPbegEzrxt` |
| request_id | 202611071021206676530008268d9d6OKdNG2QW |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_jJMxh20R7TWSppQTUHhfTmIPbegEzrxt",
    "status": "NOT_START",
    "progress": "0%",
    "model": "gpt-image-2",
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
| 任务首次查询 amount | 61,836 | 0.902806 |
| 消费日志 Quota | 61,836 | 0.902806 |
| **预扣金额** | **61,836** | **¥0.902806** |
| 预期扣减 | 61,836 | ¥0.902806 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 5 |
| 完成耗时 | 32.0s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/9d8b8a315001834827299867148/aigcImageGenFile.png


#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 61,836 | ¥0.902806 |
| 平台记录最终消费 | 61,836 | **¥0.902806** |
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
        "FileInfos": [
          {
            "FileId": "",
            "ReferenceType": "",
            "Text": "",
            "Type": "Url",
            "Url": "https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk"
          }
        ],
        "GenerationMode": "",
        "ModelName": "OG",
        "ModelVersion": "image2_medium",
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
        "Prompt": "按参考图风格画一只橘猫",
        "SceneType": ""
      },
      "Message": "",
      "Output": {
        "FileInfos": [
          {
            "ClassId": 0,
            "ExpireTime": "2026-11-14T10:22:50Z",
            "FileId": "",
            "FileType": "",
            "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/9d8b8a315001834827299867148/aigcImageGenFile.png",
            "MediaName": "",
            "MetaData": {
              "AudioDuration": 0,
              "AudioStreamSet": [],
              "Bitrate": 0,
              "Container": "png_pipe",
              "Duration": 0,
              "Height": 2048,
              "Md5": "",
              "Rotate": 0,
              "Size": 5955939,
              "VideoDuration": 0,
              "VideoStreamSet": [
                {
                  "Bitrate": 0,
                  "Codec": "png",
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
              "Category": "aigc_image_ogimage2_medium_2k",
              "Specification": "OGimage2_Medium_2K"
            },
            {
              "Category": "aigc_image_ogimage2_inputimage",
              "Specification": "OGimage2_inputimage"
            }
          ]
        }
      },
      "Progress": 100,
      "SessionContext": "",
      "SessionId": "",
      "Status": "FINISH",
      "TaskId": "1500065731-AigcImageTask-eaf77364ebc610a323fc40f5a75f49d9t"
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
    "FinishTime": "2026-11-07T10:22:43Z",
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
    "RequestId": "00f32b52-9a3c-4fac-92f9-429cf3bb1819",
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
    "task_id": "task_jJMxh20R7TWSppQTUHhfTmIPbegEzrxt",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "gpt-image-2",
    "images": [
      {
        "url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/9d8b8a315001834827299867148/aigcImageGenFile.png"
      }
    ],
    "amount": "0.902806",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794046880,
    "submit_time": 1794046880,
    "finish_time": 1794046970
  }
}
```

### 消费日志
```json
[
  {
    "id": 2,
    "user_id": 1,
    "created_at": 1794046880,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "gpt-image-2",
    "quota": 61836,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 13,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611071021206676530008268d9d6OKdNG2QW",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
