// Code generated by protoc-gen-go. DO NOT EDIT.
// source: companyfilter/companyfilter_.proto

/*
Package companyfilter is a generated protocol buffer package.

It is generated from these files:
	companyfilter/companyfilter_.proto

It has these top-level messages:
	CompanyFilter
	JobFilter
*/
package companyfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import companykindfilter "github.com/21stio/go-playground/grpc/schema/companykindfilter"
import infofilter "github.com/21stio/go-playground/grpc/schema/infofilter"
import industrykindfilter "github.com/21stio/go-playground/grpc/schema/industrykindfilter"
import locationfilter "github.com/21stio/go-playground/grpc/schema/locationfilter"
import intfilter "github.com/21stio/go-playground/grpc/schema/intfilter"
import currencyvaluefilter "github.com/21stio/go-playground/grpc/schema/currencyvaluefilter"
import personfilter "github.com/21stio/go-playground/grpc/schema/personfilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import jobkindfilter "github.com/21stio/go-playground/grpc/schema/jobkindfilter"
import stringfilter "github.com/21stio/go-playground/grpc/schema/stringfilter"
import experiencelevelkindfilter "github.com/21stio/go-playground/grpc/schema/experiencelevelkindfilter"
import boolfilter "github.com/21stio/go-playground/grpc/schema/boolfilter"
import currencyperdurationfilter "github.com/21stio/go-playground/grpc/schema/currencyperdurationfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CompanyFilter struct {
	Kind             *companykindfilter.CompanyKindFilter     `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Info             *infofilter.InfoFilter                   `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	IndustryKind     *industrykindfilter.IndustryKindFilter   `protobuf:"bytes,3,opt,name=industryKind" json:"industryKind,omitempty"`
	HeadQuarter      *locationfilter.LocationFilter           `protobuf:"bytes,4,opt,name=headQuarter" json:"headQuarter,omitempty"`
	EmployeeCount    *intfilter.IntFilter                     `protobuf:"bytes,5,opt,name=employeeCount" json:"employeeCount,omitempty"`
	Revenue          *currencyvaluefilter.CurrencyValueFilter `protobuf:"bytes,6,opt,name=revenue" json:"revenue,omitempty"`
	JobsSome         *JobFilter                               `protobuf:"bytes,7,opt,name=jobsSome" json:"jobsSome,omitempty"`
	JobsEvery        *JobFilter                               `protobuf:"bytes,8,opt,name=jobsEvery" json:"jobsEvery,omitempty"`
	JobsNone         *JobFilter                               `protobuf:"bytes,9,opt,name=jobsNone" json:"jobsNone,omitempty"`
	EmployeesSome    *personfilter.PersonFilter               `protobuf:"bytes,10,opt,name=employeesSome" json:"employeesSome,omitempty"`
	EmployeesEvery   *personfilter.PersonFilter               `protobuf:"bytes,11,opt,name=employeesEvery" json:"employeesEvery,omitempty"`
	EmployeesNone    *personfilter.PersonFilter               `protobuf:"bytes,12,opt,name=employeesNone" json:"employeesNone,omitempty"`
	IdsSome          *idfilter.IdFilter                       `protobuf:"bytes,13,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery         *idfilter.IdFilter                       `protobuf:"bytes,14,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone          *idfilter.IdFilter                       `protobuf:"bytes,15,opt,name=idsNone" json:"idsNone,omitempty"`
	Meta             *servicefilter.TypeMetaFilter            `protobuf:"bytes,16,opt,name=meta" json:"meta,omitempty"`
	And              []*CompanyFilter                         `protobuf:"bytes,17,rep,name=and" json:"and,omitempty"`
	Or               []*CompanyFilter                         `protobuf:"bytes,18,rep,name=or" json:"or,omitempty"`
	Not              []*CompanyFilter                         `protobuf:"bytes,19,rep,name=not" json:"not,omitempty"`
	Hash             *string                                  `protobuf:"bytes,20,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                   `json:"-"`
}

func (m *CompanyFilter) Reset()                    { *m = CompanyFilter{} }
func (m *CompanyFilter) String() string            { return proto.CompactTextString(m) }
func (*CompanyFilter) ProtoMessage()               {}
func (*CompanyFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CompanyFilter) GetKind() *companykindfilter.CompanyKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *CompanyFilter) GetInfo() *infofilter.InfoFilter {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *CompanyFilter) GetIndustryKind() *industrykindfilter.IndustryKindFilter {
	if m != nil {
		return m.IndustryKind
	}
	return nil
}

func (m *CompanyFilter) GetHeadQuarter() *locationfilter.LocationFilter {
	if m != nil {
		return m.HeadQuarter
	}
	return nil
}

func (m *CompanyFilter) GetEmployeeCount() *intfilter.IntFilter {
	if m != nil {
		return m.EmployeeCount
	}
	return nil
}

func (m *CompanyFilter) GetRevenue() *currencyvaluefilter.CurrencyValueFilter {
	if m != nil {
		return m.Revenue
	}
	return nil
}

func (m *CompanyFilter) GetJobsSome() *JobFilter {
	if m != nil {
		return m.JobsSome
	}
	return nil
}

func (m *CompanyFilter) GetJobsEvery() *JobFilter {
	if m != nil {
		return m.JobsEvery
	}
	return nil
}

func (m *CompanyFilter) GetJobsNone() *JobFilter {
	if m != nil {
		return m.JobsNone
	}
	return nil
}

func (m *CompanyFilter) GetEmployeesSome() *personfilter.PersonFilter {
	if m != nil {
		return m.EmployeesSome
	}
	return nil
}

func (m *CompanyFilter) GetEmployeesEvery() *personfilter.PersonFilter {
	if m != nil {
		return m.EmployeesEvery
	}
	return nil
}

func (m *CompanyFilter) GetEmployeesNone() *personfilter.PersonFilter {
	if m != nil {
		return m.EmployeesNone
	}
	return nil
}

func (m *CompanyFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *CompanyFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *CompanyFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *CompanyFilter) GetMeta() *servicefilter.TypeMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *CompanyFilter) GetAnd() []*CompanyFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *CompanyFilter) GetOr() []*CompanyFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *CompanyFilter) GetNot() []*CompanyFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *CompanyFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

type JobFilter struct {
	Kind                 *jobkindfilter.JobKindFilter                         `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Info                 *infofilter.InfoFilter                               `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	TagsSome             *stringfilter.StringFilter                           `protobuf:"bytes,3,opt,name=tagsSome" json:"tagsSome,omitempty"`
	TagsEvery            *stringfilter.StringFilter                           `protobuf:"bytes,4,opt,name=tagsEvery" json:"tagsEvery,omitempty"`
	TagsNone             *stringfilter.StringFilter                           `protobuf:"bytes,5,opt,name=tagsNone" json:"tagsNone,omitempty"`
	BenefitsSome         *stringfilter.StringFilter                           `protobuf:"bytes,6,opt,name=benefitsSome" json:"benefitsSome,omitempty"`
	BenefitsEvery        *stringfilter.StringFilter                           `protobuf:"bytes,7,opt,name=benefitsEvery" json:"benefitsEvery,omitempty"`
	BenefitsNone         *stringfilter.StringFilter                           `protobuf:"bytes,8,opt,name=benefitsNone" json:"benefitsNone,omitempty"`
	ExperienceLevelSome  *experiencelevelkindfilter.ExperienceLevelKindFilter `protobuf:"bytes,9,opt,name=experienceLevelSome" json:"experienceLevelSome,omitempty"`
	ExperienceLevelEvery *experiencelevelkindfilter.ExperienceLevelKindFilter `protobuf:"bytes,10,opt,name=experienceLevelEvery" json:"experienceLevelEvery,omitempty"`
	ExperienceLevelNone  *experiencelevelkindfilter.ExperienceLevelKindFilter `protobuf:"bytes,11,opt,name=experienceLevelNone" json:"experienceLevelNone,omitempty"`
	Remote               *boolfilter.BoolFilter                               `protobuf:"bytes,12,opt,name=remote" json:"remote,omitempty"`
	SponsorsVisa         *boolfilter.BoolFilter                               `protobuf:"bytes,13,opt,name=sponsorsVisa" json:"sponsorsVisa,omitempty"`
	PaysRelocation       *boolfilter.BoolFilter                               `protobuf:"bytes,14,opt,name=paysRelocation" json:"paysRelocation,omitempty"`
	ProvidesEquity       *boolfilter.BoolFilter                               `protobuf:"bytes,15,opt,name=providesEquity" json:"providesEquity,omitempty"`
	Salary               *currencyperdurationfilter.CurrencyPerDurationFilter `protobuf:"bytes,16,opt,name=salary" json:"salary,omitempty"`
	Company              *CompanyFilter                                       `protobuf:"bytes,17,opt,name=company" json:"company,omitempty"`
	Location             *locationfilter.LocationFilter                       `protobuf:"bytes,18,opt,name=location" json:"location,omitempty"`
	IdsSome              *idfilter.IdFilter                                   `protobuf:"bytes,19,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery             *idfilter.IdFilter                                   `protobuf:"bytes,20,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone              *idfilter.IdFilter                                   `protobuf:"bytes,21,opt,name=idsNone" json:"idsNone,omitempty"`
	Meta                 *servicefilter.TypeMetaFilter                        `protobuf:"bytes,22,opt,name=meta" json:"meta,omitempty"`
	And                  []*JobFilter                                         `protobuf:"bytes,23,rep,name=and" json:"and,omitempty"`
	Or                   []*JobFilter                                         `protobuf:"bytes,24,rep,name=or" json:"or,omitempty"`
	Not                  []*JobFilter                                         `protobuf:"bytes,25,rep,name=not" json:"not,omitempty"`
	Hash                 *string                                              `protobuf:"bytes,26,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized     []byte                                               `json:"-"`
}

func (m *JobFilter) Reset()                    { *m = JobFilter{} }
func (m *JobFilter) String() string            { return proto.CompactTextString(m) }
func (*JobFilter) ProtoMessage()               {}
func (*JobFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *JobFilter) GetKind() *jobkindfilter.JobKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *JobFilter) GetInfo() *infofilter.InfoFilter {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *JobFilter) GetTagsSome() *stringfilter.StringFilter {
	if m != nil {
		return m.TagsSome
	}
	return nil
}

func (m *JobFilter) GetTagsEvery() *stringfilter.StringFilter {
	if m != nil {
		return m.TagsEvery
	}
	return nil
}

func (m *JobFilter) GetTagsNone() *stringfilter.StringFilter {
	if m != nil {
		return m.TagsNone
	}
	return nil
}

func (m *JobFilter) GetBenefitsSome() *stringfilter.StringFilter {
	if m != nil {
		return m.BenefitsSome
	}
	return nil
}

func (m *JobFilter) GetBenefitsEvery() *stringfilter.StringFilter {
	if m != nil {
		return m.BenefitsEvery
	}
	return nil
}

func (m *JobFilter) GetBenefitsNone() *stringfilter.StringFilter {
	if m != nil {
		return m.BenefitsNone
	}
	return nil
}

func (m *JobFilter) GetExperienceLevelSome() *experiencelevelkindfilter.ExperienceLevelKindFilter {
	if m != nil {
		return m.ExperienceLevelSome
	}
	return nil
}

func (m *JobFilter) GetExperienceLevelEvery() *experiencelevelkindfilter.ExperienceLevelKindFilter {
	if m != nil {
		return m.ExperienceLevelEvery
	}
	return nil
}

func (m *JobFilter) GetExperienceLevelNone() *experiencelevelkindfilter.ExperienceLevelKindFilter {
	if m != nil {
		return m.ExperienceLevelNone
	}
	return nil
}

func (m *JobFilter) GetRemote() *boolfilter.BoolFilter {
	if m != nil {
		return m.Remote
	}
	return nil
}

func (m *JobFilter) GetSponsorsVisa() *boolfilter.BoolFilter {
	if m != nil {
		return m.SponsorsVisa
	}
	return nil
}

func (m *JobFilter) GetPaysRelocation() *boolfilter.BoolFilter {
	if m != nil {
		return m.PaysRelocation
	}
	return nil
}

func (m *JobFilter) GetProvidesEquity() *boolfilter.BoolFilter {
	if m != nil {
		return m.ProvidesEquity
	}
	return nil
}

func (m *JobFilter) GetSalary() *currencyperdurationfilter.CurrencyPerDurationFilter {
	if m != nil {
		return m.Salary
	}
	return nil
}

func (m *JobFilter) GetCompany() *CompanyFilter {
	if m != nil {
		return m.Company
	}
	return nil
}

func (m *JobFilter) GetLocation() *locationfilter.LocationFilter {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *JobFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *JobFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *JobFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *JobFilter) GetMeta() *servicefilter.TypeMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *JobFilter) GetAnd() []*JobFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *JobFilter) GetOr() []*JobFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *JobFilter) GetNot() []*JobFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *JobFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*CompanyFilter)(nil), "companyfilter.CompanyFilter")
	proto.RegisterType((*JobFilter)(nil), "companyfilter.JobFilter")
}

func init() { proto.RegisterFile("companyfilter/companyfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 967 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x57, 0xd1, 0x6e, 0x23, 0x35,
	0x14, 0x55, 0xdb, 0x6c, 0x93, 0x38, 0x69, 0x61, 0xdd, 0x52, 0x4c, 0x54, 0x50, 0x55, 0xad, 0x50,
	0x54, 0x15, 0x87, 0xad, 0xaa, 0x6a, 0x29, 0xd2, 0x6a, 0xd5, 0x52, 0xa4, 0x2e, 0x05, 0x2d, 0x5d,
	0xb4, 0x0f, 0xbc, 0xa0, 0x49, 0xe6, 0x26, 0xf1, 0x32, 0xb1, 0x07, 0x8f, 0x27, 0x62, 0x7e, 0x83,
	0x2f, 0xe3, 0x93, 0xd0, 0x78, 0x6c, 0xcf, 0x38, 0xcd, 0x6c, 0x52, 0x2d, 0x6f, 0xbe, 0xf6, 0x39,
	0xc7, 0xd7, 0xf6, 0x9d, 0x7b, 0x12, 0x74, 0x3c, 0x12, 0xb3, 0x38, 0xe0, 0xd9, 0x98, 0x45, 0x0a,
	0xe4, 0xc0, 0x8b, 0xfe, 0xa0, 0xb1, 0x14, 0x4a, 0xe0, 0x1d, 0x6f, 0xb6, 0x77, 0x62, 0xc2, 0x3f,
	0x19, 0x0f, 0x7d, 0x5a, 0x39, 0x63, 0xa8, 0xbd, 0x43, 0xc6, 0xc7, 0xc2, 0x80, 0xca, 0xa1, 0x5d,
	0x3d, 0x65, 0x3c, 0x4c, 0x13, 0x25, 0xab, 0x52, 0x0f, 0xa7, 0x2c, 0xfa, 0x59, 0x24, 0x46, 0x81,
	0x62, 0x82, 0x1b, 0xa4, 0x1f, 0x5a, 0x54, 0x8f, 0x71, 0xe5, 0xa4, 0x94, 0xbf, 0x46, 0x47, 0xa9,
	0x94, 0xc0, 0x47, 0xd9, 0x3c, 0x88, 0x52, 0xb0, 0xb9, 0x3f, 0x9c, 0xb3, 0xf8, 0xa3, 0x18, 0x64,
	0xe2, 0xf6, 0xab, 0x06, 0x16, 0x41, 0x98, 0xcb, 0x7b, 0x21, 0xdb, 0xe3, 0x04, 0xe4, 0x9c, 0x8d,
	0xec, 0x2e, 0x5e, 0xe4, 0x30, 0xef, 0xc5, 0xb0, 0x72, 0x74, 0x2f, 0x72, 0x39, 0x24, 0x4a, 0x32,
	0x3e, 0xb1, 0x32, 0x95, 0xc0, 0x22, 0x2e, 0xe1, 0xef, 0x18, 0x24, 0x03, 0x3e, 0x82, 0x08, 0xe6,
	0x10, 0x55, 0x14, 0x6b, 0x57, 0xdc, 0xfb, 0x0c, 0x85, 0x88, 0x0c, 0xb8, 0x1c, 0x3a, 0x65, 0x7b,
	0x37, 0x31, 0xc8, 0x30, 0x95, 0xd5, 0xcb, 0xaf, 0x5d, 0x31, 0xdc, 0xe3, 0x7f, 0x5b, 0x68, 0xe7,
	0xba, 0x28, 0x8b, 0x1f, 0xf5, 0x02, 0x7e, 0x81, 0x1a, 0x79, 0x02, 0x64, 0xe3, 0x68, 0xa3, 0xdf,
	0x39, 0x7b, 0x46, 0x1f, 0x14, 0x0d, 0x35, 0xf8, 0x9f, 0x18, 0x0f, 0x0b, 0xce, 0xbd, 0x66, 0xe0,
	0x13, 0xd4, 0xc8, 0x8b, 0x87, 0x6c, 0x6a, 0xe6, 0x01, 0x2d, 0x2b, 0x89, 0xde, 0xf2, 0xb1, 0xb0,
	0xd8, 0x7c, 0x1a, 0xbf, 0x46, 0x5d, 0x5b, 0x42, 0xb9, 0x0e, 0xd9, 0xd2, 0x9c, 0xaf, 0xe9, 0xc3,
	0xba, 0xa2, 0xb7, 0x15, 0x9c, 0xd1, 0xf0, 0xb8, 0xf8, 0x15, 0xea, 0x4c, 0x21, 0x08, 0x7f, 0x4d,
	0x03, 0xa9, 0x40, 0x92, 0x86, 0x96, 0xfa, 0x8a, 0xfa, 0x85, 0x47, 0xef, 0x4c, 0x68, 0x24, 0xaa,
	0x14, 0x7c, 0x89, 0x76, 0x60, 0x16, 0x47, 0x22, 0x03, 0xb8, 0x16, 0x29, 0x57, 0xe4, 0x89, 0xd6,
	0xd8, 0xa7, 0xae, 0x36, 0xe9, 0x2d, 0x57, 0x86, 0xe9, 0x43, 0xf1, 0x15, 0x6a, 0x4a, 0x98, 0x03,
	0x4f, 0x81, 0x6c, 0x6b, 0x56, 0x7f, 0x59, 0xfd, 0xd2, 0x6b, 0x33, 0xf7, 0x2e, 0x9f, 0x33, 0x4a,
	0x96, 0x88, 0xcf, 0x51, 0xeb, 0xbd, 0x18, 0x26, 0x6f, 0xc5, 0x0c, 0x48, 0x53, 0x8b, 0x10, 0xea,
	0x7d, 0xcd, 0xf4, 0xb5, 0x18, 0x1a, 0x92, 0x43, 0xe2, 0x0b, 0xd4, 0xce, 0xc7, 0x37, 0x73, 0x90,
	0x19, 0x69, 0xad, 0xa0, 0x95, 0x50, 0xbb, 0xdb, 0x2f, 0x82, 0x03, 0x69, 0xaf, 0xb3, 0x5b, 0x8e,
	0xc4, 0xaf, 0xca, 0x3b, 0x2a, 0x12, 0x45, 0x9a, 0xda, 0xa3, 0xd5, 0x0f, 0x8e, 0xbe, 0xd1, 0xc1,
	0xe2, 0x4d, 0x15, 0xf9, 0x5e, 0xa1, 0x5d, 0x37, 0x51, 0x24, 0xdd, 0x59, 0x29, 0xb1, 0xc0, 0xf0,
	0xb2, 0xd0, 0x07, 0xe8, 0x3e, 0x22, 0x0b, 0x7d, 0x8e, 0x53, 0xd4, 0x64, 0x61, 0x71, 0x82, 0x1d,
	0xcd, 0xc5, 0x94, 0xb9, 0x52, 0xb3, 0x05, 0x66, 0x21, 0x98, 0xa2, 0x16, 0x0b, 0x4d, 0xb6, 0xbb,
	0xb5, 0x70, 0x87, 0x31, 0xea, 0x3a, 0xb3, 0x4f, 0x3e, 0xa8, 0xae, 0x73, 0x79, 0x8e, 0x1a, 0x33,
	0x50, 0x01, 0xf9, 0x54, 0x43, 0xbf, 0xa4, 0x5e, 0xfb, 0xa1, 0xbf, 0x65, 0x31, 0xfc, 0x0c, 0x2a,
	0xb0, 0x1f, 0x4e, 0x0e, 0xc5, 0x14, 0x6d, 0x05, 0x3c, 0x24, 0x4f, 0x8f, 0xb6, 0xfa, 0x9d, 0xb3,
	0xc3, 0x85, 0x77, 0xf3, 0xbe, 0xe4, 0xfb, 0x1c, 0x88, 0x4f, 0xd1, 0xa6, 0x90, 0x04, 0xaf, 0x01,
	0xdf, 0x14, 0x32, 0x57, 0xe7, 0x42, 0x91, 0xbd, 0x75, 0xd4, 0xb9, 0x50, 0x18, 0xa3, 0xc6, 0x34,
	0x48, 0xa6, 0x64, 0xff, 0x68, 0xa3, 0xdf, 0xbe, 0xd7, 0xe3, 0xe3, 0x7f, 0xba, 0xa8, 0xed, 0x0a,
	0x08, 0x7f, 0xeb, 0xb5, 0x93, 0x43, 0xea, 0x75, 0xcf, 0xbc, 0xd0, 0x3e, 0xaa, 0x8d, 0x5c, 0xa0,
	0x96, 0x0a, 0x26, 0xc5, 0x6b, 0x6e, 0x99, 0x4a, 0xa8, 0x36, 0x5f, 0xfa, 0x56, 0x07, 0xf6, 0x99,
	0x2c, 0x16, 0xbf, 0x40, 0xed, 0x7c, 0x5c, 0xbc, 0x6b, 0x63, 0x25, 0xb1, 0x04, 0xdb, 0x1d, 0xf5,
	0x0b, 0x3f, 0x59, 0x6f, 0x47, 0xfd, 0xd4, 0x2f, 0x51, 0x77, 0x08, 0x1c, 0xc6, 0x4c, 0x15, 0xd9,
	0x6e, 0xaf, 0xe4, 0x7a, 0xf8, 0xbc, 0xf0, 0x6d, 0x5c, 0x64, 0xdd, 0x5c, 0x29, 0xe0, 0x13, 0xaa,
	0x19, 0xe8, 0xec, 0x5b, 0xeb, 0x67, 0xa0, 0x4f, 0x30, 0x46, 0x7b, 0xa5, 0x51, 0xdd, 0xe5, 0x46,
	0xa5, 0x0f, 0x52, 0x74, 0x90, 0x73, 0x5a, 0x6b, 0x62, 0xf4, 0xc6, 0x67, 0x55, 0x1e, 0x7c, 0x99,
	0x20, 0x9e, 0xa2, 0xfd, 0x85, 0xe9, 0xe2, 0xc0, 0xe8, 0x23, 0x36, 0x5a, 0xaa, 0xb8, 0xe4, 0x44,
	0xfa, 0x62, 0x3a, 0xff, 0xe3, 0x89, 0xf4, 0xcd, 0x51, 0xb4, 0x2d, 0x61, 0x26, 0x94, 0xed, 0x56,
	0x07, 0xb4, 0x34, 0x71, 0x7a, 0x25, 0x44, 0x64, 0xc8, 0x06, 0x85, 0x2f, 0x51, 0x37, 0x89, 0x05,
	0x4f, 0x84, 0x4c, 0xde, 0xb1, 0x24, 0x30, 0x7d, 0xaa, 0x8e, 0xe5, 0x61, 0xf1, 0x4b, 0xb4, 0x1b,
	0x07, 0x59, 0x72, 0x0f, 0xd6, 0xfe, 0x4c, 0xdb, 0xaa, 0x63, 0x2f, 0xa0, 0x35, 0x5f, 0x8a, 0x39,
	0x0b, 0x21, 0xb9, 0xf9, 0x2b, 0x65, 0x2a, 0x33, 0x7d, 0xac, 0x9e, 0xef, 0xa1, 0xf1, 0x1d, 0xda,
	0x4e, 0x82, 0x28, 0x90, 0x99, 0x69, 0x6a, 0xe7, 0xb4, 0xf6, 0x37, 0x88, 0xf3, 0xc4, 0x37, 0x20,
	0x7f, 0x30, 0x2b, 0xf6, 0x26, 0x0a, 0x0d, 0x7c, 0x81, 0x9a, 0xa6, 0x07, 0x91, 0xa7, 0xa6, 0x81,
	0x7c, 0xa8, 0x27, 0x59, 0x30, 0xbe, 0x44, 0x2d, 0x77, 0x7e, 0xbc, 0xd6, 0xef, 0x01, 0x87, 0xaf,
	0x1a, 0xc4, 0xde, 0xe3, 0x0c, 0x62, 0xff, 0x71, 0x06, 0xf1, 0xd9, 0xfa, 0x06, 0x71, 0xb0, 0xbe,
	0x41, 0x9c, 0x14, 0x06, 0xf1, 0xb9, 0x6e, 0xe1, 0xf5, 0xc6, 0xae, 0xcd, 0xa1, 0xaf, 0xcd, 0x81,
	0xac, 0x80, 0xe6, 0xc6, 0x70, 0x52, 0x18, 0xc3, 0x17, 0xab, 0x54, 0xab, 0xa6, 0xd0, 0x2b, 0x4d,
	0xe1, 0xea, 0xfb, 0xdf, 0xbf, 0x9b, 0x30, 0x35, 0x4d, 0x87, 0x39, 0x75, 0x70, 0xf6, 0x3c, 0x51,
	0x4c, 0x0c, 0x26, 0xe2, 0x9b, 0x38, 0x0a, 0xb2, 0x89, 0x14, 0x29, 0x0f, 0x07, 0x13, 0x19, 0x8f,
	0x06, 0xc9, 0x68, 0x0a, 0xb3, 0xc0, 0xff, 0x83, 0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa8,
	0x92, 0x8b, 0x06, 0xfe, 0x0c, 0x00, 0x00,
}
