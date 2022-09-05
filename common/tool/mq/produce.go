package mq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

const rocketmqAddr = "http://rocketmq-0.rocketmq.xiaoyou-index.svc.cluster.local:9876"

func CreateProduceService(group string) *ProduceService {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer(primitive.NamesrvAddr{rocketmqAddr}),
		producer.WithRetry(2),
		producer.WithGroupName("index"),
	)
	if err != nil {
		panic(any(err))
	}
	// 启动服务
	err = p.Start()
	if err != nil {
		panic(any(err))
	}
	return &ProduceService{produce: &p}
}

type ProduceService struct {
	produce *rocketmq.Producer
}

func (s ProduceService) SendSync(ctx context.Context, topic string, data interface{}) error {
	// 解析为json数据
	row, err := json.Marshal(data)
	if err != nil {
		return err
	}
	res, err := (*s.produce).SendSync(context.Background(), &primitive.Message{
		Topic: topic,
		Body:  row,
	})
	if res.Status != primitive.SendOK {
		return errors.New(fmt.Sprintf("send message erorr %v", err))
	}
	return nil
}
