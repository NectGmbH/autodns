// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContactBirthExtensions contact birth extensions
// swagger:model ContactBirthExtensions
type ContactBirthExtensions struct {

	// The country.
	Country string `json:"country,omitempty"`

	// The day.
	// Format: date-time
	Day strfmt.DateTime `json:"day,omitempty"`

	// The place.
	Place string `json:"place,omitempty"`

	// The postalCode.
	PostalCode string `json:"postalCode,omitempty"`
}

// Validate validates this contact birth extensions
func (m *ContactBirthExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactBirthExtensions) validateDay(formats strfmt.Registry) error {

	if swag.IsZero(m.Day) { // not required
		return nil
	}

	if err := validate.FormatOf("day", "body", "date-time", m.Day.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactBirthExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactBirthExtensions) UnmarshalBinary(b []byte) error {
	var res ContactBirthExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}