package tencentvod

// ──────────────────────────────────────────────────────────────────────────────
// 腾讯混元 3D AI3D 服务常量
// 参考：docs/3D文档/3D模型渠道接入规范.md §5
// ──────────────────────────────────────────────────────────────────────────────

const (
	tencentAI3DHost    = "ai3d.tencentcloudapi.com"
	tencentAI3DService = "ai3d"
	tencentAI3DVersion = "2025-05-13"
	tencentAI3DRegion  = "ap-guangzhou"
)

// ai3DModelMeta 模型名 → 提交/查询 Action 及上游 Model 参数
type ai3DModelMetaEntry struct {
	SubmitAction string
	QueryAction  string
	ModelVersion string // 上游 Model 字段值；空字符串表示不传
	IsProVersion bool   // 是否为专业版（查询返回 ResultCreditConsumed，允许完成时积分校准）
}

var ai3DModelMeta = map[string]ai3DModelMetaEntry{
	"hunyuan-3d-rapid": {
		SubmitAction: "SubmitHunyuanTo3DRapidJob",
		QueryAction:  "QueryHunyuanTo3DRapidJob",
	},
	"hunyuan-3d-pro-3.0": {
		SubmitAction: "SubmitHunyuanTo3DProJob",
		QueryAction:  "QueryHunyuanTo3DProJob",
		ModelVersion: "3.0",
		IsProVersion: true,
	},
	"hunyuan-3d-pro-3.1": {
		SubmitAction: "SubmitHunyuanTo3DProJob",
		QueryAction:  "QueryHunyuanTo3DProJob",
		ModelVersion: "3.1",
		IsProVersion: true,
	},
	"hunyuan-3d-reduce-face": {
		SubmitAction: "SubmitReduceFaceJob",
		QueryAction:  "DescribeReduceFaceJob",
	},
	"hunyuan-3d-texture-3.0": {
		SubmitAction: "SubmitTextureTo3DJob",
		QueryAction:  "DescribeTextureTo3DJob",
		ModelVersion: "3.0",
	},
	"hunyuan-3d-texture-3.1": {
		SubmitAction: "SubmitTextureTo3DJob",
		QueryAction:  "DescribeTextureTo3DJob",
		ModelVersion: "3.1",
	},
	"hunyuan-3d-profile": {
		SubmitAction: "SubmitProfileTo3DJob",
		QueryAction:  "DescribeProfileTo3DJob",
	},
	"hunyuan-3d-auto-rigging": {
		SubmitAction: "SubmitAutoRiggingJob",
		QueryAction:  "DescribeAutoRiggingJob",
	},
	"hunyuan-3d-motion": {
		SubmitAction: "SubmitHunyuanTo3DMotionJob",
		QueryAction:  "DescribeHunyuanTo3DMotionJob",
		ModelVersion: "HY-Motion-1.0",
	},
}

// getAI3DModelMeta 查找模型元数据，找不到则返回 false。
func getAI3DModelMeta(modelName string) (ai3DModelMetaEntry, bool) {
	m, ok := ai3DModelMeta[modelName]
	return m, ok
}

// is3DModel 判断模型名是否为腾讯 AI3D 模型。
func isAI3DModel(modelName string) bool {
	_, ok := ai3DModelMeta[modelName]
	return ok
}

// ai3DModelList 腾讯 AI3D 模型名列表，供 GetModelList 使用。
var ai3DModelList = func() []string {
	list := make([]string, 0, len(ai3DModelMeta))
	for k := range ai3DModelMeta {
		list = append(list, k)
	}
	return list
}()

// ──────────────────────────────────────────────────────────────────────────────
// 积分估算常量
// 参考：docs/3D文档/3D模型渠道接入规范.md §6.3
// ──────────────────────────────────────────────────────────────────────────────

// 固定积分（极速版基础、后处理类模型）
const (
	creditsRapidBase      = 15
	creditsRapidPBR       = 10
	creditsReduceFace     = 50
	creditsTexture        = 30
	creditsProfile        = 30
	creditsAutoRigging    = 10
	creditsMotion         = 10
	// 专业版基础积分
	creditsProNormal      = 20
	creditsProLowPoly     = 25
	creditsProGeometry    = 15
	creditsProSketch      = 25
	// 专业版附加积分
	creditsProMultiView   = 10
	creditsProPBR         = 10
	creditsProFaceCount   = 10
	creditsProFileFormat  = 5
)
