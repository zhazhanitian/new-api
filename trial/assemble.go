package trial

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func Assemble(kind APIKind, modelName string, nodes []Node, values map[string]any) (map[string]any, error) {
	return assemble(kind, modelName, nodes, values, false)
}

// Preview 用和正式组装相同的规则生成请求体。未填的必填项会跳过，方便对照表单查看结构。
func Preview(kind APIKind, modelName string, nodes []Node, values map[string]any) (map[string]any, string, error) {
	path, err := CanonicalPath(kind)
	if err != nil {
		return nil, "", err
	}
	body, err := assemble(kind, modelName, nodes, values, true)
	if err != nil {
		return nil, path, err
	}
	return body, path, nil
}

func assemble(kind APIKind, modelName string, nodes []Node, values map[string]any, preview bool) (map[string]any, error) {
	if err := Validate(kind, nodes); err != nil {
		return nil, err
	}
	if strings.TrimSpace(modelName) == "" {
		return nil, fmt.Errorf("模型名为空")
	}
	if values == nil {
		values = map[string]any{}
	}
	body := map[string]any{}
	if err := applyNodes(kind, nodes, "", body, values, preview); err != nil {
		return nil, err
	}
	body["model"] = modelName
	return body, nil
}

func applyNodes(kind APIKind, nodes []Node, prefix string, body map[string]any, values map[string]any, preview bool) error {
	for _, node := range nodes {
		switch node.Type {
		case NodeField:
			if err := applyField(node, prefix, body, values, preview); err != nil {
				return err
			}
		case NodeSelect, NodeBranch:
			if err := applyChoice(kind, node, prefix, body, values, preview); err != nil {
				return err
			}
		case NodeObject:
			next := prefix
			if node.ShouldSubmit() {
				path, err := resolvePath(prefix, node.Key)
				if err != nil {
					return err
				}
				next = path
			}
			if err := applyNodes(kind, node.Nodes, next, body, values, preview); err != nil {
				return err
			}
		default:
			return fmt.Errorf("节点 %s 的类型无效", node.Label)
		}
	}
	return nil
}

func applyField(node Node, prefix string, body map[string]any, values map[string]any, preview bool) error {
	value, ok, err := fieldValue(node, values)
	if err != nil {
		if preview {
			return nil
		}
		return err
	}
	if !ok {
		if node.Required && !preview {
			return fmt.Errorf("请填写%s", node.Label)
		}
		return nil
	}
	if !node.ShouldSubmit() {
		return nil
	}
	path, err := resolvePath(prefix, node.Key)
	if err != nil {
		return err
	}
	return setPath(body, path, value)
}

func fieldValue(node Node, values map[string]any) (any, bool, error) {
	raw, ok := values[node.ID]
	if !ok || raw == nil {
		if len(node.Default) > 0 && string(node.Default) != "null" {
			var decoded any
			if err := json.Unmarshal(node.Default, &decoded); err != nil {
				return nil, false, fmt.Errorf("字段 %s 的默认值无效", node.Label)
			}
			raw = decoded
			ok = true
		}
	}
	if !ok || raw == nil {
		return nil, false, nil
	}
	value, err := coerceWidget(node, raw)
	if err != nil {
		return nil, false, err
	}
	if value == nil {
		return nil, false, nil
	}
	return value, true, nil
}

func coerceWidget(node Node, raw any) (any, error) {
	switch node.Widget {
	case "text", "textarea", "image_url", "video_url", "audio_url", "file_url":
		text, ok := raw.(string)
		if !ok {
			return nil, fmt.Errorf("%s 必须是文本", node.Label)
		}
		text = strings.TrimSpace(text)
		if text == "" {
			return nil, nil
		}
		return text, nil
	case "number":
		num, ok := asFloat(raw)
		if !ok {
			return nil, fmt.Errorf("%s 必须是数字", node.Label)
		}
		if err := checkRange(node, num); err != nil {
			return nil, err
		}
		return num, nil
	case "integer":
		num, ok := asFloat(raw)
		if !ok || num != float64(int64(num)) {
			return nil, fmt.Errorf("%s 必须是整数", node.Label)
		}
		if err := checkRange(node, num); err != nil {
			return nil, err
		}
		return int64(num), nil
	case "slider":
		num, ok := asFloat(raw)
		if !ok {
			return nil, fmt.Errorf("%s 必须是数字", node.Label)
		}
		if err := checkRange(node, num); err != nil {
			return nil, err
		}
		if sliderUsesInteger(node) {
			if num != float64(int64(num)) {
				return nil, fmt.Errorf("%s 必须是整数", node.Label)
			}
			return int64(num), nil
		}
		return num, nil
	case "boolean":
		flag, ok := raw.(bool)
		if !ok {
			return nil, fmt.Errorf("%s 必须是布尔值", node.Label)
		}
		return flag, nil
	case "image_urls":
		items, ok := raw.([]any)
		if !ok {
			return nil, fmt.Errorf("%s 必须是图片列表", node.Label)
		}
		out := make([]string, 0, len(items))
		for _, item := range items {
			text, ok := item.(string)
			if !ok || strings.TrimSpace(text) == "" {
				return nil, fmt.Errorf("%s 中的每一项都必须是图片地址", node.Label)
			}
			out = append(out, strings.TrimSpace(text))
		}
		if len(out) == 0 {
			return nil, nil
		}
		return out, nil
	case "messages", "input", "json":
		return raw, nil
	default:
		return nil, fmt.Errorf("字段 %s 的控件类型无效", node.Label)
	}
}

func sliderUsesInteger(node Node) bool {
	if node.Min == nil || node.Max == nil {
		return false
	}
	return *node.Min == float64(int64(*node.Min)) && *node.Max == float64(int64(*node.Max))
}

func checkRange(node Node, num float64) error {
	if node.Min != nil && num < *node.Min {
		return fmt.Errorf("%s 不能小于 %v", node.Label, *node.Min)
	}
	if node.Max != nil && num > *node.Max {
		return fmt.Errorf("%s 不能大于 %v", node.Label, *node.Max)
	}
	return nil
}

func applyChoice(kind APIKind, node Node, prefix string, body map[string]any, values map[string]any, preview bool) error {
	raw, ok := values[node.ID]
	selected := ""
	if ok && raw != nil {
		text, isString := raw.(string)
		if !isString {
			return fmt.Errorf("%s 的选项必须是文本", node.Label)
		}
		selected = text
	}
	if selected == "" {
		if node.Required && !preview {
			return fmt.Errorf("请选择%s", node.Label)
		}
		return nil
	}
	var chosen *Option
	for i := range node.Options {
		if node.Options[i].Value == selected {
			chosen = &node.Options[i]
			break
		}
	}
	if chosen == nil {
		if !(node.Type == NodeSelect && node.AllowCustom) {
			if preview {
				return nil
			}
			return fmt.Errorf("%s 的选项不存在", node.Label)
		}
	}
	if node.ShouldSubmit() {
		path, err := resolvePath(prefix, node.Key)
		if err != nil {
			return err
		}
		if err := setPath(body, path, selected); err != nil {
			return err
		}
	}
	if chosen == nil {
		return nil
	}
	return applyNodes(kind, chosen.Nodes, prefix, body, values, preview)
}

func asFloat(raw any) (float64, bool) {
	switch value := raw.(type) {
	case float64:
		return value, true
	case float32:
		return float64(value), true
	case int:
		return float64(value), true
	case int64:
		return float64(value), true
	case json.Number:
		parsed, err := value.Float64()
		return parsed, err == nil
	default:
		return 0, false
	}
}

func setPath(body map[string]any, path string, value any) error {
	segments := strings.Split(path, ".")
	return setSegments(body, segments, value)
}

func setSegments(current map[string]any, segments []string, value any) error {
	if len(segments) == 0 {
		return fmt.Errorf("路径为空")
	}
	head := segments[0]
	if len(segments) == 1 {
		current[head] = value
		return nil
	}
	if isIndex(segments[1]) {
		index, _ := strconv.Atoi(segments[1])
		existing, _ := current[head].([]any)
		existing = growSlice(existing, index)
		if len(segments) == 2 {
			existing[index] = value
			current[head] = existing
			return nil
		}
		child, ok := existing[index].(map[string]any)
		if !ok || child == nil {
			child = map[string]any{}
			existing[index] = child
		}
		current[head] = existing
		return setSegments(child, segments[2:], value)
	}
	child, ok := current[head].(map[string]any)
	if !ok || child == nil {
		child = map[string]any{}
		current[head] = child
	}
	return setSegments(child, segments[1:], value)
}

func isIndex(seg string) bool {
	if seg == "" {
		return false
	}
	for _, r := range seg {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func growSlice(items []any, index int) []any {
	for len(items) <= index {
		items = append(items, nil)
	}
	return items
}
