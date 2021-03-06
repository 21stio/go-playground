// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getwhateverlistbyidresponse/getwhateverlistbyidresponse_.proto

/*
Package getwhateverlistbyidresponse is a generated protocol buffer package.

It is generated from these files:
	getwhateverlistbyidresponse/getwhateverlistbyidresponse_.proto

It has these top-level messages:
	GetWhateverListByIdResponse
*/
package getwhateverlistbyidresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import whatever "github.com/21stio/go-playground/grpc/schema/whatever"
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

type GetWhateverListByIdResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Whatevers        []*whatever.Whatever       `protobuf:"bytes,2,rep,name=whatevers" json:"whatevers,omitempty"`
	Pagination       *pagination.Pagination     `protobuf:"bytes,3,opt,name=pagination" json:"pagination,omitempty"`
	Hash             *string                    `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *GetWhateverListByIdResponse) Reset()                    { *m = GetWhateverListByIdResponse{} }
func (m *GetWhateverListByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*GetWhateverListByIdResponse) ProtoMessage()               {}
func (*GetWhateverListByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetWhateverListByIdResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetWhateverListByIdResponse) GetWhatevers() []*whatever.Whatever {
	if m != nil {
		return m.Whatevers
	}
	return nil
}

func (m *GetWhateverListByIdResponse) GetPagination() *pagination.Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *GetWhateverListByIdResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWhateverListByIdResponse)(nil), "getwhateverlistbyidresponse.GetWhateverListByIdResponse")
}

func init() {
	proto.RegisterFile("getwhateverlistbyidresponse/getwhateverlistbyidresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x6b, 0xf3, 0x30,
	0x10, 0x86, 0xf1, 0x17, 0x2f, 0x51, 0x36, 0x0d, 0x1f, 0xc2, 0xe9, 0x60, 0x3a, 0x79, 0xa9, 0xd4,
	0x7a, 0xe8, 0xd8, 0x21, 0x4b, 0x08, 0xb4, 0x50, 0xb4, 0x14, 0xba, 0x14, 0xc5, 0x3e, 0x24, 0x41,
	0x6c, 0x09, 0xe9, 0xd2, 0xe2, 0xdf, 0xd8, 0x3f, 0x55, 0x6a, 0xac, 0xda, 0x93, 0xb7, 0x57, 0x77,
	0xcf, 0xcb, 0x73, 0x20, 0xf2, 0xa4, 0x01, 0xbf, 0x8c, 0x42, 0xf8, 0x84, 0x70, 0xb1, 0x11, 0xcf,
	0x83, 0x6d, 0x03, 0x44, 0xef, 0xfa, 0x08, 0x62, 0x65, 0xf7, 0xc1, 0x7d, 0x70, 0xe8, 0xe8, 0x7e,
	0x85, 0x29, 0xca, 0x94, 0x3a, 0x40, 0x25, 0x96, 0x8f, 0xa9, 0x5e, 0xb0, 0xd4, 0x15, 0x29, 0xa4,
	0xcd, 0x8d, 0x57, 0xda, 0xf6, 0x0a, 0xad, 0xeb, 0xc5, 0x1c, 0xa7, 0xed, 0xed, 0x77, 0x46, 0xf6,
	0x47, 0xc0, 0xb7, 0xa9, 0xf4, 0x6c, 0x23, 0x1e, 0x86, 0x53, 0x2b, 0x27, 0x05, 0xe5, 0x24, 0xff,
	0xd5, 0xb0, 0xac, 0xcc, 0xaa, 0x5d, 0x5d, 0xf0, 0xa5, 0x9b, 0x27, 0xea, 0x05, 0x50, 0xc9, 0x91,
	0xa3, 0xf7, 0x64, 0x9b, 0x0e, 0x88, 0xec, 0x5f, 0xb9, 0xa9, 0x76, 0x35, 0xe5, 0x69, 0xc2, 0x93,
	0x46, 0xce, 0x10, 0x7d, 0x24, 0x64, 0x3e, 0x8b, 0x6d, 0x46, 0xcf, 0x7f, 0x3e, 0x8f, 0xf8, 0xeb,
	0x5f, 0x94, 0x0b, 0x92, 0x52, 0x92, 0x1b, 0x15, 0x0d, 0xcb, 0xcb, 0xac, 0xda, 0xca, 0x31, 0x1f,
	0x4e, 0xef, 0x47, 0x6d, 0xd1, 0x5c, 0xcf, 0xbc, 0x71, 0x9d, 0xa8, 0x1f, 0x22, 0x5a, 0x27, 0xb4,
	0xbb, 0xf3, 0x17, 0x35, 0xe8, 0xe0, 0xae, 0x7d, 0x2b, 0x74, 0xf0, 0x8d, 0x88, 0x8d, 0x81, 0x4e,
	0xad, 0x7d, 0xcb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xf4, 0x6e, 0x23, 0xd0, 0x01, 0x00,
	0x00,
}
