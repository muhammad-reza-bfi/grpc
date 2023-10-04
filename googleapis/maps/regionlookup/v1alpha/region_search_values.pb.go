// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/maps/regionlookup/v1alpha/region_search_values.proto

package regionlookup

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible place types to match to.
type RegionSearchValue_PlaceType int32

const (
	// Default value. This value is unused.
	RegionSearchValue_PLACE_TYPE_UNSPECIFIED RegionSearchValue_PlaceType = 0
	// Postal code.
	RegionSearchValue_POSTAL_CODE RegionSearchValue_PlaceType = 1
	// Administrative area level 1 (State in the US).
	RegionSearchValue_ADMINISTRATIVE_AREA_LEVEL_1 RegionSearchValue_PlaceType = 2
	// Administrative area level 2 (County in the US).
	RegionSearchValue_ADMINISTRATIVE_AREA_LEVEL_2 RegionSearchValue_PlaceType = 3
	// Locality (City).
	RegionSearchValue_LOCALITY RegionSearchValue_PlaceType = 4
	// Neighborhood.
	RegionSearchValue_NEIGHBORHOOD RegionSearchValue_PlaceType = 5
	// Country.
	RegionSearchValue_COUNTRY RegionSearchValue_PlaceType = 6
	// Sublocality.
	RegionSearchValue_SUBLOCALITY RegionSearchValue_PlaceType = 7
	// Administrative area level 3.
	RegionSearchValue_ADMINISTRATIVE_AREA_LEVEL_3 RegionSearchValue_PlaceType = 8
	// Administrative area level 4.
	RegionSearchValue_ADMINISTRATIVE_AREA_LEVEL_4 RegionSearchValue_PlaceType = 9
	// School district.
	RegionSearchValue_SCHOOL_DISTRICT RegionSearchValue_PlaceType = 10
)

// Enum value maps for RegionSearchValue_PlaceType.
var (
	RegionSearchValue_PlaceType_name = map[int32]string{
		0:  "PLACE_TYPE_UNSPECIFIED",
		1:  "POSTAL_CODE",
		2:  "ADMINISTRATIVE_AREA_LEVEL_1",
		3:  "ADMINISTRATIVE_AREA_LEVEL_2",
		4:  "LOCALITY",
		5:  "NEIGHBORHOOD",
		6:  "COUNTRY",
		7:  "SUBLOCALITY",
		8:  "ADMINISTRATIVE_AREA_LEVEL_3",
		9:  "ADMINISTRATIVE_AREA_LEVEL_4",
		10: "SCHOOL_DISTRICT",
	}
	RegionSearchValue_PlaceType_value = map[string]int32{
		"PLACE_TYPE_UNSPECIFIED":      0,
		"POSTAL_CODE":                 1,
		"ADMINISTRATIVE_AREA_LEVEL_1": 2,
		"ADMINISTRATIVE_AREA_LEVEL_2": 3,
		"LOCALITY":                    4,
		"NEIGHBORHOOD":                5,
		"COUNTRY":                     6,
		"SUBLOCALITY":                 7,
		"ADMINISTRATIVE_AREA_LEVEL_3": 8,
		"ADMINISTRATIVE_AREA_LEVEL_4": 9,
		"SCHOOL_DISTRICT":             10,
	}
)

func (x RegionSearchValue_PlaceType) Enum() *RegionSearchValue_PlaceType {
	p := new(RegionSearchValue_PlaceType)
	*p = x
	return p
}

func (x RegionSearchValue_PlaceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegionSearchValue_PlaceType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_regionlookup_v1alpha_region_search_values_proto_enumTypes[0].Descriptor()
}

func (RegionSearchValue_PlaceType) Type() protoreflect.EnumType {
	return &file_google_maps_regionlookup_v1alpha_region_search_values_proto_enumTypes[0]
}

func (x RegionSearchValue_PlaceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegionSearchValue_PlaceType.Descriptor instead.
func (RegionSearchValue_PlaceType) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescGZIP(), []int{0, 0}
}

// Region Search Values.
//
// Desired search values of a single region.
//
// Location must be specified by one of the following: address, latlng or
// place_id. If none is specified, an INVALID_ARGUMENT error is returned.
// region_code must also be provided when address is specified.
//
// The fields address, latlng and place_id specify a location contained inside
// the region to match. For example if address is "1600 Amphitheatre Pkwy,
// Mountain View, CA 94043" the API returns the following matched_place_id
// results when the following place_types are specified:
//
// place_type:                   matched_place_id results:
// postal_code                   Place ID for "94043"
// administrative_area_level_1   Place ID for The State of California
// administrative_area_level_2   Place ID for Santa Clara County
// etc.
//
// More Examples:
//
// If latlng is "latitude: 37.4220656 longitude: -122.0862784" and place_type
// is "locality", the result contains the Place ID (of type "locality") for
// that location (the Place ID of Mountain View, CA, in this case).
//
// If place_id is "ChIJj61dQgK6j4AR4GeTYWZsKWw" (Place ID for Google office in
// Mountain view, CA) and place_type is "locality", the result contains the
// Place ID (of type "locality") for that location (the Place ID of Mountain
// View, CA, in this case).
//
// If no match is found, matched_place_id is not set.
//
// Candidates Place IDs are returned when a search finds multiple Place
// IDs for the location specified. For example if the API is searching for
// region Place IDs of type neighboorhood for a location that is contained
// within multiple neighboords. The Place Ids will be returned as candidates in
// the candidate_place_ids field.
//
// Next available tag: 10
type RegionSearchValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location must be specified by one of the following:
	//
	// Types that are assignable to Location:
	//	*RegionSearchValue_Address
	//	*RegionSearchValue_Latlng
	//	*RegionSearchValue_PlaceId
	Location isRegionSearchValue_Location `protobuf_oneof:"location"`
	// Required. The type of the place to match.
	PlaceType RegionSearchValue_PlaceType `protobuf:"varint,6,opt,name=place_type,json=placeType,proto3,enum=google.maps.regionlookup.v1alpha.RegionSearchValue_PlaceType" json:"place_type,omitempty"`
	// The BCP-47 language code, such as "en-US" or "sr-Latn", corresponding to
	// the language in which the place name and address is requested. If none is
	// requested, then it defaults to English.
	LanguageCode string `protobuf:"bytes,7,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Two-letter ISO-3166 country/region code for the location you're trying to
	// match. region_code is required when address is specified.
	RegionCode string `protobuf:"bytes,8,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
}

func (x *RegionSearchValue) Reset() {
	*x = RegionSearchValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_regionlookup_v1alpha_region_search_values_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionSearchValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionSearchValue) ProtoMessage() {}

func (x *RegionSearchValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_regionlookup_v1alpha_region_search_values_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionSearchValue.ProtoReflect.Descriptor instead.
func (*RegionSearchValue) Descriptor() ([]byte, []int) {
	return file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescGZIP(), []int{0}
}

func (m *RegionSearchValue) GetLocation() isRegionSearchValue_Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (x *RegionSearchValue) GetAddress() string {
	if x, ok := x.GetLocation().(*RegionSearchValue_Address); ok {
		return x.Address
	}
	return ""
}

func (x *RegionSearchValue) GetLatlng() *latlng.LatLng {
	if x, ok := x.GetLocation().(*RegionSearchValue_Latlng); ok {
		return x.Latlng
	}
	return nil
}

func (x *RegionSearchValue) GetPlaceId() string {
	if x, ok := x.GetLocation().(*RegionSearchValue_PlaceId); ok {
		return x.PlaceId
	}
	return ""
}

func (x *RegionSearchValue) GetPlaceType() RegionSearchValue_PlaceType {
	if x != nil {
		return x.PlaceType
	}
	return RegionSearchValue_PLACE_TYPE_UNSPECIFIED
}

func (x *RegionSearchValue) GetLanguageCode() string {
	if x != nil {
		return x.LanguageCode
	}
	return ""
}

func (x *RegionSearchValue) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

type isRegionSearchValue_Location interface {
	isRegionSearchValue_Location()
}

type RegionSearchValue_Address struct {
	// The unstructured street address that is contained inside a region to
	// match. region_code is required when address is specified.
	Address string `protobuf:"bytes,1,opt,name=address,proto3,oneof"`
}

type RegionSearchValue_Latlng struct {
	// The latitude and longitude that is contained inside a region to match.
	Latlng *latlng.LatLng `protobuf:"bytes,2,opt,name=latlng,proto3,oneof"`
}

type RegionSearchValue_PlaceId struct {
	// The Place ID that is contained inside a region to match.
	PlaceId string `protobuf:"bytes,3,opt,name=place_id,json=placeId,proto3,oneof"`
}

func (*RegionSearchValue_Address) isRegionSearchValue_Location() {}

func (*RegionSearchValue_Latlng) isRegionSearchValue_Location() {}

func (*RegionSearchValue_PlaceId) isRegionSearchValue_Location() {}

var File_google_maps_regionlookup_v1alpha_region_search_values_proto protoreflect.FileDescriptor

var file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6c, 0x61,
	0x74, 0x6c, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x04, 0x0a, 0x11, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1a, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x06,
	0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x4c, 0x61, 0x74, 0x4c, 0x6e,
	0x67, 0x48, 0x00, 0x52, 0x06, 0x6c, 0x61, 0x74, 0x6c, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x61, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x09, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64,
	0x65, 0x22, 0x8f, 0x02, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x16, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50,
	0x4f, 0x53, 0x54, 0x41, 0x4c, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x41,
	0x52, 0x45, 0x41, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x31, 0x10, 0x02, 0x12, 0x1f, 0x0a,
	0x1b, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x41, 0x52, 0x45, 0x41, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x32, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c,
	0x4e, 0x45, 0x49, 0x47, 0x48, 0x42, 0x4f, 0x52, 0x48, 0x4f, 0x4f, 0x44, 0x10, 0x05, 0x12, 0x0b,
	0x0a, 0x07, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x53,
	0x55, 0x42, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f, 0x41,
	0x52, 0x45, 0x41, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x33, 0x10, 0x08, 0x12, 0x1f, 0x0a,
	0x1b, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x49, 0x53, 0x54, 0x52, 0x41, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x41, 0x52, 0x45, 0x41, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x34, 0x10, 0x09, 0x12, 0x13,
	0x0a, 0x0f, 0x53, 0x43, 0x48, 0x4f, 0x4f, 0x4c, 0x5f, 0x44, 0x49, 0x53, 0x54, 0x52, 0x49, 0x43,
	0x54, 0x10, 0x0a, 0x42, 0x0a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0xe1, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d,
	0x61, 0x70, 0x73, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x17, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x06, 0x4d, 0x52, 0x4c, 0x56, 0x31, 0x41, 0xaa, 0x02, 0x20,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61,
	0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescOnce sync.Once
	file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescData = file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDesc
)

func file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescGZIP() []byte {
	file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescOnce.Do(func() {
		file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescData)
	})
	return file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDescData
}

var file_google_maps_regionlookup_v1alpha_region_search_values_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_regionlookup_v1alpha_region_search_values_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_maps_regionlookup_v1alpha_region_search_values_proto_goTypes = []interface{}{
	(RegionSearchValue_PlaceType)(0), // 0: google.maps.regionlookup.v1alpha.RegionSearchValue.PlaceType
	(*RegionSearchValue)(nil),        // 1: google.maps.regionlookup.v1alpha.RegionSearchValue
	(*latlng.LatLng)(nil),            // 2: google.type.LatLng
}
var file_google_maps_regionlookup_v1alpha_region_search_values_proto_depIdxs = []int32{
	2, // 0: google.maps.regionlookup.v1alpha.RegionSearchValue.latlng:type_name -> google.type.LatLng
	0, // 1: google.maps.regionlookup.v1alpha.RegionSearchValue.place_type:type_name -> google.maps.regionlookup.v1alpha.RegionSearchValue.PlaceType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_maps_regionlookup_v1alpha_region_search_values_proto_init() }
func file_google_maps_regionlookup_v1alpha_region_search_values_proto_init() {
	if File_google_maps_regionlookup_v1alpha_region_search_values_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_maps_regionlookup_v1alpha_region_search_values_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionSearchValue); i {
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
	file_google_maps_regionlookup_v1alpha_region_search_values_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RegionSearchValue_Address)(nil),
		(*RegionSearchValue_Latlng)(nil),
		(*RegionSearchValue_PlaceId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_regionlookup_v1alpha_region_search_values_proto_goTypes,
		DependencyIndexes: file_google_maps_regionlookup_v1alpha_region_search_values_proto_depIdxs,
		EnumInfos:         file_google_maps_regionlookup_v1alpha_region_search_values_proto_enumTypes,
		MessageInfos:      file_google_maps_regionlookup_v1alpha_region_search_values_proto_msgTypes,
	}.Build()
	File_google_maps_regionlookup_v1alpha_region_search_values_proto = out.File
	file_google_maps_regionlookup_v1alpha_region_search_values_proto_rawDesc = nil
	file_google_maps_regionlookup_v1alpha_region_search_values_proto_goTypes = nil
	file_google_maps_regionlookup_v1alpha_region_search_values_proto_depIdxs = nil
}
