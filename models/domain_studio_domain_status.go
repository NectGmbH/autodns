// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DomainStudioDomainStatus domain studio domain status
// swagger:model DomainStudioDomainStatus
type DomainStudioDomainStatus string

const (

	// DomainStudioDomainStatusFREE captures enum value "FREE"
	DomainStudioDomainStatusFREE DomainStudioDomainStatus = "FREE"

	// DomainStudioDomainStatusASSIGNED captures enum value "ASSIGNED"
	DomainStudioDomainStatusASSIGNED DomainStudioDomainStatus = "ASSIGNED"

	// DomainStudioDomainStatusMARKET captures enum value "MARKET"
	DomainStudioDomainStatusMARKET DomainStudioDomainStatus = "MARKET"

	// DomainStudioDomainStatusPREMIUM captures enum value "PREMIUM"
	DomainStudioDomainStatusPREMIUM DomainStudioDomainStatus = "PREMIUM"

	// DomainStudioDomainStatusINVALID captures enum value "INVALID"
	DomainStudioDomainStatusINVALID DomainStudioDomainStatus = "INVALID"

	// DomainStudioDomainStatusERROR captures enum value "ERROR"
	DomainStudioDomainStatusERROR DomainStudioDomainStatus = "ERROR"

	// DomainStudioDomainStatusTIMEOUT captures enum value "TIMEOUT"
	DomainStudioDomainStatusTIMEOUT DomainStudioDomainStatus = "TIMEOUT"
)

// for schema
var domainStudioDomainStatusEnum []interface{}

func init() {
	var res []DomainStudioDomainStatus
	if err := json.Unmarshal([]byte(`["FREE","ASSIGNED","MARKET","PREMIUM","INVALID","ERROR","TIMEOUT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainStudioDomainStatusEnum = append(domainStudioDomainStatusEnum, v)
	}
}

func (m DomainStudioDomainStatus) validateDomainStudioDomainStatusEnum(path, location string, value DomainStudioDomainStatus) error {
	if err := validate.Enum(path, location, value, domainStudioDomainStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this domain studio domain status
func (m DomainStudioDomainStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDomainStudioDomainStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
