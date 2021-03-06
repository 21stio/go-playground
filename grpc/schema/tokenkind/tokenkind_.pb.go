// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tokenkind/tokenkind_.proto

/*
Package tokenkind is a generated protocol buffer package.

It is generated from these files:
	tokenkind/tokenkind_.proto

It has these top-level messages:
*/
package tokenkind

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

type TokenKind int32

const (
	TokenKind_TOKEN_JWT_TOKEN    TokenKind = 0
	TokenKind_TOKEN_BEARER_TOKEN TokenKind = 1
)

var TokenKind_name = map[int32]string{
	0: "TOKEN_JWT_TOKEN",
	1: "TOKEN_BEARER_TOKEN",
}
var TokenKind_value = map[string]int32{
	"TOKEN_JWT_TOKEN":    0,
	"TOKEN_BEARER_TOKEN": 1,
}

func (x TokenKind) Enum() *TokenKind {
	p := new(TokenKind)
	*p = x
	return p
}
func (x TokenKind) String() string {
	return proto.EnumName(TokenKind_name, int32(x))
}
func (x *TokenKind) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TokenKind_value, data, "TokenKind")
	if err != nil {
		return err
	}
	*x = TokenKind(value)
	return nil
}
func (TokenKind) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("tokenkind.TokenKind", TokenKind_name, TokenKind_value)
}

func init() { proto.RegisterFile("tokenkind/tokenkind_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0xc9, 0xcf, 0x4e,
	0xcd, 0xcb, 0xce, 0xcc, 0x4b, 0xd1, 0x87, 0xb3, 0xe2, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x38, 0xe1, 0x22, 0x5a, 0x16, 0x5c, 0x9c, 0x21, 0x20, 0x8e, 0x77, 0x66, 0x5e, 0x8a, 0x90, 0x30,
	0x17, 0x7f, 0x88, 0xbf, 0xb7, 0xab, 0x5f, 0xbc, 0x57, 0x78, 0x48, 0x3c, 0x98, 0x25, 0xc0, 0x20,
	0x24, 0xc6, 0x25, 0x04, 0x11, 0x74, 0x72, 0x75, 0x0c, 0x72, 0x0d, 0x82, 0x8a, 0x33, 0x3a, 0x99,
	0x47, 0x99, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x1b, 0x19, 0x16,
	0x97, 0x64, 0xe6, 0xeb, 0xa7, 0xe7, 0xeb, 0x16, 0xe4, 0x24, 0x56, 0xa6, 0x17, 0xe5, 0x97, 0xe6,
	0xa5, 0xe8, 0xa7, 0x17, 0x15, 0x24, 0xeb, 0x17, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0x22, 0x1c, 0x01,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xb1, 0x6f, 0xcb, 0x9a, 0x00, 0x00, 0x00,
}
