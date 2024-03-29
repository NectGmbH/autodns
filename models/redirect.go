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

// Redirect redirect
// swagger:model Redirect
type Redirect struct {

	// Lorem Ipsum
	// Max Items: 3
	// Min Items: 0
	Backups []string `json:"backups"`

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The domain, which the redirect belongs.
	Domain string `json:"domain,omitempty"`

	// The last dns connection.
	// Format: date-time
	LastSeen strfmt.DateTime `json:"lastSeen,omitempty"`

	// Redirect mode
	// Required: true
	Mode RedirectModeConstants `json:"mode"`

	// The owner of the object.
	Owner *BasicUser `json:"owner,omitempty"`

	// The domain to be redirected. Enter the domain with or without "www".
	Source string `json:"source,omitempty"`

	// The IDN version of the domain to be redirected. Enter the domain with or without "www".
	SourceIdn string `json:"sourceIdn,omitempty"`

	// The URL of the target domain. Enter the domain without "http://"
	Target string `json:"target,omitempty"`

	// The IDN version of the URL of the target domain. Enter the domain without "http://"
	TargetIdn string `json:"targetIdn,omitempty"`

	// Lorem Ipsum
	Title string `json:"title,omitempty"`

	// Redirect type
	// Required: true
	Type RedirectTypeConstants `json:"type"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updater of the object.
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this redirect
func (m *Redirect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastSeen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *Redirect) validateBackups(formats strfmt.Registry) error {

	if swag.IsZero(m.Backups) { // not required
		return nil
	}

	iBackupsSize := int64(len(m.Backups))

	if err := validate.MinItems("backups", "body", iBackupsSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("backups", "body", iBackupsSize, 3); err != nil {
		return err
	}

	return nil
}

func (m *Redirect) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Redirect) validateLastSeen(formats strfmt.Registry) error {

	if swag.IsZero(m.LastSeen) { // not required
		return nil
	}

	if err := validate.FormatOf("lastSeen", "body", "date-time", m.LastSeen.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Redirect) validateMode(formats strfmt.Registry) error {

	if err := m.Mode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("mode")
		}
		return err
	}

	return nil
}

func (m *Redirect) validateOwner(formats strfmt.Registry) error {

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

func (m *Redirect) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *Redirect) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Redirect) validateUpdater(formats strfmt.Registry) error {

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
func (m *Redirect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Redirect) UnmarshalBinary(b []byte) error {
	var res Redirect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
