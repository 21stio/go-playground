// Code generated by protoc-gen-go. DO NOT EDIT.
// source: urlsort/urlsort_.proto

/*
Package urlsort is a generated protocol buffer package.

It is generated from these files:
	urlsort/urlsort_.proto

It has these top-level messages:
	UrlSort
*/
package urlsort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import sortkind "github.com/21stio/go-playground/grpc/schema/sortkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UrlSort struct {
	Value            *sortkind.SortKind `protobuf:"varint,1,opt,name=value,enum=sortkind.SortKind" json:"value,omitempty"`
	Hash             *string            `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *UrlSort) Reset()                    { *m = UrlSort{} }
func (m *UrlSort) String() string            { return proto.CompactTextString(m) }
func (*UrlSort) ProtoMessage()               {}
func (*UrlSort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UrlSort) GetValue() sortkind.SortKind {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return sortkind.SortKind_ASC
}

func (m *UrlSort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UrlSort)(nil), "urlsort.UrlSort")
}

func init() { proto.RegisterFile("urlsort/urlsort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 166 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x2d, 0xca, 0x29,
	0xce, 0x2f, 0x2a, 0xd1, 0x87, 0xd2, 0xf1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50,
	0xbe, 0x94, 0x04, 0x88, 0xcc, 0xce, 0xcc, 0x4b, 0xd1, 0x87, 0x31, 0xa0, 0x4a, 0x94, 0xdc, 0xb9,
	0xd8, 0x43, 0x8b, 0x72, 0x82, 0xf3, 0x8b, 0x4a, 0x84, 0x34, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0xf8, 0x8c, 0x84, 0xf4, 0x60, 0x6a, 0xf5, 0x40, 0xd2, 0xde,
	0x99, 0x79, 0x29, 0x41, 0x10, 0x05, 0x42, 0x42, 0x5c, 0x2c, 0x19, 0x89, 0xc5, 0x19, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x93, 0x69, 0x94, 0x71, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0xbe, 0x91, 0x61, 0x71, 0x49, 0x66, 0xbe, 0x7e, 0x7a, 0xbe, 0x6e,
	0x41, 0x4e, 0x62, 0x65, 0x7a, 0x51, 0x7e, 0x69, 0x5e, 0x8a, 0x7e, 0x7a, 0x51, 0x41, 0xb2, 0x7e,
	0x71, 0x72, 0x46, 0x6a, 0x6e, 0x22, 0xcc, 0xa5, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x90,
	0xc7, 0x65, 0xbb, 0x00, 0x00, 0x00,
}
