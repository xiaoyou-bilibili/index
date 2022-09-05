package server

import (
	"context"
	"encoding/json"
	"git.xiaoyou.host/index/common/proto/app"
	"git.xiaoyou.host/index/common/service"
	"git.xiaoyou.host/index/common/tool/tools"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"index.search/internal/object"
)

var appService app.AppClient

func InitRpc() {
	ctx := context.Background()
	// 初始化服务发现
	service.InitConsul()
	conn, err := grpc.DialInsecure(ctx, grpc.WithEndpoint("discovery:///index.app"), grpc.WithDiscovery(service.GetRegistry()))
	if err != nil {
		panic(any(err))
	}
	appService = app.NewAppClient(conn)
}

type PageBody struct {
	Type      string        `json:"type"`
	Direction string        `json:"direction"`
	ClassName string        `json:"className"`
	Items     []interface{} `json:"items"`
}

type PageInfo struct {
	Type string   `json:"type"`
	Body PageBody `json:"body"`
}

func SearchNode(ctx context.Context, keyword string) (map[string]interface{}, error) {
	// 调用ES获取结果
	res, err := object.SearchKeyword(ctx, keyword)
	if err != nil {
		return nil, tools.ReturnError("search data error %v", err)
	}
	// 调用下游服务获取卡片信息
	infos := make([]*app.CardInfo, 0, len(res))
	for _, node := range res {
		infos = append(infos, &app.CardInfo{NodeId: node.NodeID, App: node.App})
	}
	nodes, err := appService.GetNodeInfo(ctx, &app.GetNodeInfoReq{Infos: infos})
	if err != nil {
		return nil, tools.ReturnError("get node info err %v", err)
	}
	// 遍历获取所有卡片
	cards := make([]interface{}, 0, len(nodes.Cards))
	for _, info := range nodes.Cards {
		var card interface{}
		if err := json.Unmarshal([]byte(info.Info), &card); err == nil {
			cards = append(cards, card)
		}
	}
	// 输出最后结果
	return map[string]interface{}{"view": PageInfo{Type: "page", Body: PageBody{Type: "flex", Direction: "column", ClassName: "search-box", Items: cards}}}, nil
}
