// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getservicesrequestfilter/getservicesrequestfilter_.proto

/*
Package getservicesrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getservicesrequestfilter/getservicesrequestfilter_.proto

It has these top-level messages:
	GetServicesRequestFilter
*/
package getservicesrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
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

type GetServicesRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	LocationQuery    *locationqueryfilter.LocationQueryFilter `protobuf:"bytes,2,opt,name=locationQuery" json:"locationQuery,omitempty"`
	And              []*GetServicesRequestFilter              `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*GetServicesRequestFilter              `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*GetServicesRequestFilter              `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                                  `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *GetServicesRequestFilter) Reset()                    { *m = GetServicesRequestFilter{} }
func (m *GetServicesRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*GetServicesRequestFilter) ProtoMessage()               {}
func (*GetServicesRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetServicesRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetServicesRequestFilter) GetLocationQuery() *locationqueryfilter.LocationQueryFilter {
	if m != nil {
		return m.LocationQuery
	}
	return nil
}

func (m *GetServicesRequestFilter) GetAnd() []*GetServicesRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetServicesRequestFilter) GetOr() []*GetServicesRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetServicesRequestFilter) GetNot() []*GetServicesRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetServicesRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetServicesRequestFilter)(nil), "getservicesrequestfilter.GetServicesRequestFilter")
}

func init() {
	proto.RegisterFile("getservicesrequestfilter/getservicesrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0x1f, 0x05, 0x33, 0xbc, 0xe4, 0x14, 0x76, 0x2a, 0xe2, 0xa1, 0x08, 0xa6, 0xd8,
	0xd3, 0xce, 0x63, 0xcc, 0x8b, 0x0a, 0xd6, 0x9b, 0x17, 0x89, 0xdd, 0x6b, 0x5b, 0x68, 0xf3, 0x76,
	0xc9, 0x5b, 0x61, 0xdf, 0xc0, 0x8f, 0x2d, 0xed, 0x22, 0x38, 0xb6, 0x5c, 0x76, 0x4b, 0xde, 0x3c,
	0xcf, 0xef, 0x79, 0x12, 0xc2, 0x16, 0x15, 0x90, 0x05, 0xf3, 0xdd, 0x94, 0x60, 0x0d, 0x6c, 0x07,
	0xb0, 0xf4, 0xd5, 0xb4, 0x04, 0x26, 0xf3, 0x1d, 0x7c, 0xc8, 0xde, 0x20, 0x21, 0x17, 0x3e, 0xc1,
	0xfc, 0xce, 0x6d, 0x3b, 0x20, 0xe5, 0x60, 0x47, 0x13, 0x47, 0x99, 0xcb, 0x16, 0x4b, 0x45, 0x0d,
	0xea, 0xed, 0x00, 0x66, 0xe7, 0xd4, 0x27, 0x66, 0x4e, 0x7f, 0xf3, 0x13, 0x31, 0xf1, 0x08, 0xf4,
	0xe6, 0x82, 0x8b, 0x3d, 0x77, 0x3d, 0x69, 0xf8, 0x82, 0xc5, 0x63, 0x82, 0x08, 0x92, 0x20, 0x9d,
	0xe5, 0xb7, 0xf2, 0x28, 0x55, 0x3a, 0xfd, 0x33, 0x90, 0xda, 0x7b, 0x8a, 0xc9, 0xc1, 0x5f, 0xd8,
	0xf5, 0x5f, 0xe8, 0xeb, 0x18, 0x2a, 0xc2, 0x09, 0x91, 0x9e, 0xaa, 0x27, 0x9f, 0xfe, 0x2b, 0x1d,
	0xe6, 0xd0, 0xce, 0x57, 0x2c, 0x52, 0x7a, 0x23, 0xa2, 0x24, 0x4a, 0x67, 0x79, 0x2e, 0x7d, 0x4f,
	0x25, 0x7d, 0x57, 0x29, 0x46, 0x3b, 0x5f, 0xb2, 0x10, 0x8d, 0x88, 0xcf, 0x86, 0x84, 0x68, 0xc6,
	0x26, 0x1a, 0x49, 0x5c, 0x9c, 0xdf, 0x44, 0x23, 0x71, 0xce, 0xe2, 0x5a, 0xd9, 0x5a, 0x5c, 0x26,
	0x41, 0x7a, 0x55, 0x4c, 0xeb, 0xe5, 0xfa, 0x7d, 0x55, 0x35, 0x54, 0x0f, 0x9f, 0xb2, 0xc4, 0x2e,
	0xcb, 0x1f, 0x2c, 0x35, 0x98, 0x55, 0x78, 0xdf, 0xb7, 0x6a, 0x57, 0x19, 0x1c, 0xf4, 0x26, 0xab,
	0x4c, 0x5f, 0x66, 0xb6, 0xac, 0xa1, 0x53, 0xde, 0xff, 0xf4, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x48,
	0x41, 0x81, 0x6a, 0x83, 0x02, 0x00, 0x00,
}
