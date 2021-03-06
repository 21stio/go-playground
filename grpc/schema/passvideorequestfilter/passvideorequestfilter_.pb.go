// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passvideorequestfilter/passvideorequestfilter_.proto

/*
Package passvideorequestfilter is a generated protocol buffer package.

It is generated from these files:
	passvideorequestfilter/passvideorequestfilter_.proto

It has these top-level messages:
	PassVideoRequestFilter
*/
package passvideorequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import videofilter "github.com/21stio/go-playground/grpc/schema/videofilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassVideoRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Video            *videofilter.VideoFilter             `protobuf:"bytes,2,opt,name=video" json:"video,omitempty"`
	And              []*PassVideoRequestFilter            `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or               []*PassVideoRequestFilter            `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not              []*PassVideoRequestFilter            `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *PassVideoRequestFilter) Reset()                    { *m = PassVideoRequestFilter{} }
func (m *PassVideoRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*PassVideoRequestFilter) ProtoMessage()               {}
func (*PassVideoRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PassVideoRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassVideoRequestFilter) GetVideo() *videofilter.VideoFilter {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *PassVideoRequestFilter) GetAnd() []*PassVideoRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassVideoRequestFilter) GetOr() []*PassVideoRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassVideoRequestFilter) GetNot() []*PassVideoRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassVideoRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassVideoRequestFilter)(nil), "passvideorequestfilter.PassVideoRequestFilter")
}

func init() {
	proto.RegisterFile("passvideorequestfilter/passvideorequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x3f, 0x4b, 0xc4, 0x30,
	0x14, 0xa7, 0x7f, 0x4e, 0x30, 0xb7, 0x65, 0x38, 0xc2, 0x0d, 0x52, 0xc4, 0xa1, 0x08, 0xbe, 0x60,
	0x71, 0x70, 0x12, 0x39, 0xc4, 0x4d, 0x90, 0x0e, 0x0e, 0x2e, 0x12, 0xdb, 0xd8, 0x16, 0xae, 0x7d,
	0x35, 0x79, 0x15, 0xfc, 0xaa, 0x7e, 0x1a, 0x69, 0x9a, 0xe1, 0xe0, 0x7a, 0x4b, 0xb7, 0x97, 0xf7,
	0xfb, 0x1b, 0x78, 0xec, 0xae, 0x57, 0xd6, 0xfe, 0x34, 0xa5, 0x46, 0xa3, 0xbf, 0x07, 0x6d, 0xe9,
	0xab, 0xd9, 0x93, 0x36, 0x72, 0x7e, 0xfd, 0x01, 0xbd, 0x41, 0x42, 0xbe, 0x99, 0x87, 0xb7, 0xd7,
	0xfe, 0xd9, 0x6a, 0x52, 0xde, 0xe8, 0x68, 0xe3, 0x3d, 0xb6, 0x17, 0x4e, 0xef, 0x59, 0x07, 0xb3,
	0xc7, 0x2f, 0xff, 0x42, 0xb6, 0x79, 0x55, 0xd6, 0xbe, 0x8d, 0x50, 0x3e, 0xb9, 0x3c, 0x3b, 0x06,
	0xbf, 0x67, 0xf1, 0xe8, 0x27, 0x82, 0x24, 0x48, 0xd7, 0xd9, 0x15, 0x1c, 0x65, 0x80, 0xe7, 0xbf,
	0x68, 0x52, 0x93, 0x26, 0x77, 0x0a, 0x0e, 0x6c, 0xe5, 0xa2, 0x44, 0xe8, 0xa4, 0x02, 0x0e, 0x82,
	0xc1, 0x25, 0x79, 0xfa, 0x44, 0xe3, 0x8f, 0x2c, 0x52, 0x5d, 0x29, 0xa2, 0x24, 0x4a, 0xd7, 0x19,
	0xc0, 0xfc, 0xb7, 0x61, 0xbe, 0x66, 0x3e, 0x4a, 0xf9, 0x03, 0x0b, 0xd1, 0x88, 0x78, 0x91, 0x41,
	0x88, 0x66, 0x6c, 0xd0, 0x21, 0x89, 0xd5, 0xb2, 0x06, 0x1d, 0x12, 0xe7, 0x2c, 0xae, 0x95, 0xad,
	0xc5, 0x59, 0x12, 0xa4, 0xe7, 0xb9, 0x9b, 0x77, 0x4f, 0xef, 0xbb, 0xaa, 0xa1, 0x7a, 0xf8, 0x84,
	0x02, 0x5b, 0x99, 0xdd, 0x5a, 0x6a, 0x50, 0x56, 0x78, 0xd3, 0xef, 0xd5, 0x6f, 0x65, 0x70, 0xe8,
	0x4a, 0x59, 0x99, 0xbe, 0x90, 0xb6, 0xa8, 0x75, 0xab, 0x4e, 0x5c, 0xc3, 0x7f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd3, 0xeb, 0x83, 0x3f, 0x3d, 0x02, 0x00, 0x00,
}
