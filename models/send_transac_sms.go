// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SendTransacSms send transac sms
// swagger:model sendTransacSms
type SendTransacSms struct {

	// Content of the message. If more than 160 characters long, will be sent as multiple text messages
	// Required: true
	Content *string `json:"content"`

	// Mobile number to send SMS with the country code
	// Required: true
	Recipient *string `json:"recipient"`

	// Name of the sender. Only alphanumeric characters. No more than 11 characters
	// Required: true
	// Max Length: 11
	Sender *string `json:"sender"`

	// Tag of the message
	Tag string `json:"tag,omitempty"`

	// Type of the SMS. Marketing SMS messages are those sent typically with marketing content. Transactional SMS messages are sent to individuals and are triggered in response to some action, such as a sign-up, purchase, etc.
	// Enum: [transactional marketing]
	Type *string `json:"type,omitempty"`

	// Webhook to call for each event triggered by the message (delivered etc.)
	WebURL string `json:"webUrl,omitempty"`
}

// Validate validates this send transac sms
func (m *SendTransacSms) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendTransacSms) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *SendTransacSms) validateRecipient(formats strfmt.Registry) error {

	if err := validate.Required("recipient", "body", m.Recipient); err != nil {
		return err
	}

	return nil
}

func (m *SendTransacSms) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	if err := validate.MaxLength("sender", "body", string(*m.Sender), 11); err != nil {
		return err
	}

	return nil
}

var sendTransacSmsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["transactional","marketing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sendTransacSmsTypeTypePropEnum = append(sendTransacSmsTypeTypePropEnum, v)
	}
}

const (

	// SendTransacSmsTypeTransactional captures enum value "transactional"
	SendTransacSmsTypeTransactional string = "transactional"

	// SendTransacSmsTypeMarketing captures enum value "marketing"
	SendTransacSmsTypeMarketing string = "marketing"
)

// prop value enum
func (m *SendTransacSms) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, sendTransacSmsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SendTransacSms) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendTransacSms) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendTransacSms) UnmarshalBinary(b []byte) error {
	var res SendTransacSms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
