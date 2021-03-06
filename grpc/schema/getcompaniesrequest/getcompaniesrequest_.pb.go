// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getcompaniesrequest/getcompaniesrequest_.proto

/*
Package getcompaniesrequest is a generated protocol buffer package.

It is generated from these files:
	getcompaniesrequest/getcompaniesrequest_.proto

It has these top-level messages:
	GetCompaniesRequest
*/
package getcompaniesrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import getcompaniesresponseselect "github.com/21stio/go-playground/grpc/schema/getcompaniesresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetCompaniesRequest struct {
	Meta             *requestmeta.RequestMeta                               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                           `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                           `protobuf:"bytes,3,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                `protobuf:"bytes,4,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                           `protobuf:"bytes,5,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *getcompaniesresponseselect.GetCompaniesResponseSelect `protobuf:"bytes,6,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                 `json:"-"`
}

func (m *GetCompaniesRequest) Reset()                    { *m = GetCompaniesRequest{} }
func (m *GetCompaniesRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCompaniesRequest) ProtoMessage()               {}
func (*GetCompaniesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetCompaniesRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetCompaniesRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *GetCompaniesRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetCompaniesRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *GetCompaniesRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetCompaniesRequest) GetSelect() *getcompaniesresponseselect.GetCompaniesResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetCompaniesRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCompaniesRequest)(nil), "getcompaniesrequest.GetCompaniesRequest")
}

func init() { proto.RegisterFile("getcompaniesrequest/getcompaniesrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x4b, 0x33, 0x31,
	0x14, 0xa4, 0xfd, 0xda, 0x7e, 0x18, 0xf5, 0x92, 0x7a, 0x08, 0x45, 0xa4, 0xf4, 0xd4, 0x83, 0x26,
	0xd8, 0x83, 0x27, 0x41, 0x6c, 0x41, 0x2f, 0x2a, 0x98, 0xde, 0xbc, 0x48, 0x8c, 0xcf, 0xdd, 0x85,
	0xdd, 0xcd, 0x36, 0xc9, 0x0a, 0xfd, 0xc9, 0xfe, 0x0b, 0xe9, 0xdb, 0xac, 0x6c, 0xa4, 0x5e, 0xc2,
	0x9b, 0x9d, 0x99, 0xcd, 0x9b, 0x21, 0x84, 0x27, 0xe0, 0xb5, 0x29, 0x2a, 0x55, 0x66, 0xe0, 0x2c,
	0x6c, 0x6a, 0x70, 0x5e, 0xec, 0xf9, 0xf6, 0xca, 0x2b, 0x6b, 0xbc, 0xa1, 0xe3, 0x3d, 0xdc, 0xe4,
	0x2c, 0x0c, 0x05, 0x78, 0x25, 0x3a, 0x73, 0x30, 0x4d, 0x66, 0x0e, 0xec, 0x67, 0xa6, 0xe1, 0x23,
	0xcb, 0x3d, 0x58, 0x11, 0xa1, 0x56, 0x43, 0xb5, 0xd2, 0x29, 0x08, 0x3c, 0x7f, 0x7c, 0xb9, 0xd1,
	0xca, 0x67, 0xa6, 0xdc, 0xd4, 0x60, 0xb7, 0x22, 0x42, 0xad, 0xe6, 0x3a, 0x5e, 0xc8, 0x55, 0xa6,
	0x74, 0xe0, 0x20, 0x07, 0xfd, 0x3b, 0x47, 0x97, 0x0a, 0xee, 0xd9, 0x57, 0x9f, 0x8c, 0xef, 0xc1,
	0xaf, 0x5a, 0x95, 0x6c, 0x96, 0xa7, 0xe7, 0x64, 0xb0, 0x0b, 0xc0, 0x7a, 0xd3, 0xde, 0xfc, 0x70,
	0xc1, 0x78, 0x27, 0x14, 0x0f, 0x9a, 0x47, 0xf0, 0x4a, 0xa2, 0x8a, 0x2e, 0xc9, 0x71, 0xc8, 0x74,
	0x87, 0x99, 0x58, 0x1f, 0x6d, 0xa7, 0x3c, 0x4a, 0xca, 0xd7, 0x5d, 0x8d, 0x8c, 0x2d, 0x74, 0x46,
	0x86, 0x98, 0x9d, 0xfd, 0x43, 0xef, 0x11, 0x47, 0xc4, 0x57, 0xbb, 0x53, 0x36, 0x14, 0x3d, 0x21,
	0x43, 0xcc, 0xce, 0x06, 0xd3, 0xde, 0xfc, 0x40, 0x36, 0x60, 0x77, 0x7b, 0xdb, 0xcc, 0x33, 0xb2,
	0xc3, 0x70, 0x7b, 0xd4, 0x17, 0x7f, 0xe8, 0x6a, 0x64, 0x6c, 0xa1, 0x4f, 0x64, 0xd4, 0x14, 0xc3,
	0x46, 0x68, 0xbe, 0xe2, 0x7f, 0x77, 0xc7, 0xe3, 0xc2, 0x1a, 0x6a, 0x8d, 0x94, 0x0c, 0x7f, 0xa1,
	0x94, 0x0c, 0x52, 0xe5, 0x52, 0xf6, 0x1f, 0x17, 0xc5, 0x79, 0x79, 0xfb, 0x72, 0x93, 0x64, 0x3e,
	0xad, 0xdf, 0xb8, 0x36, 0x85, 0x58, 0x5c, 0x3a, 0x9f, 0x19, 0x91, 0x98, 0x8b, 0x2a, 0x57, 0xdb,
	0xc4, 0x9a, 0xba, 0x7c, 0x17, 0x89, 0xad, 0xb4, 0x70, 0x3a, 0x85, 0x42, 0xed, 0x7b, 0x84, 0xdf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0xa0, 0x42, 0x2c, 0xae, 0x02, 0x00, 0x00,
}
