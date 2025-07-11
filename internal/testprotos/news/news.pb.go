// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/news/news.proto

package news

import (
	protoreflect "github.com/gucooing/bexor/reflect/protoreflect"
	protoimpl "github.com/gucooing/bexor/runtime/protoimpl"
	anypb "github.com/gucooing/bexor/types/known/anypb"
	timestamppb "github.com/gucooing/bexor/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

type Article_Status int32

const (
	Article_DRAFT     Article_Status = 0
	Article_PUBLISHED Article_Status = 1
	Article_REVOKED   Article_Status = 2
)

// Enum value maps for Article_Status.
var (
	Article_Status_name = map[int32]string{
		0: "DRAFT",
		1: "PUBLISHED",
		2: "REVOKED",
	}
	Article_Status_value = map[string]int32{
		"DRAFT":     0,
		"PUBLISHED": 1,
		"REVOKED":   2,
	}
)

func (x Article_Status) Enum() *Article_Status {
	p := new(Article_Status)
	*p = x
	return p
}

func (x Article_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Article_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_news_news_proto_enumTypes[0].Descriptor()
}

func (Article_Status) Type() protoreflect.EnumType {
	return &file_internal_testprotos_news_news_proto_enumTypes[0]
}

func (x Article_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Article_Status.Descriptor instead.
func (Article_Status) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_news_news_proto_rawDescGZIP(), []int{0, 0}
}

type Article struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Author        string                 `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Date          *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Title         string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Status        Article_Status         `protobuf:"varint,8,opt,name=status,proto3,enum=google.golang.org.Article_Status" json:"status,omitempty"`
	Tags          []string               `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Attachments   []*anypb.Any           `protobuf:"bytes,6,rep,name=attachments,proto3" json:"attachments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Article) Reset() {
	*x = Article{}
	mi := &file_internal_testprotos_news_news_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_news_news_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_news_news_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Article) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetStatus() Article_Status {
	if x != nil {
		return x.Status
	}
	return Article_DRAFT
}

func (x *Article) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Article) GetAttachments() []*anypb.Any {
	if x != nil {
		return x.Attachments
	}
	return nil
}

type BinaryAttachment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BinaryAttachment) Reset() {
	*x = BinaryAttachment{}
	mi := &file_internal_testprotos_news_news_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinaryAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryAttachment) ProtoMessage() {}

func (x *BinaryAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_news_news_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryAttachment.ProtoReflect.Descriptor instead.
func (*BinaryAttachment) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_news_news_proto_rawDescGZIP(), []int{1}
}

func (x *BinaryAttachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BinaryAttachment) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type KeyValueAttachment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data          map[string]string      `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeyValueAttachment) Reset() {
	*x = KeyValueAttachment{}
	mi := &file_internal_testprotos_news_news_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeyValueAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueAttachment) ProtoMessage() {}

func (x *KeyValueAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_news_news_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueAttachment.ProtoReflect.Descriptor instead.
func (*KeyValueAttachment) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_news_news_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValueAttachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KeyValueAttachment) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_internal_testprotos_news_news_proto protoreflect.FileDescriptor

const file_internal_testprotos_news_news_proto_rawDesc = "" +
	"\n" +
	"#internal/testprotos/news/news.proto\x12\x11google.golang.org\x1a\x19google/protobuf/any.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb9\x02\n" +
	"\aArticle\x12\x16\n" +
	"\x06author\x18\x01 \x01(\tR\x06author\x12.\n" +
	"\x04date\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\x04date\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\x129\n" +
	"\x06status\x18\b \x01(\x0e2!.google.golang.org.Article.StatusR\x06status\x12\x12\n" +
	"\x04tags\x18\a \x03(\tR\x04tags\x126\n" +
	"\vattachments\x18\x06 \x03(\v2\x14.google.protobuf.AnyR\vattachments\"/\n" +
	"\x06Status\x12\t\n" +
	"\x05DRAFT\x10\x00\x12\r\n" +
	"\tPUBLISHED\x10\x01\x12\v\n" +
	"\aREVOKED\x10\x02\":\n" +
	"\x10BinaryAttachment\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x12\n" +
	"\x04data\x18\x02 \x01(\fR\x04data\"\xa6\x01\n" +
	"\x12KeyValueAttachment\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12C\n" +
	"\x04data\x18\x02 \x03(\v2/.google.golang.org.KeyValueAttachment.DataEntryR\x04data\x1a7\n" +
	"\tDataEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B5Z3github.com/gucooing/bexor/internal/testprotos/newsb\x06proto3"

var (
	file_internal_testprotos_news_news_proto_rawDescOnce sync.Once
	file_internal_testprotos_news_news_proto_rawDescData []byte
)

func file_internal_testprotos_news_news_proto_rawDescGZIP() []byte {
	file_internal_testprotos_news_news_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_news_news_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_testprotos_news_news_proto_rawDesc), len(file_internal_testprotos_news_news_proto_rawDesc)))
	})
	return file_internal_testprotos_news_news_proto_rawDescData
}

var file_internal_testprotos_news_news_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_testprotos_news_news_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_testprotos_news_news_proto_goTypes = []any{
	(Article_Status)(0),           // 0: google.golang.org.Article.Status
	(*Article)(nil),               // 1: google.golang.org.Article
	(*BinaryAttachment)(nil),      // 2: google.golang.org.BinaryAttachment
	(*KeyValueAttachment)(nil),    // 3: google.golang.org.KeyValueAttachment
	nil,                           // 4: google.golang.org.KeyValueAttachment.DataEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 6: google.protobuf.Any
}
var file_internal_testprotos_news_news_proto_depIdxs = []int32{
	5, // 0: google.golang.org.Article.date:type_name -> google.protobuf.Timestamp
	0, // 1: google.golang.org.Article.status:type_name -> google.golang.org.Article.Status
	6, // 2: google.golang.org.Article.attachments:type_name -> google.protobuf.Any
	4, // 3: google.golang.org.KeyValueAttachment.data:type_name -> google.golang.org.KeyValueAttachment.DataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_testprotos_news_news_proto_init() }
func file_internal_testprotos_news_news_proto_init() {
	if File_internal_testprotos_news_news_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_testprotos_news_news_proto_rawDesc), len(file_internal_testprotos_news_news_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_news_news_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_news_news_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_news_news_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_news_news_proto_msgTypes,
	}.Build()
	File_internal_testprotos_news_news_proto = out.File
	file_internal_testprotos_news_news_proto_goTypes = nil
	file_internal_testprotos_news_news_proto_depIdxs = nil
}
