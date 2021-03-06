// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createcompaniesresponse/createcompaniesresponse_.proto

/*
Package createcompaniesresponse is a generated protocol buffer package.

It is generated from these files:
	createcompaniesresponse/createcompaniesresponse_.proto

It has these top-level messages:
	CreateCompaniesResponse
*/
package createcompaniesresponse

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

type CreateCompaniesResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                    `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *CreateCompaniesResponse) Reset()                    { *m = CreateCompaniesResponse{} }
func (m *CreateCompaniesResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateCompaniesResponse) ProtoMessage()               {}
func (*CreateCompaniesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateCompaniesResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateCompaniesResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateCompaniesResponse)(nil), "createcompaniesresponse.CreateCompaniesResponse")
}

func init() {
	proto.RegisterFile("createcompaniesresponse/createcompaniesresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x31, 0xcb, 0xc2, 0x30,
	0x10, 0x40, 0xe9, 0x47, 0x97, 0x2f, 0x6e, 0x59, 0x5a, 0x3a, 0x15, 0xa7, 0x2e, 0x26, 0xd8, 0xc1,
	0x1f, 0x60, 0x71, 0x74, 0xe9, 0x28, 0x88, 0x9c, 0xf1, 0x48, 0x0a, 0xa6, 0x17, 0x92, 0x74, 0xf0,
	0xdf, 0x8b, 0xa5, 0x85, 0x2e, 0xd9, 0xde, 0xdd, 0xbd, 0x37, 0x1c, 0x3b, 0x29, 0x8f, 0x10, 0x51,
	0x91, 0x75, 0x30, 0x0e, 0x18, 0x3c, 0x06, 0x47, 0x63, 0x40, 0x99, 0xd8, 0x3f, 0x84, 0xf3, 0x14,
	0x89, 0x17, 0x89, 0x7b, 0x55, 0xaf, 0x64, 0x31, 0x82, 0xdc, 0x0e, 0x4b, 0xba, 0xbf, 0xb3, 0xa2,
	0x9b, 0xe3, 0x6e, 0x8d, 0xfb, 0xc5, 0xe2, 0x82, 0xe5, 0x3f, 0xb3, 0xcc, 0xea, 0xac, 0xd9, 0xb5,
	0x95, 0xd8, 0xe6, 0x62, 0xb5, 0xae, 0x18, 0xa1, 0x9f, 0x3d, 0xce, 0x59, 0x6e, 0x20, 0x98, 0xf2,
	0xaf, 0xce, 0x9a, 0xff, 0x7e, 0xe6, 0xf3, 0xe5, 0xd6, 0xe9, 0x21, 0x9a, 0xe9, 0x29, 0x14, 0x59,
	0xd9, 0x1e, 0x43, 0x1c, 0x48, 0x6a, 0x3a, 0xb8, 0x37, 0x7c, 0xb4, 0xa7, 0x69, 0x7c, 0x49, 0xed,
	0x9d, 0x92, 0x41, 0x19, 0xb4, 0x90, 0xfa, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x60, 0x3e, 0x8b,
	0xd1, 0x19, 0x01, 0x00, 0x00,
}
