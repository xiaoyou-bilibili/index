package web

import (
	"git.xiaoyou.host/index/common/web"
	"github.com/gin-gonic/gin"
	"index.search/internal/server"
)

func Search(ctx *gin.Context) {
	// 获取关键词
	q := ctx.Query("q")
	res, err := server.SearchNode(ctx, q)
	if err != nil {
		web.Fail(ctx, "get node err %v", err)
		return
	}
	web.Success(ctx, res)
}
