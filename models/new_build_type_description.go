// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewBuildTypeDescription new build type description
// swagger:model newBuildTypeDescription
type NewBuildTypeDescription struct {

	// build types ids map
	BuildTypesIdsMap *Properties `json:"buildTypesIdsMap,omitempty"`

	// copy all associated settings
	CopyAllAssociatedSettings *bool `json:"copyAllAssociatedSettings,omitempty" xml:"copyAllAssociatedSettings"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// projects ids map
	ProjectsIdsMap *Properties `json:"projectsIdsMap,omitempty"`

	// source build type
	SourceBuildType *BuildType `json:"sourceBuildType,omitempty"`

	// source build type locator
	SourceBuildTypeLocator string `json:"sourceBuildTypeLocator,omitempty" xml:"sourceBuildTypeLocator"`

	// vcs roots ids map
	VcsRootsIdsMap *Properties `json:"vcsRootsIdsMap,omitempty"`
}

// Validate validates this new build type description
func (m *NewBuildTypeDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildTypesIdsMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectsIdsMap(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceBuildType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsRootsIdsMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewBuildTypeDescription) validateBuildTypesIdsMap(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildTypesIdsMap) { // not required
		return nil
	}

	if m.BuildTypesIdsMap != nil {
		if err := m.BuildTypesIdsMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildTypesIdsMap")
			}
			return err
		}
	}

	return nil
}

func (m *NewBuildTypeDescription) validateProjectsIdsMap(formats strfmt.Registry) error {

	if swag.IsZero(m.ProjectsIdsMap) { // not required
		return nil
	}

	if m.ProjectsIdsMap != nil {
		if err := m.ProjectsIdsMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("projectsIdsMap")
			}
			return err
		}
	}

	return nil
}

func (m *NewBuildTypeDescription) validateSourceBuildType(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceBuildType) { // not required
		return nil
	}

	if m.SourceBuildType != nil {
		if err := m.SourceBuildType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sourceBuildType")
			}
			return err
		}
	}

	return nil
}

func (m *NewBuildTypeDescription) validateVcsRootsIdsMap(formats strfmt.Registry) error {

	if swag.IsZero(m.VcsRootsIdsMap) { // not required
		return nil
	}

	if m.VcsRootsIdsMap != nil {
		if err := m.VcsRootsIdsMap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsRootsIdsMap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewBuildTypeDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewBuildTypeDescription) UnmarshalBinary(b []byte) error {
	var res NewBuildTypeDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
