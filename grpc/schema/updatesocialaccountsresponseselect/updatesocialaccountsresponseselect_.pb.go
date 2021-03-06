// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatesocialaccountsresponseselect/updatesocialaccountsresponseselect_.proto

/*
Package updatesocialaccountsresponseselect is a generated protocol buffer package.

It is generated from these files:
	updatesocialaccountsresponseselect/updatesocialaccountsresponseselect_.proto

It has these top-level messages:
	UpdateSocialAccountsResponseSelect
*/
package updatesocialaccountsresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateSocialAccountsResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *UpdateSocialAccountsResponseSelect) Reset()         { *m = UpdateSocialAccountsResponseSelect{} }
func (m *UpdateSocialAccountsResponseSelect) String() string { return proto.CompactTextString(m) }
func (*UpdateSocialAccountsResponseSelect) ProtoMessage()    {}
func (*UpdateSocialAccountsResponseSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *UpdateSocialAccountsResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *UpdateSocialAccountsResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateSocialAccountsResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *UpdateSocialAccountsResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateSocialAccountsResponseSelect)(nil), "updatesocialaccountsresponseselect.UpdateSocialAccountsResponseSelect")
}

func init() {
	proto.RegisterFile("updatesocialaccountsresponseselect/updatesocialaccountsresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0x89, 0x76, 0xf0, 0xe2, 0x96, 0xa9, 0x88, 0x48, 0xe9, 0x20, 0x1d, 0xb4, 0xc1, 0x1b,
	0xdd, 0xce, 0xc9, 0x41, 0x45, 0x7a, 0xb8, 0xb8, 0xc8, 0x33, 0xf7, 0x68, 0x0b, 0x69, 0x5f, 0xe8,
	0x7b, 0x1d, 0xfc, 0xb7, 0xfc, 0x0b, 0xe5, 0xd2, 0x1e, 0x0a, 0x37, 0xd4, 0x2d, 0xf9, 0xf2, 0xfb,
	0x7e, 0x1f, 0x44, 0x3f, 0x8d, 0x61, 0x07, 0x82, 0x4c, 0xae, 0x05, 0x0f, 0xce, 0xd1, 0xd8, 0x0b,
	0x0f, 0xc8, 0x81, 0x7a, 0x46, 0x46, 0x8f, 0x4e, 0xec, 0x32, 0xf2, 0x51, 0x86, 0x81, 0x84, 0x4c,
	0xbe, 0x8c, 0x5e, 0xdc, 0x1c, 0xee, 0x1d, 0x0a, 0xcc, 0x0b, 0xc7, 0xd1, 0x6c, 0xcc, 0xbf, 0x95,
	0xce, 0xdf, 0xa2, 0x74, 0x1b, 0xa5, 0x9b, 0x59, 0x5a, 0xcd, 0x8d, 0x6d, 0xa4, 0xcd, 0xa5, 0x5e,
	0x4d, 0xbd, 0x8d, 0xf7, 0xa9, 0xca, 0x54, 0x71, 0x56, 0xfd, 0x06, 0xe6, 0x5e, 0x27, 0x7b, 0x73,
	0x7a, 0x92, 0xa9, 0xe2, 0x7c, 0x7d, 0x5d, 0x1e, 0xcf, 0x95, 0x07, 0xdf, 0x33, 0x0a, 0x4c, 0xce,
	0x2a, 0x76, 0xcc, 0x95, 0xd6, 0x13, 0xf2, 0x08, 0xdc, 0xa4, 0xa7, 0x51, 0xfd, 0x27, 0x31, 0x46,
	0x27, 0xcd, 0xfe, 0x25, 0xc9, 0x54, 0xb1, 0xaa, 0xe2, 0xf9, 0xe1, 0xf5, 0xfd, 0xa5, 0x6e, 0xa5,
	0x19, 0x3f, 0x4b, 0x47, 0x9d, 0x5d, 0xdf, 0xb1, 0xb4, 0x64, 0x6b, 0xba, 0x0d, 0x1e, 0xbe, 0xea,
	0x81, 0xc6, 0x7e, 0x67, 0xeb, 0x21, 0x38, 0xcb, 0xae, 0xc1, 0x0e, 0xfe, 0xf1, 0xbf, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x14, 0x6c, 0x0c, 0xdb, 0xa7, 0x01, 0x00, 0x00,
}
