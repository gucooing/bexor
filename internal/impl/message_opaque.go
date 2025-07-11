// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync/atomic"

	"github.com/gucooing/bexor/internal/filedesc"
	"github.com/gucooing/bexor/reflect/protoreflect"
)

type opaqueStructInfo struct {
	structInfo
}

// isOpaque determines whether a protobuf message type is on the Opaque API.  It
// checks whether the type is a Go struct that protoc-gen-go would generate.
//
// This function only detects newly generated messages from the v2
// implementation of protoc-gen-go. It is unable to classify generated messages
// that are too old or those that are generated by a different generator
// such as protoc-gen-gogo.
func isOpaque(t reflect.Type) bool {
	// The current detection mechanism is to simply check the first field
	// for a struct tag with the "protogen" key.
	if t.Kind() == reflect.Struct && t.NumField() > 0 {
		pgt := t.Field(0).Tag.Get("protogen")
		return strings.HasPrefix(pgt, "opaque.")
	}
	return false
}

func opaqueInitHook(mi *MessageInfo) bool {
	mt := mi.GoReflectType.Elem()
	si := opaqueStructInfo{
		structInfo: mi.makeStructInfo(mt),
	}

	if !isOpaque(mt) {
		return false
	}

	defer atomic.StoreUint32(&mi.initDone, 1)

	mi.fields = map[protoreflect.FieldNumber]*fieldInfo{}
	fds := mi.Desc.Fields()
	for i := 0; i < fds.Len(); i++ {
		fd := fds.Get(i)
		fs := si.fieldsByNumber[fd.Number()]
		var fi fieldInfo
		usePresence, _ := filedesc.UsePresenceForField(fd)

		switch {
		case fd.ContainingOneof() != nil && !fd.ContainingOneof().IsSynthetic():
			// Oneofs are no different for opaque.
			fi = fieldInfoForOneof(fd, si.oneofsByName[fd.ContainingOneof().Name()], mi.Exporter, si.oneofWrappersByNumber[fd.Number()])
		case fd.IsMap():
			fi = mi.fieldInfoForMapOpaque(si, fd, fs)
		case fd.IsList() && fd.Message() == nil && usePresence:
			fi = mi.fieldInfoForScalarListOpaque(si, fd, fs)
		case fd.IsList() && fd.Message() == nil:
			// Proto3 lists without presence can use same access methods as open
			fi = fieldInfoForList(fd, fs, mi.Exporter)
		case fd.IsList() && usePresence:
			fi = mi.fieldInfoForMessageListOpaque(si, fd, fs)
		case fd.IsList():
			// Proto3 opaque messages that does not need presence bitmap.
			// Different representation than open struct, but same logic
			fi = mi.fieldInfoForMessageListOpaqueNoPresence(si, fd, fs)
		case fd.Message() != nil && usePresence:
			fi = mi.fieldInfoForMessageOpaque(si, fd, fs)
		case fd.Message() != nil:
			// Proto3 messages without presence can use same access methods as open
			fi = fieldInfoForMessage(fd, fs, mi.Exporter)
		default:
			fi = mi.fieldInfoForScalarOpaque(si, fd, fs)
		}
		mi.fields[fd.Number()] = &fi
	}
	mi.oneofs = map[protoreflect.Name]*oneofInfo{}
	for i := 0; i < mi.Desc.Oneofs().Len(); i++ {
		od := mi.Desc.Oneofs().Get(i)
		mi.oneofs[od.Name()] = makeOneofInfoOpaque(mi, od, si.structInfo, mi.Exporter)
	}

	mi.denseFields = make([]*fieldInfo, fds.Len()*2)
	for i := 0; i < fds.Len(); i++ {
		if fd := fds.Get(i); int(fd.Number()) < len(mi.denseFields) {
			mi.denseFields[fd.Number()] = mi.fields[fd.Number()]
		}
	}

	for i := 0; i < fds.Len(); {
		fd := fds.Get(i)
		if od := fd.ContainingOneof(); od != nil && !fd.ContainingOneof().IsSynthetic() {
			mi.rangeInfos = append(mi.rangeInfos, mi.oneofs[od.Name()])
			i += od.Fields().Len()
		} else {
			mi.rangeInfos = append(mi.rangeInfos, mi.fields[fd.Number()])
			i++
		}
	}

	mi.makeExtensionFieldsFunc(mt, si.structInfo)
	mi.makeUnknownFieldsFunc(mt, si.structInfo)
	mi.makeOpaqueCoderMethods(mt, si)
	mi.makeFieldTypes(si.structInfo)

	return true
}

func makeOneofInfoOpaque(mi *MessageInfo, od protoreflect.OneofDescriptor, si structInfo, x exporter) *oneofInfo {
	oi := &oneofInfo{oneofDesc: od}
	if od.IsSynthetic() {
		fd := od.Fields().Get(0)
		index, _ := presenceIndex(mi.Desc, fd)
		oi.which = func(p pointer) protoreflect.FieldNumber {
			if p.IsNil() {
				return 0
			}
			if !mi.present(p, index) {
				return 0
			}
			return od.Fields().Get(0).Number()
		}
		return oi
	}
	// Dispatch to non-opaque oneof implementation for non-synthetic oneofs.
	return makeOneofInfo(od, si, x)
}

func (mi *MessageInfo) fieldInfoForMapOpaque(si opaqueStructInfo, fd protoreflect.FieldDescriptor, fs reflect.StructField) fieldInfo {
	ft := fs.Type
	if ft.Kind() != reflect.Map {
		panic(fmt.Sprintf("invalid type: got %v, want map kind", ft))
	}
	fieldOffset := offsetOf(fs)
	conv := NewConverter(ft, fd)
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			// Don't bother checking presence bits, since we need to
			// look at the map length even if the presence bit is set.
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			return rv.Len() > 0
		},
		clear: func(p pointer) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(reflect.Zero(rv.Type()))
		},
		get: func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.Len() == 0 {
				return conv.Zero()
			}
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v protoreflect.Value) {
			pv := conv.GoValueOf(v)
			if pv.IsNil() {
				panic(fmt.Sprintf("invalid value: setting map field to read-only value"))
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(pv)
		},
		mutable: func(p pointer) protoreflect.Value {
			v := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if v.IsNil() {
				v.Set(reflect.MakeMap(fs.Type))
			}
			return conv.PBValueOf(v)
		},
		newField: func() protoreflect.Value {
			return conv.New()
		},
	}
}

func (mi *MessageInfo) fieldInfoForScalarListOpaque(si opaqueStructInfo, fd protoreflect.FieldDescriptor, fs reflect.StructField) fieldInfo {
	ft := fs.Type
	if ft.Kind() != reflect.Slice {
		panic(fmt.Sprintf("invalid type: got %v, want slice kind", ft))
	}
	conv := NewConverter(reflect.PtrTo(ft), fd)
	fieldOffset := offsetOf(fs)
	index, _ := presenceIndex(mi.Desc, fd)
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			return rv.Len() > 0
		},
		clear: func(p pointer) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(reflect.Zero(rv.Type()))
		},
		get: func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type)
			if rv.Elem().Len() == 0 {
				return conv.Zero()
			}
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v protoreflect.Value) {
			pv := conv.GoValueOf(v)
			if pv.IsNil() {
				panic(fmt.Sprintf("invalid value: setting repeated field to read-only value"))
			}
			mi.setPresent(p, index)
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(pv.Elem())
		},
		mutable: func(p pointer) protoreflect.Value {
			mi.setPresent(p, index)
			return conv.PBValueOf(p.Apply(fieldOffset).AsValueOf(fs.Type))
		},
		newField: func() protoreflect.Value {
			return conv.New()
		},
	}
}

func (mi *MessageInfo) fieldInfoForMessageListOpaque(si opaqueStructInfo, fd protoreflect.FieldDescriptor, fs reflect.StructField) fieldInfo {
	ft := fs.Type
	if ft.Kind() != reflect.Ptr || ft.Elem().Kind() != reflect.Slice {
		panic(fmt.Sprintf("invalid type: got %v, want slice kind", ft))
	}
	conv := NewConverter(ft, fd)
	fieldOffset := offsetOf(fs)
	index, _ := presenceIndex(mi.Desc, fd)
	fieldNumber := fd.Number()
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			if !mi.present(p, index) {
				return false
			}
			sp := p.Apply(fieldOffset).AtomicGetPointer()
			if sp.IsNil() {
				// Lazily unmarshal this field.
				mi.lazyUnmarshal(p, fieldNumber)
				sp = p.Apply(fieldOffset).AtomicGetPointer()
			}
			rv := sp.AsValueOf(fs.Type.Elem())
			return rv.Elem().Len() > 0
		},
		clear: func(p pointer) {
			fp := p.Apply(fieldOffset)
			sp := fp.AtomicGetPointer()
			if sp.IsNil() {
				sp = fp.AtomicSetPointerIfNil(pointerOfValue(reflect.New(fs.Type.Elem())))
				mi.setPresent(p, index)
			}
			rv := sp.AsValueOf(fs.Type.Elem())
			rv.Elem().Set(reflect.Zero(rv.Type().Elem()))
		},
		get: func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			if !mi.present(p, index) {
				return conv.Zero()
			}
			sp := p.Apply(fieldOffset).AtomicGetPointer()
			if sp.IsNil() {
				// Lazily unmarshal this field.
				mi.lazyUnmarshal(p, fieldNumber)
				sp = p.Apply(fieldOffset).AtomicGetPointer()
			}
			rv := sp.AsValueOf(fs.Type.Elem())
			if rv.Elem().Len() == 0 {
				return conv.Zero()
			}
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v protoreflect.Value) {
			fp := p.Apply(fieldOffset)
			sp := fp.AtomicGetPointer()
			if sp.IsNil() {
				sp = fp.AtomicSetPointerIfNil(pointerOfValue(reflect.New(fs.Type.Elem())))
				mi.setPresent(p, index)
			}
			rv := sp.AsValueOf(fs.Type.Elem())
			val := conv.GoValueOf(v)
			if val.IsNil() {
				panic(fmt.Sprintf("invalid value: setting repeated field to read-only value"))
			} else {
				rv.Elem().Set(val.Elem())
			}
		},
		mutable: func(p pointer) protoreflect.Value {
			fp := p.Apply(fieldOffset)
			sp := fp.AtomicGetPointer()
			if sp.IsNil() {
				if mi.present(p, index) {
					// Lazily unmarshal this field.
					mi.lazyUnmarshal(p, fieldNumber)
					sp = p.Apply(fieldOffset).AtomicGetPointer()
				} else {
					sp = fp.AtomicSetPointerIfNil(pointerOfValue(reflect.New(fs.Type.Elem())))
					mi.setPresent(p, index)
				}
			}
			rv := sp.AsValueOf(fs.Type.Elem())
			return conv.PBValueOf(rv)
		},
		newField: func() protoreflect.Value {
			return conv.New()
		},
	}
}

func (mi *MessageInfo) fieldInfoForMessageListOpaqueNoPresence(si opaqueStructInfo, fd protoreflect.FieldDescriptor, fs reflect.StructField) fieldInfo {
	ft := fs.Type
	if ft.Kind() != reflect.Ptr || ft.Elem().Kind() != reflect.Slice {
		panic(fmt.Sprintf("invalid type: got %v, want slice kind", ft))
	}
	conv := NewConverter(ft, fd)
	fieldOffset := offsetOf(fs)
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() {
				return false
			}
			return rv.Elem().Len() > 0
		},
		clear: func(p pointer) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if !rv.IsNil() {
				rv.Elem().Set(reflect.Zero(rv.Type().Elem()))
			}
		},
		get: func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() {
				return conv.Zero()
			}
			if rv.Elem().Len() == 0 {
				return conv.Zero()
			}
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v protoreflect.Value) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() {
				rv.Set(reflect.New(fs.Type.Elem()))
			}
			val := conv.GoValueOf(v)
			if val.IsNil() {
				panic(fmt.Sprintf("invalid value: setting repeated field to read-only value"))
			} else {
				rv.Elem().Set(val.Elem())
			}
		},
		mutable: func(p pointer) protoreflect.Value {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if rv.IsNil() {
				rv.Set(reflect.New(fs.Type.Elem()))
			}
			return conv.PBValueOf(rv)
		},
		newField: func() protoreflect.Value {
			return conv.New()
		},
	}
}

func (mi *MessageInfo) fieldInfoForScalarOpaque(si opaqueStructInfo, fd protoreflect.FieldDescriptor, fs reflect.StructField) fieldInfo {
	ft := fs.Type
	nullable := fd.HasPresence()
	if oneof := fd.ContainingOneof(); oneof != nil && oneof.IsSynthetic() {
		nullable = true
	}
	deref := false
	if nullable && ft.Kind() == reflect.Ptr {
		ft = ft.Elem()
		deref = true
	}
	conv := NewConverter(ft, fd)
	fieldOffset := offsetOf(fs)
	index, _ := presenceIndex(mi.Desc, fd)
	var getter func(p pointer) protoreflect.Value
	if !nullable {
		getter = getterForDirectScalar(fd, fs, conv, fieldOffset)
	} else {
		getter = getterForOpaqueNullableScalar(mi, index, fd, fs, conv, fieldOffset)
	}
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			if nullable {
				return mi.present(p, index)
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			switch rv.Kind() {
			case reflect.Bool:
				return rv.Bool()
			case reflect.Int32, reflect.Int64:
				return rv.Int() != 0
			case reflect.Uint32, reflect.Uint64:
				return rv.Uint() != 0
			case reflect.Float32, reflect.Float64:
				return rv.Float() != 0 || math.Signbit(rv.Float())
			case reflect.String, reflect.Slice:
				return rv.Len() > 0
			default:
				panic(fmt.Sprintf("invalid type: %v", rv.Type())) // should never happen
			}
		},
		clear: func(p pointer) {
			if nullable {
				mi.clearPresent(p, index)
			}
			// This is only valuable for bytes and strings, but we do it unconditionally.
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			rv.Set(reflect.Zero(rv.Type()))
		},
		get: getter,
		// TODO: Implement unsafe fast path for set?
		set: func(p pointer, v protoreflect.Value) {
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			if deref {
				if rv.IsNil() {
					rv.Set(reflect.New(ft))
				}
				rv = rv.Elem()
			}

			rv.Set(conv.GoValueOf(v))
			if nullable && rv.Kind() == reflect.Slice && rv.IsNil() {
				rv.Set(emptyBytes)
			}
			if nullable {
				mi.setPresent(p, index)
			}
		},
		newField: func() protoreflect.Value {
			return conv.New()
		},
	}
}

func (mi *MessageInfo) fieldInfoForMessageOpaque(si opaqueStructInfo, fd protoreflect.FieldDescriptor, fs reflect.StructField) fieldInfo {
	ft := fs.Type
	conv := NewConverter(ft, fd)
	fieldOffset := offsetOf(fs)
	index, _ := presenceIndex(mi.Desc, fd)
	fieldNumber := fd.Number()
	elemType := fs.Type.Elem()
	return fieldInfo{
		fieldDesc: fd,
		has: func(p pointer) bool {
			if p.IsNil() {
				return false
			}
			return mi.present(p, index)
		},
		clear: func(p pointer) {
			mi.clearPresent(p, index)
			p.Apply(fieldOffset).AtomicSetNilPointer()
		},
		get: func(p pointer) protoreflect.Value {
			if p.IsNil() || !mi.present(p, index) {
				return conv.Zero()
			}
			fp := p.Apply(fieldOffset)
			mp := fp.AtomicGetPointer()
			if mp.IsNil() {
				// Lazily unmarshal this field.
				mi.lazyUnmarshal(p, fieldNumber)
				mp = fp.AtomicGetPointer()
			}
			rv := mp.AsValueOf(elemType)
			return conv.PBValueOf(rv)
		},
		set: func(p pointer, v protoreflect.Value) {
			val := pointerOfValue(conv.GoValueOf(v))
			if val.IsNil() {
				panic("invalid nil pointer")
			}
			p.Apply(fieldOffset).AtomicSetPointer(val)
			mi.setPresent(p, index)
		},
		mutable: func(p pointer) protoreflect.Value {
			fp := p.Apply(fieldOffset)
			mp := fp.AtomicGetPointer()
			if mp.IsNil() {
				if mi.present(p, index) {
					// Lazily unmarshal this field.
					mi.lazyUnmarshal(p, fieldNumber)
					mp = fp.AtomicGetPointer()
				} else {
					mp = pointerOfValue(conv.GoValueOf(conv.New()))
					fp.AtomicSetPointer(mp)
					mi.setPresent(p, index)
				}
			}
			return conv.PBValueOf(mp.AsValueOf(fs.Type.Elem()))
		},
		newMessage: func() protoreflect.Message {
			return conv.New().Message()
		},
		newField: func() protoreflect.Value {
			return conv.New()
		},
	}
}

// A presenceList wraps a List, updating presence bits as necessary when the
// list contents change.
type presenceList struct {
	pvalueList
	setPresence func(bool)
}
type pvalueList interface {
	protoreflect.List
	//Unwrapper
}

func (list presenceList) Append(v protoreflect.Value) {
	list.pvalueList.Append(v)
	list.setPresence(true)
}
func (list presenceList) Truncate(i int) {
	list.pvalueList.Truncate(i)
	list.setPresence(i > 0)
}

// presenceIndex returns the index to pass to presence functions.
//
// TODO: field.Desc.Index() would be simpler, and would give space to record the presence of oneof fields.
func presenceIndex(md protoreflect.MessageDescriptor, fd protoreflect.FieldDescriptor) (uint32, presenceSize) {
	found := false
	var index, numIndices uint32
	for i := 0; i < md.Fields().Len(); i++ {
		f := md.Fields().Get(i)
		if f == fd {
			found = true
			index = numIndices
		}
		if f.ContainingOneof() == nil || isLastOneofField(f) {
			numIndices++
		}
	}
	if !found {
		panic(fmt.Sprintf("BUG: %v not in %v", fd.Name(), md.FullName()))
	}
	return index, presenceSize(numIndices)
}

func isLastOneofField(fd protoreflect.FieldDescriptor) bool {
	fields := fd.ContainingOneof().Fields()
	return fields.Get(fields.Len()-1) == fd
}

func (mi *MessageInfo) setPresent(p pointer, index uint32) {
	p.Apply(mi.presenceOffset).PresenceInfo().SetPresent(index, mi.presenceSize)
}

func (mi *MessageInfo) clearPresent(p pointer, index uint32) {
	p.Apply(mi.presenceOffset).PresenceInfo().ClearPresent(index)
}

func (mi *MessageInfo) present(p pointer, index uint32) bool {
	return p.Apply(mi.presenceOffset).PresenceInfo().Present(index)
}
