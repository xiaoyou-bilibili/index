package server

import (
	"context"
	"fmt"
	"git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/tool/tools"
	"index.app/internal/util"
	"io/ioutil"
)

const NodeView = "view"

func ExportNode(ctx context.Context, name string, content string, ext string) ([]byte, error) {
	// 获取所有的页面
	views, _, err := RelationServer.FindNode(ctx, name, "name", ".*", 0, 0)
	if err != nil {
		return nil, tools.ReturnError("get node err %v", err)
	}
	var fileList []string
	// 写入文件tmp
	for _, v := range views {
		attribute := v.Info.Attribute
		// 获取原始视图信息
		_, content, err := DataServer.GetText(ctx, attribute[content])
		if err != nil || content == "" {
			return nil, tools.ReturnError("const is nil err %v", err)
		}
		fmt.Println(attribute["name"], attribute[content])
		filename := fmt.Sprintf("%s.%s", attribute["name"], ext)
		fileList = append(fileList, filename)
		err = ioutil.WriteFile(filename, []byte(content), 0775)
		if err != nil {
			log.CtxLogError(ctx, "write file err %v", err)
		}
	}
	// 压缩文件
	err = util.ZipFiles("tmp.zip", fileList)
	if err != nil {
		return nil, tools.ReturnError("zip file err %v", err)
	}
	data, err := ioutil.ReadFile("tmp.zip")
	if err != nil {
		return nil, tools.ReturnError("open file err %v", err)
	}
	return data, nil
}
