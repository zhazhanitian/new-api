# 19. hunyuan-image-3.0 固定价

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:21:20  |  模型：`hunyuan-image-3.0`  |  档位：固定 0.200

### 调用参数
```json
{
  "model": "hunyuan-image-3.0",
  "prompt": "元宵节灯笼",
  "size": "1024x1024",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 未传 | 分辨率档位 |
| ratio | 未传 | 宽高比 |
| size | 1024x1024 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 无 | 参考图 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **13,699** |
| 预期 USD | $0.0274 |
| 预期 RMB | ¥0.200005 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.11s |
| task_id | `task_6nwixBShAj92SNBV4JpcCTfLnmtghbLd` |
| request_id | 202611071021206383840008268d9d6MMB7kX3s |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_6nwixBShAj92SNBV4JpcCTfLnmtghbLd",
    "status": "NOT_START",
    "progress": "0%",
    "model": "hunyuan-image-3.0",
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
| 任务首次查询 amount | 13,699 | 0.200005 |
| 消费日志 Quota | 13,699 | 0.200005 |
| **预扣金额** | **13,699** | **¥0.200005** |
| 预期扣减 | 13,699 | ¥0.200005 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | 1 | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 2 |
| 完成耗时 | 8.0s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/4945bd2a5001834827298520966/aigcImageGenFile.png


#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 13,699 | ¥0.200005 |
| 平台记录最终消费 | 13,699 | **¥0.200005** |
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
        "ExtInfo": "{\"AdditionalParameters\":\"{\\\"size\\\":\\\"1024x1024\\\"}\"}",
        "FileInfos": [],
        "GenerationMode": "",
        "ModelName": "Hunyuan",
        "ModelVersion": "3.0",
        "NegativePrompt": "",
        "OutputConfig": {
          "AspectRatio": "",
          "ClassId": 0,
          "ExpireTime": "0000-00-00T00:00:00Z",
          "InputComplianceCheck": "",
          "LogoAdd": "",
          "MediaName": "",
          "OutputComplianceCheck": "",
          "OutputFormat": "",
          "PersonGeneration": "",
          "Resolution": "",
          "StorageMode": "Temporary"
        },
        "Prompt": "元宵节灯笼",
        "SceneType": ""
      },
      "Message": "",
      "Output": {
        "FileInfos": [
          {
            "ClassId": 0,
            "ExpireTime": "2026-11-14T10:22:24Z",
            "FileId": "",
            "FileType": "",
            "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/4945bd2a5001834827298520966/aigcImageGenFile.png",
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
              "Size": 3153662,
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
              "Category": "aigc_image_hy3_1k",
              "Specification": "Hunyuan3.0_1K"
            }
          ]
        }
      },
      "Progress": 100,
      "SessionContext": "",
      "SessionId": "",
      "Status": "FINISH",
      "TaskId": "1500065731-AigcImageTask-38d51b68b1654c476d42d7eb93f84869t"
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
    "FinishTime": "2026-11-07T10:21:46Z",
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
    "RequestId": "385d5c74-b96c-4fb1-946a-7181a81e723f",
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
    "task_id": "task_6nwixBShAj92SNBV4JpcCTfLnmtghbLd",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "hunyuan-image-3.0",
    "images": [
      {
        "url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/4945bd2a5001834827298520966/aigcImageGenFile.png"
      }
    ],
    "amount": "0.200005",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794046880,
    "submit_time": 1794046880,
    "finish_time": 1794046944
  }
}
```

### 消费日志
```json
[
  {
    "id": 5,
    "user_id": 1,
    "created_at": 1794046880,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "hunyuan-image-3.0",
    "quota": 13699,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 13,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611071021206383840008268d9d6MMB7kX3s",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
