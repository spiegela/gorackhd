// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

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

// NewNodesPostParamsWithContext creates a new NodesPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewNodesPostParamsWithContext(ctx context.Context) *NodesPostParams {
	var ()
	return &NodesPostParams{

		Context: ctx,
	}
}

// NewNodesPostParamsWithHTTPClient creates a new NodesPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNodesPostParamsWithHTTPClient(client *http.Client) *NodesPostParams {
	var ()
	return &NodesPostParams{
		HTTPClient: client,
	}
}

/*NodesPostParams contains all the parameters to send to the API endpoint
for the nodes post operation typically these are written to a http.Request
*/
type NodesPostParams struct {

	/*Identifiers
	  The properties of the new node

	*/
	Identifiers *models.PartialNode

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the nodes post params
func (o *NodesPostParams) WithTimeout(timeout time.Duration) *NodesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nodes post params
func (o *NodesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nodes post params
func (o *NodesPostParams) WithContext(ctx context.Context) *NodesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nodes post params
func (o *NodesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nodes post params
func (o *NodesPostParams) WithHTTPClient(client *http.Client) *NodesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nodes post params
func (o *NodesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifiers adds the identifiers to the nodes post params
func (o *NodesPostParams) WithIdentifiers(identifiers *models.PartialNode) *NodesPostParams {
	o.SetIdentifiers(identifiers)
	return o
}

// SetIdentifiers adds the identifiers to the nodes post params
func (o *NodesPostParams) SetIdentifiers(identifiers *models.PartialNode) {
	o.Identifiers = identifiers
}

// WriteToRequest writes these params to a swagger request
func (o *NodesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Identifiers == nil {
		o.Identifiers = new(models.PartialNode)
	}

	if err := r.SetBodyParam(o.Identifiers); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
