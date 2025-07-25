// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto_test

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gucooing/bexor/encoding/prototext"
	"github.com/gucooing/bexor/proto"
	"github.com/gucooing/bexor/reflect/protoreflect"
	"github.com/gucooing/bexor/testing/protopack"
	"github.com/gucooing/bexor/types/known/durationpb"

	testpb "github.com/gucooing/bexor/internal/testprotos/test"
	test3pb "github.com/gucooing/bexor/internal/testprotos/test3"
)

func TestDecode(t *testing.T) {
	for _, test := range testValidMessages {
		if len(test.decodeTo) == 0 {
			t.Errorf("%v: no test message types", test.desc)
		}
		for _, want := range test.decodeTo {
			t.Run(fmt.Sprintf("%s (%T)", test.desc, want), func(t *testing.T) {
				opts := test.unmarshalOptions
				opts.AllowPartial = test.partial
				wire := append(([]byte)(nil), test.wire...)
				got := reflect.New(reflect.TypeOf(want).Elem()).Interface().(proto.Message)
				if err := opts.Unmarshal(wire, got); err != nil {
					t.Errorf("Unmarshal error: %v\nMessage:\n%v", err, prototext.Format(want))
					return
				}

				// Aliasing check: Unmarshal shouldn't modify the original wire
				// bytes, and modifying the original wire bytes shouldn't affect
				// the unmarshaled message.
				if !bytes.Equal(test.wire, wire) {
					t.Errorf("Unmarshal unexpectedly modified its input")
				}
				for i := range wire {
					wire[i] = 0
				}
				if !proto.Equal(got, want) && got.ProtoReflect().IsValid() && want.ProtoReflect().IsValid() {
					t.Errorf("Unmarshal returned unexpected result; got:\n%v\nwant:\n%v", prototext.Format(got), prototext.Format(want))
				}
			})
		}
	}
}

func TestDecodeRequiredFieldChecks(t *testing.T) {
	for _, test := range testValidMessages {
		if !test.partial {
			continue
		}
		for _, m := range test.decodeTo {
			t.Run(fmt.Sprintf("%s (%T)", test.desc, m), func(t *testing.T) {
				opts := test.unmarshalOptions
				opts.AllowPartial = false
				got := reflect.New(reflect.TypeOf(m).Elem()).Interface().(proto.Message)
				if err := proto.Unmarshal(test.wire, got); err == nil {
					t.Fatalf("Unmarshal succeeded (want error)\nMessage:\n%v", prototext.Format(got))
				}
			})
		}
	}
}

func TestDecodeInvalidMessages(t *testing.T) {
	for _, test := range testInvalidMessages {
		if len(test.decodeTo) == 0 {
			t.Errorf("%v: no test message types", test.desc)
		}
		for _, want := range test.decodeTo {
			t.Run(fmt.Sprintf("%s (%T)", test.desc, want), func(t *testing.T) {
				opts := test.unmarshalOptions
				opts.AllowPartial = test.partial
				got := want.ProtoReflect().New().Interface()
				if err := opts.Unmarshal(test.wire, got); err == nil {
					t.Errorf("Unmarshal unexpectedly succeeded\ninput bytes: [%x]\nMessage:\n%v", test.wire, prototext.Format(got))
				} else if !errors.Is(err, proto.Error) {
					t.Errorf("Unmarshal error is not a proto.Error: %v", err)
				}
			})
		}
	}
}

func TestDecodeZeroLengthBytes(t *testing.T) {
	// Verify that proto3 bytes fields don't give the mistaken
	// impression that they preserve presence.
	wire := protopack.Message{
		protopack.Tag{94, protopack.BytesType}, protopack.Bytes(nil),
	}.Marshal()
	m := &test3pb.TestAllTypes{}
	if err := proto.Unmarshal(wire, m); err != nil {
		t.Fatal(err)
	}
	if m.OptionalBytes != nil {
		t.Errorf("unmarshal zero-length proto3 bytes field: got %v, want nil", m.OptionalBytes)
	}
}

func TestDecodeOneofNilWrapper(t *testing.T) {
	wire := protopack.Message{
		protopack.Tag{111, protopack.VarintType}, protopack.Varint(1111),
	}.Marshal()
	m := &testpb.TestAllTypes{OneofField: (*testpb.TestAllTypes_OneofUint32)(nil)}
	if err := proto.Unmarshal(wire, m); err != nil {
		t.Fatal(err)
	}
	if got := m.GetOneofUint32(); got != 1111 {
		t.Errorf("GetOneofUint32() = %v, want %v", got, 1111)
	}
}

func TestDecodeEmptyBytes(t *testing.T) {
	// There's really nothing wrong with a nil entry in a [][]byte,
	// but we take care to produce non-nil []bytes for zero-length
	// byte strings, so test for it.
	m := &testpb.TestAllTypes{}
	b := protopack.Message{
		protopack.Tag{45, protopack.BytesType}, protopack.Bytes(nil),
	}.Marshal()
	if err := proto.Unmarshal(b, m); err != nil {
		t.Fatal(err)
	}
	if m.RepeatedBytes[0] == nil {
		t.Errorf("unmarshaling repeated bytes field containing zero-length value: Got nil bytes, want non-nil")
	}
}

func build(m proto.Message, opts ...buildOpt) proto.Message {
	for _, opt := range opts {
		opt(m)
	}
	return m
}

type buildOpt func(proto.Message)

func unknown(raw protoreflect.RawFields) buildOpt {
	return func(m proto.Message) {
		m.ProtoReflect().SetUnknown(raw)
	}
}

func extend(desc protoreflect.ExtensionType, value any) buildOpt {
	return func(m proto.Message) {
		proto.SetExtension(m, desc, value)
	}
}

// This example illustrates how to unmarshal (decode) wire format encoding into
// a Protobuf message.
func ExampleUnmarshal() {
	// This is the wire format encoding produced by the Marshal example.
	// Typically you would read from the network, from disk, etc.
	b := []byte{0x10, 0x7d}

	var dur durationpb.Duration
	if err := proto.Unmarshal(b, &dur); err != nil {
		panic(err)
	}

	fmt.Printf("Protobuf wire format decoded to duration %v\n", dur.AsDuration())

	// Output: Protobuf wire format decoded to duration 125ns
}
