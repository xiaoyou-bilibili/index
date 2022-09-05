package moudle

import (
	"fmt"
	"git.xiaoyou.host/index/common/common"
	"git.xiaoyou.host/index/common/tool/tools"
	"git.xiaoyou.host/index/common/web"
	"github.com/gin-gonic/gin"
)

// 全局路由路径信息，用于存储对应路由的处理函数
var globalRouterMap = map[string]*func(c *gin.Context){}

// 路由代理,用于全局处理
func proxy(c *gin.Context) {
	// 获取路径
	path := c.FullPath()
	// 获取请求方式
	method := c.Request.Method
	// 全局路径
	path = method + path
	if fun, ok := globalRouterMap[path]; ok && fun != nil {
		// 直接执行对应的函数
		(*fun)(c)
	} else {
		web.Fail(c, "路由不存在")
	}
}

type GinServer struct {
	engine *gin.Engine // ginEngine信息
	name   string      // group名称
	handle map[string]func(c *gin.Context)
}

// Handle 用于代理函数调用
func (r GinServer) Handle(method string, url string, function func(c *gin.Context)) {
	// 路由错误处理
	fun := func(c *gin.Context) {
		// 异常处理
		defer func() {
			if err := recover(); err != nil {
				web.Fail(c, "插件执行错误!错误信息: %v", err)
			}
		}()
		// 调用js的回调函数
		function(c)
	}
	url = fmt.Sprintf("/app/%s%s", r.name, url)
	path := method + url
	// 判断路由是否注册过，如果没注册就手动注册
	if _, ok := globalRouterMap[path]; !ok {
		switch method {
		case "GET":
			r.engine.GET(url, proxy)
		case "POST":
			r.engine.POST(url, proxy)
		case "PUT":
			r.engine.PUT(url, proxy)
		case "DELETE":
			r.engine.DELETE(url, proxy)
		}
	}
	// 自动更新函数
	globalRouterMap[path] = &fun
}

func (r GinServer) Success(ctx *gin.Context, data interface{}) { web.Success(ctx, data) }

func (r GinServer) Fail(ctx *gin.Context, format string, a ...any) { web.Fail(ctx, format, a) }

func (r GinServer) GetPathInt(ctx *gin.Context, field string, okCallback func(data int)) {
	if id, ok := tools.GinGetFiledInt(ctx, field); ok {
		okCallback(id)
	}
}

func (r GinServer) GetQueryInt(ctx *gin.Context, field string) int {
	return tools.GinGetQueryFiledInt(ctx, field)
}

func (r GinServer) GetPathIntList(ctx *gin.Context, field string) []int64 {
	return tools.GinGetFiledIntList(ctx, field)
}

func (r GinServer) BindData(ctx *gin.Context, req interface{}, callback func(data interface{})) {
	if tools.GinBindData(ctx, &req) {
		callback(req)
	}
}

func (r GinServer) GetFindField(ctx *gin.Context) map[string]interface{} {
	// 获取关键词和字段
	keyword := fmt.Sprintf(".*%s.*", ctx.Query("search_keyword"))
	field := ctx.Query("search_type")
	// 获取页数
	current := tools.GinGetQueryFiledInt(ctx, "current")
	size := tools.GinGetQueryFiledInt(ctx, "size")
	return map[string]interface{}{"keyword": keyword, "field": field, "current": current, "size": size}
}

func (r GinServer) GetPageField(ctx *gin.Context) map[string]interface{} {
	// 获取页数
	current := tools.GinGetQueryFiledInt(ctx, "current")
	size := tools.GinGetQueryFiledInt(ctx, "size")
	return map[string]interface{}{"current": current, "size": size}
}

// ReturnPageInfo 返回分页信息
func (r GinServer) ReturnPageInfo(ctx *gin.Context, current int, total int64, list interface{}) {
	web.Success(ctx, common.PageInfo{
		Current: current,
		Total:   total,
		List:    list,
	})
}
