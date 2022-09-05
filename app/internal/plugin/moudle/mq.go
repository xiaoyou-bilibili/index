package moudle

import (
	"context"
	"git.xiaoyou.host/index/common/service"
	"git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/tool/mq"
)

type MqServer struct {
	server *mq.ProduceService
	topic  string
}

func (s MqServer) Add(ctx context.Context, id int64, app string, name string, content string, tags []string) {
	err := s.server.SendSync(ctx, s.topic, service.EsObjectOption{
		Option: 1,
		EsObject: service.EsObject{
			NodeID:  id,
			Name:    name,
			App:     app,
			Content: content,
			Tags:    tags,
		},
	})
	if err != nil {
		log.CtxLogError(ctx, "send data error %v", err)
	}
}

func (s MqServer) Delete(ctx context.Context, id int64) {
	err := s.server.SendSync(ctx, s.topic, service.EsObjectOption{
		Option:   2,
		EsObject: service.EsObject{NodeID: id},
	})
	if err != nil {
		log.CtxLogError(ctx, "send data error %v", err)
	}
}

func (s MqServer) Update(ctx context.Context, id int64, app string, name string, content string, tags []string) {
	err := s.server.SendSync(ctx, s.topic, service.EsObjectOption{
		Option: 3,
		EsObject: service.EsObject{
			NodeID:  id,
			Name:    name,
			App:     app,
			Content: content,
			Tags:    tags,
		},
	})
	if err != nil {
		log.CtxLogError(ctx, "send data error %v", err)
	}
}

func (s MqServer) DeleteAll(ctx context.Context, app string) {
	err := s.server.SendSync(ctx, s.topic, service.EsObjectOption{
		Option:   4,
		EsObject: service.EsObject{App: app},
	})
	if err != nil {
		log.CtxLogError(ctx, "send data error %v", err)
	}
}
