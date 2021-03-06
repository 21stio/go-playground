// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nutritionsfilter/nutritionsfilter_.proto

/*
Package nutritionsfilter is a generated protocol buffer package.

It is generated from these files:
	nutritionsfilter/nutritionsfilter_.proto

It has these top-level messages:
	NutritionsFilter
*/
package nutritionsfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import weightperweightfilter "github.com/21stio/go-playground/grpc/schema/weightperweightfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NutritionsFilter struct {
	Kcal             *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,1,opt,name=kcal" json:"kcal,omitempty"`
	Fat              *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,2,opt,name=fat" json:"fat,omitempty"`
	SaturatedFat     *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,3,opt,name=saturatedFat" json:"saturatedFat,omitempty"`
	Carbohydrates    *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,4,opt,name=carbohydrates" json:"carbohydrates,omitempty"`
	Sugar            *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,5,opt,name=sugar" json:"sugar,omitempty"`
	Roughages        *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,6,opt,name=roughages" json:"roughages,omitempty"`
	Proteins         *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,7,opt,name=proteins" json:"proteins,omitempty"`
	Salt             *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,8,opt,name=salt" json:"salt,omitempty"`
	Thiamin          *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,9,opt,name=thiamin" json:"thiamin,omitempty"`
	Magnesium        *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,10,opt,name=magnesium" json:"magnesium,omitempty"`
	Iron             *weightperweightfilter.WeightPerWeightFilter `protobuf:"bytes,11,opt,name=iron" json:"iron,omitempty"`
	And              []*NutritionsFilter                          `protobuf:"bytes,12,rep,name=and" json:"and,omitempty"`
	Or               []*NutritionsFilter                          `protobuf:"bytes,13,rep,name=or" json:"or,omitempty"`
	Not              []*NutritionsFilter                          `protobuf:"bytes,14,rep,name=not" json:"not,omitempty"`
	Hash             *string                                      `protobuf:"bytes,15,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                       `json:"-"`
}

func (m *NutritionsFilter) Reset()                    { *m = NutritionsFilter{} }
func (m *NutritionsFilter) String() string            { return proto.CompactTextString(m) }
func (*NutritionsFilter) ProtoMessage()               {}
func (*NutritionsFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NutritionsFilter) GetKcal() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Kcal
	}
	return nil
}

func (m *NutritionsFilter) GetFat() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Fat
	}
	return nil
}

func (m *NutritionsFilter) GetSaturatedFat() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.SaturatedFat
	}
	return nil
}

func (m *NutritionsFilter) GetCarbohydrates() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Carbohydrates
	}
	return nil
}

func (m *NutritionsFilter) GetSugar() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Sugar
	}
	return nil
}

func (m *NutritionsFilter) GetRoughages() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Roughages
	}
	return nil
}

func (m *NutritionsFilter) GetProteins() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Proteins
	}
	return nil
}

func (m *NutritionsFilter) GetSalt() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *NutritionsFilter) GetThiamin() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Thiamin
	}
	return nil
}

func (m *NutritionsFilter) GetMagnesium() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Magnesium
	}
	return nil
}

func (m *NutritionsFilter) GetIron() *weightperweightfilter.WeightPerWeightFilter {
	if m != nil {
		return m.Iron
	}
	return nil
}

func (m *NutritionsFilter) GetAnd() []*NutritionsFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *NutritionsFilter) GetOr() []*NutritionsFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *NutritionsFilter) GetNot() []*NutritionsFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *NutritionsFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*NutritionsFilter)(nil), "nutritionsfilter.NutritionsFilter")
}

func init() { proto.RegisterFile("nutritionsfilter/nutritionsfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x8b, 0xe2, 0x30,
	0x14, 0x80, 0xd1, 0xea, 0xaa, 0x51, 0x77, 0x25, 0xa7, 0xb0, 0x27, 0xf1, 0xd4, 0xc3, 0x6e, 0xcb,
	0x96, 0x3d, 0x2e, 0xb2, 0x78, 0x90, 0x61, 0x0e, 0x83, 0xf4, 0x32, 0x30, 0x97, 0x21, 0xb6, 0x31,
	0x09, 0xd3, 0x26, 0x25, 0x79, 0x65, 0xf0, 0xdf, 0xce, 0x4f, 0x19, 0xd2, 0x3a, 0x0e, 0x76, 0x3c,
	0x48, 0x6e, 0xaf, 0x2f, 0xfd, 0xbe, 0xf0, 0x5e, 0xdf, 0x2b, 0x0a, 0x55, 0x0d, 0x46, 0x82, 0xd4,
	0xca, 0x1e, 0x64, 0x01, 0xcc, 0xc4, 0xdd, 0xc4, 0x73, 0x54, 0x19, 0x0d, 0x1a, 0x2f, 0xba, 0x07,
	0x3f, 0x93, 0x57, 0x26, 0xb9, 0x80, 0x8a, 0x99, 0x36, 0x38, 0x09, 0xae, 0x66, 0x4f, 0x96, 0xd5,
	0xdb, 0x08, 0x2d, 0x1e, 0xce, 0xa2, 0x6d, 0x73, 0x86, 0xff, 0xa3, 0xc1, 0x4b, 0x46, 0x0b, 0xd2,
	0x5b, 0xf6, 0xc2, 0x69, 0xf2, 0x2b, 0xba, 0x6a, 0x88, 0x1e, 0x9b, 0x87, 0x1d, 0x33, 0x6d, 0xd0,
	0xb2, 0x69, 0x43, 0xe2, 0x35, 0x0a, 0x0e, 0x14, 0x48, 0xdf, 0x43, 0xe0, 0x40, 0xbc, 0x43, 0x33,
	0x4b, 0xa1, 0x36, 0x14, 0x58, 0xbe, 0xa5, 0x40, 0x02, 0x0f, 0xd1, 0x85, 0x01, 0xa7, 0x68, 0x9e,
	0x51, 0xb3, 0xd7, 0xe2, 0x98, 0xbb, 0x9c, 0x25, 0x03, 0x0f, 0xe5, 0xa5, 0x02, 0x6f, 0xd0, 0xd0,
	0xd6, 0x9c, 0x1a, 0x32, 0xf4, 0x70, 0xb5, 0x28, 0xbe, 0x47, 0x13, 0xa3, 0x6b, 0x2e, 0x28, 0x67,
	0x96, 0x7c, 0xf3, 0xf0, 0x7c, 0xe2, 0xf8, 0x0e, 0x8d, 0xdd, 0x57, 0x65, 0x52, 0x59, 0x32, 0xf2,
	0x50, 0x9d, 0x69, 0x37, 0x01, 0x96, 0x16, 0x40, 0xc6, 0x3e, 0x13, 0xe0, 0x48, 0xbc, 0x45, 0x23,
	0x10, 0x92, 0x96, 0x52, 0x91, 0x89, 0x87, 0xe4, 0x03, 0x76, 0xfd, 0x29, 0x29, 0x57, 0xcc, 0xca,
	0xba, 0x24, 0xc8, 0xa7, 0x3f, 0x67, 0xdc, 0x55, 0x25, 0x8d, 0x56, 0x64, 0xea, 0x53, 0x95, 0x23,
	0xf1, 0x5f, 0x14, 0x50, 0x95, 0x93, 0xd9, 0x32, 0x08, 0xa7, 0xc9, 0x2a, 0xea, 0xae, 0x60, 0xd4,
	0x5d, 0xa5, 0xd4, 0xbd, 0x8e, 0x13, 0xd4, 0xd7, 0x86, 0xcc, 0x6f, 0x86, 0xfa, 0xda, 0xb8, 0x9b,
	0x94, 0x06, 0xf2, 0xfd, 0xf6, 0x9b, 0x94, 0x06, 0x8c, 0xd1, 0x40, 0x50, 0x2b, 0xc8, 0x8f, 0x65,
	0x2f, 0x9c, 0xa4, 0x4d, 0xbc, 0x59, 0x3f, 0xfd, 0xe3, 0x12, 0x44, 0xbd, 0x8f, 0x32, 0x5d, 0xc6,
	0xc9, 0x1f, 0x0b, 0x52, 0xc7, 0x5c, 0xff, 0xae, 0x0a, 0x7a, 0xe4, 0x46, 0xd7, 0x2a, 0x8f, 0xb9,
	0xa9, 0xb2, 0xd8, 0x66, 0x82, 0x95, 0xf4, 0xcb, 0xff, 0xe6, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x9f,
	0x25, 0x3c, 0x65, 0x93, 0x04, 0x00, 0x00,
}
