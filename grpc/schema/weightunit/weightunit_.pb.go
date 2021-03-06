// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weightunit/weightunit_.proto

/*
Package weightunit is a generated protocol buffer package.

It is generated from these files:
	weightunit/weightunit_.proto

It has these top-level messages:
*/
package weightunit

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

type WeightUnit int32

const (
	WeightUnit_MILLIGRAMM WeightUnit = 0
	WeightUnit_GRAMM      WeightUnit = 1
	WeightUnit_KILOGRAMM  WeightUnit = 2
)

var WeightUnit_name = map[int32]string{
	0: "MILLIGRAMM",
	1: "GRAMM",
	2: "KILOGRAMM",
}
var WeightUnit_value = map[string]int32{
	"MILLIGRAMM": 0,
	"GRAMM":      1,
	"KILOGRAMM":  2,
}

func (x WeightUnit) Enum() *WeightUnit {
	p := new(WeightUnit)
	*p = x
	return p
}
func (x WeightUnit) String() string {
	return proto.EnumName(WeightUnit_name, int32(x))
}
func (x *WeightUnit) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WeightUnit_value, data, "WeightUnit")
	if err != nil {
		return err
	}
	*x = WeightUnit(value)
	return nil
}
func (WeightUnit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("weightunit.WeightUnit", WeightUnit_name, WeightUnit_value)
}

func init() { proto.RegisterFile("weightunit/weightunit_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x4f, 0xcd, 0x4c,
	0xcf, 0x28, 0x29, 0xcd, 0xcb, 0x2c, 0xd1, 0x47, 0x30, 0xe3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xb8, 0x10, 0x42, 0x5a, 0x66, 0x5c, 0x5c, 0xe1, 0x60, 0x5e, 0x68, 0x5e, 0x66, 0x89, 0x10,
	0x1f, 0x17, 0x97, 0xaf, 0xa7, 0x8f, 0x8f, 0xa7, 0x7b, 0x90, 0xa3, 0xaf, 0xaf, 0x00, 0x83, 0x10,
	0x27, 0x17, 0x2b, 0x84, 0xc9, 0x28, 0xc4, 0xcb, 0xc5, 0xe9, 0xed, 0xe9, 0xe3, 0x0f, 0xe1, 0x32,
	0x39, 0x59, 0x44, 0x99, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x1b,
	0x19, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16, 0xe4, 0x24, 0x56, 0xa6, 0x17, 0xe5,
	0x97, 0xe6, 0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0x22,
	0x39, 0x02, 0x10, 0x00, 0x00, 0xff, 0xff, 0xce, 0xab, 0x60, 0xdf, 0x9c, 0x00, 0x00, 0x00,
}
