package server

import (
	"context"
	"fmt"
	"testing"
)

var relation1 *RelationService
var data1 *DataService

func TestMain(t *testing.M) {
	//InitConsul()
	relation1 = CreateRelationService()
	data1 = CreateDataService()
	t.Run()
}

func TestGetAttribute(t *testing.T) {
	//var relation = CreateRelationService()
	//relation.AddN
	//ode("people", PeopleInfo{
	//	Name: "小游",
	//	Pic: "123456",
	//	Desc: "描述信息",
	//	Birth: "生日",
	//})
}

func TestGetNode(t *testing.T) {
	//var music MusicInfo
	//err := relation1.GetNode(context.Background(), 7, &music)
	//fmt.Println(err)
	//fmt.Println(music)
}

func TestRelationService_UpdateNode(t *testing.T) {
	err := relation1.UpdateNode(context.Background(), 901, map[string]string{"name": "666"})
	fmt.Println(err)
}
