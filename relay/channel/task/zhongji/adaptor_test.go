package zhongji

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/QuantumNous/new-api/service"

	"github.com/gin-gonic/gin"
)

func TestBuildRequestURL(t *testing.T) {
	a := &TaskAdaptor{}
	a.Init(&relaycommon.RelayInfo{ChannelMeta: &relaycommon.ChannelMeta{
		ChannelBaseUrl: "https://aiplay.gee.cn/new-api/",
	}})
	got, err := a.BuildRequestURL(nil)
	if err != nil {
		t.Fatal(err)
	}
	if got != "https://aiplay.gee.cn/new-api/v1/contents/generations/tasks" {
		t.Fatalf("url = %s", got)
	}
}

func TestDoRequestPostsV1Path(t *testing.T) {
	var gotPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"cgt-1"}`))
	}))
	defer server.Close()
	service.InitHttpClient()

	info := &relaycommon.RelayInfo{ChannelMeta: &relaycommon.ChannelMeta{
		ChannelBaseUrl: server.URL,
		ApiKey:         "secret",
	}}
	a := &TaskAdaptor{}
	a.Init(info)

	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Request = httptest.NewRequest(http.MethodPost, "/v1/video/generations", nil)

	resp, err := a.DoRequest(c, info, strings.NewReader(`{"model":"x"}`))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)
	if gotPath != "/v1/contents/generations/tasks" {
		t.Fatalf("path = %s", gotPath)
	}
}

func TestFetchTaskUsesV1Path(t *testing.T) {
	var gotPath, gotAuth string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id":"task_1","status":"queued"}`))
	}))
	defer server.Close()
	service.InitHttpClient()

	a := &TaskAdaptor{}
	resp, err := a.FetchTask(server.URL+"/", "secret", map[string]any{"task_id": "task_1"}, "")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)
	if gotPath != "/v1/contents/generations/tasks/task_1" {
		t.Fatalf("path = %s", gotPath)
	}
	if gotAuth != "Bearer secret" {
		t.Fatalf("auth = %s", gotAuth)
	}
	if strings.Contains(gotPath, "/api/v3/") {
		t.Fatalf("path still uses volc prefix: %s", gotPath)
	}
}

func TestParseTaskResultStatusAliases(t *testing.T) {
	a := &TaskAdaptor{}
	cases := []struct {
		name   string
		body   string
		status string
		url    string
	}{
		{name: "queued", body: `{"status":"queued"}`, status: string(model.TaskStatusQueued)},
		{name: "in_progress", body: `{"status":"in_progress"}`, status: string(model.TaskStatusInProgress)},
		{name: "completed", body: `{"status":"completed","content":{"video_url":"https://example.com/a.mp4"},"usage":{"completion_tokens":10,"total_tokens":10}}`, status: string(model.TaskStatusSuccess), url: "https://example.com/a.mp4"},
		{name: "succeeded", body: `{"status":"succeeded","content":{"video_url":"https://example.com/b.mp4"}}`, status: string(model.TaskStatusSuccess), url: "https://example.com/b.mp4"},
		{name: "failed", body: `{"status":"failed","error":{"message":"rejected"}}`, status: string(model.TaskStatusFailure)},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := a.ParseTaskResult([]byte(tc.body))
			if err != nil {
				t.Fatal(err)
			}
			if string(info.Status) != tc.status {
				t.Fatalf("status = %s, want %s", info.Status, tc.status)
			}
			if tc.url != "" && info.Url != tc.url {
				t.Fatalf("url = %s", info.Url)
			}
			if tc.name == "failed" && info.Reason != "rejected" {
				t.Fatalf("reason = %s", info.Reason)
			}
		})
	}
}
