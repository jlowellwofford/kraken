// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rfthermal.proto

package rfthermal

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

type Temp_CPUThermalState int32

const (
	Temp_CPU_TEMP_NONE     Temp_CPUThermalState = 0
	Temp_CPU_TEMP_NORMAL   Temp_CPUThermalState = 1
	Temp_CPU_TEMP_HIGH     Temp_CPUThermalState = 2
	Temp_CPU_TEMP_CRITICAL Temp_CPUThermalState = 3
)

var Temp_CPUThermalState_name = map[int32]string{
	0: "CPU_TEMP_NONE",
	1: "CPU_TEMP_NORMAL",
	2: "CPU_TEMP_HIGH",
	3: "CPU_TEMP_CRITICAL",
}

var Temp_CPUThermalState_value = map[string]int32{
	"CPU_TEMP_NONE":     0,
	"CPU_TEMP_NORMAL":   1,
	"CPU_TEMP_HIGH":     2,
	"CPU_TEMP_CRITICAL": 3,
}

func (x Temp_CPUThermalState) String() string {
	return proto.EnumName(Temp_CPUThermalState_name, int32(x))
}

func (Temp_CPUThermalState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb37779bcf7d5ac7, []int{0, 0}
}

type Temp struct {
	State                Temp_CPUThermalState `protobuf:"varint,1,opt,name=state,proto3,enum=RFThermal.Temp_CPUThermalState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Temp) Reset()         { *m = Temp{} }
func (m *Temp) String() string { return proto.CompactTextString(m) }
func (*Temp) ProtoMessage()    {}
func (*Temp) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb37779bcf7d5ac7, []int{0}
}
func (m *Temp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Temp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Temp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Temp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Temp.Merge(m, src)
}
func (m *Temp) XXX_Size() int {
	return m.Size()
}
func (m *Temp) XXX_DiscardUnknown() {
	xxx_messageInfo_Temp.DiscardUnknown(m)
}

var xxx_messageInfo_Temp proto.InternalMessageInfo

func (m *Temp) GetState() Temp_CPUThermalState {
	if m != nil {
		return m.State
	}
	return Temp_CPU_TEMP_NONE
}

func (*Temp) XXX_MessageName() string {
	return "RFThermal.Temp"
}
func init() {
	proto.RegisterEnum("RFThermal.Temp_CPUThermalState", Temp_CPUThermalState_name, Temp_CPUThermalState_value)
	golang_proto.RegisterEnum("RFThermal.Temp_CPUThermalState", Temp_CPUThermalState_name, Temp_CPUThermalState_value)
	proto.RegisterType((*Temp)(nil), "RFThermal.Temp")
	golang_proto.RegisterType((*Temp)(nil), "RFThermal.Temp")
}

func init() { proto.RegisterFile("rfthermal.proto", fileDescriptor_cb37779bcf7d5ac7) }
func init() { golang_proto.RegisterFile("rfthermal.proto", fileDescriptor_cb37779bcf7d5ac7) }

var fileDescriptor_cb37779bcf7d5ac7 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x2b, 0xc9,
	0x48, 0x2d, 0xca, 0x4d, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0c, 0x72, 0x0b,
	0x81, 0x08, 0x48, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7,
	0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x55, 0x24, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1,
	0xa9, 0xb4, 0x88, 0x91, 0x8b, 0x25, 0x24, 0x35, 0xb7, 0x40, 0xc8, 0x94, 0x8b, 0xb5, 0xb8, 0x24,
	0xb1, 0x24, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x5e, 0x0f, 0x6e, 0xa4, 0x1e, 0x48,
	0x5e, 0xcf, 0x39, 0x20, 0x14, 0xca, 0x0f, 0x06, 0x29, 0x0b, 0x82, 0xa8, 0x56, 0x4a, 0xe6, 0xe2,
	0x47, 0x93, 0x11, 0x12, 0xe4, 0xe2, 0x75, 0x0e, 0x08, 0x8d, 0x0f, 0x71, 0xf5, 0x0d, 0x88, 0xf7,
	0xf3, 0xf7, 0x73, 0x15, 0x60, 0x10, 0x12, 0x06, 0xab, 0x82, 0x09, 0x05, 0xf9, 0x3a, 0xfa, 0x08,
	0x30, 0xa2, 0xa8, 0xf3, 0xf0, 0x74, 0xf7, 0x10, 0x60, 0x12, 0x12, 0xe5, 0x12, 0x84, 0x0b, 0x39,
	0x07, 0x79, 0x86, 0x78, 0x3a, 0x3b, 0xfa, 0x08, 0x30, 0x3b, 0x29, 0x9e, 0x78, 0x24, 0xc7, 0x78,
	0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x07, 0x1e, 0xcb, 0x31, 0x9e, 0x78, 0x2c, 0xc7,
	0x18, 0xc5, 0xad, 0x67, 0x0d, 0x0f, 0x87, 0x24, 0x36, 0xb0, 0x77, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0x07, 0xd4, 0xa9, 0x1b, 0x01, 0x00, 0x00,
}

func (m *Temp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Temp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Temp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.State != 0 {
		i = encodeVarintRfthermal(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRfthermal(dAtA []byte, offset int, v uint64) int {
	offset -= sovRfthermal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Temp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovRfthermal(uint64(m.State))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRfthermal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRfthermal(x uint64) (n int) {
	return sovRfthermal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Temp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRfthermal
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
			return fmt.Errorf("proto: Temp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Temp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRfthermal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= Temp_CPUThermalState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRfthermal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRfthermal
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
func skipRfthermal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRfthermal
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
					return 0, ErrIntOverflowRfthermal
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
					return 0, ErrIntOverflowRfthermal
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
				return 0, ErrInvalidLengthRfthermal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRfthermal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRfthermal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRfthermal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRfthermal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRfthermal = fmt.Errorf("proto: unexpected end of group")
)
