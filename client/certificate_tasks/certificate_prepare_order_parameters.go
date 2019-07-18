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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// NewCertificatePrepareOrderParams creates a new CertificatePrepareOrderParams object
// with the default values initialized.
func NewCertificatePrepareOrderParams() *CertificatePrepareOrderParams {
	var ()
	return &CertificatePrepareOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCertificatePrepareOrderParamsWithTimeout creates a new CertificatePrepareOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCertificatePrepareOrderParamsWithTimeout(timeout time.Duration) *CertificatePrepareOrderParams {
	var ()
	return &CertificatePrepareOrderParams{

		timeout: timeout,
	}
}

// NewCertificatePrepareOrderParamsWithContext creates a new CertificatePrepareOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCertificatePrepareOrderParamsWithContext(ctx context.Context) *CertificatePrepareOrderParams {
	var ()
	return &CertificatePrepareOrderParams{

		Context: ctx,
	}
}

// NewCertificatePrepareOrderParamsWithHTTPClient creates a new CertificatePrepareOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCertificatePrepareOrderParamsWithHTTPClient(client *http.Client) *CertificatePrepareOrderParams {
	var ()
	return &CertificatePrepareOrderParams{
		HTTPClient: client,
	}
}

/*CertificatePrepareOrderParams contains all the parameters to send to the API endpoint
for the certificate prepare order operation typically these are written to a http.Request
*/
type CertificatePrepareOrderParams struct {

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
	  certificateData

	*/
	Body *models.CertificateData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithTimeout(timeout time.Duration) *CertificatePrepareOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithContext(ctx context.Context) *CertificatePrepareOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithHTTPClient(client *http.Client) *CertificatePrepareOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithXDomainrobotContext(xDomainrobotContext *string) *CertificatePrepareOrderParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *CertificatePrepareOrderParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *CertificatePrepareOrderParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *CertificatePrepareOrderParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *CertificatePrepareOrderParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *CertificatePrepareOrderParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithXDomainrobotWS(xDomainrobotWS *string) *CertificatePrepareOrderParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the certificate prepare order params
func (o *CertificatePrepareOrderParams) WithBody(body *models.CertificateData) *CertificatePrepareOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the certificate prepare order params
func (o *CertificatePrepareOrderParams) SetBody(body *models.CertificateData) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CertificatePrepareOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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