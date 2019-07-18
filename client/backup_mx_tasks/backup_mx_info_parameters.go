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

	strfmt "github.com/go-openapi/strfmt"
)

// NewBackupMxInfoParams creates a new BackupMxInfoParams object
// with the default values initialized.
func NewBackupMxInfoParams() *BackupMxInfoParams {
	var ()
	return &BackupMxInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBackupMxInfoParamsWithTimeout creates a new BackupMxInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBackupMxInfoParamsWithTimeout(timeout time.Duration) *BackupMxInfoParams {
	var ()
	return &BackupMxInfoParams{

		timeout: timeout,
	}
}

// NewBackupMxInfoParamsWithContext creates a new BackupMxInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewBackupMxInfoParamsWithContext(ctx context.Context) *BackupMxInfoParams {
	var ()
	return &BackupMxInfoParams{

		Context: ctx,
	}
}

// NewBackupMxInfoParamsWithHTTPClient creates a new BackupMxInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBackupMxInfoParamsWithHTTPClient(client *http.Client) *BackupMxInfoParams {
	var ()
	return &BackupMxInfoParams{
		HTTPClient: client,
	}
}

/*BackupMxInfoParams contains all the parameters to send to the API endpoint
for the backup mx info operation typically these are written to a http.Request
*/
type BackupMxInfoParams struct {

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
	/*Domain*/
	Domain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the backup mx info params
func (o *BackupMxInfoParams) WithTimeout(timeout time.Duration) *BackupMxInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup mx info params
func (o *BackupMxInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup mx info params
func (o *BackupMxInfoParams) WithContext(ctx context.Context) *BackupMxInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup mx info params
func (o *BackupMxInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup mx info params
func (o *BackupMxInfoParams) WithHTTPClient(client *http.Client) *BackupMxInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup mx info params
func (o *BackupMxInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the backup mx info params
func (o *BackupMxInfoParams) WithXDomainrobotContext(xDomainrobotContext *string) *BackupMxInfoParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the backup mx info params
func (o *BackupMxInfoParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the backup mx info params
func (o *BackupMxInfoParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *BackupMxInfoParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the backup mx info params
func (o *BackupMxInfoParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the backup mx info params
func (o *BackupMxInfoParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *BackupMxInfoParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the backup mx info params
func (o *BackupMxInfoParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the backup mx info params
func (o *BackupMxInfoParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *BackupMxInfoParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the backup mx info params
func (o *BackupMxInfoParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the backup mx info params
func (o *BackupMxInfoParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *BackupMxInfoParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the backup mx info params
func (o *BackupMxInfoParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the backup mx info params
func (o *BackupMxInfoParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *BackupMxInfoParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the backup mx info params
func (o *BackupMxInfoParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the backup mx info params
func (o *BackupMxInfoParams) WithXDomainrobotWS(xDomainrobotWS *string) *BackupMxInfoParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the backup mx info params
func (o *BackupMxInfoParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithDomain adds the domain to the backup mx info params
func (o *BackupMxInfoParams) WithDomain(domain string) *BackupMxInfoParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the backup mx info params
func (o *BackupMxInfoParams) SetDomain(domain string) {
	o.Domain = domain
}

// WriteToRequest writes these params to a swagger request
func (o *BackupMxInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}