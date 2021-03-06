// Code generated by go-swagger; DO NOT EDIT.

package obms

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

// NewObmsPostLedParams creates a new ObmsPostLedParams object
// with the default values initialized.
func NewObmsPostLedParams() *ObmsPostLedParams {
	var ()
	return &ObmsPostLedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObmsPostLedParamsWithTimeout creates a new ObmsPostLedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObmsPostLedParamsWithTimeout(timeout time.Duration) *ObmsPostLedParams {
	var ()
	return &ObmsPostLedParams{

		timeout: timeout,
	}
}

// NewObmsPostLedParamsWithContext creates a new ObmsPostLedParams object
// with the default values initialized, and the ability to set a context for a request
func NewObmsPostLedParamsWithContext(ctx context.Context) *ObmsPostLedParams {
	var ()
	return &ObmsPostLedParams{

		Context: ctx,
	}
}

// NewObmsPostLedParamsWithHTTPClient creates a new ObmsPostLedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObmsPostLedParamsWithHTTPClient(client *http.Client) *ObmsPostLedParams {
	var ()
	return &ObmsPostLedParams{
		HTTPClient: client,
	}
}

/*ObmsPostLedParams contains all the parameters to send to the API endpoint
for the obms post led operation typically these are written to a http.Request
*/
type ObmsPostLedParams struct {

	/*Body
	  If the body contains the property value: true, the LED will be lit. If the value property does not exist, the LED will be turned off. The body must contain the property nodeId set to the correct node identifier.


	*/
	Body *models.ObmLed

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the obms post led params
func (o *ObmsPostLedParams) WithTimeout(timeout time.Duration) *ObmsPostLedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the obms post led params
func (o *ObmsPostLedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the obms post led params
func (o *ObmsPostLedParams) WithContext(ctx context.Context) *ObmsPostLedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the obms post led params
func (o *ObmsPostLedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the obms post led params
func (o *ObmsPostLedParams) WithHTTPClient(client *http.Client) *ObmsPostLedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the obms post led params
func (o *ObmsPostLedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the obms post led params
func (o *ObmsPostLedParams) WithBody(body *models.ObmLed) *ObmsPostLedParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the obms post led params
func (o *ObmsPostLedParams) SetBody(body *models.ObmLed) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ObmsPostLedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.ObmLed)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
