// Code generated by protoc-gen-go. DO NOT EDIT.
// source: durationunitfilter/durationunitfilter_.proto

/*
Package durationunitfilter is a generated protocol buffer package.

It is generated from these files:
	durationunitfilter/durationunitfilter_.proto

It has these top-level messages:
	DurationUnitFilter
*/
package durationunitfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import durationunit "github.com/21stio/go-playground/grpc/schema/durationunit"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DurationUnitFilter struct {
	Is               *durationunit.DurationUnit  `protobuf:"varint,1,opt,name=is,enum=durationunit.DurationUnit" json:"is,omitempty"`
	Not              *durationunit.DurationUnit  `protobuf:"varint,2,opt,name=not,enum=durationunit.DurationUnit" json:"not,omitempty"`
	In               []durationunit.DurationUnit `protobuf:"varint,3,rep,name=in,enum=durationunit.DurationUnit" json:"in,omitempty"`
	NotIn            []durationunit.DurationUnit `protobuf:"varint,4,rep,name=notIn,enum=durationunit.DurationUnit" json:"notIn,omitempty"`
	Hash             *string                     `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                      `json:"-"`
}

func (m *DurationUnitFilter) Reset()                    { *m = DurationUnitFilter{} }
func (m *DurationUnitFilter) String() string            { return proto.CompactTextString(m) }
func (*DurationUnitFilter) ProtoMessage()               {}
func (*DurationUnitFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DurationUnitFilter) GetIs() durationunit.DurationUnit {
	if m != nil && m.Is != nil {
		return *m.Is
	}
	return durationunit.DurationUnit_NANOSECOND
}

func (m *DurationUnitFilter) GetNot() durationunit.DurationUnit {
	if m != nil && m.Not != nil {
		return *m.Not
	}
	return durationunit.DurationUnit_NANOSECOND
}

func (m *DurationUnitFilter) GetIn() []durationunit.DurationUnit {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *DurationUnitFilter) GetNotIn() []durationunit.DurationUnit {
	if m != nil {
		return m.NotIn
	}
	return nil
}

func (m *DurationUnitFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DurationUnitFilter)(nil), "durationunitfilter.DurationUnitFilter")
}

func init() { proto.RegisterFile("durationunitfilter/durationunitfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x29, 0x2d, 0x4a,
	0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0xcd, 0xcb, 0x2c, 0x49, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0xc7,
	0x14, 0x8a, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc2, 0x94, 0x92, 0x52, 0x40, 0x16,
	0x43, 0xd1, 0x0b, 0xd5, 0xa5, 0xf4, 0x80, 0x91, 0x4b, 0xc8, 0x05, 0x2a, 0x1e, 0x9a, 0x97, 0x59,
	0xe2, 0x06, 0xd6, 0x28, 0xa4, 0xc5, 0xc5, 0x94, 0x59, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x67,
	0x24, 0xa5, 0x87, 0xac, 0x51, 0x0f, 0x59, 0x75, 0x10, 0x53, 0x66, 0xb1, 0x90, 0x0e, 0x17, 0x73,
	0x5e, 0x7e, 0x89, 0x04, 0x13, 0x41, 0xc5, 0x20, 0x65, 0x60, 0x93, 0xf3, 0x24, 0x98, 0x15, 0x98,
	0x09, 0x9a, 0x9c, 0x27, 0x64, 0xc0, 0xc5, 0x9a, 0x97, 0x5f, 0xe2, 0x99, 0x27, 0xc1, 0x42, 0x50,
	0x39, 0x44, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x46, 0x62, 0x71, 0x86, 0x04, 0xab, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x98, 0xed, 0xe4, 0x10, 0x65, 0x97, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c,
	0x9f, 0xab, 0x6f, 0x64, 0x58, 0x5c, 0x92, 0x99, 0xaf, 0x9f, 0x9e, 0xaf, 0x5b, 0x90, 0x93, 0x58,
	0x99, 0x5e, 0x94, 0x5f, 0x9a, 0x97, 0xa2, 0x9f, 0x5e, 0x54, 0x90, 0xac, 0x5f, 0x9c, 0x9c, 0x91,
	0x9a, 0x9b, 0x88, 0x25, 0x84, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x18, 0xdd, 0xb2, 0x09, 0x89,
	0x01, 0x00, 0x00,
}
