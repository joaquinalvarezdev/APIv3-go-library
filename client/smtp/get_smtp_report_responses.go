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

// GetSMTPReportReader is a Reader for the GetSMTPReport structure.
type GetSMTPReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSMTPReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSMTPReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSMTPReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSMTPReportOK creates a GetSMTPReportOK with default headers values
func NewGetSMTPReportOK() *GetSMTPReportOK {
	return &GetSMTPReportOK{}
}

/*GetSMTPReportOK handles this case with default header values.

Aggregated report informations
*/
type GetSMTPReportOK struct {
	Payload *models.GetReports
}

func (o *GetSMTPReportOK) Error() string {
	return fmt.Sprintf("[GET /smtp/statistics/reports][%d] getSmtpReportOK  %+v", 200, o.Payload)
}

func (o *GetSMTPReportOK) GetPayload() *models.GetReports {
	return o.Payload
}

func (o *GetSMTPReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetReports)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSMTPReportBadRequest creates a GetSMTPReportBadRequest with default headers values
func NewGetSMTPReportBadRequest() *GetSMTPReportBadRequest {
	return &GetSMTPReportBadRequest{}
}

/*GetSMTPReportBadRequest handles this case with default header values.

bad request
*/
type GetSMTPReportBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetSMTPReportBadRequest) Error() string {
	return fmt.Sprintf("[GET /smtp/statistics/reports][%d] getSmtpReportBadRequest  %+v", 400, o.Payload)
}

func (o *GetSMTPReportBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetSMTPReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
