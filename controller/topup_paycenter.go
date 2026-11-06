package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/service/paycenter"
	"github.com/QuantumNous/new-api/setting"
	"github.com/QuantumNous/new-api/setting/operation_setting"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

const (
	payCenterExpireMinutes = 30
	payCenterSubject       = "额度充值"
	payCenterNotifyMaxBody = 1 << 20
)

// payCenterPayRequest 在人民币展示下，amount 是用户要支付的整数元。
type payCenterPayRequest struct {
	Amount int64 `json:"amount"`
}

type payCenterNotify struct {
	Event           string `json:"event"`
	OrderNo         string `json:"orderNo"`
	BusinessOrderNo string `json:"businessOrderNo"`
	ProductCode     string `json:"productCode"`
	Amount          int64  `json:"amount"`
	Status          string `json:"status"`
	RefundNo        string `json:"refundNo"`
	RefundAmount    int64  `json:"refundAmount"`
	Timestamp       int64  `json:"timestamp"`
}

func payCenterConfig() paycenter.Config {
	return paycenter.Config{
		BaseURL: setting.PayCenterBaseURL,
		APIKey:  setting.PayCenterApiKey,
	}
}

// payCenterPayableYuan 计算支付中心实付金额。
// 人民币展示时，用户选择的就是整数元，不再把美元数量乘单价。
func payCenterPayableYuan(amount int64, group string) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("充值金额过低")
	}
	if !operation_setting.IsCNYDisplay() {
		return getPayMoneyWithPrice(amount, group, setting.PayCenterUnitPrice), nil
	}
	if setting.PayCenterUnitPrice <= 0 {
		return 0, errors.New("支付中心单价未配置")
	}
	discount := 1.0
	if ds, ok := operation_setting.GetPaymentSetting().AmountDiscount[int(amount)]; ok && ds > 0 {
		discount = ds
	}
	groupRatio := common.GetTopupGroupRatio(group)
	if groupRatio == 0 {
		groupRatio = 1
	}
	money := decimal.NewFromInt(amount).
		Mul(decimal.NewFromFloat(discount)).
		Mul(decimal.NewFromFloat(groupRatio))
	return money.InexactFloat64(), nil
}

func RequestPayCenterAmount(c *gin.Context) {
	if !isPayCenterTopUpEnabled() {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "当前管理员未配置支付中心"})
		return
	}
	var req payCenterPayRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "参数错误"})
		return
	}
	if req.Amount < getPayCenterMinTopup() {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": fmt.Sprintf("充值数量不能小于 %d", getPayCenterMinTopup())})
		return
	}
	group, err := model.GetUserGroup(c.GetInt("id"), true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "获取用户分组失败"})
		return
	}
	payYuan, err := payCenterPayableYuan(req.Amount, group)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": err.Error()})
		return
	}
	amountFen := paycenter.YuanToFen(payYuan)
	if amountFen < 1 {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "充值金额过低"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    strconv.FormatFloat(paycenter.FenToYuan(amountFen), 'f', 2, 64),
	})
}

func RequestPayCenterPay(c *gin.Context) {
	if !isPayCenterTopUpEnabled() {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "当前管理员未配置支付中心"})
		return
	}

	var req payCenterPayRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "参数错误"})
		return
	}
	if req.Amount < getPayCenterMinTopup() {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": fmt.Sprintf("充值数量不能小于 %d", getPayCenterMinTopup())})
		return
	}

	userID := c.GetInt("id")
	group, err := model.GetUserGroup(userID, true)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "获取用户分组失败"})
		return
	}
	payMoney, err := payCenterPayableYuan(req.Amount, group)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": err.Error()})
		return
	}
	amountFen := paycenter.YuanToFen(payMoney)
	if amountFen < 1 {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "充值金额过低"})
		return
	}
	money := paycenter.FenToYuan(amountFen)

	tradeNo := fmt.Sprintf("USR%dNO%s%d", userID, common.GetRandomString(6), time.Now().Unix())
	quotaAmount := req.Amount
	if operation_setting.GetQuotaDisplayType() == operation_setting.QuotaDisplayTypeTokens {
		dAmount := decimal.NewFromInt(quotaAmount)
		dQuotaPerUnit := decimal.NewFromFloat(common.QuotaPerUnit)
		quotaAmount = dAmount.Div(dQuotaPerUnit).IntPart()
	}

	topUp := &model.TopUp{
		UserId:          userID,
		Amount:          quotaAmount,
		Money:           money,
		TradeNo:         tradeNo,
		PaymentMethod:   "wxpay",
		PaymentProvider: model.PaymentProviderPayCenter,
		CreateTime:      time.Now().Unix(),
		Status:          common.TopUpStatusPending,
	}
	if err := topUp.Insert(); err != nil {
		logger.LogError(c.Request.Context(), fmt.Sprintf("支付中心 创建充值订单失败 user_id=%d trade_no=%s amount_fen=%d error=%q", userID, tradeNo, amountFen, err.Error()))
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "创建订单失败"})
		return
	}

	result, err := payCenterConfig().CreateNativeOrder(c.Request.Context(), paycenter.CreateNativeOrderRequest{
		BusinessOrderNo: tradeNo,
		AmountFen:       amountFen,
		Subject:         payCenterSubject,
		Attach:          fmt.Sprintf("userId=%d", userID),
		ExpireMinutes:   payCenterExpireMinutes,
	})
	if err != nil {
		var apiErr *paycenter.APIError
		if errors.As(err, &apiErr) {
			markPayCenterFailed(c, tradeNo, userID)
			logger.LogError(c.Request.Context(), fmt.Sprintf("支付中心 下单被拒绝 user_id=%d trade_no=%s amount_fen=%d error=%q", userID, tradeNo, amountFen, apiErr.Error()))
			c.JSON(http.StatusOK, gin.H{"message": "error", "data": apiErr.Error()})
			return
		}
		logger.LogError(c.Request.Context(), fmt.Sprintf("支付中心 下单无响应 user_id=%d trade_no=%s amount_fen=%d error=%q", userID, tradeNo, amountFen, err.Error()))
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "支付中心没有返回结果，请稍后重试"})
		return
	}
	if result.CodeURL == "" || (result.Amount != 0 && result.Amount != amountFen) {
		markPayCenterFailed(c, tradeNo, userID)
		logger.LogError(c.Request.Context(), fmt.Sprintf("支付中心 下单结果无效 user_id=%d trade_no=%s pay_order_no=%s amount_fen=%d response_amount=%d", userID, tradeNo, result.OrderNo, amountFen, result.Amount))
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "拉起支付失败"})
		return
	}

	logger.LogInfo(c.Request.Context(), fmt.Sprintf("支付中心 充值订单创建成功 user_id=%d trade_no=%s pay_order_no=%s amount_fen=%d", userID, tradeNo, result.OrderNo, amountFen))
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"code_url": result.CodeURL,
			"trade_no": tradeNo,
		},
	})
}

func GetPayCenterOrder(c *gin.Context) {
	tradeNo := strings.TrimSpace(c.Param("trade_no"))
	userID := c.GetInt("id")
	topUp := payCenterOrderForUser(tradeNo, userID)
	if topUp == nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "订单不存在"})
		return
	}
	status, stop := syncPayCenterOrder(c, topUp)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"status":   status,
			"trade_no": topUp.TradeNo,
			"stop":     stop,
		},
	})
}

func ClosePayCenterOrder(c *gin.Context) {
	tradeNo := strings.TrimSpace(c.Param("trade_no"))
	userID := c.GetInt("id")
	topUp := payCenterOrderForUser(tradeNo, userID)
	if topUp == nil {
		c.JSON(http.StatusOK, gin.H{"message": "error", "data": "订单不存在"})
		return
	}
	if topUp.Status == common.TopUpStatusPending && isPayCenterTopUpEnabled() {
		order, err := payCenterConfig().QueryByBusinessOrderNo(c.Request.Context(), topUp.TradeNo)
		if err != nil {
			logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 关单前查单失败 trade_no=%s error=%q", topUp.TradeNo, err.Error()))
		} else if order != nil {
			switch order.Status {
			case "SUCCESS":
				if err := model.CompletePayCenterTopUp(topUp.TradeNo, order.Amount, c.ClientIP()); err != nil {
					logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 关单时入账失败 trade_no=%s pay_order_no=%s error=%q", topUp.TradeNo, order.OrderNo, err.Error()))
				}
			case "CREATED", "PAYING":
				if order.OrderNo != "" {
					if err := payCenterConfig().CloseOrder(c.Request.Context(), order.OrderNo); err != nil {
						logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 关单失败 trade_no=%s pay_order_no=%s error=%q", topUp.TradeNo, order.OrderNo, err.Error()))
					}
				}
			case "FAILED", "REFUNDING", "REFUNDED":
				_ = model.MarkPayCenterTopUpIfPending(topUp.TradeNo, userID, common.TopUpStatusFailed)
			case "CLOSED":
				_ = model.MarkPayCenterTopUpIfPending(topUp.TradeNo, userID, common.TopUpStatusExpired)
			}
		}
	}
	status := topUp.Status
	if fresh := model.GetTopUpByTradeNo(topUp.TradeNo); fresh != nil {
		status = fresh.Status
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": gin.H{
			"status":   status,
			"trade_no": topUp.TradeNo,
		},
	})
}

func PayCenterNotify(c *gin.Context) {
	if !isPayCenterTopUpEnabled() {
		c.String(http.StatusUnauthorized, "disabled")
		return
	}
	raw, err := io.ReadAll(io.LimitReader(c.Request.Body, payCenterNotifyMaxBody))
	if err != nil {
		c.String(http.StatusBadRequest, "invalid body")
		return
	}
	if !paycenter.VerifyNotifySignature(raw, setting.PayCenterNotifyKey, c.GetHeader("X-Notify-Signature")) {
		logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 回调验签失败 client_ip=%s", c.ClientIP()))
		c.String(http.StatusUnauthorized, "invalid signature")
		return
	}

	var notify payCenterNotify
	if err := json.Unmarshal(raw, &notify); err != nil {
		c.String(http.StatusBadRequest, "invalid body")
		return
	}
	headerTS, _ := strconv.ParseInt(strings.TrimSpace(c.GetHeader("X-Notify-Timestamp")), 10, 64)
	if !paycenter.TimestampWithinWindow(headerTS, notify.Timestamp, time.Now().Unix(), paycenter.NotifyTimeWindowSeconds) {
		logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 回调时间戳无效 event=%s business_order_no=%s", notify.Event, notify.BusinessOrderNo))
		c.String(http.StatusUnauthorized, "invalid timestamp")
		return
	}
	if strings.TrimSpace(notify.ProductCode) != strings.TrimSpace(setting.PayCenterProductCode) {
		logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 回调产品编码不匹配 event=%s product_code=%s business_order_no=%s", notify.Event, notify.ProductCode, notify.BusinessOrderNo))
		c.String(http.StatusUnauthorized, "product mismatch")
		return
	}

	switch notify.Event {
	case "payment.success":
		handlePayCenterPaymentNotify(c, &notify)
	case "refund.success":
		logger.LogInfo(c.Request.Context(), fmt.Sprintf("支付中心 收到退款成功通知 order_no=%s business_order_no=%s refund_no=%s refund_amount=%d", notify.OrderNo, notify.BusinessOrderNo, notify.RefundNo, notify.RefundAmount))
		c.String(http.StatusOK, "success")
	default:
		logger.LogInfo(c.Request.Context(), fmt.Sprintf("支付中心 忽略未知通知 event=%s business_order_no=%s", notify.Event, notify.BusinessOrderNo))
		c.String(http.StatusOK, "success")
	}
}

func handlePayCenterPaymentNotify(c *gin.Context, notify *payCenterNotify) {
	if notify.BusinessOrderNo == "" || notify.Status != "SUCCESS" {
		logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 支付通知缺少单号或状态不是成功 order_no=%s business_order_no=%s status=%s", notify.OrderNo, notify.BusinessOrderNo, notify.Status))
		c.String(http.StatusOK, "success")
		return
	}
	err := model.CompletePayCenterTopUp(notify.BusinessOrderNo, notify.Amount, c.ClientIP())
	if err == nil {
		c.String(http.StatusOK, "success")
		return
	}
	if errors.Is(err, model.ErrTopUpNotFound) {
		logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 支付通知订单不存在 business_order_no=%s pay_order_no=%s", notify.BusinessOrderNo, notify.OrderNo))
		c.String(http.StatusNotFound, "order not found")
		return
	}
	logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 支付通知未入账 business_order_no=%s pay_order_no=%s amount=%d error=%q", notify.BusinessOrderNo, notify.OrderNo, notify.Amount, err.Error()))
	c.String(http.StatusOK, "success")
}

func payCenterOrderForUser(tradeNo string, userID int) *model.TopUp {
	if tradeNo == "" || userID <= 0 {
		return nil
	}
	topUp := model.GetTopUpByTradeNo(tradeNo)
	if topUp == nil || topUp.UserId != userID || topUp.PaymentProvider != model.PaymentProviderPayCenter {
		return nil
	}
	return topUp
}

func syncPayCenterOrder(c *gin.Context, topUp *model.TopUp) (status string, stop bool) {
	if topUp.Status != common.TopUpStatusPending || !isPayCenterTopUpEnabled() {
		return topUp.Status, topUp.Status != common.TopUpStatusPending
	}
	order, err := payCenterConfig().QueryByBusinessOrderNo(c.Request.Context(), topUp.TradeNo)
	if err != nil || order == nil {
		if err != nil {
			logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 查单失败 trade_no=%s error=%q", topUp.TradeNo, err.Error()))
		}
		return common.TopUpStatusPending, false
	}

	switch order.Status {
	case "SUCCESS":
		err := model.CompletePayCenterTopUp(topUp.TradeNo, order.Amount, c.ClientIP())
		if err == nil {
			return common.TopUpStatusSuccess, true
		}
		if errors.Is(err, model.ErrPayCenterAmountMismatch) {
			logger.LogWarn(c.Request.Context(), fmt.Sprintf("支付中心 查单金额不一致 trade_no=%s amount=%d", topUp.TradeNo, order.Amount))
			return common.TopUpStatusFailed, true
		}
		fresh := model.GetTopUpByTradeNo(topUp.TradeNo)
		if fresh != nil {
			return fresh.Status, fresh.Status != common.TopUpStatusPending
		}
		return common.TopUpStatusPending, false
	case "CREATED", "PAYING":
		return common.TopUpStatusPending, false
	case "CLOSED":
		_ = model.MarkPayCenterTopUpIfPending(topUp.TradeNo, topUp.UserId, common.TopUpStatusExpired)
		return common.TopUpStatusExpired, true
	case "FAILED", "REFUNDING", "REFUNDED":
		_ = model.MarkPayCenterTopUpIfPending(topUp.TradeNo, topUp.UserId, common.TopUpStatusFailed)
		return common.TopUpStatusFailed, true
	default:
		return common.TopUpStatusPending, false
	}
}

func markPayCenterFailed(c *gin.Context, tradeNo string, userID int) {
	if err := model.MarkPayCenterTopUpIfPending(tradeNo, userID, common.TopUpStatusFailed); err != nil {
		logger.LogError(c.Request.Context(), fmt.Sprintf("支付中心 标记订单失败 trade_no=%s error=%q", tradeNo, err.Error()))
	}
}
