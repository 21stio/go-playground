// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cache/cache_.proto

/*
Package cache is a generated protocol buffer package.

It is generated from these files:
	cache/cache_.proto

It has these top-level messages:
	Cache
*/
package cache

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import cachestage "github.com/21stio/go-playground/grpc/schema/cachestage"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Cache struct {
	Client           *cachestage.CacheStage `protobuf:"bytes,1,opt,name=client" json:"client,omitempty"`
	Hash             *string                `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Cache) Reset()                    { *m = Cache{} }
func (m *Cache) String() string            { return proto.CompactTextString(m) }
func (*Cache) ProtoMessage()               {}
func (*Cache) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Cache) GetClient() *cachestage.CacheStage {
	if m != nil {
		return m.Client
	}
	return nil
}

func (m *Cache) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Cache)(nil), "cache.Cache")
}

func init() { proto.RegisterFile("cache/cache_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4e, 0x4c, 0xce,
	0x48, 0xd5, 0x07, 0x93, 0xf1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x9e, 0x94,
	0x0c, 0x98, 0x2a, 0x2e, 0x49, 0x4c, 0x87, 0xca, 0x83, 0x99, 0x50, 0x45, 0x4a, 0xde, 0x5c, 0xac,
	0xce, 0x20, 0x41, 0x21, 0x3d, 0x2e, 0xb6, 0xe4, 0x9c, 0xcc, 0xd4, 0xbc, 0x12, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x6e, 0x23, 0x31, 0x3d, 0x84, 0x62, 0x3d, 0xb0, 0x92, 0x60, 0x10, 0x33, 0x08, 0xaa,
	0x4a, 0x48, 0x88, 0x8b, 0x25, 0x23, 0xb1, 0x38, 0x43, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08,
	0xcc, 0x76, 0x32, 0x8e, 0x32, 0x4c, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x37, 0x32, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x4f, 0xcf, 0xd7, 0x2d, 0xc8, 0x49, 0xac, 0x4c, 0x2f,
	0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d,
	0x84, 0xb8, 0x06, 0x10, 0x00, 0x00, 0xff, 0xff, 0x2f, 0x5c, 0xfb, 0xa1, 0xbb, 0x00, 0x00, 0x00,
}
