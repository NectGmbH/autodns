// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContactUkExtensions contact uk extensions
// swagger:model ContactUkExtensions
type ContactUkExtensions struct {

	// The entity type.
	EntityType UkTypeConstants `json:"entityType,omitempty"`
}

// Validate validates this contact uk extensions
func (m *ContactUkExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactUkExtensions) validateEntityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	if err := m.EntityType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("entityType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactUkExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactUkExtensions) UnmarshalBinary(b []byte) error {
	var res ContactUkExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
