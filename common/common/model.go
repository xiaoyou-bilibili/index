package common

// PageInfo 分页信息
type PageInfo struct {
	Current int         `json:"current"` // 当前第几页
	Total   int64       `json:"total"`   // 总页数
	List    interface{} `json:"list"`    // 列表
}
