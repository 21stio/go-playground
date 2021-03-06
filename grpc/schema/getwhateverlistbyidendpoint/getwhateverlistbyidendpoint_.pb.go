// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getwhateverlistbyidendpoint/getwhateverlistbyidendpoint_.proto

/*
Package getwhateverlistbyidendpoint is a generated protocol buffer package.

It is generated from these files:
	getwhateverlistbyidendpoint/getwhateverlistbyidendpoint_.proto

It has these top-level messages:
	GetWhateverListByIdEndpoint
*/
package getwhateverlistbyidendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import getwhateverlistbyidrequestfilter "github.com/21stio/go-playground/grpc/schema/getwhateverlistbyidrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetWhateverListByIdEndpoint struct {
	Filter           *getwhateverlistbyidrequestfilter.GetWhateverListByIdRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                            `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                             `json:"-"`
}

func (m *GetWhateverListByIdEndpoint) Reset()                    { *m = GetWhateverListByIdEndpoint{} }
func (m *GetWhateverListByIdEndpoint) String() string            { return proto.CompactTextString(m) }
func (*GetWhateverListByIdEndpoint) ProtoMessage()               {}
func (*GetWhateverListByIdEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetWhateverListByIdEndpoint) GetFilter() *getwhateverlistbyidrequestfilter.GetWhateverListByIdRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *GetWhateverListByIdEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetWhateverListByIdEndpoint)(nil), "getwhateverlistbyidendpoint.GetWhateverListByIdEndpoint")
}

func init() {
	proto.RegisterFile("getwhateverlistbyidendpoint/getwhateverlistbyidendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0x4f, 0x2d, 0x29,
	0xcf, 0x48, 0x2c, 0x49, 0x2d, 0x4b, 0x2d, 0xca, 0xc9, 0x2c, 0x2e, 0x49, 0xaa, 0xcc, 0x4c, 0x49,
	0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0xc7, 0x23, 0x17, 0xaf, 0x57, 0x50, 0x94, 0x5f,
	0x92, 0x2f, 0x24, 0x8d, 0x47, 0x8d, 0x94, 0x07, 0x16, 0xc9, 0xa2, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x92, 0xb4, 0xcc, 0x9c, 0x92, 0xd4, 0x22, 0x7d, 0x42, 0x0a, 0xa0, 0xd6, 0x28, 0xf5, 0x32, 0x72,
	0x49, 0xbb, 0xa7, 0x96, 0x84, 0x43, 0xd5, 0xfa, 0x64, 0x16, 0x97, 0x38, 0x55, 0x7a, 0xa6, 0xb8,
	0x42, 0x6d, 0x12, 0x8a, 0xe2, 0x62, 0x83, 0x68, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x72,
	0xd2, 0x23, 0x64, 0xb2, 0x1e, 0x16, 0xe3, 0x82, 0x20, 0x0a, 0xdc, 0xc0, 0x0a, 0x82, 0xa0, 0x26,
	0x0a, 0x09, 0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x4e, 0x9e, 0x51, 0xee, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0x46, 0x86, 0xc5, 0x25, 0x99, 0xf9, 0xfa, 0xe9, 0xf9, 0xba, 0x05, 0x39, 0x89, 0x95, 0xe9, 0x45,
	0xf9, 0xa5, 0x79, 0x29, 0xfa, 0xe9, 0x45, 0x05, 0xc9, 0xfa, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89,
	0xf8, 0x02, 0x12, 0x10, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x0f, 0x79, 0x4a, 0x82, 0x01, 0x00, 0x00,
}
