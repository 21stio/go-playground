// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getsocialaccountbyidrequestfilter/getsocialaccountbyidrequestfilter_.proto

/*
Package getsocialaccountbyidrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getsocialaccountbyidrequestfilter/getsocialaccountbyidrequestfilter_.proto

It has these top-level messages:
	GetSocialAccountByIdRequestFilter
*/
package getsocialaccountbyidrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSocialAccountByIdRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Id               *idfilter.IdFilter                   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	And              []*GetSocialAccountByIdRequestFilter `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*GetSocialAccountByIdRequestFilter `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*GetSocialAccountByIdRequestFilter `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *GetSocialAccountByIdRequestFilter) Reset()         { *m = GetSocialAccountByIdRequestFilter{} }
func (m *GetSocialAccountByIdRequestFilter) String() string { return proto.CompactTextString(m) }
func (*GetSocialAccountByIdRequestFilter) ProtoMessage()    {}
func (*GetSocialAccountByIdRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *GetSocialAccountByIdRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetSocialAccountByIdRequestFilter) GetId() *idfilter.IdFilter {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetSocialAccountByIdRequestFilter) GetAnd() []*GetSocialAccountByIdRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetSocialAccountByIdRequestFilter) GetOr() []*GetSocialAccountByIdRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetSocialAccountByIdRequestFilter) GetNot() []*GetSocialAccountByIdRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetSocialAccountByIdRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSocialAccountByIdRequestFilter)(nil), "getsocialaccountbyidrequestfilter.GetSocialAccountByIdRequestFilter")
}

func init() {
	proto.RegisterFile("getsocialaccountbyidrequestfilter/getsocialaccountbyidrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbd, 0x6a, 0xf3, 0x30,
	0x14, 0x86, 0xf1, 0xcf, 0xf7, 0x41, 0x95, 0x4d, 0x93, 0xc8, 0xe4, 0x84, 0x0e, 0xa6, 0x50, 0x99,
	0x7a, 0xea, 0xda, 0x50, 0x5a, 0x52, 0x08, 0x05, 0xb7, 0x74, 0xe8, 0x52, 0x14, 0x49, 0xb5, 0x05,
	0xb6, 0x8f, 0x2b, 0x1f, 0x0f, 0xbe, 0x96, 0xde, 0x6c, 0xb1, 0xad, 0x4c, 0x19, 0xb4, 0x64, 0x3b,
	0xb6, 0x9e, 0xf3, 0xbc, 0xaf, 0x40, 0xe4, 0xa5, 0xd4, 0xd8, 0x83, 0x34, 0xa2, 0x16, 0x52, 0xc2,
	0xd0, 0xe2, 0x71, 0x34, 0xca, 0xea, 0x9f, 0x41, 0xf7, 0xf8, 0x6d, 0x6a, 0xd4, 0x36, 0xf3, 0x12,
	0x5f, 0xbc, 0xb3, 0x80, 0x40, 0x37, 0x5e, 0x72, 0x7d, 0xe3, 0x3e, 0x1b, 0x8d, 0xc2, 0xe9, 0xcf,
	0xfe, 0x38, 0xdd, 0x9a, 0x19, 0xe5, 0x90, 0xd3, 0xe0, 0x4e, 0xb6, 0xbf, 0x11, 0xd9, 0x3c, 0x6b,
	0x7c, 0x9b, 0xb3, 0x1e, 0x96, 0xac, 0xdd, 0xb8, 0x57, 0xc5, 0xa2, 0x7a, 0x9a, 0x61, 0x7a, 0x4f,
	0xe2, 0x49, 0xca, 0x82, 0x24, 0x48, 0x57, 0xf9, 0x35, 0x3f, 0x0b, 0xe2, 0x8e, 0x3f, 0x68, 0x14,
	0xcb, 0x4e, 0x31, 0x6f, 0xd0, 0x2d, 0x09, 0x8d, 0x62, 0xe1, 0xbc, 0x47, 0xf9, 0x29, 0x9d, 0xef,
	0x95, 0xa3, 0x42, 0xa3, 0xe8, 0x07, 0x89, 0x44, 0xab, 0x58, 0x94, 0x44, 0xe9, 0x2a, 0x7f, 0xe4,
	0xde, 0xab, 0x73, 0x6f, 0xe1, 0x62, 0x12, 0xd2, 0x77, 0x12, 0x82, 0x65, 0xf1, 0x05, 0xb5, 0x21,
	0xd8, 0xa9, 0x6d, 0x0b, 0xc8, 0xfe, 0x5d, 0xb2, 0x6d, 0x0b, 0x48, 0x29, 0x89, 0x2b, 0xd1, 0x57,
	0xec, 0x7f, 0x12, 0xa4, 0x57, 0xc5, 0x3c, 0xef, 0x5e, 0x3f, 0x0f, 0xa5, 0xc1, 0x6a, 0x38, 0x72,
	0x09, 0x4d, 0x96, 0xdf, 0xf5, 0x68, 0x20, 0x2b, 0xe1, 0xb6, 0xab, 0xc5, 0x58, 0x5a, 0x18, 0x5a,
	0x95, 0x95, 0xb6, 0x93, 0x59, 0x2f, 0x2b, 0xdd, 0x08, 0xff, 0xf3, 0xfa, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x73, 0x71, 0xac, 0xdb, 0xa4, 0x02, 0x00, 0x00,
}