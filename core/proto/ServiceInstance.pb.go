// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ServiceInstance.proto

package proto

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ServiceInstance_ServiceState int32

const (
	ServiceInstance_UNKNOWN ServiceInstance_ServiceState = 0
	ServiceInstance_INIT    ServiceInstance_ServiceState = 1
	ServiceInstance_STOP    ServiceInstance_ServiceState = 2
	ServiceInstance_RUN     ServiceInstance_ServiceState = 3
	ServiceInstance_ERROR   ServiceInstance_ServiceState = 4
)

var ServiceInstance_ServiceState_name = map[int32]string{
	0: "UNKNOWN",
	1: "INIT",
	2: "STOP",
	3: "RUN",
	4: "ERROR",
}

var ServiceInstance_ServiceState_value = map[string]int32{
	"UNKNOWN": 0,
	"INIT":    1,
	"STOP":    2,
	"RUN":     3,
	"ERROR":   4,
}

func (x ServiceInstance_ServiceState) String() string {
	return proto.EnumName(ServiceInstance_ServiceState_name, int32(x))
}

func (ServiceInstance_ServiceState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d00b0fd20962128, []int{0, 0}
}

type ServiceInstance struct {
	Id                   string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Module               string                       `protobuf:"bytes,2,opt,name=module,proto3" json:"module,omitempty"`
	State                ServiceInstance_ServiceState `protobuf:"varint,3,opt,name=state,proto3,enum=proto.ServiceInstance_ServiceState" json:"state,omitempty"`
	Config               *types.Any                   `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	ErrorMsg             string                       `protobuf:"bytes,5,opt,name=error_msg,json=errorMsg,proto3" json:"error_msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ServiceInstance) Reset()         { *m = ServiceInstance{} }
func (m *ServiceInstance) String() string { return proto.CompactTextString(m) }
func (*ServiceInstance) ProtoMessage()    {}
func (*ServiceInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d00b0fd20962128, []int{0}
}
func (m *ServiceInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceInstance.Unmarshal(m, b)
}
func (m *ServiceInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceInstance.Marshal(b, m, deterministic)
}
func (m *ServiceInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceInstance.Merge(m, src)
}
func (m *ServiceInstance) XXX_Size() int {
	return xxx_messageInfo_ServiceInstance.Size(m)
}
func (m *ServiceInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceInstance.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceInstance proto.InternalMessageInfo

func (m *ServiceInstance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceInstance) GetModule() string {
	if m != nil {
		return m.Module
	}
	return ""
}

func (m *ServiceInstance) GetState() ServiceInstance_ServiceState {
	if m != nil {
		return m.State
	}
	return ServiceInstance_UNKNOWN
}

func (m *ServiceInstance) GetConfig() *types.Any {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ServiceInstance) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.ServiceInstance_ServiceState", ServiceInstance_ServiceState_name, ServiceInstance_ServiceState_value)
	proto.RegisterType((*ServiceInstance)(nil), "proto.ServiceInstance")
}

func init() { proto.RegisterFile("ServiceInstance.proto", fileDescriptor_3d00b0fd20962128) }

var fileDescriptor_3d00b0fd20962128 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xcf, 0x4b, 0xfb, 0x40,
	0x14, 0xc4, 0xbf, 0xf9, 0xdd, 0xbc, 0x7e, 0xa9, 0xe1, 0xa1, 0x12, 0xf5, 0x12, 0xea, 0x25, 0x07,
	0xd9, 0x42, 0x3d, 0x89, 0x27, 0x15, 0x0f, 0x41, 0x4c, 0x64, 0xd3, 0x22, 0x78, 0x91, 0x98, 0x6c,
	0x97, 0x40, 0xbb, 0x2b, 0x9b, 0x54, 0xe8, 0xdd, 0x3f, 0x5c, 0xba, 0x1b, 0x41, 0x7a, 0x7a, 0x3b,
	0xc3, 0xcc, 0xce, 0x07, 0x4e, 0x4a, 0xa6, 0xbe, 0xda, 0x9a, 0x65, 0xa2, 0xeb, 0x2b, 0x51, 0x33,
	0xf2, 0xa9, 0x64, 0x2f, 0xd1, 0xd3, 0xe7, 0xfc, 0x8c, 0x4b, 0xc9, 0xd7, 0x6c, 0xa6, 0xd5, 0xc7,
	0x76, 0x35, 0xab, 0xc4, 0xce, 0x24, 0xa6, 0xdf, 0x36, 0x1c, 0x1d, 0x74, 0x71, 0x02, 0x76, 0xdb,
	0xc4, 0x56, 0x62, 0xa5, 0x21, 0xb5, 0xdb, 0x06, 0x4f, 0xc1, 0xdf, 0xc8, 0x66, 0xbb, 0x66, 0xb1,
	0xad, 0xbd, 0x41, 0xe1, 0x0d, 0x78, 0x5d, 0x5f, 0xf5, 0x2c, 0x76, 0x12, 0x2b, 0x9d, 0xcc, 0x2f,
	0xcd, 0x97, 0xe4, 0x10, 0x65, 0xd0, 0xe5, 0x3e, 0x4a, 0x4d, 0x03, 0xaf, 0xc0, 0xaf, 0xa5, 0x58,
	0xb5, 0x3c, 0x76, 0x13, 0x2b, 0x1d, 0xcf, 0x8f, 0x89, 0x41, 0x24, 0xbf, 0x88, 0xe4, 0x4e, 0xec,
	0xe8, 0x90, 0xc1, 0x0b, 0x08, 0x99, 0x52, 0x52, 0xbd, 0x6f, 0x3a, 0x1e, 0x7b, 0x9a, 0x61, 0xa4,
	0x8d, 0xe7, 0x8e, 0x4f, 0x1f, 0xe0, 0xff, 0xdf, 0x05, 0x1c, 0x43, 0xb0, 0xcc, 0x9f, 0xf2, 0xe2,
	0x35, 0x8f, 0xfe, 0xe1, 0x08, 0xdc, 0x2c, 0xcf, 0x16, 0x91, 0xb5, 0x7f, 0x95, 0x8b, 0xe2, 0x25,
	0xb2, 0x31, 0x00, 0x87, 0x2e, 0xf3, 0xc8, 0xc1, 0x10, 0xbc, 0x47, 0x4a, 0x0b, 0x1a, 0xb9, 0xf7,
	0xe1, 0x5b, 0x40, 0x6e, 0xcd, 0xb6, 0xaf, 0xcf, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a,
	0x20, 0x73, 0x9e, 0x53, 0x01, 0x00, 0x00,
}
