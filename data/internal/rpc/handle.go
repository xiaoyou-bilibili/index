package rpc

import (
	"bytes"
	"context"
	"fmt"
	pb "git.xiaoyou.host/index/common/proto/data"
	"git.xiaoyou.host/index/common/tool/log"
	"index.data/internal/service"
	"io"
)

type DataService struct {
	pb.UnimplementedDataServer
	dataCache map[string][][]byte // 数据缓存，缓存当前上传的文件二进制内容
	sizeCache map[string]int64    // 大小缓存，缓存文件的大小信息
}

func NewDataService() *DataService {
	return &DataService{
		dataCache: map[string][][]byte{},
		sizeCache: map[string]int64{},
	}
}

// GetObject 获取对象
func (s *DataService) GetObject(req *pb.ObjectMeta, conn pb.Data_GetObjectServer) error {
	for {
		err := conn.Send(&pb.ObjectInfo{})
		if err != nil {
			return err
		}
	}
}

// PutObject 上传对象
func (s *DataService) PutObject(conn pb.Data_PutObjectServer) error {
	// 最终缓存好的数据
	buff := bytes.NewBuffer([]byte{})
	// 用于计算sha256的buffer
	shaBuff := bytes.NewBuffer([]byte{})
	var total int64
	var filename string
	for {
		req, err := conn.Recv()
		if req != nil {
			// 不断把数据写入到buffer中
			buff.Write(req.Data)
			shaBuff.Write(req.Data)
			total, filename = req.Size, req.Filename
		}
		if err == io.EOF {
			// 这里表示上传文件完毕
			data, err := service.AddObject(context.Background(), filename, total, buff, service.GetReaderSha256(shaBuff))
			fmt.Println(err)
			return conn.SendAndClose(&pb.ObjectMeta{
				Id: []string{data},
			})
		}
		if err != nil {
			return err
		}
	}
}

// DeleteObject 删除对象
func (s *DataService) DeleteObject(ctx context.Context, req *pb.ObjectMeta) (*pb.UpdateObjectResp, error) {
	// 目前使用最简单的遍历删除
	for _, v := range req.Id {
		err := service.DeleteObject(ctx, v)
		if err != nil {
			fmt.Println(err)
		}
	}
	return &pb.UpdateObjectResp{Status: true}, nil

}

// GetText 获取文本
func (s *DataService) GetText(ctx context.Context, req *pb.ObjectMeta) (*pb.GetTextResp, error) {
	res := make(map[string]*pb.TextInfo)
	for _, id := range req.Id {
		name, content, err := service.GetTextObject(ctx, req.Id[0])
		if err != nil {
			log.CtxLogError(ctx, "get text error %v", err)
			continue
		}
		res[id] = &pb.TextInfo{Name: name, Content: content}
	}

	return &pb.GetTextResp{List: res}, nil
}

// UpdateText 更新文本
func (s *DataService) UpdateText(ctx context.Context, req *pb.UpdateTextInfo) (*pb.UpdateObjectResp, error) {
	err := service.UpdateTextObject(ctx, req.Id, req.Info.Name, req.Info.Content)
	if err != nil {
		return &pb.UpdateObjectResp{Status: false, Msg: err.Error()}, nil
	}
	return &pb.UpdateObjectResp{Status: true}, nil
}

// AddText 添加文本
func (s *DataService) AddText(ctx context.Context, req *pb.TextInfo) (*pb.ObjectMeta, error) {
	id, err := service.AddTextObject(ctx, req.Name, req.Content)
	if err != nil {
		return &pb.ObjectMeta{Msg: err.Error()}, nil
	}
	return &pb.ObjectMeta{Id: []string{id}}, nil
}

func (s *DataService) PutObjectChunk(ctx context.Context, req *pb.PutObjectChunkReq) (*pb.PutObjectChunkResp, error) {
	// 因为上传的文件是乱序的，如果不在map中，就需要自己手动创建一个数组并初始化来保存数据
	if _, ok := s.dataCache[req.Md5]; !ok {
		s.dataCache[req.Md5] = make([][]byte, req.TotalChunk)
	}
	// 把文件的二进制信息保存到数组中，同时记录已经保存的文件大小
	s.dataCache[req.Md5][req.Current-1] = req.Data
	s.sizeCache[req.Md5] += int64(len(req.Data))
	// 如果上传完成就可以直接上传到对象存储了
	if req.TotalSize == s.sizeCache[req.Md5] {
		// 使用两个tmp来存储文件,一个用来计算sha256，一个用来保存最后的文件
		shaTmp := bytes.NewBuffer([]byte{})
		dataTmp := bytes.NewBuffer([]byte{})
		for _, data := range s.dataCache[req.Md5] {
			shaTmp.Write(data)
			dataTmp.Write(data)
		}
		id, err := service.AddObject(context.Background(), req.Name, int64(dataTmp.Len()), dataTmp, service.GetReaderSha256(shaTmp))
		if err != nil {
			return &pb.PutObjectChunkResp{Status: false, Msg: fmt.Sprintf("add object error %v", err)}, nil
		}
		// 把map给删掉，及时释放内存
		delete(s.dataCache, req.Md5)
		delete(s.sizeCache, req.Md5)
		return &pb.PutObjectChunkResp{
			Receive: s.sizeCache[req.Md5],
			Status:  true,
			Id:      id,
		}, nil
	}
	return &pb.PutObjectChunkResp{Receive: s.sizeCache[req.Md5], Status: true}, nil
}

func (s *DataService) GetObjectDownloadLink(ctx context.Context, req *pb.ObjectMeta) (*pb.GetObjectDownloadLinkResp, error) {
	res := make(map[string]string)
	for _, id := range req.Id {
		link, err := service.GetObjectLink(ctx, id)
		if err != nil {
			log.CtxLogError(ctx, "get link error %v", err)
			continue
		}
		res[id] = link
	}

	return &pb.GetObjectDownloadLinkResp{Links: res}, nil
}
