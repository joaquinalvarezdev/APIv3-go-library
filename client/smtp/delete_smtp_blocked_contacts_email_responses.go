// Code generated by go-swagger; DO NOT EDIT.

package smtp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// DeleteSMTPBlockedContactsEmailReader is a Reader for the DeleteSMTPBlockedContactsEmail structure.
type DeleteSMTPBlockedContactsEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSMTPBlockedContactsEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteSMTPBlockedContactsEmailNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSMTPBlockedContactsEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSMTPBlockedContactsEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSMTPBlockedContactsEmailNoContent creates a DeleteSMTPBlockedContactsEmailNoContent with default headers values
func NewDeleteSMTPBlockedContactsEmailNoContent() *DeleteSMTPBlockedContactsEmailNoContent {
	return &DeleteSMTPBlockedContactsEmailNoContent{}
}

/*DeleteSMTPBlockedContactsEmailNoContent handles this case with default header values.

Contact unblocked
*/
type DeleteSMTPBlockedContactsEmailNoContent struct {
}

func (o *DeleteSMTPBlockedContactsEmailNoContent) Error() string {
	return fmt.Sprintf("[DELETE /smtp/blockedContacts/{email}][%d] deleteSmtpBlockedContactsEmailNoContent ", 204)
}

func (o *DeleteSMTPBlockedContactsEmailNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSMTPBlockedContactsEmailBadRequest creates a DeleteSMTPBlockedContactsEmailBadRequest with default headers values
func NewDeleteSMTPBlockedContactsEmailBadRequest() *DeleteSMTPBlockedContactsEmailBadRequest {
	return &DeleteSMTPBlockedContactsEmailBadRequest{}
}

/*DeleteSMTPBlockedContactsEmailBadRequest handles this case with default header values.

bad request
*/
type DeleteSMTPBlockedContactsEmailBadRequest struct {
	Payload *models.ErrorModel
}

func (o *DeleteSMTPBlockedContactsEmailBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /smtp/blockedContacts/{email}][%d] deleteSmtpBlockedContactsEmailBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSMTPBlockedContactsEmailBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *DeleteSMTPBlockedContactsEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSMTPBlockedContactsEmailNotFound creates a DeleteSMTPBlockedContactsEmailNotFound with default headers values
func NewDeleteSMTPBlockedContactsEmailNotFound() *DeleteSMTPBlockedContactsEmailNotFound {
	return &DeleteSMTPBlockedContactsEmailNotFound{}
}

/*DeleteSMTPBlockedContactsEmailNotFound handles this case with default header values.

Contact email not found
*/
type DeleteSMTPBlockedContactsEmailNotFound struct {
	Payload *models.ErrorModel
}

func (o *DeleteSMTPBlockedContactsEmailNotFound) Error() string {
	return fmt.Sprintf("[DELETE /smtp/blockedContacts/{email}][%d] deleteSmtpBlockedContactsEmailNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSMTPBlockedContactsEmailNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *DeleteSMTPBlockedContactsEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
