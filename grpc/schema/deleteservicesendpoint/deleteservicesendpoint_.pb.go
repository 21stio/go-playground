// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deleteservicesendpoint/deleteservicesendpoint_.proto

/*
Package deleteservicesendpoint is a generated protocol buffer package.

It is generated from these files:
	deleteservicesendpoint/deleteservicesendpoint_.proto

It has these top-level messages:
	DeleteServicesEndpoint
*/
package deleteservicesendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import deleteservicesrequestfilter "github.com/21stio/go-playground/grpc/schema/deleteservicesrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteServicesEndpoint struct {
	Filter           *deleteservicesrequestfilter.DeleteServicesRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                  `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                   `json:"-"`
}

func (m *DeleteServicesEndpoint) Reset()                    { *m = DeleteServicesEndpoint{} }
func (m *DeleteServicesEndpoint) String() string            { return proto.CompactTextString(m) }
func (*DeleteServicesEndpoint) ProtoMessage()               {}
func (*DeleteServicesEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteServicesEndpoint) GetFilter() *deleteservicesrequestfilter.DeleteServicesRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *DeleteServicesEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteServicesEndpoint)(nil), "deleteservicesendpoint.DeleteServicesEndpoint")
}

func init() {
	proto.RegisterFile("deleteservicesendpoint/deleteservicesendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0x4e, 0xcd, 0x4b, 0x29, 0xc8, 0xcf,
	0xcc, 0x2b, 0xd1, 0xc7, 0x2e, 0x1c, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24, 0x86, 0x5d,
	0x5a, 0xca, 0x0e, 0x55, 0xbc, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x24, 0x2d, 0x33, 0xa7, 0x24,
	0xb5, 0x48, 0x1f, 0x8f, 0x1c, 0xd4, 0x5c, 0xa5, 0x3a, 0x2e, 0x31, 0x17, 0xb0, 0xaa, 0x60, 0xa8,
	0x2a, 0x57, 0xa8, 0xc9, 0x42, 0x01, 0x5c, 0x6c, 0x10, 0xa5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0x16, 0x7a, 0x78, 0x8c, 0xd3, 0x43, 0x35, 0x24, 0x08, 0x22, 0xe7, 0x06, 0x96, 0x0b, 0x82,
	0x9a, 0x23, 0x24, 0xc4, 0xc5, 0x92, 0x91, 0x58, 0x9c, 0x21, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19,
	0x04, 0x66, 0x3b, 0xb9, 0x44, 0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0x1b, 0x19, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16, 0xe4, 0x24, 0x56, 0xa6,
	0x17, 0xe5, 0x97, 0xe6, 0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17, 0x27, 0x67, 0xa4, 0xe6,
	0x26, 0xe2, 0x08, 0x24, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x94, 0x44, 0x73, 0x54, 0x01,
	0x00, 0x00,
}
