// Code generated by protoc-gen-go. DO NOT EDIT.
// source: locationfilter/locationfilter_.proto

/*
Package locationfilter is a generated protocol buffer package.

It is generated from these files:
	locationfilter/locationfilter_.proto

It has these top-level messages:
	LocationFilter
*/
package locationfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import infofilter "github.com/21stio/go-playground/grpc/schema/infofilter"
import lengthvaluefilter "github.com/21stio/go-playground/grpc/schema/lengthvaluefilter"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LocationFilter struct {
	Info                 *infofilter.InfoFilter               `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Radius               *lengthvaluefilter.LengthValueFilter `protobuf:"bytes,2,opt,name=radius" json:"radius,omitempty"`
	Street               *stringfilter.StringFilter           `protobuf:"bytes,3,opt,name=street" json:"street,omitempty"`
	ZipCode              *stringfilter.StringFilter           `protobuf:"bytes,4,opt,name=zipCode" json:"zipCode,omitempty"`
	City                 *stringfilter.StringFilter           `protobuf:"bytes,5,opt,name=city" json:"city,omitempty"`
	CityDistrict         *stringfilter.StringFilter           `protobuf:"bytes,6,opt,name=cityDistrict" json:"cityDistrict,omitempty"`
	Country              *stringfilter.StringFilter           `protobuf:"bytes,7,opt,name=country" json:"country,omitempty"`
	CountryState         *stringfilter.StringFilter           `protobuf:"bytes,8,opt,name=countryState" json:"countryState,omitempty"`
	CountryStateDistrict *stringfilter.StringFilter           `protobuf:"bytes,9,opt,name=countryStateDistrict" json:"countryStateDistrict,omitempty"`
	IdsSome              *idfilter.IdFilter                   `protobuf:"bytes,10,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery             *idfilter.IdFilter                   `protobuf:"bytes,11,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone              *idfilter.IdFilter                   `protobuf:"bytes,12,opt,name=idsNone" json:"idsNone,omitempty"`
	Meta                 *servicefilter.TypeMetaFilter        `protobuf:"bytes,13,opt,name=meta" json:"meta,omitempty"`
	And                  []*LocationFilter                    `protobuf:"bytes,14,rep,name=and" json:"and,omitempty"`
	Or                   []*LocationFilter                    `protobuf:"bytes,15,rep,name=or" json:"or,omitempty"`
	Not                  []*LocationFilter                    `protobuf:"bytes,16,rep,name=not" json:"not,omitempty"`
	Hash                 *string                              `protobuf:"bytes,17,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized     []byte                               `json:"-"`
}

func (m *LocationFilter) Reset()                    { *m = LocationFilter{} }
func (m *LocationFilter) String() string            { return proto.CompactTextString(m) }
func (*LocationFilter) ProtoMessage()               {}
func (*LocationFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LocationFilter) GetInfo() *infofilter.InfoFilter {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *LocationFilter) GetRadius() *lengthvaluefilter.LengthValueFilter {
	if m != nil {
		return m.Radius
	}
	return nil
}

func (m *LocationFilter) GetStreet() *stringfilter.StringFilter {
	if m != nil {
		return m.Street
	}
	return nil
}

func (m *LocationFilter) GetZipCode() *stringfilter.StringFilter {
	if m != nil {
		return m.ZipCode
	}
	return nil
}

func (m *LocationFilter) GetCity() *stringfilter.StringFilter {
	if m != nil {
		return m.City
	}
	return nil
}

func (m *LocationFilter) GetCityDistrict() *stringfilter.StringFilter {
	if m != nil {
		return m.CityDistrict
	}
	return nil
}

func (m *LocationFilter) GetCountry() *stringfilter.StringFilter {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *LocationFilter) GetCountryState() *stringfilter.StringFilter {
	if m != nil {
		return m.CountryState
	}
	return nil
}

func (m *LocationFilter) GetCountryStateDistrict() *stringfilter.StringFilter {
	if m != nil {
		return m.CountryStateDistrict
	}
	return nil
}

func (m *LocationFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *LocationFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *LocationFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *LocationFilter) GetMeta() *servicefilter.TypeMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *LocationFilter) GetAnd() []*LocationFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *LocationFilter) GetOr() []*LocationFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *LocationFilter) GetNot() []*LocationFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *LocationFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*LocationFilter)(nil), "locationfilter.LocationFilter")
}

func init() { proto.RegisterFile("locationfilter/locationfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xb5, 0x36, 0xb4, 0x9b, 0x3b, 0x0a, 0x58, 0x08, 0x59, 0x15, 0xa0, 0x6a, 0xda, 0xa1,
	0x9a, 0xc0, 0x61, 0x15, 0x27, 0x34, 0x71, 0xe0, 0x97, 0x34, 0x69, 0xec, 0xd0, 0x22, 0x0e, 0x5c,
	0x90, 0x49, 0xdc, 0xc4, 0x52, 0xea, 0x17, 0x39, 0x2f, 0x95, 0xc2, 0x8d, 0xff, 0x1c, 0xd9, 0xb5,
	0xb7, 0x66, 0xfc, 0x48, 0x4f, 0xf1, 0x7b, 0xfe, 0x7c, 0xdf, 0xf7, 0x6b, 0xc5, 0x09, 0x39, 0x2d,
	0x20, 0x11, 0xa8, 0x40, 0xaf, 0x54, 0x81, 0xd2, 0xc4, 0xed, 0xf2, 0x3b, 0x2f, 0x0d, 0x20, 0xd0,
	0x71, 0xbb, 0x3d, 0x79, 0xaa, 0xf4, 0x0a, 0xbc, 0xe2, 0x76, 0xe9, 0xe9, 0xc9, 0x59, 0x21, 0x75,
	0x86, 0xf9, 0x46, 0x14, 0xb5, 0x0c, 0x63, 0xef, 0x76, 0x02, 0x3b, 0xad, 0xd0, 0x28, 0x9d, 0x79,
	0x6c, 0xb7, 0x08, 0x04, 0x53, 0x69, 0x70, 0x4a, 0xdb, 0x3b, 0x27, 0x95, 0x34, 0x1b, 0x95, 0x04,
	0x8f, 0x56, 0xe5, 0x99, 0x93, 0x5f, 0x43, 0x32, 0xbe, 0xf2, 0xe1, 0x3f, 0xb9, 0x1d, 0x7a, 0x46,
	0x22, 0x9b, 0x99, 0x1d, 0x4c, 0x0f, 0x66, 0xa3, 0xf9, 0x13, 0x7e, 0x7b, 0x00, 0x7e, 0xa9, 0x57,
	0xb0, 0xa5, 0x16, 0x8e, 0xa1, 0x17, 0x64, 0x60, 0x44, 0xaa, 0xea, 0x8a, 0xf5, 0x1c, 0x7d, 0xca,
	0xff, 0x38, 0x09, 0xbf, 0x72, 0x9d, 0xaf, 0xb6, 0xe3, 0xb5, 0x5e, 0x43, 0xe7, 0x64, 0x50, 0xa1,
	0x91, 0x12, 0x59, 0xdf, 0xa9, 0x27, 0x7c, 0xf7, 0x80, 0x7c, 0xe9, 0x8a, 0xa0, 0xd9, 0x92, 0xf4,
	0x35, 0x19, 0xfe, 0x54, 0xe5, 0x7b, 0x48, 0x25, 0x8b, 0x3a, 0x45, 0x01, 0xa5, 0x9c, 0x44, 0x89,
	0xc2, 0x86, 0xdd, 0xeb, 0x94, 0x38, 0x8e, 0xbe, 0x25, 0xc7, 0xf6, 0xf9, 0x41, 0x59, 0x30, 0x41,
	0x36, 0xe8, 0xd4, 0xb5, 0x78, 0x9b, 0x32, 0x81, 0x5a, 0xa3, 0x69, 0xd8, 0xb0, 0x3b, 0xa5, 0x47,
	0x9d, 0xeb, 0x76, 0xb9, 0x44, 0x81, 0x92, 0x1d, 0xee, 0xe1, 0xba, 0xc3, 0xd3, 0x6b, 0xf2, 0x78,
	0xb7, 0xbe, 0x49, 0x7f, 0xd4, 0x39, 0xe7, 0xaf, 0x3a, 0xfa, 0x82, 0x0c, 0x55, 0x5a, 0x2d, 0x61,
	0x2d, 0x19, 0x71, 0x23, 0x28, 0x0f, 0x77, 0x8c, 0x5f, 0xa6, 0x21, 0xbd, 0x47, 0x28, 0x27, 0x87,
	0x2a, 0xad, 0x3e, 0x6e, 0xa4, 0x69, 0xd8, 0xe8, 0x9f, 0xf8, 0x0d, 0xe3, 0xa7, 0x5f, 0x83, 0x96,
	0xec, 0xf8, 0xbf, 0xd3, 0x2d, 0x42, 0xcf, 0x49, 0xb4, 0x96, 0x28, 0xd8, 0x7d, 0x87, 0x3e, 0xe3,
	0xad, 0xdb, 0xcc, 0xbf, 0x34, 0xa5, 0xfc, 0x2c, 0x51, 0x84, 0x97, 0x68, 0x51, 0xfa, 0x8a, 0xf4,
	0x85, 0x4e, 0xd9, 0x78, 0xda, 0x9f, 0x8d, 0xe6, 0xcf, 0x79, 0xfb, 0x1b, 0xe5, 0xed, 0x5b, 0xbf,
	0xb0, 0x28, 0xe5, 0xa4, 0x07, 0x86, 0x3d, 0xd8, 0x4b, 0xd0, 0x03, 0x63, 0x1d, 0x34, 0x20, 0x7b,
	0xb8, 0x9f, 0x83, 0x06, 0xa4, 0x94, 0x44, 0xb9, 0xa8, 0x72, 0xf6, 0x68, 0x7a, 0x30, 0x3b, 0x5a,
	0xb8, 0xf5, 0xbb, 0x8b, 0x6f, 0x6f, 0x32, 0x85, 0x79, 0xfd, 0x83, 0x27, 0xb0, 0x8e, 0xe7, 0xe7,
	0x15, 0x2a, 0x88, 0x33, 0x78, 0x59, 0x16, 0xa2, 0xc9, 0x0c, 0xd4, 0x3a, 0x8d, 0x33, 0x53, 0x26,
	0x71, 0x95, 0xe4, 0x72, 0x2d, 0xee, 0xfc, 0x82, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xac,
	0xf2, 0xba, 0xa2, 0x04, 0x00, 0x00,
}
