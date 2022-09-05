package service

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"index.data/ent"
	"index.data/ent/object"
	"index.data/internal/tools"
	"log"
	"time"
)

// CreateDatabaseService 数据库服务
func CreateDatabaseService(user string, pass string, host string, port int, database string) *DatabaseService {
	// 与数据库建立连接
	// 建立数据库连接
	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", user, pass, host, port, database))
	if err != nil {
		panic(any(err))
	}
	// 自动创建表
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return &DatabaseService{
		client: client,
	}
}

type Object struct {
	ID string `json:"id"` // 对象ID
	ObjectInfo
}

type ObjectInfo struct {
	Name        string         `json:"name"`         // 对象名字
	ContentType string         `json:"content_type"` // 对象类型
	Sha256      string         `json:"sha_256"`      // 文件sha256的值
	Size        int64          `json:"size"`         // 文件大小
	Location    ObjectLocation `json:"location"`     // 对象存储的位置
	CreateTime  time.Time      `json:"create_time"`  // 对象创建时间
	UpdateTime  time.Time      `json:"update_time"`  // 对象更新时间
}

type ObjectLocation uint8

const (
	LocationObjectStorage ObjectLocation = 1 // 存储在对象存储中
	LocationKvRock        ObjectLocation = 2 // 存储在kvRock中
)

type DatabaseService struct {
	client *ent.Client
}

// AddObject 添加对象
func (s DatabaseService) AddObject(ctx context.Context, object Object) error {
	_, err := s.client.Object.Create().
		SetID(object.ID).
		SetObjectName(object.Name).
		SetContentType(object.ContentType).
		SetObjectLocation(uint8(object.Location)).
		SetObjectSha256(object.Sha256).
		SetObjectSize(object.Size).
		Save(ctx)
	return err
}

// GetObjectInfo 获取对象信息
func (s DatabaseService) GetObjectInfo(ctx context.Context, id string) (Object, error) {
	info, err := s.client.Object.Query().Where(object.ID(id)).First(ctx)
	if err != nil {
		return Object{}, err
	}
	return Object{
		ID: info.ID,
		ObjectInfo: ObjectInfo{
			Name:        info.ObjectName,
			ContentType: info.ContentType,
			Location:    ObjectLocation(info.ObjectLocation),
			Sha256:      info.ObjectSha256,
			Size:        info.ObjectSize,
			CreateTime:  info.CreatedAt,
			UpdateTime:  info.UpdatedAt,
		},
	}, nil
}

// GetObjectInfoBySha256AndSize 根据文件的sha256和大小判断文件是否已经上传过了
func (s DatabaseService) GetObjectInfoBySha256AndSize(ctx context.Context, sha256 string, size int64) (Object, error) {
	info, err := s.client.Object.Query().Where(object.ObjectSha256(sha256), object.ObjectSize(size)).First(ctx)
	if err != nil {
		return Object{}, err
	}
	return Object{
		ID: info.ID,
		ObjectInfo: ObjectInfo{
			Name:        info.ObjectName,
			ContentType: info.ContentType,
			Location:    ObjectLocation(info.ObjectLocation),
			Sha256:      info.ObjectSha256,
			Size:        info.ObjectSize,
			CreateTime:  info.CreatedAt,
			UpdateTime:  info.UpdatedAt,
		},
	}, nil
}

// DeleteObject 删除对象信息
func (s DatabaseService) DeleteObject(ctx context.Context, id string) error {
	return s.client.Object.DeleteOneID(id).Exec(ctx)
}

// UpdateObject 更新对象信息
func (s DatabaseService) UpdateObject(ctx context.Context, id string, info ObjectInfo) error {
	return s.client.Object.UpdateOneID(id).
		SetUpdatedAt(time.Now()).
		SetObjectName(info.Name).
		SetContentType(info.ContentType).
		SetObjectLocation(uint8(info.Location)).
		SetObjectSha256(info.Sha256).
		SetObjectSize(info.Size).
		Exec(ctx)
}

// GetAllObject 获取所有的对象
func (s DatabaseService) GetAllObject(ctx context.Context) ([]*ent.Object, error) {
	return s.client.Object.Query().All(ctx)
}

// UpdateSha256 更新对象的sha56值
func (s DatabaseService) UpdateSha256(ctx context.Context, id string, sha256 string, size int64) error {
	return s.client.Object.UpdateOneID(id).SetObjectSha256(sha256).SetObjectSize(size).Exec(ctx)
}

// GetLocation 根据内容大小来决定存储位置
func (s DatabaseService) GetLocation(contentType string, size int64) ObjectLocation {
	fmt.Printf("file size %.2fMB", float64(size)/float64(1024*1024))
	// 首先获取一下文件类型
	switch contentType {
	case tools.TypeText:
		// 大于10M放对象存储
		if size/(1024*1024) >= 10 {
			return LocationObjectStorage
		} else {
			return LocationKvRock
		}
	default:
		return LocationObjectStorage
	}
}
