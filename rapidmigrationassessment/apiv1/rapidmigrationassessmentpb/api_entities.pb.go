// Copyright 2023 Google LLC
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
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: google/cloud/rapidmigrationassessment/v1/api_entities.proto

package rapidmigrationassessmentpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// -- Using suggestion from API Linter Analyzer for nesting enum --
// -- https://linter.aip.dev/216/nesting --
// State of a Collector (server_side).
// States are used for internal purposes and named to keep
// convention of legacy product:
// https://cloud.google.com/migrate/stratozone/docs/about-stratoprobe.
type Collector_State int32

const (
	// Collector state is not recognized.
	Collector_STATE_UNSPECIFIED Collector_State = 0
	// Collector started to create, but hasn't been completed MC source creation
	// and db object creation.
	Collector_STATE_INITIALIZING Collector_State = 1
	// Collector has been created, MC source creation and db object creation
	// completed.
	Collector_STATE_READY_TO_USE Collector_State = 2
	// Collector client has been registered with client.
	Collector_STATE_REGISTERED Collector_State = 3
	// Collector client is actively scanning.
	Collector_STATE_ACTIVE Collector_State = 4
	// Collector is not actively scanning.
	Collector_STATE_PAUSED Collector_State = 5
	// Collector is starting background job for deletion.
	Collector_STATE_DELETING Collector_State = 6
	// Collector completed all tasks for deletion.
	Collector_STATE_DECOMMISSIONED Collector_State = 7
	// Collector is in error state.
	Collector_STATE_ERROR Collector_State = 8
)

// Enum value maps for Collector_State.
var (
	Collector_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_INITIALIZING",
		2: "STATE_READY_TO_USE",
		3: "STATE_REGISTERED",
		4: "STATE_ACTIVE",
		5: "STATE_PAUSED",
		6: "STATE_DELETING",
		7: "STATE_DECOMMISSIONED",
		8: "STATE_ERROR",
	}
	Collector_State_value = map[string]int32{
		"STATE_UNSPECIFIED":    0,
		"STATE_INITIALIZING":   1,
		"STATE_READY_TO_USE":   2,
		"STATE_REGISTERED":     3,
		"STATE_ACTIVE":         4,
		"STATE_PAUSED":         5,
		"STATE_DELETING":       6,
		"STATE_DECOMMISSIONED": 7,
		"STATE_ERROR":          8,
	}
)

func (x Collector_State) Enum() *Collector_State {
	p := new(Collector_State)
	*p = x
	return p
}

func (x Collector_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Collector_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_enumTypes[0].Descriptor()
}

func (Collector_State) Type() protoreflect.EnumType {
	return &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_enumTypes[0]
}

func (x Collector_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Collector_State.Descriptor instead.
func (Collector_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescGZIP(), []int{2, 0}
}

// Types for project level setting.
type Annotation_Type int32

const (
	// Unknown type
	Annotation_TYPE_UNSPECIFIED Annotation_Type = 0
	// Indicates that this project has opted into StratoZone export.
	Annotation_TYPE_LEGACY_EXPORT_CONSENT Annotation_Type = 1
	// Indicates that this project is created by Qwiklab.
	Annotation_TYPE_QWIKLAB Annotation_Type = 2
)

// Enum value maps for Annotation_Type.
var (
	Annotation_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_LEGACY_EXPORT_CONSENT",
		2: "TYPE_QWIKLAB",
	}
	Annotation_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":           0,
		"TYPE_LEGACY_EXPORT_CONSENT": 1,
		"TYPE_QWIKLAB":               2,
	}
)

func (x Annotation_Type) Enum() *Annotation_Type {
	p := new(Annotation_Type)
	*p = x
	return p
}

func (x Annotation_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Annotation_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_enumTypes[1].Descriptor()
}

func (Annotation_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_enumTypes[1]
}

func (x Annotation_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Annotation_Type.Descriptor instead.
func (Annotation_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescGZIP(), []int{3, 0}
}

// Message describing a MC Source of type Guest OS Scan.
type GuestOsScan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reference to the corresponding Guest OS Scan in MC Source.
	CoreSource string `protobuf:"bytes,1,opt,name=core_source,json=coreSource,proto3" json:"core_source,omitempty"`
}

func (x *GuestOsScan) Reset() {
	*x = GuestOsScan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuestOsScan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuestOsScan) ProtoMessage() {}

func (x *GuestOsScan) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuestOsScan.ProtoReflect.Descriptor instead.
func (*GuestOsScan) Descriptor() ([]byte, []int) {
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescGZIP(), []int{0}
}

func (x *GuestOsScan) GetCoreSource() string {
	if x != nil {
		return x.CoreSource
	}
	return ""
}

// Message describing a MC Source of type VSphere Scan.
type VSphereScan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// reference to the corresponding VSphere Scan in MC Source.
	CoreSource string `protobuf:"bytes,1,opt,name=core_source,json=coreSource,proto3" json:"core_source,omitempty"`
}

func (x *VSphereScan) Reset() {
	*x = VSphereScan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VSphereScan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VSphereScan) ProtoMessage() {}

func (x *VSphereScan) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VSphereScan.ProtoReflect.Descriptor instead.
func (*VSphereScan) Descriptor() ([]byte, []int) {
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescGZIP(), []int{1}
}

func (x *VSphereScan) GetCoreSource() string {
	if x != nil {
		return x.CoreSource
	}
	return ""
}

// Message describing Collector object.
type Collector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Create time stamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Update time stamp.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Labels as key value pairs.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User specified name of the Collector.
	DisplayName string `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// User specified description of the Collector.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Service Account email used to ingest data to this Collector.
	ServiceAccount string `protobuf:"bytes,7,opt,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// Output only. Store cloud storage bucket name (which is a guid) created with
	// this Collector.
	Bucket string `protobuf:"bytes,8,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// User specified expected asset count.
	ExpectedAssetCount int64 `protobuf:"varint,9,opt,name=expected_asset_count,json=expectedAssetCount,proto3" json:"expected_asset_count,omitempty"`
	// Output only. State of the Collector.
	State Collector_State `protobuf:"varint,10,opt,name=state,proto3,enum=google.cloud.rapidmigrationassessment.v1.Collector_State" json:"state,omitempty"`
	// Output only. Client version.
	ClientVersion string `protobuf:"bytes,11,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	// Output only. Reference to MC Source Guest Os Scan.
	GuestOsScan *GuestOsScan `protobuf:"bytes,12,opt,name=guest_os_scan,json=guestOsScan,proto3" json:"guest_os_scan,omitempty"`
	// Output only. Reference to MC Source vsphere_scan.
	VsphereScan *VSphereScan `protobuf:"bytes,13,opt,name=vsphere_scan,json=vsphereScan,proto3" json:"vsphere_scan,omitempty"`
	// How many days to collect data.
	CollectionDays int32 `protobuf:"varint,14,opt,name=collection_days,json=collectionDays,proto3" json:"collection_days,omitempty"`
	// Uri for EULA (End User License Agreement) from customer.
	EulaUri string `protobuf:"bytes,15,opt,name=eula_uri,json=eulaUri,proto3" json:"eula_uri,omitempty"`
}

func (x *Collector) Reset() {
	*x = Collector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collector) ProtoMessage() {}

func (x *Collector) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collector.ProtoReflect.Descriptor instead.
func (*Collector) Descriptor() ([]byte, []int) {
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescGZIP(), []int{2}
}

func (x *Collector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collector) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Collector) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Collector) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Collector) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Collector) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Collector) GetServiceAccount() string {
	if x != nil {
		return x.ServiceAccount
	}
	return ""
}

func (x *Collector) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Collector) GetExpectedAssetCount() int64 {
	if x != nil {
		return x.ExpectedAssetCount
	}
	return 0
}

func (x *Collector) GetState() Collector_State {
	if x != nil {
		return x.State
	}
	return Collector_STATE_UNSPECIFIED
}

func (x *Collector) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *Collector) GetGuestOsScan() *GuestOsScan {
	if x != nil {
		return x.GuestOsScan
	}
	return nil
}

func (x *Collector) GetVsphereScan() *VSphereScan {
	if x != nil {
		return x.VsphereScan
	}
	return nil
}

func (x *Collector) GetCollectionDays() int32 {
	if x != nil {
		return x.CollectionDays
	}
	return 0
}

func (x *Collector) GetEulaUri() string {
	if x != nil {
		return x.EulaUri
	}
	return ""
}

// Message describing an Annotation
type Annotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of resource.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Output only. Create time stamp.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Update time stamp.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Labels as key value pairs.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Type of an annotation.
	Type Annotation_Type `protobuf:"varint,5,opt,name=type,proto3,enum=google.cloud.rapidmigrationassessment.v1.Annotation_Type" json:"type,omitempty"`
}

func (x *Annotation) Reset() {
	*x = Annotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Annotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Annotation) ProtoMessage() {}

func (x *Annotation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Annotation.ProtoReflect.Descriptor instead.
func (*Annotation) Descriptor() ([]byte, []int) {
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescGZIP(), []int{3}
}

func (x *Annotation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Annotation) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Annotation) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Annotation) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Annotation) GetType() Annotation_Type {
	if x != nil {
		return x.Type
	}
	return Annotation_TYPE_UNSPECIFIED
}

var File_google_cloud_rapidmigrationassessment_v1_api_entities_proto protoreflect.FileDescriptor

var file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x61, 0x70, 0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x61, 0x70, 0x69,
	0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0b, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x73, 0x53,
	0x63, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x2e, 0x0a, 0x0b, 0x56, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x53,
	0x63, 0x61, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0xbb, 0x09, 0x0a, 0x09, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x30,
	0x0a, 0x14, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x65, 0x78,
	0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x73, 0x73, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x54, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72,
	0x61, 0x70, 0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x0d, 0x67, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6f, 0x73, 0x5f, 0x73,
	0x63, 0x61, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x73, 0x53, 0x63, 0x61, 0x6e,
	0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x67, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x73, 0x53, 0x63,
	0x61, 0x6e, 0x12, 0x5d, 0x0a, 0x0c, 0x76, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x5f, 0x73, 0x63,
	0x61, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x53, 0x70, 0x68, 0x65, 0x72, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x42,
	0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x76, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x53, 0x63, 0x61,
	0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x79, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x79, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x75,
	0x6c, 0x61, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x75,
	0x6c, 0x61, 0x55, 0x72, 0x69, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xc7, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49,
	0x41, 0x4c, 0x49, 0x5a, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x54, 0x4f, 0x5f, 0x55, 0x53, 0x45, 0x10,
	0x02, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53,
	0x54, 0x45, 0x52, 0x45, 0x44, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12,
	0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x45, 0x43, 0x4f, 0x4d, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x45, 0x44, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x08, 0x3a, 0x76, 0xea, 0x41, 0x73, 0x0a,
	0x31, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x3e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x7d, 0x22, 0xd3, 0x04, 0x0a, 0x0a, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x58, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x4d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4e, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x5f, 0x45, 0x58, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x51, 0x57, 0x49, 0x4b, 0x4c, 0x41, 0x42, 0x10, 0x02, 0x3a, 0x79, 0xea,
	0x41, 0x76, 0x0a, 0x32, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x42, 0xb0, 0x02, 0x0a, 0x2c, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x61,
	0x70, 0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73, 0x65,
	0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x41, 0x70, 0x69, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x68, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x6f, 0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x2f, 0x72, 0x61, 0x70, 0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x3b, 0x72, 0x61, 0x70,
	0x69, 0x64, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x73, 0x73, 0x65, 0x73,
	0x73, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xaa, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x61, 0x70, 0x69, 0x64, 0x4d, 0x69, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x5c, 0x52, 0x61, 0x70, 0x69, 0x64, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x73, 0x73, 0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x2b,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52,
	0x61, 0x70, 0x69, 0x64, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x73, 0x73,
	0x65, 0x73, 0x73, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescOnce sync.Once
	file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescData = file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDesc
)

func file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescGZIP() []byte {
	file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescOnce.Do(func() {
		file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescData)
	})
	return file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDescData
}

var file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_goTypes = []interface{}{
	(Collector_State)(0),          // 0: google.cloud.rapidmigrationassessment.v1.Collector.State
	(Annotation_Type)(0),          // 1: google.cloud.rapidmigrationassessment.v1.Annotation.Type
	(*GuestOsScan)(nil),           // 2: google.cloud.rapidmigrationassessment.v1.GuestOsScan
	(*VSphereScan)(nil),           // 3: google.cloud.rapidmigrationassessment.v1.VSphereScan
	(*Collector)(nil),             // 4: google.cloud.rapidmigrationassessment.v1.Collector
	(*Annotation)(nil),            // 5: google.cloud.rapidmigrationassessment.v1.Annotation
	nil,                           // 6: google.cloud.rapidmigrationassessment.v1.Collector.LabelsEntry
	nil,                           // 7: google.cloud.rapidmigrationassessment.v1.Annotation.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_depIdxs = []int32{
	8,  // 0: google.cloud.rapidmigrationassessment.v1.Collector.create_time:type_name -> google.protobuf.Timestamp
	8,  // 1: google.cloud.rapidmigrationassessment.v1.Collector.update_time:type_name -> google.protobuf.Timestamp
	6,  // 2: google.cloud.rapidmigrationassessment.v1.Collector.labels:type_name -> google.cloud.rapidmigrationassessment.v1.Collector.LabelsEntry
	0,  // 3: google.cloud.rapidmigrationassessment.v1.Collector.state:type_name -> google.cloud.rapidmigrationassessment.v1.Collector.State
	2,  // 4: google.cloud.rapidmigrationassessment.v1.Collector.guest_os_scan:type_name -> google.cloud.rapidmigrationassessment.v1.GuestOsScan
	3,  // 5: google.cloud.rapidmigrationassessment.v1.Collector.vsphere_scan:type_name -> google.cloud.rapidmigrationassessment.v1.VSphereScan
	8,  // 6: google.cloud.rapidmigrationassessment.v1.Annotation.create_time:type_name -> google.protobuf.Timestamp
	8,  // 7: google.cloud.rapidmigrationassessment.v1.Annotation.update_time:type_name -> google.protobuf.Timestamp
	7,  // 8: google.cloud.rapidmigrationassessment.v1.Annotation.labels:type_name -> google.cloud.rapidmigrationassessment.v1.Annotation.LabelsEntry
	1,  // 9: google.cloud.rapidmigrationassessment.v1.Annotation.type:type_name -> google.cloud.rapidmigrationassessment.v1.Annotation.Type
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_init() }
func file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_init() {
	if File_google_cloud_rapidmigrationassessment_v1_api_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuestOsScan); i {
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
		file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VSphereScan); i {
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
		file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collector); i {
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
		file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Annotation); i {
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
			RawDescriptor: file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_goTypes,
		DependencyIndexes: file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_depIdxs,
		EnumInfos:         file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_enumTypes,
		MessageInfos:      file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_msgTypes,
	}.Build()
	File_google_cloud_rapidmigrationassessment_v1_api_entities_proto = out.File
	file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_rawDesc = nil
	file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_goTypes = nil
	file_google_cloud_rapidmigrationassessment_v1_api_entities_proto_depIdxs = nil
}
