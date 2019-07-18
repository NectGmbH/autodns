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

// IdentityStatus identity status
// swagger:model IdentityStatus
type IdentityStatus string

const (

	// IdentityStatusVERIFY captures enum value "VERIFY"
	IdentityStatusVERIFY IdentityStatus = "VERIFY"

	// IdentityStatusSUCCESS captures enum value "SUCCESS"
	IdentityStatusSUCCESS IdentityStatus = "SUCCESS"
)

// for schema
var identityStatusEnum []interface{}

func init() {
	var res []IdentityStatus
	if err := json.Unmarshal([]byte(`["VERIFY","SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		identityStatusEnum = append(identityStatusEnum, v)
	}
}

func (m IdentityStatus) validateIdentityStatusEnum(path, location string, value IdentityStatus) error {
	if err := validate.Enum(path, location, value, identityStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this identity status
func (m IdentityStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIdentityStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
