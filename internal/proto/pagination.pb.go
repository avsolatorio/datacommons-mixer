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
// source: v1/pagination.proto

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

type ImportGroupCursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The index of the import group.
	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	// The page index of the cursor, starts from 0.
	Page int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	// The position of the next read item in the current page, starts from 0.
	Pos int32 `protobuf:"varint,3,opt,name=pos,proto3" json:"pos,omitempty"`
}

func (x *ImportGroupCursor) Reset() {
	*x = ImportGroupCursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportGroupCursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportGroupCursor) ProtoMessage() {}

func (x *ImportGroupCursor) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportGroupCursor.ProtoReflect.Descriptor instead.
func (*ImportGroupCursor) Descriptor() ([]byte, []int) {
	return file_v1_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *ImportGroupCursor) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ImportGroupCursor) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ImportGroupCursor) GetPos() int32 {
	if x != nil {
		return x.Pos
	}
	return 0
}

type ImportGroupCursorGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Entity DCID or other information that identifies the ImportGroupCursorGroup.
	Key     string               `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Cursors []*ImportGroupCursor `protobuf:"bytes,2,rep,name=cursors,proto3" json:"cursors,omitempty"`
}

func (x *ImportGroupCursorGroup) Reset() {
	*x = ImportGroupCursorGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportGroupCursorGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportGroupCursorGroup) ProtoMessage() {}

func (x *ImportGroupCursorGroup) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportGroupCursorGroup.ProtoReflect.Descriptor instead.
func (*ImportGroupCursorGroup) Descriptor() ([]byte, []int) {
	return file_v1_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *ImportGroupCursorGroup) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ImportGroupCursorGroup) GetCursors() []*ImportGroupCursor {
	if x != nil {
		return x.Cursors
	}
	return nil
}

type PaginationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CursorGroups []*ImportGroupCursorGroup `protobuf:"bytes,1,rep,name=cursor_groups,json=cursorGroups,proto3" json:"cursor_groups,omitempty"`
}

func (x *PaginationInfo) Reset() {
	*x = PaginationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_pagination_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationInfo) ProtoMessage() {}

func (x *PaginationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_pagination_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationInfo.ProtoReflect.Descriptor instead.
func (*PaginationInfo) Descriptor() ([]byte, []int) {
	return file_v1_pagination_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationInfo) GetCursorGroups() []*ImportGroupCursorGroup {
	if x != nil {
		return x.CursorGroups
	}
	return nil
}

var File_v1_pagination_proto protoreflect.FileDescriptor

var file_v1_pagination_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x4f, 0x0a, 0x11, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x70, 0x6f, 0x73, 0x22, 0x67, 0x0a, 0x16, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x3b, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x07, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x73, 0x22,
	0x5d, 0x0a, 0x0e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x4b, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x0c, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_v1_pagination_proto_rawDescOnce sync.Once
	file_v1_pagination_proto_rawDescData = file_v1_pagination_proto_rawDesc
)

func file_v1_pagination_proto_rawDescGZIP() []byte {
	file_v1_pagination_proto_rawDescOnce.Do(func() {
		file_v1_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_pagination_proto_rawDescData)
	})
	return file_v1_pagination_proto_rawDescData
}

var file_v1_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1_pagination_proto_goTypes = []interface{}{
	(*ImportGroupCursor)(nil),      // 0: datacommons.v1.ImportGroupCursor
	(*ImportGroupCursorGroup)(nil), // 1: datacommons.v1.ImportGroupCursorGroup
	(*PaginationInfo)(nil),         // 2: datacommons.v1.PaginationInfo
}
var file_v1_pagination_proto_depIdxs = []int32{
	0, // 0: datacommons.v1.ImportGroupCursorGroup.cursors:type_name -> datacommons.v1.ImportGroupCursor
	1, // 1: datacommons.v1.PaginationInfo.cursor_groups:type_name -> datacommons.v1.ImportGroupCursorGroup
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1_pagination_proto_init() }
func file_v1_pagination_proto_init() {
	if File_v1_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportGroupCursor); i {
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
		file_v1_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportGroupCursorGroup); i {
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
		file_v1_pagination_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationInfo); i {
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
			RawDescriptor: file_v1_pagination_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_pagination_proto_goTypes,
		DependencyIndexes: file_v1_pagination_proto_depIdxs,
		MessageInfos:      file_v1_pagination_proto_msgTypes,
	}.Build()
	File_v1_pagination_proto = out.File
	file_v1_pagination_proto_rawDesc = nil
	file_v1_pagination_proto_goTypes = nil
	file_v1_pagination_proto_depIdxs = nil
}
