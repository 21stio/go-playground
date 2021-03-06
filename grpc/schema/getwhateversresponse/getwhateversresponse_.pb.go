// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getwhateversresponse/getwhateversresponse_.proto

/*
Package getwhateversresponse is a generated protocol buffer package.

It is generated from these files:
	getwhateversresponse/getwhateversresponse_.proto

It has these top-level messages:
	GetWhateversResponse
*/
package getwhateversresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import whatever "github.com/21stio/go-playground/grpc/schema/whatever"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetWhateversResponse struct {
	Meta             *responsemeta.ResponseMeta `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	TotalCount       *int32                     `protobuf:"varint,2,opt,name=totalCount" json:"totalCount,omitempty"`
	Whatevers        []*whatever.Whatever       `protobuf:"bytes,3,rep,name=whatevers" json:"whatevers,omitempty"`
	Hash             *string                    `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *GetWhateversResponse) Reset()                    { *m = GetWhateversResponse{} }
func (m *GetWhateversResponse) String() string            { return proto.CompactTextString(m) }
func (*GetWhateversResponse) ProtoMessage()               {}
func (*GetWhateversResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetWhateversResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetWhateversResponse) GetTotalCount() int32 {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return 0
}

func (m *GetWhateversResponse) GetWhatevers() []*whatever.Whatever {
	if m != nil {
		return m.Whatevers
	}
	return nil
}

func (m *GetWhateversResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWhateversResponse)(nil), "getwhateversresponse.GetWhateversResponse")
}

func init() { proto.RegisterFile("getwhateversresponse/getwhateversresponse_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x96, 0x69, 0x18, 0xea, 0x6e, 0x56, 0x07, 0x2b, 0x03, 0xb2, 0x98, 0xb2, 0x60, 0x97, 0x3c,
	0x01, 0x2a, 0x03, 0x13, 0x8b, 0x17, 0x24, 0x16, 0x64, 0xc2, 0xc9, 0xae, 0xd4, 0xe4, 0x2c, 0xfb,
	0x02, 0xe2, 0x9d, 0x78, 0x48, 0x44, 0x14, 0xd3, 0x0c, 0xd9, 0xbe, 0xfb, 0xfe, 0xee, 0x74, 0xfc,
	0xe0, 0x81, 0xbe, 0x82, 0x23, 0xf8, 0x84, 0x94, 0x13, 0xe4, 0x88, 0x43, 0x06, 0xb3, 0x46, 0xbe,
	0xe9, 0x98, 0x90, 0x50, 0xec, 0xd7, 0xc4, 0x5a, 0x15, 0xd4, 0x03, 0x39, 0xb3, 0x1c, 0xe6, 0x5c,
	0x2d, 0x4b, 0xc8, 0x14, 0x30, 0x2b, 0xb7, 0x3f, 0x8c, 0xef, 0x9f, 0x80, 0x5e, 0x4a, 0xa9, 0x9d,
	0xd3, 0x42, 0xf3, 0xea, 0xaf, 0x41, 0x32, 0xc5, 0x9a, 0x5d, 0x5b, 0xeb, 0x65, 0xad, 0x2e, 0xae,
	0x67, 0x20, 0x67, 0x27, 0x9f, 0xb8, 0xe1, 0x9c, 0x90, 0xdc, 0xf9, 0x11, 0xc7, 0x81, 0xe4, 0x95,
	0x62, 0xcd, 0xb5, 0x5d, 0x30, 0xe2, 0xc0, 0xb7, 0xff, 0x97, 0xcb, 0x8d, 0xda, 0x34, 0xbb, 0x56,
	0xe8, 0xc2, 0xe8, 0xb2, 0xdf, 0x5e, 0x4c, 0x42, 0xf0, 0x2a, 0xb8, 0x1c, 0x64, 0xa5, 0x58, 0xb3,
	0xb5, 0x13, 0x3e, 0x1e, 0x5f, 0x1f, 0xfc, 0x89, 0xc2, 0xf8, 0xae, 0x3b, 0xec, 0x4d, 0x7b, 0x9f,
	0xe9, 0x84, 0xc6, 0xe3, 0x5d, 0x3c, 0xbb, 0x6f, 0x9f, 0x70, 0x1c, 0x3e, 0x8c, 0x4f, 0xb1, 0x33,
	0xb9, 0x0b, 0xd0, 0xbb, 0xd5, 0x5f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x83, 0x97, 0x4c, 0x82,
	0x77, 0x01, 0x00, 0x00,
}
