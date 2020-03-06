// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateContact update contact
// swagger:model updateContact
type UpdateContact struct {

	// Pass the set of attributes to be updated. These attributes must be present in your account. To update existing email address of a contact with the new one please pass EMAIL in attribtes. For example, `{ "EMAIL":"newemail@domain.com", "FNAME":"Ellie", "LNAME":"Roger"}`. Keep in mind transactional attributes can be updated the same way as normal attributes.
	Attributes interface{} `json:"attributes,omitempty"`

	// Set/unset this field to blacklist/allow the contact for emails (emailBlacklisted = true)
	EmailBlacklisted bool `json:"emailBlacklisted,omitempty"`

	// Ids of the lists to add the contact to
	ListIds []int64 `json:"listIds"`

	// Set/unset this field to blacklist/allow the contact for SMS (smsBlacklisted = true)
	SmsBlacklisted bool `json:"smsBlacklisted,omitempty"`

	// transactional email forbidden sender for contact. Use only for email Contact
	SMTPBlacklistSender []strfmt.Email `json:"smtpBlacklistSender"`

	// Ids of the lists to remove the contact from
	UnlinkListIds []int64 `json:"unlinkListIds"`
}

// Validate validates this update contact
func (m *UpdateContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSMTPBlacklistSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateContact) validateSMTPBlacklistSender(formats strfmt.Registry) error {

	if swag.IsZero(m.SMTPBlacklistSender) { // not required
		return nil
	}

	for i := 0; i < len(m.SMTPBlacklistSender); i++ {

		if err := validate.FormatOf("smtpBlacklistSender"+"."+strconv.Itoa(i), "body", "email", m.SMTPBlacklistSender[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateContact) UnmarshalBinary(b []byte) error {
	var res UpdateContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
