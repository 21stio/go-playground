// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getsocialaccountlistbyidresponse/getsocialaccountlistbyidresponse_.proto

/*
Package getsocialaccountlistbyidresponse is a generated protocol buffer package.

It is generated from these files:
	getsocialaccountlistbyidresponse/getsocialaccountlistbyidresponse_.proto

It has these top-level messages:
	GetSocialAccountListByIdResponse
*/
package getsocialaccountlistbyidresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import feed "github.com/21stio/go-playground/grpc/schema/feed"
import pagination "github.com/21stio/go-playground/grpc/schema/pagination"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSocialAccountListByIdResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	SocialAccounts   []*feed.SocialAccount      `protobuf:"bytes,2,rep,name=socialAccounts" json:"socialAccounts,omitempty"`
	Pagination       *pagination.Pagination     `protobuf:"bytes,3,opt,name=pagination" json:"pagination,omitempty"`
	Hash             *string                    `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *GetSocialAccountListByIdResponse) Reset()         { *m = GetSocialAccountListByIdResponse{} }
func (m *GetSocialAccountListByIdResponse) String() string { return proto.CompactTextString(m) }
func (*GetSocialAccountListByIdResponse) ProtoMessage()    {}
func (*GetSocialAccountListByIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *GetSocialAccountListByIdResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetSocialAccountListByIdResponse) GetSocialAccounts() []*feed.SocialAccount {
	if m != nil {
		return m.SocialAccounts
	}
	return nil
}

func (m *GetSocialAccountListByIdResponse) GetPagination() *pagination.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *GetSocialAccountListByIdResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSocialAccountListByIdResponse)(nil), "getsocialaccountlistbyidresponse.GetSocialAccountListByIdResponse")
}

func init() {
	proto.RegisterFile("getsocialaccountlistbyidresponse/getsocialaccountlistbyidresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0xc7, 0xa9, 0xd7, 0xc5, 0x1c, 0x88, 0x44, 0x90, 0x52, 0x1c, 0x82, 0x53, 0x17, 0x13, 0xec,
	0xe0, 0xe2, 0xe4, 0x2d, 0x2a, 0x9c, 0x22, 0x75, 0x73, 0x91, 0x5c, 0x1a, 0xd3, 0x40, 0xdb, 0x84,
	0xbe, 0xd7, 0xa1, 0x9f, 0xd5, 0x2f, 0x23, 0xad, 0xad, 0xad, 0xb7, 0x74, 0x09, 0xff, 0xbc, 0xfc,
	0xf8, 0xff, 0x1e, 0x84, 0x3c, 0x19, 0x8d, 0xe0, 0x94, 0x95, 0xa5, 0x54, 0xca, 0xb5, 0x35, 0x96,
	0x16, 0xf0, 0xd0, 0xd9, 0xbc, 0xd1, 0xe0, 0x5d, 0x0d, 0x5a, 0xac, 0x01, 0x9f, 0xdc, 0x37, 0x0e,
	0x1d, 0x65, 0x6b, 0x60, 0xcc, 0xa6, 0x54, 0x69, 0x94, 0x62, 0x79, 0x19, 0x3b, 0xe2, 0xf3, 0x2f,
	0xad, 0x73, 0xd1, 0x1f, 0xd3, 0xe4, 0xca, 0x4b, 0x63, 0x6b, 0x89, 0xd6, 0xd5, 0x62, 0x8e, 0xe3,
	0xeb, 0xf5, 0x77, 0x40, 0xd8, 0xa3, 0xc6, 0xf7, 0x41, 0xfb, 0xf0, 0xab, 0xdd, 0x5b, 0xc0, 0x5d,
	0xf7, 0x9c, 0x67, 0x63, 0x3f, 0xe5, 0x24, 0xec, 0x1d, 0x51, 0xc0, 0x82, 0x64, 0x9b, 0xc6, 0x7c,
	0x29, 0xe6, 0x13, 0xf5, 0xa2, 0x51, 0x66, 0x03, 0x47, 0xef, 0xc9, 0x19, 0x2c, 0x0b, 0x21, 0x3a,
	0x61, 0x9b, 0x64, 0x9b, 0x5e, 0xf0, 0x7e, 0x31, 0xfe, 0x4f, 0x96, 0x1d, 0xa1, 0xf4, 0x8e, 0x90,
	0x79, 0xcd, 0x68, 0x33, 0x28, 0x2f, 0xf9, 0x3c, 0xe2, 0x6f, 0x7f, 0x31, 0x5b, 0x90, 0x94, 0x92,
	0xb0, 0x90, 0x50, 0x44, 0x21, 0x0b, 0x92, 0xd3, 0x6c, 0xc8, 0xbb, 0xd7, 0x8f, 0xbd, 0xb1, 0x58,
	0xb4, 0x07, 0xae, 0x5c, 0x25, 0xd2, 0x5b, 0x40, 0xeb, 0x84, 0x71, 0x37, 0xbe, 0x94, 0x9d, 0x69,
	0x5c, 0x5b, 0xe7, 0xc2, 0x34, 0x5e, 0x09, 0x50, 0x85, 0xae, 0xe4, 0xea, 0x47, 0xfd, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xeb, 0xdc, 0xf5, 0xf0, 0xec, 0x01, 0x00, 0x00,
}
