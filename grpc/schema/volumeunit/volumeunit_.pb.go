// Code generated by protoc-gen-go. DO NOT EDIT.
// source: volumeunit/volumeunit_.proto

/*
Package volumeunit is a generated protocol buffer package.

It is generated from these files:
	volumeunit/volumeunit_.proto

It has these top-level messages:
*/
package volumeunit

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

type VolumeUnit int32

const (
	VolumeUnit_LITER      VolumeUnit = 0
	VolumeUnit_MILLILITER VolumeUnit = 1
)

var VolumeUnit_name = map[int32]string{
	0: "LITER",
	1: "MILLILITER",
}
var VolumeUnit_value = map[string]int32{
	"LITER":      0,
	"MILLILITER": 1,
}

func (x VolumeUnit) Enum() *VolumeUnit {
	p := new(VolumeUnit)
	*p = x
	return p
}
func (x VolumeUnit) String() string {
	return proto.EnumName(VolumeUnit_name, int32(x))
}
func (x *VolumeUnit) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(VolumeUnit_value, data, "VolumeUnit")
	if err != nil {
		return err
	}
	*x = VolumeUnit(value)
	return nil
}
func (VolumeUnit) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("volumeunit.VolumeUnit", VolumeUnit_name, VolumeUnit_value)
}

func init() { proto.RegisterFile("volumeunit/volumeunit_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0xcb, 0xcf, 0x29,
	0xcd, 0x4d, 0x2d, 0xcd, 0xcb, 0x2c, 0xd1, 0x47, 0x30, 0xe3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xb8, 0x10, 0x42, 0x5a, 0xea, 0x5c, 0x5c, 0x61, 0x60, 0x5e, 0x68, 0x5e, 0x66, 0x89, 0x10,
	0x27, 0x17, 0xab, 0x8f, 0x67, 0x88, 0x6b, 0x90, 0x00, 0x83, 0x10, 0x1f, 0x17, 0x97, 0xaf, 0xa7,
	0x8f, 0x8f, 0x27, 0x84, 0xcf, 0xe8, 0x64, 0x11, 0x65, 0x96, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4,
	0x97, 0x9c, 0x9f, 0xab, 0x6f, 0x64, 0x58, 0x5c, 0x92, 0x99, 0xaf, 0x9f, 0x9e, 0xaf, 0x5b, 0x90,
	0x93, 0x58, 0x99, 0x5e, 0x94, 0x5f, 0x9a, 0x97, 0xa2, 0x9f, 0x5e, 0x54, 0x90, 0xac, 0x5f, 0x9c,
	0x9c, 0x91, 0x9a, 0x9b, 0x88, 0x64, 0x2b, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x1f, 0x18, 0x96,
	0x8d, 0x00, 0x00, 0x00,
}
