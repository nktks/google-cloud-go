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
// source: google/cloud/securitycenter/v1/access.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents an access event.
type Access struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Associated email, such as "foo@google.com".
	//
	// The email address of the authenticated user or a service account acting on
	// behalf of a third party principal making the request. For third party
	// identity callers, the `principal_subject` field is populated instead of
	// this field. For privacy reasons, the principal email address is sometimes
	// redacted. For more information, see [Caller identities in audit
	// logs](https://cloud.google.com/logging/docs/audit#user-id).
	PrincipalEmail string `protobuf:"bytes,1,opt,name=principal_email,json=principalEmail,proto3" json:"principal_email,omitempty"`
	// Caller's IP address, such as "1.1.1.1".
	CallerIp string `protobuf:"bytes,2,opt,name=caller_ip,json=callerIp,proto3" json:"caller_ip,omitempty"`
	// The caller IP's geolocation, which identifies where the call came from.
	CallerIpGeo *Geolocation `protobuf:"bytes,3,opt,name=caller_ip_geo,json=callerIpGeo,proto3" json:"caller_ip_geo,omitempty"`
	// Type of user agent associated with the finding. For example, an operating
	// system shell or an embedded or standalone application.
	UserAgentFamily string `protobuf:"bytes,4,opt,name=user_agent_family,json=userAgentFamily,proto3" json:"user_agent_family,omitempty"`
	// The caller's user agent string associated with the finding.
	UserAgent string `protobuf:"bytes,12,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// This is the API service that the service account made a call to, e.g.
	// "iam.googleapis.com"
	ServiceName string `protobuf:"bytes,5,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The method that the service account called, e.g. "SetIamPolicy".
	MethodName string `protobuf:"bytes,6,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// A string that represents the principal_subject that is associated with the
	// identity. Unlike `principal_email`, `principal_subject` supports principals
	// that aren't associated with email addresses, such as third party
	// principals. For most identities, the format is
	// `principal://iam.googleapis.com/{identity pool name}/subject/{subject}`.
	// Some GKE identities, such as GKE_WORKLOAD, FREEFORM, and GKE_HUB_WORKLOAD,
	// still use the legacy format `serviceAccount:{identity pool
	// name}[{subject}]`.
	PrincipalSubject string `protobuf:"bytes,7,opt,name=principal_subject,json=principalSubject,proto3" json:"principal_subject,omitempty"`
	// The name of the service account key that was used to create or exchange
	// credentials when authenticating the service account that made the request.
	// This is a scheme-less URI full resource name. For example:
	//
	// "//iam.googleapis.com/projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{key}".
	ServiceAccountKeyName string `protobuf:"bytes,8,opt,name=service_account_key_name,json=serviceAccountKeyName,proto3" json:"service_account_key_name,omitempty"`
	// The identity delegation history of an authenticated service account that
	// made the request. The `serviceAccountDelegationInfo[]` object contains
	// information about the real authorities that try to access Google Cloud
	// resources by delegating on a service account. When multiple authorities are
	// present, they are guaranteed to be sorted based on the original ordering of
	// the identity delegation events.
	ServiceAccountDelegationInfo []*ServiceAccountDelegationInfo `protobuf:"bytes,9,rep,name=service_account_delegation_info,json=serviceAccountDelegationInfo,proto3" json:"service_account_delegation_info,omitempty"`
	// A string that represents a username. The username provided depends on the
	// type of the finding and is likely not an IAM principal. For example, this
	// can be a system username if the finding is related to a virtual machine, or
	// it can be an application login username.
	UserName string `protobuf:"bytes,11,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *Access) Reset() {
	*x = Access{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_access_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Access) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access) ProtoMessage() {}

func (x *Access) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_access_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access.ProtoReflect.Descriptor instead.
func (*Access) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_access_proto_rawDescGZIP(), []int{0}
}

func (x *Access) GetPrincipalEmail() string {
	if x != nil {
		return x.PrincipalEmail
	}
	return ""
}

func (x *Access) GetCallerIp() string {
	if x != nil {
		return x.CallerIp
	}
	return ""
}

func (x *Access) GetCallerIpGeo() *Geolocation {
	if x != nil {
		return x.CallerIpGeo
	}
	return nil
}

func (x *Access) GetUserAgentFamily() string {
	if x != nil {
		return x.UserAgentFamily
	}
	return ""
}

func (x *Access) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

func (x *Access) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Access) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *Access) GetPrincipalSubject() string {
	if x != nil {
		return x.PrincipalSubject
	}
	return ""
}

func (x *Access) GetServiceAccountKeyName() string {
	if x != nil {
		return x.ServiceAccountKeyName
	}
	return ""
}

func (x *Access) GetServiceAccountDelegationInfo() []*ServiceAccountDelegationInfo {
	if x != nil {
		return x.ServiceAccountDelegationInfo
	}
	return nil
}

func (x *Access) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

// Identity delegation history of an authenticated service account.
type ServiceAccountDelegationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The email address of a Google account.
	PrincipalEmail string `protobuf:"bytes,1,opt,name=principal_email,json=principalEmail,proto3" json:"principal_email,omitempty"`
	// A string representing the principal_subject associated with the identity.
	// As compared to `principal_email`, supports principals that aren't
	// associated with email addresses, such as third party principals. For most
	// identities, the format will be `principal://iam.googleapis.com/{identity
	// pool name}/subjects/{subject}` except for some GKE identities
	// (GKE_WORKLOAD, FREEFORM, GKE_HUB_WORKLOAD) that are still in the legacy
	// format `serviceAccount:{identity pool name}[{subject}]`
	PrincipalSubject string `protobuf:"bytes,2,opt,name=principal_subject,json=principalSubject,proto3" json:"principal_subject,omitempty"`
}

func (x *ServiceAccountDelegationInfo) Reset() {
	*x = ServiceAccountDelegationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_access_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceAccountDelegationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceAccountDelegationInfo) ProtoMessage() {}

func (x *ServiceAccountDelegationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_access_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceAccountDelegationInfo.ProtoReflect.Descriptor instead.
func (*ServiceAccountDelegationInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_access_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceAccountDelegationInfo) GetPrincipalEmail() string {
	if x != nil {
		return x.PrincipalEmail
	}
	return ""
}

func (x *ServiceAccountDelegationInfo) GetPrincipalSubject() string {
	if x != nil {
		return x.PrincipalSubject
	}
	return ""
}

// Represents a geographical location for a given access.
type Geolocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A CLDR.
	RegionCode string `protobuf:"bytes,1,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
}

func (x *Geolocation) Reset() {
	*x = Geolocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_access_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Geolocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geolocation) ProtoMessage() {}

func (x *Geolocation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_access_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geolocation.ProtoReflect.Descriptor instead.
func (*Geolocation) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_access_proto_rawDescGZIP(), []int{2}
}

func (x *Geolocation) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

var File_google_cloud_securitycenter_v1_access_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_access_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xb7, 0x04,
	0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x6e,
	0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x70, 0x12, 0x4f,
	0x0a, 0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x5f, 0x67, 0x65, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x70, 0x47, 0x65, 0x6f, 0x12,
	0x2a, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x37, 0x0a, 0x18, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x83, 0x01, 0x0a, 0x1f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x1c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x74, 0x0a, 0x1c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2e, 0x0a,
	0x0b, 0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x42, 0xe5, 0x01,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0xaa,
	0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_access_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_access_proto_rawDescData = file_google_cloud_securitycenter_v1_access_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_access_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_access_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_access_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_access_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_access_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_access_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_securitycenter_v1_access_proto_goTypes = []interface{}{
	(*Access)(nil),                       // 0: google.cloud.securitycenter.v1.Access
	(*ServiceAccountDelegationInfo)(nil), // 1: google.cloud.securitycenter.v1.ServiceAccountDelegationInfo
	(*Geolocation)(nil),                  // 2: google.cloud.securitycenter.v1.Geolocation
}
var file_google_cloud_securitycenter_v1_access_proto_depIdxs = []int32{
	2, // 0: google.cloud.securitycenter.v1.Access.caller_ip_geo:type_name -> google.cloud.securitycenter.v1.Geolocation
	1, // 1: google.cloud.securitycenter.v1.Access.service_account_delegation_info:type_name -> google.cloud.securitycenter.v1.ServiceAccountDelegationInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_access_proto_init() }
func file_google_cloud_securitycenter_v1_access_proto_init() {
	if File_google_cloud_securitycenter_v1_access_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_access_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Access); i {
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
		file_google_cloud_securitycenter_v1_access_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceAccountDelegationInfo); i {
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
		file_google_cloud_securitycenter_v1_access_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Geolocation); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_access_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_access_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_access_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_access_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_access_proto = out.File
	file_google_cloud_securitycenter_v1_access_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_access_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_access_proto_depIdxs = nil
}
