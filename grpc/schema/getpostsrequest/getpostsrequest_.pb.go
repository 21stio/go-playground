// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpostsrequest/getpostsrequest_.proto

/*
Package getpostsrequest is a generated protocol buffer package.

It is generated from these files:
	getpostsrequest/getpostsrequest_.proto

It has these top-level messages:
	GetPostsRequest
*/
package getpostsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import getpostsresponseselect "github.com/21stio/go-playground/grpc/schema/getpostsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPostsRequest struct {
	Meta             *requestmeta.RequestMeta                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                   `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                   `protobuf:"bytes,3,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                        `protobuf:"bytes,4,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                   `protobuf:"bytes,5,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *getpostsresponseselect.GetPostsResponseSelect `protobuf:"bytes,6,opt,name=select" json:"select,omitempty"`
	Hash             *string                                        `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                         `json:"-"`
}

func (m *GetPostsRequest) Reset()                    { *m = GetPostsRequest{} }
func (m *GetPostsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetPostsRequest) ProtoMessage()               {}
func (*GetPostsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPostsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPostsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *GetPostsRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetPostsRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *GetPostsRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetPostsRequest) GetSelect() *getpostsresponseselect.GetPostsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetPostsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPostsRequest)(nil), "getpostsrequest.GetPostsRequest")
}

func init() { proto.RegisterFile("getpostsrequest/getpostsrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xc1, 0x4a, 0x03, 0x31,
	0x14, 0xa4, 0xb5, 0xad, 0x18, 0x95, 0x42, 0xf0, 0x10, 0x8a, 0x48, 0xe9, 0x41, 0x7a, 0xd0, 0x04,
	0x8b, 0x37, 0xf1, 0x52, 0xa1, 0x5e, 0x14, 0x34, 0xbd, 0x79, 0x91, 0x18, 0x9f, 0xbb, 0x0b, 0xdb,
	0x66, 0x9b, 0x64, 0x85, 0x7e, 0xac, 0xff, 0x22, 0xfb, 0x36, 0xb5, 0x9b, 0xc5, 0x4b, 0x78, 0xf3,
	0x32, 0x43, 0x66, 0x86, 0x90, 0xcb, 0x04, 0x7c, 0x61, 0x9c, 0x77, 0x16, 0x36, 0x25, 0x38, 0x2f,
	0x5a, 0xf8, 0x9d, 0x17, 0xd6, 0x78, 0x43, 0x87, 0xad, 0xfd, 0xe8, 0x22, 0x0c, 0x2b, 0xf0, 0x4a,
	0x34, 0xe6, 0x20, 0x18, 0x4d, 0x1c, 0xd8, 0xef, 0x4c, 0xc3, 0x57, 0x96, 0x7b, 0xb0, 0x22, 0x42,
	0x3b, 0x0e, 0xd5, 0x4a, 0xa7, 0x20, 0xf0, 0xfc, 0xd3, 0xe5, 0x46, 0x2b, 0x9f, 0x99, 0xf5, 0xa6,
	0x04, 0xbb, 0x15, 0x11, 0xda, 0x71, 0x6e, 0xf7, 0x66, 0x5c, 0x61, 0xd6, 0x0e, 0x1c, 0xe4, 0xa0,
	0x9b, 0xde, 0x9b, 0xeb, 0xa0, 0x9a, 0xfc, 0x74, 0xc9, 0xf0, 0x11, 0xfc, 0x4b, 0xc5, 0x90, 0xb5,
	0x61, 0x7a, 0x45, 0x7a, 0x95, 0x69, 0xd6, 0x19, 0x77, 0xa6, 0xc7, 0x33, 0xc6, 0x1b, 0x41, 0x78,
	0xe0, 0x3c, 0x83, 0x57, 0x12, 0x59, 0x74, 0x4e, 0x4e, 0x43, 0x8e, 0x05, 0xe6, 0x60, 0x5d, 0x94,
	0x9d, 0xf3, 0x28, 0x1d, 0x5f, 0x36, 0x39, 0x32, 0x96, 0xd0, 0x09, 0xe9, 0x63, 0x5e, 0x76, 0x80,
	0xda, 0x13, 0x8e, 0x88, 0x3f, 0x54, 0xa7, 0xac, 0xaf, 0xe8, 0x19, 0xe9, 0x63, 0x5e, 0xd6, 0x1b,
	0x77, 0xa6, 0x47, 0xb2, 0x06, 0xd5, 0xeb, 0xbb, 0x36, 0x5e, 0xf1, 0xb6, 0x1f, 0x5e, 0x8f, 0x3a,
	0xe2, 0x4f, 0x4d, 0x8e, 0x8c, 0x25, 0x74, 0x41, 0x06, 0x75, 0x29, 0x6c, 0x80, 0x62, 0xce, 0xff,
	0xef, 0x8c, 0xef, 0x8b, 0xaa, 0xd7, 0x4b, 0x5c, 0xcb, 0xa0, 0xa6, 0x94, 0xf4, 0x52, 0xe5, 0x52,
	0x76, 0x88, 0x06, 0x71, 0x9e, 0xdf, 0xbf, 0xdd, 0x25, 0x99, 0x4f, 0xcb, 0x0f, 0xae, 0xcd, 0x4a,
	0xcc, 0x6e, 0x9c, 0xcf, 0x8c, 0x48, 0xcc, 0x75, 0x91, 0xab, 0x6d, 0x62, 0x4d, 0xb9, 0xfe, 0x14,
	0x89, 0x2d, 0xb4, 0x70, 0x3a, 0x85, 0x95, 0x6a, 0x7f, 0xb4, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7e, 0x01, 0xfb, 0xda, 0x8a, 0x02, 0x00, 0x00,
}