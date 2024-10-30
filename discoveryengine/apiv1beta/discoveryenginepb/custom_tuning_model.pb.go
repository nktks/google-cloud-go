// Copyright 2024 Google LLC
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
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/discoveryengine/v1beta/custom_tuning_model.proto

package discoveryenginepb

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

// The state of the model.
type CustomTuningModel_ModelState int32

const (
	// Default value.
	CustomTuningModel_MODEL_STATE_UNSPECIFIED CustomTuningModel_ModelState = 0
	// The model is in a paused training state.
	CustomTuningModel_TRAINING_PAUSED CustomTuningModel_ModelState = 1
	// The model is currently training.
	CustomTuningModel_TRAINING CustomTuningModel_ModelState = 2
	// The model has successfully completed training.
	CustomTuningModel_TRAINING_COMPLETE CustomTuningModel_ModelState = 3
	// The model is ready for serving.
	CustomTuningModel_READY_FOR_SERVING CustomTuningModel_ModelState = 4
	// The model training failed.
	CustomTuningModel_TRAINING_FAILED CustomTuningModel_ModelState = 5
	// The model training finished successfully but metrics did not improve.
	CustomTuningModel_NO_IMPROVEMENT CustomTuningModel_ModelState = 6
	// Input data validation failed. Model training didn't start.
	CustomTuningModel_INPUT_VALIDATION_FAILED CustomTuningModel_ModelState = 7
)

// Enum value maps for CustomTuningModel_ModelState.
var (
	CustomTuningModel_ModelState_name = map[int32]string{
		0: "MODEL_STATE_UNSPECIFIED",
		1: "TRAINING_PAUSED",
		2: "TRAINING",
		3: "TRAINING_COMPLETE",
		4: "READY_FOR_SERVING",
		5: "TRAINING_FAILED",
		6: "NO_IMPROVEMENT",
		7: "INPUT_VALIDATION_FAILED",
	}
	CustomTuningModel_ModelState_value = map[string]int32{
		"MODEL_STATE_UNSPECIFIED": 0,
		"TRAINING_PAUSED":         1,
		"TRAINING":                2,
		"TRAINING_COMPLETE":       3,
		"READY_FOR_SERVING":       4,
		"TRAINING_FAILED":         5,
		"NO_IMPROVEMENT":          6,
		"INPUT_VALIDATION_FAILED": 7,
	}
)

func (x CustomTuningModel_ModelState) Enum() *CustomTuningModel_ModelState {
	p := new(CustomTuningModel_ModelState)
	*p = x
	return p
}

func (x CustomTuningModel_ModelState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CustomTuningModel_ModelState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_enumTypes[0].Descriptor()
}

func (CustomTuningModel_ModelState) Type() protoreflect.EnumType {
	return &file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_enumTypes[0]
}

func (x CustomTuningModel_ModelState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CustomTuningModel_ModelState.Descriptor instead.
func (CustomTuningModel_ModelState) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescGZIP(), []int{0, 0}
}

// Metadata that describes a custom tuned model.
type CustomTuningModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The fully qualified resource name of the model.
	//
	// Format:
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/customTuningModels/{custom_tuning_model}`.
	//
	// Model must be an alpha-numerical string with limit of 40 characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The display name of the model.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The version of the model.
	ModelVersion int64 `protobuf:"varint,3,opt,name=model_version,json=modelVersion,proto3" json:"model_version,omitempty"`
	// The state that the model is in (e.g.`TRAINING` or `TRAINING_FAILED`).
	ModelState CustomTuningModel_ModelState `protobuf:"varint,4,opt,name=model_state,json=modelState,proto3,enum=google.cloud.discoveryengine.v1beta.CustomTuningModel_ModelState" json:"model_state,omitempty"`
	// Deprecated: Timestamp the Model was created at.
	//
	// Deprecated: Marked as deprecated in google/cloud/discoveryengine/v1beta/custom_tuning_model.proto.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Timestamp the model training was initiated.
	TrainingStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=training_start_time,json=trainingStartTime,proto3" json:"training_start_time,omitempty"`
	// The metrics of the trained model.
	Metrics map[string]float64 `protobuf:"bytes,7,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	// Currently this is only populated if the model state is
	// `INPUT_VALIDATION_FAILED`.
	ErrorMessage string `protobuf:"bytes,8,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *CustomTuningModel) Reset() {
	*x = CustomTuningModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomTuningModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomTuningModel) ProtoMessage() {}

func (x *CustomTuningModel) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomTuningModel.ProtoReflect.Descriptor instead.
func (*CustomTuningModel) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescGZIP(), []int{0}
}

func (x *CustomTuningModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomTuningModel) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *CustomTuningModel) GetModelVersion() int64 {
	if x != nil {
		return x.ModelVersion
	}
	return 0
}

func (x *CustomTuningModel) GetModelState() CustomTuningModel_ModelState {
	if x != nil {
		return x.ModelState
	}
	return CustomTuningModel_MODEL_STATE_UNSPECIFIED
}

// Deprecated: Marked as deprecated in google/cloud/discoveryengine/v1beta/custom_tuning_model.proto.
func (x *CustomTuningModel) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *CustomTuningModel) GetTrainingStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TrainingStartTime
	}
	return nil
}

func (x *CustomTuningModel) GetMetrics() map[string]float64 {
	if x != nil {
		return x.Metrics
	}
	return nil
}

func (x *CustomTuningModel) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x08, 0x0a, 0x11, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x4c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x38, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x32, 0x0a, 0x30, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a,
	0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x4a, 0x0a, 0x13, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5d,
	0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x43, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc0,
	0x01, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a,
	0x17, 0x4d, 0x4f, 0x44, 0x45, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x52,
	0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x0c, 0x0a, 0x08, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x15, 0x0a,
	0x11, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45,
	0x54, 0x45, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x46, 0x4f,
	0x52, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05,
	0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x5f, 0x49, 0x4d, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x4d, 0x45,
	0x4e, 0x54, 0x10, 0x06, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x50, 0x55, 0x54, 0x5f, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x07, 0x3a, 0xa4, 0x02, 0xea, 0x41, 0xa0, 0x02, 0x0a, 0x30, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x68, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x7b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x7d, 0x12, 0x81, 0x01, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x2f, 0x7b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x75, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x7d, 0x42, 0x9d, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x42, 0x16, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70,
	0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45, 0x4e, 0x47,
	0x49, 0x4e, 0x45, 0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0xca, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0xea,
	0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescData = file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_goTypes = []any{
	(CustomTuningModel_ModelState)(0), // 0: google.cloud.discoveryengine.v1beta.CustomTuningModel.ModelState
	(*CustomTuningModel)(nil),         // 1: google.cloud.discoveryengine.v1beta.CustomTuningModel
	nil,                               // 2: google.cloud.discoveryengine.v1beta.CustomTuningModel.MetricsEntry
	(*timestamppb.Timestamp)(nil),     // 3: google.protobuf.Timestamp
}
var file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_depIdxs = []int32{
	0, // 0: google.cloud.discoveryengine.v1beta.CustomTuningModel.model_state:type_name -> google.cloud.discoveryengine.v1beta.CustomTuningModel.ModelState
	3, // 1: google.cloud.discoveryengine.v1beta.CustomTuningModel.create_time:type_name -> google.protobuf.Timestamp
	3, // 2: google.cloud.discoveryengine.v1beta.CustomTuningModel.training_start_time:type_name -> google.protobuf.Timestamp
	2, // 3: google.cloud.discoveryengine.v1beta.CustomTuningModel.metrics:type_name -> google.cloud.discoveryengine.v1beta.CustomTuningModel.MetricsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_init() }
func file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_init() {
	if File_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CustomTuningModel); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_depIdxs,
		EnumInfos:         file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_enumTypes,
		MessageInfos:      file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto = out.File
	file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1beta_custom_tuning_model_proto_depIdxs = nil
}
