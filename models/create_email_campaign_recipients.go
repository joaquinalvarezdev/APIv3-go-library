// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CreateEmailCampaignRecipients create email campaign recipients
// swagger:model createEmailCampaignRecipients
type CreateEmailCampaignRecipients struct {

	// List ids to exclude from the campaign
	ExclusionListIds []int64 `json:"exclusionListIds"`

	// Mandatory if scheduledAt is not empty. List Ids to send the campaign to
	ListIds []int64 `json:"listIds"`
}

// Validate validates this create email campaign recipients
func (m *CreateEmailCampaignRecipients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExclusionListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateListIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateEmailCampaignRecipients) validateExclusionListIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ExclusionListIds) { // not required
		return nil
	}

	return nil
}

func (m *CreateEmailCampaignRecipients) validateListIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ListIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateEmailCampaignRecipients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateEmailCampaignRecipients) UnmarshalBinary(b []byte) error {
	var res CreateEmailCampaignRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
