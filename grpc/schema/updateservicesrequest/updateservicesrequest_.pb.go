// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updateservicesrequest/updateservicesrequest_.proto

/*
Package updateservicesrequest is a generated protocol buffer package.

It is generated from these files:
	updateservicesrequest/updateservicesrequest_.proto

It has these top-level messages:
	UpdateServicesRequest
*/
package updateservicesrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import service "github.com/21stio/go-playground/grpc/schema/service"
import updateservicesresponseselect "github.com/21stio/go-playground/grpc/schema/updateservicesresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateServicesRequest struct {
	Meta             *requestmeta.RequestMeta                                   `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                               `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Services         []*service.Service                                         `protobuf:"bytes,3,rep,name=services" json:"services,omitempty"`
	Select           *updateservicesresponseselect.UpdateServicesResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                    `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                     `json:"-"`
}

func (m *UpdateServicesRequest) Reset()                    { *m = UpdateServicesRequest{} }
func (m *UpdateServicesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateServicesRequest) ProtoMessage()               {}
func (*UpdateServicesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateServicesRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateServicesRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *UpdateServicesRequest) GetServices() []*service.Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *UpdateServicesRequest) GetSelect() *updateservicesresponseselect.UpdateServicesResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *UpdateServicesRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateServicesRequest)(nil), "updateservicesrequest.UpdateServicesRequest")
}

func init() { proto.RegisterFile("updateservicesrequest/updateservicesrequest_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xe9, 0x4c, 0xff, 0x1f, 0xcd, 0x20, 0x48, 0x60, 0x24, 0x0c, 0x22, 0x65, 0x56, 0x5d,
	0x8c, 0x09, 0x76, 0xe9, 0x4a, 0x47, 0x70, 0xe7, 0x26, 0xc5, 0x8d, 0x1b, 0x89, 0x9d, 0x6b, 0x5b,
	0x68, 0x27, 0x31, 0x49, 0x05, 0x9f, 0xc6, 0x57, 0x15, 0xd3, 0x44, 0xda, 0xb1, 0xb8, 0xbb, 0xbd,
	0xe7, 0x3b, 0xb9, 0xe7, 0x40, 0x51, 0xd6, 0xa9, 0x9d, 0xb0, 0x60, 0x40, 0xbf, 0xd7, 0x05, 0x18,
	0x0d, 0x6f, 0x1d, 0x18, 0xcb, 0x26, 0xb7, 0xcf, 0x54, 0x69, 0x69, 0x25, 0x5e, 0x4e, 0xaa, 0xab,
	0x0b, 0x3f, 0xb4, 0x60, 0x05, 0x1b, 0xcc, 0xde, 0xb6, 0x5a, 0x7b, 0xc3, 0x6b, 0xdd, 0x58, 0xd0,
	0x6c, 0xf4, 0x15, 0x98, 0x33, 0xbf, 0x0d, 0x6a, 0xd8, 0xdf, 0x1c, 0x9e, 0x34, 0x4a, 0xee, 0x0d,
	0x18, 0x68, 0xa0, 0xf8, 0x9d, 0x76, 0x28, 0xfa, 0x17, 0xd6, 0x9f, 0x33, 0xb4, 0x7c, 0x74, 0x5c,
	0xee, 0x39, 0xde, 0x47, 0xc4, 0x1b, 0x14, 0x7f, 0xc7, 0x24, 0x51, 0x12, 0xa5, 0x8b, 0x8c, 0xd0,
	0x41, 0x74, 0xea, 0x99, 0x07, 0xb0, 0x82, 0x3b, 0x0a, 0x6f, 0xd1, 0x89, 0x3f, 0x74, 0xef, 0x92,
	0x93, 0x99, 0xb3, 0x9d, 0xd3, 0x51, 0x1f, 0x9a, 0x0f, 0x19, 0x3e, 0xb6, 0xe0, 0x0d, 0x3a, 0x0a,
	0x61, 0xc9, 0x3c, 0x99, 0xa7, 0x8b, 0xec, 0x34, 0xd8, 0x83, 0x91, 0xff, 0x10, 0x98, 0xa3, 0xff,
	0x7d, 0x15, 0x12, 0xbb, 0x53, 0xd7, 0xf4, 0xaf, 0xbe, 0xf4, 0xb0, 0x64, 0x2f, 0xe6, 0x4e, 0xe4,
	0xfe, 0x25, 0x8c, 0x51, 0x5c, 0x09, 0x53, 0x91, 0x7f, 0x49, 0x94, 0x1e, 0x73, 0x37, 0x6f, 0xef,
	0x9e, 0x6e, 0xcb, 0xda, 0x56, 0xdd, 0x0b, 0x2d, 0x64, 0xcb, 0xb2, 0x2b, 0x63, 0x6b, 0xc9, 0x4a,
	0x79, 0xa9, 0x1a, 0xf1, 0x51, 0x6a, 0xd9, 0xed, 0x77, 0xac, 0xd4, 0xaa, 0x60, 0xa6, 0xa8, 0xa0,
	0x15, 0xd3, 0xbf, 0xc8, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x05, 0x10, 0x2e, 0x50, 0x02,
	0x00, 0x00,
}
