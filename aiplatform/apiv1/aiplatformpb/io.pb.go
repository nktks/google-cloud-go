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
// source: google/cloud/aiplatform/v1/io.proto

package aiplatformpb

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

// The storage details for Avro input content.
type AvroSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage location.
	GcsSource *GcsSource `protobuf:"bytes,1,opt,name=gcs_source,json=gcsSource,proto3" json:"gcs_source,omitempty"`
}

func (x *AvroSource) Reset() {
	*x = AvroSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvroSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvroSource) ProtoMessage() {}

func (x *AvroSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvroSource.ProtoReflect.Descriptor instead.
func (*AvroSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{0}
}

func (x *AvroSource) GetGcsSource() *GcsSource {
	if x != nil {
		return x.GcsSource
	}
	return nil
}

// The storage details for CSV input content.
type CsvSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage location.
	GcsSource *GcsSource `protobuf:"bytes,1,opt,name=gcs_source,json=gcsSource,proto3" json:"gcs_source,omitempty"`
}

func (x *CsvSource) Reset() {
	*x = CsvSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsvSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsvSource) ProtoMessage() {}

func (x *CsvSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsvSource.ProtoReflect.Descriptor instead.
func (*CsvSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{1}
}

func (x *CsvSource) GetGcsSource() *GcsSource {
	if x != nil {
		return x.GcsSource
	}
	return nil
}

// The Google Cloud Storage location for the input content.
type GcsSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage URI(-s) to the input file(s). May contain
	// wildcards. For more information on wildcards, see
	// https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames.
	Uris []string `protobuf:"bytes,1,rep,name=uris,proto3" json:"uris,omitempty"`
}

func (x *GcsSource) Reset() {
	*x = GcsSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsSource) ProtoMessage() {}

func (x *GcsSource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsSource.ProtoReflect.Descriptor instead.
func (*GcsSource) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{2}
}

func (x *GcsSource) GetUris() []string {
	if x != nil {
		return x.Uris
	}
	return nil
}

// The Google Cloud Storage location where the output is to be written to.
type GcsDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage URI to output directory. If the uri doesn't
	// end with
	// '/', a '/' will be automatically appended. The directory is created if it
	// doesn't exist.
	OutputUriPrefix string `protobuf:"bytes,1,opt,name=output_uri_prefix,json=outputUriPrefix,proto3" json:"output_uri_prefix,omitempty"`
}

func (x *GcsDestination) Reset() {
	*x = GcsDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsDestination) ProtoMessage() {}

func (x *GcsDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsDestination.ProtoReflect.Descriptor instead.
func (*GcsDestination) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{3}
}

func (x *GcsDestination) GetOutputUriPrefix() string {
	if x != nil {
		return x.OutputUriPrefix
	}
	return ""
}

// The BigQuery location for the input content.
type BigQuerySource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. BigQuery URI to a table, up to 2000 characters long.
	// Accepted forms:
	//
	// *  BigQuery path. For example: `bq://projectId.bqDatasetId.bqTableId`.
	InputUri string `protobuf:"bytes,1,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
}

func (x *BigQuerySource) Reset() {
	*x = BigQuerySource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQuerySource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQuerySource) ProtoMessage() {}

func (x *BigQuerySource) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQuerySource.ProtoReflect.Descriptor instead.
func (*BigQuerySource) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{4}
}

func (x *BigQuerySource) GetInputUri() string {
	if x != nil {
		return x.InputUri
	}
	return ""
}

// The BigQuery location for the output content.
type BigQueryDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. BigQuery URI to a project or table, up to 2000 characters long.
	//
	// When only the project is specified, the Dataset and Table is created.
	// When the full table reference is specified, the Dataset must exist and
	// table must not exist.
	//
	// Accepted forms:
	//
	// *  BigQuery path. For example:
	// `bq://projectId` or `bq://projectId.bqDatasetId` or
	// `bq://projectId.bqDatasetId.bqTableId`.
	OutputUri string `protobuf:"bytes,1,opt,name=output_uri,json=outputUri,proto3" json:"output_uri,omitempty"`
}

func (x *BigQueryDestination) Reset() {
	*x = BigQueryDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryDestination) ProtoMessage() {}

func (x *BigQueryDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryDestination.ProtoReflect.Descriptor instead.
func (*BigQueryDestination) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{5}
}

func (x *BigQueryDestination) GetOutputUri() string {
	if x != nil {
		return x.OutputUri
	}
	return ""
}

// The storage details for CSV output content.
type CsvDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage location.
	GcsDestination *GcsDestination `protobuf:"bytes,1,opt,name=gcs_destination,json=gcsDestination,proto3" json:"gcs_destination,omitempty"`
}

func (x *CsvDestination) Reset() {
	*x = CsvDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsvDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsvDestination) ProtoMessage() {}

func (x *CsvDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsvDestination.ProtoReflect.Descriptor instead.
func (*CsvDestination) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{6}
}

func (x *CsvDestination) GetGcsDestination() *GcsDestination {
	if x != nil {
		return x.GcsDestination
	}
	return nil
}

// The storage details for TFRecord output content.
type TFRecordDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Google Cloud Storage location.
	GcsDestination *GcsDestination `protobuf:"bytes,1,opt,name=gcs_destination,json=gcsDestination,proto3" json:"gcs_destination,omitempty"`
}

func (x *TFRecordDestination) Reset() {
	*x = TFRecordDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TFRecordDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TFRecordDestination) ProtoMessage() {}

func (x *TFRecordDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TFRecordDestination.ProtoReflect.Descriptor instead.
func (*TFRecordDestination) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{7}
}

func (x *TFRecordDestination) GetGcsDestination() *GcsDestination {
	if x != nil {
		return x.GcsDestination
	}
	return nil
}

// The Container Registry location for the container image.
type ContainerRegistryDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Container Registry URI of a container image.
	// Only Google Container Registry and Artifact Registry are supported now.
	// Accepted forms:
	//
	//   - Google Container Registry path. For example:
	//     `gcr.io/projectId/imageName:tag`.
	//
	//   - Artifact Registry path. For example:
	//     `us-central1-docker.pkg.dev/projectId/repoName/imageName:tag`.
	//
	// If a tag is not specified, "latest" will be used as the default tag.
	OutputUri string `protobuf:"bytes,1,opt,name=output_uri,json=outputUri,proto3" json:"output_uri,omitempty"`
}

func (x *ContainerRegistryDestination) Reset() {
	*x = ContainerRegistryDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContainerRegistryDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerRegistryDestination) ProtoMessage() {}

func (x *ContainerRegistryDestination) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_io_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerRegistryDestination.ProtoReflect.Descriptor instead.
func (*ContainerRegistryDestination) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP(), []int{8}
}

func (x *ContainerRegistryDestination) GetOutputUri() string {
	if x != nil {
		return x.OutputUri
	}
	return ""
}

var File_google_cloud_aiplatform_v1_io_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_io_proto_rawDesc = []byte{
	0x0a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0a, 0x41, 0x76, 0x72, 0x6f, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x49, 0x0a, 0x0a, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x09, 0x67, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x09, 0x43,
	0x73, 0x76, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x67, 0x63, 0x73, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x67, 0x63, 0x73, 0x53, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x22, 0x24, 0x0a, 0x09, 0x47, 0x63, 0x73, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x17, 0x0a, 0x04, 0x75, 0x72, 0x69, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x04, 0x75, 0x72, 0x69, 0x73, 0x22, 0x41, 0x0a, 0x0e, 0x47, 0x63, 0x73,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x11, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x55, 0x72, 0x69, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x32, 0x0a, 0x0e,
	0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20,
	0x0a, 0x09, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x55, 0x72, 0x69,
	0x22, 0x39, 0x0a, 0x13, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x55, 0x72, 0x69, 0x22, 0x6a, 0x0a, 0x0e, 0x43,
	0x73, 0x76, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a,
	0x0f, 0x67, 0x63, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x73, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x67, 0x63, 0x73, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6f, 0x0a, 0x13, 0x54, 0x46, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58,
	0x0a, 0x0f, 0x67, 0x63, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x73, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0e, 0x67, 0x63, 0x73, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0a, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x09, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x55, 0x72, 0x69, 0x42, 0xc5, 0x01, 0x0a,
	0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x07, 0x49, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31,
	0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0x3b, 0x61, 0x69,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_io_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_io_proto_rawDescData = file_google_cloud_aiplatform_v1_io_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_io_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_io_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_io_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_io_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_io_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_io_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_google_cloud_aiplatform_v1_io_proto_goTypes = []interface{}{
	(*AvroSource)(nil),                   // 0: google.cloud.aiplatform.v1.AvroSource
	(*CsvSource)(nil),                    // 1: google.cloud.aiplatform.v1.CsvSource
	(*GcsSource)(nil),                    // 2: google.cloud.aiplatform.v1.GcsSource
	(*GcsDestination)(nil),               // 3: google.cloud.aiplatform.v1.GcsDestination
	(*BigQuerySource)(nil),               // 4: google.cloud.aiplatform.v1.BigQuerySource
	(*BigQueryDestination)(nil),          // 5: google.cloud.aiplatform.v1.BigQueryDestination
	(*CsvDestination)(nil),               // 6: google.cloud.aiplatform.v1.CsvDestination
	(*TFRecordDestination)(nil),          // 7: google.cloud.aiplatform.v1.TFRecordDestination
	(*ContainerRegistryDestination)(nil), // 8: google.cloud.aiplatform.v1.ContainerRegistryDestination
}
var file_google_cloud_aiplatform_v1_io_proto_depIdxs = []int32{
	2, // 0: google.cloud.aiplatform.v1.AvroSource.gcs_source:type_name -> google.cloud.aiplatform.v1.GcsSource
	2, // 1: google.cloud.aiplatform.v1.CsvSource.gcs_source:type_name -> google.cloud.aiplatform.v1.GcsSource
	3, // 2: google.cloud.aiplatform.v1.CsvDestination.gcs_destination:type_name -> google.cloud.aiplatform.v1.GcsDestination
	3, // 3: google.cloud.aiplatform.v1.TFRecordDestination.gcs_destination:type_name -> google.cloud.aiplatform.v1.GcsDestination
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1_io_proto_init() }
func file_google_cloud_aiplatform_v1_io_proto_init() {
	if File_google_cloud_aiplatform_v1_io_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvroSource); i {
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
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsvSource); i {
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
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsSource); i {
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
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsDestination); i {
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
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQuerySource); i {
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
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryDestination); i {
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
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsvDestination); i {
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
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TFRecordDestination); i {
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
		file_google_cloud_aiplatform_v1_io_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContainerRegistryDestination); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_io_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_io_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_io_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_io_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_io_proto = out.File
	file_google_cloud_aiplatform_v1_io_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_io_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_io_proto_depIdxs = nil
}
