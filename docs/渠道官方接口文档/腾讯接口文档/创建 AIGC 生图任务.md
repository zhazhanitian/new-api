## 1. 接口描述

接口请求域名： vod.tencentcloudapi.com 。

该接口用于[生成 AIGC 图片](https://cloud.tencent.com/document/product/266/124473)。默认限制1个并发处理，接口调用会产生实际费用，请参考点播 [AIGC 生图片计费文档](https://cloud.tencent.com/document/product/266/95125#9c4dc6ff-4b3f-4b25-bf2d-393889dfb9ac)。该功能结算模式为[后付费](https://cloud.tencent.com/document/product/266/2838)，日结客户当天使用将在第二天出账，月结客户将在次月1日统一出上月使用费用。

默认接口请求频率限制：20次/秒。

<div class="rno-api-explorer">
    <div class="rno-api-explorer-inner">
        <div class="rno-api-explorer-hd">
            <div class="rno-api-explorer-title">
                推荐使用 API Explorer
            </div>
            <a href="https://console.cloud.tencent.com/api/explorer?Product=vod&Version=2018-07-17&Action=CreateAigcImageTask" class="rno-api-explorer-btn" hotrep="doc.api.explorerbtn"><i class="rno-icon-explorer"></i>点击调试</a>
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
| Action | 是 | String | [公共参数](/document/api/266/31756)，本接口取值：CreateAigcImageTask。 |
| Version | 是 | String | [公共参数](/document/api/266/31756)，本接口取值：2018-07-17。 |
| Region | 否 | String | [公共参数](/document/api/266/31756)，此参数为可选参数。 |
| SubAppId | 是 | Integer | <p><b>点播<a href="https://cloud.tencent.com/document/product/266/14574">应用</a> ID。从2023年12月25日起开通点播的客户，如访问点播应用中的资源（无论是默认应用还是新创建的应用），必须将该字段填写为应用 ID。</b></p><br/>示例值：251007502 |
| ModelName | 是 | String | <p>模型名称。取值：</p><li>OG</li><li>GG</li><li>Hunyuan</li><li>Vidu</li><li>Kling</li><br/>示例值：Hunyuan |
| ModelVersion | 是 | String | <p>模型版本。取值：</p><li>当 ModelName 是 OG，可选值为 image2_low、image2_medium、image2_high；</li><li>当 ModelName 是 GG，可选值为 2.5、3.0、3.1、3.1-lite；</li><li>当 ModelName 是 Hunyuan，可选值为 3.0、3d_2.0、3.5-preview；</li><li>当 ModelName 是 Vidu，可选值为 q2；</li><li>当 ModelName 是 Kling，可选值为 2.1、3.0、3.0-Omni、O1、scene；</li><li>当 ModelName 是Mingmou，可选值为 1.0；</li><br/>示例值：3.0 |
| FileInfos.N | 否 | Array of [AigcImageTaskInputFileInfo](/document/api/266/31773#AigcImageTaskInputFileInfo) | <p>AIGC 生图任务的输入图片的文件信息。各模型支持最大参考图数量：</p><ul><li>GG 2.5： 3张；</li><li>GG 3.0：14张；</li><li>GG 3.1：14张；</li><li>Kling 2.1：4张；</li><li>Kling 3.0：1张；</li><li>Kling 3.0-Omni：10张；</li><li>Kling O1：10张；</li><li>Vidu q2：7张；</li><li>Hunyuan 3.0：3张；</li></ul> |
| Prompt | 否 | String | <p>生成图片的提示词。当 FileInfos 为空时，此参数必填。</p><br/>示例值：generate a car |
| NegativePrompt | 否 | String | <p>要阻止模型生成图片的提示词。</p><br/>示例值：red |
| EnhancePrompt | 否 | String | <p>是否自动优化提示词。开启时将自动优化传入的 Prompt，以提升生成质量。取值有： <li>Enabled：开启；</li> <li>Disabled：关闭；</li></p><br/>示例值：Enabled |
| OutputConfig | 否 | [AigcImageOutputConfig](/document/api/266/31773#AigcImageOutputConfig) | <p>生图任务的输出媒体文件配置。</p> |
| InputRegion | 否 | String | <p>输入的区域信息。可选值：</p><ul><li>Mainland：中国大陆；</li><li>Oversea：海外；</li><li>OverseaUSWest：海外-美西；</li></ul><br/>示例值：Mainland |
| SceneType | 否 | String | <p>场景类型。取值如下：</p><li>当 ModelName 为 Hunyuan 时：   3d_panorama 表示全景图；</li><li>当 ModelName 为 Kling 时：   image_expand 表示扩图；</li><li>其他 ModelName 暂不支持。</li><br/>示例值：3d_panorama |
| Seed | 否 | Integer | <p>模型随机种子。</p><br/>示例值：123 |
| SessionId | 否 | String | <p>用于去重的识别码，如果三天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。</p><br/>示例值：mysession |
| SessionContext | 否 | String | <p>来源上下文，用于透传用户请求信息，音画质重生完成回调将返回该字段值，最长 1000 个字符。</p><br/>示例值：mySessionContext |
| TasksPriority | 否 | Integer | <p>任务的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。</p><br/>示例值：10 |
| ExtInfo | 否 | String | <p>保留字段，特殊用途时使用。</p><ul><li><p>Hunyuan 3.0</p><ul><li>支持自由设置分辨率宽高，宽、高均在 [512, 2048] 像素范围内，宽高乘积 ≤ 1024x1024 像素。示例：<code>{&quot;AdditionalParameters&quot;: &quot;{\&quot;size\&quot;:\&quot;728x1024\&quot;}&quot;}</code></li></ul></li><li><p>OG</p><ul><li>支持自由设置分辨率宽高：<ul><li>计算像素大小，需要被16整除</li><li>总像素数必须至少为655,360，且不得超过 8,294,400</li><li>示例：<code>{&quot;AdditionalParameters&quot;: &quot;{\&quot;size\&quot;:\&quot;728x1024\&quot;}&quot;}</code></li></ul></li><li>支持设置透明图层：<ul><li>示例：<code>{&quot;AdditionalParameters&quot;: &quot;{\&quot;background\&quot;:\&quot;transparent\&quot;}&quot;}</code></li></ul></li></ul></li><li><p>Kling</p><ul><li>支持设置扩图参数，示例：<code>{&quot;AdditionalParameters&quot;:&quot;{\&quot;down_expansion_ratio\&quot;:0.2,\&quot;left_expansion_ratio\&quot;:0.3,\&quot;right_expansion_ratio\&quot;:0.4,\&quot;up_expansion_ratio\&quot;:0.1}}</code><ul><li>通用约束：<ul><li>取值范围：[0, 2]；</li><li>新图片整体面积不得超过原图片的 3 倍；</li><li>可以通过 <code>Prompt</code> 字段传入正向提示词。</li><li>示例说明：<ul><li>up_expansion_ratio：向上扩充范围，基于原图高度的倍数计算。若原图高 20，参数值为 0.1，则原图顶边距离新图顶边为 20 × 0.1 = 2，该区域为扩图范围。</li><li>down_expansion_ratio：向下扩充范围，基于原图高度的倍数计算。若原图高 20，参数值为 0.2，则原图底边距离新图底边为 20 × 0.2 = 4，该区域为扩图范围。</li><li>left_expansion_ratio：向左扩充范围，基于原图宽度的倍数计算。若原图宽 30，参数值为 0.3，则原图左边距离新图左边为 30 × 0.3 = 9，该区域为扩图范围。</li><li>right_expansion_ratio：向右扩充范围，基于原图宽度的倍数计算。若原图宽 30，参数值为 0.4，则原图右边距离新图右边为 30 × 0.4 = 12，该区域为扩图范围。</li></ul></li></ul></li></ul></li></ul></li></ul><br/>示例值：{"AdditionalParameters": ""} |

## 3. 输出参数

| 参数名称 | 类型 | 描述 |
|---------|---------|---------|
| TaskId | String | <p>任务 ID。</p><br/>示例值：251007502-AigcImage***2782aff1e896673f1ft|
| RequestId | String | 唯一请求 ID，由服务端生成，每次请求都会返回（若请求因其他原因未能抵达服务端，则该次请求不会获得 RequestId）。定位问题时需要提供该次请求的 RequestId。|

## 4. 示例

### 示例1 发起生成 AIGC 图片的任务

#### 输入示例

```
POST / HTTP/1.1
Host: vod.tencentcloudapi.com
Content-Type: application/json
X-TC-Action: CreateAigcImageTask
<公共请求参数>

{
    "SubAppId": 251007502,
    "ModelName": "GEM",
    "ModelVersion": "2.5",
    "FileInfos": [
        {
            "FileId": "3704211***509819"
        }
    ],
    "Prompt": "generate a car",
    "NegativePrompt": "red",
    "EnhancePrompt": "Enabled",
    "OutputConfig": {
        "StorageMode": "Temporary",
        "AspectRatio": "16:9",
        "PersonGeneration": "AllowAdult",
        "InputComplianceCheck": "Enabled",
        "OutputComplianceCheck": "Enabled"
    },
    "SessionId": "mysession",
    "SessionContext": "mySessionContext",
    "TasksPriority": "10",
    "ExtInfo": "myextinfo"
}
```

#### 输出示例

```json
{
    "Response": {
        "TaskId": "251007502-AigcImage***2782aff1e896673f1ft",
        "RequestId": "f50d7667-72d8-46bb-a7e3-0613588971b6"
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
| InvalidParameterValue.SessionId | 去重识别码重复，请求被去重。 |
| InvalidParameterValue.SessionIdTooLong | SessionId 过长。 |
| InvalidParameterValue.SubAppId | 参数值错误：应用 ID。 |
| LimitExceeded.Quota | 超出配额限制 |
| UnauthorizedOperation | 未授权操作。 |
