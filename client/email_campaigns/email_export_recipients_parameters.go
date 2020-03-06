// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sendinblue/APIv3-go-library/models"
)

// NewEmailExportRecipientsParams creates a new EmailExportRecipientsParams object
// with the default values initialized.
func NewEmailExportRecipientsParams() *EmailExportRecipientsParams {
	var ()
	return &EmailExportRecipientsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEmailExportRecipientsParamsWithTimeout creates a new EmailExportRecipientsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEmailExportRecipientsParamsWithTimeout(timeout time.Duration) *EmailExportRecipientsParams {
	var ()
	return &EmailExportRecipientsParams{

		timeout: timeout,
	}
}

// NewEmailExportRecipientsParamsWithContext creates a new EmailExportRecipientsParams object
// with the default values initialized, and the ability to set a context for a request
func NewEmailExportRecipientsParamsWithContext(ctx context.Context) *EmailExportRecipientsParams {
	var ()
	return &EmailExportRecipientsParams{

		Context: ctx,
	}
}

// NewEmailExportRecipientsParamsWithHTTPClient creates a new EmailExportRecipientsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEmailExportRecipientsParamsWithHTTPClient(client *http.Client) *EmailExportRecipientsParams {
	var ()
	return &EmailExportRecipientsParams{
		HTTPClient: client,
	}
}

/*EmailExportRecipientsParams contains all the parameters to send to the API endpoint
for the email export recipients operation typically these are written to a http.Request
*/
type EmailExportRecipientsParams struct {

	/*CampaignID
	  Id of the campaign

	*/
	CampaignID int64
	/*RecipientExport
	  Values to send for a recipient export request

	*/
	RecipientExport *models.EmailExportRecipients

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the email export recipients params
func (o *EmailExportRecipientsParams) WithTimeout(timeout time.Duration) *EmailExportRecipientsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the email export recipients params
func (o *EmailExportRecipientsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the email export recipients params
func (o *EmailExportRecipientsParams) WithContext(ctx context.Context) *EmailExportRecipientsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the email export recipients params
func (o *EmailExportRecipientsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the email export recipients params
func (o *EmailExportRecipientsParams) WithHTTPClient(client *http.Client) *EmailExportRecipientsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the email export recipients params
func (o *EmailExportRecipientsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the email export recipients params
func (o *EmailExportRecipientsParams) WithCampaignID(campaignID int64) *EmailExportRecipientsParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the email export recipients params
func (o *EmailExportRecipientsParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithRecipientExport adds the recipientExport to the email export recipients params
func (o *EmailExportRecipientsParams) WithRecipientExport(recipientExport *models.EmailExportRecipients) *EmailExportRecipientsParams {
	o.SetRecipientExport(recipientExport)
	return o
}

// SetRecipientExport adds the recipientExport to the email export recipients params
func (o *EmailExportRecipientsParams) SetRecipientExport(recipientExport *models.EmailExportRecipients) {
	o.RecipientExport = recipientExport
}

// WriteToRequest writes these params to a swagger request
func (o *EmailExportRecipientsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.RecipientExport != nil {
		if err := r.SetBodyParam(o.RecipientExport); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
