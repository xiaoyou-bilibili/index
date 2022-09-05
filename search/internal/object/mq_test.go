package object

import (
	"context"
	"fmt"
	"git.xiaoyou.host/index/common/service"
	"git.xiaoyou.host/index/common/tool/mq"
	"testing"
)

func TestProduce(t *testing.T) {
	produce := mq.CreateProduceService("index")
	err := produce.SendSync(context.Background(), "object", service.EsObjectOption{
		Option: 2,
		EsObject: service.EsObject{
			NodeID:  1,
			Name:    "22222",
			Content: "4522226",
			Tags:    []string{"22222", "33333"},
		},
	})
	fmt.Println(err)
}
