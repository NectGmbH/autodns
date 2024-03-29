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

// TimeUnitConstants time unit constants
// swagger:model TimeUnitConstants
type TimeUnitConstants string

const (

	// TimeUnitConstantsMILLISECOND captures enum value "MILLISECOND"
	TimeUnitConstantsMILLISECOND TimeUnitConstants = "MILLISECOND"

	// TimeUnitConstantsSECOND captures enum value "SECOND"
	TimeUnitConstantsSECOND TimeUnitConstants = "SECOND"

	// TimeUnitConstantsMINUTE captures enum value "MINUTE"
	TimeUnitConstantsMINUTE TimeUnitConstants = "MINUTE"

	// TimeUnitConstantsHOUR captures enum value "HOUR"
	TimeUnitConstantsHOUR TimeUnitConstants = "HOUR"

	// TimeUnitConstantsDAY captures enum value "DAY"
	TimeUnitConstantsDAY TimeUnitConstants = "DAY"

	// TimeUnitConstantsWEEK captures enum value "WEEK"
	TimeUnitConstantsWEEK TimeUnitConstants = "WEEK"

	// TimeUnitConstantsMONTH captures enum value "MONTH"
	TimeUnitConstantsMONTH TimeUnitConstants = "MONTH"

	// TimeUnitConstantsQUARTER captures enum value "QUARTER"
	TimeUnitConstantsQUARTER TimeUnitConstants = "QUARTER"

	// TimeUnitConstantsYEAR captures enum value "YEAR"
	TimeUnitConstantsYEAR TimeUnitConstants = "YEAR"
)

// for schema
var timeUnitConstantsEnum []interface{}

func init() {
	var res []TimeUnitConstants
	if err := json.Unmarshal([]byte(`["MILLISECOND","SECOND","MINUTE","HOUR","DAY","WEEK","MONTH","QUARTER","YEAR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeUnitConstantsEnum = append(timeUnitConstantsEnum, v)
	}
}

func (m TimeUnitConstants) validateTimeUnitConstantsEnum(path, location string, value TimeUnitConstants) error {
	if err := validate.Enum(path, location, value, timeUnitConstantsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this time unit constants
func (m TimeUnitConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTimeUnitConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
