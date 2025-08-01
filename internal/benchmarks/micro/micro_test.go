// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This package contains microbenchmarks exercising specific areas of interest.
// The benchmarks here are not comprehensive and are not necessarily indicative
// real-world performance.

package micro_test

import (
	"testing"

	"github.com/gucooing/bexor/internal/impl"
	"github.com/gucooing/bexor/proto"
	"github.com/gucooing/bexor/runtime/protoiface"
	"github.com/gucooing/bexor/types/known/emptypb"

	micropb "github.com/gucooing/bexor/internal/testprotos/benchmarks/micro"
	testpb "github.com/gucooing/bexor/internal/testprotos/test"
)

// BenchmarkEmptyMessage tests a google.protobuf.Empty.
//
// It measures per-operation overhead.
func BenchmarkEmptyMessage(b *testing.B) {
	b.Run("Wire/Marshal", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			m := &emptypb.Empty{}
			for pb.Next() {
				if _, err := proto.Marshal(m); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("Wire/Unmarshal", func(b *testing.B) {
		opts := proto.UnmarshalOptions{
			Merge: true,
		}
		b.RunParallel(func(pb *testing.PB) {
			m := &emptypb.Empty{}
			for pb.Next() {
				if err := opts.Unmarshal([]byte{}, m); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("Wire/Validate", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			mt := (&emptypb.Empty{}).ProtoReflect().Type()
			for pb.Next() {
				_, got := impl.Validate(mt, protoiface.UnmarshalInput{})
				want := impl.ValidationValid
				if got != want {
					b.Fatalf("Validate = %v, want %v", got, want)
				}
			}
		})
	})
	b.Run("Clone", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			m := &emptypb.Empty{}
			for pb.Next() {
				proto.Clone(m)
			}
		})
	})
	b.Run("New", func(b *testing.B) {
		mt := (&emptypb.Empty{}).ProtoReflect().Type()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mt.New()
			}
		})
	})
}

// BenchmarkRepeatedInt32 tests a message containing 500 non-packed repeated int32s.
//
// For unmarshal operations, it measures the cost of the field decode loop, since each
// item in the repeated field has an individual tag and value.
func BenchmarkRepeatedInt32(b *testing.B) {
	m := &testpb.TestAllTypes{}
	for i := int32(0); i < 500; i++ {
		m.RepeatedInt32 = append(m.RepeatedInt32, i)
	}
	w, err := proto.Marshal(m)
	if err != nil {
		b.Fatal(err)
	}
	b.Run("Wire/Marshal", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := proto.Marshal(m); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("Wire/Unmarshal", func(b *testing.B) {
		opts := proto.UnmarshalOptions{
			Merge: true,
		}
		b.RunParallel(func(pb *testing.PB) {
			m := &testpb.TestAllTypes{}
			for pb.Next() {
				m.RepeatedInt32 = m.RepeatedInt32[:0]
				if err := opts.Unmarshal(w, m); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("Wire/Validate", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			mt := (&testpb.TestAllTypes{}).ProtoReflect().Type()
			for pb.Next() {
				_, got := impl.Validate(mt, protoiface.UnmarshalInput{
					Buf: w,
				})
				want := impl.ValidationValid
				if got != want {
					b.Fatalf("Validate = %v, want %v", got, want)
				}
			}
		})
	})
	b.Run("Clone", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				proto.Clone(m)
			}
		})
	})
}

// BenchmarkRequired tests a message containing a required field.
func BenchmarkRequired(b *testing.B) {
	m := &micropb.SixteenRequired{
		F1:  proto.Int32(1),
		F2:  proto.Int32(1),
		F3:  proto.Int32(1),
		F4:  proto.Int32(1),
		F5:  proto.Int32(1),
		F6:  proto.Int32(1),
		F7:  proto.Int32(1),
		F8:  proto.Int32(1),
		F9:  proto.Int32(1),
		F10: proto.Int32(1),
		F11: proto.Int32(1),
		F12: proto.Int32(1),
		F13: proto.Int32(1),
		F14: proto.Int32(1),
		F15: proto.Int32(1),
		F16: proto.Int32(1),
	}
	w, err := proto.Marshal(m)
	if err != nil {
		b.Fatal(err)
	}
	b.Run("Wire/Marshal", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := proto.Marshal(m); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("Wire/Unmarshal", func(b *testing.B) {
		opts := proto.UnmarshalOptions{
			Merge: true,
		}
		b.RunParallel(func(pb *testing.PB) {
			m := &micropb.SixteenRequired{}
			for pb.Next() {
				if err := opts.Unmarshal(w, m); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
	b.Run("Wire/Validate", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			mt := (&micropb.SixteenRequired{}).ProtoReflect().Type()
			for pb.Next() {
				_, got := impl.Validate(mt, protoiface.UnmarshalInput{
					Buf: w,
				})
				want := impl.ValidationValid
				if got != want {
					b.Fatalf("Validate = %v, want %v", got, want)
				}
			}
		})
	})
	b.Run("Clone", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				proto.Clone(m)
			}
		})
	})
}

func BenchmarkExtension(b *testing.B) {
	b.Run("Has/None", func(b *testing.B) {
		m := &testpb.TestAllExtensions{}
		for i := 0; i < b.N; i++ {
			proto.HasExtension(m, testpb.E_OptionalNestedMessage)
		}
	})
	b.Run("Has/Set", func(b *testing.B) {
		m := &testpb.TestAllExtensions{}
		ext := &testpb.TestAllExtensions_NestedMessage{A: proto.Int32(-32)}
		proto.SetExtension(m, testpb.E_OptionalNestedMessage, ext)
		for i := 0; i < b.N; i++ {
			proto.HasExtension(m, testpb.E_OptionalNestedMessage)
		}
	})
	b.Run("Get/None", func(b *testing.B) {
		m := &testpb.TestAllExtensions{}
		for i := 0; i < b.N; i++ {
			proto.GetExtension(m, testpb.E_OptionalNestedMessage)
		}
	})
	b.Run("Get/Set", func(b *testing.B) {
		m := &testpb.TestAllExtensions{}
		ext := &testpb.TestAllExtensions_NestedMessage{A: proto.Int32(-32)}
		proto.SetExtension(m, testpb.E_OptionalNestedMessage, ext)
		for i := 0; i < b.N; i++ {
			proto.GetExtension(m, testpb.E_OptionalNestedMessage)
		}
	})
	b.Run("Set", func(b *testing.B) {
		m := &testpb.TestAllExtensions{}
		ext := &testpb.TestAllExtensions_NestedMessage{A: proto.Int32(-32)}
		for i := 0; i < b.N; i++ {
			proto.SetExtension(m, testpb.E_OptionalNestedMessage, ext)
		}
	})
}
