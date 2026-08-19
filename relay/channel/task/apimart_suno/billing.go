package apimart_suno

import (
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
)

// EstimateBilling returns nil: all versions of a tool currently share the same
// base model price defined in defaultModelPrice. No extra OtherRatios are needed.
//
// Future version-specific pricing: add a configurable VersionPriceMultiplier map
// (keyed by "modelID:version") and return {"version_ratio": multiplier} here.
func EstimateBilling(_ *gin.Context, _ *relaycommon.RelayInfo) map[string]float64 {
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
