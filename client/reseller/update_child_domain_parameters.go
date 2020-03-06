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

// NewUpdateChildDomainParams creates a new UpdateChildDomainParams object
// with the default values initialized.
func NewUpdateChildDomainParams() *UpdateChildDomainParams {
	var ()
	return &UpdateChildDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateChildDomainParamsWithTimeout creates a new UpdateChildDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateChildDomainParamsWithTimeout(timeout time.Duration) *UpdateChildDomainParams {
	var ()
	return &UpdateChildDomainParams{

		timeout: timeout,
	}
}

// NewUpdateChildDomainParamsWithContext creates a new UpdateChildDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateChildDomainParamsWithContext(ctx context.Context) *UpdateChildDomainParams {
	var ()
	return &UpdateChildDomainParams{

		Context: ctx,
	}
}

// NewUpdateChildDomainParamsWithHTTPClient creates a new UpdateChildDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateChildDomainParamsWithHTTPClient(client *http.Client) *UpdateChildDomainParams {
	var ()
	return &UpdateChildDomainParams{
		HTTPClient: client,
	}
}

/*UpdateChildDomainParams contains all the parameters to send to the API endpoint
for the update child domain operation typically these are written to a http.Request
*/
type UpdateChildDomainParams struct {

	/*ChildAuthKey
	  auth key of reseller's child

	*/
	ChildAuthKey string
	/*DomainName
	  Pass the existing domain that needs to be updated

	*/
	DomainName string
	/*UpdateChildDomain
	  value to update for sender domain

	*/
	UpdateChildDomain *models.UpdateChildDomain

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update child domain params
func (o *UpdateChildDomainParams) WithTimeout(timeout time.Duration) *UpdateChildDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update child domain params
func (o *UpdateChildDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update child domain params
func (o *UpdateChildDomainParams) WithContext(ctx context.Context) *UpdateChildDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update child domain params
func (o *UpdateChildDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update child domain params
func (o *UpdateChildDomainParams) WithHTTPClient(client *http.Client) *UpdateChildDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update child domain params
func (o *UpdateChildDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChildAuthKey adds the childAuthKey to the update child domain params
func (o *UpdateChildDomainParams) WithChildAuthKey(childAuthKey string) *UpdateChildDomainParams {
	o.SetChildAuthKey(childAuthKey)
	return o
}

// SetChildAuthKey adds the childAuthKey to the update child domain params
func (o *UpdateChildDomainParams) SetChildAuthKey(childAuthKey string) {
	o.ChildAuthKey = childAuthKey
}

// WithDomainName adds the domainName to the update child domain params
func (o *UpdateChildDomainParams) WithDomainName(domainName string) *UpdateChildDomainParams {
	o.SetDomainName(domainName)
	return o
}

// SetDomainName adds the domainName to the update child domain params
func (o *UpdateChildDomainParams) SetDomainName(domainName string) {
	o.DomainName = domainName
}

// WithUpdateChildDomain adds the updateChildDomain to the update child domain params
func (o *UpdateChildDomainParams) WithUpdateChildDomain(updateChildDomain *models.UpdateChildDomain) *UpdateChildDomainParams {
	o.SetUpdateChildDomain(updateChildDomain)
	return o
}

// SetUpdateChildDomain adds the updateChildDomain to the update child domain params
func (o *UpdateChildDomainParams) SetUpdateChildDomain(updateChildDomain *models.UpdateChildDomain) {
	o.UpdateChildDomain = updateChildDomain
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateChildDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param childAuthKey
	if err := r.SetPathParam("childAuthKey", o.ChildAuthKey); err != nil {
		return err
	}

	// path param domainName
	if err := r.SetPathParam("domainName", o.DomainName); err != nil {
		return err
	}

	if o.UpdateChildDomain != nil {
		if err := r.SetBodyParam(o.UpdateChildDomain); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
