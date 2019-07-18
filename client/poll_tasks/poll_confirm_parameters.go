// Code generated by go-swagger; DO NOT EDIT.

package poll_tasks

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

// NewPollConfirmParams creates a new PollConfirmParams object
// with the default values initialized.
func NewPollConfirmParams() *PollConfirmParams {
	var ()
	return &PollConfirmParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPollConfirmParamsWithTimeout creates a new PollConfirmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPollConfirmParamsWithTimeout(timeout time.Duration) *PollConfirmParams {
	var ()
	return &PollConfirmParams{

		timeout: timeout,
	}
}

// NewPollConfirmParamsWithContext creates a new PollConfirmParams object
// with the default values initialized, and the ability to set a context for a request
func NewPollConfirmParamsWithContext(ctx context.Context) *PollConfirmParams {
	var ()
	return &PollConfirmParams{

		Context: ctx,
	}
}

// NewPollConfirmParamsWithHTTPClient creates a new PollConfirmParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPollConfirmParamsWithHTTPClient(client *http.Client) *PollConfirmParams {
	var ()
	return &PollConfirmParams{
		HTTPClient: client,
	}
}

/*PollConfirmParams contains all the parameters to send to the API endpoint
for the poll confirm operation typically these are written to a http.Request
*/
type PollConfirmParams struct {

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
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the poll confirm params
func (o *PollConfirmParams) WithTimeout(timeout time.Duration) *PollConfirmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the poll confirm params
func (o *PollConfirmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the poll confirm params
func (o *PollConfirmParams) WithContext(ctx context.Context) *PollConfirmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the poll confirm params
func (o *PollConfirmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the poll confirm params
func (o *PollConfirmParams) WithHTTPClient(client *http.Client) *PollConfirmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the poll confirm params
func (o *PollConfirmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the poll confirm params
func (o *PollConfirmParams) WithXDomainrobotContext(xDomainrobotContext *string) *PollConfirmParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the poll confirm params
func (o *PollConfirmParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the poll confirm params
func (o *PollConfirmParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *PollConfirmParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the poll confirm params
func (o *PollConfirmParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the poll confirm params
func (o *PollConfirmParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *PollConfirmParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the poll confirm params
func (o *PollConfirmParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the poll confirm params
func (o *PollConfirmParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *PollConfirmParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the poll confirm params
func (o *PollConfirmParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the poll confirm params
func (o *PollConfirmParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *PollConfirmParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the poll confirm params
func (o *PollConfirmParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the poll confirm params
func (o *PollConfirmParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *PollConfirmParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the poll confirm params
func (o *PollConfirmParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the poll confirm params
func (o *PollConfirmParams) WithXDomainrobotWS(xDomainrobotWS *string) *PollConfirmParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the poll confirm params
func (o *PollConfirmParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithID adds the id to the poll confirm params
func (o *PollConfirmParams) WithID(id string) *PollConfirmParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the poll confirm params
func (o *PollConfirmParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PollConfirmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
