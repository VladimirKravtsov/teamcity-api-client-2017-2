// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BuildType build type
// swagger:model buildType
type BuildType struct {

	// agent requirements
	AgentRequirements *AgentRequirements `json:"agent-requirements,omitempty"`

	// artifact dependencies
	ArtifactDependencies *ArtifactDependencies `json:"artifact-dependencies,omitempty"`

	// branches
	Branches *Branches `json:"branches,omitempty"`

	// builds
	Builds *Builds `json:"builds,omitempty"`

	// compatible agents
	CompatibleAgents *Agents `json:"compatibleAgents,omitempty"`

	// description
	Description string `json:"description,omitempty" xml:"description"`

	// features
	Features *Features `json:"features,omitempty"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// id
	ID string `json:"id,omitempty" xml:"id"`

	// inherited
	Inherited *bool `json:"inherited,omitempty" xml:"inherited"`

	// internal Id
	InternalID string `json:"internalId,omitempty" xml:"internalId"`

	// investigations
	Investigations *Investigations `json:"investigations,omitempty"`

	// links
	Links *Links `json:"links,omitempty"`

	// locator
	Locator string `json:"locator,omitempty" xml:"locator"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// parameters
	Parameters *Properties `json:"parameters,omitempty"`

	// paused
	Paused *bool `json:"paused,omitempty" xml:"paused"`

	// project
	Project *Project `json:"project,omitempty"`

	// project Id
	ProjectID string `json:"projectId,omitempty" xml:"projectId"`

	// project internal Id
	ProjectInternalID string `json:"projectInternalId,omitempty" xml:"projectInternalId"`

	// project name
	ProjectName string `json:"projectName,omitempty" xml:"projectName"`

	// settings
	Settings *Properties `json:"settings,omitempty"`

	// snapshot dependencies
	SnapshotDependencies *SnapshotDependencies `json:"snapshot-dependencies,omitempty"`

	// steps
	Steps *Steps `json:"steps,omitempty"`

	// template
	Template *BuildType `json:"template,omitempty"`

	// template flag
	TemplateFlag *bool `json:"templateFlag,omitempty" xml:"templateFlag"`

	// templates
	Templates *BuildTypes `json:"templates,omitempty"`

	// triggers
	Triggers *Triggers `json:"triggers,omitempty"`

	// type
	Type string `json:"type,omitempty" xml:"type"`

	// uuid
	UUID string `json:"uuid,omitempty" xml:"uuid"`

	// vcs root entries
	VcsRootEntries *VcsRootEntries `json:"vcs-root-entries,omitempty"`

	// web Url
	WebURL string `json:"webUrl,omitempty" xml:"webUrl"`
}

// Validate validates this build type
func (m *BuildType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentRequirements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArtifactDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBranches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompatibleAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvestigations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSteps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcsRootEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildType) validateAgentRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentRequirements) { // not required
		return nil
	}

	if m.AgentRequirements != nil {
		if err := m.AgentRequirements.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent-requirements")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateArtifactDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.ArtifactDependencies) { // not required
		return nil
	}

	if m.ArtifactDependencies != nil {
		if err := m.ArtifactDependencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("artifact-dependencies")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateBranches(formats strfmt.Registry) error {

	if swag.IsZero(m.Branches) { // not required
		return nil
	}

	if m.Branches != nil {
		if err := m.Branches.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("branches")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateBuilds(formats strfmt.Registry) error {

	if swag.IsZero(m.Builds) { // not required
		return nil
	}

	if m.Builds != nil {
		if err := m.Builds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("builds")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateCompatibleAgents(formats strfmt.Registry) error {

	if swag.IsZero(m.CompatibleAgents) { // not required
		return nil
	}

	if m.CompatibleAgents != nil {
		if err := m.CompatibleAgents.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("compatibleAgents")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateInvestigations(formats strfmt.Registry) error {

	if swag.IsZero(m.Investigations) { // not required
		return nil
	}

	if m.Investigations != nil {
		if err := m.Investigations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("investigations")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateProject(formats strfmt.Registry) error {

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

func (m *BuildType) validateSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateSnapshotDependencies(formats strfmt.Registry) error {

	if swag.IsZero(m.SnapshotDependencies) { // not required
		return nil
	}

	if m.SnapshotDependencies != nil {
		if err := m.SnapshotDependencies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshot-dependencies")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.Steps) { // not required
		return nil
	}

	if m.Steps != nil {
		if err := m.Steps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("steps")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateTemplates(formats strfmt.Registry) error {

	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	if m.Templates != nil {
		if err := m.Templates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templates")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateTriggers(formats strfmt.Registry) error {

	if swag.IsZero(m.Triggers) { // not required
		return nil
	}

	if m.Triggers != nil {
		if err := m.Triggers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggers")
			}
			return err
		}
	}

	return nil
}

func (m *BuildType) validateVcsRootEntries(formats strfmt.Registry) error {

	if swag.IsZero(m.VcsRootEntries) { // not required
		return nil
	}

	if m.VcsRootEntries != nil {
		if err := m.VcsRootEntries.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcs-root-entries")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildType) UnmarshalBinary(b []byte) error {
	var res BuildType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
