// Code generated by go-swagger; DO NOT EDIT.

package views

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

// NewViewsGetByIDParams creates a new ViewsGetByIDParams object
// with the default values initialized.
func NewViewsGetByIDParams() *ViewsGetByIDParams {
	var ()
	return &ViewsGetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewViewsGetByIDParamsWithTimeout creates a new ViewsGetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewViewsGetByIDParamsWithTimeout(timeout time.Duration) *ViewsGetByIDParams {
	var ()
	return &ViewsGetByIDParams{

		timeout: timeout,
	}
}

// NewViewsGetByIDParamsWithContext creates a new ViewsGetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewViewsGetByIDParamsWithContext(ctx context.Context) *ViewsGetByIDParams {
	var ()
	return &ViewsGetByIDParams{

		Context: ctx,
	}
}

// NewViewsGetByIDParamsWithHTTPClient creates a new ViewsGetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewViewsGetByIDParamsWithHTTPClient(client *http.Client) *ViewsGetByIDParams {
	var ()
	return &ViewsGetByIDParams{
		HTTPClient: client,
	}
}

/*ViewsGetByIDParams contains all the parameters to send to the API endpoint
for the views get by Id operation typically these are written to a http.Request
*/
type ViewsGetByIDParams struct {

	/*Identifier
	  The view name

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the views get by Id params
func (o *ViewsGetByIDParams) WithTimeout(timeout time.Duration) *ViewsGetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the views get by Id params
func (o *ViewsGetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the views get by Id params
func (o *ViewsGetByIDParams) WithContext(ctx context.Context) *ViewsGetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the views get by Id params
func (o *ViewsGetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the views get by Id params
func (o *ViewsGetByIDParams) WithHTTPClient(client *http.Client) *ViewsGetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the views get by Id params
func (o *ViewsGetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the views get by Id params
func (o *ViewsGetByIDParams) WithIdentifier(identifier string) *ViewsGetByIDParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the views get by Id params
func (o *ViewsGetByIDParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *ViewsGetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
