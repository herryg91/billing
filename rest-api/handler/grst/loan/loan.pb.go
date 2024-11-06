// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: loan.proto

package loan

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/herryg91/cdd/protoc-gen-cdd/ext/cddapis/cdd/api"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type LoanSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoanCount        int32   `protobuf:"varint,1,opt,name=LoanCount,json=loan_count,proto3" json:"loan_count,omitempty"`
	TotalOutstanding float64 `protobuf:"fixed64,2,opt,name=TotalOutstanding,json=total_outstanding,proto3" json:"total_outstanding,omitempty"`
	IsDelinquent     bool    `protobuf:"varint,3,opt,name=IsDelinquent,json=is_delinquent,proto3" json:"is_delinquent,omitempty"`
}

func (x *LoanSummary) Reset() {
	*x = LoanSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanSummary) ProtoMessage() {}

func (x *LoanSummary) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanSummary.ProtoReflect.Descriptor instead.
func (*LoanSummary) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{0}
}

func (x *LoanSummary) GetLoanCount() int32 {
	if x != nil {
		return x.LoanCount
	}
	return 0
}

func (x *LoanSummary) GetTotalOutstanding() float64 {
	if x != nil {
		return x.TotalOutstanding
	}
	return 0
}

func (x *LoanSummary) GetIsDelinquent() bool {
	if x != nil {
		return x.IsDelinquent
	}
	return false
}

type SimulateLoanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallmentLength int32   `protobuf:"varint,1,opt,name=InstallmentLength,json=installment_length,proto3" json:"installment_length,omitempty" validate:"required"`
	Principal         float64 `protobuf:"fixed64,2,opt,name=Principal,json=principal,proto3" json:"principal,omitempty" validate:"required"`
}

func (x *SimulateLoanReq) Reset() {
	*x = SimulateLoanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SimulateLoanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimulateLoanReq) ProtoMessage() {}

func (x *SimulateLoanReq) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimulateLoanReq.ProtoReflect.Descriptor instead.
func (*SimulateLoanReq) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{1}
}

func (x *SimulateLoanReq) GetInstallmentLength() int32 {
	if x != nil {
		return x.InstallmentLength
	}
	return 0
}

func (x *SimulateLoanReq) GetPrincipal() float64 {
	if x != nil {
		return x.Principal
	}
	return 0
}

type LoanSimulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallmentLength int32                    `protobuf:"varint,1,opt,name=InstallmentLength,json=installment_length,proto3" json:"installment_length,omitempty"`
	Principal         float64                  `protobuf:"fixed64,2,opt,name=Principal,json=principal,proto3" json:"principal,omitempty"`
	InterestAmount    float64                  `protobuf:"fixed64,3,opt,name=InterestAmount,json=interest_amount,proto3" json:"interest_amount,omitempty"`
	TotalAmount       float64                  `protobuf:"fixed64,4,opt,name=TotalAmount,json=total_amount,proto3" json:"total_amount,omitempty"`
	Billings          []*LoanBillingSimulation `protobuf:"bytes,5,rep,name=Billings,json=billings,proto3" json:"billings,omitempty"`
}

func (x *LoanSimulation) Reset() {
	*x = LoanSimulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanSimulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanSimulation) ProtoMessage() {}

func (x *LoanSimulation) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanSimulation.ProtoReflect.Descriptor instead.
func (*LoanSimulation) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{2}
}

func (x *LoanSimulation) GetInstallmentLength() int32 {
	if x != nil {
		return x.InstallmentLength
	}
	return 0
}

func (x *LoanSimulation) GetPrincipal() float64 {
	if x != nil {
		return x.Principal
	}
	return 0
}

func (x *LoanSimulation) GetInterestAmount() float64 {
	if x != nil {
		return x.InterestAmount
	}
	return 0
}

func (x *LoanSimulation) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *LoanSimulation) GetBillings() []*LoanBillingSimulation {
	if x != nil {
		return x.Billings
	}
	return nil
}

type LoanBillingSimulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstallmentNumber int32   `protobuf:"varint,1,opt,name=InstallmentNumber,json=installment_number,proto3" json:"installment_number,omitempty"`
	Principal         float64 `protobuf:"fixed64,2,opt,name=Principal,json=principal,proto3" json:"principal,omitempty"`
	InterestAmount    float64 `protobuf:"fixed64,3,opt,name=InterestAmount,json=interest_amount,proto3" json:"interest_amount,omitempty"`
	TotalAmount       float64 `protobuf:"fixed64,4,opt,name=TotalAmount,json=total_amount,proto3" json:"total_amount,omitempty"`
}

func (x *LoanBillingSimulation) Reset() {
	*x = LoanBillingSimulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoanBillingSimulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoanBillingSimulation) ProtoMessage() {}

func (x *LoanBillingSimulation) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoanBillingSimulation.ProtoReflect.Descriptor instead.
func (*LoanBillingSimulation) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{3}
}

func (x *LoanBillingSimulation) GetInstallmentNumber() int32 {
	if x != nil {
		return x.InstallmentNumber
	}
	return 0
}

func (x *LoanBillingSimulation) GetPrincipal() float64 {
	if x != nil {
		return x.Principal
	}
	return 0
}

func (x *LoanBillingSimulation) GetInterestAmount() float64 {
	if x != nil {
		return x.InterestAmount
	}
	return 0
}

func (x *LoanBillingSimulation) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

type CreateLoanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description       string  `protobuf:"bytes,1,opt,name=Description,json=description,proto3" json:"description,omitempty" validate:"required"`
	InstallmentLength int32   `protobuf:"varint,2,opt,name=InstallmentLength,json=installment_length,proto3" json:"installment_length,omitempty" validate:"required"`
	Principal         float64 `protobuf:"fixed64,3,opt,name=Principal,json=principal,proto3" json:"principal,omitempty" validate:"required"`
}

func (x *CreateLoanReq) Reset() {
	*x = CreateLoanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLoanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLoanReq) ProtoMessage() {}

func (x *CreateLoanReq) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLoanReq.ProtoReflect.Descriptor instead.
func (*CreateLoanReq) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{4}
}

func (x *CreateLoanReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateLoanReq) GetInstallmentLength() int32 {
	if x != nil {
		return x.InstallmentLength
	}
	return 0
}

func (x *CreateLoanReq) GetPrincipal() float64 {
	if x != nil {
		return x.Principal
	}
	return 0
}

type Loans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Loans []*Loan `protobuf:"bytes,1,rep,name=Loans,json=loans,proto3" json:"loans,omitempty"`
}

func (x *Loans) Reset() {
	*x = Loans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loans) ProtoMessage() {}

func (x *Loans) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loans.ProtoReflect.Descriptor instead.
func (*Loans) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{5}
}

func (x *Loans) GetLoans() []*Loan {
	if x != nil {
		return x.Loans
	}
	return nil
}

type Loan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code              string               `protobuf:"bytes,1,opt,name=Code,json=code,proto3" json:"code,omitempty"`
	Description       string               `protobuf:"bytes,2,opt,name=Description,json=description,proto3" json:"description,omitempty"`
	InstallmentCycle  string               `protobuf:"bytes,3,opt,name=InstallmentCycle,json=installment_cycle,proto3" json:"installment_cycle,omitempty"`
	InstallmentLength int32                `protobuf:"varint,4,opt,name=InstallmentLength,json=installment_length,proto3" json:"installment_length,omitempty"`
	InterestType      string               `protobuf:"bytes,5,opt,name=InterestType,json=interest_type,proto3" json:"interest_type,omitempty"`
	InterestPercent   float64              `protobuf:"fixed64,6,opt,name=InterestPercent,json=interest_percent,proto3" json:"interest_percent,omitempty"`
	Principal         float64              `protobuf:"fixed64,7,opt,name=Principal,json=principal,proto3" json:"principal,omitempty"`
	InterestAmount    float64              `protobuf:"fixed64,8,opt,name=InterestAmount,json=interest_amount,proto3" json:"interest_amount,omitempty"`
	TotalAmount       float64              `protobuf:"fixed64,9,opt,name=TotalAmount,json=total_amount,proto3" json:"total_amount,omitempty"`
	Outstanding       float64              `protobuf:"fixed64,10,opt,name=Outstanding,json=outstanding,proto3" json:"outstanding,omitempty"`
	Status            string               `protobuf:"bytes,11,opt,name=Status,json=status,proto3" json:"status,omitempty" validate:"enum=PENDING|APPROVED|ACTIVE|DONE"`
	CreatedAt         *timestamp.Timestamp `protobuf:"bytes,13,opt,name=CreatedAt,json=created_at,proto3" json:"created_at,omitempty"`
}

func (x *Loan) Reset() {
	*x = Loan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Loan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Loan) ProtoMessage() {}

func (x *Loan) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Loan.ProtoReflect.Descriptor instead.
func (*Loan) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{6}
}

func (x *Loan) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Loan) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Loan) GetInstallmentCycle() string {
	if x != nil {
		return x.InstallmentCycle
	}
	return ""
}

func (x *Loan) GetInstallmentLength() int32 {
	if x != nil {
		return x.InstallmentLength
	}
	return 0
}

func (x *Loan) GetInterestType() string {
	if x != nil {
		return x.InterestType
	}
	return ""
}

func (x *Loan) GetInterestPercent() float64 {
	if x != nil {
		return x.InterestPercent
	}
	return 0
}

func (x *Loan) GetPrincipal() float64 {
	if x != nil {
		return x.Principal
	}
	return 0
}

func (x *Loan) GetInterestAmount() float64 {
	if x != nil {
		return x.InterestAmount
	}
	return 0
}

func (x *Loan) GetTotalAmount() float64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *Loan) GetOutstanding() float64 {
	if x != nil {
		return x.Outstanding
	}
	return 0
}

func (x *Loan) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Loan) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetLoanByCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=Code,json=code,proto3" json:"code,omitempty" validate:"required"`
}

func (x *GetLoanByCodeReq) Reset() {
	*x = GetLoanByCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loan_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoanByCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoanByCodeReq) ProtoMessage() {}

func (x *GetLoanByCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_loan_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoanByCodeReq.ProtoReflect.Descriptor instead.
func (*GetLoanByCodeReq) Descriptor() ([]byte, []int) {
	return file_loan_proto_rawDescGZIP(), []int{7}
}

func (x *GetLoanByCodeReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_loan_proto protoreflect.FileDescriptor

var file_loan_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f,
	0x61, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x63, 0x64, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x64, 0x64, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x0b, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x09, 0x4c, 0x6f, 0x61, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x61, 0x6e, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x10, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4f, 0x75, 0x74, 0x73, 0x74,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x11, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x23, 0x0a, 0x0c, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x69, 0x6e, 0x71, 0x75, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x6e, 0x71,
	0x75, 0x65, 0x6e, 0x74, 0x22, 0x7a, 0x0a, 0x0f, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x3b, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0c, 0xc2, 0x8a, 0x3b, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0c, 0xc2, 0x8a, 0x3b, 0x08, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x22, 0xe2, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c,
	0x12, 0x27, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65,
	0x73, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x0a, 0x08,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x15, 0x4c, 0x6f, 0x61, 0x6e, 0x42, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xc2, 0x8a, 0x3b, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x11, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0c, 0xc2, 0x8a, 0x3b, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x2a, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x42, 0x0c, 0xc2, 0x8a, 0x3b, 0x08,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x61, 0x6c, 0x22, 0x29, 0x0a, 0x05, 0x4c, 0x6f, 0x61, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x05,
	0x4c, 0x6f, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x6f,
	0x61, 0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x05, 0x6c, 0x6f, 0x61, 0x6e, 0x73, 0x22, 0xee,
	0x03, 0x0a, 0x04, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a,
	0x10, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x79, 0x63, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x79, 0x63, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x11, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6c, 0x6c, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x23, 0x0a, 0x0c, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x29,
	0x0a, 0x0f, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73,
	0x74, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x69,
	0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x70, 0x72,
	0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x65, 0x73, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x21, 0x0a, 0x0b, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x75, 0x74, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x73, 0x74, 0x61,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25, 0xc2, 0x8a, 0x3b, 0x21, 0x65, 0x6e, 0x75, 0x6d, 0x3d,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x2c, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44,
	0x2c, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x2c, 0x44, 0x4f, 0x4e, 0x45, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22,
	0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xc2, 0x8a, 0x3b, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xc7, 0x03, 0x0a, 0x07, 0x4c, 0x6f, 0x61, 0x6e, 0x41, 0x70,
	0x69, 0x12, 0x60, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72,
	0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x6c, 0x6f, 0x61, 0x6e,
	0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x25, 0x82, 0xbd,
	0x3f, 0x09, 0x88, 0xbd, 0x3f, 0x01, 0x9a, 0xbd, 0x3f, 0x01, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x12, 0x12, 0x10, 0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x2f, 0x6d, 0x79, 0x2d, 0x73, 0x75, 0x6d, 0x6d,
	0x61, 0x72, 0x79, 0x12, 0x56, 0x0a, 0x0c, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x61, 0x6e, 0x12, 0x15, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x6c, 0x6f, 0x61,
	0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x6c, 0x6f,
	0x61, 0x6e, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x58, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x61, 0x6e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xbd, 0x3f, 0x09, 0x88, 0xbd, 0x3f, 0x01,
	0x9a, 0xbd, 0x3f, 0x01, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05,
	0x2f, 0x6c, 0x6f, 0x61, 0x6e, 0x12, 0x4b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x61, 0x6e,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0b, 0x2e, 0x6c, 0x6f, 0x61, 0x6e,
	0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x73, 0x22, 0x1a, 0x82, 0xbd, 0x3f, 0x09, 0x88, 0xbd, 0x3f, 0x01,
	0x9a, 0xbd, 0x3f, 0x01, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x6c, 0x6f,
	0x61, 0x6e, 0x12, 0x5b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x61, 0x6e, 0x42, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x6c, 0x6f, 0x61, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x61, 0x6e, 0x42, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x6c, 0x6f,
	0x61, 0x6e, 0x2e, 0x4c, 0x6f, 0x61, 0x6e, 0x22, 0x26, 0x82, 0xbd, 0x3f, 0x09, 0x88, 0xbd, 0x3f,
	0x01, 0x9a, 0xbd, 0x3f, 0x01, 0x2a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x6c,
	0x6f, 0x61, 0x6e, 0x2f, 0x7b, 0x43, 0x6f, 0x64, 0x65, 0x7d, 0x2f, 0x76, 0x69, 0x65, 0x77, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6c, 0x6f, 0x61, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_loan_proto_rawDescOnce sync.Once
	file_loan_proto_rawDescData = file_loan_proto_rawDesc
)

func file_loan_proto_rawDescGZIP() []byte {
	file_loan_proto_rawDescOnce.Do(func() {
		file_loan_proto_rawDescData = protoimpl.X.CompressGZIP(file_loan_proto_rawDescData)
	})
	return file_loan_proto_rawDescData
}

var file_loan_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_loan_proto_goTypes = []interface{}{
	(*LoanSummary)(nil),           // 0: loan.LoanSummary
	(*SimulateLoanReq)(nil),       // 1: loan.SimulateLoanReq
	(*LoanSimulation)(nil),        // 2: loan.LoanSimulation
	(*LoanBillingSimulation)(nil), // 3: loan.LoanBillingSimulation
	(*CreateLoanReq)(nil),         // 4: loan.CreateLoanReq
	(*Loans)(nil),                 // 5: loan.Loans
	(*Loan)(nil),                  // 6: loan.Loan
	(*GetLoanByCodeReq)(nil),      // 7: loan.GetLoanByCodeReq
	(*timestamp.Timestamp)(nil),   // 8: google.protobuf.Timestamp
	(*empty.Empty)(nil),           // 9: google.protobuf.Empty
}
var file_loan_proto_depIdxs = []int32{
	3, // 0: loan.LoanSimulation.Billings:type_name -> loan.LoanBillingSimulation
	6, // 1: loan.Loans.Loans:type_name -> loan.Loan
	8, // 2: loan.Loan.CreatedAt:type_name -> google.protobuf.Timestamp
	9, // 3: loan.LoanApi.GetMySummary:input_type -> google.protobuf.Empty
	1, // 4: loan.LoanApi.SimulateLoan:input_type -> loan.SimulateLoanReq
	4, // 5: loan.LoanApi.CreateLoan:input_type -> loan.CreateLoanReq
	9, // 6: loan.LoanApi.GetLoans:input_type -> google.protobuf.Empty
	7, // 7: loan.LoanApi.GetLoanByCode:input_type -> loan.GetLoanByCodeReq
	0, // 8: loan.LoanApi.GetMySummary:output_type -> loan.LoanSummary
	2, // 9: loan.LoanApi.SimulateLoan:output_type -> loan.LoanSimulation
	9, // 10: loan.LoanApi.CreateLoan:output_type -> google.protobuf.Empty
	5, // 11: loan.LoanApi.GetLoans:output_type -> loan.Loans
	6, // 12: loan.LoanApi.GetLoanByCode:output_type -> loan.Loan
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_loan_proto_init() }
func file_loan_proto_init() {
	if File_loan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanSummary); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SimulateLoanReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanSimulation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loan_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoanBillingSimulation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loan_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLoanReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loan_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loans); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loan_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Loan); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_loan_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoanByCodeReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_loan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loan_proto_goTypes,
		DependencyIndexes: file_loan_proto_depIdxs,
		MessageInfos:      file_loan_proto_msgTypes,
	}.Build()
	File_loan_proto = out.File
	file_loan_proto_rawDesc = nil
	file_loan_proto_goTypes = nil
	file_loan_proto_depIdxs = nil
}