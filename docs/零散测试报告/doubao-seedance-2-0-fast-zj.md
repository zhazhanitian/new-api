# doubao-seedance-2-0-fast-zj 人像库到生视频

调用地址：`http://127.0.0.1:9006`

令牌：`sk-xxx`

模型名尾缀是 `-zj`，创建素材组传 `portrait_channel=zj`。

## 结论

链路完成。图片素材和视频素材第一次查询时已经是 `Active`。生视频任务 `task_qzLrKuE7TBdXas4isf9WQ7BqIJmcS6MS`（上游 `cgt-20261030102312-mnpw2`）状态 `SUCCESS`。

- 素材组：`group-20261030102007-jqgcf`，`portrait_channel=zj`
- 图片素材：`asset-20261030102007-65wrs`
- 视频素材：`asset-20261030102009-s9grg`
- 人像库对象前缀账号：`2103726075`（与官方模型的 `2122754587` 不是同一个账号）
- 成片：`https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332699292600000000000000000000ffffac18056ac81fff.mp4`
- 计费：渠道 #20，模型重定向到 `doubao-seedance-2-0-fast-260128`。预扣 ¥5.774972，完成后按 `tokens=216900`、`modelRatio=2.66`、`otherMultiplier=0.5946` 结算为 ¥5.010370，退回 ¥0.764602。日志有一句 `model price not found: doubao-seedance-2-0-fast-zj`，随后仍按重定向后的官方模型系数扣费。

过程中有两次失败，都记在下面：

1. 第一次创建素材组返回 `人像库渠道标识不存在或未配置`。`.env` 里已有 `PORTRAIT_ZJ_*`，但当时进程只导出了 `VOLC_PORTRAIT_*`。重启并带上 `PORTRAIT_ZJ_*` 后，第二次创建成功。
2. 第一次生视频打到 `https://aiplay.gee.cn/new-api/api/v3/contents/generations/tasks`，上游返回 404 `Invalid URL (POST /api/v3/contents/generations/tasks)`。原因是中纪适配器嵌入火山适配器后，`DoRequest` 仍走火山的 `/api/v3`。补上外层 `DoRequest` 并重启后，第二次提交返回 200，上游任务 `cgt-20261030102312-mnpw2`。服务日志里的 `submit request` 仍打印 `/api/v3`，那一行来自嵌入的火山 `BuildRequestBody`，不是实际发出的地址；实际请求能创建任务，说明走的是 `/v1/contents/generations/tasks`。

第一次成片画面停在参考图（桥上骑车），没有跟上参考视频。同一参数又提交了两次，见文末复测。两次成片都是室内喝茶，跟上了参考视频。中间轮询没有写入。

## 1. 创建素材组（进程尚未加载 zj 凭据，失败）

```text
2026-10-30T10:18:26+08:00
POST /v1/portrait/groups
http 200
elapsed 0.006s
```

请求体：

```json
{
  "name": "group-fast-zj-1030",
  "description": "doubao-seedance-2-0-fast-zj portrait chain",
  "group_type": "AIGC",
  "portrait_channel": "zj"
}
```

响应头：

```text
Connection: close
Content-Length: 72
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:18:26 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300218262808180008268d9d6v8J7yNZ4
```

响应体：

```json
{
  "message": "人像库渠道标识不存在或未配置",
  "success": false
}
```

## 2. 创建素材组（重启后，成功）

```text
2026-10-30T10:20:06+08:00
POST /v1/portrait/groups
http 200
elapsed 0.779s
```

请求体：

```json
{
  "name": "group-fast-zj-1030",
  "description": "doubao-seedance-2-0-fast-zj portrait chain",
  "group_type": "AIGC",
  "portrait_channel": "zj"
}
```

响应头：

```text
Connection: close
Content-Length: 113
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:07 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220062961980008268d9d6Y91momYm
```

响应体：

```json
{
  "data": {
    "group_id": "group-20261030102007-jqgcf",
    "id": 5,
    "name": "group-fast-zj-1030"
  },
  "message": "",
  "success": true
}
```

## 3. 创建真人图片素材

```text
2026-10-30T10:20:07+08:00
POST /v1/portrait/assets
http 200
elapsed 1.925s
```

请求体：

```json
{
  "group_id": "group-20261030102007-jqgcf",
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
Date: Fri, 30 Oct 2026 02:20:08 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 20261030022007746040008268d9d6zz7N5SYN
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030102007-65wrs",
    "id": 7,
    "remote_group_id": "group-20261030102007-jqgcf",
    "status": "Submitted"
  },
  "message": "",
  "success": true
}
```

## 4. 创建真人视频素材

```text
2026-10-30T10:20:08+08:00
POST /v1/portrait/assets
http 200
elapsed 0.810s
```

请求体：

```json
{
  "group_id": "group-20261030102007-jqgcf",
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
Date: Fri, 30 Oct 2026 02:20:09 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220089999730008268d9d6uDMCEJbz
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030102009-s9grg",
    "id": 8,
    "remote_group_id": "group-20261030102007-jqgcf",
    "status": "Submitted"
  },
  "message": "",
  "success": true
}
```

## 5. 查询素材组列表

```text
2026-10-30T10:20:20+08:00
GET /v1/portrait/groups?p=1&page_size=20
http 200
elapsed 0.005s
```

响应头：

```text
Connection: close
Content-Length: 1264
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:20 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 20261030022020666990008268d9d662VDN07f
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

## 6. 查询图片素材

```text
2026-10-30T10:20:20+08:00
GET /v1/portrait/assets/asset-20261030102007-65wrs
http 200
elapsed 0.164s
```

响应头：

```text
Connection: close
Content-Length: 1271
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:20 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220203908870008268d9d6RWxufmTG
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030102007-65wrs",
    "asset_type": "Image",
    "created_at": "2026-10-30T10:20:08.990917+08:00",
    "name": "real-person-image",
    "portrait_channel": "zj",
    "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2103726075/103010200751840335.",
    "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906",
    "status": "Active",
    "updated_at": "2026-10-30T10:20:20.544976+08:00"
  },
  "message": "",
  "success": true
}
```

## 7. 查询视频素材

```text
2026-10-30T10:20:20+08:00
GET /v1/portrait/assets/asset-20261030102009-s9grg
http 200
elapsed 0.203s
```

响应头：

```text
Connection: close
Content-Length: 1271
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:20 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220205550110008268d9d6e1X7WAcJ
```

响应体：

```json
{
  "data": {
    "asset_id": "asset-20261030102009-s9grg",
    "asset_type": "Video",
    "created_at": "2026-10-30T10:20:09.802584+08:00",
    "name": "real-person-video",
    "portrait_channel": "zj",
    "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2103726075/103010200900702122.",
    "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/cf/cf3c317917f5209f5f8bdc5b9300377a6c13eddeeb3090cdfe8e294fea95ffcf",
    "status": "Active",
    "updated_at": "2026-10-30T10:20:20.749992+08:00"
  },
  "message": "",
  "success": true
}
```

## 8. 按素材组查询素材列表

```text
2026-10-30T10:20:38+08:00
GET /v1/portrait/assets?group_id=group-20261030102007-jqgcf&p=1&page_size=20
http 200
elapsed 0.007s
```

响应头：

```text
Connection: close
Transfer-Encoding: chunked
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:38 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220385288700008268d9d6pmcbH0UG
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
        "id": 8,
        "user_id": 1,
        "remote_group_id": "group-20261030102007-jqgcf",
        "remote_asset_id": "asset-20261030102009-s9grg",
        "portrait_channel": "zj",
        "name": "real-person-video",
        "asset_type": "Video",
        "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/cf/cf3c317917f5209f5f8bdc5b9300377a6c13eddeeb3090cdfe8e294fea95ffcf",
        "status": "Active",
        "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2103726075/103010200900702122.",
        "created_at": "2026-10-30T10:20:09.802584+08:00",
        "updated_at": "2026-10-30T10:20:20.749992+08:00"
      },
      {
        "id": 7,
        "user_id": 1,
        "remote_group_id": "group-20261030102007-jqgcf",
        "remote_asset_id": "asset-20261030102007-65wrs",
        "portrait_channel": "zj",
        "name": "real-person-image",
        "asset_type": "Image",
        "source_url": "https://aigc-bucket.su.bcebos.com/v1/objects/f5/f55c3ffd3b8fbf7a219b1f5519d7186d03cd37833df142b8ded81ec25897c906",
        "status": "Active",
        "resolved_url": "https://ark-media-asset.tos-cn-beijing.volces.com/2103726075/103010200751840335.",
        "created_at": "2026-10-30T10:20:08.990917+08:00",
        "updated_at": "2026-10-30T10:20:20.544976+08:00"
      }
    ]
  },
  "message": "",
  "success": true
}
```

## 9. 提交生视频（仍走 /api/v3，失败）

```text
2026-10-30T10:20:38+08:00
POST /v1/video/generations
http 404
elapsed 0.265s
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
          "url": "asset://asset-20261030102007-65wrs"
        }
      },
      {
        "type": "video_url",
        "role": "reference_video",
        "video_url": {
          "url": "asset://asset-20261030102009-s9grg"
        }
      }
    ]
  }
}
```

响应头：

```text
Connection: close
Content-Length: 201
Content-Type: application/json; charset=utf-8
Date: Fri, 30 Oct 2026 02:20:38 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300220385351740008268d9d63XRoyyts
```

响应体：

```json
{
  "code": "fail_to_fetch_task",
  "message": "{\"error\":{\"message\":\"Invalid URL (POST /api/v3/contents/generations/tasks)\",\"type\":\"invalid_request_error\",\"param\":\"\",\"code\":\"\"}}",
  "data": null
}
```

## 10. 提交生视频（修正后，成功）

```text
2026-10-30T10:23:11+08:00
POST /v1/video/generations
http 200
elapsed 1.104s
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
          "url": "asset://asset-20261030102007-65wrs"
        }
      },
      {
        "type": "video_url",
        "role": "reference_video",
        "video_url": {
          "url": "asset://asset-20261030102009-s9grg"
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
Date: Fri, 30 Oct 2026 02:23:12 GMT
X-New-Api-Other-Ratios: {"video_input":0.5945945945945946}
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300223117875750008268d9d6wKEvn8k5
```

响应体：

```json
{
  "id": "task_qzLrKuE7TBdXas4isf9WQ7BqIJmcS6MS",
  "task_id": "task_qzLrKuE7TBdXas4isf9WQ7BqIJmcS6MS",
  "object": "video",
  "model": "doubao-seedance-2-0-fast-zj",
  "status": "queued",
  "progress": 0,
  "created_at": 1793326992
}
```

## 11. 查询生视频任务（终态）

```text
2026-10-30T10:26:13+08:00
GET /v1/video/generations/task_qzLrKuE7TBdXas4isf9WQ7BqIJmcS6MS
http 200
elapsed 0.016s
```

响应头：

```text
Connection: close
Content-Length: 1878
Content-Type: application/json
Date: Fri, 30 Oct 2026 02:26:13 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300226132925010008268d9d6i5LTu1wn
```

响应体：

```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 288,
    "created_at": 1793326992,
    "updated_at": 1793327169,
    "task_id": "task_qzLrKuE7TBdXas4isf9WQ7BqIJmcS6MS",
    "platform": "61",
    "user_id": 1,
    "group": "default",
    "channel_id": 20,
    "amount": "5.010370",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332699292600000000000000000000ffffac18056ac81fff.mp4",
    "submit_time": 1793326992,
    "start_time": 1793327008,
    "finish_time": 1793327169,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-0-fast-260128",
      "origin_model_name": "doubao-seedance-2-0-fast-zj"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332699292600000000000000000000ffffac18056ac81fff.mp4"
      },
      "created_at": 1793326992,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261030102312-mnpw2",
      "model": "doubao-seedance-2-0-fast-260128",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "720p",
      "seed": 33559,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793327163,
      "usage": {
        "completion_tokens": 216900,
        "total_tokens": 216900
      }
    }
  }
}
```

## 服务日志摘录

```text
[SYS] 2026/10/30 - 10:23:11 | model price not found: doubao-seedance-2-0-fast-zj 
[INFO] 2026/10/30 - 10:23:12 | 202610300223117875750008268d9d6wKEvn8k5 | [doubao-video] submit response: status=200 body={"id":"cgt-20261030102312-mnpw2"} 
[INFO] 2026/10/30 - 10:23:12 | 202610300223117875750008268d9d6wKEvn8k5 | record consume log: userId=1, params={"channel_id":20,"prompt_tokens":0,"completion_tokens":0,"model_name":"doubao-seedance-2-0-fast-zj","token_name":"开发部署模型测试","quota":395546,"content":"操作 generate, 计算参数：video_input: 0.59","token_id":1,"use_time_seconds":0,"is_stream":false,"group":"default","other":{"group_ratio":1,"is_model_mapped":true,"is_task":true,"model_price":-1,"model_ratio":2.66095,"request_path":"/v1/video/generations","upstream_model_name":"doubao-seedance-2-0-fast-260128"}} 
[SYS] 2026/10/30 - 10:26:09 | [doubao] poll response: body={"id": "cgt-20261030102312-mnpw2", "seed": 33559, "draft": false, "model": "doubao-seedance-2-0-fast-260128", "ratio": "16:9", "usage": {"total_tokens": 216900, "completion_tokens": 216900}, "status": "succeeded", "content": {"video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332699292600000000000000000000ffffac18056ac81fff.mp4"}, "duration": 5, "priority": 0, "created_at": 1793326992, "resolution": "720p", "updated_at": 1793327163, "service_tier": "default", "output_format": "mp4", "generate_audio": true, "framespersecond": 24, "execution_expires_after": 172800} 
[INFO] 2026/10/30 - 10:26:09 | SYSTEM | 任务 task_qzLrKuE7TBdXas4isf9WQ7BqIJmcS6MS 差额结算：delta=¥-0.764602 额度（实际：¥5.010370 额度，预扣：¥5.774972 额度，token重算：tokens=216900, modelRatio=2.66, groupRatio=1.00, otherMultiplier=0.5946） 
```

## 复测：同一参数再提交两次生视频

素材、提示词、分辨率与第一次成功的生视频相同。中间查询不记录，只保留创建和终态。

抽帧看画面：两次都是窗边喝茶，跟上了参考视频，没有停在参考图的桥上骑车。

## 12. 第一次复测：提交生视频

```text
2026-10-30T10:45:42+08:00
POST /v1/video/generations
http 200
elapsed 0.735s
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
          "url": "asset://asset-20261030102007-65wrs"
        }
      },
      {
        "type": "video_url",
        "role": "reference_video",
        "video_url": {
          "url": "asset://asset-20261030102009-s9grg"
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
Date: Fri, 30 Oct 2026 02:45:43 GMT
X-New-Api-Other-Ratios: {"video_input":0.5945945945945946}
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 202610300245423332170008268d9d6RxHnPcM9
```

响应体：

```json
{
  "id": "task_HuFbPgosWsYtBnX0NdutqU0HuBO6EhVm",
  "task_id": "task_HuFbPgosWsYtBnX0NdutqU0HuBO6EhVm",
  "object": "video",
  "model": "doubao-seedance-2-0-fast-zj",
  "status": "queued",
  "progress": 0,
  "created_at": 1793328343
}
```

## 13. 第一次复测：查询终态

```text
2026-10-30T10:48:34+08:00
GET /v1/video/generations/task_HuFbPgosWsYtBnX0NdutqU0HuBO6EhVm
http 200
elapsed 0.012s
```

响应头：

```text
Connection: close
Content-Length: 1878
Content-Type: application/json
Date: Fri, 30 Oct 2026 02:48:34 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 20261030024834804220008268d9d6AvgjaGcW
```

响应体：

```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 289,
    "created_at": 1793328343,
    "updated_at": 1793328508,
    "task_id": "task_HuFbPgosWsYtBnX0NdutqU0HuBO6EhVm",
    "platform": "61",
    "user_id": 1,
    "group": "default",
    "channel_id": 20,
    "amount": "5.010370",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332834307600000000000000000000ffffac18056a991ecc.mp4",
    "submit_time": 1793328343,
    "start_time": 1793328355,
    "finish_time": 1793328508,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-0-fast-260128",
      "origin_model_name": "doubao-seedance-2-0-fast-zj"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332834307600000000000000000000ffffac18056a991ecc.mp4"
      },
      "created_at": 1793328343,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261030104542-zffd3",
      "model": "doubao-seedance-2-0-fast-260128",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "720p",
      "seed": 51831,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793328506,
      "usage": {
        "completion_tokens": 216900,
        "total_tokens": 216900
      }
    }
  }
}
```

## 14. 第二次复测：提交生视频

```text
2026-10-30T10:45:43+08:00
POST /v1/video/generations
http 200
elapsed 0.586s
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
          "url": "asset://asset-20261030102007-65wrs"
        }
      },
      {
        "type": "video_url",
        "role": "reference_video",
        "video_url": {
          "url": "asset://asset-20261030102009-s9grg"
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
Date: Fri, 30 Oct 2026 02:45:43 GMT
X-New-Api-Other-Ratios: {"video_input":0.5945945945945946}
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 20261030024543644830008268d9d6kb07AvCF
```

响应体：

```json
{
  "id": "task_IWS0pysm6wKDrQPJ5Mn2sDAIxK60Z8JG",
  "task_id": "task_IWS0pysm6wKDrQPJ5Mn2sDAIxK60Z8JG",
  "object": "video",
  "model": "doubao-seedance-2-0-fast-zj",
  "status": "queued",
  "progress": 0,
  "created_at": 1793328343
}
```

## 15. 第二次复测：查询终态

```text
2026-10-30T10:48:04+08:00
GET /v1/video/generations/task_IWS0pysm6wKDrQPJ5Mn2sDAIxK60Z8JG
http 200
elapsed 0.010s
```

响应头：

```text
Connection: close
Content-Length: 1878
Content-Type: application/json
Date: Fri, 30 Oct 2026 02:48:04 GMT
X-New-Api-Version: v0.0.0
X-Oneapi-Request-Id: 20261030024804212060008268d9d6M59BDlJe
```

响应体：

```json
{
  "code": "success",
  "message": "",
  "data": {
    "id": 290,
    "created_at": 1793328343,
    "updated_at": 1793328476,
    "task_id": "task_IWS0pysm6wKDrQPJ5Mn2sDAIxK60Z8JG",
    "platform": "61",
    "user_id": 1,
    "group": "default",
    "channel_id": 20,
    "amount": "5.010370",
    "action": "generate",
    "status": "SUCCESS",
    "fail_reason": "",
    "result_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332834368200000000000000000000ffffac183e71d480fd.mp4",
    "submit_time": 1793328343,
    "start_time": 1793328356,
    "finish_time": 1793328476,
    "progress": "100%",
    "properties": {
      "input": "",
      "upstream_model_name": "doubao-seedance-2-0-fast-260128",
      "origin_model_name": "doubao-seedance-2-0-fast-zj"
    },
    "data": {
      "content": {
        "video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332834368200000000000000000000ffffac183e71d480fd.mp4"
      },
      "created_at": 1793328343,
      "draft": false,
      "duration": 5,
      "execution_expires_after": 172800,
      "framespersecond": 24,
      "generate_audio": true,
      "id": "cgt-20261030104543-kcbr6",
      "model": "doubao-seedance-2-0-fast-260128",
      "output_format": "mp4",
      "priority": 0,
      "ratio": "16:9",
      "resolution": "720p",
      "seed": 10776,
      "service_tier": "default",
      "status": "succeeded",
      "updated_at": 1793328453,
      "usage": {
        "completion_tokens": 216900,
        "total_tokens": 216900
      }
    }
  }
}
```

## 复测服务日志摘录

```text
[INFO] 2026/10/30 - 10:45:43 | 202610300245423332170008268d9d6RxHnPcM9 | [doubao-video] submit response: status=200 body={"id":"cgt-20261030104542-zffd3"} 
[INFO] 2026/10/30 - 10:45:43 | 20261030024543644830008268d9d6kb07AvCF | [doubao-video] submit response: status=200 body={"id":"cgt-20261030104543-kcbr6"} 
[SYS] 2026/10/30 - 10:47:56 | [doubao] poll response: body={"id": "cgt-20261030104543-kcbr6", "seed": 10776, "draft": false, "model": "doubao-seedance-2-0-fast-260128", "ratio": "16:9", "usage": {"total_tokens": 216900, "completion_tokens": 216900}, "status": "succeeded", "content": {"video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332834368200000000000000000000ffffac183e71d480fd.mp4"}, "duration": 5, "priority": 0, "created_at": 1793328343, "resolution": "720p", "updated_at": 1793328453, "service_tier": "default", "output_format": "mp4", "generate_audio": true, "framespersecond": 24, "execution_expires_after": 172800} 
[INFO] 2026/10/30 - 10:47:56 | SYSTEM | 任务 task_IWS0pysm6wKDrQPJ5Mn2sDAIxK60Z8JG 差额结算：delta=¥-0.764602 额度（实际：¥5.010370 额度，预扣：¥5.774972 额度，token重算：tokens=216900, modelRatio=2.66, groupRatio=1.00, otherMultiplier=0.5946） 
[SYS] 2026/10/30 - 10:48:28 | [doubao] poll response: body={"id": "cgt-20261030104542-zffd3", "seed": 51831, "draft": false, "model": "doubao-seedance-2-0-fast-260128", "ratio": "16:9", "usage": {"total_tokens": 216900, "completion_tokens": 216900}, "status": "succeeded", "content": {"video_url": "https://ark-acg-cn-beijing.tos-cn-beijing.volces.com/doubao-seedance-2-0-fast/02179332834307600000000000000000000ffffac18056a991ecc.mp4"}, "duration": 5, "priority": 0, "created_at": 1793328343, "resolution": "720p", "updated_at": 1793328506, "service_tier": "default", "output_format": "mp4", "generate_audio": true, "framespersecond": 24, "execution_expires_after": 172800} 
[INFO] 2026/10/30 - 10:48:28 | SYSTEM | 任务 task_HuFbPgosWsYtBnX0NdutqU0HuBO6EhVm 差额结算：delta=¥-0.764602 额度（实际：¥5.010370 额度，预扣：¥5.774972 额度，token重算：tokens=216900, modelRatio=2.66, groupRatio=1.00, otherMultiplier=0.5946） 
```
