package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/service"
	"github.com/QuantumNous/new-api/trial"

	"github.com/gin-gonic/gin"
)

type modelTrialSubmit struct {
	Group  string         `json:"group"`
	Values map[string]any `json:"values"`
}

// PrepareModelTrial 按已保存的节点树组装固定接口请求，再交给渠道分发。
func PrepareModelTrial() gin.HandlerFunc {
	return func(c *gin.Context) {
		modelName := strings.Trim(strings.TrimPrefix(c.Param("model"), "/"), "/")
		if modelName == "" {
			abortWithOpenAiMessage(c, http.StatusBadRequest, "缺少模型名")
			return
		}
		item, err := model.GetModelTrialByName(modelName)
		if err != nil {
			abortWithOpenAiMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		if item == nil || !item.Enabled {
			abortWithOpenAiMessage(c, http.StatusNotFound, "该模型还没有可用的在线体验")
			return
		}
		var req modelTrialSubmit
		if err := common.UnmarshalBodyReusable(c, &req); err != nil {
			abortWithOpenAiMessage(c, http.StatusBadRequest, "请求体无效")
			return
		}
		var nodes []trial.Node
		if err := common.Unmarshal([]byte(item.Nodes), &nodes); err != nil {
			abortWithOpenAiMessage(c, http.StatusBadRequest, "体验配置无效")
			return
		}
		body, err := trial.Assemble(trial.APIKind(item.APIKind), item.ModelName, nodes, req.Values)
		if err != nil {
			abortWithOpenAiMessage(c, http.StatusBadRequest, err.Error())
			return
		}
		raw, err := json.Marshal(body)
		if err != nil {
			abortWithOpenAiMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		path, err := trial.CanonicalPath(trial.APIKind(item.APIKind))
		if err != nil {
			abortWithOpenAiMessage(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := common.ReplaceRequestBody(c, raw); err != nil {
			abortWithOpenAiMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.Request.URL.Path = path
		// 路径改成正式接口后，计费仍按操练场处理：用登录账号扣费，不查 API Key。
		common.SetContextKey(c, constant.ContextKeyPlayground, true)
		usingGroup := common.GetContextKeyString(c, constant.ContextKeyUsingGroup)
		if req.Group != "" {
			if !service.GroupInUserUsableGroups(usingGroup, req.Group) && req.Group != usingGroup {
				abortWithOpenAiMessage(c, http.StatusForbidden, "无权使用该分组")
				return
			}
			usingGroup = req.Group
			common.SetContextKey(c, constant.ContextKeyUsingGroup, usingGroup)
			common.SetContextKey(c, constant.ContextKeyTokenGroup, usingGroup)
		}
		hash, err := trial.HashValues(req.Values)
		if err != nil {
			abortWithOpenAiMessage(c, http.StatusBadRequest, err.Error())
			return
		}
		proof, err := model.UpsertPendingTrialProof(c.GetInt("id"), item.ModelName, hash)
		if err != nil {
			abortWithOpenAiMessage(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.Set(trial.CtxModelName, item.ModelName)
		c.Set(trial.CtxValuesHash, hash)
		c.Set(trial.CtxAPIKind, item.APIKind)
		c.Set(trial.CtxProofID, proof.Id)
		c.Next()
	}
}
