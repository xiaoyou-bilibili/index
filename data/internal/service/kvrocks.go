package service

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// CreateKvRocksService 文本存储服务
func CreateKvRocksService(addr string) *KvRocksService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if rdb.Ping(context.Background()).Val() != "PONG" {
		panic(any("redis client error"))
	}
	return &KvRocksService{client: rdb}
}

type KvRocksService struct {
	client *redis.Client
}

// SetText 设置文本信息
func (s KvRocksService) SetText(ctx context.Context, key string, value interface{}) error {
	// 需要对文本进行base4编码
	return s.client.Set(ctx, key, value, 0).Err()
}

// GetTextByte 获取文本信息
func (s KvRocksService) GetTextByte(ctx context.Context, key string) ([]byte, error) {
	return s.client.Get(ctx, key).Bytes()
}

// GetText 获取文本信息
func (s KvRocksService) GetText(ctx context.Context, key string) (string, error) {
	cmd := s.client.Get(ctx, key)
	return cmd.Val(), cmd.Err()
}

// DeleteText 删除文本
func (s KvRocksService) DeleteText(ctx context.Context, key string) error {
	return s.client.Del(ctx, key).Err()
}
