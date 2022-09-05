package es

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"git.xiaoyou.host/index/common/tool/log"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"strings"
)

// CreateEsService 创建服务
func CreateEsService() *Service {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://192.168.1.40:30940"},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Error("init es error %v", err)
		panic(any(err))
	}
	return &Service{client: client}
}

type Service struct {
	client *elasticsearch.Client
}

// GetEsInfo 获取Es的信息
func (s Service) GetEsInfo() {
	res, err := s.client.Info()
	if err != nil {
		log.Error("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Error("Error: %s", res.String())
	}
	var r map[string]interface{}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Error("Error parsing the response body: %s", err)
	}
	// Print client and server version numbers.
	log.Info("Client: %s", elasticsearch.Version)
	log.Info("data: %v", r)
}

// 错误处理
func (s Service) processError(res *esapi.Response, err error) error {
	if err != nil {
		return err
	}
	if res.IsError() {
		return errors.New(res.String())
	}
	return nil
}

// body获取
func (s Service) getBody(data interface{}) io.Reader {
	body, err := json.Marshal(data)
	if err != nil {
		log.Error("marshal data error %v", err)
		return bytes.NewReader([]byte{})
	}
	return bytes.NewReader(body)
}

// InsertData ES 插入数据
func (s Service) InsertData(ctx context.Context, index string, data interface{}) error {
	req := esapi.IndexRequest{
		Index:   index,
		Body:    s.getBody(data),
		Refresh: "true",
	}
	return s.processError(req.Do(ctx, s.client))
}

// DeleteRecordByNodeId 根据ID删除数据
func (s Service) DeleteRecordByNodeId(ctx context.Context, index string, nodeID int64) error {
	req := esapi.DeleteByQueryRequest{
		Index: []string{index},
		Body:  bytes.NewReader([]byte(fmt.Sprintf(`{"query":{"match":{"node_id":%d}}}`, nodeID))),
	}
	return s.processError(req.Do(ctx, s.client))
}

// SearchNode 查找节点
func (s Service) SearchNode(ctx context.Context, index string, keyword string) (*HintRes, error) {
	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader([]byte(fmt.Sprintf(`{"query":{"multi_match":{"query":"%s","fields":["name^3","tags^2","content"]}}}`, keyword))),
	}
	r, err := req.Do(ctx, s.client)
	if err = s.processError(r, err); err != nil {
		return nil, err
	}
	// 对节点进行解析
	defer r.Body.Close()
	var res HintRes
	if err = json.NewDecoder(r.Body).Decode(&res); err != nil {
		return nil, err
	}
	// 返回结果
	return &res, nil
}

// UpdateNode 更新节点
func (s Service) UpdateNode(ctx context.Context, index string, id int64, name string, content string, tags []string) error {
	// 数组转换为json
	row, err := json.Marshal(tags)
	if err != nil {
		return err
	}
	script := fmt.Sprintf(
		`{"script":{"source":"ctx._source['name']='%s';ctx._source['content']='%s';ctx._source['tags']=%s"},"query":{"term":{"node_id": %d}}}`,
		name, content, strings.Replace(string(row), "\"", "'", -1), id,
	)
	// 字符串转义
	script = strings.Replace(script, "\n", " ", -1)
	//fmt.Println(script)
	req := esapi.UpdateByQueryRequest{
		Index: []string{index},
		Body:  bytes.NewReader([]byte(script))}
	return s.processError(req.Do(ctx, s.client))
}

// DeleteAllApp 删除所有app
func (s Service) DeleteAllApp(ctx context.Context, index string, key string, value string) error {
	// 数组转换为json
	script := fmt.Sprintf(`{"query":{"match":{"%s":"%s"}}}`, key, value)
	// 字符串转义
	req := esapi.DeleteByQueryRequest{
		Index: []string{index},
		Body:  bytes.NewReader([]byte(script))}
	return s.processError(req.Do(ctx, s.client))
}
