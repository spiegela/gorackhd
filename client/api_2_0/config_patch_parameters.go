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

/*ConfigPatchParams contains all the parameters to send to the API endpoint
for the config patch operation typically these are written to a http.Request
*/
type ConfigPatchParams struct {

	/*Config
	  The configuration properties to be modified

	*/
	Config models.GenericObj

	timeout time.Duration
}

// WithConfig adds the config to the config patch params
func (o *ConfigPatchParams) WithConfig(config models.GenericObj) *ConfigPatchParams {
	o.Config = config
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Config); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}