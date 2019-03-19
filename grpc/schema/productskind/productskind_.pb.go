// Code generated by protoc-gen-go. DO NOT EDIT.
// source: productskind/productskind_.proto

/*
Package productskind is a generated protocol buffer package.

It is generated from these files:
	productskind/productskind_.proto

It has these top-level messages:
*/
package productskind

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProductsKind int32

const (
	ProductsKind_PRODUCT_SIMILAR ProductsKind = 0
)

var ProductsKind_name = map[int32]string{
	0: "PRODUCT_SIMILAR",
}
var ProductsKind_value = map[string]int32{
	"PRODUCT_SIMILAR": 0,
}

func (x ProductsKind) Enum() *ProductsKind {
	p := new(ProductsKind)
	*p = x
	return p
}
func (x ProductsKind) String() string {
	return proto.EnumName(ProductsKind_name, int32(x))
}
func (x *ProductsKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ProductsKind_value, data, "ProductsKind")
	if err != nil {
		return err
	}
	*x = ProductsKind(value)
	return nil
}
func (ProductsKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("productskind.ProductsKind", ProductsKind_name, ProductsKind_value)
}

func init() { proto.RegisterFile("productskind/productskind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0x29, 0xce, 0xce, 0xcc, 0x4b, 0xd1, 0x47, 0xe6, 0xc4, 0xeb, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0xf1, 0x20, 0x0b, 0x6a, 0x29, 0x73, 0xf1, 0x04, 0x40, 0xf9, 0xde, 0x99, 0x79,
	0x29, 0x42, 0xc2, 0x5c, 0xfc, 0x01, 0x41, 0xfe, 0x2e, 0xa1, 0xce, 0x21, 0xf1, 0xc1, 0x9e, 0xbe,
	0x9e, 0x3e, 0x8e, 0x41, 0x02, 0x0c, 0x4e, 0x56, 0x51, 0x16, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49,
	0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x46, 0x86, 0xc5, 0x25, 0x99, 0xf9, 0xfa, 0xe9, 0xf9, 0xba, 0x05,
	0x39, 0x89, 0x95, 0xe9, 0x45, 0xf9, 0xa5, 0x79, 0x29, 0xfa, 0xe9, 0x45, 0x05, 0xc9, 0xfa, 0xc5,
	0xc9, 0x19, 0xa9, 0xb9, 0x89, 0x28, 0xb6, 0x02, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x2b, 0x75,
	0xfa, 0x91, 0x00, 0x00, 0x00,
}