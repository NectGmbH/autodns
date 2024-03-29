// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CertAuthentication cert authentication
// swagger:model CertAuthentication
type CertAuthentication struct {

	// The approver addresses for email based authentication.
	ApproverEmails []string `json:"approverEmails"`

	// The dns entry for dns based authentication.
	DNS string `json:"dns,omitempty"`

	// The content for file based authentication.
	FileContent string `json:"fileContent,omitempty"`

	// The filename for file based authentication.
	FileName string `json:"fileName,omitempty"`

	// The authentication method.
	Method AuthMethodConstants `json:"method,omitempty"`

	// Activates automatic provisioning of the zone for dns based authentication.
	Provisioning bool `json:"provisioning,omitempty"`
}

// Validate validates this cert authentication
func (m *CertAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertAuthentication) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	if err := m.Method.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("method")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CertAuthentication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertAuthentication) UnmarshalBinary(b []byte) error {
	var res CertAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
