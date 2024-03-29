// Code generated by go-swagger; DO NOT EDIT.

package id4me_identity_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new id4me identity tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for id4me identity tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ConfirmId4MeIdentity confirms an existing id4me identity not available yet coming soon bang

Confirms an existing id4me identity
*/
func (a *Client) ConfirmId4MeIdentity(params *ConfirmId4MeIdentityParams) (*ConfirmId4MeIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfirmId4MeIdentityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "confirmId4MeIdentity",
		Method:             "PUT",
		PathPattern:        "/id4MeIdentity/{name}/_confirm",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConfirmId4MeIdentityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConfirmId4MeIdentityOK), nil

}

/*
CreateId4MeIdentity creates a new id4me identity not available yet coming soon bang

Creates a new id4me identity
*/
func (a *Client) CreateId4MeIdentity(params *CreateId4MeIdentityParams) (*CreateId4MeIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateId4MeIdentityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createId4MeIdentity",
		Method:             "POST",
		PathPattern:        "/id4MeIdentity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateId4MeIdentityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateId4MeIdentityOK), nil

}

/*
DeleteId4MeIdentity deletes an existing id4me identity not available yet coming soon bang

Deletes an existing id4me identity
*/
func (a *Client) DeleteId4MeIdentity(params *DeleteId4MeIdentityParams) (*DeleteId4MeIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteId4MeIdentityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteId4MeIdentity",
		Method:             "DELETE",
		PathPattern:        "/id4MeIdentity/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteId4MeIdentityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteId4MeIdentityOK), nil

}

/*
Id4MeIdentityInfo gets an id4me identity not available yet coming soon bang

Fetch the data for the specified name.
*/
func (a *Client) Id4MeIdentityInfo(params *Id4MeIdentityInfoParams) (*Id4MeIdentityInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewId4MeIdentityInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "id4MeIdentityInfo",
		Method:             "GET",
		PathPattern:        "/id4MeIdentity/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Id4MeIdentityInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*Id4MeIdentityInfoOK), nil

}

/*
Id4MeIdentityList id4mes identity list task not available yet coming soon bang

Inquire a list of id4me identities with certain details.
*/
func (a *Client) Id4MeIdentityList(params *Id4MeIdentityListParams) (*Id4MeIdentityListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewId4MeIdentityListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "id4MeIdentityList",
		Method:             "POST",
		PathPattern:        "/id4MeIdentity/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Id4MeIdentityListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*Id4MeIdentityListOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
