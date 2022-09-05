package neo4j

// GetNodeCount 计算节点个数
func (s Service) GetNodeCount(cypher string) int64 {
	session := s.newSession()
	// 构建标签
	res, err := session.Run(cypher+" as total", map[string]interface{}{})
	if err != nil {
		return 0
	}

	single, err := res.Single()
	if err != nil {
		return 0
	}
	if count, ok := single.Get("total"); ok {
		if total, ok := count.(int64); ok {
			return total
		}
	}

	return 0
}
