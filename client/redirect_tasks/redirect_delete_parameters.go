// Code generated by go-swagger; DO NOT EDIT.

package redirect_tasks

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

// NewRedirectDeleteParams creates a new RedirectDeleteParams object
// with the default values initialized.
func NewRedirectDeleteParams() *RedirectDeleteParams {
	var ()
	return &RedirectDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRedirectDeleteParamsWithTimeout creates a new RedirectDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRedirectDeleteParamsWithTimeout(timeout time.Duration) *RedirectDeleteParams {
	var ()
	return &RedirectDeleteParams{

		timeout: timeout,
	}
}

// NewRedirectDeleteParamsWithContext creates a new RedirectDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewRedirectDeleteParamsWithContext(ctx context.Context) *RedirectDeleteParams {
	var ()
	return &RedirectDeleteParams{

		Context: ctx,
	}
}

// NewRedirectDeleteParamsWithHTTPClient creates a new RedirectDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRedirectDeleteParamsWithHTTPClient(client *http.Client) *RedirectDeleteParams {
	var ()
	return &RedirectDeleteParams{
		HTTPClient: client,
	}
}

/*RedirectDeleteParams contains all the parameters to send to the API endpoint
for the redirect delete operation typically these are written to a http.Request
*/
type RedirectDeleteParams struct {

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
	/*Keys
	  If the dns should be provisioned if available.

	*/
	Keys []string
	/*Source*/
	Source string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the redirect delete params
func (o *RedirectDeleteParams) WithTimeout(timeout time.Duration) *RedirectDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the redirect delete params
func (o *RedirectDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the redirect delete params
func (o *RedirectDeleteParams) WithContext(ctx context.Context) *RedirectDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the redirect delete params
func (o *RedirectDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the redirect delete params
func (o *RedirectDeleteParams) WithHTTPClient(client *http.Client) *RedirectDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the redirect delete params
func (o *RedirectDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotContext(xDomainrobotContext *string) *RedirectDeleteParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *RedirectDeleteParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *RedirectDeleteParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *RedirectDeleteParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *RedirectDeleteParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *RedirectDeleteParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotWS(xDomainrobotWS *string) *RedirectDeleteParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithKeys adds the keys to the redirect delete params
func (o *RedirectDeleteParams) WithKeys(keys []string) *RedirectDeleteParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the redirect delete params
func (o *RedirectDeleteParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithSource adds the source to the redirect delete params
func (o *RedirectDeleteParams) WithSource(source string) *RedirectDeleteParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the redirect delete params
func (o *RedirectDeleteParams) SetSource(source string) {
	o.Source = source
}

// WriteToRequest writes these params to a swagger request
func (o *RedirectDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	// path param source
	if err := r.SetPathParam("source", o.Source); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}