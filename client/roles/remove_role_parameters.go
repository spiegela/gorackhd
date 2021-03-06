// Code generated by go-swagger; DO NOT EDIT.

package roles

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
)

// NewRemoveRoleParams creates a new RemoveRoleParams object
// with the default values initialized.
func NewRemoveRoleParams() *RemoveRoleParams {
	var ()
	return &RemoveRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveRoleParamsWithTimeout creates a new RemoveRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveRoleParamsWithTimeout(timeout time.Duration) *RemoveRoleParams {
	var ()
	return &RemoveRoleParams{

		timeout: timeout,
	}
}

// NewRemoveRoleParamsWithContext creates a new RemoveRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveRoleParamsWithContext(ctx context.Context) *RemoveRoleParams {
	var ()
	return &RemoveRoleParams{

		Context: ctx,
	}
}

// NewRemoveRoleParamsWithHTTPClient creates a new RemoveRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveRoleParamsWithHTTPClient(client *http.Client) *RemoveRoleParams {
	var ()
	return &RemoveRoleParams{
		HTTPClient: client,
	}
}

/*RemoveRoleParams contains all the parameters to send to the API endpoint
for the remove role operation typically these are written to a http.Request
*/
type RemoveRoleParams struct {

	/*Name
	  role name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove role params
func (o *RemoveRoleParams) WithTimeout(timeout time.Duration) *RemoveRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove role params
func (o *RemoveRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove role params
func (o *RemoveRoleParams) WithContext(ctx context.Context) *RemoveRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove role params
func (o *RemoveRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove role params
func (o *RemoveRoleParams) WithHTTPClient(client *http.Client) *RemoveRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove role params
func (o *RemoveRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the remove role params
func (o *RemoveRoleParams) WithName(name string) *RemoveRoleParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the remove role params
func (o *RemoveRoleParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
