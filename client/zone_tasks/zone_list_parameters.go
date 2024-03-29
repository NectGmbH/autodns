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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// NewZoneListParams creates a new ZoneListParams object
// with the default values initialized.
func NewZoneListParams() *ZoneListParams {
	var ()
	return &ZoneListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewZoneListParamsWithTimeout creates a new ZoneListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewZoneListParamsWithTimeout(timeout time.Duration) *ZoneListParams {
	var ()
	return &ZoneListParams{

		timeout: timeout,
	}
}

// NewZoneListParamsWithContext creates a new ZoneListParams object
// with the default values initialized, and the ability to set a context for a request
func NewZoneListParamsWithContext(ctx context.Context) *ZoneListParams {
	var ()
	return &ZoneListParams{

		Context: ctx,
	}
}

// NewZoneListParamsWithHTTPClient creates a new ZoneListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewZoneListParamsWithHTTPClient(client *http.Client) *ZoneListParams {
	var ()
	return &ZoneListParams{
		HTTPClient: client,
	}
}

/*ZoneListParams contains all the parameters to send to the API endpoint
for the zone list operation typically these are written to a http.Request
*/
type ZoneListParams struct {

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
	  query

	*/
	Body *models.Query
	/*Keys
	  The query parameter to fetch additional details.

	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the zone list params
func (o *ZoneListParams) WithTimeout(timeout time.Duration) *ZoneListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the zone list params
func (o *ZoneListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the zone list params
func (o *ZoneListParams) WithContext(ctx context.Context) *ZoneListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the zone list params
func (o *ZoneListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the zone list params
func (o *ZoneListParams) WithHTTPClient(client *http.Client) *ZoneListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the zone list params
func (o *ZoneListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the zone list params
func (o *ZoneListParams) WithXDomainrobotContext(xDomainrobotContext *string) *ZoneListParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the zone list params
func (o *ZoneListParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the zone list params
func (o *ZoneListParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *ZoneListParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the zone list params
func (o *ZoneListParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the zone list params
func (o *ZoneListParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *ZoneListParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the zone list params
func (o *ZoneListParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the zone list params
func (o *ZoneListParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *ZoneListParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the zone list params
func (o *ZoneListParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the zone list params
func (o *ZoneListParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *ZoneListParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the zone list params
func (o *ZoneListParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the zone list params
func (o *ZoneListParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *ZoneListParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the zone list params
func (o *ZoneListParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the zone list params
func (o *ZoneListParams) WithXDomainrobotWS(xDomainrobotWS *string) *ZoneListParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the zone list params
func (o *ZoneListParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the zone list params
func (o *ZoneListParams) WithBody(body *models.Query) *ZoneListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the zone list params
func (o *ZoneListParams) SetBody(body *models.Query) {
	o.Body = body
}

// WithKeys adds the keys to the zone list params
func (o *ZoneListParams) WithKeys(keys []string) *ZoneListParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the zone list params
func (o *ZoneListParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *ZoneListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
