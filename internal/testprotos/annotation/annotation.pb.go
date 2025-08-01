// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/annotation/annotation.proto

package annotation

import (
	protoreflect "github.com/gucooing/bexor/reflect/protoreflect"
	protoimpl "github.com/gucooing/bexor/runtime/protoimpl"
	descriptorpb "github.com/gucooing/bexor/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

var file_internal_testprotos_annotation_annotation_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         37383685,
		Name:          "go_annotation.track_field_use",
		Tag:           "varint,37383685,opt,name=track_field_use",
		Filename:      "internal/testprotos/annotation/annotation.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Setting this on a message enables tracking of which fields in the message
	// a specific binary might access. As a consequence, it also disables the use
	// of the message accessor methods to satisfy interfaces: they can only be
	// called directly.
	//
	// optional bool track_field_use = 37383685;
	E_TrackFieldUse = &file_internal_testprotos_annotation_annotation_proto_extTypes[0]
)

var File_internal_testprotos_annotation_annotation_proto protoreflect.FileDescriptor

const file_internal_testprotos_annotation_annotation_proto_rawDesc = "" +
	"\n" +
	"/internal/testprotos/annotation/annotation.proto\x12\rgo_annotation\x1a google/protobuf/descriptor.proto:J\n" +
	"\x0ftrack_field_use\x12\x1f.google.protobuf.MessageOptions\x18\x85\xdc\xe9\x11 \x01(\bR\rtrackFieldUseB;Z9github.com/gucooing/bexor/internal/testprotos/annotation"

var file_internal_testprotos_annotation_annotation_proto_goTypes = []any{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
}
var file_internal_testprotos_annotation_annotation_proto_depIdxs = []int32{
	0, // 0: go_annotation.track_field_use:extendee -> google.protobuf.MessageOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_annotation_annotation_proto_init() }
func file_internal_testprotos_annotation_annotation_proto_init() {
	if File_internal_testprotos_annotation_annotation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_annotation_annotation_proto_rawDesc), len(file_internal_testprotos_annotation_annotation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_annotation_annotation_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_annotation_annotation_proto_depIdxs,
		ExtensionInfos:    file_internal_testprotos_annotation_annotation_proto_extTypes,
	}.Build()
	File_internal_testprotos_annotation_annotation_proto = out.File
	file_internal_testprotos_annotation_annotation_proto_goTypes = nil
	file_internal_testprotos_annotation_annotation_proto_depIdxs = nil
}
