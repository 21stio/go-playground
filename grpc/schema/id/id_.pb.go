// Code generated by protoc-gen-go. DO NOT EDIT.
// source: id/id_.proto

/*
Package id is a generated protocol buffer package.

It is generated from these files:
	id/id_.proto

It has these top-level messages:
	Id
*/
package id

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import idkind "github.com/21stio/go-playground/grpc/schema/idkind"
import url "github.com/21stio/go-playground/grpc/schema/url"
import serviceid "github.com/21stio/go-playground/grpc/schema/serviceid"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Id struct {
	Kind             *idkind.IdKind       `protobuf:"varint,1,opt,name=kind,enum=idkind.IdKind" json:"kind,omitempty"`
	Url              *url.Url             `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Name             *string              `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Username         *string              `protobuf:"bytes,4,opt,name=username" json:"username,omitempty"`
	Ean              *string              `protobuf:"bytes,5,opt,name=ean" json:"ean,omitempty"`
	ServiceId        *serviceid.ServiceId `protobuf:"bytes,6,opt,name=serviceId" json:"serviceId,omitempty"`
	Local            *string              `protobuf:"bytes,7,opt,name=local" json:"local,omitempty"`
	Hash             *string              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Id) GetKind() idkind.IdKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return idkind.IdKind_ID_URL
}

func (m *Id) GetUrl() *url.Url {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Id) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Id) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Id) GetEan() string {
	if m != nil && m.Ean != nil {
		return *m.Ean
	}
	return ""
}

func (m *Id) GetServiceId() *serviceid.ServiceId {
	if m != nil {
		return m.ServiceId
	}
	return nil
}

func (m *Id) GetLocal() string {
	if m != nil && m.Local != nil {
		return *m.Local
	}
	return ""
}

func (m *Id) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Id)(nil), "id.Id")
}

func init() { proto.RegisterFile("id/id_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8e, 0xc1, 0x4e, 0xb4, 0x30,
	0x14, 0x85, 0xc3, 0x0c, 0xf3, 0xff, 0x4c, 0x35, 0xc4, 0x34, 0x2c, 0x1a, 0x56, 0x64, 0x56, 0x6c,
	0x6c, 0x95, 0x47, 0x70, 0x47, 0xdc, 0x61, 0xdc, 0xb8, 0x31, 0x95, 0xdb, 0xc0, 0x8d, 0x85, 0x92,
	0x42, 0x4d, 0x7c, 0x5f, 0x1f, 0xc4, 0xb4, 0x30, 0xb8, 0xea, 0x39, 0xdf, 0x69, 0xce, 0x3d, 0xe4,
	0x16, 0x41, 0x20, 0xbc, 0xf3, 0xc9, 0x9a, 0xc5, 0xd0, 0x03, 0x42, 0x9e, 0x21, 0x7c, 0xe2, 0xe8,
	0xa9, 0x7f, 0xb6, 0x24, 0x4f, 0x9d, 0xd5, 0xc2, 0x59, 0x7d, 0xf5, 0xf9, 0xac, 0xec, 0x17, 0xb6,
	0x0a, 0x41, 0xec, 0x6a, 0xcb, 0x2e, 0x3f, 0x11, 0x39, 0xd4, 0x40, 0x2f, 0x24, 0xf6, 0x0d, 0x2c,
	0x2a, 0xa2, 0x32, 0xad, 0x52, 0xbe, 0x16, 0xf2, 0x1a, 0x9e, 0x71, 0x84, 0x26, 0x64, 0x34, 0x27,
	0x47, 0x67, 0x35, 0x3b, 0x14, 0x51, 0x79, 0x53, 0x25, 0xdc, 0x59, 0xcd, 0x5f, 0xad, 0x6e, 0x3c,
	0xa4, 0x94, 0xc4, 0xa3, 0x1c, 0x14, 0x3b, 0x16, 0x51, 0x79, 0x6e, 0x82, 0xa6, 0x39, 0x49, 0xdc,
	0xac, 0x6c, 0xe0, 0x71, 0xe0, 0xbb, 0xa7, 0x77, 0xe4, 0xa8, 0xe4, 0xc8, 0x4e, 0x01, 0x7b, 0x49,
	0x2b, 0x72, 0xde, 0xc6, 0xd5, 0xc0, 0xfe, 0x85, 0x1b, 0x19, 0xdf, 0xe7, 0xf2, 0x97, 0x6b, 0xd6,
	0xfc, 0x7d, 0xa3, 0x19, 0x39, 0x69, 0xd3, 0x4a, 0xcd, 0xfe, 0x87, 0x9e, 0xd5, 0xf8, 0x2d, 0xbd,
	0x9c, 0x7b, 0x96, 0xac, 0x5b, 0xbc, 0x7e, 0x7a, 0x78, 0xe3, 0x1d, 0x2e, 0xbd, 0xfb, 0xe0, 0xad,
	0x19, 0x44, 0xf5, 0x38, 0x2f, 0x68, 0x44, 0x67, 0xee, 0x27, 0x2d, 0xbf, 0x3b, 0x6b, 0xdc, 0x08,
	0xa2, 0xb3, 0x53, 0x2b, 0xe6, 0xb6, 0x57, 0x83, 0x14, 0x08, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xd0, 0x41, 0xda, 0x7d, 0x6d, 0x01, 0x00, 0x00,
}
