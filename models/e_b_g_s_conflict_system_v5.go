// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EBGSConflictSystemV5 e b g s conflict system v5
//
// swagger:model EBGSConflictSystemV5
type EBGSConflictSystemV5 struct {

	// faction1
	Faction1 *EBGSConflictSystemFactionV5 `json:"faction1,omitempty"`

	// faction2
	Faction2 *EBGSConflictSystemFactionV5 `json:"faction2,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this e b g s conflict system v5
func (m *EBGSConflictSystemV5) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFaction1(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFaction2(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EBGSConflictSystemV5) validateFaction1(formats strfmt.Registry) error {
	if swag.IsZero(m.Faction1) { // not required
		return nil
	}

	if m.Faction1 != nil {
		if err := m.Faction1.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("faction1")
			}
			return err
		}
	}

	return nil
}

func (m *EBGSConflictSystemV5) validateFaction2(formats strfmt.Registry) error {
	if swag.IsZero(m.Faction2) { // not required
		return nil
	}

	if m.Faction2 != nil {
		if err := m.Faction2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("faction2")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this e b g s conflict system v5 based on the context it is used
func (m *EBGSConflictSystemV5) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFaction1(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFaction2(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EBGSConflictSystemV5) contextValidateFaction1(ctx context.Context, formats strfmt.Registry) error {

	if m.Faction1 != nil {
		if err := m.Faction1.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("faction1")
			}
			return err
		}
	}

	return nil
}

func (m *EBGSConflictSystemV5) contextValidateFaction2(ctx context.Context, formats strfmt.Registry) error {

	if m.Faction2 != nil {
		if err := m.Faction2.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("faction2")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EBGSConflictSystemV5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EBGSConflictSystemV5) UnmarshalBinary(b []byte) error {
	var res EBGSConflictSystemV5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}