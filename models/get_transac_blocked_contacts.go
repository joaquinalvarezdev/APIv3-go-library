// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetTransacBlockedContacts get transac blocked contacts
// swagger:model getTransacBlockedContacts
type GetTransacBlockedContacts struct {

	// contacts
	Contacts []*GetTransacBlockedContactsContactsItems0 `json:"contacts"`

	// Count of blocked or unsubscribed contact
	Count int64 `json:"count,omitempty"`
}

// Validate validates this get transac blocked contacts
func (m *GetTransacBlockedContacts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContacts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTransacBlockedContacts) validateContacts(formats strfmt.Registry) error {

	if swag.IsZero(m.Contacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Contacts); i++ {
		if swag.IsZero(m.Contacts[i]) { // not required
			continue
		}

		if m.Contacts[i] != nil {
			if err := m.Contacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacBlockedContacts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacBlockedContacts) UnmarshalBinary(b []byte) error {
	var res GetTransacBlockedContacts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetTransacBlockedContactsContactsItems0 get transac blocked contacts contacts items0
// swagger:model GetTransacBlockedContactsContactsItems0
type GetTransacBlockedContactsContactsItems0 struct {

	// Date when the contact was blocked or unsubscribed on
	// Required: true
	// Format: date
	BlockedAt *strfmt.Date `json:"blockedAt"`

	// Email address of the blocked or unsubscribed contact
	// Required: true
	// Format: email
	Email *strfmt.Email `json:"email"`

	// reason
	// Required: true
	Reason *GetTransacBlockedContactsContactsItems0Reason `json:"reason"`

	// Sender email address of the blocked or unsubscribed contact
	// Required: true
	// Format: email
	SenderEmail *strfmt.Email `json:"senderEmail"`
}

// Validate validates this get transac blocked contacts contacts items0
func (m *GetTransacBlockedContactsContactsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSenderEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTransacBlockedContactsContactsItems0) validateBlockedAt(formats strfmt.Registry) error {

	if err := validate.Required("blockedAt", "body", m.BlockedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("blockedAt", "body", "date", m.BlockedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacBlockedContactsContactsItems0) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	if err := validate.FormatOf("email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacBlockedContactsContactsItems0) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	if m.Reason != nil {
		if err := m.Reason.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reason")
			}
			return err
		}
	}

	return nil
}

func (m *GetTransacBlockedContactsContactsItems0) validateSenderEmail(formats strfmt.Registry) error {

	if err := validate.Required("senderEmail", "body", m.SenderEmail); err != nil {
		return err
	}

	if err := validate.FormatOf("senderEmail", "body", "email", m.SenderEmail.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacBlockedContactsContactsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacBlockedContactsContactsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransacBlockedContactsContactsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetTransacBlockedContactsContactsItems0Reason Reason for blocking / unsubscribing
// swagger:model GetTransacBlockedContactsContactsItems0Reason
type GetTransacBlockedContactsContactsItems0Reason struct {

	// Reason code for blocking / unsubscribing (This code is safe for comparison)
	// Enum: [unsubscribedViaMA unsubscribedViaEmail adminBlocked unsubscribedViaApi hardBounce contactFlaggedAsSpam]
	Code string `json:"code,omitempty"`

	// Reason for blocking / unsubscribing (This string is not safe for comparison)
	Message string `json:"message,omitempty"`
}

// Validate validates this get transac blocked contacts contacts items0 reason
func (m *GetTransacBlockedContactsContactsItems0Reason) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getTransacBlockedContactsContactsItems0ReasonTypeCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["unsubscribedViaMA","unsubscribedViaEmail","adminBlocked","unsubscribedViaApi","hardBounce","contactFlaggedAsSpam"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getTransacBlockedContactsContactsItems0ReasonTypeCodePropEnum = append(getTransacBlockedContactsContactsItems0ReasonTypeCodePropEnum, v)
	}
}

const (

	// GetTransacBlockedContactsContactsItems0ReasonCodeUnsubscribedViaMA captures enum value "unsubscribedViaMA"
	GetTransacBlockedContactsContactsItems0ReasonCodeUnsubscribedViaMA string = "unsubscribedViaMA"

	// GetTransacBlockedContactsContactsItems0ReasonCodeUnsubscribedViaEmail captures enum value "unsubscribedViaEmail"
	GetTransacBlockedContactsContactsItems0ReasonCodeUnsubscribedViaEmail string = "unsubscribedViaEmail"

	// GetTransacBlockedContactsContactsItems0ReasonCodeAdminBlocked captures enum value "adminBlocked"
	GetTransacBlockedContactsContactsItems0ReasonCodeAdminBlocked string = "adminBlocked"

	// GetTransacBlockedContactsContactsItems0ReasonCodeUnsubscribedViaAPI captures enum value "unsubscribedViaApi"
	GetTransacBlockedContactsContactsItems0ReasonCodeUnsubscribedViaAPI string = "unsubscribedViaApi"

	// GetTransacBlockedContactsContactsItems0ReasonCodeHardBounce captures enum value "hardBounce"
	GetTransacBlockedContactsContactsItems0ReasonCodeHardBounce string = "hardBounce"

	// GetTransacBlockedContactsContactsItems0ReasonCodeContactFlaggedAsSpam captures enum value "contactFlaggedAsSpam"
	GetTransacBlockedContactsContactsItems0ReasonCodeContactFlaggedAsSpam string = "contactFlaggedAsSpam"
)

// prop value enum
func (m *GetTransacBlockedContactsContactsItems0Reason) validateCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getTransacBlockedContactsContactsItems0ReasonTypeCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetTransacBlockedContactsContactsItems0Reason) validateCode(formats strfmt.Registry) error {

	if swag.IsZero(m.Code) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodeEnum("reason"+"."+"code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacBlockedContactsContactsItems0Reason) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacBlockedContactsContactsItems0Reason) UnmarshalBinary(b []byte) error {
	var res GetTransacBlockedContactsContactsItems0Reason
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
