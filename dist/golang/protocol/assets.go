package protocol

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	// AssetTypeLen is the size in bytes of all asset type variants.
	AssetTypeLen = 152

	// CodeCoupon identifies data as a Coupon message.
	CodeCoupon = "COU"

	// CodeCurrency identifies data as a Currency message.
	CodeCurrency = "CUR"

	// CodeLoyaltyPoints identifies data as a LoyaltyPoints message.
	CodeLoyaltyPoints = "LOY"

	// CodeMembership identifies data as a Membership message.
	CodeMembership = "MEM"

	// CodeShareCommon identifies data as a ShareCommon message.
	CodeShareCommon = "SHC"

	// CodeTicketAdmission identifies data as a TicketAdmission message.
	CodeTicketAdmission = "TIC"
)

// Coupon asset type.
type Coupon struct {
	Version            uint8
	TradingRestriction []byte
	RedeemingEntity    string
	IssueDate          uint64
	ExpiryDate         uint64
	Value              uint64
	Currency           string
	Description        string
}

// NewCoupon returns a new Coupon.
func NewCoupon() Coupon {
	return Coupon{}
}

// Type returns the type identifer for this message.
func (m Coupon) Type() string {
	return CodeCoupon
}

// Len returns the byte size of this message.
func (m Coupon) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Coupon) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Coupon) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.RedeemingEntity, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.IssueDate); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExpiryDate); err != nil {
		return nil, err
	}

	if err := write(buf, m.Value); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.Currency, 3); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Description, 16); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeCoupon), len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m Coupon) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.RedeemingEntity); err != nil {
		return 0, err
	}

	if err := read(buf, &m.IssueDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExpiryDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Value); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Currency); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Description); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m Coupon) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("RedeemingEntity:%#+v", m.RedeemingEntity))
	vals = append(vals, fmt.Sprintf("IssueDate:%v", m.IssueDate))
	vals = append(vals, fmt.Sprintf("ExpiryDate:%v", m.ExpiryDate))
	vals = append(vals, fmt.Sprintf("Value:%v", m.Value))
	vals = append(vals, fmt.Sprintf("Currency:%#+v", m.Currency))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Currency asset type.
type Currency struct {
	Version            uint8
	TradingRestriction []byte
	ISOCode            string
	MonetaryAuthority  string
	Description        string
}

// NewCurrency returns a new Currency.
func NewCurrency() Currency {
	return Currency{}
}

// Type returns the type identifer for this message.
func (m Currency) Type() string {
	return CodeCurrency
}

// Len returns the byte size of this message.
func (m Currency) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Currency) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Currency) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.ISOCode, 3); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.MonetaryAuthority, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Description, 255); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeCurrency), len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m Currency) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ISOCode); err != nil {
		return 0, err
	}

	if err := read(buf, &m.MonetaryAuthority); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Description); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m Currency) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("ISOCode:%#+v", m.ISOCode))
	vals = append(vals, fmt.Sprintf("MonetaryAuthority:%#+v", m.MonetaryAuthority))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// LoyaltyPoints asset type.
type LoyaltyPoints struct {
	Version             uint8
	TradingRestriction  []byte
	AgeRestriction      []byte
	OfferType           byte
	OfferName           string
	ValidFrom           uint64
	ExpirationTimestamp uint64
	Description         string
}

// NewLoyaltyPoints returns a new LoyaltyPoints.
func NewLoyaltyPoints() LoyaltyPoints {
	return LoyaltyPoints{}
}

// Type returns the type identifer for this message.
func (m LoyaltyPoints) Type() string {
	return CodeLoyaltyPoints
}

// Len returns the byte size of this message.
func (m LoyaltyPoints) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m LoyaltyPoints) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m LoyaltyPoints) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.OfferType); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.OfferName, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Description, 16); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeLoyaltyPoints), len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m LoyaltyPoints) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.OfferType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.OfferName); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Description); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m LoyaltyPoints) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("AgeRestriction:%#x", m.AgeRestriction))
	vals = append(vals, fmt.Sprintf("OfferType:%#+v", m.OfferType))
	vals = append(vals, fmt.Sprintf("OfferName:%#+v", m.OfferName))
	vals = append(vals, fmt.Sprintf("ValidFrom:%#+v", m.ValidFrom))
	vals = append(vals, fmt.Sprintf("ExpirationTimestamp:%#+v", m.ExpirationTimestamp))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// Membership asset type.
type Membership struct {
	Version             uint8
	TradingRestriction  []byte
	AgeRestriction      []byte
	ValidFrom           uint64
	ExpirationTimestamp uint64
	ID                  string
	MembershipType      string
	Description         string
}

// NewMembership returns a new Membership.
func NewMembership() Membership {
	return Membership{}
}

// Type returns the type identifer for this message.
func (m Membership) Type() string {
	return CodeMembership
}

// Len returns the byte size of this message.
func (m Membership) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m Membership) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m Membership) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.ID, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.MembershipType, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Description, 16); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeMembership), len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m Membership) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ID); err != nil {
		return 0, err
	}

	if err := read(buf, &m.MembershipType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Description); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m Membership) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("AgeRestriction:%#x", m.AgeRestriction))
	vals = append(vals, fmt.Sprintf("ValidFrom:%#+v", m.ValidFrom))
	vals = append(vals, fmt.Sprintf("ExpirationTimestamp:%#+v", m.ExpirationTimestamp))
	vals = append(vals, fmt.Sprintf("ID:%#+v", m.ID))
	vals = append(vals, fmt.Sprintf("MembershipType:%#+v", m.MembershipType))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// ShareCommon asset type.
type ShareCommon struct {
	Version            uint8
	TradingRestriction []byte
	TransferLockout    uint64
	Ticker             string
	ISIN               string
	Description        string
}

// NewShareCommon returns a new ShareCommon.
func NewShareCommon() ShareCommon {
	return ShareCommon{}
}

// Type returns the type identifer for this message.
func (m ShareCommon) Type() string {
	return CodeShareCommon
}

// Len returns the byte size of this message.
func (m ShareCommon) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m ShareCommon) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m ShareCommon) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, m.TransferLockout); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.Ticker, 5); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.ISIN, 12); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Description, 113); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeShareCommon), len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m ShareCommon) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.TransferLockout); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Ticker); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ISIN); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Description); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m ShareCommon) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("TransferLockout:%#+v", m.TransferLockout))
	vals = append(vals, fmt.Sprintf("Ticker:%#+v", m.Ticker))
	vals = append(vals, fmt.Sprintf("ISIN:%#+v", m.ISIN))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}

// TicketAdmission asset type.
type TicketAdmission struct {
	Version             uint8
	TradingRestriction  []byte
	AgeRestriction      []byte
	AdmissionType       string
	Venue               string
	Class               string
	Area                string
	Seat                string
	StartTimeDate       uint64
	ValidFrom           uint64
	ExpirationTimestamp uint64
	Description         string
}

// NewTicketAdmission returns a new TicketAdmission.
func NewTicketAdmission() TicketAdmission {
	return TicketAdmission{}
}

// Type returns the type identifer for this message.
func (m TicketAdmission) Type() string {
	return CodeTicketAdmission
}

// Len returns the byte size of this message.
func (m TicketAdmission) Len() int64 {
	return AssetTypeLen
}

// Read implements the io.Reader interface, writing the receiver to the
// []byte.
func (m TicketAdmission) Read(b []byte) (int, error) {
	data, err := m.Serialize()

	if err != nil {
		return 0, err
	}

	copy(b, data)

	return len(b), nil
}

// Serialize returns the full OP_RETURN payload bytes.
func (m TicketAdmission) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := write(buf, m.Version); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.TradingRestriction, 5)); err != nil {
		return nil, err
	}

	if err := write(buf, pad(m.AgeRestriction, 5)); err != nil {
		return nil, err
	}

	if err := WriteFixedChar(buf, m.AdmissionType, 3); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Venue, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Class, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Area, 255); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Seat, 255); err != nil {
		return nil, err
	}

	if err := write(buf, m.StartTimeDate); err != nil {
		return nil, err
	}

	if err := write(buf, m.ValidFrom); err != nil {
		return nil, err
	}

	if err := write(buf, m.ExpirationTimestamp); err != nil {
		return nil, err
	}

	if err := WriteVarChar(buf, m.Description, 16); err != nil {
		return nil, err
	}

	b := buf.Bytes()

	header, err := NewHeaderForCode([]byte(CodeTicketAdmission), len(b))
	if err != nil {
		return nil, err
	}

	h, err := header.Serialize()
	if err != nil {
		return nil, err
	}

	out := append(h, b...)

	return out, nil
}

// Write implements the io.Writer interface, writing the data in []byte to
// the receiver.
func (m TicketAdmission) Write(b []byte) (int, error) {
	buf := bytes.NewBuffer(b)

	if err := read(buf, &m.Version); err != nil {
		return 0, err
	}

	m.TradingRestriction = make([]byte, 5)
	if err := readLen(buf, m.TradingRestriction); err != nil {
		return 0, err
	}

	m.AgeRestriction = make([]byte, 5)
	if err := readLen(buf, m.AgeRestriction); err != nil {
		return 0, err
	}

	if err := read(buf, &m.AdmissionType); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Venue); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Class); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Area); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Seat); err != nil {
		return 0, err
	}

	if err := read(buf, &m.StartTimeDate); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ValidFrom); err != nil {
		return 0, err
	}

	if err := read(buf, &m.ExpirationTimestamp); err != nil {
		return 0, err
	}

	if err := read(buf, &m.Description); err != nil {
		return 0, err
	}

	return len(b), nil
}

func (m TicketAdmission) String() string {
	vals := []string{}

	vals = append(vals, fmt.Sprintf("Version:%v", m.Version))
	vals = append(vals, fmt.Sprintf("TradingRestriction:%#x", m.TradingRestriction))
	vals = append(vals, fmt.Sprintf("AgeRestriction:%#x", m.AgeRestriction))
	vals = append(vals, fmt.Sprintf("AdmissionType:%#+v", m.AdmissionType))
	vals = append(vals, fmt.Sprintf("Venue:%#+v", m.Venue))
	vals = append(vals, fmt.Sprintf("Class:%#+v", m.Class))
	vals = append(vals, fmt.Sprintf("Area:%#+v", m.Area))
	vals = append(vals, fmt.Sprintf("Seat:%#+v", m.Seat))
	vals = append(vals, fmt.Sprintf("StartTimeDate:%#+v", m.StartTimeDate))
	vals = append(vals, fmt.Sprintf("ValidFrom:%#+v", m.ValidFrom))
	vals = append(vals, fmt.Sprintf("ExpirationTimestamp:%#+v", m.ExpirationTimestamp))
	vals = append(vals, fmt.Sprintf("Description:%#+v", m.Description))

	return fmt.Sprintf("{%s}", strings.Join(vals, " "))
}
