package mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func CreateConsumerService(group string) *ConsumerService {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(primitive.NamesrvAddr{rocketmqAddr}),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName(group),
	)
	if err != nil {
		panic(any(err))
	}
	return &ConsumerService{consumer: &c}
}

type ConsumerService struct {
	consumer *rocketmq.PushConsumer
}

func (s ConsumerService) Subscribe(topic string, handle func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error)) error {
	err := (*s.consumer).Subscribe(topic, consumer.MessageSelector{}, handle)
	if err != nil {
		return err
	}
	// 启动服务
	err = (*s.consumer).Start()
	if err != nil {
		return err
	}
	return nil
}
