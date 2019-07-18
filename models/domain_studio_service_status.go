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

// DomainStudioServiceStatus domain studio service status
// swagger:model DomainStudioServiceStatus
type DomainStudioServiceStatus string

const (

	// DomainStudioServiceStatusRUNNING captures enum value "RUNNING"
	DomainStudioServiceStatusRUNNING DomainStudioServiceStatus = "RUNNING"

	// DomainStudioServiceStatusSUCCESS captures enum value "SUCCESS"
	DomainStudioServiceStatusSUCCESS DomainStudioServiceStatus = "SUCCESS"

	// DomainStudioServiceStatusFAILED captures enum value "FAILED"
	DomainStudioServiceStatusFAILED DomainStudioServiceStatus = "FAILED"
)

// for schema
var domainStudioServiceStatusEnum []interface{}

func init() {
	var res []DomainStudioServiceStatus
	if err := json.Unmarshal([]byte(`["RUNNING","SUCCESS","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainStudioServiceStatusEnum = append(domainStudioServiceStatusEnum, v)
	}
}

func (m DomainStudioServiceStatus) validateDomainStudioServiceStatusEnum(path, location string, value DomainStudioServiceStatus) error {
	if err := validate.Enum(path, location, value, domainStudioServiceStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this domain studio service status
func (m DomainStudioServiceStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDomainStudioServiceStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
