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

// SendEmail send email
// swagger:model sendEmail
type SendEmail struct {

	// attachment
	Attachment SendEmailAttachment `json:"attachment"`

	// Absolute url of the attachment (no local file). Extension allowed: xlsx, xls, ods, docx, docm, doc, csv, pdf, txt, gif, jpg, jpeg, png, tif, tiff, rtf, bmp, cgm, css, shtml, html, htm, zip, xml, ppt, pptx, tar, ez, ics, mobi, msg, pub and eps
	AttachmentURL string `json:"attachmentUrl,omitempty"`

	// attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// Email addresses of the recipients in bcc
	EmailBcc []strfmt.Email `json:"emailBcc"`

	// Email addresses of the recipients in cc
	EmailCc []strfmt.Email `json:"emailCc"`

	// Email addresses of the recipients
	// Required: true
	EmailTo []strfmt.Email `json:"emailTo"`

	// headers
	Headers map[string]string `json:"headers,omitempty"`

	// Email on which campaign recipients will be able to reply to
	ReplyTo strfmt.Email `json:"replyTo,omitempty"`
}

// Validate validates this send email
func (m *SendEmail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailBcc(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmailCc(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmailTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendEmail) validateEmailBcc(formats strfmt.Registry) error {

	if swag.IsZero(m.EmailBcc) { // not required
		return nil
	}

	return nil
}

func (m *SendEmail) validateEmailCc(formats strfmt.Registry) error {

	if swag.IsZero(m.EmailCc) { // not required
		return nil
	}

	return nil
}

func (m *SendEmail) validateEmailTo(formats strfmt.Registry) error {

	if err := validate.Required("emailTo", "body", m.EmailTo); err != nil {
		return err
	}

	return nil
}

func (m *SendEmail) validateReplyTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyTo) { // not required
		return nil
	}

	if err := validate.FormatOf("replyTo", "body", "email", m.ReplyTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendEmail) UnmarshalBinary(b []byte) error {
	var res SendEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
