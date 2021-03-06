// Code generated by protoc-gen-go. DO NOT EDIT.
// source: whateverunionsort/whateverunionsort_.proto

/*
Package whateverunionsort is a generated protocol buffer package.

It is generated from these files:
	whateverunionsort/whateverunionsort_.proto

It has these top-level messages:
	WhateverUnionSort
*/
package whateverunionsort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sortkind "github.com/21stio/go-playground/grpc/schema/sortkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WhateverUnionSort struct {
	StringField      *sortkind.SortKind `protobuf:"varint,1,opt,name=stringField,enum=sortkind.SortKind" json:"stringField,omitempty"`
	IntField         *sortkind.SortKind `protobuf:"varint,2,opt,name=intField,enum=sortkind.SortKind" json:"intField,omitempty"`
	FloatField       *sortkind.SortKind `protobuf:"varint,3,opt,name=floatField,enum=sortkind.SortKind" json:"floatField,omitempty"`
	BoolField        *sortkind.SortKind `protobuf:"varint,4,opt,name=boolField,enum=sortkind.SortKind" json:"boolField,omitempty"`
	Hash             *string            `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *WhateverUnionSort) Reset()                    { *m = WhateverUnionSort{} }
func (m *WhateverUnionSort) String() string            { return proto.CompactTextString(m) }
func (*WhateverUnionSort) ProtoMessage()               {}
func (*WhateverUnionSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WhateverUnionSort) GetStringField() sortkind.SortKind {
	if m != nil && m.StringField != nil {
		return *m.StringField
	}
	return sortkind.SortKind_ASC
}

func (m *WhateverUnionSort) GetIntField() sortkind.SortKind {
	if m != nil && m.IntField != nil {
		return *m.IntField
	}
	return sortkind.SortKind_ASC
}

func (m *WhateverUnionSort) GetFloatField() sortkind.SortKind {
	if m != nil && m.FloatField != nil {
		return *m.FloatField
	}
	return sortkind.SortKind_ASC
}

func (m *WhateverUnionSort) GetBoolField() sortkind.SortKind {
	if m != nil && m.BoolField != nil {
		return *m.BoolField
	}
	return sortkind.SortKind_ASC
}

func (m *WhateverUnionSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*WhateverUnionSort)(nil), "whateverunionsort.WhateverUnionSort")
}

func init() { proto.RegisterFile("whateverunionsort/whateverunionsort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x71, 0xaa, 0x13, 0xdc, 0x13, 0x84, 0xbd, 0x53, 0xf1, 0x34, 0x3c, 0x0d, 0xc1, 0x44,
	0x8b, 0x57, 0x11, 0x3c, 0x78, 0xf1, 0x36, 0x11, 0xc1, 0x8b, 0x64, 0x6d, 0x4d, 0x1e, 0x66, 0x79,
	0x25, 0x79, 0x55, 0xfc, 0xbb, 0xfd, 0x07, 0xa4, 0xae, 0x73, 0x83, 0xd2, 0x4b, 0xf8, 0x91, 0x7c,
	0xbe, 0x97, 0xc0, 0xc5, 0x97, 0x33, 0x52, 0x7f, 0xd6, 0xb1, 0x0d, 0xc4, 0x21, 0x71, 0x14, 0x3d,
	0xb8, 0x79, 0x53, 0x4d, 0x64, 0x61, 0x9c, 0x0d, 0x5e, 0xce, 0xf2, 0xee, 0xfc, 0xa0, 0x50, 0xe9,
	0xed, 0xe8, 0xf1, 0xf9, 0x4f, 0x06, 0xb3, 0x97, 0xde, 0x3f, 0x77, 0xfe, 0x89, 0xa3, 0xe0, 0x0d,
	0x9c, 0x24, 0x89, 0x14, 0xec, 0x03, 0xd5, 0xbe, 0xca, 0xb3, 0x79, 0xb6, 0x38, 0x2d, 0x50, 0x6d,
	0x63, 0xd5, 0xa1, 0x47, 0x0a, 0xd5, 0x72, 0x9f, 0xa1, 0x82, 0x63, 0x0a, 0xb2, 0x49, 0x0e, 0x46,
	0x93, 0x7f, 0x83, 0x05, 0xc0, 0xbb, 0x67, 0xd3, 0x17, 0x87, 0xa3, 0xc5, 0x9e, 0xc2, 0x2b, 0x98,
	0xae, 0x98, 0xfd, 0x26, 0x99, 0x8c, 0x26, 0x3b, 0x84, 0x08, 0x13, 0x67, 0x92, 0xcb, 0x8f, 0xe6,
	0xd9, 0x62, 0xba, 0xfc, 0xdb, 0xf7, 0x77, 0xaf, 0xb7, 0x96, 0xc4, 0xb5, 0x2b, 0x55, 0xf2, 0x5a,
	0x17, 0xd7, 0x49, 0x88, 0xb5, 0xe5, 0xcb, 0xc6, 0x9b, 0x6f, 0x1b, 0xb9, 0x0d, 0x95, 0xb6, 0xb1,
	0x29, 0x75, 0x2a, 0x5d, 0xbd, 0x36, 0xc3, 0xaf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xe8,
	0x34, 0x04, 0x90, 0x01, 0x00, 0x00,
}
