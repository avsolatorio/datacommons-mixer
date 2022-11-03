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
//    /internal/place (/landing-page)
//    /internal/bio
// ========================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: internal.proto

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

type Place struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dcid string `protobuf:"bytes,1,opt,name=dcid,proto3" json:"dcid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Pop  int32  `protobuf:"varint,3,opt,name=pop,proto3" json:"pop,omitempty"`
}

func (x *Place) Reset() {
	*x = Place{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Place) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Place) ProtoMessage() {}

func (x *Place) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Place.ProtoReflect.Descriptor instead.
func (*Place) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{0}
}

func (x *Place) GetDcid() string {
	if x != nil {
		return x.Dcid
	}
	return ""
}

func (x *Place) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Place) GetPop() int32 {
	if x != nil {
		return x.Pop
	}
	return 0
}

type Places struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Places []*Place `protobuf:"bytes,1,rep,name=places,proto3" json:"places,omitempty"`
}

func (x *Places) Reset() {
	*x = Places{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Places) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Places) ProtoMessage() {}

func (x *Places) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Places.ProtoReflect.Descriptor instead.
func (*Places) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{1}
}

func (x *Places) GetPlaces() []*Place {
	if x != nil {
		return x.Places
	}
	return nil
}

// Request to get all data in place page.
type GetPlacePageDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The dcid of the place.
	Place string `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
	// A list of additional stat vars need to be fetched in addition to the
	// data in cache. This is to be used in local development, where new chart is
	// added to chart config but data is not added to cache (delay from cache
	// build).
	NewStatVars []string `protobuf:"bytes,4,rep,name=new_stat_vars,json=newStatVars,proto3" json:"new_stat_vars,omitempty"`
	// Seed value for random selection. Used by test to get deterministic result.
	Seed int64 `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (x *GetPlacePageDataRequest) Reset() {
	*x = GetPlacePageDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlacePageDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlacePageDataRequest) ProtoMessage() {}

func (x *GetPlacePageDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlacePageDataRequest.ProtoReflect.Descriptor instead.
func (*GetPlacePageDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{2}
}

func (x *GetPlacePageDataRequest) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *GetPlacePageDataRequest) GetNewStatVars() []string {
	if x != nil {
		return x.NewStatVars
	}
	return nil
}

func (x *GetPlacePageDataRequest) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

// Response to get place page info for a place.
type GetPlacePageDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Keyed by place dcid.
	StatVarSeries    map[string]*StatVarSeries `protobuf:"bytes,1,rep,name=stat_var_series,json=statVarSeries,proto3" json:"stat_var_series,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AllChildPlaces   map[string]*Places        `protobuf:"bytes,2,rep,name=all_child_places,json=allChildPlaces,proto3" json:"all_child_places,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	LatestPopulation map[string]*PointStat     `protobuf:"bytes,8,rep,name=latest_population,json=latestPopulation,proto3" json:"latest_population,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ValidCategories  map[string]*ObsCategories `protobuf:"bytes,9,rep,name=valid_categories,json=validCategories,proto3" json:"valid_categories,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ChildPlacesType  string                    `protobuf:"bytes,3,opt,name=child_places_type,json=childPlacesType,proto3" json:"child_places_type,omitempty"`
	ChildPlaces      []string                  `protobuf:"bytes,4,rep,name=child_places,json=childPlaces,proto3" json:"child_places,omitempty"`
	ParentPlaces     []string                  `protobuf:"bytes,5,rep,name=parent_places,json=parentPlaces,proto3" json:"parent_places,omitempty"`
	SimilarPlaces    []string                  `protobuf:"bytes,6,rep,name=similar_places,json=similarPlaces,proto3" json:"similar_places,omitempty"`
	NearbyPlaces     []string                  `protobuf:"bytes,7,rep,name=nearby_places,json=nearbyPlaces,proto3" json:"nearby_places,omitempty"`
}

func (x *GetPlacePageDataResponse) Reset() {
	*x = GetPlacePageDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlacePageDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlacePageDataResponse) ProtoMessage() {}

func (x *GetPlacePageDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlacePageDataResponse.ProtoReflect.Descriptor instead.
func (*GetPlacePageDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlacePageDataResponse) GetStatVarSeries() map[string]*StatVarSeries {
	if x != nil {
		return x.StatVarSeries
	}
	return nil
}

func (x *GetPlacePageDataResponse) GetAllChildPlaces() map[string]*Places {
	if x != nil {
		return x.AllChildPlaces
	}
	return nil
}

func (x *GetPlacePageDataResponse) GetLatestPopulation() map[string]*PointStat {
	if x != nil {
		return x.LatestPopulation
	}
	return nil
}

func (x *GetPlacePageDataResponse) GetValidCategories() map[string]*ObsCategories {
	if x != nil {
		return x.ValidCategories
	}
	return nil
}

func (x *GetPlacePageDataResponse) GetChildPlacesType() string {
	if x != nil {
		return x.ChildPlacesType
	}
	return ""
}

func (x *GetPlacePageDataResponse) GetChildPlaces() []string {
	if x != nil {
		return x.ChildPlaces
	}
	return nil
}

func (x *GetPlacePageDataResponse) GetParentPlaces() []string {
	if x != nil {
		return x.ParentPlaces
	}
	return nil
}

func (x *GetPlacePageDataResponse) GetSimilarPlaces() []string {
	if x != nil {
		return x.SimilarPlaces
	}
	return nil
}

func (x *GetPlacePageDataResponse) GetNearbyPlaces() []string {
	if x != nil {
		return x.NearbyPlaces
	}
	return nil
}

// Request to get all data in bio page.
type GetBioPageDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The dcid of the entity
	Dcid string `protobuf:"bytes,1,opt,name=dcid,proto3" json:"dcid,omitempty"`
}

func (x *GetBioPageDataRequest) Reset() {
	*x = GetBioPageDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBioPageDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBioPageDataRequest) ProtoMessage() {}

func (x *GetBioPageDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBioPageDataRequest.ProtoReflect.Descriptor instead.
func (*GetBioPageDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{4}
}

func (x *GetBioPageDataRequest) GetDcid() string {
	if x != nil {
		return x.Dcid
	}
	return ""
}

// List of categories of observations for a given place. These correspond to the
// sections of the Place Explorer (https://datacommons.org/place/geoId/06).
// This is a copied proto from g3.
type ObsCategories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category  []string `protobuf:"bytes,1,rep,name=category,proto3" json:"category,omitempty"`
	PlaceName *string  `protobuf:"bytes,2,opt,name=place_name,json=placeName,proto3,oneof" json:"place_name,omitempty"`
	PlaceDcid *string  `protobuf:"bytes,3,opt,name=place_dcid,json=placeDcid,proto3,oneof" json:"place_dcid,omitempty"`
}

func (x *ObsCategories) Reset() {
	*x = ObsCategories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObsCategories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObsCategories) ProtoMessage() {}

func (x *ObsCategories) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObsCategories.ProtoReflect.Descriptor instead.
func (*ObsCategories) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{5}
}

func (x *ObsCategories) GetCategory() []string {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *ObsCategories) GetPlaceName() string {
	if x != nil && x.PlaceName != nil {
		return *x.PlaceName
	}
	return ""
}

func (x *ObsCategories) GetPlaceDcid() string {
	if x != nil && x.PlaceDcid != nil {
		return *x.PlaceDcid
	}
	return ""
}

// Data received from the cache for the landing page.
type LandingPageCache struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is stat var dcid.
	Data map[string]*ObsTimeSeries `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LandingPageCache) Reset() {
	*x = LandingPageCache{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LandingPageCache) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LandingPageCache) ProtoMessage() {}

func (x *LandingPageCache) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LandingPageCache.ProtoReflect.Descriptor instead.
func (*LandingPageCache) Descriptor() ([]byte, []int) {
	return file_internal_proto_rawDescGZIP(), []int{6}
}

func (x *LandingPageCache) GetData() map[string]*ObsTimeSeries {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_internal_proto protoreflect.FileDescriptor

var file_internal_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x63, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x6f, 0x70, 0x22, 0x34, 0x0a, 0x06,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x22, 0x6d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x5f,
	0x76, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x65, 0x65, 0x64, 0x4a, 0x04, 0x08, 0x02, 0x10,
	0x03, 0x22, 0xe5, 0x07, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60,
	0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x50, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x56, 0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x63, 0x0a, 0x10, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x68, 0x0a, 0x11, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x70, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f,
	0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x65, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x69,
	0x6d, 0x69, 0x6c, 0x61, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x69, 0x6d, 0x69, 0x6c, 0x61, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x61, 0x72, 0x62, 0x79, 0x5f, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x61, 0x72, 0x62, 0x79,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x1a, 0x5c, 0x0a, 0x12, 0x53, 0x74, 0x61, 0x74, 0x56, 0x61,
	0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x56, 0x61, 0x72, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x56, 0x0a, 0x13, 0x41, 0x6c, 0x6c, 0x43, 0x68, 0x69, 0x6c, 0x64,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65,
	0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5b, 0x0a, 0x15,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6f, 0x70, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x5e, 0x0a, 0x14, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73,
	0x2e, 0x4f, 0x62, 0x73, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x42, 0x69, 0x6f, 0x50, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x63, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x63, 0x69, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x0d, 0x4f, 0x62, 0x73, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x5f, 0x64, 0x63, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x09,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x44, 0x63, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x5f,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x5f, 0x64, 0x63, 0x69, 0x64, 0x22, 0xaa, 0x01, 0x0a, 0x10, 0x4c,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x61, 0x67, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12,
	0x3b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x50, 0x61, 0x67, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x53, 0x0a, 0x09,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x62, 0x73, 0x54, 0x69, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_rawDescOnce sync.Once
	file_internal_proto_rawDescData = file_internal_proto_rawDesc
)

func file_internal_proto_rawDescGZIP() []byte {
	file_internal_proto_rawDescOnce.Do(func() {
		file_internal_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_rawDescData)
	})
	return file_internal_proto_rawDescData
}

var file_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_proto_goTypes = []interface{}{
	(*Place)(nil),                    // 0: datacommons.Place
	(*Places)(nil),                   // 1: datacommons.Places
	(*GetPlacePageDataRequest)(nil),  // 2: datacommons.GetPlacePageDataRequest
	(*GetPlacePageDataResponse)(nil), // 3: datacommons.GetPlacePageDataResponse
	(*GetBioPageDataRequest)(nil),    // 4: datacommons.GetBioPageDataRequest
	(*ObsCategories)(nil),            // 5: datacommons.ObsCategories
	(*LandingPageCache)(nil),         // 6: datacommons.LandingPageCache
	nil,                              // 7: datacommons.GetPlacePageDataResponse.StatVarSeriesEntry
	nil,                              // 8: datacommons.GetPlacePageDataResponse.AllChildPlacesEntry
	nil,                              // 9: datacommons.GetPlacePageDataResponse.LatestPopulationEntry
	nil,                              // 10: datacommons.GetPlacePageDataResponse.ValidCategoriesEntry
	nil,                              // 11: datacommons.LandingPageCache.DataEntry
	(*StatVarSeries)(nil),            // 12: datacommons.StatVarSeries
	(*PointStat)(nil),                // 13: datacommons.PointStat
	(*ObsTimeSeries)(nil),            // 14: datacommons.ObsTimeSeries
}
var file_internal_proto_depIdxs = []int32{
	0,  // 0: datacommons.Places.places:type_name -> datacommons.Place
	7,  // 1: datacommons.GetPlacePageDataResponse.stat_var_series:type_name -> datacommons.GetPlacePageDataResponse.StatVarSeriesEntry
	8,  // 2: datacommons.GetPlacePageDataResponse.all_child_places:type_name -> datacommons.GetPlacePageDataResponse.AllChildPlacesEntry
	9,  // 3: datacommons.GetPlacePageDataResponse.latest_population:type_name -> datacommons.GetPlacePageDataResponse.LatestPopulationEntry
	10, // 4: datacommons.GetPlacePageDataResponse.valid_categories:type_name -> datacommons.GetPlacePageDataResponse.ValidCategoriesEntry
	11, // 5: datacommons.LandingPageCache.data:type_name -> datacommons.LandingPageCache.DataEntry
	12, // 6: datacommons.GetPlacePageDataResponse.StatVarSeriesEntry.value:type_name -> datacommons.StatVarSeries
	1,  // 7: datacommons.GetPlacePageDataResponse.AllChildPlacesEntry.value:type_name -> datacommons.Places
	13, // 8: datacommons.GetPlacePageDataResponse.LatestPopulationEntry.value:type_name -> datacommons.PointStat
	5,  // 9: datacommons.GetPlacePageDataResponse.ValidCategoriesEntry.value:type_name -> datacommons.ObsCategories
	14, // 10: datacommons.LandingPageCache.DataEntry.value:type_name -> datacommons.ObsTimeSeries
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_internal_proto_init() }
func file_internal_proto_init() {
	if File_internal_proto != nil {
		return
	}
	file_stat_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Place); i {
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
		file_internal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Places); i {
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
		file_internal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlacePageDataRequest); i {
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
		file_internal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlacePageDataResponse); i {
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
		file_internal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBioPageDataRequest); i {
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
		file_internal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObsCategories); i {
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
		file_internal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LandingPageCache); i {
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
	file_internal_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_goTypes,
		DependencyIndexes: file_internal_proto_depIdxs,
		MessageInfos:      file_internal_proto_msgTypes,
	}.Build()
	File_internal_proto = out.File
	file_internal_proto_rawDesc = nil
	file_internal_proto_goTypes = nil
	file_internal_proto_depIdxs = nil
}
