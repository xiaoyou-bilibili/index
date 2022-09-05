package service

import (
	"context"
	"git.xiaoyou.host/index/common/tool/log"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
)

var registry *consul.Registry
var client *api.Client

// InitConsul 初始化consul服务
func InitConsul() {
	config := api.DefaultConfig()
	config.Address = "consul-1.consul.xiaoyou-index.svc.cluster.local:8500"
	consulClient, err := api.NewClient(config)
	if err != nil {
		panic(any(err))
	}
	client = consulClient
	// 注册consul服务
	registry = consul.New(consulClient)
}

func GetRegistry() *consul.Registry {
	return registry
}

func GetClient() *api.Client {
	return client
}

func GetStringValue(ctx context.Context, key string) string {
	if client == nil {
		return ""
	}
	conf, _, err := client.KV().Get(key, nil)
	if err != nil || conf == nil {
		log.CtxLogError(ctx, "get config error %v", err)
		return ""
	}
	return string(conf.Value)
}
