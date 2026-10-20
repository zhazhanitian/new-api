package service

import (
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/model"
)

func TestApplySyncImageRecordKeepsUpstreamBody(t *testing.T) {
	task := &model.Task{}
	body := []byte(`{"Response":{"AigcImageTask":{"Status":"FINISH","Output":{"FileInfos":[{"FileUrl":"https://example.com/a.png"}]}}}}`)
	applySyncImageRecord(task, &syncImageRecord{
		UpstreamBody: body,
		URLs:         []string{"https://example.com/a.png"},
	})

	if task.PrivateData.ResultURL != "https://example.com/a.png" {
		t.Fatalf("result url = %q", task.PrivateData.ResultURL)
	}
	if !strings.Contains(string(task.Data), `"AigcImageTask"`) {
		t.Fatalf("data = %s", task.Data)
	}
}

func TestApplySyncImageRecordMultiImageUsesURLList(t *testing.T) {
	task := &model.Task{}
	body := []byte(`{"Response":{"AigcImageTask":{"Status":"FINISH"}}}`)
	applySyncImageRecord(task, &syncImageRecord{
		UpstreamBody: body,
		URLs:         []string{"https://example.com/a.png", "https://example.com/b.png"},
	})

	if task.PrivateData.ResultURL != "https://example.com/a.png" {
		t.Fatalf("result url = %q", task.PrivateData.ResultURL)
	}
	var images dto.ImageResponse
	if err := common.Unmarshal(task.Data, &images); err != nil {
		t.Fatal(err)
	}
	if len(images.Data) != 2 || images.Data[1].Url != "https://example.com/b.png" {
		t.Fatalf("images = %+v", images.Data)
	}
}
