package plugin

import (
	"context"
	"fmt"
	"git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/tool/tools"
	"github.com/gin-gonic/gin"
	"index.app/internal/server"
)

var (
	pluginList = map[string]*ScriptContainer{}
	engine     *gin.Engine
)

// GetPluginList 获取插件列表
func GetPluginList() []*server.Plugin {
	res := make([]*server.Plugin, 0, len(pluginList))
	for _, v := range pluginList {
		res = append(res, &server.Plugin{
			Id: v.id,
			PluginInfo: server.PluginInfo{
				Name:   v.name,
				Unique: v.unique,
			},
		})
	}
	return res
}

// GetPluginInfo 获取插件的信息
func GetPluginInfo(name string) *ScriptContainer {
	return pluginList[name]
}

// InitPlugin 插件初始化
func InitPlugin(e *gin.Engine) {
	engine = e
	// 获取所有的插件
	ctx := context.Background()
	plugins, _, err := server.RelationServer.FindNode(ctx, server.NodePlugin, "name", ".*", 0, 0)
	if err != nil {
		log.Error("get plug err %v", err)
		return
	}
	// 遍历所有插件
	for _, plugin := range plugins {
		attribute := plugin.Info.Attribute
		fmt.Println("获取插件", attribute["name"], attribute["code"])
		_, content, err := server.DataServer.GetText(ctx, attribute["code"])
		if err != nil {
			log.Error("get content err %v", err)
			continue
		}
		// 创建一个插件
		if err = AddNewPlugin(plugin.Id, attribute["name"], attribute["unique"], content, attribute["code"]); err != nil {
			log.Error("add plugin err %v", err)
		}
	}
}

// AddNewPlugin 添加一个新插件
func AddNewPlugin(id int64, name, unique, code, codeId string) error {
	container := &ScriptContainer{id: id, name: name, unique: unique, code: code, engine: engine, codeId: codeId}
	_, err := container.start()
	if err != nil {
		return tools.ReturnError("start container error %v", err)
	}
	pluginList[container.unique] = container
	return nil
}

// DeletePlugin 删除某个插件
func DeletePlugin(ctx context.Context, unique string) error {
	if container, ok := pluginList[unique]; ok && container != nil {
		// 删除插件内容
		if err := server.DataServer.DeleteObject(ctx, []string{container.codeId}); err != nil {
			return err
		}
		// 删除节点
		if err := server.RelationServer.DeleteNode(ctx, []int64{container.id}); err != nil {
			return err
		}
		// 删除map
		delete(pluginList, unique)
	}
	return nil
}
