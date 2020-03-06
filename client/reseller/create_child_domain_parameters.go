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

	"github.com/sendinblue/APIv3-go-library/models"
)

// NewCreateChildDomainParams creates a new CreateChildDomainParams object
// with the default values initialized.
func NewCreateChildDomainParams() *CreateChildDomainParams {
	var ()
	return &CreateChildDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateChildDomainParamsWithTimeout creates a new CreateChildDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateChildDomainParamsWithTimeout(timeout time.Duration) *CreateChildDomainParams {
	var ()
	return &CreateChildDomainParams{

		timeout: timeout,
	}
}

// NewCreateChildDomainParamsWithContext creates a new CreateChildDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateChildDomainParamsWithContext(ctx context.Context) *CreateChildDomainParams {
	var ()
	return &CreateChildDomainParams{

		Context: ctx,
	}
}

// NewCreateChildDomainParamsWithHTTPClient creates a new CreateChildDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateChildDomainParamsWithHTTPClient(client *http.Client) *CreateChildDomainParams {
	var ()
	return &CreateChildDomainParams{
		HTTPClient: client,
	}
}

/*CreateChildDomainParams contains all the parameters to send to the API endpoint
for the create child domain operation typically these are written to a http.Request
*/
type CreateChildDomainParams struct {

	/*AddChildDomain
	  Sender domain to add for a specific child account. This will not be displayed to the parent account.

	*/
	AddChildDomain *models.AddChildDomain
	/*ChildAuthKey
	  auth key of reseller's child

	*/
	ChildAuthKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create child domain params
func (o *CreateChildDomainParams) WithTimeout(timeout time.Duration) *CreateChildDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create child domain params
func (o *CreateChildDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create child domain params
func (o *CreateChildDomainParams) WithContext(ctx context.Context) *CreateChildDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create child domain params
func (o *CreateChildDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create child domain params
func (o *CreateChildDomainParams) WithHTTPClient(client *http.Client) *CreateChildDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create child domain params
func (o *CreateChildDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddChildDomain adds the addChildDomain to the create child domain params
func (o *CreateChildDomainParams) WithAddChildDomain(addChildDomain *models.AddChildDomain) *CreateChildDomainParams {
	o.SetAddChildDomain(addChildDomain)
	return o
}

// SetAddChildDomain adds the addChildDomain to the create child domain params
func (o *CreateChildDomainParams) SetAddChildDomain(addChildDomain *models.AddChildDomain) {
	o.AddChildDomain = addChildDomain
}

// WithChildAuthKey adds the childAuthKey to the create child domain params
func (o *CreateChildDomainParams) WithChildAuthKey(childAuthKey string) *CreateChildDomainParams {
	o.SetChildAuthKey(childAuthKey)
	return o
}

// SetChildAuthKey adds the childAuthKey to the create child domain params
func (o *CreateChildDomainParams) SetChildAuthKey(childAuthKey string) {
	o.ChildAuthKey = childAuthKey
}

// WriteToRequest writes these params to a swagger request
func (o *CreateChildDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AddChildDomain != nil {
		if err := r.SetBodyParam(o.AddChildDomain); err != nil {
			return err
		}
	}

	// path param childAuthKey
	if err := r.SetPathParam("childAuthKey", o.ChildAuthKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
