package tools

import (
	"fmt"
)

// GetObjectLink 获取节点直链
func GetObjectLink(id string) string {
	return fmt.Sprintf("https://index.xiaoyou.host/data/object/%s", id)
}
