package test

import (
	"context"
	"crypto/sha256"
	"fmt"
	"index.data/internal/service"
	"index.data/internal/tools"
	"index.data/pkg/storage"
	"io"
	"testing"
)

var (
	database   = service.CreateDatabaseService("root", "xiaoyou", "192.168.1.40", 31676, "data")
	fileObject = storage.NewS3Storage("192.168.1.40:32561", "D8E4K4KAKX71QO2X79IG", "L9RiS7xMOVjCCcHfP4bPxGPGHd1MpiUSEHuODVYR", false)
	kvRock     = service.CreateKvRocksService("192.168.1.40:32262")
)

func getFileSha256(ctx context.Context, id string, contentType string, location int) (string, int64) {
	if location == 2 {
		object, err := kvRock.GetTextByte(ctx, id)
		if err != nil {
			return "", 0
		}
		h := sha256.New()
		h.Write(object)
		return fmt.Sprintf("%X", h.Sum(nil)), int64(len(object))
	}
	// 先获取对象
	object, err := fileObject.GetObjet(ctx, tools.GetBucketFromContentType(contentType), id)
	if err != nil {
		fmt.Println(err)
		return "", 0
	}
	ha := sha256.New()
	if _, err := io.Copy(ha, object.Data); err != nil {
		fmt.Println("error2")
	}
	return fmt.Sprintf("%X", ha.Sum(nil)), object.Size
}

func TestGetAllFileSha256(t *testing.T) {
	ctx := context.Background()
	// 先获取所有的文件
	object, err := database.GetAllObject(ctx)
	fmt.Println(err)
	for _, v := range object {
		fmt.Println("计算", v.ID)
		sha, size := getFileSha256(ctx, v.ID, v.ContentType, int(v.ObjectLocation))
		fmt.Println(sha, size)
		err = database.UpdateSha256(ctx, v.ID, sha, size)
		fmt.Println(err)
		//break
	}
}
