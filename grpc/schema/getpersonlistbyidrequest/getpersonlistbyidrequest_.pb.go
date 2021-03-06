// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpersonlistbyidrequest/getpersonlistbyidrequest_.proto

/*
Package getpersonlistbyidrequest is a generated protocol buffer package.

It is generated from these files:
	getpersonlistbyidrequest/getpersonlistbyidrequest_.proto

It has these top-level messages:
	GetPersonListByIdRequest
*/
package getpersonlistbyidrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import id "github.com/21stio/go-playground/grpc/schema/id"
import personskind "github.com/21stio/go-playground/grpc/schema/personskind"
import getpersonlistbyidresponseselect "github.com/21stio/go-playground/grpc/schema/getpersonlistbyidresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPersonListByIdRequest struct {
	Meta             *requestmeta.RequestMeta                                         `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Cache            *cache.Cache                                                     `protobuf:"bytes,2,opt,name=cache" json:"cache,omitempty"`
	Id               *id.Id                                                           `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Kind             *personskind.PersonsKind                                         `protobuf:"varint,4,opt,name=kind,enum=personskind.PersonsKind" json:"kind,omitempty"`
	Select           *getpersonlistbyidresponseselect.GetPersonListByIdResponseSelect `protobuf:"bytes,5,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                          `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                           `json:"-"`
}

func (m *GetPersonListByIdRequest) Reset()                    { *m = GetPersonListByIdRequest{} }
func (m *GetPersonListByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPersonListByIdRequest) ProtoMessage()               {}
func (*GetPersonListByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPersonListByIdRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPersonListByIdRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetPersonListByIdRequest) GetId() *id.Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetPersonListByIdRequest) GetKind() personskind.PersonsKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return personskind.PersonsKind_COMPANY_EMPLOYEES
}

func (m *GetPersonListByIdRequest) GetSelect() *getpersonlistbyidresponseselect.GetPersonListByIdResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetPersonListByIdRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPersonListByIdRequest)(nil), "getpersonlistbyidrequest.GetPersonListByIdRequest")
}

func init() {
	proto.RegisterFile("getpersonlistbyidrequest/getpersonlistbyidrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x25, 0xb1, 0x2d, 0xb8, 0x16, 0x0f, 0x7b, 0x90, 0xa5, 0x87, 0x52, 0x7a, 0xea, 0x41, 0x67,
	0xb1, 0x27, 0x8f, 0x52, 0xa5, 0x52, 0x54, 0x90, 0xf5, 0x22, 0x5e, 0x24, 0xcd, 0x0e, 0xc9, 0x60,
	0x9b, 0x8d, 0x99, 0xed, 0xa1, 0x7f, 0xe3, 0xa7, 0x4a, 0x36, 0x29, 0xe4, 0xd0, 0xe2, 0x65, 0x78,
	0x93, 0xf7, 0xf2, 0xe6, 0x3d, 0x58, 0x71, 0x97, 0xa1, 0x2f, 0xb1, 0x62, 0x57, 0x6c, 0x88, 0xfd,
	0x7a, 0x4f, 0xb6, 0xc2, 0x9f, 0x1d, 0xb2, 0xd7, 0xa7, 0x88, 0x2f, 0x28, 0x2b, 0xe7, 0x9d, 0x54,
	0xa7, 0x04, 0xa3, 0x71, 0x0b, 0xb6, 0xe8, 0x13, 0xdd, 0xc1, 0xed, 0x9f, 0x23, 0x99, 0x26, 0x69,
	0x8e, 0x3a, 0xcc, 0xc3, 0xb7, 0x21, 0x59, 0x4d, 0xf6, 0xb0, 0x8d, 0x1b, 0x63, 0xfe, 0xa6, 0xc2,
	0xea, 0x0e, 0x3e, 0xf0, 0xcb, 0x23, 0xb7, 0xb9, 0x74, 0x05, 0x23, 0xe3, 0x06, 0xd3, 0xa3, 0xe1,
	0xbb, 0x7c, 0xeb, 0x33, 0xfd, 0x8d, 0x85, 0x7a, 0x42, 0xff, 0x16, 0xa4, 0x2f, 0xc4, 0x7e, 0xb1,
	0x5f, 0x59, 0xd3, 0x24, 0x96, 0xd7, 0xa2, 0x57, 0xa7, 0x56, 0xd1, 0x24, 0x9a, 0x5d, 0xcc, 0x15,
	0x74, 0x9a, 0x40, 0xab, 0x79, 0x45, 0x9f, 0x98, 0xa0, 0x92, 0x53, 0xd1, 0x0f, 0x85, 0x54, 0x1c,
	0xe4, 0x43, 0x08, 0x1b, 0x3c, 0xd4, 0xd3, 0x34, 0x94, 0xbc, 0x12, 0x31, 0x59, 0x75, 0x16, 0x04,
	0x03, 0x20, 0x0b, 0x2b, 0x6b, 0x62, 0xb2, 0xf5, 0xa5, 0xba, 0x9d, 0xea, 0x4d, 0xa2, 0xd9, 0xe5,
	0x5c, 0x41, 0xa7, 0x31, 0x34, 0xd9, 0xf8, 0x99, 0x0a, 0x6b, 0x82, 0x4a, 0x7e, 0x88, 0x41, 0xd3,
	0x42, 0xf5, 0x83, 0xd3, 0x3d, 0xfc, 0xd3, 0x16, 0x8e, 0x54, 0x6c, 0xf8, 0xf7, 0xc0, 0x9b, 0xd6,
	0x4f, 0x4a, 0xd1, 0xcb, 0x13, 0xce, 0xd5, 0x60, 0x12, 0xcd, 0xce, 0x4d, 0xc0, 0x8b, 0xe5, 0xe7,
	0x63, 0x46, 0x3e, 0xdf, 0xad, 0x21, 0x75, 0x5b, 0x3d, 0xbf, 0x65, 0x4f, 0x4e, 0x67, 0xee, 0xa6,
	0xdc, 0x24, 0xfb, 0xac, 0x72, 0xbb, 0xc2, 0xea, 0xac, 0x2a, 0x53, 0xcd, 0x69, 0x8e, 0xdb, 0xe4,
	0xe4, 0xab, 0xf9, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x87, 0xee, 0xa4, 0x6e, 0x69, 0x02, 0x00, 0x00,
}
