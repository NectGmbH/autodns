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

// ContactIdentificationExtensions contact identification extensions
// swagger:model ContactIdentificationExtensions
type ContactIdentificationExtensions struct {

	// The authority.
	Authority string `json:"authority,omitempty"`

	// The date of issue.
	// Format: date-time
	DateOfIssue strfmt.DateTime `json:"dateOfIssue,omitempty"`

	// The number.
	Number string `json:"number,omitempty"`

	// The valid till date.
	// Format: date-time
	ValidTill strfmt.DateTime `json:"validTill,omitempty"`
}

// Validate validates this contact identification extensions
func (m *ContactIdentificationExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateOfIssue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidTill(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactIdentificationExtensions) validateDateOfIssue(formats strfmt.Registry) error {

	if swag.IsZero(m.DateOfIssue) { // not required
		return nil
	}

	if err := validate.FormatOf("dateOfIssue", "body", "date-time", m.DateOfIssue.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactIdentificationExtensions) validateValidTill(formats strfmt.Registry) error {

	if swag.IsZero(m.ValidTill) { // not required
		return nil
	}

	if err := validate.FormatOf("validTill", "body", "date-time", m.ValidTill.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactIdentificationExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactIdentificationExtensions) UnmarshalBinary(b []byte) error {
	var res ContactIdentificationExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
