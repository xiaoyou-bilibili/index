package neo4j

import (
	"fmt"
	"testing"
)

var service = CreateService()

func TestDeleteNode(t *testing.T) {
	err := service.DeleteNode(894)
	fmt.Println(err)
}

func TestDeleteRelation(t *testing.T) {
	err := service.DeleteRelation(0)
	fmt.Println(err)
}

func TestInsert(t *testing.T) {
	_, err := service.InsertNode(&NodeInfo{
		NodeLabel: []string{"People"},
		Properties: map[string]interface{}{
			"name":   "周杰伦",
			"avatar": "xxx",
		}})
	fmt.Println(err)
}

func TestRelation(t *testing.T) {
	_, err := service.AddRelation(&RelationInfo{
		8, 7, "创作", map[string]interface{}{"time": "2021"},
	})
	fmt.Println(err)
}

func TestGetNodeInfo(t *testing.T) {
	res, err := service.GetNode(7)
	fmt.Println(err)
	for k, v := range res.Properties {
		fmt.Println(k, v)
	}
}

func TestGetRelation(t *testing.T) {
	res, err := service.GetRelation(0)
	fmt.Println(err)
	for k, v := range res.Properties {
		fmt.Println(k, v)
	}
}

func TestUpdateNode(t *testing.T) {
	err := service.UpdateNode(0, &NodeInfo{Properties: map[string]interface{}{"name": "", "singer": ""}})
	fmt.Println(err)
}

func TestUpdateRelation(t *testing.T) {
	err := service.UpdateRelation(0, &RelationInfo{Properties: map[string]interface{}{"aaa": "123456"}})
	fmt.Println(err)
}

func TestService_DeleteRelationWithNode(t *testing.T) {
	err := service.DeleteRelationWithNode(8, 1)
	fmt.Println(err)
}

func TestService_GetNodeChild(t *testing.T) {
	node, child, count, err := service.GetNodeChild(901, "file", 1, 2)
	fmt.Println(node)
	fmt.Println(child)
	fmt.Println(count)
	fmt.Println(err)
}

func TestService_GetNodeParent(t *testing.T) {
	node, child, count, err := service.GetNodeParent(1001, "tag", 1, 2)
	fmt.Println(node)
	fmt.Println(child)
	fmt.Println(count)
	fmt.Println(err)
}

func TestService_FindNode(t *testing.T) {
	node, total, err := service.FindNode("music", "name", ".*世界.*", 3, 2)
	fmt.Println(node)
	fmt.Println(total)
	fmt.Println(err)
}

func TestService_GetNodeCount(t *testing.T) {
	count := service.GetNodeCount("MATCH (n:music) RETURN count(n)")
	fmt.Println(count)
}
