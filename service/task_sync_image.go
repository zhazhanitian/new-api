package service

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/setting/operation_setting"
	"github.com/QuantumNous/new-api/types"
	"github.com/gin-gonic/gin"
)

const syncImageRecordKey = "sync_image_record"
const syncImageClientResponseKey = "sync_image_client_response"

type syncImageRecord struct {
	UpstreamBody []byte
	URLs         []string
}

// KeepSyncImageClientResponse 暂存尚未写出的同步生图响应。
// 金额要等结算完成后才能填，所以这里先不写给客户端。
func KeepSyncImageClientResponse(c *gin.Context, imageResponse *dto.ImageResponse) {
	if c == nil || imageResponse == nil {
		return
	}
	c.Set(syncImageClientResponseKey, imageResponse)
}

// WriteSyncImageClientResponse 在结算后把实扣金额写入响应并返回给客户端。
// 没有暂存响应时不做任何事，其它渠道仍由各自的 DoResponse 写出。
func WriteSyncImageClientResponse(c *gin.Context, quota int) bool {
	if c == nil {
		return false
	}
	raw, ok := c.Get(syncImageClientResponseKey)
	if !ok {
		return false
	}
	imageResponse, ok := raw.(*dto.ImageResponse)
	if !ok || imageResponse == nil {
		return false
	}
	usd := float64(quota) / common.QuotaPerUnit
	imageResponse.Amount = fmt.Sprintf("%.6f", usd*operation_setting.USDExchangeRate)
	body, err := common.Marshal(imageResponse)
	if err != nil {
		logger.LogWarn(c, "WriteSyncImageClientResponse: "+err.Error())
		return false
	}
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	_, _ = c.Writer.Write(body)
	return true
}

// KeepSyncImageRecord 保存同步生图拿到的上游原文和图片地址，供本路径写任务日志。
// 单图保留原文；多图改为图片地址列表。返回给调用方的响应不从这里取。
func KeepSyncImageRecord(c *gin.Context, upstreamBody []byte, urls []string) {
	if c == nil {
		return
	}
	c.Set(syncImageRecordKey, &syncImageRecord{
		UpstreamBody: append([]byte(nil), upstreamBody...),
		URLs:         append([]string(nil), urls...),
	})
}

// RecordSyncImageTask 为同步生图路径（POST /v1/images/generations）写入任务记录，
// 使其与异步任务路径对称，在任务日志中可见。
//
// 调用时机：
//   - 成功：在 ImageHelper 末尾 apiErr=nil 调用（已结算，Quota 准确）
//   - 失败：在 Relay() 的 defer 内、所有重试耗尽后 apiErr!=nil 调用（已退款，Quota=0）
//
// Quota 说明：
//   - 成功：使用 info.PriceData.Quota。阶梯计费会在 PostTextConsumeQuota 里把实扣写回该字段。
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
		task.Status = model.TaskStatusSuccess
		task.Progress = "100%"
		task.Quota = info.PriceData.Quota
		if raw, ok := c.Get(syncImageRecordKey); ok {
			if rec, ok := raw.(*syncImageRecord); ok {
				applySyncImageRecord(task, rec)
			}
		}
	}

	if insertErr := task.Insert(); insertErr != nil {
		logger.LogWarn(c, "RecordSyncImageTask: failed to insert task record: "+insertErr.Error())
	}
}

func applySyncImageRecord(task *model.Task, rec *syncImageRecord) {
	if task == nil || rec == nil {
		return
	}
	urls := make([]string, 0, len(rec.URLs))
	for _, raw := range rec.URLs {
		url := strings.TrimSpace(raw)
		if url == "" {
			continue
		}
		urls = append(urls, url)
		if task.PrivateData.ResultURL == "" && (strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")) {
			task.PrivateData.ResultURL = url
		}
	}
	if len(rec.UpstreamBody) > 0 {
		task.Data = redactVideoResponseBody(rec.UpstreamBody)
	}
	if len(urls) <= 1 {
		return
	}
	images := make([]dto.ImageData, 0, len(urls))
	for _, url := range urls {
		images = append(images, dto.ImageData{Url: url})
	}
	if payload, err := common.Marshal(dto.ImageResponse{Data: images}); err == nil {
		task.Data = payload
	}
}
