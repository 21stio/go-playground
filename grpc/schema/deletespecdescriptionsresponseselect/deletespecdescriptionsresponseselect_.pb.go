// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletespecdescriptionsresponseselect/deletespecdescriptionsresponseselect_.proto

/*
Package deletespecdescriptionsresponseselect is a generated protocol buffer package.

It is generated from these files:
	deletespecdescriptionsresponseselect/deletespecdescriptionsresponseselect_.proto

It has these top-level messages:
	DeleteSpecDescriptionsResponseSelect
*/
package deletespecdescriptionsresponseselect

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

type DeleteSpecDescriptionsResponseSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Meta             *responsemetaselect.ResponseMetaSelect `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,3,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *DeleteSpecDescriptionsResponseSelect) Reset()         { *m = DeleteSpecDescriptionsResponseSelect{} }
func (m *DeleteSpecDescriptionsResponseSelect) String() string { return proto.CompactTextString(m) }
func (*DeleteSpecDescriptionsResponseSelect) ProtoMessage()    {}
func (*DeleteSpecDescriptionsResponseSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *DeleteSpecDescriptionsResponseSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *DeleteSpecDescriptionsResponseSelect) GetMeta() *responsemetaselect.ResponseMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteSpecDescriptionsResponseSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *DeleteSpecDescriptionsResponseSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteSpecDescriptionsResponseSelect)(nil), "deletespecdescriptionsresponseselect.DeleteSpecDescriptionsResponseSelect")
}

func init() {
	proto.RegisterFile("deletespecdescriptionsresponseselect/deletespecdescriptionsresponseselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0x89, 0x76, 0xf0, 0xe2, 0x96, 0xa9, 0x88, 0x48, 0x91, 0x43, 0x3a, 0x68, 0x83, 0x37,
	0xba, 0x29, 0x37, 0xb8, 0x08, 0x47, 0x6e, 0x73, 0x91, 0x98, 0x3e, 0xda, 0x42, 0xda, 0x17, 0xf2,
	0xde, 0x0d, 0xfe, 0x5f, 0xfe, 0x81, 0x72, 0x69, 0x0f, 0x0f, 0xba, 0x74, 0x4b, 0xbe, 0xfc, 0xbe,
	0xdf, 0x07, 0x91, 0xbb, 0x1a, 0x3c, 0x30, 0x50, 0x00, 0x57, 0x03, 0xb9, 0xd8, 0x05, 0xee, 0x70,
	0xa0, 0x08, 0x14, 0x70, 0x20, 0x20, 0xf0, 0xe0, 0x58, 0x2f, 0x81, 0xbe, 0xaa, 0x10, 0x91, 0x51,
	0xad, 0x97, 0xc0, 0x37, 0x8f, 0xa7, 0x7b, 0x0f, 0x6c, 0xa7, 0x95, 0x79, 0x34, 0x39, 0xef, 0x7f,
	0x85, 0x5c, 0x6f, 0x93, 0x76, 0x1f, 0xc0, 0x6d, 0xcf, 0xb4, 0x66, 0xea, 0xec, 0x13, 0xaf, 0x6e,
	0xe5, 0x6a, 0x6c, 0xbe, 0x7a, 0x9f, 0x8b, 0x42, 0x94, 0x57, 0xe6, 0x3f, 0x50, 0x2f, 0x32, 0x3b,
	0xba, 0xf3, 0x8b, 0x42, 0x94, 0xd7, 0x9b, 0x87, 0x6a, 0x3e, 0x58, 0x9d, 0x7c, 0x1f, 0xc0, 0x76,
	0x74, 0x9a, 0xd4, 0x51, 0x77, 0x52, 0x8e, 0xc8, 0xbb, 0xa5, 0x36, 0xbf, 0x4c, 0xea, 0xb3, 0x44,
	0x29, 0x99, 0xb5, 0xc7, 0x97, 0xac, 0x10, 0xe5, 0xca, 0xa4, 0xf3, 0x9b, 0xf9, 0xdc, 0x35, 0x1d,
	0xb7, 0x87, 0xef, 0xca, 0x61, 0xaf, 0x37, 0xcf, 0xc4, 0x1d, 0xea, 0x06, 0x9f, 0x82, 0xb7, 0x3f,
	0x4d, 0xc4, 0xc3, 0x50, 0xeb, 0x26, 0x06, 0xa7, 0xc9, 0xb5, 0xd0, 0xdb, 0x45, 0xbf, 0xfc, 0x17,
	0x00, 0x00, 0xff, 0xff, 0x72, 0x48, 0x91, 0xfc, 0xb1, 0x01, 0x00, 0x00,
}