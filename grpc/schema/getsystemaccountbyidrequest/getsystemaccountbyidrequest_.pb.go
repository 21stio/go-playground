// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getsystemaccountbyidrequest/getsystemaccountbyidrequest_.proto

/*
Package getsystemaccountbyidrequest is a generated protocol buffer package.

It is generated from these files:
	getsystemaccountbyidrequest/getsystemaccountbyidrequest_.proto

It has these top-level messages:
	GetSystemAccountByIdRequest
*/
package getsystemaccountbyidrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import id "github.com/21stio/go-playground/grpc/schema/id"
import getsystemaccountbyidresponseselect "github.com/21stio/go-playground/grpc/schema/getsystemaccountbyidresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetSystemAccountByIdRequest struct {
	Meta             *requestmeta.RequestMeta                                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Cache            *cache.Cache                                                           `protobuf:"bytes,2,opt,name=cache" json:"cache,omitempty"`
	Id               *id.Id                                                                 `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Select           *getsystemaccountbyidresponseselect.GetSystemAccountByIdResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                 `json:"-"`
}

func (m *GetSystemAccountByIdRequest) Reset()                    { *m = GetSystemAccountByIdRequest{} }
func (m *GetSystemAccountByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetSystemAccountByIdRequest) ProtoMessage()               {}
func (*GetSystemAccountByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetSystemAccountByIdRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetSystemAccountByIdRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetSystemAccountByIdRequest) GetId() *id.Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetSystemAccountByIdRequest) GetSelect() *getsystemaccountbyidresponseselect.GetSystemAccountByIdResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetSystemAccountByIdRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetSystemAccountByIdRequest)(nil), "getsystemaccountbyidrequest.GetSystemAccountByIdRequest")
}

func init() {
	proto.RegisterFile("getsystemaccountbyidrequest/getsystemaccountbyidrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0xed, 0x06, 0xc6, 0x9d, 0x72, 0x90, 0xb0, 0x81, 0x8c, 0x9d, 0x76, 0xd0, 0x04,
	0xf7, 0x01, 0x04, 0x27, 0x38, 0x0a, 0x7a, 0xc9, 0x6e, 0x1e, 0x94, 0x2c, 0x79, 0xb4, 0x81, 0xb5,
	0xa9, 0x4d, 0x7a, 0xe8, 0x67, 0xf7, 0x22, 0x7b, 0xed, 0xa0, 0x87, 0x59, 0x2f, 0x8f, 0xf7, 0xf2,
	0x7e, 0xf9, 0xf3, 0x0b, 0x21, 0x4f, 0x19, 0x04, 0xdf, 0xfa, 0x00, 0x85, 0xd2, 0xda, 0x35, 0x65,
	0x38, 0xb4, 0xd6, 0xd4, 0xf0, 0xdd, 0x80, 0x0f, 0x62, 0x64, 0xf7, 0xc5, 0xab, 0xda, 0x05, 0x47,
	0x17, 0x23, 0xcc, 0xfc, 0xae, 0x6f, 0x0a, 0x08, 0x4a, 0x0c, 0xfa, 0xfe, 0xf2, 0x9c, 0x6a, 0xa5,
	0x73, 0x10, 0x58, 0xcf, 0x67, 0x33, 0x6b, 0x84, 0x35, 0xe7, 0xe9, 0xed, 0x72, 0xbc, 0xaf, 0x5c,
	0xe9, 0xc1, 0xc3, 0x11, 0xf4, 0x5f, 0x96, 0x43, 0xa4, 0x4f, 0x5b, 0xfd, 0x44, 0x64, 0xb1, 0x83,
	0xb0, 0x47, 0xfa, 0xb9, 0xa3, 0xb7, 0x6d, 0x6a, 0x64, 0xa7, 0x46, 0xef, 0x49, 0x72, 0xd2, 0x63,
	0xd1, 0x32, 0x5a, 0xdf, 0x6c, 0x18, 0x1f, 0x28, 0xf3, 0x9e, 0x79, 0x87, 0xa0, 0x24, 0x52, 0x74,
	0x45, 0x26, 0x68, 0xce, 0x62, 0xc4, 0x67, 0x1c, 0x27, 0xfe, 0x72, 0xaa, 0xb2, 0x5b, 0xd1, 0x5b,
	0x12, 0x5b, 0xc3, 0xae, 0x10, 0x98, 0x72, 0x6b, 0x78, 0x6a, 0x64, 0x6c, 0x0d, 0xfd, 0x24, 0xd3,
	0x4e, 0x8d, 0x25, 0xb8, 0x7b, 0xe5, 0xff, 0xbf, 0x82, 0x5f, 0x56, 0xef, 0x90, 0x3d, 0x22, 0xb2,
	0x4f, 0xa5, 0x94, 0x24, 0xb9, 0xf2, 0x39, 0x9b, 0x2c, 0xa3, 0xf5, 0xb5, 0xc4, 0x7e, 0x9b, 0x7e,
	0xec, 0x32, 0x1b, 0xf2, 0xe6, 0xc0, 0xb5, 0x2b, 0xc4, 0xe6, 0xd1, 0x07, 0xeb, 0x44, 0xe6, 0x1e,
	0xaa, 0xa3, 0x6a, 0xb3, 0xda, 0x35, 0xa5, 0x11, 0x59, 0x5d, 0x69, 0xe1, 0x75, 0x0e, 0x85, 0x1a,
	0xfb, 0xfc, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xb4, 0xfb, 0x71, 0x36, 0x02, 0x00, 0x00,
}
