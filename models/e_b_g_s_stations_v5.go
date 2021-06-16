// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EBGSStationsV5 e b g s stations v5
//
// swagger:model EBGSStationsV5
type EBGSStationsV5 struct {

	// v
	V int64 `json:"__v,omitempty"`

	// id
	ID string `json:"_id,omitempty"`

	// all economies
	AllEconomies []*EBGSAllEconomiesV5 `json:"all_economies"`

	// allegiance
	Allegiance string `json:"allegiance,omitempty"`

	// controlling minor faction
	ControllingMinorFaction string `json:"controlling_minor_faction,omitempty"`

	// controlling minor faction cased
	ControllingMinorFactionCased string `json:"controlling_minor_faction_cased,omitempty"`

	// controlling minor faction id
	ControllingMinorFactionID string `json:"controlling_minor_faction_id,omitempty"`

	// distance from star
	DistanceFromStar int64 `json:"distance_from_star,omitempty"`

	// economy
	Economy string `json:"economy,omitempty"`

	// eddb id
	EddbID int64 `json:"eddb_id,omitempty"`

	// government
	Government string `json:"government,omitempty"`

	// history
	History []*EBGSStationHistoryV5 `json:"history"`

	// market id
	MarketID string `json:"market_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// name aliases
	NameAliases []*EBGSNameAliasV5 `json:"name_aliases"`

	// name lower
	NameLower string `json:"name_lower,omitempty"`

	// services
	Services []*EBGSStationServicesV5 `json:"services"`

	// state
	State string `json:"state,omitempty"`

	// system
	System string `json:"system,omitempty"`

	// system id
	SystemID string `json:"system_id,omitempty"`

	// system lower
	SystemLower string `json:"system_lower,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Validate validates this e b g s stations v5
func (m *EBGSStationsV5) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllEconomies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameAliases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EBGSStationsV5) validateAllEconomies(formats strfmt.Registry) error {
	if swag.IsZero(m.AllEconomies) { // not required
		return nil
	}

	for i := 0; i < len(m.AllEconomies); i++ {
		if swag.IsZero(m.AllEconomies[i]) { // not required
			continue
		}

		if m.AllEconomies[i] != nil {
			if err := m.AllEconomies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("all_economies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EBGSStationsV5) validateHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.History) { // not required
		return nil
	}

	for i := 0; i < len(m.History); i++ {
		if swag.IsZero(m.History[i]) { // not required
			continue
		}

		if m.History[i] != nil {
			if err := m.History[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EBGSStationsV5) validateNameAliases(formats strfmt.Registry) error {
	if swag.IsZero(m.NameAliases) { // not required
		return nil
	}

	for i := 0; i < len(m.NameAliases); i++ {
		if swag.IsZero(m.NameAliases[i]) { // not required
			continue
		}

		if m.NameAliases[i] != nil {
			if err := m.NameAliases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("name_aliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EBGSStationsV5) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {
		if swag.IsZero(m.Services[i]) { // not required
			continue
		}

		if m.Services[i] != nil {
			if err := m.Services[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this e b g s stations v5 based on the context it is used
func (m *EBGSStationsV5) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllEconomies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNameAliases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EBGSStationsV5) contextValidateAllEconomies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AllEconomies); i++ {

		if m.AllEconomies[i] != nil {
			if err := m.AllEconomies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("all_economies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EBGSStationsV5) contextValidateHistory(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.History); i++ {

		if m.History[i] != nil {
			if err := m.History[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("history" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EBGSStationsV5) contextValidateNameAliases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NameAliases); i++ {

		if m.NameAliases[i] != nil {
			if err := m.NameAliases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("name_aliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EBGSStationsV5) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Services); i++ {

		if m.Services[i] != nil {
			if err := m.Services[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("services" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EBGSStationsV5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EBGSStationsV5) UnmarshalBinary(b []byte) error {
	var res EBGSStationsV5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}