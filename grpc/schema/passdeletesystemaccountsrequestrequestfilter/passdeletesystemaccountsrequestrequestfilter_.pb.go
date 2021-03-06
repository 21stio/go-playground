// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletesystemaccountsrequestrequestfilter/passdeletesystemaccountsrequestrequestfilter_.proto

/*
Package passdeletesystemaccountsrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passdeletesystemaccountsrequestrequestfilter/passdeletesystemaccountsrequestrequestfilter_.proto

It has these top-level messages:
	PassDeleteSystemAccountsRequestRequestFilter
*/
package passdeletesystemaccountsrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import deletesystemaccountsrequestfilter "github.com/21stio/go-playground/grpc/schema/deletesystemaccountsrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteSystemAccountsRequestRequestFilter struct {
	Meta                        *requestmetafilter.RequestMetaFilter                                 `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	DeleteSystemAccountsRequest *deletesystemaccountsrequestfilter.DeleteSystemAccountsRequestFilter `protobuf:"bytes,2,opt,name=deleteSystemAccountsRequest" json:"deleteSystemAccountsRequest,omitempty"`
	And                         []*PassDeleteSystemAccountsRequestRequestFilter                      `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                          []*PassDeleteSystemAccountsRequestRequestFilter                      `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                         []*PassDeleteSystemAccountsRequestRequestFilter                      `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash                        *string                                                              `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized            []byte                                                               `json:"-"`
}

func (m *PassDeleteSystemAccountsRequestRequestFilter) Reset() {
	*m = PassDeleteSystemAccountsRequestRequestFilter{}
}
func (m *PassDeleteSystemAccountsRequestRequestFilter) String() string {
	return proto.CompactTextString(m)
}
func (*PassDeleteSystemAccountsRequestRequestFilter) ProtoMessage() {}
func (*PassDeleteSystemAccountsRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteSystemAccountsRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteSystemAccountsRequestRequestFilter) GetDeleteSystemAccountsRequest() *deletesystemaccountsrequestfilter.DeleteSystemAccountsRequestFilter {
	if m != nil {
		return m.DeleteSystemAccountsRequest
	}
	return nil
}

func (m *PassDeleteSystemAccountsRequestRequestFilter) GetAnd() []*PassDeleteSystemAccountsRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassDeleteSystemAccountsRequestRequestFilter) GetOr() []*PassDeleteSystemAccountsRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassDeleteSystemAccountsRequestRequestFilter) GetNot() []*PassDeleteSystemAccountsRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassDeleteSystemAccountsRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteSystemAccountsRequestRequestFilter)(nil), "passdeletesystemaccountsrequestrequestfilter.PassDeleteSystemAccountsRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passdeletesystemaccountsrequestrequestfilter/passdeletesystemaccountsrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 295 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0xe9, 0xb5, 0x0a, 0xe6, 0xb6, 0x4c, 0xe1, 0x5c, 0x8a, 0x38, 0x14, 0x39, 0x53, 0xec,
	0xe4, 0xaa, 0x1c, 0x0e, 0x82, 0x20, 0x75, 0x91, 0x2e, 0x1a, 0xdb, 0xd8, 0x56, 0xda, 0xbe, 0x9a,
	0xf7, 0x3a, 0xdc, 0x17, 0xf0, 0xdb, 0xf9, 0x9d, 0xa4, 0xb9, 0x6c, 0x87, 0x3d, 0x3a, 0xdc, 0xf4,
	0x92, 0xf0, 0x7e, 0xff, 0xff, 0xff, 0x25, 0x84, 0xbd, 0xf7, 0x0a, 0xb1, 0xd0, 0x8d, 0x26, 0x8d,
	0x5b, 0x24, 0xdd, 0xaa, 0x3c, 0x87, 0xa1, 0x23, 0x34, 0xfa, 0x7b, 0xd0, 0x48, 0xae, 0x7c, 0xd6,
	0x0d, 0x69, 0x13, 0xcf, 0x69, 0x7e, 0x93, 0xbd, 0x01, 0x02, 0xbe, 0x9e, 0x03, 0xad, 0xae, 0xdc,
	0xb6, 0xd5, 0xa4, 0x9c, 0xe9, 0xde, 0x89, 0x53, 0x5e, 0x3d, 0x4e, 0xa8, 0x3a, 0xf6, 0x60, 0x87,
	0xd3, 0xba, 0xf8, 0x0d, 0xd8, 0xfa, 0x59, 0x21, 0x6e, 0x2c, 0xf0, 0x62, 0x81, 0x3b, 0x07, 0xa4,
	0x3b, 0xc0, 0x95, 0x07, 0xcb, 0xf1, 0x5b, 0x16, 0x8c, 0x89, 0x84, 0x17, 0x7a, 0xd1, 0x32, 0xb9,
	0x94, 0x7b, 0x29, 0xa5, 0xeb, 0x7f, 0xd2, 0xa4, 0x76, 0x4c, 0x6a, 0x09, 0xfe, 0xe3, 0xb1, 0xf3,
	0xe2, 0x7f, 0x1b, 0xb1, 0xb0, 0x8a, 0x1b, 0x79, 0x30, 0xbb, 0x9c, 0x08, 0xeb, 0x1c, 0xa7, 0x8c,
	0x78, 0xc3, 0x7c, 0xd5, 0x15, 0xc2, 0x0f, 0xfd, 0x68, 0x99, 0x64, 0x72, 0xce, 0x3b, 0xc9, 0x39,
	0x77, 0x95, 0x8e, 0x36, 0xfc, 0x8b, 0x2d, 0xc0, 0x88, 0xe0, 0xe8, 0x66, 0x0b, 0x30, 0xe3, 0x64,
	0x1d, 0x90, 0x38, 0x39, 0xfe, 0x64, 0x1d, 0x10, 0xe7, 0x2c, 0xa8, 0x14, 0x56, 0xe2, 0x34, 0xf4,
	0xa2, 0xb3, 0xd4, 0xae, 0xef, 0xb3, 0xec, 0xb5, 0xac, 0xa9, 0x1a, 0x3e, 0x64, 0x0e, 0x6d, 0x9c,
	0xdc, 0x20, 0xd5, 0x10, 0x97, 0x70, 0xdd, 0x37, 0x6a, 0x5b, 0x1a, 0x18, 0xba, 0x22, 0x2e, 0x4d,
	0x9f, 0xc7, 0x98, 0x57, 0xba, 0x55, 0xb3, 0x3e, 0xd6, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21,
	0x51, 0x4a, 0xf2, 0xb4, 0x03, 0x00, 0x00,
}
