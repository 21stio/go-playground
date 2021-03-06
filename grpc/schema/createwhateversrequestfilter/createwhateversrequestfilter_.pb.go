// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createwhateversrequestfilter/createwhateversrequestfilter_.proto

/*
Package createwhateversrequestfilter is a generated protocol buffer package.

It is generated from these files:
	createwhateversrequestfilter/createwhateversrequestfilter_.proto

It has these top-level messages:
	CreateWhateversRequestFilter
*/
package createwhateversrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import whateverfilter "github.com/21stio/go-playground/grpc/schema/whateverfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateWhateversRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	WhateversSome    *whateverfilter.WhateverFilter       `protobuf:"bytes,2,opt,name=whateversSome" json:"whateversSome,omitempty"`
	WhateversEvery   *whateverfilter.WhateverFilter       `protobuf:"bytes,3,opt,name=whateversEvery" json:"whateversEvery,omitempty"`
	WhateversNone    *whateverfilter.WhateverFilter       `protobuf:"bytes,4,opt,name=whateversNone" json:"whateversNone,omitempty"`
	And              []*CreateWhateversRequestFilter      `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*CreateWhateversRequestFilter      `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*CreateWhateversRequestFilter      `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *CreateWhateversRequestFilter) Reset()                    { *m = CreateWhateversRequestFilter{} }
func (m *CreateWhateversRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*CreateWhateversRequestFilter) ProtoMessage()               {}
func (*CreateWhateversRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateWhateversRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateWhateversRequestFilter) GetWhateversSome() *whateverfilter.WhateverFilter {
	if m != nil {
		return m.WhateversSome
	}
	return nil
}

func (m *CreateWhateversRequestFilter) GetWhateversEvery() *whateverfilter.WhateverFilter {
	if m != nil {
		return m.WhateversEvery
	}
	return nil
}

func (m *CreateWhateversRequestFilter) GetWhateversNone() *whateverfilter.WhateverFilter {
	if m != nil {
		return m.WhateversNone
	}
	return nil
}

func (m *CreateWhateversRequestFilter) GetAnd() []*CreateWhateversRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *CreateWhateversRequestFilter) GetOr() []*CreateWhateversRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *CreateWhateversRequestFilter) GetNot() []*CreateWhateversRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *CreateWhateversRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateWhateversRequestFilter)(nil), "createwhateversrequestfilter.CreateWhateversRequestFilter")
}

func init() {
	proto.RegisterFile("createwhateversrequestfilter/createwhateversrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x69, 0x13, 0xff, 0x6d, 0xd1, 0xc3, 0x9e, 0x96, 0x52, 0x24, 0x48, 0x0f, 0x41, 0x70,
	0x83, 0x39, 0x89, 0x27, 0xf1, 0x4f, 0x91, 0xa2, 0x1e, 0xe2, 0x41, 0xf0, 0x22, 0x6b, 0x3a, 0x26,
	0x81, 0x26, 0x13, 0x37, 0x93, 0x4a, 0xbf, 0x91, 0x1f, 0x53, 0x92, 0x6e, 0x0a, 0xb1, 0x10, 0x94,
	0xde, 0x32, 0x93, 0xf7, 0x7e, 0xef, 0x25, 0xbb, 0xec, 0x2a, 0xd4, 0xa0, 0x08, 0xbe, 0x62, 0x45,
	0xb0, 0x00, 0x5d, 0x68, 0xf8, 0x2c, 0xa1, 0xa0, 0x8f, 0x64, 0x4e, 0xa0, 0xbd, 0xae, 0x97, 0x6f,
	0x32, 0xd7, 0x48, 0xc8, 0x47, 0x5d, 0xa2, 0xe1, 0xa9, 0x19, 0x53, 0x20, 0x65, 0xa0, 0x1b, 0x1b,
	0x43, 0x1a, 0x8e, 0x1b, 0x86, 0x11, 0xb6, 0x47, 0xa3, 0x3a, 0xf9, 0xb6, 0xd9, 0xe8, 0xa6, 0x8e,
	0x7c, 0x69, 0x22, 0x83, 0x15, 0x71, 0x52, 0xeb, 0xf8, 0x05, 0xb3, 0x2b, 0xb6, 0xe8, 0x39, 0x3d,
	0x77, 0xe0, 0x8f, 0xe5, 0x46, 0x9e, 0x34, 0xfa, 0x47, 0x20, 0xb5, 0xf2, 0x04, 0xb5, 0x83, 0xdf,
	0xb2, 0xc3, 0xf5, 0x67, 0x3c, 0x63, 0x0a, 0xa2, 0x5f, 0x23, 0x8e, 0x65, 0xbb, 0x89, 0x6c, 0x82,
	0x8d, 0xb9, 0x6d, 0xe2, 0x13, 0x76, 0xb4, 0x5e, 0xdc, 0x2d, 0x40, 0x2f, 0x85, 0xf5, 0x27, 0xcc,
	0x2f, 0x57, 0xab, 0xcd, 0x13, 0x66, 0x20, 0xec, 0x7f, 0xb6, 0xa9, 0x4c, 0xfc, 0x81, 0x59, 0x2a,
	0x9b, 0x89, 0x1d, 0xc7, 0x72, 0x07, 0xfe, 0xa5, 0xec, 0x3a, 0x2c, 0xd9, 0xf5, 0x5b, 0x83, 0x0a,
	0xc3, 0xa7, 0xac, 0x8f, 0x5a, 0xec, 0x6e, 0x0d, 0xeb, 0xa3, 0xae, 0x9a, 0x65, 0x48, 0x62, 0x6f,
	0xfb, 0x66, 0x19, 0x12, 0xe7, 0xcc, 0x8e, 0x55, 0x11, 0x8b, 0x7d, 0xa7, 0xe7, 0x1e, 0x04, 0xf5,
	0xf3, 0xf5, 0xf4, 0xf5, 0x3e, 0x4a, 0x28, 0x2e, 0xdf, 0x65, 0x88, 0xa9, 0xe7, 0x9f, 0x17, 0x94,
	0xa0, 0x17, 0xe1, 0x59, 0x3e, 0x57, 0xcb, 0x48, 0x63, 0x99, 0xcd, 0xbc, 0x48, 0xe7, 0xa1, 0x57,
	0x84, 0x31, 0xa4, 0xaa, 0xf3, 0xb6, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x80, 0xa8, 0x1e, 0x26,
	0x29, 0x03, 0x00, 0x00,
}
