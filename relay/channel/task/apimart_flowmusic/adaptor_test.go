package apimart_flowmusic

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
)

func newTestContext(body any) *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	data, _ := json.Marshal(body)
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/music/generations", bytes.NewReader(data))
	c.Request.Header.Set("Content-Type", "application/json")
	return c
}

func newRelayInfo(modelID string) *relaycommon.RelayInfo {
	return &relaycommon.RelayInfo{
		OriginModelName: modelID,
		ChannelMeta: &relaycommon.ChannelMeta{
			ChannelBaseUrl: "https://apib.ai",
			ApiKey:         "test-key",
		},
		TaskRelayInfo: &relaycommon.TaskRelayInfo{
			PublicTaskID: "task_public123",
		},
	}
}

func TestChannelType(t *testing.T) {
	if constant.ChannelTypeAPIMartFlowMusic != 63 {
		t.Fatalf("ChannelTypeAPIMartFlowMusic = %d, want 63", constant.ChannelTypeAPIMartFlowMusic)
	}
}

func TestGetToolDef_AllModelsRegistered(t *testing.T) {
	if len(ModelList) != len(toolDefs) {
		t.Fatalf("ModelList has %d models, toolDefs has %d", len(ModelList), len(toolDefs))
	}
	for _, m := range ModelList {
		if _, ok := GetToolDef(m); !ok {
			t.Errorf("model %q has no ToolDef", m)
		}
	}
}

func TestBuildRequestURL(t *testing.T) {
	cases := []struct {
		model   string
		wantURL string
	}{
		{"flowmusic", "https://apib.ai/v1/music/generations"},
		{"flowmusic-lyrics", "https://apib.ai/v1/music/generations/lyricsFlowMusic"},
		{"flowmusic-extend", "https://apib.ai/v1/music/generations/extendFlowMusic"},
		{"flowmusic-replace", "https://apib.ai/v1/music/generations/replaceFlowMusic"},
		{"flowmusic-cover", "https://apib.ai/v1/music/generations/coverFlowMusic"},
		{"flowmusic-stems", "https://apib.ai/v1/music/generations/stemsFlowMusic"},
		{"flowmusic-upload", "https://apib.ai/v1/music/generations/uploadAudioFlowMusic"},
		{"flowmusic-download", "https://apib.ai/v1/music/generations/downloadAudioFlowMusic"},
		{"flowmusic-video", "https://apib.ai/v1/music/generations/videoClipFlowMusic"},
	}
	a := &TaskAdaptor{}
	for _, tc := range cases {
		t.Run(tc.model, func(t *testing.T) {
			url, err := a.BuildRequestURL(newRelayInfo(tc.model))
			if err != nil {
				t.Fatalf("BuildRequestURL: %v", err)
			}
			if url != tc.wantURL {
				t.Fatalf("url = %s, want %s", url, tc.wantURL)
			}
		})
	}
}

func TestBuildRequestBody_ForwardsFieldsAndDropsTaskID(t *testing.T) {
	c := newTestContext(map[string]any{
		"model":        "flowmusic",
		"version":      "lyria-3.5",
		"title":        "坚持",
		"lyrics":       "[Verse 1]\n黑夜再长也会天亮",
		"sound_prompt": "energetic rock",
		"bpm":          "120",
		"length":       120,
		"task_id":      "task_public_origin",
	})
	info := newRelayInfo("flowmusic")
	a := &TaskAdaptor{}
	if err := a.ValidateRequestAndSetAction(c, info); err != nil {
		t.Fatalf("validate: %v", err)
	}
	if info.Action != "flowmusic" {
		t.Fatalf("action = %s", info.Action)
	}

	reader, err := a.BuildRequestBody(c, info)
	if err != nil {
		t.Fatalf("build body: %v", err)
	}
	var got map[string]json.RawMessage
	if err := json.NewDecoder(reader).Decode(&got); err != nil {
		t.Fatalf("decode upstream body: %v", err)
	}
	if string(got["model"]) != `"flowmusic"` {
		t.Fatalf("model = %s", got["model"])
	}
	if string(got["version"]) != `"lyria-3.5"` {
		t.Fatalf("version = %s", got["version"])
	}
	if string(got["bpm"]) != `"120"` {
		t.Fatalf("bpm = %s", got["bpm"])
	}
	if string(got["length"]) != "120" {
		t.Fatalf("length = %s", got["length"])
	}
	if string(got["sound_prompt"]) != `"energetic rock"` {
		t.Fatalf("sound_prompt = %s", got["sound_prompt"])
	}
	if _, ok := got["task_id"]; ok {
		t.Fatal("task_id was forwarded")
	}
}

func TestBuildRequestBody_DefaultsLyria3Pro(t *testing.T) {
	c := newTestContext(map[string]any{
		"model":        "flowmusic",
		"sound_prompt": "upbeat pop",
	})
	info := newRelayInfo("flowmusic")
	a := &TaskAdaptor{}
	if err := a.ValidateRequestAndSetAction(c, info); err != nil {
		t.Fatalf("validate: %v", err)
	}
	reader, err := a.BuildRequestBody(c, info)
	if err != nil {
		t.Fatalf("build body: %v", err)
	}
	var got map[string]json.RawMessage
	if err := json.NewDecoder(reader).Decode(&got); err != nil {
		t.Fatalf("decode upstream body: %v", err)
	}
	if string(got["version"]) != `"lyria-3-pro"` {
		t.Fatalf("version = %s, want default lyria-3-pro", got["version"])
	}
}

func TestBuildRequestBody_LyricsDoesNotDefaultVersion(t *testing.T) {
	c := newTestContext(map[string]any{
		"model":  "flowmusic-lyrics",
		"prompt": "一首关于坚持的摇滚歌曲",
	})
	info := newRelayInfo("flowmusic-lyrics")
	a := &TaskAdaptor{}
	if err := a.ValidateRequestAndSetAction(c, info); err != nil {
		t.Fatalf("validate: %v", err)
	}
	reader, err := a.BuildRequestBody(c, info)
	if err != nil {
		t.Fatalf("build body: %v", err)
	}
	var got map[string]json.RawMessage
	if err := json.NewDecoder(reader).Decode(&got); err != nil {
		t.Fatalf("decode upstream body: %v", err)
	}
	if _, ok := got["version"]; ok {
		t.Fatalf("lyrics forwarded version %s", got["version"])
	}
}

func TestBuildRequestBody_ForwardsVersionOnAnyTool(t *testing.T) {
	c := newTestContext(map[string]any{
		"model":   "flowmusic-lyrics",
		"version": "lyria-next",
		"prompt":  "一首关于坚持的摇滚歌曲",
	})
	info := newRelayInfo("flowmusic-lyrics")
	a := &TaskAdaptor{}
	if err := a.ValidateRequestAndSetAction(c, info); err != nil {
		t.Fatalf("validate: %v", err)
	}
	reader, err := a.BuildRequestBody(c, info)
	if err != nil {
		t.Fatalf("build body: %v", err)
	}
	var got map[string]json.RawMessage
	if err := json.NewDecoder(reader).Decode(&got); err != nil {
		t.Fatalf("decode upstream body: %v", err)
	}
	if string(got["model"]) != `"flowmusic"` {
		t.Fatalf("model = %s", got["model"])
	}
	if string(got["version"]) != `"lyria-next"` {
		t.Fatalf("version = %s, want forwarded as-is", got["version"])
	}
	var prompt string
	if err := json.Unmarshal(got["prompt"], &prompt); err != nil {
		t.Fatalf("prompt: %v", err)
	}
	if prompt != "一首关于坚持的摇滚歌曲" {
		t.Fatalf("prompt = %s", prompt)
	}
}

func TestValidate_UnknownModel(t *testing.T) {
	c := newTestContext(map[string]any{"model": "flowmusic-nope"})
	info := newRelayInfo("flowmusic-nope")
	if err := (&TaskAdaptor{}).ValidateRequestAndSetAction(c, info); err == nil {
		t.Fatal("expected unknown model error")
	}
}

func TestValidate_RejectsNonObject(t *testing.T) {
	c := newTestContext([]any{"not", "an", "object"})
	info := newRelayInfo("flowmusic")
	if err := (&TaskAdaptor{}).ValidateRequestAndSetAction(c, info); err == nil {
		t.Fatal("expected invalid request")
	}
}

func TestParseTaskResult_ResultURLPriority(t *testing.T) {
	cases := []struct {
		name    string
		body    string
		wantURL string
		status  string
	}{
		{
			name:    "audio",
			body:    `{"code":200,"data":{"id":"t1","status":"completed","progress":100,"result":{"music":[{"audio_url":"https://cdn/a.m4a","video_url":"https://cdn/a.mp4","file_url":"https://cdn/a.zip"}]}}}`,
			wantURL: "https://cdn/a.m4a",
			status:  string(model.TaskStatusSuccess),
		},
		{
			name:    "video",
			body:    `{"code":200,"data":{"id":"t2","status":"completed","result":{"music":[{"video_url":"https://cdn/a.mp4","file_url":"https://cdn/a.zip"}]}}}`,
			wantURL: "https://cdn/a.mp4",
			status:  string(model.TaskStatusSuccess),
		},
		{
			name:    "file",
			body:    `{"code":200,"data":{"id":"t3","status":"completed","result":{"music":[{"file_url":"https://cdn/a.zip"}]}}}`,
			wantURL: "https://cdn/a.zip",
			status:  string(model.TaskStatusSuccess),
		},
		{
			name:    "lyrics",
			body:    `{"code":200,"data":{"id":"t4","status":"completed","result":{"lyrics":[{"title":"坚持","lyrics":"黑夜"}]}}}`,
			wantURL: "",
			status:  string(model.TaskStatusSuccess),
		},
		{
			name:   "failed",
			body:   `{"code":200,"data":{"id":"t5","status":"failed","error":{"message":"clip not found"}}}`,
			status: string(model.TaskStatusFailure),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := ParseTaskResult([]byte(tc.body))
			if err != nil {
				t.Fatalf("parse: %v", err)
			}
			if info.Status != tc.status {
				t.Fatalf("status = %s, want %s", info.Status, tc.status)
			}
			if info.Url != tc.wantURL {
				t.Fatalf("url = %s, want %s", info.Url, tc.wantURL)
			}
			if tc.name == "failed" && info.Reason != "clip not found" {
				t.Fatalf("reason = %s", info.Reason)
			}
		})
	}
}
