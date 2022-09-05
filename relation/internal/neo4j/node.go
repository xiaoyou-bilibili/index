package neo4j

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/db"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"strings"
)

// NodeInfo 节点信息
type NodeInfo struct {
	ID         int64                  `json:"id"`         // 节点ID
	NodeLabel  []string               `json:"node_label"` // 节点标签
	Properties map[string]interface{} `json:"properties"` // 节点属性
}

// DeleteNode 删除节点
func (s Service) DeleteNode(id int64) error {
	session := s.newSession()
	info, err := session.Run("MATCH (n) WHERE id(n)=$id DELETE n", map[string]interface{}{
		"id": id,
	})
	if err != nil || info == nil {
		return err
	}
	// 使用上面这种方式删除存在存在关系无法删除的问题，这里换另外一种方式把父关系连带删除
	_, err = info.Consume()
	if err != nil {
		info, err = session.Run("MATCH (a)-[r]->(b) WHERE id(b)=$id DELETE r,b", map[string]interface{}{
			"id": id,
		})
		if err != nil || info == nil {
			return err
		}
		_, err = info.Consume()
		return err
	}
	return err
}

// InsertNode 插入节点
func (s Service) InsertNode(info *NodeInfo) (int64, error) {
	session := s.newSession()
	// 构建标签
	cypher := fmt.Sprintf("CREATE (n:%s %s) RETURN n", strings.Join(info.NodeLabel, ":"), s.buildAttribute(info.Properties))
	res, err := session.Run(cypher, info.Properties)
	if err == nil {
		var sign *db.Record
		sign, err = res.Single()
		if err == nil {
			if i, ok := sign.Get("n"); ok {
				if node, ok := i.(dbtype.Node); ok {
					return node.Id, nil

				}
			}
		}
	}
	return 0, err
}

// UpdateNode 更新节点
func (s Service) UpdateNode(id int64, info *NodeInfo) error {
	session := s.newSession()
	// 构建标签
	cypher := fmt.Sprintf("MATCH (n) WHERE id(n)=%d SET %s RETURN n", id, s.buildUpdate("n", info.Properties))
	_, err := session.Run(cypher, info.Properties)
	return err
}

// GetNode 获取节点
func (s Service) GetNode(id int64) (*NodeInfo, error) {
	session := s.newSession()
	// 查询记录
	res, err := session.Run("MATCH (n) WHERE id(n)=$id RETURN n", map[string]interface{}{"id": id})
	if err == nil {
		var sign *db.Record
		sign, err = res.Single()
		if err == nil {
			if i, ok := sign.Get("n"); ok {
				if node, ok := i.(dbtype.Node); ok {
					return &NodeInfo{
						NodeLabel:  node.Labels,
						Properties: node.Props,
					}, nil

				}
			}
		}
	}
	return &NodeInfo{}, err
}

func (s Service) findNodes(cypher string, totalCypher string) (*NodeInfo, []*NodeInfo, int64, error) {
	session := s.newSession()
	// 构建标签
	res, err := session.Run(cypher, map[string]interface{}{})
	if err != nil || res == nil {
		return nil, nil, 0, err
	}
	record := &db.Record{}
	var node *NodeInfo
	var childs []*NodeInfo
	// 遍历结果
	for res.NextRecord(&record) {
		// node节点只解析一次
		if node == nil {
			if i, ok := record.Get("a"); ok {
				if n, ok := i.(dbtype.Node); ok {
					node = &NodeInfo{NodeLabel: n.Labels, Properties: n.Props}
				}
			}
		}
		if i, ok := record.Get("b"); ok {
			if n, ok := i.(dbtype.Node); ok {
				childs = append(childs, &NodeInfo{ID: n.Id, NodeLabel: n.Labels, Properties: n.Props})
			}
		}
	}

	return node, childs, s.GetNodeCount(totalCypher), nil
}

// GetNodeChild 获取某个节点下所有子节点
func (s Service) GetNodeChild(id int64, relation string, current int32, limit int32) (*NodeInfo, []*NodeInfo, int64, error) {
	return s.findNodes(
		fmt.Sprintf("MATCH (a)-[r:%s]->(b) WHERE id(a)=%d RETURN a,b skip %d limit %d", relation, id, limit*(current-1), limit),
		fmt.Sprintf("MATCH (a)-[r:%s]->(b) WHERE id(a)=%d RETURN count(b)", relation, id))
}

// GetNodeParent 获取某个节点所有父节点
func (s Service) GetNodeParent(id int64, relation string, current int32, limit int32) (*NodeInfo, []*NodeInfo, int64, error) {
	return s.findNodes(
		fmt.Sprintf("MATCH (b)-[r:%s]->(a) WHERE id(a)=%d RETURN a,b skip %d limit %d", relation, id, limit*(current-1), limit),
		fmt.Sprintf("MATCH (b)-[r:%s]->(a) WHERE id(a)=%d RETURN count(b)", relation, id))
}

func (s Service) FindNode(label string, filed string, keyword string, current int32, limit int32) ([]*NodeInfo, int64, error) {
	session := s.newSession()
	// 构建标签
	cypher := fmt.Sprintf("MATCH (n:%s) WHERE n.%s=~\"%s\" RETURN n skip %d limit %d", label, filed, keyword, limit*(current-1), limit)
	res, err := session.Run(cypher, map[string]interface{}{})
	if err != nil || res == nil {
		return nil, 0, err
	}
	record := &db.Record{}
	var nodes []*NodeInfo
	// 遍历结果
	for res.NextRecord(&record) {
		if i, ok := record.Get("n"); ok {
			if n, ok := i.(dbtype.Node); ok {
				nodes = append(nodes, &NodeInfo{ID: n.Id, NodeLabel: n.Labels, Properties: n.Props})
			}
		}
	}

	return nodes, s.GetNodeCount(fmt.Sprintf("MATCH (n:%s) WHERE n.%s=~\"%s\" RETURN count(n)", label, filed, keyword)), nil
}
