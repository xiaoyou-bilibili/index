package service

import (
	"bytes"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"index.data/internal/tools"
	"index.data/pkg/storage"
	"io"
	"io/ioutil"
)

var (
	database   = CreateDatabaseService("xxx", "xxx", "mysql-0.mysql.xiaoyou-index.svc.cluster.local", 3306, "data")
	fileObject = storage.NewS3Storage("192.168.1.40:32561", "xxx", "xxx", false)
	kvRock     = CreateKvRocksService("kvrocks-0.kvrocks.xiaoyou-index.svc.cluster.local:6666")
)

const host = "https://index.xiaoyou.host/data/object/"

// AddObject 添加对象
func AddObject(ctx context.Context, name string, size int64, data io.Reader, sha256 string) (string, error) {
	objectID := primitive.NewObjectID().Hex()
	contentType := tools.GetContentType(name)
	location := database.GetLocation(contentType, size)
	// 添加对象时先判断文件是否存在，如果存在那么就直接返回已经存在的对象
	info, err2 := database.GetObjectInfoBySha256AndSize(ctx, sha256, size)
	if err2 == nil && info.ID != "" {
		return info.ID, nil
	}
	var err error
	if location == LocationObjectStorage {
		err = fileObject.PutObjet(ctx, tools.GetBucketFromContentType(contentType), storage.Object{
			ObjectID:    objectID,
			ContentType: contentType,
			Data:        data,
			Size:        size,
		})
	} else {
		data, err := ioutil.ReadAll(data)
		if err != nil {
			return "", err
		}
		err = kvRock.SetText(ctx, objectID, data)
	}
	if err != nil {
		return "", err
	}
	// 如果存储成功就更新到数据库
	return objectID, database.AddObject(ctx, Object{
		ID: objectID,
		ObjectInfo: ObjectInfo{
			Name:        name,
			ContentType: contentType,
			Location:    location,
			Size:        size,
			Sha256:      sha256,
		},
	})
}

// PutBigObject 上传大文件对象
func PutBigObject(ctx context.Context, name string, size int64, filename string, sha256 string) (string, error) {
	objectID := primitive.NewObjectID().Hex()
	contentType := tools.GetContentType(name)
	location := LocationObjectStorage
	// 添加对象时先判断文件是否存在，如果存在那么就直接返回已经存在的对象
	info, err := database.GetObjectInfoBySha256AndSize(ctx, sha256, size)
	if err == nil && info.ID != "" {
		return info.ID, nil
	}
	// 上传文件对象
	err = fileObject.PutLargeObject(ctx, tools.GetBucketFromContentType(contentType), storage.Object{
		ObjectID:    objectID,
		ContentType: contentType,
	}, filename)
	if err != nil {
		return "", err
	}
	// 如果存储成功就更新到数据库
	return objectID, database.AddObject(ctx, Object{
		ID: objectID,
		ObjectInfo: ObjectInfo{
			Name:        name,
			ContentType: contentType,
			Location:    location,
			Size:        size,
			Sha256:      sha256,
		},
	})
}

// GetObject 获取对象
func GetObject(ctx context.Context, id string) (string, storage.Object, error) {
	// 首先获取一下object信息
	object, err := database.GetObjectInfo(ctx, id)
	if err != nil {
		return "", storage.Object{}, err
	}
	// 判断存储的位置
	if object.Location == LocationObjectStorage {
		// 直接获取文件
		info, err := fileObject.GetObjet(ctx, tools.GetBucketFromContentType(object.ContentType), id)
		return object.Name, info, err
	} else {
		// 查询redis
		text, err := kvRock.GetTextByte(ctx, id)
		data := bytes.NewReader(text)
		return object.Name, storage.Object{
			ObjectID:    id,
			ContentType: tools.TypeText,
			Data:        data,
			Size:        data.Size(),
		}, err
	}

}

// FindObjectBySha256 根据sha256查找文件
func FindObjectBySha256(ctx context.Context, sha256 string, size int64) (Object, error) {
	// 添加对象时先判断文件是否存在，如果存在那么就直接返回已经存在的对象
	info, err := database.GetObjectInfoBySha256AndSize(ctx, sha256, size)
	if err == nil && info.ID != "" {
		return info, err
	}
	return info, nil
}

// DeleteObject 删除对象
func DeleteObject(ctx context.Context, id string) error {
	// 先获取object信息
	object, err := database.GetObjectInfo(ctx, id)
	if err != nil {
		return err
	}
	// 判断不同的存储位置
	if object.Location == LocationObjectStorage {
		err = fileObject.DeleteObject(ctx, tools.GetBucketFromContentType(object.ContentType), id)
	} else {
		err = kvRock.DeleteText(ctx, id)
	}
	if err != nil {
		return err
	}
	// 删除数据库
	return database.DeleteObject(ctx, id)
}

// GetTextObject 获取文本对象
func GetTextObject(ctx context.Context, id string) (name string, content string, err error) {
	// 目前文本对象直接从kvRocks中去拿
	object, err := database.GetObjectInfo(ctx, id)
	if err != nil {
		return
	}
	name = object.Name
	content, err = kvRock.GetText(ctx, id)
	return
}

// AddTextObject 添加文本对象
func AddTextObject(ctx context.Context, name string, content string) (string, error) {
	// 先存rocks
	objectID := primitive.NewObjectID().Hex()
	err := kvRock.SetText(ctx, objectID, content)
	if err != nil {
		return "", err
	}
	sha256, size := getByteSha256([]byte(content))
	err = database.AddObject(ctx, Object{
		ID: objectID,
		ObjectInfo: ObjectInfo{
			Name:        name,
			ContentType: tools.TypeText,
			Location:    LocationKvRock,
			Sha256:      sha256,
			Size:        size,
		},
	})
	if err != nil {
		return "", err
	}
	return objectID, nil
}

// UpdateTextObject 更新文本对象
func UpdateTextObject(ctx context.Context, id string, name string, content string) error {
	// 先更新rocks
	err := kvRock.SetText(ctx, id, content)
	if err != nil {
		return err
	}
	sha256, size := getByteSha256([]byte(content))
	return database.UpdateObject(ctx, id, ObjectInfo{
		Name:        name,
		ContentType: tools.TypeText,
		Location:    LocationKvRock,
		Sha256:      sha256,
		Size:        size,
	})
}

// GetObjectInfo 获取对象信息
func GetObjectInfo(ctx context.Context, id string) (Object, error) {
	return database.GetObjectInfo(ctx, id)
}

// GetObjectLink 获取文件的下载链接
func GetObjectLink(ctx context.Context, id string) (string, error) {
	// 首先去查询数据库
	info, err := database.GetObjectInfo(ctx, id)
	if err != nil {
		return "", err
	}
	// 判断文件的位置
	if info.Location == LocationKvRock {
		// 文本对象直接返回下载服务的链接
		return host + id, nil
	}
	// 二进制对象则从数据库中查找
	link, err := fileObject.GetDownloadLink(ctx, tools.GetBucketFromContentType(info.ContentType), id, info.Name)
	if err != nil {
		return "", err
	}
	return link, nil
}
