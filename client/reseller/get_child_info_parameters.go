// Code generated by go-swagger; DO NOT EDIT.

package reseller

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
)

// NewGetChildInfoParams creates a new GetChildInfoParams object
// with the default values initialized.
func NewGetChildInfoParams() *GetChildInfoParams {
	var ()
	return &GetChildInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChildInfoParamsWithTimeout creates a new GetChildInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChildInfoParamsWithTimeout(timeout time.Duration) *GetChildInfoParams {
	var ()
	return &GetChildInfoParams{

		timeout: timeout,
	}
}

// NewGetChildInfoParamsWithContext creates a new GetChildInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChildInfoParamsWithContext(ctx context.Context) *GetChildInfoParams {
	var ()
	return &GetChildInfoParams{

		Context: ctx,
	}
}

// NewGetChildInfoParamsWithHTTPClient creates a new GetChildInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChildInfoParamsWithHTTPClient(client *http.Client) *GetChildInfoParams {
	var ()
	return &GetChildInfoParams{
		HTTPClient: client,
	}
}

/*GetChildInfoParams contains all the parameters to send to the API endpoint
for the get child info operation typically these are written to a http.Request
*/
type GetChildInfoParams struct {

	/*ChildAuthKey
	  auth key of reseller's child

	*/
	ChildAuthKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get child info params
func (o *GetChildInfoParams) WithTimeout(timeout time.Duration) *GetChildInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get child info params
func (o *GetChildInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get child info params
func (o *GetChildInfoParams) WithContext(ctx context.Context) *GetChildInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get child info params
func (o *GetChildInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get child info params
func (o *GetChildInfoParams) WithHTTPClient(client *http.Client) *GetChildInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get child info params
func (o *GetChildInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChildAuthKey adds the childAuthKey to the get child info params
func (o *GetChildInfoParams) WithChildAuthKey(childAuthKey string) *GetChildInfoParams {
	o.SetChildAuthKey(childAuthKey)
	return o
}

// SetChildAuthKey adds the childAuthKey to the get child info params
func (o *GetChildInfoParams) SetChildAuthKey(childAuthKey string) {
	o.ChildAuthKey = childAuthKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetChildInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param childAuthKey
	if err := r.SetPathParam("childAuthKey", o.ChildAuthKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
