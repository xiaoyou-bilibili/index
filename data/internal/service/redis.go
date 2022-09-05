package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func CreateRedisService(addr string) *RedisService {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if rdb.Ping(context.Background()).Val() != "PONG" {
		panic(any("redis client error"))
	}
	return &RedisService{client: rdb}
}

type RedisService struct {
	client *redis.Client
}

// SetByte 设置字节
func (r RedisService) SetByte(ctx context.Context, key string, current int32, value []byte) error {
	return r.client.Set(ctx, fmt.Sprintf("%s:%d", key, current), value, time.Hour).Err()
}

// GetByte 获取字节数据
func (r RedisService) GetByte(ctx context.Context, key string, current int32) ([]byte, error) {
	return r.client.Get(ctx, fmt.Sprintf("%s:%d", key, current)).Bytes()
}

// DeleteByte 删除key
func (r RedisService) DeleteByte(ctx context.Context, key string, current int32) error {
	return r.client.Del(ctx, fmt.Sprintf("%s:%d", key, current)).Err()
}
