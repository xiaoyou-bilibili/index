package web

import "github.com/gin-gonic/gin"

// RegisterRouter 注册http路由
func RegisterRouter(r *gin.Engine) {
	// 获取对象
	r.GET("/data/object/:id", getObject)
	// 获取对象信息
	r.GET("/data/object/:id/info", getObjectInfo)
	// 上传对象
	r.POST("/data/object", addObject)
	// 删除对象
	r.DELETE("/data/object/:id", deleteObject)
	// 获取文本对象
	r.GET("/data/text/:id", getObjectText)
	// 添加文本对象
	r.POST("/data/text", addObjectText)
	// 更新文本对象
	r.PUT("/data/text/:id", updateObjectText)

	// 上传大文件
	r.POST("/data/object/big", uploadBigObject)
	// 根据sha256返回文件内容
	r.GET("/data/object/sha256", FindFileBySha256)
}
