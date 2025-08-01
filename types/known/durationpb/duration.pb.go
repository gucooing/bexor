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
// source: google/protobuf/duration.proto

// Package durationpb contains generated types for google/protobuf/duration.proto.
//
// The Duration message represents a signed span of time.
//
// # Conversion to a Go Duration
//
// The AsDuration method can be used to convert a Duration message to a
// standard Go time.Duration value:
//
//	d := dur.AsDuration()
//	... // make use of d as a time.Duration
//
// Converting to a time.Duration is a common operation so that the extensive
// set of time-based operations provided by the time package can be leveraged.
// See https://golang.org/pkg/time for more information.
//
// The AsDuration method performs the conversion on a best-effort basis.
// Durations with denormal values (e.g., nanoseconds beyond -99999999 and
// +99999999, inclusive; or seconds and nanoseconds with opposite signs)
// are normalized during the conversion to a time.Duration. To manually check for
// invalid Duration per the documented limitations in duration.proto,
// additionally call the CheckValid method:
//
//	if err := dur.CheckValid(); err != nil {
//		... // handle error
//	}
//
// Note that the documented limitations in duration.proto does not protect a
// Duration from overflowing the representable range of a time.Duration in Go.
// The AsDuration method uses saturation arithmetic such that an overflow clamps
// the resulting value to the closest representable value (e.g., math.MaxInt64
// for positive overflow and math.MinInt64 for negative overflow).
//
// # Conversion from a Go Duration
//
// The durationpb.New function can be used to construct a Duration message
// from a standard Go time.Duration value:
//
//	dur := durationpb.New(d)
//	... // make use of d as a *durationpb.Duration
package durationpb

import (
	protoreflect "github.com/gucooing/bexor/reflect/protoreflect"
	protoimpl "github.com/gucooing/bexor/runtime/protoimpl"
	math "math"
	reflect "reflect"
	sync "sync"
	time "time"
	unsafe "unsafe"
)

// A Duration represents a signed, fixed-length span of time represented
// as a count of seconds and fractions of seconds at nanosecond
// resolution. It is independent of any calendar and concepts like "day"
// or "month". It is related to Timestamp in that the difference between
// two Timestamp values is a Duration and it can be added or subtracted
// from a Timestamp. Range is approximately +-10,000 years.
//
// # Examples
//
// Example 1: Compute Duration from two Timestamps in pseudo code.
//
//	Timestamp start = ...;
//	Timestamp end = ...;
//	Duration duration = ...;
//
//	duration.seconds = end.seconds - start.seconds;
//	duration.nanos = end.nanos - start.nanos;
//
//	if (duration.seconds < 0 && duration.nanos > 0) {
//	  duration.seconds += 1;
//	  duration.nanos -= 1000000000;
//	} else if (duration.seconds > 0 && duration.nanos < 0) {
//	  duration.seconds -= 1;
//	  duration.nanos += 1000000000;
//	}
//
// Example 2: Compute Timestamp from Timestamp + Duration in pseudo code.
//
//	Timestamp start = ...;
//	Duration duration = ...;
//	Timestamp end = ...;
//
//	end.seconds = start.seconds + duration.seconds;
//	end.nanos = start.nanos + duration.nanos;
//
//	if (end.nanos < 0) {
//	  end.seconds -= 1;
//	  end.nanos += 1000000000;
//	} else if (end.nanos >= 1000000000) {
//	  end.seconds += 1;
//	  end.nanos -= 1000000000;
//	}
//
// Example 3: Compute Duration from datetime.timedelta in Python.
//
//	td = datetime.timedelta(days=3, minutes=10)
//	duration = Duration()
//	duration.FromTimedelta(td)
//
// # JSON Mapping
//
// In JSON format, the Duration type is encoded as a string rather than an
// object, where the string ends in the suffix "s" (indicating seconds) and
// is preceded by the number of seconds, with nanoseconds expressed as
// fractional seconds. For example, 3 seconds with 0 nanoseconds should be
// encoded in JSON format as "3s", while 3 seconds and 1 nanosecond should
// be expressed in JSON format as "3.000000001s", and 3 seconds and 1
// microsecond should be expressed in JSON format as "3.000001s".
type Duration struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Signed seconds of the span of time. Must be from -315,576,000,000
	// to +315,576,000,000 inclusive. Note: these bounds are computed from:
	// 60 sec/min * 60 min/hr * 24 hr/day * 365.25 days/year * 10000 years
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Signed fractions of a second at nanosecond resolution of the span
	// of time. Durations less than one second are represented with a 0
	// `seconds` field and a positive or negative `nanos` field. For durations
	// of one second or more, a non-zero value for the `nanos` field must be
	// of the same sign as the `seconds` field. Must be from -999,999,999
	// to +999,999,999 inclusive.
	Nanos         int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

// New constructs a new Duration from the provided time.Duration.
func New(d time.Duration) *Duration {
	nanos := d.Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9
	return &Duration{Seconds: int64(secs), Nanos: int32(nanos)}
}

// AsDuration converts x to a time.Duration,
// returning the closest duration value in the event of overflow.
func (x *Duration) AsDuration() time.Duration {
	secs := x.GetSeconds()
	nanos := x.GetNanos()
	d := time.Duration(secs) * time.Second
	overflow := d/time.Second != time.Duration(secs)
	d += time.Duration(nanos) * time.Nanosecond
	overflow = overflow || (secs < 0 && nanos < 0 && d > 0)
	overflow = overflow || (secs > 0 && nanos > 0 && d < 0)
	if overflow {
		switch {
		case secs < 0:
			return time.Duration(math.MinInt64)
		case secs > 0:
			return time.Duration(math.MaxInt64)
		}
	}
	return d
}

// IsValid reports whether the duration is valid.
// It is equivalent to CheckValid == nil.
func (x *Duration) IsValid() bool {
	return x.check() == 0
}

// CheckValid returns an error if the duration is invalid.
// In particular, it checks whether the value is within the range of
// -10000 years to +10000 years inclusive.
// An error is reported for a nil Duration.
func (x *Duration) CheckValid() error {
	switch x.check() {
	case invalidNil:
		return protoimpl.X.NewError("invalid nil Duration")
	case invalidUnderflow:
		return protoimpl.X.NewError("duration (%v) exceeds -10000 years", x)
	case invalidOverflow:
		return protoimpl.X.NewError("duration (%v) exceeds +10000 years", x)
	case invalidNanosRange:
		return protoimpl.X.NewError("duration (%v) has out-of-range nanos", x)
	case invalidNanosSign:
		return protoimpl.X.NewError("duration (%v) has seconds and nanos with different signs", x)
	default:
		return nil
	}
}

const (
	_ = iota
	invalidNil
	invalidUnderflow
	invalidOverflow
	invalidNanosRange
	invalidNanosSign
)

func (x *Duration) check() uint {
	const absDuration = 315576000000 // 10000yr * 365.25day/yr * 24hr/day * 60min/hr * 60sec/min
	secs := x.GetSeconds()
	nanos := x.GetNanos()
	switch {
	case x == nil:
		return invalidNil
	case secs < -absDuration:
		return invalidUnderflow
	case secs > +absDuration:
		return invalidOverflow
	case nanos <= -1e9 || nanos >= +1e9:
		return invalidNanosRange
	case (secs > 0 && nanos < 0) || (secs < 0 && nanos > 0):
		return invalidNanosSign
	default:
		return 0
	}
}

func (x *Duration) Reset() {
	*x = Duration{}
	mi := &file_google_protobuf_duration_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Duration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Duration) ProtoMessage() {}

func (x *Duration) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_duration_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Duration.ProtoReflect.Descriptor instead.
func (*Duration) Descriptor() ([]byte, []int) {
	return file_google_protobuf_duration_proto_rawDescGZIP(), []int{0}
}

func (x *Duration) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Duration) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_google_protobuf_duration_proto protoreflect.FileDescriptor

const file_google_protobuf_duration_proto_rawDesc = "" +
	"\n" +
	"\x1egoogle/protobuf/duration.proto\x12\x0fgoogle.protobuf\":\n" +
	"\bDuration\x12\x18\n" +
	"\aseconds\x18\x01 \x01(\x03R\aseconds\x12\x14\n" +
	"\x05nanos\x18\x02 \x01(\x05R\x05nanosB\x83\x01\n" +
	"\x13com.google.protobufB\rDurationProtoP\x01Z1github.com/gucooing/bexor/types/known/durationpb\xf8\x01\x01\xa2\x02\x03GPB\xaa\x02\x1eGoogle.Protobuf.WellKnownTypesb\x06proto3"

var (
	file_google_protobuf_duration_proto_rawDescOnce sync.Once
	file_google_protobuf_duration_proto_rawDescData []byte
)

func file_google_protobuf_duration_proto_rawDescGZIP() []byte {
	file_google_protobuf_duration_proto_rawDescOnce.Do(func() {
		file_google_protobuf_duration_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_protobuf_duration_proto_rawDesc), len(file_google_protobuf_duration_proto_rawDesc)))
	})
	return file_google_protobuf_duration_proto_rawDescData
}

var file_google_protobuf_duration_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_duration_proto_goTypes = []any{
	(*Duration)(nil), // 0: google.protobuf.Duration
}
var file_google_protobuf_duration_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_protobuf_duration_proto_init() }
func file_google_protobuf_duration_proto_init() {
	if File_google_protobuf_duration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_protobuf_duration_proto_rawDesc), len(file_google_protobuf_duration_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_duration_proto_goTypes,
		DependencyIndexes: file_google_protobuf_duration_proto_depIdxs,
		MessageInfos:      file_google_protobuf_duration_proto_msgTypes,
	}.Build()
	File_google_protobuf_duration_proto = out.File
	file_google_protobuf_duration_proto_goTypes = nil
	file_google_protobuf_duration_proto_depIdxs = nil
}
