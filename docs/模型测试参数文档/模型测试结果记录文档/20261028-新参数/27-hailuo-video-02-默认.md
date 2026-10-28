# hailuo-video-02 测试报告

## 27. hailuo-video-02 默认
> 执行时间：2026-10-28  |  模型：`hailuo-video-02`

> 表达式值 271230，quota 135615

### 调用参数
```json
{
  "model": "hailuo-video-02",
  "prompt": "雨中的咖啡馆橱窗，行人撑伞经过"
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
| 预期 Quota | **135,615** |
| 预期 RMB | ¥1.9800 |

### 提交结果
| 项目 | 值 |
|---|---|
| HTTP 状态 | 200（预期 200）✅ |
| task_id | `task_Zuio0lqK5H3VTZxBG5wpa8WyAobsGTrS` |

```json
{
  "id": "task_Zuio0lqK5H3VTZxBG5wpa8WyAobsGTrS",
  "task_id": "task_Zuio0lqK5H3VTZxBG5wpa8WyAobsGTrS",
  "object": "video",
  "model": "hailuo-video-02",
  "status": "queued",
  "progress": 0,
  "created_at": 1793172929
}
```

### 扣费分析（提交时）
| 项目 | Quota | RMB |
|---|---|---|
| 预期扣减 | 135,615 | ¥1.9800 |
| 日志预扣 | 135,615 | ¥1.9800 |
| 平台记录最终消费 | 132,789 | ¥1.9387 |
| 是否符合 | ✅ 符合 | |
| request_id | - | |

### 任务轮询
| 项目 | 值 |
|---|---|
| 最终状态 | **SUCCESS** |
| 轮询次数 | 3 |
| progress | 100% |

**视频 URL**: https://store.vod-qcloud.com/b8c85f46vodnj251000800/91a5cfde5001834825866011195/aigcVideoGenFile.mp4

#### 实际输出元数据
| 项目 | 值 |
|---|---|
| 实际视频时长 | 5.875 |
| 输出分辨率 | 1366×768 |
| tokens | - |

#### 查询调用参数

`GET /v1/video/generations/task_Zuio0lqK5H3VTZxBG5wpa8WyAobsGTrS`

#### 最终查询响应（全量 JSON）
```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 251,
    "created_at": 1793172929,
    "updated_at": 1793173028,
    "task_id": "task_Zuio0lqK5H3VTZxBG5wpa8WyAobsGTrS",
    "platform": "58",
    "user_id": 1,
    "group": "default",
    "channel_id": 13,
    "amount": "1.938719",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/91a5cfde5001834825866011195/aigcVideoGenFile.mp4",
    "submit_time": 1793172929,
    "start_time": 1793172969,
    "finish_time": 1793173028,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "hailuo-video-02",
      "origin_model_name": "hailuo-video-02"
    },
    "data": {
      "Response": {
        "AigcAudioTask": null,
        "AigcHunyuan3DTask": null,
        "AigcImageTask": null,
        "AigcVideoRedrawTask": null,
        "AigcVideoTask": {
          "ChannelErrCode": "",
          "ErrCode": 0,
          "ErrCodeExt": "",
          "Input": {
            "EnhancePrompt": "",
            "FileInfos": [],
            "GenerationMode": "",
            "InputRegion": "",
            "LastFrameFileId": "",
            "LastFrameUrl": "",
            "ModelName": "Hailuo",
            "ModelVersion": "02",
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
            "Prompt": "雨中的咖啡馆橱窗，行人撑伞经过",
            "SceneType": "",
            "SubjectInfos": []
          },
          "Message": "",
          "Output": {
            "FileInfos": [
              {
                "ClassId": 0,
                "ExpireTime": "2026-11-04T07:37:08Z",
                "FileContent": "",
                "FileId": "",
                "FileType": "",
                "FileUrl": "https://store.vod-qcloud.com/b8c85f46vodnj251000800/91a5cfde5001834825866011195/aigcVideoGenFile.mp4",
                "MediaName": "",
                "MetaData": {
                  "AudioDuration": 0,
                  "AudioStreamSet": [],
                  "Bitrate": 850109,
                  "Container": "mov,mp4,m4a,3gp,3g2,mj2",
                  "Duration": 5.875,
                  "Height": 768,
                  "Md5": "",
                  "Rotate": 0,
                  "Size": 624299,
                  "VideoDuration": 5.875,
                  "VideoStreamSet": [
                    {
                      "Bitrate": 844850,
                      "Codec": "h264",
                      "CodecTag": "",
                      "Codecs": "",
                      "DynamicRangeInfo": {
                        "HDRType": "",
                        "Type": "Unknown"
                      },
                      "Fps": 24,
                      "Height": 768,
                      "Width": 1366
                    }
                  ],
                  "Width": 1366
                },
                "StorageMode": "Temporary",
                "UsageType": ""
              }
            ],
            "ProcedureTaskIds": [],
            "Usage": {
              "InputImageCount": 0,
              "InputSeconds": 0,
              "InputTokens": 0,
              "OutputSeconds": 5,
              "ThoughtTokens": 0,
              "TotalSeconds": 0
            }
          },
          "Progress": 100,
          "SessionContext": "",
          "SessionId": "",
          "Status": "FINISH",
          "TaskId": "1500065731-AigcVideoTask-5e2f7a218e6b33e8890499cdd6a60d6ft"
        },
        "BeginProcessTime": "2026-10-28T07:35:29Z",
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
        "CreateTime": "2026-10-28T07:35:29Z",
        "DNATask": null,
        "DescribeAigcFaceInfoAsyncTask": null,
        "DescribeFileAttributesTask": null,
        "DesignVoiceAsyncTask": null,
        "EditMediaTask": null,
        "ExtractBlindWatermarkTask": null,
        "ExtractCopyRightWatermarkTask": null,
        "ExtractTraceWatermarkTask": null,
        "FastClipMediaTask": null,
        "FinishTime": "2026-10-28T07:37:08Z",
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
        "RequestId": "fcece478-7127-4524-8e6f-0c4a7a39ff75",
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
