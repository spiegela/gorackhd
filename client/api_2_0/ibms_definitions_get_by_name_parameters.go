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

// NewIbmsDefinitionsGetByNameParams creates a new IbmsDefinitionsGetByNameParams object
// with the default values initialized.
func NewIbmsDefinitionsGetByNameParams() *IbmsDefinitionsGetByNameParams {
	var ()
	return &IbmsDefinitionsGetByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIbmsDefinitionsGetByNameParamsWithTimeout creates a new IbmsDefinitionsGetByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIbmsDefinitionsGetByNameParamsWithTimeout(timeout time.Duration) *IbmsDefinitionsGetByNameParams {
	var ()
	return &IbmsDefinitionsGetByNameParams{

		timeout: timeout,
	}
}

/*IbmsDefinitionsGetByNameParams contains all the parameters to send to the API endpoint
for the ibms definitions get by name operation typically these are written to a http.Request
*/
type IbmsDefinitionsGetByNameParams struct {

	/*Name
	  The IBMS service name

	*/
	Name string

	timeout time.Duration
}

// WithName adds the name to the ibms definitions get by name params
func (o *IbmsDefinitionsGetByNameParams) WithName(name string) *IbmsDefinitionsGetByNameParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *IbmsDefinitionsGetByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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