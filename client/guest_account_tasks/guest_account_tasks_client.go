// Code generated by go-swagger; DO NOT EDIT.

package guest_account_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new guest account tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for guest account tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GuestApplyVerify verifies guest account

applies a verification code for a guest account.
*/
func (a *Client) GuestApplyVerify(params *GuestApplyVerifyParams) (*GuestApplyVerifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGuestApplyVerifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "guestApplyVerify",
		Method:             "GET",
		PathPattern:        "/user/_guestverify/{token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GuestApplyVerifyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GuestApplyVerifyOK), nil

}

/*
GuestCreate creates guest account

creates a guest account.
*/
func (a *Client) GuestCreate(params *GuestCreateParams) (*GuestCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGuestCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "guestCreate",
		Method:             "POST",
		PathPattern:        "/user/_guest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GuestCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GuestCreateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
