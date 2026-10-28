# hailuo-video-2.3-fast 测试报告

## 30. hailuo-video-2.3-fast 默认
> 执行时间：2026-10-28  |  模型：`hailuo-video-2.3-fast`

> 表达式值 184932，quota 92466

### 调用参数
```json
{
  "model": "hailuo-video-2.3-fast",
  "prompt": "桌面上的咖啡杯，热气缓缓升起"
}
```

### 价格变量核对
| 变量 | 请求值 | 说明 |
|---|---|---|
| duration | 未传 | 表达式按请求秒数预扣，缺省走模型默认 |
| resolution | 未传 | 外层分辨率 |
| ratio | 未传 | 不参与本轮计价 |
| audio_generation | 未传 | 仅 Enabled 走有声价 |
| generate_audio | 未传 | 豆包 1.5 显式 false 才无声 |
| content | 无 | 空 role 的图按参考图，视频按参考视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **92,466** |
| 预期 RMB | ¥1.3500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_bwsBmEmBhg01APZVvnAGJk0DoduBv0DD` |

```json
{
  "id": "task_bwsBmEmBhg01APZVvnAGJk0DoduBv0DD",
  "task_id": "task_bwsBmEmBhg01APZVvnAGJk0DoduBv0DD",
  "object": "video",
  "model": "hailuo-video-2.3-fast",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172930
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 92,466 | ¥1.3500 |
| 日志预扣 | - | - |
| 平台记录最终消费 | 92,466 | ¥1.3500 |
| 是否符合 | ✅ 预扣符合，任务失败后退款 | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **FAILURE** |
| 轮询次数 | 3 |
| progress | 100% |
| 失败原因 | task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode |

**视频 URL**: task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | - |
| 输出分辨率 | - |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_bwsBmEmBhg01APZVvnAGJk0DoduBv0DD`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 254,
    "created_at": 1793172930,
    "updated_at": 1793173032,
    "task_id": "task_bwsBmEmBhg01APZVvnAGJk0DoduBv0DD",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.350004",
    "action": "generate",
    "status": "FAILURE",
    "fail_reason": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
    "result_url": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
    "submit_time": 1793172930,
    "start_time": 1793172972,
    "finish_time": 1793173032,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-2.3-fast",
      "origin_model_name": "hailuo-video-2.3-fast"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcHunyuan3DTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ChannelErrCode": "",
          "ErrCode": 70000,
          "ErrCodeExt": "InvalidParameter",
          "Input": {
            "EnhancePrompt": "",
            "FileInfos": [],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Hailuo",
            "ModelVersion": "2.3-fast",
            "NegativePrompt": "",
            "OutputConfig": {
              "AspectRatio": "",
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
              "Resolution": "768P",
              "StorageMode": "Temporary"
            },
            "Prompt": "桌面上的咖啡杯，热气缓缓升起",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
          "Output": {
            "FileInfos": [],
            "ProcedureTaskIds": [],
            "Usage": {
              "InputImageCount": 0,
              "InputSeconds": 0,
              "InputTokens": 0,
              "OutputSeconds": 0,
              "ThoughtTokens": 0,
              "TotalSeconds": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-61ffa9f34cdb4336e1201698ea8b57a1t"
        },
        "BeginProcessTime": "2026-10-28T07:35:31Z",
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
        "CreateTime": "2026-10-28T07:35:30Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:37:02Z",
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
        "RequestId": "e5087b8e-9e84-4910-8eec-22de6db01d1d",
        "ReviewAudioVideoTask": null,
        "SceneAigcImageTask": null,
        "SceneAigcVideoTask": null,
        "SearchMediaBySemanticsAsync": null,
        "SnapshotByTimeOffsetTask": null,
        "SplitMediaTask": null,
        "Status": "FINISH",
        "TaskType": "AigcVideoTask",
        "TextToSpeechAsyncTask": null,
        "TranscodeTask": null,
        "VideoDepthEstimationTask": null,
        "VideoDubbingAsyncTask": null,
        "WandAsrTask": null,
        "WechatMiniProgramPublishTask": null,
        "WechatPublishTask": null
      }
    }
  }
}
```
