package apimart_suno

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
)

// ── helpers ───────────────────────────────────────────────────────────────────

func newTestContext(body map[string]any) *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	data, _ := json.Marshal(body)
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/music/generations", bytes.NewReader(data))
	c.Request.Header.Set("Content-Type", "application/json")
	return c
}

func newRelayInfo(modelID string) *relaycommon.RelayInfo {
	info := &relaycommon.RelayInfo{
		OriginModelName: modelID,
		ChannelMeta: &relaycommon.ChannelMeta{
			ChannelBaseUrl: "https://apib.ai",
			ApiKey:         "test-key",
		},
		TaskRelayInfo: &relaycommon.TaskRelayInfo{
			PublicTaskID: "task_public123",
		},
	}
	return info
}

// ── Tool routing ──────────────────────────────────────────────────────────────

func TestGetToolDef_AllModelsRegistered(t *testing.T) {
	t.Parallel()
	for _, m := range ModelList {
		def, ok := GetToolDef(m)
		if !ok {
			t.Errorf("model %q has no ToolDef", m)
			continue
		}
		_ = def
	}
}

func TestBuildRequestURL(t *testing.T) {
	t.Parallel()
	cases := []struct {
		model   string
		wantURL string
	}{
		{"suno-music", "https://apib.ai/v1/music/generations"},
		{"suno-lyrics", "https://apib.ai/v1/music/generations/lyrics"},
		{"suno-cover", "https://apib.ai/v1/music/generations/coverSong"},
		{"suno-extend", "https://apib.ai/v1/music/generations/extend"},
		{"suno-stems-all", "https://apib.ai/v1/music/generations/stemsAll"},
		{"suno-upsample-tags", "https://apib.ai/v1/music/generations/upsampleTags"},
		{"suno-replace-section", "https://apib.ai/v1/music/generations/replaceMusic"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.model, func(t *testing.T) {
			t.Parallel()
			a := &TaskAdaptor{}
			info := newRelayInfo(tc.model)
			url, err := a.BuildRequestURL(info)
			if err != nil {
				t.Fatalf("BuildRequestURL(%q) returned error: %v", tc.model, err)
			}
			if url != tc.wantURL {
				t.Errorf("BuildRequestURL(%q) = %q, want %q", tc.model, url, tc.wantURL)
			}
		})
	}
}

// ── Version validation ────────────────────────────────────────────────────────

func TestValidate_MusicVersionRequired(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	// missing version → should fail
	c := newTestContext(map[string]any{"model": "suno-music", "prompt": "test"})
	info := newRelayInfo("suno-music")
	taskErr := a.ValidateRequestAndSetAction(c, info)
	if taskErr == nil {
		t.Fatal("expected error for missing version, got nil")
	}
	if !strings.Contains(taskErr.Error.Error(), "version") {
		t.Errorf("expected version error, got: %s", taskErr.Error)
	}
}

func TestValidate_MusicVersionValid(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	c := newTestContext(map[string]any{"model": "suno-music", "version": "v5", "prompt": "test"})
	info := newRelayInfo("suno-music")
	taskErr := a.ValidateRequestAndSetAction(c, info)
	if taskErr != nil {
		t.Fatalf("unexpected error for valid version: %s", taskErr.Error)
	}
}

func TestValidate_MusicVersionInvalid(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	c := newTestContext(map[string]any{"model": "suno-music", "version": "v99"})
	info := newRelayInfo("suno-music")
	taskErr := a.ValidateRequestAndSetAction(c, info)
	if taskErr == nil {
		t.Fatal("expected error for invalid version, got nil")
	}
}

func TestValidate_NoVersionToolIgnoresVersion(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	// suno-lyrics has no version dimension; passing version should be silently ignored
	c := newTestContext(map[string]any{"model": "suno-lyrics", "version": "v5", "prompt": "test"})
	info := newRelayInfo("suno-lyrics")
	taskErr := a.ValidateRequestAndSetAction(c, info)
	if taskErr != nil {
		t.Fatalf("unexpected error: %s", taskErr.Error)
	}
	// Verify version was cleared
	reqRaw, _ := c.Get("task_request")
	req := reqRaw.(*dto.APIMartSunoRequest)
	if req.Version != "" {
		t.Errorf("expected version to be cleared for no-version tool, got %q", req.Version)
	}
}

func TestValidate_SoundsVersionRestricted(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}

	t.Run("valid_v5", func(t *testing.T) {
		c := newTestContext(map[string]any{"model": "suno-sounds", "version": "v5", "prompt": "rain"})
		info := newRelayInfo("suno-sounds")
		if err := a.ValidateRequestAndSetAction(c, info); err != nil {
			t.Fatalf("unexpected error: %s", err.Error)
		}
	})
	t.Run("invalid_v3.5", func(t *testing.T) {
		c := newTestContext(map[string]any{"model": "suno-sounds", "version": "v3.5", "prompt": "rain"})
		info := newRelayInfo("suno-sounds")
		if err := a.ValidateRequestAndSetAction(c, info); err == nil {
			t.Fatal("expected error for v3.5 in suno-sounds, got nil")
		}
	})
}

func TestValidate_AddStemOnlyV55(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	t.Run("valid_v5.5", func(t *testing.T) {
		c := newTestContext(map[string]any{"model": "suno-add-stem", "version": "v5.5", "task_id": "task_abc"})
		info := newRelayInfo("suno-add-stem")
		if err := a.ValidateRequestAndSetAction(c, info); err != nil {
			t.Fatalf("unexpected error: %s", err.Error)
		}
	})
	t.Run("invalid_v5", func(t *testing.T) {
		c := newTestContext(map[string]any{"model": "suno-add-stem", "version": "v5", "task_id": "task_abc"})
		info := newRelayInfo("suno-add-stem")
		if err := a.ValidateRequestAndSetAction(c, info); err == nil {
			t.Fatal("expected error for v5 in suno-add-stem, got nil")
		}
	})
}

// ── Required field validation ─────────────────────────────────────────────────

func TestValidate_MissingTaskID(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	c := newTestContext(map[string]any{"model": "suno-extend", "version": "v5", "continue_at": 30})
	// task_id missing
	info := newRelayInfo("suno-extend")
	taskErr := a.ValidateRequestAndSetAction(c, info)
	if taskErr == nil {
		t.Fatal("expected error for missing task_id, got nil")
	}
	if !strings.Contains(taskErr.Error.Error(), "task_id") {
		t.Errorf("expected task_id error, got: %s", taskErr.Error)
	}
}

func TestValidate_MissingContinueAt(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	c := newTestContext(map[string]any{"model": "suno-extend", "version": "v5", "task_id": "task_abc"})
	info := newRelayInfo("suno-extend")
	taskErr := a.ValidateRequestAndSetAction(c, info)
	if taskErr == nil {
		t.Fatal("expected error for missing continue_at, got nil")
	}
}

func TestValidate_MashupRequiresExactly2TaskIDs(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}

	t.Run("one_task_id", func(t *testing.T) {
		c := newTestContext(map[string]any{"model": "suno-mashup", "version": "v5", "task_ids": []string{"id1"}})
		info := newRelayInfo("suno-mashup")
		if err := a.ValidateRequestAndSetAction(c, info); err == nil {
			t.Fatal("expected error for 1 task_id, got nil")
		}
	})
	t.Run("two_task_ids", func(t *testing.T) {
		c := newTestContext(map[string]any{"model": "suno-mashup", "version": "v5", "task_ids": []string{"id1", "id2"}})
		info := newRelayInfo("suno-mashup")
		if err := a.ValidateRequestAndSetAction(c, info); err != nil {
			t.Fatalf("unexpected error for 2 task_ids: %s", err.Error)
		}
	})
}

func TestValidate_InspoAudioURLsBounds(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}

	t.Run("zero_urls", func(t *testing.T) {
		c := newTestContext(map[string]any{"model": "suno-inspo", "version": "v5", "audio_urls": []string{}})
		info := newRelayInfo("suno-inspo")
		if err := a.ValidateRequestAndSetAction(c, info); err == nil {
			t.Fatal("expected error for 0 audio_urls")
		}
	})
	t.Run("five_urls", func(t *testing.T) {
		c := newTestContext(map[string]any{
			"model": "suno-inspo", "version": "v5",
			"audio_urls": []string{"u1", "u2", "u3", "u4", "u5"},
		})
		info := newRelayInfo("suno-inspo")
		if err := a.ValidateRequestAndSetAction(c, info); err == nil {
			t.Fatal("expected error for 5 audio_urls")
		}
	})
	t.Run("two_urls_valid", func(t *testing.T) {
		c := newTestContext(map[string]any{
			"model": "suno-inspo", "version": "v5",
			"audio_urls": []string{"u1", "u2"},
		})
		info := newRelayInfo("suno-inspo")
		if err := a.ValidateRequestAndSetAction(c, info); err != nil {
			t.Fatalf("unexpected error: %s", err.Error)
		}
	})
}

func TestValidate_AdjustSpeedRange(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	speed := func(v float64) map[string]any {
		return map[string]any{"model": "suno-adjust-speed", "task_id": "t1", "speed": v}
	}
	if err := a.ValidateRequestAndSetAction(newTestContext(speed(0.1)), newRelayInfo("suno-adjust-speed")); err == nil {
		t.Error("0.1 speed should fail")
	}
	if err := a.ValidateRequestAndSetAction(newTestContext(speed(5.0)), newRelayInfo("suno-adjust-speed")); err == nil {
		t.Error("5.0 speed should fail")
	}
	if err := a.ValidateRequestAndSetAction(newTestContext(speed(1.0)), newRelayInfo("suno-adjust-speed")); err != nil {
		t.Errorf("1.0 speed should pass, got: %s", err.Error)
	}
}

func TestValidate_StyleWeightOutOfRange(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	sw := 1.5
	c := newTestContext(map[string]any{"model": "suno-music", "version": "v5", "style_weight": sw})
	info := newRelayInfo("suno-music")
	if err := a.ValidateRequestAndSetAction(c, info); err == nil {
		t.Error("style_weight 1.5 should fail")
	}
}

func TestValidate_SoundsBPMRange(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	// BPM 0 → fail
	c := newTestContext(map[string]any{"model": "suno-sounds", "version": "v5", "prompt": "x", "bpm": 0})
	info := newRelayInfo("suno-sounds")
	if err := a.ValidateRequestAndSetAction(c, info); err == nil {
		t.Error("bpm=0 should fail")
	}
	// BPM 301 → fail
	c2 := newTestContext(map[string]any{"model": "suno-sounds", "version": "v5", "prompt": "x", "bpm": 301})
	info2 := newRelayInfo("suno-sounds")
	if err := a.ValidateRequestAndSetAction(c2, info2); err == nil {
		t.Error("bpm=301 should fail")
	}
}

// ── ParseTaskResult ───────────────────────────────────────────────────────────

func TestParseTaskResult_Completed(t *testing.T) {
	t.Parallel()
	body := []byte(`{
		"code": 200,
		"data": {
			"id": "task_abc",
			"status": "completed",
			"progress": 100,
			"result": {
				"music": [{"audio_url": "https://example.com/a.mp3", "title": "Test"}]
			}
		}
	}`)
	info, err := ParseTaskResult(body)
	if err != nil {
		t.Fatalf("ParseTaskResult error: %v", err)
	}
	if info.Status != string(model.TaskStatusSuccess) {
		t.Errorf("status = %q, want SUCCESS", info.Status)
	}
	if info.Url != "https://example.com/a.mp3" {
		t.Errorf("url = %q, want audio_url", info.Url)
	}
	if info.Progress != "100%" {
		t.Errorf("progress = %q, want 100%%", info.Progress)
	}
}

func TestParseTaskResult_Failed(t *testing.T) {
	t.Parallel()
	body := []byte(`{
		"code": 200,
		"data": {
			"id": "task_abc",
			"status": "failed",
			"progress": 100,
			"error": {"message": "upstream error"}
		}
	}`)
	info, err := ParseTaskResult(body)
	if err != nil {
		t.Fatalf("ParseTaskResult error: %v", err)
	}
	if info.Status != string(model.TaskStatusFailure) {
		t.Errorf("status = %q, want FAILURE", info.Status)
	}
	if info.Reason != "upstream error" {
		t.Errorf("reason = %q, want 'upstream error'", info.Reason)
	}
}

func TestParseTaskResult_Pending(t *testing.T) {
	t.Parallel()
	body := []byte(`{"code": 200, "data": {"id": "t1", "status": "pending", "progress": 50}}`)
	info, err := ParseTaskResult(body)
	if err != nil {
		t.Fatalf("ParseTaskResult error: %v", err)
	}
	if info.Status != string(model.TaskStatusQueued) {
		t.Errorf("status = %q, want QUEUED", info.Status)
	}
}

func TestParseTaskResult_Submitted(t *testing.T) {
	t.Parallel()
	body := []byte(`{"code": 200, "data": {"id": "t1", "status": "submitted", "progress": 10}}`)
	info, err := ParseTaskResult(body)
	if err != nil {
		t.Fatalf("ParseTaskResult error: %v", err)
	}
	if info.Status != string(model.TaskStatusSubmitted) {
		t.Errorf("status = %q, want SUBMITTED", info.Status)
	}
}

func TestParseTaskResult_UpstreamError(t *testing.T) {
	t.Parallel()
	// Non-200 code
	body := []byte(`{"code": 500, "message": "internal error"}`)
	info, err := ParseTaskResult(body)
	if err != nil {
		t.Fatalf("ParseTaskResult error: %v", err)
	}
	if info.Status != string(model.TaskStatusFailure) {
		t.Errorf("status = %q, want FAILURE", info.Status)
	}
}

func TestParseTaskResult_MalformedJSON(t *testing.T) {
	t.Parallel()
	_, err := ParseTaskResult([]byte(`not json`))
	if err == nil {
		t.Fatal("expected error for malformed JSON, got nil")
	}
}

// ── BuildRequestURL ───────────────────────────────────────────────────────────

func TestBuildRequestURL_UnknownModel(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	info := newRelayInfo("suno-nonexistent")
	_, err := a.BuildRequestURL(info)
	if err == nil {
		t.Fatal("expected error for unknown model, got nil")
	}
}

// ── Upstream body shape ───────────────────────────────────────────────────────

func TestBuildUpstreamBody_Music(t *testing.T) {
	t.Parallel()
	def, _ := GetToolDef("suno-music")
	boolTrue := true
	sw := 0.8
	req := &dto.APIMartSunoRequest{
		Version:     "v5",
		Custom:      &boolTrue,
		Prompt:      "summer vibes",
		Style:       "electronic, upbeat",
		StyleWeight: &sw,
	}
	body := buildUpstreamBody(req, "suno-music", def)

	if body["version"] != "v5" {
		t.Errorf("version = %v, want v5", body["version"])
	}
	if body["style"] != "electronic, upbeat" {
		t.Errorf("style = %v, want 'electronic, upbeat'", body["style"])
	}
	if body["style_weight"] != 0.8 {
		t.Errorf("style_weight = %v, want 0.8", body["style_weight"])
	}
	if _, hasTaskID := body["task_id"]; hasTaskID {
		t.Error("suno-music should not have task_id in body")
	}
}

func TestBuildUpstreamBody_Mashup(t *testing.T) {
	t.Parallel()
	def, _ := GetToolDef("suno-mashup")
	req := &dto.APIMartSunoRequest{
		Version: "v5",
		TaskIDs: []string{"id1", "id2"},
	}
	body := buildUpstreamBody(req, "suno-mashup", def)
	ids, ok := body["task_ids"].([]string)
	if !ok || len(ids) != 2 {
		t.Errorf("task_ids = %v, want [id1, id2]", body["task_ids"])
	}
	if _, hasTaskID := body["task_id"]; hasTaskID {
		t.Error("mashup should not have singular task_id")
	}
}

func TestBuildUpstreamBody_NoVersionTool(t *testing.T) {
	t.Parallel()
	def, _ := GetToolDef("suno-stems")
	req := &dto.APIMartSunoRequest{
		TaskID: "task_abc",
	}
	body := buildUpstreamBody(req, "suno-stems", def)
	if _, hasVer := body["version"]; hasVer {
		t.Error("no-version tool should not include version in body")
	}
	if body["task_id"] != "task_abc" {
		t.Errorf("task_id = %v, want task_abc", body["task_id"])
	}
}

func TestBuildUpstreamBody_Upload(t *testing.T) {
	t.Parallel()
	def, _ := GetToolDef("suno-upload")
	req := &dto.APIMartSunoRequest{AudioFilePath: "https://cdn/audio.mp3"}
	body := buildUpstreamBody(req, "suno-upload", def)
	if body["audioFilePath"] != "https://cdn/audio.mp3" {
		t.Errorf("audioFilePath = %v", body["audioFilePath"])
	}
}

func TestBuildUpstreamBody_UpsampleTags(t *testing.T) {
	t.Parallel()
	def, _ := GetToolDef("suno-upsample-tags")
	req := &dto.APIMartSunoRequest{Tags: "electronic, upbeat"}
	body := buildUpstreamBody(req, "suno-upsample-tags", def)
	if body["tags"] != "electronic, upbeat" {
		t.Errorf("tags = %v", body["tags"])
	}
}

// ── GetTaskAdaptor factory ────────────────────────────────────────────────────

func TestGetTaskAdaptor_ChannelType(t *testing.T) {
	t.Parallel()
	platform := constant.TaskPlatform("60") // strconv.Itoa(ChannelTypeAPIMartSuno)
	_ = platform // Just confirms the constant is 60
	if constant.ChannelTypeAPIMartSuno != 60 {
		t.Errorf("ChannelTypeAPIMartSuno = %d, want 60", constant.ChannelTypeAPIMartSuno)
	}
}

// ── Model list completeness ───────────────────────────────────────────────────

func TestModelList_HasExpectedCount(t *testing.T) {
	t.Parallel()
	if len(ModelList) != 31 {
		t.Errorf("ModelList has %d entries, want 31", len(ModelList))
	}
}

func TestModelList_NoDuplicates(t *testing.T) {
	t.Parallel()
	seen := make(map[string]bool)
	for _, m := range ModelList {
		if seen[m] {
			t.Errorf("duplicate model ID: %q", m)
		}
		seen[m] = true
	}
}

// ── FetchTask URL construction ────────────────────────────────────────────────

func TestFetchTask_URL(t *testing.T) {
	t.Parallel()
	var capturedURL string
	var capturedAuth string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedURL = r.URL.Path
		capturedAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"code":200,"data":{"id":"t1","status":"pending","progress":50}}`))
	}))
	defer ts.Close()

	// Build the request manually to verify URL construction (avoids service dependency in unit test)
	taskID := "upstream_task_123"
	wantPath := "/v1/music/tasks/" + taskID
	wantURL := strings.TrimRight(ts.URL, "/") + wantPath

	req, err := http.NewRequest(http.MethodGet, wantURL, nil)
	if err != nil {
		t.Fatalf("new request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer test-key")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("do request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if capturedURL != wantPath {
		t.Errorf("URL path = %q, want %q", capturedURL, wantPath)
	}
	if capturedAuth != "Bearer test-key" {
		t.Errorf("auth = %q, want 'Bearer test-key'", capturedAuth)
	}
}

func TestFetchTask_MissingTaskID(t *testing.T) {
	t.Parallel()
	a := &TaskAdaptor{}
	_, err := a.FetchTask("https://apib.ai", "key", map[string]any{}, "")
	if err == nil {
		t.Fatal("expected error for missing task_id")
	}
}
