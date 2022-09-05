package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
	"index.search/internal/object"
	"index.search/internal/server"
	"index.search/internal/web"
	//"index.file/internal/web"
)

func main() {
	server.InitRpc()
	object.StartHandleMessage()
	// 注册rpc路由
	router := gin.Default()
	// 注册静态文件路由
	web.RegisterRouter(router)
	httpSrv := http.NewServer(http.Address(":8005"))
	httpSrv.HandlePrefix("/", router)
	app := kratos.New(kratos.Name("index.search"), kratos.Server(httpSrv))
	if err := app.Run(); err != nil {
		panic(any(err))
	}
}
