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
// source: v2/observation.proto

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

// Represent a list of entities passed in as plain list or graph expression.
type DcidOrExpression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dcids []string `protobuf:"bytes,1,rep,name=dcids,proto3" json:"dcids,omitempty"`
	// An arrow notation expression
	// Ex: country/USA<-containedInPlace{typeOf: State}
	Expression string `protobuf:"bytes,2,opt,name=expression,proto3" json:"expression,omitempty"`
	// Supported operators for the formula expression: +, -, *, /, along with parentheses.
	// The items of the formula have a form of "StatVar[OptionalFilters]".
	//
	// The following filters are supported:
	// 1. mm: MeasurementMethod
	// 2. op: ObservationPeriod
	// 3. ut: Unit
	// 4. sf: ScalingFactor
	//
	// For example:
	// 1. Person_Count - Person_Count_Female
	// 2. Person_Count_Female / Person_Count[mm=CensusACS5yrSurvey;op=P5Y]
	// 3. Person_Count - Person_Count_Female - Person_Count_Male[op=P5Y]
	// 4. (Person_Count_Male - Person_Count_Female) / Person_Count
	Formula string `protobuf:"bytes,3,opt,name=formula,proto3" json:"formula,omitempty"`
}

func (x *DcidOrExpression) Reset() {
	*x = DcidOrExpression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DcidOrExpression) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DcidOrExpression) ProtoMessage() {}

func (x *DcidOrExpression) ProtoReflect() protoreflect.Message {
	mi := &file_v2_observation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DcidOrExpression.ProtoReflect.Descriptor instead.
func (*DcidOrExpression) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{0}
}

func (x *DcidOrExpression) GetDcids() []string {
	if x != nil {
		return x.Dcids
	}
	return nil
}

func (x *DcidOrExpression) GetExpression() string {
	if x != nil {
		return x.Expression
	}
	return ""
}

func (x *DcidOrExpression) GetFormula() string {
	if x != nil {
		return x.Formula
	}
	return ""
}

// Holds all observations of a particular variable.
type VariableObservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keyed by entity DCID
	ByEntity map[string]*EntityObservation `protobuf:"bytes,1,rep,name=by_entity,json=byEntity,proto3" json:"by_entity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VariableObservation) Reset() {
	*x = VariableObservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariableObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariableObservation) ProtoMessage() {}

func (x *VariableObservation) ProtoReflect() protoreflect.Message {
	mi := &file_v2_observation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariableObservation.ProtoReflect.Descriptor instead.
func (*VariableObservation) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{1}
}

func (x *VariableObservation) GetByEntity() map[string]*EntityObservation {
	if x != nil {
		return x.ByEntity
	}
	return nil
}

type EntityObservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Facet are orderred by preference.
	OrderedFacets []*FacetObservation `protobuf:"bytes,1,rep,name=ordered_facets,json=orderedFacets,proto3" json:"ordered_facets,omitempty"`
}

func (x *EntityObservation) Reset() {
	*x = EntityObservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityObservation) ProtoMessage() {}

func (x *EntityObservation) ProtoReflect() protoreflect.Message {
	mi := &file_v2_observation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityObservation.ProtoReflect.Descriptor instead.
func (*EntityObservation) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{2}
}

func (x *EntityObservation) GetOrderedFacets() []*FacetObservation {
	if x != nil {
		return x.OrderedFacets
	}
	return nil
}

type FacetObservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FacetId string `protobuf:"bytes,1,opt,name=facet_id,json=facetId,proto3" json:"facet_id,omitempty"`
	// Observations are sorted by date
	Observations []*proto.PointStat `protobuf:"bytes,2,rep,name=observations,proto3" json:"observations,omitempty"`
}

func (x *FacetObservation) Reset() {
	*x = FacetObservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FacetObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FacetObservation) ProtoMessage() {}

func (x *FacetObservation) ProtoReflect() protoreflect.Message {
	mi := &file_v2_observation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FacetObservation.ProtoReflect.Descriptor instead.
func (*FacetObservation) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{3}
}

func (x *FacetObservation) GetFacetId() string {
	if x != nil {
		return x.FacetId
	}
	return ""
}

func (x *FacetObservation) GetObservations() []*proto.PointStat {
	if x != nil {
		return x.Observations
	}
	return nil
}

type FacetFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When specified, only observations with provenance in these domains are
	// returned.
	Domains []string `protobuf:"bytes,2,rep,name=domains,proto3" json:"domains,omitempty"`
	// When specified, only observations with this facet id are returned
	FacetId string `protobuf:"bytes,3,opt,name=facet_id,json=facetId,proto3" json:"facet_id,omitempty"`
}

func (x *FacetFilter) Reset() {
	*x = FacetFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FacetFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FacetFilter) ProtoMessage() {}

func (x *FacetFilter) ProtoReflect() protoreflect.Message {
	mi := &file_v2_observation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FacetFilter.ProtoReflect.Descriptor instead.
func (*FacetFilter) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{4}
}

func (x *FacetFilter) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *FacetFilter) GetFacetId() string {
	if x != nil {
		return x.FacetId
	}
	return ""
}

// Generic observation request
type ObservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of statistical variable DCIDs, or a formula expression.
	Variable *DcidOrExpression `protobuf:"bytes,1,opt,name=variable,proto3" json:"variable,omitempty"`
	// A list of entity DCIDs, or an arrow expression that covers a list of entity
	// DCIDs.
	Entity *DcidOrExpression `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	// Date of the observation
	//   - Not specified: all observations are returned
	//   - "LATEST": latest obseration of each facet is returned
	//   - "<DATE>": a speficied valid ISO 8601 date. Observation corresponding to
	//     this date is returned.
	Date string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	// Value of the observation
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	// [Optional] filter returned observations by facet
	Filter *FacetFilter `protobuf:"bytes,5,opt,name=filter,proto3" json:"filter,omitempty"`
	// Fields to return, valid values are: "variable", "entity", "date", "value", "facet"
	Select []string `protobuf:"bytes,6,rep,name=select,proto3" json:"select,omitempty"`
}

func (x *ObservationRequest) Reset() {
	*x = ObservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationRequest) ProtoMessage() {}

func (x *ObservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v2_observation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationRequest.ProtoReflect.Descriptor instead.
func (*ObservationRequest) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{5}
}

func (x *ObservationRequest) GetVariable() *DcidOrExpression {
	if x != nil {
		return x.Variable
	}
	return nil
}

func (x *ObservationRequest) GetEntity() *DcidOrExpression {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *ObservationRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ObservationRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ObservationRequest) GetFilter() *FacetFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ObservationRequest) GetSelect() []string {
	if x != nil {
		return x.Select
	}
	return nil
}

type ObservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keyed by variable DCID
	ByVariable map[string]*VariableObservation `protobuf:"bytes,1,rep,name=by_variable,json=byVariable,proto3" json:"by_variable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Keyed by facet ID
	Facets map[string]*proto.Facet `protobuf:"bytes,2,rep,name=facets,proto3" json:"facets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ObservationResponse) Reset() {
	*x = ObservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationResponse) ProtoMessage() {}

func (x *ObservationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v2_observation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationResponse.ProtoReflect.Descriptor instead.
func (*ObservationResponse) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{6}
}

func (x *ObservationResponse) GetByVariable() map[string]*VariableObservation {
	if x != nil {
		return x.ByVariable
	}
	return nil
}

func (x *ObservationResponse) GetFacets() map[string]*proto.Facet {
	if x != nil {
		return x.Facets
	}
	return nil
}

var File_v2_observation_proto protoreflect.FileDescriptor

var file_v2_observation_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x32, 0x2f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x62, 0x0a, 0x10, 0x44, 0x63, 0x69, 0x64, 0x4f, 0x72, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x63, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x64, 0x63, 0x69, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x22, 0xc5, 0x01, 0x0a, 0x13, 0x56, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e,
	0x0a, 0x09, 0x62, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x62, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x5e,
	0x0a, 0x0d, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x37, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5c,
	0x0a, 0x11, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x66,
	0x61, 0x63, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x61, 0x63,
	0x65, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x46, 0x61, 0x63, 0x65, 0x74, 0x73, 0x22, 0x69, 0x0a, 0x10,
	0x46, 0x61, 0x63, 0x65, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x74, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0c, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x42, 0x0a, 0x0b, 0x46, 0x61, 0x63, 0x65, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x74, 0x49, 0x64, 0x22, 0x83, 0x02, 0x0a, 0x12,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x63, 0x69, 0x64, 0x4f, 0x72, 0x45, 0x78, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x38, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x44, 0x63, 0x69, 0x64, 0x4f, 0x72, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x22, 0xe7, 0x02, 0x0a, 0x13, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x62, 0x79, 0x5f,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x42, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x62, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x47, 0x0a, 0x06, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32,
	0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x1a, 0x62, 0x0a, 0x0f, 0x42, 0x79, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a, 0x0b,
	0x46, 0x61, 0x63, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x78, 0x65, 0x72, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v2_observation_proto_rawDescOnce sync.Once
	file_v2_observation_proto_rawDescData = file_v2_observation_proto_rawDesc
)

func file_v2_observation_proto_rawDescGZIP() []byte {
	file_v2_observation_proto_rawDescOnce.Do(func() {
		file_v2_observation_proto_rawDescData = protoimpl.X.CompressGZIP(file_v2_observation_proto_rawDescData)
	})
	return file_v2_observation_proto_rawDescData
}

var file_v2_observation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_v2_observation_proto_goTypes = []interface{}{
	(*DcidOrExpression)(nil),    // 0: datacommons.v2.DcidOrExpression
	(*VariableObservation)(nil), // 1: datacommons.v2.VariableObservation
	(*EntityObservation)(nil),   // 2: datacommons.v2.EntityObservation
	(*FacetObservation)(nil),    // 3: datacommons.v2.FacetObservation
	(*FacetFilter)(nil),         // 4: datacommons.v2.FacetFilter
	(*ObservationRequest)(nil),  // 5: datacommons.v2.ObservationRequest
	(*ObservationResponse)(nil), // 6: datacommons.v2.ObservationResponse
	nil,                         // 7: datacommons.v2.VariableObservation.ByEntityEntry
	nil,                         // 8: datacommons.v2.ObservationResponse.ByVariableEntry
	nil,                         // 9: datacommons.v2.ObservationResponse.FacetsEntry
	(*proto.PointStat)(nil),     // 10: datacommons.PointStat
	(*proto.Facet)(nil),         // 11: datacommons.Facet
}
var file_v2_observation_proto_depIdxs = []int32{
	7,  // 0: datacommons.v2.VariableObservation.by_entity:type_name -> datacommons.v2.VariableObservation.ByEntityEntry
	3,  // 1: datacommons.v2.EntityObservation.ordered_facets:type_name -> datacommons.v2.FacetObservation
	10, // 2: datacommons.v2.FacetObservation.observations:type_name -> datacommons.PointStat
	0,  // 3: datacommons.v2.ObservationRequest.variable:type_name -> datacommons.v2.DcidOrExpression
	0,  // 4: datacommons.v2.ObservationRequest.entity:type_name -> datacommons.v2.DcidOrExpression
	4,  // 5: datacommons.v2.ObservationRequest.filter:type_name -> datacommons.v2.FacetFilter
	8,  // 6: datacommons.v2.ObservationResponse.by_variable:type_name -> datacommons.v2.ObservationResponse.ByVariableEntry
	9,  // 7: datacommons.v2.ObservationResponse.facets:type_name -> datacommons.v2.ObservationResponse.FacetsEntry
	2,  // 8: datacommons.v2.VariableObservation.ByEntityEntry.value:type_name -> datacommons.v2.EntityObservation
	1,  // 9: datacommons.v2.ObservationResponse.ByVariableEntry.value:type_name -> datacommons.v2.VariableObservation
	11, // 10: datacommons.v2.ObservationResponse.FacetsEntry.value:type_name -> datacommons.Facet
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_v2_observation_proto_init() }
func file_v2_observation_proto_init() {
	if File_v2_observation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_observation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DcidOrExpression); i {
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
		file_v2_observation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariableObservation); i {
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
		file_v2_observation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityObservation); i {
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
		file_v2_observation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FacetObservation); i {
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
		file_v2_observation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FacetFilter); i {
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
		file_v2_observation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationRequest); i {
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
		file_v2_observation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObservationResponse); i {
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
			RawDescriptor: file_v2_observation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v2_observation_proto_goTypes,
		DependencyIndexes: file_v2_observation_proto_depIdxs,
		MessageInfos:      file_v2_observation_proto_msgTypes,
	}.Build()
	File_v2_observation_proto = out.File
	file_v2_observation_proto_rawDesc = nil
	file_v2_observation_proto_goTypes = nil
	file_v2_observation_proto_depIdxs = nil
}
