package api_2_0

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// IbmsGetReader is a Reader for the IbmsGet structure.
type IbmsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *IbmsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIbmsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewIbmsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewIbmsGetOK creates a IbmsGetOK with default headers values
func NewIbmsGetOK() *IbmsGetOK {
	return &IbmsGetOK{}
}

/*IbmsGetOK handles this case with default header values.

Successfully retrieved the list of IBMS service instances
*/
type IbmsGetOK struct {
	Payload IbmsGetOKBodyBody
}

func (o *IbmsGetOK) Error() string {
	return fmt.Sprintf("[GET /ibms][%d] ibmsGetOK  %+v", 200, o.Payload)
}

func (o *IbmsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIbmsGetDefault creates a IbmsGetDefault with default headers values
func NewIbmsGetDefault(code int) *IbmsGetDefault {
	return &IbmsGetDefault{
		_statusCode: code,
	}
}

/*IbmsGetDefault handles this case with default header values.

Unexpected error
*/
type IbmsGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the ibms get default response
func (o *IbmsGetDefault) Code() int {
	return o._statusCode
}

func (o *IbmsGetDefault) Error() string {
	return fmt.Sprintf("[GET /ibms][%d] ibmsGet default  %+v", o._statusCode, o.Payload)
}

func (o *IbmsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IbmsGetOKBodyBody ibms get o k body body

swagger:model IbmsGetOKBodyBody
*/
type IbmsGetOKBodyBody interface{}