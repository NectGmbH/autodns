// Code generated by go-swagger; DO NOT EDIT.

package id4me_agent_tasks

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

// NewUpdateId4MeAgentParams creates a new UpdateId4MeAgentParams object
// with the default values initialized.
func NewUpdateId4MeAgentParams() *UpdateId4MeAgentParams {
	var ()
	return &UpdateId4MeAgentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateId4MeAgentParamsWithTimeout creates a new UpdateId4MeAgentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateId4MeAgentParamsWithTimeout(timeout time.Duration) *UpdateId4MeAgentParams {
	var ()
	return &UpdateId4MeAgentParams{

		timeout: timeout,
	}
}

// NewUpdateId4MeAgentParamsWithContext creates a new UpdateId4MeAgentParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateId4MeAgentParamsWithContext(ctx context.Context) *UpdateId4MeAgentParams {
	var ()
	return &UpdateId4MeAgentParams{

		Context: ctx,
	}
}

// NewUpdateId4MeAgentParamsWithHTTPClient creates a new UpdateId4MeAgentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateId4MeAgentParamsWithHTTPClient(client *http.Client) *UpdateId4MeAgentParams {
	var ()
	return &UpdateId4MeAgentParams{
		HTTPClient: client,
	}
}

/*UpdateId4MeAgentParams contains all the parameters to send to the API endpoint
for the update id4 me agent operation typically these are written to a http.Request
*/
type UpdateId4MeAgentParams struct {

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
	  agent

	*/
	Body *models.Id4MeAgent

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithTimeout(timeout time.Duration) *UpdateId4MeAgentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithContext(ctx context.Context) *UpdateId4MeAgentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithHTTPClient(client *http.Client) *UpdateId4MeAgentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithXDomainrobotContext(xDomainrobotContext *string) *UpdateId4MeAgentParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *UpdateId4MeAgentParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *UpdateId4MeAgentParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *UpdateId4MeAgentParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *UpdateId4MeAgentParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *UpdateId4MeAgentParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithXDomainrobotWS(xDomainrobotWS *string) *UpdateId4MeAgentParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the update id4 me agent params
func (o *UpdateId4MeAgentParams) WithBody(body *models.Id4MeAgent) *UpdateId4MeAgentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update id4 me agent params
func (o *UpdateId4MeAgentParams) SetBody(body *models.Id4MeAgent) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateId4MeAgentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
