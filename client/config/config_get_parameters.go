// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewConfigGetParams creates a new ConfigGetParams object
// with the default values initialized.
func NewConfigGetParams() *ConfigGetParams {

	return &ConfigGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfigGetParamsWithTimeout creates a new ConfigGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfigGetParamsWithTimeout(timeout time.Duration) *ConfigGetParams {

	return &ConfigGetParams{

		timeout: timeout,
	}
}

// NewConfigGetParamsWithContext creates a new ConfigGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfigGetParamsWithContext(ctx context.Context) *ConfigGetParams {

	return &ConfigGetParams{

		Context: ctx,
	}
}

// NewConfigGetParamsWithHTTPClient creates a new ConfigGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConfigGetParamsWithHTTPClient(client *http.Client) *ConfigGetParams {

	return &ConfigGetParams{
		HTTPClient: client,
	}
}

/*ConfigGetParams contains all the parameters to send to the API endpoint
for the config get operation typically these are written to a http.Request
*/
type ConfigGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the config get params
func (o *ConfigGetParams) WithTimeout(timeout time.Duration) *ConfigGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the config get params
func (o *ConfigGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the config get params
func (o *ConfigGetParams) WithContext(ctx context.Context) *ConfigGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the config get params
func (o *ConfigGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the config get params
func (o *ConfigGetParams) WithHTTPClient(client *http.Client) *ConfigGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the config get params
func (o *ConfigGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
