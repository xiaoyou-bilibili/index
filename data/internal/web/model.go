package web

type addObjectResp struct {
	ObjectID string `json:"object_id"` // 对象ID
}

type getTextResp struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

// FileUploadReq 文件上传携带的值
type FileUploadReq struct {
	TotalChunk int32  `json:"totalChunk" form:"totalChunks"` // 总分片大小
	Current    int32  `json:"current" form:"chunkNumber"`    // 当前几个分片
	TotalSize  int64  `json:"totalSize" form:"totalSize"`    // 总大小
	Name       string `json:"name" form:"filename"`          // 文件名
	Data       string `json:"data" form:"data"`              // 文件数据（使用base4）
}
