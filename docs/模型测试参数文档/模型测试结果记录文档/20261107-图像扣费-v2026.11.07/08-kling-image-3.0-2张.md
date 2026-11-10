# 8. kling-image-3.0 2张

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:21:20  |  模型：`kling-image-3.0`  |  档位：1k x 2

### 调用参数
```json
{
  "model": "kling-image-3.0",
  "prompt": "一杯咖啡的顶视角海报",
  "resolution": "1k",
  "ratio": "1:1",
  "n": 2
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 1k | 分辨率档位 |
| ratio | 1:1 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 2 | 张数 |
| image | 无 | 参考图 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **27,397** |
| 预期 USD | $0.0548 |
| 预期 RMB | ¥0.399996 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.31s |
| task_id | `task_88N7Zvlxq25MrlMhszmkX4faNmKxfAqf` |
| request_id | 202611071021202231360008268d9d6Jtpey9A9 |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_88N7Zvlxq25MrlMhszmkX4faNmKxfAqf",
    "status": "NOT_START",
    "progress": "0%",
    "model": "kling-image-3.0",
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
| 任务首次查询 amount | 27,397 | 0.399996 |
| 消费日志 Quota | - | - |
| **预扣金额** | **27,397** | **¥0.399996** |
| 预期扣减 | 27,397 | ¥0.399996 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 17 |
| 完成耗时 | 128.3s |
| progress | 100% |
| 输出张数 | 2 |

**图片 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/07c73a5b5001834827300546635/aigcImageGenFile.png  
https://store.vod-qcloud.com/b8c85f46vodnj251000800/07c73a5b5001834827300546635/aigcImageGenFile_1.png


#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 27,397 | ¥0.399996 |
| 平台记录最终消费 | 27,397 | **¥0.399996** |
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
        "ModelName": "Kling",
        "ModelVersion": "3.0",
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
          "OutputImageCount": 2,
          "PersonGeneration": "",
          "Resolution": "1k",
          "StorageMode": "Temporary"
        },
        "Prompt": "一杯咖啡的顶视角海报",
        "SceneType": ""
      },
      "Message": "",
      "Output": {
        "FileInfos": [
          {
            "ClassId": 0,
            "ExpireTime": "2026-11-14T10:23:28Z",
            "FileId": "",
            "FileType": "",
            "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/07c73a5b5001834827300546635/aigcImageGenFile.png",
            "MediaName": "",
            "MetaData": {
              "AudioDuration": 0,
              "AudioStreamSet": [],
              "Bitrate": 0,
              "Container": "png_pipe",
              "Duration": 0,
              "Height": 1024,
              "Md5": "",
              "Rotate": 0,
              "Size": 1236042,
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
                  "Height": 1024,
                  "Width": 1024
                }
              ],
              "Width": 1024
            },
            "StorageMode": "Temporary"
          },
          {
            "ClassId": 0,
            "ExpireTime": "2026-11-14T10:23:28Z",
            "FileId": "",
            "FileType": "",
            "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/07c73a5b5001834827300546635/aigcImageGenFile_1.png",
            "MediaName": "",
            "MetaData": {
              "AudioDuration": 0,
              "AudioStreamSet": [],
              "Bitrate": 0,
              "Container": "png_pipe",
              "Duration": 0,
              "Height": 1024,
              "Md5": "",
              "Rotate": 0,
              "Size": 1113547,
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
                  "Height": 1024,
                  "Width": 1024
                }
              ],
              "Width": 1024
            },
            "StorageMode": "Temporary"
          }
        ],
        "Usage": {
          "SpecificationUsages": [
            {
              "Category": "aigc_image_kl3_1k2k",
              "Specification": "Kling3.0_1K2K"
            },
            {
              "Category": "aigc_image_kl3_1k2k",
              "Specification": "Kling3.0_1K2K"
            }
          ]
        }
      },
      "Progress": 100,
      "SessionContext": "",
      "SessionId": "",
      "Status": "FINISH",
      "TaskId": "1500065731-AigcImageTask-7f316f1450a90c99cd8fc76616e9ccect"
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
    "FinishTime": "2026-11-07T10:23:08Z",
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
    "RequestId": "bdbe9431-50a3-45c2-94dd-3b5a5c24ddfb",
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
    "task_id": "task_88N7Zvlxq25MrlMhszmkX4faNmKxfAqf",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "kling-image-3.0",
    "images": [
      {
        "url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/07c73a5b5001834827300546635/aigcImageGenFile.png"
      },
      {
        "url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/07c73a5b5001834827300546635/aigcImageGenFile_1.png"
      }
    ],
    "amount": "0.399996",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794046880,
    "submit_time": 1794046880,
    "finish_time": 1794047008
  }
}
```

### 消费日志
```json
[]
```
