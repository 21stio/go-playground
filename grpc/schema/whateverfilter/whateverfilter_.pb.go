// Code generated by protoc-gen-go. DO NOT EDIT.
// source: whateverfilter/whateverfilter_.proto

/*
Package whateverfilter is a generated protocol buffer package.

It is generated from these files:
	whateverfilter/whateverfilter_.proto

It has these top-level messages:
	WhateverFilter
*/
package whateverfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import whateverkindfilter "github.com/21stio/go-playground/grpc/schema/whateverkindfilter"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"
import intfilter "github.com/21stio/go-playground/grpc/schema/intfilter"
import floatfilter "github.com/21stio/go-playground/grpc/schema/floatfilter"
import boolfilter "github.com/21stio/go-playground/grpc/schema/boolfilter"
import whateverunionfilter "github.com/21stio/go-playground/grpc/schema/whateverunionfilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WhateverFilter struct {
	EnumField        *whateverkindfilter.WhateverKindFilter   `protobuf:"bytes,1,opt,name=enumField" json:"enumField,omitempty"`
	StringField      *stringfilter.StringFilter               `protobuf:"bytes,2,opt,name=stringField" json:"stringField,omitempty"`
	IntField         *intfilter.IntFilter                     `protobuf:"bytes,3,opt,name=intField" json:"intField,omitempty"`
	FloatField       *floatfilter.FloatFilter                 `protobuf:"bytes,4,opt,name=floatField" json:"floatField,omitempty"`
	BoolField        *boolfilter.BoolFilter                   `protobuf:"bytes,5,opt,name=boolField" json:"boolField,omitempty"`
	ObjectField      *WhateverFilter                          `protobuf:"bytes,6,opt,name=objectField" json:"objectField,omitempty"`
	UnionField       *whateverunionfilter.WhateverUnionFilter `protobuf:"bytes,7,opt,name=unionField" json:"unionField,omitempty"`
	EnumListSome     *whateverkindfilter.WhateverKindFilter   `protobuf:"bytes,8,opt,name=enumListSome" json:"enumListSome,omitempty"`
	EnumListEvery    *whateverkindfilter.WhateverKindFilter   `protobuf:"bytes,9,opt,name=enumListEvery" json:"enumListEvery,omitempty"`
	EnumListNone     *whateverkindfilter.WhateverKindFilter   `protobuf:"bytes,10,opt,name=enumListNone" json:"enumListNone,omitempty"`
	StringListSome   *stringfilter.StringFilter               `protobuf:"bytes,11,opt,name=stringListSome" json:"stringListSome,omitempty"`
	StringListEvery  *stringfilter.StringFilter               `protobuf:"bytes,12,opt,name=stringListEvery" json:"stringListEvery,omitempty"`
	StringListNone   *stringfilter.StringFilter               `protobuf:"bytes,13,opt,name=stringListNone" json:"stringListNone,omitempty"`
	IntListSome      *intfilter.IntFilter                     `protobuf:"bytes,14,opt,name=intListSome" json:"intListSome,omitempty"`
	IntListEvery     *intfilter.IntFilter                     `protobuf:"bytes,15,opt,name=intListEvery" json:"intListEvery,omitempty"`
	IntListNone      *intfilter.IntFilter                     `protobuf:"bytes,16,opt,name=intListNone" json:"intListNone,omitempty"`
	FloatListSome    *floatfilter.FloatFilter                 `protobuf:"bytes,17,opt,name=floatListSome" json:"floatListSome,omitempty"`
	FloatListEvery   *floatfilter.FloatFilter                 `protobuf:"bytes,18,opt,name=floatListEvery" json:"floatListEvery,omitempty"`
	FloatListNone    *floatfilter.FloatFilter                 `protobuf:"bytes,19,opt,name=floatListNone" json:"floatListNone,omitempty"`
	BoolListSome     *boolfilter.BoolFilter                   `protobuf:"bytes,20,opt,name=boolListSome" json:"boolListSome,omitempty"`
	BoolListEvery    *boolfilter.BoolFilter                   `protobuf:"bytes,21,opt,name=boolListEvery" json:"boolListEvery,omitempty"`
	BoolListNone     *boolfilter.BoolFilter                   `protobuf:"bytes,22,opt,name=boolListNone" json:"boolListNone,omitempty"`
	ObjectListSome   *WhateverFilter                          `protobuf:"bytes,23,opt,name=objectListSome" json:"objectListSome,omitempty"`
	ObjectListEvery  *WhateverFilter                          `protobuf:"bytes,24,opt,name=objectListEvery" json:"objectListEvery,omitempty"`
	ObjectListNone   *WhateverFilter                          `protobuf:"bytes,25,opt,name=objectListNone" json:"objectListNone,omitempty"`
	UnionListSome    *whateverunionfilter.WhateverUnionFilter `protobuf:"bytes,26,opt,name=unionListSome" json:"unionListSome,omitempty"`
	UnionListEvery   *whateverunionfilter.WhateverUnionFilter `protobuf:"bytes,27,opt,name=unionListEvery" json:"unionListEvery,omitempty"`
	UnionListNone    *whateverunionfilter.WhateverUnionFilter `protobuf:"bytes,28,opt,name=unionListNone" json:"unionListNone,omitempty"`
	IdsSome          *idfilter.IdFilter                       `protobuf:"bytes,29,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery         *idfilter.IdFilter                       `protobuf:"bytes,30,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone          *idfilter.IdFilter                       `protobuf:"bytes,31,opt,name=idsNone" json:"idsNone,omitempty"`
	Meta             *servicefilter.TypeMetaFilter            `protobuf:"bytes,32,opt,name=meta" json:"meta,omitempty"`
	And              []*WhateverFilter                        `protobuf:"bytes,33,rep,name=and" json:"and,omitempty"`
	Or               []*WhateverFilter                        `protobuf:"bytes,34,rep,name=or" json:"or,omitempty"`
	Not              []*WhateverFilter                        `protobuf:"bytes,35,rep,name=not" json:"not,omitempty"`
	Hash             *string                                  `protobuf:"bytes,36,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *WhateverFilter) Reset()                    { *m = WhateverFilter{} }
func (m *WhateverFilter) String() string            { return proto.CompactTextString(m) }
func (*WhateverFilter) ProtoMessage()               {}
func (*WhateverFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WhateverFilter) GetEnumField() *whateverkindfilter.WhateverKindFilter {
	if m != nil {
		return m.EnumField
	}
	return nil
}

func (m *WhateverFilter) GetStringField() *stringfilter.StringFilter {
	if m != nil {
		return m.StringField
	}
	return nil
}

func (m *WhateverFilter) GetIntField() *intfilter.IntFilter {
	if m != nil {
		return m.IntField
	}
	return nil
}

func (m *WhateverFilter) GetFloatField() *floatfilter.FloatFilter {
	if m != nil {
		return m.FloatField
	}
	return nil
}

func (m *WhateverFilter) GetBoolField() *boolfilter.BoolFilter {
	if m != nil {
		return m.BoolField
	}
	return nil
}

func (m *WhateverFilter) GetObjectField() *WhateverFilter {
	if m != nil {
		return m.ObjectField
	}
	return nil
}

func (m *WhateverFilter) GetUnionField() *whateverunionfilter.WhateverUnionFilter {
	if m != nil {
		return m.UnionField
	}
	return nil
}

func (m *WhateverFilter) GetEnumListSome() *whateverkindfilter.WhateverKindFilter {
	if m != nil {
		return m.EnumListSome
	}
	return nil
}

func (m *WhateverFilter) GetEnumListEvery() *whateverkindfilter.WhateverKindFilter {
	if m != nil {
		return m.EnumListEvery
	}
	return nil
}

func (m *WhateverFilter) GetEnumListNone() *whateverkindfilter.WhateverKindFilter {
	if m != nil {
		return m.EnumListNone
	}
	return nil
}

func (m *WhateverFilter) GetStringListSome() *stringfilter.StringFilter {
	if m != nil {
		return m.StringListSome
	}
	return nil
}

func (m *WhateverFilter) GetStringListEvery() *stringfilter.StringFilter {
	if m != nil {
		return m.StringListEvery
	}
	return nil
}

func (m *WhateverFilter) GetStringListNone() *stringfilter.StringFilter {
	if m != nil {
		return m.StringListNone
	}
	return nil
}

func (m *WhateverFilter) GetIntListSome() *intfilter.IntFilter {
	if m != nil {
		return m.IntListSome
	}
	return nil
}

func (m *WhateverFilter) GetIntListEvery() *intfilter.IntFilter {
	if m != nil {
		return m.IntListEvery
	}
	return nil
}

func (m *WhateverFilter) GetIntListNone() *intfilter.IntFilter {
	if m != nil {
		return m.IntListNone
	}
	return nil
}

func (m *WhateverFilter) GetFloatListSome() *floatfilter.FloatFilter {
	if m != nil {
		return m.FloatListSome
	}
	return nil
}

func (m *WhateverFilter) GetFloatListEvery() *floatfilter.FloatFilter {
	if m != nil {
		return m.FloatListEvery
	}
	return nil
}

func (m *WhateverFilter) GetFloatListNone() *floatfilter.FloatFilter {
	if m != nil {
		return m.FloatListNone
	}
	return nil
}

func (m *WhateverFilter) GetBoolListSome() *boolfilter.BoolFilter {
	if m != nil {
		return m.BoolListSome
	}
	return nil
}

func (m *WhateverFilter) GetBoolListEvery() *boolfilter.BoolFilter {
	if m != nil {
		return m.BoolListEvery
	}
	return nil
}

func (m *WhateverFilter) GetBoolListNone() *boolfilter.BoolFilter {
	if m != nil {
		return m.BoolListNone
	}
	return nil
}

func (m *WhateverFilter) GetObjectListSome() *WhateverFilter {
	if m != nil {
		return m.ObjectListSome
	}
	return nil
}

func (m *WhateverFilter) GetObjectListEvery() *WhateverFilter {
	if m != nil {
		return m.ObjectListEvery
	}
	return nil
}

func (m *WhateverFilter) GetObjectListNone() *WhateverFilter {
	if m != nil {
		return m.ObjectListNone
	}
	return nil
}

func (m *WhateverFilter) GetUnionListSome() *whateverunionfilter.WhateverUnionFilter {
	if m != nil {
		return m.UnionListSome
	}
	return nil
}

func (m *WhateverFilter) GetUnionListEvery() *whateverunionfilter.WhateverUnionFilter {
	if m != nil {
		return m.UnionListEvery
	}
	return nil
}

func (m *WhateverFilter) GetUnionListNone() *whateverunionfilter.WhateverUnionFilter {
	if m != nil {
		return m.UnionListNone
	}
	return nil
}

func (m *WhateverFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *WhateverFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *WhateverFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *WhateverFilter) GetMeta() *servicefilter.TypeMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *WhateverFilter) GetAnd() []*WhateverFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *WhateverFilter) GetOr() []*WhateverFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *WhateverFilter) GetNot() []*WhateverFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *WhateverFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*WhateverFilter)(nil), "whateverfilter.WhateverFilter")
}

func init() { proto.RegisterFile("whateverfilter/whateverfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 732 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xef, 0x4f, 0x13, 0x31,
	0x18, 0xc7, 0xc3, 0x0f, 0x85, 0x3d, 0xfb, 0x81, 0x54, 0xc4, 0x73, 0x02, 0x4e, 0x24, 0x86, 0x17,
	0xd8, 0x01, 0x31, 0x86, 0x10, 0x62, 0x08, 0xc1, 0x05, 0x14, 0x89, 0x01, 0x8d, 0x89, 0x6f, 0xcc,
	0xb1, 0x2b, 0x5b, 0x75, 0x6b, 0x97, 0xbb, 0x0e, 0xc3, 0xdf, 0xe8, 0x3f, 0x65, 0xee, 0xe9, 0xb5,
	0x77, 0x3d, 0x75, 0xbb, 0xf1, 0xee, 0xda, 0xe7, 0xfb, 0x79, 0xfa, 0xfd, 0xb6, 0x77, 0xdd, 0x60,
	0xe3, 0x57, 0xd7, 0x57, 0xec, 0x86, 0x85, 0xd7, 0xbc, 0xa7, 0x58, 0xd8, 0x74, 0x87, 0xdf, 0xe9,
	0x20, 0x94, 0x4a, 0x92, 0x9a, 0x3b, 0x5d, 0xdf, 0x32, 0xe3, 0x9f, 0x5c, 0x04, 0x39, 0x32, 0x9d,
	0x4a, 0xe8, 0x7a, 0x23, 0x52, 0x21, 0x17, 0x9d, 0x44, 0x97, 0x1d, 0x18, 0x45, 0x9d, 0x0b, 0x95,
	0x94, 0xed, 0x93, 0xa9, 0xad, 0x5d, 0xf7, 0xa4, 0x6f, 0xaa, 0x99, 0x67, 0x53, 0x5f, 0xb9, 0x92,
	0xb2, 0x97, 0x94, 0xd3, 0x47, 0x53, 0xa5, 0xc6, 0xd6, 0x50, 0x70, 0x29, 0x72, 0x56, 0x33, 0x73,
	0x46, 0xef, 0x71, 0x93, 0x87, 0xe7, 0x52, 0xac, 0x47, 0x2c, 0xbc, 0xe1, 0x6d, 0x66, 0x62, 0x64,
	0x47, 0x89, 0x66, 0xfd, 0xf7, 0x22, 0xd4, 0xbe, 0x26, 0xcd, 0x5b, 0x58, 0x21, 0xc7, 0x50, 0x62,
	0x62, 0xd8, 0x6f, 0x71, 0xd6, 0x0b, 0xbc, 0xa9, 0xc6, 0xd4, 0x66, 0x79, 0xf7, 0x25, 0xfd, 0x7b,
	0xaf, 0xa8, 0xc1, 0x3e, 0x70, 0x11, 0x68, 0xf4, 0x22, 0x05, 0xc9, 0x01, 0x94, 0xf5, 0xbe, 0xe9,
	0x3e, 0xd3, 0xd8, 0xa7, 0x4e, 0xb3, 0x7b, 0x49, 0x2f, 0x13, 0x01, 0xb2, 0x59, 0x39, 0xd9, 0x86,
	0x79, 0x2e, 0x94, 0x46, 0x67, 0x10, 0x5d, 0xa2, 0x76, 0x9f, 0xe9, 0x69, 0x5c, 0x42, 0xc8, 0xaa,
	0xc8, 0x1e, 0x00, 0x6e, 0xb5, 0x66, 0x66, 0x91, 0xf1, 0x68, 0x66, 0xf7, 0x69, 0x4b, 0x97, 0x91,
	0xcb, 0x68, 0xc9, 0x6b, 0x28, 0xc5, 0xa7, 0xa0, 0xc1, 0x7b, 0x08, 0x2e, 0xd3, 0xf4, 0x5c, 0xe8,
	0x11, 0x16, 0x75, 0x3e, 0x2b, 0x24, 0x87, 0x50, 0x96, 0x57, 0x3f, 0x58, 0x3b, 0x59, 0xf0, 0x3e,
	0x72, 0x6b, 0xd4, 0x7d, 0xed, 0xa8, 0xbb, 0xb5, 0x17, 0x59, 0x84, 0x9c, 0x00, 0xe0, 0x71, 0xea,
	0x06, 0x73, 0xd8, 0x60, 0xf3, 0x5f, 0xa7, 0x6f, 0xbb, 0x7c, 0xd1, 0x72, 0x9d, 0x20, 0x65, 0xc9,
	0x7b, 0xa8, 0xc4, 0x1b, 0x7f, 0xc6, 0x23, 0x75, 0x29, 0xfb, 0xcc, 0x9b, 0x9f, 0xe8, 0xd0, 0x1c,
	0x96, 0x9c, 0x41, 0xd5, 0x8c, 0xdf, 0xdd, 0xb0, 0xf0, 0xd6, 0x2b, 0x4d, 0xd4, 0xcc, 0x85, 0xb3,
	0xce, 0xce, 0xa5, 0x60, 0x1e, 0xdc, 0xcd, 0x59, 0xcc, 0x92, 0x23, 0xa8, 0xe9, 0x57, 0xc4, 0xe6,
	0x2c, 0x8f, 0x7d, 0xa9, 0x72, 0x04, 0x39, 0x86, 0x85, 0x74, 0x46, 0xe7, 0xab, 0x8c, 0x6d, 0x92,
	0x47, 0x5c, 0x27, 0x98, 0xab, 0x3a, 0x89, 0x13, 0x4c, 0xf3, 0x06, 0xca, 0x5c, 0x28, 0x1b, 0xa5,
	0x36, 0xe2, 0x25, 0xcf, 0x0a, 0xc9, 0x1e, 0x54, 0x92, 0xa1, 0xb6, 0xbf, 0x30, 0x02, 0x74, 0x94,
	0x99, 0x15, 0xd1, 0xf2, 0x83, 0x02, 0x2b, 0xa2, 0xd3, 0xb7, 0x50, 0xc5, 0xaf, 0xc5, 0x7a, 0x5d,
	0x1c, 0xf3, 0x71, 0xb9, 0x72, 0x72, 0x08, 0x35, 0x3b, 0xa1, 0x3d, 0x93, 0x31, 0x0d, 0x72, 0x7a,
	0xc7, 0x01, 0x7a, 0x7f, 0x58, 0xd8, 0x01, 0x26, 0xd8, 0x87, 0x4a, 0xfc, 0xe1, 0xda, 0x00, 0x4b,
	0x23, 0x3f, 0x72, 0x47, 0x4b, 0x0e, 0xa0, 0x6a, 0xc6, 0xda, 0xfc, 0xa3, 0x91, 0xb0, 0x2b, 0xce,
	0xae, 0x8c, 0xc6, 0x97, 0x8b, 0xad, 0x8c, 0xae, 0x5b, 0x50, 0xd3, 0xd7, 0x85, 0xf5, 0xfd, 0xb8,
	0xd0, 0x25, 0x93, 0xa3, 0xc8, 0x09, 0x2c, 0xa4, 0x33, 0x3a, 0x83, 0x57, 0xa8, 0x51, 0x1e, 0x73,
	0x1d, 0x61, 0x9e, 0x27, 0x93, 0x3a, 0xc2, 0x64, 0xe7, 0x50, 0xc5, 0xdb, 0xcb, 0x06, 0xab, 0x4f,
	0x78, 0xf9, 0xb9, 0x38, 0xf9, 0x04, 0x35, 0x3b, 0xa1, 0x03, 0x3e, 0x9d, 0xb0, 0x61, 0x8e, 0x77,
	0x1c, 0x62, 0xd0, 0x95, 0x3b, 0x3b, 0xc4, 0xc4, 0x5b, 0x30, 0xc7, 0x83, 0x08, 0xb3, 0xae, 0x62,
	0x27, 0x42, 0xcd, 0xaf, 0x35, 0x3d, 0x35, 0xd7, 0x9d, 0x91, 0x10, 0x0a, 0xf3, 0x3c, 0x88, 0x74,
	0x92, 0xb5, 0xff, 0xca, 0xad, 0x26, 0xe9, 0x8e, 0x3e, 0x9f, 0x8d, 0xec, 0x8e, 0x5e, 0x76, 0x60,
	0xb6, 0xcf, 0x94, 0xef, 0x35, 0x50, 0xba, 0x4a, 0x9d, 0xff, 0x05, 0xf4, 0xf3, 0xed, 0x80, 0x7d,
	0x64, 0xca, 0x4f, 0x28, 0x94, 0x92, 0x6d, 0x98, 0xf1, 0x45, 0xe0, 0x3d, 0x6f, 0xcc, 0x14, 0x38,
	0xed, 0x58, 0x4a, 0x28, 0x4c, 0xcb, 0xd0, 0x5b, 0x2f, 0x04, 0x4c, 0xcb, 0x30, 0x5e, 0x41, 0x48,
	0xe5, 0xbd, 0x28, 0xb6, 0x82, 0x90, 0x8a, 0x10, 0x98, 0xed, 0xfa, 0x51, 0xd7, 0xdb, 0x68, 0x4c,
	0x6d, 0x96, 0x2e, 0xf0, 0xf9, 0xe8, 0xe0, 0xdb, 0x7e, 0x87, 0xab, 0xee, 0xf0, 0x8a, 0xb6, 0x65,
	0xbf, 0xb9, 0xbb, 0x13, 0x29, 0x2e, 0x9b, 0x1d, 0xf9, 0x6a, 0xd0, 0xf3, 0x6f, 0x3b, 0xa1, 0x1c,
	0x8a, 0xa0, 0xd9, 0x09, 0x07, 0xed, 0x66, 0xd4, 0xee, 0xb2, 0xbe, 0x9f, 0xfb, 0xeb, 0xf8, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x70, 0x08, 0xf9, 0xfb, 0x5a, 0x0a, 0x00, 0x00,
}
