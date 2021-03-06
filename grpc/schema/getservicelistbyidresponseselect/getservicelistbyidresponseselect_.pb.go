// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getservicelistbyidresponseselect/getservicelistbyidresponseselect_.proto

/*
Package getservicelistbyidresponseselect is a generated protocol buffer package.

It is generated from these files:
	getservicelistbyidresponseselect/getservicelistbyidresponseselect_.proto

It has these top-level messages:
	GetServiceListByIdResponseSelect
*/
package getservicelistbyidresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import servicesselect "github.com/21stio/go-playground/grpc/schema/servicesselect"
import paginationselect "github.com/21stio/go-playground/grpc/schema/paginationselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetServiceListByIdResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Services         *servicesselect.ServicesSelect         `protobuf:"bytes,3,opt,name=services" json:"services,omitempty"`
	Pagination       *paginationselect.PaginationSelect     `protobuf:"bytes,4,opt,name=pagination" json:"pagination,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,5,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *GetServiceListByIdResponseSelect) Reset()         { *m = GetServiceListByIdResponseSelect{} }
func (m *GetServiceListByIdResponseSelect) String() string { return proto.CompactTextString(m) }
func (*GetServiceListByIdResponseSelect) ProtoMessage()    {}
func (*GetServiceListByIdResponseSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *GetServiceListByIdResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetServiceListByIdResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetServiceListByIdResponseSelect) GetServices() *servicesselect.ServicesSelect {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *GetServiceListByIdResponseSelect) GetPagination() *paginationselect.PaginationSelect {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func (m *GetServiceListByIdResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetServiceListByIdResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetServiceListByIdResponseSelect)(nil), "getservicelistbyidresponseselect.GetServiceListByIdResponseSelect")
}

func init() {
	proto.RegisterFile("getservicelistbyidresponseselect/getservicelistbyidresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x3f, 0x6b, 0xc3, 0x30,
	0x10, 0xc5, 0x71, 0xea, 0x96, 0x44, 0xdd, 0x34, 0x99, 0x50, 0x82, 0x09, 0xa5, 0x78, 0x68, 0x65,
	0x9a, 0x31, 0x5b, 0xbd, 0x34, 0x85, 0xb4, 0x14, 0x67, 0xeb, 0x52, 0x14, 0x5b, 0xc8, 0x02, 0xdb,
	0x32, 0xbe, 0x4b, 0xc1, 0xdf, 0xaa, 0x1f, 0x31, 0x58, 0x96, 0xf3, 0x77, 0xf0, 0xa6, 0x7b, 0xfa,
	0xbd, 0xa7, 0x7b, 0x20, 0xb2, 0x92, 0x02, 0x41, 0xd4, 0x7f, 0x2a, 0x11, 0xb9, 0x02, 0xdc, 0x36,
	0x2a, 0xad, 0x05, 0x54, 0xba, 0x04, 0x01, 0x22, 0x17, 0x09, 0x86, 0x43, 0xc0, 0x2f, 0xab, 0x6a,
	0x8d, 0x9a, 0xfa, 0x43, 0xe0, 0xf4, 0xb9, 0x9f, 0x0b, 0x81, 0xdc, 0xa6, 0x5f, 0x4b, 0x36, 0x6f,
	0xfa, 0x68, 0xc3, 0xc0, 0x92, 0xe7, 0x63, 0x4f, 0x05, 0x15, 0x97, 0xaa, 0xe4, 0xa8, 0xda, 0x10,
	0xc3, 0x5d, 0x0a, 0x96, 0x9c, 0xff, 0x8f, 0x88, 0xff, 0x2e, 0x70, 0xd3, 0xc5, 0xac, 0x15, 0x60,
	0xd4, 0x7c, 0xa4, 0xb1, 0x7d, 0x7f, 0x63, 0x58, 0xfa, 0x40, 0x26, 0x9d, 0xeb, 0x2d, 0xcf, 0x3d,
	0xc7, 0x77, 0x82, 0x71, 0x7c, 0x14, 0xe8, 0x92, 0xb8, 0xed, 0x9e, 0xde, 0xc8, 0x77, 0x82, 0xfb,
	0xc5, 0x13, 0xbb, 0x5e, 0x9e, 0xf5, 0x79, 0x9f, 0x02, 0x79, 0x97, 0x19, 0x1b, 0x0f, 0x5d, 0x92,
	0x71, 0xdf, 0xc0, 0xbb, 0x31, 0xfe, 0x19, 0x3b, 0xaf, 0xc4, 0xec, 0x6a, 0x60, 0x7d, 0x07, 0x9e,
	0x46, 0x84, 0x1c, 0x5b, 0x79, 0xae, 0x71, 0xcf, 0xd9, 0x65, 0x51, 0xf6, 0x7d, 0x10, 0x6c, 0xc2,
	0x89, 0x8b, 0xce, 0x08, 0xe9, 0xb0, 0x15, 0x87, 0xcc, 0xbb, 0x35, 0xd5, 0x4e, 0x14, 0x4a, 0x89,
	0x9b, 0xb5, 0x37, 0x77, 0xbe, 0x13, 0x4c, 0x62, 0x73, 0x8e, 0xbe, 0x7e, 0xd6, 0x52, 0x61, 0xb6,
	0xdb, 0xb2, 0x44, 0x17, 0xe1, 0xe2, 0x15, 0x50, 0xe9, 0x50, 0xea, 0x97, 0x2a, 0xe7, 0x8d, 0xac,
	0xf5, 0xae, 0x4c, 0x43, 0x59, 0x57, 0x49, 0x08, 0x49, 0x26, 0x0a, 0x3e, 0xf8, 0x53, 0xf6, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0x44, 0x5f, 0x59, 0x6d, 0x02, 0x00, 0x00,
}
