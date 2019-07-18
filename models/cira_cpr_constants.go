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

// CiraCprConstants cira cpr constants
// swagger:model CiraCprConstants
type CiraCprConstants string

const (

	// CiraCprConstantsCCT captures enum value "CCT"
	CiraCprConstantsCCT CiraCprConstants = "CCT"

	// CiraCprConstantsRES captures enum value "RES"
	CiraCprConstantsRES CiraCprConstants = "RES"

	// CiraCprConstantsCCO captures enum value "CCO"
	CiraCprConstantsCCO CiraCprConstants = "CCO"

	// CiraCprConstantsABO captures enum value "ABO"
	CiraCprConstantsABO CiraCprConstants = "ABO"

	// CiraCprConstantsTDM captures enum value "TDM"
	CiraCprConstantsTDM CiraCprConstants = "TDM"

	// CiraCprConstantsMAJ captures enum value "MAJ"
	CiraCprConstantsMAJ CiraCprConstants = "MAJ"

	// CiraCprConstantsGOV captures enum value "GOV"
	CiraCprConstantsGOV CiraCprConstants = "GOV"

	// CiraCprConstantsLGR captures enum value "LGR"
	CiraCprConstantsLGR CiraCprConstants = "LGR"

	// CiraCprConstantsTRS captures enum value "TRS"
	CiraCprConstantsTRS CiraCprConstants = "TRS"

	// CiraCprConstantsPRT captures enum value "PRT"
	CiraCprConstantsPRT CiraCprConstants = "PRT"

	// CiraCprConstantsASS captures enum value "ASS"
	CiraCprConstantsASS CiraCprConstants = "ASS"

	// CiraCprConstantsTRD captures enum value "TRD"
	CiraCprConstantsTRD CiraCprConstants = "TRD"

	// CiraCprConstantsPLT captures enum value "PLT"
	CiraCprConstantsPLT CiraCprConstants = "PLT"

	// CiraCprConstantsEDU captures enum value "EDU"
	CiraCprConstantsEDU CiraCprConstants = "EDU"

	// CiraCprConstantsLAM captures enum value "LAM"
	CiraCprConstantsLAM CiraCprConstants = "LAM"

	// CiraCprConstantsHOP captures enum value "HOP"
	CiraCprConstantsHOP CiraCprConstants = "HOP"

	// CiraCprConstantsINB captures enum value "INB"
	CiraCprConstantsINB CiraCprConstants = "INB"

	// CiraCprConstantsOMK captures enum value "OMK"
	CiraCprConstantsOMK CiraCprConstants = "OMK"
)

// for schema
var ciraCprConstantsEnum []interface{}

func init() {
	var res []CiraCprConstants
	if err := json.Unmarshal([]byte(`["CCT","RES","CCO","ABO","TDM","MAJ","GOV","LGR","TRS","PRT","ASS","TRD","PLT","EDU","LAM","HOP","INB","OMK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ciraCprConstantsEnum = append(ciraCprConstantsEnum, v)
	}
}

func (m CiraCprConstants) validateCiraCprConstantsEnum(path, location string, value CiraCprConstants) error {
	if err := validate.Enum(path, location, value, ciraCprConstantsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this cira cpr constants
func (m CiraCprConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCiraCprConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
