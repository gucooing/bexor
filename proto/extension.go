// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"github.com/gucooing/bexor/reflect/protoreflect"
)

// HasExtension reports whether an extension field is populated.
// It returns false if m is invalid or if xt does not extend m.
func HasExtension(m Message, xt protoreflect.ExtensionType) bool {
	// Treat nil message interface or descriptor as an empty message; no populated
	// fields.
	if m == nil || xt == nil {
		return false
	}

	// As a special-case, we reports invalid or mismatching descriptors
	// as always not being populated (since they aren't).
	mr := m.ProtoReflect()
	xd := xt.TypeDescriptor()
	if mr.Descriptor() != xd.ContainingMessage() {
		return false
	}

	return mr.Has(xd)
}

// ClearExtension clears an extension field such that subsequent
// [HasExtension] calls return false.
// It panics if m is invalid or if xt does not extend m.
func ClearExtension(m Message, xt protoreflect.ExtensionType) {
	m.ProtoReflect().Clear(xt.TypeDescriptor())
}

// GetExtension retrieves the value for an extension field.
// If the field is unpopulated, it returns the default value for
// scalars and an immutable, empty value for lists or messages.
// It panics if xt does not extend m.
//
// The type of the value is dependent on the field type of the extension.
// For extensions generated by protoc-gen-go, the Go type is as follows:
//
//	╔═══════════════════╤═════════════════════════╗
//	║ Go type           │ Protobuf kind           ║
//	╠═══════════════════╪═════════════════════════╣
//	║ bool              │ bool                    ║
//	║ int32             │ int32, sint32, sfixed32 ║
//	║ int64             │ int64, sint64, sfixed64 ║
//	║ uint32            │ uint32, fixed32         ║
//	║ uint64            │ uint64, fixed64         ║
//	║ float32           │ float                   ║
//	║ float64           │ double                  ║
//	║ string            │ string                  ║
//	║ []byte            │ bytes                   ║
//	║ protoreflect.Enum │ enum                    ║
//	║ proto.Message     │ message, group          ║
//	╚═══════════════════╧═════════════════════════╝
//
// The protoreflect.Enum and proto.Message types are the concrete Go type
// associated with the named enum or message. Repeated fields are represented
// using a Go slice of the base element type.
//
// If a generated extension descriptor variable is directly passed to
// GetExtension, then the call should be followed immediately by a
// type assertion to the expected output value. For example:
//
//	mm := proto.GetExtension(m, foopb.E_MyExtension).(*foopb.MyMessage)
//
// This pattern enables static analysis tools to verify that the asserted type
// matches the Go type associated with the extension field and
// also enables a possible future migration to a type-safe extension API.
//
// Since singular messages are the most common extension type, the pattern of
// calling HasExtension followed by GetExtension may be simplified to:
//
//	if mm := proto.GetExtension(m, foopb.E_MyExtension).(*foopb.MyMessage); mm != nil {
//	    ... // make use of mm
//	}
//
// The mm variable is non-nil if and only if HasExtension reports true.
func GetExtension(m Message, xt protoreflect.ExtensionType) any {
	// Treat nil message interface as an empty message; return the default.
	if m == nil {
		return xt.InterfaceOf(xt.Zero())
	}

	return xt.InterfaceOf(m.ProtoReflect().Get(xt.TypeDescriptor()))
}

// SetExtension stores the value of an extension field.
// It panics if m is invalid, xt does not extend m, or if type of v
// is invalid for the specified extension field.
//
// The type of the value is dependent on the field type of the extension.
// For extensions generated by protoc-gen-go, the Go type is as follows:
//
//	╔═══════════════════╤═════════════════════════╗
//	║ Go type           │ Protobuf kind           ║
//	╠═══════════════════╪═════════════════════════╣
//	║ bool              │ bool                    ║
//	║ int32             │ int32, sint32, sfixed32 ║
//	║ int64             │ int64, sint64, sfixed64 ║
//	║ uint32            │ uint32, fixed32         ║
//	║ uint64            │ uint64, fixed64         ║
//	║ float32           │ float                   ║
//	║ float64           │ double                  ║
//	║ string            │ string                  ║
//	║ []byte            │ bytes                   ║
//	║ protoreflect.Enum │ enum                    ║
//	║ proto.Message     │ message, group          ║
//	╚═══════════════════╧═════════════════════════╝
//
// The protoreflect.Enum and proto.Message types are the concrete Go type
// associated with the named enum or message. Repeated fields are represented
// using a Go slice of the base element type.
//
// If a generated extension descriptor variable is directly passed to
// SetExtension (e.g., foopb.E_MyExtension), then the value should be a
// concrete type that matches the expected Go type for the extension descriptor
// so that static analysis tools can verify type correctness.
// This also enables a possible future migration to a type-safe extension API.
func SetExtension(m Message, xt protoreflect.ExtensionType, v any) {
	xd := xt.TypeDescriptor()
	pv := xt.ValueOf(v)

	// Specially treat an invalid list, map, or message as clear.
	isValid := true
	switch {
	case xd.IsList():
		isValid = pv.List().IsValid()
	case xd.IsMap():
		isValid = pv.Map().IsValid()
	case xd.Message() != nil:
		isValid = pv.Message().IsValid()
	}
	if !isValid {
		m.ProtoReflect().Clear(xd)
		return
	}

	m.ProtoReflect().Set(xd, pv)
}

// RangeExtensions iterates over every populated extension field in m in an
// undefined order, calling f for each extension type and value encountered.
// It returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current extension field.
func RangeExtensions(m Message, f func(protoreflect.ExtensionType, any) bool) {
	// Treat nil message interface as an empty message; nothing to range over.
	if m == nil {
		return
	}

	m.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		if fd.IsExtension() {
			xt := fd.(protoreflect.ExtensionTypeDescriptor).Type()
			vi := xt.InterfaceOf(v)
			return f(xt, vi)
		}
		return true
	})
}
