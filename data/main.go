package main

import (
	"git.xiaoyou.host/index/common/proto/data"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
	"index.data/internal/rpc"
	"index.data/internal/web"
)

func main() {
	// 服务发现
	config := api.DefaultConfig()
	// 这里配置成我们的consul的agent地址
	config.Address = "consul-0.consul.xiaoyou-index.svc.cluster.local:8500"
	consulClient, err := api.NewClient(config)
	if err != nil {
		panic(any(err))
	}
	r := consul.New(consulClient)
	// 路由注册
	router := gin.Default()
	web.RegisterRouter(router)
	// 启动web服务
	httpSrv := http.NewServer(http.Address(":8001"))
	httpSrv.HandlePrefix("/", router)
	// 启动rpc服务
	s := rpc.NewDataService()
	grpcSrv := grpc.NewServer(
		grpc.Address(":9001"),
		grpc.Middleware(recovery.Recovery()))
	data.RegisterDataServer(grpcSrv, s)
	// 启动总服务
	app := kratos.New(
		kratos.Name("index.data"),
		kratos.Server(httpSrv, grpcSrv),
		kratos.Registrar(r))
	if err := app.Run(); err != nil {
		panic(any(err))
	}
}
