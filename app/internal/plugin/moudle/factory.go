package moudle

import (
	"github.com/gin-gonic/gin"
	"index.app/internal/server"
)

func CreateToolsServer() *ToolsServer {
	return &ToolsServer{}
}

func CreateGinServer(engine *gin.Engine, name string) *GinServer {
	return &GinServer{engine: engine, name: name}
}

func CreateDataServer() *DataServer {
	return &DataServer{server: server.DataServer}
}

func CreateRelationServer() *RelationServer {
	return &RelationServer{server: server.RelationServer}
}

func CreateMqServer() *MqServer {
	return &MqServer{server: server.MqServer, topic: "object"}
}
