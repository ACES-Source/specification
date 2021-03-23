// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assets.proto

package assets

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Message - Membership (MBR)
type Membership struct {
	AgeRestriction       *AgeRestrictionField `protobuf:"bytes,1,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	ValidFrom            uint64               `protobuf:"varint,2,opt,name=ValidFrom,proto3" json:"ValidFrom,omitempty"`
	ExpirationTimestamp  uint64               `protobuf:"varint,3,opt,name=ExpirationTimestamp,proto3" json:"ExpirationTimestamp,omitempty"`
	ID                   string               `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty"`
	MembershipClass      string               `protobuf:"bytes,5,opt,name=MembershipClass,proto3" json:"MembershipClass,omitempty"`
	RoleType             string               `protobuf:"bytes,6,opt,name=RoleType,proto3" json:"RoleType,omitempty"`
	MembershipType       string               `protobuf:"bytes,7,opt,name=MembershipType,proto3" json:"MembershipType,omitempty"`
	Description          string               `protobuf:"bytes,8,opt,name=Description,proto3" json:"Description,omitempty"`
	TransfersPermitted   bool                 `protobuf:"varint,9,opt,name=TransfersPermitted,proto3" json:"TransfersPermitted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Membership) Reset()         { *m = Membership{} }
func (m *Membership) String() string { return proto.CompactTextString(m) }
func (*Membership) ProtoMessage()    {}
func (*Membership) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{0}
}

func (m *Membership) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Membership.Unmarshal(m, b)
}
func (m *Membership) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Membership.Marshal(b, m, deterministic)
}
func (m *Membership) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Membership.Merge(m, src)
}
func (m *Membership) XXX_Size() int {
	return xxx_messageInfo_Membership.Size(m)
}
func (m *Membership) XXX_DiscardUnknown() {
	xxx_messageInfo_Membership.DiscardUnknown(m)
}

var xxx_messageInfo_Membership proto.InternalMessageInfo

func (m *Membership) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *Membership) GetValidFrom() uint64 {
	if m != nil {
		return m.ValidFrom
	}
	return 0
}

func (m *Membership) GetExpirationTimestamp() uint64 {
	if m != nil {
		return m.ExpirationTimestamp
	}
	return 0
}

func (m *Membership) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Membership) GetMembershipClass() string {
	if m != nil {
		return m.MembershipClass
	}
	return ""
}

func (m *Membership) GetRoleType() string {
	if m != nil {
		return m.RoleType
	}
	return ""
}

func (m *Membership) GetMembershipType() string {
	if m != nil {
		return m.MembershipType
	}
	return ""
}

func (m *Membership) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Membership) GetTransfersPermitted() bool {
	if m != nil {
		return m.TransfersPermitted
	}
	return false
}

// Message - Currency (CCY)
type Currency struct {
	CurrencyCode         string   `protobuf:"bytes,1,opt,name=CurrencyCode,proto3" json:"CurrencyCode,omitempty"`
	MonetaryAuthority    string   `protobuf:"bytes,2,opt,name=MonetaryAuthority,proto3" json:"MonetaryAuthority,omitempty"`
	Precision            uint64   `protobuf:"varint,4,opt,name=Precision,proto3" json:"Precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{1}
}

func (m *Currency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Currency.Unmarshal(m, b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
}
func (m *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(m, src)
}
func (m *Currency) XXX_Size() int {
	return xxx_messageInfo_Currency.Size(m)
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *Currency) GetMonetaryAuthority() string {
	if m != nil {
		return m.MonetaryAuthority
	}
	return ""
}

func (m *Currency) GetPrecision() uint64 {
	if m != nil {
		return m.Precision
	}
	return 0
}

// Message - Share - Common (SHC)
type ShareCommon struct {
	Ticker               string   `protobuf:"bytes,1,opt,name=Ticker,proto3" json:"Ticker,omitempty"`
	ISIN                 string   `protobuf:"bytes,2,opt,name=ISIN,proto3" json:"ISIN,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	TransfersPermitted   bool     `protobuf:"varint,4,opt,name=TransfersPermitted,proto3" json:"TransfersPermitted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShareCommon) Reset()         { *m = ShareCommon{} }
func (m *ShareCommon) String() string { return proto.CompactTextString(m) }
func (*ShareCommon) ProtoMessage()    {}
func (*ShareCommon) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{2}
}

func (m *ShareCommon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShareCommon.Unmarshal(m, b)
}
func (m *ShareCommon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShareCommon.Marshal(b, m, deterministic)
}
func (m *ShareCommon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShareCommon.Merge(m, src)
}
func (m *ShareCommon) XXX_Size() int {
	return xxx_messageInfo_ShareCommon.Size(m)
}
func (m *ShareCommon) XXX_DiscardUnknown() {
	xxx_messageInfo_ShareCommon.DiscardUnknown(m)
}

var xxx_messageInfo_ShareCommon proto.InternalMessageInfo

func (m *ShareCommon) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *ShareCommon) GetISIN() string {
	if m != nil {
		return m.ISIN
	}
	return ""
}

func (m *ShareCommon) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ShareCommon) GetTransfersPermitted() bool {
	if m != nil {
		return m.TransfersPermitted
	}
	return false
}

// Message - Bond - Fixed Rate (BFR)
type BondFixedRate struct {
	Name                       string               `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	BondType                   string               `protobuf:"bytes,2,opt,name=BondType,proto3" json:"BondType,omitempty"`
	ISIN                       string               `protobuf:"bytes,3,opt,name=ISIN,proto3" json:"ISIN,omitempty"`
	Collateral                 string               `protobuf:"bytes,4,opt,name=Collateral,proto3" json:"Collateral,omitempty"`
	ParValue                   *CurrencyValueField  `protobuf:"bytes,5,opt,name=ParValue,proto3" json:"ParValue,omitempty"`
	InterestRate               *RateField           `protobuf:"bytes,6,opt,name=InterestRate,proto3" json:"InterestRate,omitempty"`
	InterestPaymentInitialDate uint64               `protobuf:"varint,7,opt,name=InterestPaymentInitialDate,proto3" json:"InterestPaymentInitialDate,omitempty"`
	InterestPaymentDateDeltas  []uint64             `protobuf:"varint,8,rep,packed,name=InterestPaymentDateDeltas,proto3" json:"InterestPaymentDateDeltas,omitempty"`
	LatePaymentPenaltyRate     *RateField           `protobuf:"bytes,9,opt,name=LatePaymentPenaltyRate,proto3" json:"LatePaymentPenaltyRate,omitempty"`
	LatePaymentWindow          uint64               `protobuf:"varint,10,opt,name=LatePaymentWindow,proto3" json:"LatePaymentWindow,omitempty"`
	LatePaymentPenaltyPeriod   uint64               `protobuf:"varint,11,opt,name=LatePaymentPenaltyPeriod,proto3" json:"LatePaymentPenaltyPeriod,omitempty"`
	MaturityDate               uint64               `protobuf:"varint,12,opt,name=MaturityDate,proto3" json:"MaturityDate,omitempty"`
	AgeRestriction             *AgeRestrictionField `protobuf:"bytes,13,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	TransfersPermitted         bool                 `protobuf:"varint,14,opt,name=TransfersPermitted,proto3" json:"TransfersPermitted,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}             `json:"-"`
	XXX_unrecognized           []byte               `json:"-"`
	XXX_sizecache              int32                `json:"-"`
}

func (m *BondFixedRate) Reset()         { *m = BondFixedRate{} }
func (m *BondFixedRate) String() string { return proto.CompactTextString(m) }
func (*BondFixedRate) ProtoMessage()    {}
func (*BondFixedRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{3}
}

func (m *BondFixedRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BondFixedRate.Unmarshal(m, b)
}
func (m *BondFixedRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BondFixedRate.Marshal(b, m, deterministic)
}
func (m *BondFixedRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BondFixedRate.Merge(m, src)
}
func (m *BondFixedRate) XXX_Size() int {
	return xxx_messageInfo_BondFixedRate.Size(m)
}
func (m *BondFixedRate) XXX_DiscardUnknown() {
	xxx_messageInfo_BondFixedRate.DiscardUnknown(m)
}

var xxx_messageInfo_BondFixedRate proto.InternalMessageInfo

func (m *BondFixedRate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BondFixedRate) GetBondType() string {
	if m != nil {
		return m.BondType
	}
	return ""
}

func (m *BondFixedRate) GetISIN() string {
	if m != nil {
		return m.ISIN
	}
	return ""
}

func (m *BondFixedRate) GetCollateral() string {
	if m != nil {
		return m.Collateral
	}
	return ""
}

func (m *BondFixedRate) GetParValue() *CurrencyValueField {
	if m != nil {
		return m.ParValue
	}
	return nil
}

func (m *BondFixedRate) GetInterestRate() *RateField {
	if m != nil {
		return m.InterestRate
	}
	return nil
}

func (m *BondFixedRate) GetInterestPaymentInitialDate() uint64 {
	if m != nil {
		return m.InterestPaymentInitialDate
	}
	return 0
}

func (m *BondFixedRate) GetInterestPaymentDateDeltas() []uint64 {
	if m != nil {
		return m.InterestPaymentDateDeltas
	}
	return nil
}

func (m *BondFixedRate) GetLatePaymentPenaltyRate() *RateField {
	if m != nil {
		return m.LatePaymentPenaltyRate
	}
	return nil
}

func (m *BondFixedRate) GetLatePaymentWindow() uint64 {
	if m != nil {
		return m.LatePaymentWindow
	}
	return 0
}

func (m *BondFixedRate) GetLatePaymentPenaltyPeriod() uint64 {
	if m != nil {
		return m.LatePaymentPenaltyPeriod
	}
	return 0
}

func (m *BondFixedRate) GetMaturityDate() uint64 {
	if m != nil {
		return m.MaturityDate
	}
	return 0
}

func (m *BondFixedRate) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *BondFixedRate) GetTransfersPermitted() bool {
	if m != nil {
		return m.TransfersPermitted
	}
	return false
}

// Message - Coupon (COU)
type Coupon struct {
	RedeemingEntity      string              `protobuf:"bytes,1,opt,name=RedeemingEntity,proto3" json:"RedeemingEntity,omitempty"`
	ValidFromTimestamp   uint64              `protobuf:"varint,2,opt,name=ValidFromTimestamp,proto3" json:"ValidFromTimestamp,omitempty"`
	ExpiryTimestamp      uint64              `protobuf:"varint,3,opt,name=ExpiryTimestamp,proto3" json:"ExpiryTimestamp,omitempty"`
	CouponName           string              `protobuf:"bytes,6,opt,name=CouponName,proto3" json:"CouponName,omitempty"`
	TransfersPermitted   bool                `protobuf:"varint,8,opt,name=TransfersPermitted,proto3" json:"TransfersPermitted,omitempty"`
	FaceValue            *CurrencyValueField `protobuf:"bytes,9,opt,name=FaceValue,proto3" json:"FaceValue,omitempty"`
	RedemptionVenue      string              `protobuf:"bytes,10,opt,name=RedemptionVenue,proto3" json:"RedemptionVenue,omitempty"`
	Details              string              `protobuf:"bytes,11,opt,name=Details,proto3" json:"Details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Coupon) Reset()         { *m = Coupon{} }
func (m *Coupon) String() string { return proto.CompactTextString(m) }
func (*Coupon) ProtoMessage()    {}
func (*Coupon) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{4}
}

func (m *Coupon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coupon.Unmarshal(m, b)
}
func (m *Coupon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coupon.Marshal(b, m, deterministic)
}
func (m *Coupon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coupon.Merge(m, src)
}
func (m *Coupon) XXX_Size() int {
	return xxx_messageInfo_Coupon.Size(m)
}
func (m *Coupon) XXX_DiscardUnknown() {
	xxx_messageInfo_Coupon.DiscardUnknown(m)
}

var xxx_messageInfo_Coupon proto.InternalMessageInfo

func (m *Coupon) GetRedeemingEntity() string {
	if m != nil {
		return m.RedeemingEntity
	}
	return ""
}

func (m *Coupon) GetValidFromTimestamp() uint64 {
	if m != nil {
		return m.ValidFromTimestamp
	}
	return 0
}

func (m *Coupon) GetExpiryTimestamp() uint64 {
	if m != nil {
		return m.ExpiryTimestamp
	}
	return 0
}

func (m *Coupon) GetCouponName() string {
	if m != nil {
		return m.CouponName
	}
	return ""
}

func (m *Coupon) GetTransfersPermitted() bool {
	if m != nil {
		return m.TransfersPermitted
	}
	return false
}

func (m *Coupon) GetFaceValue() *CurrencyValueField {
	if m != nil {
		return m.FaceValue
	}
	return nil
}

func (m *Coupon) GetRedemptionVenue() string {
	if m != nil {
		return m.RedemptionVenue
	}
	return ""
}

func (m *Coupon) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

// Message - Loyalty Points (LOY)
type LoyaltyPoints struct {
	AgeRestriction       *AgeRestrictionField `protobuf:"bytes,1,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	ProgramName          string               `protobuf:"bytes,2,opt,name=ProgramName,proto3" json:"ProgramName,omitempty"`
	ExpirationTimestamp  uint64               `protobuf:"varint,4,opt,name=ExpirationTimestamp,proto3" json:"ExpirationTimestamp,omitempty"`
	Details              string               `protobuf:"bytes,5,opt,name=Details,proto3" json:"Details,omitempty"`
	TransfersPermitted   bool                 `protobuf:"varint,6,opt,name=TransfersPermitted,proto3" json:"TransfersPermitted,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *LoyaltyPoints) Reset()         { *m = LoyaltyPoints{} }
func (m *LoyaltyPoints) String() string { return proto.CompactTextString(m) }
func (*LoyaltyPoints) ProtoMessage()    {}
func (*LoyaltyPoints) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{5}
}

func (m *LoyaltyPoints) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoyaltyPoints.Unmarshal(m, b)
}
func (m *LoyaltyPoints) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoyaltyPoints.Marshal(b, m, deterministic)
}
func (m *LoyaltyPoints) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoyaltyPoints.Merge(m, src)
}
func (m *LoyaltyPoints) XXX_Size() int {
	return xxx_messageInfo_LoyaltyPoints.Size(m)
}
func (m *LoyaltyPoints) XXX_DiscardUnknown() {
	xxx_messageInfo_LoyaltyPoints.DiscardUnknown(m)
}

var xxx_messageInfo_LoyaltyPoints proto.InternalMessageInfo

func (m *LoyaltyPoints) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *LoyaltyPoints) GetProgramName() string {
	if m != nil {
		return m.ProgramName
	}
	return ""
}

func (m *LoyaltyPoints) GetExpirationTimestamp() uint64 {
	if m != nil {
		return m.ExpirationTimestamp
	}
	return 0
}

func (m *LoyaltyPoints) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *LoyaltyPoints) GetTransfersPermitted() bool {
	if m != nil {
		return m.TransfersPermitted
	}
	return false
}

// Message - Ticket (Admission) (TIC)
type TicketAdmission struct {
	AgeRestriction       *AgeRestrictionField `protobuf:"bytes,1,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	Venue                string               `protobuf:"bytes,3,opt,name=Venue,proto3" json:"Venue,omitempty"`
	Area                 string               `protobuf:"bytes,5,opt,name=Area,proto3" json:"Area,omitempty"`
	Seat                 string               `protobuf:"bytes,6,opt,name=Seat,proto3" json:"Seat,omitempty"`
	EventTimestamp       uint64               `protobuf:"varint,7,opt,name=EventTimestamp,proto3" json:"EventTimestamp,omitempty"`
	EventName            string               `protobuf:"bytes,10,opt,name=EventName,proto3" json:"EventName,omitempty"`
	TransfersPermitted   bool                 `protobuf:"varint,11,opt,name=TransfersPermitted,proto3" json:"TransfersPermitted,omitempty"`
	Details              string               `protobuf:"bytes,12,opt,name=Details,proto3" json:"Details,omitempty"`
	Section              string               `protobuf:"bytes,13,opt,name=Section,proto3" json:"Section,omitempty"`
	Row                  string               `protobuf:"bytes,14,opt,name=Row,proto3" json:"Row,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TicketAdmission) Reset()         { *m = TicketAdmission{} }
func (m *TicketAdmission) String() string { return proto.CompactTextString(m) }
func (*TicketAdmission) ProtoMessage()    {}
func (*TicketAdmission) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{6}
}

func (m *TicketAdmission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TicketAdmission.Unmarshal(m, b)
}
func (m *TicketAdmission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TicketAdmission.Marshal(b, m, deterministic)
}
func (m *TicketAdmission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TicketAdmission.Merge(m, src)
}
func (m *TicketAdmission) XXX_Size() int {
	return xxx_messageInfo_TicketAdmission.Size(m)
}
func (m *TicketAdmission) XXX_DiscardUnknown() {
	xxx_messageInfo_TicketAdmission.DiscardUnknown(m)
}

var xxx_messageInfo_TicketAdmission proto.InternalMessageInfo

func (m *TicketAdmission) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *TicketAdmission) GetVenue() string {
	if m != nil {
		return m.Venue
	}
	return ""
}

func (m *TicketAdmission) GetArea() string {
	if m != nil {
		return m.Area
	}
	return ""
}

func (m *TicketAdmission) GetSeat() string {
	if m != nil {
		return m.Seat
	}
	return ""
}

func (m *TicketAdmission) GetEventTimestamp() uint64 {
	if m != nil {
		return m.EventTimestamp
	}
	return 0
}

func (m *TicketAdmission) GetEventName() string {
	if m != nil {
		return m.EventName
	}
	return ""
}

func (m *TicketAdmission) GetTransfersPermitted() bool {
	if m != nil {
		return m.TransfersPermitted
	}
	return false
}

func (m *TicketAdmission) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *TicketAdmission) GetSection() string {
	if m != nil {
		return m.Section
	}
	return ""
}

func (m *TicketAdmission) GetRow() string {
	if m != nil {
		return m.Row
	}
	return ""
}

// Message - Casino Chip (CHP)
type CasinoChip struct {
	CurrencyCode         string               `protobuf:"bytes,1,opt,name=CurrencyCode,proto3" json:"CurrencyCode,omitempty"`
	UseType              string               `protobuf:"bytes,2,opt,name=UseType,proto3" json:"UseType,omitempty"`
	AgeRestriction       *AgeRestrictionField `protobuf:"bytes,3,opt,name=AgeRestriction,proto3" json:"AgeRestriction,omitempty"`
	ExpirationTimestamp  uint64               `protobuf:"varint,5,opt,name=ExpirationTimestamp,proto3" json:"ExpirationTimestamp,omitempty"`
	Precision            uint64               `protobuf:"varint,6,opt,name=Precision,proto3" json:"Precision,omitempty"`
	TransfersPermitted   bool                 `protobuf:"varint,7,opt,name=TransfersPermitted,proto3" json:"TransfersPermitted,omitempty"`
	CasinoName           string               `protobuf:"bytes,8,opt,name=CasinoName,proto3" json:"CasinoName,omitempty"`
	FaceValue            *CurrencyValueField  `protobuf:"bytes,9,opt,name=FaceValue,proto3" json:"FaceValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CasinoChip) Reset()         { *m = CasinoChip{} }
func (m *CasinoChip) String() string { return proto.CompactTextString(m) }
func (*CasinoChip) ProtoMessage()    {}
func (*CasinoChip) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{7}
}

func (m *CasinoChip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CasinoChip.Unmarshal(m, b)
}
func (m *CasinoChip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CasinoChip.Marshal(b, m, deterministic)
}
func (m *CasinoChip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CasinoChip.Merge(m, src)
}
func (m *CasinoChip) XXX_Size() int {
	return xxx_messageInfo_CasinoChip.Size(m)
}
func (m *CasinoChip) XXX_DiscardUnknown() {
	xxx_messageInfo_CasinoChip.DiscardUnknown(m)
}

var xxx_messageInfo_CasinoChip proto.InternalMessageInfo

func (m *CasinoChip) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *CasinoChip) GetUseType() string {
	if m != nil {
		return m.UseType
	}
	return ""
}

func (m *CasinoChip) GetAgeRestriction() *AgeRestrictionField {
	if m != nil {
		return m.AgeRestriction
	}
	return nil
}

func (m *CasinoChip) GetExpirationTimestamp() uint64 {
	if m != nil {
		return m.ExpirationTimestamp
	}
	return 0
}

func (m *CasinoChip) GetPrecision() uint64 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *CasinoChip) GetTransfersPermitted() bool {
	if m != nil {
		return m.TransfersPermitted
	}
	return false
}

func (m *CasinoChip) GetCasinoName() string {
	if m != nil {
		return m.CasinoName
	}
	return ""
}

func (m *CasinoChip) GetFaceValue() *CurrencyValueField {
	if m != nil {
		return m.FaceValue
	}
	return nil
}

// Field - Age Restriction
type AgeRestrictionField struct {
	Lower                uint32   `protobuf:"varint,1,opt,name=Lower,proto3" json:"Lower,omitempty"`
	Upper                uint32   `protobuf:"varint,2,opt,name=Upper,proto3" json:"Upper,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgeRestrictionField) Reset()         { *m = AgeRestrictionField{} }
func (m *AgeRestrictionField) String() string { return proto.CompactTextString(m) }
func (*AgeRestrictionField) ProtoMessage()    {}
func (*AgeRestrictionField) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{8}
}

func (m *AgeRestrictionField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgeRestrictionField.Unmarshal(m, b)
}
func (m *AgeRestrictionField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgeRestrictionField.Marshal(b, m, deterministic)
}
func (m *AgeRestrictionField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgeRestrictionField.Merge(m, src)
}
func (m *AgeRestrictionField) XXX_Size() int {
	return xxx_messageInfo_AgeRestrictionField.Size(m)
}
func (m *AgeRestrictionField) XXX_DiscardUnknown() {
	xxx_messageInfo_AgeRestrictionField.DiscardUnknown(m)
}

var xxx_messageInfo_AgeRestrictionField proto.InternalMessageInfo

func (m *AgeRestrictionField) GetLower() uint32 {
	if m != nil {
		return m.Lower
	}
	return 0
}

func (m *AgeRestrictionField) GetUpper() uint32 {
	if m != nil {
		return m.Upper
	}
	return 0
}

// Field - Currency Value
type CurrencyValueField struct {
	Value                uint64   `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	CurrencyCode         string   `protobuf:"bytes,2,opt,name=CurrencyCode,proto3" json:"CurrencyCode,omitempty"`
	Precision            uint32   `protobuf:"varint,3,opt,name=Precision,proto3" json:"Precision,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CurrencyValueField) Reset()         { *m = CurrencyValueField{} }
func (m *CurrencyValueField) String() string { return proto.CompactTextString(m) }
func (*CurrencyValueField) ProtoMessage()    {}
func (*CurrencyValueField) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{9}
}

func (m *CurrencyValueField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CurrencyValueField.Unmarshal(m, b)
}
func (m *CurrencyValueField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CurrencyValueField.Marshal(b, m, deterministic)
}
func (m *CurrencyValueField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CurrencyValueField.Merge(m, src)
}
func (m *CurrencyValueField) XXX_Size() int {
	return xxx_messageInfo_CurrencyValueField.Size(m)
}
func (m *CurrencyValueField) XXX_DiscardUnknown() {
	xxx_messageInfo_CurrencyValueField.DiscardUnknown(m)
}

var xxx_messageInfo_CurrencyValueField proto.InternalMessageInfo

func (m *CurrencyValueField) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CurrencyValueField) GetCurrencyCode() string {
	if m != nil {
		return m.CurrencyCode
	}
	return ""
}

func (m *CurrencyValueField) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

// Field - Rate
type RateField struct {
	Precision            uint32   `protobuf:"varint,1,opt,name=Precision,proto3" json:"Precision,omitempty"`
	Value                uint64   `protobuf:"varint,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateField) Reset()         { *m = RateField{} }
func (m *RateField) String() string { return proto.CompactTextString(m) }
func (*RateField) ProtoMessage()    {}
func (*RateField) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{10}
}

func (m *RateField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateField.Unmarshal(m, b)
}
func (m *RateField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateField.Marshal(b, m, deterministic)
}
func (m *RateField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateField.Merge(m, src)
}
func (m *RateField) XXX_Size() int {
	return xxx_messageInfo_RateField.Size(m)
}
func (m *RateField) XXX_DiscardUnknown() {
	xxx_messageInfo_RateField.DiscardUnknown(m)
}

var xxx_messageInfo_RateField proto.InternalMessageInfo

func (m *RateField) GetPrecision() uint32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *RateField) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Membership)(nil), "assets.Membership")
	proto.RegisterType((*Currency)(nil), "assets.Currency")
	proto.RegisterType((*ShareCommon)(nil), "assets.ShareCommon")
	proto.RegisterType((*BondFixedRate)(nil), "assets.BondFixedRate")
	proto.RegisterType((*Coupon)(nil), "assets.Coupon")
	proto.RegisterType((*LoyaltyPoints)(nil), "assets.LoyaltyPoints")
	proto.RegisterType((*TicketAdmission)(nil), "assets.TicketAdmission")
	proto.RegisterType((*CasinoChip)(nil), "assets.CasinoChip")
	proto.RegisterType((*AgeRestrictionField)(nil), "assets.AgeRestrictionField")
	proto.RegisterType((*CurrencyValueField)(nil), "assets.CurrencyValueField")
	proto.RegisterType((*RateField)(nil), "assets.RateField")
}

func init() { proto.RegisterFile("assets.proto", fileDescriptor_610ca40ce07a87fe) }

var fileDescriptor_610ca40ce07a87fe = []byte{
	// 1048 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x41, 0x6f, 0xdb, 0x46,
	0x13, 0x85, 0x24, 0x4a, 0x22, 0x47, 0xb6, 0xa3, 0x6c, 0x3e, 0x04, 0xfc, 0xdc, 0x22, 0x10, 0x74,
	0x28, 0x74, 0x28, 0xec, 0xc2, 0x45, 0xda, 0xa2, 0x28, 0x5a, 0x38, 0x92, 0x0d, 0x48, 0xb0, 0x03,
	0x81, 0x76, 0x5c, 0xa0, 0xb7, 0xb5, 0x38, 0x91, 0x17, 0x21, 0xb9, 0xc2, 0xee, 0xaa, 0x8e, 0x7a,
	0xea, 0xbd, 0xff, 0xa5, 0xbf, 0xab, 0x87, 0x02, 0x45, 0x6f, 0x3d, 0x16, 0x3b, 0x24, 0x25, 0x8a,
	0x16, 0xdb, 0xa4, 0xc8, 0x6d, 0xe7, 0xcd, 0x2c, 0x77, 0xf6, 0xcd, 0xdb, 0x19, 0xc2, 0x1e, 0xd7,
	0x1a, 0x8d, 0x3e, 0x5a, 0x28, 0x69, 0x24, 0x6b, 0xa5, 0x56, 0xff, 0xcf, 0x3a, 0xc0, 0x25, 0xc6,
	0xb7, 0xa8, 0xf4, 0x9d, 0x58, 0xb0, 0x21, 0x1c, 0x9c, 0xce, 0x31, 0x40, 0x6d, 0x94, 0x98, 0x19,
	0x21, 0x13, 0xbf, 0xd6, 0xab, 0x0d, 0x3a, 0x27, 0x1f, 0x1d, 0x65, 0xbb, 0xb7, 0xbd, 0xe7, 0x02,
	0xa3, 0x30, 0x28, 0x6d, 0x61, 0x1f, 0x83, 0x77, 0xc3, 0x23, 0x11, 0x9e, 0x2b, 0x19, 0xfb, 0xf5,
	0x5e, 0x6d, 0xe0, 0x04, 0x1b, 0x80, 0x7d, 0x06, 0x4f, 0xce, 0xde, 0x2e, 0x84, 0xe2, 0x36, 0xf6,
	0x5a, 0xc4, 0xa8, 0x0d, 0x8f, 0x17, 0x7e, 0x83, 0xe2, 0x76, 0xb9, 0xd8, 0x01, 0xd4, 0xc7, 0x23,
	0xdf, 0xe9, 0xd5, 0x06, 0x5e, 0x50, 0x1f, 0x8f, 0xd8, 0x00, 0x1e, 0x6d, 0x52, 0x1e, 0x46, 0x5c,
	0x6b, 0xbf, 0x49, 0xce, 0x32, 0xcc, 0x0e, 0xc1, 0x0d, 0x64, 0x84, 0xd7, 0xab, 0x05, 0xfa, 0x2d,
	0x0a, 0x59, 0xdb, 0xec, 0x13, 0x38, 0xd8, 0x84, 0x53, 0x44, 0x9b, 0x22, 0x4a, 0x28, 0xeb, 0x41,
	0x67, 0x84, 0x7a, 0xa6, 0xc4, 0x82, 0xf8, 0x70, 0x29, 0xa8, 0x08, 0xb1, 0x23, 0x60, 0xd7, 0x8a,
	0x27, 0xfa, 0x35, 0x2a, 0x3d, 0x45, 0x15, 0x0b, 0x63, 0x30, 0xf4, 0xbd, 0x5e, 0x6d, 0xe0, 0x06,
	0x3b, 0x3c, 0xfd, 0x9f, 0x6b, 0xe0, 0x0e, 0x97, 0x4a, 0x61, 0x32, 0x5b, 0xb1, 0x3e, 0xec, 0xe5,
	0xeb, 0xa1, 0x0c, 0x91, 0xf8, 0xf6, 0x82, 0x2d, 0x8c, 0x7d, 0x0a, 0x8f, 0x2f, 0x65, 0x82, 0x86,
	0xab, 0xd5, 0xe9, 0xd2, 0xdc, 0x49, 0x25, 0xcc, 0x8a, 0x88, 0xf5, 0x82, 0x87, 0x0e, 0x4b, 0xff,
	0x54, 0xe1, 0x4c, 0x68, 0x9b, 0xae, 0x93, 0xd2, 0xbf, 0x06, 0x26, 0x8e, 0xdb, 0xe8, 0x3a, 0xfd,
	0x5f, 0x6a, 0xd0, 0xb9, 0xba, 0xe3, 0x0a, 0x87, 0x32, 0x8e, 0x65, 0xc2, 0x9e, 0x42, 0xeb, 0x5a,
	0xcc, 0xde, 0xa0, 0xca, 0xce, 0xcf, 0x2c, 0xc6, 0xc0, 0x19, 0x5f, 0x8d, 0x5f, 0x66, 0x87, 0xd1,
	0xba, 0x4c, 0x48, 0xe3, 0x5d, 0x09, 0x71, 0x2a, 0x09, 0xf9, 0xb5, 0x09, 0xfb, 0x2f, 0x64, 0x12,
	0x9e, 0x8b, 0xb7, 0x18, 0x06, 0xdc, 0xa0, 0x3d, 0xf7, 0x25, 0x8f, 0x73, 0x36, 0x68, 0x6d, 0x8b,
	0x69, 0x83, 0xa8, 0x54, 0x69, 0x3e, 0x6b, 0x7b, 0x9d, 0x67, 0xa3, 0x90, 0xe7, 0x33, 0x80, 0xa1,
	0x8c, 0x22, 0x6e, 0x50, 0xf1, 0x28, 0x93, 0x4f, 0x01, 0x61, 0x5f, 0x80, 0x3b, 0xe5, 0xea, 0x86,
	0x47, 0x4b, 0x24, 0xfd, 0x74, 0x4e, 0x0e, 0x73, 0x95, 0xe7, 0xec, 0x93, 0x33, 0x15, 0xf9, 0x3a,
	0x96, 0x3d, 0x87, 0xbd, 0x71, 0x62, 0x50, 0xa1, 0x36, 0x36, 0x57, 0x12, 0x56, 0xe7, 0xe4, 0x71,
	0xbe, 0xd7, 0x62, 0xe9, 0x96, 0xad, 0x30, 0xf6, 0x2d, 0x1c, 0xe6, 0xf6, 0x94, 0xaf, 0x62, 0x4c,
	0xcc, 0x38, 0x11, 0x46, 0xf0, 0x68, 0x64, 0x3f, 0xd2, 0xa6, 0x3a, 0xfd, 0x43, 0x04, 0xfb, 0x06,
	0xfe, 0x5f, 0xf2, 0x5a, 0x78, 0x84, 0x91, 0xe1, 0xda, 0x77, 0x7b, 0x8d, 0x81, 0x13, 0x54, 0x07,
	0xb0, 0x31, 0x3c, 0xbd, 0xe0, 0x06, 0x33, 0xc7, 0x14, 0x13, 0x1e, 0x99, 0x15, 0xa5, 0xef, 0x55,
	0xa5, 0x5f, 0xb1, 0xc1, 0xaa, 0xb1, 0xe0, 0xf9, 0x5e, 0x24, 0xa1, 0xbc, 0xf7, 0x81, 0xf2, 0x7f,
	0xe8, 0x60, 0x5f, 0x83, 0xff, 0xf0, 0x3b, 0x53, 0x54, 0x42, 0x86, 0x7e, 0x87, 0x36, 0x55, 0xfa,
	0xed, 0xdb, 0xb8, 0xe4, 0x66, 0x69, 0x55, 0x4d, 0x24, 0xed, 0x51, 0xfc, 0x16, 0xb6, 0xa3, 0x63,
	0xed, 0xbf, 0x7f, 0xc7, 0xda, 0x2d, 0xd8, 0x83, 0x4a, 0xc1, 0xfe, 0x5e, 0x87, 0xd6, 0x50, 0x2e,
	0x17, 0x32, 0xb1, 0xcd, 0x28, 0xc0, 0x10, 0x31, 0x16, 0xc9, 0xfc, 0x2c, 0x31, 0xf6, 0x65, 0xa6,
	0xa2, 0x2d, 0xc3, 0xf6, 0x90, 0x75, 0x17, 0xdc, 0xf4, 0xbd, 0xb4, 0x3f, 0xee, 0xf0, 0xd8, 0x2f,
	0x53, 0x37, 0x5c, 0x95, 0x9b, 0x64, 0x19, 0x4e, 0x95, 0x6e, 0xb3, 0xa1, 0x37, 0xd3, 0xca, 0x95,
	0x9e, 0x23, 0x15, 0xd7, 0x73, 0xab, 0xae, 0xc7, 0xbe, 0x02, 0xef, 0x9c, 0xcf, 0x30, 0x7d, 0x1a,
	0xde, 0xbf, 0x3e, 0x8d, 0x4d, 0x70, 0xce, 0x46, 0x4c, 0x7d, 0xe0, 0x06, 0x93, 0x25, 0x92, 0x32,
	0x32, 0x36, 0x0a, 0x30, 0xf3, 0xa1, 0x3d, 0x42, 0xc3, 0x45, 0xa4, 0x49, 0x06, 0x5e, 0x90, 0x9b,
	0x13, 0xc7, 0x75, 0xba, 0xcd, 0x89, 0xe3, 0x36, 0xbb, 0xad, 0x89, 0xe3, 0xb6, 0xbb, 0x6e, 0xff,
	0xaf, 0x1a, 0xec, 0x5f, 0xc8, 0x15, 0x29, 0x43, 0x8a, 0xc4, 0xe8, 0x0f, 0x33, 0xa7, 0x7a, 0xd0,
	0x99, 0x2a, 0x39, 0x57, 0x3c, 0x26, 0xde, 0xd2, 0x9e, 0x52, 0x84, 0xaa, 0x66, 0x95, 0x53, 0x3d,
	0xab, 0x0a, 0xd7, 0x6a, 0x6e, 0x5d, 0xab, 0xa2, 0x08, 0xad, 0xaa, 0x22, 0x64, 0x8d, 0xfa, 0xb7,
	0x3a, 0x3c, 0xa2, 0x5e, 0x6c, 0x4e, 0xc3, 0x58, 0x68, 0xdb, 0xc2, 0x3f, 0xcc, 0xe5, 0xff, 0x07,
	0xcd, 0xb4, 0x3e, 0x69, 0xcb, 0x4c, 0x0d, 0xdb, 0x47, 0x4f, 0x15, 0xf2, 0x2c, 0x77, 0x5a, 0x5b,
	0xec, 0x0a, 0xb9, 0xc9, 0x74, 0x45, 0x6b, 0x3b, 0x3c, 0xcf, 0x7e, 0xc4, 0xc4, 0x6c, 0x38, 0x49,
	0x1b, 0x58, 0x09, 0xb5, 0xb3, 0x88, 0x10, 0x22, 0x38, 0x55, 0xc2, 0x06, 0xa8, 0xa0, 0xa4, 0x53,
	0xa9, 0xcb, 0x02, 0xb9, 0x7b, 0xdb, 0xe4, 0xfa, 0xd0, 0xbe, 0xc2, 0xcd, 0xf3, 0xf7, 0x82, 0xdc,
	0x64, 0x5d, 0x68, 0x04, 0xf2, 0x9e, 0xde, 0xb2, 0x17, 0xd8, 0xe5, 0xc4, 0x71, 0xeb, 0xdd, 0xc6,
	0x5a, 0x65, 0x6e, 0xd7, 0x9b, 0x38, 0xae, 0xd7, 0x85, 0xfe, 0x1f, 0x75, 0x80, 0x21, 0xd7, 0x22,
	0x91, 0x43, 0xfb, 0x2b, 0xf4, 0x2e, 0x83, 0xd9, 0x87, 0xf6, 0x2b, 0x8d, 0x85, 0x89, 0x94, 0x9b,
	0x3b, 0x6a, 0xd4, 0x78, 0xff, 0x1a, 0x55, 0xc8, 0xaf, 0x59, 0x2d, 0xbf, 0xad, 0xd9, 0xdf, 0x2a,
	0xcd, 0xfe, 0x0a, 0xbe, 0xdb, 0x95, 0x7c, 0x3f, 0xcb, 0x09, 0xa1, 0xf2, 0xb9, 0x59, 0x5f, 0x59,
	0x23, 0xff, 0xbd, 0x4f, 0xa4, 0xec, 0xf7, 0x4f, 0xe1, 0xc9, 0x0e, 0x1a, 0xac, 0x34, 0x2f, 0xe4,
	0x7d, 0xf6, 0x2f, 0xb2, 0x1f, 0xa4, 0x86, 0x45, 0x5f, 0x2d, 0x16, 0xa8, 0x88, 0xe9, 0xfd, 0x20,
	0x35, 0xfa, 0x11, 0xb0, 0x87, 0x27, 0x91, 0xb8, 0x29, 0xa9, 0x1a, 0x51, 0x90, 0x1a, 0x0f, 0x2a,
	0x5a, 0xdf, 0x51, 0xd1, 0x2d, 0x02, 0x1b, 0x74, 0xd2, 0x06, 0xe8, 0x7f, 0x07, 0xde, 0x7a, 0x3e,
	0x6e, 0x87, 0xd6, 0x4a, 0xa1, 0x9b, 0x14, 0xea, 0x85, 0x14, 0x5e, 0x7c, 0xf9, 0xc3, 0xf3, 0xb9,
	0x30, 0x77, 0xcb, 0xdb, 0xa3, 0x99, 0x8c, 0x8f, 0x8d, 0x7c, 0x83, 0x89, 0xf8, 0x09, 0xc3, 0x63,
	0xbd, 0xc0, 0x99, 0x78, 0x2d, 0x66, 0x54, 0xce, 0xe3, 0x50, 0x68, 0x73, 0x3c, 0x97, 0x11, 0x4f,
	0xe6, 0xc7, 0x29, 0x9d, 0xb7, 0x2d, 0xfa, 0x6d, 0xff, 0xfc, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x97, 0xbb, 0x59, 0x80, 0xc6, 0x0b, 0x00, 0x00,
}
