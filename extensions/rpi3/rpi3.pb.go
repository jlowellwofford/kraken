// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rpi3.proto

package rpi3

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Pi_Model int32

const (
	Pi_ThreeB     Pi_Model = 0
	Pi_ThreeBPlus Pi_Model = 1
)

var Pi_Model_name = map[int32]string{
	0: "ThreeB",
	1: "ThreeBPlus",
}

var Pi_Model_value = map[string]int32{
	"ThreeB":     0,
	"ThreeBPlus": 1,
}

func (x Pi_Model) String() string {
	return proto.EnumName(Pi_Model_name, int32(x))
}

func (Pi_Model) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4e4dbb54cb6b2047, []int{0, 0}
}

type Pi_PXE int32

const (
	Pi_NONE Pi_PXE = 0
	Pi_WAIT Pi_PXE = 1
	Pi_INIT Pi_PXE = 2
)

var Pi_PXE_name = map[int32]string{
	0: "NONE",
	1: "WAIT",
	2: "INIT",
}

var Pi_PXE_value = map[string]int32{
	"NONE": 0,
	"WAIT": 1,
	"INIT": 2,
}

func (x Pi_PXE) String() string {
	return proto.EnumName(Pi_PXE_name, int32(x))
}

func (Pi_PXE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4e4dbb54cb6b2047, []int{0, 1}
}

type Pi struct {
	Chassis              string   `protobuf:"bytes,1,opt,name=chassis,proto3" json:"chassis,omitempty"`
	Rank                 uint32   `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Model                Pi_Model `protobuf:"varint,3,opt,name=model,proto3,enum=RPi3.Pi_Model" json:"model,omitempty"`
	Pxe                  Pi_PXE   `protobuf:"varint,4,opt,name=pxe,proto3,enum=RPi3.Pi_PXE" json:"pxe,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pi) Reset()         { *m = Pi{} }
func (m *Pi) String() string { return proto.CompactTextString(m) }
func (*Pi) ProtoMessage()    {}
func (*Pi) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e4dbb54cb6b2047, []int{0}
}
func (m *Pi) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pi) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pi.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pi) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pi.Merge(m, src)
}
func (m *Pi) XXX_Size() int {
	return m.Size()
}
func (m *Pi) XXX_DiscardUnknown() {
	xxx_messageInfo_Pi.DiscardUnknown(m)
}

var xxx_messageInfo_Pi proto.InternalMessageInfo

func (m *Pi) GetChassis() string {
	if m != nil {
		return m.Chassis
	}
	return ""
}

func (m *Pi) GetRank() uint32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

func (m *Pi) GetModel() Pi_Model {
	if m != nil {
		return m.Model
	}
	return Pi_ThreeB
}

func (m *Pi) GetPxe() Pi_PXE {
	if m != nil {
		return m.Pxe
	}
	return Pi_NONE
}

func (*Pi) XXX_MessageName() string {
	return "RPi3.Pi"
}
func init() {
	proto.RegisterEnum("RPi3.Pi_Model", Pi_Model_name, Pi_Model_value)
	golang_proto.RegisterEnum("RPi3.Pi_Model", Pi_Model_name, Pi_Model_value)
	proto.RegisterEnum("RPi3.Pi_PXE", Pi_PXE_name, Pi_PXE_value)
	golang_proto.RegisterEnum("RPi3.Pi_PXE", Pi_PXE_name, Pi_PXE_value)
	proto.RegisterType((*Pi)(nil), "RPi3.Pi")
	golang_proto.RegisterType((*Pi)(nil), "RPi3.Pi")
}

func init() { proto.RegisterFile("rpi3.proto", fileDescriptor_4e4dbb54cb6b2047) }
func init() { golang_proto.RegisterFile("rpi3.proto", fileDescriptor_4e4dbb54cb6b2047) }

var fileDescriptor_4e4dbb54cb6b2047 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x3b, 0x4d, 0x1a, 0x75, 0xd0, 0x10, 0xf6, 0x14, 0x44, 0x96, 0x90, 0x7a, 0xc8, 0xc5,
	0x2d, 0x98, 0xa3, 0x27, 0x0b, 0x39, 0xe4, 0x60, 0x5c, 0x96, 0x80, 0xc5, 0x5b, 0x53, 0x63, 0xb2,
	0xd8, 0xba, 0x21, 0x69, 0xc0, 0xc7, 0x13, 0x4f, 0x3d, 0xfa, 0x08, 0x92, 0xbe, 0x88, 0xec, 0x06,
	0xbd, 0x7d, 0xff, 0xfe, 0xdf, 0x32, 0xcc, 0x20, 0xb6, 0x8d, 0x8c, 0x59, 0xd3, 0xaa, 0xbd, 0x22,
	0xb6, 0xe0, 0x32, 0xbe, 0xbc, 0xa9, 0xe4, 0xbe, 0xee, 0x0b, 0xb6, 0x51, 0xbb, 0x45, 0xa5, 0x2a,
	0xb5, 0x30, 0x65, 0xd1, 0xbf, 0x9a, 0x64, 0x82, 0xa1, 0xf1, 0x53, 0xf8, 0x05, 0x38, 0xe5, 0x92,
	0xf8, 0x78, 0xb2, 0xa9, 0xd7, 0x5d, 0x27, 0x3b, 0x1f, 0x02, 0x88, 0xce, 0xc4, 0x5f, 0x24, 0x04,
	0xed, 0x76, 0xfd, 0xfe, 0xe6, 0x4f, 0x03, 0x88, 0x2e, 0x84, 0x61, 0x72, 0x8d, 0xb3, 0x9d, 0x7a,
	0x29, 0xb7, 0xbe, 0x15, 0x40, 0xe4, 0xde, 0xba, 0x4c, 0x4f, 0x66, 0x5c, 0xb2, 0x07, 0xfd, 0x2a,
	0xc6, 0x92, 0x50, 0xb4, 0x9a, 0x8f, 0xd2, 0xb7, 0x8d, 0x73, 0xfe, 0xef, 0xf0, 0x55, 0x22, 0x74,
	0x11, 0xce, 0x71, 0x66, 0x7c, 0x82, 0xe8, 0xe4, 0x75, 0x5b, 0x96, 0x4b, 0x6f, 0x42, 0x5c, 0xc4,
	0x91, 0xf9, 0xb6, 0xef, 0x3c, 0x08, 0xe7, 0x68, 0xf1, 0x55, 0x42, 0x4e, 0xd1, 0xce, 0x1e, 0xb3,
	0xc4, 0x9b, 0x68, 0x7a, 0xba, 0x4f, 0x73, 0x0f, 0x34, 0xa5, 0x59, 0x9a, 0x7b, 0xd3, 0xe5, 0xd5,
	0x61, 0xa0, 0xf0, 0x3d, 0x50, 0xf8, 0x19, 0x28, 0x7c, 0x1e, 0x29, 0x1c, 0x8e, 0x14, 0x9e, 0x1d,
	0x76, 0xa7, 0xaf, 0x53, 0x38, 0x66, 0xd3, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xbc, 0x2d,
	0x6e, 0x2c, 0x01, 0x00, 0x00,
}

func (m *Pi) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pi) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pi) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Pxe != 0 {
		i = encodeVarintRpi3(dAtA, i, uint64(m.Pxe))
		i--
		dAtA[i] = 0x20
	}
	if m.Model != 0 {
		i = encodeVarintRpi3(dAtA, i, uint64(m.Model))
		i--
		dAtA[i] = 0x18
	}
	if m.Rank != 0 {
		i = encodeVarintRpi3(dAtA, i, uint64(m.Rank))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chassis) > 0 {
		i -= len(m.Chassis)
		copy(dAtA[i:], m.Chassis)
		i = encodeVarintRpi3(dAtA, i, uint64(len(m.Chassis)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRpi3(dAtA []byte, offset int, v uint64) int {
	offset -= sovRpi3(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pi) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chassis)
	if l > 0 {
		n += 1 + l + sovRpi3(uint64(l))
	}
	if m.Rank != 0 {
		n += 1 + sovRpi3(uint64(m.Rank))
	}
	if m.Model != 0 {
		n += 1 + sovRpi3(uint64(m.Model))
	}
	if m.Pxe != 0 {
		n += 1 + sovRpi3(uint64(m.Pxe))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRpi3(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRpi3(x uint64) (n int) {
	return sovRpi3(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pi) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRpi3
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Pi: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pi: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chassis", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpi3
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRpi3
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRpi3
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chassis = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			m.Rank = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpi3
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rank |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Model", wireType)
			}
			m.Model = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpi3
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Model |= Pi_Model(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pxe", wireType)
			}
			m.Pxe = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRpi3
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Pxe |= Pi_PXE(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRpi3(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRpi3
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRpi3(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRpi3
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpi3
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRpi3
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRpi3
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRpi3
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRpi3
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRpi3        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRpi3          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRpi3 = fmt.Errorf("proto: unexpected end of group")
)
