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

// DomainStudioDomainSource domain studio domain source
// swagger:model DomainStudioDomainSource
type DomainStudioDomainSource string

const (

	// DomainStudioDomainSourceINITIAL captures enum value "INITIAL"
	DomainStudioDomainSourceINITIAL DomainStudioDomainSource = "INITIAL"

	// DomainStudioDomainSourceSUGGESTION captures enum value "SUGGESTION"
	DomainStudioDomainSourceSUGGESTION DomainStudioDomainSource = "SUGGESTION"

	// DomainStudioDomainSourcePREMIUM captures enum value "PREMIUM"
	DomainStudioDomainSourcePREMIUM DomainStudioDomainSource = "PREMIUM"

	// DomainStudioDomainSourceGEO captures enum value "GEO"
	DomainStudioDomainSourceGEO DomainStudioDomainSource = "GEO"

	// DomainStudioDomainSourceSIMILAR captures enum value "SIMILAR"
	DomainStudioDomainSourceSIMILAR DomainStudioDomainSource = "SIMILAR"

	// DomainStudioDomainSourceRECOMMENDED captures enum value "RECOMMENDED"
	DomainStudioDomainSourceRECOMMENDED DomainStudioDomainSource = "RECOMMENDED"
)

// for schema
var domainStudioDomainSourceEnum []interface{}

func init() {
	var res []DomainStudioDomainSource
	if err := json.Unmarshal([]byte(`["INITIAL","SUGGESTION","PREMIUM","GEO","SIMILAR","RECOMMENDED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainStudioDomainSourceEnum = append(domainStudioDomainSourceEnum, v)
	}
}

func (m DomainStudioDomainSource) validateDomainStudioDomainSourceEnum(path, location string, value DomainStudioDomainSource) error {
	if err := validate.Enum(path, location, value, domainStudioDomainSourceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this domain studio domain source
func (m DomainStudioDomainSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDomainStudioDomainSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
