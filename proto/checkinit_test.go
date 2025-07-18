// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gucooing/bexor/encoding/prototext"
	"github.com/gucooing/bexor/proto"

	testpb "github.com/gucooing/bexor/internal/testprotos/test"
	testeditionspb "github.com/gucooing/bexor/internal/testprotos/testeditions"
)

func TestCheckInitializedErrors(t *testing.T) {
	type test struct {
		m    proto.Message
		want string
		skip bool
	}
	tests := []test{{
		m:    &testpb.TestRequired{},
		want: `goproto.proto.test.TestRequired.required_field`,
	}, {
		m: &testpb.TestRequiredForeign{
			OptionalMessage: &testpb.TestRequired{},
		},
		want: `goproto.proto.test.TestRequired.required_field`,
	}, {
		m: &testpb.TestRequiredForeign{
			RepeatedMessage: []*testpb.TestRequired{
				{RequiredField: proto.Int32(1)},
				{},
			},
		},
		want: `goproto.proto.test.TestRequired.required_field`,
	}, {
		m: &testpb.TestRequiredForeign{
			MapMessage: map[int32]*testpb.TestRequired{
				1: {},
			},
		},
		want: `goproto.proto.test.TestRequired.required_field`,
	}, {
		m:    &testeditionspb.TestRequired{},
		want: `goproto.proto.testeditions.TestRequired.required_field`,
	}, {
		m: &testeditionspb.TestRequiredForeign{
			OptionalMessage: &testeditionspb.TestRequired{},
		},
		want: `goproto.proto.testeditions.TestRequired.required_field`,
	}, {
		m: &testeditionspb.TestRequiredForeign{
			RepeatedMessage: []*testeditionspb.TestRequired{
				{RequiredField: proto.Int32(1)},
				{},
			},
		},
		want: `goproto.proto.testeditions.TestRequired.required_field`,
	}, {
		m: &testeditionspb.TestRequiredForeign{
			MapMessage: map[int32]*testeditionspb.TestRequired{
				1: {},
			},
		},
		want: `goproto.proto.testeditions.TestRequired.required_field`,
	}}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if tt.skip {
				t.SkipNow()
			}

			err := proto.CheckInitialized(tt.m)
			got := "<nil>"
			if err != nil {
				got = fmt.Sprintf("%q", err)
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("CheckInitialized(m):\n got: %v\nwant contains: %v\nMessage:\n%v", got, tt.want, prototext.Format(tt.m))
			}
		})
	}
}
