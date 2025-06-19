// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prototest_test

import (
	"fmt"
	"testing"

	"github.com/gucooing/bexor/proto"
	"github.com/gucooing/bexor/runtime/protoimpl"
	"github.com/gucooing/bexor/testing/prototest"

	irregularpb "github.com/gucooing/bexor/internal/testprotos/irregular"
	legacypb "github.com/gucooing/bexor/internal/testprotos/legacy"
	legacy1pb "github.com/gucooing/bexor/internal/testprotos/legacy/proto2_20160225_2fc053c5"
	testpb "github.com/gucooing/bexor/internal/testprotos/test"
	test3pb "github.com/gucooing/bexor/internal/testprotos/test3"
	testeditionspb "github.com/gucooing/bexor/internal/testprotos/testeditions"
)

func Test(t *testing.T) {
	ms := []proto.Message{
		(*testpb.TestAllTypes)(nil),
		(*test3pb.TestAllTypes)(nil),
		(*testeditionspb.TestAllTypes)(nil),
		(*testpb.TestRequired)(nil),
		(*testeditionspb.TestRequired)(nil),
		(*irregularpb.Message)(nil),
		(*testpb.TestAllExtensions)(nil),
		(*testeditionspb.TestAllExtensions)(nil),
		(*legacypb.Legacy)(nil),
		protoimpl.X.MessageOf((*legacy1pb.Message)(nil)).Interface(),
	}

	for _, m := range ms {
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			prototest.Message{}.Test(t, m.ProtoReflect().Type())
		})
	}
}
