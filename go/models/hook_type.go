// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HookType hook type
//
// swagger:model hookType
type HookType struct {

	// events
	Events []string `json:"events"`

	// fields
	Fields []interface{} `json:"fields"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this hook type
func (m *HookType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HookType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HookType) UnmarshalBinary(b []byte) error {
	var res HookType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
