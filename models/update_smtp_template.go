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

// UpdateSMTPTemplate update Smtp template
// swagger:model updateSmtpTemplate
type UpdateSMTPTemplate struct {

	// Absolute url of the attachment (no local file). Extension allowed: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff, rtf, bmp, cgm, css, shtml, html, htm, zip, xml, ppt, pptx, tar, ez, ics, mobi, msg, pub and eps
	AttachmentURL string `json:"attachmentUrl,omitempty"`

	// Required if htmlUrl is empty. Body of the message (HTML must have more than 10 characters)
	HTMLContent string `json:"htmlContent,omitempty"`

	// Required if htmlContent is empty. URL to the body of the email (HTML)
	HTMLURL string `json:"htmlUrl,omitempty"`

	// Status of the template. isActive = false means template is inactive, isActive = true means template is active
	IsActive bool `json:"isActive,omitempty"`

	// Email on which campaign recipients will be able to reply to
	// Format: email
	ReplyTo strfmt.Email `json:"replyTo,omitempty"`

	// sender
	Sender *UpdateSMTPTemplateSender `json:"sender,omitempty"`

	// Subject of the email
	Subject string `json:"subject,omitempty"`

	// Tag of the template
	Tag string `json:"tag,omitempty"`

	// Name of the template
	TemplateName string `json:"templateName,omitempty"`

	// To personalize the «To» Field. If you want to include the first name and last name of your recipient, add {FNAME} {LNAME}. These contact attributes must already exist in your SendinBlue account. If input parameter 'params' used please use {{contact.FNAME}} {{contact.LNAME}} for personalization
	ToField string `json:"toField,omitempty"`
}

// Validate validates this update Smtp template
func (m *UpdateSMTPTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplyTo(formats); err != nil {
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

func (m *UpdateSMTPTemplate) validateReplyTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyTo) { // not required
		return nil
	}

	if err := validate.FormatOf("replyTo", "body", "email", m.ReplyTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSMTPTemplate) validateSender(formats strfmt.Registry) error {

	if swag.IsZero(m.Sender) { // not required
		return nil
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

// MarshalBinary interface implementation
func (m *UpdateSMTPTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSMTPTemplate) UnmarshalBinary(b []byte) error {
	var res UpdateSMTPTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// UpdateSMTPTemplateSender Sender details including id or email and name (optional). Only one of either Sender's email or Sender's ID shall be passed in one request at a time. For example `{"name":"xyz", "email":"example@abc.com"}` , `{"name":"xyz", "id":123}`
// swagger:model UpdateSMTPTemplateSender
type UpdateSMTPTemplateSender struct {

	// Email of the sender
	// Format: email
	Email strfmt.Email `json:"email,omitempty"`

	// Select the sender for the template on the basis of sender id. In order to select a sender with specific pool of IP’s, dedicated ip users shall pass id (instead of email).
	ID int64 `json:"id,omitempty"`

	// Name of the sender
	Name string `json:"name,omitempty"`
}

// Validate validates this update SMTP template sender
func (m *UpdateSMTPTemplateSender) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateSMTPTemplateSender) validateEmail(formats strfmt.Registry) error {

	if swag.IsZero(m.Email) { // not required
		return nil
	}

	if err := validate.FormatOf("sender"+"."+"email", "body", "email", m.Email.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateSMTPTemplateSender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateSMTPTemplateSender) UnmarshalBinary(b []byte) error {
	var res UpdateSMTPTemplateSender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
