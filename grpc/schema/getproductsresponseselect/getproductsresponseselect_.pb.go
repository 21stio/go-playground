// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getproductsresponseselect/getproductsresponseselect_.proto

/*
Package getproductsresponseselect is a generated protocol buffer package.

It is generated from these files:
	getproductsresponseselect/getproductsresponseselect_.proto

It has these top-level messages:
	GetProductsResponseSelect
*/
package getproductsresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import productsselect "github.com/21stio/go-playground/grpc/schema/productsselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetProductsResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	TotalCount       *bool                                  `protobuf:"varint,3,opt,name=totalCount" json:"totalCount,omitempty"`
	Products         *productsselect.ProductsSelect         `protobuf:"bytes,4,opt,name=products" json:"products,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,5,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *GetProductsResponseSelect) Reset()                    { *m = GetProductsResponseSelect{} }
func (m *GetProductsResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*GetProductsResponseSelect) ProtoMessage()               {}
func (*GetProductsResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetProductsResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetProductsResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetProductsResponseSelect) GetTotalCount() bool {
	if m != nil && m.TotalCount != nil {
		return *m.TotalCount
	}
	return false
}

func (m *GetProductsResponseSelect) GetProducts() *productsselect.ProductsSelect {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *GetProductsResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetProductsResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductsResponseSelect)(nil), "getproductsresponseselect.GetProductsResponseSelect")
}

func init() {
	proto.RegisterFile("getproductsresponseselect/getproductsresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0xc7, 0xe9, 0x7e, 0xfd, 0xc9, 0x16, 0x6f, 0x39, 0x75, 0x43, 0x46, 0x11, 0x91, 0x1e, 0x34,
	0xc5, 0x1d, 0x7b, 0x53, 0x91, 0x79, 0x11, 0xa4, 0xde, 0xbc, 0x48, 0xcc, 0x42, 0x3a, 0x48, 0xfb,
	0x42, 0xf2, 0x7a, 0xf0, 0x0f, 0xf0, 0xff, 0x96, 0xa6, 0x89, 0x73, 0x8c, 0xde, 0xfa, 0xbe, 0xef,
	0xd3, 0x4f, 0xbe, 0x21, 0xa4, 0x52, 0x12, 0x8d, 0x85, 0x5d, 0x2f, 0xd0, 0x59, 0xe9, 0x0c, 0x74,
	0x4e, 0x3a, 0xa9, 0xa5, 0xc0, 0x72, 0x72, 0xf3, 0xc1, 0x8c, 0x05, 0x04, 0xba, 0x9c, 0x24, 0x56,
	0x37, 0x71, 0x6e, 0x25, 0xf2, 0xe0, 0x3b, 0x8d, 0x82, 0x68, 0x75, 0x15, 0x2d, 0x81, 0x3c, 0x1e,
	0x03, 0x75, 0xf9, 0x3d, 0x23, 0xcb, 0xad, 0xc4, 0xd7, 0xb0, 0xac, 0x83, 0xee, 0xcd, 0x43, 0xf4,
	0x82, 0x2c, 0x46, 0xfc, 0x5e, 0xeb, 0x2c, 0xc9, 0x93, 0x62, 0x5e, 0x1f, 0x02, 0x5a, 0x91, 0x74,
	0x38, 0x36, 0x9b, 0xe5, 0x49, 0x71, 0xbe, 0xb9, 0x66, 0xa7, 0x5d, 0x58, 0xf4, 0xbd, 0x48, 0xe4,
	0xa3, 0xb3, 0xf6, 0xff, 0xd0, 0x35, 0x21, 0x08, 0xc8, 0xf5, 0x23, 0xf4, 0x1d, 0x66, 0xff, 0xbc,
	0xfa, 0x4f, 0x42, 0x2b, 0x32, 0x8f, 0x85, 0xb3, 0xd4, 0xfb, 0xd7, 0xec, 0xf8, 0x06, 0x2c, 0x76,
	0x0e, 0xde, 0x5f, 0x7e, 0x70, 0x8f, 0xc8, 0x33, 0x77, 0x4d, 0xf6, 0x7f, 0x74, 0x1f, 0x12, 0x4a,
	0x49, 0xda, 0x0c, 0x9b, 0xb3, 0x3c, 0x29, 0x16, 0xb5, 0xff, 0x7e, 0xd8, 0xbe, 0x3f, 0xa9, 0x3d,
	0x36, 0xfd, 0x27, 0x13, 0xd0, 0x96, 0x9b, 0x3b, 0x87, 0x7b, 0x28, 0x15, 0xdc, 0x1a, 0xcd, 0xbf,
	0x94, 0x85, 0xbe, 0xdb, 0x95, 0xca, 0x1a, 0x51, 0x3a, 0xd1, 0xc8, 0x96, 0x4f, 0x3f, 0xe3, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x78, 0x8d, 0x2a, 0xfc, 0x01, 0x00, 0x00,
}