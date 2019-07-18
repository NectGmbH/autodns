// Code generated by go-swagger; DO NOT EDIT.

package id4me_identity_tasks

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

// NewId4MeIdentityInfoParams creates a new Id4MeIdentityInfoParams object
// with the default values initialized.
func NewId4MeIdentityInfoParams() *Id4MeIdentityInfoParams {
	var ()
	return &Id4MeIdentityInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewId4MeIdentityInfoParamsWithTimeout creates a new Id4MeIdentityInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewId4MeIdentityInfoParamsWithTimeout(timeout time.Duration) *Id4MeIdentityInfoParams {
	var ()
	return &Id4MeIdentityInfoParams{

		timeout: timeout,
	}
}

// NewId4MeIdentityInfoParamsWithContext creates a new Id4MeIdentityInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewId4MeIdentityInfoParamsWithContext(ctx context.Context) *Id4MeIdentityInfoParams {
	var ()
	return &Id4MeIdentityInfoParams{

		Context: ctx,
	}
}

// NewId4MeIdentityInfoParamsWithHTTPClient creates a new Id4MeIdentityInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewId4MeIdentityInfoParamsWithHTTPClient(client *http.Client) *Id4MeIdentityInfoParams {
	var ()
	return &Id4MeIdentityInfoParams{
		HTTPClient: client,
	}
}

/*Id4MeIdentityInfoParams contains all the parameters to send to the API endpoint
for the id4 me identity info operation typically these are written to a http.Request
*/
type Id4MeIdentityInfoParams struct {

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
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithTimeout(timeout time.Duration) *Id4MeIdentityInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithContext(ctx context.Context) *Id4MeIdentityInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithHTTPClient(client *http.Client) *Id4MeIdentityInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithXDomainrobotContext(xDomainrobotContext *string) *Id4MeIdentityInfoParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *Id4MeIdentityInfoParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *Id4MeIdentityInfoParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *Id4MeIdentityInfoParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *Id4MeIdentityInfoParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *Id4MeIdentityInfoParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithXDomainrobotWS(xDomainrobotWS *string) *Id4MeIdentityInfoParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithName adds the name to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) WithName(name string) *Id4MeIdentityInfoParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the id4 me identity info params
func (o *Id4MeIdentityInfoParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *Id4MeIdentityInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
