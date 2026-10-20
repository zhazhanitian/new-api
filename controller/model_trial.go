package controller

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/trial"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type modelTrialPayload struct {
	ModelName string          `json:"model_name"`
	APIKind   string          `json:"api_kind"`
	Enabled   *bool           `json:"enabled"`
	Nodes     json.RawMessage `json:"nodes"`
	Examples  json.RawMessage `json:"examples"`
}

type modelTrialExamplePayload struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Values      map[string]any `json:"values"`
}

type modelTrialView struct {
	Id          int                 `json:"id"`
	ModelName   string              `json:"model_name"`
	APIKind     string              `json:"api_kind"`
	Enabled     bool                `json:"enabled"`
	Tags        string              `json:"tags,omitempty"`
	Nodes       []trial.Node        `json:"nodes"`
	Examples    []modelTrialExample `json:"examples"`
	CreatedTime int64               `json:"created_time"`
	UpdatedTime int64               `json:"updated_time"`
}

type modelTrialExample struct {
	trial.Example
	Stale bool `json:"stale"`
}

func ListModelTrials(c *gin.Context) {
	items, err := model.ListModelTrials()
	if err != nil {
		common.ApiError(c, err)
		return
	}
	views := make([]modelTrialView, 0, len(items))
	for _, item := range items {
		views = append(views, toModelTrialView(item))
	}
	common.ApiSuccess(c, views)
}

func ListEnabledModelTrials(c *gin.Context) {
	items, err := model.ListEnabledModelTrials()
	if err != nil {
		common.ApiError(c, err)
		return
	}
	views := make([]modelTrialView, 0, len(items))
	for _, item := range items {
		views = append(views, toModelTrialView(item))
	}
	if err := attachModelTags(views); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, views)
}

func GetModelTrial(c *gin.Context) {
	item, ok := loadModelTrial(c)
	if !ok {
		return
	}
	common.ApiSuccess(c, toModelTrialView(item))
}

func CreateModelTrial(c *gin.Context) {
	item, err := trialFromPayload(c, nil)
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}
	existing, err := model.GetModelTrialByName(item.ModelName)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	if existing != nil {
		common.ApiErrorMsg(c, "该模型已经有体验配置")
		return
	}
	if err := item.Insert(); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, toModelTrialView(item))
}

func UpdateModelTrial(c *gin.Context) {
	current, ok := loadModelTrial(c)
	if !ok {
		return
	}
	item, err := trialFromPayload(c, current)
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}
	item.Id = current.Id
	item.CreatedTime = current.CreatedTime
	if item.Examples == "" {
		item.Examples = current.Examples
	}
	other, err := model.GetModelTrialByName(item.ModelName)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	if other != nil && other.Id != current.Id {
		common.ApiErrorMsg(c, "该模型已经有体验配置")
		return
	}
	if err := item.Update(); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, toModelTrialView(item))
}

func DeleteModelTrial(c *gin.Context) {
	item, ok := loadModelTrial(c)
	if !ok {
		return
	}
	if err := model.DeleteModelTrial(item.Id); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, nil)
}

func SaveModelTrialExample(c *gin.Context) {
	item, ok := loadModelTrial(c)
	if !ok {
		return
	}
	var req modelTrialExamplePayload
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ApiErrorMsg(c, "请求体无效")
		return
	}
	if strings.TrimSpace(req.Title) == "" {
		common.ApiErrorMsg(c, "请填写示例标题")
		return
	}
	nodes, err := decodeTrialNodes(item.Nodes)
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}
	if _, err := trial.Assemble(trial.APIKind(item.APIKind), item.ModelName, nodes, req.Values); err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}
	hash, err := trial.HashValues(req.Values)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	verified, err := model.HasVerifiedTrialProof(c.GetInt("id"), item.ModelName, hash)
	if err != nil {
		common.ApiError(c, err)
		return
	}
	if !verified {
		common.ApiError(c, errTrialNotVerified)
		return
	}
	examples, err := decodeTrialExamples(item.Examples)
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}
	examples = append(examples, trial.Example{
		ID:          uuid.NewString(),
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		Values:      req.Values,
		VerifiedAt:  common.GetTimestamp(),
	})
	if err := saveTrialExamples(item, examples); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, toModelTrialView(item))
}

func UpdateModelTrialExample(c *gin.Context) {
	item, ok := loadModelTrial(c)
	if !ok {
		return
	}
	var req modelTrialExamplePayload
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ApiErrorMsg(c, "请求体无效")
		return
	}
	examples, err := decodeTrialExamples(item.Examples)
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}
	exampleID := c.Param("exampleId")
	found := false
	for i := range examples {
		if examples[i].ID != exampleID {
			continue
		}
		if strings.TrimSpace(req.Title) != "" {
			examples[i].Title = strings.TrimSpace(req.Title)
		}
		examples[i].Description = strings.TrimSpace(req.Description)
		if req.Values != nil {
			nodes, nodeErr := decodeTrialNodes(item.Nodes)
			if nodeErr != nil {
				common.ApiErrorMsg(c, nodeErr.Error())
				return
			}
			if _, assembleErr := trial.Assemble(trial.APIKind(item.APIKind), item.ModelName, nodes, req.Values); assembleErr != nil {
				common.ApiErrorMsg(c, assembleErr.Error())
				return
			}
			hash, hashErr := trial.HashValues(req.Values)
			if hashErr != nil {
				common.ApiError(c, hashErr)
				return
			}
			verified, verifyErr := model.HasVerifiedTrialProof(c.GetInt("id"), item.ModelName, hash)
			if verifyErr != nil {
				common.ApiError(c, verifyErr)
				return
			}
			if !verified {
				common.ApiError(c, errTrialNotVerified)
				return
			}
			examples[i].Values = req.Values
			examples[i].VerifiedAt = common.GetTimestamp()
		}
		found = true
		break
	}
	if !found {
		common.ApiErrorMsg(c, "示例不存在")
		return
	}
	if err := saveTrialExamples(item, examples); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, toModelTrialView(item))
}

func DeleteModelTrialExample(c *gin.Context) {
	item, ok := loadModelTrial(c)
	if !ok {
		return
	}
	examples, err := decodeTrialExamples(item.Examples)
	if err != nil {
		common.ApiErrorMsg(c, err.Error())
		return
	}
	exampleID := c.Param("exampleId")
	next := make([]trial.Example, 0, len(examples))
	found := false
	for _, example := range examples {
		if example.ID == exampleID {
			found = true
			continue
		}
		next = append(next, example)
	}
	if !found {
		common.ApiErrorMsg(c, "示例不存在")
		return
	}
	if err := saveTrialExamples(item, next); err != nil {
		common.ApiError(c, err)
		return
	}
	common.ApiSuccess(c, toModelTrialView(item))
}

func loadModelTrial(c *gin.Context) (*model.ModelTrial, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.ApiErrorMsg(c, "配置不存在")
		return nil, false
	}
	item, err := model.GetModelTrialByID(id)
	if err != nil {
		common.ApiErrorMsg(c, "配置不存在")
		return nil, false
	}
	return item, true
}

func trialFromPayload(c *gin.Context, current *model.ModelTrial) (*model.ModelTrial, error) {
	var req modelTrialPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	item := &model.ModelTrial{}
	if current != nil {
		*item = *current
	}
	if strings.TrimSpace(req.ModelName) != "" {
		item.ModelName = strings.TrimSpace(req.ModelName)
	}
	if req.APIKind != "" {
		item.APIKind = req.APIKind
	}
	if req.Enabled != nil {
		item.Enabled = *req.Enabled
	} else if current == nil {
		item.Enabled = true
	}
	if len(req.Nodes) > 0 {
		var nodes []trial.Node
		if err := json.Unmarshal(req.Nodes, &nodes); err != nil {
			return nil, err
		}
		if err := trial.Validate(trial.APIKind(item.APIKind), nodes); err != nil {
			return nil, err
		}
		raw, err := json.Marshal(nodes)
		if err != nil {
			return nil, err
		}
		item.Nodes = string(raw)
	}
	if len(req.Examples) > 0 {
		var examples []trial.Example
		if err := json.Unmarshal(req.Examples, &examples); err != nil {
			return nil, err
		}
		raw, err := json.Marshal(examples)
		if err != nil {
			return nil, err
		}
		item.Examples = string(raw)
	}
	if strings.TrimSpace(item.ModelName) == "" {
		return nil, errMissingTrialModel
	}
	if item.Nodes == "" {
		return nil, errMissingTrialNodes
	}
	return item, nil
}

var (
	errMissingTrialModel = errString("请选择模型")
	errMissingTrialNodes = errString("请配置节点")
)

type errString string

func (e errString) Error() string { return string(e) }

func attachModelTags(views []modelTrialView) error {
	names := make([]string, 0, len(views))
	for _, view := range views {
		names = append(names, view.ModelName)
	}
	tags, err := model.GetModelTagsByNames(names)
	if err != nil {
		return err
	}
	for i := range views {
		views[i].Tags = tags[views[i].ModelName]
	}
	return nil
}

func toModelTrialView(item *model.ModelTrial) modelTrialView {
	nodes, _ := decodeTrialNodes(item.Nodes)
	examples, _ := decodeTrialExamples(item.Examples)
	viewExamples := make([]modelTrialExample, 0, len(examples))
	for _, example := range examples {
		viewExamples = append(viewExamples, modelTrialExample{
			Example: example,
			Stale:   trial.ExampleStale(nodes, example.Values),
		})
	}
	if nodes == nil {
		nodes = []trial.Node{}
	}
	return modelTrialView{
		Id:          item.Id,
		ModelName:   item.ModelName,
		APIKind:     item.APIKind,
		Enabled:     item.Enabled,
		Nodes:       nodes,
		Examples:    viewExamples,
		CreatedTime: item.CreatedTime,
		UpdatedTime: item.UpdatedTime,
	}
}

func decodeTrialNodes(raw string) ([]trial.Node, error) {
	if strings.TrimSpace(raw) == "" {
		return []trial.Node{}, nil
	}
	var nodes []trial.Node
	if err := json.Unmarshal([]byte(raw), &nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}

func decodeTrialExamples(raw string) ([]trial.Example, error) {
	if strings.TrimSpace(raw) == "" {
		return []trial.Example{}, nil
	}
	var examples []trial.Example
	if err := json.Unmarshal([]byte(raw), &examples); err != nil {
		return nil, err
	}
	return examples, nil
}

func saveTrialExamples(item *model.ModelTrial, examples []trial.Example) error {
	raw, err := json.Marshal(examples)
	if err != nil {
		return err
	}
	item.Examples = string(raw)
	return item.Update()
}
