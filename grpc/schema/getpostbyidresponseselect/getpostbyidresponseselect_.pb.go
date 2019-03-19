// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getpostbyidresponseselect/getpostbyidresponseselect_.proto

/*
Package getpostbyidresponseselect is a generated protocol buffer package.

It is generated from these files:
	getpostbyidresponseselect/getpostbyidresponseselect_.proto

It has these top-level messages:
	GetPostByIdResponseSelect
*/
package getpostbyidresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"
import feedselect "github.com/21stio/go-playground/grpc/schema/feedselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetPostByIdResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	Post             *feedselect.PostSelect                 `protobuf:"bytes,3,opt,name=post" json:"post,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,4,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *GetPostByIdResponseSelect) Reset()                    { *m = GetPostByIdResponseSelect{} }
func (m *GetPostByIdResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*GetPostByIdResponseSelect) ProtoMessage()               {}
func (*GetPostByIdResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetPostByIdResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *GetPostByIdResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *GetPostByIdResponseSelect) GetPost() *feedselect.PostSelect {
	if m != nil {
		return m.Post
	}
	return nil
}

func (m *GetPostByIdResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *GetPostByIdResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*GetPostByIdResponseSelect)(nil), "getpostbyidresponseselect.GetPostByIdResponseSelect")
}

func init() {
	proto.RegisterFile("getpostbyidresponseselect/getpostbyidresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0x56, 0x70, 0xf1, 0x96, 0x83, 0x74, 0x63, 0x48, 0xf1, 0x20, 0x45, 0x34, 0xc1,
	0x1d, 0x77, 0x73, 0x20, 0xd3, 0x83, 0x20, 0xf5, 0xe6, 0x45, 0xb2, 0xf6, 0x99, 0x16, 0xda, 0x25,
	0x34, 0x6f, 0x87, 0xfe, 0xa1, 0xfe, 0x3f, 0x92, 0x34, 0xa5, 0x42, 0xe9, 0xed, 0xe5, 0xfb, 0x7e,
	0xef, 0xcb, 0x17, 0x42, 0xb7, 0x0a, 0xd0, 0x68, 0x8b, 0x87, 0xae, 0x2a, 0x5a, 0xb0, 0x46, 0x1f,
	0x2d, 0x58, 0xa8, 0x21, 0x47, 0x31, 0xeb, 0x7c, 0x73, 0xd3, 0x6a, 0xd4, 0x6c, 0x39, 0x4b, 0xac,
	0x1e, 0x86, 0x73, 0x03, 0x28, 0x43, 0xde, 0x54, 0x0a, 0x41, 0xab, 0xf5, 0x0f, 0x40, 0x11, 0xa8,
	0x71, 0x0c, 0xee, 0xed, 0x2f, 0xa1, 0xcb, 0x3d, 0xe0, 0x87, 0xb6, 0xb8, 0xeb, 0xde, 0x8a, 0x2c,
	0xc4, 0x7c, 0x7a, 0x88, 0xad, 0xe9, 0xa2, 0xc7, 0x9f, 0xeb, 0x3a, 0x26, 0x09, 0x49, 0x2f, 0xb3,
	0x51, 0x60, 0x5b, 0x1a, 0xb9, 0xeb, 0xe2, 0xb3, 0x84, 0xa4, 0x57, 0x9b, 0x3b, 0x3e, 0xed, 0xc0,
	0x87, 0xbc, 0x77, 0x40, 0xd9, 0x67, 0x66, 0x7e, 0x87, 0xdd, 0xd3, 0xc8, 0xbd, 0x2e, 0x3e, 0xf7,
	0xbb, 0xd7, 0x7c, 0x6c, 0xc6, 0x5d, 0x97, 0x81, 0x75, 0x0c, 0xbb, 0xa1, 0xb4, 0xb7, 0x5e, 0xa5,
	0x2d, 0xe3, 0xc8, 0xd7, 0xf8, 0xa7, 0x30, 0x46, 0xa3, 0xd2, 0x39, 0x17, 0x09, 0x49, 0x17, 0x99,
	0x9f, 0x77, 0xfb, 0xaf, 0x17, 0x55, 0x61, 0x79, 0x3a, 0xf0, 0x5c, 0x37, 0x62, 0xf3, 0x64, 0xb1,
	0xd2, 0x42, 0xe9, 0x47, 0x53, 0xcb, 0x4e, 0xb5, 0xfa, 0x74, 0x2c, 0x84, 0x6a, 0x4d, 0x2e, 0x6c,
	0x5e, 0x42, 0x23, 0xe7, 0xbf, 0xe3, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x2c, 0x4e, 0xb9, 0xc4,
	0x01, 0x00, 0x00,
}