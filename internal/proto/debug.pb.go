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

//  **** IMPORTANT NOTE ****
//
//  The proto of BT data has to match exactly the g3 proto, including tag
//  number.

// *** IMPORTANT NOTE ***
// Keep this in sync with the same proto in datacommonsorg/import.
// Eventually we need to move the shared protos out to a separate repo.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: debug.proto

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

// Severity level of the message.
type Log_Level int32

const (
	Log_LEVEL_UNSPECIFIED Log_Level = 0
	Log_LEVEL_INFO        Log_Level = 1
	Log_LEVEL_WARNING     Log_Level = 2
	Log_LEVEL_ERROR       Log_Level = 3
	Log_LEVEL_FATAL       Log_Level = 4
)

// Enum value maps for Log_Level.
var (
	Log_Level_name = map[int32]string{
		0: "LEVEL_UNSPECIFIED",
		1: "LEVEL_INFO",
		2: "LEVEL_WARNING",
		3: "LEVEL_ERROR",
		4: "LEVEL_FATAL",
	}
	Log_Level_value = map[string]int32{
		"LEVEL_UNSPECIFIED": 0,
		"LEVEL_INFO":        1,
		"LEVEL_WARNING":     2,
		"LEVEL_ERROR":       3,
		"LEVEL_FATAL":       4,
	}
)

func (x Log_Level) Enum() *Log_Level {
	p := new(Log_Level)
	*p = x
	return p
}

func (x Log_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Log_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_debug_proto_enumTypes[0].Descriptor()
}

func (Log_Level) Type() protoreflect.EnumType {
	return &file_debug_proto_enumTypes[0]
}

func (x Log_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Log_Level) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Log_Level(num)
	return nil
}

// Deprecated: Use Log_Level.Descriptor instead.
func (Log_Level) EnumDescriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{0, 0}
}

//
// A log of import processing with details on any warnings, errors, etc.
//
type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries    []*Log_Entry    `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
	CounterSet *Log_CounterSet `protobuf:"bytes,2,opt,name=counter_set,json=counterSet" json:"counter_set,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetEntries() []*Log_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *Log) GetCounterSet() *Log_CounterSet {
	if x != nil {
		return x.CounterSet
	}
	return nil
}

type Log_Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the file. Could be a full path or not. Typically, this is
	// an MCF file or CSV.
	File *string `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	// Line number in file, starting from 1.  Matches the row number in
	// CSV (since we do not handle newlines in values).
	LineNumber *int64 `protobuf:"varint,2,opt,name=line_number,json=lineNumber" json:"line_number,omitempty"`
	// Column number in line, starting from 1. Can be empty.
	ColumnNumber *int64 `protobuf:"varint,3,opt,name=column_number,json=columnNumber" json:"column_number,omitempty"`
	// Column name set when the input is a CSV. Can be empty.
	ColumnName *string `protobuf:"bytes,4,opt,name=column_name,json=columnName" json:"column_name,omitempty"`
}

func (x *Log_Location) Reset() {
	*x = Log_Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log_Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log_Location) ProtoMessage() {}

func (x *Log_Location) ProtoReflect() protoreflect.Message {
	mi := &file_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log_Location.ProtoReflect.Descriptor instead.
func (*Log_Location) Descriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Log_Location) GetFile() string {
	if x != nil && x.File != nil {
		return *x.File
	}
	return ""
}

func (x *Log_Location) GetLineNumber() int64 {
	if x != nil && x.LineNumber != nil {
		return *x.LineNumber
	}
	return 0
}

func (x *Log_Location) GetColumnNumber() int64 {
	if x != nil && x.ColumnNumber != nil {
		return *x.ColumnNumber
	}
	return 0
}

func (x *Log_Location) GetColumnName() string {
	if x != nil && x.ColumnName != nil {
		return *x.ColumnName
	}
	return ""
}

type Log_CounterSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is the name of a counter.
	Counters map[string]int64 `protobuf:"bytes,1,rep,name=counters" json:"counters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (x *Log_CounterSet) Reset() {
	*x = Log_CounterSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log_CounterSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log_CounterSet) ProtoMessage() {}

func (x *Log_CounterSet) ProtoReflect() protoreflect.Message {
	mi := &file_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log_CounterSet.ProtoReflect.Descriptor instead.
func (*Log_CounterSet) Descriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Log_CounterSet) GetCounters() map[string]int64 {
	if x != nil {
		return x.Counters
	}
	return nil
}

// One log entry. This could be a sample of messages for a particular
// counter.
type Log_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level    *Log_Level    `protobuf:"varint,1,opt,name=level,enum=datacommons.Log_Level" json:"level,omitempty"`
	Location *Log_Location `protobuf:"bytes,2,opt,name=location" json:"location,omitempty"`
	// This must be user understandable.
	UserMessage *string `protobuf:"bytes,3,opt,name=user_message,json=userMessage" json:"user_message,omitempty"`
	// A counter key in CounterSet.
	CounterKey *string `protobuf:"bytes,4,opt,name=counter_key,json=counterKey" json:"counter_key,omitempty"`
}

func (x *Log_Entry) Reset() {
	*x = Log_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log_Entry) ProtoMessage() {}

func (x *Log_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_debug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log_Entry.ProtoReflect.Descriptor instead.
func (*Log_Entry) Descriptor() ([]byte, []int) {
	return file_debug_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Log_Entry) GetLevel() Log_Level {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return Log_LEVEL_UNSPECIFIED
}

func (x *Log_Entry) GetLocation() *Log_Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Log_Entry) GetUserMessage() string {
	if x != nil && x.UserMessage != nil {
		return *x.UserMessage
	}
	return ""
}

func (x *Log_Entry) GetCounterKey() string {
	if x != nil && x.CounterKey != nil {
		return *x.CounterKey
	}
	return ""
}

var File_debug_proto protoreflect.FileDescriptor

var file_debug_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x22, 0xa8, 0x05, 0x0a, 0x03, 0x4c,
	0x6f, 0x67, 0x12, 0x30, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x74, 0x1a, 0x85, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x90, 0x01, 0x0a, 0x0a, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x08, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x1a, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xb0, 0x01,
	0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x4b, 0x65, 0x79,
	0x22, 0x63, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x45, 0x56,
	0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x46, 0x41,
	0x54, 0x41, 0x4c, 0x10, 0x04, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_debug_proto_rawDescOnce sync.Once
	file_debug_proto_rawDescData = file_debug_proto_rawDesc
)

func file_debug_proto_rawDescGZIP() []byte {
	file_debug_proto_rawDescOnce.Do(func() {
		file_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_debug_proto_rawDescData)
	})
	return file_debug_proto_rawDescData
}

var file_debug_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_debug_proto_goTypes = []interface{}{
	(Log_Level)(0),         // 0: datacommons.Log.Level
	(*Log)(nil),            // 1: datacommons.Log
	(*Log_Location)(nil),   // 2: datacommons.Log.Location
	(*Log_CounterSet)(nil), // 3: datacommons.Log.CounterSet
	(*Log_Entry)(nil),      // 4: datacommons.Log.Entry
	nil,                    // 5: datacommons.Log.CounterSet.CountersEntry
}
var file_debug_proto_depIdxs = []int32{
	4, // 0: datacommons.Log.entries:type_name -> datacommons.Log.Entry
	3, // 1: datacommons.Log.counter_set:type_name -> datacommons.Log.CounterSet
	5, // 2: datacommons.Log.CounterSet.counters:type_name -> datacommons.Log.CounterSet.CountersEntry
	0, // 3: datacommons.Log.Entry.level:type_name -> datacommons.Log.Level
	2, // 4: datacommons.Log.Entry.location:type_name -> datacommons.Log.Location
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_debug_proto_init() }
func file_debug_proto_init() {
	if File_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log_Location); i {
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
		file_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log_CounterSet); i {
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
		file_debug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log_Entry); i {
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
			RawDescriptor: file_debug_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_debug_proto_goTypes,
		DependencyIndexes: file_debug_proto_depIdxs,
		EnumInfos:         file_debug_proto_enumTypes,
		MessageInfos:      file_debug_proto_msgTypes,
	}.Build()
	File_debug_proto = out.File
	file_debug_proto_rawDesc = nil
	file_debug_proto_goTypes = nil
	file_debug_proto_depIdxs = nil
}
