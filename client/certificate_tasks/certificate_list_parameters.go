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

// NewCertificateListParams creates a new CertificateListParams object
// with the default values initialized.
func NewCertificateListParams() *CertificateListParams {
	var ()
	return &CertificateListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCertificateListParamsWithTimeout creates a new CertificateListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCertificateListParamsWithTimeout(timeout time.Duration) *CertificateListParams {
	var ()
	return &CertificateListParams{

		timeout: timeout,
	}
}

// NewCertificateListParamsWithContext creates a new CertificateListParams object
// with the default values initialized, and the ability to set a context for a request
func NewCertificateListParamsWithContext(ctx context.Context) *CertificateListParams {
	var ()
	return &CertificateListParams{

		Context: ctx,
	}
}

// NewCertificateListParamsWithHTTPClient creates a new CertificateListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCertificateListParamsWithHTTPClient(client *http.Client) *CertificateListParams {
	var ()
	return &CertificateListParams{
		HTTPClient: client,
	}
}

/*CertificateListParams contains all the parameters to send to the API endpoint
for the certificate list operation typically these are written to a http.Request
*/
type CertificateListParams struct {

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

// WithTimeout adds the timeout to the certificate list params
func (o *CertificateListParams) WithTimeout(timeout time.Duration) *CertificateListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the certificate list params
func (o *CertificateListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the certificate list params
func (o *CertificateListParams) WithContext(ctx context.Context) *CertificateListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the certificate list params
func (o *CertificateListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the certificate list params
func (o *CertificateListParams) WithHTTPClient(client *http.Client) *CertificateListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the certificate list params
func (o *CertificateListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotContext adds the xDomainrobotContext to the certificate list params
func (o *CertificateListParams) WithXDomainrobotContext(xDomainrobotContext *string) *CertificateListParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the certificate list params
func (o *CertificateListParams) SetXDomainrobotContext(xDomainrobotContext *string) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the certificate list params
func (o *CertificateListParams) WithXDomainrobotDemo(xDomainrobotDemo *string) *CertificateListParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the certificate list params
func (o *CertificateListParams) SetXDomainrobotDemo(xDomainrobotDemo *string) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the certificate list params
func (o *CertificateListParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) *CertificateListParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the certificate list params
func (o *CertificateListParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *string) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the certificate list params
func (o *CertificateListParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *CertificateListParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the certificate list params
func (o *CertificateListParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotPIN adds the xDomainrobotPIN to the certificate list params
func (o *CertificateListParams) WithXDomainrobotPIN(xDomainrobotPIN *string) *CertificateListParams {
	o.SetXDomainrobotPIN(xDomainrobotPIN)
	return o
}

// SetXDomainrobotPIN adds the xDomainrobotPIN to the certificate list params
func (o *CertificateListParams) SetXDomainrobotPIN(xDomainrobotPIN *string) {
	o.XDomainrobotPIN = xDomainrobotPIN
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the certificate list params
func (o *CertificateListParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *CertificateListParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the certificate list params
func (o *CertificateListParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the certificate list params
func (o *CertificateListParams) WithXDomainrobotWS(xDomainrobotWS *string) *CertificateListParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the certificate list params
func (o *CertificateListParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the certificate list params
func (o *CertificateListParams) WithBody(body *models.Query) *CertificateListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the certificate list params
func (o *CertificateListParams) SetBody(body *models.Query) {
	o.Body = body
}

// WithKeys adds the keys to the certificate list params
func (o *CertificateListParams) WithKeys(keys []string) *CertificateListParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the certificate list params
func (o *CertificateListParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *CertificateListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
