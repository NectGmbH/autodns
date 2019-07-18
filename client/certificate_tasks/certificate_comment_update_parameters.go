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

// NewCertificateCommentUpdateParams creates a new CertificateCommentUpdateParams object
// with the default values initialized.
func NewCertificateCommentUpdateParams() *CertificateCommentUpdateParams {
	var ()
	return &CertificateCommentUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCertificateCommentUpdateParamsWithTimeout creates a new CertificateCommentUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCertificateCommentUpdateParamsWithTimeout(timeout time.Duration) *CertificateCommentUpdateParams {
	var ()
	return &CertificateCommentUpdateParams{

		timeout: timeout,
	}
}

// NewCertificateCommentUpdateParamsWithContext creates a new CertificateCommentUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCertificateCommentUpdateParamsWithContext(ctx context.Context) *CertificateCommentUpdateParams {
	var ()
	return &CertificateCommentUpdateParams{

		Context: ctx,
	}
}

// NewCertificateCommentUpdateParamsWithHTTPClient creates a new CertificateCommentUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCertificateCommentUpdateParamsWithHTTPClient(client *http.Client) *CertificateCommentUpdateParams {
	var ()
	return &CertificateCommentUpdateParams{
		HTTPClient: client,
	}
}

/*CertificateCommentUpdateParams contains all the parameters to send to the API endpoint
for the certificate comment update operation typically these are written to a http.Request
*/
type CertificateCommentUpdateParams struct {

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

// WithTimeout adds the timeout to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithTimeout(timeout time.Duration) *CertificateCommentUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithContext(ctx context.Context) *CertificateCommentUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithHTTPClient(client *http.Client) *CertificateCommentUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithXDomainrobotContext(xDomainrobotContext *string) *CertificateCommentUpdateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *CertificateCommentUpdateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *CertificateCommentUpdateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *CertificateCommentUpdateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *CertificateCommentUpdateParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *CertificateCommentUpdateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithXDomainrobotWS(xDomainrobotWS *string) *CertificateCommentUpdateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithBody(body *models.Certificate) *CertificateCommentUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetBody(body *models.Certificate) {
	o.Body = body
}

// WithID adds the id to the certificate comment update params
func (o *CertificateCommentUpdateParams) WithID(id int32) *CertificateCommentUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the certificate comment update params
func (o *CertificateCommentUpdateParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CertificateCommentUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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