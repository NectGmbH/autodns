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

// NewDomainCancelationUpdateParams creates a new DomainCancelationUpdateParams object
// with the default values initialized.
func NewDomainCancelationUpdateParams() *DomainCancelationUpdateParams {
	var ()
	return &DomainCancelationUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainCancelationUpdateParamsWithTimeout creates a new DomainCancelationUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainCancelationUpdateParamsWithTimeout(timeout time.Duration) *DomainCancelationUpdateParams {
	var ()
	return &DomainCancelationUpdateParams{

		timeout: timeout,
	}
}

// NewDomainCancelationUpdateParamsWithContext creates a new DomainCancelationUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainCancelationUpdateParamsWithContext(ctx context.Context) *DomainCancelationUpdateParams {
	var ()
	return &DomainCancelationUpdateParams{

		Context: ctx,
	}
}

// NewDomainCancelationUpdateParamsWithHTTPClient creates a new DomainCancelationUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainCancelationUpdateParamsWithHTTPClient(client *http.Client) *DomainCancelationUpdateParams {
	var ()
	return &DomainCancelationUpdateParams{
		HTTPClient: client,
	}
}

/*DomainCancelationUpdateParams contains all the parameters to send to the API endpoint
for the domain cancelation update operation typically these are written to a http.Request
*/
type DomainCancelationUpdateParams struct {

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

// WithTimeout adds the timeout to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithTimeout(timeout time.Duration) *DomainCancelationUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithContext(ctx context.Context) *DomainCancelationUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithHTTPClient(client *http.Client) *DomainCancelationUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithXDomainrobotContext(xDomainrobotContext *string) *DomainCancelationUpdateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *DomainCancelationUpdateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *DomainCancelationUpdateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainCancelationUpdateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *DomainCancelationUpdateParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainCancelationUpdateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainCancelationUpdateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithBody(body *models.Domain) *DomainCancelationUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetBody(body *models.Domain) {
	o.Body = body
}

// WithName adds the name to the domain cancelation update params
func (o *DomainCancelationUpdateParams) WithName(name string) *DomainCancelationUpdateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the domain cancelation update params
func (o *DomainCancelationUpdateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DomainCancelationUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
