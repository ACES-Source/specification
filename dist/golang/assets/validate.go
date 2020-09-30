package assets

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/tokenized/pkg/bitcoin"
)

const (
	max1ByteInteger = 255
	max2ByteInteger = 65535
	max4ByteInteger = 4294967295
)

func (a *Membership) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field ID - varchar
	if len(a.ID) > max1ByteInteger {
		return fmt.Errorf("ID over max size : %d > %d", len(a.ID), max1ByteInteger)
	}

	// Field MembershipClass - varchar
	if len(a.MembershipClass) > max1ByteInteger {
		return fmt.Errorf("MembershipClass over max size : %d > %d", len(a.MembershipClass), max1ByteInteger)
	}

	// Field RoleType - varchar
	if len(a.RoleType) > max1ByteInteger {
		return fmt.Errorf("RoleType over max size : %d > %d", len(a.RoleType), max1ByteInteger)
	}

	// Field MembershipType - varchar
	if len(a.MembershipType) > max1ByteInteger {
		return fmt.Errorf("MembershipType over max size : %d > %d", len(a.MembershipType), max1ByteInteger)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("Description over max size : %d > %d", len(a.Description), max2ByteInteger)
	}
	if len(a.Description) == 0 {
		return fmt.Errorf("Description required")
	}

	return nil
}

func (a *Currency) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field CurrencyCode - fixedchar
	if len(a.CurrencyCode) != 0 && len(a.CurrencyCode) != 3 {
		return fmt.Errorf("CurrencyCode fixed width field wrong size : %d should be %d",
			len(a.CurrencyCode), 3)
	}
	if len(a.CurrencyCode) == 0 {
		return fmt.Errorf("CurrencyCode required")
	}

	// Field MonetaryAuthority - varchar
	if len(a.MonetaryAuthority) > max1ByteInteger {
		return fmt.Errorf("MonetaryAuthority over max size : %d > %d", len(a.MonetaryAuthority), max1ByteInteger)
	}

	// Field Precision - uint
	if a.Precision == 0 {
		return fmt.Errorf("Precision required")
	}

	return nil
}

func (a *ShareCommon) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field Ticker - fixedchar
	if len(a.Ticker) != 0 && len(a.Ticker) != 5 {
		return fmt.Errorf("Ticker fixed width field wrong size : %d should be %d",
			len(a.Ticker), 5)
	}

	// Field ISIN - fixedchar
	if len(a.ISIN) != 0 && len(a.ISIN) != 12 {
		return fmt.Errorf("ISIN fixed width field wrong size : %d should be %d",
			len(a.ISIN), 12)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("Description over max size : %d > %d", len(a.Description), max2ByteInteger)
	}
	if len(a.Description) == 0 {
		return fmt.Errorf("Description required")
	}

	return nil
}

func (a *Coupon) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	ValueFieldIsEmpty := a.Value == 0

	// Field RedeemingEntity - varchar
	if len(a.RedeemingEntity) > max1ByteInteger {
		return fmt.Errorf("RedeemingEntity over max size : %d > %d", len(a.RedeemingEntity), max1ByteInteger)
	}

	// Field IssueDate - uint

	// Field ExpiryDate - uint

	// Field Value - uint

	// Field Currency - fixedchar
	if len(a.Currency) != 0 && len(a.Currency) != 3 {
		return fmt.Errorf("Currency fixed width field wrong size : %d should be %d",
			len(a.Currency), 3)
	}
	if ValueFieldIsEmpty && len(a.Currency) != 0 {
		return fmt.Errorf("Currency is only allowed when Value is specified : %v", a.Value)
	}
	if !ValueFieldIsEmpty && len(a.Currency) == 0 {
		return fmt.Errorf("Currency is required when Value is specified : %v", a.Value)
	}

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("Description over max size : %d > %d", len(a.Description), max2ByteInteger)
	}
	if len(a.Description) == 0 {
		return fmt.Errorf("Description required")
	}

	// Field Precision - uint
	if ValueFieldIsEmpty && a.Precision != 0 {
		return fmt.Errorf("Precision is only allowed when Value is specified : %v", a.Value)
	}
	if !ValueFieldIsEmpty && a.Precision == 0 {
		return fmt.Errorf("Precision is required when Value is specified : %v", a.Value)
	}

	return nil
}

func (a *LoyaltyPoints) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field OfferName - varchar
	if len(a.OfferName) > max1ByteInteger {
		return fmt.Errorf("OfferName over max size : %d > %d", len(a.OfferName), max1ByteInteger)
	}

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("Description over max size : %d > %d", len(a.Description), max2ByteInteger)
	}
	if len(a.Description) == 0 {
		return fmt.Errorf("Description required")
	}

	return nil
}

func (a *TicketAdmission) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field AdmissionType - fixedchar
	if len(a.AdmissionType) != 0 && len(a.AdmissionType) != 3 {
		return fmt.Errorf("AdmissionType fixed width field wrong size : %d should be %d",
			len(a.AdmissionType), 3)
	}

	// Field Venue - varchar
	if len(a.Venue) > max1ByteInteger {
		return fmt.Errorf("Venue over max size : %d > %d", len(a.Venue), max1ByteInteger)
	}

	// Field Class - varchar
	if len(a.Class) > max1ByteInteger {
		return fmt.Errorf("Class over max size : %d > %d", len(a.Class), max1ByteInteger)
	}

	// Field Area - varchar
	if len(a.Area) > max1ByteInteger {
		return fmt.Errorf("Area over max size : %d > %d", len(a.Area), max1ByteInteger)
	}

	// Field Seat - varchar
	if len(a.Seat) > max1ByteInteger {
		return fmt.Errorf("Seat over max size : %d > %d", len(a.Seat), max1ByteInteger)
	}

	// Field StartTimeDate - uint

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field Description - varchar
	if len(a.Description) > max2ByteInteger {
		return fmt.Errorf("Description over max size : %d > %d", len(a.Description), max2ByteInteger)
	}
	if len(a.Description) == 0 {
		return fmt.Errorf("Description required")
	}

	return nil
}

func (a *CasinoChip) Validate() error {
	if a == nil {
		return errors.New("Empty")
	}

	// Field CurrencyCode - fixedchar
	if len(a.CurrencyCode) != 0 && len(a.CurrencyCode) != 3 {
		return fmt.Errorf("CurrencyCode fixed width field wrong size : %d should be %d",
			len(a.CurrencyCode), 3)
	}
	if len(a.CurrencyCode) == 0 {
		return fmt.Errorf("CurrencyCode required")
	}

	// Field UseType - fixedchar
	if len(a.UseType) != 0 && len(a.UseType) != 1 {
		return fmt.Errorf("UseType fixed width field wrong size : %d should be %d",
			len(a.UseType), 1)
	}

	// Field AgeRestriction - AgeRestriction
	if err := a.AgeRestriction.Validate(); err != nil {
		return errors.Wrap(err, "AgeRestriction")
	}

	// Field ValidFrom - uint

	// Field ExpirationTimestamp - uint

	// Field Precision - uint
	if a.Precision == 0 {
		return fmt.Errorf("Precision required")
	}

	return nil
}

func (a *AgeRestrictionField) Validate() error {
	if a == nil {
		return nil
	}

	// Field Lower - uint
	if a.Lower > uint32(max1ByteInteger) {
		return fmt.Errorf("Lower over max value : %d > %d", a.Lower, max1ByteInteger)
	}

	// Field Upper - uint
	if a.Upper > uint32(max1ByteInteger) {
		return fmt.Errorf("Upper over max value : %d > %d", a.Upper, max1ByteInteger)
	}

	return nil
}

// AddressIsValid returns true if an "Address" alias field is valid.
func AddressIsValid(b []byte) error {
	_, err := bitcoin.DecodeRawAddress(b)
	return err
}

// PublicKeyIsValid returns true if a "PublicKey" alias field is valid.
func PublicKeyIsValid(b []byte) error {
	_, err := bitcoin.PublicKeyFromBytes(b)
	return err
}

// SignatureIsValid returns true if a "Signature" alias field is valid.
func SignatureIsValid(b []byte) error {
	_, err := bitcoin.SignatureFromBytes(b)
	return err
}
