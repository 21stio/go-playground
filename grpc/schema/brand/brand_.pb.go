// Code generated by protoc-gen-go. DO NOT EDIT.
// source: brand/brand_.proto

/*
Package brand is a generated protocol buffer package.

It is generated from these files:
	brand/brand_.proto

It has these top-level messages:
	Brand
*/
package brand

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import info "github.com/21stio/go-playground/grpc/schema/info"
import company "github.com/21stio/go-playground/grpc/schema/company"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Brand struct {
	Info             *info.Info       `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Company          *company.Company `protobuf:"bytes,2,opt,name=company" json:"company,omitempty"`
	Hash             *string          `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Brand) Reset()                    { *m = Brand{} }
func (m *Brand) String() string            { return proto.CompactTextString(m) }
func (*Brand) ProtoMessage()               {}
func (*Brand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Brand) GetInfo() *info.Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Brand) GetCompany() *company.Company {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *Brand) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Brand)(nil), "brand.Brand")
}

func init() { proto.RegisterFile("brand/brand_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2a, 0x4a, 0xcc,
	0x4b, 0xd1, 0x07, 0x93, 0xf1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x9e, 0x94,
	0x40, 0x66, 0x5e, 0x5a, 0xbe, 0x3e, 0x88, 0x80, 0x4a, 0x48, 0x89, 0x25, 0xe7, 0xe7, 0x16, 0x24,
	0xe6, 0x55, 0xea, 0x43, 0x69, 0xa8, 0xb8, 0x52, 0x3a, 0x17, 0xab, 0x13, 0x48, 0x8b, 0x90, 0x1c,
	0x17, 0x0b, 0x48, 0xbd, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0x97, 0x1e, 0x88, 0xa3, 0xe7,
	0x99, 0x97, 0x96, 0x1f, 0x04, 0x16, 0x17, 0xd2, 0xe2, 0x62, 0x87, 0x6a, 0x95, 0x60, 0x02, 0x2b,
	0x11, 0xd0, 0x83, 0xf2, 0xf5, 0x9c, 0x21, 0x74, 0x10, 0x4c, 0x81, 0x90, 0x10, 0x17, 0x4b, 0x46,
	0x62, 0x71, 0x86, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xed, 0x64, 0x1c, 0x65, 0x98,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0x04, 0xd2, 0xa6, 0x6f, 0x64, 0x58, 0x5c, 0x92, 0x99, 0xaf, 0x9f,
	0x9e, 0xaf, 0x5b, 0x90, 0x93, 0x58, 0x99, 0x5e, 0x94, 0x5f, 0x9a, 0x97, 0xa2, 0x9f, 0x5e, 0x54,
	0x90, 0xac, 0x5f, 0x9c, 0x9c, 0x91, 0x9a, 0x9b, 0x08, 0xf1, 0x15, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xf7, 0x63, 0x4c, 0x03, 0xe3, 0x00, 0x00, 0x00,
}