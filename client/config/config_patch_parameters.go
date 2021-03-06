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

	"github.com/spiegela/gorackhd/models"
)

// NewConfigPatchParams creates a new ConfigPatchParams object
// with the default values initialized.
func NewConfigPatchParams() *ConfigPatchParams {
	var ()
	return &ConfigPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfigPatchParamsWithTimeout creates a new ConfigPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfigPatchParamsWithTimeout(timeout time.Duration) *ConfigPatchParams {
	var ()
	return &ConfigPatchParams{

		timeout: timeout,
	}
}

// NewConfigPatchParamsWithContext creates a new ConfigPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfigPatchParamsWithContext(ctx context.Context) *ConfigPatchParams {
	var ()
	return &ConfigPatchParams{

		Context: ctx,
	}
}

// NewConfigPatchParamsWithHTTPClient creates a new ConfigPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConfigPatchParamsWithHTTPClient(client *http.Client) *ConfigPatchParams {
	var ()
	return &ConfigPatchParams{
		HTTPClient: client,
	}
}

/*ConfigPatchParams contains all the parameters to send to the API endpoint
for the config patch operation typically these are written to a http.Request
*/
type ConfigPatchParams struct {

	/*Config
	  The configuration properties to be modified

	*/
	Config models.GenericObj

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the config patch params
func (o *ConfigPatchParams) WithTimeout(timeout time.Duration) *ConfigPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the config patch params
func (o *ConfigPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the config patch params
func (o *ConfigPatchParams) WithContext(ctx context.Context) *ConfigPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the config patch params
func (o *ConfigPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the config patch params
func (o *ConfigPatchParams) WithHTTPClient(client *http.Client) *ConfigPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the config patch params
func (o *ConfigPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfig adds the config to the config patch params
func (o *ConfigPatchParams) WithConfig(config models.GenericObj) *ConfigPatchParams {
	o.SetConfig(config)
	return o
}

// SetConfig adds the config to the config patch params
func (o *ConfigPatchParams) SetConfig(config models.GenericObj) {
	o.Config = config
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Config); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
