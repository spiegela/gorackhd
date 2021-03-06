// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostNodeWorkflowOptions post node workflow options
// swagger:model postNodeWorkflowOptions

type PostNodeWorkflowOptions struct {

	// defaults
	Defaults interface{} `json:"defaults,omitempty"`
}

/* polymorph postNodeWorkflowOptions defaults false */

// Validate validates this post node workflow options
func (m *PostNodeWorkflowOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PostNodeWorkflowOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostNodeWorkflowOptions) UnmarshalBinary(b []byte) error {
	var res PostNodeWorkflowOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
