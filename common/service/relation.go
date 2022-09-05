package service

import (
	"context"
	"errors"
	"git.xiaoyou.host/index/common/proto/relation"
	"git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/web"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"reflect"
)

// CreateRelationService 初始化关系服务
func CreateRelationService() *RelationService {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("discovery:///index.relation"), grpc.WithDiscovery(registry))
	if err != nil {
		panic(any(err))
	}
	return &RelationService{client: relation.NewRelationClient(conn)}
}

type RelationService struct {
	client relation.RelationClient
}

// GetMapFromStruct 把结构体转换为map
func (s RelationService) GetMapFromStruct(info interface{}) map[string]string {
	// 属性信息
	attribute := make(map[string]string)
	typ := reflect.TypeOf(info)
	val := reflect.ValueOf(info)
	// 如果不是结构体直接退出
	if val.Kind() != reflect.Struct {
		return attribute
	}
	// 遍历所有字段
	for i := 0; i < val.NumField(); i++ {
		// 获取标签的值
		tag := typ.Field(i).Tag.Get("json")
		// 获取对应字段的内容
		attribute[tag] = val.Field(i).String()
	}
	return attribute
}

// GetStructFromMap 把map的值映射到结构体中
func (s RelationService) GetStructFromMap(info map[string]string, res interface{}) error {
	if res == nil {
		return nil
	}
	// 因为我们传入的是指针，可以通过Elem来获取到传入的值
	typ := reflect.TypeOf(res).Elem()
	val := reflect.ValueOf(res).Elem()
	// 如果不是结构体直接退出
	if val.Kind() != reflect.Struct {
		return errors.New("必须为结构体")
	}
	// 遍历的所有的字段
	for i := 0; i < val.NumField(); i++ {
		// 获取标签的值
		tag := typ.Field(i).Tag.Get("json")
		// 获取对应字段的内容
		val.Field(i).SetString(info[tag])
	}
	return nil
}

// AddNode 添加一个节点
func (s RelationService) AddNode(ctx context.Context, name string, info interface{}) (int64, error) {
	resp, err := s.client.AddNode(ctx, &relation.NodeInfo{
		NodeLabel: []string{name},
		Attribute: s.GetMapFromStruct(info),
	})
	if err != nil {
		log.CtxLogError(ctx, "add node error %s", err.Error())
		return 0, err
	}
	if len(resp.Id) == 0 {
		return 0, errors.New(resp.Msg)
	}
	return resp.Id[0], nil
}

// GinAddNode 快速添加节点
func (s RelationService) GinAddNode(ctx *gin.Context, name string, data interface{}) {
	// 添加一个节点
	id, err := s.AddNode(ctx, name, data)
	if err != nil {
		web.Fail(ctx, "add node error %v", err)
		return
	}
	// 返回添加结果
	web.Success(ctx, map[string]interface{}{"id": id})
}

// GetNode 获取节点内容
func (s RelationService) GetNode(ctx context.Context, id int64, res interface{}) error {
	node, err := s.client.GetNode(ctx, &relation.NodeMeta{Id: []int64{id}})
	if err != nil {
		return err
	}
	if node == nil || node.Info == nil {
		return nil
	}
	return s.GetStructFromMap(node.Info.Attribute, res)
}

// DeleteNode 删除节点
func (s RelationService) DeleteNode(ctx context.Context, id []int64) error {
	res, err := s.client.DeleteNode(ctx, &relation.NodeMeta{Id: id})
	if err != nil {
		return err
	}
	if !res.Status {
		return errors.New(res.Msg)
	}
	return nil
}

// AddRelation 添加一个关系
func (s RelationService) AddRelation(ctx context.Context, start int64, end int64, relationInfo string, info interface{}) (int64, error) {
	resp, err := s.client.AddRelation(ctx, &relation.RelationInfo{
		Start:     start,
		End:       end,
		Relation:  relationInfo,
		Attribute: s.GetMapFromStruct(info),
	})
	if err != nil {
		log.CtxLogError(ctx, "add relation error %s", err.Error())
		return 0, err
	}
	if len(resp.Id) == 0 {
		return 0, errors.New(resp.Msg)
	}
	return resp.Id[0], nil
}

// GetNodeChild 获取一个节点的所有子节点
func (s RelationService) GetNodeChild(ctx context.Context, id int64, relationType string, parent interface{}, current int32, size int32) ([]*relation.Node, int64, error) {
	resp, err := s.client.GetNodeChild(ctx, &relation.GetNodeChildReq{
		Id:       id,
		Relation: relationType,
		Current:  current,
		Size:     size,
	})
	if err != nil {
		return nil, 0, err
	}
	if resp == nil || resp.Info == nil || resp.Nodes == nil { // 表示无结果
		log.CtxLogError(ctx, "no result")
		return nil, 0, nil
	}
	return resp.Nodes, resp.Total, s.GetStructFromMap(resp.Info.Attribute, parent)
}

// GetNodeParent 获取节点所有的父节点
func (s RelationService) GetNodeParent(ctx context.Context, id int64, relationType string, parent interface{}, current int32, size int32) ([]*relation.Node, int64, error) {
	resp, err := s.client.GetNodeParent(ctx, &relation.GetNodeChildReq{
		Id:       id,
		Relation: relationType,
		Current:  current,
		Size:     size,
	})
	if err != nil {
		return nil, 0, err
	}
	if resp == nil || resp.Info == nil || resp.Nodes == nil { // 表示无结果
		log.CtxLogError(ctx, "no result")
		return nil, 0, nil
	}
	return resp.Nodes, resp.Total, s.GetStructFromMap(resp.Info.Attribute, parent)
}

// DeleteRelationWithNode 根据开始和结束点删除联系
func (s RelationService) DeleteRelationWithNode(ctx context.Context, start int64, end int64) error {
	resp, err := s.client.DeleteRelationWithNode(ctx, &relation.DeleteNodeWithRelationReq{
		Start: start,
		End:   end,
	})
	if err != nil {
		return err
	}
	if !resp.Status {
		return errors.New(resp.Msg)
	}
	return nil
}

// FindNode 查找节点
func (s RelationService) FindNode(ctx context.Context, label string, field string, keyword string, current int32, size int32) ([]*relation.Node, int64, error) {
	resp, err := s.client.FindNode(ctx, &relation.FindNodeReq{
		NodeLabel: label,
		Field:     field,
		Keyword:   keyword,
		Current:   current,
		Size:      size,
	})
	if err != nil {
		return nil, 0, err
	}
	return resp.NodeList, resp.Total, nil
}

// UpdateNode 更新节点
func (s RelationService) UpdateNode(ctx context.Context, id int64, attribute map[string]string) error {
	resp, err := s.client.UpdateNode(ctx, &relation.Node{
		Id: id,
		Info: &relation.NodeInfo{
			Attribute: attribute,
		},
	})
	if err != nil {
		return err
	}
	if resp != nil && !resp.Status {
		return errors.New(resp.Msg)
	}
	return nil
}
