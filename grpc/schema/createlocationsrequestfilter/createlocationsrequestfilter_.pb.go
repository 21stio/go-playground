// Code generated by protoc-gen-go. DO NOT EDIT.
// source: createlocationsrequestfilter/createlocationsrequestfilter_.proto

/*
Package createlocationsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	createlocationsrequestfilter/createlocationsrequestfilter_.proto

It has these top-level messages:
	CreateLocationsRequestFilter
*/
package createlocationsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import locationfilter "github.com/21stio/go-playground/grpc/schema/locationfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateLocationsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	LocationsSome    *locationfilter.LocationFilter       `protobuf:"bytes,2,opt,name=locationsSome" json:"locationsSome,omitempty"`
	LocationsEvery   *locationfilter.LocationFilter       `protobuf:"bytes,3,opt,name=locationsEvery" json:"locationsEvery,omitempty"`
	LocationsNone    *locationfilter.LocationFilter       `protobuf:"bytes,4,opt,name=locationsNone" json:"locationsNone,omitempty"`
	And              []*CreateLocationsRequestFilter      `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*CreateLocationsRequestFilter      `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*CreateLocationsRequestFilter      `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *CreateLocationsRequestFilter) Reset()                    { *m = CreateLocationsRequestFilter{} }
func (m *CreateLocationsRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*CreateLocationsRequestFilter) ProtoMessage()               {}
func (*CreateLocationsRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateLocationsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CreateLocationsRequestFilter) GetLocationsSome() *locationfilter.LocationFilter {
	if m != nil {
		return m.LocationsSome
	}
	return nil
}

func (m *CreateLocationsRequestFilter) GetLocationsEvery() *locationfilter.LocationFilter {
	if m != nil {
		return m.LocationsEvery
	}
	return nil
}

func (m *CreateLocationsRequestFilter) GetLocationsNone() *locationfilter.LocationFilter {
	if m != nil {
		return m.LocationsNone
	}
	return nil
}

func (m *CreateLocationsRequestFilter) GetAnd() []*CreateLocationsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *CreateLocationsRequestFilter) GetOr() []*CreateLocationsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *CreateLocationsRequestFilter) GetNot() []*CreateLocationsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *CreateLocationsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateLocationsRequestFilter)(nil), "createlocationsrequestfilter.CreateLocationsRequestFilter")
}

func init() {
	proto.RegisterFile("createlocationsrequestfilter/createlocationsrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4b, 0x4b, 0xc3, 0x40,
	0x14, 0x85, 0xc9, 0xc3, 0xd7, 0x14, 0x5d, 0xcc, 0x6a, 0x28, 0x45, 0x82, 0x74, 0x11, 0x04, 0x27,
	0x98, 0x95, 0xb8, 0x12, 0x1f, 0x45, 0x4a, 0x75, 0x11, 0x77, 0x6e, 0x64, 0x4c, 0xc7, 0x24, 0x90,
	0xcc, 0x8d, 0x93, 0x1b, 0xa1, 0xff, 0xc8, 0x9f, 0x29, 0x49, 0x27, 0x85, 0x58, 0x08, 0x4a, 0x77,
	0xb9, 0x37, 0xe7, 0x7c, 0xe7, 0x24, 0x33, 0xe4, 0x26, 0xd6, 0x52, 0xa0, 0xcc, 0x21, 0x16, 0x98,
	0x81, 0xaa, 0xb4, 0xfc, 0xac, 0x65, 0x85, 0x1f, 0x59, 0x8e, 0x52, 0x07, 0x43, 0x2f, 0xdf, 0x78,
	0xa9, 0x01, 0x81, 0x4e, 0x86, 0x44, 0xe3, 0x73, 0x33, 0x16, 0x12, 0x85, 0x81, 0x6e, 0x6d, 0x0c,
	0x69, 0x3c, 0xed, 0x18, 0x46, 0xd8, 0x1f, 0x8d, 0xea, 0xec, 0xdb, 0x25, 0x93, 0xbb, 0x36, 0x72,
	0xd1, 0x45, 0x46, 0x6b, 0xe2, 0xac, 0xd5, 0xd1, 0x2b, 0xe2, 0x36, 0x6c, 0x66, 0x79, 0x96, 0x3f,
	0x0a, 0xa7, 0x7c, 0x2b, 0x8f, 0x1b, 0xfd, 0x93, 0x44, 0xb1, 0xf6, 0x44, 0xad, 0x83, 0xde, 0x93,
	0xe3, 0xcd, 0x67, 0xbc, 0x40, 0x21, 0x99, 0xdd, 0x22, 0x4e, 0x79, 0xbf, 0x09, 0xef, 0x82, 0x8d,
	0xb9, 0x6f, 0xa2, 0x33, 0x72, 0xb2, 0x59, 0x3c, 0x7c, 0x49, 0xbd, 0x62, 0xce, 0x9f, 0x30, 0xbf,
	0x5c, 0xbd, 0x36, 0xcf, 0xa0, 0x24, 0x73, 0xff, 0xd9, 0xa6, 0x31, 0xd1, 0x05, 0x71, 0x84, 0x5a,
	0xb2, 0x3d, 0xcf, 0xf1, 0x47, 0xe1, 0x35, 0x1f, 0x3a, 0x2c, 0x3e, 0xf4, 0x5b, 0xa3, 0x06, 0x43,
	0xe7, 0xc4, 0x06, 0xcd, 0xf6, 0x77, 0x86, 0xd9, 0xa0, 0x9b, 0x66, 0x0a, 0x90, 0x1d, 0xec, 0xde,
	0x4c, 0x01, 0x52, 0x4a, 0xdc, 0x54, 0x54, 0x29, 0x3b, 0xf4, 0x2c, 0xff, 0x28, 0x6a, 0x9f, 0x6f,
	0xe7, 0xaf, 0x8f, 0x49, 0x86, 0x69, 0xfd, 0xce, 0x63, 0x28, 0x82, 0xf0, 0xb2, 0xc2, 0x0c, 0x82,
	0x04, 0x2e, 0xca, 0x5c, 0xac, 0x12, 0x0d, 0xb5, 0x5a, 0x06, 0x89, 0x2e, 0xe3, 0xa0, 0x8a, 0x53,
	0x59, 0x88, 0xc1, 0xdb, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x76, 0x89, 0xac, 0x6f, 0x29, 0x03,
	0x00, 0x00,
}
