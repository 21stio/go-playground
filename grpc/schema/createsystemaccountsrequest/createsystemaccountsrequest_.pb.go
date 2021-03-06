// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createsystemaccountsrequest/createsystemaccountsrequest_.proto

/*
Package createsystemaccountsrequest is a generated protocol buffer package.

It is generated from these files:
	createsystemaccountsrequest/createsystemaccountsrequest_.proto

It has these top-level messages:
	CreateSystemAccountsRequest
*/
package createsystemaccountsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import systemaccount "github.com/21stio/go-playground/grpc/schema/systemaccount"
import createsystemaccountsresponseselect "github.com/21stio/go-playground/grpc/schema/createsystemaccountsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateSystemAccountsRequest struct {
	Meta             *requestmeta.RequestMeta                                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                           `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	SystemAccounts   []*systemaccount.SystemAccount                                         `protobuf:"bytes,3,rep,name=systemAccounts" json:"systemAccounts,omitempty"`
	Select           *createsystemaccountsresponseselect.CreateSystemAccountsResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                 `json:"-"`
}

func (m *CreateSystemAccountsRequest) Reset()                    { *m = CreateSystemAccountsRequest{} }
func (m *CreateSystemAccountsRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSystemAccountsRequest) ProtoMessage()               {}
func (*CreateSystemAccountsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateSystemAccountsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateSystemAccountsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *CreateSystemAccountsRequest) GetSystemAccounts() []*systemaccount.SystemAccount {
	if m != nil {
		return m.SystemAccounts
	}
	return nil
}

func (m *CreateSystemAccountsRequest) GetSelect() *createsystemaccountsresponseselect.CreateSystemAccountsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *CreateSystemAccountsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateSystemAccountsRequest)(nil), "createsystemaccountsrequest.CreateSystemAccountsRequest")
}

func init() {
	proto.RegisterFile("createsystemaccountsrequest/createsystemaccountsrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0x86, 0xe9, 0xcf, 0xf7, 0x81, 0x11, 0x5d, 0x64, 0x15, 0x5a, 0x91, 0xd2, 0x55, 0x17, 0x9a,
	0x60, 0x2f, 0x40, 0xb0, 0x4a, 0x45, 0xd0, 0x4d, 0xba, 0x73, 0xa1, 0xc4, 0x78, 0x9c, 0x19, 0x98,
	0x99, 0x8c, 0x39, 0x19, 0xa1, 0x97, 0xe8, 0x5d, 0x89, 0x99, 0x88, 0x89, 0xd4, 0x71, 0x77, 0xc2,
	0x3c, 0xe7, 0xe5, 0x7d, 0x0e, 0x43, 0xce, 0xb5, 0x05, 0xe5, 0x00, 0xb7, 0xe8, 0xa0, 0x52, 0x5a,
	0x9b, 0xb6, 0x76, 0x68, 0xe1, 0xb5, 0x05, 0x74, 0xa2, 0xe7, 0xdb, 0x23, 0x6f, 0xac, 0x71, 0x86,
	0x4e, 0x7b, 0x98, 0xc9, 0x71, 0x18, 0x2a, 0x70, 0x4a, 0x44, 0x73, 0x58, 0x9e, 0xcc, 0x11, 0xec,
	0x5b, 0xa1, 0xe1, 0xa5, 0x28, 0x1d, 0x58, 0x91, 0xbc, 0xbe, 0x99, 0x38, 0x5a, 0x24, 0xaf, 0x2f,
	0xe6, 0x76, 0x77, 0x09, 0x6c, 0x4c, 0x8d, 0x80, 0x50, 0x82, 0xfe, 0xcd, 0x25, 0x46, 0x42, 0xda,
	0xfc, 0x7d, 0x48, 0xa6, 0x97, 0x9e, 0xde, 0x78, 0xfa, 0x22, 0xd0, 0xb2, 0x13, 0xa0, 0x27, 0x64,
	0xfc, 0x29, 0xc1, 0x06, 0xb3, 0xc1, 0x62, 0x7f, 0xc9, 0x78, 0x24, 0xc6, 0x03, 0x73, 0x07, 0x4e,
	0x49, 0x4f, 0xd1, 0x15, 0x39, 0x08, 0x5e, 0x6b, 0xef, 0xc5, 0x86, 0x7e, 0xed, 0x88, 0x27, 0xb6,
	0x7c, 0x13, 0x33, 0x32, 0x5d, 0xa1, 0x57, 0xe4, 0x10, 0x93, 0x2a, 0x6c, 0x34, 0x1b, 0x75, 0x21,
	0xb1, 0x0f, 0x4f, 0xfa, 0xca, 0x1f, 0x3b, 0xf4, 0x81, 0xfc, 0xef, 0x44, 0xd9, 0xd8, 0x57, 0x58,
	0xf3, 0xbf, 0x6f, 0xc2, 0x77, 0x1f, 0xa2, 0x43, 0x36, 0x1e, 0x91, 0x21, 0x95, 0x52, 0x32, 0xce,
	0x15, 0xe6, 0xec, 0xdf, 0x6c, 0xb0, 0xd8, 0x93, 0x7e, 0x5e, 0xdd, 0xdc, 0x5f, 0x67, 0x85, 0xcb,
	0xdb, 0x27, 0xae, 0x4d, 0x25, 0x96, 0x67, 0xe8, 0x0a, 0x23, 0x32, 0x73, 0xda, 0x94, 0x6a, 0x9b,
	0x59, 0xd3, 0xd6, 0xcf, 0x22, 0xb3, 0x8d, 0x16, 0xa8, 0x73, 0xa8, 0x54, 0xdf, 0x0f, 0xf7, 0x11,
	0x00, 0x00, 0xff, 0xff, 0xf7, 0xcd, 0x47, 0x3d, 0xaa, 0x02, 0x00, 0x00,
}
