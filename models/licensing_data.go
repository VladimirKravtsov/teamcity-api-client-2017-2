// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LicensingData licensing data
// swagger:model licensingData
type LicensingData struct {

	// agents left
	AgentsLeft int32 `json:"agentsLeft,omitempty" xml:"agentsLeft"`

	// build types left
	BuildTypesLeft int32 `json:"buildTypesLeft,omitempty" xml:"buildTypesLeft"`

	// license keys
	LicenseKeys *LicenseKeys `json:"licenseKeys,omitempty"`

	// license use exceeded
	LicenseUseExceeded *bool `json:"licenseUseExceeded,omitempty" xml:"licenseUseExceeded"`

	// max agents
	MaxAgents int32 `json:"maxAgents,omitempty" xml:"maxAgents"`

	// max build types
	MaxBuildTypes int32 `json:"maxBuildTypes,omitempty" xml:"maxBuildTypes"`

	// server effective release date
	ServerEffectiveReleaseDate string `json:"serverEffectiveReleaseDate,omitempty" xml:"serverEffectiveReleaseDate"`

	// server license type
	ServerLicenseType string `json:"serverLicenseType,omitempty" xml:"serverLicenseType"`

	// unlimited agents
	UnlimitedAgents *bool `json:"unlimitedAgents,omitempty" xml:"unlimitedAgents"`

	// unlimited build types
	UnlimitedBuildTypes *bool `json:"unlimitedBuildTypes,omitempty" xml:"unlimitedBuildTypes"`
}

// Validate validates this licensing data
func (m *LicensingData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLicenseKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LicensingData) validateLicenseKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.LicenseKeys) { // not required
		return nil
	}

	if m.LicenseKeys != nil {
		if err := m.LicenseKeys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("licenseKeys")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LicensingData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LicensingData) UnmarshalBinary(b []byte) error {
	var res LicensingData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
