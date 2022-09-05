package server

import (
	"context"
	"errors"
	"git.xiaoyou.host/index/common/proto/data"
	"git.xiaoyou.host/index/common/service"
	"git.xiaoyou.host/index/common/tool/log"
	"git.xiaoyou.host/index/common/tool/tools"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

// CreateDataService 创建数据服务
func CreateDataService() *DataService {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithTimeout(time.Minute*1), // 设置超时时间为1分钟
		grpc.WithEndpoint("discovery:///index.data"),
		grpc.WithDiscovery(service.GetRegistry()))
	if err != nil {
		panic(any(err))
	}
	return &DataService{client: data.NewDataClient(conn)}
}

// DataService 数据服务
type DataService struct {
	client data.DataClient
}

// DownloadLinkAndUpload 自动下载url并上传文件
func (s DataService) DownloadLinkAndUpload(ctx context.Context, url string, name string, header map[string]string) (string, error) {
	if url == "" {
		return "", nil
	}
	if !strings.HasPrefix(url, "http") {
		url = "http:" + url
	}
	// 初始化request
	req, err := http.NewRequest("GET", url, nil)
	for k, v := range header {
		req.Header.Set(k, v)
	}
	// 构建并发送请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", tools.ReturnError("send data error %v", err)
	}
	defer res.Body.Close()
	id, err := s.UploadFromReader(ctx, res.Body, name, res.ContentLength)
	if err != nil {
		return "", tools.ReturnError("upload data error %v", err)
	}
	return id, nil
}

// UploadFromReader 从reader中上传
func (s DataService) UploadFromReader(ctx context.Context, reader io.Reader, name string, size int64) (string, error) {
	// 获取流式上传的链接
	client, err := s.client.PutObject(ctx)
	if err != nil {
		log.CtxLogError(ctx, "rpc服务连接失败 %v", err)
		return "", err
	}
	var offset int64
	// 每次分片读取2M的文件
	tmp := make([]byte, 1024*1024*2)
	for {
		// 安装offset来进行读取
		n, err := reader.Read(tmp)
		if err == nil || err == io.EOF {
			// 当n大于0时才进行流式传输
			if n > 0 {
				err2 := client.Send(&data.ObjectPutInfo{Filename: name, Size: size, Offset: offset, Data: tmp[:n]})
				if err2 != nil {
					log.CtxLogError(ctx, "上传错误 %v", err2)
					return "", err
				}
			}
			// 如果是读取完毕，那么就停止流式传输
			if err == io.EOF {
				info, err := client.CloseAndRecv()
				if len(info.Id) > 0 {
					return info.Id[0], err
				}
				return "", errors.New("没有id")
			}
		} else {
			client.CloseAndRecv()
			log.CtxLogError(ctx, "读取文件失败 %v", err)
			return "", err
		}
		offset += int64(n)
	}
}

// UploadObjectFromFile 从gin的文件对象中上传文件
func (s DataService) UploadObjectFromFile(ctx context.Context, file *multipart.FileHeader) (string, error) {
	if file == nil {
		return "", errors.New("没有文件")
	}
	// 流式读取文件
	f, err := file.Open()
	if err != nil {
		log.CtxLogError(ctx, "打开文件失败 %v", err)
		return "", err
	}
	defer f.Close()
	return s.UploadFromReader(ctx, f, file.Filename, file.Size)
}

// AddText 上传文本对象
func (s DataService) AddText(ctx context.Context, name string, content string) (string, error) {
	res, err := s.client.AddText(ctx, &data.TextInfo{
		Name:    name,
		Content: content,
	})
	if err != nil {
		return "", err
	}
	if res.Msg != "" || len(res.Id) == 0 {
		return "", errors.New(res.Msg)
	}
	return res.Id[0], nil
}

// UpdateText 更新文本对象
func (s DataService) UpdateText(ctx context.Context, id string, name string, content string) error {
	res, err := s.client.UpdateText(ctx, &data.UpdateTextInfo{
		Id: id,
		Info: &data.TextInfo{
			Name:    name,
			Content: content,
		},
	})
	if err != nil {
		return err
	}
	if !res.Status {
		return errors.New(res.Msg)
	}
	return nil
}

// GetText 获取文本对象
func (s DataService) GetText(ctx context.Context, id string) (name string, content string, err error) {
	info, err := s.client.GetText(ctx, &data.ObjectMeta{Id: []string{id}})
	if err != nil {
		return "", "", err
	}
	if info != nil && info.List[id] != nil {
		return info.List[id].Name, info.List[id].Content, nil
	}
	return "", "", errors.New("文件不存在")
}

// DeleteObject 删除对象
func (s DataService) DeleteObject(ctx context.Context, id []string) error {
	res, err := s.client.DeleteObject(ctx, &data.ObjectMeta{Id: id})
	if err != nil {
		return err
	}
	if !res.Status {
		return errors.New(res.Msg)
	}
	return nil
}

// UploadFileChunk 分片上传文件
func (s DataService) UploadFileChunk(ctx context.Context, req *data.PutObjectChunkReq) (*data.PutObjectChunkResp, error) {
	return s.client.PutObjectChunk(ctx, req)
}

// GetDownloadLink 获取下载链接
func (s DataService) GetDownloadLink(ctx context.Context, id string) (string, error) {
	info, err := s.client.GetObjectDownloadLink(ctx, &data.ObjectMeta{Id: []string{id}})
	if err != nil {
		return "", err
	}
	if info != nil && info.Links[id] != "" {
		return info.Links[id], nil
	}
	return "", errors.New("文件不存在")
}
