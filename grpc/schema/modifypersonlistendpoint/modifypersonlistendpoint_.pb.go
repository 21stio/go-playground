// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifypersonlistendpoint/modifypersonlistendpoint_.proto

/*
Package modifypersonlistendpoint is a generated protocol buffer package.

It is generated from these files:
	modifypersonlistendpoint/modifypersonlistendpoint_.proto

It has these top-level messages:
	ModifyPersonListEndpoint
*/
package modifypersonlistendpoint

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import modifypersonlistrequestfilter "github.com/21stio/go-playground/grpc/schema/modifypersonlistrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ModifyPersonListEndpoint struct {
	Filter           *modifypersonlistrequestfilter.ModifyPersonListRequestFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Hash             *string                                                      `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                       `json:"-"`
}

func (m *ModifyPersonListEndpoint) Reset()                    { *m = ModifyPersonListEndpoint{} }
func (m *ModifyPersonListEndpoint) String() string            { return proto.CompactTextString(m) }
func (*ModifyPersonListEndpoint) ProtoMessage()               {}
func (*ModifyPersonListEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ModifyPersonListEndpoint) GetFilter() *modifypersonlistrequestfilter.ModifyPersonListRequestFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *ModifyPersonListEndpoint) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifyPersonListEndpoint)(nil), "modifypersonlistendpoint.ModifyPersonListEndpoint")
}

func init() {
	proto.RegisterFile("modifypersonlistendpoint/modifypersonlistendpoint_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc8, 0xcd, 0x4f, 0xc9,
	0x4c, 0xab, 0x2c, 0x48, 0x2d, 0x2a, 0xce, 0xcf, 0xcb, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x29,
	0xc8, 0xcf, 0xcc, 0x2b, 0xd1, 0xc7, 0x25, 0x11, 0xaf, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f, 0x24,
	0x81, 0x4b, 0x81, 0x94, 0x13, 0xba, 0x4c, 0x51, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x49, 0x5a, 0x66,
	0x4e, 0x49, 0x6a, 0x91, 0x3e, 0x5e, 0x59, 0xa8, 0xe9, 0x4a, 0x2d, 0x8c, 0x5c, 0x12, 0xbe, 0x60,
	0x85, 0x01, 0x60, 0x85, 0x3e, 0x99, 0xc5, 0x25, 0xae, 0x50, 0x0b, 0x84, 0x42, 0xb8, 0xd8, 0x20,
	0xaa, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x6c, 0xf4, 0xf0, 0x9a, 0xa9, 0x87, 0x6e, 0x50,
	0x10, 0x44, 0xd6, 0x0d, 0x2c, 0x1b, 0x04, 0x35, 0x4b, 0x48, 0x88, 0x8b, 0x25, 0x23, 0xb1, 0x38,
	0x43, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x76, 0x72, 0x8b, 0x72, 0x49, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x37, 0x32, 0x2c, 0x2e, 0xc9, 0xcc, 0xd7, 0x4f,
	0xcf, 0xd7, 0x2d, 0xc8, 0x49, 0xac, 0x4c, 0x2f, 0xca, 0x2f, 0xcd, 0x4b, 0xd1, 0x4f, 0x2f, 0x2a,
	0x48, 0xd6, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0xc4, 0x19, 0x66, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x37, 0x0c, 0x7a, 0x22, 0x67, 0x01, 0x00, 0x00,
}
