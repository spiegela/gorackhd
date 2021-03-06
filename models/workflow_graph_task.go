// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// WorkflowGraphTask workflow graph task
// swagger:model workflow_graph_task

type WorkflowGraphTask struct {

	// ignore failure
	IgnoreFailure bool `json:"ignoreFailure,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// task name
	TaskName string `json:"taskName,omitempty"`

	// wait on
	WaitOn interface{} `json:"waitOn,omitempty"`
}

/* polymorph workflow_graph_task ignoreFailure false */

/* polymorph workflow_graph_task label false */

/* polymorph workflow_graph_task taskName false */

/* polymorph workflow_graph_task waitOn false */

// Validate validates this workflow graph task
func (m *WorkflowGraphTask) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowGraphTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowGraphTask) UnmarshalBinary(b []byte) error {
	var res WorkflowGraphTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
