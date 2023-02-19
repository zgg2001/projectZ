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

type ParkingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PId         int32 `protobuf:"varint,1,opt,name=p_id,json=pId,proto3" json:"p_id,omitempty"`
	Temperature int32 `protobuf:"varint,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    int32 `protobuf:"varint,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Weather     int32 `protobuf:"varint,4,opt,name=weather,proto3" json:"weather,omitempty"`
}

func (x *ParkingInfo) Reset() {
	*x = ParkingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkingInfo) ProtoMessage() {}

func (x *ParkingInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ParkingInfo.ProtoReflect.Descriptor instead.
func (*ParkingInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *ParkingInfo) GetPId() int32 {
	if x != nil {
		return x.PId
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

	SId         int32 `protobuf:"varint,1,opt,name=s_id,json=sId,proto3" json:"s_id,omitempty"`
	Temperature int32 `protobuf:"varint,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humidity    int32 `protobuf:"varint,3,opt,name=humidity,proto3" json:"humidity,omitempty"`
	Alarm       int32 `protobuf:"varint,4,opt,name=alarm,proto3" json:"alarm,omitempty"`
}

func (x *ParkingSpaceInfo) Reset() {
	*x = ParkingSpaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParkingSpaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParkingSpaceInfo) ProtoMessage() {}

func (x *ParkingSpaceInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ParkingSpaceInfo.ProtoReflect.Descriptor instead.
func (*ParkingSpaceInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *ParkingSpaceInfo) GetSId() int32 {
	if x != nil {
		return x.SId
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

type CarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PTemperature int32  `protobuf:"varint,1,opt,name=p_temperature,json=pTemperature,proto3" json:"p_temperature,omitempty"`
	PHumidity    int32  `protobuf:"varint,2,opt,name=p_humidity,json=pHumidity,proto3" json:"p_humidity,omitempty"`
	PWeather     int32  `protobuf:"varint,3,opt,name=p_weather,json=pWeather,proto3" json:"p_weather,omitempty"`
	PAddress     string `protobuf:"bytes,4,opt,name=p_address,json=pAddress,proto3" json:"p_address,omitempty"`
	SId          int32  `protobuf:"varint,5,opt,name=s_id,json=sId,proto3" json:"s_id,omitempty"`
	STemperature int32  `protobuf:"varint,6,opt,name=s_temperature,json=sTemperature,proto3" json:"s_temperature,omitempty"`
	SHumidity    int32  `protobuf:"varint,7,opt,name=s_humidity,json=sHumidity,proto3" json:"s_humidity,omitempty"`
	SAlarm       int32  `protobuf:"varint,8,opt,name=s_alarm,json=sAlarm,proto3" json:"s_alarm,omitempty"`
}

func (x *CarInfo) Reset() {
	*x = CarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarInfo) ProtoMessage() {}

func (x *CarInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use CarInfo.ProtoReflect.Descriptor instead.
func (*CarInfo) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *CarInfo) GetPTemperature() int32 {
	if x != nil {
		return x.PTemperature
	}
	return 0
}

func (x *CarInfo) GetPHumidity() int32 {
	if x != nil {
		return x.PHumidity
	}
	return 0
}

func (x *CarInfo) GetPWeather() int32 {
	if x != nil {
		return x.PWeather
	}
	return 0
}

func (x *CarInfo) GetPAddress() string {
	if x != nil {
		return x.PAddress
	}
	return ""
}

func (x *CarInfo) GetSId() int32 {
	if x != nil {
		return x.SId
	}
	return 0
}

func (x *CarInfo) GetSTemperature() int32 {
	if x != nil {
		return x.STemperature
	}
	return 0
}

func (x *CarInfo) GetSHumidity() int32 {
	if x != nil {
		return x.SHumidity
	}
	return 0
}

func (x *CarInfo) GetSAlarm() int32 {
	if x != nil {
		return x.SAlarm
	}
	return 0
}

// raspbery
// license plate check
type LPCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model          int32  `protobuf:"varint,1,opt,name=model,proto3" json:"model,omitempty"`
	ParkingId      int32  `protobuf:"varint,2,opt,name=parking_id,json=parkingId,proto3" json:"parking_id,omitempty"`
	ParkingSpaceId int32  `protobuf:"varint,3,opt,name=parking_space_id,json=parkingSpaceId,proto3" json:"parking_space_id,omitempty"`
	License        string `protobuf:"bytes,4,opt,name=license,proto3" json:"license,omitempty"`
}

func (x *LPCheckRequest) Reset() {
	*x = LPCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LPCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LPCheckRequest) ProtoMessage() {}

func (x *LPCheckRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LPCheckRequest.ProtoReflect.Descriptor instead.
func (*LPCheckRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
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
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LPCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LPCheckResponse) ProtoMessage() {}

func (x *LPCheckResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use LPCheckResponse.ProtoReflect.Descriptor instead.
func (*LPCheckResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
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
type UploadInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PInfo    *ParkingInfo        `protobuf:"bytes,1,opt,name=p_info,json=pInfo,proto3" json:"p_info,omitempty"`
	SInfoArr []*ParkingSpaceInfo `protobuf:"bytes,2,rep,name=s_info_arr,json=sInfoArr,proto3" json:"s_info_arr,omitempty"`
}

func (x *UploadInfoRequest) Reset() {
	*x = UploadInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInfoRequest) ProtoMessage() {}

func (x *UploadInfoRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use UploadInfoRequest.ProtoReflect.Descriptor instead.
func (*UploadInfoRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *UploadInfoRequest) GetPInfo() *ParkingInfo {
	if x != nil {
		return x.PInfo
	}
	return nil
}

func (x *UploadInfoRequest) GetSInfoArr() []*ParkingSpaceInfo {
	if x != nil {
		return x.SInfoArr
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
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadInfoResponse) ProtoMessage() {}

func (x *UploadInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
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
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *UploadInfoResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

// client
// get User data
type GetUserDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UId int32 `protobuf:"varint,1,opt,name=u_id,json=uId,proto3" json:"u_id,omitempty"`
}

func (x *GetUserDataRequest) Reset() {
	*x = GetUserDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDataRequest) ProtoMessage() {}

func (x *GetUserDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDataRequest.ProtoReflect.Descriptor instead.
func (*GetUserDataRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserDataRequest) GetUId() int32 {
	if x != nil {
		return x.UId
	}
	return 0
}

type GetUserDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarInfoArr []*CarInfo `protobuf:"bytes,1,rep,name=car_info_arr,json=carInfoArr,proto3" json:"car_info_arr,omitempty"`
}

func (x *GetUserDataResponse) Reset() {
	*x = GetUserDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDataResponse) ProtoMessage() {}

func (x *GetUserDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDataResponse.ProtoReflect.Descriptor instead.
func (*GetUserDataResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{8}
}

func (x *GetUserDataResponse) GetCarInfoArr() []*CarInfo {
	if x != nil {
		return x.CarInfoArr
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x78, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11,
	0x0a, 0x04, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x22, 0x79, 0x0a, 0x10, 0x50, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x11, 0x0a,
	0x04, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61,
	0x6c, 0x61, 0x72, 0x6d, 0x22, 0xf7, 0x01, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x23, 0x0a, 0x0d, 0x70, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x54, 0x65, 0x6d, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x5f, 0x68, 0x75, 0x6d, 0x69, 0x64,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x48, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x5f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x11,
	0x0a, 0x04, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x54, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x5f, 0x68, 0x75, 0x6d, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x48, 0x75, 0x6d,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x22, 0x89,
	0x01, 0x0a, 0x0e, 0x4c, 0x50, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x22, 0x43, 0x0a, 0x0f, 0x4c, 0x50,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22,
	0x69, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x06, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x0a, 0x73, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x08, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x72, 0x72, 0x22, 0x2c, 0x0a, 0x12, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x27, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x11,
	0x0a, 0x04, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x49,
	0x64, 0x22, 0x41, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x72, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x43, 0x61, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x41, 0x72, 0x72, 0x32, 0xc0, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x11, 0x4c, 0x69, 0x63, 0x65, 0x6e,
	0x63, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0f, 0x2e, 0x4c,
	0x50, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x4c, 0x50, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3c, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x13, 0x2e, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_proto_goTypes = []interface{}{
	(*ParkingInfo)(nil),         // 0: ParkingInfo
	(*ParkingSpaceInfo)(nil),    // 1: ParkingSpaceInfo
	(*CarInfo)(nil),             // 2: CarInfo
	(*LPCheckRequest)(nil),      // 3: LPCheckRequest
	(*LPCheckResponse)(nil),     // 4: LPCheckResponse
	(*UploadInfoRequest)(nil),   // 5: UploadInfoRequest
	(*UploadInfoResponse)(nil),  // 6: UploadInfoResponse
	(*GetUserDataRequest)(nil),  // 7: GetUserDataRequest
	(*GetUserDataResponse)(nil), // 8: GetUserDataResponse
}
var file_service_proto_depIdxs = []int32{
	0, // 0: UploadInfoRequest.p_info:type_name -> ParkingInfo
	1, // 1: UploadInfoRequest.s_info_arr:type_name -> ParkingSpaceInfo
	2, // 2: GetUserDataResponse.car_info_arr:type_name -> CarInfo
	3, // 3: ProjectService.LicencePlateCheck:input_type -> LPCheckRequest
	5, // 4: ProjectService.UploadParkingInfo:input_type -> UploadInfoRequest
	7, // 5: ProjectService.GetUserData:input_type -> GetUserDataRequest
	4, // 6: ProjectService.LicencePlateCheck:output_type -> LPCheckResponse
	6, // 7: ProjectService.UploadParkingInfo:output_type -> UploadInfoResponse
	8, // 8: ProjectService.GetUserData:output_type -> GetUserDataResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarInfo); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDataRequest); i {
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
		file_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserDataResponse); i {
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
			NumMessages:   9,
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
