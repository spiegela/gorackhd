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

// NewGetAllTagsParams creates a new GetAllTagsParams object
// with the default values initialized.
func NewGetAllTagsParams() *GetAllTagsParams {
	var ()
	return &GetAllTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllTagsParamsWithTimeout creates a new GetAllTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllTagsParamsWithTimeout(timeout time.Duration) *GetAllTagsParams {
	var ()
	return &GetAllTagsParams{

		timeout: timeout,
	}
}

/*GetAllTagsParams contains all the parameters to send to the API endpoint
for the get all tags operation typically these are written to a http.Request
*/
type GetAllTagsParams struct {

	/*Sort
	  Query string specifying properties to sort with

	*/
	Sort *string

	timeout time.Duration
}

// WithSort adds the sort to the get all tags params
func (o *GetAllTagsParams) WithSort(sort *string) *GetAllTagsParams {
	o.Sort = sort
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}