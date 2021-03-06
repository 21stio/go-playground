// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passcreatefeedsrequestrequest/passcreatefeedsrequestrequest_.proto

/*
Package passcreatefeedsrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passcreatefeedsrequestrequest/passcreatefeedsrequestrequest_.proto

It has these top-level messages:
	PassCreateFeedsRequestRequest
*/
package passcreatefeedsrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import createfeedsrequest "github.com/21stio/go-playground/grpc/schema/createfeedsrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassCreateFeedsRequestRequest struct {
	Meta               *requestmeta.RequestMeta               `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter      *servicefilter.ServiceFilter           `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	CreateFeedsRequest *createfeedsrequest.CreateFeedsRequest `protobuf:"bytes,3,opt,name=createFeedsRequest" json:"createFeedsRequest,omitempty"`
	Hash               *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized   []byte                                 `json:"-"`
}

func (m *PassCreateFeedsRequestRequest) Reset()                    { *m = PassCreateFeedsRequestRequest{} }
func (m *PassCreateFeedsRequestRequest) String() string            { return proto.CompactTextString(m) }
func (*PassCreateFeedsRequestRequest) ProtoMessage()               {}
func (*PassCreateFeedsRequestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PassCreateFeedsRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassCreateFeedsRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassCreateFeedsRequestRequest) GetCreateFeedsRequest() *createfeedsrequest.CreateFeedsRequest {
	if m != nil {
		return m.CreateFeedsRequest
	}
	return nil
}

func (m *PassCreateFeedsRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassCreateFeedsRequestRequest)(nil), "passcreatefeedsrequestrequest.PassCreateFeedsRequestRequest")
}

func init() {
	proto.RegisterFile("passcreatefeedsrequestrequest/passcreatefeedsrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0xf6, 0x62, 0xc4, 0x4b, 0x4e, 0x61, 0x38, 0x19, 0x3b, 0xc8, 0x0e, 0x33, 0xc1,
	0x7d, 0x84, 0x0a, 0x03, 0x11, 0x41, 0x2a, 0x78, 0xf0, 0x22, 0xcf, 0xec, 0xad, 0x2d, 0xac, 0xa6,
	0xe6, 0xbd, 0x0a, 0x7e, 0x78, 0x41, 0x16, 0x23, 0xb4, 0xb4, 0xf4, 0xf4, 0x5e, 0x9b, 0xdf, 0xff,
	0xcf, 0x2f, 0x44, 0x64, 0x0d, 0x10, 0x59, 0x8f, 0xc0, 0xb8, 0x47, 0xdc, 0x91, 0xc7, 0xcf, 0x16,
	0x89, 0xe3, 0x30, 0x93, 0xa7, 0x6f, 0xba, 0xf1, 0x8e, 0x9d, 0x9c, 0x4f, 0x52, 0xb3, 0xab, 0xb8,
	0xd4, 0xc8, 0x60, 0x3a, 0x7b, 0x8c, 0xcf, 0x96, 0x84, 0xfe, 0xab, 0xb2, 0xb8, 0xaf, 0x0e, 0x8c,
	0xde, 0xf4, 0xbe, 0xfe, 0x99, 0xf5, 0xb0, 0xde, 0x0c, 0x7f, 0x45, 0x7a, 0xf9, 0x93, 0x88, 0xf9,
	0x13, 0x10, 0xdd, 0x05, 0x62, 0x7b, 0x24, 0xf2, 0x3f, 0x22, 0x0e, 0xb9, 0x16, 0xe9, 0x51, 0x41,
	0x25, 0x8b, 0x64, 0x75, 0xbe, 0x51, 0xba, 0xa3, 0xa5, 0x23, 0xf3, 0x88, 0x0c, 0x79, 0xa0, 0x64,
	0x26, 0x2e, 0xa2, 0xd5, 0x36, 0x58, 0xa9, 0x93, 0x10, 0xbb, 0xd4, 0x3d, 0x57, 0xfd, 0xdc, 0x65,
	0xf2, 0x7e, 0x44, 0xbe, 0x08, 0x69, 0x07, 0x3a, 0xea, 0x34, 0x14, 0x5d, 0xeb, 0xe1, 0x5d, 0xf4,
	0x88, 0xfc, 0x48, 0x83, 0x94, 0x22, 0x2d, 0x81, 0x4a, 0x95, 0x2e, 0x92, 0xd5, 0x59, 0x1e, 0xf6,
	0xec, 0xe1, 0xf5, 0xbe, 0xa8, 0xb8, 0x6c, 0xdf, 0xb5, 0x75, 0xb5, 0xd9, 0xdc, 0x12, 0x57, 0xce,
	0x14, 0xee, 0xa6, 0x39, 0xc0, 0x77, 0xe1, 0x5d, 0xfb, 0xb1, 0x33, 0x85, 0x6f, 0xac, 0x21, 0x5b,
	0x62, 0x0d, 0xd3, 0x8f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x25, 0x3b, 0x3a, 0x2f, 0x22, 0x02,
	0x00, 0x00,
}
