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

// NewIbmsGetByIDParams creates a new IbmsGetByIDParams object
// with the default values initialized.
func NewIbmsGetByIDParams() *IbmsGetByIDParams {
	var ()
	return &IbmsGetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIbmsGetByIDParamsWithTimeout creates a new IbmsGetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIbmsGetByIDParamsWithTimeout(timeout time.Duration) *IbmsGetByIDParams {
	var ()
	return &IbmsGetByIDParams{

		timeout: timeout,
	}
}

/*IbmsGetByIDParams contains all the parameters to send to the API endpoint
for the ibms get by Id operation typically these are written to a http.Request
*/
type IbmsGetByIDParams struct {

	/*Identifier
	  The IBMS service identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the ibms get by Id params
func (o *IbmsGetByIDParams) WithIdentifier(identifier string) *IbmsGetByIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *IbmsGetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
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