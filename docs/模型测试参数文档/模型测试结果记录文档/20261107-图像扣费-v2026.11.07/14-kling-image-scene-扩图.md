# 14. kling-image-scene 扩图

> 本文件由 `image_billing_batch.js` 生成。价格文档 `v2026.11.07`。

> 执行时间：2026/11/7 18:21:20  |  模型：`kling-image-scene`  |  档位：固定扩图

### 调用参数
```json
{
  "model": "kling-image-scene",
  "prompt": "以原图风格自然延展画面",
  "n": 1,
  "image": [
    "https://fastly.picsum.photos/id/40/512/512.jpg?hmac=PIXIOTze0IyV8x16iptEpCgTW8UtUw1URQ7dp4YpXWk"
  ],
  "extra_fields": {
    "scene_type": "image_expand",
    "expansion": {
      "up": 0.2,
      "down": 0.2,
      "left": 0.4,
      "right": 0.4
    }
  }
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| resolution | 未传 | 分辨率档位 |
| ratio | 未传 | 宽高比 |
| size | 未传 | 宽x高 |
| quality | 未传 | 渲染精度 |
| n | 1 | 张数 |
| image | 1 张 | 参考图 |

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
| task_id | `task_9YVjHxeCs0ybno3Bc47xDxY2U90CJPD5` |
| request_id | 202611071021205356020008268d9d6VKdDD1w3 |

```json
{
  "code": "success",
  "message": "",
  "data": {
    "task_id": "task_9YVjHxeCs0ybno3Bc47xDxY2U90CJPD5",
    "status": "NOT_START",
    "progress": "0%",
    "model": "kling-image-scene",
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
| 消费日志 Quota | - | - |
| **预扣金额** | **13,699** | **¥0.200005** |
| 预期扣减 | 13,699 | ¥0.200005 |
| 预扣是否符合 | ✅ 符合 | — |
| 分组倍率 | - | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 8 |
| 完成耗时 | 56.1s |
| progress | 100% |
| 输出张数 | 1 |

**图片 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/8d00b5e95001834827299136862/aigcImageGenFile.png


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
        "ExtInfo": "{\"AdditionalParameters\":\"{\\\"down_expansion_ratio\\\":0.2,\\\"left_expansion_ratio\\\":0.4,\\\"right_expansion_ratio\\\":0.4,\\\"up_expansion_ratio\\\":0.2}\"}",
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
        "ModelName": "Kling",
        "ModelVersion": "scene",
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
        "Prompt": "以原图风格自然延展画面",
        "SceneType": "image_expand"
      },
      "Message": "",
      "Output": {
        "FileInfos": [
          {
            "ClassId": 0,
            "ExpireTime": "2026-11-14T10:23:10Z",
            "FileId": "",
            "FileType": "",
            "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/8d00b5e95001834827299136862/aigcImageGenFile.png",
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
              "Size": 1529258,
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
                  "Width": 1344
                }
              ],
              "Width": 1344
            },
            "StorageMode": "Temporary"
          }
        ],
        "Usage": {
          "SpecificationUsages": [
            {
              "Category": "aigc_image_kl_expand",
              "Specification": "Kling_Expand"
            }
          ]
        }
      },
      "Progress": 100,
      "SessionContext": "",
      "SessionId": "",
      "Status": "FINISH",
      "TaskId": "1500065731-AigcImageTask-c10e7929fd74bdf275eb01704f694666t"
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
    "FinishTime": "2026-11-07T10:23:00Z",
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
    "RequestId": "f0648270-8cb6-4588-83b7-e254fd363069",
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
    "task_id": "task_9YVjHxeCs0ybno3Bc47xDxY2U90CJPD5",
    "status": "SUCCESS",
    "progress": "100%",
    "model": "kling-image-scene",
    "images": [
      {
        "url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/8d00b5e95001834827299136862/aigcImageGenFile.png"
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
    "finish_time": 1794046990
  }
}
```

### 消费日志
```json
[]
```
