// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowTask workflow task
// swagger:model workflow_task

type WorkflowTask struct {

	// friendly name
	FriendlyName string `json:"friendlyName,omitempty"`

	// implements task
	ImplementsTask string `json:"implementsTask,omitempty"`

	// injectable name
	InjectableName string `json:"injectableName,omitempty"`

	// options
	Options interface{} `json:"options,omitempty"`

	// properties
	Properties interface{} `json:"properties,omitempty"`
}

/* polymorph workflow_task friendlyName false */

/* polymorph workflow_task implementsTask false */

/* polymorph workflow_task injectableName false */

/* polymorph workflow_task options false */

/* polymorph workflow_task properties false */

// Validate validates this workflow task
func (m *WorkflowTask) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowTask) UnmarshalBinary(b []byte) error {
	var res WorkflowTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
