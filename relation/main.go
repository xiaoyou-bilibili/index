package main

import (
	"git.xiaoyou.host/index/common/proto/relation"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"index.relation/internal/rpc"
	"log"
)

func main() {
	// 服务发现
	config := api.DefaultConfig()
	// 这里配置成我们的consul的agent地址
	config.Address = "consul-1.consul.xiaoyou-index.svc.cluster.local:8500"
	consulClient, err := api.NewClient(config)
	if err != nil {
		panic(any(err))
	}
	r := consul.New(consulClient)
	// RPC服务注册
	s := rpc.NewRelationService()
	grpcSrv := grpc.NewServer(grpc.Address(":9000"), grpc.Middleware(recovery.Recovery()))
	relation.RegisterRelationServer(grpcSrv, s)
	// 启动应用
	app := kratos.New(
		kratos.Name("index.relation"),
		kratos.Server(grpcSrv),
		kratos.Registrar(r),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
