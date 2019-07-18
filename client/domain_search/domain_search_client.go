// Code generated by go-swagger; DO NOT EDIT.

package domain_search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new domain search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for domain search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DomainSearch powerfuls search api for free domains premium domains and alternate domain names

Configurable search results
*/
func (a *Client) DomainSearch(params *DomainSearchParams) (*DomainSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDomainSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "domainSearch",
		Method:             "POST",
		PathPattern:        "/domainstudio",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DomainSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DomainSearchOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
