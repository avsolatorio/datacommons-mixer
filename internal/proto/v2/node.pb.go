// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: v2/node.proto

package v2

import (
	proto "github.com/datacommonsorg/mixer/internal/proto"
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

type Nodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*proto.EntityInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *Nodes) Reset() {
	*x = Nodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nodes) ProtoMessage() {}

func (x *Nodes) ProtoReflect() protoreflect.Message {
	mi := &file_v2_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nodes.ProtoReflect.Descriptor instead.
func (*Nodes) Descriptor() ([]byte, []int) {
	return file_v2_node_proto_rawDescGZIP(), []int{0}
}

func (x *Nodes) GetNodes() []*proto.EntityInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type LinkedGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keyed by property, with values being a list of nodes
	Arcs map[string]*Nodes `protobuf:"bytes,1,rep,name=arcs,proto3" json:"arcs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Keyed by property, with values being another LinkedGraph
	Neighbor map[string]*LinkedGraph `protobuf:"bytes,2,rep,name=neighbor,proto3" json:"neighbor,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Only set when querying for properties only, with "[nodes]->" or "[nodes]<-]".
	Properties []string `protobuf:"bytes,3,rep,name=properties,proto3" json:"properties,omitempty"`
}

func (x *LinkedGraph) Reset() {
	*x = LinkedGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkedGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkedGraph) ProtoMessage() {}

func (x *LinkedGraph) ProtoReflect() protoreflect.Message {
	mi := &file_v2_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkedGraph.ProtoReflect.Descriptor instead.
func (*LinkedGraph) Descriptor() ([]byte, []int) {
	return file_v2_node_proto_rawDescGZIP(), []int{1}
}

func (x *LinkedGraph) GetArcs() map[string]*Nodes {
	if x != nil {
		return x.Arcs
	}
	return nil
}

func (x *LinkedGraph) GetNeighbor() map[string]*LinkedGraph {
	if x != nil {
		return x.Neighbor
	}
	return nil
}

func (x *LinkedGraph) GetProperties() []string {
	if x != nil {
		return x.Properties
	}
	return nil
}

type NodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Nodes to query for.
	Nodes []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// The expression for the property or chained properties.
	// The property includes arrow notation, name property(ies) and filter
	// expression. For example:
	// <-containedInPlace+{typeOf: City}->name
	Property string `protobuf:"bytes,2,opt,name=property,proto3" json:"property,omitempty"`
	// Max number of result nodes to be returned for each query node.
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Pagination token
	NextToken string `protobuf:"bytes,4,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *NodeRequest) Reset() {
	*x = NodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeRequest) ProtoMessage() {}

func (x *NodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeRequest.ProtoReflect.Descriptor instead.
func (*NodeRequest) Descriptor() ([]byte, []int) {
	return file_v2_node_proto_rawDescGZIP(), []int{2}
}

func (x *NodeRequest) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *NodeRequest) GetProperty() string {
	if x != nil {
		return x.Property
	}
	return ""
}

func (x *NodeRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *NodeRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type NodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keyed by the query node dcid.
	Data map[string]*LinkedGraph `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The pagination token for getting the next set of entries.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *NodeResponse) Reset() {
	*x = NodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeResponse) ProtoMessage() {}

func (x *NodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeResponse.ProtoReflect.Descriptor instead.
func (*NodeResponse) Descriptor() ([]byte, []int) {
	return file_v2_node_proto_rawDescGZIP(), []int{3}
}

func (x *NodeResponse) GetData() map[string]*LinkedGraph {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *NodeResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

var File_v2_node_proto protoreflect.FileDescriptor

var file_v2_node_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x32, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x1a,
	0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x05, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xd9, 0x02, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x39, 0x0a, 0x04, 0x61, 0x72, 0x63, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x2e, 0x41, 0x72, 0x63, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x72, 0x63, 0x73,
	0x12, 0x45, 0x0a, 0x08, 0x6e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x4e, 0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x4e, 0x0a, 0x09, 0x41, 0x72, 0x63, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x58, 0x0a, 0x0d, 0x4e, 0x65, 0x69, 0x67, 0x68,
	0x62, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65,
	0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x74, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65,
	0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xbf, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x1a, 0x54, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x78, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_node_proto_rawDescOnce sync.Once
	file_v2_node_proto_rawDescData = file_v2_node_proto_rawDesc
)

func file_v2_node_proto_rawDescGZIP() []byte {
	file_v2_node_proto_rawDescOnce.Do(func() {
		file_v2_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_node_proto_rawDescData)
	})
	return file_v2_node_proto_rawDescData
}

var file_v2_node_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_v2_node_proto_goTypes = []interface{}{
	(*Nodes)(nil),            // 0: datacommons.v2.Nodes
	(*LinkedGraph)(nil),      // 1: datacommons.v2.LinkedGraph
	(*NodeRequest)(nil),      // 2: datacommons.v2.NodeRequest
	(*NodeResponse)(nil),     // 3: datacommons.v2.NodeResponse
	nil,                      // 4: datacommons.v2.LinkedGraph.ArcsEntry
	nil,                      // 5: datacommons.v2.LinkedGraph.NeighborEntry
	nil,                      // 6: datacommons.v2.NodeResponse.DataEntry
	(*proto.EntityInfo)(nil), // 7: datacommons.EntityInfo
}
var file_v2_node_proto_depIdxs = []int32{
	7, // 0: datacommons.v2.Nodes.nodes:type_name -> datacommons.EntityInfo
	4, // 1: datacommons.v2.LinkedGraph.arcs:type_name -> datacommons.v2.LinkedGraph.ArcsEntry
	5, // 2: datacommons.v2.LinkedGraph.neighbor:type_name -> datacommons.v2.LinkedGraph.NeighborEntry
	6, // 3: datacommons.v2.NodeResponse.data:type_name -> datacommons.v2.NodeResponse.DataEntry
	0, // 4: datacommons.v2.LinkedGraph.ArcsEntry.value:type_name -> datacommons.v2.Nodes
	1, // 5: datacommons.v2.LinkedGraph.NeighborEntry.value:type_name -> datacommons.v2.LinkedGraph
	1, // 6: datacommons.v2.NodeResponse.DataEntry.value:type_name -> datacommons.v2.LinkedGraph
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_v2_node_proto_init() }
func file_v2_node_proto_init() {
	if File_v2_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nodes); i {
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
		file_v2_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkedGraph); i {
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
		file_v2_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeRequest); i {
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
		file_v2_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeResponse); i {
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
			RawDescriptor: file_v2_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_node_proto_goTypes,
		DependencyIndexes: file_v2_node_proto_depIdxs,
		MessageInfos:      file_v2_node_proto_msgTypes,
	}.Build()
	File_v2_node_proto = out.File
	file_v2_node_proto_rawDesc = nil
	file_v2_node_proto_goTypes = nil
	file_v2_node_proto_depIdxs = nil
}
