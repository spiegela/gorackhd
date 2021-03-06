// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RelationsObj A key value set of node relations.
// swagger:model relations_obj

type RelationsObj struct {

	// info
	Info string `json:"info,omitempty"`

	// Relation Type with the node.
	RelationType string `json:"relationType,omitempty"`

	// Array of targets.
	Targets []string `json:"targets"`
}

/* polymorph relations_obj info false */

/* polymorph relations_obj relationType false */

/* polymorph relations_obj targets false */

// Validate validates this relations obj
func (m *RelationsObj) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationsObj) validateTargets(formats strfmt.Registry) error {

	if swag.IsZero(m.Targets) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationsObj) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationsObj) UnmarshalBinary(b []byte) error {
	var res RelationsObj
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
