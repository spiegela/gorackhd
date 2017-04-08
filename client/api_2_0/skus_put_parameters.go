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

// NewSkusPutParams creates a new SkusPutParams object
// with the default values initialized.
func NewSkusPutParams() *SkusPutParams {
	var ()
	return &SkusPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSkusPutParamsWithTimeout creates a new SkusPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSkusPutParamsWithTimeout(timeout time.Duration) *SkusPutParams {
	var ()
	return &SkusPutParams{

		timeout: timeout,
	}
}

/*SkusPutParams contains all the parameters to send to the API endpoint
for the skus put operation typically these are written to a http.Request
*/
type SkusPutParams struct {

	/*Body
	  The properties used to define the SKU

	*/
	Body *models.Skus20SkusUpsert

	timeout time.Duration
}

// WithBody adds the body to the skus put params
func (o *SkusPutParams) WithBody(body *models.Skus20SkusUpsert) *SkusPutParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SkusPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.Skus20SkusUpsert)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}