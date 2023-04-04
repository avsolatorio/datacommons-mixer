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
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
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

// Holds all observations of a particular variable.
type VariableObservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keyed by entity DCID
	ObservationsByEntity map[string]*EntityObservation `protobuf:"bytes,1,rep,name=observations_by_entity,json=observationsByEntity,proto3" json:"observations_by_entity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *VariableObservation) Reset() {
	*x = VariableObservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariableObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariableObservation) ProtoMessage() {}

func (x *VariableObservation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use VariableObservation.ProtoReflect.Descriptor instead.
func (*VariableObservation) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{0}
}

func (x *VariableObservation) GetObservationsByEntity() map[string]*EntityObservation {
	if x != nil {
		return x.ObservationsByEntity
	}
	return nil
}

type EntityObservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Facet are orderred by preference.
	OrderedFacetObservations []*FacetObservation `protobuf:"bytes,1,rep,name=ordered_facet_observations,json=orderedFacetObservations,proto3" json:"ordered_facet_observations,omitempty"`
}

func (x *EntityObservation) Reset() {
	*x = EntityObservation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityObservation) ProtoMessage() {}

func (x *EntityObservation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EntityObservation.ProtoReflect.Descriptor instead.
func (*EntityObservation) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{1}
}

func (x *EntityObservation) GetOrderedFacetObservations() []*FacetObservation {
	if x != nil {
		return x.OrderedFacetObservations
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
		mi := &file_v2_observation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FacetObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FacetObservation) ProtoMessage() {}

func (x *FacetObservation) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FacetObservation.ProtoReflect.Descriptor instead.
func (*FacetObservation) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{2}
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

// Generic observation request
type ObservationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of entity DCIDs
	Entities []string `protobuf:"bytes,1,rep,name=entities,proto3" json:"entities,omitempty"`
	// An arrow notation expression of entities
	// Ex: country/USA<-containedInPlace{typeOf: State}
	EntitiesExpression string `protobuf:"bytes,2,opt,name=entities_expression,json=entitiesExpression,proto3" json:"entities_expression,omitempty"`
	// A list of statistical variable DCIDs
	Variables []string `protobuf:"bytes,3,rep,name=variables,proto3" json:"variables,omitempty"`
	// An arrow notation expression of variables
	// Ex: dc/g/Root<-memberOf
	VariablesExpression string `protobuf:"bytes,4,opt,name=variables_expression,json=variablesExpression,proto3" json:"variables_expression,omitempty"`
	// Date of the observation
	//   - Not specified: all observations are returned
	//   - "LATEST": latest obseration of each facet is returned
	//   - "<DATE>": a speficied valid ISO 8601 date. Observation corresponding to
	//     this date is returned.
	Date string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
	// When specified, only observation with this unit is returned.
	Unit string `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	// When specified, only observation with this measurement method is returned.
	MeasurementMethod string `protobuf:"bytes,7,opt,name=measurement_method,json=measurementMethod,proto3" json:"measurement_method,omitempty"`
	// When specified, only observation with this observation period is returned.
	Period string `protobuf:"bytes,8,opt,name=period,proto3" json:"period,omitempty"`
}

func (x *ObservationRequest) Reset() {
	*x = ObservationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationRequest) ProtoMessage() {}

func (x *ObservationRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ObservationRequest.ProtoReflect.Descriptor instead.
func (*ObservationRequest) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{3}
}

func (x *ObservationRequest) GetEntities() []string {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *ObservationRequest) GetEntitiesExpression() string {
	if x != nil {
		return x.EntitiesExpression
	}
	return ""
}

func (x *ObservationRequest) GetVariables() []string {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *ObservationRequest) GetVariablesExpression() string {
	if x != nil {
		return x.VariablesExpression
	}
	return ""
}

func (x *ObservationRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *ObservationRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *ObservationRequest) GetMeasurementMethod() string {
	if x != nil {
		return x.MeasurementMethod
	}
	return ""
}

func (x *ObservationRequest) GetPeriod() string {
	if x != nil {
		return x.Period
	}
	return ""
}

type ObservationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keyed by variable DCID
	ObservationsByVariable map[string]*VariableObservation `protobuf:"bytes,1,rep,name=observations_by_variable,json=observationsByVariable,proto3" json:"observations_by_variable,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Keyed by facet ID
	Facets map[string]*proto.Facet `protobuf:"bytes,2,rep,name=facets,proto3" json:"facets,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ObservationResponse) Reset() {
	*x = ObservationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v2_observation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObservationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationResponse) ProtoMessage() {}

func (x *ObservationResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ObservationResponse.ProtoReflect.Descriptor instead.
func (*ObservationResponse) Descriptor() ([]byte, []int) {
	return file_v2_observation_proto_rawDescGZIP(), []int{4}
}

func (x *ObservationResponse) GetObservationsByVariable() map[string]*VariableObservation {
	if x != nil {
		return x.ObservationsByVariable
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
	0x74, 0x6f, 0x22, 0xf6, 0x01, 0x0a, 0x13, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x73, 0x0a, 0x16, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x14, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a,
	0x6a, 0x0a, 0x19, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x37,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x73, 0x0a, 0x11, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5e, 0x0a, 0x1a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x61, 0x63, 0x65,
	0x74, 0x5f, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x18, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x64, 0x46,
	0x61, 0x63, 0x65, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x69, 0x0a, 0x10, 0x46, 0x61, 0x63, 0x65, 0x74, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x61, 0x63, 0x65, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x63, 0x65, 0x74, 0x49, 0x64, 0x12,
	0x3a, 0x0a, 0x0c, 0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x0c, 0x6f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa1, 0x02, 0x0a, 0x12,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2f,
	0x0a, 0x13, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1c, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x31, 0x0a,
	0x14, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x6d, 0x65, 0x61, 0x73,
	0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x22,
	0x98, 0x03, 0x0a, 0x13, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a, 0x18, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x16, 0x6f, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x47, 0x0a, 0x06, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x1a, 0x6e, 0x0a, 0x1b, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4d, 0x0a, 0x0b, 0x46,
	0x61, 0x63, 0x65, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x67, 0x2f, 0x6d, 0x69, 0x78, 0x65, 0x72, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x32, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_v2_observation_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v2_observation_proto_goTypes = []interface{}{
	(*VariableObservation)(nil), // 0: datacommons.v2.VariableObservation
	(*EntityObservation)(nil),   // 1: datacommons.v2.EntityObservation
	(*FacetObservation)(nil),    // 2: datacommons.v2.FacetObservation
	(*ObservationRequest)(nil),  // 3: datacommons.v2.ObservationRequest
	(*ObservationResponse)(nil), // 4: datacommons.v2.ObservationResponse
	nil,                         // 5: datacommons.v2.VariableObservation.ObservationsByEntityEntry
	nil,                         // 6: datacommons.v2.ObservationResponse.ObservationsByVariableEntry
	nil,                         // 7: datacommons.v2.ObservationResponse.FacetsEntry
	(*proto.PointStat)(nil),     // 8: datacommons.PointStat
	(*proto.Facet)(nil),         // 9: datacommons.Facet
}
var file_v2_observation_proto_depIdxs = []int32{
	5, // 0: datacommons.v2.VariableObservation.observations_by_entity:type_name -> datacommons.v2.VariableObservation.ObservationsByEntityEntry
	2, // 1: datacommons.v2.EntityObservation.ordered_facet_observations:type_name -> datacommons.v2.FacetObservation
	8, // 2: datacommons.v2.FacetObservation.observations:type_name -> datacommons.PointStat
	6, // 3: datacommons.v2.ObservationResponse.observations_by_variable:type_name -> datacommons.v2.ObservationResponse.ObservationsByVariableEntry
	7, // 4: datacommons.v2.ObservationResponse.facets:type_name -> datacommons.v2.ObservationResponse.FacetsEntry
	1, // 5: datacommons.v2.VariableObservation.ObservationsByEntityEntry.value:type_name -> datacommons.v2.EntityObservation
	0, // 6: datacommons.v2.ObservationResponse.ObservationsByVariableEntry.value:type_name -> datacommons.v2.VariableObservation
	9, // 7: datacommons.v2.ObservationResponse.FacetsEntry.value:type_name -> datacommons.Facet
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_v2_observation_proto_init() }
func file_v2_observation_proto_init() {
	if File_v2_observation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v2_observation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_v2_observation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_v2_observation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_v2_observation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_v2_observation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
			NumMessages:   8,
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
