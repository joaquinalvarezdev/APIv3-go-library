// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetSmsEventReport get sms event report
// swagger:model getSmsEventReport
type GetSmsEventReport struct {

	// events
	Events []*GetSmsEventReportEventsItems0 `json:"events"`
}

// Validate validates this get sms event report
func (m *GetSmsEventReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSmsEventReport) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSmsEventReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSmsEventReport) UnmarshalBinary(b []byte) error {
	var res GetSmsEventReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GetSmsEventReportEventsItems0 get sms event report events items0
// swagger:model GetSmsEventReportEventsItems0
type GetSmsEventReportEventsItems0 struct {

	// UTC date-time on which the event has been generated
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// Event which occurred
	// Required: true
	// Enum: [bounces hardBounces softBounces delivered sent accepted unsubscription replies blocked]
	Event *string `json:"event"`

	// Message ID which generated the event
	// Required: true
	MessageID *string `json:"messageId"`

	// Phone number which has generated the event
	// Required: true
	PhoneNumber *string `json:"phoneNumber"`

	// Reason of bounce (only available if the event is hardbounce or softbounce)
	Reason string `json:"reason,omitempty"`

	// reply
	Reply string `json:"reply,omitempty"`

	// Tag of the SMS which generated the event
	Tag string `json:"tag,omitempty"`
}

// Validate validates this get sms event report events items0
func (m *GetSmsEventReportEventsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSmsEventReportEventsItems0) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

var getSmsEventReportEventsItems0TypeEventPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bounces","hardBounces","softBounces","delivered","sent","accepted","unsubscription","replies","blocked"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getSmsEventReportEventsItems0TypeEventPropEnum = append(getSmsEventReportEventsItems0TypeEventPropEnum, v)
	}
}

const (

	// GetSmsEventReportEventsItems0EventBounces captures enum value "bounces"
	GetSmsEventReportEventsItems0EventBounces string = "bounces"

	// GetSmsEventReportEventsItems0EventHardBounces captures enum value "hardBounces"
	GetSmsEventReportEventsItems0EventHardBounces string = "hardBounces"

	// GetSmsEventReportEventsItems0EventSoftBounces captures enum value "softBounces"
	GetSmsEventReportEventsItems0EventSoftBounces string = "softBounces"

	// GetSmsEventReportEventsItems0EventDelivered captures enum value "delivered"
	GetSmsEventReportEventsItems0EventDelivered string = "delivered"

	// GetSmsEventReportEventsItems0EventSent captures enum value "sent"
	GetSmsEventReportEventsItems0EventSent string = "sent"

	// GetSmsEventReportEventsItems0EventAccepted captures enum value "accepted"
	GetSmsEventReportEventsItems0EventAccepted string = "accepted"

	// GetSmsEventReportEventsItems0EventUnsubscription captures enum value "unsubscription"
	GetSmsEventReportEventsItems0EventUnsubscription string = "unsubscription"

	// GetSmsEventReportEventsItems0EventReplies captures enum value "replies"
	GetSmsEventReportEventsItems0EventReplies string = "replies"

	// GetSmsEventReportEventsItems0EventBlocked captures enum value "blocked"
	GetSmsEventReportEventsItems0EventBlocked string = "blocked"
)

// prop value enum
func (m *GetSmsEventReportEventsItems0) validateEventEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getSmsEventReportEventsItems0TypeEventPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetSmsEventReportEventsItems0) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event); err != nil {
		return err
	}

	// value enum
	if err := m.validateEventEnum("event", "body", *m.Event); err != nil {
		return err
	}

	return nil
}

func (m *GetSmsEventReportEventsItems0) validateMessageID(formats strfmt.Registry) error {

	if err := validate.Required("messageId", "body", m.MessageID); err != nil {
		return err
	}

	return nil
}

func (m *GetSmsEventReportEventsItems0) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phoneNumber", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSmsEventReportEventsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSmsEventReportEventsItems0) UnmarshalBinary(b []byte) error {
	var res GetSmsEventReportEventsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
