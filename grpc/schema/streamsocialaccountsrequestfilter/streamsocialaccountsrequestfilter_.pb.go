// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamsocialaccountsrequestfilter/streamsocialaccountsrequestfilter_.proto

/*
Package streamsocialaccountsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	streamsocialaccountsrequestfilter/streamsocialaccountsrequestfilter_.proto

It has these top-level messages:
	StreamSocialAccountsRequestFilter
*/
package streamsocialaccountsrequestfilter

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

type StreamSocialAccountsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	StreamKind       *streamkindfilter.StreamKindFilter         `protobuf:"bytes,2,opt,name=streamKind" json:"streamKind,omitempty"`
	PollInterval     *durationscalarfilter.DurationScalarFilter `protobuf:"bytes,3,opt,name=pollInterval" json:"pollInterval,omitempty"`
	LocationQuery    *locationqueryfilter.LocationQueryFilter   `protobuf:"bytes,4,opt,name=locationQuery" json:"locationQuery,omitempty"`
	And              []*StreamSocialAccountsRequestFilter       `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*StreamSocialAccountsRequestFilter       `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*StreamSocialAccountsRequestFilter       `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                                    `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                     `json:"-"`
}

func (m *StreamSocialAccountsRequestFilter) Reset()         { *m = StreamSocialAccountsRequestFilter{} }
func (m *StreamSocialAccountsRequestFilter) String() string { return proto.CompactTextString(m) }
func (*StreamSocialAccountsRequestFilter) ProtoMessage()    {}
func (*StreamSocialAccountsRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *StreamSocialAccountsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *StreamSocialAccountsRequestFilter) GetStreamKind() *streamkindfilter.StreamKindFilter {
	if m != nil {
		return m.StreamKind
	}
	return nil
}

func (m *StreamSocialAccountsRequestFilter) GetPollInterval() *durationscalarfilter.DurationScalarFilter {
	if m != nil {
		return m.PollInterval
	}
	return nil
}

func (m *StreamSocialAccountsRequestFilter) GetLocationQuery() *locationqueryfilter.LocationQueryFilter {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *StreamSocialAccountsRequestFilter) GetAnd() []*StreamSocialAccountsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *StreamSocialAccountsRequestFilter) GetOr() []*StreamSocialAccountsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *StreamSocialAccountsRequestFilter) GetNot() []*StreamSocialAccountsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *StreamSocialAccountsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamSocialAccountsRequestFilter)(nil), "streamsocialaccountsrequestfilter.StreamSocialAccountsRequestFilter")
}

func init() {
	proto.RegisterFile("streamsocialaccountsrequestfilter/streamsocialaccountsrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xbd, 0x4e, 0xf3, 0x30,
	0x18, 0x85, 0xd5, 0x36, 0xdf, 0x9f, 0xfb, 0xb1, 0x78, 0xb2, 0x3a, 0xb5, 0x15, 0x43, 0x54, 0x09,
	0x07, 0x3a, 0xb1, 0x52, 0x55, 0x48, 0xfc, 0xb4, 0x88, 0x04, 0x31, 0xb0, 0x20, 0xe3, 0x98, 0x24,
	0xc2, 0xb1, 0x53, 0xdb, 0x41, 0xea, 0xb5, 0x72, 0x33, 0x28, 0x8e, 0x0b, 0x2d, 0x8d, 0x94, 0xa5,
	0x5b, 0x72, 0x72, 0xce, 0x93, 0xf7, 0xbc, 0xb2, 0xc1, 0xb5, 0x36, 0x8a, 0x91, 0x5c, 0x4b, 0x9a,
	0x11, 0x4e, 0x28, 0x95, 0xa5, 0x30, 0x5a, 0xb1, 0x55, 0xc9, 0xb4, 0x79, 0xcd, 0xb8, 0x61, 0x2a,
	0x68, 0x75, 0x3c, 0xe3, 0x42, 0x49, 0x23, 0xe1, 0xa8, 0xd5, 0x39, 0x98, 0xb8, 0xd7, 0x9c, 0x19,
	0xe2, 0xf0, 0x7b, 0x8a, 0xc3, 0x0d, 0xfc, 0x1a, 0xf7, 0x96, 0x89, 0x78, 0x67, 0x92, 0x6f, 0x61,
	0xe3, 0x3c, 0x8d, 0x4b, 0x45, 0x4c, 0x26, 0x85, 0xa6, 0x84, 0x13, 0xe5, 0xdc, 0x4d, 0xe2, 0x26,
	0x81, 0xb9, 0xa4, 0xf6, 0xe3, 0xaa, 0x64, 0x6a, 0xed, 0x02, 0x0d, 0x9a, 0xf3, 0x8f, 0x3f, 0x3c,
	0x30, 0x8a, 0xec, 0xdf, 0x23, 0xdb, 0xee, 0xc2, 0xb5, 0x0b, 0xeb, 0xe1, 0x2f, 0xad, 0x19, 0x9e,
	0x03, 0xaf, 0xaa, 0x81, 0x3a, 0xc3, 0x8e, 0xdf, 0x9f, 0x1e, 0xe3, 0xbd, 0x6a, 0xd8, 0xf9, 0x17,
	0xcc, 0x90, 0x3a, 0x13, 0xda, 0x04, 0x9c, 0x01, 0x50, 0x97, 0xbb, 0xc9, 0x44, 0x8c, 0xba, 0x36,
	0x3f, 0xc6, 0x3f, 0xfb, 0xe2, 0xe8, 0xcb, 0xe3, 0xd2, 0x5b, 0x29, 0xb8, 0x04, 0xff, 0x0b, 0xc9,
	0xf9, 0x95, 0x30, 0x4c, 0xbd, 0x13, 0x8e, 0x7a, 0x96, 0x32, 0xc1, 0x4d, 0x7b, 0xc0, 0x73, 0x27,
	0x46, 0x56, 0x74, 0xb4, 0x9d, 0x3c, 0x5c, 0x82, 0xa3, 0xcd, 0x46, 0xee, 0xab, 0x8d, 0x20, 0xcf,
	0x02, 0xfd, 0xa6, 0xdd, 0xe1, 0xdb, 0x6d, 0xa7, 0xc3, 0xed, 0xc6, 0xe1, 0x23, 0xe8, 0x11, 0x11,
	0xa3, 0x5f, 0xc3, 0x9e, 0xdf, 0x9f, 0xce, 0x71, 0xeb, 0x61, 0xc1, 0xad, 0x0b, 0x0f, 0x2b, 0x20,
	0x7c, 0x00, 0x5d, 0xa9, 0xd0, 0xef, 0x03, 0x62, 0xbb, 0x52, 0x55, 0xd3, 0x0a, 0x69, 0xd0, 0x9f,
	0x43, 0x4e, 0x2b, 0xa4, 0x81, 0x10, 0x78, 0x29, 0xd1, 0x29, 0xfa, 0x3b, 0xec, 0xf8, 0xff, 0x42,
	0xfb, 0x3c, 0xbb, 0x7b, 0x5a, 0x24, 0x99, 0x49, 0xcb, 0x17, 0x4c, 0x65, 0x1e, 0x4c, 0xcf, 0xb4,
	0xc9, 0x64, 0x90, 0xc8, 0x93, 0x82, 0x93, 0x75, 0xa2, 0x64, 0x29, 0xe2, 0x20, 0x51, 0x05, 0x0d,
	0x34, 0x4d, 0x59, 0x4e, 0xda, 0x2f, 0xe4, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xbc, 0x85,
	0xab, 0xd6, 0x03, 0x00, 0x00,
}
