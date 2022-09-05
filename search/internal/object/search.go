package object

import (
	"context"
	"git.xiaoyou.host/index/common/service"
)

func SearchKeyword(ctx context.Context, keyword string) ([]service.EsObject, error) {
	res, err := esService.SearchNode(ctx, "object", keyword)
	if err != nil {
		return nil, err
	}
	objects := make([]service.EsObject, 0, len(res.Hits.Hits))
	for _, v := range res.Hits.Hits {
		objects = append(objects, v.Source)
	}
	return objects, nil
}
