// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatepersonsrequestendpointselect/passupdatepersonsrequestendpointselect_.proto

/*
Package passupdatepersonsrequestendpointselect is a generated protocol buffer package.

It is generated from these files:
	passupdatepersonsrequestendpointselect/passupdatepersonsrequestendpointselect_.proto

It has these top-level messages:
	PassUpdatePersonsRequestEndpointSelect
*/
package passupdatepersonsrequestendpointselect

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

type PassUpdatePersonsRequestEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassUpdatePersonsRequestEndpointSelect) Reset() {
	*m = PassUpdatePersonsRequestEndpointSelect{}
}
func (m *PassUpdatePersonsRequestEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*PassUpdatePersonsRequestEndpointSelect) ProtoMessage()    {}
func (*PassUpdatePersonsRequestEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdatePersonsRequestEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *PassUpdatePersonsRequestEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *PassUpdatePersonsRequestEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdatePersonsRequestEndpointSelect)(nil), "passupdatepersonsrequestendpointselect.PassUpdatePersonsRequestEndpointSelect")
}

func init() {
	proto.RegisterFile("passupdatepersonsrequestendpointselect/passupdatepersonsrequestendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd0, 0xb1, 0x4a, 0xc7, 0x30,
	0x10, 0x06, 0x70, 0xaa, 0x0e, 0x36, 0x63, 0xa6, 0x0e, 0x22, 0xc5, 0xa1, 0x74, 0xb1, 0x41, 0xdf,
	0x40, 0x41, 0x70, 0x2c, 0xb1, 0x2e, 0x2e, 0x12, 0xd3, 0x23, 0x29, 0xa4, 0xb9, 0x98, 0x4b, 0x06,
	0x7d, 0x7a, 0x31, 0x11, 0xfc, 0x8f, 0xdd, 0xee, 0xbe, 0xef, 0x77, 0xcb, 0xb1, 0x25, 0x28, 0xa2,
	0x1c, 0x56, 0x95, 0x20, 0x40, 0x24, 0xf4, 0x14, 0xe1, 0x33, 0x03, 0x25, 0xf0, 0x6b, 0xc0, 0xcd,
	0x27, 0x02, 0x07, 0x3a, 0x89, 0x63, 0xec, 0x7d, 0x0a, 0x11, 0x13, 0xf2, 0xe1, 0x18, 0xbf, 0xf9,
	0x66, 0xc3, 0xac, 0x88, 0x5e, 0x8b, 0x9c, 0xab, 0x94, 0x55, 0x3e, 0xfd, 0xc9, 0x97, 0x22, 0xf9,
	0x15, 0x6b, 0xeb, 0xcd, 0x83, 0x73, 0x5d, 0xd3, 0x37, 0xe3, 0xa5, 0xfc, 0x0f, 0xf8, 0x35, 0x63,
	0x75, 0x79, 0x56, 0x64, 0xbb, 0xb3, 0x52, 0x9f, 0x24, 0x9c, 0xb3, 0x0b, 0xfb, 0xdb, 0x9c, 0xf7,
	0xcd, 0xd8, 0xca, 0x32, 0x3f, 0x2e, 0x6f, 0xd2, 0x6c, 0xc9, 0xe6, 0x8f, 0x49, 0xe3, 0x2e, 0xee,
	0xef, 0x28, 0x6d, 0x28, 0x0c, 0xde, 0x06, 0xa7, 0xbe, 0x4c, 0xc4, 0xec, 0x57, 0x61, 0x62, 0xd0,
	0x82, 0xb4, 0x85, 0x5d, 0x1d, 0x7c, 0xc0, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0x10, 0xb9,
	0x5e, 0x50, 0x01, 0x00, 0x00,
}
