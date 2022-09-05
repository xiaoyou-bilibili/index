package storage

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"time"
)

// NewS3Storage 初始化一个S3对象
func NewS3Storage(endpoint string, access string, secret string, ssl bool) Storage {
	// 建立连接
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(access, secret, ""),
		Secure: ssl,
	})
	if err != nil {
		panic(any(err))
	}
	return &s3Storage{
		client: client,
	}
}

type s3Storage struct {
	client *minio.Client
}

func (s s3Storage) PutObjet(ctx context.Context, bucket string, req Object) error {
	_, err := s.client.PutObject(ctx, bucket, req.ObjectID, req.Data, req.Size, minio.PutObjectOptions{ContentType: req.ContentType})
	return err
}

func (s s3Storage) GetObjet(ctx context.Context, bucket string, objectID string) (Object, error) {
	object, err := s.client.GetObject(ctx, bucket, objectID, minio.GetObjectOptions{})
	if err != nil {
		return Object{}, err
	}
	state, err := object.Stat()
	if err != nil {
		return Object{}, err
	}
	return Object{
		ObjectID:    objectID,
		ContentType: state.ContentType,
		Data:        object,
		Size:        state.Size,
	}, nil
}

func (s s3Storage) DeleteObject(ctx context.Context, bucket string, objectID string) error {
	return s.client.RemoveObject(ctx, bucket, objectID, minio.RemoveObjectOptions{ForceDelete: true})
}

func (s s3Storage) GetDownloadLink(ctx context.Context, bucket string, objectID string, filename string) (string, error) {
	reqParams := make(url.Values)
	// 设置下载的文件名字
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	// 默认设置链接有效期为1个小时
	link, err := s.client.PresignedGetObject(ctx, bucket, objectID, time.Hour, reqParams)
	if err != nil {
		return "", err
	}
	return link.String(), nil
}

func (s s3Storage) PutLargeObject(ctx context.Context, bucket string, obj Object, filepath string) error {
	_, err := s.client.FPutObject(ctx, bucket, obj.ObjectID, filepath, minio.PutObjectOptions{ContentType: obj.ContentType})
	return err
}
