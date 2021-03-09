package assets

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/tokenized/pkg/bitcoin"
	"github.com/tokenized/specification/dist/golang/internal"
	"github.com/tokenized/specification/dist/golang/permissions"

	"github.com/pkg/errors"
)

const (
	// AmendmentOperationModify specifies an amendment is modifying a value.
	AmendmentOperationModify = uint32(0)

	// AmendmentOperationAddElement specifies an amendment is adding a new element to a list.
	AmendmentOperationAddElement = uint32(1)

	// AmendmentOperationRemoveElement specifies an amendment is removing an element from a list.
	AmendmentOperationRemoveElement = uint32(2)
)

// Membership Permission / Amendment Field Indices
const (
	MembershipFieldAgeRestriction      = uint32(1)
	MembershipFieldValidFrom           = uint32(2)
	MembershipFieldExpirationTimestamp = uint32(3)
	MembershipFieldID                  = uint32(4)
	MembershipFieldMembershipClass     = uint32(5)
	MembershipFieldRoleType            = uint32(6)
	MembershipFieldMembershipType      = uint32(7)
	MembershipFieldDescription         = uint32(8)
	MembershipFieldTransfersPermitted  = uint32(9)
)

// ApplyAmendment updates a Membership based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *Membership) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case MembershipFieldAgeRestriction: // AgeRestrictionField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case MembershipFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldID: // string
		a.ID = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldMembershipClass: // string
		a.MembershipClass = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldRoleType: // string
		a.RoleType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldMembershipType: // string
		a.MembershipType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldDescription: // string
		a.Description = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case MembershipFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Membership amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Memberships and returns
// amendment data.
func (a *Membership) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *Membership) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, MembershipFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// ValidFrom uint64
	fip = append(ofip, MembershipFieldValidFrom)
	if a.ValidFrom != newValue.ValidFrom {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ValidFrom)); err != nil {
			return nil, errors.Wrap(err, "ValidFrom")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, MembershipFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ID string
	fip = append(ofip, MembershipFieldID)
	if a.ID != newValue.ID {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ID),
		})
	}

	// MembershipClass string
	fip = append(ofip, MembershipFieldMembershipClass)
	if a.MembershipClass != newValue.MembershipClass {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.MembershipClass),
		})
	}

	// RoleType string
	fip = append(ofip, MembershipFieldRoleType)
	if a.RoleType != newValue.RoleType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.RoleType),
		})
	}

	// MembershipType string
	fip = append(ofip, MembershipFieldMembershipType)
	if a.MembershipType != newValue.MembershipType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.MembershipType),
		})
	}

	// Description string
	fip = append(ofip, MembershipFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, MembershipFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// Currency Permission / Amendment Field Indices
const (
	CurrencyFieldCurrencyCode          = uint32(1)
	CurrencyFieldMonetaryAuthority     = uint32(2)
	DeprecatedCurrencyFieldDescription = uint32(3)
	CurrencyFieldPrecision             = uint32(4)
)

// ApplyAmendment updates a Currency based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *Currency) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case CurrencyFieldCurrencyCode: // string
		if CurrenciesData(a.CurrencyCode) == nil {
			return nil, fmt.Errorf("Currencies resource value not defined : %v", a.CurrencyCode)
		}
		a.CurrencyCode = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case CurrencyFieldMonetaryAuthority: // string
		a.MonetaryAuthority = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedCurrencyFieldDescription: // deprecated

	case CurrencyFieldPrecision: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown Currency amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Currencys and returns
// amendment data.
func (a *Currency) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *Currency) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// CurrencyCode string
	fip = append(ofip, CurrencyFieldCurrencyCode)
	if a.CurrencyCode != newValue.CurrencyCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CurrencyCode),
		})
	}

	// MonetaryAuthority string
	fip = append(ofip, CurrencyFieldMonetaryAuthority)
	if a.MonetaryAuthority != newValue.MonetaryAuthority {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.MonetaryAuthority),
		})
	}

	// deprecated Description deprecated

	// Precision uint64
	fip = append(ofip, CurrencyFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// ShareCommon Permission / Amendment Field Indices
const (
	ShareCommonFieldTicker             = uint32(1)
	ShareCommonFieldISIN               = uint32(2)
	ShareCommonFieldDescription        = uint32(3)
	ShareCommonFieldTransfersPermitted = uint32(4)
)

// ApplyAmendment updates a ShareCommon based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *ShareCommon) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case ShareCommonFieldTicker: // string
		a.Ticker = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ShareCommonFieldISIN: // string
		a.ISIN = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ShareCommonFieldDescription: // string
		a.Description = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case ShareCommonFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown ShareCommon amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two ShareCommons and returns
// amendment data.
func (a *ShareCommon) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *ShareCommon) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Ticker string
	fip = append(ofip, ShareCommonFieldTicker)
	if a.Ticker != newValue.Ticker {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Ticker),
		})
	}

	// ISIN string
	fip = append(ofip, ShareCommonFieldISIN)
	if a.ISIN != newValue.ISIN {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ISIN),
		})
	}

	// Description string
	fip = append(ofip, ShareCommonFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, ShareCommonFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// BondFixedRate Permission / Amendment Field Indices
const (
	BondFixedRateFieldName                     = uint32(1)
	BondFixedRateFieldType                     = uint32(2)
	BondFixedRateFieldISIN                     = uint32(3)
	BondFixedRateFieldCollateral               = uint32(4)
	BondFixedRateFieldParValue                 = uint32(5)
	BondFixedRateFieldInterestRate             = uint32(6)
	BondFixedRateFieldInterestPaymentPeriod    = uint32(7)
	BondFixedRateFieldLatePaymentPenaltyRate   = uint32(8)
	BondFixedRateFieldLatePaymentPenaltyPeriod = uint32(9)
	BondFixedRateFieldLatePaymentWindow        = uint32(10)
	BondFixedRateFieldMaturityDate             = uint32(11)
	BondFixedRateFieldAgeRestriction           = uint32(12)
	BondFixedRateFieldTransfersPermitted       = uint32(13)
)

// ApplyAmendment updates a BondFixedRate based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *BondFixedRate) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case BondFixedRateFieldName: // string
		a.Name = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldType: // string
		a.Type = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldISIN: // string
		a.ISIN = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldCollateral: // string
		a.Collateral = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldParValue: // CurrencyValueField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.ParValue.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case BondFixedRateFieldInterestRate: // RateField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.InterestRate.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case BondFixedRateFieldInterestPaymentPeriod: // PeriodField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.InterestPaymentPeriod.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case BondFixedRateFieldLatePaymentPenaltyRate: // RateField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.LatePaymentPenaltyRate.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case BondFixedRateFieldLatePaymentPenaltyPeriod: // PeriodField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.LatePaymentPenaltyPeriod.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case BondFixedRateFieldLatePaymentWindow: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for LatePaymentWindow : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("LatePaymentWindow amendment value failed to deserialize : %s", err)
		} else {
			a.LatePaymentWindow = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldMaturityDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for MaturityDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("MaturityDate amendment value failed to deserialize : %s", err)
		} else {
			a.MaturityDate = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case BondFixedRateFieldAgeRestriction: // AgeRestrictionField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case BondFixedRateFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown BondFixedRate amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two BondFixedRates and returns
// amendment data.
func (a *BondFixedRate) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *BondFixedRate) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Name string
	fip = append(ofip, BondFixedRateFieldName)
	if a.Name != newValue.Name {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Name),
		})
	}

	// Type string
	fip = append(ofip, BondFixedRateFieldType)
	if a.Type != newValue.Type {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Type),
		})
	}

	// ISIN string
	fip = append(ofip, BondFixedRateFieldISIN)
	if a.ISIN != newValue.ISIN {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.ISIN),
		})
	}

	// Collateral string
	fip = append(ofip, BondFixedRateFieldCollateral)
	if a.Collateral != newValue.Collateral {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Collateral),
		})
	}

	// ParValue CurrencyValueField
	fip = append(ofip, BondFixedRateFieldParValue)

	ParValueAmendments, err := a.ParValue.CreateAmendments(fip, newValue.ParValue)
	if err != nil {
		return nil, errors.Wrap(err, "ParValue")
	}
	result = append(result, ParValueAmendments...)

	// InterestRate RateField
	fip = append(ofip, BondFixedRateFieldInterestRate)

	InterestRateAmendments, err := a.InterestRate.CreateAmendments(fip, newValue.InterestRate)
	if err != nil {
		return nil, errors.Wrap(err, "InterestRate")
	}
	result = append(result, InterestRateAmendments...)

	// InterestPaymentPeriod PeriodField
	fip = append(ofip, BondFixedRateFieldInterestPaymentPeriod)

	InterestPaymentPeriodAmendments, err := a.InterestPaymentPeriod.CreateAmendments(fip, newValue.InterestPaymentPeriod)
	if err != nil {
		return nil, errors.Wrap(err, "InterestPaymentPeriod")
	}
	result = append(result, InterestPaymentPeriodAmendments...)

	// LatePaymentPenaltyRate RateField
	fip = append(ofip, BondFixedRateFieldLatePaymentPenaltyRate)

	LatePaymentPenaltyRateAmendments, err := a.LatePaymentPenaltyRate.CreateAmendments(fip, newValue.LatePaymentPenaltyRate)
	if err != nil {
		return nil, errors.Wrap(err, "LatePaymentPenaltyRate")
	}
	result = append(result, LatePaymentPenaltyRateAmendments...)

	// LatePaymentPenaltyPeriod PeriodField
	fip = append(ofip, BondFixedRateFieldLatePaymentPenaltyPeriod)

	LatePaymentPenaltyPeriodAmendments, err := a.LatePaymentPenaltyPeriod.CreateAmendments(fip, newValue.LatePaymentPenaltyPeriod)
	if err != nil {
		return nil, errors.Wrap(err, "LatePaymentPenaltyPeriod")
	}
	result = append(result, LatePaymentPenaltyPeriodAmendments...)

	// LatePaymentWindow uint64
	fip = append(ofip, BondFixedRateFieldLatePaymentWindow)
	if a.LatePaymentWindow != newValue.LatePaymentWindow {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.LatePaymentWindow)); err != nil {
			return nil, errors.Wrap(err, "LatePaymentWindow")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// MaturityDate uint64
	fip = append(ofip, BondFixedRateFieldMaturityDate)
	if a.MaturityDate != newValue.MaturityDate {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.MaturityDate)); err != nil {
			return nil, errors.Wrap(err, "MaturityDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, BondFixedRateFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// TransfersPermitted bool
	fip = append(ofip, BondFixedRateFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// Coupon Permission / Amendment Field Indices
const (
	CouponFieldRedeemingEntity     = uint32(1)
	CouponFieldIssueDate           = uint32(2)
	CouponFieldExpiryDate          = uint32(3)
	DeprecatedCouponFieldValue     = uint32(4)
	DeprecatedCouponFieldCurrency  = uint32(5)
	CouponFieldDescription         = uint32(6)
	DeprecatedCouponFieldPrecision = uint32(7)
	CouponFieldTransfersPermitted  = uint32(8)
	CouponFieldValue               = uint32(9)
)

// ApplyAmendment updates a Coupon based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *Coupon) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case CouponFieldRedeemingEntity: // string
		a.RedeemingEntity = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case CouponFieldIssueDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for IssueDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("IssueDate amendment value failed to deserialize : %s", err)
		} else {
			a.IssueDate = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case CouponFieldExpiryDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpiryDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpiryDate amendment value failed to deserialize : %s", err)
		} else {
			a.ExpiryDate = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedCouponFieldValue: // deprecated

	case DeprecatedCouponFieldCurrency: // deprecated

	case CouponFieldDescription: // string
		a.Description = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case DeprecatedCouponFieldPrecision: // deprecated

	case CouponFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	case CouponFieldValue: // CurrencyValueField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.Value.ApplyAmendment(fip[1:], operation, data, subPermissions)

	}

	return nil, fmt.Errorf("Unknown Coupon amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Coupons and returns
// amendment data.
func (a *Coupon) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *Coupon) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// RedeemingEntity string
	fip = append(ofip, CouponFieldRedeemingEntity)
	if a.RedeemingEntity != newValue.RedeemingEntity {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.RedeemingEntity),
		})
	}

	// IssueDate uint64
	fip = append(ofip, CouponFieldIssueDate)
	if a.IssueDate != newValue.IssueDate {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.IssueDate)); err != nil {
			return nil, errors.Wrap(err, "IssueDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpiryDate uint64
	fip = append(ofip, CouponFieldExpiryDate)
	if a.ExpiryDate != newValue.ExpiryDate {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpiryDate)); err != nil {
			return nil, errors.Wrap(err, "ExpiryDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// deprecated Value deprecated

	// deprecated Currency deprecated

	// Description string
	fip = append(ofip, CouponFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	// deprecated Precision deprecated

	// TransfersPermitted bool
	fip = append(ofip, CouponFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Value CurrencyValueField
	fip = append(ofip, CouponFieldValue)

	ValueAmendments, err := a.Value.CreateAmendments(fip, newValue.Value)
	if err != nil {
		return nil, errors.Wrap(err, "Value")
	}
	result = append(result, ValueAmendments...)

	return result, nil
}

// LoyaltyPoints Permission / Amendment Field Indices
const (
	LoyaltyPointsFieldAgeRestriction      = uint32(1)
	LoyaltyPointsFieldOfferName           = uint32(2)
	LoyaltyPointsFieldValidFrom           = uint32(3)
	LoyaltyPointsFieldExpirationTimestamp = uint32(4)
	LoyaltyPointsFieldDescription         = uint32(5)
	LoyaltyPointsFieldTransfersPermitted  = uint32(6)
)

// ApplyAmendment updates a LoyaltyPoints based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *LoyaltyPoints) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case LoyaltyPointsFieldAgeRestriction: // AgeRestrictionField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case LoyaltyPointsFieldOfferName: // string
		a.OfferName = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case LoyaltyPointsFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case LoyaltyPointsFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case LoyaltyPointsFieldDescription: // string
		a.Description = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case LoyaltyPointsFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown LoyaltyPoints amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two LoyaltyPointss and returns
// amendment data.
func (a *LoyaltyPoints) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *LoyaltyPoints) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, LoyaltyPointsFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// OfferName string
	fip = append(ofip, LoyaltyPointsFieldOfferName)
	if a.OfferName != newValue.OfferName {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.OfferName),
		})
	}

	// ValidFrom uint64
	fip = append(ofip, LoyaltyPointsFieldValidFrom)
	if a.ValidFrom != newValue.ValidFrom {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ValidFrom)); err != nil {
			return nil, errors.Wrap(err, "ValidFrom")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, LoyaltyPointsFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Description string
	fip = append(ofip, LoyaltyPointsFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, LoyaltyPointsFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// TicketAdmission Permission / Amendment Field Indices
const (
	TicketAdmissionFieldAgeRestriction      = uint32(1)
	TicketAdmissionFieldAdmissionType       = uint32(2)
	TicketAdmissionFieldVenue               = uint32(3)
	TicketAdmissionFieldClass               = uint32(4)
	TicketAdmissionFieldArea                = uint32(5)
	TicketAdmissionFieldSeat                = uint32(6)
	TicketAdmissionFieldStartTimeDate       = uint32(7)
	TicketAdmissionFieldValidFrom           = uint32(8)
	TicketAdmissionFieldExpirationTimestamp = uint32(9)
	TicketAdmissionFieldDescription         = uint32(10)
	TicketAdmissionFieldTransfersPermitted  = uint32(11)
)

// ApplyAmendment updates a TicketAdmission based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *TicketAdmission) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case TicketAdmissionFieldAgeRestriction: // AgeRestrictionField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case TicketAdmissionFieldAdmissionType: // string
		a.AdmissionType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldVenue: // string
		a.Venue = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldClass: // string
		a.Class = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldArea: // string
		a.Area = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldSeat: // string
		a.Seat = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldStartTimeDate: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for StartTimeDate : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("StartTimeDate amendment value failed to deserialize : %s", err)
		} else {
			a.StartTimeDate = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldDescription: // string
		a.Description = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case TicketAdmissionFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown TicketAdmission amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two TicketAdmissions and returns
// amendment data.
func (a *TicketAdmission) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *TicketAdmission) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, TicketAdmissionFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// AdmissionType string
	fip = append(ofip, TicketAdmissionFieldAdmissionType)
	if a.AdmissionType != newValue.AdmissionType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.AdmissionType),
		})
	}

	// Venue string
	fip = append(ofip, TicketAdmissionFieldVenue)
	if a.Venue != newValue.Venue {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Venue),
		})
	}

	// Class string
	fip = append(ofip, TicketAdmissionFieldClass)
	if a.Class != newValue.Class {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Class),
		})
	}

	// Area string
	fip = append(ofip, TicketAdmissionFieldArea)
	if a.Area != newValue.Area {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Area),
		})
	}

	// Seat string
	fip = append(ofip, TicketAdmissionFieldSeat)
	if a.Seat != newValue.Seat {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Seat),
		})
	}

	// StartTimeDate uint64
	fip = append(ofip, TicketAdmissionFieldStartTimeDate)
	if a.StartTimeDate != newValue.StartTimeDate {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.StartTimeDate)); err != nil {
			return nil, errors.Wrap(err, "StartTimeDate")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ValidFrom uint64
	fip = append(ofip, TicketAdmissionFieldValidFrom)
	if a.ValidFrom != newValue.ValidFrom {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ValidFrom)); err != nil {
			return nil, errors.Wrap(err, "ValidFrom")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, TicketAdmissionFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Description string
	fip = append(ofip, TicketAdmissionFieldDescription)
	if a.Description != newValue.Description {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Description),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, TicketAdmissionFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// CasinoChip Permission / Amendment Field Indices
const (
	CasinoChipFieldCurrencyCode        = uint32(1)
	CasinoChipFieldUseType             = uint32(2)
	CasinoChipFieldAgeRestriction      = uint32(3)
	CasinoChipFieldValidFrom           = uint32(4)
	CasinoChipFieldExpirationTimestamp = uint32(5)
	CasinoChipFieldPrecision           = uint32(6)
	CasinoChipFieldTransfersPermitted  = uint32(7)
)

// ApplyAmendment updates a CasinoChip based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *CasinoChip) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty asset amendment field index path")
	}

	switch fip[0] {
	case CasinoChipFieldCurrencyCode: // string
		if CurrenciesData(a.CurrencyCode) == nil {
			return nil, fmt.Errorf("Currencies resource value not defined : %v", a.CurrencyCode)
		}
		a.CurrencyCode = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case CasinoChipFieldUseType: // string
		a.UseType = string(data)
		return permissions.SubPermissions(fip, operation, false)

	case CasinoChipFieldAgeRestriction: // AgeRestrictionField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.AgeRestriction.ApplyAmendment(fip[1:], operation, data, subPermissions)

	case CasinoChipFieldValidFrom: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ValidFrom : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ValidFrom amendment value failed to deserialize : %s", err)
		} else {
			a.ValidFrom = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case CasinoChipFieldExpirationTimestamp: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for ExpirationTimestamp : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("ExpirationTimestamp amendment value failed to deserialize : %s", err)
		} else {
			a.ExpirationTimestamp = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case CasinoChipFieldPrecision: // uint64
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)

	case CasinoChipFieldTransfersPermitted: // bool
		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TransfersPermitted : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TransfersPermitted amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TransfersPermitted); err != nil {
			return nil, fmt.Errorf("TransfersPermitted amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)

	}

	return nil, fmt.Errorf("Unknown CasinoChip amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two CasinoChips and returns
// amendment data.
func (a *CasinoChip) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *CasinoChip) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// CurrencyCode string
	fip = append(ofip, CasinoChipFieldCurrencyCode)
	if a.CurrencyCode != newValue.CurrencyCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CurrencyCode),
		})
	}

	// UseType string
	fip = append(ofip, CasinoChipFieldUseType)
	if a.UseType != newValue.UseType {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.UseType),
		})
	}

	// AgeRestriction AgeRestrictionField
	fip = append(ofip, CasinoChipFieldAgeRestriction)

	AgeRestrictionAmendments, err := a.AgeRestriction.CreateAmendments(fip, newValue.AgeRestriction)
	if err != nil {
		return nil, errors.Wrap(err, "AgeRestriction")
	}
	result = append(result, AgeRestrictionAmendments...)

	// ValidFrom uint64
	fip = append(ofip, CasinoChipFieldValidFrom)
	if a.ValidFrom != newValue.ValidFrom {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ValidFrom)); err != nil {
			return nil, errors.Wrap(err, "ValidFrom")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// ExpirationTimestamp uint64
	fip = append(ofip, CasinoChipFieldExpirationTimestamp)
	if a.ExpirationTimestamp != newValue.ExpirationTimestamp {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.ExpirationTimestamp)); err != nil {
			return nil, errors.Wrap(err, "ExpirationTimestamp")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Precision uint64
	fip = append(ofip, CasinoChipFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// TransfersPermitted bool
	fip = append(ofip, CasinoChipFieldTransfersPermitted)
	if a.TransfersPermitted != newValue.TransfersPermitted {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TransfersPermitted); err != nil {
			return nil, errors.Wrap(err, "TransfersPermitted")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// AgeRestrictionField Permission / Amendment Field Indices
const (
	AgeRestrictionFieldLower = uint32(1)
	AgeRestrictionFieldUpper = uint32(2)
)

// ApplyAmendment updates a AgeRestrictionField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *AgeRestrictionField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty AgeRestriction amendment field index path")
	}

	switch fip[0] {
	case AgeRestrictionFieldLower: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Lower : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Lower amendment value failed to deserialize : %s", err)
		} else {
			a.Lower = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case AgeRestrictionFieldUpper: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Upper : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Upper amendment value failed to deserialize : %s", err)
		} else {
			a.Upper = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	}

	return nil, fmt.Errorf("Unknown AgeRestriction amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two AgeRestrictions and returns
// amendment data.
func (a *AgeRestrictionField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *AgeRestrictionField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Lower uint32
	fip = append(ofip, AgeRestrictionFieldLower)
	if a.Lower != newValue.Lower {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Lower)); err != nil {
			return nil, errors.Wrap(err, "Lower")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Upper uint32
	fip = append(ofip, AgeRestrictionFieldUpper)
	if a.Upper != newValue.Upper {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Upper)); err != nil {
			return nil, errors.Wrap(err, "Upper")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// CurrencyValueField Permission / Amendment Field Indices
const (
	CurrencyValueFieldValue        = uint32(1)
	CurrencyValueFieldCurrencyCode = uint32(2)
	CurrencyValueFieldPrecision    = uint32(3)
)

// ApplyAmendment updates a CurrencyValueField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *CurrencyValueField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty CurrencyValue amendment field index path")
	}

	switch fip[0] {
	case CurrencyValueFieldValue: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Value : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Value amendment value failed to deserialize : %s", err)
		} else {
			a.Value = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case CurrencyValueFieldCurrencyCode: // string

		if CurrenciesData(a.CurrencyCode) == nil {
			return nil, fmt.Errorf("Currencies resource value not defined : %v", a.CurrencyCode)
		}
		a.CurrencyCode = string(data)
		return permissions.SubPermissions(fip, operation, false)
	case CurrencyValueFieldPrecision: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	}

	return nil, fmt.Errorf("Unknown CurrencyValue amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two CurrencyValues and returns
// amendment data.
func (a *CurrencyValueField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *CurrencyValueField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Value uint64
	fip = append(ofip, CurrencyValueFieldValue)
	if a.Value != newValue.Value {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Value)); err != nil {
			return nil, errors.Wrap(err, "Value")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// CurrencyCode string
	fip = append(ofip, CurrencyValueFieldCurrencyCode)
	if a.CurrencyCode != newValue.CurrencyCode {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.CurrencyCode),
		})
	}

	// Precision uint32
	fip = append(ofip, CurrencyValueFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// RateField Permission / Amendment Field Indices
const (
	RateFieldPrecision = uint32(1)
	RateFieldValue     = uint32(2)
)

// ApplyAmendment updates a RateField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *RateField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Rate amendment field index path")
	}

	switch fip[0] {
	case RateFieldPrecision: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Precision : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Precision amendment value failed to deserialize : %s", err)
		} else {
			a.Precision = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case RateFieldValue: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Value : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Value amendment value failed to deserialize : %s", err)
		} else {
			a.Value = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	}

	return nil, fmt.Errorf("Unknown Rate amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Rates and returns
// amendment data.
func (a *RateField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *RateField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Precision uint32
	fip = append(ofip, RateFieldPrecision)
	if a.Precision != newValue.Precision {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Precision)); err != nil {
			return nil, errors.Wrap(err, "Precision")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// Value uint64
	fip = append(ofip, RateFieldValue)
	if a.Value != newValue.Value {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Value)); err != nil {
			return nil, errors.Wrap(err, "Value")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// PeriodField Permission / Amendment Field Indices
const (
	PeriodFieldUnit      = uint32(1)
	PeriodFieldUnitCount = uint32(2)
	PeriodFieldEndUnit   = uint32(3)
	PeriodFieldEndIndex  = uint32(4)
	PeriodFieldEndCount  = uint32(5)
	PeriodFieldEndTime   = uint32(6)
)

// ApplyAmendment updates a PeriodField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *PeriodField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Period amendment field index path")
	}

	switch fip[0] {
	case PeriodFieldUnit: // string

		if TimeUnitData(a.Unit) == nil {
			return nil, fmt.Errorf("TimeUnit resource value not defined : %v", a.Unit)
		}
		a.Unit = string(data)
		return permissions.SubPermissions(fip, operation, false)
	case PeriodFieldUnitCount: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for UnitCount : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("UnitCount amendment value failed to deserialize : %s", err)
		} else {
			a.UnitCount = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case PeriodFieldEndUnit: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for EndUnit : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("EndUnit amendment value failed to deserialize : %s", err)
		} else {
			a.EndUnit = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case PeriodFieldEndIndex: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for EndIndex : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("EndIndex amendment value failed to deserialize : %s", err)
		} else {
			a.EndIndex = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case PeriodFieldEndCount: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for EndCount : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("EndCount amendment value failed to deserialize : %s", err)
		} else {
			a.EndCount = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case PeriodFieldEndTime: // TimeField

		subPermissions, err := permissions.SubPermissions(fip, operation, false)
		if err != nil {
			return nil, errors.Wrap(err, "sub permissions")
		}

		return a.EndTime.ApplyAmendment(fip[1:], operation, data, subPermissions)
	}

	return nil, fmt.Errorf("Unknown Period amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Periods and returns
// amendment data.
func (a *PeriodField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *PeriodField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Unit string
	fip = append(ofip, PeriodFieldUnit)
	if a.Unit != newValue.Unit {
		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: []byte(newValue.Unit),
		})
	}

	// UnitCount uint64
	fip = append(ofip, PeriodFieldUnitCount)
	if a.UnitCount != newValue.UnitCount {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.UnitCount)); err != nil {
			return nil, errors.Wrap(err, "UnitCount")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// EndUnit uint64
	fip = append(ofip, PeriodFieldEndUnit)
	if a.EndUnit != newValue.EndUnit {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.EndUnit)); err != nil {
			return nil, errors.Wrap(err, "EndUnit")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// EndIndex uint64
	fip = append(ofip, PeriodFieldEndIndex)
	if a.EndIndex != newValue.EndIndex {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.EndIndex)); err != nil {
			return nil, errors.Wrap(err, "EndIndex")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// EndCount uint64
	fip = append(ofip, PeriodFieldEndCount)
	if a.EndCount != newValue.EndCount {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.EndCount)); err != nil {
			return nil, errors.Wrap(err, "EndCount")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// EndTime TimeField
	fip = append(ofip, PeriodFieldEndTime)

	EndTimeAmendments, err := a.EndTime.CreateAmendments(fip, newValue.EndTime)
	if err != nil {
		return nil, errors.Wrap(err, "EndTime")
	}
	result = append(result, EndTimeAmendments...)

	return result, nil
}

// TimeField Permission / Amendment Field Indices
const (
	TimeFieldTime             = uint32(1)
	TimeFieldTimeZoneOffset   = uint32(2)
	TimeFieldTimeZonePositive = uint32(3)
)

// ApplyAmendment updates a TimeField based on amendment data.
// Note: This does not check permissions or data validity. This does check data format.
// fip must have at least one value.
func (a *TimeField) ApplyAmendment(fip permissions.FieldIndexPath, operation uint32,
	data []byte, permissions permissions.Permissions) (permissions.Permissions, error) {

	if len(fip) == 0 {
		return nil, errors.New("Empty Time amendment field index path")
	}

	switch fip[0] {
	case TimeFieldTime: // uint64

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for Time : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("Time amendment value failed to deserialize : %s", err)
		} else {
			a.Time = uint64(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case TimeFieldTimeZoneOffset: // uint32

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TimeZoneOffset : %v", fip)
		}
		buf := bytes.NewBuffer(data)
		if value, err := bitcoin.ReadBase128VarInt(buf); err != nil {
			return nil, fmt.Errorf("TimeZoneOffset amendment value failed to deserialize : %s", err)
		} else {
			a.TimeZoneOffset = uint32(value)
		}
		return permissions.SubPermissions(fip, operation, false)
	case TimeFieldTimeZonePositive: // bool

		if len(fip) > 1 {
			return nil, fmt.Errorf("Amendment field index path too deep for TimeZonePositive : %v", fip)
		}
		if len(data) != 1 {
			return nil, fmt.Errorf("TimeZonePositive amendment value is wrong size : %d", len(data))
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.LittleEndian, &a.TimeZonePositive); err != nil {
			return nil, fmt.Errorf("TimeZonePositive amendment value failed to deserialize : %s", err)
		}
		return permissions.SubPermissions(fip, operation, false)
	}

	return nil, fmt.Errorf("Unknown Time amendment field index : %v", fip)
}

// CreateAmendments determines the differences between two Times and returns
// amendment data.
func (a *TimeField) CreateAmendments(fip permissions.FieldIndexPath,
	newValue *TimeField) ([]*internal.Amendment, error) {

	if a.Equal(newValue) {
		return nil, nil
	}

	var result []*internal.Amendment
	ofip := fip.Copy() // save original to be appended for each field

	// Time uint64
	fip = append(ofip, TimeFieldTime)
	if a.Time != newValue.Time {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.Time)); err != nil {
			return nil, errors.Wrap(err, "Time")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// TimeZoneOffset uint32
	fip = append(ofip, TimeFieldTimeZoneOffset)
	if a.TimeZoneOffset != newValue.TimeZoneOffset {
		var buf bytes.Buffer
		if err := bitcoin.WriteBase128VarInt(&buf, uint64(newValue.TimeZoneOffset)); err != nil {
			return nil, errors.Wrap(err, "TimeZoneOffset")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	// TimeZonePositive bool
	fip = append(ofip, TimeFieldTimeZonePositive)
	if a.TimeZonePositive != newValue.TimeZonePositive {
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, newValue.TimeZonePositive); err != nil {
			return nil, errors.Wrap(err, "TimeZonePositive")
		}

		result = append(result, &internal.Amendment{
			FIP:  fip,
			Data: buf.Bytes(),
		})
	}

	return result, nil
}

// CreatePayloadAmendments deserializes asset payloads and create amendments for them.
func CreatePayloadAmendments(fip permissions.FieldIndexPath,
	assetType, payload, newPayload []byte) ([]*internal.Amendment, error) {

	current, err := Deserialize(assetType, payload)
	if err != nil {
		return nil, errors.Wrap(err, "deserialize payload")
	}

	new, err := Deserialize(assetType, newPayload)
	if err != nil {
		return nil, errors.Wrap(err, "deserialize payload")
	}

	var result []*internal.Amendment
	switch c := current.(type) {
	case *Membership:
		result, err = c.CreateAmendments(fip, new.(*Membership))

	case *Currency:
		result, err = c.CreateAmendments(fip, new.(*Currency))

	case *ShareCommon:
		result, err = c.CreateAmendments(fip, new.(*ShareCommon))

	case *BondFixedRate:
		result, err = c.CreateAmendments(fip, new.(*BondFixedRate))

	case *Coupon:
		result, err = c.CreateAmendments(fip, new.(*Coupon))

	case *LoyaltyPoints:
		result, err = c.CreateAmendments(fip, new.(*LoyaltyPoints))

	case *TicketAdmission:
		result, err = c.CreateAmendments(fip, new.(*TicketAdmission))

	case *CasinoChip:
		result, err = c.CreateAmendments(fip, new.(*CasinoChip))

	}

	return result, nil
}
