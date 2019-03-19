// Code generated by protoc-gen-go. DO NOT EDIT.
// source: translationsselect/translationsselect_.proto

/*
Package translationsselect is a generated protocol buffer package.

It is generated from these files:
	translationsselect/translationsselect_.proto

It has these top-level messages:
	TranslationsSelect
*/
package translationsselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import translationfilter "github.com/21stio/go-playground/grpc/schema/translationfilter"
import translationsort "github.com/21stio/go-playground/grpc/schema/translationsort"
import page "github.com/21stio/go-playground/grpc/schema/page"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TranslationsSelect struct {
	Filter           *translationfilter.TranslationFilter `protobuf:"bytes,1,opt,name=filter" json:"filter,omitempty"`
	Sort             *translationsort.TranslationSort     `protobuf:"bytes,2,opt,name=sort" json:"sort,omitempty"`
	Page             *page.Page                           `protobuf:"bytes,3,opt,name=page" json:"page,omitempty"`
	SelectAll        *bool                                `protobuf:"varint,4,opt,name=selectAll" json:"selectAll,omitempty"`
	Language         *bool                                `protobuf:"varint,5,opt,name=language" json:"language,omitempty"`
	Formatting       *bool                                `protobuf:"varint,6,opt,name=formatting" json:"formatting,omitempty"`
	Value            *bool                                `protobuf:"varint,7,opt,name=value" json:"value,omitempty"`
	SelectHash       *bool                                `protobuf:"varint,8,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                              `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *TranslationsSelect) Reset()                    { *m = TranslationsSelect{} }
func (m *TranslationsSelect) String() string            { return proto.CompactTextString(m) }
func (*TranslationsSelect) ProtoMessage()               {}
func (*TranslationsSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TranslationsSelect) GetFilter() *translationfilter.TranslationFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *TranslationsSelect) GetSort() *translationsort.TranslationSort {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *TranslationsSelect) GetPage() *page.Page {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *TranslationsSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *TranslationsSelect) GetLanguage() bool {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return false
}

func (m *TranslationsSelect) GetFormatting() bool {
	if m != nil && m.Formatting != nil {
		return *m.Formatting
	}
	return false
}

func (m *TranslationsSelect) GetValue() bool {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return false
}

func (m *TranslationsSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *TranslationsSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*TranslationsSelect)(nil), "translationsselect.TranslationsSelect")
}

func init() { proto.RegisterFile("translationsselect/translationsselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe9, 0xec, 0xe6, 0xf6, 0xbc, 0xc8, 0xc3, 0x43, 0x28, 0x32, 0x8a, 0x88, 0x0c, 0xd1,
	0x04, 0x87, 0x47, 0x11, 0xf5, 0x20, 0x1e, 0xa5, 0xf3, 0xe4, 0x45, 0x62, 0xcd, 0xd2, 0x42, 0xd6,
	0x94, 0x34, 0x15, 0xfc, 0x8e, 0x7e, 0x28, 0x49, 0x5a, 0x5d, 0xb4, 0x5e, 0x4a, 0xdf, 0xef, 0x9f,
	0x5f, 0xfe, 0x8f, 0x52, 0x38, 0xb3, 0x86, 0x57, 0x8d, 0xe2, 0xb6, 0xd4, 0x55, 0xd3, 0x08, 0x25,
	0x72, 0xcb, 0x86, 0xe8, 0x85, 0xd6, 0x46, 0x5b, 0x8d, 0x38, 0x8c, 0x92, 0xd3, 0x80, 0xad, 0x4b,
	0x65, 0x85, 0x61, 0x03, 0xd2, 0xfb, 0xc9, 0x49, 0xe8, 0x6b, 0xf3, 0xbb, 0x4a, 0x9b, 0xef, 0x9e,
	0x64, 0xbf, 0xe6, 0x52, 0x30, 0xf7, 0xe8, 0xc9, 0xd1, 0xe7, 0x08, 0xf0, 0x29, 0x38, 0xbc, 0xf2,
	0xe5, 0x78, 0x05, 0x93, 0xae, 0x81, 0x44, 0x69, 0xb4, 0xd8, 0x5b, 0x1e, 0xd3, 0x41, 0x37, 0x0d,
	0xb4, 0x7b, 0x4f, 0xb2, 0xde, 0xc1, 0x4b, 0x88, 0x5d, 0x2b, 0x19, 0x79, 0x37, 0xa5, 0x7f, 0xb6,
	0x09, 0xcd, 0x95, 0x36, 0x36, 0xf3, 0xa7, 0x71, 0x0e, 0xb1, 0xdb, 0x8c, 0xec, 0x78, 0x0b, 0xa8,
	0x1b, 0xe8, 0x23, 0x97, 0x22, 0xf3, 0x1c, 0x0f, 0x61, 0xd6, 0x7d, 0x9a, 0x5b, 0xa5, 0x48, 0x9c,
	0x46, 0x8b, 0x69, 0xb6, 0x05, 0x98, 0xc0, 0x54, 0xf1, 0x4a, 0xb6, 0xee, 0x86, 0xb1, 0x0f, 0x7f,
	0x66, 0x9c, 0x03, 0xac, 0xb5, 0xd9, 0x70, 0x6b, 0xcb, 0x4a, 0x92, 0x89, 0x4f, 0x03, 0x82, 0x07,
	0x30, 0x7e, 0xe7, 0xaa, 0x15, 0x64, 0xd7, 0x47, 0xdd, 0xe0, 0xac, 0xee, 0xfa, 0x07, 0xde, 0x14,
	0x64, 0xda, 0x59, 0x5b, 0x82, 0x08, 0x71, 0xe1, 0x92, 0x59, 0x1a, 0x2d, 0x66, 0x99, 0x7f, 0xbf,
	0xbb, 0x79, 0xbe, 0x96, 0xa5, 0x2d, 0xda, 0x57, 0x9a, 0xeb, 0x0d, 0x5b, 0x5e, 0x34, 0xb6, 0xd4,
	0x4c, 0xea, 0xf3, 0x5a, 0xf1, 0x0f, 0x69, 0x74, 0x5b, 0xbd, 0x31, 0x69, 0xea, 0x9c, 0x35, 0x79,
	0x21, 0x36, 0xfc, 0x9f, 0x3f, 0xe2, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x93, 0x8c, 0x7a, 0x6f, 0x39,
	0x02, 0x00, 0x00,
}