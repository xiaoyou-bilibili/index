package storage

import (
	"context"
	"io"
)

// Storage 标准存储接口
type Storage interface {
	PutObjet(ctx context.Context, bucket string, obj Object) error                                        // 存储对象
	GetObjet(ctx context.Context, bucket string, objectID string) (Object, error)                         // 获取对象
	DeleteObject(ctx context.Context, bucket string, objectID string) error                               // 删除对象
	GetDownloadLink(ctx context.Context, bucket string, objectID string, filename string) (string, error) // 获取对象的临时下载链接
	PutLargeObject(ctx context.Context, bucket string, obj Object, filepath string) error
}

// Object 对象内容
type Object struct {
	ObjectID    string    `json:"object_id"`    // 对象ID
	ContentType string    `json:"content_type"` // 对象类型
	Data        io.Reader `json:"data"`         // 对象内容
	Size        int64     `json:"size"`         // 文件
}
