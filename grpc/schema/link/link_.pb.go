// Code generated by protoc-gen-go. DO NOT EDIT.
// source: link/link_.proto

/*
Package link is a generated protocol buffer package.

It is generated from these files:
	link/link_.proto

It has these top-level messages:
	Link
*/
package link

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import url "github.com/21stio/go-playground/grpc/schema/url"
import linkkind "github.com/21stio/go-playground/grpc/schema/linkkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Link struct {
	Url              *url.Url           `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Kind             *linkkind.LinkKind `protobuf:"varint,2,opt,name=kind,enum=linkkind.LinkKind" json:"kind,omitempty"`
	Hash             *string            `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Link) GetUrl() *url.Url {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Link) GetKind() linkkind.LinkKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return linkkind.LinkKind_DETAILS
}

func (m *Link) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Link)(nil), "link.Link")
}

func init() { proto.RegisterFile("link/link_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xc1, 0x8a, 0x83, 0x30,
	0x10, 0x86, 0xc9, 0x9a, 0xc3, 0x6e, 0x16, 0x64, 0xc9, 0x29, 0x78, 0x92, 0x3d, 0x14, 0x2f, 0x4d,
	0xda, 0x3c, 0x42, 0xaf, 0xed, 0x49, 0xe8, 0xa5, 0x87, 0x16, 0xab, 0x25, 0x09, 0x8e, 0x89, 0x8c,
	0xe6, 0xd0, 0xb7, 0x2f, 0x11, 0xbd, 0x0c, 0xff, 0x7c, 0x3f, 0xdf, 0x30, 0xec, 0x0f, 0x9c, 0xef,
	0x55, 0x1a, 0x0f, 0x39, 0x62, 0x98, 0x03, 0xa7, 0x69, 0x29, 0xf2, 0x88, 0xa0, 0x22, 0xc2, 0x4a,
	0x0b, 0x91, 0x68, 0xef, 0x7c, 0xa7, 0xb6, 0xb0, 0x36, 0xff, 0x77, 0x46, 0x2f, 0xce, 0xf7, 0xbc,
	0x60, 0x59, 0x44, 0x10, 0xa4, 0x24, 0xd5, 0xaf, 0xfe, 0x96, 0x11, 0x41, 0x5e, 0x11, 0xea, 0x04,
	0xf9, 0x8e, 0xd1, 0xa4, 0x88, 0xaf, 0x92, 0x54, 0xb9, 0xe6, 0x72, 0xbb, 0x21, 0x93, 0x79, 0x76,
	0xbe, 0xab, 0x97, 0x9e, 0x73, 0x46, 0x6d, 0x33, 0x59, 0x91, 0x95, 0xa4, 0xfa, 0xa9, 0x97, 0x7c,
	0xd2, 0xb7, 0x83, 0x71, 0xb3, 0x8d, 0x4f, 0xd9, 0x86, 0x41, 0xe9, 0xe3, 0x34, 0xbb, 0xa0, 0x4c,
	0xd8, 0x8f, 0xd0, 0xbc, 0x0d, 0x86, 0xe8, 0x3b, 0x65, 0x70, 0x6c, 0xd5, 0xd4, 0xda, 0xd7, 0xd0,
	0x2c, 0xef, 0x7d, 0x02, 0x00, 0x00, 0xff, 0xff, 0x88, 0x92, 0x11, 0xbe, 0xd6, 0x00, 0x00, 0x00,
}
