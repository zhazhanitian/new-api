package trial

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const MaxNodes = 200

type APIKind string

const (
	APIKindChatCompletions APIKind = "chat_completions"
	APIKindResponses       APIKind = "responses"
	APIKindImageGeneration APIKind = "image_generation"
	APIKindImageTask       APIKind = "image_task"
	APIKindVideo           APIKind = "video"
	APIKindMusic           APIKind = "music"
	APIKindThreeD          APIKind = "threed"
)

type NodeType string

const (
	NodeField  NodeType = "field"
	NodeSelect NodeType = "select"
	NodeBranch NodeType = "branch"
	NodeObject NodeType = "object"
)

type Node struct {
	ID          string          `json:"id"`
	Type        NodeType        `json:"type"`
	Key         string          `json:"key,omitempty"`
	Label       string          `json:"label"`
	Help        string          `json:"help,omitempty"`
	Submit      *bool           `json:"submit,omitempty"`
	Required    bool            `json:"required,omitempty"`
	Default     json.RawMessage `json:"default,omitempty"`
	Widget      string          `json:"widget,omitempty"`
	AllowCustom bool            `json:"allow_custom,omitempty"`
	Min         *float64        `json:"min,omitempty"`
	Max         *float64        `json:"max,omitempty"`
	Options     []Option        `json:"options,omitempty"`
	Nodes       []Node          `json:"nodes,omitempty"`
}

type Option struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Help  string `json:"help,omitempty"`
	Nodes []Node `json:"nodes,omitempty"`
}

type Example struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description,omitempty"`
	Values      map[string]any `json:"values"`
	VerifiedAt  int64          `json:"verified_at"`
}

func (n Node) ShouldSubmit() bool {
	if n.Submit == nil {
		return true
	}
	return *n.Submit
}

func ValidAPIKind(kind APIKind) bool {
	switch kind {
	case APIKindChatCompletions, APIKindResponses, APIKindImageGeneration, APIKindImageTask, APIKindVideo, APIKindMusic, APIKindThreeD:
		return true
	default:
		return false
	}
}

func CanonicalPath(kind APIKind) (string, error) {
	switch kind {
	case APIKindChatCompletions:
		return "/v1/chat/completions", nil
	case APIKindResponses:
		return "/v1/responses", nil
	case APIKindImageGeneration:
		return "/v1/images/generations", nil
	case APIKindImageTask:
		return "/v2/image-tasks", nil
	case APIKindVideo:
		return "/v1/video/generations", nil
	case APIKindMusic:
		return "/v1/music/generations", nil
	case APIKindThreeD:
		return "/v1/3d/generations", nil
	default:
		return "", fmt.Errorf("未知的接口类型 %s", kind)
	}
}

func IsAsync(kind APIKind) bool {
	switch kind {
	case APIKindImageTask, APIKindVideo, APIKindMusic, APIKindThreeD:
		return true
	default:
		return false
	}
}

func Validate(kind APIKind, nodes []Node) error {
	if !ValidAPIKind(kind) {
		return fmt.Errorf("未知的接口类型 %s", kind)
	}
	if len(nodes) == 0 {
		return fmt.Errorf("至少配置一个节点")
	}
	seen := map[string]struct{}{}
	count := 0
	if err := validateNodes(kind, nodes, "", &count, seen); err != nil {
		return err
	}
	if count > MaxNodes {
		return fmt.Errorf("节点总数不能超过 %d", MaxNodes)
	}
	switch kind {
	case APIKindChatCompletions:
		if !hasWidget(nodes, "messages") {
			return fmt.Errorf("对话补全必须包含 messages 控件")
		}
	case APIKindResponses:
		if !hasWidget(nodes, "input") {
			return fmt.Errorf("Responses 必须包含 input 控件")
		}
	}
	return nil
}

func validateNodes(kind APIKind, nodes []Node, prefix string, count *int, seen map[string]struct{}) error {
	for _, node := range nodes {
		*count++
		if strings.TrimSpace(node.ID) == "" {
			return fmt.Errorf("节点缺少 id")
		}
		if _, ok := seen[node.ID]; ok {
			return fmt.Errorf("节点 id 重复: %s", node.ID)
		}
		seen[node.ID] = struct{}{}
		if strings.TrimSpace(node.Label) == "" {
			return fmt.Errorf("节点 %s 缺少标题", node.ID)
		}
		switch node.Type {
		case NodeField:
			if err := validateField(kind, node, prefix); err != nil {
				return err
			}
		case NodeSelect, NodeBranch:
			if node.Type == NodeBranch && node.AllowCustom {
				return fmt.Errorf("分支节点 %s 不能开启自定义，自定义只用于下拉", node.Label)
			}
			if node.Type == NodeSelect && node.Widget != "" {
				return fmt.Errorf("下拉节点 %s 不使用控件类型", node.Label)
			}
			if err := validateChoice(kind, node, prefix, count, seen); err != nil {
				return err
			}
		case NodeObject:
			if node.ShouldSubmit() && strings.TrimSpace(node.Key) == "" {
				return fmt.Errorf("参与提交的参数对象 %s 必须填写写入路径", node.Label)
			}
			if node.ShouldSubmit() {
				if _, err := resolvePath(prefix, node.Key); err != nil {
					return fmt.Errorf("参数对象 %s 的路径无效: %s", node.Label, err.Error())
				}
			}
			next := prefix
			if node.ShouldSubmit() {
				path, err := resolvePath(prefix, node.Key)
				if err != nil {
					return err
				}
				if err := checkRoot(kind, path); err != nil {
					return fmt.Errorf("参数对象 %s: %s", node.Label, err.Error())
				}
				next = path
			}
			if err := validateNodes(kind, node.Nodes, next, count, seen); err != nil {
				return err
			}
		default:
			return fmt.Errorf("节点 %s 的类型无效", node.Label)
		}
	}
	return nil
}

func validateField(kind APIKind, node Node, prefix string) error {
	if !validWidget(node.Widget) {
		return fmt.Errorf("字段 %s 的控件类型无效", node.Label)
	}
	if node.Widget == "messages" && kind != APIKindChatCompletions {
		return fmt.Errorf("messages 控件只能用于对话补全")
	}
	if node.Widget == "input" && kind != APIKindResponses {
		return fmt.Errorf("input 控件只能用于 Responses")
	}
	if node.AllowCustom {
		return fmt.Errorf("自定义只用于下拉，字段 %s 不能开启", node.Label)
	}
	if node.Widget == "slider" {
		if node.Min == nil || node.Max == nil {
			return fmt.Errorf("拖动条 %s 必须填写最小值和最大值", node.Label)
		}
		if *node.Min >= *node.Max {
			return fmt.Errorf("拖动条 %s 的最小值必须小于最大值", node.Label)
		}
	}
	if node.ShouldSubmit() {
		if strings.TrimSpace(node.Key) == "" {
			return fmt.Errorf("参与提交的字段 %s 必须填写写入路径", node.Label)
		}
		path, err := resolvePath(prefix, node.Key)
		if err != nil {
			return fmt.Errorf("字段 %s 的路径无效: %s", node.Label, err.Error())
		}
		if err := checkRoot(kind, path); err != nil {
			return fmt.Errorf("字段 %s: %s", node.Label, err.Error())
		}
	} else if strings.TrimSpace(node.Key) != "" {
		if _, err := resolvePath(prefix, node.Key); err != nil {
			return fmt.Errorf("字段 %s 的路径无效: %s", node.Label, err.Error())
		}
	}
	return nil
}

func validateChoice(kind APIKind, node Node, prefix string, count *int, seen map[string]struct{}) error {
	if len(node.Options) == 0 {
		return fmt.Errorf("%s 至少需要一个选项", node.Label)
	}
	if node.ShouldSubmit() && strings.TrimSpace(node.Key) == "" {
		return fmt.Errorf("参与提交的%s %s 必须填写写入路径", choiceName(node.Type), node.Label)
	}
	if !node.ShouldSubmit() && strings.TrimSpace(node.Key) != "" {
		return fmt.Errorf("不参与提交的%s %s 不要填写写入路径", choiceName(node.Type), node.Label)
	}
	if node.ShouldSubmit() {
		path, err := resolvePath(prefix, node.Key)
		if err != nil {
			return fmt.Errorf("%s %s 的路径无效: %s", choiceName(node.Type), node.Label, err.Error())
		}
		if err := checkRoot(kind, path); err != nil {
			return fmt.Errorf("%s %s: %s", choiceName(node.Type), node.Label, err.Error())
		}
	}
	seenValue := map[string]struct{}{}
	for _, option := range node.Options {
		if strings.TrimSpace(option.Value) == "" || strings.TrimSpace(option.Label) == "" {
			return fmt.Errorf("%s 的选项必须有值和标题", node.Label)
		}
		if _, ok := seenValue[option.Value]; ok {
			return fmt.Errorf("%s 的选项值重复: %s", node.Label, option.Value)
		}
		seenValue[option.Value] = struct{}{}
		if err := validateNodes(kind, option.Nodes, prefix, count, seen); err != nil {
			return err
		}
	}
	return nil
}

func choiceName(t NodeType) string {
	if t == NodeBranch {
		return "分支"
	}
	return "下拉"
}

func validWidget(widget string) bool {
	switch widget {
	case "text", "textarea", "number", "integer", "slider", "boolean",
		"image_url", "image_urls", "video_url", "audio_url", "file_url",
		"messages", "input", "json":
		return true
	default:
		return false
	}
}

func hasWidget(nodes []Node, widget string) bool {
	for _, node := range nodes {
		if node.Type == NodeField && node.Widget == widget {
			return true
		}
		if hasWidget(node.Nodes, widget) {
			return true
		}
		for _, option := range node.Options {
			if hasWidget(option.Nodes, widget) {
				return true
			}
		}
	}
	return false
}

func resolvePath(prefix, key string) (string, error) {
	key = strings.TrimSpace(key)
	if key == "" {
		return "", fmt.Errorf("路径为空")
	}
	path := key
	if strings.HasPrefix(key, "/") {
		path = strings.TrimPrefix(key, "/")
	} else if prefix != "" {
		path = prefix + "." + key
	}
	if path == "" || strings.Contains(path, "..") {
		return "", fmt.Errorf("路径 %s 无效", key)
	}
	for _, seg := range strings.Split(path, ".") {
		if !validSegment(seg) {
			return "", fmt.Errorf("路径片段 %s 无效", seg)
		}
	}
	return path, nil
}

func validSegment(seg string) bool {
	if seg == "" {
		return false
	}
	for _, r := range seg {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' || r == '-' {
			continue
		}
		return false
	}
	return true
}

func checkRoot(kind APIKind, path string) error {
	allowed := allowedRoots(kind)
	if allowed == nil {
		return nil
	}
	root := strings.Split(path, ".")[0]
	if !allowed[root] {
		return fmt.Errorf("路径 %s 不在该接口允许的字段内", path)
	}
	return nil
}

func allowedRoots(kind APIKind) map[string]bool {
	switch kind {
	case APIKindChatCompletions:
		return map[string]bool{
			"model": true, "messages": true, "stream": true, "temperature": true, "top_p": true,
			"max_tokens": true, "max_completion_tokens": true, "frequency_penalty": true,
			"presence_penalty": true, "stop": true, "n": true, "seed": true,
			"response_format": true, "tools": true, "tool_choice": true, "reasoning_effort": true,
		}
	case APIKindResponses:
		return map[string]bool{
			"model": true, "input": true, "instructions": true, "stream": true,
			"temperature": true, "top_p": true, "max_output_tokens": true,
			"tools": true, "reasoning": true, "previous_response_id": true,
		}
	case APIKindImageGeneration, APIKindImageTask:
		return map[string]bool{
			"model": true, "prompt": true, "image": true, "images": true, "n": true,
			"size": true, "quality": true, "response_format": true, "output_format": true,
			"watermark": true, "extra_fields": true, "metadata": true,
		}
	case APIKindVideo:
		return map[string]bool{
			"model": true, "prompt": true, "images": true, "image": true, "duration": true,
			"seconds": true, "size": true, "seed": true, "n": true, "metadata": true,
		}
	case APIKindThreeD:
		return map[string]bool{"model": true, "metadata": true}
	default:
		return nil
	}
}

func ExampleStale(nodes []Node, values map[string]any) bool {
	if values == nil {
		return false
	}
	ids := map[string]Node{}
	collectNodes(nodes, ids)
	for id := range values {
		node, ok := ids[id]
		if !ok {
			return true
		}
		if node.Type == NodeSelect || node.Type == NodeBranch {
			raw, ok := values[id].(string)
			if !ok {
				return true
			}
			if !optionExists(node, raw) && !(node.Type == NodeSelect && node.AllowCustom) {
				return true
			}
		}
	}
	return false
}

func collectNodes(nodes []Node, ids map[string]Node) {
	for _, node := range nodes {
		ids[node.ID] = node
		collectNodes(node.Nodes, ids)
		for _, option := range node.Options {
			collectNodes(option.Nodes, ids)
		}
	}
}

func optionExists(node Node, value string) bool {
	for _, option := range node.Options {
		if option.Value == value {
			return true
		}
	}
	return false
}

func HashValues(values map[string]any) (string, error) {
	if values == nil {
		values = map[string]any{}
	}
	raw, err := json.Marshal(values)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:]), nil
}
