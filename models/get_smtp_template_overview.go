// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetSMTPTemplateOverview get Smtp template overview
// swagger:model getSmtpTemplateOverview
type GetSMTPTemplateOverview struct {

	// Creation UTC date-time of the template (YYYY-MM-DDTHH:mm:ss.SSSZ)
	// Required: true
	CreatedAt *strfmt.DateTime `json:"createdAt"`

	// HTML content of the template
	// Required: true
	HTMLContent *string `json:"htmlContent"`

	// ID of the template
	// Required: true
	ID *int64 `json:"id"`

	// Status of template (true=active, false=inactive)
	// Required: true
	IsActive *bool `json:"isActive"`

	// Last modification UTC date-time of the template (YYYY-MM-DDTHH:mm:ss.SSSZ)
	// Required: true
	ModifiedAt *strfmt.DateTime `json:"modifiedAt"`

	// Name of the template
	// Required: true
	Name *string `json:"name"`

	// Email defined as the "Reply to" for the template
	// Required: true
	ReplyTo *strfmt.Email `json:"replyTo"`

	// sender
	// Required: true
	Sender *GetSMTPTemplateOverviewSender `json:"sender"`

	// Subject of the template
	// Required: true
	Subject *string `json:"subject"`

	// Tag of the template
	// Required: true
	Tag *string `json:"tag"`

	// Status of test sending for the template (true=test email has been sent, false=test email has not been sent)
	// Required: true
	TestSent *bool `json:"testSent"`

	// Customisation of the "to" field for the template
	// Required: true
	ToField *string `json:"toField"`
}

// Validate validates this get Smtp template overview
func (m *GetSMTPTemplateOverview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHTMLContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIsActive(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModifiedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTag(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTestSent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateToField(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSMTPTemplateOverview) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateHTMLContent(formats strfmt.Registry) error {

	if err := validate.Required("htmlContent", "body", m.HTMLContent); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateIsActive(formats strfmt.Registry) error {

	if err := validate.Required("isActive", "body", m.IsActive); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateModifiedAt(formats strfmt.Registry) error {

	if err := validate.Required("modifiedAt", "body", m.ModifiedAt); err != nil {
		return err
	}

	if err := validate.FormatOf("modifiedAt", "body", "date-time", m.ModifiedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateReplyTo(formats strfmt.Registry) error {

	if err := validate.Required("replyTo", "body", m.ReplyTo); err != nil {
		return err
	}

	if err := validate.FormatOf("replyTo", "body", "email", m.ReplyTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	if m.Sender != nil {

		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateTag(formats strfmt.Registry) error {

	if err := validate.Required("tag", "body", m.Tag); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateTestSent(formats strfmt.Registry) error {

	if err := validate.Required("testSent", "body", m.TestSent); err != nil {
		return err
	}

	return nil
}

func (m *GetSMTPTemplateOverview) validateToField(formats strfmt.Registry) error {

	if err := validate.Required("toField", "body", m.ToField); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSMTPTemplateOverview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSMTPTemplateOverview) UnmarshalBinary(b []byte) error {
	var res GetSMTPTemplateOverview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
