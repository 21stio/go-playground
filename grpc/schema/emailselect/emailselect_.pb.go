// Code generated by protoc-gen-go. DO NOT EDIT.
// source: emailselect/emailselect_.proto

/*
Package emailselect is a generated protocol buffer package.

It is generated from these files:
	emailselect/emailselect_.proto

It has these top-level messages:
	EmailSelect
*/
package emailselect

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

type EmailSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Value            *bool   `protobuf:"varint,2,opt,name=value" json:"value,omitempty"`
	SelectHash       *bool   `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EmailSelect) Reset()                    { *m = EmailSelect{} }
func (m *EmailSelect) String() string            { return proto.CompactTextString(m) }
func (*EmailSelect) ProtoMessage()               {}
func (*EmailSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EmailSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *EmailSelect) GetValue() bool {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return false
}

func (m *EmailSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *EmailSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*EmailSelect)(nil), "emailselect.EmailSelect")
}

func init() { proto.RegisterFile("emailselect/emailselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcd, 0x4d, 0xcc,
	0xcc, 0x29, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1, 0x47, 0x62, 0xc7, 0xeb, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0x71, 0x23, 0x89, 0x29, 0x95, 0x72, 0x71, 0xbb, 0x82, 0xb8, 0xc1, 0x60, 0xae, 0x90,
	0x0c, 0x17, 0x27, 0x44, 0xc2, 0x31, 0x27, 0x47, 0x82, 0x51, 0x81, 0x51, 0x83, 0x23, 0x08, 0x21,
	0x20, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x96, 0x81, 0x70, 0x84,
	0xe4, 0xb8, 0xb8, 0x20, 0x4a, 0x3c, 0x12, 0x8b, 0x33, 0x24, 0x98, 0xc1, 0x52, 0x48, 0x22, 0x42,
	0x42, 0x5c, 0x2c, 0x19, 0x20, 0x19, 0x16, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0xdb, 0xc9, 0x32,
	0xca, 0x3c, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0xdf, 0xc8, 0xb0, 0xb8,
	0x24, 0x33, 0x5f, 0x3f, 0x3d, 0x5f, 0xb7, 0x20, 0x27, 0xb1, 0x32, 0xbd, 0x28, 0xbf, 0x34, 0x2f,
	0x45, 0x3f, 0xbd, 0xa8, 0x20, 0x59, 0xbf, 0x38, 0x39, 0x23, 0x35, 0x37, 0x11, 0xd9, 0x17, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x4d, 0x3a, 0x30, 0xdf, 0x00, 0x00, 0x00,
}
