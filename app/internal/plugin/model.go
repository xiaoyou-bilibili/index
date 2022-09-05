package plugin

// ManegeView 管理界面
type ManegeView struct {
	ID   int64                    `json:"id"`   // 页面ID
	Name string                   `json:"name"` // 页面名字
	Data func() map[string]string `json:"data"` // 待映射的数据
}
