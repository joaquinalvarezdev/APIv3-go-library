// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SendSms send sms
// swagger:model sendSms
type SendSms struct {

	// message Id
	// Required: true
	MessageID *int64 `json:"messageId"`

	// reference
	// Required: true
	Reference *string `json:"reference"`

	// Remaining SMS credits of the user
	RemainingCredits float32 `json:"remainingCredits,omitempty"`

	// Count of SMS's to send multiple text messages
	SmsCount int64 `json:"smsCount,omitempty"`

	// SMS credits used per text message
	UsedCredits float32 `json:"usedCredits,omitempty"`
}

// Validate validates this send sms
func (m *SendSms) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSms) validateMessageID(formats strfmt.Registry) error {

	if err := validate.Required("messageId", "body", m.MessageID); err != nil {
		return err
	}

	return nil
}

func (m *SendSms) validateReference(formats strfmt.Registry) error {

	if err := validate.Required("reference", "body", m.Reference); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSms) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSms) UnmarshalBinary(b []byte) error {
	var res SendSms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
