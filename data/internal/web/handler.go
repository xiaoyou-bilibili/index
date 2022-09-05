package web

import (
	"bytes"
	"encoding/base64"
	"fmt"
	log2 "git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/web"
	"github.com/gin-gonic/gin"
	"index.data/internal/service"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
)

var redisClient = service.CreateRedisService("redis-0.redis.xiaoyou-index.svc.cluster.local:6379")

func getObject(ctx *gin.Context) {
	id := ctx.Param("id")
	_, data, err := service.GetObject(ctx, id)
	if err != nil {
		fail(ctx, "没有这个文件 %v", err)
		return
	}
	ctx.DataFromReader(200, data.Size, data.ContentType, data.Data, map[string]string{})
}

func getObjectInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	info, err := service.GetObjectInfo(ctx, id)
	if err != nil {
		fail(ctx, "获取信息失败 %v", err)
		return
	}
	success(ctx, info)
}

func addObject(ctx *gin.Context) {
	// 获取上传的文件
	file, _ := ctx.FormFile("file")
	if file == nil {
		fail(ctx, "请上传文件")
		return
	}
	f, err := file.Open()
	if err != nil {
		fail(ctx, "打开文件失败 %v", err)
		return
	}
	defer f.Close()
	tmp, err := ioutil.ReadAll(f)
	if err != nil {
		fail(ctx, "读取字节失败 %v", err)
		return
	}
	id, err := service.AddObject(ctx, file.Filename, file.Size, bytes.NewReader(tmp), service.GetReaderSha256(bytes.NewReader(tmp)))
	if err != nil {
		fmt.Println(err)
		fail(ctx, "保存文件失败 %v", err)
		return
	}
	success(ctx, addObjectResp{ObjectID: id})
}

func deleteObject(ctx *gin.Context) {
	id := ctx.Param("id")
	err := service.DeleteObject(ctx, id)
	if err != nil {
		fail(ctx, "删除失败 %v", err)
		return
	}
	success(ctx, nil)
}

func getObjectText(ctx *gin.Context) {
	id := ctx.Param("id")
	name, content, err := service.GetTextObject(ctx, id)
	if err != nil {
		fail(ctx, "获取文本失败 %v", err)
		return
	}
	success(ctx, getTextResp{
		Name:    name,
		Content: content,
	})
}

func addObjectText(ctx *gin.Context) {
	req := getTextResp{}
	err := ctx.BindJSON(&req)
	if err != nil {
		fail(ctx, "参数错误 %v", err)
		return
	}
	id, err := service.AddTextObject(ctx, req.Name, req.Content)
	if err != nil {
		fail(ctx, "保存失败 %v", err)
		return
	}
	success(ctx, addObjectResp{ObjectID: id})
}

func updateObjectText(ctx *gin.Context) {
	id := ctx.Param("id")
	req := getTextResp{}
	err := ctx.BindJSON(&req)
	if err != nil {
		fail(ctx, "获取参数失败 %v", err)
		return
	}
	// 更新文本对象
	err = service.UpdateTextObject(ctx, id, req.Name, req.Content)
	if err != nil {
		fail(ctx, "更新失败 %v", err)
		return
	}
	success(ctx, nil)
}

// 使用并发安全map
var sizeCache sync.Map

// 上传大对象
func uploadBigObject(ctx *gin.Context) {
	var req FileUploadReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		web.Fail(ctx, "bind data error %v", err)
		return
	}
	var data []byte
	if req.Data != "" {
		// base64解码
		data, err = base64.StdEncoding.DecodeString(req.Data)
	} else {
		// 从form表单中读取数据
		file, err := ctx.FormFile("file")
		if err != nil {
			web.Fail(ctx, "没有文件")
			return
		}
		f, err := file.Open()
		if err != nil {
			web.Fail(ctx, "打开文件失败")
			return
		}
		defer f.Close()
		data, err = ioutil.ReadAll(f)
	}
	if err != nil {
		web.Fail(ctx, "获取二进制信息失败")
		return
	}

	// 把二进制信息直接保存到redis中
	err = redisClient.SetByte(ctx, req.Name, req.Current, data)
	if err != nil {
		web.Fail(ctx, "set redis error")
		return
	}

	// 上传文件
	uploadFile := func() {
		// 删除大小缓存
		sizeCache.Delete(req.Name)
		// 读取redis数据
		_ = os.Remove("tmp")
		file, err := os.OpenFile("tmp", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			web.Fail(ctx, "打开文件失败")
			return
		}
		// 遍历读取文件
		var i int32
		for i = 1; i <= req.TotalChunk; i++ {
			log2.CtxLogInfo(ctx, "正在写入第 %d,总计 %d\n", i, req.TotalChunk)
			tmp, err := redisClient.GetByte(ctx, req.Name, i)
			if err != nil {
				web.Fail(ctx, "获取文件失败")
				return
			}
			// 及时删除，避免占用太多内容
			_ = redisClient.DeleteByte(ctx, req.Name, i)
			_, _ = file.Write(tmp)
		}
		// 上传文件并计算sha256
		_ = file.Close()
		f, err := os.Open("tmp")
		if err != nil {
			web.Fail(ctx, "打开文件失败 %v", err)
			return
		}
		sha256 := service.GetReaderSha256(f)
		log2.CtxLogInfo(ctx, "sha256 %s", sha256)
		_ = f.Close()
		id, err := service.PutBigObject(ctx, req.Name, req.TotalSize, "tmp", sha256)
		if err != nil {
			web.Fail(ctx, "添加对象失败 %v", err)
			return
		}
		log2.CtxLogInfo(ctx, "id %s \n", id)
		web.Success(ctx, map[string]string{"id": id})
		return
	}

	if size, ok := sizeCache.Load(req.Name); ok {
		newSize := size.(int64) + int64(len(data))
		sizeCache.Store(req.Name, newSize)
		// 如果上传完成就可以直接上传到对象存储了
		if req.TotalSize == newSize {
			uploadFile()
			return
		}
	} else {
		// 如果一个分片就可以解决那么就直接调用快速删除
		if int64(len(data)) >= req.TotalSize {
			id, err := service.AddObject(ctx, req.Name, req.TotalSize, bytes.NewReader(data), service.GetReaderSha256(bytes.NewReader(data)))
			if err != nil {
				fmt.Println(err)
				fail(ctx, "保存文件失败 %v", err)
				return
			}
			success(ctx, addObjectResp{ObjectID: id})
			return
		}
		sizeCache.Store(req.Name, int64(len(data)))
	}
	web.Success(ctx, nil)
}

// FindFileBySha256 根据SHa256查找文件
func FindFileBySha256(ctx *gin.Context) {
	total := ctx.Query("size")
	sha256 := ctx.Query("sha256")
	size, err := strconv.ParseInt(total, 10, 64)
	if err != nil || sha256 == "" || size == 0 {
		web.Fail(ctx, "参数错误")
		return
	}
	info, err := service.FindObjectBySha256(ctx, sha256, size)
	if err != nil {
		web.Fail(ctx, "查找对象失败 %v", err)
		return
	}
	web.Success(ctx, info)
}
