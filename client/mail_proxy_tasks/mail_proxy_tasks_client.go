// Code generated by go-swagger; DO NOT EDIT.

package mail_proxy_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new mail proxy tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mail proxy tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
MailProxyCreate creates mail proxy

Create a new mail proxy.
*/
func (a *Client) MailProxyCreate(params *MailProxyCreateParams) (*MailProxyCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyCreate",
		Method:             "POST",
		PathPattern:        "/mailProxy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MailProxyCreateOK), nil

}

/*
MailProxyDelete deletes mail proxy

Delete an existing mail proxy.
*/
func (a *Client) MailProxyDelete(params *MailProxyDeleteParams) (*MailProxyDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyDelete",
		Method:             "DELETE",
		PathPattern:        "/mailProxy/{domain}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MailProxyDeleteOK), nil

}

/*
MailProxyInfo gets mail proxy

Inquire the data for the specified mail proxy.
*/
func (a *Client) MailProxyInfo(params *MailProxyInfoParams) (*MailProxyInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyInfo",
		Method:             "GET",
		PathPattern:        "/mailProxy/{domain}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MailProxyInfoOK), nil

}

/*
MailProxyList lists mail proxy

Inquire a list of mail proxies with certain details. The data to be displayed can be extended per url paremeter.
*/
func (a *Client) MailProxyList(params *MailProxyListParams) (*MailProxyListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyList",
		Method:             "POST",
		PathPattern:        "/mailProxy/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MailProxyListOK), nil

}

/*
MailProxyUpdate updates mail proxy

Update an existing mail proxy.
*/
func (a *Client) MailProxyUpdate(params *MailProxyUpdateParams) (*MailProxyUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyUpdate",
		Method:             "PUT",
		PathPattern:        "/mailProxy/{domain}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*MailProxyUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}