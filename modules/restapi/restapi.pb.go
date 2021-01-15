// restapi.proto: describes the RestAPIConfig object
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
// source: restapi.proto

package restapi

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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Port int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_restapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_restapi_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Config) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_restapi_proto protoreflect.FileDescriptor

var file_restapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x22, 0x30, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_restapi_proto_rawDescOnce sync.Once
	file_restapi_proto_rawDescData = file_restapi_proto_rawDesc
)

func file_restapi_proto_rawDescGZIP() []byte {
	file_restapi_proto_rawDescOnce.Do(func() {
		file_restapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_restapi_proto_rawDescData)
	})
	return file_restapi_proto_rawDescData
}

var file_restapi_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_restapi_proto_goTypes = []interface{}{
	(*Config)(nil), // 0: restapi.Config
}
var file_restapi_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_restapi_proto_init() }
func file_restapi_proto_init() {
	if File_restapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_restapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_restapi_proto_goTypes,
		DependencyIndexes: file_restapi_proto_depIdxs,
		MessageInfos:      file_restapi_proto_msgTypes,
	}.Build()
	File_restapi_proto = out.File
	file_restapi_proto_rawDesc = nil
	file_restapi_proto_goTypes = nil
	file_restapi_proto_depIdxs = nil
}