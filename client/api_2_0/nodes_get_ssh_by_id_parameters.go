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

// NewNodesGetSSHByIDParams creates a new NodesGetSSHByIDParams object
// with the default values initialized.
func NewNodesGetSSHByIDParams() *NodesGetSSHByIDParams {
	var ()
	return &NodesGetSSHByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesGetSSHByIDParamsWithTimeout creates a new NodesGetSSHByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesGetSSHByIDParamsWithTimeout(timeout time.Duration) *NodesGetSSHByIDParams {
	var ()
	return &NodesGetSSHByIDParams{

		timeout: timeout,
	}
}

/*NodesGetSSHByIDParams contains all the parameters to send to the API endpoint
for the nodes get Ssh by Id operation typically these are written to a http.Request
*/
type NodesGetSSHByIDParams struct {

	/*Identifier
	  The node identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the nodes get Ssh by Id params
func (o *NodesGetSSHByIDParams) WithIdentifier(identifier string) *NodesGetSSHByIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *NodesGetSSHByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}