// Code generated by go-swagger; DO NOT EDIT.

package certificate_tasks

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

// NewCertificateRevokeParams creates a new CertificateRevokeParams object
// with the default values initialized.
func NewCertificateRevokeParams() *CertificateRevokeParams {
	var ()
	return &CertificateRevokeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCertificateRevokeParamsWithTimeout creates a new CertificateRevokeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCertificateRevokeParamsWithTimeout(timeout time.Duration) *CertificateRevokeParams {
	var ()
	return &CertificateRevokeParams{

		timeout: timeout,
	}
}

// NewCertificateRevokeParamsWithContext creates a new CertificateRevokeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCertificateRevokeParamsWithContext(ctx context.Context) *CertificateRevokeParams {
	var ()
	return &CertificateRevokeParams{

		Context: ctx,
	}
}

// NewCertificateRevokeParamsWithHTTPClient creates a new CertificateRevokeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCertificateRevokeParamsWithHTTPClient(client *http.Client) *CertificateRevokeParams {
	var ()
	return &CertificateRevokeParams{
		HTTPClient: client,
	}
}

/*CertificateRevokeParams contains all the parameters to send to the API endpoint
for the certificate revoke operation typically these are written to a http.Request
*/
type CertificateRevokeParams struct {

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
	  certificate

	*/
	Body *models.Certificate
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the certificate revoke params
func (o *CertificateRevokeParams) WithTimeout(timeout time.Duration) *CertificateRevokeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the certificate revoke params
func (o *CertificateRevokeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the certificate revoke params
func (o *CertificateRevokeParams) WithContext(ctx context.Context) *CertificateRevokeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the certificate revoke params
func (o *CertificateRevokeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the certificate revoke params
func (o *CertificateRevokeParams) WithHTTPClient(client *http.Client) *CertificateRevokeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the certificate revoke params
func (o *CertificateRevokeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the certificate revoke params
func (o *CertificateRevokeParams) WithXDomainrobotContext(xDomainrobotContext *string) *CertificateRevokeParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the certificate revoke params
func (o *CertificateRevokeParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the certificate revoke params
func (o *CertificateRevokeParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *CertificateRevokeParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the certificate revoke params
func (o *CertificateRevokeParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the certificate revoke params
func (o *CertificateRevokeParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *CertificateRevokeParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the certificate revoke params
func (o *CertificateRevokeParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the certificate revoke params
func (o *CertificateRevokeParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *CertificateRevokeParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the certificate revoke params
func (o *CertificateRevokeParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the certificate revoke params
func (o *CertificateRevokeParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *CertificateRevokeParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the certificate revoke params
func (o *CertificateRevokeParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the certificate revoke params
func (o *CertificateRevokeParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *CertificateRevokeParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the certificate revoke params
func (o *CertificateRevokeParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the certificate revoke params
func (o *CertificateRevokeParams) WithXDomainrobotWS(xDomainrobotWS *string) *CertificateRevokeParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the certificate revoke params
func (o *CertificateRevokeParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the certificate revoke params
func (o *CertificateRevokeParams) WithBody(body *models.Certificate) *CertificateRevokeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the certificate revoke params
func (o *CertificateRevokeParams) SetBody(body *models.Certificate) {
	o.Body = body
}

// WithID adds the id to the certificate revoke params
func (o *CertificateRevokeParams) WithID(id int32) *CertificateRevokeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the certificate revoke params
func (o *CertificateRevokeParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CertificateRevokeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}