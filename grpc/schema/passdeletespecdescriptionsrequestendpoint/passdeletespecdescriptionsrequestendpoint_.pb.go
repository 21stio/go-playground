// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletespecdescriptionsrequestendpoint/passdeletespecdescriptionsrequestendpoint_.proto

/*
Package passdeletespecdescriptionsrequestendpoint is a generated protocol buffer package.

It is generated from these files:
	passdeletespecdescriptionsrequestendpoint/passdeletespecdescriptionsrequestendpoint_.proto

It has these top-level messages:
	PassDeleteSpecDescriptionsRequestEndpoint
*/
package passdeletespecdescriptionsrequestendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import passdeletespecdescriptionsrequestrequestfilter "github.com/21stio/go-playground/grpc/schema/passdeletespecdescriptionsrequestrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteSpecDescriptionsRequestEndpoint struct {
	Filter           *passdeletespecdescriptionsrequestrequestfilter.PassDeleteSpecDescriptionsRequestRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                                                        `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                                                         `json:"-"`
}

func (m *PassDeleteSpecDescriptionsRequestEndpoint) Reset() {
	*m = PassDeleteSpecDescriptionsRequestEndpoint{}
}
func (m *PassDeleteSpecDescriptionsRequestEndpoint) String() string { return proto.CompactTextString(m) }
func (*PassDeleteSpecDescriptionsRequestEndpoint) ProtoMessage()    {}
func (*PassDeleteSpecDescriptionsRequestEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteSpecDescriptionsRequestEndpoint) GetFilter() *passdeletespecdescriptionsrequestrequestfilter.PassDeleteSpecDescriptionsRequestRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *PassDeleteSpecDescriptionsRequestEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteSpecDescriptionsRequestEndpoint)(nil), "passdeletespecdescriptionsrequestendpoint.PassDeleteSpecDescriptionsRequestEndpoint")
}

func init() {
	proto.RegisterFile("passdeletespecdescriptionsrequestendpoint/passdeletespecdescriptionsrequestendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0x59, 0x11, 0xc1, 0xd8, 0xa5, 0x3a, 0xac, 0x0e, 0xab, 0xbb, 0xc2, 0x04, 0xef, 0x11,
	0xe4, 0xb4, 0x96, 0x15, 0x11, 0xae, 0x50, 0x62, 0x32, 0x26, 0x81, 0xbd, 0xcc, 0x98, 0x99, 0x15,
	0x7c, 0x32, 0x5f, 0x4f, 0xd8, 0x4b, 0x61, 0x77, 0x7b, 0xd5, 0x3f, 0xc5, 0xf7, 0xf1, 0xff, 0x30,
	0x6a, 0x47, 0x8e, 0x39, 0xc0, 0x00, 0x02, 0x4c, 0xe0, 0x03, 0xb0, 0xaf, 0x99, 0x24, 0x63, 0xe1,
	0x0a, 0x5f, 0x23, 0xb0, 0x40, 0x09, 0x84, 0xb9, 0x88, 0x9d, 0x4d, 0xbe, 0x1b, 0xaa, 0x28, 0xa8,
	0xd7, 0xb3, 0x8d, 0xeb, 0x70, 0x14, 0x6d, 0xf1, 0x99, 0x07, 0x81, 0x6a, 0x4f, 0xc3, 0xdb, 0xa0,
	0x9b, 0xdf, 0x4e, 0xad, 0x9f, 0x1c, 0xf3, 0x76, 0x32, 0x9f, 0x09, 0xfc, 0xf6, 0x9f, 0xd9, 0x1f,
	0x94, 0x87, 0xb6, 0x49, 0x7f, 0xab, 0x8b, 0x83, 0xbe, 0xe8, 0x96, 0xdd, 0xea, 0x6a, 0xf3, 0x66,
	0x4e, 0x6b, 0x35, 0x47, 0xab, 0x5a, 0x3c, 0x4e, 0x78, 0xdf, 0xda, 0xb4, 0x56, 0xe7, 0xc9, 0x71,
	0x5a, 0x9c, 0x2d, 0xbb, 0xd5, 0x65, 0x3f, 0xdd, 0xf7, 0xaf, 0xbb, 0x97, 0x98, 0x25, 0x8d, 0x1f,
	0xc6, 0xe3, 0xde, 0x6e, 0xee, 0x58, 0x32, 0xda, 0x88, 0xb7, 0x34, 0xb8, 0x9f, 0x58, 0x71, 0x2c,
	0xc1, 0xc6, 0x4a, 0xde, 0xb2, 0x4f, 0xb0, 0x77, 0xf3, 0x5f, 0xf5, 0x17, 0x00, 0x00, 0xff, 0xff,
	0xfd, 0x3d, 0x01, 0x60, 0x00, 0x02, 0x00, 0x00,
}
