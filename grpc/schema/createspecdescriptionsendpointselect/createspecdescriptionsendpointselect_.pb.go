// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createspecdescriptionsendpointselect/createspecdescriptionsendpointselect_.proto

/*
Package createspecdescriptionsendpointselect is a generated protocol buffer package.

It is generated from these files:
	createspecdescriptionsendpointselect/createspecdescriptionsendpointselect_.proto

It has these top-level messages:
	CreateSpecDescriptionsEndpointSelect
*/
package createspecdescriptionsendpointselect

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

type CreateSpecDescriptionsEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateSpecDescriptionsEndpointSelect) Reset()         { *m = CreateSpecDescriptionsEndpointSelect{} }
func (m *CreateSpecDescriptionsEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*CreateSpecDescriptionsEndpointSelect) ProtoMessage()    {}
func (*CreateSpecDescriptionsEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *CreateSpecDescriptionsEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *CreateSpecDescriptionsEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *CreateSpecDescriptionsEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSpecDescriptionsEndpointSelect)(nil), "createspecdescriptionsendpointselect.CreateSpecDescriptionsEndpointSelect")
}

func init() {
	proto.RegisterFile("createspecdescriptionsendpointselect/createspecdescriptionsendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcf, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0x06, 0x70, 0xdc, 0x76, 0xa8, 0x35, 0x6a, 0xf2, 0x50, 0x8a, 0x29, 0x1e, 0xbc, 0xd4, 0xa2,
	0x7d, 0x83, 0xfe, 0x83, 0x8e, 0xc6, 0xde, 0xba, 0x14, 0xe5, 0x7c, 0x48, 0x02, 0x59, 0x27, 0x24,
	0x19, 0x92, 0xb7, 0x0f, 0x91, 0x03, 0xf1, 0xe8, 0xed, 0xee, 0xfb, 0x7e, 0x37, 0x1c, 0xeb, 0x21,
	0xa0, 0x4c, 0x18, 0x3d, 0xc2, 0x84, 0x11, 0x82, 0xf1, 0xc9, 0x90, 0x8b, 0xe8, 0x26, 0x4f, 0xc6,
	0xa5, 0x88, 0x16, 0x21, 0x89, 0x3d, 0xe8, 0xbf, 0xf3, 0x81, 0x12, 0xf1, 0x66, 0x0f, 0x7e, 0x39,
	0xb2, 0xe6, 0x2b, 0xbb, 0xd1, 0x23, 0x7c, 0x6f, 0xdc, 0xcf, 0xd5, 0x8d, 0xd9, 0xf1, 0x27, 0x56,
	0xae, 0x17, 0x1f, 0xd6, 0x56, 0x45, 0x5d, 0xb4, 0x8f, 0xc3, 0x2d, 0xe0, 0xcf, 0x8c, 0xad, 0xcb,
	0xaf, 0x8c, 0xba, 0xba, 0xcb, 0xf5, 0x26, 0xe1, 0x9c, 0x3d, 0xe8, 0x4b, 0x73, 0x5f, 0x17, 0x6d,
	0x39, 0xe4, 0xf9, 0x73, 0xf8, 0xeb, 0x95, 0x49, 0x7a, 0x39, 0x74, 0x40, 0xb3, 0x78, 0x7f, 0x8b,
	0xc9, 0x90, 0x50, 0xf4, 0xea, 0xad, 0x3c, 0xa9, 0x40, 0x8b, 0x9b, 0x84, 0x0a, 0x1e, 0x44, 0x04,
	0x8d, 0xb3, 0xdc, 0xf5, 0xfa, 0x39, 0x00, 0x00, 0xff, 0xff, 0x80, 0x2b, 0xa9, 0x4e, 0x46, 0x01,
	0x00, 0x00,
}