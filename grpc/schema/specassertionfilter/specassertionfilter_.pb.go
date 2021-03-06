// Code generated by protoc-gen-go. DO NOT EDIT.
// source: specassertionfilter/specassertionfilter_.proto

/*
Package specassertionfilter is a generated protocol buffer package.

It is generated from these files:
	specassertionfilter/specassertionfilter_.proto

It has these top-level messages:
	SpecAssertionFilter
*/
package specassertionfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"
import boolfilter "github.com/21stio/go-playground/grpc/schema/boolfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SpecAssertionFilter struct {
	Name             *stringfilter.StringFilter `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Pending          *boolfilter.BoolFilter     `protobuf:"bytes,2,opt,name=pending" json:"pending,omitempty"`
	Passed           *boolfilter.BoolFilter     `protobuf:"bytes,3,opt,name=passed" json:"passed,omitempty"`
	And              []*SpecAssertionFilter     `protobuf:"bytes,4,rep,name=and" json:"and,omitempty"`
	Or               []*SpecAssertionFilter     `protobuf:"bytes,5,rep,name=or" json:"or,omitempty"`
	Not              []*SpecAssertionFilter     `protobuf:"bytes,6,rep,name=not" json:"not,omitempty"`
	Hash             *string                    `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *SpecAssertionFilter) Reset()                    { *m = SpecAssertionFilter{} }
func (m *SpecAssertionFilter) String() string            { return proto.CompactTextString(m) }
func (*SpecAssertionFilter) ProtoMessage()               {}
func (*SpecAssertionFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpecAssertionFilter) GetName() *stringfilter.StringFilter {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SpecAssertionFilter) GetPending() *boolfilter.BoolFilter {
	if m != nil {
		return m.Pending
	}
	return nil
}

func (m *SpecAssertionFilter) GetPassed() *boolfilter.BoolFilter {
	if m != nil {
		return m.Passed
	}
	return nil
}

func (m *SpecAssertionFilter) GetAnd() []*SpecAssertionFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *SpecAssertionFilter) GetOr() []*SpecAssertionFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *SpecAssertionFilter) GetNot() []*SpecAssertionFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *SpecAssertionFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*SpecAssertionFilter)(nil), "specassertionfilter.SpecAssertionFilter")
}

func init() { proto.RegisterFile("specassertionfilter/specassertionfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0xc6, 0xf1, 0x9f, 0x26, 0x54, 0xd9, 0x14, 0x28, 0xc2, 0x74, 0x30, 0x9d, 0xbc, 0xf4, 0xdc,
	0x7a, 0x2a, 0x5d, 0x4a, 0x32, 0xf4, 0x01, 0x9c, 0xad, 0x4b, 0x51, 0x6c, 0x55, 0x16, 0xd8, 0x3a,
	0x21, 0x29, 0x43, 0x1f, 0xba, 0xef, 0x50, 0xec, 0xda, 0x24, 0x83, 0x28, 0x64, 0xbb, 0xd3, 0xf7,
	0xfd, 0xee, 0xe3, 0xc4, 0x11, 0x70, 0x46, 0x34, 0xdc, 0x39, 0x61, 0xbd, 0x42, 0xfd, 0xa5, 0x7a,
	0x2f, 0x6c, 0x19, 0x78, 0xfb, 0x04, 0x63, 0xd1, 0x23, 0xdd, 0x06, 0xb4, 0x2c, 0x77, 0xde, 0x2a,
	0x2d, 0x17, 0xfa, 0xa2, 0x99, 0xb1, 0xec, 0xfe, 0x88, 0xd8, 0xcf, 0xfa, 0xb9, 0x9c, 0xd5, 0x87,
	0x9f, 0x98, 0x6c, 0x0f, 0x46, 0x34, 0xbb, 0x65, 0xee, 0xfb, 0x24, 0x53, 0x20, 0xa9, 0xe6, 0x83,
	0x60, 0x51, 0x1e, 0x15, 0x9b, 0x2a, 0x83, 0xcb, 0xc9, 0x70, 0x98, 0x9a, 0x3f, 0x67, 0x3d, 0xf9,
	0xe8, 0x13, 0x59, 0x1b, 0xa1, 0x5b, 0xa5, 0x25, 0x8b, 0x27, 0xe4, 0x0e, 0xce, 0x61, 0xb0, 0x47,
	0xec, 0x67, 0xfb, 0x62, 0xa3, 0x40, 0x56, 0x66, 0xdc, 0xa6, 0x65, 0xc9, 0xbf, 0xc0, 0xec, 0xa2,
	0xaf, 0x24, 0xe1, 0xba, 0x65, 0x69, 0x9e, 0x14, 0x9b, 0xaa, 0x08, 0x7d, 0x1e, 0x04, 0x16, 0xa9,
	0x47, 0x88, 0xbe, 0x90, 0x18, 0x2d, 0xbb, 0xb9, 0x12, 0x8d, 0xd1, 0x8e, 0xa9, 0x1a, 0x3d, 0x5b,
	0x5d, 0x9b, 0xaa, 0xd1, 0x53, 0x4a, 0xd2, 0x8e, 0xbb, 0x8e, 0xad, 0xf3, 0xa8, 0xb8, 0xad, 0xa7,
	0x7a, 0xbf, 0xfb, 0x78, 0x93, 0xca, 0x77, 0xa7, 0x23, 0x34, 0x38, 0x94, 0xd5, 0xb3, 0xf3, 0x0a,
	0x4b, 0x89, 0x8f, 0xa6, 0xe7, 0xdf, 0xd2, 0xe2, 0x49, 0xb7, 0xa5, 0xb4, 0xa6, 0x29, 0x5d, 0xd3,
	0x89, 0x81, 0x87, 0xce, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfa, 0xc2, 0x74, 0x4e, 0x38, 0x02,
	0x00, 0x00,
}
