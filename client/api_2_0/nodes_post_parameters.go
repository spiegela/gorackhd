package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NewNodesPostParams creates a new NodesPostParams object
// with the default values initialized.
func NewNodesPostParams() *NodesPostParams {
	var ()
	return &NodesPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNodesPostParamsWithTimeout creates a new NodesPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNodesPostParamsWithTimeout(timeout time.Duration) *NodesPostParams {
	var ()
	return &NodesPostParams{

		timeout: timeout,
	}
}

/*NodesPostParams contains all the parameters to send to the API endpoint
for the nodes post operation typically these are written to a http.Request
*/
type NodesPostParams struct {

	/*Identifiers
	  The properties of the new node

	*/
	Identifiers *models.Node20PartialNode

	timeout time.Duration
}

// WithIdentifiers adds the identifiers to the nodes post params
func (o *NodesPostParams) WithIdentifiers(identifiers *models.Node20PartialNode) *NodesPostParams {
	o.Identifiers = identifiers
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *NodesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Identifiers == nil {
		o.Identifiers = new(models.Node20PartialNode)
	}

	if err := r.SetBodyParam(o.Identifiers); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}