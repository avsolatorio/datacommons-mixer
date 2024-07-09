// Copyright 2022 Google LLC
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
// source: v1/triples.proto

package v1

import (
	proto "github.com/avsolatorio/datacommons-mixer/internal/proto"
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

// Basic info for a collection of nodes.
type NodeInfoCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*proto.EntityInfo `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *NodeInfoCollection) Reset() {
	*x = NodeInfoCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_triples_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfoCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfoCollection) ProtoMessage() {}

func (x *NodeInfoCollection) ProtoReflect() protoreflect.Message {
	mi := &file_v1_triples_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfoCollection.ProtoReflect.Descriptor instead.
func (*NodeInfoCollection) Descriptor() ([]byte, []int) {
	return file_v1_triples_proto_rawDescGZIP(), []int{0}
}

func (x *NodeInfoCollection) GetNodes() []*proto.EntityInfo {
	if x != nil {
		return x.Nodes
	}
	return nil
}

type TriplesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Direction can only be "in" and "out"
	Direction string `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
	// [Optional]
	// The pagination token for getting the next set of entries. This is empty
	// for the first request and needs to be set in the subsequent request.
	// This is the value returned from a prior call to TriplesRequest
	NextToken string `protobuf:"bytes,3,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *TriplesRequest) Reset() {
	*x = TriplesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_triples_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriplesRequest) ProtoMessage() {}

func (x *TriplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_triples_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriplesRequest.ProtoReflect.Descriptor instead.
func (*TriplesRequest) Descriptor() ([]byte, []int) {
	return file_v1_triples_proto_rawDescGZIP(), []int{1}
}

func (x *TriplesRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *TriplesRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *TriplesRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type TriplesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is property.
	Triples map[string]*NodeInfoCollection `protobuf:"bytes,1,rep,name=triples,proto3" json:"triples,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The pagination token for getting the next set of entries.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *TriplesResponse) Reset() {
	*x = TriplesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_triples_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TriplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TriplesResponse) ProtoMessage() {}

func (x *TriplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_triples_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TriplesResponse.ProtoReflect.Descriptor instead.
func (*TriplesResponse) Descriptor() ([]byte, []int) {
	return file_v1_triples_proto_rawDescGZIP(), []int{2}
}

func (x *TriplesResponse) GetTriples() map[string]*NodeInfoCollection {
	if x != nil {
		return x.Triples
	}
	return nil
}

func (x *TriplesResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type BulkTriplesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []string `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	// Direction can only be "in" and "out"
	Direction string `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
	// [Optional]
	// The pagination token for getting the next set of entries. This is empty
	// for the first request and needs to be set in the subsequent request.
	// This is the value returned from a prior call to BulkTriplesRequest
	NextToken string `protobuf:"bytes,3,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *BulkTriplesRequest) Reset() {
	*x = BulkTriplesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_triples_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkTriplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkTriplesRequest) ProtoMessage() {}

func (x *BulkTriplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_triples_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkTriplesRequest.ProtoReflect.Descriptor instead.
func (*BulkTriplesRequest) Descriptor() ([]byte, []int) {
	return file_v1_triples_proto_rawDescGZIP(), []int{3}
}

func (x *BulkTriplesRequest) GetNodes() []string {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *BulkTriplesRequest) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *BulkTriplesRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type BulkTriplesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BulkTriplesResponse_NodeTriples `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	// The pagination token for getting the next set of entries.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *BulkTriplesResponse) Reset() {
	*x = BulkTriplesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_triples_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkTriplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkTriplesResponse) ProtoMessage() {}

func (x *BulkTriplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_triples_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkTriplesResponse.ProtoReflect.Descriptor instead.
func (*BulkTriplesResponse) Descriptor() ([]byte, []int) {
	return file_v1_triples_proto_rawDescGZIP(), []int{4}
}

func (x *BulkTriplesResponse) GetData() []*BulkTriplesResponse_NodeTriples {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *BulkTriplesResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

type BulkTriplesResponse_NodeTriples struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Node string `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// Key is property.
	Triples map[string]*NodeInfoCollection `protobuf:"bytes,2,rep,name=triples,proto3" json:"triples,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BulkTriplesResponse_NodeTriples) Reset() {
	*x = BulkTriplesResponse_NodeTriples{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_triples_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkTriplesResponse_NodeTriples) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkTriplesResponse_NodeTriples) ProtoMessage() {}

func (x *BulkTriplesResponse_NodeTriples) ProtoReflect() protoreflect.Message {
	mi := &file_v1_triples_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkTriplesResponse_NodeTriples.ProtoReflect.Descriptor instead.
func (*BulkTriplesResponse_NodeTriples) Descriptor() ([]byte, []int) {
	return file_v1_triples_proto_rawDescGZIP(), []int{4, 0}
}

func (x *BulkTriplesResponse_NodeTriples) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *BulkTriplesResponse_NodeTriples) GetTriples() map[string]*NodeInfoCollection {
	if x != nil {
		return x.Triples
	}
	return nil
}

var File_v1_triples_proto protoreflect.FileDescriptor

var file_v1_triples_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x43, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x61, 0x0a, 0x0e, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd8, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x07,
	0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54,
	0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x1a, 0x5e, 0x0a, 0x0c, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x67, 0x0a, 0x12, 0x42, 0x75, 0x6c, 0x6b, 0x54, 0x72, 0x69, 0x70, 0x6c,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd5, 0x02, 0x0a,
	0x13, 0x42, 0x75, 0x6c, 0x6b, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x69, 0x70,
	0x6c, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0xd9, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64,
	0x65, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x56, 0x0a, 0x07,
	0x74, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x75, 0x6c, 0x6b, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x54,
	0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x72, 0x69,
	0x70, 0x6c, 0x65, 0x73, 0x1a, 0x5e, 0x0a, 0x0c, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x6f, 0x72,
	0x67, 0x2f, 0x6d, 0x69, 0x78, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_triples_proto_rawDescOnce sync.Once
	file_v1_triples_proto_rawDescData = file_v1_triples_proto_rawDesc
)

func file_v1_triples_proto_rawDescGZIP() []byte {
	file_v1_triples_proto_rawDescOnce.Do(func() {
		file_v1_triples_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_triples_proto_rawDescData)
	})
	return file_v1_triples_proto_rawDescData
}

var file_v1_triples_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_triples_proto_goTypes = []interface{}{
	(*NodeInfoCollection)(nil),              // 0: datacommons.v1.NodeInfoCollection
	(*TriplesRequest)(nil),                  // 1: datacommons.v1.TriplesRequest
	(*TriplesResponse)(nil),                 // 2: datacommons.v1.TriplesResponse
	(*BulkTriplesRequest)(nil),              // 3: datacommons.v1.BulkTriplesRequest
	(*BulkTriplesResponse)(nil),             // 4: datacommons.v1.BulkTriplesResponse
	nil,                                     // 5: datacommons.v1.TriplesResponse.TriplesEntry
	(*BulkTriplesResponse_NodeTriples)(nil), // 6: datacommons.v1.BulkTriplesResponse.NodeTriples
	nil,                                     // 7: datacommons.v1.BulkTriplesResponse.NodeTriples.TriplesEntry
	(*proto.EntityInfo)(nil),                // 8: datacommons.EntityInfo
}
var file_v1_triples_proto_depIdxs = []int32{
	8, // 0: datacommons.v1.NodeInfoCollection.nodes:type_name -> datacommons.EntityInfo
	5, // 1: datacommons.v1.TriplesResponse.triples:type_name -> datacommons.v1.TriplesResponse.TriplesEntry
	6, // 2: datacommons.v1.BulkTriplesResponse.data:type_name -> datacommons.v1.BulkTriplesResponse.NodeTriples
	0, // 3: datacommons.v1.TriplesResponse.TriplesEntry.value:type_name -> datacommons.v1.NodeInfoCollection
	7, // 4: datacommons.v1.BulkTriplesResponse.NodeTriples.triples:type_name -> datacommons.v1.BulkTriplesResponse.NodeTriples.TriplesEntry
	0, // 5: datacommons.v1.BulkTriplesResponse.NodeTriples.TriplesEntry.value:type_name -> datacommons.v1.NodeInfoCollection
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_v1_triples_proto_init() }
func file_v1_triples_proto_init() {
	if File_v1_triples_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_triples_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfoCollection); i {
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
		file_v1_triples_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriplesRequest); i {
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
		file_v1_triples_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TriplesResponse); i {
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
		file_v1_triples_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkTriplesRequest); i {
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
		file_v1_triples_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkTriplesResponse); i {
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
		file_v1_triples_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkTriplesResponse_NodeTriples); i {
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
			RawDescriptor: file_v1_triples_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_triples_proto_goTypes,
		DependencyIndexes: file_v1_triples_proto_depIdxs,
		MessageInfos:      file_v1_triples_proto_msgTypes,
	}.Build()
	File_v1_triples_proto = out.File
	file_v1_triples_proto_rawDesc = nil
	file_v1_triples_proto_goTypes = nil
	file_v1_triples_proto_depIdxs = nil
}
