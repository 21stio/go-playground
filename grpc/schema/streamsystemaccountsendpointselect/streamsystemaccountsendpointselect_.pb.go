// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamsystemaccountsendpointselect/streamsystemaccountsendpointselect_.proto

/*
Package streamsystemaccountsendpointselect is a generated protocol buffer package.

It is generated from these files:
	streamsystemaccountsendpointselect/streamsystemaccountsendpointselect_.proto

It has these top-level messages:
	StreamSystemAccountsEndpointSelect
*/
package streamsystemaccountsendpointselect

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

type StreamSystemAccountsEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StreamSystemAccountsEndpointSelect) Reset()         { *m = StreamSystemAccountsEndpointSelect{} }
func (m *StreamSystemAccountsEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*StreamSystemAccountsEndpointSelect) ProtoMessage()    {}
func (*StreamSystemAccountsEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *StreamSystemAccountsEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *StreamSystemAccountsEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *StreamSystemAccountsEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamSystemAccountsEndpointSelect)(nil), "streamsystemaccountsendpointselect.StreamSystemAccountsEndpointSelect")
}

func init() {
	proto.RegisterFile("streamsystemaccountsendpointselect/streamsystemaccountsendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x29, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0x2d, 0xae, 0x2c, 0x2e, 0x49, 0xcd, 0x4d, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29,
	0x4e, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0x04, 0x31, 0x72, 0x52, 0x93, 0x4b, 0xf4, 0x09, 0x2b, 0x89,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x22, 0xac, 0x54, 0xa9, 0x8c, 0x4b, 0x29, 0x18,
	0xac, 0x2a, 0x18, 0xac, 0xca, 0x11, 0xaa, 0xca, 0x15, 0xaa, 0x2a, 0x18, 0xac, 0x4a, 0x48, 0x86,
	0x8b, 0x13, 0xa2, 0xde, 0x31, 0x27, 0x47, 0x82, 0x51, 0x81, 0x51, 0x83, 0x23, 0x08, 0x21, 0x20,
	0x24, 0xc7, 0xc5, 0x05, 0xe1, 0x78, 0x24, 0x16, 0x67, 0x48, 0x30, 0x81, 0xa5, 0x91, 0x44, 0x84,
	0x84, 0xb8, 0x58, 0x32, 0x40, 0x32, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x53, 0x40,
	0x94, 0x5f, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x91, 0x61, 0x71,
	0x49, 0x66, 0xbe, 0x7e, 0x7a, 0xbe, 0x6e, 0x41, 0x4e, 0x62, 0x65, 0x7a, 0x51, 0x7e, 0x69, 0x5e,
	0x8a, 0x7e, 0x7a, 0x51, 0x41, 0xb2, 0x7e, 0x71, 0x72, 0x46, 0x6a, 0x6e, 0x22, 0x11, 0x9e, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x2a, 0x1e, 0x43, 0x3c, 0x01, 0x00, 0x00,
}