package dto

// ──────────────────────────────────────────────────────────────────────────────
// 3D 任务统一 DTO
// 字段来源：docs/3D文档/3D模型渠道接入规范.md §3.1 / §3.2
// ──────────────────────────────────────────────────────────────────────────────

// Image3D 统一图片结构，同时承载 URL/Base64 和多视角语义。
// url 字段接受 HTTP URL 或 data:image/<格式>;base64,<数据> 两种形式。
// view 枚举：front/back/left/right/top/bottom/left_front/right_front
type Image3D struct {
	URL  string `json:"url"`
	View string `json:"view,omitempty"`
}

// Input3DFile 统一 3D 文件输入结构。
// 智能拓扑/纹理/绑骨蒙皮/文生动作均通过此结构传入上游文件；
// 对应上游字段：智能拓扑→File3D(File3D结构)，纹理→File3D(File3D结构)，
// 绑骨蒙皮→File3D(InputFile3D结构)，文生动作→RetargetFile(InputFile3D结构)。
type Input3DFile struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}

// OpenAI3DRequest 统一 3D 任务提交请求。
// model 在顶层传入；其余字段通过 metadata 传入（与现有视频/图片任务格式一致）。
type OpenAI3DRequest struct {
	Model    string                 `json:"model"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// OpenAI3DMetadata 从 Metadata map 解析出的强类型结构。
// 适配器内部使用，不作为对外协议。
type OpenAI3DMetadata struct {
	Prompt    string      `json:"prompt,omitempty"`
	Images    []Image3D   `json:"images,omitempty"`
	InputFile *Input3DFile `json:"input_file,omitempty"`

	// 输出格式：各模型枚举不同，适配器按模型分支转换
	FileFormat string `json:"file_format,omitempty"`

	// 腾讯极速/专业/纹理生成支持
	EnablePBR *bool `json:"enable_pbr,omitempty"`

	// 数美/腾讯极速/专业版支持
	GeometryOnly *bool `json:"geometry_only,omitempty"`

	// 影眸专用材质：pbr/shaded/all/none
	Material string `json:"material,omitempty"`

	// 面数控制
	FaceCount *int `json:"face_count,omitempty"`

	// 质量级别：high/medium/low（seed3d/影眸）
	QualityLevel string `json:"quality_level,omitempty"`

	// 数美专用分辨率：1536/1536pro
	Resolution string `json:"resolution,omitempty"`

	// 影眸专用网格类型：raw/quad
	MeshMode string `json:"mesh_mode,omitempty"`

	// 影眸专用随机种子：[0,65535]
	Seed *int64 `json:"seed,omitempty"`

	// 腾讯专业版生成模式：normal/low_poly/sketch
	GenerateMode string `json:"generate_mode,omitempty"`

	// 多边形类型：triangle/quadrilateral（腾讯专业LowPoly/智能拓扑）
	PolygonType string `json:"polygon_type,omitempty"`

	// 智能拓扑面数级别：high/medium/low
	FaceLevel string `json:"face_level,omitempty"`

	// 3D人物生成模板枚举（完整枚举见规范§5.5）
	Template string `json:"template,omitempty"`

	// 绑骨蒙皮动作模板：[1,48]
	MotionType *int `json:"motion_type,omitempty"`

	// 纹理生成贴图边长：[720,4096]
	TextureSize *int `json:"texture_size,omitempty"`

	// 纹理生成：保留UV
	EnableKeepUV *bool `json:"enable_keep_uv,omitempty"`

	// 文生动作专用字段
	Duration          *int  `json:"duration,omitempty"`
	EnableMesh        *bool `json:"enable_mesh,omitempty"`
	EnableRewrite     *bool `json:"enable_rewrite,omitempty"`
	EnableDurationEst *bool `json:"enable_duration_est,omitempty"`

	// 影眸专用字段
	UseOriginalAlpha *bool    `json:"use_original_alpha,omitempty"`
	HdTexture        *bool    `json:"hd_texture,omitempty"`
	Addons           string   `json:"addons,omitempty"`
	TaPose           *bool    `json:"ta_pose,omitempty"`
	BboxCondition    []int    `json:"bbox_condition,omitempty"`
}

// ──────────────────────────────────────────────────────────────────────────────
// 查询响应 DTO
// ──────────────────────────────────────────────────────────────────────────────

// Task3DFile 单个 3D 结果文件。
// 火山：content.file_url → url，format 从请求快照或查询字段回填。
// 腾讯：ResultFile3Ds[].Url → url，Type 转小写 → format，PreviewImageUrl → preview_url。
type Task3DFile struct {
	URL        string `json:"url"`
	Format     string `json:"format,omitempty"`
	PreviewURL string `json:"preview_url,omitempty"`
}

// Task3DUsage 任务用量。
// 火山：total_tokens 填写；腾讯非专业版为0；腾讯专业版填 credits。
type Task3DUsage struct {
	InputTokens  int     `json:"input_tokens"`
	OutputTokens int     `json:"output_tokens"`
	TotalTokens  int     `json:"total_tokens"`
	Credits      float64 `json:"credits,omitempty"`
}

// OpenAI3DTask 统一 3D 任务查询响应。
type OpenAI3DTask struct {
	TaskID     string       `json:"task_id"`
	Object     string       `json:"object"`
	Model      string       `json:"model"`
	Status     string       `json:"status"`
	Progress   string       `json:"progress,omitempty"`
	Files      []Task3DFile `json:"files"`
	FailReason string       `json:"fail_reason,omitempty"`
	Usage      *Task3DUsage `json:"usage,omitempty"`
	Amount     string       `json:"amount"`
	CreatedAt  int64        `json:"created_at"`
	SubmitTime int64        `json:"submit_time"`
	FinishTime int64        `json:"finish_time,omitempty"`
}
