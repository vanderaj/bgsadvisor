// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EBGSConflictSystemFactionV5 e b g s conflict system faction v5
//
// swagger:model EBGSConflictSystemFactionV5
type EBGSConflictSystemFactionV5 struct {

	// days won
	DaysWon int64 `json:"days_won,omitempty"`

	// faction id
	FactionID string `json:"faction_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// name lower
	NameLower string `json:"name_lower,omitempty"`

	// stake
	Stake string `json:"stake,omitempty"`

	// stake lower
	StakeLower string `json:"stake_lower,omitempty"`

	// station id
	StationID string `json:"station_id,omitempty"`
}

// Validate validates this e b g s conflict system faction v5
func (m *EBGSConflictSystemFactionV5) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this e b g s conflict system faction v5 based on context it is used
func (m *EBGSConflictSystemFactionV5) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EBGSConflictSystemFactionV5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EBGSConflictSystemFactionV5) UnmarshalBinary(b []byte) error {
	var res EBGSConflictSystemFactionV5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
