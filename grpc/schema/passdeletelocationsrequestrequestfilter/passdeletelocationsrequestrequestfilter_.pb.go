// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passdeletelocationsrequestrequestfilter/passdeletelocationsrequestrequestfilter_.proto

/*
Package passdeletelocationsrequestrequestfilter is a generated protocol buffer package.

It is generated from these files:
	passdeletelocationsrequestrequestfilter/passdeletelocationsrequestrequestfilter_.proto

It has these top-level messages:
	PassDeleteLocationsRequestRequestFilter
*/
package passdeletelocationsrequestrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import deletelocationsrequestfilter "github.com/21stio/go-playground/grpc/schema/deletelocationsrequestfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassDeleteLocationsRequestRequestFilter struct {
	Meta                   *requestmetafilter.RequestMetaFilter                       `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	DeleteLocationsRequest *deletelocationsrequestfilter.DeleteLocationsRequestFilter `protobuf:"bytes,2,opt,name=deleteLocationsRequest" json:"deleteLocationsRequest,omitempty"`
	And                    []*PassDeleteLocationsRequestRequestFilter                 `protobuf:"bytes,3,rep,name=and" json:"and,omitempty"`
	Or                     []*PassDeleteLocationsRequestRequestFilter                 `protobuf:"bytes,4,rep,name=or" json:"or,omitempty"`
	Not                    []*PassDeleteLocationsRequestRequestFilter                 `protobuf:"bytes,5,rep,name=not" json:"not,omitempty"`
	Hash                   *string                                                    `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized       []byte                                                     `json:"-"`
}

func (m *PassDeleteLocationsRequestRequestFilter) Reset() {
	*m = PassDeleteLocationsRequestRequestFilter{}
}
func (m *PassDeleteLocationsRequestRequestFilter) String() string { return proto.CompactTextString(m) }
func (*PassDeleteLocationsRequestRequestFilter) ProtoMessage()    {}
func (*PassDeleteLocationsRequestRequestFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassDeleteLocationsRequestRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassDeleteLocationsRequestRequestFilter) GetDeleteLocationsRequest() *deletelocationsrequestfilter.DeleteLocationsRequestFilter {
	if m != nil {
		return m.DeleteLocationsRequest
	}
	return nil
}

func (m *PassDeleteLocationsRequestRequestFilter) GetAnd() []*PassDeleteLocationsRequestRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PassDeleteLocationsRequestRequestFilter) GetOr() []*PassDeleteLocationsRequestRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PassDeleteLocationsRequestRequestFilter) GetNot() []*PassDeleteLocationsRequestRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PassDeleteLocationsRequestRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassDeleteLocationsRequestRequestFilter)(nil), "passdeletelocationsrequestrequestfilter.PassDeleteLocationsRequestRequestFilter")
}

func init() {
	proto.RegisterFile("passdeletelocationsrequestrequestfilter/passdeletelocationsrequestrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xbd, 0x4b, 0xc4, 0x30,
	0x14, 0xa7, 0x1f, 0x0a, 0xe6, 0xb6, 0x0c, 0x12, 0x6e, 0x2a, 0x22, 0x5c, 0x11, 0x4c, 0xb1, 0x93,
	0x38, 0x89, 0x88, 0x93, 0xc2, 0x51, 0xd1, 0xc1, 0x45, 0x73, 0x6d, 0x6c, 0x0b, 0x6d, 0x5e, 0x4d,
	0x5e, 0x07, 0xff, 0x10, 0xff, 0x5f, 0x69, 0x2e, 0x4e, 0x57, 0x8f, 0x2c, 0x37, 0xbd, 0x7c, 0xfc,
	0xbe, 0xde, 0x83, 0x47, 0x5e, 0x07, 0x61, 0x4c, 0x25, 0x3b, 0x89, 0xb2, 0x83, 0x52, 0x60, 0x0b,
	0xca, 0x68, 0xf9, 0x35, 0x4a, 0x83, 0xae, 0x7c, 0xb6, 0x1d, 0x4a, 0x9d, 0x79, 0xe2, 0xde, 0xf9,
	0xa0, 0x01, 0x81, 0xae, 0x3c, 0xf1, 0xcb, 0x0b, 0x77, 0xed, 0x25, 0x0a, 0x67, 0xb5, 0xf3, 0xe2,
	0x44, 0x97, 0xb7, 0xf3, 0x82, 0x8e, 0xb6, 0xef, 0xd3, 0x29, 0x9c, 0xfd, 0xc4, 0x64, 0xb5, 0x16,
	0xc6, 0xdc, 0x5b, 0xec, 0xe3, 0x1f, 0xb6, 0xd8, 0x62, 0x5d, 0x79, 0xb0, 0x14, 0x7a, 0x4d, 0xe2,
	0x29, 0x02, 0x0b, 0x92, 0x20, 0x5d, 0xe4, 0xe7, 0x7c, 0x27, 0x16, 0x77, 0xf8, 0x27, 0x89, 0x62,
	0xcb, 0x29, 0x2c, 0x83, 0x6a, 0x72, 0x5a, 0xcd, 0x1a, 0xb0, 0xd0, 0x6a, 0xdd, 0xf0, 0x7d, 0x59,
	0xf9, 0x7c, 0x38, 0xe7, 0xf0, 0x8f, 0x32, 0xdd, 0x90, 0x48, 0xa8, 0x8a, 0x45, 0x49, 0x94, 0x2e,
	0xf2, 0x35, 0xf7, 0x1c, 0x3f, 0xf7, 0x1c, 0x46, 0x31, 0x89, 0xd3, 0x0f, 0x12, 0x82, 0x66, 0xf1,
	0x81, 0x2c, 0x42, 0xd0, 0x53, 0x17, 0x0a, 0x90, 0x1d, 0x1d, 0xaa, 0x0b, 0x05, 0x48, 0x29, 0x89,
	0x1b, 0x61, 0x1a, 0x76, 0x9c, 0x04, 0xe9, 0x49, 0x61, 0xcf, 0x77, 0x2f, 0x6f, 0xcf, 0x75, 0x8b,
	0xcd, 0xb8, 0xe1, 0x25, 0xf4, 0x59, 0x7e, 0x65, 0xb0, 0x85, 0xac, 0x86, 0xcb, 0xa1, 0x13, 0xdf,
	0xb5, 0x86, 0x51, 0x55, 0x59, 0xad, 0x87, 0x32, 0x33, 0x65, 0x23, 0x7b, 0xe1, 0xbb, 0x0c, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xce, 0x5d, 0x0a, 0x5e, 0x03, 0x00, 0x00,
}
