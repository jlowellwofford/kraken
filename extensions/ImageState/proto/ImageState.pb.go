// ImageState.proto: generic data for tracking node images states
//
// Author: J. Lowell Wofford <lowell@lanl.gov>
//
// This software is open source software available under the BSD-3 license.
// Copyright (c) 2020, Triad National Security, LLC
// See LICENSE file for details.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.5.1
// source: ImageState.proto

package proto

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

type ImageState_State int32

const (
	ImageState_NONE     ImageState_State = 0 // current image state is unknown
	ImageState_INACTIVE ImageState_State = 1 // the image is not currently loaded, no action has been taken
	ImageState_LOADING  ImageState_State = 2 // in the process of attaching/loading image
	ImageState_ACTIVE   ImageState_State = 3 // on the correct image and active
	ImageState_UPDATE   ImageState_State = 4 // image needs to be changed
	ImageState_ERROR    ImageState_State = 5 // unrecoverable error
)

// Enum value maps for ImageState_State.
var (
	ImageState_State_name = map[int32]string{
		0: "NONE",
		1: "INACTIVE",
		2: "LOADING",
		3: "ACTIVE",
		4: "UPDATE",
		5: "ERROR",
	}
	ImageState_State_value = map[string]int32{
		"NONE":     0,
		"INACTIVE": 1,
		"LOADING":  2,
		"ACTIVE":   3,
		"UPDATE":   4,
		"ERROR":    5,
	}
)

func (x ImageState_State) Enum() *ImageState_State {
	p := new(ImageState_State)
	*p = x
	return p
}

func (x ImageState_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageState_State) Descriptor() protoreflect.EnumDescriptor {
	return file_ImageState_proto_enumTypes[0].Descriptor()
}

func (ImageState_State) Type() protoreflect.EnumType {
	return &file_ImageState_proto_enumTypes[0]
}

func (x ImageState_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageState_State.Descriptor instead.
func (ImageState_State) EnumDescriptor() ([]byte, []int) {
	return file_ImageState_proto_rawDescGZIP(), []int{0, 0}
}

type ImageState_Busy int32

const (
	ImageState_UNKNOWN ImageState_Busy = 0
	ImageState_IDLE    ImageState_Busy = 1
	ImageState_BUSY    ImageState_Busy = 2
)

// Enum value maps for ImageState_Busy.
var (
	ImageState_Busy_name = map[int32]string{
		0: "UNKNOWN",
		1: "IDLE",
		2: "BUSY",
	}
	ImageState_Busy_value = map[string]int32{
		"UNKNOWN": 0,
		"IDLE":    1,
		"BUSY":    2,
	}
)

func (x ImageState_Busy) Enum() *ImageState_Busy {
	p := new(ImageState_Busy)
	*p = x
	return p
}

func (x ImageState_Busy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageState_Busy) Descriptor() protoreflect.EnumDescriptor {
	return file_ImageState_proto_enumTypes[1].Descriptor()
}

func (ImageState_Busy) Type() protoreflect.EnumType {
	return &file_ImageState_proto_enumTypes[1]
}

func (x ImageState_Busy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageState_Busy.Descriptor instead.
func (ImageState_Busy) EnumDescriptor() ([]byte, []int) {
	return file_ImageState_proto_rawDescGZIP(), []int{0, 1}
}

type ImageState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State     ImageState_State `protobuf:"varint,1,opt,name=state,proto3,enum=proto.ImageState_State" json:"state,omitempty"` // current state of the image
	Busy      ImageState_Busy  `protobuf:"varint,2,opt,name=busy,proto3,enum=proto.ImageState_Busy" json:"busy,omitempty"`    // is the image busy?
	ImageName string           `protobuf:"bytes,3,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`     // unique identifier for image
	Error     string           `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`                              // last image error message
}

func (x *ImageState) Reset() {
	*x = ImageState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ImageState_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageState) ProtoMessage() {}

func (x *ImageState) ProtoReflect() protoreflect.Message {
	mi := &file_ImageState_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageState.ProtoReflect.Descriptor instead.
func (*ImageState) Descriptor() ([]byte, []int) {
	return file_ImageState_proto_rawDescGZIP(), []int{0}
}

func (x *ImageState) GetState() ImageState_State {
	if x != nil {
		return x.State
	}
	return ImageState_NONE
}

func (x *ImageState) GetBusy() ImageState_Busy {
	if x != nil {
		return x.Busy
	}
	return ImageState_UNKNOWN
}

func (x *ImageState) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *ImageState) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_ImageState_proto protoreflect.FileDescriptor

var file_ImageState_proto_rawDesc = []byte{
	0x0a, 0x10, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x02, 0x0a, 0x0a, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x75, 0x73, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x42, 0x75, 0x73, 0x79, 0x52, 0x04, 0x62,
	0x75, 0x73, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x4f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49,
	0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x41,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x09,
	0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x22, 0x27, 0x0a, 0x04, 0x42, 0x75, 0x73,
	0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x42, 0x55, 0x53, 0x59,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ImageState_proto_rawDescOnce sync.Once
	file_ImageState_proto_rawDescData = file_ImageState_proto_rawDesc
)

func file_ImageState_proto_rawDescGZIP() []byte {
	file_ImageState_proto_rawDescOnce.Do(func() {
		file_ImageState_proto_rawDescData = protoimpl.X.CompressGZIP(file_ImageState_proto_rawDescData)
	})
	return file_ImageState_proto_rawDescData
}

var file_ImageState_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ImageState_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ImageState_proto_goTypes = []interface{}{
	(ImageState_State)(0), // 0: proto.ImageState.State
	(ImageState_Busy)(0),  // 1: proto.ImageState.Busy
	(*ImageState)(nil),    // 2: proto.ImageState
}
var file_ImageState_proto_depIdxs = []int32{
	0, // 0: proto.ImageState.state:type_name -> proto.ImageState.State
	1, // 1: proto.ImageState.busy:type_name -> proto.ImageState.Busy
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ImageState_proto_init() }
func file_ImageState_proto_init() {
	if File_ImageState_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ImageState_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageState); i {
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
			RawDescriptor: file_ImageState_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ImageState_proto_goTypes,
		DependencyIndexes: file_ImageState_proto_depIdxs,
		EnumInfos:         file_ImageState_proto_enumTypes,
		MessageInfos:      file_ImageState_proto_msgTypes,
	}.Build()
	File_ImageState_proto = out.File
	file_ImageState_proto_rawDesc = nil
	file_ImageState_proto_goTypes = nil
	file_ImageState_proto_depIdxs = nil
}
