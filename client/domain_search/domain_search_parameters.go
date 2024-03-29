// Code generated by go-swagger; DO NOT EDIT.

package domain_search

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

// NewDomainSearchParams creates a new DomainSearchParams object
// with the default values initialized.
func NewDomainSearchParams() *DomainSearchParams {
	var ()
	return &DomainSearchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainSearchParamsWithTimeout creates a new DomainSearchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainSearchParamsWithTimeout(timeout time.Duration) *DomainSearchParams {
	var ()
	return &DomainSearchParams{

		timeout: timeout,
	}
}

// NewDomainSearchParamsWithContext creates a new DomainSearchParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainSearchParamsWithContext(ctx context.Context) *DomainSearchParams {
	var ()
	return &DomainSearchParams{

		Context: ctx,
	}
}

// NewDomainSearchParamsWithHTTPClient creates a new DomainSearchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainSearchParamsWithHTTPClient(client *http.Client) *DomainSearchParams {
	var ()
	return &DomainSearchParams{
		HTTPClient: client,
	}
}

/*DomainSearchParams contains all the parameters to send to the API endpoint
for the domain search operation typically these are written to a http.Request
*/
type DomainSearchParams struct {

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
	  search

	*/
	Body *models.DomainEnvelopeSearchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domain search params
func (o *DomainSearchParams) WithTimeout(timeout time.Duration) *DomainSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain search params
func (o *DomainSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain search params
func (o *DomainSearchParams) WithContext(ctx context.Context) *DomainSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain search params
func (o *DomainSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain search params
func (o *DomainSearchParams) WithHTTPClient(client *http.Client) *DomainSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain search params
func (o *DomainSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain search params
func (o *DomainSearchParams) WithXDomainrobotContext(xDomainrobotContext *string) *DomainSearchParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain search params
func (o *DomainSearchParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain search params
func (o *DomainSearchParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *DomainSearchParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain search params
func (o *DomainSearchParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain search params
func (o *DomainSearchParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *DomainSearchParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain search params
func (o *DomainSearchParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain search params
func (o *DomainSearchParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainSearchParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain search params
func (o *DomainSearchParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the domain search params
func (o *DomainSearchParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *DomainSearchParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the domain search params
func (o *DomainSearchParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain search params
func (o *DomainSearchParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainSearchParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain search params
func (o *DomainSearchParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain search params
func (o *DomainSearchParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainSearchParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain search params
func (o *DomainSearchParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the domain search params
func (o *DomainSearchParams) WithBody(body *models.DomainEnvelopeSearchRequest) *DomainSearchParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domain search params
func (o *DomainSearchParams) SetBody(body *models.DomainEnvelopeSearchRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DomainSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
