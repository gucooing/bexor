// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/proto2/enum.proto

package proto2

import (
	protoreflect "github.com/gucooing/bexor/reflect/protoreflect"
	protoimpl "github.com/gucooing/bexor/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

// EnumType1 comment.
type EnumType1 int32

const (
	// EnumType1_ONE comment.
	EnumType1_ONE EnumType1 = 1
	// EnumType1_TWO comment.
	EnumType1_TWO EnumType1 = 2
)

// Enum value maps for EnumType1.
var (
	EnumType1_name = map[int32]string{
		1: "ONE",
		2: "TWO",
	}
	EnumType1_value = map[string]int32{
		"ONE": 1,
		"TWO": 2,
	}
)

func (x EnumType1) Enum() *EnumType1 {
	p := new(EnumType1)
	*p = x
	return p
}

func (x EnumType1) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumType1) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[0].Descriptor()
}

func (EnumType1) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[0]
}

func (x EnumType1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumType1) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumType1(num)
	return nil
}

// Deprecated: Use EnumType1.Descriptor instead.
func (EnumType1) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP(), []int{0}
}

type EnumType2 int32

const (
	EnumType2_duplicate1 EnumType2 = 1
	EnumType2_duplicate2 EnumType2 = 1
)

// Enum value maps for EnumType2.
var (
	EnumType2_name = map[int32]string{
		1: "duplicate1",
		// Duplicate value: 1: "duplicate2",
	}
	EnumType2_value = map[string]int32{
		"duplicate1": 1,
		"duplicate2": 1,
	}
)

func (x EnumType2) Enum() *EnumType2 {
	p := new(EnumType2)
	*p = x
	return p
}

func (x EnumType2) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumType2) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[1].Descriptor()
}

func (EnumType2) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[1]
}

func (x EnumType2) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumType2) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumType2(num)
	return nil
}

// Deprecated: Use EnumType2.Descriptor instead.
func (EnumType2) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP(), []int{1}
}

// NestedEnumType1A comment.
type EnumContainerMessage1_NestedEnumType1A int32

const (
	// NestedEnumType1A_VALUE comment.
	EnumContainerMessage1_NESTED_1A_VALUE EnumContainerMessage1_NestedEnumType1A = 0
)

// Enum value maps for EnumContainerMessage1_NestedEnumType1A.
var (
	EnumContainerMessage1_NestedEnumType1A_name = map[int32]string{
		0: "NESTED_1A_VALUE",
	}
	EnumContainerMessage1_NestedEnumType1A_value = map[string]int32{
		"NESTED_1A_VALUE": 0,
	}
)

func (x EnumContainerMessage1_NestedEnumType1A) Enum() *EnumContainerMessage1_NestedEnumType1A {
	p := new(EnumContainerMessage1_NestedEnumType1A)
	*p = x
	return p
}

func (x EnumContainerMessage1_NestedEnumType1A) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumContainerMessage1_NestedEnumType1A) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[2].Descriptor()
}

func (EnumContainerMessage1_NestedEnumType1A) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[2]
}

func (x EnumContainerMessage1_NestedEnumType1A) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumContainerMessage1_NestedEnumType1A) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_NestedEnumType1A(num)
	return nil
}

// Deprecated: Use EnumContainerMessage1_NestedEnumType1A.Descriptor instead.
func (EnumContainerMessage1_NestedEnumType1A) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP(), []int{0, 0}
}

type EnumContainerMessage1_NestedEnumType1B int32

const (
	EnumContainerMessage1_NESTED_1B_VALUE EnumContainerMessage1_NestedEnumType1B = 0
)

// Enum value maps for EnumContainerMessage1_NestedEnumType1B.
var (
	EnumContainerMessage1_NestedEnumType1B_name = map[int32]string{
		0: "NESTED_1B_VALUE",
	}
	EnumContainerMessage1_NestedEnumType1B_value = map[string]int32{
		"NESTED_1B_VALUE": 0,
	}
)

func (x EnumContainerMessage1_NestedEnumType1B) Enum() *EnumContainerMessage1_NestedEnumType1B {
	p := new(EnumContainerMessage1_NestedEnumType1B)
	*p = x
	return p
}

func (x EnumContainerMessage1_NestedEnumType1B) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumContainerMessage1_NestedEnumType1B) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[3].Descriptor()
}

func (EnumContainerMessage1_NestedEnumType1B) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[3]
}

func (x EnumContainerMessage1_NestedEnumType1B) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumContainerMessage1_NestedEnumType1B) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_NestedEnumType1B(num)
	return nil
}

// Deprecated: Use EnumContainerMessage1_NestedEnumType1B.Descriptor instead.
func (EnumContainerMessage1_NestedEnumType1B) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP(), []int{0, 1}
}

// NestedEnumType2A comment.
type EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A int32

const (
	// NestedEnumType2A_VALUE comment.
	EnumContainerMessage1_EnumContainerMessage2_NESTED_2A_VALUE EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A = 0
)

// Enum value maps for EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A.
var (
	EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_name = map[int32]string{
		0: "NESTED_2A_VALUE",
	}
	EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_value = map[string]int32{
		"NESTED_2A_VALUE": 0,
	}
)

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Enum() *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A {
	p := new(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A)
	*p = x
	return p
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[4].Descriptor()
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[4]
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A(num)
	return nil
}

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A.Descriptor instead.
func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP(), []int{0, 0, 0}
}

type EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B int32

const (
	EnumContainerMessage1_EnumContainerMessage2_NESTED_2B_VALUE EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B = 0
)

// Enum value maps for EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B.
var (
	EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_name = map[int32]string{
		0: "NESTED_2B_VALUE",
	}
	EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_value = map[string]int32{
		"NESTED_2B_VALUE": 0,
	}
)

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Enum() *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B {
	p := new(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B)
	*p = x
	return p
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[5].Descriptor()
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes[5]
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B(num)
	return nil
}

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B.Descriptor instead.
func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP(), []int{0, 0, 1}
}

type EnumContainerMessage1 struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	DefaultDuplicate1 *EnumType2             `protobuf:"varint,1,opt,name=default_duplicate1,json=defaultDuplicate1,enum=goproto.protoc.proto2.EnumType2,def=1" json:"default_duplicate1,omitempty"`
	DefaultDuplicate2 *EnumType2             `protobuf:"varint,2,opt,name=default_duplicate2,json=defaultDuplicate2,enum=goproto.protoc.proto2.EnumType2,def=1" json:"default_duplicate2,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

// Default values for EnumContainerMessage1 fields.
const (
	Default_EnumContainerMessage1_DefaultDuplicate1 = EnumType2_duplicate1
	Default_EnumContainerMessage1_DefaultDuplicate2 = EnumType2_duplicate2
)

func (x *EnumContainerMessage1) Reset() {
	*x = EnumContainerMessage1{}
	mi := &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnumContainerMessage1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumContainerMessage1) ProtoMessage() {}

func (x *EnumContainerMessage1) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumContainerMessage1.ProtoReflect.Descriptor instead.
func (*EnumContainerMessage1) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP(), []int{0}
}

func (x *EnumContainerMessage1) GetDefaultDuplicate1() EnumType2 {
	if x != nil && x.DefaultDuplicate1 != nil {
		return *x.DefaultDuplicate1
	}
	return Default_EnumContainerMessage1_DefaultDuplicate1
}

func (x *EnumContainerMessage1) GetDefaultDuplicate2() EnumType2 {
	if x != nil && x.DefaultDuplicate2 != nil {
		return *x.DefaultDuplicate2
	}
	return Default_EnumContainerMessage1_DefaultDuplicate2
}

type EnumContainerMessage1_EnumContainerMessage2 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnumContainerMessage1_EnumContainerMessage2) Reset() {
	*x = EnumContainerMessage1_EnumContainerMessage2{}
	mi := &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnumContainerMessage1_EnumContainerMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumContainerMessage1_EnumContainerMessage2) ProtoMessage() {}

func (x *EnumContainerMessage1_EnumContainerMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_proto2_enum_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumContainerMessage1_EnumContainerMessage2.ProtoReflect.Descriptor instead.
func (*EnumContainerMessage1_EnumContainerMessage2) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP(), []int{0, 0}
}

var File_cmd_protoc_gen_go_testdata_proto2_enum_proto protoreflect.FileDescriptor

const file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDesc = "" +
	"\n" +
	",cmd/protoc-gen-go/testdata/proto2/enum.proto\x12\x15goproto.protoc.proto2\"\x8e\x03\n" +
	"\x15EnumContainerMessage1\x12[\n" +
	"\x12default_duplicate1\x18\x01 \x01(\x0e2 .goproto.protoc.proto2.EnumType2:\n" +
	"duplicate1R\x11defaultDuplicate1\x12[\n" +
	"\x12default_duplicate2\x18\x02 \x01(\x0e2 .goproto.protoc.proto2.EnumType2:\n" +
	"duplicate2R\x11defaultDuplicate2\x1ai\n" +
	"\x15EnumContainerMessage2\"'\n" +
	"\x10NestedEnumType2A\x12\x13\n" +
	"\x0fNESTED_2A_VALUE\x10\x00\"'\n" +
	"\x10NestedEnumType2B\x12\x13\n" +
	"\x0fNESTED_2B_VALUE\x10\x00\"'\n" +
	"\x10NestedEnumType1A\x12\x13\n" +
	"\x0fNESTED_1A_VALUE\x10\x00\"'\n" +
	"\x10NestedEnumType1B\x12\x13\n" +
	"\x0fNESTED_1B_VALUE\x10\x00*\x1d\n" +
	"\tEnumType1\x12\a\n" +
	"\x03ONE\x10\x01\x12\a\n" +
	"\x03TWO\x10\x02*Q\n" +
	"\tEnumType2\x12\x0e\n" +
	"\n" +
	"duplicate1\x10\x01\x12\x0e\n" +
	"\n" +
	"duplicate2\x10\x01\x1a\x02\x10\x01\"\x04\b\x02\x10\x02\"\x04\b\x03\x10\x03*\tRESERVED1*\tRESERVED2B>Z<github.com/gucooing/bexor/cmd/protoc-gen-go/testdata/proto2"

var (
	file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescData []byte
)

func file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDesc)))
	})
	return file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_cmd_protoc_gen_go_testdata_proto2_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_protoc_gen_go_testdata_proto2_enum_proto_goTypes = []any{
	(EnumType1)(0), // 0: goproto.protoc.proto2.EnumType1
	(EnumType2)(0), // 1: goproto.protoc.proto2.EnumType2
	(EnumContainerMessage1_NestedEnumType1A)(0),                       // 2: goproto.protoc.proto2.EnumContainerMessage1.NestedEnumType1A
	(EnumContainerMessage1_NestedEnumType1B)(0),                       // 3: goproto.protoc.proto2.EnumContainerMessage1.NestedEnumType1B
	(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A)(0), // 4: goproto.protoc.proto2.EnumContainerMessage1.EnumContainerMessage2.NestedEnumType2A
	(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B)(0), // 5: goproto.protoc.proto2.EnumContainerMessage1.EnumContainerMessage2.NestedEnumType2B
	(*EnumContainerMessage1)(nil),                                     // 6: goproto.protoc.proto2.EnumContainerMessage1
	(*EnumContainerMessage1_EnumContainerMessage2)(nil),               // 7: goproto.protoc.proto2.EnumContainerMessage1.EnumContainerMessage2
}
var file_cmd_protoc_gen_go_testdata_proto2_enum_proto_depIdxs = []int32{
	1, // 0: goproto.protoc.proto2.EnumContainerMessage1.default_duplicate1:type_name -> goproto.protoc.proto2.EnumType2
	1, // 1: goproto.protoc.proto2.EnumContainerMessage1.default_duplicate2:type_name -> goproto.protoc.proto2.EnumType2
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_proto2_enum_proto_init() }
func file_cmd_protoc_gen_go_testdata_proto2_enum_proto_init() {
	if File_cmd_protoc_gen_go_testdata_proto2_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDesc), len(file_cmd_protoc_gen_go_testdata_proto2_enum_proto_rawDesc)),
			NumEnums:      6,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_proto2_enum_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_proto2_enum_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_proto2_enum_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_proto2_enum_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_proto2_enum_proto = out.File
	file_cmd_protoc_gen_go_testdata_proto2_enum_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_proto2_enum_proto_depIdxs = nil
}
