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
// source: google/cloud/cloudcontrolspartner/v1beta/partners.proto

package cloudcontrolspartnerpb

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

// Represents Google Cloud supported external key management partners
// [Google Cloud EKM partners
// docs](https://cloud.google.com/kms/docs/ekm#supported_partners).
type EkmMetadata_EkmSolution int32

const (
	// Unspecified EKM solution
	EkmMetadata_EKM_SOLUTION_UNSPECIFIED EkmMetadata_EkmSolution = 0
	// EKM Partner Fortanix
	EkmMetadata_FORTANIX EkmMetadata_EkmSolution = 1
	// EKM Partner FutureX
	EkmMetadata_FUTUREX EkmMetadata_EkmSolution = 2
	// EKM Partner Thales
	EkmMetadata_THALES EkmMetadata_EkmSolution = 3
	// EKM Partner Virtu
	EkmMetadata_VIRTRU EkmMetadata_EkmSolution = 4
)

// Enum value maps for EkmMetadata_EkmSolution.
var (
	EkmMetadata_EkmSolution_name = map[int32]string{
		0: "EKM_SOLUTION_UNSPECIFIED",
		1: "FORTANIX",
		2: "FUTUREX",
		3: "THALES",
		4: "VIRTRU",
	}
	EkmMetadata_EkmSolution_value = map[string]int32{
		"EKM_SOLUTION_UNSPECIFIED": 0,
		"FORTANIX":                 1,
		"FUTUREX":                  2,
		"THALES":                   3,
		"VIRTRU":                   4,
	}
)

func (x EkmMetadata_EkmSolution) Enum() *EkmMetadata_EkmSolution {
	p := new(EkmMetadata_EkmSolution)
	*p = x
	return p
}

func (x EkmMetadata_EkmSolution) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EkmMetadata_EkmSolution) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_enumTypes[0].Descriptor()
}

func (EkmMetadata_EkmSolution) Type() protoreflect.EnumType {
	return &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_enumTypes[0]
}

func (x EkmMetadata_EkmSolution) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EkmMetadata_EkmSolution.Descriptor instead.
func (EkmMetadata_EkmSolution) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescGZIP(), []int{3, 0}
}

// Message describing Partner resource
type Partner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier. The resource name of the partner.
	// Format: organizations/{organization}/locations/{location}/partner
	// Example: "organizations/123456/locations/us-central1/partner"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of SKUs the partner is offering
	Skus []*Sku `protobuf:"bytes,3,rep,name=skus,proto3" json:"skus,omitempty"`
	// List of Google Cloud supported EKM partners supported by the partner
	EkmSolutions []*EkmMetadata `protobuf:"bytes,4,rep,name=ekm_solutions,json=ekmSolutions,proto3" json:"ekm_solutions,omitempty"`
	// List of Google Cloud regions that the partner sells services to customers.
	// Valid Google Cloud regions found here:
	// https://cloud.google.com/compute/docs/regions-zones
	OperatedCloudRegions []string `protobuf:"bytes,5,rep,name=operated_cloud_regions,json=operatedCloudRegions,proto3" json:"operated_cloud_regions,omitempty"`
	// Google Cloud project ID in the partner's Google Cloud organization for
	// receiving enhanced Logs for Partners.
	PartnerProjectId string `protobuf:"bytes,7,opt,name=partner_project_id,json=partnerProjectId,proto3" json:"partner_project_id,omitempty"`
	// Output only. Time the resource was created
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The last time the resource was updated
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Partner) Reset() {
	*x = Partner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Partner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partner) ProtoMessage() {}

func (x *Partner) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partner.ProtoReflect.Descriptor instead.
func (*Partner) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescGZIP(), []int{0}
}

func (x *Partner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Partner) GetSkus() []*Sku {
	if x != nil {
		return x.Skus
	}
	return nil
}

func (x *Partner) GetEkmSolutions() []*EkmMetadata {
	if x != nil {
		return x.EkmSolutions
	}
	return nil
}

func (x *Partner) GetOperatedCloudRegions() []string {
	if x != nil {
		return x.OperatedCloudRegions
	}
	return nil
}

func (x *Partner) GetPartnerProjectId() string {
	if x != nil {
		return x.PartnerProjectId
	}
	return ""
}

func (x *Partner) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Partner) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// Message for getting a Partner
type GetPartnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Format: organizations/{organization}/locations/{location}/partner
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPartnerRequest) Reset() {
	*x = GetPartnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerRequest) ProtoMessage() {}

func (x *GetPartnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerRequest.ProtoReflect.Descriptor instead.
func (*GetPartnerRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescGZIP(), []int{1}
}

func (x *GetPartnerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Represents the SKU a partner owns inside Google Cloud to sell to customers.
type Sku struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Argentum product SKU, that is associated with the partner offerings to
	// customers used by Syntro for billing purposes. SKUs can represent resold
	// Google products or support services.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display name of the product identified by the SKU. A partner may want to
	// show partner branded names for their offerings such as local sovereign
	// cloud solutions.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *Sku) Reset() {
	*x = Sku{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sku) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sku) ProtoMessage() {}

func (x *Sku) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sku.ProtoReflect.Descriptor instead.
func (*Sku) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescGZIP(), []int{2}
}

func (x *Sku) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sku) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

// Holds information needed by Mudbray to use partner EKMs for workloads.
type EkmMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Cloud EKM partner.
	EkmSolution EkmMetadata_EkmSolution `protobuf:"varint,1,opt,name=ekm_solution,json=ekmSolution,proto3,enum=google.cloud.cloudcontrolspartner.v1beta.EkmMetadata_EkmSolution" json:"ekm_solution,omitempty"`
	// Endpoint for sending requests to the EKM for key provisioning during
	// Assured Workload creation.
	EkmEndpointUri string `protobuf:"bytes,2,opt,name=ekm_endpoint_uri,json=ekmEndpointUri,proto3" json:"ekm_endpoint_uri,omitempty"`
}

func (x *EkmMetadata) Reset() {
	*x = EkmMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EkmMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EkmMetadata) ProtoMessage() {}

func (x *EkmMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EkmMetadata.ProtoReflect.Descriptor instead.
func (*EkmMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescGZIP(), []int{3}
}

func (x *EkmMetadata) GetEkmSolution() EkmMetadata_EkmSolution {
	if x != nil {
		return x.EkmSolution
	}
	return EkmMetadata_EKM_SOLUTION_UNSPECIFIED
}

func (x *EkmMetadata) GetEkmEndpointUri() string {
	if x != nil {
		return x.EkmEndpointUri
	}
	return ""
}

var File_google_cloud_cloudcontrolspartner_v1beta_partners_proto protoreflect.FileDescriptor

var file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9f, 0x04, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x08, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x53,
	0x6b, 0x75, 0x52, 0x04, 0x73, 0x6b, 0x75, 0x73, 0x12, 0x5a, 0x0a, 0x0d, 0x65, 0x6b, 0x6d, 0x5f,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x45, 0x6b, 0x6d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x65, 0x6b, 0x6d, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x0a, 0x16, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x14, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x74, 0xea, 0x41,
	0x71, 0x0a, 0x2b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x39,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x7d, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x32, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x22, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x33, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2d, 0x0a, 0x2b, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x38, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xfd, 0x01, 0x0a, 0x0b, 0x45,
	0x6b, 0x6d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x64, 0x0a, 0x0c, 0x65, 0x6b,
	0x6d, 0x5f, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x45, 0x6b, 0x6d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x6b, 0x6d, 0x53, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x65, 0x6b, 0x6d, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x10, 0x65, 0x6b, 0x6d, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6b, 0x6d, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x72, 0x69, 0x22, 0x5e, 0x0a, 0x0b, 0x45, 0x6b,
	0x6d, 0x53, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x4b, 0x4d,
	0x5f, 0x53, 0x4f, 0x4c, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4f, 0x52, 0x54, 0x41,
	0x4e, 0x49, 0x58, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x55, 0x54, 0x55, 0x52, 0x45, 0x58,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x48, 0x41, 0x4c, 0x45, 0x53, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x56, 0x49, 0x52, 0x54, 0x52, 0x55, 0x10, 0x04, 0x42, 0xa5, 0x02, 0x0a, 0x2c, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0d, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x60, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x73, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x70, 0x62, 0xaa, 0x02,
	0x28, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x28, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x2b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x73, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescOnce sync.Once
	file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescData = file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDesc
)

func file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescGZIP() []byte {
	file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescOnce.Do(func() {
		file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescData)
	})
	return file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDescData
}

var file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_goTypes = []interface{}{
	(EkmMetadata_EkmSolution)(0),  // 0: google.cloud.cloudcontrolspartner.v1beta.EkmMetadata.EkmSolution
	(*Partner)(nil),               // 1: google.cloud.cloudcontrolspartner.v1beta.Partner
	(*GetPartnerRequest)(nil),     // 2: google.cloud.cloudcontrolspartner.v1beta.GetPartnerRequest
	(*Sku)(nil),                   // 3: google.cloud.cloudcontrolspartner.v1beta.Sku
	(*EkmMetadata)(nil),           // 4: google.cloud.cloudcontrolspartner.v1beta.EkmMetadata
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_depIdxs = []int32{
	3, // 0: google.cloud.cloudcontrolspartner.v1beta.Partner.skus:type_name -> google.cloud.cloudcontrolspartner.v1beta.Sku
	4, // 1: google.cloud.cloudcontrolspartner.v1beta.Partner.ekm_solutions:type_name -> google.cloud.cloudcontrolspartner.v1beta.EkmMetadata
	5, // 2: google.cloud.cloudcontrolspartner.v1beta.Partner.create_time:type_name -> google.protobuf.Timestamp
	5, // 3: google.cloud.cloudcontrolspartner.v1beta.Partner.update_time:type_name -> google.protobuf.Timestamp
	0, // 4: google.cloud.cloudcontrolspartner.v1beta.EkmMetadata.ekm_solution:type_name -> google.cloud.cloudcontrolspartner.v1beta.EkmMetadata.EkmSolution
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_init() }
func file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_init() {
	if File_google_cloud_cloudcontrolspartner_v1beta_partners_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Partner); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartnerRequest); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sku); i {
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
		file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EkmMetadata); i {
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
			RawDescriptor: file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_goTypes,
		DependencyIndexes: file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_depIdxs,
		EnumInfos:         file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_enumTypes,
		MessageInfos:      file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_msgTypes,
	}.Build()
	File_google_cloud_cloudcontrolspartner_v1beta_partners_proto = out.File
	file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_rawDesc = nil
	file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_goTypes = nil
	file_google_cloud_cloudcontrolspartner_v1beta_partners_proto_depIdxs = nil
}
