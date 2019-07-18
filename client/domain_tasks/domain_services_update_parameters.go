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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// NewDomainServicesUpdateParams creates a new DomainServicesUpdateParams object
// with the default values initialized.
func NewDomainServicesUpdateParams() *DomainServicesUpdateParams {
	var ()
	return &DomainServicesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainServicesUpdateParamsWithTimeout creates a new DomainServicesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainServicesUpdateParamsWithTimeout(timeout time.Duration) *DomainServicesUpdateParams {
	var ()
	return &DomainServicesUpdateParams{

		timeout: timeout,
	}
}

// NewDomainServicesUpdateParamsWithContext creates a new DomainServicesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainServicesUpdateParamsWithContext(ctx context.Context) *DomainServicesUpdateParams {
	var ()
	return &DomainServicesUpdateParams{

		Context: ctx,
	}
}

// NewDomainServicesUpdateParamsWithHTTPClient creates a new DomainServicesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainServicesUpdateParamsWithHTTPClient(client *http.Client) *DomainServicesUpdateParams {
	var ()
	return &DomainServicesUpdateParams{
		HTTPClient: client,
	}
}

/*DomainServicesUpdateParams contains all the parameters to send to the API endpoint
for the domain services update operation typically these are written to a http.Request
*/
type DomainServicesUpdateParams struct {

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
	  services

	*/
	Body *models.DomainServicesUpdate
	/*Keys
	  If the dns should be provisioned if available.

	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domain services update params
func (o *DomainServicesUpdateParams) WithTimeout(timeout time.Duration) *DomainServicesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain services update params
func (o *DomainServicesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain services update params
func (o *DomainServicesUpdateParams) WithContext(ctx context.Context) *DomainServicesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain services update params
func (o *DomainServicesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain services update params
func (o *DomainServicesUpdateParams) WithHTTPClient(client *http.Client) *DomainServicesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain services update params
func (o *DomainServicesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain services update params
func (o *DomainServicesUpdateParams) WithXDomainrobotContext(xDomainrobotContext *string) *DomainServicesUpdateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain services update params
func (o *DomainServicesUpdateParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain services update params
func (o *DomainServicesUpdateParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *DomainServicesUpdateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain services update params
func (o *DomainServicesUpdateParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain services update params
func (o *DomainServicesUpdateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *DomainServicesUpdateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain services update params
func (o *DomainServicesUpdateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain services update params
func (o *DomainServicesUpdateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainServicesUpdateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain services update params
func (o *DomainServicesUpdateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the domain services update params
func (o *DomainServicesUpdateParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *DomainServicesUpdateParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the domain services update params
func (o *DomainServicesUpdateParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain services update params
func (o *DomainServicesUpdateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainServicesUpdateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain services update params
func (o *DomainServicesUpdateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain services update params
func (o *DomainServicesUpdateParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainServicesUpdateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain services update params
func (o *DomainServicesUpdateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the domain services update params
func (o *DomainServicesUpdateParams) WithBody(body *models.DomainServicesUpdate) *DomainServicesUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domain services update params
func (o *DomainServicesUpdateParams) SetBody(body *models.DomainServicesUpdate) {
	o.Body = body
}

// WithKeys adds the keys to the domain services update params
func (o *DomainServicesUpdateParams) WithKeys(keys []string) *DomainServicesUpdateParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the domain services update params
func (o *DomainServicesUpdateParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *DomainServicesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}