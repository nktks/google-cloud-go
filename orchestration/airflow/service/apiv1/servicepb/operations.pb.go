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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: google/cloud/orchestration/airflow/service/v1/operations.proto

package servicepb

import (
	reflect "reflect"
	sync "sync"

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

// An enum describing the overall state of an operation.
type OperationMetadata_State int32

const (
	// Unused.
	OperationMetadata_STATE_UNSPECIFIED OperationMetadata_State = 0
	// The operation has been created but is not yet started.
	OperationMetadata_PENDING OperationMetadata_State = 1
	// The operation is underway.
	OperationMetadata_RUNNING OperationMetadata_State = 2
	// The operation completed successfully.
	OperationMetadata_SUCCEEDED  OperationMetadata_State = 3
	OperationMetadata_SUCCESSFUL OperationMetadata_State = 3
	// The operation is no longer running but did not succeed.
	OperationMetadata_FAILED OperationMetadata_State = 4
)

// Enum value maps for OperationMetadata_State.
var (
	OperationMetadata_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "PENDING",
		2: "RUNNING",
		3: "SUCCEEDED",
		// Duplicate value: 3: "SUCCESSFUL",
		4: "FAILED",
	}
	OperationMetadata_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"PENDING":           1,
		"RUNNING":           2,
		"SUCCEEDED":         3,
		"SUCCESSFUL":        3,
		"FAILED":            4,
	}
)

func (x OperationMetadata_State) Enum() *OperationMetadata_State {
	p := new(OperationMetadata_State)
	*p = x
	return p
}

func (x OperationMetadata_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationMetadata_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_orchestration_airflow_service_v1_operations_proto_enumTypes[0].Descriptor()
}

func (OperationMetadata_State) Type() protoreflect.EnumType {
	return &file_google_cloud_orchestration_airflow_service_v1_operations_proto_enumTypes[0]
}

func (x OperationMetadata_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationMetadata_State.Descriptor instead.
func (OperationMetadata_State) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescGZIP(), []int{0, 0}
}

// Type of longrunning operation.
type OperationMetadata_Type int32

const (
	// Unused.
	OperationMetadata_TYPE_UNSPECIFIED OperationMetadata_Type = 0
	// A resource creation operation.
	OperationMetadata_CREATE OperationMetadata_Type = 1
	// A resource deletion operation.
	OperationMetadata_DELETE OperationMetadata_Type = 2
	// A resource update operation.
	OperationMetadata_UPDATE OperationMetadata_Type = 3
	// A resource check operation.
	OperationMetadata_CHECK OperationMetadata_Type = 4
	// Saves snapshot of the resource operation.
	OperationMetadata_SAVE_SNAPSHOT OperationMetadata_Type = 5
	// Loads snapshot of the resource operation.
	OperationMetadata_LOAD_SNAPSHOT OperationMetadata_Type = 6
	// Triggers failover of environment's Cloud SQL instance (only for highly
	// resilient environments).
	OperationMetadata_DATABASE_FAILOVER OperationMetadata_Type = 7
)

// Enum value maps for OperationMetadata_Type.
var (
	OperationMetadata_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "CREATE",
		2: "DELETE",
		3: "UPDATE",
		4: "CHECK",
		5: "SAVE_SNAPSHOT",
		6: "LOAD_SNAPSHOT",
		7: "DATABASE_FAILOVER",
	}
	OperationMetadata_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":  0,
		"CREATE":            1,
		"DELETE":            2,
		"UPDATE":            3,
		"CHECK":             4,
		"SAVE_SNAPSHOT":     5,
		"LOAD_SNAPSHOT":     6,
		"DATABASE_FAILOVER": 7,
	}
)

func (x OperationMetadata_Type) Enum() *OperationMetadata_Type {
	p := new(OperationMetadata_Type)
	*p = x
	return p
}

func (x OperationMetadata_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationMetadata_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_orchestration_airflow_service_v1_operations_proto_enumTypes[1].Descriptor()
}

func (OperationMetadata_Type) Type() protoreflect.EnumType {
	return &file_google_cloud_orchestration_airflow_service_v1_operations_proto_enumTypes[1]
}

func (x OperationMetadata_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationMetadata_Type.Descriptor instead.
func (OperationMetadata_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescGZIP(), []int{0, 1}
}

// Metadata describing an operation.
type OperationMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The current operation state.
	State OperationMetadata_State `protobuf:"varint,1,opt,name=state,proto3,enum=google.cloud.orchestration.airflow.service.v1.OperationMetadata_State" json:"state,omitempty"`
	// Output only. The type of operation being performed.
	OperationType OperationMetadata_Type `protobuf:"varint,2,opt,name=operation_type,json=operationType,proto3,enum=google.cloud.orchestration.airflow.service.v1.OperationMetadata_Type" json:"operation_type,omitempty"`
	// Output only. The resource being operated on, as a [relative resource name](
	// /apis/design/resource_names#relative_resource_name).
	Resource string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	// Output only. The UUID of the resource being operated on.
	ResourceUuid string `protobuf:"bytes,4,opt,name=resource_uuid,json=resourceUuid,proto3" json:"resource_uuid,omitempty"`
	// Output only. The time the operation was submitted to the server.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The time when the operation terminated, regardless of its
	// success. This field is unset if the operation is still ongoing.
	EndTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *OperationMetadata) Reset() {
	*x = OperationMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_orchestration_airflow_service_v1_operations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationMetadata) ProtoMessage() {}

func (x *OperationMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_orchestration_airflow_service_v1_operations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationMetadata.ProtoReflect.Descriptor instead.
func (*OperationMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescGZIP(), []int{0}
}

func (x *OperationMetadata) GetState() OperationMetadata_State {
	if x != nil {
		return x.State
	}
	return OperationMetadata_STATE_UNSPECIFIED
}

func (x *OperationMetadata) GetOperationType() OperationMetadata_Type {
	if x != nil {
		return x.OperationType
	}
	return OperationMetadata_TYPE_UNSPECIFIED
}

func (x *OperationMetadata) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *OperationMetadata) GetResourceUuid() string {
	if x != nil {
		return x.ResourceUuid
	}
	return ""
}

func (x *OperationMetadata) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *OperationMetadata) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

var File_google_cloud_orchestration_airflow_service_v1_operations_proto protoreflect.FileDescriptor

var file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x69, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x69, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x88, 0x05, 0x0a, 0x11, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x46, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x61, 0x69, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x6c, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x45, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61, 0x69, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x35, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x67, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55, 0x4c, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x1a, 0x02, 0x10, 0x01,
	0x22, 0x88, 0x01, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54,
	0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x04, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x41, 0x56, 0x45, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x10,
	0x05, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48,
	0x4f, 0x54, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x41, 0x54, 0x41, 0x42, 0x41, 0x53, 0x45,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x07, 0x42, 0x93, 0x01, 0x0a, 0x31,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x61,
	0x69, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0f, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x69, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescOnce sync.Once
	file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescData = file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDesc
)

func file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescGZIP() []byte {
	file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescOnce.Do(func() {
		file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescData)
	})
	return file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDescData
}

var file_google_cloud_orchestration_airflow_service_v1_operations_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_orchestration_airflow_service_v1_operations_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_orchestration_airflow_service_v1_operations_proto_goTypes = []interface{}{
	(OperationMetadata_State)(0),  // 0: google.cloud.orchestration.airflow.service.v1.OperationMetadata.State
	(OperationMetadata_Type)(0),   // 1: google.cloud.orchestration.airflow.service.v1.OperationMetadata.Type
	(*OperationMetadata)(nil),     // 2: google.cloud.orchestration.airflow.service.v1.OperationMetadata
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_google_cloud_orchestration_airflow_service_v1_operations_proto_depIdxs = []int32{
	0, // 0: google.cloud.orchestration.airflow.service.v1.OperationMetadata.state:type_name -> google.cloud.orchestration.airflow.service.v1.OperationMetadata.State
	1, // 1: google.cloud.orchestration.airflow.service.v1.OperationMetadata.operation_type:type_name -> google.cloud.orchestration.airflow.service.v1.OperationMetadata.Type
	3, // 2: google.cloud.orchestration.airflow.service.v1.OperationMetadata.create_time:type_name -> google.protobuf.Timestamp
	3, // 3: google.cloud.orchestration.airflow.service.v1.OperationMetadata.end_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_orchestration_airflow_service_v1_operations_proto_init() }
func file_google_cloud_orchestration_airflow_service_v1_operations_proto_init() {
	if File_google_cloud_orchestration_airflow_service_v1_operations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_orchestration_airflow_service_v1_operations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationMetadata); i {
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
			RawDescriptor: file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_orchestration_airflow_service_v1_operations_proto_goTypes,
		DependencyIndexes: file_google_cloud_orchestration_airflow_service_v1_operations_proto_depIdxs,
		EnumInfos:         file_google_cloud_orchestration_airflow_service_v1_operations_proto_enumTypes,
		MessageInfos:      file_google_cloud_orchestration_airflow_service_v1_operations_proto_msgTypes,
	}.Build()
	File_google_cloud_orchestration_airflow_service_v1_operations_proto = out.File
	file_google_cloud_orchestration_airflow_service_v1_operations_proto_rawDesc = nil
	file_google_cloud_orchestration_airflow_service_v1_operations_proto_goTypes = nil
	file_google_cloud_orchestration_airflow_service_v1_operations_proto_depIdxs = nil
}
