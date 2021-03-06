// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serviceselect/serviceselect_.proto

/*
Package serviceselect is a generated protocol buffer package.

It is generated from these files:
	serviceselect/serviceselect_.proto

It has these top-level messages:
	TypeMetaSelect
	ServiceSelect
*/
package serviceselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import ratingselect "github.com/21stio/go-playground/grpc/schema/ratingselect"
import durationvalueselect "github.com/21stio/go-playground/grpc/schema/durationvalueselect"
import timestampselect "github.com/21stio/go-playground/grpc/schema/timestampselect"
import endpointsselect "github.com/21stio/go-playground/grpc/schema/endpointsselect"
import serviceportsselect "github.com/21stio/go-playground/grpc/schema/serviceportsselect"
import idsselect "github.com/21stio/go-playground/grpc/schema/idsselect"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TypeMetaSelect struct {
	SelectAll        *bool                                    `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Service          *ServiceSelect                           `protobuf:"bytes,2,opt,name=service" json:"service,omitempty"`
	Relevance        *bool                                    `protobuf:"varint,3,opt,name=relevance" json:"relevance,omitempty"`
	Rating           *ratingselect.RatingSelect               `protobuf:"bytes,4,opt,name=rating" json:"rating,omitempty"`
	Age              *durationvalueselect.DurationValueSelect `protobuf:"bytes,5,opt,name=age" json:"age,omitempty"`
	CreatedAt        *timestampselect.TimestampSelect         `protobuf:"bytes,6,opt,name=createdAt" json:"createdAt,omitempty"`
	UpdatedAt        *timestampselect.TimestampSelect         `protobuf:"bytes,7,opt,name=updatedAt" json:"updatedAt,omitempty"`
	SelectHash       *bool                                    `protobuf:"varint,8,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string                                  `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *TypeMetaSelect) Reset()                    { *m = TypeMetaSelect{} }
func (m *TypeMetaSelect) String() string            { return proto.CompactTextString(m) }
func (*TypeMetaSelect) ProtoMessage()               {}
func (*TypeMetaSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TypeMetaSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *TypeMetaSelect) GetService() *ServiceSelect {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *TypeMetaSelect) GetRelevance() bool {
	if m != nil && m.Relevance != nil {
		return *m.Relevance
	}
	return false
}

func (m *TypeMetaSelect) GetRating() *ratingselect.RatingSelect {
	if m != nil {
		return m.Rating
	}
	return nil
}

func (m *TypeMetaSelect) GetAge() *durationvalueselect.DurationValueSelect {
	if m != nil {
		return m.Age
	}
	return nil
}

func (m *TypeMetaSelect) GetCreatedAt() *timestampselect.TimestampSelect {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TypeMetaSelect) GetUpdatedAt() *timestampselect.TimestampSelect {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *TypeMetaSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *TypeMetaSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

type ServiceSelect struct {
	SelectAll        *bool                                  `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	Name             *bool                                  `protobuf:"varint,2,opt,name=name" json:"name,omitempty"`
	Host             *bool                                  `protobuf:"varint,3,opt,name=host" json:"host,omitempty"`
	Endpoints        *endpointsselect.EndpointsSelect       `protobuf:"bytes,4,opt,name=endpoints" json:"endpoints,omitempty"`
	Ports            *serviceportsselect.ServicePortsSelect `protobuf:"bytes,5,opt,name=ports" json:"ports,omitempty"`
	SelectHash       *bool                                  `protobuf:"varint,6,opt,name=selectHash" json:"selectHash,omitempty"`
	Ids              *idsselect.IdsSelect                   `protobuf:"bytes,7,opt,name=ids" json:"ids,omitempty"`
	Meta             *TypeMetaSelect                        `protobuf:"bytes,8,opt,name=meta" json:"meta,omitempty"`
	Hash             *string                                `protobuf:"bytes,9,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                 `json:"-"`
}

func (m *ServiceSelect) Reset()                    { *m = ServiceSelect{} }
func (m *ServiceSelect) String() string            { return proto.CompactTextString(m) }
func (*ServiceSelect) ProtoMessage()               {}
func (*ServiceSelect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServiceSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *ServiceSelect) GetName() bool {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return false
}

func (m *ServiceSelect) GetHost() bool {
	if m != nil && m.Host != nil {
		return *m.Host
	}
	return false
}

func (m *ServiceSelect) GetEndpoints() *endpointsselect.EndpointsSelect {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *ServiceSelect) GetPorts() *serviceportsselect.ServicePortsSelect {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *ServiceSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *ServiceSelect) GetIds() *idsselect.IdsSelect {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *ServiceSelect) GetMeta() *TypeMetaSelect {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ServiceSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*TypeMetaSelect)(nil), "serviceselect.TypeMetaSelect")
	proto.RegisterType((*ServiceSelect)(nil), "serviceselect.ServiceSelect")
}

func init() { proto.RegisterFile("serviceselect/serviceselect_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x8b, 0xdb, 0x30,
	0x14, 0x24, 0x9b, 0x6c, 0x36, 0xd1, 0xb2, 0x3d, 0x88, 0x1e, 0x4c, 0xd8, 0x16, 0x93, 0x43, 0xc8,
	0xa1, 0x95, 0xd9, 0x1c, 0x0a, 0xfd, 0xa0, 0xb0, 0xa5, 0x85, 0xf6, 0x50, 0x28, 0xde, 0xa5, 0x87,
	0x5e, 0x8a, 0x6a, 0x3f, 0x6c, 0x83, 0x6d, 0x19, 0x49, 0x0e, 0xec, 0x3f, 0xea, 0xef, 0xea, 0x2f,
	0x29, 0x92, 0x9e, 0xb4, 0xb1, 0x13, 0x68, 0x6f, 0x7a, 0xa3, 0x99, 0x49, 0x34, 0x6f, 0x30, 0x59,
	0x2b, 0x90, 0xfb, 0x2a, 0x03, 0x05, 0x35, 0x64, 0x3a, 0x19, 0x4c, 0x3f, 0x59, 0x27, 0x85, 0x16,
	0xf4, 0x6a, 0x80, 0xae, 0x62, 0xc9, 0x75, 0xd5, 0x16, 0xa8, 0x38, 0x1c, 0x50, 0xb0, 0x62, 0x79,
	0x6f, 0x60, 0xd1, 0xee, 0x79, 0xdd, 0x7b, 0xeb, 0x13, 0x98, 0xe7, 0x6f, 0x74, 0xd5, 0x80, 0xd2,
	0xbc, 0xe9, 0x90, 0x3b, 0x9a, 0x03, 0x0f, 0xda, 0xbc, 0x13, 0x55, 0xab, 0x15, 0xf2, 0x46, 0xb3,
	0xe7, 0xbd, 0xc0, 0x3f, 0xdc, 0x09, 0x19, 0xa8, 0xc7, 0x90, 0x67, 0xaf, 0xaa, 0xdc, 0x93, 0xc2,
	0x09, 0xef, 0xd6, 0xbf, 0xa7, 0xe4, 0xc9, 0xfd, 0x43, 0x07, 0x5f, 0x41, 0xf3, 0x3b, 0x7b, 0x43,
	0xaf, 0xc9, 0xd2, 0x71, 0x6e, 0xeb, 0x3a, 0x9a, 0xc4, 0x93, 0xed, 0x22, 0x7d, 0x04, 0xe8, 0x2b,
	0x72, 0x81, 0xbf, 0x14, 0x9d, 0xc5, 0x93, 0xed, 0xe5, 0xee, 0x9a, 0x0d, 0xd2, 0x63, 0x77, 0x6e,
	0x72, 0x66, 0xa9, 0x27, 0x1b, 0x57, 0x09, 0x35, 0xec, 0x79, 0x9b, 0x41, 0x34, 0x75, 0xae, 0x01,
	0xa0, 0x3b, 0x32, 0x77, 0x39, 0x47, 0x33, 0x6b, 0xba, 0x62, 0x87, 0xb1, 0xb3, 0xd4, 0x0e, 0x68,
	0x89, 0x4c, 0xfa, 0x86, 0x4c, 0x79, 0x01, 0xd1, 0xb9, 0x15, 0x6c, 0x4f, 0xad, 0x84, 0x7d, 0x44,
	0xec, 0xbb, 0xc1, 0x50, 0x6e, 0x44, 0xf4, 0x3d, 0x59, 0x66, 0x12, 0xb8, 0x86, 0xfc, 0x56, 0x47,
	0x73, 0xeb, 0x10, 0xb3, 0xd1, 0x52, 0xd8, 0xbd, 0x9f, 0x51, 0xf9, 0x28, 0x31, 0xfa, 0xbe, 0xcb,
	0x51, 0x7f, 0xf1, 0xbf, 0xfa, 0x20, 0xa1, 0xcf, 0x09, 0x71, 0xa4, 0xcf, 0x5c, 0x95, 0xd1, 0xc2,
	0xc6, 0x71, 0x80, 0x50, 0x4a, 0x66, 0xa5, 0xb9, 0x59, 0xc6, 0x93, 0xed, 0x32, 0xb5, 0xe7, 0xf5,
	0x9f, 0x33, 0x72, 0x35, 0x08, 0xf7, 0x1f, 0x9b, 0xa2, 0x64, 0xd6, 0xf2, 0xc6, 0xad, 0x69, 0x91,
	0xda, 0xb3, 0xf5, 0x15, 0x4a, 0xe3, 0x02, 0xec, 0xd9, 0xbc, 0x25, 0xd4, 0x0c, 0xe3, 0x8f, 0xd9,
	0xa8, 0x78, 0xec, 0x93, 0x9f, 0xfd, 0x5b, 0x02, 0x81, 0xbe, 0x23, 0xe7, 0xb6, 0x74, 0xb8, 0x89,
	0x0d, 0x3b, 0x6e, 0xa2, 0x2f, 0xc5, 0x37, 0x03, 0xa1, 0x83, 0x13, 0x8d, 0x92, 0x98, 0x1f, 0x25,
	0xb1, 0x21, 0xd3, 0x2a, 0x57, 0x98, 0xf1, 0x53, 0x16, 0x0a, 0xcc, 0xbe, 0xe4, 0xde, 0xc9, 0x10,
	0xe8, 0x0d, 0x99, 0x35, 0xa0, 0xb9, 0xcd, 0xf2, 0x72, 0xf7, 0x6c, 0x54, 0xca, 0x61, 0xc5, 0x53,
	0x4b, 0x3d, 0x15, 0xf2, 0x87, 0xb7, 0x3f, 0x5e, 0x17, 0x95, 0x2e, 0xfb, 0x5f, 0x2c, 0x13, 0x4d,
	0xb2, 0xbb, 0x51, 0xba, 0x12, 0x49, 0x21, 0x5e, 0x76, 0x35, 0x7f, 0x28, 0xa4, 0xe8, 0xdb, 0x3c,
	0x29, 0x64, 0x97, 0x25, 0x2a, 0x2b, 0xa1, 0xe1, 0xc3, 0xcf, 0xc9, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0x73, 0xad, 0x7a, 0x6c, 0x04, 0x00, 0x00,
}
