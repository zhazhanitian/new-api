# 15. vidu-image-q2 文生图1080p

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:21:20  |  模型：`vidu-image-q2`  |  档位：1080p 文生图

### 调用参数
```json
{
  "model": "vidu-image-q2",
  "prompt": "仙鹤在荷花池边展翅",
  "resolution": "1080p",
  "ratio": "9:16",
  "n": 1
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 1080p | 分辨率档位 |
| ratio | 9:16 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 无 | 参考图 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **12,877** |
| 预期 USD | $0.0258 |
| 预期 RMB | ¥0.188004 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.13s |
| task_id | `task_0fdQuLULYhw1CgmSgEByBEY84e89A2wH` |
| request_id | 202611071021205367530008268d9d6goMhQGgu |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_0fdQuLULYhw1CgmSgEByBEY84e89A2wH",
    "status": "NOT_START",
    "progress": "0%",
    "model": "vidu-image-q2",
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
| 任务首次查询 amount | 12,877 | 0.188004 |
| 消费日志 Quota | 12,877 | 0.188004 |
| **预扣金额** | **12,877** | **¥0.188004** |
| 预期扣减 | 12,877 | ¥0.188004 |
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

**图片 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/3d7b85e25001834827298007508/aigcImageGenFile.png


#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 12,877 | ¥0.188004 |
| 平台记录最终消费 | 12,877 | **¥0.188004** |
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
        "ModelName": "Vidu",
        "ModelVersion": "q2",
        "NegativePrompt": "",
        "OutputConfig": {
          "AspectRatio": "9:16",
          "ClassId": 0,
          "ExpireTime": "0000-00-00T00:00:00Z",
          "InputComplianceCheck": "",
          "LogoAdd": "",
          "MediaName": "",
          "OutputComplianceCheck": "",
          "OutputFormat": "",
          "PersonGeneration": "",
          "Resolution": "1080p",
          "StorageMode": "Temporary"
        },
        "Prompt": "仙鹤在荷花池边展翅",
        "SceneType": ""
      },
      "Message": "",
      "Output": {
        "FileInfos": [
          {
            "ClassId": 0,
            "ExpireTime": "2026-11-14T10:22:22Z",
            "FileId": "",
            "FileType": "",
            "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/3d7b85e25001834827298007508/aigcImageGenFile.png",
            "MediaName": "",
            "MetaData": {
              "AudioDuration": 0,
              "AudioStreamSet": [],
              "Bitrate": 0,
              "Container": "png_pipe",
              "Duration": 0,
              "Height": 1920,
              "Md5": "",
              "Rotate": 0,
              "Size": 2608471,
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
                  "Height": 1920,
                  "Width": 1080
                }
              ],
              "Width": 1080
            },
            "StorageMode": "Temporary"
          }
        ],
        "Usage": {
          "SpecificationUsages": [
            {
              "Category": "aigc_image_viduq2_t2i_1080",
              "Specification": "ViduQ2_T2i_1080P"
            }
          ]
        }
      },
      "Progress": 100,
      "SessionContext": "",
      "SessionId": "",
      "Status": "FINISH",
      "TaskId": "1500065731-AigcImageTask-64a5e0bca2084a1dfb0ad7b7cfc23fbat"
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
    "FinishTime": "2026-11-07T10:21:50Z",
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
    "RequestId": "7789ea0d-ddb2-44af-8ff8-42eba8ce33a8",
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
    "task_id": "task_0fdQuLULYhw1CgmSgEByBEY84e89A2wH",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "vidu-image-q2",
    "images": [
      {
        "url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/3d7b85e25001834827298007508/aigcImageGenFile.png"
      }
    ],
    "amount": "0.188004",
    "usage": {
      "input_tokens": 0,
      "output_tokens": 0,
      "total_tokens": 0
    },
    "created_at": 1794046880,
    "submit_time": 1794046880,
    "finish_time": 1794046942
  }
}
```

### 消费日志
```json
[
  {
    "id": 10,
    "user_id": 1,
    "created_at": 1794046880,
    "type": 2,
    "content": "操作 imageGenerate",
    "username": "admin",
    "token_name": "开发部署模型测试",
    "model_name": "vidu-image-q2",
    "quota": 12877,
    "prompt_tokens": 0,
    "completion_tokens": 0,
    "use_time": 0,
    "is_stream": false,
    "channel": 13,
    "channel_name": "",
    "token_id": 1,
    "group": "default",
    "ip": "",
    "request_id": "202611071021205367530008268d9d6goMhQGgu",
    "other": "{\"group_ratio\":1,\"is_task\":true,\"model_price\":0,\"request_path\":\"/v2/image-tasks\"}"
  }
]
```
