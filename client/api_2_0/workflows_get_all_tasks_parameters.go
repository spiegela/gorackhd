package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWorkflowsGetAllTasksParams creates a new WorkflowsGetAllTasksParams object
// with the default values initialized.
func NewWorkflowsGetAllTasksParams() *WorkflowsGetAllTasksParams {

	return &WorkflowsGetAllTasksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowsGetAllTasksParamsWithTimeout creates a new WorkflowsGetAllTasksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowsGetAllTasksParamsWithTimeout(timeout time.Duration) *WorkflowsGetAllTasksParams {

	return &WorkflowsGetAllTasksParams{

		timeout: timeout,
	}
}

/*WorkflowsGetAllTasksParams contains all the parameters to send to the API endpoint
for the workflows get all tasks operation typically these are written to a http.Request
*/
type WorkflowsGetAllTasksParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowsGetAllTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}