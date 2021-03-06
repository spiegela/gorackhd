// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewGetBootstrapParams creates a new GetBootstrapParams object
// with the default values initialized.
func NewGetBootstrapParams() *GetBootstrapParams {
	var ()
	return &GetBootstrapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBootstrapParamsWithTimeout creates a new GetBootstrapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBootstrapParamsWithTimeout(timeout time.Duration) *GetBootstrapParams {
	var ()
	return &GetBootstrapParams{

		timeout: timeout,
	}
}

// NewGetBootstrapParamsWithContext creates a new GetBootstrapParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBootstrapParamsWithContext(ctx context.Context) *GetBootstrapParams {
	var ()
	return &GetBootstrapParams{

		Context: ctx,
	}
}

// NewGetBootstrapParamsWithHTTPClient creates a new GetBootstrapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetBootstrapParamsWithHTTPClient(client *http.Client) *GetBootstrapParams {
	var ()
	return &GetBootstrapParams{
		HTTPClient: client,
	}
}

/*GetBootstrapParams contains all the parameters to send to the API endpoint
for the get bootstrap operation typically these are written to a http.Request
*/
type GetBootstrapParams struct {

	/*MacAddress
	  Query string containing the mac address

	*/
	MacAddress *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get bootstrap params
func (o *GetBootstrapParams) WithTimeout(timeout time.Duration) *GetBootstrapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get bootstrap params
func (o *GetBootstrapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get bootstrap params
func (o *GetBootstrapParams) WithContext(ctx context.Context) *GetBootstrapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get bootstrap params
func (o *GetBootstrapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get bootstrap params
func (o *GetBootstrapParams) WithHTTPClient(client *http.Client) *GetBootstrapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get bootstrap params
func (o *GetBootstrapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMacAddress adds the macAddress to the get bootstrap params
func (o *GetBootstrapParams) WithMacAddress(macAddress *string) *GetBootstrapParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the get bootstrap params
func (o *GetBootstrapParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WriteToRequest writes these params to a swagger request
func (o *GetBootstrapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MacAddress != nil {

		// query param macAddress
		var qrMacAddress string
		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {
			if err := r.SetQueryParam("macAddress", qMacAddress); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
