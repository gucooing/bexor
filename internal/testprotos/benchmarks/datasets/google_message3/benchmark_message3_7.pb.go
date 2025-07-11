// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datasets/google_message3/benchmark_message3_7.proto

package google_message3

import (
	protoreflect "github.com/gucooing/bexor/reflect/protoreflect"
	protoimpl "github.com/gucooing/bexor/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Message11018 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message11018) Reset() {
	*x = Message11018{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message11018) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message11018) ProtoMessage() {}

func (x *Message11018) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message11018.ProtoReflect.Descriptor instead.
func (*Message11018) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{0}
}

type Message10800 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field10808 *string  `protobuf:"bytes,1,opt,name=field10808" json:"field10808,omitempty"`
	Field10809 *int64   `protobuf:"varint,2,opt,name=field10809" json:"field10809,omitempty"`
	Field10810 *bool    `protobuf:"varint,3,opt,name=field10810" json:"field10810,omitempty"`
	Field10811 *float32 `protobuf:"fixed32,4,opt,name=field10811" json:"field10811,omitempty"`
}

func (x *Message10800) Reset() {
	*x = Message10800{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message10800) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message10800) ProtoMessage() {}

func (x *Message10800) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message10800.ProtoReflect.Descriptor instead.
func (*Message10800) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{1}
}

func (x *Message10800) GetField10808() string {
	if x != nil && x.Field10808 != nil {
		return *x.Field10808
	}
	return ""
}

func (x *Message10800) GetField10809() int64 {
	if x != nil && x.Field10809 != nil {
		return *x.Field10809
	}
	return 0
}

func (x *Message10800) GetField10810() bool {
	if x != nil && x.Field10810 != nil {
		return *x.Field10810
	}
	return false
}

func (x *Message10800) GetField10811() float32 {
	if x != nil && x.Field10811 != nil {
		return *x.Field10811
	}
	return 0
}

type Message10802 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message10802) Reset() {
	*x = Message10802{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message10802) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message10802) ProtoMessage() {}

func (x *Message10802) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message10802.ProtoReflect.Descriptor instead.
func (*Message10802) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{2}
}

type Message10748 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field10750 *string `protobuf:"bytes,1,opt,name=field10750" json:"field10750,omitempty"`
	Field10751 *int32  `protobuf:"varint,2,opt,name=field10751" json:"field10751,omitempty"`
	Field10752 *int32  `protobuf:"varint,3,opt,name=field10752" json:"field10752,omitempty"`
	Field10753 *int32  `protobuf:"varint,4,opt,name=field10753" json:"field10753,omitempty"`
}

func (x *Message10748) Reset() {
	*x = Message10748{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message10748) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message10748) ProtoMessage() {}

func (x *Message10748) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message10748.ProtoReflect.Descriptor instead.
func (*Message10748) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{3}
}

func (x *Message10748) GetField10750() string {
	if x != nil && x.Field10750 != nil {
		return *x.Field10750
	}
	return ""
}

func (x *Message10748) GetField10751() int32 {
	if x != nil && x.Field10751 != nil {
		return *x.Field10751
	}
	return 0
}

func (x *Message10748) GetField10752() int32 {
	if x != nil && x.Field10752 != nil {
		return *x.Field10752
	}
	return 0
}

func (x *Message10748) GetField10753() int32 {
	if x != nil && x.Field10753 != nil {
		return *x.Field10753
	}
	return 0
}

type Message7966 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field7969 *string `protobuf:"bytes,1,opt,name=field7969" json:"field7969,omitempty"`
	Field7970 *bool   `protobuf:"varint,2,opt,name=field7970" json:"field7970,omitempty"`
}

func (x *Message7966) Reset() {
	*x = Message7966{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message7966) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message7966) ProtoMessage() {}

func (x *Message7966) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message7966.ProtoReflect.Descriptor instead.
func (*Message7966) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{4}
}

func (x *Message7966) GetField7969() string {
	if x != nil && x.Field7969 != nil {
		return *x.Field7969
	}
	return ""
}

func (x *Message7966) GetField7970() bool {
	if x != nil && x.Field7970 != nil {
		return *x.Field7970
	}
	return false
}

type Message708 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field823 *Message741 `protobuf:"bytes,1,opt,name=field823" json:"field823,omitempty"`
	Field824 []string    `protobuf:"bytes,6,rep,name=field824" json:"field824,omitempty"`
	Field825 *string     `protobuf:"bytes,2,opt,name=field825" json:"field825,omitempty"`
	Field826 *string     `protobuf:"bytes,3,opt,name=field826" json:"field826,omitempty"`
	Field827 []string    `protobuf:"bytes,4,rep,name=field827" json:"field827,omitempty"`
	Field828 []string    `protobuf:"bytes,5,rep,name=field828" json:"field828,omitempty"`
}

func (x *Message708) Reset() {
	*x = Message708{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message708) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message708) ProtoMessage() {}

func (x *Message708) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message708.ProtoReflect.Descriptor instead.
func (*Message708) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{5}
}

func (x *Message708) GetField823() *Message741 {
	if x != nil {
		return x.Field823
	}
	return nil
}

func (x *Message708) GetField824() []string {
	if x != nil {
		return x.Field824
	}
	return nil
}

func (x *Message708) GetField825() string {
	if x != nil && x.Field825 != nil {
		return *x.Field825
	}
	return ""
}

func (x *Message708) GetField826() string {
	if x != nil && x.Field826 != nil {
		return *x.Field826
	}
	return ""
}

func (x *Message708) GetField827() []string {
	if x != nil {
		return x.Field827
	}
	return nil
}

func (x *Message708) GetField828() []string {
	if x != nil {
		return x.Field828
	}
	return nil
}

type Message8942 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Message8942) Reset() {
	*x = Message8942{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message8942) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message8942) ProtoMessage() {}

func (x *Message8942) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message8942.ProtoReflect.Descriptor instead.
func (*Message8942) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{6}
}

type Message11011 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field11752 []byte `protobuf:"bytes,1,req,name=field11752" json:"field11752,omitempty"`
	Field11753 []byte `protobuf:"bytes,2,req,name=field11753" json:"field11753,omitempty"`
}

func (x *Message11011) Reset() {
	*x = Message11011{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message11011) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message11011) ProtoMessage() {}

func (x *Message11011) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message11011.ProtoReflect.Descriptor instead.
func (*Message11011) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{7}
}

func (x *Message11011) GetField11752() []byte {
	if x != nil {
		return x.Field11752
	}
	return nil
}

func (x *Message11011) GetField11753() []byte {
	if x != nil {
		return x.Field11753
	}
	return nil
}

type UnusedEmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnusedEmptyMessage) Reset() {
	*x = UnusedEmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnusedEmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnusedEmptyMessage) ProtoMessage() {}

func (x *UnusedEmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnusedEmptyMessage.ProtoReflect.Descriptor instead.
func (*UnusedEmptyMessage) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{8}
}

type Message741 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field936 []string `protobuf:"bytes,1,rep,name=field936" json:"field936,omitempty"`
}

func (x *Message741) Reset() {
	*x = Message741{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message741) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message741) ProtoMessage() {}

func (x *Message741) ProtoReflect() protoreflect.Message {
	mi := &file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message741.ProtoReflect.Descriptor instead.
func (*Message741) Descriptor() ([]byte, []int) {
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP(), []int{9}
}

func (x *Message741) GetField936() []string {
	if x != nil {
		return x.Field936
	}
	return nil
}

var File_datasets_google_message3_benchmark_message3_7_proto protoreflect.FileDescriptor

var file_datasets_google_message3_benchmark_message3_7_proto_rawDesc = []byte{
	0x0a, 0x33, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x2f, 0x62, 0x65, 0x6e, 0x63, 0x68,
	0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x5f, 0x37, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b,
	0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x33, 0x22, 0x0e, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x31, 0x30, 0x31,
	0x38, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x30, 0x38,
	0x30, 0x30, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x38, 0x30, 0x38,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x38,
	0x30, 0x38, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x38, 0x30, 0x39,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x38,
	0x30, 0x39, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x38, 0x31, 0x30,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x38,
	0x31, 0x30, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x38, 0x31, 0x31,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x38,
	0x31, 0x31, 0x22, 0x0e, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x30, 0x38,
	0x30, 0x32, 0x22, 0x8e, 0x01, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x30,
	0x37, 0x34, 0x38, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x37, 0x35,
	0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x37, 0x35, 0x30, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x37, 0x35,
	0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x37, 0x35, 0x31, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x37, 0x35,
	0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x37, 0x35, 0x32, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30, 0x37, 0x35,
	0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x30,
	0x37, 0x35, 0x33, 0x22, 0x49, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x37, 0x39,
	0x36, 0x36, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x39, 0x36, 0x39, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x39, 0x36, 0x39,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x39, 0x37, 0x30, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x37, 0x39, 0x37, 0x30, 0x22, 0xdc,
	0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x37, 0x30, 0x38, 0x12, 0x42, 0x0a,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32, 0x33, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x33, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x37, 0x34, 0x31, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32,
	0x33, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32, 0x34, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32, 0x34, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32, 0x35, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x38, 0x32, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x38, 0x32, 0x36, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32,
	0x37, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32,
	0x37, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32, 0x38, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x38, 0x32, 0x38, 0x22, 0x0d, 0x0a,
	0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x38, 0x39, 0x34, 0x32, 0x22, 0x4e, 0x0a, 0x0c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x31, 0x31, 0x30, 0x31, 0x31, 0x12, 0x1e, 0x0a, 0x0a,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x31, 0x37, 0x35, 0x32, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c,
	0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x31, 0x37, 0x35, 0x32, 0x12, 0x1e, 0x0a, 0x0a,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x31, 0x37, 0x35, 0x33, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0c,
	0x52, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x31, 0x37, 0x35, 0x33, 0x22, 0x14, 0x0a, 0x12,
	0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x28, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x37, 0x34, 0x31,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x33, 0x36, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x39, 0x33, 0x36, 0x42, 0x23, 0x0a, 0x1e,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x62, 0x65, 0x6e, 0x63, 0x68, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0xf8, 0x01,
	0x01,
}

var (
	file_datasets_google_message3_benchmark_message3_7_proto_rawDescOnce sync.Once
	file_datasets_google_message3_benchmark_message3_7_proto_rawDescData = file_datasets_google_message3_benchmark_message3_7_proto_rawDesc
)

func file_datasets_google_message3_benchmark_message3_7_proto_rawDescGZIP() []byte {
	file_datasets_google_message3_benchmark_message3_7_proto_rawDescOnce.Do(func() {
		file_datasets_google_message3_benchmark_message3_7_proto_rawDescData = protoimpl.X.CompressGZIP(file_datasets_google_message3_benchmark_message3_7_proto_rawDescData)
	})
	return file_datasets_google_message3_benchmark_message3_7_proto_rawDescData
}

var file_datasets_google_message3_benchmark_message3_7_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_datasets_google_message3_benchmark_message3_7_proto_goTypes = []interface{}{
	(*Message11018)(nil),       // 0: benchmarks.google_message3.Message11018
	(*Message10800)(nil),       // 1: benchmarks.google_message3.Message10800
	(*Message10802)(nil),       // 2: benchmarks.google_message3.Message10802
	(*Message10748)(nil),       // 3: benchmarks.google_message3.Message10748
	(*Message7966)(nil),        // 4: benchmarks.google_message3.Message7966
	(*Message708)(nil),         // 5: benchmarks.google_message3.Message708
	(*Message8942)(nil),        // 6: benchmarks.google_message3.Message8942
	(*Message11011)(nil),       // 7: benchmarks.google_message3.Message11011
	(*UnusedEmptyMessage)(nil), // 8: benchmarks.google_message3.UnusedEmptyMessage
	(*Message741)(nil),         // 9: benchmarks.google_message3.Message741
}
var file_datasets_google_message3_benchmark_message3_7_proto_depIdxs = []int32{
	9, // 0: benchmarks.google_message3.Message708.field823:type_name -> benchmarks.google_message3.Message741
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_datasets_google_message3_benchmark_message3_7_proto_init() }
func file_datasets_google_message3_benchmark_message3_7_proto_init() {
	if File_datasets_google_message3_benchmark_message3_7_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message11018); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message10800); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message10802); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message10748); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message7966); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message708); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message8942); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message11011); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnusedEmptyMessage); i {
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
		file_datasets_google_message3_benchmark_message3_7_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message741); i {
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
			RawDescriptor: file_datasets_google_message3_benchmark_message3_7_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datasets_google_message3_benchmark_message3_7_proto_goTypes,
		DependencyIndexes: file_datasets_google_message3_benchmark_message3_7_proto_depIdxs,
		MessageInfos:      file_datasets_google_message3_benchmark_message3_7_proto_msgTypes,
	}.Build()
	File_datasets_google_message3_benchmark_message3_7_proto = out.File
	file_datasets_google_message3_benchmark_message3_7_proto_rawDesc = nil
	file_datasets_google_message3_benchmark_message3_7_proto_goTypes = nil
	file_datasets_google_message3_benchmark_message3_7_proto_depIdxs = nil
}
