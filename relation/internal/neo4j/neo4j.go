package neo4j

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"strings"
)

func CreateService() *Service {
	driver, err := neo4j.NewDriver("neo4j://192.168.1.40:32017", neo4j.BasicAuth("neo4j", "xiaoyou", ""))
	if err != nil {
		panic(any("client neo4j error"))
	}
	return &Service{
		driver:  driver,
		session: driver.NewSession(neo4j.SessionConfig{}),
	}
}

// Service 图数据库服务
type Service struct {
	driver  neo4j.Driver
	session neo4j.Session
}

// 初始化一个session服务
func (s Service) newSession() neo4j.Session {
	return s.session
}

// 构建属性关系
func (s Service) buildAttribute(attribute map[string]interface{}) string {
	tags := make([]string, 0, len(attribute))
	for k := range attribute {
		tags = append(tags, fmt.Sprintf("%s:$%s", k, k))
	}
	return fmt.Sprintf("{%s}", strings.Join(tags, ","))
}

// 构建更新关系
func (s Service) buildUpdate(name string, attribute map[string]interface{}) string {
	tags := make([]string, 0, len(attribute))
	for k := range attribute {
		tags = append(tags, fmt.Sprintf("%s.%s=$%s", name, k, k))
	}
	return strings.Join(tags, ",")
}
