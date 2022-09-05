package web

import (
	"git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/tool/tools"
	"git.xiaoyou.host/index/common/web"
	"github.com/gin-gonic/gin"
	"index.app/internal/plugin"
	"index.app/internal/server"
	"io/ioutil"
)

func GetPlugin(key string, ctx *gin.Context) *plugin.ScriptContainer {
	info := plugin.GetPluginInfo(ctx.Param(key))
	if info == nil {
		web.Fail(ctx, "no plugin found")
		return nil
	}
	return info
}

// AddPlugin 新建插件
func AddPlugin(ctx *gin.Context) {
	var req server.PluginInfo
	if !tools.GinBindData(ctx, &req) {
		return
	}
	code := req.Code
	// 如果插件已经存在就提示已存在
	info := plugin.GetPluginInfo(req.Unique)
	if info != nil || req.Unique == "" {
		web.Fail(ctx, "标识非法或已存在")
		return
	}
	// 保存页面代码
	var err error
	req.Code, err = server.DataServer.AddText(ctx, req.Name, req.Code)
	if err != nil {
		web.Fail(ctx, "add view err %v", err)
		return
	}
	id, err := server.RelationServer.AddNode(ctx, server.NodePlugin, server.RelationServer.GetMapFromStruct(req))
	if err != nil {
		web.Fail(ctx, "add plugin node err %v", err)
		return
	}
	// 添加新插件
	if err = plugin.AddNewPlugin(id, req.Name, req.Unique, code, req.Code); err != nil {
		web.Fail(ctx, "start plugin err %v", err)
		return
	}
	web.Success(ctx, map[string]interface{}{"id": id})
}

// GetPluginInterface 获取插件语法定义
func GetPluginInterface(ctx *gin.Context) {
	// 直接读取配置文件
	data, err := ioutil.ReadFile("plugins/core.d.ts")
	if err != nil {
		web.Fail(ctx, "read data error %v", err)
		return
	}
	web.Success(ctx, map[string]string{"data": string(data)})
}

// GetPluginList 获取插件列表
func GetPluginList(context *gin.Context) {
	web.Success(context, plugin.GetPluginList())
}

// GetPluginInfo 获取插件信息
func GetPluginInfo(ctx *gin.Context) {
	container := GetPlugin("name", ctx)
	if container == nil {
		return
	}
	web.Success(ctx, server.PluginInfo{
		Name:   container.GetName(),
		Unique: container.GetUnique(),
		Code:   container.GetCode(),
	})
}

// UpdatePlugin 更新插件内容
func UpdatePlugin(ctx *gin.Context) {
	var req server.PluginInfo
	if !tools.GinBindData(ctx, &req) {
		return
	}
	container := GetPlugin("name", ctx)
	if container == nil {
		return
	}
	if err := container.Update(ctx, req); err != nil {
		web.Fail(ctx, "update plugin err %v", err)
		return
	}
	// 自动重载内容
	err := container.Reload()
	if err != nil {
		web.Fail(ctx, "reload plugin err %v", err)
		return
	}
	web.Success(ctx, nil)
}

// PluginReload 插件重载
func PluginReload(ctx *gin.Context) {
	container := GetPlugin("name", ctx)
	if container == nil {
		return
	}
	err := container.Reload()
	if err != nil {
		web.Fail(ctx, "%v", err)
		return
	}
	web.Success(ctx, nil)
}

// PluginDelete 删除插件
func PluginDelete(ctx *gin.Context) {
	err := plugin.DeletePlugin(ctx, ctx.Param("name"))
	if err != nil {
		web.Fail(ctx, "%v", err)
		return
	}
	web.Success(ctx, nil)
}

// ExportPlugin 导出所有插件
func ExportPlugin(ctx *gin.Context) {
	// 获取所有的页面
	data, err := server.ExportNode(ctx, server.NodePlugin, "code", "js")
	if err != nil {
		web.Fail(ctx, "export file err %v", err)
		return
	}
	ctx.Header("content-disposition", `attachment; filename=plugin.zip`)
	ctx.Data(200, "application/octet-stream", data)
}

// AddView 新建页面
func AddView(ctx *gin.Context) {
	var req ViewInfo
	if !tools.GinBindData(ctx, &req) {
		return
	}
	// 保存页面代码
	id, err := server.DataServer.AddText(ctx, req.Name, req.View)
	if err != nil {
		web.Fail(ctx, "add view err %v", err)
		return
	}
	// 保存关系
	req.View = id
	id2, err := server.RelationServer.AddNode(ctx, server.NodeView, server.RelationServer.GetMapFromStruct(req))
	if err != nil {
		web.Fail(ctx, "add view node err %v", err)
		return
	}
	web.Success(ctx, map[string]interface{}{"id": id2})
}

// GetViews 获取所有页面
func GetViews(ctx *gin.Context) {
	data, _, err := server.RelationServer.FindNode(ctx, server.NodeView, "name", ".*", 0, 0)
	if err != nil {
		web.Fail(ctx, "get node err %v", err)
		return
	}
	res := make([]*View, 0, len(data))
	for _, v := range data {
		attribute := v.Info.Attribute
		view := &View{
			Id:       v.Id,
			ViewInfo: ViewInfo{Name: attribute["name"], View: attribute["view"]},
		}
		res = append(res, view)
	}
	web.Success(ctx, res)
}

// UpdateView 修改页面
func UpdateView(ctx *gin.Context) {
	id, ok := tools.GinGetFiledInt(ctx, "id")
	if !ok {
		return
	}
	var req ViewInfo
	if !tools.GinBindData(ctx, &req) {
		return
	}
	// 获取页面id
	row, err := server.RelationServer.GetNode(ctx, int64(id))
	if err != nil {
		web.Fail(ctx, "get node err %v", err)
		return
	}
	var node ViewInfo
	_ = server.RelationServer.GetStructFromMap(row, &node)
	// 更新文本
	if err = server.DataServer.UpdateText(ctx, node.View, req.Name, req.View); err != nil {
		web.Fail(ctx, "update text err %v", err)
		return
	}
	// 如果名字没变就不更新页面
	if req.Name != node.Name {
		if err = server.RelationServer.UpdateNode(ctx, int64(id), map[string]string{"name": req.Name}); err != nil {
			web.Fail(ctx, "update view err %v", err)
			return
		}
	}
	web.Success(ctx, nil)
}

// DeleteView 删除页面
func DeleteView(ctx *gin.Context) {
	id, ok := tools.GinGetFiledInt(ctx, "id")
	if !ok {
		return
	}
	// 获取页面id
	row, err := server.RelationServer.GetNode(ctx, int64(id))
	if err != nil {
		web.Fail(ctx, "get node err %v", err)
		return
	}
	var node ViewInfo
	_ = server.RelationServer.GetStructFromMap(row, &node)
	if err = server.DataServer.DeleteObject(ctx, []string{node.View}); err != nil {
		log.CtxLogError(ctx, "update node err %v", err)
	}
	if err = server.RelationServer.DeleteNode(ctx, []int64{int64(id)}); err != nil {
		web.Fail(ctx, "delete view err %v", err)
		return
	}
	web.Success(ctx, nil)
}

// GetView 获取某个页面
func GetView(ctx *gin.Context) {
	id, ok := tools.GinGetFiledInt(ctx, "id")
	if !ok {
		return
	}
	// 获取页面id
	row, err := server.RelationServer.GetNode(ctx, int64(id))
	if err != nil {
		web.Fail(ctx, "get node err %v", err)
		return
	}
	var node ViewInfo
	_ = server.RelationServer.GetStructFromMap(row, &node)
	_, content, err := server.DataServer.GetText(ctx, node.View)
	if err != nil {
		web.Fail(ctx, "get text err %v", err)
		return
	}
	web.Success(ctx, ViewInfo{Name: node.Name, View: content})
}

// GetManageViews 获取所有管理界面
func GetManageViews(ctx *gin.Context) {
	var manageViewList []*map[string]interface{}
	// 遍历所有插件
	plugins := plugin.GetPluginList()
	for _, v := range plugins {
		if container := plugin.GetPluginInfo(v.Unique); container != nil {
			for _, view := range container.GetAllManageViews() {
				manageViewList = append(manageViewList, &map[string]interface{}{
					"id":   view.ID,
					"name": view.Name,
				})
			}
		}
	}
	web.Success(ctx, manageViewList)
}

// GetCustomView 获取自定义页面
func GetCustomView(ctx *gin.Context) {
	// 分别获取应用名称，资源id和展示类型
	name := ctx.Param("name")
	view := ctx.Param("view")
	if name == "" || view == "" {
		web.Fail(ctx, "param err")
	}
	id, ok := tools.GinGetFiledInt(ctx, "id")
	if !ok {
		return
	}
	// 首先寻找应用
	container := plugin.GetPluginInfo(name)
	if container == nil {
		web.Fail(ctx, "no plugin found")
		return
	}
	web.Success(ctx, container.GetCustomView(ctx, int64(id), view))
}

// ExportView 导出所有视图
func ExportView(ctx *gin.Context) {
	// 获取所有的页面
	data, err := server.ExportNode(ctx, server.NodeView, server.NodeView, "json")
	if err != nil {
		web.Fail(ctx, "export file err %v", err)
		return
	}
	ctx.Header("content-disposition", `attachment; filename=view.zip`)
	ctx.Data(200, "application/octet-stream", data)
}
