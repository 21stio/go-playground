// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getsocialaccountbyidendpointselect/getsocialaccountbyidendpointselect_.proto

/*
Package getsocialaccountbyidendpointselect is a generated protocol buffer package.

It is generated from these files:
	getsocialaccountbyidendpointselect/getsocialaccountbyidendpointselect_.proto

It has these top-level messages:
	GetSocialAccountByIdEndpointSelect
*/
package getsocialaccountbyidendpointselect

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

type GetSocialAccountByIdEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetSocialAccountByIdEndpointSelect) Reset()         { *m = GetSocialAccountByIdEndpointSelect{} }
func (m *GetSocialAccountByIdEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*GetSocialAccountByIdEndpointSelect) ProtoMessage()    {}
func (*GetSocialAccountByIdEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *GetSocialAccountByIdEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetSocialAccountByIdEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetSocialAccountByIdEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSocialAccountByIdEndpointSelect)(nil), "getsocialaccountbyidendpointselect.GetSocialAccountByIdEndpointSelect")
}

func init() {
	proto.RegisterFile("getsocialaccountbyidendpointselect/getsocialaccountbyidendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x49, 0x4f, 0x2d, 0x29,
	0xce, 0x4f, 0xce, 0x4c, 0xcc, 0x49, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x49, 0xaa, 0xcc, 0x4c,
	0x49, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0x29, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1, 0x27,
	0xac, 0x24, 0x5e, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x48, 0x89, 0xb0, 0x52, 0xa5, 0x32, 0x2e,
	0x25, 0xf7, 0xd4, 0x92, 0x60, 0xb0, 0x2a, 0x47, 0x88, 0x2a, 0xa7, 0x4a, 0xcf, 0x14, 0x57, 0xa8,
	0xaa, 0x60, 0xb0, 0x2a, 0x21, 0x19, 0x2e, 0x4e, 0x88, 0x7a, 0xc7, 0x9c, 0x1c, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x8e, 0x20, 0x84, 0x80, 0x90, 0x1c, 0x17, 0x17, 0x84, 0xe3, 0x91, 0x58, 0x9c, 0x21,
	0xc1, 0x04, 0x96, 0x46, 0x12, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x00, 0xc9, 0x30, 0x2b, 0x30, 0x6a,
	0x70, 0x06, 0x81, 0xd9, 0x4e, 0x01, 0x51, 0x7e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0x46, 0x86, 0xc5, 0x25, 0x99, 0xf9, 0xfa, 0xe9, 0xf9, 0xba, 0x05, 0x39, 0x89,
	0x95, 0xe9, 0x45, 0xf9, 0xa5, 0x79, 0x29, 0xfa, 0xe9, 0x45, 0x05, 0xc9, 0xfa, 0xc5, 0xc9, 0x19,
	0xa9, 0xb9, 0x89, 0x44, 0x78, 0x1a, 0x10, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x8c, 0x48, 0xb1, 0x3c,
	0x01, 0x00, 0x00,
}
