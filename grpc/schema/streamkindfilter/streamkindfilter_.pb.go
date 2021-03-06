// Code generated by protoc-gen-go. DO NOT EDIT.
// source: streamkindfilter/streamkindfilter_.proto

/*
Package streamkindfilter is a generated protocol buffer package.

It is generated from these files:
	streamkindfilter/streamkindfilter_.proto

It has these top-level messages:
	StreamKindFilter
*/
package streamkindfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import streamkind "github.com/21stio/go-playground/grpc/schema/streamkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StreamKindFilter struct {
	Is               *streamkind.StreamKind  `protobuf:"varint,1,opt,name=is,enum=streamkind.StreamKind" json:"is,omitempty"`
	Not              *streamkind.StreamKind  `protobuf:"varint,2,opt,name=not,enum=streamkind.StreamKind" json:"not,omitempty"`
	In               []streamkind.StreamKind `protobuf:"varint,3,rep,name=in,enum=streamkind.StreamKind" json:"in,omitempty"`
	NotIn            []streamkind.StreamKind `protobuf:"varint,4,rep,name=notIn,enum=streamkind.StreamKind" json:"notIn,omitempty"`
	Hash             *string                 `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *StreamKindFilter) Reset()                    { *m = StreamKindFilter{} }
func (m *StreamKindFilter) String() string            { return proto.CompactTextString(m) }
func (*StreamKindFilter) ProtoMessage()               {}
func (*StreamKindFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamKindFilter) GetIs() streamkind.StreamKind {
	if m != nil && m.Is != nil {
		return *m.Is
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamKindFilter) GetNot() streamkind.StreamKind {
	if m != nil && m.Not != nil {
		return *m.Not
	}
	return streamkind.StreamKind_WATCH
}

func (m *StreamKindFilter) GetIn() []streamkind.StreamKind {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *StreamKindFilter) GetNotIn() []streamkind.StreamKind {
	if m != nil {
		return m.NotIn
	}
	return nil
}

func (m *StreamKindFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*StreamKindFilter)(nil), "streamkindfilter.StreamKindFilter")
}

func init() { proto.RegisterFile("streamkindfilter/streamkindfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x2e, 0x29, 0x4a,
	0x4d, 0xcc, 0xcd, 0xce, 0xcc, 0x4b, 0x49, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2, 0x47, 0x17, 0x88,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x40, 0x97, 0x90, 0x92, 0x41, 0x88, 0x20, 0xe9,
	0x82, 0xaa, 0x57, 0x3a, 0xc7, 0xc8, 0x25, 0x10, 0x0c, 0x16, 0xf5, 0xce, 0xcc, 0x4b, 0x71, 0x03,
	0x6b, 0x11, 0x52, 0xe3, 0x62, 0xca, 0x2c, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0x12, 0xd3,
	0x43, 0x68, 0xd2, 0x43, 0xa8, 0x0c, 0x62, 0xca, 0x2c, 0x16, 0xd2, 0xe0, 0x62, 0xce, 0xcb, 0x2f,
	0x91, 0x60, 0xc2, 0xab, 0x10, 0xa4, 0x04, 0x6c, 0x62, 0x9e, 0x04, 0xb3, 0x02, 0x33, 0x5e, 0x13,
	0xf3, 0x84, 0x74, 0xb8, 0x58, 0xf3, 0xf2, 0x4b, 0x3c, 0xf3, 0x24, 0x58, 0xf0, 0x2a, 0x85, 0x28,
	0x12, 0x12, 0xe2, 0x62, 0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x55, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0xb3, 0x9d, 0xec, 0xa2, 0x6c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x8d, 0x0c, 0x8b, 0x4b, 0x32, 0xf3, 0xf5, 0xd3, 0xf3, 0x75, 0x0b, 0x72, 0x12, 0x2b, 0xd3, 0x8b,
	0xf2, 0x4b, 0xf3, 0x52, 0xf4, 0xd3, 0x8b, 0x0a, 0x92, 0xf5, 0x8b, 0x93, 0x33, 0x52, 0x73, 0x13,
	0x31, 0xc2, 0x11, 0x10, 0x00, 0x00, 0xff, 0xff, 0x28, 0xcf, 0x02, 0xe9, 0x6b, 0x01, 0x00, 0x00,
}
