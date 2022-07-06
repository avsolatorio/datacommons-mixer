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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: v1/page.proto

package proto

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

type BioPageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity string `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
}

func (x *BioPageRequest) Reset() {
	*x = BioPageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_page_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BioPageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BioPageRequest) ProtoMessage() {}

func (x *BioPageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_page_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BioPageRequest.ProtoReflect.Descriptor instead.
func (*BioPageRequest) Descriptor() ([]byte, []int) {
	return file_v1_page_proto_rawDescGZIP(), []int{0}
}

func (x *BioPageRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

type PlacePageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity string `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	// A list of additional stat vars need to be fetched in addition to the
	// data in cache. This is to be used in local development, where new chart is
	// added to chart config but data is not added to cache (delay from cache
	// build).
	NewStatVars []string `protobuf:"bytes,2,rep,name=new_stat_vars,json=newStatVars,proto3" json:"new_stat_vars,omitempty"`
	// Seed value for random selection. Used by test to get deterministic result.
	Seed int64 `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (x *PlacePageRequest) Reset() {
	*x = PlacePageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_page_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacePageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacePageRequest) ProtoMessage() {}

func (x *PlacePageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_page_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacePageRequest.ProtoReflect.Descriptor instead.
func (*PlacePageRequest) Descriptor() ([]byte, []int) {
	return file_v1_page_proto_rawDescGZIP(), []int{1}
}

func (x *PlacePageRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *PlacePageRequest) GetNewStatVars() []string {
	if x != nil {
		return x.NewStatVars
	}
	return nil
}

func (x *PlacePageRequest) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

var File_v1_page_proto protoreflect.FileDescriptor

var file_v1_page_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22,
	0x28, 0x0a, 0x0e, 0x42, 0x69, 0x6f, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x62, 0x0a, 0x10, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x76, 0x61, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65,
	0x77, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_page_proto_rawDescOnce sync.Once
	file_v1_page_proto_rawDescData = file_v1_page_proto_rawDesc
)

func file_v1_page_proto_rawDescGZIP() []byte {
	file_v1_page_proto_rawDescOnce.Do(func() {
		file_v1_page_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_page_proto_rawDescData)
	})
	return file_v1_page_proto_rawDescData
}

var file_v1_page_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_page_proto_goTypes = []interface{}{
	(*BioPageRequest)(nil),   // 0: datacommons.v1.BioPageRequest
	(*PlacePageRequest)(nil), // 1: datacommons.v1.PlacePageRequest
}
var file_v1_page_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_page_proto_init() }
func file_v1_page_proto_init() {
	if File_v1_page_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_page_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BioPageRequest); i {
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
		file_v1_page_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlacePageRequest); i {
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
			RawDescriptor: file_v1_page_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_page_proto_goTypes,
		DependencyIndexes: file_v1_page_proto_depIdxs,
		MessageInfos:      file_v1_page_proto_msgTypes,
	}.Build()
	File_v1_page_proto = out.File
	file_v1_page_proto_rawDesc = nil
	file_v1_page_proto_goTypes = nil
	file_v1_page_proto_depIdxs = nil
}
