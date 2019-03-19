// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getproductlistbyidrequest/getproductlistbyidrequest_.proto

/*
Package getproductlistbyidrequest is a generated protocol buffer package.

It is generated from these files:
	getproductlistbyidrequest/getproductlistbyidrequest_.proto

It has these top-level messages:
	GetProductListByIdRequest
*/
package getproductlistbyidrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import cache "github.com/21stio/go-playground/grpc/schema/cache"
import id "github.com/21stio/go-playground/grpc/schema/id"
import productskind "github.com/21stio/go-playground/grpc/schema/productskind"
import getproductlistbyidresponseselect "github.com/21stio/go-playground/grpc/schema/getproductlistbyidresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetProductListByIdRequest struct {
	Meta             *requestmeta.RequestMeta                                           `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Cache            *cache.Cache                                                       `protobuf:"bytes,2,opt,name=cache" json:"cache,omitempty"`
	Id               *id.Id                                                             `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Kind             *productskind.ProductsKind                                         `protobuf:"varint,4,opt,name=kind,enum=productskind.ProductsKind" json:"kind,omitempty"`
	Select           *getproductlistbyidresponseselect.GetProductListByIdResponseSelect `protobuf:"bytes,5,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                            `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                             `json:"-"`
}

func (m *GetProductListByIdRequest) Reset()                    { *m = GetProductListByIdRequest{} }
func (m *GetProductListByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProductListByIdRequest) ProtoMessage()               {}
func (*GetProductListByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetProductListByIdRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetProductListByIdRequest) GetCache() *cache.Cache {
	if m != nil {
		return m.Cache
	}
	return nil
}

func (m *GetProductListByIdRequest) GetId() *id.Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetProductListByIdRequest) GetKind() productskind.ProductsKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return productskind.ProductsKind_PRODUCT_SIMILAR
}

func (m *GetProductListByIdRequest) GetSelect() *getproductlistbyidresponseselect.GetProductListByIdResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *GetProductListByIdRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductListByIdRequest)(nil), "getproductlistbyidrequest.GetProductListByIdRequest")
}

func init() {
	proto.RegisterFile("getproductlistbyidrequest/getproductlistbyidrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0xed, 0x06, 0xc6, 0xe1, 0x21, 0x07, 0xc9, 0x76, 0x90, 0xb2, 0xd3, 0x0e, 0xfa,
	0x82, 0x3b, 0x7a, 0x9c, 0xc8, 0x1c, 0x2a, 0x48, 0xbc, 0xed, 0x22, 0x59, 0x12, 0xda, 0x87, 0x5b,
	0x53, 0x9b, 0xf4, 0xb0, 0x0f, 0xe4, 0xf7, 0x94, 0xa5, 0x19, 0xf4, 0xb0, 0xe1, 0xe5, 0xf1, 0x5e,
	0xdf, 0xaf, 0xff, 0xbe, 0x1f, 0x94, 0x3c, 0x16, 0xc6, 0xd7, 0x8d, 0xd5, 0xad, 0xf2, 0x5b, 0x74,
	0x7e, 0xb3, 0x47, 0xdd, 0x98, 0x9f, 0xd6, 0x38, 0xcf, 0xcf, 0x6e, 0xbe, 0xa0, 0x6e, 0xac, 0xb7,
	0x74, 0x7c, 0x96, 0x98, 0xdc, 0xc6, 0x66, 0x67, 0xbc, 0xe4, 0xbd, 0x3e, 0xbe, 0x3a, 0xa1, 0x4a,
	0xaa, 0xd2, 0xf0, 0x50, 0x8f, 0xcf, 0x46, 0xa8, 0x39, 0xea, 0xe3, 0x94, 0xc7, 0x64, 0xf7, 0x8d,
	0x95, 0xe6, 0xfd, 0xe1, 0x48, 0xbc, 0x9c, 0xfa, 0xbc, 0xab, 0x6d, 0xe5, 0x8c, 0x33, 0x5b, 0xa3,
	0x4e, 0x1b, 0xf4, 0x81, 0x98, 0x34, 0xfd, 0x4d, 0xc9, 0x78, 0x69, 0xfc, 0x47, 0xc7, 0xbe, 0xa1,
	0xf3, 0x8b, 0xfd, 0x4a, 0x8b, 0xee, 0x6c, 0x7a, 0x47, 0xb2, 0xc3, 0xe9, 0x2c, 0xc9, 0x93, 0xd9,
	0xd5, 0x9c, 0x41, 0x4f, 0x07, 0x22, 0xf3, 0x6e, 0xbc, 0x14, 0x81, 0xa2, 0x53, 0x32, 0x08, 0x56,
	0x2c, 0x0d, 0xf8, 0x08, 0xc2, 0x04, 0x4f, 0x87, 0x2a, 0xba, 0x15, 0xbd, 0x21, 0x29, 0x6a, 0x76,
	0x11, 0x80, 0x21, 0xa0, 0x86, 0x95, 0x16, 0x29, 0x6a, 0x0a, 0x24, 0x3b, 0x08, 0xb2, 0x2c, 0x4f,
	0x66, 0xd7, 0xf3, 0x09, 0xf4, 0xad, 0x21, 0x5e, 0xe7, 0x5e, 0xb1, 0xd2, 0x22, 0x70, 0x74, 0x4d,
	0x86, 0x9d, 0x08, 0x1b, 0x84, 0xac, 0x05, 0xfc, 0x67, 0x0c, 0xa7, 0x34, 0x3b, 0xe0, 0x33, 0x00,
	0x22, 0x26, 0x52, 0x4a, 0xb2, 0x52, 0xba, 0x92, 0x0d, 0xf3, 0x64, 0x76, 0x29, 0x42, 0xbf, 0x58,
	0xae, 0x9f, 0x0b, 0xf4, 0x65, 0xbb, 0x01, 0x65, 0x77, 0x7c, 0xfe, 0xe0, 0x3c, 0x5a, 0x5e, 0xd8,
	0xfb, 0x7a, 0x2b, 0xf7, 0x45, 0x63, 0xdb, 0x4a, 0xf3, 0xa2, 0xa9, 0x15, 0x77, 0xaa, 0x34, 0x3b,
	0x79, 0xfe, 0x07, 0xfa, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x74, 0x5c, 0xd9, 0x76, 0x02, 0x00,
	0x00,
}