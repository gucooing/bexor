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
// source: google/protobuf/field_mask.proto

// Package fieldmaskpb contains generated types for google/protobuf/field_mask.proto.
//
// The FieldMask message represents a set of symbolic field paths.
// The paths are specific to some target message type,
// which is not stored within the FieldMask message itself.
//
// # Constructing a FieldMask
//
// The New function is used construct a FieldMask:
//
//	var messageType *descriptorpb.DescriptorProto
//	fm, err := fieldmaskpb.New(messageType, "field.name", "field.number")
//	if err != nil {
//		... // handle error
//	}
//	... // make use of fm
//
// The "field.name" and "field.number" paths are valid paths according to the
// google.protobuf.DescriptorProto message. Use of a path that does not correlate
// to valid fields reachable from DescriptorProto would result in an error.
//
// Once a FieldMask message has been constructed,
// the Append method can be used to insert additional paths to the path set:
//
//	var messageType *descriptorpb.DescriptorProto
//	if err := fm.Append(messageType, "options"); err != nil {
//		... // handle error
//	}
//
// # Type checking a FieldMask
//
// In order to verify that a FieldMask represents a set of fields that are
// reachable from some target message type, use the IsValid method:
//
//	var messageType *descriptorpb.DescriptorProto
//	if fm.IsValid(messageType) {
//		... // make use of fm
//	}
//
// IsValid needs to be passed the target message type as an input since the
// FieldMask message itself does not store the message type that the set of paths
// are for.
package fieldmaskpb

import (
	proto "github.com/gucooing/bexor/proto"
	protoreflect "github.com/gucooing/bexor/reflect/protoreflect"
	protoimpl "github.com/gucooing/bexor/runtime/protoimpl"
	reflect "reflect"
	sort "sort"
	strings "strings"
	sync "sync"
	unsafe "unsafe"
)

// `FieldMask` represents a set of symbolic field paths, for example:
//
//	paths: "f.a"
//	paths: "f.b.d"
//
// Here `f` represents a field in some root message, `a` and `b`
// fields in the message found in `f`, and `d` a field found in the
// message in `f.b`.
//
// Field masks are used to specify a subset of fields that should be
// returned by a get operation or modified by an update operation.
// Field masks also have a custom JSON encoding (see below).
//
// # Field Masks in Projections
//
// When used in the context of a projection, a response message or
// sub-message is filtered by the API to only contain those fields as
// specified in the mask. For example, if the mask in the previous
// example is applied to a response message as follows:
//
//	f {
//	  a : 22
//	  b {
//	    d : 1
//	    x : 2
//	  }
//	  y : 13
//	}
//	z: 8
//
// The result will not contain specific values for fields x,y and z
// (their value will be set to the default, and omitted in proto text
// output):
//
//	f {
//	  a : 22
//	  b {
//	    d : 1
//	  }
//	}
//
// A repeated field is not allowed except at the last position of a
// paths string.
//
// If a FieldMask object is not present in a get operation, the
// operation applies to all fields (as if a FieldMask of all fields
// had been specified).
//
// Note that a field mask does not necessarily apply to the
// top-level response message. In case of a REST get operation, the
// field mask applies directly to the response, but in case of a REST
// list operation, the mask instead applies to each individual message
// in the returned resource list. In case of a REST custom method,
// other definitions may be used. Where the mask applies will be
// clearly documented together with its declaration in the API.  In
// any case, the effect on the returned resource/resources is required
// behavior for APIs.
//
// # Field Masks in Update Operations
//
// A field mask in update operations specifies which fields of the
// targeted resource are going to be updated. The API is required
// to only change the values of the fields as specified in the mask
// and leave the others untouched. If a resource is passed in to
// describe the updated values, the API ignores the values of all
// fields not covered by the mask.
//
// If a repeated field is specified for an update operation, new values will
// be appended to the existing repeated field in the target resource. Note that
// a repeated field is only allowed in the last position of a `paths` string.
//
// If a sub-message is specified in the last position of the field mask for an
// update operation, then new value will be merged into the existing sub-message
// in the target resource.
//
// For example, given the target message:
//
//	f {
//	  b {
//	    d: 1
//	    x: 2
//	  }
//	  c: [1]
//	}
//
// And an update message:
//
//	f {
//	  b {
//	    d: 10
//	  }
//	  c: [2]
//	}
//
// then if the field mask is:
//
//	paths: ["f.b", "f.c"]
//
// then the result will be:
//
//	f {
//	  b {
//	    d: 10
//	    x: 2
//	  }
//	  c: [1, 2]
//	}
//
// An implementation may provide options to override this default behavior for
// repeated and message fields.
//
// In order to reset a field's value to the default, the field must
// be in the mask and set to the default value in the provided resource.
// Hence, in order to reset all fields of a resource, provide a default
// instance of the resource and set all fields in the mask, or do
// not provide a mask as described below.
//
// If a field mask is not present on update, the operation applies to
// all fields (as if a field mask of all fields has been specified).
// Note that in the presence of schema evolution, this may mean that
// fields the client does not know and has therefore not filled into
// the request will be reset to their default. If this is unwanted
// behavior, a specific service may require a client to always specify
// a field mask, producing an error if not.
//
// As with get operations, the location of the resource which
// describes the updated values in the request message depends on the
// operation kind. In any case, the effect of the field mask is
// required to be honored by the API.
//
// ## Considerations for HTTP REST
//
// The HTTP kind of an update operation which uses a field mask must
// be set to PATCH instead of PUT in order to satisfy HTTP semantics
// (PUT must only be used for full updates).
//
// # JSON Encoding of Field Masks
//
// In JSON, a field mask is encoded as a single string where paths are
// separated by a comma. Fields name in each path are converted
// to/from lower-camel naming conventions.
//
// As an example, consider the following message declarations:
//
//	message Profile {
//	  User user = 1;
//	  Photo photo = 2;
//	}
//	message User {
//	  string display_name = 1;
//	  string address = 2;
//	}
//
// In proto a field mask for `Profile` may look as such:
//
//	mask {
//	  paths: "user.display_name"
//	  paths: "photo"
//	}
//
// In JSON, the same mask is represented as below:
//
//	{
//	  mask: "user.displayName,photo"
//	}
//
// # Field Masks and Oneof Fields
//
// Field masks treat fields in oneofs just as regular fields. Consider the
// following message:
//
//	message SampleMessage {
//	  oneof test_oneof {
//	    string name = 4;
//	    SubMessage sub_message = 9;
//	  }
//	}
//
// The field mask can be:
//
//	mask {
//	  paths: "name"
//	}
//
// Or:
//
//	mask {
//	  paths: "sub_message"
//	}
//
// Note that oneof type names ("test_oneof" in this case) cannot be used in
// paths.
//
// ## Field Mask Verification
//
// The implementation of any API method which has a FieldMask type field in the
// request should verify the included field paths, and return an
// `INVALID_ARGUMENT` error if any path is unmappable.
type FieldMask struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The set of field mask paths.
	Paths         []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// New constructs a field mask from a list of paths and verifies that
// each one is valid according to the specified message type.
func New(m proto.Message, paths ...string) (*FieldMask, error) {
	x := new(FieldMask)
	return x, x.Append(m, paths...)
}

// Union returns the union of all the paths in the input field masks.
func Union(mx *FieldMask, my *FieldMask, ms ...*FieldMask) *FieldMask {
	var out []string
	out = append(out, mx.GetPaths()...)
	out = append(out, my.GetPaths()...)
	for _, m := range ms {
		out = append(out, m.GetPaths()...)
	}
	return &FieldMask{Paths: normalizePaths(out)}
}

// Intersect returns the intersection of all the paths in the input field masks.
func Intersect(mx *FieldMask, my *FieldMask, ms ...*FieldMask) *FieldMask {
	var ss1, ss2 []string // reused buffers for performance
	intersect := func(out, in []string) []string {
		ss1 = normalizePaths(append(ss1[:0], in...))
		ss2 = normalizePaths(append(ss2[:0], out...))
		out = out[:0]
		for i1, i2 := 0, 0; i1 < len(ss1) && i2 < len(ss2); {
			switch s1, s2 := ss1[i1], ss2[i2]; {
			case hasPathPrefix(s1, s2):
				out = append(out, s1)
				i1++
			case hasPathPrefix(s2, s1):
				out = append(out, s2)
				i2++
			case lessPath(s1, s2):
				i1++
			case lessPath(s2, s1):
				i2++
			}
		}
		return out
	}

	out := Union(mx, my, ms...).GetPaths()
	out = intersect(out, mx.GetPaths())
	out = intersect(out, my.GetPaths())
	for _, m := range ms {
		out = intersect(out, m.GetPaths())
	}
	return &FieldMask{Paths: normalizePaths(out)}
}

// IsValid reports whether all the paths are syntactically valid and
// refer to known fields in the specified message type.
// It reports false for a nil FieldMask.
func (x *FieldMask) IsValid(m proto.Message) bool {
	paths := x.GetPaths()
	return x != nil && numValidPaths(m, paths) == len(paths)
}

// Append appends a list of paths to the mask and verifies that each one
// is valid according to the specified message type.
// An invalid path is not appended and breaks insertion of subsequent paths.
func (x *FieldMask) Append(m proto.Message, paths ...string) error {
	numValid := numValidPaths(m, paths)
	x.Paths = append(x.Paths, paths[:numValid]...)
	paths = paths[numValid:]
	if len(paths) > 0 {
		name := m.ProtoReflect().Descriptor().FullName()
		return protoimpl.X.NewError("invalid path %q for message %q", paths[0], name)
	}
	return nil
}

func numValidPaths(m proto.Message, paths []string) int {
	md0 := m.ProtoReflect().Descriptor()
	for i, path := range paths {
		md := md0
		if !rangeFields(path, func(field string) bool {
			// Search the field within the message.
			if md == nil {
				return false // not within a message
			}
			fd := md.Fields().ByName(protoreflect.Name(field))
			// The real field name of a group is the message name.
			if fd == nil {
				gd := md.Fields().ByName(protoreflect.Name(strings.ToLower(field)))
				if gd != nil && gd.Kind() == protoreflect.GroupKind && string(gd.Message().Name()) == field {
					fd = gd
				}
			} else if fd.Kind() == protoreflect.GroupKind && string(fd.Message().Name()) != field {
				fd = nil
			}
			if fd == nil {
				return false // message has does not have this field
			}

			// Identify the next message to search within.
			md = fd.Message() // may be nil

			// Repeated fields are only allowed at the last position.
			if fd.IsList() || fd.IsMap() {
				md = nil
			}

			return true
		}) {
			return i
		}
	}
	return len(paths)
}

// Normalize converts the mask to its canonical form where all paths are sorted
// and redundant paths are removed.
func (x *FieldMask) Normalize() {
	x.Paths = normalizePaths(x.Paths)
}

func normalizePaths(paths []string) []string {
	sort.Slice(paths, func(i, j int) bool {
		return lessPath(paths[i], paths[j])
	})

	// Elide any path that is a prefix match on the previous.
	out := paths[:0]
	for _, path := range paths {
		if len(out) > 0 && hasPathPrefix(path, out[len(out)-1]) {
			continue
		}
		out = append(out, path)
	}
	return out
}

// hasPathPrefix is like strings.HasPrefix, but further checks for either
// an exact matche or that the prefix is delimited by a dot.
func hasPathPrefix(path, prefix string) bool {
	return strings.HasPrefix(path, prefix) && (len(path) == len(prefix) || path[len(prefix)] == '.')
}

// lessPath is a lexicographical comparison where dot is specially treated
// as the smallest symbol.
func lessPath(x, y string) bool {
	for i := 0; i < len(x) && i < len(y); i++ {
		if x[i] != y[i] {
			return (x[i] - '.') < (y[i] - '.')
		}
	}
	return len(x) < len(y)
}

// rangeFields is like strings.Split(path, "."), but avoids allocations by
// iterating over each field in place and calling a iterator function.
func rangeFields(path string, f func(field string) bool) bool {
	for {
		var field string
		if i := strings.IndexByte(path, '.'); i >= 0 {
			field, path = path[:i], path[i:]
		} else {
			field, path = path, ""
		}

		if !f(field) {
			return false
		}

		if len(path) == 0 {
			return true
		}
		path = strings.TrimPrefix(path, ".")
	}
}

func (x *FieldMask) Reset() {
	*x = FieldMask{}
	mi := &file_google_protobuf_field_mask_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMask) ProtoMessage() {}

func (x *FieldMask) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_field_mask_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMask.ProtoReflect.Descriptor instead.
func (*FieldMask) Descriptor() ([]byte, []int) {
	return file_google_protobuf_field_mask_proto_rawDescGZIP(), []int{0}
}

func (x *FieldMask) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

var File_google_protobuf_field_mask_proto protoreflect.FileDescriptor

const file_google_protobuf_field_mask_proto_rawDesc = "" +
	"\n" +
	" google/protobuf/field_mask.proto\x12\x0fgoogle.protobuf\"!\n" +
	"\tFieldMask\x12\x14\n" +
	"\x05paths\x18\x01 \x03(\tR\x05pathsB\x85\x01\n" +
	"\x13com.google.protobufB\x0eFieldMaskProtoP\x01Z2github.com/gucooing/bexor/types/known/fieldmaskpb\xf8\x01\x01\xa2\x02\x03GPB\xaa\x02\x1eGoogle.Protobuf.WellKnownTypesb\x06proto3"

var (
	file_google_protobuf_field_mask_proto_rawDescOnce sync.Once
	file_google_protobuf_field_mask_proto_rawDescData []byte
)

func file_google_protobuf_field_mask_proto_rawDescGZIP() []byte {
	file_google_protobuf_field_mask_proto_rawDescOnce.Do(func() {
		file_google_protobuf_field_mask_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_protobuf_field_mask_proto_rawDesc), len(file_google_protobuf_field_mask_proto_rawDesc)))
	})
	return file_google_protobuf_field_mask_proto_rawDescData
}

var file_google_protobuf_field_mask_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_field_mask_proto_goTypes = []any{
	(*FieldMask)(nil), // 0: google.protobuf.FieldMask
}
var file_google_protobuf_field_mask_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_protobuf_field_mask_proto_init() }
func file_google_protobuf_field_mask_proto_init() {
	if File_google_protobuf_field_mask_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_protobuf_field_mask_proto_rawDesc), len(file_google_protobuf_field_mask_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_field_mask_proto_goTypes,
		DependencyIndexes: file_google_protobuf_field_mask_proto_depIdxs,
		MessageInfos:      file_google_protobuf_field_mask_proto_msgTypes,
	}.Build()
	File_google_protobuf_field_mask_proto = out.File
	file_google_protobuf_field_mask_proto_goTypes = nil
	file_google_protobuf_field_mask_proto_depIdxs = nil
}
