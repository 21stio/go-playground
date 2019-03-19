// Code generated by protoc-gen-go. DO NOT EDIT.
// source: whateversort/whateversort_.proto

/*
Package whateversort is a generated protocol buffer package.

It is generated from these files:
	whateversort/whateversort_.proto

It has these top-level messages:
	WhateverSort
*/
package whateversort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sortkind "github.com/21stio/go-playground/grpc/schema/sortkind"
import whateverunionsort "github.com/21stio/go-playground/grpc/schema/whateverunionsort"
import servicesort "github.com/21stio/go-playground/grpc/schema/servicesort"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WhateverSort struct {
	StringField      *sortkind.SortKind                   `protobuf:"varint,1,opt,name=stringField,enum=sortkind.SortKind" json:"stringField,omitempty"`
	IntField         *sortkind.SortKind                   `protobuf:"varint,2,opt,name=intField,enum=sortkind.SortKind" json:"intField,omitempty"`
	FloatField       *sortkind.SortKind                   `protobuf:"varint,3,opt,name=floatField,enum=sortkind.SortKind" json:"floatField,omitempty"`
	BoolField        *sortkind.SortKind                   `protobuf:"varint,4,opt,name=boolField,enum=sortkind.SortKind" json:"boolField,omitempty"`
	ObjectField      *WhateverSort                        `protobuf:"bytes,5,opt,name=objectField" json:"objectField,omitempty"`
	UnionField       *whateverunionsort.WhateverUnionSort `protobuf:"bytes,6,opt,name=unionField" json:"unionField,omitempty"`
	Meta             *servicesort.TypeMetaSort            `protobuf:"bytes,7,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *WhateverSort) Reset()                    { *m = WhateverSort{} }
func (m *WhateverSort) String() string            { return proto.CompactTextString(m) }
func (*WhateverSort) ProtoMessage()               {}
func (*WhateverSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WhateverSort) GetStringField() sortkind.SortKind {
	if m != nil && m.StringField != nil {
		return *m.StringField
	}
	return sortkind.SortKind_ASC
}

func (m *WhateverSort) GetIntField() sortkind.SortKind {
	if m != nil && m.IntField != nil {
		return *m.IntField
	}
	return sortkind.SortKind_ASC
}

func (m *WhateverSort) GetFloatField() sortkind.SortKind {
	if m != nil && m.FloatField != nil {
		return *m.FloatField
	}
	return sortkind.SortKind_ASC
}

func (m *WhateverSort) GetBoolField() sortkind.SortKind {
	if m != nil && m.BoolField != nil {
		return *m.BoolField
	}
	return sortkind.SortKind_ASC
}

func (m *WhateverSort) GetObjectField() *WhateverSort {
	if m != nil {
		return m.ObjectField
	}
	return nil
}

func (m *WhateverSort) GetUnionField() *whateverunionsort.WhateverUnionSort {
	if m != nil {
		return m.UnionField
	}
	return nil
}

func (m *WhateverSort) GetMeta() *servicesort.TypeMetaSort {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *WhateverSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*WhateverSort)(nil), "whateversort.WhateverSort")
}

func init() { proto.RegisterFile("whateversort/whateversort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0x80, 0xd9, 0xb7, 0x7d, 0xb5, 0x9d, 0x16, 0x0f, 0x39, 0xad, 0x3d, 0xc8, 0x22, 0x1e, 0x8a,
	0xd0, 0x44, 0x17, 0x0f, 0x22, 0x9e, 0x44, 0xbc, 0x88, 0x97, 0xaa, 0x08, 0x5e, 0x24, 0xdd, 0x8d,
	0xbb, 0xd1, 0x6d, 0x66, 0x49, 0xd2, 0x4a, 0x7f, 0x8e, 0xff, 0x54, 0x9a, 0xfd, 0x30, 0x45, 0xf6,
	0x12, 0x26, 0x93, 0xe7, 0x99, 0x24, 0x93, 0x40, 0xf4, 0x95, 0x73, 0x2b, 0xd6, 0x42, 0x1b, 0xd4,
	0x96, 0xf9, 0x93, 0x37, 0x5a, 0x6a, 0xb4, 0x48, 0xc6, 0x7e, 0x72, 0x12, 0x6e, 0xc7, 0x4f, 0xa9,
	0x52, 0xd6, 0x04, 0x35, 0x37, 0x39, 0x6d, 0xb8, 0x95, 0x92, 0xa8, 0x76, 0xca, 0xb5, 0x99, 0x86,
	0x3d, 0x32, 0x42, 0xaf, 0x65, 0x22, 0x1c, 0xe5, 0xc5, 0xf5, 0xfa, 0xf1, 0x77, 0x0f, 0xc6, 0x2f,
	0xb5, 0xfc, 0x88, 0xda, 0x92, 0x0b, 0x18, 0x19, 0xab, 0xa5, 0xca, 0xee, 0xa4, 0x28, 0xd2, 0x30,
	0x88, 0x82, 0xe9, 0x41, 0x4c, 0x68, 0x73, 0x06, 0xba, 0x85, 0xee, 0xa5, 0x4a, 0xe7, 0x3e, 0x46,
	0x28, 0x0c, 0xa4, 0xb2, 0x95, 0xf2, 0xaf, 0x53, 0x69, 0x19, 0x12, 0x03, 0xbc, 0x17, 0xc8, 0x6b,
	0xa3, 0xd7, 0x69, 0x78, 0x14, 0x39, 0x83, 0xe1, 0x02, 0xb1, 0xa8, 0x94, 0x7e, 0xa7, 0xf2, 0x0b,
	0x91, 0x6b, 0x18, 0xe1, 0xe2, 0x43, 0x24, 0xf5, 0x36, 0xff, 0xa3, 0x60, 0x3a, 0x8a, 0x27, 0xd4,
	0x6f, 0x33, 0xf5, 0x2f, 0x3f, 0xf7, 0x71, 0x72, 0x0b, 0xe0, 0xda, 0x59, 0xc9, 0x7b, 0x4e, 0x3e,
	0xa1, 0x7f, 0x3a, 0xdd, 0x56, 0x78, 0xde, 0x66, 0x5c, 0x19, 0xcf, 0x23, 0x33, 0xe8, 0x2f, 0x85,
	0xe5, 0xe1, 0xbe, 0xf3, 0x0f, 0xa9, 0xf7, 0x06, 0xf4, 0x69, 0x53, 0x8a, 0x07, 0x61, 0xb9, 0x93,
	0x1c, 0x46, 0x08, 0xf4, 0x73, 0x6e, 0xf2, 0x70, 0x10, 0x05, 0xd3, 0xe1, 0xdc, 0xc5, 0x37, 0x57,
	0xaf, 0x97, 0x99, 0xb4, 0xf9, 0x6a, 0x41, 0x13, 0x5c, 0xb2, 0xf8, 0xdc, 0x58, 0x89, 0x2c, 0xc3,
	0x59, 0x59, 0xf0, 0x4d, 0xa6, 0x71, 0xa5, 0x52, 0x96, 0xe9, 0x32, 0x61, 0x26, 0xc9, 0xc5, 0x92,
	0xef, 0x7c, 0xad, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x70, 0x7f, 0x22, 0x76, 0x02, 0x00,
	0x00,
}