// Code generated by protoc-gen-go. DO NOT EDIT.
// source: speccontext/speccontext_.proto

/*
Package speccontext is a generated protocol buffer package.

It is generated from these files:
	speccontext/speccontext_.proto

It has these top-level messages:
	SpecContext
*/
package speccontext

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import text "github.com/21stio/go-playground/grpc/schema/text"
import error "github.com/21stio/go-playground/grpc/schema/error"
import specassertion "github.com/21stio/go-playground/grpc/schema/specassertion"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SpecContext struct {
	Name             *string                        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Pending          *bool                          `protobuf:"varint,2,opt,name=pending" json:"pending,omitempty"`
	Comment          *text.Text                     `protobuf:"bytes,3,opt,name=comment" json:"comment,omitempty"`
	ExampleInput     *text.Text                     `protobuf:"bytes,4,opt,name=exampleInput" json:"exampleInput,omitempty"`
	ExampleOutput    *text.Text                     `protobuf:"bytes,5,opt,name=exampleOutput" json:"exampleOutput,omitempty"`
	Errors           []*error.Error                 `protobuf:"bytes,6,rep,name=errors" json:"errors,omitempty"`
	Assertions       []*specassertion.SpecAssertion `protobuf:"bytes,7,rep,name=assertions" json:"assertions,omitempty"`
	Hash             *string                        `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *SpecContext) Reset()                    { *m = SpecContext{} }
func (m *SpecContext) String() string            { return proto.CompactTextString(m) }
func (*SpecContext) ProtoMessage()               {}
func (*SpecContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SpecContext) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SpecContext) GetPending() bool {
	if m != nil && m.Pending != nil {
		return *m.Pending
	}
	return false
}

func (m *SpecContext) GetComment() *text.Text {
	if m != nil {
		return m.Comment
	}
	return nil
}

func (m *SpecContext) GetExampleInput() *text.Text {
	if m != nil {
		return m.ExampleInput
	}
	return nil
}

func (m *SpecContext) GetExampleOutput() *text.Text {
	if m != nil {
		return m.ExampleOutput
	}
	return nil
}

func (m *SpecContext) GetErrors() []*error.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *SpecContext) GetAssertions() []*specassertion.SpecAssertion {
	if m != nil {
		return m.Assertions
	}
	return nil
}

func (m *SpecContext) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*SpecContext)(nil), "speccontext.SpecContext")
}

func init() { proto.RegisterFile("speccontext/speccontext_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4b, 0x33, 0x31,
	0x10, 0xc6, 0xd9, 0xb6, 0x6f, 0xdb, 0x37, 0x5b, 0x41, 0x72, 0x0a, 0x45, 0x64, 0x29, 0x3d, 0xec,
	0xc5, 0x44, 0x7b, 0x11, 0xc1, 0x8b, 0x8a, 0x07, 0x4f, 0xc2, 0xea, 0xc9, 0x8b, 0xc4, 0x74, 0xd8,
	0x5d, 0x68, 0xfe, 0x90, 0x64, 0x61, 0xfd, 0x6c, 0x7e, 0x39, 0x49, 0xba, 0x2b, 0xbb, 0xf4, 0x12,
	0xe6, 0x99, 0xe7, 0x37, 0x03, 0xf3, 0x04, 0x5d, 0x3a, 0x03, 0x42, 0x68, 0xe5, 0xa1, 0xf5, 0x6c,
	0x50, 0x7f, 0x52, 0x63, 0xb5, 0xd7, 0x38, 0x1d, 0xf4, 0xd6, 0xe7, 0x91, 0x1a, 0xd8, 0x6b, 0x0c,
	0xd6, 0x6a, 0xcb, 0xe2, 0xdb, 0xf7, 0x36, 0x61, 0x84, 0x3b, 0x07, 0xd6, 0xd7, 0x5a, 0xb1, 0x91,
	0xea, 0x98, 0xcd, 0xcf, 0x04, 0xa5, 0x6f, 0x06, 0xc4, 0xd3, 0x71, 0x33, 0xc6, 0x68, 0xa6, 0xb8,
	0x04, 0x92, 0x64, 0x49, 0xfe, 0xbf, 0x88, 0x35, 0x26, 0x68, 0x61, 0x40, 0xed, 0x6b, 0x55, 0x92,
	0x49, 0x96, 0xe4, 0xcb, 0xa2, 0x97, 0x78, 0x8b, 0x16, 0x42, 0x4b, 0x09, 0xca, 0x93, 0x69, 0x96,
	0xe4, 0xe9, 0x0e, 0xd1, 0xb0, 0x85, 0xbe, 0x43, 0xeb, 0x8b, 0xde, 0xc2, 0x14, 0xad, 0xa0, 0xe5,
	0xd2, 0x1c, 0xe0, 0x45, 0x99, 0xc6, 0x93, 0xd9, 0x09, 0x3a, 0xf2, 0xf1, 0x35, 0x3a, 0xeb, 0xf4,
	0x6b, 0xe3, 0xc3, 0xc0, 0xbf, 0x93, 0x81, 0x31, 0x80, 0xb7, 0x68, 0x1e, 0x2f, 0x77, 0x64, 0x9e,
	0x4d, 0xf3, 0x74, 0xb7, 0xa2, 0x51, 0xd2, 0xe7, 0xf0, 0x16, 0x9d, 0x87, 0xef, 0x11, 0xfa, 0xbb,
	0xdf, 0x91, 0x45, 0x24, 0x2f, 0xe8, 0x28, 0x16, 0x1a, 0xb2, 0x78, 0xe8, 0x55, 0x31, 0xe0, 0x43,
	0x32, 0x15, 0x77, 0x15, 0x59, 0x1e, 0x93, 0x09, 0xf5, 0xe3, 0xdd, 0xc7, 0x6d, 0x59, 0xfb, 0xaa,
	0xf9, 0xa2, 0x42, 0x4b, 0xb6, 0xbb, 0x71, 0xbe, 0xd6, 0xac, 0xd4, 0x57, 0xe6, 0xc0, 0xbf, 0x4b,
	0xab, 0x1b, 0xb5, 0x67, 0xa5, 0x35, 0x82, 0x39, 0x51, 0x81, 0xe4, 0xc3, 0x6f, 0xfd, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0x1f, 0xa0, 0xe2, 0xa0, 0xf0, 0x01, 0x00, 0x00,
}
