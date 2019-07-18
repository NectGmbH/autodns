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

// AuRegistrantIDTypeConstants au registrant Id type constants
// swagger:model AuRegistrantIdTypeConstants
type AuRegistrantIDTypeConstants string

const (

	// AuRegistrantIDTypeConstantsACN captures enum value "ACN"
	AuRegistrantIDTypeConstantsACN AuRegistrantIDTypeConstants = "ACN"

	// AuRegistrantIDTypeConstantsABN captures enum value "ABN"
	AuRegistrantIDTypeConstantsABN AuRegistrantIDTypeConstants = "ABN"

	// AuRegistrantIDTypeConstantsOTHER captures enum value "OTHER"
	AuRegistrantIDTypeConstantsOTHER AuRegistrantIDTypeConstants = "OTHER"
)

// for schema
var auRegistrantIdTypeConstantsEnum []interface{}

func init() {
	var res []AuRegistrantIDTypeConstants
	if err := json.Unmarshal([]byte(`["ACN","ABN","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auRegistrantIdTypeConstantsEnum = append(auRegistrantIdTypeConstantsEnum, v)
	}
}

func (m AuRegistrantIDTypeConstants) validateAuRegistrantIDTypeConstantsEnum(path, location string, value AuRegistrantIDTypeConstants) error {
	if err := validate.Enum(path, location, value, auRegistrantIdTypeConstantsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this au registrant Id type constants
func (m AuRegistrantIDTypeConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuRegistrantIDTypeConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
