// Code generated by go-swagger; DO NOT EDIT.

package document_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new document tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for document tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DocumentInfoWithAlias gets document

Fetch the document for the given alias. The alias can be an ID, UUID or a script name like pricelist.xml.
*/
func (a *Client) DocumentInfoWithAlias(params *DocumentInfoWithAliasParams) (*DocumentInfoWithAliasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentInfoWithAliasParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "documentInfoWithAlias",
		Method:             "GET",
		PathPattern:        "/document/{alias}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentInfoWithAliasReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DocumentInfoWithAliasOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
