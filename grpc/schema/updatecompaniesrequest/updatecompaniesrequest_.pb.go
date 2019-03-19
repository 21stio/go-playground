// Code generated by protoc-gen-go. DO NOT EDIT.
// source: updatecompaniesrequest/updatecompaniesrequest_.proto

/*
Package updatecompaniesrequest is a generated protocol buffer package.

It is generated from these files:
	updatecompaniesrequest/updatecompaniesrequest_.proto

It has these top-level messages:
	UpdateCompaniesRequest
*/
package updatecompaniesrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import company "github.com/21stio/go-playground/grpc/schema/company"
import updatecompaniesresponseselect "github.com/21stio/go-playground/grpc/schema/updatecompaniesresponseselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpdateCompaniesRequest struct {
	Meta             *requestmeta.RequestMeta                                     `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	ServiceFilter    *servicefilter.ServiceFilter                                 `protobuf:"bytes,2,opt,name=serviceFilter" json:"serviceFilter,omitempty"`
	Companies        []*company.Company                                           `protobuf:"bytes,3,rep,name=companies" json:"companies,omitempty"`
	Select           *updatecompaniesresponseselect.UpdateCompaniesResponseSelect `protobuf:"bytes,4,opt,name=select" json:"select,omitempty"`
	Hash             *string                                                      `protobuf:"bytes,5,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                       `json:"-"`
}

func (m *UpdateCompaniesRequest) Reset()                    { *m = UpdateCompaniesRequest{} }
func (m *UpdateCompaniesRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateCompaniesRequest) ProtoMessage()               {}
func (*UpdateCompaniesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UpdateCompaniesRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *UpdateCompaniesRequest) GetServiceFilter() *servicefilter.ServiceFilter {
	if m != nil {
		return m.ServiceFilter
	}
	return nil
}

func (m *UpdateCompaniesRequest) GetCompanies() []*company.Company {
	if m != nil {
		return m.Companies
	}
	return nil
}

func (m *UpdateCompaniesRequest) GetSelect() *updatecompaniesresponseselect.UpdateCompaniesResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *UpdateCompaniesRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateCompaniesRequest)(nil), "updatecompaniesrequest.UpdateCompaniesRequest")
}

func init() {
	proto.RegisterFile("updatecompaniesrequest/updatecompaniesrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xf4, 0x30,
	0x10, 0xc6, 0xd9, 0x3f, 0xef, 0x0b, 0x9b, 0x45, 0x90, 0x1c, 0x4a, 0x58, 0x44, 0xca, 0x9e, 0x7a,
	0xd0, 0x04, 0x8b, 0x47, 0x4f, 0x55, 0xbc, 0x79, 0xc9, 0xea, 0xc5, 0x8b, 0xc4, 0xee, 0xd8, 0x16,
	0xda, 0x26, 0x26, 0xa9, 0xd0, 0xef, 0xe3, 0x07, 0x95, 0x4d, 0x53, 0x6d, 0xd9, 0xea, 0xa9, 0xd3,
	0x79, 0x9e, 0x99, 0xf9, 0x3d, 0x10, 0x74, 0xdd, 0xa8, 0xbd, 0xb0, 0x90, 0xca, 0x4a, 0x89, 0xba,
	0x00, 0xa3, 0xe1, 0xbd, 0x01, 0x63, 0xd9, 0x74, 0xfb, 0x85, 0x2a, 0x2d, 0xad, 0xc4, 0xc1, 0xb4,
	0xbc, 0x39, 0xf7, 0x45, 0x05, 0x56, 0xb0, 0x41, 0xed, 0xe7, 0x36, 0x5b, 0x03, 0xfa, 0xa3, 0x48,
	0xe1, 0xad, 0x28, 0x2d, 0x68, 0x36, 0xfa, 0xeb, 0x3d, 0x41, 0xb7, 0xb5, 0x65, 0xfe, 0xdb, 0xf7,
	0x93, 0xa3, 0x9b, 0x46, 0xc9, 0xda, 0x80, 0x81, 0x12, 0xd2, 0x09, 0xe0, 0xa1, 0xea, 0x77, 0x6c,
	0x3f, 0xe7, 0x28, 0x78, 0x72, 0xc6, 0xdb, 0xde, 0xc8, 0x3b, 0x4a, 0x7c, 0x81, 0x96, 0x07, 0x52,
	0x32, 0x0b, 0x67, 0xd1, 0x3a, 0x26, 0x74, 0x40, 0x4f, 0xbd, 0xe7, 0x01, 0xac, 0xe0, 0xce, 0x85,
	0x13, 0x74, 0xe2, 0xe1, 0xef, 0x1d, 0x3c, 0x99, 0xbb, 0xb1, 0x33, 0x3a, 0x8a, 0x44, 0x77, 0x43,
	0x0f, 0x1f, 0x8f, 0x60, 0x8a, 0x56, 0xdf, 0xb8, 0x64, 0x11, 0x2e, 0xa2, 0x75, 0x7c, 0x4a, 0x7d,
	0x68, 0xda, 0xf1, 0xb5, 0xfc, 0xc7, 0x82, 0x1f, 0xd1, 0xff, 0x2e, 0x0d, 0x59, 0xba, 0x63, 0x37,
	0xf4, 0xcf, 0xcc, 0xf4, 0x28, 0x68, 0xa7, 0xee, 0x9c, 0xca, 0xfd, 0x2e, 0x8c, 0xd1, 0x32, 0x17,
	0x26, 0x27, 0xff, 0xc2, 0x59, 0xb4, 0xe2, 0xae, 0x4e, 0xee, 0x9e, 0x93, 0xac, 0xb0, 0x79, 0xf3,
	0x7a, 0xc0, 0x61, 0xf1, 0x95, 0xb1, 0x85, 0x64, 0x99, 0xbc, 0x54, 0xa5, 0x68, 0x33, 0x2d, 0x9b,
	0x7a, 0xcf, 0x32, 0xad, 0x52, 0x66, 0xd2, 0x1c, 0x2a, 0xf1, 0xcb, 0x5b, 0xf9, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xba, 0x16, 0xc4, 0xa1, 0x5b, 0x02, 0x00, 0x00,
}