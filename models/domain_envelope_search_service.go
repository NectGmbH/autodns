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

// DomainEnvelopeSearchService domain envelope search service
// swagger:model DomainEnvelopeSearchService
type DomainEnvelopeSearchService string

const (

	// DomainEnvelopeSearchServiceWHOIS captures enum value "WHOIS"
	DomainEnvelopeSearchServiceWHOIS DomainEnvelopeSearchService = "WHOIS"

	// DomainEnvelopeSearchServicePRICE captures enum value "PRICE"
	DomainEnvelopeSearchServicePRICE DomainEnvelopeSearchService = "PRICE"

	// DomainEnvelopeSearchServiceESTIMATION captures enum value "ESTIMATION"
	DomainEnvelopeSearchServiceESTIMATION DomainEnvelopeSearchService = "ESTIMATION"
)

// for schema
var domainEnvelopeSearchServiceEnum []interface{}

func init() {
	var res []DomainEnvelopeSearchService
	if err := json.Unmarshal([]byte(`["WHOIS","PRICE","ESTIMATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainEnvelopeSearchServiceEnum = append(domainEnvelopeSearchServiceEnum, v)
	}
}

func (m DomainEnvelopeSearchService) validateDomainEnvelopeSearchServiceEnum(path, location string, value DomainEnvelopeSearchService) error {
	if err := validate.Enum(path, location, value, domainEnvelopeSearchServiceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this domain envelope search service
func (m DomainEnvelopeSearchService) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDomainEnvelopeSearchServiceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
