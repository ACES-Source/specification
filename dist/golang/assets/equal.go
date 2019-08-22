package assets

import proto "github.com/golang/protobuf/proto"

func (l *Membership) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Membership)
	if !ok {
		return false
	}

	// Field AgeRestriction - AgeRestriction

	// Field ValidFrom - uint
	if l.ValidFrom != r.ValidFrom {
		return false // fmt.Errorf("ValidFrom integer mismatched")
	}

	// Field ExpirationTimestamp - uint
	if l.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field ID - varchar
	if l.ID != r.ID {
		return false // fmt.Errorf("ID string mismatched")
	}

	// Field MembershipClass - varchar
	if l.MembershipClass != r.MembershipClass {
		return false // fmt.Errorf("MembershipClass string mismatched")
	}

	// Field RoleType - varchar
	if l.RoleType != r.RoleType {
		return false // fmt.Errorf("RoleType string mismatched")
	}

	// Field MembershipType - varchar
	if l.MembershipType != r.MembershipType {
		return false // fmt.Errorf("MembershipType string mismatched")
	}

	// Field Description - varchar
	if l.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	return true
}

func (l *Currency) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Currency)
	if !ok {
		return false
	}

	// Field ISOCode - fixedchar
	if l.ISOCode != r.ISOCode {
		return false // fmt.Errorf("ISOCode string mismatched")
	}

	// Field MonetaryAuthority - varchar
	if l.MonetaryAuthority != r.MonetaryAuthority {
		return false // fmt.Errorf("MonetaryAuthority string mismatched")
	}

	// Field Description - varchar
	if l.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	return true
}

func (l *ShareCommon) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*ShareCommon)
	if !ok {
		return false
	}

	// Field TransferLockout - uint
	if l.TransferLockout != r.TransferLockout {
		return false // fmt.Errorf("TransferLockout integer mismatched")
	}

	// Field Ticker - fixedchar
	if l.Ticker != r.Ticker {
		return false // fmt.Errorf("Ticker string mismatched")
	}

	// Field ISIN - fixedchar
	if l.ISIN != r.ISIN {
		return false // fmt.Errorf("ISIN string mismatched")
	}

	// Field Description - varchar
	if l.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	return true
}

func (l *Coupon) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*Coupon)
	if !ok {
		return false
	}

	// Field RedeemingEntity - varchar
	if l.RedeemingEntity != r.RedeemingEntity {
		return false // fmt.Errorf("RedeemingEntity string mismatched")
	}

	// Field IssueDate - uint
	if l.IssueDate != r.IssueDate {
		return false // fmt.Errorf("IssueDate integer mismatched")
	}

	// Field ExpiryDate - uint
	if l.ExpiryDate != r.ExpiryDate {
		return false // fmt.Errorf("ExpiryDate integer mismatched")
	}

	// Field Value - uint
	if l.Value != r.Value {
		return false // fmt.Errorf("Value integer mismatched")
	}

	// Field Currency - fixedchar
	if l.Currency != r.Currency {
		return false // fmt.Errorf("Currency string mismatched")
	}

	// Field Description - varchar
	if l.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	return true
}

func (l *LoyaltyPoints) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*LoyaltyPoints)
	if !ok {
		return false
	}

	// Field AgeRestriction - AgeRestriction

	// Field OfferName - varchar
	if l.OfferName != r.OfferName {
		return false // fmt.Errorf("OfferName string mismatched")
	}

	// Field ValidFrom - uint
	if l.ValidFrom != r.ValidFrom {
		return false // fmt.Errorf("ValidFrom integer mismatched")
	}

	// Field ExpirationTimestamp - uint
	if l.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field Description - varchar
	if l.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	return true
}

func (l *TicketAdmission) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*TicketAdmission)
	if !ok {
		return false
	}

	// Field AgeRestriction - AgeRestriction

	// Field AdmissionType - fixedchar
	if l.AdmissionType != r.AdmissionType {
		return false // fmt.Errorf("AdmissionType string mismatched")
	}

	// Field Venue - varchar
	if l.Venue != r.Venue {
		return false // fmt.Errorf("Venue string mismatched")
	}

	// Field Class - varchar
	if l.Class != r.Class {
		return false // fmt.Errorf("Class string mismatched")
	}

	// Field Area - varchar
	if l.Area != r.Area {
		return false // fmt.Errorf("Area string mismatched")
	}

	// Field Seat - varchar
	if l.Seat != r.Seat {
		return false // fmt.Errorf("Seat string mismatched")
	}

	// Field StartTimeDate - uint
	if l.StartTimeDate != r.StartTimeDate {
		return false // fmt.Errorf("StartTimeDate integer mismatched")
	}

	// Field ValidFrom - uint
	if l.ValidFrom != r.ValidFrom {
		return false // fmt.Errorf("ValidFrom integer mismatched")
	}

	// Field ExpirationTimestamp - uint
	if l.ExpirationTimestamp != r.ExpirationTimestamp {
		return false // fmt.Errorf("ExpirationTimestamp integer mismatched")
	}

	// Field Description - varchar
	if l.Description != r.Description {
		return false // fmt.Errorf("Description string mismatched")
	}

	return true
}

func (l *AgeRestrictionField) Equal(right proto.Message) bool {
	if l == nil {
		return right == nil
	}
	r, ok := right.(*AgeRestrictionField)
	if !ok {
		return false
	}

	// Field Lower - uint
	if l.Lower != r.Lower {
		return false // fmt.Errorf("Lower integer mismatched")
	}

	// Field Upper - uint
	if l.Upper != r.Upper {
		return false // fmt.Errorf("Upper integer mismatched")
	}

	return true
}
