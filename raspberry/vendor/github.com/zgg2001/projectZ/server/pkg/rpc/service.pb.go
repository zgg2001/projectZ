// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.8
// source: service.proto

package rpc

import (
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

// license plate check
type LPCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model          int32  `protobuf:"varint,1,opt,name=model,proto3" json:"model,omitempty"`
	ParkingId      int32  `protobuf:"varint,2,opt,name=parkingId,proto3" json:"parkingId,omitempty"`
	ParkingSpaceId int32  `protobuf:"varint,3,opt,name=parkingSpaceId,proto3" json:"parkingSpaceId,omitempty"`
	License        string `protobuf:"bytes,4,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *LPCheckRequest) Reset() {
	*x = LPCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LPCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LPCheckRequest) ProtoMessage() {}

func (x *LPCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LPCheckRequest.ProtoReflect.Descriptor instead.
func (*LPCheckRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *LPCheckRequest) GetModel() int32 {
	if x != nil {
		return x.Model
	}
	return 0
}

func (x *LPCheckRequest) GetParkingId() int32 {
	if x != nil {
		return x.ParkingId
	}
	return 0
}

func (x *LPCheckRequest) GetParkingSpaceId() int32 {
	if x != nil {
		return x.ParkingSpaceId
	}
	return 0
}

func (x *LPCheckRequest) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

type LPCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  bool  `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	Balance int32 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *LPCheckResponse) Reset() {
	*x = LPCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LPCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LPCheckResponse) ProtoMessage() {}

func (x *LPCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LPCheckResponse.ProtoReflect.Descriptor instead.
func (*LPCheckResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *LPCheckResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *LPCheckResponse) GetBalance() int32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

// upload parking info
type ParkingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Temperature int32 `protobuf:"varint,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    int32 `protobuf:"varint,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Weather     int32 `protobuf:"varint,4,opt,name=weather,proto3" json:"weather,omitempty"`
}

func (x *ParkingInfo) Reset() {
	*x = ParkingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkingInfo) ProtoMessage() {}

func (x *ParkingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParkingInfo.ProtoReflect.Descriptor instead.
func (*ParkingInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *ParkingInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ParkingInfo) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *ParkingInfo) GetHumidity() int32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *ParkingInfo) GetWeather() int32 {
	if x != nil {
		return x.Weather
	}
	return 0
}

type ParkingSpaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Temperature int32 `protobuf:"varint,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    int32 `protobuf:"varint,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Alarm       int32 `protobuf:"varint,4,opt,name=alarm,proto3" json:"alarm,omitempty"`
}

func (x *ParkingSpaceInfo) Reset() {
	*x = ParkingSpaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkingSpaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkingSpaceInfo) ProtoMessage() {}

func (x *ParkingSpaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParkingSpaceInfo.ProtoReflect.Descriptor instead.
func (*ParkingSpaceInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *ParkingSpaceInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ParkingSpaceInfo) GetTemperature() int32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *ParkingSpaceInfo) GetHumidity() int32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *ParkingSpaceInfo) GetAlarm() int32 {
	if x != nil {
		return x.Alarm
	}
	return 0
}

type UploadInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PInfo   *ParkingInfo        `protobuf:"bytes,1,opt,name=pInfo,proto3" json:"pInfo,omitempty"`
	InfoArr []*ParkingSpaceInfo `protobuf:"bytes,2,rep,name=infoArr,proto3" json:"infoArr,omitempty"`
}

func (x *UploadInfoRequest) Reset() {
	*x = UploadInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInfoRequest) ProtoMessage() {}

func (x *UploadInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInfoRequest.ProtoReflect.Descriptor instead.
func (*UploadInfoRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *UploadInfoRequest) GetPInfo() *ParkingInfo {
	if x != nil {
		return x.PInfo
	}
	return nil
}

func (x *UploadInfoRequest) GetInfoArr() []*ParkingSpaceInfo {
	if x != nil {
		return x.InfoArr
	}
	return nil
}

type UploadInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UploadInfoResponse) Reset() {
	*x = UploadInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInfoResponse) ProtoMessage() {}

func (x *UploadInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadInfoResponse.ProtoReflect.Descriptor instead.
func (*UploadInfoResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *UploadInfoResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x86, 0x01, 0x0a, 0x0e, 0x4c, 0x50, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e,
	0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x43, 0x0a, 0x0f, 0x4c, 0x50, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x75, 0x0a,
	0x0b, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65,
	0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x77, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x10, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x22, 0x64, 0x0a, 0x11,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x07, 0x69, 0x6e, 0x66, 0x6f, 0x41, 0x72, 0x72,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x69, 0x6e, 0x66, 0x6f, 0x41,
	0x72, 0x72, 0x22, 0x2c, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x86, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x11, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x63, 0x65, 0x50, 0x6c,
	0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x4c, 0x50, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4c, 0x50, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x11, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_proto_goTypes = []interface{}{
	(*LPCheckRequest)(nil),     // 0: LPCheckRequest
	(*LPCheckResponse)(nil),    // 1: LPCheckResponse
	(*ParkingInfo)(nil),        // 2: ParkingInfo
	(*ParkingSpaceInfo)(nil),   // 3: ParkingSpaceInfo
	(*UploadInfoRequest)(nil),  // 4: UploadInfoRequest
	(*UploadInfoResponse)(nil), // 5: UploadInfoResponse
}
var file_service_proto_depIdxs = []int32{
	2, // 0: UploadInfoRequest.pInfo:type_name -> ParkingInfo
	3, // 1: UploadInfoRequest.infoArr:type_name -> ParkingSpaceInfo
	0, // 2: ProjectService.LicencePlateCheck:input_type -> LPCheckRequest
	4, // 3: ProjectService.UploadParkingInfo:input_type -> UploadInfoRequest
	1, // 4: ProjectService.LicencePlateCheck:output_type -> LPCheckResponse
	5, // 5: ProjectService.UploadParkingInfo:output_type -> UploadInfoResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LPCheckRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LPCheckResponse); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParkingInfo); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParkingSpaceInfo); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadInfoRequest); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadInfoResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}
