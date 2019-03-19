// Code generated by protoc-gen-go. DO NOT EDIT.
// source: locationquery/locationquery_.proto

/*
Package locationquery is a generated protocol buffer package.

It is generated from these files:
	locationquery/locationquery_.proto

It has these top-level messages:
	LocationQuery
*/
package locationquery

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import lengthvalue "github.com/21stio/go-playground/grpc/schema/lengthvalue"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LocationQuery struct {
	RadiusLt             *lengthvalue.LengthValue `protobuf:"bytes,1,opt,name=radiusLt" json:"radiusLt,omitempty"`
	Street               *string                  `protobuf:"bytes,2,opt,name=street" json:"street,omitempty"`
	ZipCode              *string                  `protobuf:"bytes,3,opt,name=zipCode" json:"zipCode,omitempty"`
	City                 *string                  `protobuf:"bytes,4,opt,name=city" json:"city,omitempty"`
	CityDistrict         *string                  `protobuf:"bytes,5,opt,name=cityDistrict" json:"cityDistrict,omitempty"`
	Country              *string                  `protobuf:"bytes,6,opt,name=country" json:"country,omitempty"`
	CountryState         *string                  `protobuf:"bytes,7,opt,name=countryState" json:"countryState,omitempty"`
	CountryStateDistrict *string                  `protobuf:"bytes,8,opt,name=countryStateDistrict" json:"countryStateDistrict,omitempty"`
	Hash                 *string                  `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized     []byte                   `json:"-"`
}

func (m *LocationQuery) Reset()                    { *m = LocationQuery{} }
func (m *LocationQuery) String() string            { return proto.CompactTextString(m) }
func (*LocationQuery) ProtoMessage()               {}
func (*LocationQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LocationQuery) GetRadiusLt() *lengthvalue.LengthValue {
	if m != nil {
		return m.RadiusLt
	}
	return nil
}

func (m *LocationQuery) GetStreet() string {
	if m != nil && m.Street != nil {
		return *m.Street
	}
	return ""
}

func (m *LocationQuery) GetZipCode() string {
	if m != nil && m.ZipCode != nil {
		return *m.ZipCode
	}
	return ""
}

func (m *LocationQuery) GetCity() string {
	if m != nil && m.City != nil {
		return *m.City
	}
	return ""
}

func (m *LocationQuery) GetCityDistrict() string {
	if m != nil && m.CityDistrict != nil {
		return *m.CityDistrict
	}
	return ""
}

func (m *LocationQuery) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *LocationQuery) GetCountryState() string {
	if m != nil && m.CountryState != nil {
		return *m.CountryState
	}
	return ""
}

func (m *LocationQuery) GetCountryStateDistrict() string {
	if m != nil && m.CountryStateDistrict != nil {
		return *m.CountryStateDistrict
	}
	return ""
}

func (m *LocationQuery) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationQuery)(nil), "locationquery.LocationQuery")
}

func init() { proto.RegisterFile("locationquery/locationquery_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x4a, 0xf4, 0x30,
	0x14, 0xc5, 0xe9, 0x7c, 0xf3, 0xcd, 0x9f, 0xe8, 0x6c, 0x82, 0x48, 0x70, 0x21, 0x43, 0x57, 0xb3,
	0xb1, 0xc1, 0xe2, 0x46, 0xdc, 0xa9, 0xcb, 0x6e, 0x1c, 0xc1, 0x85, 0x1b, 0x89, 0x69, 0x68, 0x03,
	0x9d, 0xde, 0x9a, 0xdc, 0x08, 0xf5, 0xbd, 0x7c, 0x3f, 0x49, 0x1a, 0x87, 0x16, 0x5c, 0xf5, 0x77,
	0xce, 0x3d, 0xe7, 0xf6, 0x42, 0x48, 0xda, 0x80, 0x14, 0xa8, 0xa1, 0xfd, 0x70, 0xca, 0xf4, 0x7c,
	0xa2, 0xde, 0xb2, 0xce, 0x00, 0x02, 0xdd, 0x4c, 0xdc, 0x8b, 0xcb, 0x46, 0xb5, 0x15, 0xd6, 0x9f,
	0xa2, 0x71, 0x8a, 0x8f, 0x38, 0xc6, 0xd3, 0xef, 0x19, 0xd9, 0x14, 0xb1, 0xf1, 0xe4, 0x1b, 0xf4,
	0x86, 0xac, 0x8c, 0x28, 0xb5, 0xb3, 0x05, 0xb2, 0x64, 0x9b, 0xec, 0x4e, 0x72, 0x96, 0x8d, 0x8a,
	0x59, 0x11, 0xf8, 0xc5, 0xf3, 0xfe, 0x98, 0xa4, 0xe7, 0x64, 0x61, 0xd1, 0x28, 0x85, 0x6c, 0xb6,
	0x4d, 0x76, 0xeb, 0x7d, 0x54, 0x94, 0x91, 0xe5, 0x97, 0xee, 0x1e, 0xa0, 0x54, 0xec, 0x5f, 0x18,
	0xfc, 0x4a, 0x4a, 0xc9, 0x5c, 0x6a, 0xec, 0xd9, 0x3c, 0xd8, 0x81, 0x69, 0x4a, 0x4e, 0xfd, 0xf7,
	0x51, 0x5b, 0x34, 0x5a, 0x22, 0xfb, 0x1f, 0x66, 0x13, 0xcf, 0x6f, 0x94, 0xe0, 0x5a, 0x34, 0x3d,
	0x5b, 0x0c, 0x1b, 0xa3, 0x0c, 0xed, 0x01, 0x9f, 0x51, 0xa0, 0x62, 0xcb, 0xd8, 0x1e, 0x79, 0x34,
	0x27, 0x67, 0x63, 0x7d, 0xfc, 0xd3, 0x2a, 0x64, 0xff, 0x9c, 0xf9, 0x4b, 0x6b, 0x61, 0x6b, 0xb6,
	0x1e, 0x2e, 0xf5, 0x7c, 0x7f, 0xf7, 0x7a, 0x5b, 0x69, 0xac, 0xdd, 0x7b, 0x26, 0xe1, 0xc0, 0xf3,
	0x6b, 0x8b, 0x1a, 0x78, 0x05, 0x57, 0x5d, 0x23, 0xfa, 0xca, 0x80, 0x6b, 0x4b, 0x5e, 0x99, 0x4e,
	0x72, 0x2b, 0x6b, 0x75, 0x10, 0xd3, 0xa7, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd2, 0xf1, 0xa2,
	0x33, 0xc8, 0x01, 0x00, 0x00,
}