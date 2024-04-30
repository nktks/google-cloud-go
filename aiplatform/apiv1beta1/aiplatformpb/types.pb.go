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
// source: google/cloud/aiplatform/v1beta1/types.proto

package aiplatformpb

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

// Data type of the tensor.
type Tensor_DataType int32

const (
	// Not a legal value for DataType. Used to indicate a DataType field has not
	// been set.
	Tensor_DATA_TYPE_UNSPECIFIED Tensor_DataType = 0
	// Data types that all computation devices are expected to be
	// capable to support.
	Tensor_BOOL   Tensor_DataType = 1
	Tensor_STRING Tensor_DataType = 2
	Tensor_FLOAT  Tensor_DataType = 3
	Tensor_DOUBLE Tensor_DataType = 4
	Tensor_INT8   Tensor_DataType = 5
	Tensor_INT16  Tensor_DataType = 6
	Tensor_INT32  Tensor_DataType = 7
	Tensor_INT64  Tensor_DataType = 8
	Tensor_UINT8  Tensor_DataType = 9
	Tensor_UINT16 Tensor_DataType = 10
	Tensor_UINT32 Tensor_DataType = 11
	Tensor_UINT64 Tensor_DataType = 12
)

// Enum value maps for Tensor_DataType.
var (
	Tensor_DataType_name = map[int32]string{
		0:  "DATA_TYPE_UNSPECIFIED",
		1:  "BOOL",
		2:  "STRING",
		3:  "FLOAT",
		4:  "DOUBLE",
		5:  "INT8",
		6:  "INT16",
		7:  "INT32",
		8:  "INT64",
		9:  "UINT8",
		10: "UINT16",
		11: "UINT32",
		12: "UINT64",
	}
	Tensor_DataType_value = map[string]int32{
		"DATA_TYPE_UNSPECIFIED": 0,
		"BOOL":                  1,
		"STRING":                2,
		"FLOAT":                 3,
		"DOUBLE":                4,
		"INT8":                  5,
		"INT16":                 6,
		"INT32":                 7,
		"INT64":                 8,
		"UINT8":                 9,
		"UINT16":                10,
		"UINT32":                11,
		"UINT64":                12,
	}
)

func (x Tensor_DataType) Enum() *Tensor_DataType {
	p := new(Tensor_DataType)
	*p = x
	return p
}

func (x Tensor_DataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Tensor_DataType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_aiplatform_v1beta1_types_proto_enumTypes[0].Descriptor()
}

func (Tensor_DataType) Type() protoreflect.EnumType {
	return &file_google_cloud_aiplatform_v1beta1_types_proto_enumTypes[0]
}

func (x Tensor_DataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Tensor_DataType.Descriptor instead.
func (Tensor_DataType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_types_proto_rawDescGZIP(), []int{4, 0}
}

// A list of boolean values.
type BoolArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of bool values.
	Values []bool `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *BoolArray) Reset() {
	*x = BoolArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoolArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoolArray) ProtoMessage() {}

func (x *BoolArray) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoolArray.ProtoReflect.Descriptor instead.
func (*BoolArray) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_types_proto_rawDescGZIP(), []int{0}
}

func (x *BoolArray) GetValues() []bool {
	if x != nil {
		return x.Values
	}
	return nil
}

// A list of double values.
type DoubleArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of double values.
	Values []float64 `protobuf:"fixed64,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *DoubleArray) Reset() {
	*x = DoubleArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoubleArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoubleArray) ProtoMessage() {}

func (x *DoubleArray) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoubleArray.ProtoReflect.Descriptor instead.
func (*DoubleArray) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_types_proto_rawDescGZIP(), []int{1}
}

func (x *DoubleArray) GetValues() []float64 {
	if x != nil {
		return x.Values
	}
	return nil
}

// A list of int64 values.
type Int64Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of int64 values.
	Values []int64 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *Int64Array) Reset() {
	*x = Int64Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int64Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64Array) ProtoMessage() {}

func (x *Int64Array) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int64Array.ProtoReflect.Descriptor instead.
func (*Int64Array) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Int64Array) GetValues() []int64 {
	if x != nil {
		return x.Values
	}
	return nil
}

// A list of string values.
type StringArray struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of string values.
	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *StringArray) Reset() {
	*x = StringArray{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringArray) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringArray) ProtoMessage() {}

func (x *StringArray) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringArray.ProtoReflect.Descriptor instead.
func (*StringArray) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_types_proto_rawDescGZIP(), []int{3}
}

func (x *StringArray) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// A tensor value type.
type Tensor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The data type of tensor.
	Dtype Tensor_DataType `protobuf:"varint,1,opt,name=dtype,proto3,enum=google.cloud.aiplatform.v1beta1.Tensor_DataType" json:"dtype,omitempty"`
	// Shape of the tensor.
	Shape []int64 `protobuf:"varint,2,rep,packed,name=shape,proto3" json:"shape,omitempty"`
	// Type specific representations that make it easy to create tensor protos in
	// all languages.  Only the representation corresponding to "dtype" can
	// be set.  The values hold the flattened representation of the tensor in
	// row major order.
	//
	// [BOOL][google.aiplatform.master.Tensor.DataType.BOOL]
	BoolVal []bool `protobuf:"varint,3,rep,packed,name=bool_val,json=boolVal,proto3" json:"bool_val,omitempty"`
	// [STRING][google.aiplatform.master.Tensor.DataType.STRING]
	StringVal []string `protobuf:"bytes,14,rep,name=string_val,json=stringVal,proto3" json:"string_val,omitempty"`
	// [STRING][google.aiplatform.master.Tensor.DataType.STRING]
	BytesVal [][]byte `protobuf:"bytes,15,rep,name=bytes_val,json=bytesVal,proto3" json:"bytes_val,omitempty"`
	// [FLOAT][google.aiplatform.master.Tensor.DataType.FLOAT]
	FloatVal []float32 `protobuf:"fixed32,5,rep,packed,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	// [DOUBLE][google.aiplatform.master.Tensor.DataType.DOUBLE]
	DoubleVal []float64 `protobuf:"fixed64,6,rep,packed,name=double_val,json=doubleVal,proto3" json:"double_val,omitempty"`
	// [INT_8][google.aiplatform.master.Tensor.DataType.INT8]
	// [INT_16][google.aiplatform.master.Tensor.DataType.INT16]
	// [INT_32][google.aiplatform.master.Tensor.DataType.INT32]
	IntVal []int32 `protobuf:"varint,7,rep,packed,name=int_val,json=intVal,proto3" json:"int_val,omitempty"`
	// [INT64][google.aiplatform.master.Tensor.DataType.INT64]
	Int64Val []int64 `protobuf:"varint,8,rep,packed,name=int64_val,json=int64Val,proto3" json:"int64_val,omitempty"`
	// [UINT8][google.aiplatform.master.Tensor.DataType.UINT8]
	// [UINT16][google.aiplatform.master.Tensor.DataType.UINT16]
	// [UINT32][google.aiplatform.master.Tensor.DataType.UINT32]
	UintVal []uint32 `protobuf:"varint,9,rep,packed,name=uint_val,json=uintVal,proto3" json:"uint_val,omitempty"`
	// [UINT64][google.aiplatform.master.Tensor.DataType.UINT64]
	Uint64Val []uint64 `protobuf:"varint,10,rep,packed,name=uint64_val,json=uint64Val,proto3" json:"uint64_val,omitempty"`
	// A list of tensor values.
	ListVal []*Tensor `protobuf:"bytes,11,rep,name=list_val,json=listVal,proto3" json:"list_val,omitempty"`
	// A map of string to tensor.
	StructVal map[string]*Tensor `protobuf:"bytes,12,rep,name=struct_val,json=structVal,proto3" json:"struct_val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Serialized raw tensor content.
	TensorVal []byte `protobuf:"bytes,13,opt,name=tensor_val,json=tensorVal,proto3" json:"tensor_val,omitempty"`
}

func (x *Tensor) Reset() {
	*x = Tensor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tensor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tensor) ProtoMessage() {}

func (x *Tensor) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tensor.ProtoReflect.Descriptor instead.
func (*Tensor) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1beta1_types_proto_rawDescGZIP(), []int{4}
}

func (x *Tensor) GetDtype() Tensor_DataType {
	if x != nil {
		return x.Dtype
	}
	return Tensor_DATA_TYPE_UNSPECIFIED
}

func (x *Tensor) GetShape() []int64 {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *Tensor) GetBoolVal() []bool {
	if x != nil {
		return x.BoolVal
	}
	return nil
}

func (x *Tensor) GetStringVal() []string {
	if x != nil {
		return x.StringVal
	}
	return nil
}

func (x *Tensor) GetBytesVal() [][]byte {
	if x != nil {
		return x.BytesVal
	}
	return nil
}

func (x *Tensor) GetFloatVal() []float32 {
	if x != nil {
		return x.FloatVal
	}
	return nil
}

func (x *Tensor) GetDoubleVal() []float64 {
	if x != nil {
		return x.DoubleVal
	}
	return nil
}

func (x *Tensor) GetIntVal() []int32 {
	if x != nil {
		return x.IntVal
	}
	return nil
}

func (x *Tensor) GetInt64Val() []int64 {
	if x != nil {
		return x.Int64Val
	}
	return nil
}

func (x *Tensor) GetUintVal() []uint32 {
	if x != nil {
		return x.UintVal
	}
	return nil
}

func (x *Tensor) GetUint64Val() []uint64 {
	if x != nil {
		return x.Uint64Val
	}
	return nil
}

func (x *Tensor) GetListVal() []*Tensor {
	if x != nil {
		return x.ListVal
	}
	return nil
}

func (x *Tensor) GetStructVal() map[string]*Tensor {
	if x != nil {
		return x.StructVal
	}
	return nil
}

func (x *Tensor) GetTensorVal() []byte {
	if x != nil {
		return x.TensorVal
	}
	return nil
}

var File_google_cloud_aiplatform_v1beta1_types_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1beta1_types_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x22, 0x23,
	0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x6c, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x41, 0x72, 0x72,
	0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x01, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x0a, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0x25, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xb9, 0x06, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x12, 0x46, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x05, 0x73, 0x68, 0x61, 0x70, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x08, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x08, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x02, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61,
	0x74, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x01, 0x52, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65,
	0x56, 0x61, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x06, 0x69, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x03, 0x28, 0x03, 0x52,
	0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x69, 0x6e,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x07, 0x75, 0x69, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76,
	0x61, 0x6c, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x12, 0x42, 0x0a, 0x08, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x07,
	0x6c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x55, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x56, 0x61, 0x6c, 0x1a, 0x65, 0x0a,
	0x0e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x4f, 0x41, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a,
	0x06, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x54,
	0x38, 0x10, 0x05, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x06, 0x12, 0x09,
	0x0a, 0x05, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x54,
	0x36, 0x34, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x55, 0x49, 0x4e, 0x54, 0x38, 0x10, 0x09, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x31, 0x36, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x55,
	0x49, 0x4e, 0x54, 0x33, 0x32, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x49, 0x4e, 0x54, 0x36,
	0x34, 0x10, 0x0c, 0x42, 0xe1, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70,
	0x62, 0x3b, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x70, 0x62, 0xaa, 0x02,
	0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x41, 0x49,
	0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x1f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0xea, 0x02, 0x22, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x41, 0x49, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1beta1_types_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1beta1_types_proto_rawDescData = file_google_cloud_aiplatform_v1beta1_types_proto_rawDesc
)

func file_google_cloud_aiplatform_v1beta1_types_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1beta1_types_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1beta1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1beta1_types_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1beta1_types_proto_rawDescData
}

var file_google_cloud_aiplatform_v1beta1_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_google_cloud_aiplatform_v1beta1_types_proto_goTypes = []interface{}{
	(Tensor_DataType)(0), // 0: google.cloud.aiplatform.v1beta1.Tensor.DataType
	(*BoolArray)(nil),    // 1: google.cloud.aiplatform.v1beta1.BoolArray
	(*DoubleArray)(nil),  // 2: google.cloud.aiplatform.v1beta1.DoubleArray
	(*Int64Array)(nil),   // 3: google.cloud.aiplatform.v1beta1.Int64Array
	(*StringArray)(nil),  // 4: google.cloud.aiplatform.v1beta1.StringArray
	(*Tensor)(nil),       // 5: google.cloud.aiplatform.v1beta1.Tensor
	nil,                  // 6: google.cloud.aiplatform.v1beta1.Tensor.StructValEntry
}
var file_google_cloud_aiplatform_v1beta1_types_proto_depIdxs = []int32{
	0, // 0: google.cloud.aiplatform.v1beta1.Tensor.dtype:type_name -> google.cloud.aiplatform.v1beta1.Tensor.DataType
	5, // 1: google.cloud.aiplatform.v1beta1.Tensor.list_val:type_name -> google.cloud.aiplatform.v1beta1.Tensor
	6, // 2: google.cloud.aiplatform.v1beta1.Tensor.struct_val:type_name -> google.cloud.aiplatform.v1beta1.Tensor.StructValEntry
	5, // 3: google.cloud.aiplatform.v1beta1.Tensor.StructValEntry.value:type_name -> google.cloud.aiplatform.v1beta1.Tensor
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_aiplatform_v1beta1_types_proto_init() }
func file_google_cloud_aiplatform_v1beta1_types_proto_init() {
	if File_google_cloud_aiplatform_v1beta1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoolArray); i {
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
		file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoubleArray); i {
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
		file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int64Array); i {
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
		file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringArray); i {
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
		file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tensor); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1beta1_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1beta1_types_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1beta1_types_proto_depIdxs,
		EnumInfos:         file_google_cloud_aiplatform_v1beta1_types_proto_enumTypes,
		MessageInfos:      file_google_cloud_aiplatform_v1beta1_types_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1beta1_types_proto = out.File
	file_google_cloud_aiplatform_v1beta1_types_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1beta1_types_proto_goTypes = nil
	file_google_cloud_aiplatform_v1beta1_types_proto_depIdxs = nil
}
