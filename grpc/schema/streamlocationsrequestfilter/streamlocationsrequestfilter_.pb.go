// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamlocationsrequestfilter/streamlocationsrequestfilter_.proto

/*
Package streamlocationsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	streamlocationsrequestfilter/streamlocationsrequestfilter_.proto

It has these top-level messages:
	StreamLocationsRequestFilter
*/
package streamlocationsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import streamkindfilter "github.com/21stio/go-playground/grpc/schema/streamkindfilter"
import durationscalarfilter "github.com/21stio/go-playground/grpc/schema/durationscalarfilter"
import locationqueryfilter "github.com/21stio/go-playground/grpc/schema/locationqueryfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamLocationsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkindfilter.StreamKindFilter         `protobuf:"bytes,2,opt,name=streamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalarfilter.DurationScalarFilter `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	LocationQuery    *locationqueryfilter.LocationQueryFilter   `protobuf:"bytes,4,opt,name=locationQuery" json:"locationQuery,omitempty"`
	And              []*StreamLocationsRequestFilter            `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*StreamLocationsRequestFilter            `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*StreamLocationsRequestFilter            `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                                    `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *StreamLocationsRequestFilter) Reset()                    { *m = StreamLocationsRequestFilter{} }
func (m *StreamLocationsRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*StreamLocationsRequestFilter) ProtoMessage()               {}
func (*StreamLocationsRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamLocationsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamLocationsRequestFilter) GetStreamKind() *streamkindfilter.StreamKindFilter {
	if m != nil {
		return m.StreamKind
	}
	return nil
}

func (m *StreamLocationsRequestFilter) GetPollInterval() *durationscalarfilter.DurationScalarFilter {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamLocationsRequestFilter) GetLocationQuery() *locationqueryfilter.LocationQueryFilter {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamLocationsRequestFilter) GetAnd() []*StreamLocationsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *StreamLocationsRequestFilter) GetOr() []*StreamLocationsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *StreamLocationsRequestFilter) GetNot() []*StreamLocationsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *StreamLocationsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamLocationsRequestFilter)(nil), "streamlocationsrequestfilter.StreamLocationsRequestFilter")
}

func init() {
	proto.RegisterFile("streamlocationsrequestfilter/streamlocationsrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x3f, 0x4f, 0xfb, 0x30,
	0x14, 0x54, 0xdb, 0xfc, 0xfe, 0xe0, 0xc2, 0xe2, 0xc9, 0xaa, 0x3a, 0x54, 0x15, 0x43, 0x54, 0x09,
	0x07, 0x3a, 0x21, 0x26, 0x54, 0x21, 0x04, 0xa5, 0x54, 0x22, 0xdd, 0x58, 0x90, 0x49, 0x4c, 0x12,
	0xe1, 0xf8, 0xa5, 0x8e, 0x83, 0xd4, 0x6f, 0xc6, 0xc7, 0x43, 0x71, 0x1c, 0x68, 0x69, 0x94, 0xa5,
	0x5b, 0x72, 0xbe, 0xbb, 0xf7, 0xee, 0x64, 0xa3, 0xeb, 0x5c, 0x2b, 0xce, 0x52, 0x01, 0x01, 0xd3,
	0x09, 0xc8, 0x5c, 0xf1, 0x75, 0xc1, 0x73, 0xfd, 0x96, 0x08, 0xcd, 0x95, 0xd7, 0x76, 0xf8, 0x42,
	0x33, 0x05, 0x1a, 0xf0, 0xb0, 0x8d, 0x34, 0x98, 0xd8, 0xdf, 0x94, 0x6b, 0x66, 0x4d, 0xf7, 0x10,
	0xeb, 0x34, 0x70, 0x2b, 0xa7, 0xf7, 0x44, 0x86, 0x3b, 0xf3, 0x7f, 0x80, 0x9a, 0x79, 0x1e, 0x16,
	0xaa, 0x9a, 0x16, 0x30, 0xc1, 0x94, 0x65, 0x37, 0x81, 0xb5, 0x82, 0xd6, 0xfb, 0xad, 0x0b, 0xae,
	0x36, 0x56, 0xd0, 0x80, 0x59, 0xfe, 0xf8, 0xd3, 0x41, 0xc3, 0x95, 0x99, 0xbe, 0xa8, 0x83, 0xf9,
	0xd5, 0xde, 0xb7, 0x86, 0x87, 0x2f, 0x91, 0x53, 0x26, 0x20, 0x9d, 0x51, 0xc7, 0xed, 0x4f, 0x4f,
	0xe9, 0x5e, 0x2a, 0x6a, 0xf9, 0x8f, 0x5c, 0xb3, 0x4a, 0xe3, 0x1b, 0x05, 0x9e, 0x21, 0x54, 0xe5,
	0x7a, 0x48, 0x64, 0x48, 0xba, 0x46, 0x3f, 0xa6, 0xbf, 0xa3, 0xd2, 0xd5, 0x37, 0xc7, 0xaa, 0xb7,
	0x54, 0x78, 0x89, 0x8e, 0x33, 0x10, 0xe2, 0x5e, 0x6a, 0xae, 0x3e, 0x98, 0x20, 0x3d, 0xe3, 0x32,
	0xa1, 0x4d, 0x15, 0xd0, 0x1b, 0x0b, 0xae, 0x0c, 0x68, 0xdd, 0x76, 0xf4, 0x78, 0x89, 0x4e, 0xea,
	0x32, 0x9e, 0xca, 0x32, 0x88, 0x63, 0x0c, 0xdd, 0xa6, 0xda, 0xe8, 0x62, 0x9b, 0x69, 0xed, 0x76,
	0xe5, 0x78, 0x81, 0x7a, 0x4c, 0x86, 0xe4, 0xcf, 0xa8, 0xe7, 0xf6, 0xa7, 0x57, 0xb4, 0xed, 0x8a,
	0xd0, 0xb6, 0x9a, 0xfd, 0xd2, 0x06, 0xcf, 0x51, 0x17, 0x14, 0xf9, 0x7b, 0xb0, 0x59, 0x17, 0x54,
	0xb9, 0x99, 0x04, 0x4d, 0xfe, 0x1d, 0xbe, 0x99, 0x04, 0x8d, 0x31, 0x72, 0x62, 0x96, 0xc7, 0xe4,
	0xff, 0xa8, 0xe3, 0x1e, 0xf9, 0xe6, 0x7b, 0x36, 0x7f, 0xbe, 0x8b, 0x12, 0x1d, 0x17, 0xaf, 0x34,
	0x80, 0xd4, 0x9b, 0x5e, 0xe4, 0x3a, 0x01, 0x2f, 0x82, 0xb3, 0x4c, 0xb0, 0x4d, 0xa4, 0xa0, 0x90,
	0xa1, 0x17, 0xa9, 0x2c, 0xf0, 0xf2, 0x20, 0xe6, 0x29, 0x6b, 0x7d, 0x63, 0x5f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x07, 0xc0, 0x52, 0x6b, 0x9f, 0x03, 0x00, 0x00,
}
