// Code generated by go-swagger; DO NOT EDIT.

package backup_mx_tasks

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

	models "github.com/NectGmbH/autodns/models"
)

// NewBackupMxCreateParams creates a new BackupMxCreateParams object
// with the default values initialized.
func NewBackupMxCreateParams() *BackupMxCreateParams {
	var ()
	return &BackupMxCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBackupMxCreateParamsWithTimeout creates a new BackupMxCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBackupMxCreateParamsWithTimeout(timeout time.Duration) *BackupMxCreateParams {
	var ()
	return &BackupMxCreateParams{

		timeout: timeout,
	}
}

// NewBackupMxCreateParamsWithContext creates a new BackupMxCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewBackupMxCreateParamsWithContext(ctx context.Context) *BackupMxCreateParams {
	var ()
	return &BackupMxCreateParams{

		Context: ctx,
	}
}

// NewBackupMxCreateParamsWithHTTPClient creates a new BackupMxCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBackupMxCreateParamsWithHTTPClient(client *http.Client) *BackupMxCreateParams {
	var ()
	return &BackupMxCreateParams{
		HTTPClient: client,
	}
}

/*BackupMxCreateParams contains all the parameters to send to the API endpoint
for the backup mx create operation typically these are written to a http.Request
*/
type BackupMxCreateParams struct {

	/*XDomainrobotContext*/
	XDomainrobotContext *string
	/*XDomainrobotDemo*/
	XDomainrobotDemo *string
	/*XDomainrobotOwnerContext*/
	XDomainrobotOwnerContext *string
	/*XDomainrobotOwnerUser*/
	XDomainrobotOwnerUser *string
	/*XDomainrobotPIN*/
	XDomainrobotPIN *string
	/*XDomainrobotSessionID*/
	XDomainrobotSessionID *string
	/*XDomainrobotWS*/
	XDomainrobotWS *string
	/*Body
	  backupMx

	*/
	Body *models.BackupMx
	/*Keys
	  If the dns should be provisioned if available.

	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the backup mx create params
func (o *BackupMxCreateParams) WithTimeout(timeout time.Duration) *BackupMxCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup mx create params
func (o *BackupMxCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup mx create params
func (o *BackupMxCreateParams) WithContext(ctx context.Context) *BackupMxCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup mx create params
func (o *BackupMxCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup mx create params
func (o *BackupMxCreateParams) WithHTTPClient(client *http.Client) *BackupMxCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup mx create params
func (o *BackupMxCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the backup mx create params
func (o *BackupMxCreateParams) WithXDomainrobotContext(xDomainrobotContext *string) *BackupMxCreateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the backup mx create params
func (o *BackupMxCreateParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the backup mx create params
func (o *BackupMxCreateParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *BackupMxCreateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the backup mx create params
func (o *BackupMxCreateParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the backup mx create params
func (o *BackupMxCreateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *BackupMxCreateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the backup mx create params
func (o *BackupMxCreateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the backup mx create params
func (o *BackupMxCreateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *BackupMxCreateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the backup mx create params
func (o *BackupMxCreateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the backup mx create params
func (o *BackupMxCreateParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *BackupMxCreateParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the backup mx create params
func (o *BackupMxCreateParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the backup mx create params
func (o *BackupMxCreateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *BackupMxCreateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the backup mx create params
func (o *BackupMxCreateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the backup mx create params
func (o *BackupMxCreateParams) WithXDomainrobotWS(xDomainrobotWS *string) *BackupMxCreateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the backup mx create params
func (o *BackupMxCreateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the backup mx create params
func (o *BackupMxCreateParams) WithBody(body *models.BackupMx) *BackupMxCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the backup mx create params
func (o *BackupMxCreateParams) SetBody(body *models.BackupMx) {
	o.Body = body
}

// WithKeys adds the keys to the backup mx create params
func (o *BackupMxCreateParams) WithKeys(keys []string) *BackupMxCreateParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the backup mx create params
func (o *BackupMxCreateParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *BackupMxCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XDomainrobotContext != nil {

		// header param X-Domainrobot-Context
		if err := r.SetHeaderParam("X-Domainrobot-Context", *o.XDomainrobotContext); err != nil {
			return err
		}

	}

	if o.XDomainrobotDemo != nil {

		// header param X-Domainrobot-Demo
		if err := r.SetHeaderParam("X-Domainrobot-Demo", *o.XDomainrobotDemo); err != nil {
			return err
		}

	}

	if o.XDomainrobotOwnerContext != nil {

		// header param X-Domainrobot-Owner-Context
		if err := r.SetHeaderParam("X-Domainrobot-Owner-Context", *o.XDomainrobotOwnerContext); err != nil {
			return err
		}

	}

	if o.XDomainrobotOwnerUser != nil {

		// header param X-Domainrobot-Owner-User
		if err := r.SetHeaderParam("X-Domainrobot-Owner-User", *o.XDomainrobotOwnerUser); err != nil {
			return err
		}

	}

	if o.XDomainrobotPIN != nil {

		// header param X-Domainrobot-PIN
		if err := r.SetHeaderParam("X-Domainrobot-PIN", *o.XDomainrobotPIN); err != nil {
			return err
		}

	}

	if o.XDomainrobotSessionID != nil {

		// header param X-Domainrobot-SessionId
		if err := r.SetHeaderParam("X-Domainrobot-SessionId", *o.XDomainrobotSessionID); err != nil {
			return err
		}

	}

	if o.XDomainrobotWS != nil {

		// header param X-Domainrobot-WS
		if err := r.SetHeaderParam("X-Domainrobot-WS", *o.XDomainrobotWS); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
