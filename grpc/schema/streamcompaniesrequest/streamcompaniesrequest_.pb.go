// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamcompaniesrequest/streamcompaniesrequest_.proto

/*
Package streamcompaniesrequest is a generated protocol buffer package.

It is generated from these files:
	streamcompaniesrequest/streamcompaniesrequest_.proto

It has these top-level messages:
	StreamCompaniesRequest
*/
package streamcompaniesrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import streamkind "github.com/21stio/go-playground/grpc/schema/streamkind"
import durationscalar "github.com/21stio/go-playground/grpc/schema/durationscalar"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import streamcompaniesresponseselect "github.com/21stio/go-playground/grpc/schema/streamcompaniesresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamCompaniesRequest struct {
	Meta             *requestmeta.RequestMeta                                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkind.StreamKind                                       `protobuf:"varint,2,opt,name=streamKind,enum=streamkind.StreamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalar.DurationScalar                               `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                 `protobuf:"bytes,4,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                                 `protobuf:"bytes,5,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                      `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                                 `protobuf:"bytes,7,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *streamcompaniesresponseselect.StreamCompaniesResponseSelect `protobuf:"bytes,8,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                      `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                       `json:"-"`
}

func (m *StreamCompaniesRequest) Reset()                    { *m = StreamCompaniesRequest{} }
func (m *StreamCompaniesRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamCompaniesRequest) ProtoMessage()               {}
func (*StreamCompaniesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamCompaniesRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamCompaniesRequest) GetStreamKind() streamkind.StreamKind {
	if m != nil && m.StreamKind != nil {
		return *m.StreamKind
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamCompaniesRequest) GetPollInterval() *durationscalar.DurationScalar {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamCompaniesRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *StreamCompaniesRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *StreamCompaniesRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *StreamCompaniesRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamCompaniesRequest) GetSelect() *streamcompaniesresponseselect.StreamCompaniesResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StreamCompaniesRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamCompaniesRequest)(nil), "streamcompaniesrequest.StreamCompaniesRequest")
}

func init() {
	proto.RegisterFile("streamcompaniesrequest/streamcompaniesrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0x13, 0x31,
	0x14, 0x54, 0x68, 0x52, 0xa8, 0x29, 0x1c, 0x2c, 0x14, 0x59, 0x55, 0x55, 0x45, 0x11, 0x87, 0x1c,
	0xc0, 0x16, 0x11, 0xe2, 0xc4, 0x69, 0x5b, 0x21, 0x21, 0xe0, 0x80, 0x97, 0x13, 0x17, 0x64, 0x9c,
	0xc7, 0xae, 0x85, 0x77, 0xbd, 0xb5, 0xbd, 0x95, 0xfa, 0xd3, 0x7c, 0x03, 0xca, 0x5b, 0x6f, 0xbb,
	0x56, 0x03, 0x17, 0xeb, 0x4d, 0x66, 0xe6, 0xbd, 0x64, 0x46, 0x21, 0x6f, 0x43, 0xf4, 0xa0, 0x1a,
	0xed, 0x9a, 0x4e, 0xb5, 0x06, 0x82, 0x87, 0xeb, 0x1e, 0x42, 0x14, 0x87, 0x3f, 0xfe, 0xc1, 0x3b,
	0xef, 0xa2, 0xa3, 0xcb, 0xc3, 0xf4, 0xd9, 0x45, 0x1a, 0x1a, 0x88, 0x4a, 0x4c, 0xe6, 0xe4, 0x3b,
	0x3b, 0x1f, 0x7c, 0xbf, 0x4d, 0xbb, 0x13, 0xf7, 0xe3, 0xc8, 0xbe, 0xdc, 0xf5, 0x5e, 0x45, 0xe3,
	0xda, 0xa0, 0x95, 0x55, 0x5e, 0xe4, 0x70, 0x54, 0xad, 0x03, 0xf8, 0x1b, 0xa3, 0xe1, 0x97, 0xb1,
	0x11, 0xbc, 0xc8, 0xd0, 0xa8, 0xa1, 0x5a, 0xe9, 0x1a, 0x04, 0xbe, 0x77, 0x3e, 0xeb, 0x34, 0xae,
	0xbb, 0xee, 0xc1, 0xdf, 0x8a, 0x0c, 0x8d, 0x9a, 0xe2, 0xc1, 0xef, 0x0a, 0x9d, 0x6b, 0x03, 0x04,
	0xb0, 0xa0, 0x0f, 0x84, 0x32, 0x65, 0xd3, 0x8e, 0xf5, 0x9f, 0x23, 0xb2, 0x2c, 0x51, 0x78, 0x39,
	0x0a, 0xe5, 0x90, 0x04, 0x7d, 0x45, 0xe6, 0xfb, 0x34, 0xd8, 0x6c, 0x35, 0xdb, 0x3c, 0xdd, 0x32,
	0x3e, 0x49, 0x88, 0x27, 0xcd, 0x17, 0x88, 0x4a, 0xa2, 0x8a, 0xbe, 0x23, 0x64, 0x38, 0xf8, 0xc9,
	0xb4, 0x3b, 0xf6, 0x68, 0x35, 0xdb, 0x3c, 0xdf, 0x2e, 0xf9, 0x7d, 0x6c, 0xbc, 0xbc, 0x63, 0xe5,
	0x44, 0x49, 0x0b, 0x72, 0xda, 0x39, 0x6b, 0x3f, 0xb6, 0x11, 0xfc, 0x8d, 0xb2, 0xec, 0x08, 0xaf,
	0x5d, 0xf0, 0x3c, 0x4e, 0x7e, 0x95, 0x60, 0x89, 0x50, 0x66, 0x1e, 0x5a, 0x90, 0x67, 0x29, 0xd8,
	0x0f, 0x18, 0x2c, 0x9b, 0xe3, 0x92, 0x73, 0x9e, 0xc5, 0xcd, 0xcb, 0xa9, 0x46, 0xe6, 0x16, 0xba,
	0x26, 0x0b, 0x2c, 0x80, 0x2d, 0xd0, 0x7b, 0xca, 0x11, 0xf1, 0xcb, 0xfd, 0x2b, 0x07, 0x8a, 0xbe,
	0x20, 0x0b, 0x2c, 0x80, 0x1d, 0xaf, 0x66, 0x9b, 0x13, 0x39, 0x80, 0xfd, 0xf5, 0xb1, 0x9e, 0xaf,
	0xc8, 0x3e, 0x4e, 0xd7, 0xb3, 0xd2, 0xf8, 0xe7, 0xa9, 0x46, 0xe6, 0x16, 0xfa, 0x8d, 0x1c, 0x0f,
	0xbd, 0xb0, 0x27, 0x68, 0x7e, 0xcf, 0xff, 0xdb, 0x1e, 0x7f, 0x50, 0xd9, 0xc0, 0x96, 0xc8, 0xca,
	0xb4, 0x8b, 0x52, 0x32, 0xaf, 0x55, 0xa8, 0xd9, 0x09, 0x7e, 0x5d, 0x9c, 0x8b, 0xab, 0xef, 0x45,
	0x65, 0x62, 0xdd, 0xff, 0xe4, 0xda, 0x35, 0x62, 0xfb, 0x26, 0x44, 0xe3, 0x44, 0xe5, 0x5e, 0x77,
	0x56, 0xdd, 0x56, 0xde, 0xf5, 0xed, 0x4e, 0x54, 0xbe, 0xd3, 0x22, 0xe8, 0x1a, 0x1a, 0xf5, 0x8f,
	0x7f, 0xd6, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x9d, 0xfc, 0xdf, 0x89, 0x03, 0x00, 0x00,
}
