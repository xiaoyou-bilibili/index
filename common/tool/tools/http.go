package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"git.xiaoyou.host/index/common/tool/log"
	"io"
	"io/ioutil"
	"net/http"
)

// HttpSendRequest 发送HTTP请求
func HttpSendRequest(ctx context.Context, url string, method string, header map[string]string, data interface{}, resp interface{}) {
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
		return
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// 解析json数据
	err = json.Unmarshal(body, resp)
	if err != nil {
		log.CtxLogError(ctx, "unmarshal data err %v, body %s", err, string(body))
	}
}
