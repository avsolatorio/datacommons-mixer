// Copyright 2021 Google LLC
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

//  **** IMPORTANT NOTE ****
//
//  The proto of BT data has to match exactly the g3 proto, including tag
//  number.

// REST API URL from the proto in this file:
// ========================================
//    /translate
// ========================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: translate.proto

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

// Request to translate a graph query.
type TranslateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String representaion of schema mappings used to translate.
	SchemaMapping string `protobuf:"bytes,1,opt,name=schema_mapping,json=schemaMapping,proto3" json:"schema_mapping,omitempty"`
	// String representation of sparql query.
	Sparql string `protobuf:"bytes,2,opt,name=sparql,proto3" json:"sparql,omitempty"`
}

func (x *TranslateRequest) Reset() {
	*x = TranslateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_translate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateRequest) ProtoMessage() {}

func (x *TranslateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_translate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateRequest.ProtoReflect.Descriptor instead.
func (*TranslateRequest) Descriptor() ([]byte, []int) {
	return file_translate_proto_rawDescGZIP(), []int{0}
}

func (x *TranslateRequest) GetSchemaMapping() string {
	if x != nil {
		return x.SchemaMapping
	}
	return ""
}

func (x *TranslateRequest) GetSparql() string {
	if x != nil {
		return x.Sparql
	}
	return ""
}

// Response of a translate request.
type TranslateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The translated sql.
	Sql string `protobuf:"bytes,1,opt,name=sql,proto3" json:"sql,omitempty"`
	// Serialized json string of the translation result
	Translation string `protobuf:"bytes,2,opt,name=translation,proto3" json:"translation,omitempty"`
}

func (x *TranslateResponse) Reset() {
	*x = TranslateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_translate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslateResponse) ProtoMessage() {}

func (x *TranslateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_translate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslateResponse.ProtoReflect.Descriptor instead.
func (*TranslateResponse) Descriptor() ([]byte, []int) {
	return file_translate_proto_rawDescGZIP(), []int{1}
}

func (x *TranslateResponse) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

func (x *TranslateResponse) GetTranslation() string {
	if x != nil {
		return x.Translation
	}
	return ""
}

var File_translate_proto protoreflect.FileDescriptor

var file_translate_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x22, 0x51,
	0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x70, 0x61,
	0x72, 0x71, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x61, 0x72, 0x71,
	0x6c, 0x22, 0x47, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x71, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_translate_proto_rawDescOnce sync.Once
	file_translate_proto_rawDescData = file_translate_proto_rawDesc
)

func file_translate_proto_rawDescGZIP() []byte {
	file_translate_proto_rawDescOnce.Do(func() {
		file_translate_proto_rawDescData = protoimpl.X.CompressGZIP(file_translate_proto_rawDescData)
	})
	return file_translate_proto_rawDescData
}

var file_translate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_translate_proto_goTypes = []interface{}{
	(*TranslateRequest)(nil),  // 0: datacommons.TranslateRequest
	(*TranslateResponse)(nil), // 1: datacommons.TranslateResponse
}
var file_translate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_translate_proto_init() }
func file_translate_proto_init() {
	if File_translate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_translate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateRequest); i {
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
		file_translate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslateResponse); i {
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
			RawDescriptor: file_translate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_translate_proto_goTypes,
		DependencyIndexes: file_translate_proto_depIdxs,
		MessageInfos:      file_translate_proto_msgTypes,
	}.Build()
	File_translate_proto = out.File
	file_translate_proto_rawDesc = nil
	file_translate_proto_goTypes = nil
	file_translate_proto_depIdxs = nil
}
