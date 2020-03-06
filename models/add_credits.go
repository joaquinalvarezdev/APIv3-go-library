// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddCredits add credits
// swagger:model addCredits
type AddCredits struct {

	// Required if sms credits are empty. Email credits to be added to the child account
	Email int64 `json:"email,omitempty"`

	// Required if email credits are empty. SMS credits to be added to the child account
	Sms int64 `json:"sms,omitempty"`
}

// Validate validates this add credits
func (m *AddCredits) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddCredits) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddCredits) UnmarshalBinary(b []byte) error {
	var res AddCredits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
