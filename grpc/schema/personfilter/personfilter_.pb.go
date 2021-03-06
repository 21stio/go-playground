// Code generated by protoc-gen-go. DO NOT EDIT.
// source: personfilter/personfilter_.proto

/*
Package personfilter is a generated protocol buffer package.

It is generated from these files:
	personfilter/personfilter_.proto

It has these top-level messages:
	PersonFilter
*/
package personfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import infofilter "github.com/21stio/go-playground/grpc/schema/infofilter"
import feedfilter "github.com/21stio/go-playground/grpc/schema/feedfilter"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"
import emailfilter "github.com/21stio/go-playground/grpc/schema/emailfilter"
import phonenumberfilter "github.com/21stio/go-playground/grpc/schema/phonenumberfilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PersonFilter struct {
	Info              *infofilter.InfoFilter               `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	AccountsSome      *feedfilter.SocialAccountFilter      `protobuf:"bytes,2,opt,name=accountsSome" json:"accountsSome,omitempty"`
	AccountsEvery     *feedfilter.SocialAccountFilter      `protobuf:"bytes,3,opt,name=accountsEvery" json:"accountsEvery,omitempty"`
	AccountsNone      *feedfilter.SocialAccountFilter      `protobuf:"bytes,4,opt,name=accountsNone" json:"accountsNone,omitempty"`
	FirstName         *stringfilter.StringFilter           `protobuf:"bytes,5,opt,name=firstName" json:"firstName,omitempty"`
	LastName          *stringfilter.StringFilter           `protobuf:"bytes,6,opt,name=lastName" json:"lastName,omitempty"`
	Location          *stringfilter.StringFilter           `protobuf:"bytes,7,opt,name=location" json:"location,omitempty"`
	EmailsSome        *emailfilter.EmailFilter             `protobuf:"bytes,8,opt,name=emailsSome" json:"emailsSome,omitempty"`
	EmailsEvery       *emailfilter.EmailFilter             `protobuf:"bytes,9,opt,name=emailsEvery" json:"emailsEvery,omitempty"`
	EmailsNone        *emailfilter.EmailFilter             `protobuf:"bytes,10,opt,name=emailsNone" json:"emailsNone,omitempty"`
	PhoneNumbersSome  *phonenumberfilter.PhoneNumberFilter `protobuf:"bytes,11,opt,name=phoneNumbersSome" json:"phoneNumbersSome,omitempty"`
	PhoneNumbersEvery *phonenumberfilter.PhoneNumberFilter `protobuf:"bytes,12,opt,name=phoneNumbersEvery" json:"phoneNumbersEvery,omitempty"`
	PhoneNumbersNone  *phonenumberfilter.PhoneNumberFilter `protobuf:"bytes,13,opt,name=phoneNumbersNone" json:"phoneNumbersNone,omitempty"`
	IdsSome           *idfilter.IdFilter                   `protobuf:"bytes,14,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery          *idfilter.IdFilter                   `protobuf:"bytes,15,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone           *idfilter.IdFilter                   `protobuf:"bytes,16,opt,name=idsNone" json:"idsNone,omitempty"`
	Meta              *servicefilter.TypeMetaFilter        `protobuf:"bytes,17,opt,name=meta" json:"meta,omitempty"`
	And               []*PersonFilter                      `protobuf:"bytes,18,rep,name=and" json:"and,omitempty"`
	Or                []*PersonFilter                      `protobuf:"bytes,19,rep,name=or" json:"or,omitempty"`
	Not               []*PersonFilter                      `protobuf:"bytes,20,rep,name=not" json:"not,omitempty"`
	Hash              *string                              `protobuf:"bytes,21,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized  []byte                               `json:"-"`
}

func (m *PersonFilter) Reset()                    { *m = PersonFilter{} }
func (m *PersonFilter) String() string            { return proto.CompactTextString(m) }
func (*PersonFilter) ProtoMessage()               {}
func (*PersonFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PersonFilter) GetInfo() *infofilter.InfoFilter {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PersonFilter) GetAccountsSome() *feedfilter.SocialAccountFilter {
	if m != nil {
		return m.AccountsSome
	}
	return nil
}

func (m *PersonFilter) GetAccountsEvery() *feedfilter.SocialAccountFilter {
	if m != nil {
		return m.AccountsEvery
	}
	return nil
}

func (m *PersonFilter) GetAccountsNone() *feedfilter.SocialAccountFilter {
	if m != nil {
		return m.AccountsNone
	}
	return nil
}

func (m *PersonFilter) GetFirstName() *stringfilter.StringFilter {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *PersonFilter) GetLastName() *stringfilter.StringFilter {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *PersonFilter) GetLocation() *stringfilter.StringFilter {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *PersonFilter) GetEmailsSome() *emailfilter.EmailFilter {
	if m != nil {
		return m.EmailsSome
	}
	return nil
}

func (m *PersonFilter) GetEmailsEvery() *emailfilter.EmailFilter {
	if m != nil {
		return m.EmailsEvery
	}
	return nil
}

func (m *PersonFilter) GetEmailsNone() *emailfilter.EmailFilter {
	if m != nil {
		return m.EmailsNone
	}
	return nil
}

func (m *PersonFilter) GetPhoneNumbersSome() *phonenumberfilter.PhoneNumberFilter {
	if m != nil {
		return m.PhoneNumbersSome
	}
	return nil
}

func (m *PersonFilter) GetPhoneNumbersEvery() *phonenumberfilter.PhoneNumberFilter {
	if m != nil {
		return m.PhoneNumbersEvery
	}
	return nil
}

func (m *PersonFilter) GetPhoneNumbersNone() *phonenumberfilter.PhoneNumberFilter {
	if m != nil {
		return m.PhoneNumbersNone
	}
	return nil
}

func (m *PersonFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *PersonFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *PersonFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *PersonFilter) GetMeta() *servicefilter.TypeMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PersonFilter) GetAnd() []*PersonFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PersonFilter) GetOr() []*PersonFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PersonFilter) GetNot() []*PersonFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PersonFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*PersonFilter)(nil), "personfilter.PersonFilter")
}

func init() { proto.RegisterFile("personfilter/personfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x6f, 0x1a, 0x31,
	0x10, 0x86, 0xc5, 0x47, 0x1b, 0x30, 0xa4, 0x4d, 0xdc, 0x0f, 0x59, 0x28, 0x6d, 0x51, 0xd4, 0x43,
	0x84, 0x52, 0xa3, 0xe4, 0x50, 0xa1, 0xdc, 0xda, 0x8a, 0x4a, 0x39, 0x14, 0x21, 0xe8, 0xa9, 0x97,
	0xca, 0xd9, 0x35, 0x60, 0x89, 0xf5, 0xac, 0xbc, 0x26, 0x12, 0xff, 0xac, 0x3f, 0xaf, 0xda, 0x59,
	0x3b, 0x6b, 0x40, 0x2d, 0x24, 0xb7, 0x19, 0xfb, 0x79, 0x87, 0x77, 0x86, 0x59, 0x93, 0x6e, 0x2a,
	0x4d, 0x06, 0x7a, 0xa6, 0x96, 0x56, 0x9a, 0x7e, 0x98, 0xfc, 0xe6, 0xa9, 0x01, 0x0b, 0xb4, 0x1d,
	0x1e, 0x76, 0xce, 0x94, 0x9e, 0x81, 0xa3, 0xcb, 0xd0, 0xb1, 0x9d, 0xb3, 0x99, 0x94, 0xb1, 0xbb,
	0x2d, 0x43, 0x7f, 0xdb, 0xcd, 0xac, 0x51, 0x7a, 0xee, 0xee, 0xc3, 0xc4, 0x13, 0xef, 0x65, 0x22,
	0xd4, 0xd2, 0x01, 0x41, 0xec, 0xef, 0x7b, 0xe9, 0x02, 0xb4, 0xd4, 0xab, 0xe4, 0x4e, 0x1a, 0x6f,
	0x79, 0xfb, 0xc4, 0xb3, 0x4c, 0x79, 0x27, 0x6a, 0xcb, 0xc7, 0x79, 0x26, 0xcd, 0xbd, 0x8a, 0xa4,
	0x37, 0x12, 0x66, 0x8e, 0x39, 0xff, 0xd3, 0x24, 0xed, 0x31, 0x36, 0xfe, 0x1d, 0xcf, 0x69, 0x8f,
	0xd4, 0xf3, 0x7e, 0x59, 0xa5, 0x5b, 0xb9, 0x68, 0x5d, 0xbf, 0xe5, 0x65, 0xf3, 0xfc, 0x56, 0xcf,
	0xa0, 0xa0, 0x26, 0xc8, 0xd0, 0x6f, 0xa4, 0x2d, 0xa2, 0x08, 0x56, 0xda, 0x66, 0x53, 0x48, 0x24,
	0xab, 0xa2, 0xe6, 0x03, 0x2f, 0x47, 0xc2, 0xa7, 0x10, 0x29, 0xb1, 0xfc, 0x52, 0x50, 0x4e, 0xbc,
	0x21, 0xa2, 0x43, 0x72, 0xec, 0xf3, 0xe1, 0xbd, 0x34, 0x6b, 0x56, 0x3b, 0xac, 0xca, 0xa6, 0x2a,
	0xf4, 0x32, 0x02, 0x2d, 0x59, 0xfd, 0x91, 0x5e, 0x72, 0x11, 0x1d, 0x90, 0xe6, 0x4c, 0x99, 0xcc,
	0x8e, 0x44, 0x22, 0xd9, 0x33, 0xac, 0xd0, 0xe1, 0xe1, 0x1f, 0xc8, 0xa7, 0x98, 0x38, 0x71, 0x09,
	0xd3, 0xcf, 0xa4, 0xb1, 0x14, 0x4e, 0xf8, 0x7c, 0xaf, 0xf0, 0x81, 0x45, 0x1d, 0x44, 0xc2, 0x2a,
	0xd0, 0xec, 0xe8, 0x00, 0x9d, 0x63, 0xe9, 0x80, 0x10, 0xdc, 0x9b, 0x62, 0xf0, 0x0d, 0x54, 0x32,
	0x1e, 0xac, 0x12, 0x1f, 0xe6, 0xb1, 0xd3, 0x05, 0x2c, 0xbd, 0x21, 0xad, 0x22, 0x2b, 0xa6, 0xdd,
	0xdc, 0x23, 0x0d, 0xe1, 0xf2, 0x57, 0x71, 0xc4, 0xe4, 0xb0, 0x5f, 0xc5, 0xc9, 0x8e, 0xc9, 0x09,
	0x6e, 0xf0, 0x08, 0x37, 0xb8, 0x70, 0xdd, 0x42, 0xfd, 0x47, 0xbe, 0xb3, 0xda, 0x7c, 0x5c, 0xa2,
	0xae, 0xd6, 0x8e, 0x9a, 0x4e, 0xc8, 0x69, 0x78, 0x56, 0x74, 0xd3, 0x7e, 0x44, 0xc9, 0x5d, 0xf9,
	0xb6, 0x4b, 0xec, 0xf2, 0xf8, 0xa9, 0x2e, 0xb1, 0xef, 0x4b, 0x72, 0xa4, 0xe2, 0xa2, 0xdd, 0x17,
	0x58, 0x88, 0x72, 0xff, 0x99, 0xf2, 0xdb, 0xd8, 0xc9, 0x3c, 0x42, 0x39, 0x69, 0xa8, 0xd8, 0xb5,
	0xf2, 0xf2, 0x9f, 0xf8, 0x03, 0xe3, 0xaa, 0xa3, 0xcd, 0x93, 0xff, 0x56, 0x47, 0x2f, 0x57, 0xa4,
	0x9e, 0x48, 0x2b, 0xd8, 0x29, 0xa2, 0xef, 0xf8, 0xc6, 0x83, 0xc0, 0x7f, 0xae, 0x53, 0xf9, 0x43,
	0x5a, 0xe1, 0xbf, 0xf0, 0x1c, 0xa5, 0x97, 0xa4, 0x26, 0x74, 0xcc, 0x68, 0xb7, 0x86, 0x9b, 0x19,
	0x3e, 0x91, 0x3c, 0x7c, 0x36, 0x26, 0x39, 0x46, 0x7b, 0xa4, 0x0a, 0x86, 0xbd, 0xda, 0x0b, 0x57,
	0xc1, 0xe4, 0x95, 0x35, 0x58, 0xf6, 0x7a, 0x7f, 0x65, 0x0d, 0x96, 0x52, 0x52, 0x5f, 0x88, 0x6c,
	0xc1, 0xde, 0x74, 0x2b, 0x17, 0xcd, 0x09, 0xc6, 0x5f, 0x6f, 0x7e, 0x0d, 0xe6, 0xca, 0x2e, 0x56,
	0x77, 0x3c, 0x82, 0xa4, 0x7f, 0x7d, 0x95, 0x59, 0x05, 0xfd, 0x39, 0x7c, 0x4a, 0x97, 0x62, 0x3d,
	0x37, 0xb0, 0xd2, 0x71, 0x7f, 0x6e, 0xd2, 0xa8, 0x9f, 0x45, 0x0b, 0x99, 0x88, 0x8d, 0x37, 0xff,
	0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x10, 0x83, 0xd7, 0xd1, 0x0f, 0x06, 0x00, 0x00,
}
