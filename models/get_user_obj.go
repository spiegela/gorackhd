// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetUserObj get user obj
// swagger:model get_user_obj

type GetUserObj struct {

	// role
	Role string `json:"role,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

/* polymorph get_user_obj role false */

/* polymorph get_user_obj username false */

// Validate validates this get user obj
func (m *GetUserObj) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetUserObj) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserObj) UnmarshalBinary(b []byte) error {
	var res GetUserObj
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
