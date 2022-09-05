package tools

import (
	"git.xiaoyou.host/index/common/web"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

// GinGetFiledInt 快速获取某一个字段的int类型
func GinGetFiledInt(ctx *gin.Context, filed string) (int, bool) {
	// 获取id参数
	p := ctx.Param(filed)
	id, err := strconv.Atoi(p)
	if id == 0 || err != nil {
		web.Fail(ctx, "参数错误")
		return 0, false
	}
	return id, true
}

// GinGetQueryFiledInt 获取query字段
func GinGetQueryFiledInt(ctx *gin.Context, filed string) int {
	id, _ := strconv.Atoi(ctx.Query(filed))
	return id
}

// GinGetFiledIntList 快速获取某一个字段列表
func GinGetFiledIntList(ctx *gin.Context, filed string) []int64 {
	// 获取id参数
	strs := strings.Split(ctx.Param(filed), ",")
	res := make([]int64, 0, len(strs))
	for _, str := range strs {
		id, _ := strconv.Atoi(str)
		if id > 0 {
			res = append(res, int64(id))
		}
	}

	return res
}

// GinBindData 绑定数据
func GinBindData(ctx *gin.Context, req interface{}) bool {
	if err := ctx.BindJSON(req); err != nil {
		web.Fail(ctx, "参数错误")
		return false
	}
	return true
}
