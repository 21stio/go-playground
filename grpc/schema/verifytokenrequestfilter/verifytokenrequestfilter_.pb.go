// Code generated by protoc-gen-go. DO NOT EDIT.
// source: verifytokenrequestfilter/verifytokenrequestfilter_.proto

/*
Package verifytokenrequestfilter is a generated protocol buffer package.

It is generated from these files:
	verifytokenrequestfilter/verifytokenrequestfilter_.proto

It has these top-level messages:
	VerifyTokenRequestFilter
*/
package verifytokenrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import tokenfilter "github.com/21stio/go-playground/grpc/schema/tokenfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VerifyTokenRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Token            *tokenfilter.TokenFilter             `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	And              []*VerifyTokenRequestFilter          `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*VerifyTokenRequestFilter          `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*VerifyTokenRequestFilter          `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *VerifyTokenRequestFilter) Reset()                    { *m = VerifyTokenRequestFilter{} }
func (m *VerifyTokenRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*VerifyTokenRequestFilter) ProtoMessage()               {}
func (*VerifyTokenRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VerifyTokenRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *VerifyTokenRequestFilter) GetToken() *tokenfilter.TokenFilter {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *VerifyTokenRequestFilter) GetAnd() []*VerifyTokenRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *VerifyTokenRequestFilter) GetOr() []*VerifyTokenRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *VerifyTokenRequestFilter) GetNot() []*VerifyTokenRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *VerifyTokenRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifyTokenRequestFilter)(nil), "verifytokenrequestfilter.VerifyTokenRequestFilter")
}

func init() {
	proto.RegisterFile("verifytokenrequestfilter/verifytokenrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x6a, 0xc3, 0x30,
	0x10, 0xc7, 0xf1, 0x47, 0x0a, 0x55, 0x36, 0x4d, 0x22, 0x43, 0x31, 0xa5, 0x83, 0x29, 0xf4, 0x44,
	0x3d, 0x65, 0x0e, 0x21, 0x5b, 0x17, 0x53, 0x3a, 0x74, 0x29, 0xaa, 0xa3, 0xd8, 0xa6, 0xb1, 0xce,
	0x95, 0xcf, 0x85, 0xbc, 0x71, 0x1f, 0x23, 0x58, 0xd6, 0x10, 0x08, 0x5a, 0xb2, 0x9d, 0x74, 0xbf,
	0xff, 0x87, 0x40, 0x6c, 0xfd, 0xa7, 0x6d, 0x7b, 0x38, 0x11, 0xfe, 0x68, 0x63, 0xf5, 0xef, 0xa8,
	0x07, 0x3a, 0xb4, 0x47, 0xd2, 0x56, 0x86, 0x16, 0x5f, 0xd0, 0x5b, 0x24, 0xe4, 0x22, 0x04, 0xac,
	0x9e, 0xfd, 0xb1, 0xd3, 0xa4, 0xbc, 0xd9, 0xd5, 0x8d, 0x77, 0x59, 0x3d, 0x38, 0xbd, 0xa7, 0x2e,
	0x66, 0xbf, 0x7f, 0xfc, 0x8f, 0x99, 0xf8, 0x70, 0x41, 0xef, 0xd3, 0xb2, 0x9c, 0x7d, 0x76, 0x8e,
	0xe1, 0x6b, 0x96, 0x4e, 0x8e, 0x22, 0xca, 0xa2, 0x7c, 0x59, 0x3c, 0xc1, 0x55, 0x0a, 0x78, 0xfe,
	0x4d, 0x93, 0x9a, 0x35, 0xa5, 0x53, 0x70, 0x60, 0x0b, 0x17, 0x26, 0x62, 0x27, 0x15, 0x70, 0x11,
	0x0d, 0x2e, 0xc9, 0xe3, 0x33, 0xc6, 0xb7, 0x2c, 0x51, 0x66, 0x2f, 0x92, 0x2c, 0xc9, 0x97, 0x45,
	0x01, 0xa1, 0xa7, 0x43, 0xa8, 0x6a, 0x39, 0xc9, 0xf9, 0x86, 0xc5, 0x68, 0x45, 0x7a, 0xb3, 0x49,
	0x8c, 0x76, 0x6a, 0x62, 0x90, 0xc4, 0xe2, 0xf6, 0x26, 0x06, 0x89, 0x73, 0x96, 0x36, 0x6a, 0x68,
	0xc4, 0x5d, 0x16, 0xe5, 0xf7, 0xa5, 0x9b, 0x37, 0xbb, 0xcf, 0x6d, 0xdd, 0x52, 0x33, 0x7e, 0x43,
	0x85, 0x9d, 0x2c, 0x5e, 0x07, 0x6a, 0x51, 0xd6, 0xf8, 0xd2, 0x1f, 0xd5, 0xa9, 0xb6, 0x38, 0x9a,
	0xbd, 0xac, 0x6d, 0x5f, 0xc9, 0xa1, 0x6a, 0x74, 0xa7, 0x82, 0xff, 0xe3, 0x1c, 0x00, 0x00, 0xff,
	0xff, 0x30, 0x67, 0xdb, 0xff, 0x53, 0x02, 0x00, 0x00,
}
