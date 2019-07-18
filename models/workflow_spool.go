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

// WorkflowSpool workflow spool
// swagger:model WorkflowSpool
type WorkflowSpool struct {

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The next execution date of the job.
	// Format: date-time
	Execution strfmt.DateTime `json:"execution,omitempty"`

	// The job id.
	ID int64 `json:"id,omitempty"`

	// The owner of the object.
	Owner *BasicUser `json:"owner,omitempty"`

	// The status of the job.
	Status JobStatusConstants `json:"status,omitempty"`

	// The substatus of the job.
	SubStatus string `json:"subStatus,omitempty"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updater of the object.
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this workflow spool
func (m *WorkflowSpool) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *WorkflowSpool) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSpool) validateExecution(formats strfmt.Registry) error {

	if swag.IsZero(m.Execution) { // not required
		return nil
	}

	if err := validate.FormatOf("execution", "body", "date-time", m.Execution.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSpool) validateOwner(formats strfmt.Registry) error {

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

func (m *WorkflowSpool) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *WorkflowSpool) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkflowSpool) validateUpdater(formats strfmt.Registry) error {

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
func (m *WorkflowSpool) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowSpool) UnmarshalBinary(b []byte) error {
	var res WorkflowSpool
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
