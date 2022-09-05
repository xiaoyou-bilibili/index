package rpc

import (
	"context"
	pb "git.xiaoyou.host/index/common/proto/app"
	"index.app/internal/plugin"
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
	res := make([]*pb.CardInfo, 0, len(req.Infos))
	// 遍历所有的卡片
	for _, card := range req.Infos {
		// 判断卡片是否存在
		if container := plugin.GetPluginInfo(card.App); container != nil {
			// 调用搜索接口
			if info := container.GetSearchCard(ctx, card.NodeId); info != "" {
				res = append(res, &pb.CardInfo{
					NodeId: card.NodeId,
					App:    card.App,
					Info:   info,
				})
			}
		}
	}
	return &pb.GetNodeInfoResp{
		Cards: res,
	}, nil
}
