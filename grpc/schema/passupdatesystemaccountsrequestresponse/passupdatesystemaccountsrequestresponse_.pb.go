// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatesystemaccountsrequestresponse/passupdatesystemaccountsrequestresponse_.proto

/*
Package passupdatesystemaccountsrequestresponse is a generated protocol buffer package.

It is generated from these files:
	passupdatesystemaccountsrequestresponse/passupdatesystemaccountsrequestresponse_.proto

It has these top-level messages:
	PassUpdateSystemAccountsRequestResponse
*/
package passupdatesystemaccountsrequestresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import updatesystemaccountsrequest "github.com/21stio/go-playground/grpc/schema/updatesystemaccountsrequest"
import error "github.com/21stio/go-playground/grpc/schema/error"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdateSystemAccountsRequestResponse struct {
	Meta                        *responsemeta.ResponseMeta                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	UpdateSystemAccountsRequest *updatesystemaccountsrequest.UpdateSystemAccountsRequest `protobuf:"bytes,2,opt,name=updateSystemAccountsRequest" json:"updateSystemAccountsRequest,omitempty"`
	Errors                      []*error.Error                                           `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Hash                        *string                                                  `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized            []byte                                                   `json:"-"`
}

func (m *PassUpdateSystemAccountsRequestResponse) Reset() {
	*m = PassUpdateSystemAccountsRequestResponse{}
}
func (m *PassUpdateSystemAccountsRequestResponse) String() string { return proto.CompactTextString(m) }
func (*PassUpdateSystemAccountsRequestResponse) ProtoMessage()    {}
func (*PassUpdateSystemAccountsRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateSystemAccountsRequestResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassUpdateSystemAccountsRequestResponse) GetUpdateSystemAccountsRequest() *updatesystemaccountsrequest.UpdateSystemAccountsRequest {
	if m != nil {
		return m.UpdateSystemAccountsRequest
	}
	return nil
}

func (m *PassUpdateSystemAccountsRequestResponse) GetErrors() []*error.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *PassUpdateSystemAccountsRequestResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateSystemAccountsRequestResponse)(nil), "passupdatesystemaccountsrequestresponse.PassUpdateSystemAccountsRequestResponse")
}

func init() {
	proto.RegisterFile("passupdatesystemaccountsrequestresponse/passupdatesystemaccountsrequestresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xe9, 0x6e, 0x11, 0xcc, 0x7a, 0xca, 0xa9, 0xd4, 0x4b, 0x11, 0x61, 0x7b, 0x31, 0xc1,
	0x9e, 0x3c, 0x09, 0x0a, 0x1e, 0x05, 0xe9, 0xb2, 0x1e, 0xbc, 0x48, 0xcc, 0x0e, 0xed, 0x82, 0x6d,
	0x62, 0x66, 0x72, 0x58, 0xff, 0x84, 0x7f, 0x59, 0x36, 0xa6, 0xb0, 0x17, 0x43, 0x2f, 0xc3, 0x4c,
	0xe6, 0xbd, 0xc7, 0x17, 0x18, 0xf6, 0x6a, 0x15, 0xa2, 0xb7, 0x3b, 0x45, 0x80, 0x07, 0x24, 0x18,
	0x94, 0xd6, 0xc6, 0x8f, 0x84, 0x0e, 0xbe, 0x3c, 0x20, 0x39, 0x40, 0x6b, 0x46, 0x04, 0x39, 0x53,
	0xf7, 0x2e, 0xac, 0x33, 0x64, 0xf8, 0x7a, 0xa6, 0xbe, 0xac, 0xa6, 0x6e, 0x00, 0x52, 0xf2, 0x74,
	0x88, 0x51, 0xe5, 0x7d, 0x22, 0x46, 0x26, 0x76, 0x93, 0x9f, 0x83, 0x73, 0xc6, 0xc9, 0x50, 0xe3,
	0xdb, 0xd5, 0xcf, 0x82, 0xad, 0x5f, 0x14, 0xe2, 0x36, 0xd8, 0x37, 0xc1, 0xfe, 0x10, 0xed, 0xed,
	0x9f, 0xbd, 0x8d, 0x28, 0x5c, 0xb0, 0xfc, 0x88, 0x53, 0x64, 0x55, 0x56, 0xaf, 0x9a, 0x52, 0x9c,
	0x32, 0x8a, 0x49, 0xf5, 0x0c, 0xa4, 0xda, 0xa0, 0xe3, 0xdf, 0xec, 0xd2, 0xff, 0x1f, 0x5b, 0x2c,
	0x42, 0xcc, 0x9d, 0x48, 0x90, 0x8b, 0x14, 0x56, 0x2a, 0x9c, 0x5f, 0xb3, 0xb3, 0xf0, 0x4f, 0x2c,
	0x96, 0xd5, 0xb2, 0x5e, 0x35, 0x17, 0x22, 0x8c, 0xe2, 0xe9, 0x58, 0xdb, 0xb8, 0xe3, 0x9c, 0xe5,
	0xbd, 0xc2, 0xbe, 0xc8, 0xab, 0xac, 0x3e, 0x6f, 0x43, 0xff, 0xb8, 0x7d, 0xdb, 0x74, 0x7b, 0xea,
	0xfd, 0x87, 0xd0, 0x66, 0x90, 0xcd, 0x2d, 0xd2, 0xde, 0xc8, 0xce, 0xdc, 0xd8, 0x4f, 0x75, 0xe8,
	0x9c, 0xf1, 0xe3, 0x4e, 0x76, 0xce, 0x6a, 0x89, 0xba, 0x87, 0x41, 0xcd, 0x3d, 0x87, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbd, 0x98, 0x98, 0xab, 0x60, 0x02, 0x00, 0x00,
}