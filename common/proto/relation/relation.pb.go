// 定义我们接口的版本

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: proto/relation/relation.proto

// 定义包名称

package relation

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 删除响应
type UpdateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`        // 消息提示
}

func (x *UpdateResp) Reset() {
	*x = UpdateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResp) ProtoMessage() {}

func (x *UpdateResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResp.ProtoReflect.Descriptor instead.
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateResp) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *UpdateResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 添加节点响应
type NodeMeta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"` // 节点或联系ID
	Msg string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`       // 消息提示
}

func (x *NodeMeta) Reset() {
	*x = NodeMeta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMeta) ProtoMessage() {}

func (x *NodeMeta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMeta.ProtoReflect.Descriptor instead.
func (*NodeMeta) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{1}
}

func (x *NodeMeta) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *NodeMeta) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// 对象信息
type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeLabel []string          `protobuf:"bytes,1,rep,name=node_label,json=nodeLabel,proto3" json:"node_label,omitempty"`                                                                        // 节点标签
	Attribute map[string]string `protobuf:"bytes,2,rep,name=attribute,proto3" json:"attribute,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 节点属性
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInfo) GetNodeLabel() []string {
	if x != nil {
		return x.NodeLabel
	}
	return nil
}

func (x *NodeInfo) GetAttribute() map[string]string {
	if x != nil {
		return x.Attribute
	}
	return nil
}

// 联系信息
type RelationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start     int64             `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`                                                                                                // 起点
	End       int64             `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`                                                                                                    // 终点
	Relation  string            `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`                                                                                           // 节点关系
	Attribute map[string]string `protobuf:"bytes,4,rep,name=attribute,proto3" json:"attribute,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // 节点属性
}

func (x *RelationInfo) Reset() {
	*x = RelationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationInfo) ProtoMessage() {}

func (x *RelationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationInfo.ProtoReflect.Descriptor instead.
func (*RelationInfo) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{3}
}

func (x *RelationInfo) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *RelationInfo) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

func (x *RelationInfo) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *RelationInfo) GetAttribute() map[string]string {
	if x != nil {
		return x.Attribute
	}
	return nil
}

// 更新联系信息
type UpdateRelationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    // 联系id
	Info *RelationInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"` // 联系信息
}

func (x *UpdateRelationReq) Reset() {
	*x = UpdateRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRelationReq) ProtoMessage() {}

func (x *UpdateRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRelationReq.ProtoReflect.Descriptor instead.
func (*UpdateRelationReq) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateRelationReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRelationReq) GetInfo() *RelationInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// 根据开始和结束节点删除ID
type DeleteNodeWithRelationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"` // 开始ID
	End   int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`     // 结束ID
}

func (x *DeleteNodeWithRelationReq) Reset() {
	*x = DeleteNodeWithRelationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNodeWithRelationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNodeWithRelationReq) ProtoMessage() {}

func (x *DeleteNodeWithRelationReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNodeWithRelationReq.ProtoReflect.Descriptor instead.
func (*DeleteNodeWithRelationReq) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNodeWithRelationReq) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *DeleteNodeWithRelationReq) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type GetNodeChildReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`            // 节点ID
	Relation string `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty"` // 过滤联系的类型
	Current  int32  `protobuf:"varint,3,opt,name=current,proto3" json:"current,omitempty"`  // 当前第几页
	Size     int32  `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`        // 每页数量
}

func (x *GetNodeChildReq) Reset() {
	*x = GetNodeChildReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeChildReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeChildReq) ProtoMessage() {}

func (x *GetNodeChildReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeChildReq.ProtoReflect.Descriptor instead.
func (*GetNodeChildReq) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{6}
}

func (x *GetNodeChildReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetNodeChildReq) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *GetNodeChildReq) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *GetNodeChildReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// 查找某个节点下所有的内容
type GetNodeChildResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node   `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`  // 所有的子节点
	Info  *NodeInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`    // 当前节点信息
	Total int64     `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"` // 总数据条数
}

func (x *GetNodeChildResp) Reset() {
	*x = GetNodeChildResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeChildResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeChildResp) ProtoMessage() {}

func (x *GetNodeChildResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeChildResp.ProtoReflect.Descriptor instead.
func (*GetNodeChildResp) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{7}
}

func (x *GetNodeChildResp) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *GetNodeChildResp) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *GetNodeChildResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

// 节点
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    // 节点id
	Info *NodeInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"` // 节点信息
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{8}
}

func (x *Node) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetInfo() *NodeInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// 模糊查找节点请求
type FindNodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeLabel string `protobuf:"bytes,1,opt,name=nodeLabel,proto3" json:"nodeLabel,omitempty"` // 节点标签
	Field     string `protobuf:"bytes,2,opt,name=field,proto3" json:"field,omitempty"`         // 待查找的字段
	Keyword   string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`     // 查找的关键词
	Current   int32  `protobuf:"varint,4,opt,name=current,proto3" json:"current,omitempty"`    // 当前第几页
	Size      int32  `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`          // 每页数量
}

func (x *FindNodeReq) Reset() {
	*x = FindNodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNodeReq) ProtoMessage() {}

func (x *FindNodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNodeReq.ProtoReflect.Descriptor instead.
func (*FindNodeReq) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{9}
}

func (x *FindNodeReq) GetNodeLabel() string {
	if x != nil {
		return x.NodeLabel
	}
	return ""
}

func (x *FindNodeReq) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *FindNodeReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *FindNodeReq) GetCurrent() int32 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *FindNodeReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

// 模糊查找节点返回结果
type FindNodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeList []*Node `protobuf:"bytes,1,rep,name=nodeList,proto3" json:"nodeList,omitempty"` // 查找到的节点
	Total    int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`      // 总数据条数
}

func (x *FindNodeResp) Reset() {
	*x = FindNodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_relation_relation_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNodeResp) ProtoMessage() {}

func (x *FindNodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_relation_relation_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNodeResp.ProtoReflect.Descriptor instead.
func (*FindNodeResp) Descriptor() ([]byte, []int) {
	return file_proto_relation_relation_proto_rawDescGZIP(), []int{10}
}

func (x *FindNodeResp) GetNodeList() []*Node {
	if x != nil {
		return x.NodeList
	}
	return nil
}

func (x *FindNodeResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_relation_relation_proto protoreflect.FileDescriptor

var file_proto_relation_relation_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x36, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2c, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xa3, 0x01, 0x0a, 0x08, 0x4e,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xd0, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x1a, 0x3c, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0x43, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x65, 0x6e, 0x64, 0x22, 0x6b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x43,
	0x68, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x22, 0x6c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x69, 0x6c,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22,
	0x39, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x46,
	0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x6f,
	0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4b, 0x0a, 0x0c, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x32, 0xf6, 0x04, 0x0a, 0x08, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f,
	0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x23, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x09, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x50,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d,
	0x65, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x49, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x39, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x1d, 0x5a, 0x1b,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_relation_relation_proto_rawDescOnce sync.Once
	file_proto_relation_relation_proto_rawDescData = file_proto_relation_relation_proto_rawDesc
)

func file_proto_relation_relation_proto_rawDescGZIP() []byte {
	file_proto_relation_relation_proto_rawDescOnce.Do(func() {
		file_proto_relation_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_relation_relation_proto_rawDescData)
	})
	return file_proto_relation_relation_proto_rawDescData
}

var file_proto_relation_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_relation_relation_proto_goTypes = []interface{}{
	(*UpdateResp)(nil),                // 0: api.UpdateResp
	(*NodeMeta)(nil),                  // 1: api.NodeMeta
	(*NodeInfo)(nil),                  // 2: api.NodeInfo
	(*RelationInfo)(nil),              // 3: api.RelationInfo
	(*UpdateRelationReq)(nil),         // 4: api.UpdateRelationReq
	(*DeleteNodeWithRelationReq)(nil), // 5: api.DeleteNodeWithRelationReq
	(*GetNodeChildReq)(nil),           // 6: api.GetNodeChildReq
	(*GetNodeChildResp)(nil),          // 7: api.GetNodeChildResp
	(*Node)(nil),                      // 8: api.Node
	(*FindNodeReq)(nil),               // 9: api.FindNodeReq
	(*FindNodeResp)(nil),              // 10: api.FindNodeResp
	nil,                               // 11: api.NodeInfo.AttributeEntry
	nil,                               // 12: api.RelationInfo.AttributeEntry
}
var file_proto_relation_relation_proto_depIdxs = []int32{
	11, // 0: api.NodeInfo.attribute:type_name -> api.NodeInfo.AttributeEntry
	12, // 1: api.RelationInfo.attribute:type_name -> api.RelationInfo.AttributeEntry
	3,  // 2: api.UpdateRelationReq.info:type_name -> api.RelationInfo
	8,  // 3: api.GetNodeChildResp.nodes:type_name -> api.Node
	2,  // 4: api.GetNodeChildResp.info:type_name -> api.NodeInfo
	2,  // 5: api.Node.info:type_name -> api.NodeInfo
	8,  // 6: api.FindNodeResp.nodeList:type_name -> api.Node
	2,  // 7: api.Relation.AddNode:input_type -> api.NodeInfo
	1,  // 8: api.Relation.DeleteNode:input_type -> api.NodeMeta
	8,  // 9: api.Relation.UpdateNode:input_type -> api.Node
	1,  // 10: api.Relation.GetNode:input_type -> api.NodeMeta
	6,  // 11: api.Relation.GetNodeChild:input_type -> api.GetNodeChildReq
	9,  // 12: api.Relation.FindNode:input_type -> api.FindNodeReq
	6,  // 13: api.Relation.GetNodeParent:input_type -> api.GetNodeChildReq
	3,  // 14: api.Relation.AddRelation:input_type -> api.RelationInfo
	1,  // 15: api.Relation.DeleteRelation:input_type -> api.NodeMeta
	5,  // 16: api.Relation.DeleteRelationWithNode:input_type -> api.DeleteNodeWithRelationReq
	4,  // 17: api.Relation.UpdateRelation:input_type -> api.UpdateRelationReq
	1,  // 18: api.Relation.GetRelation:input_type -> api.NodeMeta
	1,  // 19: api.Relation.AddNode:output_type -> api.NodeMeta
	0,  // 20: api.Relation.DeleteNode:output_type -> api.UpdateResp
	0,  // 21: api.Relation.UpdateNode:output_type -> api.UpdateResp
	8,  // 22: api.Relation.GetNode:output_type -> api.Node
	7,  // 23: api.Relation.GetNodeChild:output_type -> api.GetNodeChildResp
	10, // 24: api.Relation.FindNode:output_type -> api.FindNodeResp
	7,  // 25: api.Relation.GetNodeParent:output_type -> api.GetNodeChildResp
	1,  // 26: api.Relation.AddRelation:output_type -> api.NodeMeta
	0,  // 27: api.Relation.DeleteRelation:output_type -> api.UpdateResp
	0,  // 28: api.Relation.DeleteRelationWithNode:output_type -> api.UpdateResp
	0,  // 29: api.Relation.UpdateRelation:output_type -> api.UpdateResp
	3,  // 30: api.Relation.GetRelation:output_type -> api.RelationInfo
	19, // [19:31] is the sub-list for method output_type
	7,  // [7:19] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_relation_relation_proto_init() }
func file_proto_relation_relation_proto_init() {
	if File_proto_relation_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_relation_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMeta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelationInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRelationReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNodeWithRelationReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeChildReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeChildResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNodeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_relation_relation_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNodeResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_relation_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_relation_relation_proto_goTypes,
		DependencyIndexes: file_proto_relation_relation_proto_depIdxs,
		MessageInfos:      file_proto_relation_relation_proto_msgTypes,
	}.Build()
	File_proto_relation_relation_proto = out.File
	file_proto_relation_relation_proto_rawDesc = nil
	file_proto_relation_relation_proto_goTypes = nil
	file_proto_relation_relation_proto_depIdxs = nil
}
