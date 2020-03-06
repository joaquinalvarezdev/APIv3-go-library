// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// DeleteWebhookReader is a Reader for the DeleteWebhook structure.
type DeleteWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteWebhookNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteWebhookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteWebhookNoContent creates a DeleteWebhookNoContent with default headers values
func NewDeleteWebhookNoContent() *DeleteWebhookNoContent {
	return &DeleteWebhookNoContent{}
}

/*DeleteWebhookNoContent handles this case with default header values.

Webhook deleted
*/
type DeleteWebhookNoContent struct {
}

func (o *DeleteWebhookNoContent) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{webhookId}][%d] deleteWebhookNoContent ", 204)
}

func (o *DeleteWebhookNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteWebhookBadRequest creates a DeleteWebhookBadRequest with default headers values
func NewDeleteWebhookBadRequest() *DeleteWebhookBadRequest {
	return &DeleteWebhookBadRequest{}
}

/*DeleteWebhookBadRequest handles this case with default header values.

bad request
*/
type DeleteWebhookBadRequest struct {
	Payload *models.ErrorModel
}

func (o *DeleteWebhookBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{webhookId}][%d] deleteWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteWebhookBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *DeleteWebhookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebhookNotFound creates a DeleteWebhookNotFound with default headers values
func NewDeleteWebhookNotFound() *DeleteWebhookNotFound {
	return &DeleteWebhookNotFound{}
}

/*DeleteWebhookNotFound handles this case with default header values.

Webhook ID not found
*/
type DeleteWebhookNotFound struct {
	Payload *models.ErrorModel
}

func (o *DeleteWebhookNotFound) Error() string {
	return fmt.Sprintf("[DELETE /webhooks/{webhookId}][%d] deleteWebhookNotFound  %+v", 404, o.Payload)
}

func (o *DeleteWebhookNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *DeleteWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
