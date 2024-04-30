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
// source: google/cloud/alloydb/connectors/v1beta/resources.proto

package connectorspb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AuthType contains all supported authentication types.
type MetadataExchangeRequest_AuthType int32

const (
	// Authentication type is unspecified and DB_NATIVE is used by default
	MetadataExchangeRequest_AUTH_TYPE_UNSPECIFIED MetadataExchangeRequest_AuthType = 0
	// Database native authentication (user/password)
	MetadataExchangeRequest_DB_NATIVE MetadataExchangeRequest_AuthType = 1
	// Automatic IAM authentication
	MetadataExchangeRequest_AUTO_IAM MetadataExchangeRequest_AuthType = 2
)

// Enum value maps for MetadataExchangeRequest_AuthType.
var (
	MetadataExchangeRequest_AuthType_name = map[int32]string{
		0: "AUTH_TYPE_UNSPECIFIED",
		1: "DB_NATIVE",
		2: "AUTO_IAM",
	}
	MetadataExchangeRequest_AuthType_value = map[string]int32{
		"AUTH_TYPE_UNSPECIFIED": 0,
		"DB_NATIVE":             1,
		"AUTO_IAM":              2,
	}
)

func (x MetadataExchangeRequest_AuthType) Enum() *MetadataExchangeRequest_AuthType {
	p := new(MetadataExchangeRequest_AuthType)
	*p = x
	return p
}

func (x MetadataExchangeRequest_AuthType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetadataExchangeRequest_AuthType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_alloydb_connectors_v1beta_resources_proto_enumTypes[0].Descriptor()
}

func (MetadataExchangeRequest_AuthType) Type() protoreflect.EnumType {
	return &file_google_cloud_alloydb_connectors_v1beta_resources_proto_enumTypes[0]
}

func (x MetadataExchangeRequest_AuthType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetadataExchangeRequest_AuthType.Descriptor instead.
func (MetadataExchangeRequest_AuthType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescGZIP(), []int{0, 0}
}

// Response code.
type MetadataExchangeResponse_ResponseCode int32

const (
	// Unknown response code
	MetadataExchangeResponse_RESPONSE_CODE_UNSPECIFIED MetadataExchangeResponse_ResponseCode = 0
	// Success
	MetadataExchangeResponse_OK MetadataExchangeResponse_ResponseCode = 1
	// Failure
	MetadataExchangeResponse_ERROR MetadataExchangeResponse_ResponseCode = 2
)

// Enum value maps for MetadataExchangeResponse_ResponseCode.
var (
	MetadataExchangeResponse_ResponseCode_name = map[int32]string{
		0: "RESPONSE_CODE_UNSPECIFIED",
		1: "OK",
		2: "ERROR",
	}
	MetadataExchangeResponse_ResponseCode_value = map[string]int32{
		"RESPONSE_CODE_UNSPECIFIED": 0,
		"OK":                        1,
		"ERROR":                     2,
	}
)

func (x MetadataExchangeResponse_ResponseCode) Enum() *MetadataExchangeResponse_ResponseCode {
	p := new(MetadataExchangeResponse_ResponseCode)
	*p = x
	return p
}

func (x MetadataExchangeResponse_ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetadataExchangeResponse_ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_alloydb_connectors_v1beta_resources_proto_enumTypes[1].Descriptor()
}

func (MetadataExchangeResponse_ResponseCode) Type() protoreflect.EnumType {
	return &file_google_cloud_alloydb_connectors_v1beta_resources_proto_enumTypes[1]
}

func (x MetadataExchangeResponse_ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetadataExchangeResponse_ResponseCode.Descriptor instead.
func (MetadataExchangeResponse_ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescGZIP(), []int{1, 0}
}

// Message used by AlloyDB connectors to exchange client and connection metadata
// with the server after a successful TLS handshake. This metadata includes an
// IAM token, which is used to authenticate users based on their IAM identity.
// The sole purpose of this message is for the use of AlloyDB connectors.
// Clients should not rely on this message directly as there can be breaking
// changes in the future.
type MetadataExchangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Connector information.
	UserAgent string `protobuf:"bytes,1,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// Authentication type.
	AuthType MetadataExchangeRequest_AuthType `protobuf:"varint,2,opt,name=auth_type,json=authType,proto3,enum=google.cloud.alloydb.connectors.v1beta.MetadataExchangeRequest_AuthType" json:"auth_type,omitempty"`
	// IAM token used for both IAM user authentiation and
	// `alloydb.instances.connect` permission check.
	Oauth2Token string `protobuf:"bytes,3,opt,name=oauth2_token,json=oauth2Token,proto3" json:"oauth2_token,omitempty"`
}

func (x *MetadataExchangeRequest) Reset() {
	*x = MetadataExchangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_alloydb_connectors_v1beta_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataExchangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataExchangeRequest) ProtoMessage() {}

func (x *MetadataExchangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_alloydb_connectors_v1beta_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataExchangeRequest.ProtoReflect.Descriptor instead.
func (*MetadataExchangeRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataExchangeRequest) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *MetadataExchangeRequest) GetAuthType() MetadataExchangeRequest_AuthType {
	if x != nil {
		return x.AuthType
	}
	return MetadataExchangeRequest_AUTH_TYPE_UNSPECIFIED
}

func (x *MetadataExchangeRequest) GetOauth2Token() string {
	if x != nil {
		return x.Oauth2Token
	}
	return ""
}

// Message for response to metadata exchange request. The sole purpose of this
// message is for the use of AlloyDB connectors. Clients should not rely on this
// message directly as there can be breaking changes in the future.
type MetadataExchangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Response code.
	ResponseCode MetadataExchangeResponse_ResponseCode `protobuf:"varint,1,opt,name=response_code,json=responseCode,proto3,enum=google.cloud.alloydb.connectors.v1beta.MetadataExchangeResponse_ResponseCode" json:"response_code,omitempty"`
	// Optional. Error message.
	Error string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MetadataExchangeResponse) Reset() {
	*x = MetadataExchangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_alloydb_connectors_v1beta_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataExchangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataExchangeResponse) ProtoMessage() {}

func (x *MetadataExchangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_alloydb_connectors_v1beta_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataExchangeResponse.ProtoReflect.Descriptor instead.
func (*MetadataExchangeResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataExchangeResponse) GetResponseCode() MetadataExchangeResponse_ResponseCode {
	if x != nil {
		return x.ResponseCode
	}
	return MetadataExchangeResponse_RESPONSE_CODE_UNSPECIFIED
}

func (x *MetadataExchangeResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_google_cloud_alloydb_connectors_v1beta_resources_proto protoreflect.FileDescriptor

var file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x17, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x65, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x32, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x08, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x55, 0x54, 0x48, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x42, 0x5f, 0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x49, 0x41, 0x4d, 0x10, 0x02, 0x22,
	0xeb, 0x01, 0x0a, 0x18, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x4d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x40, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52,
	0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x42, 0x89, 0x02,
	0x0a, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x6c, 0x6c, 0x6f, 0x79, 0x64, 0x62, 0x2f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x70, 0x62, 0x3b, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x70, 0x62, 0xaa, 0x02, 0x26, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x44,
	0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x2e, 0x56, 0x31, 0x42,
	0x65, 0x74, 0x61, 0xca, 0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x5c, 0x41, 0x6c, 0x6c, 0x6f, 0x79, 0x44, 0x62, 0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea, 0x02, 0x2a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x6c,
	0x6c, 0x6f, 0x79, 0x44, 0x62, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescOnce sync.Once
	file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescData = file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDesc
)

func file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescGZIP() []byte {
	file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescOnce.Do(func() {
		file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescData)
	})
	return file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDescData
}

var file_google_cloud_alloydb_connectors_v1beta_resources_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_alloydb_connectors_v1beta_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_alloydb_connectors_v1beta_resources_proto_goTypes = []interface{}{
	(MetadataExchangeRequest_AuthType)(0),      // 0: google.cloud.alloydb.connectors.v1beta.MetadataExchangeRequest.AuthType
	(MetadataExchangeResponse_ResponseCode)(0), // 1: google.cloud.alloydb.connectors.v1beta.MetadataExchangeResponse.ResponseCode
	(*MetadataExchangeRequest)(nil),            // 2: google.cloud.alloydb.connectors.v1beta.MetadataExchangeRequest
	(*MetadataExchangeResponse)(nil),           // 3: google.cloud.alloydb.connectors.v1beta.MetadataExchangeResponse
}
var file_google_cloud_alloydb_connectors_v1beta_resources_proto_depIdxs = []int32{
	0, // 0: google.cloud.alloydb.connectors.v1beta.MetadataExchangeRequest.auth_type:type_name -> google.cloud.alloydb.connectors.v1beta.MetadataExchangeRequest.AuthType
	1, // 1: google.cloud.alloydb.connectors.v1beta.MetadataExchangeResponse.response_code:type_name -> google.cloud.alloydb.connectors.v1beta.MetadataExchangeResponse.ResponseCode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_alloydb_connectors_v1beta_resources_proto_init() }
func file_google_cloud_alloydb_connectors_v1beta_resources_proto_init() {
	if File_google_cloud_alloydb_connectors_v1beta_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_alloydb_connectors_v1beta_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataExchangeRequest); i {
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
		file_google_cloud_alloydb_connectors_v1beta_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataExchangeResponse); i {
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
			RawDescriptor: file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_alloydb_connectors_v1beta_resources_proto_goTypes,
		DependencyIndexes: file_google_cloud_alloydb_connectors_v1beta_resources_proto_depIdxs,
		EnumInfos:         file_google_cloud_alloydb_connectors_v1beta_resources_proto_enumTypes,
		MessageInfos:      file_google_cloud_alloydb_connectors_v1beta_resources_proto_msgTypes,
	}.Build()
	File_google_cloud_alloydb_connectors_v1beta_resources_proto = out.File
	file_google_cloud_alloydb_connectors_v1beta_resources_proto_rawDesc = nil
	file_google_cloud_alloydb_connectors_v1beta_resources_proto_goTypes = nil
	file_google_cloud_alloydb_connectors_v1beta_resources_proto_depIdxs = nil
}
