// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// GetEmailCampaignReader is a Reader for the GetEmailCampaign structure.
type GetEmailCampaignReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailCampaignReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmailCampaignOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEmailCampaignBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEmailCampaignNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmailCampaignOK creates a GetEmailCampaignOK with default headers values
func NewGetEmailCampaignOK() *GetEmailCampaignOK {
	return &GetEmailCampaignOK{}
}

/*GetEmailCampaignOK handles this case with default header values.

Email campaign informations
*/
type GetEmailCampaignOK struct {
	Payload *models.GetEmailCampaign
}

func (o *GetEmailCampaignOK) Error() string {
	return fmt.Sprintf("[GET /emailCampaigns/{campaignId}][%d] getEmailCampaignOK  %+v", 200, o.Payload)
}

func (o *GetEmailCampaignOK) GetPayload() *models.GetEmailCampaign {
	return o.Payload
}

func (o *GetEmailCampaignOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetEmailCampaign)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailCampaignBadRequest creates a GetEmailCampaignBadRequest with default headers values
func NewGetEmailCampaignBadRequest() *GetEmailCampaignBadRequest {
	return &GetEmailCampaignBadRequest{}
}

/*GetEmailCampaignBadRequest handles this case with default header values.

bad request
*/
type GetEmailCampaignBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetEmailCampaignBadRequest) Error() string {
	return fmt.Sprintf("[GET /emailCampaigns/{campaignId}][%d] getEmailCampaignBadRequest  %+v", 400, o.Payload)
}

func (o *GetEmailCampaignBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetEmailCampaignBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailCampaignNotFound creates a GetEmailCampaignNotFound with default headers values
func NewGetEmailCampaignNotFound() *GetEmailCampaignNotFound {
	return &GetEmailCampaignNotFound{}
}

/*GetEmailCampaignNotFound handles this case with default header values.

Campaign ID not found
*/
type GetEmailCampaignNotFound struct {
	Payload *models.ErrorModel
}

func (o *GetEmailCampaignNotFound) Error() string {
	return fmt.Sprintf("[GET /emailCampaigns/{campaignId}][%d] getEmailCampaignNotFound  %+v", 404, o.Payload)
}

func (o *GetEmailCampaignNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetEmailCampaignNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
