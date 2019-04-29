// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AuthorizedInfo authorized info
// swagger:model authorizedInfo
type AuthorizedInfo struct {

	// comment
	Comment *Comment `json:"comment,omitempty"`

	// status
	Status *bool `json:"status,omitempty" xml:"status"`
}

// Validate validates this authorized info
func (m *AuthorizedInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthorizedInfo) validateComment(formats strfmt.Registry) error {

	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if m.Comment != nil {
		if err := m.Comment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("comment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthorizedInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthorizedInfo) UnmarshalBinary(b []byte) error {
	var res AuthorizedInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}