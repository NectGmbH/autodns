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

// HkDocumentTypeConstants hk document type constants
// swagger:model HkDocumentTypeConstants
type HkDocumentTypeConstants string

const (

	// HkDocumentTypeConstantsHKID captures enum value "HKID"
	HkDocumentTypeConstantsHKID HkDocumentTypeConstants = "HKID"

	// HkDocumentTypeConstantsOTHID captures enum value "OTHID"
	HkDocumentTypeConstantsOTHID HkDocumentTypeConstants = "OTHID"

	// HkDocumentTypeConstantsPASSNO captures enum value "PASSNO"
	HkDocumentTypeConstantsPASSNO HkDocumentTypeConstants = "PASSNO"

	// HkDocumentTypeConstantsBIRTHCERT captures enum value "BIRTHCERT"
	HkDocumentTypeConstantsBIRTHCERT HkDocumentTypeConstants = "BIRTHCERT"

	// HkDocumentTypeConstantsOTHIDV captures enum value "OTHIDV"
	HkDocumentTypeConstantsOTHIDV HkDocumentTypeConstants = "OTHIDV"

	// HkDocumentTypeConstantsBR captures enum value "BR"
	HkDocumentTypeConstantsBR HkDocumentTypeConstants = "BR"

	// HkDocumentTypeConstantsCI captures enum value "CI"
	HkDocumentTypeConstantsCI HkDocumentTypeConstants = "CI"

	// HkDocumentTypeConstantsCRS captures enum value "CRS"
	HkDocumentTypeConstantsCRS HkDocumentTypeConstants = "CRS"

	// HkDocumentTypeConstantsHKSARG captures enum value "HKSARG"
	HkDocumentTypeConstantsHKSARG HkDocumentTypeConstants = "HKSARG"

	// HkDocumentTypeConstantsHKORDINANCE captures enum value "HKORDINANCE"
	HkDocumentTypeConstantsHKORDINANCE HkDocumentTypeConstants = "HKORDINANCE"

	// HkDocumentTypeConstantsOTHORG captures enum value "OTHORG"
	HkDocumentTypeConstantsOTHORG HkDocumentTypeConstants = "OTHORG"
)

// for schema
var hkDocumentTypeConstantsEnum []interface{}

func init() {
	var res []HkDocumentTypeConstants
	if err := json.Unmarshal([]byte(`["HKID","OTHID","PASSNO","BIRTHCERT","OTHIDV","BR","CI","CRS","HKSARG","HKORDINANCE","OTHORG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hkDocumentTypeConstantsEnum = append(hkDocumentTypeConstantsEnum, v)
	}
}

func (m HkDocumentTypeConstants) validateHkDocumentTypeConstantsEnum(path, location string, value HkDocumentTypeConstants) error {
	if err := validate.Enum(path, location, value, hkDocumentTypeConstantsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this hk document type constants
func (m HkDocumentTypeConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHkDocumentTypeConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}