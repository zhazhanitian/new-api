package doubao

// ──────────────────────────────────────────────────────────────────────────────
// 火山方舟 3D 模型常量
// 参考：docs/3D文档/3D模型渠道接入规范.md §4
// ──────────────────────────────────────────────────────────────────────────────

const (
	ModelSeed3D  = "doubao-seed3d-2-0-260328"
	ModelHyper3D = "hyper3d-gen2-260112"
	ModelHitem3D = "hitem3d-2-0-251223"
)

// is3DModel 判断模型名是否为火山 3D 模型。
func is3DModel(modelName string) bool {
	return modelName == ModelSeed3D ||
		modelName == ModelHyper3D ||
		modelName == ModelHitem3D
}

// fileFormatIntMap 数美 hitem3d 的文件格式整数映射。
// 规范：obj/glb/stl/fbx/usdz → 1/2/3/4/5，默认1
var fileFormatIntMap = map[string]int{
	"obj":  1,
	"glb":  2,
	"stl":  3,
	"fbx":  4,
	"usdz": 5,
}

// viewBitIndex 数美多视角位图：front=位0, back=位1, left=位2, right=位3
var viewBitIndex = map[string]int{
	"front": 0,
	"back":  1,
	"left":  2,
	"right": 3,
}
