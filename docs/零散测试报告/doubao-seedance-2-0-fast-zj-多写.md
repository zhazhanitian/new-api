# doubao-seedance-2-0-fast-zj 人像库到生视频（多写）

调用地址：`http://127.0.0.1:9006`

令牌：`sk-xxx`

不传 `portrait_channel`。素材组、图片素材、视频素材与火山模型是同一次创建，不另建一套。`model` 使用 `doubao-seedance-2-0-fast-zj`，`asset://` 仍是上面的逻辑素材 ID。

## 结论

链路完成。素材审核过程与火山模型那份记录相同：视频素材在中纪尚未 `Active` 时，汇总状态保持 `Processing`。两边都 `Active` 之后，用同一对 `asset://` 提交。生视频任务 `task_9OppsFEqzMmPV3iwBRIhFJzNlZNChGN6`（上游 `cgt-20261030141659-y46qs`）状态 `SUCCESS`。

- 素材组：`group-20261030141637-wd9qc`。中纪上游组 `group-20261030141637-zcthm`
- 图片素材：`asset://asset-20261030141637-k9fr4`，发往中纪时换成 `asset-20261030141639-mfbb4`
- 视频素材：`asset://asset-20261030141640-rhvwc`，发往中纪时换成 `asset-20261030141641-cldjs`
- 人像库对象前缀账号：中纪 `2103726075`（火山是 `2122754587`）
- 成片：`https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179334102097500000000000000000000ffffac182c4fe6f08c.mp4`
- 计费：渠道 #20，模型重定向到 `doubao-seedance-2-0-fast-260128`。预扣 ¥5.774972，完成后按 `tokens=216900`、`modelRatio=2.66`、`otherMultiplier=0.5946` 结算为 ¥5.010370，退回 ¥0.764602

本次提交一次成功，没有再出现 `/api/v3` 的 404。中间轮询没有写入。

## 1. 创建素材组

```text
2026-10-30T14:16:36+08:00
POST /v1/portrait/groups
http 200
elapsed 0.83s
```

请求体：

```json
{
  "name": "group-fast-fanout-1030",
  "description": "fanout portrait chain for fast 260128 and zj",
  "group_type": "AIGC"
}
```

响应头：

```text
Connection: close
Content-Length: 117
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 06:16:37 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300616365056660008268d9d6uB6kvU0X
```

响应体：

```json
{
  "data": {
    "group_id": "group-20261030141637-wd9qc",
    "id": 6,
    "name": "group-fast-fanout-1030"
  },
  "message": "",
  "success": true
}
```

## 2. 创建真人图片素材

```text
2026-10-30T14:16:37+08:00
POST /v1/portrait/assets
http 200
elapsed 3.202s
```

请求体：

```json
{
  "group_id": "group-20261030141637-wd9qc",
  "url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906",
  "asset_type": "Image",
  "name": "real-person-image"
}
```

响应头：

```text
Connection: close
Content-Length: 262
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 06:16:40 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300616373358500008268d9d6TS7uzAsu
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030141637-k9fr4",
    "channels": [
      {
        "portrait_channel": "volc",
        "status": "Submitted"
      },
      {
        "portrait_channel": "zj",
        "status": "Submitted"
      }
    ],
    "id": 9,
    "remote_group_id": "group-20261030141637-wd9qc",
    "status": "Submitted"
  },
  "message": "",
  "success": true
}
```

## 3. 创建真人视频素材

```text
2026-10-30T14:16:40+08:00
POST /v1/portrait/assets
http 200
elapsed 1.54s
```

请求体：

```json
{
  "group_id": "group-20261030141637-wd9qc",
  "url": "https://aigc-bucket.su.bcebos.com/v1/objects/cf/cf3c317917f5209f5f8bdc5b9300377a6c13eddeeb3090cdfe8e294fea95ffcf",
  "asset_type": "Video",
  "name": "real-person-video"
}
```

响应头：

```text
Connection: close
Content-Length: 263
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 06:16:42 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300616405384910008268d9d6R7Ljd6xl
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030141640-rhvwc",
    "channels": [
      {
        "portrait_channel": "volc",
        "status": "Submitted"
      },
      {
        "portrait_channel": "zj",
        "status": "Submitted"
      }
    ],
    "id": 10,
    "remote_group_id": "group-20261030141637-wd9qc",
    "status": "Submitted"
  },
  "message": "",
  "success": true
}
```

## 4. 查询素材组列表

```text
2026-10-30T14:16:58+08:00
GET /v1/portrait/groups?p=1&page_size=20
http 200
elapsed 0.008s
```

响应头：

```text
Connection: close
Content-Length: 1362
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 06:16:58 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300616586421880008268d9d6HHp38KzG
```

响应体：

```json
{
  "data": {
    "page": 1,
    "page_size": 20,
    "total": 6,
    "items": [
      {
        "id": 6,
        "user_id": 1,
        "remote_group_id": "group-20261030141637-wd9qc",
        "name": "group-fast-fanout-1030",
        "project_name": "",
        "created_at": "2026-10-30T14:16:37.313169+08:00",
        "updated_at": "2026-10-30T14:16:37.313169+08:00"
      },
      {
        "id": 5,
        "user_id": 1,
        "remote_group_id": "group-20261030102007-jqgcf",
        "name": "group-fast-zj-1030",
        "project_name": "",
        "created_at": "2026-10-30T10:20:07.06254+08:00",
        "updated_at": "2026-10-30T10:20:07.06254+08:00"
      },
      {
        "id": 4,
        "user_id": 1,
        "remote_group_id": "group-20261030101826-rhjzc",
        "name": "group-fast-260128-1030",
        "project_name": "",
        "created_at": "2026-10-30T10:18:26.256916+08:00",
        "updated_at": "2026-10-30T10:18:26.256916+08:00"
      },
      {
        "id": 3,
        "user_id": 1,
        "remote_group_id": "group-20261028153338-tkmd8",
        "name": "billing-test-1793172817493",
        "project_name": "",
        "created_at": "2026-10-28T15:33:37.975521+08:00",
        "updated_at": "2026-10-28T15:33:37.975521+08:00"
      },
      {
        "id": 2,
        "user_id": 1,
        "remote_group_id": "group-20261015095847-jnxz9",
        "name": "group-xzy-name-001",
        "project_name": "",
        "created_at": "2026-10-15T09:58:47.829039+08:00",
        "updated_at": "2026-10-15T09:58:47.829039+08:00"
      },
      {
        "id": 1,
        "user_id": 1,
        "remote_group_id": "group-20260520180204-zkcj6",
        "name": "group-test-seedance-user",
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
2026-10-30T14:16:58+08:00
GET /v1/portrait/assets/asset-20261030141637-k9fr4
http 200
elapsed 0.006s
```

响应头：

```text
Connection: close
Content-Length: 1190
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 06:16:58 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300616586503230008268d9d66bEbFwIG
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030141637-k9fr4",
    "asset_type": "Image",
    "channels": [
      {
        "portrait_channel": "volc",
        "status": "Active"
      },
      {
        "portrait_channel": "zj",
        "status": "Active"
      }
    ],
    "created_at": "2026-10-30T14:16:40.517122+08:00",
    "name": "real-person-image",
    "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2122754587/103014163770948078.",
    "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906",
    "status": "Active",
    "updated_at": "2026-10-30T14:16:42.3807+08:00"
  },
  "message": "",
  "success": true
}
```

## 6. 查询视频素材

```text
2026-10-30T14:16:58+08:00
GET /v1/portrait/assets/asset-20261030141640-rhvwc
http 200
elapsed 0.007s
```

响应头：

```text
Connection: close
Content-Length: 1192
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 06:16:58 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300616586563350008268d9d6lclnd6RS
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030141640-rhvwc",
    "asset_type": "Video",
    "channels": [
      {
        "portrait_channel": "volc",
        "status": "Active"
      },
      {
        "portrait_channel": "zj",
        "status": "Active"
      }
    ],
    "created_at": "2026-10-30T14:16:42.062839+08:00",
    "name": "real-person-video",
    "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2122754587/103014164135272946.",
    "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/cf/cf3c317917f5209f5f8bdc5b9300377a6c13eddeeb3090cdfe8e294fea95ffcf",
    "status": "Active",
    "updated_at": "2026-10-30T14:16:58.634303+08:00"
  },
  "message": "",
  "success": true
}
```

## 7. 按素材组查询素材列表

```text
2026-10-30T14:16:58+08:00
GET /v1/portrait/assets?group_id=group-20261030141637-wd9qc&p=1&page_size=20
http 200
elapsed 0.009s
```

响应头：

```text
Connection: close
Transfer-Encoding: chunked
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 06:16:58 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300616586632110008268d9d6S366HOZ0
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
        "id": 10,
        "user_id": 1,
        "remote_group_id": "group-20261030141637-wd9qc",
        "remote_asset_id": "asset-20261030141640-rhvwc",
        "name": "real-person-video",
        "asset_type": "Video",
        "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/cf/cf3c317917f5209f5f8bdc5b9300377a6c13eddeeb3090cdfe8e294fea95ffcf",
        "status": "Active",
        "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2122754587/103014164135272946.",
        "created_at": "2026-10-30T14:16:42.062839+08:00",
        "updated_at": "2026-10-30T14:16:58.634303+08:00",
        "channels": [
          {
            "portrait_channel": "volc",
            "status": "Active"
          },
          {
            "portrait_channel": "zj",
            "status": "Active"
          }
        ]
      },
      {
        "id": 9,
        "user_id": 1,
        "remote_group_id": "group-20261030141637-wd9qc",
        "remote_asset_id": "asset-20261030141637-k9fr4",
        "name": "real-person-image",
        "asset_type": "Image",
        "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906",
        "status": "Active",
        "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2122754587/103014163770948078.",
        "created_at": "2026-10-30T14:16:40.517122+08:00",
        "updated_at": "2026-10-30T14:16:42.3807+08:00",
        "channels": [
          {
            "portrait_channel": "volc",
            "status": "Active"
          },
          {
            "portrait_channel": "zj",
            "status": "Active"
          }
        ]
      }
    ]
  },
  "message": "",
  "success": true
}
```

## 8. 提交生视频

```text
2026-10-30T14:16:59+08:00
POST /v1/video/generations
http 200
elapsed 1.244s
```

请求体：

```json
{
  "model": "doubao-seedance-2-0-fast-zj",
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
          "url": "asset://asset-20261030141637-k9fr4"
        }
      },
      {
        "type": "video_url",
        "role": "reference_video",
        "video_url": {
          "url": "asset://asset-20261030141640-rhvwc"
        }
      }
    ]
  }
}
```

响应头：

```text
Connection: close
Content-Length: 206
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 06:17:00 GMT
X-New-Api-Other-Ratios: {"video_input":0.5945945945945946}
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300616596097040008268d9d6t9N0eGgY
```

响应体：

```json
{
  "id": "task_9OppsFEqzMmPV3iwBRIhFJzNlZNChGN6",
  "task_id": "task_9OppsFEqzMmPV3iwBRIhFJzNlZNChGN6",
  "object": "video",
  "model": "doubao-seedance-2-0-fast-zj",
  "status": "queued",
  "progress": 0,
  "created_at": 1793341020
}
```

## 9. 查询生视频任务（终态）

```text
2026-10-30T14:20:31+08:00
GET /v1/video/generations/task_9OppsFEqzMmPV3iwBRIhFJzNlZNChGN6
http 200
elapsed 0.008s
```

响应头：

```text
Connection: close
Content-Length: 1878
Content-Type: application/json
Date: Fri, 30 Oct 2026 06:20:31 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300620312896510008268d9d6qFHbpGNE
```

响应体：

```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 292,
    "created_at": 1793341020,
    "updated_at": 1793341222,
    "task_id": "task_9OppsFEqzMmPV3iwBRIhFJzNlZNChGN6",
    "platform": "61",
    "user_id": 1,
    "group": "default",
    "channel_id": 20,
    "amount": "5.010370",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179334102097500000000000000000000ffffac182c4fe6f08c.mp4",
    "submit_time": 1793341020,
    "start_time": 1793341037,
    "finish_time": 1793341222,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-0-fast-260128",
      "origin_model_name": "doubao-seedance-2-0-fast-zj"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179334102097500000000000000000000ffffac182c4fe6f08c.mp4"
      },
      "created_at": 1793341020,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261030141659-y46qs",
      "model": "doubao-seedance-2-0-fast-260128",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "720p",
      "seed": 35888,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793341210,
      "usage": {
        "completion_tokens": 216900,
        "total_tokens": 216900
      }
    }
  }
}
```

## 服务日志摘录

提交时的 `submit request` 和 `model price not found` 行已经滚出本地进程输出，这里只保留仍在输出里的轮询和结算。

```text
[SYS] 2026/10/30 - 14:20:22 | [doubao] poll response: body={"id": "cgt-20261030141659-y46qs", "seed": 35888, "draft": false, "model": "doubao-seedance-2-0-fast-260128", "ratio": "16:9", "usage": {"total_tokens": 216900, "completion_tokens": 216900}, "status": "succeeded", "content": {"video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179334102097500000000000000000000ffffac182c4fe6f08c.mp4"}, "duration": 5, "priority": 0, "created_at": 1793341020, "resolution": "720p", "updated_at": 1793341210, "service_tier": "default", "output_format": "mp4", "generate_audio": true, "framespersecond": 24, "execution_expires_after": 172800}
[INFO] 2026/10/30 - 14:20:22 | SYSTEM | 任务 task_9OppsFEqzMmPV3iwBRIhFJzNlZNChGN6 差额结算：delta=¥-0.764602 额度（实际：¥5.010370 额度，预扣：¥5.774972 额度，token重算：tokens=216900, modelRatio=2.66, groupRatio=1.00, otherMultiplier=0.5946） 
[GIN] 2026/10/30 - 14:20:31 | relay | 202610300620312896510008268d9d6qFHbpGNE | 200 |      4.7695ms |       127.0.0.1 |     GET /v1/video/generations/task_9OppsFEqzMmPV3iwBRIhFJzNlZNChGN6
```
