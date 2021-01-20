// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pipxe.proto

package pipxe

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Config struct {
	SrvIfaceUrl          string   `protobuf:"bytes,1,opt,name=srv_iface_url,json=srvIfaceUrl,proto3" json:"srv_iface_url,omitempty"`
	SrvIpUrl             string   `protobuf:"bytes,2,opt,name=srv_ip_url,json=srvIpUrl,proto3" json:"srv_ip_url,omitempty"`
	IpUrl                string   `protobuf:"bytes,3,opt,name=ip_url,json=ipUrl,proto3" json:"ip_url,omitempty"`
	NmUrl                string   `protobuf:"bytes,4,opt,name=nm_url,json=nmUrl,proto3" json:"nm_url,omitempty"`
	MacUrl               string   `protobuf:"bytes,5,opt,name=mac_url,json=macUrl,proto3" json:"mac_url,omitempty"`
	SubnetUrl            string   `protobuf:"bytes,6,opt,name=subnet_url,json=subnetUrl,proto3" json:"subnet_url,omitempty"`
	TftpDir              string   `protobuf:"bytes,7,opt,name=tftp_dir,json=tftpDir,proto3" json:"tftp_dir,omitempty"`
	ArpDeadline          string   `protobuf:"bytes,8,opt,name=arp_deadline,json=arpDeadline,proto3" json:"arp_deadline,omitempty"`
	DhcpRetry            uint32   `protobuf:"varint,9,opt,name=dhcp_retry,json=dhcpRetry,proto3" json:"dhcp_retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_250beee9db110b1e, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetSrvIfaceUrl() string {
	if m != nil {
		return m.SrvIfaceUrl
	}
	return ""
}

func (m *Config) GetSrvIpUrl() string {
	if m != nil {
		return m.SrvIpUrl
	}
	return ""
}

func (m *Config) GetIpUrl() string {
	if m != nil {
		return m.IpUrl
	}
	return ""
}

func (m *Config) GetNmUrl() string {
	if m != nil {
		return m.NmUrl
	}
	return ""
}

func (m *Config) GetMacUrl() string {
	if m != nil {
		return m.MacUrl
	}
	return ""
}

func (m *Config) GetSubnetUrl() string {
	if m != nil {
		return m.SubnetUrl
	}
	return ""
}

func (m *Config) GetTftpDir() string {
	if m != nil {
		return m.TftpDir
	}
	return ""
}

func (m *Config) GetArpDeadline() string {
	if m != nil {
		return m.ArpDeadline
	}
	return ""
}

func (m *Config) GetDhcpRetry() uint32 {
	if m != nil {
		return m.DhcpRetry
	}
	return 0
}

func init() {
	proto.RegisterType((*Config)(nil), "PiPXE.Config")
}

func init() { proto.RegisterFile("pipxe.proto", fileDescriptor_250beee9db110b1e) }

var fileDescriptor_250beee9db110b1e = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0xd0, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x71, 0x52, 0xcd, 0x26, 0x3b, 0xb5, 0x97, 0x80, 0x18, 0x41, 0xa1, 0xf6, 0xd4, 0x53,
	0x2f, 0x1e, 0xbd, 0x69, 0x3d, 0x78, 0x2b, 0x01, 0x41, 0xbc, 0x2c, 0xdb, 0x64, 0xa3, 0x0b, 0xc9,
	0x66, 0x98, 0x6c, 0x8b, 0x3e, 0x84, 0xef, 0x2c, 0x33, 0x69, 0x8f, 0xf9, 0xfd, 0x3f, 0x02, 0x3b,
	0x30, 0x47, 0x8f, 0x3f, 0x6e, 0x83, 0x34, 0xc4, 0xa1, 0x48, 0x77, 0x7e, 0xf7, 0xf1, 0xba, 0xfa,
	0x9b, 0x81, 0x7a, 0x19, 0x42, 0xeb, 0xbf, 0x8a, 0x15, 0x2c, 0x46, 0x3a, 0x1a, 0xdf, 0xda, 0xda,
	0x99, 0x03, 0x75, 0x65, 0xb2, 0x4c, 0xd6, 0xba, 0x9a, 0x8f, 0x74, 0x7c, 0x63, 0x7b, 0xa7, 0xae,
	0xb8, 0x03, 0x90, 0x0d, 0xca, 0x60, 0x26, 0x83, 0x9c, 0x07, 0xc8, 0xf5, 0x1a, 0xd4, 0xa9, 0x5c,
	0x48, 0x49, 0xfd, 0x99, 0x43, 0x2f, 0x7c, 0x39, 0x71, 0xe8, 0x99, 0x6f, 0x20, 0xeb, 0x6d, 0x2d,
	0x9e, 0x8a, 0xab, 0xde, 0xd6, 0x1c, 0xee, 0x01, 0xc6, 0xc3, 0x3e, 0xb8, 0x28, 0x4d, 0x49, 0xd3,
	0x93, 0x70, 0xbe, 0x85, 0x3c, 0xb6, 0x11, 0x4d, 0xe3, 0xa9, 0xcc, 0x24, 0x66, 0xfc, 0xbd, 0xf5,
	0x54, 0x3c, 0xc0, 0x95, 0x25, 0x34, 0x8d, 0xb3, 0x4d, 0xe7, 0x83, 0x2b, 0xf3, 0xe9, 0x05, 0x96,
	0x70, 0x7b, 0x22, 0xfe, 0x79, 0xf3, 0x5d, 0xa3, 0x21, 0x17, 0xe9, 0xb7, 0xd4, 0xcb, 0x64, 0xbd,
	0xa8, 0x34, 0x4b, 0xc5, 0xf0, 0xac, 0x3f, 0xb3, 0xcd, 0x93, 0xdc, 0x69, 0xaf, 0xe4, 0x50, 0x8f,
	0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x1a, 0x36, 0x4e, 0x37, 0x01, 0x00, 0x00,
}