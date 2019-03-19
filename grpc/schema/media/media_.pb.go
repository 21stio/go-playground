// Code generated by protoc-gen-go. DO NOT EDIT.
// source: media/media_.proto

/*
Package media is a generated protocol buffer package.

It is generated from these files:
	media/media_.proto

It has these top-level messages:
	Media
*/
package media

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mediakind "github.com/21stio/go-playground/grpc/schema/mediakind"
import image "github.com/21stio/go-playground/grpc/schema/image"
import video "github.com/21stio/go-playground/grpc/schema/video"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Media struct {
	Kind             *mediakind.MediaKind `protobuf:"varint,1,opt,name=kind,enum=mediakind.MediaKind" json:"kind,omitempty"`
	Image            *image.Image         `protobuf:"bytes,2,opt,name=image" json:"image,omitempty"`
	Video            *video.Video         `protobuf:"bytes,3,opt,name=video" json:"video,omitempty"`
	Hash             *string              `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Media) Reset()                    { *m = Media{} }
func (m *Media) String() string            { return proto.CompactTextString(m) }
func (*Media) ProtoMessage()               {}
func (*Media) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Media) GetKind() mediakind.MediaKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return mediakind.MediaKind_MEDIA_IMAGE
}

func (m *Media) GetImage() *image.Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Media) GetVideo() *video.Video {
	if m != nil {
		return m.Video
	}
	return nil
}

func (m *Media) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*Media)(nil), "media.Media")
}

func init() { proto.RegisterFile("media/media_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xbd, 0x4e, 0xc7, 0x20,
	0x14, 0xc5, 0x83, 0xb6, 0x83, 0x68, 0x1c, 0x88, 0x03, 0xe9, 0xd4, 0x74, 0x62, 0x91, 0x9b, 0xd6,
	0x37, 0x70, 0x33, 0xc6, 0x85, 0xc1, 0xc1, 0xc5, 0x60, 0x21, 0x40, 0x94, 0xd2, 0xf4, 0xc3, 0xc4,
	0x87, 0xf0, 0x9d, 0xff, 0xe1, 0x96, 0x74, 0xf9, 0x71, 0xcf, 0x3d, 0x87, 0x43, 0x02, 0x65, 0xd1,
	0x9a, 0xa0, 0x01, 0xf9, 0x29, 0xe7, 0x25, 0x6d, 0x89, 0xd5, 0xa8, 0x9a, 0x06, 0x8f, 0xef, 0x30,
	0x19, 0x38, 0xa7, 0x12, 0x69, 0x58, 0x88, 0xda, 0x59, 0x40, 0x9e, 0xbb, 0xdf, 0x60, 0x6c, 0x02,
	0x64, 0xd9, 0x75, 0xff, 0x84, 0xd6, 0x6f, 0xf9, 0x32, 0x13, 0xb4, 0xca, 0x05, 0x9c, 0xb4, 0x44,
	0xdc, 0x0f, 0x0f, 0xf2, 0xac, 0x94, 0xe8, 0xbf, 0x86, 0xc9, 0x28, 0x4c, 0xb0, 0x8e, 0xd6, 0xd8,
	0xcb, 0xaf, 0x5a, 0x22, 0x6e, 0x87, 0x3b, 0x89, 0x4a, 0xbe, 0x64, 0xaa, 0xc3, 0xca, 0x19, 0x7c,
	0x87, 0x5f, 0x97, 0x0c, 0x2a, 0xf9, 0x9e, 0xa9, 0x0e, 0x8b, 0x31, 0x5a, 0x79, 0xbd, 0x7a, 0x5e,
	0xb5, 0x44, 0xdc, 0x28, 0x9c, 0x9f, 0x9f, 0x3e, 0x7a, 0x17, 0x36, 0xbf, 0x7f, 0xc9, 0x31, 0x45,
	0x18, 0xfa, 0x75, 0x0b, 0x09, 0x5c, 0x7a, 0x9c, 0x7f, 0xf4, 0x9f, 0x5b, 0xd2, 0x3e, 0x19, 0x70,
	0xcb, 0x3c, 0xc2, 0x3a, 0x7a, 0x1b, 0xcb, 0xb7, 0x5c, 0x02, 0x00, 0x00, 0xff, 0xff, 0x14, 0x42,
	0xb7, 0x11, 0x24, 0x01, 0x00, 0x00,
}