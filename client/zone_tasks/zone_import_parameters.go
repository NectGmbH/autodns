// Code generated by go-swagger; DO NOT EDIT.

package zone_tasks

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

// NewZoneImportParams creates a new ZoneImportParams object
// with the default values initialized.
func NewZoneImportParams() *ZoneImportParams {
	var ()
	return &ZoneImportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewZoneImportParamsWithTimeout creates a new ZoneImportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewZoneImportParamsWithTimeout(timeout time.Duration) *ZoneImportParams {
	var ()
	return &ZoneImportParams{

		timeout: timeout,
	}
}

// NewZoneImportParamsWithContext creates a new ZoneImportParams object
// with the default values initialized, and the ability to set a context for a request
func NewZoneImportParamsWithContext(ctx context.Context) *ZoneImportParams {
	var ()
	return &ZoneImportParams{

		Context: ctx,
	}
}

// NewZoneImportParamsWithHTTPClient creates a new ZoneImportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewZoneImportParamsWithHTTPClient(client *http.Client) *ZoneImportParams {
	var ()
	return &ZoneImportParams{
		HTTPClient: client,
	}
}

/*ZoneImportParams contains all the parameters to send to the API endpoint
for the zone import operation typically these are written to a http.Request
*/
type ZoneImportParams struct {

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
	  zone

	*/
	Body *models.Zone
	/*Name*/
	Name string
	/*Nameserver*/
	Nameserver string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the zone import params
func (o *ZoneImportParams) WithTimeout(timeout time.Duration) *ZoneImportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the zone import params
func (o *ZoneImportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the zone import params
func (o *ZoneImportParams) WithContext(ctx context.Context) *ZoneImportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the zone import params
func (o *ZoneImportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the zone import params
func (o *ZoneImportParams) WithHTTPClient(client *http.Client) *ZoneImportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the zone import params
func (o *ZoneImportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the zone import params
func (o *ZoneImportParams) WithXDomainrobotContext(xDomainrobotContext *string) *ZoneImportParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the zone import params
func (o *ZoneImportParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the zone import params
func (o *ZoneImportParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *ZoneImportParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the zone import params
func (o *ZoneImportParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the zone import params
func (o *ZoneImportParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *ZoneImportParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the zone import params
func (o *ZoneImportParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the zone import params
func (o *ZoneImportParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *ZoneImportParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the zone import params
func (o *ZoneImportParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the zone import params
func (o *ZoneImportParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *ZoneImportParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the zone import params
func (o *ZoneImportParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the zone import params
func (o *ZoneImportParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *ZoneImportParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the zone import params
func (o *ZoneImportParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the zone import params
func (o *ZoneImportParams) WithXDomainrobotWS(xDomainrobotWS *string) *ZoneImportParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the zone import params
func (o *ZoneImportParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the zone import params
func (o *ZoneImportParams) WithBody(body *models.Zone) *ZoneImportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the zone import params
func (o *ZoneImportParams) SetBody(body *models.Zone) {
	o.Body = body
}

// WithName adds the name to the zone import params
func (o *ZoneImportParams) WithName(name string) *ZoneImportParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the zone import params
func (o *ZoneImportParams) SetName(name string) {
	o.Name = name
}

// WithNameserver adds the nameserver to the zone import params
func (o *ZoneImportParams) WithNameserver(nameserver string) *ZoneImportParams {
	o.SetNameserver(nameserver)
	return o
}

// SetNameserver adds the nameserver to the zone import params
func (o *ZoneImportParams) SetNameserver(nameserver string) {
	o.Nameserver = nameserver
}

// WriteToRequest writes these params to a swagger request
func (o *ZoneImportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nameserver
	if err := r.SetPathParam("nameserver", o.Nameserver); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
