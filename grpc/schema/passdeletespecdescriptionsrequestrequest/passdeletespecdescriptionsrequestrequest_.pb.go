// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletespecdescriptionsrequestrequest/passdeletespecdescriptionsrequestrequest_.proto

/*
Package passdeletespecdescriptionsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passdeletespecdescriptionsrequestrequest/passdeletespecdescriptionsrequestrequest_.proto

It has these top-level messages:
	PassDeleteSpecDescriptionsRequestRequest
*/
package passdeletespecdescriptionsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import deletespecdescriptionsrequest "github.com/21stio/go-playground/grpc/schema/deletespecdescriptionsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteSpecDescriptionsRequestRequest struct {
	Meta                          *requestmeta.RequestMeta                                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter                 *servicefilter.ServiceFilter                                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	DeleteSpecDescriptionsRequest *deletespecdescriptionsrequest.DeleteSpecDescriptionsRequest `protobuf:"bytes,3,opt,name=deleteSpecDescriptionsRequest" json:"deleteSpecDescriptionsRequest,omitempty"`
	Hash                          *string                                                      `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized              []byte                                                       `json:"-"`
}

func (m *PassDeleteSpecDescriptionsRequestRequest) Reset() {
	*m = PassDeleteSpecDescriptionsRequestRequest{}
}
func (m *PassDeleteSpecDescriptionsRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassDeleteSpecDescriptionsRequestRequest) ProtoMessage()    {}
func (*PassDeleteSpecDescriptionsRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteSpecDescriptionsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteSpecDescriptionsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassDeleteSpecDescriptionsRequestRequest) GetDeleteSpecDescriptionsRequest() *deletespecdescriptionsrequest.DeleteSpecDescriptionsRequest {
	if m != nil {
		return m.DeleteSpecDescriptionsRequest
	}
	return nil
}

func (m *PassDeleteSpecDescriptionsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteSpecDescriptionsRequestRequest)(nil), "passdeletespecdescriptionsrequestrequest.PassDeleteSpecDescriptionsRequestRequest")
}

func init() {
	proto.RegisterFile("passdeletespecdescriptionsrequestrequest/passdeletespecdescriptionsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0xe9, 0xd9, 0xc5, 0x88, 0x4b, 0xa6, 0x70, 0xa8, 0x1c, 0x37, 0x75, 0xd0, 0x04, 0x6f,
	0x76, 0x2a, 0x87, 0x9b, 0x20, 0x3d, 0x11, 0x71, 0x91, 0x98, 0x3e, 0xdb, 0x40, 0xdb, 0xc4, 0xbc,
	0x54, 0x70, 0xf5, 0x7f, 0xf2, 0xff, 0x13, 0x63, 0x84, 0x76, 0x09, 0x9d, 0xde, 0x2b, 0xfd, 0x7d,
	0xdf, 0xfb, 0x3e, 0x08, 0x79, 0xb2, 0x12, 0xb1, 0x86, 0x0e, 0x3c, 0xa0, 0x05, 0x55, 0x03, 0x2a,
	0xa7, 0xad, 0xd7, 0x66, 0x40, 0x07, 0xef, 0x23, 0xa0, 0x8f, 0x43, 0x2c, 0x05, 0x5f, 0xb8, 0x75,
	0xc6, 0x1b, 0x5a, 0x2c, 0x15, 0xac, 0x2f, 0xe2, 0xd2, 0x83, 0x97, 0x62, 0xb2, 0x47, 0xa7, 0xf5,
	0x16, 0xc1, 0x7d, 0x68, 0x05, 0x6f, 0xba, 0xf3, 0xe0, 0xc4, 0xec, 0xeb, 0x9f, 0x29, 0x93, 0x97,
	0x44, 0xf2, 0x6f, 0xf4, 0xd8, 0x7e, 0xaf, 0x48, 0x71, 0x2f, 0x11, 0xf7, 0x01, 0x3e, 0x58, 0x50,
	0xfb, 0x09, 0x5c, 0xfd, 0xc1, 0x71, 0xd0, 0x4b, 0x92, 0xff, 0x66, 0x64, 0xd9, 0x26, 0x2b, 0x4e,
	0x76, 0x8c, 0x4f, 0x72, 0xf3, 0xc8, 0xdc, 0x81, 0x97, 0x55, 0xa0, 0x68, 0x49, 0x4e, 0x63, 0xec,
	0xdb, 0x10, 0x9b, 0xad, 0x82, 0xec, 0x8c, 0xcf, 0xca, 0xf0, 0xc3, 0x94, 0xa9, 0xe6, 0x12, 0xfa,
	0x95, 0x91, 0xf3, 0x3a, 0x15, 0x8d, 0x1d, 0x05, 0xd3, 0x1b, 0x9e, 0x6c, 0xcb, 0xd3, 0xf5, 0xd2,
	0x27, 0x28, 0x25, 0x79, 0x2b, 0xb1, 0x65, 0xf9, 0x26, 0x2b, 0x8e, 0xab, 0xb0, 0x97, 0x8f, 0xcf,
	0x0f, 0x8d, 0xf6, 0xed, 0xf8, 0xca, 0x95, 0xe9, 0xc5, 0xee, 0x1a, 0xbd, 0x36, 0xa2, 0x31, 0x57,
	0xb6, 0x93, 0x9f, 0x8d, 0x33, 0xe3, 0x50, 0x8b, 0xc6, 0x59, 0x25, 0x50, 0xb5, 0xd0, 0xcb, 0xc5,
	0x0f, 0xe9, 0x27, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x9e, 0xf0, 0x49, 0x9c, 0x02, 0x00, 0x00,
}