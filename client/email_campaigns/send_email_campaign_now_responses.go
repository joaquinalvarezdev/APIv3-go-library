// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// SendEmailCampaignNowReader is a Reader for the SendEmailCampaignNow structure.
type SendEmailCampaignNowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendEmailCampaignNowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewSendEmailCampaignNowNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendEmailCampaignNowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 402:
		result := NewSendEmailCampaignNowPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendEmailCampaignNowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendEmailCampaignNowNoContent creates a SendEmailCampaignNowNoContent with default headers values
func NewSendEmailCampaignNowNoContent() *SendEmailCampaignNowNoContent {
	return &SendEmailCampaignNowNoContent{}
}

/*SendEmailCampaignNowNoContent handles this case with default header values.

Email campaign has been scheduled
*/
type SendEmailCampaignNowNoContent struct {
}

func (o *SendEmailCampaignNowNoContent) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/sendNow][%d] sendEmailCampaignNowNoContent ", 204)
}

func (o *SendEmailCampaignNowNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendEmailCampaignNowBadRequest creates a SendEmailCampaignNowBadRequest with default headers values
func NewSendEmailCampaignNowBadRequest() *SendEmailCampaignNowBadRequest {
	return &SendEmailCampaignNowBadRequest{}
}

/*SendEmailCampaignNowBadRequest handles this case with default header values.

Campaign could not be sent
*/
type SendEmailCampaignNowBadRequest struct {
	Payload *models.ErrorModel
}

func (o *SendEmailCampaignNowBadRequest) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/sendNow][%d] sendEmailCampaignNowBadRequest  %+v", 400, o.Payload)
}

func (o *SendEmailCampaignNowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendEmailCampaignNowPaymentRequired creates a SendEmailCampaignNowPaymentRequired with default headers values
func NewSendEmailCampaignNowPaymentRequired() *SendEmailCampaignNowPaymentRequired {
	return &SendEmailCampaignNowPaymentRequired{}
}

/*SendEmailCampaignNowPaymentRequired handles this case with default header values.

You don't have enough credit to send your campaign. Please update your plan
*/
type SendEmailCampaignNowPaymentRequired struct {
	Payload *models.ErrorModel
}

func (o *SendEmailCampaignNowPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/sendNow][%d] sendEmailCampaignNowPaymentRequired  %+v", 402, o.Payload)
}

func (o *SendEmailCampaignNowPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendEmailCampaignNowNotFound creates a SendEmailCampaignNowNotFound with default headers values
func NewSendEmailCampaignNowNotFound() *SendEmailCampaignNowNotFound {
	return &SendEmailCampaignNowNotFound{}
}

/*SendEmailCampaignNowNotFound handles this case with default header values.

Campaign ID not found
*/
type SendEmailCampaignNowNotFound struct {
	Payload *models.ErrorModel
}

func (o *SendEmailCampaignNowNotFound) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/sendNow][%d] sendEmailCampaignNowNotFound  %+v", 404, o.Payload)
}

func (o *SendEmailCampaignNowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
