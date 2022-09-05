package plugin

import (
	"context"
	"encoding/json"
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
	"index.app/internal/plugin/moudle"
	"index.app/internal/server"
)

// ScriptContainer 脚本容器
type ScriptContainer struct {
	id           int64                                                        // 插件id
	name         string                                                       // 脚本名字
	unique       string                                                       // 脚本标识
	code         string                                                       // 脚本代码
	codeId       string                                                       // 脚本代码资源id
	engine       *gin.Engine                                                  // gin引擎
	runtime      *goja.Runtime                                                // 脚本容器运行时
	searchHandle func(ctx context.Context, id int64) string                   // 搜索卡片
	manageViews  []*ManegeView                                                // 管理界面
	customView   func(ctx context.Context, id int64, view string) interface{} // 获取自定义页面
}

// GetAllManageViews 获取所有的管理界面
func (c *ScriptContainer) GetAllManageViews() []*ManegeView {
	return c.manageViews
}

// GetCode 获取插件代码
func (c *ScriptContainer) GetCode() string {
	return c.code
}

// GetName 获取插件名字
func (c *ScriptContainer) GetName() string {
	return c.name
}

// GetUnique 获取插件标识
func (c *ScriptContainer) GetUnique() string {
	return c.unique
}

// Update 更新插件代码
func (c *ScriptContainer) Update(ctx context.Context, req server.PluginInfo) error {
	// 更新代码
	if err := server.DataServer.UpdateText(ctx, c.codeId, req.Name, req.Code); err != nil {
		return err
	}
	// 如果名字和标识有更新，那么就更新一下
	if c.unique != req.Unique || c.name != req.Name {
		if err := server.RelationServer.UpdateNode(ctx, c.id, map[string]string{"name": req.Name, "unique": req.Unique}); err != nil {
			return err
		}
	}
	c.code = req.Code
	c.name = req.Name
	c.unique = req.Unique
	// 重载插件
	return c.Reload()
}

// Reload 插件重载
func (c *ScriptContainer) Reload() error {
	// 清空管理界面
	c.manageViews = []*ManegeView{}
	_, err := c.start()
	return err
}

// GetSearchCard 获取搜索卡片
func (c *ScriptContainer) GetSearchCard(ctx context.Context, id int64) string {
	if c.searchHandle != nil {
		return c.searchHandle(ctx, id)
	}
	return ""
}

// GetCustomView 获取自定义页面
func (c *ScriptContainer) GetCustomView(ctx context.Context, id int64, view string) interface{} {
	if c.customView != nil {
		return c.customView(ctx, id, view)
	}
	return nil
}

// 挂载函数
func (c *ScriptContainer) mountFun() {
	// 把所有接口都挂载上去
	_ = c.runtime.Set("tools", moudle.CreateToolsServer())
	_ = c.runtime.Set("gin", moudle.CreateGinServer(c.engine, c.unique))
	_ = c.runtime.Set("dataServer", moudle.CreateDataServer())
	_ = c.runtime.Set("relationServer", moudle.CreateRelationServer())
	_ = c.runtime.Set("mqServer", moudle.CreateMqServer())
	_ = c.runtime.Set("view", ViewServer{container: c})
	return
}

// 启动容器
func (c *ScriptContainer) start() (goja.Value, error) {
	c.runtime = goja.New()
	c.mountFun()
	return c.runtime.RunString(c.code)
}

type ViewServer struct {
	container *ScriptContainer
}

// HandleSearch 处理搜索请求
func (s ViewServer) HandleSearch(handle func(ctx context.Context, id int64) string) {
	s.container.searchHandle = handle
}

// HandleView 获取自定义页面
func (s ViewServer) HandleView(handle func(ctx context.Context, id int64, view string) interface{}) {
	s.container.customView = handle
}

// RegisterManage 注册管理页面
func (s ViewServer) RegisterManage(name string, id int64, handle func() map[string]string) {
	s.container.manageViews = append(s.container.manageViews, &ManegeView{
		ID:   id,
		Name: name,
		Data: handle,
	})
}

// GetView 自动获取某个页面(err为true表示显示错误页面)
func (s ViewServer) GetView(ctx context.Context, id int64, data interface{}, exception bool) interface{} {
	if exception {
		return nil
	}
	info, err := server.RelationServer.GetNode(ctx, id)
	if err != nil {
		return nil
	}
	_, view, err := server.DataServer.GetText(ctx, info["view"])
	if err != nil {
		return nil
	}
	// 获取到对象的视图
	var res map[string]interface{}
	if err = json.Unmarshal([]byte(view), &res); err != nil {
		return nil
	}
	// 直接覆盖data
	res["data"] = data
	return res
}
