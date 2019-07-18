// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SubjectAlternativeName subject alternative name
// swagger:model SubjectAlternativeName
type SubjectAlternativeName struct {

	// The approver email of the san.
	ApproverEmail string `json:"approverEmail,omitempty"`

	// The name of the san.
	Name string `json:"name,omitempty"`

	// The order id of the san.
	OrderID string `json:"orderId,omitempty"`
}

// Validate validates this subject alternative name
func (m *SubjectAlternativeName) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubjectAlternativeName) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectAlternativeName) UnmarshalBinary(b []byte) error {
	var res SubjectAlternativeName
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
