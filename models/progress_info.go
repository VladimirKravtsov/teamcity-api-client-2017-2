// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ProgressInfo progress info
// swagger:model progress-info
type ProgressInfo struct {

	// current stage text
	CurrentStageText string `json:"currentStageText,omitempty" xml:"currentStageText"`

	// elapsed seconds
	ElapsedSeconds int64 `json:"elapsedSeconds,omitempty" xml:"elapsedSeconds"`

	// estimated total seconds
	EstimatedTotalSeconds int64 `json:"estimatedTotalSeconds,omitempty" xml:"estimatedTotalSeconds"`

	// left seconds
	LeftSeconds int64 `json:"leftSeconds,omitempty" xml:"leftSeconds"`

	// outdated
	Outdated *bool `json:"outdated,omitempty" xml:"outdated"`

	// percentage complete
	PercentageComplete int32 `json:"percentageComplete,omitempty" xml:"percentageComplete"`

	// probably hanging
	ProbablyHanging *bool `json:"probablyHanging,omitempty" xml:"probablyHanging"`
}

// Validate validates this progress info
func (m *ProgressInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProgressInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressInfo) UnmarshalBinary(b []byte) error {
	var res ProgressInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
