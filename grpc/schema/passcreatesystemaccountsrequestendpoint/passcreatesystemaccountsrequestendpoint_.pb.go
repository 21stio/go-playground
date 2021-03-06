// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatesystemaccountsrequestendpoint/passcreatesystemaccountsrequestendpoint_.proto

/*
Package passcreatesystemaccountsrequestendpoint is a generated protocol buffer package.

It is generated from these files:
	passcreatesystemaccountsrequestendpoint/passcreatesystemaccountsrequestendpoint_.proto

It has these top-level messages:
	PassCreateSystemAccountsRequestEndpoint
*/
package passcreatesystemaccountsrequestendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import passcreatesystemaccountsrequestrequestfilter "github.com/21stio/go-playground/grpc/schema/passcreatesystemaccountsrequestrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateSystemAccountsRequestEndpoint struct {
	Filter           *passcreatesystemaccountsrequestrequestfilter.PassCreateSystemAccountsRequestRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                                                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                                     `json:"-"`
}

func (m *PassCreateSystemAccountsRequestEndpoint) Reset() {
	*m = PassCreateSystemAccountsRequestEndpoint{}
}
func (m *PassCreateSystemAccountsRequestEndpoint) String() string { return proto.CompactTextString(m) }
func (*PassCreateSystemAccountsRequestEndpoint) ProtoMessage()    {}
func (*PassCreateSystemAccountsRequestEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassCreateSystemAccountsRequestEndpoint) GetFilter() *passcreatesystemaccountsrequestrequestfilter.PassCreateSystemAccountsRequestRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *PassCreateSystemAccountsRequestEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateSystemAccountsRequestEndpoint)(nil), "passcreatesystemaccountsrequestendpoint.PassCreateSystemAccountsRequestEndpoint")
}

func init() {
	proto.RegisterFile("passcreatesystemaccountsrequestendpoint/passcreatesystemaccountsrequestendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0x43, 0x31,
	0x14, 0x86, 0x89, 0x88, 0x60, 0xdc, 0x32, 0x15, 0xa7, 0xe2, 0xd2, 0x2e, 0x26, 0xd8, 0x37, 0x50,
	0xd1, 0x59, 0x6e, 0xd1, 0xa1, 0x8b, 0xc6, 0xf4, 0x98, 0x04, 0x7a, 0x73, 0x62, 0xce, 0xc9, 0xd0,
	0xc7, 0xf2, 0x0d, 0x85, 0x34, 0x7b, 0xef, 0x9d, 0xfe, 0x33, 0x7c, 0x1f, 0xff, 0x0f, 0x47, 0x7e,
	0x64, 0x4b, 0xe4, 0x0a, 0x58, 0x06, 0x3a, 0x12, 0xc3, 0x68, 0x9d, 0xc3, 0x9a, 0x98, 0x0a, 0xfc,
	0x56, 0x20, 0x86, 0xb4, 0xcf, 0x18, 0x13, 0x9b, 0x89, 0xdc, 0xa7, 0xce, 0x05, 0x19, 0xd5, 0x6a,
	0x22, 0x7f, 0xfb, 0x75, 0x06, 0xec, 0xf1, 0x13, 0x0f, 0x0c, 0xc5, 0xcc, 0x81, 0xfb, 0x94, 0xbb,
	0x3f, 0x21, 0x57, 0x6f, 0x96, 0xe8, 0xb9, 0x79, 0xdb, 0xe6, 0x3d, 0x76, 0x6f, 0x38, 0x09, 0x2f,
	0x7d, 0x8d, 0x2a, 0xf2, 0xea, 0x24, 0x2f, 0xc4, 0x52, 0xac, 0x6f, 0x36, 0x3b, 0x3d, 0xa7, 0x51,
	0x9f, 0xa9, 0xe9, 0xf1, 0xda, 0xe0, 0xa1, 0x37, 0x29, 0x25, 0x2f, 0x83, 0xa5, 0xb0, 0xb8, 0x58,
	0x8a, 0xf5, 0xf5, 0xd0, 0xee, 0xa7, 0xf7, 0xdd, 0xd6, 0x47, 0x0e, 0xf5, 0x5b, 0x3b, 0x1c, 0xcd,
	0xe6, 0x81, 0x38, 0xa2, 0xf1, 0x78, 0x9f, 0x0f, 0xf6, 0xe8, 0x0b, 0xd6, 0xb4, 0x37, 0xbe, 0x64,
	0x67, 0xc8, 0x05, 0x18, 0xed, 0xd4, 0xe7, 0xfc, 0x07, 0x00, 0x00, 0xff, 0xff, 0xda, 0xbb, 0x31,
	0x2c, 0xee, 0x01, 0x00, 0x00,
}
