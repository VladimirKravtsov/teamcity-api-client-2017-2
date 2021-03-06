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

// TestOccurrences test occurrences
// swagger:model testOccurrences
type TestOccurrences struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// default
	Default *bool `json:"default,omitempty"`

	// failed
	Failed int32 `json:"failed,omitempty" xml:"failed"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// ignored
	Ignored int32 `json:"ignored,omitempty" xml:"ignored"`

	// muted
	Muted int32 `json:"muted,omitempty" xml:"muted"`

	// new failed
	NewFailed int32 `json:"newFailed,omitempty" xml:"newFailed"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref"`

	// passed
	Passed int32 `json:"passed,omitempty" xml:"passed"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref"`

	// test occurrence
	TestOccurrence []*TestOccurrence `json:"testOccurrence"`
}

// Validate validates this test occurrences
func (m *TestOccurrences) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTestOccurrence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestOccurrences) validateTestOccurrence(formats strfmt.Registry) error {

	if swag.IsZero(m.TestOccurrence) { // not required
		return nil
	}

	for i := 0; i < len(m.TestOccurrence); i++ {
		if swag.IsZero(m.TestOccurrence[i]) { // not required
			continue
		}

		if m.TestOccurrence[i] != nil {
			if err := m.TestOccurrence[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("testOccurrence" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestOccurrences) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestOccurrences) UnmarshalBinary(b []byte) error {
	var res TestOccurrences
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
