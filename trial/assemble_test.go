package trial

import (
	"encoding/json"
	"strings"
	"testing"
)

func boolPtr(v bool) *bool { return &v }

func TestAssembleDropsInactiveBranchAndHelp(t *testing.T) {
	nodes := []Node{
		{
			ID: "mode", Type: NodeBranch, Key: "metadata.scene_type", Label: "生成模式",
			Help: "HELP_SHOULD_NOT_APPEAR", Submit: boolPtr(true), Required: true,
			Options: []Option{
				{
					Value: "text2img", Label: "文生图",
					Nodes: []Node{{
						ID: "prompt-text", Type: NodeField, Key: "prompt", Label: "提示词",
						Help: "HELP_SHOULD_NOT_APPEAR", Widget: "textarea", Required: true,
					}},
				},
				{
					Value: "first_last_frame", Label: "首尾帧",
					Nodes: []Node{
						{ID: "prompt-frame", Type: NodeField, Key: "prompt", Label: "提示词", Widget: "textarea", Required: true},
						{ID: "first", Type: NodeField, Key: "images.0", Label: "首帧", Widget: "image_url", Required: true},
						{ID: "last", Type: NodeField, Key: "images.1", Label: "尾帧", Widget: "image_url", Required: true},
					},
				},
			},
		},
	}
	body, err := Assemble(APIKindImageTask, "demo-image", nodes, map[string]any{
		"mode":         "first_last_frame",
		"prompt-frame": "从白天到黑夜",
		"first":        "https://example.com/a.png",
		"last":         "https://example.com/b.png",
		"prompt-text":  "这段不该出现",
	})
	if err != nil {
		t.Fatal(err)
	}
	raw, _ := json.Marshal(body)
	if strings.Contains(string(raw), "HELP_SHOULD_NOT_APPEAR") || strings.Contains(string(raw), "这段不该出现") {
		t.Fatalf("unexpected body %s", raw)
	}
	if body["prompt"] != "从白天到黑夜" {
		t.Fatalf("prompt = %#v", body["prompt"])
	}
	meta := body["metadata"].(map[string]any)
	if meta["scene_type"] != "first_last_frame" {
		t.Fatalf("scene = %#v", meta["scene_type"])
	}
	images := body["images"].([]any)
	if len(images) != 2 || images[0] != "https://example.com/a.png" || images[1] != "https://example.com/b.png" {
		t.Fatalf("images = %#v", images)
	}

	textBody, err := Assemble(APIKindImageTask, "demo-image", nodes, map[string]any{
		"mode":        "text2img",
		"prompt-text": "一只猫",
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := textBody["images"]; ok {
		t.Fatalf("text mode leaked images: %#v", textBody["images"])
	}
}

func TestSubmitFalseBranchIsOmitted(t *testing.T) {
	nodes := []Node{{
		ID: "mode", Type: NodeBranch, Label: "展示模式", Submit: boolPtr(false), Required: true,
		Options: []Option{{
			Value: "a", Label: "甲",
			Nodes: []Node{{
				ID: "prompt", Type: NodeField, Key: "prompt", Label: "提示词", Widget: "text", Required: true,
			}},
		}},
	}}
	body, err := Assemble(APIKindImageGeneration, "demo", nodes, map[string]any{
		"mode":   "a",
		"prompt": "hello",
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := body["mode"]; ok {
		t.Fatalf("branch value leaked: %#v", body)
	}
	if body["prompt"] != "hello" || body["model"] != "demo" {
		t.Fatalf("body = %#v", body)
	}
}

func TestCustomSelectAndNestedObject(t *testing.T) {
	min := 0.0
	max := 1.0
	nodes := []Node{
		{
			ID: "ratio", Type: NodeSelect, Key: "metadata.ratio", Label: "比例", AllowCustom: true, Required: true,
			Options: []Option{{Value: "16:9", Label: "16:9"}},
		},
		{
			ID: "kind", Type: NodeBranch, Key: "metadata.kind", Label: "类型", Required: true,
			Options: []Option{{
				Value: "image", Label: "图生图",
				Nodes: []Node{{
					ID: "method", Type: NodeBranch, Key: "metadata.method", Label: "方式", Required: true,
					Options: []Option{{
						Value: "frames", Label: "首尾帧",
						Nodes: []Node{{
							ID: "frames", Type: NodeObject, Key: "metadata.frames", Label: "帧参数",
							Nodes: []Node{
								{
									ID: "first", Type: NodeObject, Key: "first", Label: "首帧",
									Nodes: []Node{
										{ID: "url", Type: NodeField, Key: "url", Label: "地址", Widget: "image_url", Required: true},
										{ID: "weight", Type: NodeField, Key: "weight", Label: "权重", Widget: "number", Min: &min, Max: &max},
									},
								},
							},
						}},
					}},
				}},
			}},
		},
	}
	body, err := Assemble(APIKindVideo, "demo-video", nodes, map[string]any{
		"ratio":  "21:9",
		"kind":   "image",
		"method": "frames",
		"url":    "https://example.com/a.png",
		"weight": 0.5,
	})
	if err != nil {
		t.Fatal(err)
	}
	meta := body["metadata"].(map[string]any)
	if meta["ratio"] != "21:9" || meta["kind"] != "image" || meta["method"] != "frames" {
		t.Fatalf("metadata = %#v", meta)
	}
	frames := meta["frames"].(map[string]any)
	first := frames["first"].(map[string]any)
	if first["url"] != "https://example.com/a.png" || first["weight"] != 0.5 {
		t.Fatalf("frames = %#v", frames)
	}
}

func TestKindSkeletons(t *testing.T) {
	cases := []struct {
		kind  APIKind
		nodes []Node
		vals  map[string]any
		key   string
	}{
		{APIKindChatCompletions, []Node{
			{ID: "m", Type: NodeField, Key: "messages", Label: "消息", Widget: "messages", Required: true},
			{ID: "s", Type: NodeField, Key: "stream", Label: "流式", Widget: "boolean"},
		}, map[string]any{"m": []any{map[string]any{"role": "user", "content": "hi"}}, "s": true}, "messages"},
		{APIKindResponses, []Node{
			{ID: "in", Type: NodeField, Key: "input", Label: "输入", Widget: "input", Required: true},
		}, map[string]any{"in": "hello"}, "input"},
		{APIKindImageGeneration, []Node{
			{ID: "p", Type: NodeField, Key: "prompt", Label: "提示词", Widget: "textarea", Required: true},
		}, map[string]any{"p": "cat"}, "prompt"},
		{APIKindMusic, []Node{
			{ID: "p", Type: NodeField, Key: "prompt", Label: "提示词", Widget: "textarea", Required: true},
			{ID: "v", Type: NodeField, Key: "version", Label: "版本", Widget: "text"},
		}, map[string]any{"p": "song", "v": "v6"}, "version"},
		{APIKindThreeD, []Node{
			{ID: "meta", Type: NodeObject, Key: "metadata", Label: "参数", Nodes: []Node{
				{ID: "p", Type: NodeField, Key: "prompt", Label: "提示词", Widget: "textarea", Required: true},
			}},
		}, map[string]any{"p": "cube"}, "metadata"},
	}
	for _, tc := range cases {
		body, err := Assemble(tc.kind, "m", tc.nodes, tc.vals)
		if err != nil {
			t.Fatalf("%s: %v", tc.kind, err)
		}
		if body["model"] != "m" {
			t.Fatalf("%s model overwritten", tc.kind)
		}
		if _, ok := body[tc.key]; !ok {
			t.Fatalf("%s missing %s in %#v", tc.kind, tc.key, body)
		}
		path, err := CanonicalPath(tc.kind)
		if err != nil || path == "" {
			t.Fatalf("%s path %v", tc.kind, err)
		}
	}
}

func TestSliderRange(t *testing.T) {
	minValue := 1.0
	maxValue := 15.0
	nodes := []Node{
		{ID: "p", Type: NodeField, Key: "prompt", Label: "提示词", Widget: "textarea", Required: true},
		{ID: "n", Type: NodeField, Key: "n", Label: "张数", Widget: "slider", Min: &minValue, Max: &maxValue},
	}
	body, err := Assemble(APIKindImageTask, "demo", nodes, map[string]any{
		"p": "一只猫",
		"n": 8,
	})
	if err != nil {
		t.Fatal(err)
	}
	if body["n"] != int64(8) {
		t.Fatalf("n = %#v", body["n"])
	}
	if _, err := Assemble(APIKindImageTask, "demo", nodes, map[string]any{"p": "一只猫", "n": 16}); err == nil {
		t.Fatal("expected out of range slider to fail")
	}
	if err := Validate(APIKindImageTask, []Node{
		{ID: "p", Type: NodeField, Key: "prompt", Label: "提示词", Widget: "textarea", Required: true},
		{ID: "n", Type: NodeField, Key: "n", Label: "张数", Widget: "slider"},
	}); err == nil {
		t.Fatal("expected slider without range to fail")
	}
}

func TestRejectsEscapingPath(t *testing.T) {
	nodes := []Node{{
		ID: "p", Type: NodeField, Key: "unknown_root", Label: "提示词", Widget: "text", Required: true,
	}}
	if err := Validate(APIKindImageGeneration, nodes); err == nil {
		t.Fatal("expected path to be rejected")
	}
}

func TestPreviewSkipsEmptyRequired(t *testing.T) {
	nodes := []Node{
		{ID: "prompt", Type: NodeField, Key: "prompt", Label: "提示词", Widget: "textarea", Required: true},
		{ID: "size", Type: NodeField, Key: "size", Label: "尺寸", Widget: "text"},
	}
	body, path, err := Preview(APIKindImageGeneration, "OG-image-2", nodes, map[string]any{
		"size": "1792x1024",
	})
	if err != nil {
		t.Fatal(err)
	}
	if path != "/v1/images/generations" {
		t.Fatalf("path = %s", path)
	}
	if body["model"] != "OG-image-2" || body["size"] != "1792x1024" {
		t.Fatalf("body = %#v", body)
	}
	if _, ok := body["prompt"]; ok {
		t.Fatalf("empty prompt leaked: %#v", body)
	}
}

func TestExampleStaleWhenOptionRemoved(t *testing.T) {
	nodes := []Node{{
		ID: "mode", Type: NodeBranch, Label: "模式", Submit: boolPtr(false),
		Options: []Option{{Value: "a", Label: "甲"}},
	}}
	if ExampleStale(nodes, map[string]any{"mode": "gone"}) == false {
		t.Fatal("expected stale example")
	}
	if ExampleStale(nodes, map[string]any{"mode": "a"}) {
		t.Fatal("expected fresh example")
	}
}
