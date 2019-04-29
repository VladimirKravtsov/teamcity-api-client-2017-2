// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Investigation investigation
// swagger:model investigation
type Investigation struct {

	// assignee
	Assignee *User `json:"assignee,omitempty"`

	// assignment
	Assignment *Comment `json:"assignment,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// resolution
	Resolution *Resolution `json:"resolution,omitempty"`

	// responsible
	Responsible *User `json:"responsible,omitempty"`

	// scope
	Scope *ProblemScope `json:"scope,omitempty"`

	// state
	State string `json:"state,omitempty" xml:"state"`

	// target
	Target *ProblemTarget `json:"target,omitempty"`
}

// Validate validates this investigation
func (m *Investigation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponsible(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Investigation) validateAssignee(formats strfmt.Registry) error {

	if swag.IsZero(m.Assignee) { // not required
		return nil
	}

	if m.Assignee != nil {
		if err := m.Assignee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignee")
			}
			return err
		}
	}

	return nil
}

func (m *Investigation) validateAssignment(formats strfmt.Registry) error {

	if swag.IsZero(m.Assignment) { // not required
		return nil
	}

	if m.Assignment != nil {
		if err := m.Assignment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignment")
			}
			return err
		}
	}

	return nil
}

func (m *Investigation) validateResolution(formats strfmt.Registry) error {

	if swag.IsZero(m.Resolution) { // not required
		return nil
	}

	if m.Resolution != nil {
		if err := m.Resolution.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resolution")
			}
			return err
		}
	}

	return nil
}

func (m *Investigation) validateResponsible(formats strfmt.Registry) error {

	if swag.IsZero(m.Responsible) { // not required
		return nil
	}

	if m.Responsible != nil {
		if err := m.Responsible.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("responsible")
			}
			return err
		}
	}

	return nil
}

func (m *Investigation) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scope")
			}
			return err
		}
	}

	return nil
}

func (m *Investigation) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	if m.Target != nil {
		if err := m.Target.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Investigation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Investigation) UnmarshalBinary(b []byte) error {
	var res Investigation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
