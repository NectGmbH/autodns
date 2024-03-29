// Code generated by go-swagger; DO NOT EDIT.

package job_tasks

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

// NewJobHistoryInfoParams creates a new JobHistoryInfoParams object
// with the default values initialized.
func NewJobHistoryInfoParams() *JobHistoryInfoParams {
	var ()
	return &JobHistoryInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobHistoryInfoParamsWithTimeout creates a new JobHistoryInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobHistoryInfoParamsWithTimeout(timeout time.Duration) *JobHistoryInfoParams {
	var ()
	return &JobHistoryInfoParams{

		timeout: timeout,
	}
}

// NewJobHistoryInfoParamsWithContext creates a new JobHistoryInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobHistoryInfoParamsWithContext(ctx context.Context) *JobHistoryInfoParams {
	var ()
	return &JobHistoryInfoParams{

		Context: ctx,
	}
}

// NewJobHistoryInfoParamsWithHTTPClient creates a new JobHistoryInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobHistoryInfoParamsWithHTTPClient(client *http.Client) *JobHistoryInfoParams {
	var ()
	return &JobHistoryInfoParams{
		HTTPClient: client,
	}
}

/*JobHistoryInfoParams contains all the parameters to send to the API endpoint
for the job history info operation typically these are written to a http.Request
*/
type JobHistoryInfoParams struct {

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
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the job history info params
func (o *JobHistoryInfoParams) WithTimeout(timeout time.Duration) *JobHistoryInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job history info params
func (o *JobHistoryInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job history info params
func (o *JobHistoryInfoParams) WithContext(ctx context.Context) *JobHistoryInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job history info params
func (o *JobHistoryInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job history info params
func (o *JobHistoryInfoParams) WithHTTPClient(client *http.Client) *JobHistoryInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job history info params
func (o *JobHistoryInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the job history info params
func (o *JobHistoryInfoParams) WithXDomainrobotContext(xDomainrobotContext *string) *JobHistoryInfoParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the job history info params
func (o *JobHistoryInfoParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the job history info params
func (o *JobHistoryInfoParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *JobHistoryInfoParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the job history info params
func (o *JobHistoryInfoParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the job history info params
func (o *JobHistoryInfoParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *JobHistoryInfoParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the job history info params
func (o *JobHistoryInfoParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the job history info params
func (o *JobHistoryInfoParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *JobHistoryInfoParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the job history info params
func (o *JobHistoryInfoParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the job history info params
func (o *JobHistoryInfoParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *JobHistoryInfoParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the job history info params
func (o *JobHistoryInfoParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the job history info params
func (o *JobHistoryInfoParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *JobHistoryInfoParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the job history info params
func (o *JobHistoryInfoParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the job history info params
func (o *JobHistoryInfoParams) WithXDomainrobotWS(xDomainrobotWS *string) *JobHistoryInfoParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the job history info params
func (o *JobHistoryInfoParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithID adds the id to the job history info params
func (o *JobHistoryInfoParams) WithID(id int32) *JobHistoryInfoParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the job history info params
func (o *JobHistoryInfoParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JobHistoryInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
