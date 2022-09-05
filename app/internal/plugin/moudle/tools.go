package moudle

import (
	"bytes"
	"context"
	"encoding/json"
	"git.xiaoyou.host/index/common/service"
	"git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/tool/tools"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ToolsServer struct{}

func (t ToolsServer) LogInfo(format string, data ...interface{}) {
	log.Info(format, data)
}

func (t ToolsServer) LogCtxInfo(ctx context.Context, format string, data ...interface{}) {
	log.CtxLogInfo(ctx, format, data)
}

func (t ToolsServer) LogCtxError(ctx context.Context, format string, data ...interface{}) {
	log.CtxLogError(ctx, format, data)
}

func (t ToolsServer) GetObjectLink(id string) string {
	return tools.GetObjectLink(id)
}

func (t ToolsServer) HttpSendRequest(ctx context.Context, url string, method string, header map[string]string, data interface{}) map[string]interface{} {
	response := map[string]interface{}{"data": nil, "err": nil}
	var b io.Reader
	if data != nil {
		// 解析data数据
		row, _ := json.Marshal(data)
		b = bytes.NewReader(row)
	} else {
		b = nil
	}

	// 初始化request
	req, err := http.NewRequest(method, url, b)
	req.Header.Set("Content-Type", "application/json")
	for k, v := range header {
		req.Header.Set(k, v)
	}
	// 构建并发送请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.CtxLogError(ctx, "send data error %v", err)
		response["err"] = err
		return response
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var resp map[string]interface{}
	// 解析json数据
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.CtxLogError(ctx, "unmarshal data err %v, body %s", err, string(body))
	}
	response["data"] = resp
	return response
}

func (t ToolsServer) GetConsulKV(ctx context.Context, key string) string {
	return service.GetStringValue(ctx, key)
}

func (t ToolsServer) ContextBackground() context.Context {
	return context.Background()
}

func (t ToolsServer) GetUnix() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func (t ToolsServer) ParseUnix(data string) string {
	t1, err := strconv.ParseInt(data, 10, 64)
	if err == nil {
		return tools.Time2String(time.Unix(t1, 0), true)
	}
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func (t ToolsServer) GetFileSuffix(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return ""
	}
	// 获取文件的后缀
	return filename[index+1:]
}

func (t ToolsServer) ReturnError(format string, data ...any) error {
	return tools.ReturnError(format, data)
}
