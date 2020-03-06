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

// NewUpdateChildAccountStatusParams creates a new UpdateChildAccountStatusParams object
// with the default values initialized.
func NewUpdateChildAccountStatusParams() *UpdateChildAccountStatusParams {
	var ()
	return &UpdateChildAccountStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateChildAccountStatusParamsWithTimeout creates a new UpdateChildAccountStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateChildAccountStatusParamsWithTimeout(timeout time.Duration) *UpdateChildAccountStatusParams {
	var ()
	return &UpdateChildAccountStatusParams{

		timeout: timeout,
	}
}

// NewUpdateChildAccountStatusParamsWithContext creates a new UpdateChildAccountStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateChildAccountStatusParamsWithContext(ctx context.Context) *UpdateChildAccountStatusParams {
	var ()
	return &UpdateChildAccountStatusParams{

		Context: ctx,
	}
}

// NewUpdateChildAccountStatusParamsWithHTTPClient creates a new UpdateChildAccountStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateChildAccountStatusParamsWithHTTPClient(client *http.Client) *UpdateChildAccountStatusParams {
	var ()
	return &UpdateChildAccountStatusParams{
		HTTPClient: client,
	}
}

/*UpdateChildAccountStatusParams contains all the parameters to send to the API endpoint
for the update child account status operation typically these are written to a http.Request
*/
type UpdateChildAccountStatusParams struct {

	/*ChildAuthKey
	  auth key of reseller's child

	*/
	ChildAuthKey string
	/*UpdateChildAccountStatus
	  values to update in child account status

	*/
	UpdateChildAccountStatus *models.UpdateChildAccountStatus

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update child account status params
func (o *UpdateChildAccountStatusParams) WithTimeout(timeout time.Duration) *UpdateChildAccountStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update child account status params
func (o *UpdateChildAccountStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update child account status params
func (o *UpdateChildAccountStatusParams) WithContext(ctx context.Context) *UpdateChildAccountStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update child account status params
func (o *UpdateChildAccountStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update child account status params
func (o *UpdateChildAccountStatusParams) WithHTTPClient(client *http.Client) *UpdateChildAccountStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update child account status params
func (o *UpdateChildAccountStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChildAuthKey adds the childAuthKey to the update child account status params
func (o *UpdateChildAccountStatusParams) WithChildAuthKey(childAuthKey string) *UpdateChildAccountStatusParams {
	o.SetChildAuthKey(childAuthKey)
	return o
}

// SetChildAuthKey adds the childAuthKey to the update child account status params
func (o *UpdateChildAccountStatusParams) SetChildAuthKey(childAuthKey string) {
	o.ChildAuthKey = childAuthKey
}

// WithUpdateChildAccountStatus adds the updateChildAccountStatus to the update child account status params
func (o *UpdateChildAccountStatusParams) WithUpdateChildAccountStatus(updateChildAccountStatus *models.UpdateChildAccountStatus) *UpdateChildAccountStatusParams {
	o.SetUpdateChildAccountStatus(updateChildAccountStatus)
	return o
}

// SetUpdateChildAccountStatus adds the updateChildAccountStatus to the update child account status params
func (o *UpdateChildAccountStatusParams) SetUpdateChildAccountStatus(updateChildAccountStatus *models.UpdateChildAccountStatus) {
	o.UpdateChildAccountStatus = updateChildAccountStatus
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateChildAccountStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param childAuthKey
	if err := r.SetPathParam("childAuthKey", o.ChildAuthKey); err != nil {
		return err
	}

	if o.UpdateChildAccountStatus != nil {
		if err := r.SetBodyParam(o.UpdateChildAccountStatus); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
