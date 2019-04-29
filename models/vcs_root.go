// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VcsRoot vcs root
// swagger:model vcs-root
type VcsRoot struct {

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// internal Id
	InternalID string `json:"internalId,omitempty" xml:"internalId"`

	// locator
	Locator string `json:"locator,omitempty" xml:"locator"`

	// modification check interval
	ModificationCheckInterval int32 `json:"modificationCheckInterval,omitempty" xml:"modificationCheckInterval"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// project
	Project *Project `json:"project,omitempty"`

	// project locator
	ProjectLocator string `json:"projectLocator,omitempty" xml:"projectLocator"`

	// properties
	Properties *Properties `json:"properties,omitempty"`

	// repository Id strings
	RepositoryIDStrings *Items `json:"repositoryIdStrings,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty" xml:"uuid"`

	// vcs name
	VcsName string `json:"vcsName,omitempty" xml:"vcsName"`

	// vcs root instances
	VcsRootInstances *VcsRootInstances `json:"vcsRootInstances,omitempty"`
}

// Validate validates this vcs root
func (m *VcsRoot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepositoryIDStrings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsRootInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsRoot) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *VcsRoot) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

func (m *VcsRoot) validateRepositoryIDStrings(formats strfmt.Registry) error {

	if swag.IsZero(m.RepositoryIDStrings) { // not required
		return nil
	}

	if m.RepositoryIDStrings != nil {
		if err := m.RepositoryIDStrings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repositoryIdStrings")
			}
			return err
		}
	}

	return nil
}

func (m *VcsRoot) validateVcsRootInstances(formats strfmt.Registry) error {

	if swag.IsZero(m.VcsRootInstances) { // not required
		return nil
	}

	if m.VcsRootInstances != nil {
		if err := m.VcsRootInstances.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsRootInstances")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcsRoot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsRoot) UnmarshalBinary(b []byte) error {
	var res VcsRoot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
