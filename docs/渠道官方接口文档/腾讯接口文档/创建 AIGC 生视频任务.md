## 1. 接口描述

接口请求域名： vod.tencentcloudapi.com 。

该接口用于[生成 AIGC 视频](https://cloud.tencent.com/document/product/266/124474)。默认限制1个并发处理，接口调用会产生实际费用，请参考点播 [AIGC 生视频计费文档](https://cloud.tencent.com/document/product/266/95125#96b3b59a-f9e1-49e9-966a-bedb70a4bf12)。该功能结算模式为[后付费](https://cloud.tencent.com/document/product/266/2838)，日结客户当天使用将在第二天出账，月结客户将在次月1日统一出上月使用费用。

默认接口请求频率限制：20次/秒。

<div class="rno-api-explorer">
    <div class="rno-api-explorer-inner">
        <div class="rno-api-explorer-hd">
            <div class="rno-api-explorer-title">
                推荐使用 API Explorer
            </div>
            <a href="https://console.cloud.tencent.com/api/explorer?Product=vod&Version=2018-07-17&Action=CreateAigcVideoTask" class="rno-api-explorer-btn" hotrep="doc.api.explorerbtn"><i class="rno-icon-explorer"></i>点击调试</a>
        </div>
        <div class="rno-api-explorer-body">
            <div class="rno-api-explorer-cont">
                API Explorer 提供了在线调用、签名验证、SDK 代码生成和快速检索接口等能力。您可查看每次调用的请求内容和返回结果以及自动生成 SDK 调用示例。
            </div>
        </div>
    </div>
</div>

## 2. 输入参数

以下请求参数列表仅列出了接口请求参数和部分公共参数，完整公共参数列表见 [公共请求参数](/document/api/266/31756)。

| 参数名称 | 必选 | 类型 | 描述 |
|---------|---------|---------|---------|
| Action | 是 | String | [公共参数](/document/api/266/31756)，本接口取值：CreateAigcVideoTask。 |
| Version | 是 | String | [公共参数](/document/api/266/31756)，本接口取值：2018-07-17。 |
| Region | 否 | String | [公共参数](/document/api/266/31756)，此参数为可选参数。 |
| SubAppId | 是 | Integer | <p><b>点播<a href="/document/product/266/14574">应用</a> ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b></p><br/>示例值：251007502 |
| ModelName | 是 | String | <p>模型名称。取值：<br>Kling：可灵；<br>Vidu；<br>Hailuo：海螺；<br>Hunyuan：混元；<br>Mingmou：明眸；<br>GV；<br>OS；<br>PixVerse;</p><br/>示例值：Kling |
| ModelVersion | 是 | String | <p>模型版本。取值：<br>当 ModelName 是 Hailuo，可选值为 02、2.3、2.3-fast、H3、H3_regen、H3-Max；<br>当 ModelName 是 Kling，可选值为 1.6、2.0、2.1、2.5、2.6、O1、3.0、3.0-Omni、3.0-turbo、scene；<br>当 ModelName 是 Vidu，可选值为 q2、q2-pro、q2-turbo、q3、q3-pro、q3-turbo、q3-mix、q3-drama、q3-ad、scene、avatar-q2-pro、avatar-q2-turbo、lip-sync；<br>当 ModelName 是 GV，可选值为 3.1、3.1-fast、3.1-lite、omni；<br>当 ModelName 是 OS，可选值为 2.0；<br>当 ModelName 是 Hunyuan，可选值为 1.5、3d_2.0；<br>当 ModelName 是 Mingmou，可选值为 1.0；<br>当 ModelName 是 PixVerse，可选值为 v5.6、v6、c1；</p><br/>示例值：O1 |
| FileInfos.N | 否 | Array of [AigcVideoTaskInputFileInfo](/document/api/266/31773#AigcVideoTaskInputFileInfo) | <p>用于描述模型在生成视频时要使用的资源文件，分为<strong>首尾帧模式、参考图片/视频/声音生成、视频编辑等模式</strong>。</p><p><strong>首尾帧视频生成</strong>：<strong>首帧图片的Usage字段为FirstFrame，尾帧图片的Usage字段为LastFrame</strong>，支持各一张，可以单独传首帧，不能单独传尾帧。<strong>首尾帧生成会参考图片比例</strong>。<br><strong>参考图片/视频/声音生成</strong>：可传入单个或者多个图片/视频/声音作为参考，<strong>Usage字段为Reference</strong>；<strong>参考模式，可以调整生成视频的宽高比例</strong>。<br><strong>视频编辑</strong>：Vidu、Kling可输入视频进行编辑。传入视频的同时也可以传入图片，<strong>图片的Usage字段为Reference</strong>。</p><p>注意：</p><ol><li>图片大小不超过10M。</li><li>支持的图片格式：jpeg、jpg、png。x0b</li><li>关于模型某个版本是否支持参考图片/视频/声音、首尾帧、视频编辑等功能，可向我们索取文档或者参考原厂文档信息。</li></ol> |
| SubjectInfos.N | 否 | Array of [AigcVideoTaskInputSubjectInfo](/document/api/266/31773#AigcVideoTaskInputSubjectInfo) | <p>主体输入信息。</p> |
| LastFrameFileId | 否 | String | <p>用于作为尾帧画面来生成视频的媒体文件 ID。该文件在云点播上的全局唯一标识符，在上传成功后由云点播后台分配。可以在 <a href="/document/product/266/7830">视频上传完成事件通知</a> 或 <a href="https://console.cloud.tencent.com/vod/media">云点播控制台</a> 获取该字段。</p><ol><li>指定该参数时，须同时通过 FileInfos 指定首帧画面。</li><li>图片大小需小于10M。</li><li>图片格式的取值为：jpeg，jpg, png, webp。</li></ol><br/>示例值：3704211***509911 |
| LastFrameUrl | 否 | String | <p>用于作为尾帧画面来生成视频的媒体文件 URL。说明：</p><ol><li>指定该参数时，须同时通过 FileInfos 指定首帧画面。</li><li>图片大小需小于5M。</li><li>图片格式的取值为：jpeg，jpg, png, webp。</li></ol><br/>示例值：https://test.com/1.png |
| Prompt | 否 | String | <p>生成视频的提示词。<br>当未传入参考文件，没有使用场景类型，ExtInfo不为空，Prompt 为必填。</p><br/>示例值：generate a car |
| NegativePrompt | 否 | String | <p>要阻止模型生成视频的提示词。</p><br/>示例值：red |
| EnhancePrompt | 否 | String | <p>是否自动优化提示词。开启时将自动优化传入的 Prompt，以提升生成质量。取值有： <li>Enabled：开启；</li> <li>Disabled：关闭；</li></p><br/>示例值：Enabled |
| OutputConfig | 否 | [AigcVideoOutputConfig](/document/api/266/31773#AigcVideoOutputConfig) | <p>生视频任务的输出媒体文件配置。</p> |
| InputRegion | 否 | String | <p>输入文件的区域信息。取值如下：</p><ul><li>当输入文件存储在海外时：Oversea；</li><li>当输入文件存储在美西时：OverseaUSWest；</li><li>当输入文件存储在国内时：Mainland。</li></ul><p>默认值：Mainland</p><br/>示例值：Mainland |
| SceneType | 否 | String | <p>场景类型。取值如下：</p><li>当 ModelName 为 Kling 时：    motion_control 表示动作控制；    avatar_i2v 表示数字人；    lip_sync 表示对口型；</li><li>当 ModelName 为 Vidu 时：    template_effect 表示特效模板；</li><li>其他 ModelName 暂不支持。</li><br/>示例值：motion_control |
| Procedure | 否 | String | <p><a href="https://cloud.tencent.com/document/product/266/33475#.E4.BB.BB.E5.8A.A1.E6.B5.81">任务流名称</a>，在需要对生成的新视频执行任务流时填写。</p><br/>示例值：testProcedure |
| Seed | 否 | Integer | <p>模型随机种子。</p><br/>示例值：123 |
| SessionId | 否 | String | <p>用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。</p><br/>示例值：mysession2 |
| SessionContext | 否 | String | <p>来源上下文，用于透传用户请求信息，音画质重生完成回调将返回该字段值，最长 1000 个字符。</p><br/>示例值：mysessionContext |
| TasksPriority | 否 | Integer | <p>任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。</p><br/>示例值：10 |
| ExtInfo | 否 | String | <p>保留字段，特殊用途时使用。<br>可用于传入模型特殊参数、分镜prompt等</p><br/>示例值：myextinfo |

## 3. 输出参数

| 参数名称 | 类型 | 描述 |
|---------|---------|---------|
| TaskId | String | <p>任务 ID。</p><br/>示例值：251007502-AigcVideo***25dacdcef7dd2b20fdt|
| RequestId | String | 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。|

## 4. 示例

### 示例1 发起生成 AIGC 视频的任务

#### 输入示例

```
POST / HTTP/1.1
Host: vod.tencentcloudapi.com
Content-Type: application/json
X-TC-Action: CreateAigcVideoTask
<公共请求参数>

{
    "SubAppId": 251007502,
    "ModelName": "GV",
    "ModelVersion": "3.1-fast",
    "FileInfos": [
        {
            "FileId": "3704211***509819"
        }
    ],
    "LastFrameFileId": "3704211***509911",
    "Prompt": "generate a car",
    "NegativePrompt": "red",
    "EnhancePrompt": "Enabled",
    "OutputConfig": {
        "StorageMode": "Temporary",
        "AspectRatio": "9:16",
        "AudioGeneration": "Enabled",
        "PersonGeneration": "AllowAdult",
        "InputComplianceCheck": "Enabled",
        "OutputComplianceCheck": "Enabled"
    },
    "SessionId": "mysession2",
    "SessionContext": "mysessionContext",
    "TasksPriority": "10",
    "ExtInfo": "myextinfo"
}
```

#### 输出示例

```json
{
    "Response": {
        "TaskId": "251007502-AigcVideo***25dacdcef7dd2b20fdt",
        "RequestId": "d68920a4-c989-4afe-ac4d-2f06de99368e"
    }
}
```


## 5. 开发者资源

### 腾讯云 API 平台

[腾讯云 API 平台](https://cloud.tencent.com/api) 是综合 API 文档、错误码、API Explorer 及 SDK 等资源的统一查询平台，方便您从同一入口查询及使用腾讯云提供的所有 API 服务。

### API Inspector

用户可通过 [API Inspector](https://cloud.tencent.com/document/product/1278/49361) 查看控制台每一步操作关联的 API 调用情况，并自动生成各语言版本的 API 代码，也可前往 [API Explorer](https://cloud.tencent.com/document/product/1278/46697) 进行在线调试。

### SDK

云 API 3.0 提供了配套的开发工具集（SDK），支持多种编程语言，能更方便的调用 API。
* Tencent Cloud SDK 3.0 for Python: [CNB](https://cnb.cool/tencent/cloud/api/sdk/tencentcloud-sdk-python/-/blob/master/tencentcloud/vod/v20180717/vod_client.py), [GitHub](https://github.com/TencentCloud/tencentcloud-sdk-python/blob/master/tencentcloud/vod/v20180717/vod_client.py), [Gitee](https://gitee.com/TencentCloud/tencentcloud-sdk-python/blob/master/tencentcloud/vod/v20180717/vod_client.py)
* Tencent Cloud SDK 3.0 for Java: [CNB](https://cnb.cool/tencent/cloud/api/sdk/tencentcloud-sdk-java/-/blob/master/src/main/java/com/tencentcloudapi/vod/v20180717/VodClient.java), [GitHub](https://github.com/TencentCloud/tencentcloud-sdk-java/blob/master/src/main/java/com/tencentcloudapi/vod/v20180717/VodClient.java), [Gitee](https://gitee.com/TencentCloud/tencentcloud-sdk-java/blob/master/src/main/java/com/tencentcloudapi/vod/v20180717/VodClient.java)
* Tencent Cloud SDK 3.0 for PHP: [CNB](https://cnb.cool/tencent/cloud/api/sdk/tencentcloud-sdk-php/-/blob/master/src/TencentCloud/Vod/V20180717/VodClient.php), [GitHub](https://github.com/TencentCloud/tencentcloud-sdk-php/blob/master/src/TencentCloud/Vod/V20180717/VodClient.php), [Gitee](https://gitee.com/TencentCloud/tencentcloud-sdk-php/blob/master/src/TencentCloud/Vod/V20180717/VodClient.php)
* Tencent Cloud SDK 3.0 for Go: [CNB](https://cnb.cool/tencent/cloud/api/sdk/tencentcloud-sdk-go/-/blob/master/tencentcloud/vod/v20180717/client.go), [GitHub](https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/vod/v20180717/client.go), [Gitee](https://gitee.com/TencentCloud/tencentcloud-sdk-go/blob/master/tencentcloud/vod/v20180717/client.go)
* Tencent Cloud SDK 3.0 for Node.js: [CNB](https://cnb.cool/tencent/cloud/api/sdk/tencentcloud-sdk-nodejs/-/blob/master/src/services/vod/v20180717/vod_client.ts), [GitHub](https://github.com/TencentCloud/tencentcloud-sdk-nodejs/blob/master/src/services/vod/v20180717/vod_client.ts), [Gitee](https://gitee.com/TencentCloud/tencentcloud-sdk-nodejs/blob/master/src/services/vod/v20180717/vod_client.ts)
* Tencent Cloud SDK 3.0 for .NET: [CNB](https://cnb.cool/tencent/cloud/api/sdk/tencentcloud-sdk-dotnet/-/blob/master/TencentCloud/Vod/V20180717/VodClient.cs), [GitHub](https://github.com/TencentCloud/tencentcloud-sdk-dotnet/blob/master/TencentCloud/Vod/V20180717/VodClient.cs), [Gitee](https://gitee.com/TencentCloud/tencentcloud-sdk-dotnet/blob/master/TencentCloud/Vod/V20180717/VodClient.cs)
* Tencent Cloud SDK 3.0 for C++: [CNB](https://cnb.cool/tencent/cloud/api/sdk/tencentcloud-sdk-cpp/-/blob/master/vod/src/v20180717/VodClient.cpp), [GitHub](https://github.com/TencentCloud/tencentcloud-sdk-cpp/blob/master/vod/src/v20180717/VodClient.cpp), [Gitee](https://gitee.com/TencentCloud/tencentcloud-sdk-cpp/blob/master/vod/src/v20180717/VodClient.cpp)
* Tencent Cloud SDK 3.0 for Ruby: [CNB](https://cnb.cool/tencent/cloud/api/sdk/tencentcloud-sdk-ruby/-/blob/master/tencentcloud-sdk-vod/lib/v20180717/client.rb), [GitHub](https://github.com/TencentCloud/tencentcloud-sdk-ruby/blob/master/tencentcloud-sdk-vod/lib/v20180717/client.rb), [Gitee](https://gitee.com/TencentCloud/tencentcloud-sdk-ruby/blob/master/tencentcloud-sdk-vod/lib/v20180717/client.rb)

### 命令行工具

* [Tencent Cloud CLI 3.0](https://cloud.tencent.com/document/product/440/6176)

## 6. 错误码

以下仅列出了接口业务逻辑相关的错误码，其他错误码详见 [公共错误码](/document/api/266/31774#.E5.85.AC.E5.85.B1.E9.94.99.E8.AF.AF.E7.A0.81)。

| 错误码 | 描述 |
|---------|---------|
| FailedOperation | 操作失败。 |
| FailedOperation.InvalidVodUser | 没有开通点播业务。 |
| InternalError | 内部错误。 |
| InvalidParameter | 参数错误。 |
| InvalidParameterValue.FileId | FileId 不存在。 |
| InvalidParameterValue.SessionContextTooLong | SessionContext 过长。 |
| InvalidParameterValue.SessionIdTooLong | SessionId 过长。 |
| InvalidParameterValue.SubAppId | 参数值错误：应用 ID。 |
| LimitExceeded.Quota | 超出配额限制 |
| UnauthorizedOperation | 未授权操作。 |
