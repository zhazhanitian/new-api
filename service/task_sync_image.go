package service

import (
	"strconv"
	"time"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
)

// RecordSyncImageTask 为同步生图路径（POST /v1/images/generations）写入任务记录，
// 使其与异步任务路径对称，在任务日志中可见。
//
// 调用时机：
//   - 成功：在 ImageHelper 末尾 apiErr=nil 调用（已结算，Quota 准确）
//   - 失败：在 Relay() 的 defer 内、所有重试耗尽后 apiErr!=nil 调用（已退款，Quota=0）
//
// Quota 说明：
//   - 成功：使用 info.PriceData.Quota（对按价格计费的生图模型完全准确）
//   - 失败：写 0（用户已退款，无实际扣费）
func RecordSyncImageTask(c *gin.Context, info *relaycommon.RelayInfo, apiErr *types.NewAPIError) {
	if info == nil || info.ChannelMeta == nil {
		return
	}

	// platform 与 GetTaskPlatform 保持一致：渠道类型整数字符串
	platform := constant.TaskPlatform(strconv.Itoa(info.ChannelType))

	now := time.Now().Unix()
	submitTime := now
	if !info.StartTime.IsZero() {
		submitTime = info.StartTime.Unix()
	}

	task := model.InitTask(platform, info)
	task.Action = constant.TaskActionImageGenerate
	task.SubmitTime = submitTime
	task.StartTime = submitTime
	task.FinishTime = now

	task.PrivateData.BillingSource = info.BillingSource
	task.PrivateData.SubscriptionId = info.SubscriptionId
	task.PrivateData.TokenId = info.TokenId
	task.PrivateData.BillingContext = &model.TaskBillingContext{
		ModelPrice:      info.PriceData.ModelPrice,
		GroupRatio:      info.PriceData.GroupRatioInfo.GroupRatio,
		ModelRatio:      info.PriceData.ModelRatio,
		OtherRatios:     info.PriceData.OtherRatios,
		OriginModelName: info.OriginModelName,
		PerCallBilling:  info.PriceData.UsePrice,
	}

	if apiErr != nil {
		// 失败：已退款，不计费
		task.Status = model.TaskStatusFailure
		task.Progress = "0%"
		task.FailReason = apiErr.Error()
		task.Quota = 0
	} else {
		// 成功：使用预估/已结算额度
		task.Status = model.TaskStatusSuccess
		task.Progress = "100%"
		task.Quota = info.PriceData.Quota
	}

	if insertErr := task.Insert(); insertErr != nil {
		logger.LogWarn(c, "RecordSyncImageTask: failed to insert task record: "+insertErr.Error())
	}
}
