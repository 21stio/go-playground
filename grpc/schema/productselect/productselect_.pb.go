// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productselect/productselect_.proto

/*
Package productselect is a generated protocol buffer package.

It is generated from these files:
	productselect/productselect_.proto

It has these top-level messages:
	ProductSelect
*/
package productselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import infoselect "github.com/21stio/go-playground/grpc/schema/infoselect"
import brandselect "github.com/21stio/go-playground/grpc/schema/brandselect"
import nutritionsselect "github.com/21stio/go-playground/grpc/schema/nutritionsselect"
import physicalsselect "github.com/21stio/go-playground/grpc/schema/physicalsselect"
import productsselect "github.com/21stio/go-playground/grpc/schema/productsselect"
import idsselect "github.com/21stio/go-playground/grpc/schema/idsselect"
import serviceselect "github.com/21stio/go-playground/grpc/schema/serviceselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProductSelect struct {
	SelectAll        *bool                              `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Info             *infoselect.InfoSelect             `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	Brand            *brandselect.BrandSelect           `protobuf:"bytes,3,opt,name=brand" json:"brand,omitempty"`
	Nutritions       *nutritionsselect.NutritionsSelect `protobuf:"bytes,4,opt,name=nutritions" json:"nutritions,omitempty"`
	Physicals        *physicalsselect.PhysicalsSelect   `protobuf:"bytes,5,opt,name=physicals" json:"physicals,omitempty"`
	Similar          *productsselect.ProductsSelect     `protobuf:"bytes,6,opt,name=similar" json:"similar,omitempty"`
	SelectHash       *bool                              `protobuf:"varint,7,opt,name=selectHash" json:"selectHash,omitempty"`
	Ids              *idsselect.IdsSelect               `protobuf:"bytes,8,opt,name=ids" json:"ids,omitempty"`
	Meta             *serviceselect.TypeMetaSelect      `protobuf:"bytes,9,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                            `protobuf:"bytes,10,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *ProductSelect) Reset()                    { *m = ProductSelect{} }
func (m *ProductSelect) String() string            { return proto.CompactTextString(m) }
func (*ProductSelect) ProtoMessage()               {}
func (*ProductSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProductSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *ProductSelect) GetInfo() *infoselect.InfoSelect {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ProductSelect) GetBrand() *brandselect.BrandSelect {
	if m != nil {
		return m.Brand
	}
	return nil
}

func (m *ProductSelect) GetNutritions() *nutritionsselect.NutritionsSelect {
	if m != nil {
		return m.Nutritions
	}
	return nil
}

func (m *ProductSelect) GetPhysicals() *physicalsselect.PhysicalsSelect {
	if m != nil {
		return m.Physicals
	}
	return nil
}

func (m *ProductSelect) GetSimilar() *productsselect.ProductsSelect {
	if m != nil {
		return m.Similar
	}
	return nil
}

func (m *ProductSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *ProductSelect) GetIds() *idsselect.IdsSelect {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ProductSelect) GetMeta() *serviceselect.TypeMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ProductSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ProductSelect)(nil), "productselect.ProductSelect")
}

func init() { proto.RegisterFile("productselect/productselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0x5d, 0xab, 0xd3, 0x30,
	0x18, 0xc7, 0xa9, 0xeb, 0xf1, 0x9c, 0x46, 0xce, 0x4d, 0x10, 0x09, 0x63, 0x8e, 0x52, 0x64, 0x14,
	0xc1, 0xa7, 0x6c, 0x57, 0x8a, 0x20, 0xb8, 0x2b, 0x77, 0xa1, 0x8c, 0xe9, 0x95, 0x37, 0x92, 0xb5,
	0x59, 0x1b, 0x68, 0x9b, 0x92, 0xa4, 0xc2, 0xbe, 0xa8, 0x9f, 0x47, 0x9a, 0x26, 0x7d, 0xbb, 0x7b,
	0x5e, 0x7e, 0xcf, 0x7f, 0xdb, 0x2f, 0x0c, 0x45, 0x8d, 0x14, 0x59, 0x9b, 0x6a, 0xc5, 0x4a, 0x96,
	0xea, 0x64, 0xd6, 0xfd, 0x81, 0x46, 0x0a, 0x2d, 0xf0, 0xf3, 0x6c, 0xba, 0xde, 0xf0, 0xfa, 0x26,
	0x2c, 0x3f, 0x96, 0x16, 0x5e, 0x6f, 0xaf, 0x92, 0xd6, 0x99, 0x5d, 0x4f, 0x6a, 0xb7, 0x8f, 0xeb,
	0x56, 0x4b, 0xae, 0xb9, 0xa8, 0x95, 0x85, 0x96, 0x03, 0x47, 0xee, 0x9a, 0xe2, 0xae, 0x78, 0x4a,
	0x4b, 0x07, 0x2e, 0x7a, 0xc7, 0xbd, 0x73, 0x5f, 0x6f, 0xf1, 0x1b, 0xe6, 0xd4, 0x9a, 0x67, 0x0e,
	0x18, 0x2a, 0xb7, 0x8b, 0x14, 0x93, 0x7f, 0x79, 0xca, 0xec, 0x7e, 0xd6, 0x59, 0x26, 0xfa, 0xb7,
	0x42, 0xcf, 0xe7, 0x3e, 0xf9, 0xa7, 0x59, 0xe0, 0x0d, 0x0a, 0x7a, 0xe4, 0x6b, 0x59, 0x12, 0x2f,
	0xf4, 0xe2, 0xa7, 0xcb, 0x38, 0xc0, 0xef, 0x91, 0xdf, 0xc9, 0x21, 0x2f, 0x42, 0x2f, 0x7e, 0x75,
	0x78, 0x03, 0xa3, 0x29, 0x38, 0xd5, 0x37, 0xd1, 0x67, 0x5c, 0x0c, 0x83, 0x01, 0x3d, 0x18, 0x53,
	0x64, 0x65, 0x60, 0x02, 0x13, 0x6f, 0x70, 0xec, 0x6a, 0x8b, 0xf7, 0x18, 0x3e, 0x22, 0x34, 0x4a,
	0x23, 0xbe, 0x39, 0x8a, 0x60, 0xe9, 0x11, 0x7e, 0x0c, 0x03, 0x7b, 0x3e, 0xb9, 0xc2, 0x5f, 0x50,
	0x30, 0xf8, 0x24, 0x0f, 0x26, 0x22, 0x84, 0x85, 0x61, 0x38, 0xbb, 0xde, 0x06, 0x8c, 0x27, 0xf8,
	0x23, 0x7a, 0x54, 0xbc, 0xe2, 0x25, 0x95, 0xe4, 0xa5, 0xb9, 0xde, 0xc2, 0x5c, 0x3c, 0x58, 0x5b,
	0xee, 0xd6, 0xe1, 0x78, 0x8b, 0x50, 0x4f, 0x7c, 0xa3, 0xaa, 0x20, 0x8f, 0x46, 0xdc, 0x64, 0x82,
	0x77, 0x68, 0xc5, 0x33, 0x45, 0x9e, 0x4c, 0xea, 0x6b, 0x18, 0x5e, 0x0b, 0x4e, 0x99, 0xcb, 0xea,
	0x00, 0xbc, 0x47, 0x7e, 0xc5, 0x34, 0x25, 0x81, 0x01, 0xdf, 0xc2, 0xec, 0xd9, 0xe0, 0xd7, 0xbd,
	0x61, 0xdf, 0x99, 0xa6, 0x4e, 0x74, 0x87, 0x62, 0x8c, 0xfc, 0xa2, 0xfb, 0x50, 0x14, 0x7a, 0x71,
	0x70, 0x31, 0xf5, 0xf1, 0xf3, 0xef, 0x4f, 0x39, 0xd7, 0x45, 0x7b, 0x85, 0x54, 0x54, 0xc9, 0x61,
	0xaf, 0x34, 0x17, 0x49, 0x2e, 0x3e, 0x34, 0x25, 0xbd, 0xe7, 0x52, 0xb4, 0x75, 0x96, 0xe4, 0xb2,
	0x49, 0x13, 0x95, 0x16, 0xac, 0xa2, 0xf3, 0x7f, 0xc8, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9,
	0x2d, 0xab, 0xdf, 0x3f, 0x03, 0x00, 0x00,
}
