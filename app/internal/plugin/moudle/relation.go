package moudle

import (
	"context"
	"git.xiaoyou.host/index/common/proto/relation"
	"git.xiaoyou.host/index/common/tool/log"
	"github.com/gin-gonic/gin"
	"index.app/internal/server"
	"math"
)

type RelationServer struct {
	server *server.RelationService
}

func (s RelationServer) AddNode(ctx context.Context, name string, info map[string]string) map[string]interface{} {
	id, err := s.server.AddNode(ctx, name, info)
	return map[string]interface{}{"id": id, "err": err}
}

func (s RelationServer) GinAddNode(ctx *gin.Context, name string, data map[string]string, okHandle func(int64)) {
	if id, ok := s.server.GinAddNode(ctx, name, data); ok {
		okHandle(id)
	}
}

func (s RelationServer) GetNode(ctx context.Context, id int64, data interface{}) map[string]interface{} {
	info, err := s.server.GetNode(ctx, id)
	return map[string]interface{}{"info": info, "err": err}
}

func (s RelationServer) DeleteNode(ctx context.Context, id []int64) error {
	return s.server.DeleteNode(ctx, id)
}

func (s RelationServer) AddRelation(ctx context.Context,
	start int64, end int64, relationInfo string, info map[string]string) map[string]interface{} {
	id, err := s.server.AddRelation(ctx, start, end, relationInfo, info)
	return map[string]interface{}{"id": id, "err": err}
}

func (s RelationServer) GetNodeChild(ctx context.Context, id int64, relationType string, current int32, size int32, data interface{}) map[string]interface{} {
	nodes, total, err := s.server.GetNodeChild(ctx, id, relationType, current, size)
	return map[string]interface{}{"nodes": nodes, "total": total, "err": err}
}

func (s RelationServer) GetNodeParent(ctx context.Context, id int64, relationType string, current int32, size int32, data interface{}) map[string]interface{} {
	nodes, total, err := s.server.GetNodeParent(ctx, id, relationType, current, size)
	return map[string]interface{}{"nodes": nodes, "total": total, "err": err}
}

func (s RelationServer) DeleteRelationWithNode(ctx context.Context, start int64, end int64) error {
	return s.server.DeleteRelationWithNode(ctx, start, end)
}

func (s RelationServer) FindNode(ctx context.Context, label string, field string, keyword string, current int32, size int32, data interface{}) map[string]interface{} {
	nodes, total, err := s.server.FindNode(ctx, label, field, keyword, current, size)
	return map[string]interface{}{"nodes": nodes, "total": total, "err": err}
}

func (s RelationServer) UpdateNode(ctx context.Context, id int64, attribute map[string]string) error {
	return s.server.UpdateNode(ctx, id, attribute)
}

func (s RelationServer) NodeRange(nodes []*relation.Node, data interface{}) []map[string]interface{} {
	res := make([]map[string]interface{}, 0, len(nodes))
	for _, v := range nodes {
		info := make(map[string]interface{})
		info["id"] = v.Id
		for k, v2 := range v.Info.Attribute {
			info[k] = v2
		}
		res = append(res, info)
	}
	return res
}

func (s RelationServer) AddTag(ctx context.Context, name string) map[string]interface{} {
	res := map[string]interface{}{"id": 0, "err": nil}
	var tagID int64 = 0
	nodes, total, err := s.server.FindNode(ctx, "tag", "name", name, 0, 0)
	if err != nil || total == 0 {
		log.CtxLogInfo(ctx, "tag not exist %v", err)
		// 新建标签
		tagID, err = s.server.AddNode(ctx, "tag", map[string]string{"name": name})
		if err != nil {
			log.CtxLogInfo(ctx, "add node error %v", err)
			res["err"] = err
			return res
		}
	} else {
		tagID = nodes[0].Id
	}
	res["id"] = tagID
	return res
}

// RangeAllNode 遍历所有节点
func (s RelationServer) RangeAllNode(ctx context.Context, name string, data interface{}, handle func(nodes []*relation.Node)) {
	// 先获取一下节点总数
	nodes, total, err := s.server.FindNode(ctx, name, "name", ".*", 1, 200)
	if err != nil {
		return
	}
	handle(nodes)
	// 然后从第二页开始遍历
	for i := 2; i <= int(math.Ceil(float64(total)/200)); i++ {
		nodes, _, err = s.server.FindNode(ctx, name, "name", ".*", int32(i), 200)
		if err != nil {
			continue
		}
		handle(nodes)
	}
}
