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

// CertificateType certificate type
// swagger:model CertificateType
type CertificateType string

const (

	// CertificateTypeFQDN captures enum value "FQDN"
	CertificateTypeFQDN CertificateType = "FQDN"

	// CertificateTypeMAIL captures enum value "MAIL"
	CertificateTypeMAIL CertificateType = "MAIL"

	// CertificateTypeCODE captures enum value "CODE"
	CertificateTypeCODE CertificateType = "CODE"
)

// for schema
var certificateTypeEnum []interface{}

func init() {
	var res []CertificateType
	if err := json.Unmarshal([]byte(`["FQDN","MAIL","CODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateTypeEnum = append(certificateTypeEnum, v)
	}
}

func (m CertificateType) validateCertificateTypeEnum(path, location string, value CertificateType) error {
	if err := validate.Enum(path, location, value, certificateTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this certificate type
func (m CertificateType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCertificateTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
