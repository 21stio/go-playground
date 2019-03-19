// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatefeedsrequestendpointselect/passupdatefeedsrequestendpointselect_.proto

/*
Package passupdatefeedsrequestendpointselect is a generated protocol buffer package.

It is generated from these files:
	passupdatefeedsrequestendpointselect/passupdatefeedsrequestendpointselect_.proto

It has these top-level messages:
	PassUpdateFeedsRequestEndpointSelect
*/
package passupdatefeedsrequestendpointselect

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

type PassUpdateFeedsRequestEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassUpdateFeedsRequestEndpointSelect) Reset()         { *m = PassUpdateFeedsRequestEndpointSelect{} }
func (m *PassUpdateFeedsRequestEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*PassUpdateFeedsRequestEndpointSelect) ProtoMessage()    {}
func (*PassUpdateFeedsRequestEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateFeedsRequestEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *PassUpdateFeedsRequestEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *PassUpdateFeedsRequestEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateFeedsRequestEndpointSelect)(nil), "passupdatefeedsrequestendpointselect.PassUpdateFeedsRequestEndpointSelect")
}

func init() {
	proto.RegisterFile("passupdatefeedsrequestendpointselect/passupdatefeedsrequestendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0x28, 0x48, 0x2c, 0x2e,
	0x2e, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x4d, 0x4b, 0x4d, 0x4d, 0x29, 0x2e, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x49, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0x29, 0x4e, 0xcd, 0x49, 0x4d, 0x2e,
	0xd1, 0x27, 0x46, 0x51, 0xbc, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x0a, 0x31, 0x8a, 0x95,
	0x2a, 0xb8, 0x54, 0x02, 0x12, 0x8b, 0x8b, 0x43, 0xc1, 0xea, 0xdc, 0x40, 0xea, 0x82, 0x20, 0xea,
	0x5c, 0xa1, 0xea, 0x82, 0xc1, 0xea, 0x84, 0x64, 0xb8, 0x38, 0x21, 0x3a, 0x1c, 0x73, 0x72, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x10, 0x02, 0x42, 0x72, 0x5c, 0x5c, 0x10, 0x8e, 0x47, 0x62,
	0x71, 0x86, 0x04, 0x13, 0x58, 0x1a, 0x49, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x03, 0x24, 0xc3, 0xac,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x3b, 0x05, 0x45, 0x05, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x1b, 0x19, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16,
	0xe4, 0x24, 0x56, 0xa6, 0x17, 0xe5, 0x97, 0xe6, 0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17,
	0x27, 0x67, 0xa4, 0xe6, 0x26, 0x12, 0xe5, 0x75, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0xbe,
	0x77, 0x7c, 0x46, 0x01, 0x00, 0x00,
}