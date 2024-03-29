// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DomainServicesUpdate domain services update
// swagger:model DomainServicesUpdate
type DomainServicesUpdate struct {

	// The name of the domain.
	Domains []*Domain `json:"domains"`

	// The services to add.
	ServicesAdd *DomainServices `json:"servicesAdd,omitempty"`

	// The services to remove.
	ServicesRem *DomainServices `json:"servicesRem,omitempty"`
}

// Validate validates this domain services update
func (m *DomainServicesUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicesAdd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServicesRem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainServicesUpdate) validateDomains(formats strfmt.Registry) error {

	if swag.IsZero(m.Domains) { // not required
		return nil
	}

	for i := 0; i < len(m.Domains); i++ {
		if swag.IsZero(m.Domains[i]) { // not required
			continue
		}

		if m.Domains[i] != nil {
			if err := m.Domains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainServicesUpdate) validateServicesAdd(formats strfmt.Registry) error {

	if swag.IsZero(m.ServicesAdd) { // not required
		return nil
	}

	if m.ServicesAdd != nil {
		if err := m.ServicesAdd.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("servicesAdd")
			}
			return err
		}
	}

	return nil
}

func (m *DomainServicesUpdate) validateServicesRem(formats strfmt.Registry) error {

	if swag.IsZero(m.ServicesRem) { // not required
		return nil
	}

	if m.ServicesRem != nil {
		if err := m.ServicesRem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("servicesRem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainServicesUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainServicesUpdate) UnmarshalBinary(b []byte) error {
	var res DomainServicesUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
