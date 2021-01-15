// IPv4.proto: describes IPv4 specific state structures
//
// Author: J. Lowell Wofford <lowell@lanl.gov>
//
// This software is open source software available under the BSD-3 license.
// Copyright (c) 2018, Triad National Security, LLC
// See LICENSE file for details.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: IPv4.proto

package IPv4

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type IPv4 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip     []byte `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Subnet []byte `protobuf:"bytes,3,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *IPv4) Reset() {
	*x = IPv4{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IPv4_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPv4) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4) ProtoMessage() {}

func (x *IPv4) ProtoReflect() protoreflect.Message {
	mi := &file_IPv4_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4.ProtoReflect.Descriptor instead.
func (*IPv4) Descriptor() ([]byte, []int) {
	return file_IPv4_proto_rawDescGZIP(), []int{0}
}

func (x *IPv4) GetIp() []byte {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *IPv4) GetSubnet() []byte {
	if x != nil {
		return x.Subnet
	}
	return nil
}

type Ethernet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iface   string `protobuf:"bytes,1,opt,name=iface,proto3" json:"iface,omitempty"` // interface name, e.g. eth0
	Mac     []byte `protobuf:"bytes,2,opt,name=mac,proto3" json:"mac,omitempty"`     // mac address of the interface
	Mtu     uint32 `protobuf:"varint,3,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Control bool   `protobuf:"varint,4,opt,name=control,proto3" json:"control,omitempty"` // should we assume control (if we can)?
}

func (x *Ethernet) Reset() {
	*x = Ethernet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IPv4_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ethernet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ethernet) ProtoMessage() {}

func (x *Ethernet) ProtoReflect() protoreflect.Message {
	mi := &file_IPv4_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ethernet.ProtoReflect.Descriptor instead.
func (*Ethernet) Descriptor() ([]byte, []int) {
	return file_IPv4_proto_rawDescGZIP(), []int{1}
}

func (x *Ethernet) GetIface() string {
	if x != nil {
		return x.Iface
	}
	return ""
}

func (x *Ethernet) GetMac() []byte {
	if x != nil {
		return x.Mac
	}
	return nil
}

func (x *Ethernet) GetMtu() uint32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

func (x *Ethernet) GetControl() bool {
	if x != nil {
		return x.Control
	}
	return false
}

type IPv4OverEthernet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ifaces         []*IPv4OverEthernet_ConfiguredInterface `protobuf:"bytes,1,rep,name=ifaces,proto3" json:"ifaces,omitempty"`
	Routes         []*IPv4                                 `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	DnsNameservers []*IPv4                                 `protobuf:"bytes,3,rep,name=dns_nameservers,json=dnsNameservers,proto3" json:"dns_nameservers,omitempty"`
	DnsDomains     []*IPv4                                 `protobuf:"bytes,4,rep,name=dns_domains,json=dnsDomains,proto3" json:"dns_domains,omitempty"`
	Hostname       *DNSA                                   `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *IPv4OverEthernet) Reset() {
	*x = IPv4OverEthernet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IPv4_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPv4OverEthernet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4OverEthernet) ProtoMessage() {}

func (x *IPv4OverEthernet) ProtoReflect() protoreflect.Message {
	mi := &file_IPv4_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4OverEthernet.ProtoReflect.Descriptor instead.
func (*IPv4OverEthernet) Descriptor() ([]byte, []int) {
	return file_IPv4_proto_rawDescGZIP(), []int{2}
}

func (x *IPv4OverEthernet) GetIfaces() []*IPv4OverEthernet_ConfiguredInterface {
	if x != nil {
		return x.Ifaces
	}
	return nil
}

func (x *IPv4OverEthernet) GetRoutes() []*IPv4 {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *IPv4OverEthernet) GetDnsNameservers() []*IPv4 {
	if x != nil {
		return x.DnsNameservers
	}
	return nil
}

func (x *IPv4OverEthernet) GetDnsDomains() []*IPv4 {
	if x != nil {
		return x.DnsDomains
	}
	return nil
}

func (x *IPv4OverEthernet) GetHostname() *DNSA {
	if x != nil {
		return x.Hostname
	}
	return nil
}

type DNSA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hostname   string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Domainname string `protobuf:"bytes,2,opt,name=domainname,proto3" json:"domainname,omitempty"`
	Ip         *IPv4  `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *DNSA) Reset() {
	*x = DNSA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IPv4_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSA) ProtoMessage() {}

func (x *DNSA) ProtoReflect() protoreflect.Message {
	mi := &file_IPv4_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSA.ProtoReflect.Descriptor instead.
func (*DNSA) Descriptor() ([]byte, []int) {
	return file_IPv4_proto_rawDescGZIP(), []int{3}
}

func (x *DNSA) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DNSA) GetDomainname() string {
	if x != nil {
		return x.Domainname
	}
	return ""
}

func (x *DNSA) GetIp() *IPv4 {
	if x != nil {
		return x.Ip
	}
	return nil
}

type DNSCNAME struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cname    string `protobuf:"bytes,1,opt,name=cname,proto3" json:"cname,omitempty"`
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *DNSCNAME) Reset() {
	*x = DNSCNAME{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IPv4_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNSCNAME) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNSCNAME) ProtoMessage() {}

func (x *DNSCNAME) ProtoReflect() protoreflect.Message {
	mi := &file_IPv4_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNSCNAME.ProtoReflect.Descriptor instead.
func (*DNSCNAME) Descriptor() ([]byte, []int) {
	return file_IPv4_proto_rawDescGZIP(), []int{4}
}

func (x *DNSCNAME) GetCname() string {
	if x != nil {
		return x.Cname
	}
	return ""
}

func (x *DNSCNAME) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type IPv4OverEthernet_ConfiguredInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eth *Ethernet `protobuf:"bytes,1,opt,name=eth,proto3" json:"eth,omitempty"`
	Ip  *IPv4     `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *IPv4OverEthernet_ConfiguredInterface) Reset() {
	*x = IPv4OverEthernet_ConfiguredInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_IPv4_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPv4OverEthernet_ConfiguredInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPv4OverEthernet_ConfiguredInterface) ProtoMessage() {}

func (x *IPv4OverEthernet_ConfiguredInterface) ProtoReflect() protoreflect.Message {
	mi := &file_IPv4_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPv4OverEthernet_ConfiguredInterface.ProtoReflect.Descriptor instead.
func (*IPv4OverEthernet_ConfiguredInterface) Descriptor() ([]byte, []int) {
	return file_IPv4_proto_rawDescGZIP(), []int{2, 0}
}

func (x *IPv4OverEthernet_ConfiguredInterface) GetEth() *Ethernet {
	if x != nil {
		return x.Eth
	}
	return nil
}

func (x *IPv4OverEthernet_ConfiguredInterface) GetIp() *IPv4 {
	if x != nil {
		return x.Ip
	}
	return nil
}

var File_IPv4_proto protoreflect.FileDescriptor

var file_IPv4_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x49, 0x50, 0x76, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x49, 0x50,
	0x76, 0x34, 0x22, 0x2e, 0x0a, 0x04, 0x49, 0x50, 0x76, 0x34, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x22, 0x5e, 0x0a, 0x08, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x66, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x66, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x22, 0xd9, 0x02, 0x0a, 0x10, 0x49, 0x50, 0x76, 0x34, 0x4f, 0x76, 0x65, 0x72, 0x45,
	0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x42, 0x0a, 0x06, 0x69, 0x66, 0x61, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x2e, 0x49,
	0x50, 0x76, 0x34, 0x4f, 0x76, 0x65, 0x72, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x52, 0x06, 0x69, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x50,
	0x76, 0x34, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12,
	0x33, 0x0a, 0x0f, 0x64, 0x6e, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x2e,
	0x49, 0x50, 0x76, 0x34, 0x52, 0x0e, 0x64, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x2b, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x50, 0x76, 0x34,
	0x2e, 0x49, 0x50, 0x76, 0x34, 0x52, 0x0a, 0x64, 0x6e, 0x73, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x12, 0x26, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x2e, 0x44, 0x4e, 0x53, 0x41, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x53, 0x0a, 0x13, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65,
	0x12, 0x20, 0x0a, 0x03, 0x65, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x49, 0x50, 0x76, 0x34, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x52, 0x03, 0x65,
	0x74, 0x68, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x49, 0x50, 0x76, 0x34, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x52, 0x02, 0x69, 0x70, 0x22, 0x5e,
	0x0a, 0x04, 0x44, 0x4e, 0x53, 0x41, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x49, 0x50, 0x76, 0x34, 0x2e, 0x49, 0x50, 0x76, 0x34, 0x52, 0x02, 0x69, 0x70, 0x22, 0x3c,
	0x0a, 0x08, 0x44, 0x4e, 0x53, 0x43, 0x4e, 0x41, 0x4d, 0x45, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x3b, 0x49, 0x50, 0x76, 0x34, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_IPv4_proto_rawDescOnce sync.Once
	file_IPv4_proto_rawDescData = file_IPv4_proto_rawDesc
)

func file_IPv4_proto_rawDescGZIP() []byte {
	file_IPv4_proto_rawDescOnce.Do(func() {
		file_IPv4_proto_rawDescData = protoimpl.X.CompressGZIP(file_IPv4_proto_rawDescData)
	})
	return file_IPv4_proto_rawDescData
}

var file_IPv4_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_IPv4_proto_goTypes = []interface{}{
	(*IPv4)(nil),             // 0: IPv4.IPv4
	(*Ethernet)(nil),         // 1: IPv4.Ethernet
	(*IPv4OverEthernet)(nil), // 2: IPv4.IPv4OverEthernet
	(*DNSA)(nil),             // 3: IPv4.DNSA
	(*DNSCNAME)(nil),         // 4: IPv4.DNSCNAME
	(*IPv4OverEthernet_ConfiguredInterface)(nil), // 5: IPv4.IPv4OverEthernet.ConfiguredInterface
}
var file_IPv4_proto_depIdxs = []int32{
	5, // 0: IPv4.IPv4OverEthernet.ifaces:type_name -> IPv4.IPv4OverEthernet.ConfiguredInterface
	0, // 1: IPv4.IPv4OverEthernet.routes:type_name -> IPv4.IPv4
	0, // 2: IPv4.IPv4OverEthernet.dns_nameservers:type_name -> IPv4.IPv4
	0, // 3: IPv4.IPv4OverEthernet.dns_domains:type_name -> IPv4.IPv4
	3, // 4: IPv4.IPv4OverEthernet.hostname:type_name -> IPv4.DNSA
	0, // 5: IPv4.DNSA.ip:type_name -> IPv4.IPv4
	1, // 6: IPv4.IPv4OverEthernet.ConfiguredInterface.eth:type_name -> IPv4.Ethernet
	0, // 7: IPv4.IPv4OverEthernet.ConfiguredInterface.ip:type_name -> IPv4.IPv4
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_IPv4_proto_init() }
func file_IPv4_proto_init() {
	if File_IPv4_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_IPv4_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPv4); i {
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
		file_IPv4_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ethernet); i {
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
		file_IPv4_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPv4OverEthernet); i {
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
		file_IPv4_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSA); i {
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
		file_IPv4_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNSCNAME); i {
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
		file_IPv4_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPv4OverEthernet_ConfiguredInterface); i {
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
			RawDescriptor: file_IPv4_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_IPv4_proto_goTypes,
		DependencyIndexes: file_IPv4_proto_depIdxs,
		MessageInfos:      file_IPv4_proto_msgTypes,
	}.Build()
	File_IPv4_proto = out.File
	file_IPv4_proto_rawDesc = nil
	file_IPv4_proto_goTypes = nil
	file_IPv4_proto_depIdxs = nil
}
