package web

import "github.com/gin-gonic/gin"

// RegisterRouter 注册http路由
func RegisterRouter(r *gin.Engine) {
	r.GET("/search", Search)
}
