// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: api/v1/machine.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DeviceRequest_Device int32

const (
	DeviceRequest_NONE  DeviceRequest_Device = 0
	DeviceRequest_BIOS  DeviceRequest_Device = 1
	DeviceRequest_CDROM DeviceRequest_Device = 2
	DeviceRequest_DISK  DeviceRequest_Device = 3
	DeviceRequest_PXE   DeviceRequest_Device = 4
)

// Enum value maps for DeviceRequest_Device.
var (
	DeviceRequest_Device_name = map[int32]string{
		0: "NONE",
		1: "BIOS",
		2: "CDROM",
		3: "DISK",
		4: "PXE",
	}
	DeviceRequest_Device_value = map[string]int32{
		"NONE":  0,
		"BIOS":  1,
		"CDROM": 2,
		"DISK":  3,
		"PXE":   4,
	}
)

func (x DeviceRequest_Device) Enum() *DeviceRequest_Device {
	p := new(DeviceRequest_Device)
	*p = x
	return p
}

func (x DeviceRequest_Device) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeviceRequest_Device) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_machine_proto_enumTypes[0].Descriptor()
}

func (DeviceRequest_Device) Type() protoreflect.EnumType {
	return &file_api_v1_machine_proto_enumTypes[0]
}

func (x DeviceRequest_Device) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeviceRequest_Device.Descriptor instead.
func (DeviceRequest_Device) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_machine_proto_rawDescGZIP(), []int{0, 0}
}

type PowerRequest_Action int32

const (
	PowerRequest_ON      PowerRequest_Action = 0
	PowerRequest_OFF     PowerRequest_Action = 1
	PowerRequest_HARDOFF PowerRequest_Action = 2
	PowerRequest_CYCLE   PowerRequest_Action = 3
	PowerRequest_RESET   PowerRequest_Action = 4
	PowerRequest_STATUS  PowerRequest_Action = 5
)

// Enum value maps for PowerRequest_Action.
var (
	PowerRequest_Action_name = map[int32]string{
		0: "ON",
		1: "OFF",
		2: "HARDOFF",
		3: "CYCLE",
		4: "RESET",
		5: "STATUS",
	}
	PowerRequest_Action_value = map[string]int32{
		"ON":      0,
		"OFF":     1,
		"HARDOFF": 2,
		"CYCLE":   3,
		"RESET":   4,
		"STATUS":  5,
	}
)

func (x PowerRequest_Action) Enum() *PowerRequest_Action {
	p := new(PowerRequest_Action)
	*p = x
	return p
}

func (x PowerRequest_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PowerRequest_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_machine_proto_enumTypes[1].Descriptor()
}

func (PowerRequest_Action) Type() protoreflect.EnumType {
	return &file_api_v1_machine_proto_enumTypes[1]
}

func (x PowerRequest_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PowerRequest_Action.Descriptor instead.
func (PowerRequest_Action) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_machine_proto_rawDescGZIP(), []int{2, 0}
}

type DeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authn      *Authn               `protobuf:"bytes,1,opt,name=authn,proto3" json:"authn,omitempty"`
	Vendor     *Vendor              `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Device     DeviceRequest_Device `protobuf:"varint,3,opt,name=device,proto3,enum=github.com.tinkerbell.pbnj.api.v1.DeviceRequest_Device" json:"device,omitempty"`
	Persistent bool                 `protobuf:"varint,4,opt,name=persistent,proto3" json:"persistent,omitempty"`
	EfiBoot    bool                 `protobuf:"varint,5,opt,name=efi_boot,json=efiBoot,proto3" json:"efi_boot,omitempty"`
}

func (x *DeviceRequest) Reset() {
	*x = DeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_machine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceRequest) ProtoMessage() {}

func (x *DeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_machine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceRequest.ProtoReflect.Descriptor instead.
func (*DeviceRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_machine_proto_rawDescGZIP(), []int{0}
}

func (x *DeviceRequest) GetAuthn() *Authn {
	if x != nil {
		return x.Authn
	}
	return nil
}

func (x *DeviceRequest) GetVendor() *Vendor {
	if x != nil {
		return x.Vendor
	}
	return nil
}

func (x *DeviceRequest) GetDevice() DeviceRequest_Device {
	if x != nil {
		return x.Device
	}
	return DeviceRequest_NONE
}

func (x *DeviceRequest) GetPersistent() bool {
	if x != nil {
		return x.Persistent
	}
	return false
}

func (x *DeviceRequest) GetEfiBoot() bool {
	if x != nil {
		return x.EfiBoot
	}
	return false
}

type DeviceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *DeviceResponse) Reset() {
	*x = DeviceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_machine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceResponse) ProtoMessage() {}

func (x *DeviceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_machine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceResponse.ProtoReflect.Descriptor instead.
func (*DeviceResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_machine_proto_rawDescGZIP(), []int{1}
}

func (x *DeviceResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type PowerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authn       *Authn              `protobuf:"bytes,1,opt,name=authn,proto3" json:"authn,omitempty"`
	Vendor      *Vendor             `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Action      PowerRequest_Action `protobuf:"varint,3,opt,name=action,proto3,enum=github.com.tinkerbell.pbnj.api.v1.PowerRequest_Action" json:"action,omitempty"`
	SoftTimeout int32               `protobuf:"varint,4,opt,name=soft_timeout,json=softTimeout,proto3" json:"soft_timeout,omitempty"`
	OffDuration int32               `protobuf:"varint,5,opt,name=off_duration,json=offDuration,proto3" json:"off_duration,omitempty"`
}

func (x *PowerRequest) Reset() {
	*x = PowerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_machine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerRequest) ProtoMessage() {}

func (x *PowerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_machine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerRequest.ProtoReflect.Descriptor instead.
func (*PowerRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_machine_proto_rawDescGZIP(), []int{2}
}

func (x *PowerRequest) GetAuthn() *Authn {
	if x != nil {
		return x.Authn
	}
	return nil
}

func (x *PowerRequest) GetVendor() *Vendor {
	if x != nil {
		return x.Vendor
	}
	return nil
}

func (x *PowerRequest) GetAction() PowerRequest_Action {
	if x != nil {
		return x.Action
	}
	return PowerRequest_ON
}

func (x *PowerRequest) GetSoftTimeout() int32 {
	if x != nil {
		return x.SoftTimeout
	}
	return 0
}

func (x *PowerRequest) GetOffDuration() int32 {
	if x != nil {
		return x.OffDuration
	}
	return 0
}

type PowerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *PowerResponse) Reset() {
	*x = PowerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_machine_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowerResponse) ProtoMessage() {}

func (x *PowerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_machine_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowerResponse.ProtoReflect.Descriptor instead.
func (*PowerResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_machine_proto_rawDescGZIP(), []int{3}
}

func (x *PowerResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

var File_api_v1_machine_proto protoreflect.FileDescriptor

var file_api_v1_machine_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62,
	0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda,
	0x02, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3e, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x28, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6e, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x12, 0x41, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x12, 0x4f, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x66, 0x69, 0x5f, 0x62, 0x6f, 0x6f, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x66, 0x69, 0x42, 0x6f, 0x6f, 0x74, 0x22,
	0x3a, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x49, 0x4f, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a,
	0x05, 0x43, 0x44, 0x52, 0x4f, 0x4d, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x49, 0x53, 0x4b,
	0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x58, 0x45, 0x10, 0x04, 0x22, 0x29, 0x0a, 0x0e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0xf1, 0x02, 0x0a, 0x0c, 0x50, 0x6f, 0x77, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70,
	0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6e,
	0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x41, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e,
	0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x4e, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65,
	0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x6f,
	0x66, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x73, 0x6f, 0x66, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x6f, 0x66, 0x66, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x6f, 0x66, 0x66, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x48, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x48,
	0x41, 0x52, 0x44, 0x4f, 0x46, 0x46, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x59, 0x43, 0x4c,
	0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x04, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x10, 0x05, 0x22, 0x28, 0x0a, 0x0d, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74,
	0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x32, 0xe8, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x12, 0x71, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x31, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69,
	0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x05, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x2f, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65, 0x72,
	0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2e, 0x70, 0x62, 0x6e, 0x6a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x65, 0x72, 0x62, 0x65, 0x6c, 0x6c, 0x2f, 0x70, 0x62, 0x6e, 0x6a, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_machine_proto_rawDescOnce sync.Once
	file_api_v1_machine_proto_rawDescData = file_api_v1_machine_proto_rawDesc
)

func file_api_v1_machine_proto_rawDescGZIP() []byte {
	file_api_v1_machine_proto_rawDescOnce.Do(func() {
		file_api_v1_machine_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_machine_proto_rawDescData)
	})
	return file_api_v1_machine_proto_rawDescData
}

var file_api_v1_machine_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v1_machine_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_machine_proto_goTypes = []interface{}{
	(DeviceRequest_Device)(0), // 0: github.com.tinkerbell.pbnj.api.v1.DeviceRequest.Device
	(PowerRequest_Action)(0),  // 1: github.com.tinkerbell.pbnj.api.v1.PowerRequest.Action
	(*DeviceRequest)(nil),     // 2: github.com.tinkerbell.pbnj.api.v1.DeviceRequest
	(*DeviceResponse)(nil),    // 3: github.com.tinkerbell.pbnj.api.v1.DeviceResponse
	(*PowerRequest)(nil),      // 4: github.com.tinkerbell.pbnj.api.v1.PowerRequest
	(*PowerResponse)(nil),     // 5: github.com.tinkerbell.pbnj.api.v1.PowerResponse
	(*Authn)(nil),             // 6: github.com.tinkerbell.pbnj.api.v1.Authn
	(*Vendor)(nil),            // 7: github.com.tinkerbell.pbnj.api.v1.Vendor
}
var file_api_v1_machine_proto_depIdxs = []int32{
	6, // 0: github.com.tinkerbell.pbnj.api.v1.DeviceRequest.authn:type_name -> github.com.tinkerbell.pbnj.api.v1.Authn
	7, // 1: github.com.tinkerbell.pbnj.api.v1.DeviceRequest.vendor:type_name -> github.com.tinkerbell.pbnj.api.v1.Vendor
	0, // 2: github.com.tinkerbell.pbnj.api.v1.DeviceRequest.device:type_name -> github.com.tinkerbell.pbnj.api.v1.DeviceRequest.Device
	6, // 3: github.com.tinkerbell.pbnj.api.v1.PowerRequest.authn:type_name -> github.com.tinkerbell.pbnj.api.v1.Authn
	7, // 4: github.com.tinkerbell.pbnj.api.v1.PowerRequest.vendor:type_name -> github.com.tinkerbell.pbnj.api.v1.Vendor
	1, // 5: github.com.tinkerbell.pbnj.api.v1.PowerRequest.action:type_name -> github.com.tinkerbell.pbnj.api.v1.PowerRequest.Action
	2, // 6: github.com.tinkerbell.pbnj.api.v1.Machine.BootDevice:input_type -> github.com.tinkerbell.pbnj.api.v1.DeviceRequest
	4, // 7: github.com.tinkerbell.pbnj.api.v1.Machine.Power:input_type -> github.com.tinkerbell.pbnj.api.v1.PowerRequest
	3, // 8: github.com.tinkerbell.pbnj.api.v1.Machine.BootDevice:output_type -> github.com.tinkerbell.pbnj.api.v1.DeviceResponse
	5, // 9: github.com.tinkerbell.pbnj.api.v1.Machine.Power:output_type -> github.com.tinkerbell.pbnj.api.v1.PowerResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1_machine_proto_init() }
func file_api_v1_machine_proto_init() {
	if File_api_v1_machine_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v1_machine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceRequest); i {
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
		file_api_v1_machine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceResponse); i {
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
		file_api_v1_machine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerRequest); i {
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
		file_api_v1_machine_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowerResponse); i {
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
			RawDescriptor: file_api_v1_machine_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_machine_proto_goTypes,
		DependencyIndexes: file_api_v1_machine_proto_depIdxs,
		EnumInfos:         file_api_v1_machine_proto_enumTypes,
		MessageInfos:      file_api_v1_machine_proto_msgTypes,
	}.Build()
	File_api_v1_machine_proto = out.File
	file_api_v1_machine_proto_rawDesc = nil
	file_api_v1_machine_proto_goTypes = nil
	file_api_v1_machine_proto_depIdxs = nil
}
