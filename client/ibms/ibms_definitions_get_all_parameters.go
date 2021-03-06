// Code generated by go-swagger; DO NOT EDIT.

package ibms

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

// NewIbmsDefinitionsGetAllParams creates a new IbmsDefinitionsGetAllParams object
// with the default values initialized.
func NewIbmsDefinitionsGetAllParams() *IbmsDefinitionsGetAllParams {

	return &IbmsDefinitionsGetAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIbmsDefinitionsGetAllParamsWithTimeout creates a new IbmsDefinitionsGetAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIbmsDefinitionsGetAllParamsWithTimeout(timeout time.Duration) *IbmsDefinitionsGetAllParams {

	return &IbmsDefinitionsGetAllParams{

		timeout: timeout,
	}
}

// NewIbmsDefinitionsGetAllParamsWithContext creates a new IbmsDefinitionsGetAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewIbmsDefinitionsGetAllParamsWithContext(ctx context.Context) *IbmsDefinitionsGetAllParams {

	return &IbmsDefinitionsGetAllParams{

		Context: ctx,
	}
}

// NewIbmsDefinitionsGetAllParamsWithHTTPClient creates a new IbmsDefinitionsGetAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIbmsDefinitionsGetAllParamsWithHTTPClient(client *http.Client) *IbmsDefinitionsGetAllParams {

	return &IbmsDefinitionsGetAllParams{
		HTTPClient: client,
	}
}

/*IbmsDefinitionsGetAllParams contains all the parameters to send to the API endpoint
for the ibms definitions get all operation typically these are written to a http.Request
*/
type IbmsDefinitionsGetAllParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ibms definitions get all params
func (o *IbmsDefinitionsGetAllParams) WithTimeout(timeout time.Duration) *IbmsDefinitionsGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ibms definitions get all params
func (o *IbmsDefinitionsGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ibms definitions get all params
func (o *IbmsDefinitionsGetAllParams) WithContext(ctx context.Context) *IbmsDefinitionsGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ibms definitions get all params
func (o *IbmsDefinitionsGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ibms definitions get all params
func (o *IbmsDefinitionsGetAllParams) WithHTTPClient(client *http.Client) *IbmsDefinitionsGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ibms definitions get all params
func (o *IbmsDefinitionsGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IbmsDefinitionsGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
