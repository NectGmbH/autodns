// Code generated by go-swagger; DO NOT EDIT.

package contact_tasks

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
)

// NewContactDeleteParams creates a new ContactDeleteParams object
// with the default values initialized.
func NewContactDeleteParams() *ContactDeleteParams {
	var ()
	return &ContactDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContactDeleteParamsWithTimeout creates a new ContactDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContactDeleteParamsWithTimeout(timeout time.Duration) *ContactDeleteParams {
	var ()
	return &ContactDeleteParams{

		timeout: timeout,
	}
}

// NewContactDeleteParamsWithContext creates a new ContactDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewContactDeleteParamsWithContext(ctx context.Context) *ContactDeleteParams {
	var ()
	return &ContactDeleteParams{

		Context: ctx,
	}
}

// NewContactDeleteParamsWithHTTPClient creates a new ContactDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContactDeleteParamsWithHTTPClient(client *http.Client) *ContactDeleteParams {
	var ()
	return &ContactDeleteParams{
		HTTPClient: client,
	}
}

/*ContactDeleteParams contains all the parameters to send to the API endpoint
for the contact delete operation typically these are written to a http.Request
*/
type ContactDeleteParams struct {

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
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the contact delete params
func (o *ContactDeleteParams) WithTimeout(timeout time.Duration) *ContactDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact delete params
func (o *ContactDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact delete params
func (o *ContactDeleteParams) WithContext(ctx context.Context) *ContactDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact delete params
func (o *ContactDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact delete params
func (o *ContactDeleteParams) WithHTTPClient(client *http.Client) *ContactDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact delete params
func (o *ContactDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the contact delete params
func (o *ContactDeleteParams) WithXDomainrobotContext(xDomainrobotContext *string) *ContactDeleteParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the contact delete params
func (o *ContactDeleteParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the contact delete params
func (o *ContactDeleteParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *ContactDeleteParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the contact delete params
func (o *ContactDeleteParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the contact delete params
func (o *ContactDeleteParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *ContactDeleteParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the contact delete params
func (o *ContactDeleteParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the contact delete params
func (o *ContactDeleteParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *ContactDeleteParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the contact delete params
func (o *ContactDeleteParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the contact delete params
func (o *ContactDeleteParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *ContactDeleteParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the contact delete params
func (o *ContactDeleteParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the contact delete params
func (o *ContactDeleteParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *ContactDeleteParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the contact delete params
func (o *ContactDeleteParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the contact delete params
func (o *ContactDeleteParams) WithXDomainrobotWS(xDomainrobotWS *string) *ContactDeleteParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the contact delete params
func (o *ContactDeleteParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithID adds the id to the contact delete params
func (o *ContactDeleteParams) WithID(id int32) *ContactDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contact delete params
func (o *ContactDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContactDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
