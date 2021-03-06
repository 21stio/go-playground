// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weightperweight/weightperweight_.proto

/*
Package weightperweight is a generated protocol buffer package.

It is generated from these files:
	weightperweight/weightperweight_.proto

It has these top-level messages:
	WeightPerWeight
*/
package weightperweight

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import weightvalue "github.com/21stio/go-playground/grpc/schema/weightvalue"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WeightPerWeight struct {
	Counter          *weightvalue.WeightValue `protobuf:"bytes,1,opt,name=counter" json:"counter,omitempty"`
	Divider          *weightvalue.WeightValue `protobuf:"bytes,2,opt,name=divider" json:"divider,omitempty"`
	Hash             *string                  `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *WeightPerWeight) Reset()                    { *m = WeightPerWeight{} }
func (m *WeightPerWeight) String() string            { return proto.CompactTextString(m) }
func (*WeightPerWeight) ProtoMessage()               {}
func (*WeightPerWeight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WeightPerWeight) GetCounter() *weightvalue.WeightValue {
	if m != nil {
		return m.Counter
	}
	return nil
}

func (m *WeightPerWeight) GetDivider() *weightvalue.WeightValue {
	if m != nil {
		return m.Divider
	}
	return nil
}

func (m *WeightPerWeight) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*WeightPerWeight)(nil), "weightperweight.WeightPerWeight")
}

func init() { proto.RegisterFile("weightperweight/weightperweight_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x4f, 0xcd, 0x4c,
	0xcf, 0x28, 0x29, 0x48, 0x2d, 0x82, 0x30, 0xf4, 0xd1, 0xf8, 0xf1, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xfc, 0x68, 0xe2, 0x52, 0x72, 0x10, 0xba, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x1f, 0x89,
	0x0d, 0xd5, 0xa0, 0xd4, 0xcb, 0xc8, 0xc5, 0x1f, 0x0e, 0x16, 0x0e, 0x48, 0x2d, 0x82, 0x30, 0x84,
	0x8c, 0xb8, 0xd8, 0x93, 0xf3, 0x4b, 0xf3, 0x4a, 0x52, 0x8b, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8,
	0x8d, 0x24, 0xf4, 0x90, 0x74, 0xea, 0x41, 0x54, 0x85, 0x81, 0xd8, 0x41, 0x30, 0x85, 0x20, 0x3d,
	0x29, 0x99, 0x65, 0x99, 0x29, 0xa9, 0x45, 0x12, 0x4c, 0x84, 0xf4, 0x40, 0x15, 0x0a, 0x09, 0x71,
	0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4e, 0xb6,
	0x51, 0xd6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x46, 0x86, 0xc5,
	0x25, 0x99, 0xf9, 0xfa, 0xe9, 0xf9, 0xba, 0x05, 0x39, 0x89, 0x95, 0xe9, 0x45, 0xf9, 0xa5, 0x79,
	0x29, 0xfa, 0xe9, 0x45, 0x05, 0xc9, 0xfa, 0xc5, 0xc9, 0x19, 0xa9, 0xb9, 0x89, 0xe8, 0xc1, 0x00,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xf9, 0x82, 0xd4, 0x28, 0x01, 0x00, 0x00,
}
