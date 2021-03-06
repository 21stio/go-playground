// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatesocialaccountsresponse/updatesocialaccountsresponse_.proto

/*
Package updatesocialaccountsresponse is a generated protocol buffer package.

It is generated from these files:
	updatesocialaccountsresponse/updatesocialaccountsresponse_.proto

It has these top-level messages:
	UpdateSocialAccountsResponse
*/
package updatesocialaccountsresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateSocialAccountsResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *UpdateSocialAccountsResponse) Reset()                    { *m = UpdateSocialAccountsResponse{} }
func (m *UpdateSocialAccountsResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateSocialAccountsResponse) ProtoMessage()               {}
func (*UpdateSocialAccountsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateSocialAccountsResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateSocialAccountsResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateSocialAccountsResponse)(nil), "updatesocialaccountsresponse.UpdateSocialAccountsResponse")
}

func init() {
	proto.RegisterFile("updatesocialaccountsresponse/updatesocialaccountsresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x3f, 0xeb, 0x83, 0x30,
	0x10, 0x40, 0xf1, 0x87, 0xcb, 0x2f, 0xdd, 0x32, 0x89, 0x38, 0x48, 0x27, 0x97, 0x26, 0xd4, 0x4f,
	0xd0, 0x76, 0x2a, 0x85, 0x2e, 0x96, 0x2e, 0x5d, 0xca, 0x19, 0x43, 0x22, 0xa8, 0x17, 0xf2, 0x67,
	0xe8, 0xb7, 0x2f, 0x0d, 0x0a, 0x4e, 0x6e, 0xef, 0xb8, 0x77, 0x0f, 0x8e, 0x9c, 0x82, 0xe9, 0xc0,
	0x4b, 0x87, 0xa2, 0x87, 0x01, 0x84, 0xc0, 0x30, 0x79, 0x67, 0xa5, 0x33, 0x38, 0x39, 0xc9, 0xb7,
	0x96, 0x6f, 0x66, 0x2c, 0x7a, 0xa4, 0xc5, 0x96, 0x94, 0x97, 0x0b, 0x8d, 0xd2, 0x03, 0x5f, 0x0f,
	0xf3, 0xfd, 0xbe, 0x25, 0xc5, 0x33, 0x16, 0x1e, 0xb1, 0x70, 0x9e, 0x0b, 0xcd, 0xac, 0x52, 0x46,
	0xd2, 0x9f, 0x9e, 0x25, 0x65, 0x52, 0xed, 0xea, 0x9c, 0xad, 0x1b, 0x6c, 0xb1, 0xee, 0xd2, 0x43,
	0x13, 0x3d, 0x4a, 0x49, 0xaa, 0xc1, 0xe9, 0xec, 0xaf, 0x4c, 0xaa, 0xff, 0x26, 0xf2, 0xe5, 0xf6,
	0xba, 0xaa, 0xde, 0xeb, 0xd0, 0x32, 0x81, 0x23, 0xaf, 0x8f, 0xce, 0xf7, 0xc8, 0x15, 0x1e, 0xcc,
	0x00, 0x1f, 0x65, 0x31, 0x4c, 0x1d, 0x57, 0xd6, 0x08, 0xee, 0x84, 0x96, 0x23, 0x6c, 0xbe, 0xfd,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x7a, 0x60, 0xc3, 0x32, 0x01, 0x00, 0x00,
}
