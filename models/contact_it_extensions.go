// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ContactItExtensions contact it extensions
// swagger:model ContactItExtensions
type ContactItExtensions struct {

	// The number of the matching entity.
	EntityType int32 `json:"entityType,omitempty"`
}

// Validate validates this contact it extensions
func (m *ContactItExtensions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactItExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactItExtensions) UnmarshalBinary(b []byte) error {
	var res ContactItExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
