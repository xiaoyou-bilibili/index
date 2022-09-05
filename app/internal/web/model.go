package web

type View struct {
	Id int64 `json:"id"` // 视图ID
	ViewInfo
}

// ViewInfo 视图信息
type ViewInfo struct {
	Name string `json:"name"` // 视图名称
	View string `json:"view"` // 视图代码
}
