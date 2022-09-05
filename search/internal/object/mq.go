package object

import (
	"context"
	"encoding/json"
	"git.xiaoyou.host/index/common/service"
	"git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/tool/mq"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"index.search/pkg/es"
)

var esService = es.CreateEsService()

const (
	index = "object"
)

// StartHandleMessage 开始处理消息队列的数据
func StartHandleMessage() {
	server := mq.CreateConsumerService("index")
	err := server.Subscribe("object", func(ctx context.Context, exts ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		// 获取到消息
		for _, ext := range exts {
			var msg service.EsObjectOption
			// 对消息进行解析
			if err := json.Unmarshal(ext.Body, &msg); err == nil {
				log.CtxLogInfo(ctx, "get message %v", msg)
				HandleEsMessage(ctx, msg)
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		panic(any(err))
	}
}

// HandleEsMessage 处理消息
func HandleEsMessage(ctx context.Context, msg service.EsObjectOption) {
	var err error
	switch msg.Option {
	case 1: // 新增操作
		log.CtxLogError(ctx, "add node %v", msg)
		err = esService.InsertData(ctx, index, msg.EsObject)
	case 2: // 删除操作
		err = esService.DeleteRecordByNodeId(ctx, index, msg.NodeID)
	case 3: // 更新操作
		log.CtxLogError(ctx, "update node %v", msg)
		err = esService.UpdateNode(ctx, index, msg.NodeID, msg.Name, msg.Content, msg.Tags)
	case 4: // 全部删除操作
		err = esService.DeleteAllApp(ctx, index, "app", msg.App)
	}
	if err != nil {
		log.CtxLogError(ctx, "option error %v", err)
	}
}
