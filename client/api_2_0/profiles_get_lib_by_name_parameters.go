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

// NewProfilesGetLibByNameParams creates a new ProfilesGetLibByNameParams object
// with the default values initialized.
func NewProfilesGetLibByNameParams() *ProfilesGetLibByNameParams {
	var ()
	return &ProfilesGetLibByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProfilesGetLibByNameParamsWithTimeout creates a new ProfilesGetLibByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProfilesGetLibByNameParamsWithTimeout(timeout time.Duration) *ProfilesGetLibByNameParams {
	var ()
	return &ProfilesGetLibByNameParams{

		timeout: timeout,
	}
}

/*ProfilesGetLibByNameParams contains all the parameters to send to the API endpoint
for the profiles get lib by name operation typically these are written to a http.Request
*/
type ProfilesGetLibByNameParams struct {

	/*Name
	  The profile name

	*/
	Name string
	/*Scope
	  The profile scope

	*/
	Scope *string

	timeout time.Duration
}

// WithName adds the name to the profiles get lib by name params
func (o *ProfilesGetLibByNameParams) WithName(name string) *ProfilesGetLibByNameParams {
	o.Name = name
	return o
}

// WithScope adds the scope to the profiles get lib by name params
func (o *ProfilesGetLibByNameParams) WithScope(scope *string) *ProfilesGetLibByNameParams {
	o.Scope = scope
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ProfilesGetLibByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Scope != nil {

		// query param scope
		var qrScope string
		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {
			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}