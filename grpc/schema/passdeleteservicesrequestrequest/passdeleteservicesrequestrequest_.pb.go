// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeleteservicesrequestrequest/passdeleteservicesrequestrequest_.proto

/*
Package passdeleteservicesrequestrequest is a generated protocol buffer package.

It is generated from these files:
	passdeleteservicesrequestrequest/passdeleteservicesrequestrequest_.proto

It has these top-level messages:
	PassDeleteServicesRequestRequest
*/
package passdeleteservicesrequestrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import deleteservicesrequest "github.com/21stio/go-playground/grpc/schema/deleteservicesrequest"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteServicesRequestRequest struct {
	Meta                  *requestmeta.RequestMeta                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter         *servicefilter.ServiceFilter                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	DeleteServicesRequest *deleteservicesrequest.DeleteServicesRequest `protobuf:"bytes,3,opt,name=deleteServicesRequest" json:"deleteServicesRequest,omitempty"`
	Hash                  *string                                      `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized      []byte                                       `json:"-"`
}

func (m *PassDeleteServicesRequestRequest) Reset()         { *m = PassDeleteServicesRequestRequest{} }
func (m *PassDeleteServicesRequestRequest) String() string { return proto.CompactTextString(m) }
func (*PassDeleteServicesRequestRequest) ProtoMessage()    {}
func (*PassDeleteServicesRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteServicesRequestRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteServicesRequestRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *PassDeleteServicesRequestRequest) GetDeleteServicesRequest() *deleteservicesrequest.DeleteServicesRequest {
	if m != nil {
		return m.DeleteServicesRequest
	}
	return nil
}

func (m *PassDeleteServicesRequestRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteServicesRequestRequest)(nil), "passdeleteservicesrequestrequest.PassDeleteServicesRequestRequest")
}

func init() {
	proto.RegisterFile("passdeleteservicesrequestrequest/passdeleteservicesrequestrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xe9, 0xda, 0x8b, 0x11, 0x2f, 0x01, 0x21, 0x2c, 0x22, 0x65, 0x4f, 0x7b, 0x58, 0x13,
	0xec, 0x4f, 0x58, 0x44, 0x3c, 0xa8, 0x48, 0xf7, 0xe6, 0x45, 0xb2, 0xed, 0xd8, 0x16, 0x5a, 0x53,
	0x33, 0x53, 0xc1, 0xff, 0xe0, 0x8f, 0x16, 0xe3, 0x08, 0x2d, 0x04, 0x7a, 0x9a, 0xd7, 0xf2, 0xbd,
	0x37, 0x6f, 0x88, 0xb8, 0x1f, 0x2c, 0x62, 0x05, 0x1d, 0x10, 0x20, 0xf8, 0xcf, 0xb6, 0x04, 0xf4,
	0xf0, 0x31, 0x02, 0x12, 0x0f, 0xb3, 0x04, 0xbc, 0xea, 0xc1, 0x3b, 0x72, 0x32, 0x5b, 0x02, 0xd7,
	0x57, 0x2c, 0x7a, 0x20, 0x6b, 0x26, 0x9a, 0x13, 0xd6, 0x1b, 0xf6, 0xbd, 0xb5, 0x1d, 0x81, 0x37,
	0xb3, 0xaf, 0x7f, 0x26, 0x8f, 0x6e, 0x30, 0xd1, 0xbf, 0xec, 0xd9, 0x7c, 0xaf, 0x44, 0xf6, 0x6c,
	0x11, 0x6f, 0x03, 0x74, 0x60, 0xa8, 0xf8, 0x83, 0x78, 0xc8, 0x9d, 0x48, 0x7f, 0xbb, 0xa8, 0x24,
	0x4b, 0xb6, 0x67, 0xb9, 0xd2, 0x93, 0x7e, 0x9a, 0x99, 0x47, 0x20, 0x5b, 0x04, 0x4a, 0xee, 0xc5,
	0x39, 0x2f, 0xbb, 0x0b, 0xf5, 0xd4, 0x2a, 0xd8, 0x2e, 0xf5, 0xac, 0xb4, 0x3e, 0x4c, 0x99, 0x62,
	0x6e, 0x91, 0x47, 0x71, 0x51, 0xc5, 0x1a, 0xa9, 0x93, 0x90, 0xb5, 0xd3, 0xd1, 0xa3, 0x74, 0xfc,
	0x8a, 0x78, 0x94, 0x94, 0x22, 0x6d, 0x2c, 0x36, 0x2a, 0xcd, 0x92, 0xed, 0x69, 0x11, 0xf4, 0xfe,
	0xe9, 0xe5, 0xa1, 0x6e, 0xa9, 0x19, 0x8f, 0xba, 0x74, 0xbd, 0xc9, 0x6f, 0x90, 0x5a, 0x67, 0x6a,
	0x77, 0x3d, 0x74, 0xf6, 0xab, 0xf6, 0x6e, 0x7c, 0xaf, 0x4c, 0xed, 0x87, 0xd2, 0x60, 0xd9, 0x40,
	0x6f, 0x17, 0xdf, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x91, 0xd1, 0x53, 0x67, 0x43, 0x02, 0x00,
	0x00,
}
