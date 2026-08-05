# hailuo-video-2.3-fast 测试报告

> 本文件由 `tencent_vod_video_test.js` 自动追加生成，禁止手动修改序号。

## TC-27（seq=27）: 默认（768P × 6s）
> 执行时间：2026/8/5 10:53:20  |  模型：`hailuo-video-2.3-fast`  |  标签：standard

> 💡 768P 30822×6=184932

### 调用参数
```json
{
  "model": "hailuo-video-2.3-fast",
  "prompt": "樱花满开的公园小径，花瓣随风飘落，慢动作，浪漫唯美"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 未传 | **6** | 未传，使用表达式默认值 |
| size（分辨率） | 未传 | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **92,466** |
| 预期 USD | $0.1849 |
| 预期 RMB | ¥1.3500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.17s |
| task_id | `task_sGs6QnmIyKyfAq6S0schdyM6NGB0VU0l` |

```json
{
  "id": "task_sGs6QnmIyKyfAq6S0schdyM6NGB0VU0l",
  "task_id": "task_sGs6QnmIyKyfAq6S0schdyM6NGB0VU0l",
  "object": "video",
  "model": "hailuo-video-2.3-fast",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898383
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 10,967,539 | — |
| 提交后可用 Quota | 10,875,073 | — |
| **预扣金额** | **92,466** | **¥1.3500** |
| 预期扣减 | 92,466 | ¥1.3500 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 92,466 | — |
| request_id | 20260805025303676828000f61X4gQRB2ONg0wE | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **FAILURE** |
| 轮询次数 | 2 |
| 完成耗时 | 15.0s |
| progress | 100% |

**视频 URL**: task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode

**失败原因**: task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 92,466 | ¥1.3500 |
| 平台记录最终消费 | 92,466 | **¥1.350004** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 128,
    "created_at": 1785898383,
    "updated_at": 1785898399,
    "task_id": "task_sGs6QnmIyKyfAq6S0schdyM6NGB0VU0l",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.350004",
    "action": "generate",
    "status": "FAILURE",
    "fail_reason": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
    "result_url": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
    "submit_time": 1785898383,
    "start_time": 0,
    "finish_time": 1785898399,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-2.3-fast",
      "origin_model_name": "hailuo-video-2.3-fast"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ErrCode": 70000,
          "ErrCodeExt": "InternalError",
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
              "Duration": 6,
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
            "Prompt": "樱花满开的公园小径，花瓣随风飘落，慢动作，浪漫唯美",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
          "Output": {
            "FileInfos": [],
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
          "TaskId": "1500065731-AigcVideoTask-ea61bbf4d13bd88c7167246cf686de9bt"
        },
        "BeginProcessTime": "2026-08-05T02:53:03Z",
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
        "CreateTime": "2026-08-05T02:53:03Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:53:07Z",
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
        "RequestId": "da8c110b-e6f6-4aae-8dac-127b2e1cf207",
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

## TC-28（seq=28）: duration=10 + size=1080P
> 执行时间：2026/8/5 10:53:37  |  模型：`hailuo-video-2.3-fast`  |  标签：standard

> 💡 1080P 52740×10=527400

### 调用参数
```json
{
  "model": "hailuo-video-2.3-fast",
  "prompt": "樱花满开的公园小径，花瓣随风飘落，慢动作，浪漫唯美",
  "duration": 10,
  "size": "1080P"
}
```

### 价格变量核对
| 变量 | 请求值 | 有效计费值 | 说明 |
|---|---|---|---|
| duration（时长/s） | 10 | **10** | 已传，直接使用 |
| size（分辨率） | 1080P | — | 未传时适配器默认 720P |
| audio_generation | 未传 | — | 未传视为无声版 |
| images（参考图数） | 无 | — | 文生视频 |

### 预期扣费
| 项目 | 值 |
|---|---|
| 预期 HTTP 状态 | 200 |
| 预期 Quota | **263,700** |
| 预期 USD | $0.5274 |
| 预期 RMB | ¥3.8500 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| 耗时 | 0.18s |
| task_id | `task_xW2EJvrzZ2z7hpVfJNtPQZRkA5xNAn4v` |

```json
{
  "id": "task_xW2EJvrzZ2z7hpVfJNtPQZRkA5xNAn4v",
  "task_id": "task_xW2EJvrzZ2z7hpVfJNtPQZRkA5xNAn4v",
  "object": "video",
  "model": "hailuo-video-2.3-fast",
  "status": "queued",
  "progress": 0,
  "created_at": 1785898400
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 提交前可用 Quota | 10,870,108 | — |
| 提交后可用 Quota | 10,606,408 | — |
| **预扣金额** | **263,700** | **¥3.8500** |
| 预期扣减 | 263,700 | ¥3.8500 |
| 预扣是否符合 | ✅ 符合 | — |
| 消费日志 Quota | 263,700 | — |
| request_id | 20260805025320391548000f61X4gQRMuDCRPmV | — |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **FAILURE** |
| 轮询次数 | 2 |
| 完成耗时 | 15.0s |
| progress | 100% |

**视频 URL**: task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode

**失败原因**: task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode

#### 最终结算（多退少补）
| 项目 | Quota | RMB |
|---|---|---|
| 预扣金额 | 263,700 | ¥3.8500 |
| 平台记录最终消费 | 263,700 | **¥3.850020** |
| 差额 | — | +¥0.0000（补扣） |

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 129,
    "created_at": 1785898400,
    "updated_at": 1785898415,
    "task_id": "task_xW2EJvrzZ2z7hpVfJNtPQZRkA5xNAn4v",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "3.850020",
    "action": "generate",
    "status": "FAILURE",
    "fail_reason": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
    "result_url": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
    "submit_time": 1785898400,
    "start_time": 0,
    "finish_time": 1785898415,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-2.3-fast",
      "origin_model_name": "hailuo-video-2.3-fast"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ErrCode": 70000,
          "ErrCodeExt": "InternalError",
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
              "Duration": 10,
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
            "Prompt": "樱花满开的公园小径，花瓣随风飘落，慢动作，浪漫唯美",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "task failed with status: FAIL, message: create gen_video task failed. ret:2013,msg:invalid params, model MiniMax-Hailuo-2.3-Fast does not support Text-to-Video mode",
          "Output": {
            "FileInfos": [],
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
          "TaskId": "1500065731-AigcVideoTask-3746f55c45b192e3c812490f2e9fdc09t"
        },
        "BeginProcessTime": "2026-08-05T02:53:20Z",
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
        "CreateTime": "2026-08-05T02:53:20Z",
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-08-05T02:53:24Z",
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
        "RequestId": "de8b93fc-5c49-43f2-b945-040950fcb0b5",
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

