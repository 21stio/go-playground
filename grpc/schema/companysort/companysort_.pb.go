// Code generated by protoc-gen-go. DO NOT EDIT.
// source: companysort/companysort_.proto

/*
Package companysort is a generated protocol buffer package.

It is generated from these files:
	companysort/companysort_.proto

It has these top-level messages:
	CompanySort
*/
package companysort

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import infosort "github.com/21stio/go-playground/grpc/schema/infosort"
import locationsort "github.com/21stio/go-playground/grpc/schema/locationsort"
import sortkind "github.com/21stio/go-playground/grpc/schema/sortkind"
import currencyvaluesort "github.com/21stio/go-playground/grpc/schema/currencyvaluesort"
import servicesort "github.com/21stio/go-playground/grpc/schema/servicesort"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CompanySort struct {
	Info             *infosort.InfoSort                   `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	HeadQuarter      *locationsort.LocationSort           `protobuf:"bytes,2,opt,name=headQuarter" json:"headQuarter,omitempty"`
	EmployeeCount    *sortkind.SortKind                   `protobuf:"varint,3,opt,name=employeeCount,enum=sortkind.SortKind" json:"employeeCount,omitempty"`
	Revenue          *currencyvaluesort.CurrencyValueSort `protobuf:"bytes,4,opt,name=revenue" json:"revenue,omitempty"`
	Meta             *servicesort.TypeMetaSort            `protobuf:"bytes,5,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                              `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *CompanySort) Reset()                    { *m = CompanySort{} }
func (m *CompanySort) String() string            { return proto.CompactTextString(m) }
func (*CompanySort) ProtoMessage()               {}
func (*CompanySort) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CompanySort) GetInfo() *infosort.InfoSort {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *CompanySort) GetHeadQuarter() *locationsort.LocationSort {
	if m != nil {
		return m.HeadQuarter
	}
	return nil
}

func (m *CompanySort) GetEmployeeCount() sortkind.SortKind {
	if m != nil && m.EmployeeCount != nil {
		return *m.EmployeeCount
	}
	return sortkind.SortKind_ASC
}

func (m *CompanySort) GetRevenue() *currencyvaluesort.CurrencyValueSort {
	if m != nil {
		return m.Revenue
	}
	return nil
}

func (m *CompanySort) GetMeta() *servicesort.TypeMetaSort {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CompanySort) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CompanySort)(nil), "companysort.CompanySort")
}

func init() { proto.RegisterFile("companysort/companysort_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xe9, 0x5a, 0x15, 0x53, 0xf4, 0x90, 0x53, 0xdc, 0xc3, 0x52, 0x44, 0x64, 0x11, 0x36,
	0xc5, 0xbd, 0xa8, 0x20, 0x1e, 0xdc, 0x93, 0xa8, 0x07, 0x57, 0xf1, 0xe0, 0x45, 0x62, 0x3a, 0xdb,
	0x16, 0xdb, 0x4c, 0x49, 0xd3, 0x85, 0xfe, 0x2b, 0x7f, 0xa2, 0x34, 0x6d, 0x20, 0xea, 0xed, 0xcd,
	0xcc, 0xf7, 0xc8, 0x7b, 0xa5, 0x64, 0x26, 0xb1, 0xaa, 0x85, 0xea, 0x1a, 0xd4, 0x26, 0xf1, 0xf4,
	0x07, 0xaf, 0x35, 0x1a, 0xa4, 0x91, 0xb7, 0x9b, 0xb2, 0x42, 0x6d, 0xd0, 0x92, 0x4e, 0x8c, 0xd8,
	0x34, 0x2e, 0x51, 0x0a, 0x53, 0xa0, 0xb2, 0x57, 0x7f, 0x70, 0x04, 0xeb, 0x87, 0xaf, 0x42, 0xa5,
	0x89, 0x13, 0xee, 0x72, 0x2e, 0x5b, 0xad, 0x41, 0xc9, 0x6e, 0x2b, 0xca, 0x16, 0x86, 0x20, 0x7f,
	0x37, 0x8e, 0x9d, 0x35, 0xa0, 0xb7, 0x85, 0x1c, 0x28, 0x4f, 0x8f, 0xf7, 0x93, 0xef, 0x09, 0x89,
	0x56, 0x43, 0xe2, 0x17, 0xd4, 0x86, 0x9e, 0x91, 0xb0, 0x8f, 0xca, 0x82, 0x38, 0x98, 0x47, 0x4b,
	0xca, 0x5d, 0x6e, 0x7e, 0xaf, 0x36, 0xd8, 0x13, 0x6b, 0x7b, 0xa7, 0x37, 0x24, 0xca, 0x41, 0xa4,
	0xcf, 0xad, 0xd0, 0x06, 0x34, 0x9b, 0x58, 0x7c, 0xca, 0xfd, 0x22, 0xfc, 0x71, 0x1c, 0xac, 0xcd,
	0xc7, 0xe9, 0x15, 0x39, 0x84, 0xaa, 0x2e, 0xb1, 0x03, 0x58, 0x61, 0xab, 0x0c, 0xdb, 0x89, 0x83,
	0xf9, 0xd1, 0x92, 0x72, 0x57, 0x95, 0xf7, 0x9e, 0x87, 0x42, 0xa5, 0xeb, 0xdf, 0x20, 0xbd, 0x25,
	0xfb, 0x1a, 0xb6, 0xa0, 0x5a, 0x60, 0xa1, 0x7d, 0xf3, 0x94, 0xff, 0xeb, 0xce, 0x57, 0xe3, 0xe6,
	0xad, 0xdf, 0xd8, 0xd7, 0x9d, 0x89, 0x2e, 0x48, 0x58, 0x81, 0x11, 0x6c, 0xd7, 0x9a, 0x8f, 0xb9,
	0xf7, 0x49, 0xf8, 0x6b, 0x57, 0xc3, 0x13, 0x18, 0x31, 0xd4, 0xec, 0x31, 0x4a, 0x49, 0x98, 0x8b,
	0x26, 0x67, 0x7b, 0x71, 0x30, 0x3f, 0x58, 0x5b, 0x7d, 0x77, 0xfd, 0x7e, 0x99, 0x15, 0x26, 0x6f,
	0x3f, 0xb9, 0xc4, 0x2a, 0x59, 0x5e, 0x34, 0xa6, 0xc0, 0x24, 0xc3, 0x45, 0x5d, 0x8a, 0x2e, 0xd3,
	0xd8, 0xaa, 0x34, 0xc9, 0x74, 0x2d, 0x93, 0x46, 0xe6, 0x50, 0x09, 0xff, 0x1f, 0xf9, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0x3a, 0xb8, 0x70, 0x3d, 0x02, 0x00, 0x00,
}
