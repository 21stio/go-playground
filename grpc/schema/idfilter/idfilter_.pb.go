// Code generated by protoc-gen-go. DO NOT EDIT.
// source: idfilter/idfilter_.proto

/*
Package idfilter is a generated protocol buffer package.

It is generated from these files:
	idfilter/idfilter_.proto

It has these top-level messages:
	IdFilter
*/
package idfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import idkindfilter "github.com/21stio/go-playground/grpc/schema/idkindfilter"
import urlfilter "github.com/21stio/go-playground/grpc/schema/urlfilter"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"
import serviceidfilter "github.com/21stio/go-playground/grpc/schema/serviceidfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IdFilter struct {
	Kind             *idkindfilter.IdKindFilter       `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Url              *urlfilter.UrlFilter             `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Name             *stringfilter.StringFilter       `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Username         *stringfilter.StringFilter       `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	Ean              *stringfilter.StringFilter       `protobuf:"bytes,5,opt,name=ean" json:"ean,omitempty"`
	ServiceId        *serviceidfilter.ServiceIdFilter `protobuf:"bytes,6,opt,name=serviceId" json:"serviceId,omitempty"`
	Local            *stringfilter.StringFilter       `protobuf:"bytes,7,opt,name=local" json:"local,omitempty"`
	And              []*IdFilter                      `protobuf:"bytes,8,rep,name=and" json:"and,omitempty"`
	Or               []*IdFilter                      `protobuf:"bytes,9,rep,name=or" json:"or,omitempty"`
	Not              []*IdFilter                      `protobuf:"bytes,10,rep,name=not" json:"not,omitempty"`
	Hash             *string                          `protobuf:"bytes,11,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *IdFilter) Reset()                    { *m = IdFilter{} }
func (m *IdFilter) String() string            { return proto.CompactTextString(m) }
func (*IdFilter) ProtoMessage()               {}
func (*IdFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IdFilter) GetKind() *idkindfilter.IdKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *IdFilter) GetUrl() *urlfilter.UrlFilter {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *IdFilter) GetName() *stringfilter.StringFilter {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *IdFilter) GetUsername() *stringfilter.StringFilter {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *IdFilter) GetEan() *stringfilter.StringFilter {
	if m != nil {
		return m.Ean
	}
	return nil
}

func (m *IdFilter) GetServiceId() *serviceidfilter.ServiceIdFilter {
	if m != nil {
		return m.ServiceId
	}
	return nil
}

func (m *IdFilter) GetLocal() *stringfilter.StringFilter {
	if m != nil {
		return m.Local
	}
	return nil
}

func (m *IdFilter) GetAnd() []*IdFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *IdFilter) GetOr() []*IdFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *IdFilter) GetNot() []*IdFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *IdFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*IdFilter)(nil), "idfilter.IdFilter")
}

func init() { proto.RegisterFile("idfilter/idfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 338 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xd9, 0xba, 0x69, 0x97, 0xdd, 0x82, 0x87, 0xd0, 0x53, 0x19, 0x32, 0x76, 0xd0, 0x54,
	0x87, 0xec, 0xe8, 0xc1, 0x83, 0x30, 0xbc, 0x6d, 0x78, 0xf1, 0x22, 0xb1, 0x89, 0x6d, 0x30, 0x4b,
	0xc6, 0x6b, 0x2a, 0xf8, 0x4f, 0xfb, 0x37, 0x48, 0x62, 0xd2, 0x8e, 0xc1, 0xd8, 0xed, 0x7d, 0x79,
	0xbf, 0xef, 0xeb, 0xc7, 0xa3, 0x88, 0x48, 0xfe, 0x29, 0x95, 0x15, 0x50, 0xc4, 0xe1, 0x9d, 0xee,
	0xc1, 0x58, 0x83, 0xd3, 0xf8, 0x90, 0xe5, 0x92, 0x7f, 0x49, 0xdd, 0x73, 0xbd, 0x08, 0x6c, 0x96,
	0xb5, 0xa0, 0xc2, 0xba, 0x9b, 0xe2, 0x2e, 0x6f, 0x2c, 0x48, 0x5d, 0x85, 0xf5, 0xa1, 0x88, 0xc4,
	0xbc, 0x11, 0xf0, 0x2d, 0x4b, 0xd1, 0x55, 0x39, 0xd2, 0x81, 0x9b, 0xfd, 0x26, 0x28, 0x5d, 0xf3,
	0x67, 0xff, 0x86, 0x29, 0x1a, 0xb9, 0x1e, 0x64, 0x90, 0x0f, 0x16, 0xd3, 0x65, 0x46, 0x0f, 0x6b,
	0xd1, 0x35, 0x7f, 0x91, 0x3a, 0x90, 0x1b, 0xcf, 0xe1, 0x39, 0x4a, 0x5a, 0x50, 0x64, 0xe8, 0xf1,
	0x2b, 0xda, 0xd5, 0xa4, 0xaf, 0xa0, 0x02, 0xe8, 0x00, 0x97, 0xab, 0xd9, 0x4e, 0x90, 0x24, 0xe4,
	0x1e, 0x16, 0xa6, 0x5b, 0x2f, 0x62, 0xae, 0xe3, 0xf0, 0x0a, 0xa5, 0x6d, 0x23, 0xc0, 0x7b, 0x46,
	0x67, 0x3d, 0x1d, 0x8b, 0x6f, 0x50, 0x22, 0x98, 0x26, 0xe3, 0xb3, 0x16, 0x87, 0xe1, 0x47, 0x34,
	0x09, 0x47, 0x59, 0x73, 0x72, 0xe1, 0x3d, 0x39, 0x3d, 0x3a, 0x13, 0xdd, 0x46, 0x22, 0x38, 0x7b,
	0x0b, 0xbe, 0x43, 0x63, 0x65, 0x4a, 0xa6, 0xc8, 0xe5, 0xd9, 0xef, 0xfd, 0x83, 0xf8, 0x1a, 0x25,
	0x4c, 0x73, 0x92, 0xe6, 0xc9, 0x62, 0xba, 0xc4, 0x54, 0xf6, 0xa7, 0x8d, 0xbd, 0x98, 0xe6, 0x78,
	0x86, 0x86, 0x06, 0xc8, 0xe4, 0x24, 0x34, 0x34, 0xe0, 0x92, 0xb4, 0xb1, 0x04, 0x9d, 0x4e, 0xd2,
	0xc6, 0x62, 0x8c, 0x46, 0x35, 0x6b, 0x6a, 0x32, 0xcd, 0x07, 0x8b, 0xc9, 0xc6, 0xcf, 0x4f, 0xab,
	0xb7, 0x87, 0x4a, 0xda, 0xba, 0xfd, 0xa0, 0xa5, 0xd9, 0x15, 0xcb, 0xfb, 0xc6, 0x4a, 0x53, 0x54,
	0xe6, 0x76, 0xaf, 0xd8, 0x4f, 0x05, 0xa6, 0xd5, 0xbc, 0xa8, 0x60, 0x5f, 0x16, 0x4d, 0x59, 0x8b,
	0x1d, 0xeb, 0xfe, 0xe0, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14, 0xaf, 0x01, 0x14, 0xd5, 0x02,
	0x00, 0x00,
}
