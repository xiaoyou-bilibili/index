package es

import (
	"context"
	"fmt"
	"testing"
)

var service = CreateEsService()

// 参考 https://github.com/elastic/go-elasticsearch#go-elasticsearch

func TestCreateEsService(t *testing.T) {
	//err := service.InsertData(context.Background(), "object", .Object{
	//	Name:    "测试",
	//	Tags:    []string{"标签1", "标签2"},
	//	Content: "这个是测试的内容",
	//	NodeID:  10,
	//})
	//fmt.Println(err)
}

func TestService_DeleteRecordByNodeId(t *testing.T) {
	err := service.DeleteRecordByNodeId(context.Background(), "object", 10)
	fmt.Println(err)
}

func TestService_SearchNode(t *testing.T) {
	res, err := service.SearchNode(context.Background(), "object", "标签")
	fmt.Println(res)
	fmt.Println(err)
}

func TestService_UpdateNode(t *testing.T) {
	err := service.UpdateNode(context.Background(), "object", 10, "名字", "内容", []string{"111", "211", "311"})
	fmt.Println(err)
}
