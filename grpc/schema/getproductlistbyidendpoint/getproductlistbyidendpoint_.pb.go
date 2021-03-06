// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getproductlistbyidendpoint/getproductlistbyidendpoint_.proto

/*
Package getproductlistbyidendpoint is a generated protocol buffer package.

It is generated from these files:
	getproductlistbyidendpoint/getproductlistbyidendpoint_.proto

It has these top-level messages:
	GetProductListByIdEndpoint
*/
package getproductlistbyidendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import getproductlistbyidrequestfilter "github.com/21stio/go-playground/grpc/schema/getproductlistbyidrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetProductListByIdEndpoint struct {
	Filter           *getproductlistbyidrequestfilter.GetProductListByIdRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                          `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                           `json:"-"`
}

func (m *GetProductListByIdEndpoint) Reset()                    { *m = GetProductListByIdEndpoint{} }
func (m *GetProductListByIdEndpoint) String() string            { return proto.CompactTextString(m) }
func (*GetProductListByIdEndpoint) ProtoMessage()               {}
func (*GetProductListByIdEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetProductListByIdEndpoint) GetFilter() *getproductlistbyidrequestfilter.GetProductListByIdRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *GetProductListByIdEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductListByIdEndpoint)(nil), "getproductlistbyidendpoint.GetProductListByIdEndpoint")
}

func init() {
	proto.RegisterFile("getproductlistbyidendpoint/getproductlistbyidendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4b, 0xc0, 0x30,
	0x10, 0x46, 0x89, 0x88, 0x60, 0xdc, 0x32, 0x95, 0x4e, 0xc5, 0xa9, 0x8b, 0x09, 0x76, 0x76, 0x90,
	0x82, 0x55, 0xc1, 0x41, 0x3a, 0x89, 0x8b, 0xb4, 0x49, 0x4c, 0x02, 0x6d, 0x2f, 0x26, 0x97, 0xa1,
	0x7f, 0xc1, 0x5f, 0x2d, 0xb4, 0xdd, 0x4a, 0xed, 0x76, 0xf0, 0x1e, 0xef, 0x83, 0xa3, 0x0f, 0x46,
	0xa3, 0x0f, 0xa0, 0x92, 0xc4, 0xc1, 0x45, 0xec, 0x67, 0xa7, 0xf4, 0xa4, 0x3c, 0xb8, 0x09, 0xc5,
	0x31, 0xfa, 0xe2, 0x3e, 0x00, 0x02, 0xcb, 0x8f, 0x95, 0xbc, 0xd9, 0xb3, 0xa0, 0x7f, 0x92, 0x8e,
	0xf8, 0xed, 0x06, 0xd4, 0x41, 0x9c, 0xf0, 0x6d, 0xe3, 0xf6, 0x97, 0xd0, 0xfc, 0x59, 0xe3, 0xfb,
	0xaa, 0xbe, 0xb9, 0x88, 0xf5, 0xfc, 0xaa, 0x9e, 0xb6, 0x19, 0xf6, 0x41, 0xaf, 0x56, 0x3f, 0x23,
	0x05, 0x29, 0x6f, 0xaa, 0x47, 0x7e, 0xd2, 0xe5, 0xfb, 0x58, 0xbb, 0xf2, 0x66, 0xe1, 0xed, 0xd6,
	0x63, 0x8c, 0x5e, 0xda, 0x2e, 0xda, 0xec, 0xa2, 0x20, 0xe5, 0x75, 0xbb, 0xdc, 0xf5, 0xcb, 0x67,
	0x63, 0x1c, 0xda, 0xd4, 0x73, 0x09, 0xa3, 0xa8, 0xee, 0x23, 0x3a, 0x10, 0x06, 0xee, 0xfc, 0xd0,
	0xcd, 0x26, 0x40, 0x9a, 0x94, 0x30, 0xc1, 0x4b, 0x11, 0xa5, 0xd5, 0x63, 0xf7, 0xcf, 0x07, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0x48, 0xb3, 0xef, 0x79, 0x01, 0x00, 0x00,
}
