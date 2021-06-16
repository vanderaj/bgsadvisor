// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EBGSAllEconomiesV5 e b g s all economies v5
//
// swagger:model EBGSAllEconomiesV5
type EBGSAllEconomiesV5 struct {

	// name
	Name string `json:"name,omitempty"`

	// proportion
	Proportion int64 `json:"proportion,omitempty"`
}

// Validate validates this e b g s all economies v5
func (m *EBGSAllEconomiesV5) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this e b g s all economies v5 based on context it is used
func (m *EBGSAllEconomiesV5) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EBGSAllEconomiesV5) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EBGSAllEconomiesV5) UnmarshalBinary(b []byte) error {
	var res EBGSAllEconomiesV5
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
