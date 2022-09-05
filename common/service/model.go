package service

// Music 音乐
type Music struct {
	ID int64 `json:"id"` // 音乐ID
	MusicInfo
}

// MusicInfo 音乐信息
type MusicInfo struct {
	Name   string `json:"name"`   // 音乐名字
	Pic    string `json:"pic"`    // 音乐图片
	Audio  string `json:"audio"`  // 播放路径
	Lrc    string `json:"lrc"`    // 歌词路径
	Singer string `json:"singer"` // 歌手
	QQMid  string `json:"qq_mid"` // QQ音乐ID
}

// People 人物
type People struct {
	ID int64 `json:"id"`
	PeopleInfo
}

// PeopleInfo 人物信息
type PeopleInfo struct {
	Name  string `json:"name"`   // 名字
	Pic   string `json:"pic"`    // 人物图片
	Desc  string `json:"desc"`   // 人物描述
	Birth string `json:"birth"`  // 出生日期
	QQMid string `json:"qq_mid"` // QQ音乐歌手ID
}

// Album 专辑信息
type Album struct {
	ID int64 `json:"id"` // 专辑ID
	AlbumInfo
}

// AlbumInfo 专辑
type AlbumInfo struct {
	Name string `json:"name"` // 专辑名字
	Pic  string `json:"pic"`  // 封面
	Desc string `json:"desc"` // 专辑描述
}

// EsObject 对象信息
type EsObject struct {
	NodeID  int64    `json:"node_id"` // 节点ID
	Name    string   `json:"name"`    // 节点名字
	App     string   `json:"app"`     // 节点所属应用
	Tags    []string `json:"tags"`    // 节点标签
	Content string   `json:"content"` // 节点内容

}

// EsObjectOption ES消息队列模型
type EsObjectOption struct {
	Option int32 `json:"option"` // 操作类型 1 新增 2 删除 3 更新
	EsObject
}
