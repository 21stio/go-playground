// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamspecdescriptionsrequest/streamspecdescriptionsrequest_.proto

/*
Package streamspecdescriptionsrequest is a generated protocol buffer package.

It is generated from these files:
	streamspecdescriptionsrequest/streamspecdescriptionsrequest_.proto

It has these top-level messages:
	StreamSpecDescriptionsRequest
*/
package streamspecdescriptionsrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import streamkind "github.com/21stio/go-playground/grpc/schema/streamkind"
import durationscalar "github.com/21stio/go-playground/grpc/schema/durationscalar"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import locationquery "github.com/21stio/go-playground/grpc/schema/locationquery"
import streamspecdescriptionsresponseselect "github.com/21stio/go-playground/grpc/schema/streamspecdescriptionsresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamSpecDescriptionsRequest struct {
	Meta             *requestmeta.RequestMeta                                                   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkind.StreamKind                                                     `protobuf:"varint,2,opt,name=streamKind,enum=streamkind.StreamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalar.DurationScalar                                             `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                               `protobuf:"bytes,4,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Cache            *cache.Cache                                                               `protobuf:"bytes,5,opt,name=cache" json:"cache,omitempty"`
	Query            *string                                                                    `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
	LocationQuery    *locationquery.LocationQuery                                               `protobuf:"bytes,7,opt,name=locationQuery" json:"locationQuery,omitempty"`
	Select           *streamspecdescriptionsresponseselect.StreamSpecDescriptionsResponseSelect `protobuf:"bytes,8,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                                    `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                     `json:"-"`
}

func (m *StreamSpecDescriptionsRequest) Reset()                    { *m = StreamSpecDescriptionsRequest{} }
func (m *StreamSpecDescriptionsRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamSpecDescriptionsRequest) ProtoMessage()               {}
func (*StreamSpecDescriptionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamSpecDescriptionsRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamSpecDescriptionsRequest) GetStreamKind() streamkind.StreamKind {
	if m != nil && m.StreamKind != nil {
		return *m.StreamKind
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamSpecDescriptionsRequest) GetPollInterval() *durationscalar.DurationScalar {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamSpecDescriptionsRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *StreamSpecDescriptionsRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *StreamSpecDescriptionsRequest) GetQuery() string {
	if m != nil && m.Query != nil {
		return *m.Query
	}
	return ""
}

func (m *StreamSpecDescriptionsRequest) GetLocationQuery() *locationquery.LocationQuery {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamSpecDescriptionsRequest) GetSelect() *streamspecdescriptionsresponseselect.StreamSpecDescriptionsResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *StreamSpecDescriptionsRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamSpecDescriptionsRequest)(nil), "streamspecdescriptionsrequest.StreamSpecDescriptionsRequest")
}

func init() {
	proto.RegisterFile("streamspecdescriptionsrequest/streamspecdescriptionsrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x6f, 0x13, 0x31,
	0x10, 0x85, 0x15, 0x48, 0x0a, 0x35, 0x85, 0x83, 0x85, 0x90, 0x55, 0xb5, 0x55, 0x14, 0x71, 0xc8,
	0x01, 0x6c, 0x91, 0x03, 0x3f, 0x20, 0x54, 0x48, 0xa5, 0x20, 0x81, 0xf7, 0xc6, 0x05, 0xb9, 0xde,
	0x21, 0x6b, 0xe1, 0xac, 0x5d, 0xdb, 0x5b, 0xa9, 0x67, 0xfe, 0x38, 0xca, 0xac, 0xb7, 0x5d, 0x1f,
	0x12, 0x71, 0xb1, 0x66, 0x34, 0xdf, 0x9b, 0x49, 0xde, 0xd3, 0x92, 0x75, 0x4c, 0x01, 0xd4, 0x36,
	0x7a, 0xd0, 0x35, 0x44, 0x1d, 0x8c, 0x4f, 0xc6, 0xb5, 0x31, 0xc0, 0x6d, 0x07, 0x31, 0x89, 0x83,
	0xd3, 0x5f, 0xdc, 0x07, 0x97, 0x1c, 0x3d, 0x3f, 0x48, 0x9d, 0x5e, 0xe4, 0x62, 0x0b, 0x49, 0x89,
	0x51, 0x9d, 0xe5, 0xa7, 0x67, 0xbd, 0xfc, 0x8f, 0x69, 0x6b, 0xf1, 0x58, 0x0e, 0xd3, 0xb7, 0x75,
	0x17, 0x14, 0xee, 0xd3, 0xca, 0xaa, 0x20, 0xca, 0x76, 0xa0, 0x16, 0x11, 0xc2, 0x9d, 0xd1, 0xf0,
	0xdb, 0xd8, 0x04, 0x41, 0x14, 0xdd, 0xc0, 0x50, 0xad, 0x74, 0x03, 0x02, 0xdf, 0x07, 0x9d, 0x75,
	0x1a, 0xd7, 0xdd, 0x76, 0x10, 0xee, 0x45, 0xd1, 0x0d, 0xcc, 0xf7, 0x7d, 0x7f, 0x2f, 0x7a, 0xd7,
	0x46, 0x88, 0x60, 0x41, 0xef, 0x77, 0x6a, 0x0c, 0xe5, 0x8d, 0x8b, 0xbf, 0x53, 0x72, 0x5e, 0x21,
	0x5f, 0x79, 0xd0, 0x97, 0x23, 0x5e, 0xf6, 0xf6, 0xd0, 0x77, 0x64, 0xba, 0xb3, 0x88, 0x4d, 0xe6,
	0x93, 0xe5, 0x8b, 0x15, 0xe3, 0x23, 0xdb, 0x78, 0x66, 0xbe, 0x41, 0x52, 0x12, 0x29, 0xfa, 0x91,
	0x90, 0xfe, 0xfc, 0xb5, 0x69, 0x6b, 0xf6, 0x64, 0x3e, 0x59, 0xbe, 0x5a, 0xbd, 0xe1, 0x8f, 0x5e,
	0xf2, 0xea, 0x61, 0x2a, 0x47, 0x24, 0x5d, 0x93, 0x13, 0xef, 0xac, 0xbd, 0x6a, 0x13, 0x84, 0x3b,
	0x65, 0xd9, 0x53, 0xbc, 0x76, 0xc1, 0x4b, 0x8f, 0xf9, 0x65, 0x6e, 0x2b, 0x6c, 0x65, 0xa1, 0xa1,
	0x6b, 0xf2, 0x32, 0xbb, 0xfd, 0x19, 0xdd, 0x66, 0x53, 0x5c, 0x72, 0xc6, 0x8b, 0x0c, 0x78, 0x35,
	0x66, 0x64, 0x29, 0xa1, 0x0b, 0x32, 0xc3, 0x54, 0xd8, 0x0c, 0xb5, 0x27, 0x1c, 0x3b, 0xfe, 0x69,
	0xf7, 0xca, 0x7e, 0x44, 0x5f, 0x93, 0x19, 0xa6, 0xc2, 0x8e, 0xe6, 0x93, 0xe5, 0xb1, 0xec, 0x9b,
	0xdd, 0xf5, 0x21, 0xb3, 0x1f, 0x38, 0x7d, 0x96, 0xaf, 0x17, 0x49, 0xf2, 0xaf, 0x63, 0x46, 0x96,
	0x12, 0x7a, 0x43, 0x8e, 0xfa, 0x78, 0xd8, 0x73, 0x14, 0x7f, 0xe1, 0xff, 0x93, 0x25, 0xdf, 0x17,
	0x60, 0x0f, 0x55, 0x08, 0xc9, 0xbc, 0x99, 0x52, 0x32, 0x6d, 0x54, 0x6c, 0xd8, 0x31, 0xfe, 0x78,
	0xac, 0xd7, 0xd7, 0x3f, 0xaf, 0x36, 0x26, 0x35, 0xdd, 0x0d, 0xd7, 0x6e, 0x2b, 0x56, 0x1f, 0x62,
	0x32, 0x4e, 0x6c, 0xdc, 0x7b, 0x6f, 0xd5, 0xfd, 0x26, 0xb8, 0xae, 0xad, 0xc5, 0x26, 0x78, 0x2d,
	0xa2, 0x6e, 0x60, 0xab, 0x0e, 0x7f, 0x8a, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba, 0x29, 0x1d,
	0xbf, 0xc8, 0x03, 0x00, 0x00,
}
