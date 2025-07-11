// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protoimpl

import (
	"github.com/gucooing/bexor/internal/version"
)

const (
	// MaxVersion is the maximum supported version for generated .pb.go files.
	// It is always the current version of the module.
	MaxVersion = version.Minor

	// GenVersion is the runtime version required by generated .pb.go files.
	// This is incremented when generated code relies on new functionality
	// in the runtime.
	GenVersion = 20

	// MinVersion is the minimum supported version for generated .pb.go files.
	// This is incremented when the runtime drops support for old code.
	MinVersion = 0
)

// EnforceVersion is used by code generated by protoc-gen-go
// to statically enforce minimum and maximum versions of this package.
// A compilation failure implies either that:
//   - the runtime package is too old and needs to be updated OR
//   - the generated code is too old and needs to be regenerated.
//
// The runtime package can be upgraded by running:
//
//	go get github.com/gucooing/bexor
//
// The generated code can be regenerated by running:
//
//	protoc --go_out=${PROTOC_GEN_GO_ARGS} ${PROTO_FILES}
//
// Example usage by generated code:
//
//	const (
//		// Verify that this generated code is sufficiently up-to-date.
//		_ = protoimpl.EnforceVersion(genVersion - protoimpl.MinVersion)
//		// Verify that runtime/protoimpl is sufficiently up-to-date.
//		_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - genVersion)
//	)
//
// The genVersion is the current minor version used to generated the code.
// This compile-time check relies on negative integer overflow of a uint
// being a compilation failure (guaranteed by the Go specification).
type EnforceVersion uint

// This enforces the following invariant:
//
//	MinVersion ≤ GenVersion ≤ MaxVersion
const (
	_ = EnforceVersion(GenVersion - MinVersion)
	_ = EnforceVersion(MaxVersion - GenVersion)
)
