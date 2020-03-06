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

// GetAggregatedSMTPReportReader is a Reader for the GetAggregatedSMTPReport structure.
type GetAggregatedSMTPReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAggregatedSMTPReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAggregatedSMTPReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAggregatedSMTPReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAggregatedSMTPReportOK creates a GetAggregatedSMTPReportOK with default headers values
func NewGetAggregatedSMTPReportOK() *GetAggregatedSMTPReportOK {
	return &GetAggregatedSMTPReportOK{}
}

/*GetAggregatedSMTPReportOK handles this case with default header values.

Aggregated report informations
*/
type GetAggregatedSMTPReportOK struct {
	Payload *models.GetAggregatedReport
}

func (o *GetAggregatedSMTPReportOK) Error() string {
	return fmt.Sprintf("[GET /smtp/statistics/aggregatedReport][%d] getAggregatedSmtpReportOK  %+v", 200, o.Payload)
}

func (o *GetAggregatedSMTPReportOK) GetPayload() *models.GetAggregatedReport {
	return o.Payload
}

func (o *GetAggregatedSMTPReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAggregatedReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAggregatedSMTPReportBadRequest creates a GetAggregatedSMTPReportBadRequest with default headers values
func NewGetAggregatedSMTPReportBadRequest() *GetAggregatedSMTPReportBadRequest {
	return &GetAggregatedSMTPReportBadRequest{}
}

/*GetAggregatedSMTPReportBadRequest handles this case with default header values.

bad request
*/
type GetAggregatedSMTPReportBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetAggregatedSMTPReportBadRequest) Error() string {
	return fmt.Sprintf("[GET /smtp/statistics/aggregatedReport][%d] getAggregatedSmtpReportBadRequest  %+v", 400, o.Payload)
}

func (o *GetAggregatedSMTPReportBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetAggregatedSMTPReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
