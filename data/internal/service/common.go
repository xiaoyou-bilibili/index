package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"index.data/internal/tools"
	"io"
)

func getObjectSha256(ctx context.Context, id string, contentType string, location int) (string, int64) {
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
		fmt.Println(err)
		return "", 0
	}
	return fmt.Sprintf("%X", ha.Sum(nil)), object.Size
}

func GetReaderSha256(data io.Reader) string {
	ha := sha256.New()
	if _, err := io.Copy(ha, data); err != nil {
		return ""
	}
	return fmt.Sprintf("%X", ha.Sum(nil))
}

func getByteSha256(data []byte) (string, int64) {
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%X", h.Sum(nil)), int64(len(data))
}
