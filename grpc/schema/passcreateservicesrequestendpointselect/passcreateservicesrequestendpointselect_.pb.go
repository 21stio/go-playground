// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreateservicesrequestendpointselect/passcreateservicesrequestendpointselect_.proto

/*
Package passcreateservicesrequestendpointselect is a generated protocol buffer package.

It is generated from these files:
	passcreateservicesrequestendpointselect/passcreateservicesrequestendpointselect_.proto

It has these top-level messages:
	PassCreateServicesRequestEndpointSelect
*/
package passcreateservicesrequestendpointselect

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

type PassCreateServicesRequestEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassCreateServicesRequestEndpointSelect) Reset() {
	*m = PassCreateServicesRequestEndpointSelect{}
}
func (m *PassCreateServicesRequestEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*PassCreateServicesRequestEndpointSelect) ProtoMessage()    {}
func (*PassCreateServicesRequestEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateServicesRequestEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *PassCreateServicesRequestEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *PassCreateServicesRequestEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateServicesRequestEndpointSelect)(nil), "passcreateservicesrequestendpointselect.PassCreateServicesRequestEndpointSelect")
}

func init() {
	proto.RegisterFile("passcreateservicesrequestendpointselect/passcreateservicesrequestendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0xdc, 0x76, 0xa8, 0x35, 0x6a, 0xf2, 0x50, 0x8a, 0xe9, 0x62, 0x2f, 0xb5, 0x68,
	0xdf, 0x20, 0x09, 0x81, 0x8c, 0xc1, 0x26, 0x19, 0xb2, 0x04, 0x45, 0x3e, 0x24, 0x81, 0x6c, 0x29,
	0x3a, 0x39, 0x10, 0xf2, 0xf2, 0x21, 0xb2, 0x21, 0x19, 0xbd, 0xdd, 0xfd, 0xff, 0x77, 0xcb, 0x91,
	0xbd, 0xe3, 0x88, 0xc2, 0x03, 0x0f, 0x80, 0xe0, 0x2f, 0x5a, 0x00, 0x7a, 0x38, 0x0f, 0x80, 0x01,
	0xfa, 0xd6, 0x59, 0xdd, 0x07, 0x04, 0x03, 0x22, 0xb0, 0x99, 0xee, 0x58, 0x39, 0x6f, 0x83, 0xa5,
	0xc5, 0x4c, 0xff, 0x73, 0x23, 0xc5, 0x96, 0x23, 0xae, 0x22, 0x6d, 0x26, 0x5a, 0x8f, 0x74, 0x3d,
	0xd1, 0x26, 0x52, 0xfa, 0x45, 0xd2, 0xf1, 0x68, 0x61, 0x4c, 0x96, 0xe4, 0x49, 0xf9, 0x59, 0x3f,
	0x03, 0xfa, 0x4d, 0xc8, 0xb8, 0x6c, 0x38, 0xaa, 0xec, 0x2d, 0xd6, 0x2f, 0x09, 0xa5, 0xe4, 0x43,
	0x3d, 0x9a, 0xf7, 0x3c, 0x29, 0xd3, 0x3a, 0xce, 0xcb, 0xdd, 0xa1, 0x91, 0x3a, 0xa8, 0xe1, 0x54,
	0x09, 0xdb, 0xb1, 0xff, 0x3f, 0x0c, 0xda, 0x32, 0x69, 0x7f, 0x9d, 0xe1, 0x57, 0xe9, 0xed, 0xd0,
	0xb7, 0x4c, 0x7a, 0x27, 0x18, 0x0a, 0x05, 0x1d, 0x9f, 0xfb, 0x83, 0x7b, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa7, 0xef, 0x89, 0xf6, 0x55, 0x01, 0x00, 0x00,
}