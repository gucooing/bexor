// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gucooing/bexor/internal/impl"
	"github.com/gucooing/bexor/internal/protobuild"
	"github.com/gucooing/bexor/reflect/protoreflect"
	"github.com/gucooing/bexor/reflect/protoregistry"
	"github.com/gucooing/bexor/testing/protopack"

	_ "github.com/gucooing/bexor/internal/testprotos/lazy"
	_ "github.com/gucooing/bexor/internal/testprotos/lazy/lazy_opaque"
	_ "github.com/gucooing/bexor/internal/testprotos/required"
	_ "github.com/gucooing/bexor/internal/testprotos/required/required_opaque"
	_ "github.com/gucooing/bexor/internal/testprotos/test"
	_ "github.com/gucooing/bexor/internal/testprotos/test3"
	_ "github.com/gucooing/bexor/internal/testprotos/testeditions/testeditions_hybrid"
	_ "github.com/gucooing/bexor/internal/testprotos/testeditions/testeditions_opaque"
)

var testLazyUnmarshal = flag.Bool("test_lazy_unmarshal", false, "test lazy proto.Unmarshal")

func TestMain(m *testing.M) {
	flag.Parse()
	impl.EnableLazyUnmarshal(*testLazyUnmarshal)
	os.Exit(m.Run())
}

var relatedMessages = func() map[protoreflect.MessageType][]protoreflect.MessageType {
	related := map[protoreflect.MessageType][]protoreflect.MessageType{}
	const opaqueNamePrefix = "opaque."
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		name := mt.Descriptor().FullName()
		if !strings.HasPrefix(string(name), opaqueNamePrefix) {
			return true
		}
		mt1, err := protoregistry.GlobalTypes.FindMessageByName(name[len(opaqueNamePrefix):])
		if err != nil {
			panic(fmt.Sprintf("%v: can't find related message", name))
		}
		related[mt1] = append(related[mt1], mt)
		return true
	})
	return related
}()

func init() {
	testValidMessages = append(testValidMessages, []testProto{
		{
			desc:          "lazy field contains wrong wire type",
			checkFastInit: true,
			decodeTo: makeMessages(protobuild.Message{
				"optional_nested_message": protobuild.Message{
					protobuild.Unknown: protopack.Message{
						protopack.Tag{2, protopack.VarintType}, protopack.Varint(3),
					}.Marshal(),
				},
			}),
			wire: protopack.Message{
				protopack.Tag{18, protopack.BytesType}, protopack.LengthPrefix(protopack.Message{
					protopack.Tag{2, protopack.VarintType}, protopack.Varint(3),
				}),
			}.Marshal(),
		}, {
			desc:          "lazy field contains right and wrong wire type",
			checkFastInit: true,
			decodeTo: makeMessages(protobuild.Message{
				"optional_nested_message": protobuild.Message{
					"corecursive": protobuild.Message{
						"optional_int32": 2,
					},
					protobuild.Unknown: protopack.Message{
						protopack.Tag{2, protopack.VarintType}, protopack.Varint(3),
					}.Marshal(),
				},
			}),
			wire: protopack.Message{
				protopack.Tag{18, protopack.BytesType}, protopack.LengthPrefix(protopack.Message{
					protopack.Tag{2, protopack.BytesType}, protopack.LengthPrefix(protopack.Message{
						protopack.Tag{1, protopack.VarintType}, protopack.Varint(2),
					}),
					protopack.Tag{2, protopack.VarintType}, protopack.Varint(3),
				}),
			}.Marshal(),
		},
	}...)
}
