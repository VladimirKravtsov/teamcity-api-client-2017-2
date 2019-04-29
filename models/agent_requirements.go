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

// AgentRequirements agent requirements
// swagger:model agent-requirements
type AgentRequirements struct {

	// agent requirement
	AgentRequirement []*AgentRequirement `json:"agent-requirement"`

	// count
	Count int32 `json:"count,omitempty" xml:"count"`
}

// Validate validates this agent requirements
func (m *AgentRequirements) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentRequirement(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentRequirements) validateAgentRequirement(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentRequirement) { // not required
		return nil
	}

	for i := 0; i < len(m.AgentRequirement); i++ {
		if swag.IsZero(m.AgentRequirement[i]) { // not required
			continue
		}

		if m.AgentRequirement[i] != nil {
			if err := m.AgentRequirement[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agent-requirement" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentRequirements) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentRequirements) UnmarshalBinary(b []byte) error {
	var res AgentRequirements
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
