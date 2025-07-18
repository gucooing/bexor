// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Messages in this file are used to test wire encoding order.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/order/order.proto

package order

import (
	protoreflect "github.com/gucooing/bexor/reflect/protoreflect"
	protoimpl "github.com/gucooing/bexor/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type Message struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Field_2 *string                `protobuf:"bytes,2,opt,name=field_2,json=field2" json:"field_2,omitempty"`
	Field_1 *string                `protobuf:"bytes,1,opt,name=field_1,json=field1" json:"field_1,omitempty"`
	// Types that are valid to be assigned to Oneof_1:
	//
	//	*Message_Field_10
	Oneof_1         isMessage_Oneof_1 `protobuf_oneof:"oneof_1"`
	Field_20        *string           `protobuf:"bytes,20,opt,name=field_20,json=field20" json:"field_20,omitempty"`
	extensionFields protoimpl.ExtensionFields
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_internal_testprotos_order_order_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_order_order_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_order_order_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetField_2() string {
	if x != nil && x.Field_2 != nil {
		return *x.Field_2
	}
	return ""
}

func (x *Message) GetField_1() string {
	if x != nil && x.Field_1 != nil {
		return *x.Field_1
	}
	return ""
}

func (x *Message) GetOneof_1() isMessage_Oneof_1 {
	if x != nil {
		return x.Oneof_1
	}
	return nil
}

func (x *Message) GetField_10() string {
	if x != nil {
		if x, ok := x.Oneof_1.(*Message_Field_10); ok {
			return x.Field_10
		}
	}
	return ""
}

func (x *Message) GetField_20() string {
	if x != nil && x.Field_20 != nil {
		return *x.Field_20
	}
	return ""
}

type isMessage_Oneof_1 interface {
	isMessage_Oneof_1()
}

type Message_Field_10 struct {
	Field_10 string `protobuf:"bytes,10,opt,name=field_10,json=field10,oneof"`
}

func (*Message_Field_10) isMessage_Oneof_1() {}

var file_internal_testprotos_order_order_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*Message)(nil),
		ExtensionType: (*string)(nil),
		Field:         30,
		Name:          "goproto.proto.order.field_30",
		Tag:           "bytes,30,opt,name=field_30",
		Filename:      "internal/testprotos/order/order.proto",
	},
	{
		ExtendedType:  (*Message)(nil),
		ExtensionType: (*string)(nil),
		Field:         31,
		Name:          "goproto.proto.order.field_31",
		Tag:           "bytes,31,opt,name=field_31",
		Filename:      "internal/testprotos/order/order.proto",
	},
	{
		ExtendedType:  (*Message)(nil),
		ExtensionType: (*string)(nil),
		Field:         32,
		Name:          "goproto.proto.order.field_32",
		Tag:           "bytes,32,opt,name=field_32",
		Filename:      "internal/testprotos/order/order.proto",
	},
}

// Extension fields to Message.
var (
	// optional string field_30 = 30;
	E_Field_30 = &file_internal_testprotos_order_order_proto_extTypes[0]
	// optional string field_31 = 31;
	E_Field_31 = &file_internal_testprotos_order_order_proto_extTypes[1]
	// optional string field_32 = 32;
	E_Field_32 = &file_internal_testprotos_order_order_proto_extTypes[2]
)

var File_internal_testprotos_order_order_proto protoreflect.FileDescriptor

const file_internal_testprotos_order_order_proto_rawDesc = "" +
	"\n" +
	"%internal/testprotos/order/order.proto\x12\x13goproto.proto.order\"\x84\x01\n" +
	"\aMessage\x12\x17\n" +
	"\afield_2\x18\x02 \x01(\tR\x06field2\x12\x17\n" +
	"\afield_1\x18\x01 \x01(\tR\x06field1\x12\x1b\n" +
	"\bfield_10\x18\n" +
	" \x01(\tH\x00R\afield10\x12\x19\n" +
	"\bfield_20\x18\x14 \x01(\tR\afield20*\x04\b\x1e\x10)B\t\n" +
	"\aoneof_1:7\n" +
	"\bfield_30\x12\x1c.goproto.proto.order.Message\x18\x1e \x01(\tR\afield30:7\n" +
	"\bfield_31\x12\x1c.goproto.proto.order.Message\x18\x1f \x01(\tR\afield31:7\n" +
	"\bfield_32\x12\x1c.goproto.proto.order.Message\x18  \x01(\tR\afield32B6Z4github.com/gucooing/bexor/internal/testprotos/order"

var (
	file_internal_testprotos_order_order_proto_rawDescOnce sync.Once
	file_internal_testprotos_order_order_proto_rawDescData []byte
)

func file_internal_testprotos_order_order_proto_rawDescGZIP() []byte {
	file_internal_testprotos_order_order_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_order_order_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_testprotos_order_order_proto_rawDesc), len(file_internal_testprotos_order_order_proto_rawDesc)))
	})
	return file_internal_testprotos_order_order_proto_rawDescData
}

var file_internal_testprotos_order_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_order_order_proto_goTypes = []any{
	(*Message)(nil), // 0: goproto.proto.order.Message
}
var file_internal_testprotos_order_order_proto_depIdxs = []int32{
	0, // 0: goproto.proto.order.field_30:extendee -> goproto.proto.order.Message
	0, // 1: goproto.proto.order.field_31:extendee -> goproto.proto.order.Message
	0, // 2: goproto.proto.order.field_32:extendee -> goproto.proto.order.Message
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_order_order_proto_init() }
func file_internal_testprotos_order_order_proto_init() {
	if File_internal_testprotos_order_order_proto != nil {
		return
	}
	file_internal_testprotos_order_order_proto_msgTypes[0].OneofWrappers = []any{
		(*Message_Field_10)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_order_order_proto_rawDesc), len(file_internal_testprotos_order_order_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_order_order_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_order_order_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_order_order_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_order_order_proto_extTypes,
	}.Build()
	File_internal_testprotos_order_order_proto = out.File
	file_internal_testprotos_order_order_proto_goTypes = nil
	file_internal_testprotos_order_order_proto_depIdxs = nil
}
