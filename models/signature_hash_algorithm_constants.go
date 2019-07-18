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

// SignatureHashAlgorithmConstants signature hash algorithm constants
// swagger:model SignatureHashAlgorithmConstants
type SignatureHashAlgorithmConstants string

const (

	// SignatureHashAlgorithmConstantsSHA1 captures enum value "SHA1"
	SignatureHashAlgorithmConstantsSHA1 SignatureHashAlgorithmConstants = "SHA1"

	// SignatureHashAlgorithmConstantsSHA2 captures enum value "SHA2"
	SignatureHashAlgorithmConstantsSHA2 SignatureHashAlgorithmConstants = "SHA2"

	// SignatureHashAlgorithmConstantsSHA2FULLCHAIN captures enum value "SHA2_FULL_CHAIN"
	SignatureHashAlgorithmConstantsSHA2FULLCHAIN SignatureHashAlgorithmConstants = "SHA2_FULL_CHAIN"

	// SignatureHashAlgorithmConstantsUNKNOWN captures enum value "UNKNOWN"
	SignatureHashAlgorithmConstantsUNKNOWN SignatureHashAlgorithmConstants = "UNKNOWN"

	// SignatureHashAlgorithmConstantsSHA384 captures enum value "SHA384"
	SignatureHashAlgorithmConstantsSHA384 SignatureHashAlgorithmConstants = "SHA384"

	// SignatureHashAlgorithmConstantsSHA512 captures enum value "SHA512"
	SignatureHashAlgorithmConstantsSHA512 SignatureHashAlgorithmConstants = "SHA512"

	// SignatureHashAlgorithmConstantsSHA256 captures enum value "SHA256"
	SignatureHashAlgorithmConstantsSHA256 SignatureHashAlgorithmConstants = "SHA256"
)

// for schema
var signatureHashAlgorithmConstantsEnum []interface{}

func init() {
	var res []SignatureHashAlgorithmConstants
	if err := json.Unmarshal([]byte(`["SHA1","SHA2","SHA2_FULL_CHAIN","UNKNOWN","SHA384","SHA512","SHA256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		signatureHashAlgorithmConstantsEnum = append(signatureHashAlgorithmConstantsEnum, v)
	}
}

func (m SignatureHashAlgorithmConstants) validateSignatureHashAlgorithmConstantsEnum(path, location string, value SignatureHashAlgorithmConstants) error {
	if err := validate.Enum(path, location, value, signatureHashAlgorithmConstantsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this signature hash algorithm constants
func (m SignatureHashAlgorithmConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSignatureHashAlgorithmConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
