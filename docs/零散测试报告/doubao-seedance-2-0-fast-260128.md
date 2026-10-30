# doubao-seedance-2-0-fast-260128 人像库到生视频

调用地址：`http://127.0.0.1:9006`

令牌：`sk-xxx`

模型名没有渠道尾缀，创建素材组不传 `portrait_channel`，服务端写入 `volc`。

## 结论

链路完成。图片素材和视频素材第一次查询时已经是 `Active`。生视频任务 `task_wJ4H2tJCbHxgAaHM4pbqmbboBUQepuQF`（上游 `cgt-20261030102038-rfmpb`）状态 `SUCCESS`。

- 素材组：`group-20261030101826-rhjzc`，`portrait_channel=volc`
- 图片素材：`asset-20261030101835-l4dtc`
- 视频素材：`asset-20261030101836-l86cs`
- 人像库对象前缀账号：`2122754587`
- 成片：`https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332683859400000000000000000000ffffac1525a3649212.mp4`
- 计费：预扣 ¥5.774972，完成后按 `tokens=216900`、`modelRatio=2.66`、`otherMultiplier=0.5946` 结算为 ¥5.010370，退回 ¥0.764602。上游地址 `https://ark.cn-beijing.volces.com/api/v3/contents/generations/tasks`，渠道 #16。

素材提交后上游很快审过，中间没有采到 `Processing`。生视频在服务重启后的第一次查询已经是 `SUCCESS`，因此客户端轮询只保留这一次终态；上游成功响应在文末服务日志里。

## 1. 创建素材组

```text
2026-10-30T10:18:25+08:00
POST /v1/portrait/groups
http 200
elapsed 0.558s
```

请求体：

```json
{
  "name": "group-fast-260128-1030",
  "description": "doubao-seedance-2-0-fast-260128 portrait chain",
  "group_type": "AIGC"
}
```

响应头：

```text
Connection: close
Content-Length: 117
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:18:26 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300218257211840008268d9d6oqubkQLT
```

响应体：

```json
{
  "data": {
    "group_id": "group-20261030101826-rhjzc",
    "id": 4,
    "name": "group-fast-260128-1030"
  },
  "message": "",
  "success": true
}
```

## 2. 创建真人图片素材

```text
2026-10-30T10:18:35+08:00
POST /v1/portrait/assets
http 200
elapsed 1.137s
```

请求体：

```json
{
  "group_id": "group-20261030101826-rhjzc",
  "url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906",
  "asset_type": "Image",
  "name": "real-person-image"
}
```

响应头：

```text
Connection: close
Content-Length: 153
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:18:36 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300218352635610008268d9d6KW4tS6tE
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030101835-l4dtc",
    "id": 5,
    "remote_group_id": "group-20261030101826-rhjzc",
    "status": "Submitted"
  },
  "message": "",
  "success": true
}
```

## 3. 创建真人视频素材

```text
2026-10-30T10:18:36+08:00
POST /v1/portrait/assets
http 200
elapsed 0.686s
```

请求体：

```json
{
  "group_id": "group-20261030101826-rhjzc",
  "url": "https://aigc-bucket.su.bcebos.com/v1/objects/cf/cf3c317917f5209f5f8bdc5b9300377a6c13eddeeb3090cdfe8e294fea95ffcf",
  "asset_type": "Video",
  "name": "real-person-video"
}
```

响应头：

```text
Connection: close
Content-Length: 153
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:18:37 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300218363994820008268d9d6DyugCOAC
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030101836-l86cs",
    "id": 6,
    "remote_group_id": "group-20261030101826-rhjzc",
    "status": "Submitted"
  },
  "message": "",
  "success": true
}
```

## 4. 查询素材组列表

```text
2026-10-30T10:20:20+08:00
GET /v1/portrait/groups?p=1&page_size=20
http 200
elapsed 0.023s
```

响应头：

```text
Connection: close
Content-Length: 1264
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:20 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 20261030022020458620008268d9d6BVXViOum
```

响应体：

```json
{
  "data": {
    "page": 1,
    "page_size": 20,
    "total": 5,
    "items": [
      {
        "id": 5,
        "user_id": 1,
        "remote_group_id": "group-20261030102007-jqgcf",
        "name": "group-fast-zj-1030",
        "portrait_channel": "zj",
        "project_name": "",
        "created_at": "2026-10-30T10:20:07.06254+08:00",
        "updated_at": "2026-10-30T10:20:07.06254+08:00"
      },
      {
        "id": 4,
        "user_id": 1,
        "remote_group_id": "group-20261030101826-rhjzc",
        "name": "group-fast-260128-1030",
        "portrait_channel": "volc",
        "project_name": "",
        "created_at": "2026-10-30T10:18:26.256916+08:00",
        "updated_at": "2026-10-30T10:18:26.256916+08:00"
      },
      {
        "id": 3,
        "user_id": 1,
        "remote_group_id": "group-20261028153338-tkmd8",
        "name": "billing-test-1793172817493",
        "portrait_channel": "",
        "project_name": "",
        "created_at": "2026-10-28T15:33:37.975521+08:00",
        "updated_at": "2026-10-28T15:33:37.975521+08:00"
      },
      {
        "id": 2,
        "user_id": 1,
        "remote_group_id": "group-20261015095847-jnxz9",
        "name": "group-xzy-name-001",
        "portrait_channel": "",
        "project_name": "",
        "created_at": "2026-10-15T09:58:47.829039+08:00",
        "updated_at": "2026-10-15T09:58:47.829039+08:00"
      },
      {
        "id": 1,
        "user_id": 1,
        "remote_group_id": "group-20260520180204-zkcj6",
        "name": "group-test-seedance-user",
        "portrait_channel": "",
        "project_name": "",
        "created_at": "2026-05-20T18:02:04.926391+08:00",
        "updated_at": "2026-05-20T19:51:11.808048+08:00"
      }
    ]
  },
  "message": "",
  "success": true
}
```

## 5. 查询图片素材

```text
2026-10-30T10:20:20+08:00
GET /v1/portrait/assets/asset-20261030101835-l4dtc
http 200
elapsed 0.160s
```

响应头：

```text
Connection: close
Content-Length: 1273
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:20 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 20261030022020713690008268d9d6IMeuxY3P
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030101835-l4dtc",
    "asset_type": "Image",
    "created_at": "2026-10-30T10:18:36.384802+08:00",
    "name": "real-person-image",
    "portrait_channel": "volc",
    "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2122754587/103010183550282465.",
    "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906",
    "status": "Active",
    "updated_at": "2026-10-30T10:20:20.211085+08:00"
  },
  "message": "",
  "success": true
}
```

## 6. 查询视频素材

```text
2026-10-30T10:20:20+08:00
GET /v1/portrait/assets/asset-20261030101836-l86cs
http 200
elapsed 0.157s
```

响应头：

```text
Connection: close
Content-Length: 1272
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:20 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220202325540008268d9d68BoYaXmQ
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030101836-l86cs",
    "asset_type": "Video",
    "created_at": "2026-10-30T10:18:37.076052+08:00",
    "name": "real-person-video",
    "portrait_channel": "volc",
    "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2122754587/103010183709769800.",
    "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/cf/cf3c317917f5209f5f8bdc5b9300377a6c13eddeeb3090cdfe8e294fea95ffcf",
    "status": "Active",
    "updated_at": "2026-10-30T10:20:20.38319+08:00"
  },
  "message": "",
  "success": true
}
```

## 7. 按素材组查询素材列表

```text
2026-10-30T10:20:37+08:00
GET /v1/portrait/assets?group_id=group-20261030101826-rhjzc&p=1&page_size=20
http 200
elapsed 0.013s
```

响应头：

```text
Connection: close
Transfer-Encoding: chunked
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:37 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220378513330008268d9d6Z6YgE6Mv
```

响应体：

```json
{
  "data": {
    "page": 1,
    "page_size": 20,
    "total": 2,
    "items": [
      {
        "id": 6,
        "user_id": 1,
        "remote_group_id": "group-20261030101826-rhjzc",
        "remote_asset_id": "asset-20261030101836-l86cs",
        "portrait_channel": "volc",
        "name": "real-person-video",
        "asset_type": "Video",
        "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/cf/cf3c317917f5209f5f8bdc5b9300377a6c13eddeeb3090cdfe8e294fea95ffcf",
        "status": "Active",
        "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2122754587/103010183709769800.",
        "created_at": "2026-10-30T10:18:37.076052+08:00",
        "updated_at": "2026-10-30T10:20:20.38319+08:00"
      },
      {
        "id": 5,
        "user_id": 1,
        "remote_group_id": "group-20261030101826-rhjzc",
        "remote_asset_id": "asset-20261030101835-l4dtc",
        "portrait_channel": "volc",
        "name": "real-person-image",
        "asset_type": "Image",
        "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906",
        "status": "Active",
        "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2122754587/103010183550282465.",
        "created_at": "2026-10-30T10:18:36.384802+08:00",
        "updated_at": "2026-10-30T10:20:20.211085+08:00"
      }
    ]
  },
  "message": "",
  "success": true
}
```

## 8. 提交生视频

```text
2026-10-30T10:20:37+08:00
POST /v1/video/generations
http 200
elapsed 0.665s
```

请求体：

```json
{
  "model": "doubao-seedance-2-0-fast-260128",
  "prompt": "图片1里的人物保持样貌，按照视频1的动作轻轻活动，镜头稳定",
  "duration": 5,
  "resolution": "720p",
  "ratio": "adaptive",
  "metadata": {
    "generate_audio": true,
    "watermark": false,
    "content": [
      {
        "type": "image_url",
        "image_url": {
          "url": "asset://asset-20261030101835-l4dtc"
        }
      },
      {
        "type": "video_url",
        "role": "reference_video",
        "video_url": {
          "url": "asset://asset-20261030101836-l86cs"
        }
      }
    ]
  }
}
```

响应头：

```text
Connection: close
Content-Length: 210
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:38 GMT
X-New-Api-Other-Ratios: {"video_input":0.5945945945945946}
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220378620160008268d9d6PCkuOFv4
```

响应体：

```json
{
  "id": "task_wJ4H2tJCbHxgAaHM4pbqmbboBUQepuQF",
  "task_id": "task_wJ4H2tJCbHxgAaHM4pbqmbboBUQepuQF",
  "object": "video",
  "model": "doubao-seedance-2-0-fast-260128",
  "status": "queued",
  "progress": 0,
  "created_at": 1793326838
}
```

## 9. 查询生视频任务（终态）

```text
2026-10-30T10:23:12+08:00
GET /v1/video/generations/task_wJ4H2tJCbHxgAaHM4pbqmbboBUQepuQF
http 200
elapsed 0.008s
```

响应头：

```text
Connection: close
Content-Length: 1882
Content-Type: application/json
Date: Fri, 30 Oct 2026 02:23:12 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300223128892590008268d9d688x9xPpy
```

响应体：

```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 287,
    "created_at": 1793326838,
    "updated_at": 1793326992,
    "task_id": "task_wJ4H2tJCbHxgAaHM4pbqmbboBUQepuQF",
    "platform": "45",
    "user_id": 1,
    "group": "default",
    "channel_id": 16,
    "amount": "5.010370",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332683859400000000000000000000ffffac1525a3649212.mp4",
    "submit_time": 1793326838,
    "start_time": 1793326852,
    "finish_time": 1793326992,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-0-fast-260128",
      "origin_model_name": "doubao-seedance-2-0-fast-260128"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332683859400000000000000000000ffffac1525a3649212.mp4"
      },
      "created_at": 1793326838,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261030102038-rfmpb",
      "model": "doubao-seedance-2-0-fast-260128",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "720p",
      "seed": 23841,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793326977,
      "usage": {
        "completion_tokens": 216900,
        "total_tokens": 216900
      }
    }
  }
}
```

## 服务日志摘录

提交时的上游请求（重启前日志，原文）：

```text
[INFO] 2026/10/30 - 10:20:37 | 202610300220378620160008268d9d6PCkuOFv4 | [doubao-video] submit request: url=https://ark.cn-beijing.volces.com/api/v3/contents/generations/tasks body={"model":"doubao-seedance-2-0-fast-260128","content":[{"type":"image_url","image_url":{"url":"asset://asset-20261030101835-l4dtc"},"role":"reference_image"},{"type":"video_url","video_url":{"url":"asset://asset-20261030101836-l86cs"},"role":"reference_video"},{"type":"text","text":"图片1里的人物保持样貌，按照视频1的动作轻轻活动，镜头稳定"}],"generate_audio":true,"resolution":"720p","ratio":"adaptive","duration":5,"watermark":false}
```

重启后的轮询与结算：

```text
[SYS] 2026/10/30 - 10:23:12 | [doubao] poll response: body={"id":"cgt-20261030102038-rfmpb","model":"doubao-seedance-2-0-fast-260128","status":"succeeded","content":{"video_url":"https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332683859400000000000000000000ffffac1525a3649212.mp4"},"usage":{"completion_tokens":216900,"total_tokens":216900},"created_at":1793326838,"updated_at":1793326977,"seed":23841,"resolution":"720p","ratio":"16:9","duration":5,"framespersecond":24,"service_tier":"default","execution_expires_after":172800,"generate_audio":true,"draft":false,"priority":0,"output_format":"mp4"}
[INFO] 2026/10/30 - 10:23:12 | SYSTEM | 任务 task_wJ4H2tJCbHxgAaHM4pbqmbboBUQepuQF 差额结算：delta=¥-0.764602 额度（实际：¥5.010370 额度，预扣：¥5.774972 额度，token重算：tokens=216900, modelRatio=2.66, groupRatio=1.00, otherMultiplier=0.5946） 
[GIN] 2026/10/30 - 10:23:12 | relay | 202610300223128892590008268d9d688x9xPpy | 200 |    4.879125ms |       127.0.0.1 |     GET /v1/video/generations/task_wJ4H2tJCbHxgAaHM4pbqmbboBUQepuQF
```
