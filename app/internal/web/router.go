package web

import (
	"github.com/gin-gonic/gin"
)

// RegisterRouter 注册http路由
func RegisterRouter(r *gin.Engine) {
	plugin := r.Group("/app/core")
	// 获取语法定义
	plugin.GET("/interface/plugin", GetPluginInterface)
	// 获取插件列表
	plugin.GET("/plugins", GetPluginList)
	// 新建插件
	plugin.POST("/plugin", AddPlugin)
	// 获取插件内容
	plugin.GET("/plugin/:name", GetPluginInfo)
	// 更新插件内容
	plugin.PUT("/plugin/:name", UpdatePlugin)
	// 插件重载
	plugin.GET("/plugin/:name/reload", PluginReload)
	// 删除插件
	plugin.DELETE("/plugin/:name", PluginDelete)
	// 插件导出
	plugin.GET("/plugin/export", ExportPlugin)

	// 获取所有管理页面
	plugin.GET("/view/manage", GetManageViews)
	// 获取自定义页面
	plugin.GET("/view/app/:name/:id/:view", GetCustomView)
	// 视图导出功能
	plugin.GET("/view/export", ExportView)

	// 新建页面
	plugin.POST("/views", AddView)
	// 获取所有页面
	plugin.GET("/views", GetViews)
	// 修改页面
	plugin.PUT("/views/:id", UpdateView)
	// 删除页面
	plugin.DELETE("/views/:id", DeleteView)
	// 获取某个页面
	plugin.GET("/views/:id", GetView)
}
