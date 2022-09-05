package service

import (
	"context"

	pb "git.xiaoyou.host/index/common/proto/app"
)

type AppService struct {
	pb.UnimplementedAppServer
}

func NewAppService() *AppService {
	return &AppService{}
}

func (s *AppService) DeleteNode(ctx context.Context, req *pb.NoeMeta) (*pb.DeleteNodeResp, error) {
	return &pb.DeleteNodeResp{}, nil
}
func (s *AppService) GetNodeInfo(ctx context.Context, req *pb.GetNodeInfoReq) (*pb.GetNodeInfoResp, error) {
	return &pb.GetNodeInfoResp{}, nil
}
