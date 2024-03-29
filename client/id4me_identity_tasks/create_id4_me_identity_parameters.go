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

	models "github.com/NectGmbH/autodns/models"
)

// NewCreateId4MeIdentityParams creates a new CreateId4MeIdentityParams object
// with the default values initialized.
func NewCreateId4MeIdentityParams() *CreateId4MeIdentityParams {
	var ()
	return &CreateId4MeIdentityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateId4MeIdentityParamsWithTimeout creates a new CreateId4MeIdentityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateId4MeIdentityParamsWithTimeout(timeout time.Duration) *CreateId4MeIdentityParams {
	var ()
	return &CreateId4MeIdentityParams{

		timeout: timeout,
	}
}

// NewCreateId4MeIdentityParamsWithContext creates a new CreateId4MeIdentityParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateId4MeIdentityParamsWithContext(ctx context.Context) *CreateId4MeIdentityParams {
	var ()
	return &CreateId4MeIdentityParams{

		Context: ctx,
	}
}

// NewCreateId4MeIdentityParamsWithHTTPClient creates a new CreateId4MeIdentityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateId4MeIdentityParamsWithHTTPClient(client *http.Client) *CreateId4MeIdentityParams {
	var ()
	return &CreateId4MeIdentityParams{
		HTTPClient: client,
	}
}

/*CreateId4MeIdentityParams contains all the parameters to send to the API endpoint
for the create id4 me identity operation typically these are written to a http.Request
*/
type CreateId4MeIdentityParams struct {

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
	  identity

	*/
	Body *models.Id4MeIdentity

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithTimeout(timeout time.Duration) *CreateId4MeIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithContext(ctx context.Context) *CreateId4MeIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithHTTPClient(client *http.Client) *CreateId4MeIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithXDomainrobotContext(xDomainrobotContext *string) *CreateId4MeIdentityParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *CreateId4MeIdentityParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *CreateId4MeIdentityParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *CreateId4MeIdentityParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *CreateId4MeIdentityParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *CreateId4MeIdentityParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithXDomainrobotWS(xDomainrobotWS *string) *CreateId4MeIdentityParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the create id4 me identity params
func (o *CreateId4MeIdentityParams) WithBody(body *models.Id4MeIdentity) *CreateId4MeIdentityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create id4 me identity params
func (o *CreateId4MeIdentityParams) SetBody(body *models.Id4MeIdentity) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateId4MeIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
