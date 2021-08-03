// Copyright 2019 Google LLC
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

// Mixer service definitions

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: mixer.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_mixer_proto protoreflect.FileDescriptor

var file_mixer_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x69, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x89, 0x20, 0x0a, 0x05, 0x4d, 0x69, 0x78, 0x65, 0x72, 0x12,
	0x5b, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x06, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5a,
	0x0b, 0x22, 0x06, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x9d, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x25, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x12, 0x15, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x5a, 0x1a, 0x22, 0x15, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x2d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x9d, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x25, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x12, 0x15, 0x2f, 0x6e, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x5a, 0x1a, 0x22, 0x15, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x2d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x69, 0x70,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x23, 0x12, 0x0d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x74, 0x72, 0x69, 0x70, 0x6c,
	0x65, 0x73, 0x5a, 0x12, 0x22, 0x0d, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x74, 0x72, 0x69, 0x70,
	0x6c, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x49, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x49, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27,
	0x12, 0x0f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x2d, 0x69,
	0x6e, 0x5a, 0x14, 0x22, 0x0f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x2d, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x79, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x4f, 0x62, 0x73, 0x12, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x62, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x56, 0x4f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x0f, 0x2f, 0x62, 0x75,
	0x6c, 0x6b, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6f, 0x62, 0x73, 0x5a, 0x14, 0x22, 0x0f,
	0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2d, 0x6f, 0x62, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x6e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1c,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x12, 0x0b, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x5a, 0x10, 0x22, 0x0b, 0x2f, 0x62, 0x75, 0x6c, 0x6b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x96, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x74,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x53, 0x65, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x13, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x73, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x5a, 0x18, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x73, 0x65,
	0x74, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x7a, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x20, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x0b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x5a, 0x10, 0x22, 0x0b, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x7f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12, 0x0c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x5a, 0x11, 0x22, 0x0c, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12,
	0x09, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x61, 0x6c, 0x6c, 0x5a, 0x0e, 0x22, 0x09, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x2f, 0x61, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0xa0, 0x01, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x74, 0x57, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x12, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x74, 0x57, 0x69, 0x74,
	0x68, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x16, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f,
	0x73, 0x65, 0x74, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2d, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x5a, 0x1b, 0x22, 0x16, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x73, 0x65, 0x74, 0x2f, 0x77, 0x69,
	0x74, 0x68, 0x69, 0x6e, 0x2d, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x53, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1b, 0x12, 0x09, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x73, 0x65, 0x74, 0x5a,
	0x0e, 0x22, 0x09, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2f, 0x73, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12,
	0xaa, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x28, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x37, 0x12, 0x17, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5a, 0x1c,
	0x22, 0x17, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67, 0x2d,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xa7, 0x01, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x12,
	0x17, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5a, 0x1c, 0x22, 0x17, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x2d, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xb3, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x52, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4c,
	0x12, 0x0d, 0x2f, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x61, 0x67, 0x65, 0x5a,
	0x12, 0x22, 0x0d, 0x2f, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x61, 0x67, 0x65,
	0x3a, 0x01, 0x2a, 0x5a, 0x11, 0x12, 0x0f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5a, 0x14, 0x22, 0x0f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x87, 0x01, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x26, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x50, 0x61, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x4e,
	0x6f, 0x64, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x11, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x5a, 0x16,
	0x22, 0x11, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x65, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x0a, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x5a, 0x0f, 0x22, 0x0a, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6c, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x52, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x09, 0x12, 0x07, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x5f, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x90, 0x01, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x61,
	0x72, 0x12, 0x24, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x56, 0x61, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x56, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x10, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x2d, 0x76, 0x61, 0x72, 0x5a, 0x15, 0x22, 0x10, 0x2f, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2d, 0x76, 0x61, 0x72, 0x3a, 0x01, 0x2a, 0x12,
	0x90, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x56, 0x61, 0x72, 0x73, 0x12, 0x24, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x56,
	0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x12, 0x10, 0x2f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x73, 0x5a, 0x15, 0x22, 0x10, 0x2f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0xb3, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x73, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x56, 0x31, 0x12, 0x29,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x73, 0x55, 0x6e, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x73, 0x55, 0x6e, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3b, 0x12, 0x19, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61,
	0x72, 0x73, 0x2f, 0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x5a, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x73, 0x2f,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0xcb, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x44, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x2f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x44, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x44, 0x61, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x49, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x43, 0x12, 0x1d, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x2f, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2d, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x5a, 0x22, 0x22, 0x1d, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x2d, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xbe, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x22, 0x6a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x64, 0x12, 0x15, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x2d, 0x76, 0x61, 0x72, 0x2d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5a, 0x1a, 0x22, 0x15, 0x2f, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x2d, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x3a, 0x01, 0x2a, 0x5a, 0x15, 0x12, 0x13, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d,
	0x76, 0x61, 0x72, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x61, 0x6c, 0x6c, 0x5a, 0x18, 0x22,
	0x13, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x61, 0x6c, 0x6c, 0x3a, 0x01, 0x2a, 0x12, 0x8c, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x27, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x4e, 0x6f, 0x64, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12,
	0x0f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5a, 0x14, 0x22, 0x0f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x56, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x56,
	0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x0e, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x2d, 0x76, 0x61, 0x72, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x5a, 0x13, 0x22, 0x0e, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x2f, 0x70, 0x61, 0x74, 0x68, 0x3a, 0x01, 0x2a, 0x12,
	0x87, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61,
	0x72, 0x12, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29,
	0x12, 0x10, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5a, 0x15, 0x22, 0x10, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x12, 0x95, 0x01, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x25, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x53,
	0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x11, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x76, 0x61,
	0x72, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x5a, 0x16, 0x22, 0x11, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x2d, 0x76, 0x61, 0x72, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x3a, 0x01,
	0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_mixer_proto_goTypes = []interface{}{
	(*QueryRequest)(nil),                        // 0: datacommons.QueryRequest
	(*GetPropertyLabelsRequest)(nil),            // 1: datacommons.GetPropertyLabelsRequest
	(*GetPropertyValuesRequest)(nil),            // 2: datacommons.GetPropertyValuesRequest
	(*GetTriplesRequest)(nil),                   // 3: datacommons.GetTriplesRequest
	(*GetPlacesInRequest)(nil),                  // 4: datacommons.GetPlacesInRequest
	(*GetPlaceObsRequest)(nil),                  // 5: datacommons.GetPlaceObsRequest
	(*GetStatsRequest)(nil),                     // 6: datacommons.GetStatsRequest
	(*GetStatSetSeriesRequest)(nil),             // 7: datacommons.GetStatSetSeriesRequest
	(*GetStatValueRequest)(nil),                 // 8: datacommons.GetStatValueRequest
	(*GetStatSeriesRequest)(nil),                // 9: datacommons.GetStatSeriesRequest
	(*GetStatAllRequest)(nil),                   // 10: datacommons.GetStatAllRequest
	(*GetStatSetWithinPlaceRequest)(nil),        // 11: datacommons.GetStatSetWithinPlaceRequest
	(*GetStatSetRequest)(nil),                   // 12: datacommons.GetStatSetRequest
	(*GetLocationsRankingsRequest)(nil),         // 13: datacommons.GetLocationsRankingsRequest
	(*GetRelatedLocationsRequest)(nil),          // 14: datacommons.GetRelatedLocationsRequest
	(*GetPlacePageDataRequest)(nil),             // 15: datacommons.GetPlacePageDataRequest
	(*GetProteinPageDataRequest)(nil),           // 16: datacommons.GetProteinPageDataRequest
	(*TranslateRequest)(nil),                    // 17: datacommons.TranslateRequest
	(*SearchRequest)(nil),                       // 18: datacommons.SearchRequest
	(*GetVersionRequest)(nil),                   // 19: datacommons.GetVersionRequest
	(*GetPlaceStatsVarRequest)(nil),             // 20: datacommons.GetPlaceStatsVarRequest
	(*GetPlaceStatVarsRequest)(nil),             // 21: datacommons.GetPlaceStatVarsRequest
	(*GetPlaceStatVarsUnionRequest)(nil),        // 22: datacommons.GetPlaceStatVarsUnionRequest
	(*GetPlaceStatDateWithinPlaceRequest)(nil),  // 23: datacommons.GetPlaceStatDateWithinPlaceRequest
	(*GetStatVarGroupRequest)(nil),              // 24: datacommons.GetStatVarGroupRequest
	(*GetStatVarGroupNodeRequest)(nil),          // 25: datacommons.GetStatVarGroupNodeRequest
	(*GetStatVarPathRequest)(nil),               // 26: datacommons.GetStatVarPathRequest
	(*SearchStatVarRequest)(nil),                // 27: datacommons.SearchStatVarRequest
	(*GetStatVarSummaryRequest)(nil),            // 28: datacommons.GetStatVarSummaryRequest
	(*QueryResponse)(nil),                       // 29: datacommons.QueryResponse
	(*GetPropertyLabelsResponse)(nil),           // 30: datacommons.GetPropertyLabelsResponse
	(*GetPropertyValuesResponse)(nil),           // 31: datacommons.GetPropertyValuesResponse
	(*GetTriplesResponse)(nil),                  // 32: datacommons.GetTriplesResponse
	(*GetPlacesInResponse)(nil),                 // 33: datacommons.GetPlacesInResponse
	(*SVOCollection)(nil),                       // 34: datacommons.SVOCollection
	(*GetStatsResponse)(nil),                    // 35: datacommons.GetStatsResponse
	(*GetStatSetSeriesResponse)(nil),            // 36: datacommons.GetStatSetSeriesResponse
	(*GetStatValueResponse)(nil),                // 37: datacommons.GetStatValueResponse
	(*GetStatSeriesResponse)(nil),               // 38: datacommons.GetStatSeriesResponse
	(*GetStatAllResponse)(nil),                  // 39: datacommons.GetStatAllResponse
	(*GetStatSetResponse)(nil),                  // 40: datacommons.GetStatSetResponse
	(*GetLocationsRankingsResponse)(nil),        // 41: datacommons.GetLocationsRankingsResponse
	(*GetRelatedLocationsResponse)(nil),         // 42: datacommons.GetRelatedLocationsResponse
	(*GetPlacePageDataResponse)(nil),            // 43: datacommons.GetPlacePageDataResponse
	(*GraphNode)(nil),                           // 44: datacommons.GraphNode
	(*TranslateResponse)(nil),                   // 45: datacommons.TranslateResponse
	(*SearchResponse)(nil),                      // 46: datacommons.SearchResponse
	(*GetVersionResponse)(nil),                  // 47: datacommons.GetVersionResponse
	(*GetPlaceStatsVarResponse)(nil),            // 48: datacommons.GetPlaceStatsVarResponse
	(*GetPlaceStatVarsResponse)(nil),            // 49: datacommons.GetPlaceStatVarsResponse
	(*GetPlaceStatVarsUnionResponse)(nil),       // 50: datacommons.GetPlaceStatVarsUnionResponse
	(*GetPlaceStatDateWithinPlaceResponse)(nil), // 51: datacommons.GetPlaceStatDateWithinPlaceResponse
	(*StatVarGroups)(nil),                       // 52: datacommons.StatVarGroups
	(*StatVarGroupNode)(nil),                    // 53: datacommons.StatVarGroupNode
	(*GetStatVarPathResponse)(nil),              // 54: datacommons.GetStatVarPathResponse
	(*SearchStatVarResponse)(nil),               // 55: datacommons.SearchStatVarResponse
	(*GetStatVarSummaryResponse)(nil),           // 56: datacommons.GetStatVarSummaryResponse
}
var file_mixer_proto_depIdxs = []int32{
	0,  // 0: datacommons.Mixer.Query:input_type -> datacommons.QueryRequest
	1,  // 1: datacommons.Mixer.GetPropertyLabels:input_type -> datacommons.GetPropertyLabelsRequest
	2,  // 2: datacommons.Mixer.GetPropertyValues:input_type -> datacommons.GetPropertyValuesRequest
	3,  // 3: datacommons.Mixer.GetTriples:input_type -> datacommons.GetTriplesRequest
	4,  // 4: datacommons.Mixer.GetPlacesIn:input_type -> datacommons.GetPlacesInRequest
	5,  // 5: datacommons.Mixer.GetPlaceObs:input_type -> datacommons.GetPlaceObsRequest
	6,  // 6: datacommons.Mixer.GetStats:input_type -> datacommons.GetStatsRequest
	7,  // 7: datacommons.Mixer.GetStatSetSeries:input_type -> datacommons.GetStatSetSeriesRequest
	8,  // 8: datacommons.Mixer.GetStatValue:input_type -> datacommons.GetStatValueRequest
	9,  // 9: datacommons.Mixer.GetStatSeries:input_type -> datacommons.GetStatSeriesRequest
	10, // 10: datacommons.Mixer.GetStatAll:input_type -> datacommons.GetStatAllRequest
	11, // 11: datacommons.Mixer.GetStatSetWithinPlace:input_type -> datacommons.GetStatSetWithinPlaceRequest
	12, // 12: datacommons.Mixer.GetStatSet:input_type -> datacommons.GetStatSetRequest
	13, // 13: datacommons.Mixer.GetLocationsRankings:input_type -> datacommons.GetLocationsRankingsRequest
	14, // 14: datacommons.Mixer.GetRelatedLocations:input_type -> datacommons.GetRelatedLocationsRequest
	15, // 15: datacommons.Mixer.GetPlacePageData:input_type -> datacommons.GetPlacePageDataRequest
	16, // 16: datacommons.Mixer.GetProteinPageData:input_type -> datacommons.GetProteinPageDataRequest
	17, // 17: datacommons.Mixer.Translate:input_type -> datacommons.TranslateRequest
	18, // 18: datacommons.Mixer.Search:input_type -> datacommons.SearchRequest
	19, // 19: datacommons.Mixer.GetVersion:input_type -> datacommons.GetVersionRequest
	20, // 20: datacommons.Mixer.GetPlaceStatsVar:input_type -> datacommons.GetPlaceStatsVarRequest
	21, // 21: datacommons.Mixer.GetPlaceStatVars:input_type -> datacommons.GetPlaceStatVarsRequest
	22, // 22: datacommons.Mixer.GetPlaceStatVarsUnionV1:input_type -> datacommons.GetPlaceStatVarsUnionRequest
	23, // 23: datacommons.Mixer.GetPlaceStatDateWithinPlace:input_type -> datacommons.GetPlaceStatDateWithinPlaceRequest
	24, // 24: datacommons.Mixer.GetStatVarGroup:input_type -> datacommons.GetStatVarGroupRequest
	25, // 25: datacommons.Mixer.GetStatVarGroupNode:input_type -> datacommons.GetStatVarGroupNodeRequest
	26, // 26: datacommons.Mixer.GetStatVarPath:input_type -> datacommons.GetStatVarPathRequest
	27, // 27: datacommons.Mixer.SearchStatVar:input_type -> datacommons.SearchStatVarRequest
	28, // 28: datacommons.Mixer.GetStatVarSummary:input_type -> datacommons.GetStatVarSummaryRequest
	29, // 29: datacommons.Mixer.Query:output_type -> datacommons.QueryResponse
	30, // 30: datacommons.Mixer.GetPropertyLabels:output_type -> datacommons.GetPropertyLabelsResponse
	31, // 31: datacommons.Mixer.GetPropertyValues:output_type -> datacommons.GetPropertyValuesResponse
	32, // 32: datacommons.Mixer.GetTriples:output_type -> datacommons.GetTriplesResponse
	33, // 33: datacommons.Mixer.GetPlacesIn:output_type -> datacommons.GetPlacesInResponse
	34, // 34: datacommons.Mixer.GetPlaceObs:output_type -> datacommons.SVOCollection
	35, // 35: datacommons.Mixer.GetStats:output_type -> datacommons.GetStatsResponse
	36, // 36: datacommons.Mixer.GetStatSetSeries:output_type -> datacommons.GetStatSetSeriesResponse
	37, // 37: datacommons.Mixer.GetStatValue:output_type -> datacommons.GetStatValueResponse
	38, // 38: datacommons.Mixer.GetStatSeries:output_type -> datacommons.GetStatSeriesResponse
	39, // 39: datacommons.Mixer.GetStatAll:output_type -> datacommons.GetStatAllResponse
	40, // 40: datacommons.Mixer.GetStatSetWithinPlace:output_type -> datacommons.GetStatSetResponse
	40, // 41: datacommons.Mixer.GetStatSet:output_type -> datacommons.GetStatSetResponse
	41, // 42: datacommons.Mixer.GetLocationsRankings:output_type -> datacommons.GetLocationsRankingsResponse
	42, // 43: datacommons.Mixer.GetRelatedLocations:output_type -> datacommons.GetRelatedLocationsResponse
	43, // 44: datacommons.Mixer.GetPlacePageData:output_type -> datacommons.GetPlacePageDataResponse
	44, // 45: datacommons.Mixer.GetProteinPageData:output_type -> datacommons.GraphNode
	45, // 46: datacommons.Mixer.Translate:output_type -> datacommons.TranslateResponse
	46, // 47: datacommons.Mixer.Search:output_type -> datacommons.SearchResponse
	47, // 48: datacommons.Mixer.GetVersion:output_type -> datacommons.GetVersionResponse
	48, // 49: datacommons.Mixer.GetPlaceStatsVar:output_type -> datacommons.GetPlaceStatsVarResponse
	49, // 50: datacommons.Mixer.GetPlaceStatVars:output_type -> datacommons.GetPlaceStatVarsResponse
	50, // 51: datacommons.Mixer.GetPlaceStatVarsUnionV1:output_type -> datacommons.GetPlaceStatVarsUnionResponse
	51, // 52: datacommons.Mixer.GetPlaceStatDateWithinPlace:output_type -> datacommons.GetPlaceStatDateWithinPlaceResponse
	52, // 53: datacommons.Mixer.GetStatVarGroup:output_type -> datacommons.StatVarGroups
	53, // 54: datacommons.Mixer.GetStatVarGroupNode:output_type -> datacommons.StatVarGroupNode
	54, // 55: datacommons.Mixer.GetStatVarPath:output_type -> datacommons.GetStatVarPathResponse
	55, // 56: datacommons.Mixer.SearchStatVar:output_type -> datacommons.SearchStatVarResponse
	56, // 57: datacommons.Mixer.GetStatVarSummary:output_type -> datacommons.GetStatVarSummaryResponse
	29, // [29:58] is the sub-list for method output_type
	0,  // [0:29] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mixer_proto_init() }
func file_mixer_proto_init() {
	if File_mixer_proto != nil {
		return
	}
	file_common_proto_init()
	file_internal_proto_init()
	file_misc_proto_init()
	file_node_proto_init()
	file_place_proto_init()
	file_query_proto_init()
	file_stat_proto_init()
	file_stat_var_proto_init()
	file_translate_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mixer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mixer_proto_goTypes,
		DependencyIndexes: file_mixer_proto_depIdxs,
	}.Build()
	File_mixer_proto = out.File
	file_mixer_proto_rawDesc = nil
	file_mixer_proto_goTypes = nil
	file_mixer_proto_depIdxs = nil
}
