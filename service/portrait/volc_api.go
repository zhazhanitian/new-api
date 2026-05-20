package portrait

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/QuantumNous/new-api/setting/portrait_setting"
)

// volcResponse 火山引擎 OpenAPI 通用响应结构
type volcResponse struct {
	ResponseMetadata struct {
		RequestId string `json:"RequestId"`
		Error     *struct {
			Code    string `json:"Code"`
			Message string `json:"Message"`
		} `json:"Error,omitempty"`
	} `json:"ResponseMetadata"`
	Result json.RawMessage `json:"Result,omitempty"`
}

func callVolcAPI(action string, body map[string]interface{}) (json.RawMessage, error) {
	cfg := portrait_setting.VolcPortraitAccessKeyId
	if cfg == "" {
		return nil, fmt.Errorf("火山引擎凭据未配置，请在系统设置中配置 VolcPortraitAccessKeyId")
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %w", err)
	}

	region := portrait_setting.VolcPortraitRegion
	if region == "" {
		region = "cn-beijing"
	}

	headers, queryStr, err := BuildRequest(
		portrait_setting.VolcPortraitAccessKeyId,
		portrait_setting.VolcPortraitSecretAccessKey,
		region,
		action,
		bodyBytes,
	)
	if err != nil {
		return nil, err
	}

	reqURL := fmt.Sprintf("https://%s/?%s", volcHost, queryStr)
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("构造请求失败: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求火山引擎失败: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	var volcResp volcResponse
	if err = json.Unmarshal(respBytes, &volcResp); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w, 原始响应: %s", err, string(respBytes))
	}

	if volcResp.ResponseMetadata.Error != nil {
		e := volcResp.ResponseMetadata.Error
		return nil, fmt.Errorf("[%s] %s (RequestId: %s)", e.Code, e.Message, volcResp.ResponseMetadata.RequestId)
	}

	return volcResp.Result, nil
}

// CreateAssetGroupResult 创建素材组响应
type CreateAssetGroupResult struct {
	Id string `json:"Id"`
}

// CreateAssetGroup 在火山引擎创建素材组
func CreateAssetGroup(name, description, groupType string) (*CreateAssetGroupResult, error) {
	projectName := portrait_setting.VolcPortraitProjectName
	body := map[string]interface{}{
		"Name":        name,
		"GroupType":   groupType,
		"ProjectName": projectName,
	}
	if description != "" {
		body["Description"] = description
	}

	raw, err := callVolcAPI("CreateAssetGroup", body)
	if err != nil {
		return nil, err
	}

	var result CreateAssetGroupResult
	if err = json.Unmarshal(raw, &result); err != nil {
		// 部分响应 Id 可能在其他字段，兜底尝试
		var fallback map[string]interface{}
		if json.Unmarshal(raw, &fallback) == nil {
			for _, key := range []string{"Id", "id", "GroupId", "groupId"} {
				if v, ok := fallback[key].(string); ok && v != "" {
					result.Id = v
					break
				}
			}
		}
	}
	if result.Id == "" {
		return nil, fmt.Errorf("CreateAssetGroup 未返回有效 Id，原始响应: %s", string(raw))
	}
	return &result, nil
}

// ListAssetGroupsResult 查询素材组响应
type ListAssetGroupsResult struct {
	TotalCount int `json:"TotalCount"`
	Items      []struct {
		Id          string `json:"Id"`
		Name        string `json:"Name"`
		ProjectName string `json:"ProjectName"`
		CreateTime  string `json:"CreateTime"`
		UpdateTime  string `json:"UpdateTime"`
	} `json:"Items"`
}

// ListAssetGroups 查询火山引擎上的素材组（按 projectName 过滤）
func ListAssetGroups(page, pageSize int) (*ListAssetGroupsResult, error) {
	projectName := portrait_setting.VolcPortraitProjectName
	body := map[string]interface{}{
		"PageNumber":  page,
		"PageSize":    pageSize,
		"Filter":      map[string]interface{}{"GroupType": "AIGC"},
		"ProjectName": projectName,
	}

	raw, err := callVolcAPI("ListAssetGroups", body)
	if err != nil {
		return nil, err
	}

	var result ListAssetGroupsResult
	if err = json.Unmarshal(raw, &result); err != nil {
		return nil, fmt.Errorf("解析 ListAssetGroups 响应失败: %w", err)
	}
	return &result, nil
}

// CreateAssetResult 创建素材响应
type CreateAssetResult struct {
	AssetId string
}

// CreateAsset 在火山引擎注册素材
func CreateAsset(groupId, assetURL, assetType, name string) (*CreateAssetResult, error) {
	projectName := portrait_setting.VolcPortraitProjectName
	body := map[string]interface{}{
		"GroupId":     groupId,
		"URL":         assetURL,
		"AssetType":   assetType,
		"ProjectName": projectName,
	}
	if name != "" {
		body["Name"] = name
	}

	raw, err := callVolcAPI("CreateAsset", body)
	if err != nil {
		return nil, err
	}

	var fallback map[string]interface{}
	if err = json.Unmarshal(raw, &fallback); err != nil {
		return nil, fmt.Errorf("解析 CreateAsset 响应失败: %w", err)
	}

	var assetId string
	for _, key := range []string{"Id", "id", "AssetId", "assetId"} {
		if v, ok := fallback[key].(string); ok && v != "" {
			assetId = v
			break
		}
	}
	if assetId == "" {
		return nil, fmt.Errorf("CreateAsset 未返回有效 AssetId，原始响应: %s", string(raw))
	}
	return &CreateAssetResult{AssetId: assetId}, nil
}

// GetAssetResult 获取素材状态响应
type GetAssetResult struct {
	Status      string `json:"status"`
	ResolvedUrl string `json:"resolved_url"`
	Error       string `json:"error"`
}

// GetAsset 查询素材审核状态
func GetAsset(remoteAssetId string) (*GetAssetResult, error) {
	projectName := portrait_setting.VolcPortraitProjectName
	body := map[string]interface{}{
		"Id":          remoteAssetId,
		"ProjectName": projectName,
	}

	raw, err := callVolcAPI("GetAsset", body)
	if err != nil {
		return nil, err
	}

	var flat map[string]interface{}
	if err = json.Unmarshal(raw, &flat); err != nil {
		return nil, fmt.Errorf("解析 GetAsset 响应失败: %w", err)
	}

	result := &GetAssetResult{}
	for _, key := range []string{"Status", "status"} {
		if v, ok := flat[key].(string); ok {
			result.Status = v
			break
		}
	}
	for _, key := range []string{"URL", "Url", "url"} {
		if v, ok := flat[key].(string); ok {
			result.ResolvedUrl = v
			break
		}
	}
	for _, key := range []string{"Error", "error"} {
		if v, ok := flat[key].(string); ok {
			result.Error = v
			break
		}
	}
	return result, nil
}
