// Code generated by go-swagger; DO NOT EDIT.

package hooks

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

// NewHooksGetByIDParams creates a new HooksGetByIDParams object
// with the default values initialized.
func NewHooksGetByIDParams() *HooksGetByIDParams {
	var ()
	return &HooksGetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHooksGetByIDParamsWithTimeout creates a new HooksGetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHooksGetByIDParamsWithTimeout(timeout time.Duration) *HooksGetByIDParams {
	var ()
	return &HooksGetByIDParams{

		timeout: timeout,
	}
}

// NewHooksGetByIDParamsWithContext creates a new HooksGetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewHooksGetByIDParamsWithContext(ctx context.Context) *HooksGetByIDParams {
	var ()
	return &HooksGetByIDParams{

		Context: ctx,
	}
}

// NewHooksGetByIDParamsWithHTTPClient creates a new HooksGetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHooksGetByIDParamsWithHTTPClient(client *http.Client) *HooksGetByIDParams {
	var ()
	return &HooksGetByIDParams{
		HTTPClient: client,
	}
}

/*HooksGetByIDParams contains all the parameters to send to the API endpoint
for the hooks get by Id operation typically these are written to a http.Request
*/
type HooksGetByIDParams struct {

	/*Identifier
	  id of the hook you wish to delete

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the hooks get by Id params
func (o *HooksGetByIDParams) WithTimeout(timeout time.Duration) *HooksGetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the hooks get by Id params
func (o *HooksGetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the hooks get by Id params
func (o *HooksGetByIDParams) WithContext(ctx context.Context) *HooksGetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the hooks get by Id params
func (o *HooksGetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the hooks get by Id params
func (o *HooksGetByIDParams) WithHTTPClient(client *http.Client) *HooksGetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the hooks get by Id params
func (o *HooksGetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the hooks get by Id params
func (o *HooksGetByIDParams) WithIdentifier(identifier string) *HooksGetByIDParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the hooks get by Id params
func (o *HooksGetByIDParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *HooksGetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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
