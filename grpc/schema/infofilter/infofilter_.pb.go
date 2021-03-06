// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infofilter/infofilter_.proto

/*
Package infofilter is a generated protocol buffer package.

It is generated from these files:
	infofilter/infofilter_.proto

It has these top-level messages:
	InfoFilter
*/
package infofilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import textfilter "github.com/21stio/go-playground/grpc/schema/textfilter"
import imagefilter "github.com/21stio/go-playground/grpc/schema/imagefilter"
import linkfilter "github.com/21stio/go-playground/grpc/schema/linkfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InfoFilter struct {
	Name             *textfilter.TextFilter   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description      *textfilter.TextFilter   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	ImagesSome       *imagefilter.ImageFilter `protobuf:"bytes,3,opt,name=imagesSome" json:"imagesSome,omitempty"`
	ImagesEvery      *imagefilter.ImageFilter `protobuf:"bytes,4,opt,name=imagesEvery" json:"imagesEvery,omitempty"`
	ImagesNone       *imagefilter.ImageFilter `protobuf:"bytes,5,opt,name=imagesNone" json:"imagesNone,omitempty"`
	LinksSome        *linkfilter.LinkFilter   `protobuf:"bytes,6,opt,name=linksSome" json:"linksSome,omitempty"`
	LinksEvery       *linkfilter.LinkFilter   `protobuf:"bytes,7,opt,name=linksEvery" json:"linksEvery,omitempty"`
	LinksNone        *linkfilter.LinkFilter   `protobuf:"bytes,8,opt,name=linksNone" json:"linksNone,omitempty"`
	And              []*InfoFilter            `protobuf:"bytes,9,rep,name=and" json:"and,omitempty"`
	Or               []*InfoFilter            `protobuf:"bytes,10,rep,name=or" json:"or,omitempty"`
	Not              []*InfoFilter            `protobuf:"bytes,11,rep,name=not" json:"not,omitempty"`
	Hash             *string                  `protobuf:"bytes,12,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *InfoFilter) Reset()                    { *m = InfoFilter{} }
func (m *InfoFilter) String() string            { return proto.CompactTextString(m) }
func (*InfoFilter) ProtoMessage()               {}
func (*InfoFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InfoFilter) GetName() *textfilter.TextFilter {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *InfoFilter) GetDescription() *textfilter.TextFilter {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *InfoFilter) GetImagesSome() *imagefilter.ImageFilter {
	if m != nil {
		return m.ImagesSome
	}
	return nil
}

func (m *InfoFilter) GetImagesEvery() *imagefilter.ImageFilter {
	if m != nil {
		return m.ImagesEvery
	}
	return nil
}

func (m *InfoFilter) GetImagesNone() *imagefilter.ImageFilter {
	if m != nil {
		return m.ImagesNone
	}
	return nil
}

func (m *InfoFilter) GetLinksSome() *linkfilter.LinkFilter {
	if m != nil {
		return m.LinksSome
	}
	return nil
}

func (m *InfoFilter) GetLinksEvery() *linkfilter.LinkFilter {
	if m != nil {
		return m.LinksEvery
	}
	return nil
}

func (m *InfoFilter) GetLinksNone() *linkfilter.LinkFilter {
	if m != nil {
		return m.LinksNone
	}
	return nil
}

func (m *InfoFilter) GetAnd() []*InfoFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *InfoFilter) GetOr() []*InfoFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *InfoFilter) GetNot() []*InfoFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *InfoFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*InfoFilter)(nil), "infofilter.InfoFilter")
}

func init() { proto.RegisterFile("infofilter/infofilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xd9, 0xd6, 0xff, 0x5f, 0xf7, 0xd6, 0x53, 0x0e, 0x12, 0xc6, 0x90, 0xe1, 0x41, 0x86,
	0x60, 0x8a, 0x43, 0xc6, 0xf0, 0x28, 0x28, 0x0c, 0xc4, 0xc3, 0xf4, 0xe4, 0x45, 0x62, 0x97, 0xb5,
	0x61, 0x6b, 0xde, 0x92, 0x66, 0xb2, 0x7d, 0x51, 0x3f, 0x8f, 0x34, 0xcd, 0x4c, 0x4e, 0x75, 0xb7,
	0x27, 0xc9, 0xef, 0x79, 0x9f, 0xa7, 0x69, 0x0b, 0x43, 0xa9, 0x56, 0xb8, 0x92, 0x1b, 0x23, 0x74,
	0xe2, 0xe5, 0x07, 0x2b, 0x35, 0x1a, 0x24, 0xe0, 0xb7, 0x06, 0x43, 0x23, 0x76, 0xc6, 0x91, 0x5e,
	0x3a, 0x72, 0x70, 0x21, 0x0b, 0x9e, 0x89, 0xc3, 0x20, 0xaf, 0x0f, 0xe7, 0xc3, 0x8d, 0x54, 0x6b,
	0x77, 0xec, 0xa5, 0x3b, 0xbd, 0xfc, 0x8e, 0x00, 0xe6, 0x6a, 0x85, 0x4f, 0x76, 0x97, 0x5c, 0x43,
	0xa4, 0x78, 0x21, 0x68, 0x67, 0xd4, 0x19, 0xc7, 0x93, 0x73, 0xe6, 0xe3, 0xd8, 0x9b, 0xd8, 0x99,
	0x86, 0x5a, 0x58, 0x86, 0xcc, 0x20, 0x5e, 0x8a, 0x2a, 0xd5, 0xb2, 0x34, 0x12, 0x15, 0xed, 0xb6,
	0x5a, 0x42, 0x94, 0xcc, 0x00, 0x6c, 0xd1, 0xea, 0x15, 0x0b, 0x41, 0x7b, 0xd6, 0x48, 0x59, 0xd0,
	0x9d, 0xcd, 0x6b, 0xed, 0xac, 0x01, 0x4b, 0xee, 0x21, 0x6e, 0x56, 0x8f, 0x5f, 0x42, 0xef, 0x69,
	0xf4, 0x87, 0x35, 0x84, 0x7d, 0xea, 0x0b, 0x2a, 0x41, 0xff, 0x1d, 0x97, 0x5a, 0xb3, 0xe4, 0x0e,
	0xfa, 0xf5, 0xcd, 0x35, 0x75, 0xff, 0xbb, 0xe7, 0xf4, 0x77, 0xc9, 0x9e, 0xa5, 0x5a, 0x3b, 0x9b,
	0x07, 0xc9, 0x14, 0xc0, 0x2e, 0x9a, 0xaa, 0x27, 0xad, 0xb6, 0x80, 0xfc, 0x4d, 0xb3, 0x35, 0x4f,
	0x8f, 0x48, 0xb3, 0x1d, 0xc7, 0xd0, 0xe3, 0x6a, 0x49, 0xfb, 0xa3, 0x9e, 0xe5, 0xfd, 0xe7, 0xc3,
	0xfc, 0xeb, 0x5d, 0xd4, 0x08, 0xb9, 0x82, 0x2e, 0x6a, 0x0a, 0xad, 0x60, 0x17, 0x75, 0x3d, 0x51,
	0xa1, 0xa1, 0x71, 0xfb, 0x44, 0x85, 0x86, 0x10, 0x88, 0x72, 0x5e, 0xe5, 0xf4, 0x6c, 0xd4, 0x19,
	0xf7, 0x17, 0x56, 0x3f, 0xcc, 0xde, 0xa7, 0x99, 0x34, 0xf9, 0xf6, 0x93, 0xa5, 0x58, 0x24, 0x93,
	0xdb, 0xca, 0x48, 0x4c, 0x32, 0xbc, 0x29, 0x37, 0x7c, 0x9f, 0x69, 0xdc, 0xaa, 0x65, 0x92, 0xe9,
	0x32, 0x4d, 0xaa, 0x34, 0x17, 0x05, 0x0f, 0xfe, 0x80, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97,
	0x48, 0xe8, 0xec, 0x19, 0x03, 0x00, 0x00,
}
