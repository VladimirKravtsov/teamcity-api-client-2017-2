// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Tests tests
// swagger:model tests
type Tests struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// default
	Default *bool `json:"default,omitempty"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref"`

	// test
	Test []*Test `json:"test"`
}

// Validate validates this tests
func (m *Tests) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTest(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Tests) validateTest(formats strfmt.Registry) error {

	if swag.IsZero(m.Test) { // not required
		return nil
	}

	for i := 0; i < len(m.Test); i++ {
		if swag.IsZero(m.Test[i]) { // not required
			continue
		}

		if m.Test[i] != nil {
			if err := m.Test[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("test" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Tests) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Tests) UnmarshalBinary(b []byte) error {
	var res Tests
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}