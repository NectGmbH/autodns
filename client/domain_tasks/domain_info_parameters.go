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
)

// NewDomainInfoParams creates a new DomainInfoParams object
// with the default values initialized.
func NewDomainInfoParams() *DomainInfoParams {
	var ()
	return &DomainInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainInfoParamsWithTimeout creates a new DomainInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainInfoParamsWithTimeout(timeout time.Duration) *DomainInfoParams {
	var ()
	return &DomainInfoParams{

		timeout: timeout,
	}
}

// NewDomainInfoParamsWithContext creates a new DomainInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainInfoParamsWithContext(ctx context.Context) *DomainInfoParams {
	var ()
	return &DomainInfoParams{

		Context: ctx,
	}
}

// NewDomainInfoParamsWithHTTPClient creates a new DomainInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainInfoParamsWithHTTPClient(client *http.Client) *DomainInfoParams {
	var ()
	return &DomainInfoParams{
		HTTPClient: client,
	}
}

/*DomainInfoParams contains all the parameters to send to the API endpoint
for the domain info operation typically these are written to a http.Request
*/
type DomainInfoParams struct {

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

// WithTimeout adds the timeout to the domain info params
func (o *DomainInfoParams) WithTimeout(timeout time.Duration) *DomainInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain info params
func (o *DomainInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain info params
func (o *DomainInfoParams) WithContext(ctx context.Context) *DomainInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain info params
func (o *DomainInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain info params
func (o *DomainInfoParams) WithHTTPClient(client *http.Client) *DomainInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain info params
func (o *DomainInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain info params
func (o *DomainInfoParams) WithXDomainrobotContext(xDomainrobotContext *string) *DomainInfoParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain info params
func (o *DomainInfoParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain info params
func (o *DomainInfoParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *DomainInfoParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain info params
func (o *DomainInfoParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain info params
func (o *DomainInfoParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *DomainInfoParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain info params
func (o *DomainInfoParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain info params
func (o *DomainInfoParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainInfoParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain info params
func (o *DomainInfoParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the domain info params
func (o *DomainInfoParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *DomainInfoParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the domain info params
func (o *DomainInfoParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain info params
func (o *DomainInfoParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainInfoParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain info params
func (o *DomainInfoParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain info params
func (o *DomainInfoParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainInfoParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain info params
func (o *DomainInfoParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithName adds the name to the domain info params
func (o *DomainInfoParams) WithName(name string) *DomainInfoParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the domain info params
func (o *DomainInfoParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DomainInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
