// Code generated by protoc-gen-go. DO NOT EDIT.
// source: passupdatepostsrequestresponse/passupdatepostsrequestresponse_.proto

/*
Package passupdatepostsrequestresponse is a generated protocol buffer package.

It is generated from these files:
	passupdatepostsrequestresponse/passupdatepostsrequestresponse_.proto

It has these top-level messages:
	PassUpdatePostsRequestResponse
*/
package passupdatepostsrequestresponse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import responsemeta "github.com/21stio/go-playground/grpc/schema/responsemeta"
import updatepostsrequest "github.com/21stio/go-playground/grpc/schema/updatepostsrequest"
import error "github.com/21stio/go-playground/grpc/schema/error"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PassUpdatePostsRequestResponse struct {
	Meta               *responsemeta.ResponseMeta             `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	UpdatePostsRequest *updatepostsrequest.UpdatePostsRequest `protobuf:"bytes,2,opt,name=updatePostsRequest" json:"updatePostsRequest,omitempty"`
	Errors             []*error.Error                         `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	Hash               *string                                `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized   []byte                                 `json:"-"`
}

func (m *PassUpdatePostsRequestResponse) Reset()                    { *m = PassUpdatePostsRequestResponse{} }
func (m *PassUpdatePostsRequestResponse) String() string            { return proto.CompactTextString(m) }
func (*PassUpdatePostsRequestResponse) ProtoMessage()               {}
func (*PassUpdatePostsRequestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PassUpdatePostsRequestResponse) GetMeta() *responsemeta.ResponseMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PassUpdatePostsRequestResponse) GetUpdatePostsRequest() *updatepostsrequest.UpdatePostsRequest {
	if m != nil {
		return m.UpdatePostsRequest
	}
	return nil
}

func (m *PassUpdatePostsRequestResponse) GetErrors() []*error.Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *PassUpdatePostsRequestResponse) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PassUpdatePostsRequestResponse)(nil), "passupdatepostsrequestresponse.PassUpdatePostsRequestResponse")
}

func init() {
	proto.RegisterFile("passupdatepostsrequestresponse/passupdatepostsrequestresponse_.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xa6, 0x6e, 0x11, 0xcc, 0x7a, 0xca, 0xa9, 0xf4, 0xb0, 0x14, 0x11, 0xe9, 0x41, 0x27, 0xd8,
	0x47, 0x10, 0xbd, 0x88, 0xc2, 0x12, 0xd0, 0x83, 0x17, 0x89, 0xdd, 0xa1, 0x5d, 0xb0, 0x3b, 0x31,
	0x93, 0x1c, 0x7c, 0x5d, 0x9f, 0x44, 0x1a, 0x5b, 0x58, 0xc8, 0xb2, 0x5e, 0x86, 0xf9, 0xf9, 0xfe,
	0x60, 0xc4, 0xbd, 0x35, 0xcc, 0xc1, 0x6e, 0x8c, 0x47, 0x4b, 0xec, 0xd9, 0xe1, 0x57, 0x40, 0xf6,
	0x0e, 0xd9, 0xd2, 0x8e, 0x51, 0x1d, 0x3f, 0xbf, 0x83, 0x75, 0xe4, 0x49, 0xae, 0x8e, 0xc3, 0xca,
	0x6a, 0xee, 0x06, 0xf4, 0x46, 0xed, 0x0f, 0x93, 0x42, 0x79, 0x9d, 0xb2, 0x55, 0xba, 0x9a, 0xd1,
	0x12, 0x9d, 0x23, 0xa7, 0x62, 0x9d, 0x76, 0x17, 0x3f, 0x99, 0x58, 0xad, 0x0d, 0xf3, 0x4b, 0x64,
	0xad, 0x47, 0x96, 0xfe, 0x63, 0xe9, 0xc9, 0x4f, 0x82, 0xc8, 0x47, 0xcf, 0x22, 0xab, 0xb2, 0x7a,
	0xd9, 0x94, 0xb0, 0x1f, 0x04, 0x66, 0xd4, 0x33, 0x7a, 0xa3, 0x23, 0x4e, 0xbe, 0x0a, 0x19, 0x12,
	0xb5, 0xe2, 0x24, 0xb2, 0xaf, 0x20, 0x8d, 0x07, 0x07, 0xbc, 0x0f, 0x28, 0xc8, 0x4b, 0x71, 0x1a,
	0xa3, 0x73, 0xb1, 0xa8, 0x16, 0xf5, 0xb2, 0x39, 0x87, 0x38, 0xc2, 0xc3, 0x58, 0xf5, 0x74, 0x93,
	0x52, 0xe4, 0xbd, 0xe1, 0xbe, 0xc8, 0xab, 0xac, 0x3e, 0xd3, 0xb1, 0xbf, 0x7b, 0x7a, 0x7b, 0xec,
	0xb6, 0xbe, 0x0f, 0x1f, 0xd0, 0xd2, 0xa0, 0x9a, 0x5b, 0xf6, 0x5b, 0x52, 0x1d, 0xdd, 0xd8, 0x4f,
	0xf3, 0xdd, 0x39, 0x0a, 0xbb, 0x8d, 0xea, 0x9c, 0x6d, 0x15, 0xb7, 0x3d, 0x0e, 0xe6, 0x9f, 0xef,
	0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x87, 0xae, 0xff, 0xfd, 0x01, 0x00, 0x00,
}
