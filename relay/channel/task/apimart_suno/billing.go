package apimart_suno

import (
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
)

// maxModeModels 是价格清单中支持 Max 档（2 倍计费）的模型。
// variety:"max" 不会触发本倍率；仅显式 max_mode=true 才翻倍。
var maxModeModels = map[string]struct{}{
	"suno-music":            {},
	"suno-extend":           {},
	"suno-cover":            {},
	"suno-sample":           {},
	"suno-mashup":           {},
	"suno-add-instrumental": {},
	"suno-add-stem":         {},
	"suno-add-vocals":       {},
	"suno-replace-section":  {},
	"suno-inspo":            {},
	"suno-upload-cover":     {},
	"suno-upload-extend":    {},
}

// EstimateBilling applies a 2× multiplier when max_mode=true on a Max-capable model.
// Other tools charge a fixed per-request fee; no extra OtherRatios are needed.
func EstimateBilling(c *gin.Context, info *relaycommon.RelayInfo) map[string]float64 {
	if info == nil {
		return nil
	}
	if _, ok := maxModeModels[info.OriginModelName]; !ok {
		return nil
	}
	v, ok := c.Get("task_request")
	if !ok {
		return nil
	}
	req, ok := v.(*dto.APIMartSunoRequest)
	if !ok {
		return nil
	}
	if req.MaxMode != nil && *req.MaxMode {
		return map[string]float64{"max_mode": 2.0}
	}
	return nil
}

// AdjustBillingOnSubmit returns nil: APIMart Suno charges a fixed per-request fee;
// there is nothing to re-estimate after the upstream submission response.
func AdjustBillingOnSubmit(_ *relaycommon.RelayInfo, _ []byte) map[string]float64 {
	return nil
}

// AdjustBillingOnComplete returns 0: keep the pre-charged amount unchanged.
// APIMart refunds automatically on failure (handled by the platform's sweepTimedOutTasks
// and RefundTaskQuota path when the task reaches FAILURE status).
func AdjustBillingOnComplete(_ *model.Task, _ *relaycommon.TaskInfo) int {
	return 0
}
