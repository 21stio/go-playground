// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getjobbyidrequestfilter/getjobbyidrequestfilter_.proto

/*
Package getjobbyidrequestfilter is a generated protocol buffer package.

It is generated from these files:
	getjobbyidrequestfilter/getjobbyidrequestfilter_.proto

It has these top-level messages:
	GetJobByIdRequestFilter
*/
package getjobbyidrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetJobByIdRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Id               *idfilter.IdFilter                   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	And              []*GetJobByIdRequestFilter           `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*GetJobByIdRequestFilter           `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*GetJobByIdRequestFilter           `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *GetJobByIdRequestFilter) Reset()                    { *m = GetJobByIdRequestFilter{} }
func (m *GetJobByIdRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*GetJobByIdRequestFilter) ProtoMessage()               {}
func (*GetJobByIdRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetJobByIdRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetJobByIdRequestFilter) GetId() *idfilter.IdFilter {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetJobByIdRequestFilter) GetAnd() []*GetJobByIdRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *GetJobByIdRequestFilter) GetOr() []*GetJobByIdRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *GetJobByIdRequestFilter) GetNot() []*GetJobByIdRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *GetJobByIdRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetJobByIdRequestFilter)(nil), "getjobbyidrequestfilter.GetJobByIdRequestFilter")
}

func init() {
	proto.RegisterFile("getjobbyidrequestfilter/getjobbyidrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x4f, 0x84, 0x30,
	0x14, 0xc6, 0x43, 0x41, 0x13, 0x7b, 0x5b, 0x97, 0x6b, 0x6e, 0x22, 0x17, 0x07, 0x62, 0x62, 0x51,
	0x06, 0xe3, 0x68, 0x30, 0x6a, 0xce, 0xc4, 0x85, 0xd1, 0xc5, 0x14, 0x5a, 0xa1, 0xe6, 0xe0, 0x61,
	0x79, 0x0c, 0xfc, 0xa1, 0xfe, 0x3f, 0x06, 0xae, 0x37, 0x5d, 0xba, 0xdc, 0xf6, 0xfa, 0xbe, 0xef,
	0xfb, 0xf5, 0x6b, 0x52, 0xfa, 0x50, 0x6b, 0xfc, 0x81, 0xb2, 0x9c, 0x8c, 0xb2, 0xfa, 0x77, 0xd4,
	0x03, 0x7e, 0x9b, 0x3d, 0x6a, 0x9b, 0x7a, 0xf6, 0x5f, 0xa2, 0xb7, 0x80, 0xc0, 0xd6, 0x1e, 0x7d,
	0x73, 0xe3, 0x8e, 0xad, 0x46, 0xe9, 0x50, 0x27, 0x1b, 0x07, 0xd9, 0x70, 0xa3, 0x9c, 0xe5, 0x38,
	0x38, 0x65, 0xfb, 0x47, 0xe8, 0xfa, 0x4d, 0xe3, 0x3b, 0x94, 0xf9, 0xb4, 0x53, 0xc5, 0x01, 0xf0,
	0xba, 0x58, 0xd8, 0x23, 0x8d, 0x66, 0x14, 0x0f, 0xe2, 0x20, 0x59, 0x65, 0xd7, 0xe2, 0x04, 0x2f,
	0x9c, 0xff, 0x43, 0xa3, 0x3c, 0x64, 0x8a, 0x25, 0xc1, 0xb6, 0x94, 0x18, 0xc5, 0xc9, 0x92, 0x63,
	0xe2, 0x78, 0xa7, 0xd8, 0x29, 0xe7, 0x22, 0x46, 0xb1, 0x9c, 0x86, 0xb2, 0x53, 0x3c, 0x8c, 0xc3,
	0x64, 0x95, 0xdd, 0x09, 0xcf, 0x33, 0x85, 0xa7, 0x5c, 0x31, 0x87, 0xd9, 0x13, 0x25, 0x60, 0x79,
	0x74, 0x26, 0x82, 0x80, 0x9d, 0x5b, 0x74, 0x80, 0xfc, 0xe2, 0xdc, 0x16, 0x1d, 0x20, 0x63, 0x34,
	0x6a, 0xe4, 0xd0, 0xf0, 0xcb, 0x38, 0x48, 0xae, 0x8a, 0x65, 0xce, 0x5f, 0x3e, 0x9f, 0x6b, 0x83,
	0xcd, 0x58, 0x8a, 0x0a, 0xda, 0x34, 0xbb, 0x1f, 0xd0, 0x40, 0x5a, 0xc3, 0x6d, 0xbf, 0x97, 0x53,
	0x6d, 0x61, 0xec, 0x54, 0x5a, 0xdb, 0xbe, 0x4a, 0x87, 0xaa, 0xd1, 0xad, 0xf4, 0x7d, 0x82, 0xff,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x7f, 0x10, 0xc0, 0x36, 0x02, 0x00, 0x00,
}