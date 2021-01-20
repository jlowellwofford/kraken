// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Node.proto

package proto

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type Node_RunState int32

const (
	Node_UNKNOWN Node_RunState = 0
	Node_INIT    Node_RunState = 1
	Node_SYNC    Node_RunState = 2
	Node_ERROR   Node_RunState = 3
)

var Node_RunState_name = map[int32]string{
	0: "UNKNOWN",
	1: "INIT",
	2: "SYNC",
	3: "ERROR",
}

var Node_RunState_value = map[string]int32{
	"UNKNOWN": 0,
	"INIT":    1,
	"SYNC":    2,
	"ERROR":   3,
}

func (x Node_RunState) String() string {
	return proto.EnumName(Node_RunState_name, int32(x))
}

func (Node_RunState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ade4b1ea954e8f8c, []int{1, 0}
}

type Node_PhysState int32

const (
	Node_PHYS_UNKNOWN Node_PhysState = 0
	Node_POWER_OFF    Node_PhysState = 1
	Node_POWER_ON     Node_PhysState = 2
	Node_POWER_CYCLE  Node_PhysState = 3
	Node_PHYS_HANG    Node_PhysState = 4
	Node_PHYS_ERROR   Node_PhysState = 5
)

var Node_PhysState_name = map[int32]string{
	0: "PHYS_UNKNOWN",
	1: "POWER_OFF",
	2: "POWER_ON",
	3: "POWER_CYCLE",
	4: "PHYS_HANG",
	5: "PHYS_ERROR",
}

var Node_PhysState_value = map[string]int32{
	"PHYS_UNKNOWN": 0,
	"POWER_OFF":    1,
	"POWER_ON":     2,
	"POWER_CYCLE":  3,
	"PHYS_HANG":    4,
	"PHYS_ERROR":   5,
}

func (x Node_PhysState) String() string {
	return proto.EnumName(Node_PhysState_name, int32(x))
}

func (Node_PhysState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ade4b1ea954e8f8c, []int{1, 1}
}

type NodeList struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeList) Reset()         { *m = NodeList{} }
func (m *NodeList) String() string { return proto.CompactTextString(m) }
func (*NodeList) ProtoMessage()    {}
func (*NodeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_ade4b1ea954e8f8c, []int{0}
}
func (m *NodeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeList.Unmarshal(m, b)
}
func (m *NodeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeList.Marshal(b, m, deterministic)
}
func (m *NodeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeList.Merge(m, src)
}
func (m *NodeList) XXX_Size() int {
	return xxx_messageInfo_NodeList.Size(m)
}
func (m *NodeList) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeList.DiscardUnknown(m)
}

var xxx_messageInfo_NodeList proto.InternalMessageInfo

func (m *NodeList) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	Id                   *NodeID            `protobuf:"bytes,1,opt,name=id,proto3,customtype=NodeID" json:"id,omitempty"`
	Nodename             string             `protobuf:"bytes,2,opt,name=nodename,proto3" json:"nodename,omitempty"`
	RunState             Node_RunState      `protobuf:"varint,3,opt,name=run_state,json=runState,proto3,enum=proto.Node_RunState" json:"run_state,omitempty"`
	PhysState            Node_PhysState     `protobuf:"varint,4,opt,name=phys_state,json=physState,proto3,enum=proto.Node_PhysState" json:"phys_state,omitempty"`
	Arch                 string             `protobuf:"bytes,5,opt,name=arch,proto3" json:"arch,omitempty"`
	Platform             string             `protobuf:"bytes,6,opt,name=platform,proto3" json:"platform,omitempty"`
	ParentId             *NodeID            `protobuf:"bytes,7,opt,name=parent_id,json=parentId,proto3,customtype=NodeID" json:"parent_id,omitempty"`
	Services             []*ServiceInstance `protobuf:"bytes,14,rep,name=services,proto3" json:"services,omitempty"`
	Extensions           []*types.Any       `protobuf:"bytes,15,rep,name=extensions,proto3" json:"extensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_ade4b1ea954e8f8c, []int{1}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetNodename() string {
	if m != nil {
		return m.Nodename
	}
	return ""
}

func (m *Node) GetRunState() Node_RunState {
	if m != nil {
		return m.RunState
	}
	return Node_UNKNOWN
}

func (m *Node) GetPhysState() Node_PhysState {
	if m != nil {
		return m.PhysState
	}
	return Node_PHYS_UNKNOWN
}

func (m *Node) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func (m *Node) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *Node) GetServices() []*ServiceInstance {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *Node) GetExtensions() []*types.Any {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.Node_RunState", Node_RunState_name, Node_RunState_value)
	proto.RegisterEnum("proto.Node_PhysState", Node_PhysState_name, Node_PhysState_value)
	proto.RegisterType((*NodeList)(nil), "proto.NodeList")
	proto.RegisterType((*Node)(nil), "proto.Node")
}

func init() { proto.RegisterFile("Node.proto", fileDescriptor_ade4b1ea954e8f8c) }

var fileDescriptor_ade4b1ea954e8f8c = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x25, 0x69, 0xb2, 0x3a, 0x5f, 0x4b, 0x67, 0x59, 0x1b, 0x0a, 0xbd, 0x50, 0x7a, 0xa1, 0x97,
	0x65, 0xa2, 0x20, 0x2e, 0x9c, 0xb6, 0xd2, 0xb1, 0xc0, 0x94, 0x56, 0x2e, 0x68, 0x2a, 0x97, 0x2a,
	0x6b, 0xbc, 0x34, 0x68, 0xb5, 0xa3, 0xd8, 0x45, 0xf4, 0x7f, 0xf0, 0xeb, 0x38, 0xf0, 0x5b, 0x90,
	0xed, 0xb4, 0xab, 0xe0, 0xe4, 0xf7, 0xfc, 0xde, 0xd3, 0xf7, 0xf9, 0x19, 0x20, 0x11, 0x19, 0x8b,
	0xca, 0x4a, 0x28, 0x41, 0x7c, 0x73, 0x74, 0xcf, 0xf2, 0x42, 0xad, 0x36, 0x77, 0xd1, 0x52, 0xac,
	0xcf, 0x73, 0x91, 0x8b, 0x73, 0x73, 0x7d, 0xb7, 0xb9, 0x37, 0xcc, 0x10, 0x83, 0x6c, 0xaa, 0xfb,
	0x3c, 0x17, 0x22, 0x7f, 0x60, 0x8f, 0xae, 0x94, 0x6f, 0x6b, 0xe9, 0x74, 0xc6, 0xaa, 0x1f, 0xc5,
	0x92, 0xc5, 0x5c, 0xaa, 0x94, 0x2f, 0xeb, 0x39, 0xfd, 0x33, 0x40, 0x7a, 0xea, 0x4d, 0x21, 0x15,
	0x79, 0x09, 0x3e, 0x17, 0x19, 0x93, 0xa1, 0xd3, 0x6b, 0x0c, 0x5a, 0xc3, 0x96, 0xb5, 0x44, 0x5a,
	0xa7, 0x56, 0xe9, 0xff, 0xf2, 0xc0, 0xd3, 0x9c, 0x74, 0xc1, 0x2d, 0xb2, 0xd0, 0xe9, 0x39, 0x83,
	0xf6, 0x25, 0xfc, 0xfe, 0xf3, 0xe2, 0x48, 0xdf, 0xc6, 0x1f, 0xa8, 0x5b, 0x64, 0xa4, 0x0b, 0x48,
	0xbb, 0x79, 0xba, 0x66, 0xa1, 0xdb, 0x73, 0x06, 0x01, 0xdd, 0x73, 0xf2, 0x1a, 0x82, 0x6a, 0xc3,
	0x17, 0x52, 0xa5, 0x8a, 0x85, 0x8d, 0x9e, 0x33, 0xe8, 0x0c, 0x4f, 0x0e, 0xe6, 0x44, 0x74, 0xc3,
	0x67, 0x5a, 0xa3, 0xa8, 0xaa, 0x11, 0x79, 0x0b, 0x50, 0xae, 0xb6, 0xb2, 0xce, 0x78, 0x26, 0x73,
	0x7a, 0x98, 0x99, 0xae, 0xb6, 0xd2, 0x86, 0x82, 0x72, 0x07, 0x09, 0x01, 0x2f, 0xad, 0x96, 0xab,
	0xd0, 0x37, 0x0b, 0x18, 0xac, 0x17, 0x2b, 0x1f, 0x52, 0x75, 0x2f, 0xaa, 0x75, 0x78, 0x64, 0x17,
	0xdb, 0x71, 0xf2, 0x0a, 0x82, 0x32, 0xad, 0x18, 0x57, 0x8b, 0x22, 0x0b, 0x9b, 0xff, 0xbd, 0x0b,
	0x59, 0x31, 0xce, 0xc8, 0x10, 0x90, 0xb4, 0x55, 0xca, 0xb0, 0x63, 0x8a, 0x7a, 0x56, 0x2f, 0xf3,
	0x4f, 0xc3, 0x74, 0xef, 0xd3, 0x4f, 0x60, 0x3f, 0x15, 0xe3, 0xb2, 0x10, 0x5c, 0x86, 0xc7, 0x26,
	0x75, 0x12, 0xd9, 0xcf, 0x8a, 0x76, 0x9f, 0x15, 0x5d, 0xf0, 0x2d, 0x3d, 0xf0, 0xf5, 0xdf, 0x01,
	0xda, 0xd5, 0x41, 0x5a, 0xd0, 0xfc, 0x9a, 0x7c, 0x4e, 0x26, 0xb7, 0x09, 0x7e, 0x42, 0x10, 0x78,
	0x71, 0x12, 0x7f, 0xc1, 0x8e, 0x46, 0xb3, 0x79, 0x32, 0xc2, 0x2e, 0x09, 0xc0, 0x1f, 0x53, 0x3a,
	0xa1, 0xb8, 0xd1, 0xff, 0x0e, 0xc1, 0xbe, 0x12, 0x82, 0xa1, 0x3d, 0xbd, 0x9e, 0xcf, 0x16, 0x8f,
	0xe9, 0xa7, 0x10, 0x4c, 0x27, 0xb7, 0x63, 0xba, 0x98, 0x5c, 0x5d, 0x61, 0x87, 0xb4, 0x01, 0xd5,
	0x34, 0xc1, 0x2e, 0x39, 0x86, 0x96, 0x65, 0xa3, 0xf9, 0xe8, 0x66, 0x8c, 0x1b, 0xc6, 0xad, 0xf3,
	0xd7, 0x17, 0xc9, 0x47, 0xec, 0x91, 0x0e, 0x80, 0xa1, 0x76, 0x96, 0xff, 0xc9, 0x43, 0x08, 0x77,
	0x2e, 0x83, 0x6f, 0xcd, 0xe8, 0xbd, 0x7d, 0xc7, 0x91, 0x39, 0xde, 0xfc, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0x15, 0xa3, 0xae, 0xcd, 0x02, 0x00, 0x00,
}
