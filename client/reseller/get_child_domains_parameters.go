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

// NewGetChildDomainsParams creates a new GetChildDomainsParams object
// with the default values initialized.
func NewGetChildDomainsParams() *GetChildDomainsParams {
	var ()
	return &GetChildDomainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChildDomainsParamsWithTimeout creates a new GetChildDomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChildDomainsParamsWithTimeout(timeout time.Duration) *GetChildDomainsParams {
	var ()
	return &GetChildDomainsParams{

		timeout: timeout,
	}
}

// NewGetChildDomainsParamsWithContext creates a new GetChildDomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChildDomainsParamsWithContext(ctx context.Context) *GetChildDomainsParams {
	var ()
	return &GetChildDomainsParams{

		Context: ctx,
	}
}

// NewGetChildDomainsParamsWithHTTPClient creates a new GetChildDomainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChildDomainsParamsWithHTTPClient(client *http.Client) *GetChildDomainsParams {
	var ()
	return &GetChildDomainsParams{
		HTTPClient: client,
	}
}

/*GetChildDomainsParams contains all the parameters to send to the API endpoint
for the get child domains operation typically these are written to a http.Request
*/
type GetChildDomainsParams struct {

	/*ChildAuthKey
	  auth key of reseller's child

	*/
	ChildAuthKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get child domains params
func (o *GetChildDomainsParams) WithTimeout(timeout time.Duration) *GetChildDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get child domains params
func (o *GetChildDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get child domains params
func (o *GetChildDomainsParams) WithContext(ctx context.Context) *GetChildDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get child domains params
func (o *GetChildDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get child domains params
func (o *GetChildDomainsParams) WithHTTPClient(client *http.Client) *GetChildDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get child domains params
func (o *GetChildDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChildAuthKey adds the childAuthKey to the get child domains params
func (o *GetChildDomainsParams) WithChildAuthKey(childAuthKey string) *GetChildDomainsParams {
	o.SetChildAuthKey(childAuthKey)
	return o
}

// SetChildAuthKey adds the childAuthKey to the get child domains params
func (o *GetChildDomainsParams) SetChildAuthKey(childAuthKey string) {
	o.ChildAuthKey = childAuthKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetChildDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
