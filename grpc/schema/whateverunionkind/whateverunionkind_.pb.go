// Code generated by protoc-gen-go. DO NOT EDIT.
// source: whateverunionkind/whateverunionkind_.proto

/*
Package whateverunionkind is a generated protocol buffer package.

It is generated from these files:
	whateverunionkind/whateverunionkind_.proto

It has these top-level messages:
*/
package whateverunionkind

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WhateverUnionKind int32

const (
	WhateverUnionKind_WHATEVER_UNION_STRING_FIELD WhateverUnionKind = 0
	WhateverUnionKind_WHATEVER_UNION_INT_FIELD    WhateverUnionKind = 1
	WhateverUnionKind_WHATEVER_UNION_FLOAT_FIELD  WhateverUnionKind = 2
	WhateverUnionKind_WHATEVER_UNION_BOOL_FIELD   WhateverUnionKind = 3
)

var WhateverUnionKind_name = map[int32]string{
	0: "WHATEVER_UNION_STRING_FIELD",
	1: "WHATEVER_UNION_INT_FIELD",
	2: "WHATEVER_UNION_FLOAT_FIELD",
	3: "WHATEVER_UNION_BOOL_FIELD",
}
var WhateverUnionKind_value = map[string]int32{
	"WHATEVER_UNION_STRING_FIELD": 0,
	"WHATEVER_UNION_INT_FIELD":    1,
	"WHATEVER_UNION_FLOAT_FIELD":  2,
	"WHATEVER_UNION_BOOL_FIELD":   3,
}

func (x WhateverUnionKind) Enum() *WhateverUnionKind {
	p := new(WhateverUnionKind)
	*p = x
	return p
}
func (x WhateverUnionKind) String() string {
	return proto.EnumName(WhateverUnionKind_name, int32(x))
}
func (x *WhateverUnionKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WhateverUnionKind_value, data, "WhateverUnionKind")
	if err != nil {
		return err
	}
	*x = WhateverUnionKind(value)
	return nil
}
func (WhateverUnionKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("whateverunionkind.WhateverUnionKind", WhateverUnionKind_name, WhateverUnionKind_value)
}

func init() { proto.RegisterFile("whateverunionkind/whateverunionkind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2a, 0xcf, 0x48, 0x2c,
	0x49, 0x2d, 0x4b, 0x2d, 0x2a, 0xcd, 0xcb, 0xcc, 0xcf, 0xcb, 0xce, 0xcc, 0x4b, 0xd1, 0xc7, 0x10,
	0x89, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc4, 0x90, 0xd1, 0x9a, 0xc8, 0xc8, 0x25,
	0x18, 0x0e, 0x15, 0x0d, 0x05, 0x89, 0x7a, 0x67, 0xe6, 0xa5, 0x08, 0xc9, 0x73, 0x49, 0x87, 0x7b,
	0x38, 0x86, 0xb8, 0x86, 0xb9, 0x06, 0xc5, 0x87, 0xfa, 0x79, 0xfa, 0xfb, 0xc5, 0x07, 0x87, 0x04,
	0x79, 0xfa, 0xb9, 0xc7, 0xbb, 0x79, 0xba, 0xfa, 0xb8, 0x08, 0x30, 0x08, 0xc9, 0x70, 0x49, 0xa0,
	0x29, 0xf0, 0xf4, 0x0b, 0x81, 0xca, 0x32, 0x0a, 0xc9, 0x71, 0x49, 0xa1, 0xc9, 0xba, 0xf9, 0xf8,
	0x3b, 0xc2, 0xe4, 0x99, 0x84, 0x64, 0xb9, 0x24, 0xd1, 0xe4, 0x9d, 0xfc, 0xfd, 0x7d, 0xa0, 0xd2,
	0xcc, 0x4e, 0xf6, 0x51, 0xb6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x46, 0x86, 0xc5, 0x25, 0x99, 0xf9, 0xfa, 0xe9, 0xf9, 0xba, 0x05, 0x39, 0x89, 0x95, 0xe9, 0x45,
	0xf9, 0xa5, 0x79, 0x29, 0xfa, 0xe9, 0x45, 0x05, 0xc9, 0xfa, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89,
	0x98, 0xde, 0x05, 0x04, 0x00, 0x00, 0xff, 0xff, 0x42, 0xc2, 0x82, 0xe8, 0x14, 0x01, 0x00, 0x00,
}
