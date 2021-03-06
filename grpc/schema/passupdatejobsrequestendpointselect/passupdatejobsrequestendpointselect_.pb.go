// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatejobsrequestendpointselect/passupdatejobsrequestendpointselect_.proto

/*
Package passupdatejobsrequestendpointselect is a generated protocol buffer package.

It is generated from these files:
	passupdatejobsrequestendpointselect/passupdatejobsrequestendpointselect_.proto

It has these top-level messages:
	PassUpdateJobsRequestEndpointSelect
*/
package passupdatejobsrequestendpointselect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdateJobsRequestEndpointSelect struct {
	SelectAll        *bool   `protobuf:"varint,1,opt,name=selectAll" json:"selectAll,omitempty"`
	SelectHash       *bool   `protobuf:"varint,2,opt,name=selectHash" json:"selectHash,omitempty"`
	Hash             *string `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PassUpdateJobsRequestEndpointSelect) Reset()         { *m = PassUpdateJobsRequestEndpointSelect{} }
func (m *PassUpdateJobsRequestEndpointSelect) String() string { return proto.CompactTextString(m) }
func (*PassUpdateJobsRequestEndpointSelect) ProtoMessage()    {}
func (*PassUpdateJobsRequestEndpointSelect) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *PassUpdateJobsRequestEndpointSelect) GetSelectAll() bool {
	if m != nil && m.SelectAll != nil {
		return *m.SelectAll
	}
	return false
}

func (m *PassUpdateJobsRequestEndpointSelect) GetSelectHash() bool {
	if m != nil && m.SelectHash != nil {
		return *m.SelectHash
	}
	return false
}

func (m *PassUpdateJobsRequestEndpointSelect) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdateJobsRequestEndpointSelect)(nil), "passupdatejobsrequestendpointselect.PassUpdateJobsRequestEndpointSelect")
}

func init() {
	proto.RegisterFile("passupdatejobsrequestendpointselect/passupdatejobsrequestendpointselect_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xf2, 0x2b, 0x48, 0x2c, 0x2e,
	0x2e, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0xcd, 0xca, 0x4f, 0x2a, 0x2e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x49, 0xcd, 0x4b, 0x29, 0xc8, 0xcf, 0xcc, 0x2b, 0x29, 0x4e, 0xcd, 0x49, 0x4d, 0x2e, 0xd1,
	0x27, 0x42, 0x4d, 0xbc, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x32, 0x11, 0x6a, 0x95, 0xca,
	0xb9, 0x94, 0x03, 0x12, 0x8b, 0x8b, 0x43, 0xc1, 0xca, 0xbc, 0xf2, 0x93, 0x8a, 0x83, 0x20, 0xca,
	0x5c, 0xa1, 0xca, 0x82, 0xc1, 0xca, 0x84, 0x64, 0xb8, 0x38, 0x21, 0x1a, 0x1c, 0x73, 0x72, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x82, 0x10, 0x02, 0x42, 0x72, 0x5c, 0x5c, 0x10, 0x8e, 0x47, 0x62,
	0x71, 0x86, 0x04, 0x13, 0x58, 0x1a, 0x49, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x03, 0x24, 0xc3, 0xac,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x3b, 0x05, 0x46, 0xf9, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x1b, 0x19, 0x16, 0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16,
	0xe4, 0x24, 0x56, 0xa6, 0x17, 0xe5, 0x97, 0xe6, 0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17,
	0x27, 0x67, 0xa4, 0xe6, 0x26, 0x12, 0xe3, 0x6f, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xd7,
	0xce, 0x34, 0x41, 0x01, 0x00, 0x00,
}
