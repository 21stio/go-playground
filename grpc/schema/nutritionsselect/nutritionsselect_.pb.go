// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nutritionsselect/nutritionsselect_.proto

/*
Package nutritionsselect is a generated protocol buffer package.

It is generated from these files:
	nutritionsselect/nutritionsselect_.proto

It has these top-level messages:
	NutritionsSelect
*/
package nutritionsselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import weightperweightselect "github.com/21stio/go-playground/grpc/schema/weightperweightselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NutritionsSelect struct {
	SelectAll        *bool                                        `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Kcal             *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,2,opt,name=kcal" json:"kcal,omitempty"`
	Fat              *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,3,opt,name=fat" json:"fat,omitempty"`
	SaturatedFat     *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,4,opt,name=saturatedFat" json:"saturatedFat,omitempty"`
	Carbohydrates    *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,5,opt,name=carbohydrates" json:"carbohydrates,omitempty"`
	Sugar            *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,6,opt,name=sugar" json:"sugar,omitempty"`
	Roughages        *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,7,opt,name=roughages" json:"roughages,omitempty"`
	Proteins         *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,8,opt,name=proteins" json:"proteins,omitempty"`
	Salt             *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,9,opt,name=salt" json:"salt,omitempty"`
	Thiamin          *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,10,opt,name=thiamin" json:"thiamin,omitempty"`
	Magnesium        *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,11,opt,name=magnesium" json:"magnesium,omitempty"`
	Iron             *weightperweightselect.WeightPerWeightSelect `protobuf:"bytes,12,opt,name=iron" json:"iron,omitempty"`
	SelectHash       *bool                                        `protobuf:"varint,13,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                      `protobuf:"bytes,14,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                       `json:"-"`
}

func (m *NutritionsSelect) Reset()                    { *m = NutritionsSelect{} }
func (m *NutritionsSelect) String() string            { return proto.CompactTextString(m) }
func (*NutritionsSelect) ProtoMessage()               {}
func (*NutritionsSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NutritionsSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *NutritionsSelect) GetKcal() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Kcal
	}
	return nil
}

func (m *NutritionsSelect) GetFat() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Fat
	}
	return nil
}

func (m *NutritionsSelect) GetSaturatedFat() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.SaturatedFat
	}
	return nil
}

func (m *NutritionsSelect) GetCarbohydrates() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Carbohydrates
	}
	return nil
}

func (m *NutritionsSelect) GetSugar() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Sugar
	}
	return nil
}

func (m *NutritionsSelect) GetRoughages() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Roughages
	}
	return nil
}

func (m *NutritionsSelect) GetProteins() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Proteins
	}
	return nil
}

func (m *NutritionsSelect) GetSalt() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Salt
	}
	return nil
}

func (m *NutritionsSelect) GetThiamin() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Thiamin
	}
	return nil
}

func (m *NutritionsSelect) GetMagnesium() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Magnesium
	}
	return nil
}

func (m *NutritionsSelect) GetIron() *weightperweightselect.WeightPerWeightSelect {
	if m != nil {
		return m.Iron
	}
	return nil
}

func (m *NutritionsSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *NutritionsSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*NutritionsSelect)(nil), "nutritionsselect.NutritionsSelect")
}

func init() { proto.RegisterFile("nutritionsselect/nutritionsselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xbb, 0x6b, 0xeb, 0x30,
	0x14, 0x87, 0xf1, 0xcd, 0xd3, 0x27, 0xc9, 0x25, 0x68, 0x12, 0x97, 0x4b, 0x09, 0x9d, 0x3c, 0xb4,
	0x36, 0xcd, 0x5c, 0x42, 0x9b, 0x21, 0x84, 0x0e, 0x25, 0xb8, 0x43, 0xa1, 0x4b, 0x51, 0x1c, 0x55,
	0x12, 0xb5, 0x2d, 0xa3, 0x07, 0x25, 0xff, 0x7c, 0x29, 0xb2, 0xd3, 0xa4, 0x49, 0x33, 0x69, 0x3b,
	0xfe, 0x1d, 0x7d, 0x1f, 0xe8, 0xf8, 0xd8, 0x10, 0x95, 0xd6, 0x28, 0x61, 0x84, 0x2c, 0xb5, 0xa6,
	0x39, 0xcd, 0x4c, 0x72, 0x1a, 0xbc, 0xc6, 0x95, 0x92, 0x46, 0xa2, 0xf1, 0x69, 0xe3, 0xdf, 0xf4,
	0x83, 0x0a, 0xc6, 0x4d, 0x45, 0x55, 0x53, 0xec, 0x04, 0x67, 0xd3, 0x9d, 0xe5, 0xf2, 0xb3, 0x0b,
	0xe3, 0xc7, 0xbd, 0xe8, 0xa9, 0xee, 0xa1, 0xff, 0x10, 0x36, 0xa7, 0xee, 0xf3, 0x1c, 0x07, 0x93,
	0x20, 0xea, 0xa7, 0x87, 0x00, 0xdd, 0x41, 0xfb, 0x3d, 0x23, 0x39, 0xfe, 0x33, 0x09, 0xa2, 0xc1,
	0xf4, 0x2a, 0x3e, 0xeb, 0x8f, 0x9f, 0xeb, 0x87, 0x15, 0x55, 0x4d, 0xd1, 0x98, 0xd3, 0x9a, 0x44,
	0x33, 0x68, 0xbd, 0x11, 0x83, 0x5b, 0x1e, 0x02, 0x07, 0xa2, 0x15, 0x0c, 0x35, 0x31, 0x56, 0x11,
	0x43, 0x37, 0x0b, 0x62, 0x70, 0xdb, 0x43, 0x74, 0x64, 0x40, 0x29, 0x8c, 0x32, 0xa2, 0xd6, 0x92,
	0x6f, 0x37, 0x2e, 0xd3, 0xb8, 0xe3, 0xa1, 0x3c, 0x56, 0xa0, 0x39, 0x74, 0xb4, 0x65, 0x44, 0xe1,
	0xae, 0x87, 0xab, 0x41, 0xd1, 0x03, 0x84, 0x4a, 0x5a, 0xc6, 0x09, 0xa3, 0x1a, 0xf7, 0x3c, 0x3c,
	0x07, 0x1c, 0x2d, 0xa1, 0xef, 0xde, 0x39, 0x15, 0xa5, 0xc6, 0x7d, 0x0f, 0xd5, 0x9e, 0x76, 0x1b,
	0xa0, 0x49, 0x6e, 0x70, 0xe8, 0xb3, 0x01, 0x8e, 0x44, 0x0b, 0xe8, 0x19, 0x2e, 0x48, 0x21, 0x4a,
	0x0c, 0x1e, 0x92, 0x6f, 0xd8, 0xcd, 0xa7, 0x20, 0xac, 0xa4, 0x5a, 0xd8, 0x02, 0x0f, 0x7c, 0xe6,
	0xb3, 0xc7, 0xdd, 0xad, 0x84, 0x92, 0x25, 0x1e, 0xfa, 0xdc, 0xca, 0x91, 0xe8, 0x02, 0xa0, 0x39,
	0xb5, 0x24, 0x9a, 0xe3, 0x51, 0xfd, 0xe1, 0xfc, 0x48, 0x10, 0x82, 0x36, 0x77, 0x9d, 0xbf, 0x93,
	0x20, 0x0a, 0xd3, 0xba, 0x9e, 0xcf, 0x5e, 0x6e, 0x99, 0x30, 0xdc, 0xae, 0xe3, 0x4c, 0x16, 0xc9,
	0xf4, 0x46, 0x1b, 0x21, 0x13, 0x26, 0xaf, 0xab, 0x9c, 0x6c, 0x99, 0x92, 0xb6, 0xdc, 0x24, 0x4c,
	0x55, 0x59, 0xa2, 0x33, 0x4e, 0x0b, 0xf2, 0xeb, 0x6f, 0xf0, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd0,
	0x0d, 0x4b, 0x77, 0x31, 0x04, 0x00, 0x00,
}
