// Code generated by protoc-gen-go. DO NOT EDIT.
// source: modifywhateverlistrequest/modifywhateverlistrequest_.proto

/*
Package modifywhateverlistrequest is a generated protocol buffer package.

It is generated from these files:
	modifywhateverlistrequest/modifywhateverlistrequest_.proto

It has these top-level messages:
	ModifyWhateverListRequest
*/
package modifywhateverlistrequest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import requestmeta "github.com/21stio/go-playground/grpc/schema/requestmeta"
import whateverskind "github.com/21stio/go-playground/grpc/schema/whateverskind"
import id "github.com/21stio/go-playground/grpc/schema/id"
import listmodifier "github.com/21stio/go-playground/grpc/schema/listmodifier"
import modifywhateverlistresponseselect "github.com/21stio/go-playground/grpc/schema/modifywhateverlistresponseselect"
import whatever "github.com/21stio/go-playground/grpc/schema/whatever"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ModifyWhateverListRequest struct {
	Meta             *requestmeta.RequestMeta                                           `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Kind             *whateverskind.WhateversKind                                       `protobuf:"varint,2,opt,name=kind,enum=whateverskind.WhateversKind" json:"kind,omitempty"`
	Id               *id.Id                                                             `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Modifier         *listmodifier.ListModifier                                         `protobuf:"varint,4,opt,name=modifier,enum=listmodifier.ListModifier" json:"modifier,omitempty"`
	Select           *modifywhateverlistresponseselect.ModifyWhateverListResponseSelect `protobuf:"bytes,5,opt,name=select" json:"select,omitempty"`
	Whatevers        []*whatever.Whatever                                               `protobuf:"bytes,6,rep,name=whatevers" json:"whatevers,omitempty"`
	Hash             *string                                                            `protobuf:"bytes,7,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                                             `json:"-"`
}

func (m *ModifyWhateverListRequest) Reset()                    { *m = ModifyWhateverListRequest{} }
func (m *ModifyWhateverListRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyWhateverListRequest) ProtoMessage()               {}
func (*ModifyWhateverListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ModifyWhateverListRequest) GetMeta() *requestmeta.RequestMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *ModifyWhateverListRequest) GetKind() whateverskind.WhateversKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return whateverskind.WhateversKind_WHATEVER_OBJECT_LIST
}

func (m *ModifyWhateverListRequest) GetId() *id.Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ModifyWhateverListRequest) GetModifier() listmodifier.ListModifier {
	if m != nil && m.Modifier != nil {
		return *m.Modifier
	}
	return listmodifier.ListModifier_ADD
}

func (m *ModifyWhateverListRequest) GetSelect() *modifywhateverlistresponseselect.ModifyWhateverListResponseSelect {
	if m != nil {
		return m.Select
	}
	return nil
}

func (m *ModifyWhateverListRequest) GetWhatevers() []*whatever.Whatever {
	if m != nil {
		return m.Whatevers
	}
	return nil
}

func (m *ModifyWhateverListRequest) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*ModifyWhateverListRequest)(nil), "modifywhateverlistrequest.ModifyWhateverListRequest")
}

func init() {
	proto.RegisterFile("modifywhateverlistrequest/modifywhateverlistrequest_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x69, 0x57, 0xa7, 0xcb, 0xc4, 0x43, 0x0e, 0x92, 0x0d, 0x91, 0xb2, 0x53, 0x0f, 0x9a,
	0xcc, 0x1e, 0x3c, 0x78, 0x1c, 0x88, 0x8a, 0xee, 0x12, 0x0f, 0xc2, 0x2e, 0x52, 0x9b, 0xd8, 0x06,
	0xd7, 0xa6, 0x36, 0x99, 0xb2, 0x0f, 0xed, 0x77, 0x90, 0x66, 0x49, 0xed, 0x60, 0xc3, 0xdb, 0x7b,
	0xbc, 0x5f, 0xfe, 0x7d, 0xbf, 0x47, 0xc1, 0x4d, 0x21, 0x99, 0x78, 0x5f, 0x7f, 0xe7, 0x89, 0xe6,
	0x5f, 0xbc, 0x5e, 0x0a, 0xa5, 0x6b, 0xfe, 0xb9, 0xe2, 0x4a, 0x93, 0xbd, 0x93, 0x57, 0x5c, 0xd5,
	0x52, 0x4b, 0x38, 0xda, 0x4b, 0x8c, 0xcf, 0x6d, 0x51, 0x70, 0x9d, 0x90, 0x4e, 0x6d, 0x9f, 0x8e,
	0x27, 0xee, 0x91, 0xfa, 0x10, 0x25, 0x23, 0x5b, 0x9d, 0x63, 0x8e, 0x05, 0x23, 0xa2, 0xed, 0xc2,
	0x26, 0xde, 0x7c, 0x50, 0xf0, 0x9a, 0x74, 0x1b, 0x47, 0xdc, 0xef, 0x5a, 0x47, 0x55, 0xb2, 0x54,
	0x5c, 0xf1, 0x25, 0x4f, 0x77, 0x1b, 0x75, 0x01, 0x97, 0x84, 0x1c, 0xd2, 0x2e, 0x66, 0x27, 0x93,
	0x1f, 0x1f, 0x8c, 0xe6, 0x26, 0xe5, 0xc5, 0x4e, 0x9e, 0x84, 0xd2, 0x74, 0x23, 0x08, 0x2f, 0x40,
	0xd0, 0x48, 0x22, 0x2f, 0xf4, 0xa2, 0x61, 0x8c, 0x70, 0x47, 0x1c, 0x5b, 0x66, 0xce, 0x75, 0x42,
	0x0d, 0x05, 0xa7, 0x20, 0x68, 0x74, 0x91, 0x1f, 0x7a, 0xd1, 0x49, 0x7c, 0x86, 0xb7, 0x8e, 0x80,
	0x5d, 0xbe, 0x7a, 0x14, 0x25, 0xa3, 0x86, 0x84, 0xa7, 0xc0, 0x17, 0x0c, 0xf5, 0x4c, 0x7a, 0x1f,
	0x0b, 0x86, 0x1f, 0x18, 0xf5, 0x05, 0x83, 0xd7, 0xe0, 0xc8, 0x1d, 0x03, 0x05, 0x26, 0x6d, 0x8c,
	0xbb, 0x17, 0xc2, 0xcd, 0x92, 0x73, 0xdb, 0xd0, 0x96, 0x85, 0x0b, 0xd0, 0xdf, 0x88, 0xa3, 0x03,
	0x93, 0x39, 0xc3, 0xff, 0x5d, 0x08, 0xef, 0x92, 0xdf, 0x00, 0xcf, 0x06, 0xa0, 0x36, 0x11, 0x4e,
	0xc1, 0xa0, 0x15, 0x42, 0xfd, 0xb0, 0x17, 0x0d, 0x63, 0xd8, 0x2a, 0xb6, 0x76, 0xf4, 0x0f, 0x82,
	0x10, 0x04, 0x79, 0xa2, 0x72, 0x74, 0x18, 0x7a, 0xd1, 0x80, 0x9a, 0x7a, 0x76, 0xb7, 0xb8, 0xcd,
	0x84, 0xce, 0x57, 0x6f, 0x38, 0x95, 0x05, 0x89, 0xaf, 0x94, 0x16, 0x92, 0x64, 0xf2, 0xb2, 0x5a,
	0x26, 0xeb, 0xac, 0x96, 0xab, 0x92, 0x91, 0xac, 0xae, 0x52, 0xa2, 0xd2, 0x9c, 0x17, 0xc9, 0xfe,
	0x5f, 0xf6, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x6e, 0x22, 0xe1, 0xe8, 0x02, 0x00, 0x00,
}