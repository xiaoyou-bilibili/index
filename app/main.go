package main

import (
	app2 "git.xiaoyou.host/index/common/proto/app"
	"git.xiaoyou.host/index/common/service"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"index.app/internal/plugin"
	"index.app/internal/rpc"
	"index.app/internal/server"
	"index.app/internal/web"
)

func main() {
	server.InitRpc()
	// http
	router := gin.Default()
	web.RegisterRouter(router)
	httpSrv := http.NewServer(http.Address(":8006"))
	httpSrv.HandlePrefix("/", router)
	// 加载插件
	plugin.InitPlugin(router)
	// rpc
	s := rpc.NewAppService()
	grpcSrv := grpc.NewServer(grpc.Address(":9006"), grpc.Middleware(recovery.Recovery()))
	app2.RegisterAppServer(grpcSrv, s)
	app := kratos.New(
		kratos.Name("index.app"),
		kratos.Server(httpSrv, grpcSrv),
		kratos.Registrar(service.GetRegistry()))
	if err := app.Run(); err != nil {
		panic(any(err))
	}
}
