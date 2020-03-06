// Code generated by go-swagger; DO NOT EDIT.

package senders

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// NewCreateSenderParams creates a new CreateSenderParams object
// with the default values initialized.
func NewCreateSenderParams() *CreateSenderParams {
	var ()
	return &CreateSenderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSenderParamsWithTimeout creates a new CreateSenderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSenderParamsWithTimeout(timeout time.Duration) *CreateSenderParams {
	var ()
	return &CreateSenderParams{

		timeout: timeout,
	}
}

// NewCreateSenderParamsWithContext creates a new CreateSenderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSenderParamsWithContext(ctx context.Context) *CreateSenderParams {
	var ()
	return &CreateSenderParams{

		Context: ctx,
	}
}

// NewCreateSenderParamsWithHTTPClient creates a new CreateSenderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSenderParamsWithHTTPClient(client *http.Client) *CreateSenderParams {
	var ()
	return &CreateSenderParams{
		HTTPClient: client,
	}
}

/*CreateSenderParams contains all the parameters to send to the API endpoint
for the create sender operation typically these are written to a http.Request
*/
type CreateSenderParams struct {

	/*Sender
	  sender's name

	*/
	Sender *models.CreateSender

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create sender params
func (o *CreateSenderParams) WithTimeout(timeout time.Duration) *CreateSenderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create sender params
func (o *CreateSenderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create sender params
func (o *CreateSenderParams) WithContext(ctx context.Context) *CreateSenderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create sender params
func (o *CreateSenderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create sender params
func (o *CreateSenderParams) WithHTTPClient(client *http.Client) *CreateSenderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create sender params
func (o *CreateSenderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSender adds the sender to the create sender params
func (o *CreateSenderParams) WithSender(sender *models.CreateSender) *CreateSenderParams {
	o.SetSender(sender)
	return o
}

// SetSender adds the sender to the create sender params
func (o *CreateSenderParams) SetSender(sender *models.CreateSender) {
	o.Sender = sender
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSenderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Sender != nil {
		if err := r.SetBodyParam(o.Sender); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
