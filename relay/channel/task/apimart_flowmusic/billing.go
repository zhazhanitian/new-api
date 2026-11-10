package apimart_flowmusic

import (
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
)

// EstimateBilling returns nil. Flow Music is a fixed per-request price.
// Lyria 3 Pro and Lyria 3.5 share the same price, so version does not change the charge.
func EstimateBilling(_ *gin.Context, _ *relaycommon.RelayInfo) map[string]float64 {
	return nil
}

// AdjustBillingOnSubmit returns nil. The pre-charge is the final charge.
func AdjustBillingOnSubmit(_ *relaycommon.RelayInfo, _ []byte) map[string]float64 {
	return nil
}

// AdjustBillingOnComplete returns 0. Failure refunds go through RefundTaskQuota,
// not a complete-time price adjustment.
func AdjustBillingOnComplete(_ *model.Task, _ *relaycommon.TaskInfo) int {
	return 0
}
