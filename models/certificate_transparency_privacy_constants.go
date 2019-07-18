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

// CertificateTransparencyPrivacyConstants certificate transparency privacy constants
// swagger:model CertificateTransparencyPrivacyConstants
type CertificateTransparencyPrivacyConstants string

const (

	// CertificateTransparencyPrivacyConstantsPUBLIC captures enum value "PUBLIC"
	CertificateTransparencyPrivacyConstantsPUBLIC CertificateTransparencyPrivacyConstants = "PUBLIC"

	// CertificateTransparencyPrivacyConstantsREDACTED captures enum value "REDACTED"
	CertificateTransparencyPrivacyConstantsREDACTED CertificateTransparencyPrivacyConstants = "REDACTED"
)

// for schema
var certificateTransparencyPrivacyConstantsEnum []interface{}

func init() {
	var res []CertificateTransparencyPrivacyConstants
	if err := json.Unmarshal([]byte(`["PUBLIC","REDACTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateTransparencyPrivacyConstantsEnum = append(certificateTransparencyPrivacyConstantsEnum, v)
	}
}

func (m CertificateTransparencyPrivacyConstants) validateCertificateTransparencyPrivacyConstantsEnum(path, location string, value CertificateTransparencyPrivacyConstants) error {
	if err := validate.Enum(path, location, value, certificateTransparencyPrivacyConstantsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this certificate transparency privacy constants
func (m CertificateTransparencyPrivacyConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCertificateTransparencyPrivacyConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}