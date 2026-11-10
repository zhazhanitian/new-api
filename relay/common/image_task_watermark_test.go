package common

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	apicommon "github.com/QuantumNous/new-api/common"
	"github.com/gin-gonic/gin"
)

func TestValidateImageTaskRequestKeepsWatermarkFalse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	body := []byte(`{
		"model": "doubao-seedream-5-0-260128",
		"prompt": "仙侠大战",
		"resolution": "2K",
		"output_format": "png",
		"n": 1,
		"watermark": false
	}`)
	c.Request = httptest.NewRequest(http.MethodPost, "/v2/image-tasks", bytes.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	if err := apicommon.ReplaceRequestBody(c, body); err != nil {
		t.Fatal(err)
	}

	info := &RelayInfo{TaskRelayInfo: &TaskRelayInfo{}}
	if taskErr := ValidateImageTaskRequest(c, info); taskErr != nil {
		t.Fatalf("validate: %s", taskErr.Message)
	}
	req, err := GetTaskRequest(c)
	if err != nil {
		t.Fatal(err)
	}
	if req.Watermark == nil || *req.Watermark {
		t.Fatalf("watermark = %v, want false", req.Watermark)
	}
	if req.Resolution != "2K" {
		t.Fatalf("resolution = %q, want 2K", req.Resolution)
	}
}
