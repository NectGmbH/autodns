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

// NewContactInfoParams creates a new ContactInfoParams object
// with the default values initialized.
func NewContactInfoParams() *ContactInfoParams {
	var ()
	return &ContactInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContactInfoParamsWithTimeout creates a new ContactInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContactInfoParamsWithTimeout(timeout time.Duration) *ContactInfoParams {
	var ()
	return &ContactInfoParams{

		timeout: timeout,
	}
}

// NewContactInfoParamsWithContext creates a new ContactInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewContactInfoParamsWithContext(ctx context.Context) *ContactInfoParams {
	var ()
	return &ContactInfoParams{

		Context: ctx,
	}
}

// NewContactInfoParamsWithHTTPClient creates a new ContactInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContactInfoParamsWithHTTPClient(client *http.Client) *ContactInfoParams {
	var ()
	return &ContactInfoParams{
		HTTPClient: client,
	}
}

/*ContactInfoParams contains all the parameters to send to the API endpoint
for the contact info operation typically these are written to a http.Request
*/
type ContactInfoParams struct {

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

// WithTimeout adds the timeout to the contact info params
func (o *ContactInfoParams) WithTimeout(timeout time.Duration) *ContactInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contact info params
func (o *ContactInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contact info params
func (o *ContactInfoParams) WithContext(ctx context.Context) *ContactInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contact info params
func (o *ContactInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contact info params
func (o *ContactInfoParams) WithHTTPClient(client *http.Client) *ContactInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contact info params
func (o *ContactInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the contact info params
func (o *ContactInfoParams) WithXDomainrobotContext(xDomainrobotContext *string) *ContactInfoParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the contact info params
func (o *ContactInfoParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the contact info params
func (o *ContactInfoParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *ContactInfoParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the contact info params
func (o *ContactInfoParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the contact info params
func (o *ContactInfoParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *ContactInfoParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the contact info params
func (o *ContactInfoParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the contact info params
func (o *ContactInfoParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *ContactInfoParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the contact info params
func (o *ContactInfoParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the contact info params
func (o *ContactInfoParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *ContactInfoParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the contact info params
func (o *ContactInfoParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the contact info params
func (o *ContactInfoParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *ContactInfoParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the contact info params
func (o *ContactInfoParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the contact info params
func (o *ContactInfoParams) WithXDomainrobotWS(xDomainrobotWS *string) *ContactInfoParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the contact info params
func (o *ContactInfoParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithID adds the id to the contact info params
func (o *ContactInfoParams) WithID(id int32) *ContactInfoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the contact info params
func (o *ContactInfoParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContactInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
