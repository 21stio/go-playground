// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getsocialaccountbyidrequest/getsocialaccountbyidrequest_.proto

/*
Package getsocialaccountbyidrequest is a generated protocol buffer package.

It is generated from these files:
	getsocialaccountbyidrequest/getsocialaccountbyidrequest_.proto

It has these top-level messages:
	GetSocialAccountByIdRequest
*/
package getsocialaccountbyidrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import id "github.com/21stio/go-playground/grpc/schema/id"
import getsocialaccountbyidresponseselect "github.com/21stio/go-playground/grpc/schema/getsocialaccountbyidresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSocialAccountByIdRequest struct {
	Meta             *requestmeta.RequestMeta                                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Cache            *cache.Cache                                                           `protobuf:"bytes,2,opt,name=cache" json:"cache,omitempty"`
	Id               *id.Id                                                                 `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Select           *getsocialaccountbyidresponseselect.GetSocialAccountByIdResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                 `json:"-"`
}

func (m *GetSocialAccountByIdRequest) Reset()                    { *m = GetSocialAccountByIdRequest{} }
func (m *GetSocialAccountByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSocialAccountByIdRequest) ProtoMessage()               {}
func (*GetSocialAccountByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetSocialAccountByIdRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetSocialAccountByIdRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetSocialAccountByIdRequest) GetId() *id.Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetSocialAccountByIdRequest) GetSelect() *getsocialaccountbyidresponseselect.GetSocialAccountByIdResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetSocialAccountByIdRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSocialAccountByIdRequest)(nil), "getsocialaccountbyidrequest.GetSocialAccountByIdRequest")
}

func init() {
	proto.RegisterFile("getsocialaccountbyidrequest/getsocialaccountbyidrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xec, 0x06, 0xc6, 0x9d, 0x72, 0x90, 0xb0, 0x81, 0x8c, 0x9d, 0x76, 0xd0, 0x04,
	0xf7, 0x01, 0x04, 0x27, 0x38, 0x0a, 0x7a, 0xc9, 0x6e, 0x1e, 0x94, 0x2c, 0x79, 0xb4, 0x81, 0xb6,
	0xa9, 0x4d, 0x7a, 0xe8, 0x67, 0xf7, 0x22, 0x7b, 0xed, 0xa0, 0x87, 0x59, 0x2f, 0x8f, 0xf7, 0xf2,
	0x7e, 0xf9, 0xf3, 0x0b, 0x21, 0x4f, 0x29, 0x04, 0xef, 0xb4, 0x55, 0xb9, 0xd2, 0xda, 0x35, 0x65,
	0x38, 0xb6, 0xd6, 0xd4, 0xf0, 0xdd, 0x80, 0x0f, 0x62, 0x64, 0xf7, 0xc5, 0xab, 0xda, 0x05, 0x47,
	0x97, 0x23, 0xcc, 0xe2, 0xae, 0x6f, 0x0a, 0x08, 0x4a, 0x0c, 0xfa, 0xfe, 0xf2, 0x82, 0x6a, 0xa5,
	0x33, 0x10, 0x58, 0xcf, 0x67, 0x73, 0x6b, 0x84, 0x35, 0xe7, 0xe9, 0xed, 0x72, 0xbc, 0xaf, 0x5c,
	0xe9, 0xc1, 0x43, 0x0e, 0xfa, 0x2f, 0xcb, 0x21, 0xd2, 0xa7, 0xad, 0x7f, 0x22, 0xb2, 0xdc, 0x43,
	0x38, 0x20, 0xfd, 0xdc, 0xd1, 0xbb, 0x36, 0x31, 0xb2, 0x53, 0xa3, 0xf7, 0x24, 0x3e, 0xe9, 0xb1,
	0x68, 0x15, 0x6d, 0x6e, 0xb6, 0x8c, 0x0f, 0x94, 0x79, 0xcf, 0xbc, 0x43, 0x50, 0x12, 0x29, 0xba,
	0x26, 0x53, 0x34, 0x67, 0x13, 0xc4, 0xe7, 0x1c, 0x27, 0xfe, 0x72, 0xaa, 0xb2, 0x5b, 0xd1, 0x5b,
	0x32, 0xb1, 0x86, 0x5d, 0x21, 0x30, 0xe3, 0xd6, 0xf0, 0xc4, 0xc8, 0x89, 0x35, 0xf4, 0x93, 0xcc,
	0x3a, 0x35, 0x16, 0xe3, 0xee, 0x95, 0xff, 0xff, 0x0a, 0x7e, 0x59, 0xbd, 0x43, 0x0e, 0x88, 0xc8,
	0x3e, 0x95, 0x52, 0x12, 0x67, 0xca, 0x67, 0x6c, 0xba, 0x8a, 0x36, 0xd7, 0x12, 0xfb, 0x5d, 0xf2,
	0xb1, 0x4f, 0x6d, 0xc8, 0x9a, 0x23, 0xd7, 0xae, 0x10, 0xdb, 0x47, 0x1f, 0xac, 0x13, 0xa9, 0x7b,
	0xa8, 0x72, 0xd5, 0xa6, 0xb5, 0x6b, 0x4a, 0x23, 0xd2, 0xba, 0xd2, 0xc2, 0xeb, 0x0c, 0x0a, 0x35,
	0xf6, 0xf9, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xd6, 0x9f, 0x4e, 0x36, 0x02, 0x00, 0x00,
}
