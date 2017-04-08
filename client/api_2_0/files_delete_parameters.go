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

// NewFilesDeleteParams creates a new FilesDeleteParams object
// with the default values initialized.
func NewFilesDeleteParams() *FilesDeleteParams {
	var ()
	return &FilesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFilesDeleteParamsWithTimeout creates a new FilesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFilesDeleteParamsWithTimeout(timeout time.Duration) *FilesDeleteParams {
	var ()
	return &FilesDeleteParams{

		timeout: timeout,
	}
}

/*FilesDeleteParams contains all the parameters to send to the API endpoint
for the files delete operation typically these are written to a http.Request
*/
type FilesDeleteParams struct {

	/*Fileidentifier
	  UUID of the file you wish to delete

	*/
	Fileidentifier string

	timeout time.Duration
}

// WithFileidentifier adds the fileidentifier to the files delete params
func (o *FilesDeleteParams) WithFileidentifier(fileidentifier string) *FilesDeleteParams {
	o.Fileidentifier = fileidentifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *FilesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param fileidentifier
	if err := r.SetPathParam("fileidentifier", o.Fileidentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}