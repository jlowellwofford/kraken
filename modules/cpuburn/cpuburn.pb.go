// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cpuburn.proto

package cpuburn

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

type Config struct {
	TempSensor           string   `protobuf:"bytes,1,opt,name=temp_sensor,json=tempSensor,proto3" json:"temp_sensor,omitempty"`
	ThermalThrottle      bool     `protobuf:"varint,2,opt,name=thermal_throttle,json=thermalThrottle,proto3" json:"thermal_throttle,omitempty"`
	ThermalPoll          uint32   `protobuf:"varint,3,opt,name=thermal_poll,json=thermalPoll,proto3" json:"thermal_poll,omitempty"`
	ThermalResume        uint32   `protobuf:"varint,4,opt,name=thermal_resume,json=thermalResume,proto3" json:"thermal_resume,omitempty"`
	ThermalCrit          uint32   `protobuf:"varint,5,opt,name=thermal_crit,json=thermalCrit,proto3" json:"thermal_crit,omitempty"`
	Workers              uint32   `protobuf:"varint,6,opt,name=workers,proto3" json:"workers,omitempty"`
	WorkersThrottled     uint32   `protobuf:"varint,7,opt,name=workers_throttled,json=workersThrottled,proto3" json:"workers_throttled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_95f28937a9fd0ca5, []int{0}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Config.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return m.Size()
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetTempSensor() string {
	if m != nil {
		return m.TempSensor
	}
	return ""
}

func (m *Config) GetThermalThrottle() bool {
	if m != nil {
		return m.ThermalThrottle
	}
	return false
}

func (m *Config) GetThermalPoll() uint32 {
	if m != nil {
		return m.ThermalPoll
	}
	return 0
}

func (m *Config) GetThermalResume() uint32 {
	if m != nil {
		return m.ThermalResume
	}
	return 0
}

func (m *Config) GetThermalCrit() uint32 {
	if m != nil {
		return m.ThermalCrit
	}
	return 0
}

func (m *Config) GetWorkers() uint32 {
	if m != nil {
		return m.Workers
	}
	return 0
}

func (m *Config) GetWorkersThrottled() uint32 {
	if m != nil {
		return m.WorkersThrottled
	}
	return 0
}

func (*Config) XXX_MessageName() string {
	return "CPUBurn.Config"
}
func init() {
	proto.RegisterType((*Config)(nil), "CPUBurn.Config")
	golang_proto.RegisterType((*Config)(nil), "CPUBurn.Config")
}

func init() { proto.RegisterFile("cpuburn.proto", fileDescriptor_95f28937a9fd0ca5) }
func init() { golang_proto.RegisterFile("cpuburn.proto", fileDescriptor_95f28937a9fd0ca5) }

var fileDescriptor_95f28937a9fd0ca5 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xe5, 0x02, 0x09, 0x75, 0x09, 0x14, 0x4f, 0x16, 0x83, 0x1b, 0x90, 0x90, 0x82, 0x10,
	0xe9, 0xc0, 0xc8, 0xd6, 0xbc, 0x40, 0x15, 0x60, 0x61, 0xa9, 0x48, 0xea, 0x26, 0x11, 0x4e, 0x2e,
	0xba, 0xd8, 0xe2, 0x15, 0x78, 0x2c, 0xc6, 0x8e, 0x3c, 0x02, 0x4a, 0x5f, 0x04, 0x61, 0x1c, 0x04,
	0xdb, 0xfd, 0xdf, 0xff, 0xdf, 0xe9, 0xf4, 0xd3, 0x20, 0x6f, 0x4d, 0x66, 0xb0, 0x89, 0x5b, 0x04,
	0x0d, 0xcc, 0x4f, 0x96, 0x8f, 0x0b, 0x83, 0xcd, 0xd9, 0x4d, 0x51, 0xe9, 0xd2, 0x64, 0x71, 0x0e,
	0xf5, 0xbc, 0x80, 0x02, 0xe6, 0xd6, 0xcf, 0xcc, 0xc6, 0x2a, 0x2b, 0xec, 0xf4, 0xb3, 0x77, 0xf1,
	0x36, 0xa2, 0x5e, 0x02, 0xcd, 0xa6, 0x2a, 0xd8, 0x8c, 0x4e, 0xb4, 0xac, 0xdb, 0x55, 0x27, 0x9b,
	0x0e, 0x90, 0x93, 0x90, 0x44, 0xe3, 0x94, 0x7e, 0xa3, 0x7b, 0x4b, 0xd8, 0x15, 0x9d, 0xea, 0x52,
	0x62, 0xfd, 0xac, 0x56, 0xba, 0x44, 0xd0, 0x5a, 0x49, 0x3e, 0x0a, 0x49, 0x74, 0x98, 0x9e, 0x38,
	0xfe, 0xe0, 0x30, 0x3b, 0xa7, 0x47, 0x43, 0xb4, 0x05, 0xa5, 0xf8, 0x5e, 0x48, 0xa2, 0x20, 0x9d,
	0x38, 0xb6, 0x04, 0xa5, 0xd8, 0x25, 0x3d, 0x1e, 0x22, 0x28, 0x3b, 0x53, 0x4b, 0xbe, 0x6f, 0x43,
	0x81, 0xa3, 0xa9, 0x85, 0x7f, 0x2f, 0xe5, 0x58, 0x69, 0x7e, 0xf0, 0xef, 0x52, 0x82, 0x95, 0x66,
	0x9c, 0xfa, 0xaf, 0x80, 0x2f, 0x12, 0x3b, 0xee, 0x59, 0x77, 0x90, 0xec, 0x9a, 0x9e, 0xba, 0xf1,
	0xf7, 0xe3, 0x35, 0xf7, 0x6d, 0x66, 0xea, 0x8c, 0xe1, 0xe5, 0xf5, 0x62, 0xb6, 0xed, 0x05, 0xf9,
	0xe8, 0x05, 0xf9, 0xec, 0x05, 0x79, 0xdf, 0x09, 0xb2, 0xdd, 0x09, 0xf2, 0x34, 0x8e, 0xef, 0x5c,
	0xd3, 0x99, 0x67, 0x2b, 0xbb, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xdf, 0x9e, 0x27, 0x7b,
	0x01, 0x00, 0x00,
}

func (m *Config) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Config) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Config) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.WorkersThrottled != 0 {
		i = encodeVarintCpuburn(dAtA, i, uint64(m.WorkersThrottled))
		i--
		dAtA[i] = 0x38
	}
	if m.Workers != 0 {
		i = encodeVarintCpuburn(dAtA, i, uint64(m.Workers))
		i--
		dAtA[i] = 0x30
	}
	if m.ThermalCrit != 0 {
		i = encodeVarintCpuburn(dAtA, i, uint64(m.ThermalCrit))
		i--
		dAtA[i] = 0x28
	}
	if m.ThermalResume != 0 {
		i = encodeVarintCpuburn(dAtA, i, uint64(m.ThermalResume))
		i--
		dAtA[i] = 0x20
	}
	if m.ThermalPoll != 0 {
		i = encodeVarintCpuburn(dAtA, i, uint64(m.ThermalPoll))
		i--
		dAtA[i] = 0x18
	}
	if m.ThermalThrottle {
		i--
		if m.ThermalThrottle {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.TempSensor) > 0 {
		i -= len(m.TempSensor)
		copy(dAtA[i:], m.TempSensor)
		i = encodeVarintCpuburn(dAtA, i, uint64(len(m.TempSensor)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCpuburn(dAtA []byte, offset int, v uint64) int {
	offset -= sovCpuburn(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Config) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TempSensor)
	if l > 0 {
		n += 1 + l + sovCpuburn(uint64(l))
	}
	if m.ThermalThrottle {
		n += 2
	}
	if m.ThermalPoll != 0 {
		n += 1 + sovCpuburn(uint64(m.ThermalPoll))
	}
	if m.ThermalResume != 0 {
		n += 1 + sovCpuburn(uint64(m.ThermalResume))
	}
	if m.ThermalCrit != 0 {
		n += 1 + sovCpuburn(uint64(m.ThermalCrit))
	}
	if m.Workers != 0 {
		n += 1 + sovCpuburn(uint64(m.Workers))
	}
	if m.WorkersThrottled != 0 {
		n += 1 + sovCpuburn(uint64(m.WorkersThrottled))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCpuburn(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCpuburn(x uint64) (n int) {
	return sovCpuburn(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Config) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCpuburn
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
			return fmt.Errorf("proto: Config: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Config: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TempSensor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpuburn
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
				return ErrInvalidLengthCpuburn
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCpuburn
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TempSensor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThermalThrottle", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpuburn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ThermalThrottle = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThermalPoll", wireType)
			}
			m.ThermalPoll = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpuburn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ThermalPoll |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThermalResume", wireType)
			}
			m.ThermalResume = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpuburn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ThermalResume |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThermalCrit", wireType)
			}
			m.ThermalCrit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpuburn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ThermalCrit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Workers", wireType)
			}
			m.Workers = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpuburn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Workers |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkersThrottled", wireType)
			}
			m.WorkersThrottled = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCpuburn
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WorkersThrottled |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCpuburn(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCpuburn
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
func skipCpuburn(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCpuburn
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
					return 0, ErrIntOverflowCpuburn
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
					return 0, ErrIntOverflowCpuburn
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
				return 0, ErrInvalidLengthCpuburn
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCpuburn
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCpuburn
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCpuburn        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCpuburn          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCpuburn = fmt.Errorf("proto: unexpected end of group")
)
