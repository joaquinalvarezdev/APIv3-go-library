// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateEmailCampaign create email campaign
// swagger:model createEmailCampaign
type CreateEmailCampaign struct {

	// Absolute url of the attachment (no local file). Extension allowed: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff, rtf, bmp, cgm, css, shtml, html, htm, zip, xml, ppt, pptx, tar, ez, ics, mobi, msg, pub and eps
	AttachmentURL string `json:"attachmentUrl,omitempty"`

	// Footer of the email campaign
	Footer string `json:"footer,omitempty"`

	// Header of the email campaign
	Header string `json:"header,omitempty"`

	// Mandatory if htmlUrl is empty. Body of the message (HTML)
	HTMLContent string `json:"htmlContent,omitempty"`

	// Mandatory if htmlContent is empty. Url to the message (HTML)
	HTMLURL string `json:"htmlUrl,omitempty"`

	// Use true to embedded the images in your email. Final size of the email should be less than 4MB. Campaigns with embedded images can not be sent to more than 5000 contacts
	InlineImageActivation *bool `json:"inlineImageActivation,omitempty"`

	// Use true to enable the mirror link
	MirrorActive bool `json:"mirrorActive,omitempty"`

	// Name of the campaign
	// Required: true
	Name *string `json:"name"`

	// recipients
	Recipients *CreateEmailCampaignRecipients `json:"recipients,omitempty"`

	// For trigger campagins use false to make sure a contact receives the same campaign only once
	Recurring *bool `json:"recurring,omitempty"`

	// Email on which the campaign recipients will be able to reply to
	ReplyTo strfmt.Email `json:"replyTo,omitempty"`

	// Sending UTC date-time (YYYY-MM-DDTHH:mm:ss.SSSZ). Prefer to pass your timezone in date-time format for accurate result.
	ScheduledAt strfmt.DateTime `json:"scheduledAt,omitempty"`

	// sender
	// Required: true
	Sender *CreateEmailCampaignSender `json:"sender"`

	// Subject of the campaign
	// Required: true
	Subject *string `json:"subject"`

	// Tag of the campaign
	Tag string `json:"tag,omitempty"`

	// To personalize the «To» Field, e.g. if you want to include the first name and last name of your recipient, use [FNAME] [LNAME]. These attributes must already exist in your contact database
	ToField string `json:"toField,omitempty"`

	// Type of the campaign
	// Required: true
	Type *string `json:"type"`

	// Customize the utm_campaign value. If this field is empty, the campaign name will be used. Only alphanumeric characters and spaces are allowed
	UtmCampaign string `json:"utmCampaign,omitempty"`
}

// Validate validates this create email campaign
func (m *CreateEmailCampaign) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScheduledAt(formats); err != nil {
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

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateEmailCampaign) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmailCampaign) validateRecipients(formats strfmt.Registry) error {

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

func (m *CreateEmailCampaign) validateReplyTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyTo) { // not required
		return nil
	}

	if err := validate.FormatOf("replyTo", "body", "email", m.ReplyTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmailCampaign) validateScheduledAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledAt) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduledAt", "body", "date-time", m.ScheduledAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateEmailCampaign) validateSender(formats strfmt.Registry) error {

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

func (m *CreateEmailCampaign) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

var createEmailCampaignTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["classic","trigger"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createEmailCampaignTypeTypePropEnum = append(createEmailCampaignTypeTypePropEnum, v)
	}
}

const (
	// CreateEmailCampaignTypeClassic captures enum value "classic"
	CreateEmailCampaignTypeClassic string = "classic"
	// CreateEmailCampaignTypeTrigger captures enum value "trigger"
	CreateEmailCampaignTypeTrigger string = "trigger"
)

// prop value enum
func (m *CreateEmailCampaign) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createEmailCampaignTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateEmailCampaign) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateEmailCampaign) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateEmailCampaign) UnmarshalBinary(b []byte) error {
	var res CreateEmailCampaign
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
