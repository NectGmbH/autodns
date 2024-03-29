// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SslContact ssl contact
// swagger:model SslContact
type SslContact struct {

	// The address of the contact.
	Address []string `json:"address"`

	// The city of the contact
	City string `json:"city,omitempty"`

	// The country of the contact
	Country string `json:"country,omitempty"`

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The email of the contact
	// Required: true
	Email *string `json:"email"`

	// The fax number of the contact
	Fax string `json:"fax,omitempty"`

	// The first name of the contact
	Fname string `json:"fname,omitempty"`

	// Unique identifier of the object
	// Required: true
	ID *int32 `json:"id"`

	// The last name of the contact
	Lname string `json:"lname,omitempty"`

	// The name of the organization
	// Required: true
	Organization *string `json:"organization"`

	// The owner of the object
	Owner *BasicUser `json:"owner,omitempty"`

	// The pcode of the contact
	Pcode string `json:"pcode,omitempty"`

	// The phone number of the contact
	// Required: true
	Phone *string `json:"phone"`

	// The local country state of the contact
	State string `json:"state,omitempty"`

	// The title of the contact
	// Required: true
	Title *string `json:"title"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updating using of the object
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this ssl contact
func (m *SslContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdater(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SslContact) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateOrganization(formats strfmt.Registry) error {

	if err := validate.Required("organization", "body", m.Organization); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *SslContact) validatePhone(formats strfmt.Registry) error {

	if err := validate.Required("phone", "body", m.Phone); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateUpdater(formats strfmt.Registry) error {

	if swag.IsZero(m.Updater) { // not required
		return nil
	}

	if m.Updater != nil {
		if err := m.Updater.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updater")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SslContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SslContact) UnmarshalBinary(b []byte) error {
	var res SslContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
