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

// HooksGetByIDReader is a Reader for the HooksGetByID structure.
type HooksGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *HooksGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHooksGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewHooksGetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewHooksGetByIDOK creates a HooksGetByIDOK with default headers values
func NewHooksGetByIDOK() *HooksGetByIDOK {
	return &HooksGetByIDOK{}
}

/*HooksGetByIDOK handles this case with default header values.

Successfully retrieved the specified hook
*/
type HooksGetByIDOK struct {
	Payload HooksGetByIDOKBodyBody
}

func (o *HooksGetByIDOK) Error() string {
	return fmt.Sprintf("[GET /hooks/{identifier}][%d] hooksGetByIdOK  %+v", 200, o.Payload)
}

func (o *HooksGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHooksGetByIDDefault creates a HooksGetByIDDefault with default headers values
func NewHooksGetByIDDefault(code int) *HooksGetByIDDefault {
	return &HooksGetByIDDefault{
		_statusCode: code,
	}
}

/*HooksGetByIDDefault handles this case with default header values.

Unexpected error
*/
type HooksGetByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the hooks get by Id default response
func (o *HooksGetByIDDefault) Code() int {
	return o._statusCode
}

func (o *HooksGetByIDDefault) Error() string {
	return fmt.Sprintf("[GET /hooks/{identifier}][%d] hooksGetById default  %+v", o._statusCode, o.Payload)
}

func (o *HooksGetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*HooksGetByIDOKBodyBody hooks get by ID o k body body

swagger:model HooksGetByIDOKBodyBody
*/
type HooksGetByIDOKBodyBody interface{}