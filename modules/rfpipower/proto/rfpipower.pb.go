// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rfpipower.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RFPiPowerConfig struct {
	Servers              map[string]*RFPiPowerServer `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Tick                 string                      `protobuf:"bytes,2,opt,name=tick,proto3" json:"tick,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RFPiPowerConfig) Reset()         { *m = RFPiPowerConfig{} }
func (m *RFPiPowerConfig) String() string { return proto.CompactTextString(m) }
func (*RFPiPowerConfig) ProtoMessage()    {}
func (*RFPiPowerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_rfpipower_659fa86591b832a1, []int{0}
}
func (m *RFPiPowerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RFPiPowerConfig.Unmarshal(m, b)
}
func (m *RFPiPowerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RFPiPowerConfig.Marshal(b, m, deterministic)
}
func (dst *RFPiPowerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RFPiPowerConfig.Merge(dst, src)
}
func (m *RFPiPowerConfig) XXX_Size() int {
	return xxx_messageInfo_RFPiPowerConfig.Size(m)
}
func (m *RFPiPowerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RFPiPowerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RFPiPowerConfig proto.InternalMessageInfo

func (m *RFPiPowerConfig) GetServers() map[string]*RFPiPowerServer {
	if m != nil {
		return m.Servers
	}
	return nil
}

func (m *RFPiPowerConfig) GetTick() string {
	if m != nil {
		return m.Tick
	}
	return ""
}

type RFPiPowerServer struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RFPiPowerServer) Reset()         { *m = RFPiPowerServer{} }
func (m *RFPiPowerServer) String() string { return proto.CompactTextString(m) }
func (*RFPiPowerServer) ProtoMessage()    {}
func (*RFPiPowerServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_rfpipower_659fa86591b832a1, []int{1}
}
func (m *RFPiPowerServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RFPiPowerServer.Unmarshal(m, b)
}
func (m *RFPiPowerServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RFPiPowerServer.Marshal(b, m, deterministic)
}
func (dst *RFPiPowerServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RFPiPowerServer.Merge(dst, src)
}
func (m *RFPiPowerServer) XXX_Size() int {
	return xxx_messageInfo_RFPiPowerServer.Size(m)
}
func (m *RFPiPowerServer) XXX_DiscardUnknown() {
	xxx_messageInfo_RFPiPowerServer.DiscardUnknown(m)
}

var xxx_messageInfo_RFPiPowerServer proto.InternalMessageInfo

func (m *RFPiPowerServer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RFPiPowerServer) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RFPiPowerServer) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*RFPiPowerConfig)(nil), "proto.RFPiPowerConfig")
	proto.RegisterMapType((map[string]*RFPiPowerServer)(nil), "proto.RFPiPowerConfig.ServersEntry")
	proto.RegisterType((*RFPiPowerServer)(nil), "proto.RFPiPowerServer")
}

func init() { proto.RegisterFile("rfpipower.proto", fileDescriptor_rfpipower_659fa86591b832a1) }

var fileDescriptor_rfpipower_659fa86591b832a1 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x2b, 0xc8,
	0x2c, 0xc8, 0x2f, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0x3b, 0x18, 0xb9, 0xf8, 0x83, 0xdc, 0x02, 0x32, 0x03, 0x40, 0x52, 0xce, 0xf9, 0x79, 0x69, 0x99,
	0xe9, 0x42, 0xb6, 0x5c, 0xec, 0xc5, 0xa9, 0x45, 0x65, 0xa9, 0x45, 0xc5, 0x12, 0x8c, 0x0a, 0xcc,
	0x1a, 0xdc, 0x46, 0xca, 0x10, 0x3d, 0x7a, 0x68, 0x0a, 0xf5, 0x82, 0x21, 0xaa, 0x5c, 0xf3, 0x4a,
	0x8a, 0x2a, 0x83, 0x60, 0x7a, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x32, 0x93, 0xb3, 0x25, 0x98, 0x14,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa9, 0x20, 0x2e, 0x1e, 0x64, 0xc5, 0x42, 0x02, 0x5c, 0xcc,
	0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x60, 0x25, 0x20, 0xa6, 0x90, 0x0e, 0x17, 0x6b, 0x59, 0x62, 0x4e,
	0x69, 0x2a, 0x58, 0x1b, 0xb7, 0x91, 0x18, 0xba, 0x95, 0x10, 0xed, 0x41, 0x10, 0x45, 0x56, 0x4c,
	0x16, 0x8c, 0x4a, 0x9e, 0x48, 0x2e, 0x87, 0xc8, 0x82, 0xac, 0xce, 0x4b, 0xcc, 0x4d, 0x85, 0x9a,
	0x0b, 0x66, 0x0b, 0xf1, 0x71, 0x31, 0x65, 0x16, 0x40, 0x1d, 0xc3, 0x94, 0x59, 0x00, 0x52, 0x53,
	0x90, 0x5f, 0x54, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x66, 0x27, 0xb1, 0x81, 0x2d,
	0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x71, 0x5b, 0x3b, 0x26, 0x01, 0x00, 0x00,
}