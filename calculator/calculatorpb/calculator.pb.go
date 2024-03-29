// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstNumber  int32 `protobuf:"varint,1,opt,name=first_number,json=firstNumber,proto3" json:"first_number,omitempty"`
	SecondNumber int32 `protobuf:"varint,2,opt,name=second_number,json=secondNumber,proto3" json:"second_number,omitempty"`
}

func (x *Values) Reset() {
	*x = Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Values) ProtoMessage() {}

func (x *Values) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Values.ProtoReflect.Descriptor instead.
func (*Values) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Values) GetFirstNumber() int32 {
	if x != nil {
		return x.FirstNumber
	}
	return 0
}

func (x *Values) GetSecondNumber() int32 {
	if x != nil {
		return x.SecondNumber
	}
	return 0
}

type CalculatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Values `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CalculatorRequest) Reset() {
	*x = CalculatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorRequest) ProtoMessage() {}

func (x *CalculatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorRequest.ProtoReflect.Descriptor instead.
func (*CalculatorRequest) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculatorRequest) GetValue() *Values {
	if x != nil {
		return x.Value
	}
	return nil
}

type CalculatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CalculatorResponse) Reset() {
	*x = CalculatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatorResponse) ProtoMessage() {}

func (x *CalculatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatorResponse.ProtoReflect.Descriptor instead.
func (*CalculatorResponse) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *CalculatorResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type PrimeDecompositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *PrimeDecompositionRequest) Reset() {
	*x = PrimeDecompositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeDecompositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeDecompositionRequest) ProtoMessage() {}

func (x *PrimeDecompositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeDecompositionRequest.ProtoReflect.Descriptor instead.
func (*PrimeDecompositionRequest) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeDecompositionRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type PrimeDecompositionReponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *PrimeDecompositionReponse) Reset() {
	*x = PrimeDecompositionReponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeDecompositionReponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeDecompositionReponse) ProtoMessage() {}

func (x *PrimeDecompositionReponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeDecompositionReponse.ProtoReflect.Descriptor instead.
func (*PrimeDecompositionReponse) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *PrimeDecompositionReponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *AverageRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float64 `protobuf:"fixed64,2,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *AverageResponse) GetAverage() float64 {
	if x != nil {
		return x.Average
	}
	return 0
}

type FindMaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FindMaximumRequest) Reset() {
	*x = FindMaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumRequest) ProtoMessage() {}

func (x *FindMaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumRequest.ProtoReflect.Descriptor instead.
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *FindMaximumRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int32 `protobuf:"varint,2,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *FindMaximumResponse) Reset() {
	*x = FindMaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatorpb_calculator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumResponse) ProtoMessage() {}

func (x *FindMaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculatorpb_calculator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumResponse.ProtoReflect.Descriptor instead.
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return file_calculatorpb_calculator_proto_rawDescGZIP(), []int{8}
}

func (x *FindMaximumResponse) GetMaximum() int32 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x50, 0x0a, 0x06, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0c, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3d, 0x0a,
	0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2c, 0x0a, 0x12,
	0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x33, 0x0a, 0x19, 0x50, 0x72,
	0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x33, 0x0a, 0x19, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x26, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2b, 0x0a, 0x0f,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x12, 0x46, 0x69, 0x6e,
	0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x32, 0xda, 0x02, 0x0a, 0x0a, 0x53, 0x75, 0x6d,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x1d,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x66, 0x0a, 0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44,
	0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x07, 0x41, 0x76, 0x65, 0x72, 0x61,
	0x67, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x76, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12,
	0x54, 0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x1e,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_calculatorpb_calculator_proto_rawDescData = file_calculatorpb_calculator_proto_rawDesc
)

func file_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculatorpb_calculator_proto_rawDescData)
	})
	return file_calculatorpb_calculator_proto_rawDescData
}

var file_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*Values)(nil),                    // 0: calculator.Values
	(*CalculatorRequest)(nil),         // 1: calculator.CalculatorRequest
	(*CalculatorResponse)(nil),        // 2: calculator.CalculatorResponse
	(*PrimeDecompositionRequest)(nil), // 3: calculator.PrimeDecompositionRequest
	(*PrimeDecompositionReponse)(nil), // 4: calculator.PrimeDecompositionReponse
	(*AverageRequest)(nil),            // 5: calculator.AverageRequest
	(*AverageResponse)(nil),           // 6: calculator.AverageResponse
	(*FindMaximumRequest)(nil),        // 7: calculator.FindMaximumRequest
	(*FindMaximumResponse)(nil),       // 8: calculator.FindMaximumResponse
}
var file_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorRequest.value:type_name -> calculator.Values
	1, // 1: calculator.SumService.Sum:input_type -> calculator.CalculatorRequest
	3, // 2: calculator.SumService.PrimeDecomposition:input_type -> calculator.PrimeDecompositionRequest
	5, // 3: calculator.SumService.Average:input_type -> calculator.AverageRequest
	7, // 4: calculator.SumService.FindMaximum:input_type -> calculator.FindMaximumRequest
	2, // 5: calculator.SumService.Sum:output_type -> calculator.CalculatorResponse
	4, // 6: calculator.SumService.PrimeDecomposition:output_type -> calculator.PrimeDecompositionReponse
	6, // 7: calculator.SumService.Average:output_type -> calculator.AverageResponse
	8, // 8: calculator.SumService.FindMaximum:output_type -> calculator.FindMaximumResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_calculatorpb_calculator_proto_init() }
func file_calculatorpb_calculator_proto_init() {
	if File_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Values); i {
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
		file_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorRequest); i {
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
		file_calculatorpb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatorResponse); i {
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
		file_calculatorpb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeDecompositionRequest); i {
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
		file_calculatorpb_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeDecompositionReponse); i {
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
		file_calculatorpb_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
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
		file_calculatorpb_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
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
		file_calculatorpb_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumRequest); i {
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
		file_calculatorpb_calculator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumResponse); i {
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
			RawDescriptor: file_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_calculatorpb_calculator_proto = out.File
	file_calculatorpb_calculator_proto_rawDesc = nil
	file_calculatorpb_calculator_proto_goTypes = nil
	file_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	//Unary
	Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error)
	//Server streaming
	PrimeDecomposition(ctx context.Context, in *PrimeDecompositionRequest, opts ...grpc.CallOption) (SumService_PrimeDecompositionClient, error)
	//Client streaming
	Average(ctx context.Context, opts ...grpc.CallOption) (SumService_AverageClient, error)
	//BiDi streaming
	FindMaximum(ctx context.Context, opts ...grpc.CallOption) (SumService_FindMaximumClient, error)
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) Sum(ctx context.Context, in *CalculatorRequest, opts ...grpc.CallOption) (*CalculatorResponse, error) {
	out := new(CalculatorResponse)
	err := c.cc.Invoke(ctx, "/calculator.SumService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) PrimeDecomposition(ctx context.Context, in *PrimeDecompositionRequest, opts ...grpc.CallOption) (SumService_PrimeDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[0], "/calculator.SumService/PrimeDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServicePrimeDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_PrimeDecompositionClient interface {
	Recv() (*PrimeDecompositionReponse, error)
	grpc.ClientStream
}

type sumServicePrimeDecompositionClient struct {
	grpc.ClientStream
}

func (x *sumServicePrimeDecompositionClient) Recv() (*PrimeDecompositionReponse, error) {
	m := new(PrimeDecompositionReponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) Average(ctx context.Context, opts ...grpc.CallOption) (SumService_AverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[1], "/calculator.SumService/Average", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceAverageClient{stream}
	return x, nil
}

type SumService_AverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type sumServiceAverageClient struct {
	grpc.ClientStream
}

func (x *sumServiceAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sumServiceClient) FindMaximum(ctx context.Context, opts ...grpc.CallOption) (SumService_FindMaximumClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[2], "/calculator.SumService/FindMaximum", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceFindMaximumClient{stream}
	return x, nil
}

type SumService_FindMaximumClient interface {
	Send(*FindMaximumRequest) error
	Recv() (*FindMaximumResponse, error)
	grpc.ClientStream
}

type sumServiceFindMaximumClient struct {
	grpc.ClientStream
}

func (x *sumServiceFindMaximumClient) Send(m *FindMaximumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumServiceFindMaximumClient) Recv() (*FindMaximumResponse, error) {
	m := new(FindMaximumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	//Unary
	Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error)
	//Server streaming
	PrimeDecomposition(*PrimeDecompositionRequest, SumService_PrimeDecompositionServer) error
	//Client streaming
	Average(SumService_AverageServer) error
	//BiDi streaming
	FindMaximum(SumService_FindMaximumServer) error
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) Sum(context.Context, *CalculatorRequest) (*CalculatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedSumServiceServer) PrimeDecomposition(*PrimeDecompositionRequest, SumService_PrimeDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeDecomposition not implemented")
}
func (*UnimplementedSumServiceServer) Average(SumService_AverageServer) error {
	return status.Errorf(codes.Unimplemented, "method Average not implemented")
}
func (*UnimplementedSumServiceServer) FindMaximum(SumService_FindMaximumServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMaximum not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalculatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.SumService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).Sum(ctx, req.(*CalculatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_PrimeDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeDecompositionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).PrimeDecomposition(m, &sumServicePrimeDecompositionServer{stream})
}

type SumService_PrimeDecompositionServer interface {
	Send(*PrimeDecompositionReponse) error
	grpc.ServerStream
}

type sumServicePrimeDecompositionServer struct {
	grpc.ServerStream
}

func (x *sumServicePrimeDecompositionServer) Send(m *PrimeDecompositionReponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SumService_Average_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).Average(&sumServiceAverageServer{stream})
}

type SumService_AverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type sumServiceAverageServer struct {
	grpc.ServerStream
}

func (x *sumServiceAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SumService_FindMaximum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumServiceServer).FindMaximum(&sumServiceFindMaximumServer{stream})
}

type SumService_FindMaximumServer interface {
	Send(*FindMaximumResponse) error
	Recv() (*FindMaximumRequest, error)
	grpc.ServerStream
}

type sumServiceFindMaximumServer struct {
	grpc.ServerStream
}

func (x *sumServiceFindMaximumServer) Send(m *FindMaximumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumServiceFindMaximumServer) Recv() (*FindMaximumRequest, error) {
	m := new(FindMaximumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _SumService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeDecomposition",
			Handler:       _SumService_PrimeDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Average",
			Handler:       _SumService_Average_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMaximum",
			Handler:       _SumService_FindMaximum_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculatorpb/calculator.proto",
}
