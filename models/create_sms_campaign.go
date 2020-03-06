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

// CreateSmsCampaign create sms campaign
// swagger:model createSmsCampaign
type CreateSmsCampaign struct {

	// Content of the message. The maximum characters used per SMS is 160, if used more than that, it will be counted as more than one SMS
	// Required: true
	Content *string `json:"content"`

	// Name of the campaign
	// Required: true
	Name *string `json:"name"`

	// recipients
	Recipients *CreateSmsCampaignRecipients `json:"recipients,omitempty"`

	// UTC date-time on which the campaign has to run (YYYY-MM-DDTHH:mm:ss.SSSZ). Prefer to pass your timezone in date-time format for accurate result.
	// Format: date-time
	ScheduledAt strfmt.DateTime `json:"scheduledAt,omitempty"`

	// Name of the sender. The number of characters is limited to 11
	// Required: true
	// Max Length: 11
	Sender *string `json:"sender"`
}

// Validate validates this create sms campaign
func (m *CreateSmsCampaign) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSmsCampaign) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	return nil
}

func (m *CreateSmsCampaign) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateSmsCampaign) validateRecipients(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipients) { // not required
		return nil
	}

	if m.Recipients != nil {
		if err := m.Recipients.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipients")
			}
			return err
		}
	}

	return nil
}

func (m *CreateSmsCampaign) validateScheduledAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledAt) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduledAt", "body", "date-time", m.ScheduledAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateSmsCampaign) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	if err := validate.MaxLength("sender", "body", string(*m.Sender), 11); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateSmsCampaign) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSmsCampaign) UnmarshalBinary(b []byte) error {
	var res CreateSmsCampaign
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CreateSmsCampaignRecipients create sms campaign recipients
// swagger:model CreateSmsCampaignRecipients
type CreateSmsCampaignRecipients struct {

	// List ids which have to be excluded from a campaign
	ExclusionListIds []int64 `json:"exclusionListIds"`

	// Lists Ids to send the campaign to. REQUIRED if scheduledAt is not empty
	// Required: true
	ListIds []int64 `json:"listIds"`
}

// Validate validates this create sms campaign recipients
func (m *CreateSmsCampaignRecipients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateListIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateSmsCampaignRecipients) validateListIds(formats strfmt.Registry) error {

	if err := validate.Required("recipients"+"."+"listIds", "body", m.ListIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateSmsCampaignRecipients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateSmsCampaignRecipients) UnmarshalBinary(b []byte) error {
	var res CreateSmsCampaignRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
