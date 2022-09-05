package server

import (
	"git.xiaoyou.host/index/common/service"
	"git.xiaoyou.host/index/common/tool/mq"
)

var (
	DataServer     *DataService
	RelationServer *RelationService
	MqServer       *mq.ProduceService
)

func InitRpc() {
	service.InitConsul()
	// 初始化服务发现
	DataServer = CreateDataService()
	RelationServer = CreateRelationService()
	MqServer = mq.CreateProduceService("index")
}
