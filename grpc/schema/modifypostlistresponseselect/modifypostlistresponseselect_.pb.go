// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifypostlistresponseselect/modifypostlistresponseselect_.proto

/*
Package modifypostlistresponseselect is a generated protocol buffer package.

It is generated from these files:
	modifypostlistresponseselect/modifypostlistresponseselect_.proto

It has these top-level messages:
	ModifyPostListResponseSelect
*/
package modifypostlistresponseselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemetaselect "github.com/21stio/go-playground/grpc/schema/responsemetaselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ModifyPostListResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ModifyPostListResponseSelect) Reset()                    { *m = ModifyPostListResponseSelect{} }
func (m *ModifyPostListResponseSelect) String() string            { return proto.CompactTextString(m) }
func (*ModifyPostListResponseSelect) ProtoMessage()               {}
func (*ModifyPostListResponseSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ModifyPostListResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *ModifyPostListResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ModifyPostListResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *ModifyPostListResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifyPostListResponseSelect)(nil), "modifypostlistresponseselect.ModifyPostListResponseSelect")
}

func init() {
	proto.RegisterFile("modifypostlistresponseselect/modifypostlistresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0x89, 0x76, 0xf0, 0xe2, 0x96, 0xa9, 0x48, 0x91, 0xe2, 0x20, 0x1d, 0xb4, 0xc1, 0x1b,
	0x9d, 0xd4, 0xe9, 0x10, 0x0f, 0xa4, 0x6e, 0x2e, 0x12, 0x7b, 0x31, 0x09, 0x24, 0xf7, 0x85, 0x7c,
	0xdf, 0x0d, 0xf7, 0x9f, 0xfc, 0x91, 0x72, 0x69, 0x0f, 0x0f, 0x0a, 0xdd, 0x92, 0xf7, 0x7d, 0xde,
	0x27, 0x10, 0xfe, 0x14, 0x60, 0xe3, 0x7e, 0xf6, 0x11, 0x90, 0xbc, 0x43, 0x4a, 0x1a, 0x23, 0x6c,
	0x51, 0xa3, 0xf6, 0xba, 0x27, 0x39, 0x57, 0x7e, 0xb5, 0x31, 0x01, 0x81, 0xa8, 0xe6, 0xa0, 0xab,
	0xbb, 0xe3, 0x3d, 0x68, 0x52, 0xa3, 0x75, 0x1a, 0x8d, 0xae, 0x9b, 0x5f, 0xc6, 0xab, 0x75, 0xd6,
	0xbd, 0x03, 0xd2, 0x9b, 0x43, 0xea, 0x46, 0xf6, 0x23, 0x73, 0xa2, 0xe2, 0x8b, 0x61, 0xf1, 0xec,
	0x7d, 0xc9, 0x6a, 0xd6, 0x5c, 0x74, 0xff, 0x81, 0x78, 0xe4, 0xc5, 0xc1, 0x59, 0x9e, 0xd5, 0xac,
	0xb9, 0x5c, 0xde, 0xb6, 0xd3, 0x87, 0xda, 0xa3, 0x6f, 0xad, 0x49, 0x0d, 0xce, 0x2e, 0x6f, 0xc4,
	0x35, 0xe7, 0x03, 0xb2, 0x52, 0x68, 0xcb, 0xf3, 0xac, 0x3e, 0x49, 0x84, 0xe0, 0x85, 0x3d, 0x34,
	0x45, 0xcd, 0x9a, 0x45, 0x97, 0xcf, 0x2f, 0xaf, 0x9f, 0x2b, 0xe3, 0xc8, 0xee, 0xbe, 0xdb, 0x1e,
	0x82, 0x5c, 0x3e, 0x20, 0x39, 0x90, 0x06, 0xee, 0xa3, 0x57, 0x7b, 0x93, 0x60, 0xb7, 0xdd, 0x48,
	0x93, 0x62, 0x2f, 0xb1, 0xb7, 0x3a, 0xa8, 0xd9, 0xdf, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xfc,
	0xba, 0x45, 0x04, 0x89, 0x01, 0x00, 0x00,
}