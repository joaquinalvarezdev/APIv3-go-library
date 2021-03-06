// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// RequestContactExportReader is a Reader for the RequestContactExport structure.
type RequestContactExportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestContactExportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewRequestContactExportAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRequestContactExportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRequestContactExportAccepted creates a RequestContactExportAccepted with default headers values
func NewRequestContactExportAccepted() *RequestContactExportAccepted {
	return &RequestContactExportAccepted{}
}

/*RequestContactExportAccepted handles this case with default header values.

Contact export request has been accepted
*/
type RequestContactExportAccepted struct {
	Payload *models.CreatedProcessID
}

func (o *RequestContactExportAccepted) Error() string {
	return fmt.Sprintf("[POST /contacts/export][%d] requestContactExportAccepted  %+v", 202, o.Payload)
}

func (o *RequestContactExportAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatedProcessID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestContactExportBadRequest creates a RequestContactExportBadRequest with default headers values
func NewRequestContactExportBadRequest() *RequestContactExportBadRequest {
	return &RequestContactExportBadRequest{}
}

/*RequestContactExportBadRequest handles this case with default header values.

bad request
*/
type RequestContactExportBadRequest struct {
	Payload *models.ErrorModel
}

func (o *RequestContactExportBadRequest) Error() string {
	return fmt.Sprintf("[POST /contacts/export][%d] requestContactExportBadRequest  %+v", 400, o.Payload)
}

func (o *RequestContactExportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
