// Code generated by go-swagger; DO NOT EDIT.

package mail_proxy_tasks

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

// NewMailProxyUpdateParams creates a new MailProxyUpdateParams object
// with the default values initialized.
func NewMailProxyUpdateParams() *MailProxyUpdateParams {
	var ()
	return &MailProxyUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMailProxyUpdateParamsWithTimeout creates a new MailProxyUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMailProxyUpdateParamsWithTimeout(timeout time.Duration) *MailProxyUpdateParams {
	var ()
	return &MailProxyUpdateParams{

		timeout: timeout,
	}
}

// NewMailProxyUpdateParamsWithContext creates a new MailProxyUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewMailProxyUpdateParamsWithContext(ctx context.Context) *MailProxyUpdateParams {
	var ()
	return &MailProxyUpdateParams{

		Context: ctx,
	}
}

// NewMailProxyUpdateParamsWithHTTPClient creates a new MailProxyUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMailProxyUpdateParamsWithHTTPClient(client *http.Client) *MailProxyUpdateParams {
	var ()
	return &MailProxyUpdateParams{
		HTTPClient: client,
	}
}

/*MailProxyUpdateParams contains all the parameters to send to the API endpoint
for the mail proxy update operation typically these are written to a http.Request
*/
type MailProxyUpdateParams struct {

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
	  mailProxy

	*/
	Body *models.MailProxy
	/*Domain*/
	Domain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the mail proxy update params
func (o *MailProxyUpdateParams) WithTimeout(timeout time.Duration) *MailProxyUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mail proxy update params
func (o *MailProxyUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mail proxy update params
func (o *MailProxyUpdateParams) WithContext(ctx context.Context) *MailProxyUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mail proxy update params
func (o *MailProxyUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mail proxy update params
func (o *MailProxyUpdateParams) WithHTTPClient(client *http.Client) *MailProxyUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mail proxy update params
func (o *MailProxyUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the mail proxy update params
func (o *MailProxyUpdateParams) WithXDomainrobotContext(xDomainrobotContext *string) *MailProxyUpdateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the mail proxy update params
func (o *MailProxyUpdateParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the mail proxy update params
func (o *MailProxyUpdateParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *MailProxyUpdateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the mail proxy update params
func (o *MailProxyUpdateParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the mail proxy update params
func (o *MailProxyUpdateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *MailProxyUpdateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the mail proxy update params
func (o *MailProxyUpdateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the mail proxy update params
func (o *MailProxyUpdateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *MailProxyUpdateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the mail proxy update params
func (o *MailProxyUpdateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the mail proxy update params
func (o *MailProxyUpdateParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *MailProxyUpdateParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the mail proxy update params
func (o *MailProxyUpdateParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the mail proxy update params
func (o *MailProxyUpdateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *MailProxyUpdateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the mail proxy update params
func (o *MailProxyUpdateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the mail proxy update params
func (o *MailProxyUpdateParams) WithXDomainrobotWS(xDomainrobotWS *string) *MailProxyUpdateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the mail proxy update params
func (o *MailProxyUpdateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the mail proxy update params
func (o *MailProxyUpdateParams) WithBody(body *models.MailProxy) *MailProxyUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the mail proxy update params
func (o *MailProxyUpdateParams) SetBody(body *models.MailProxy) {
	o.Body = body
}

// WithDomain adds the domain to the mail proxy update params
func (o *MailProxyUpdateParams) WithDomain(domain string) *MailProxyUpdateParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the mail proxy update params
func (o *MailProxyUpdateParams) SetDomain(domain string) {
	o.Domain = domain
}

// WriteToRequest writes these params to a swagger request
func (o *MailProxyUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}