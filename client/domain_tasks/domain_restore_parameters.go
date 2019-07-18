// Code generated by go-swagger; DO NOT EDIT.

package domain_tasks

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

	models "github.com/NectGmbH/autodns/models"
)

// NewDomainRestoreParams creates a new DomainRestoreParams object
// with the default values initialized.
func NewDomainRestoreParams() *DomainRestoreParams {
	var ()
	return &DomainRestoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainRestoreParamsWithTimeout creates a new DomainRestoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainRestoreParamsWithTimeout(timeout time.Duration) *DomainRestoreParams {
	var ()
	return &DomainRestoreParams{

		timeout: timeout,
	}
}

// NewDomainRestoreParamsWithContext creates a new DomainRestoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainRestoreParamsWithContext(ctx context.Context) *DomainRestoreParams {
	var ()
	return &DomainRestoreParams{

		Context: ctx,
	}
}

// NewDomainRestoreParamsWithHTTPClient creates a new DomainRestoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainRestoreParamsWithHTTPClient(client *http.Client) *DomainRestoreParams {
	var ()
	return &DomainRestoreParams{
		HTTPClient: client,
	}
}

/*DomainRestoreParams contains all the parameters to send to the API endpoint
for the domain restore operation typically these are written to a http.Request
*/
type DomainRestoreParams struct {

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
	  domain

	*/
	Body *models.Domain
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domain restore params
func (o *DomainRestoreParams) WithTimeout(timeout time.Duration) *DomainRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain restore params
func (o *DomainRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain restore params
func (o *DomainRestoreParams) WithContext(ctx context.Context) *DomainRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain restore params
func (o *DomainRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain restore params
func (o *DomainRestoreParams) WithHTTPClient(client *http.Client) *DomainRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain restore params
func (o *DomainRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain restore params
func (o *DomainRestoreParams) WithXDomainrobotContext(xDomainrobotContext *string) *DomainRestoreParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain restore params
func (o *DomainRestoreParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain restore params
func (o *DomainRestoreParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *DomainRestoreParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain restore params
func (o *DomainRestoreParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain restore params
func (o *DomainRestoreParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *DomainRestoreParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain restore params
func (o *DomainRestoreParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain restore params
func (o *DomainRestoreParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainRestoreParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain restore params
func (o *DomainRestoreParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the domain restore params
func (o *DomainRestoreParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *DomainRestoreParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the domain restore params
func (o *DomainRestoreParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain restore params
func (o *DomainRestoreParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainRestoreParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain restore params
func (o *DomainRestoreParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain restore params
func (o *DomainRestoreParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainRestoreParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain restore params
func (o *DomainRestoreParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the domain restore params
func (o *DomainRestoreParams) WithBody(body *models.Domain) *DomainRestoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domain restore params
func (o *DomainRestoreParams) SetBody(body *models.Domain) {
	o.Body = body
}

// WithName adds the name to the domain restore params
func (o *DomainRestoreParams) WithName(name string) *DomainRestoreParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the domain restore params
func (o *DomainRestoreParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DomainRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
