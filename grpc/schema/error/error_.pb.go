// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error/error_.proto

/*
Package error is a generated protocol buffer package.

It is generated from these files:
	error/error_.proto

It has these top-level messages:
	Error
*/
package error

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import errorkind "github.com/21stio/go-playground/grpc/schema/errorkind"
import endpointkind "github.com/21stio/go-playground/grpc/schema/endpointkind"
import service "github.com/21stio/go-playground/grpc/schema/service"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Error struct {
	Kind             *errorkind.ErrorKind       `protobuf:"varint,1,opt,name=kind,enum=errorkind.ErrorKind" json:"kind,omitempty"`
	Endpoint         *endpointkind.EndpointKind `protobuf:"varint,2,opt,name=endpoint,enum=endpointkind.EndpointKind" json:"endpoint,omitempty"`
	Message          *string                    `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Service          *service.Service           `protobuf:"bytes,4,opt,name=service" json:"service,omitempty"`
	Hash             *string                    `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Error) GetKind() errorkind.ErrorKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return errorkind.ErrorKind_ERROR_SERVICE
}

func (m *Error) GetEndpoint() endpointkind.EndpointKind {
	if m != nil && m.Endpoint != nil {
		return *m.Endpoint
	}
	return endpointkind.EndpointKind_ENDPOINT_AUTHENTICATE_CREDENTIAL_ENDPOINT
}

func (m *Error) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Error) GetService() *service.Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Error) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "error.Error")
}

func init() { proto.RegisterFile("error/error_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x89, 0xb6, 0xac, 0x46, 0x10, 0x09, 0x22, 0xa1, 0xa7, 0xe2, 0xa9, 0x08, 0x4e, 0xd8,
	0x0a, 0x3e, 0x80, 0xb0, 0x27, 0x6f, 0xf5, 0xe6, 0x45, 0x6a, 0x1b, 0xd2, 0xa0, 0xcd, 0x94, 0xa4,
	0x2b, 0xf8, 0x7c, 0xbe, 0x98, 0xec, 0x34, 0x29, 0xbd, 0xe4, 0x9f, 0x3f, 0xf3, 0xcd, 0xfc, 0x21,
	0x5c, 0x68, 0xef, 0xd1, 0x2b, 0x3a, 0x3f, 0x60, 0xf2, 0x38, 0xa3, 0xc8, 0xc9, 0x15, 0x05, 0xc9,
	0x97, 0x75, 0xbd, 0x5a, 0xab, 0x88, 0x14, 0xa5, 0x76, 0xfd, 0x84, 0xd6, 0xcd, 0x4b, 0x7b, 0x63,
	0x12, 0x71, 0x17, 0xb4, 0xff, 0xb1, 0x9d, 0x56, 0x51, 0xe3, 0xfd, 0xfd, 0x1f, 0xe3, 0xf9, 0xe1,
	0xb4, 0x4e, 0x54, 0x3c, 0x3b, 0x0d, 0x48, 0x56, 0xb2, 0xea, 0xba, 0xbe, 0x85, 0x35, 0x04, 0xa8,
	0xff, 0x6a, 0x5d, 0xdf, 0x10, 0x21, 0x9e, 0xf9, 0x45, 0x8a, 0x90, 0x67, 0x44, 0x17, 0xb0, 0xcd,
	0x84, 0x43, 0x34, 0x34, 0xb3, 0xb2, 0x42, 0xf2, 0xdd, 0xa8, 0x43, 0x68, 0x8d, 0x96, 0xe7, 0x25,
	0xab, 0x2e, 0x9b, 0x64, 0xc5, 0x03, 0xdf, 0xc5, 0x77, 0xc9, 0xac, 0x64, 0xd5, 0x55, 0x7d, 0x03,
	0xd1, 0xc3, 0xdb, 0xa2, 0x4d, 0x02, 0x84, 0xe0, 0xd9, 0xd0, 0x86, 0x41, 0xe6, 0xb4, 0x82, 0xea,
	0x97, 0xa7, 0xf7, 0xbd, 0xb1, 0xf3, 0x70, 0xfc, 0x84, 0x0e, 0x47, 0x55, 0xef, 0xc3, 0x6c, 0x51,
	0x19, 0x7c, 0x9c, 0xbe, 0xdb, 0x5f, 0xe3, 0xf1, 0xe8, 0x7a, 0x65, 0xfc, 0xd4, 0xa9, 0xd0, 0x0d,
	0x7a, 0x6c, 0x97, 0xff, 0xfb, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x14, 0x48, 0x55, 0xcb, 0x6c, 0x01,
	0x00, 0x00,
}
