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

/*HooksGetByIDParams contains all the parameters to send to the API endpoint
for the hooks get by Id operation typically these are written to a http.Request
*/
type HooksGetByIDParams struct {

	/*Identifier
	  id of the hook you wish to delete

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the hooks get by Id params
func (o *HooksGetByIDParams) WithIdentifier(identifier string) *HooksGetByIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *HooksGetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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