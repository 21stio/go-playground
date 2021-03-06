// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletespecdescriptionsendpointselect/deletespecdescriptionsendpointselect_.proto

/*
Package deletespecdescriptionsendpointselect is a generated protocol buffer package.

It is generated from these files:
	deletespecdescriptionsendpointselect/deletespecdescriptionsendpointselect_.proto

It has these top-level messages:
	DeleteSpecDescriptionsEndpointSelect
*/
package deletespecdescriptionsendpointselect

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

type DeleteSpecDescriptionsEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteSpecDescriptionsEndpointSelect) Reset()         { *m = DeleteSpecDescriptionsEndpointSelect{} }
func (m *DeleteSpecDescriptionsEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*DeleteSpecDescriptionsEndpointSelect) ProtoMessage()    {}
func (*DeleteSpecDescriptionsEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *DeleteSpecDescriptionsEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *DeleteSpecDescriptionsEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *DeleteSpecDescriptionsEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteSpecDescriptionsEndpointSelect)(nil), "deletespecdescriptionsendpointselect.DeleteSpecDescriptionsEndpointSelect")
}

func init() {
	proto.RegisterFile("deletespecdescriptionsendpointselect/deletespecdescriptionsendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xcf, 0xb1, 0x6a, 0xc3, 0x30,
	0x10, 0x06, 0x60, 0xdc, 0x76, 0xa8, 0x35, 0x6a, 0xf2, 0x50, 0x8a, 0x29, 0x1e, 0xbc, 0xd4, 0xa2,
	0x7d, 0x83, 0x04, 0x07, 0x32, 0x1a, 0x7b, 0xcb, 0x12, 0x1c, 0xe9, 0x90, 0x04, 0xb2, 0x4e, 0x48,
	0x32, 0x24, 0x6f, 0x1f, 0x22, 0x07, 0xe2, 0xd1, 0xdb, 0xdd, 0xff, 0x7f, 0x37, 0x1c, 0xe9, 0x04,
	0x18, 0x88, 0x10, 0x1c, 0x70, 0x01, 0x81, 0x7b, 0xed, 0xa2, 0x46, 0x1b, 0xc0, 0x0a, 0x87, 0xda,
	0xc6, 0x00, 0x06, 0x78, 0x64, 0x5b, 0xd0, 0xb9, 0x71, 0x1e, 0x23, 0xd2, 0x6a, 0x0b, 0xfe, 0xb9,
	0x92, 0xaa, 0x4d, 0x6e, 0x70, 0xc0, 0xdb, 0x95, 0x3b, 0x3c, 0xdd, 0x90, 0x1c, 0xfd, 0x22, 0xf9,
	0x72, 0xb1, 0x33, 0xa6, 0xc8, 0xca, 0xac, 0xfe, 0xec, 0x5f, 0x01, 0xfd, 0x26, 0x64, 0x59, 0x8e,
	0x63, 0x50, 0xc5, 0x5b, 0xaa, 0x57, 0x09, 0xa5, 0xe4, 0x43, 0x3d, 0x9a, 0xf7, 0x32, 0xab, 0xf3,
	0x3e, 0xcd, 0xfb, 0xfe, 0xd4, 0x49, 0x1d, 0xd5, 0x7c, 0x69, 0x38, 0x4e, 0xec, 0xff, 0x2f, 0x44,
	0x8d, 0x4c, 0xe2, 0xaf, 0x33, 0xe3, 0x4d, 0x7a, 0x9c, 0xad, 0x60, 0xd2, 0x3b, 0xce, 0x02, 0x57,
	0x30, 0x8d, 0x9b, 0x5e, 0xbf, 0x07, 0x00, 0x00, 0xff, 0xff, 0x26, 0x56, 0xab, 0x8f, 0x46, 0x01,
	0x00, 0x00,
}
