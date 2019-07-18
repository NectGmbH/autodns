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

// ContactTypeConstants contact type constants
// swagger:model ContactTypeConstants
type ContactTypeConstants string

const (

	// ContactTypeConstantsPERSON captures enum value "PERSON"
	ContactTypeConstantsPERSON ContactTypeConstants = "PERSON"

	// ContactTypeConstantsORG captures enum value "ORG"
	ContactTypeConstantsORG ContactTypeConstants = "ORG"

	// ContactTypeConstantsROLE captures enum value "ROLE"
	ContactTypeConstantsROLE ContactTypeConstants = "ROLE"
)

// for schema
var contactTypeConstantsEnum []interface{}

func init() {
	var res []ContactTypeConstants
	if err := json.Unmarshal([]byte(`["PERSON","ORG","ROLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactTypeConstantsEnum = append(contactTypeConstantsEnum, v)
	}
}

func (m ContactTypeConstants) validateContactTypeConstantsEnum(path, location string, value ContactTypeConstants) error {
	if err := validate.Enum(path, location, value, contactTypeConstantsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this contact type constants
func (m ContactTypeConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContactTypeConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
