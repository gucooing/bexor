// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gucooing/bexor/encoding/prototext"
	"github.com/gucooing/bexor/internal/genid"
	"github.com/gucooing/bexor/proto"
	"github.com/gucooing/bexor/testing/protocmp"

	"github.com/gucooing/bexor/types/descriptorpb"
)

func TestAnnotations(t *testing.T) {
	sourceFile, err := os.ReadFile("testdata/annotations/annotations.pb.go")
	if err != nil {
		t.Fatal(err)
	}
	metaFile, err := os.ReadFile("testdata/annotations/annotations.pb.go.meta")
	if err != nil {
		t.Fatal(err)
	}
	gotInfo := &descriptorpb.GeneratedCodeInfo{}
	if err := prototext.Unmarshal(metaFile, gotInfo); err != nil {
		t.Fatalf("can't parse meta file: %v", err)
	}

	wantInfo := &descriptorpb.GeneratedCodeInfo{}
	for _, want := range []struct {
		prefix, text, suffix string
		annotation           *descriptorpb.GeneratedCodeInfo_Annotation
	}{{
		"type ", "AnnotationsTestEnum", " int32",
		&descriptorpb.GeneratedCodeInfo_Annotation{
			Path: []int32{int32(genid.FileDescriptorProto_EnumType_field_number), 0},
		},
	}, {
		"\t", "AnnotationsTestEnum_ANNOTATIONS_TEST_ENUM_VALUE", " AnnotationsTestEnum = 0",
		&descriptorpb.GeneratedCodeInfo_Annotation{
			Path: []int32{int32(genid.FileDescriptorProto_EnumType_field_number), 0, int32(genid.EnumDescriptorProto_Value_field_number), 0},
		},
	}, {
		"type ", "AnnotationsTestMessage", " struct {",
		&descriptorpb.GeneratedCodeInfo_Annotation{
			Path: []int32{int32(genid.FileDescriptorProto_MessageType_field_number), 0},
		},
	}, {
		"\t", "AnnotationsTestField", " ",
		&descriptorpb.GeneratedCodeInfo_Annotation{
			Path: []int32{int32(genid.FileDescriptorProto_MessageType_field_number), 0, int32(genid.DescriptorProto_Field_field_number), 0},
		},
	}, {
		"func (x *AnnotationsTestMessage) ", "GetAnnotationsTestField", "() string {",
		&descriptorpb.GeneratedCodeInfo_Annotation{
			Path: []int32{int32(genid.FileDescriptorProto_MessageType_field_number), 0, int32(genid.DescriptorProto_Field_field_number), 0},
		},
	}} {
		s := want.prefix + want.text + want.suffix
		pos := bytes.Index(sourceFile, []byte(s))
		if pos < 0 {
			t.Errorf("source file does not contain: %v", s)
			continue
		}
		begin := pos + len(want.prefix)
		end := begin + len(want.text)
		a := &descriptorpb.GeneratedCodeInfo_Annotation{
			Begin:      proto.Int32(int32(begin)),
			End:        proto.Int32(int32(end)),
			SourceFile: proto.String("cmd/protoc-gen-go/testdata/annotations/annotations.proto"),
		}
		proto.Merge(a, want.annotation)
		wantInfo.Annotation = append(wantInfo.Annotation, a)
	}
	if diff := cmp.Diff(wantInfo, gotInfo, protocmp.Transform()); diff != "" {
		t.Fatalf("unexpected annotations for annotations.proto (-want +got):\n%s", diff)
	}
}
