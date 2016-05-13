package skus

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetSkusReader is a Reader for the GetSkus structure.
type GetSkusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSkusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSkusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetSkusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetSkusOK creates a GetSkusOK with default headers values
func NewGetSkusOK() *GetSkusOK {
	return &GetSkusOK{}
}

/*GetSkusOK handles this case with default header values.

list of skus

*/
type GetSkusOK struct {
	Payload GetSkusOKBodyBody
}

func (o *GetSkusOK) Error() string {
	return fmt.Sprintf("[GET /skus][%d] getSkusOK  %+v", 200, o.Payload)
}

func (o *GetSkusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSkusDefault creates a GetSkusDefault with default headers values
func NewGetSkusDefault(code int) *GetSkusDefault {
	return &GetSkusDefault{
		_statusCode: code,
	}
}

/*GetSkusDefault handles this case with default header values.

Unexpected error
*/
type GetSkusDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get skus default response
func (o *GetSkusDefault) Code() int {
	return o._statusCode
}

func (o *GetSkusDefault) Error() string {
	return fmt.Sprintf("[GET /skus][%d] GetSkus default  %+v", o._statusCode, o.Payload)
}

func (o *GetSkusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSkusOKBodyBody get skus o k body body

swagger:model GetSkusOKBodyBody
*/
type GetSkusOKBodyBody interface{}