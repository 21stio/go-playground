// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deletejobsrequestfilter/deletejobsrequestfilter_.proto

/*
Package deletejobsrequestfilter is a generated protocol buffer package.

It is generated from these files:
	deletejobsrequestfilter/deletejobsrequestfilter_.proto

It has these top-level messages:
	DeleteJobsRequestFilter
*/
package deletejobsrequestfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmetafilter "github.com/21stio/go-playground/grpc/schema/requestmetafilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeleteJobsRequestFilter struct {
	Meta             *requestmetafilter.RequestMetaFilter `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	IdsSome          *idfilter.IdFilter                   `protobuf:"bytes,2,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery         *idfilter.IdFilter                   `protobuf:"bytes,3,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone          *idfilter.IdFilter                   `protobuf:"bytes,4,opt,name=idsNone" json:"idsNone,omitempty"`
	And              []*DeleteJobsRequestFilter           `protobuf:"bytes,5,rep,name=and" json:"and,omitempty"`
	Or               []*DeleteJobsRequestFilter           `protobuf:"bytes,6,rep,name=or" json:"or,omitempty"`
	Not              []*DeleteJobsRequestFilter           `protobuf:"bytes,7,rep,name=not" json:"not,omitempty"`
	Hash             *string                              `protobuf:"bytes,8,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *DeleteJobsRequestFilter) Reset()                    { *m = DeleteJobsRequestFilter{} }
func (m *DeleteJobsRequestFilter) String() string            { return proto.CompactTextString(m) }
func (*DeleteJobsRequestFilter) ProtoMessage()               {}
func (*DeleteJobsRequestFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeleteJobsRequestFilter) GetMeta() *requestmetafilter.RequestMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *DeleteJobsRequestFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *DeleteJobsRequestFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *DeleteJobsRequestFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *DeleteJobsRequestFilter) GetAnd() []*DeleteJobsRequestFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *DeleteJobsRequestFilter) GetOr() []*DeleteJobsRequestFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *DeleteJobsRequestFilter) GetNot() []*DeleteJobsRequestFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *DeleteJobsRequestFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteJobsRequestFilter)(nil), "deletejobsrequestfilter.DeleteJobsRequestFilter")
}

func init() {
	proto.RegisterFile("deletejobsrequestfilter/deletejobsrequestfilter_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x31, 0x4f, 0xf3, 0x30,
	0x10, 0x86, 0xd5, 0xa6, 0xdf, 0xd7, 0xe2, 0x6e, 0x5e, 0x6a, 0x75, 0xaa, 0x10, 0x43, 0x84, 0xc0,
	0x81, 0x0c, 0x88, 0x11, 0x15, 0x8a, 0x04, 0x12, 0x0c, 0x61, 0x63, 0x41, 0x4e, 0x7c, 0x24, 0x46,
	0x49, 0x2e, 0xd8, 0x0e, 0x52, 0x7f, 0x1c, 0xff, 0x0d, 0x25, 0x75, 0x58, 0x2a, 0x33, 0x74, 0x73,
	0xee, 0x9e, 0xf7, 0xb9, 0xd3, 0x29, 0xe4, 0x4a, 0x42, 0x09, 0x16, 0x3e, 0x30, 0x35, 0x1a, 0x3e,
	0x5b, 0x30, 0xf6, 0x5d, 0x95, 0x16, 0x74, 0xe4, 0xa9, 0xbf, 0xf1, 0x46, 0xa3, 0x45, 0xba, 0xf0,
	0xf4, 0x97, 0xa7, 0xee, 0xb3, 0x02, 0x2b, 0x9c, 0x6a, 0xaf, 0xe2, 0x24, 0x4b, 0xa6, 0xa4, 0x43,
	0x86, 0x87, 0xeb, 0x1c, 0x7f, 0x07, 0x64, 0x71, 0xd7, 0x4f, 0x78, 0xc4, 0xd4, 0x24, 0x3b, 0xc1,
	0x7d, 0x8f, 0xd0, 0x6b, 0x32, 0xe9, 0x54, 0x6c, 0xb4, 0x1a, 0x85, 0xf3, 0xf8, 0x84, 0xef, 0xe9,
	0xb9, 0xe3, 0x9f, 0xc0, 0x8a, 0x5d, 0x26, 0xe9, 0x13, 0xf4, 0x8c, 0x4c, 0x95, 0x34, 0x2f, 0x58,
	0x01, 0x1b, 0xf7, 0x61, 0xca, 0x87, 0xc1, 0xfc, 0x41, 0x3a, 0x74, 0x40, 0x28, 0x27, 0x33, 0x25,
	0xcd, 0xe6, 0x0b, 0xf4, 0x96, 0x05, 0x5e, 0xfc, 0x97, 0x71, 0xf6, 0x67, 0xac, 0x81, 0x4d, 0xfe,
	0xb4, 0x77, 0x08, 0x5d, 0x93, 0x40, 0xd4, 0x92, 0xfd, 0x5b, 0x05, 0xe1, 0x3c, 0xbe, 0xe0, 0x9e,
	0x73, 0x72, 0xcf, 0x11, 0x92, 0x2e, 0x4c, 0x6f, 0xc8, 0x18, 0x35, 0xfb, 0x7f, 0xa0, 0x62, 0x8c,
	0xba, 0xdb, 0xa2, 0x46, 0xcb, 0xa6, 0x87, 0x6e, 0x51, 0xa3, 0xa5, 0x94, 0x4c, 0x0a, 0x61, 0x0a,
	0x36, 0x5b, 0x8d, 0xc2, 0xa3, 0xa4, 0x7f, 0xaf, 0x37, 0xaf, 0xb7, 0xb9, 0xb2, 0x45, 0x9b, 0xf2,
	0x0c, 0xab, 0x28, 0xbe, 0x34, 0x56, 0x61, 0x94, 0xe3, 0x79, 0x53, 0x8a, 0x6d, 0xae, 0xb1, 0xad,
	0x65, 0x94, 0xeb, 0x26, 0x8b, 0x4c, 0x56, 0x40, 0x25, 0x7c, 0x3f, 0xdb, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xd6, 0xe6, 0x86, 0xeb, 0x9e, 0x02, 0x00, 0x00,
}