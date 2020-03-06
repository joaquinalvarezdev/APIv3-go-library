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

// GetTransacSmsReport get transac sms report
// swagger:model getTransacSmsReport
type GetTransacSmsReport struct {

	// reports
	Reports []*GetTransacSmsReportReportsItems0 `json:"reports"`
}

// Validate validates this get transac sms report
func (m *GetTransacSmsReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReports(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTransacSmsReport) validateReports(formats strfmt.Registry) error {

	if swag.IsZero(m.Reports) { // not required
		return nil
	}

	for i := 0; i < len(m.Reports); i++ {
		if swag.IsZero(m.Reports[i]) { // not required
			continue
		}

		if m.Reports[i] != nil {
			if err := m.Reports[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacSmsReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacSmsReport) UnmarshalBinary(b []byte) error {
	var res GetTransacSmsReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetTransacSmsReportReportsItems0 get transac sms report reports items0
// swagger:model GetTransacSmsReportReportsItems0
type GetTransacSmsReportReportsItems0 struct {

	// Number of accepted for the date
	// Required: true
	Accepted *int64 `json:"accepted"`

	// Number of blocked contact for the date
	// Required: true
	Blocked *int64 `json:"blocked"`

	// Date for which statistics are retrieved
	// Required: true
	// Format: date
	Date *strfmt.Date `json:"date"`

	// Number of delivered SMS for the date
	// Required: true
	Delivered *int64 `json:"delivered"`

	// Number of hardbounces for the date
	// Required: true
	HardBounces *int64 `json:"hardBounces"`

	// Number of rejected for the date
	// Required: true
	Rejected *int64 `json:"rejected"`

	// Number of answered SMS for the date
	// Required: true
	Replied *int64 `json:"replied"`

	// Number of requests for the date
	// Required: true
	Requests *int64 `json:"requests"`

	// Number of softbounces for the date
	// Required: true
	SoftBounces *int64 `json:"softBounces"`

	// Number of unsubscription for the date
	// Required: true
	Unsubscribed *int64 `json:"unsubscribed"`
}

// Validate validates this get transac sms report reports items0
func (m *GetTransacSmsReportReportsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccepted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelivered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHardBounces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejected(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplied(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftBounces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnsubscribed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateAccepted(formats strfmt.Registry) error {

	if err := validate.Required("accepted", "body", m.Accepted); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateBlocked(formats strfmt.Registry) error {

	if err := validate.Required("blocked", "body", m.Blocked); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateDelivered(formats strfmt.Registry) error {

	if err := validate.Required("delivered", "body", m.Delivered); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateHardBounces(formats strfmt.Registry) error {

	if err := validate.Required("hardBounces", "body", m.HardBounces); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateRejected(formats strfmt.Registry) error {

	if err := validate.Required("rejected", "body", m.Rejected); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateReplied(formats strfmt.Registry) error {

	if err := validate.Required("replied", "body", m.Replied); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateRequests(formats strfmt.Registry) error {

	if err := validate.Required("requests", "body", m.Requests); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateSoftBounces(formats strfmt.Registry) error {

	if err := validate.Required("softBounces", "body", m.SoftBounces); err != nil {
		return err
	}

	return nil
}

func (m *GetTransacSmsReportReportsItems0) validateUnsubscribed(formats strfmt.Registry) error {

	if err := validate.Required("unsubscribed", "body", m.Unsubscribed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTransacSmsReportReportsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTransacSmsReportReportsItems0) UnmarshalBinary(b []byte) error {
	var res GetTransacSmsReportReportsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
