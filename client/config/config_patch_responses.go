// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// ConfigPatchReader is a Reader for the ConfigPatch structure.
type ConfigPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewConfigPatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewConfigPatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigPatchOK creates a ConfigPatchOK with default headers values
func NewConfigPatchOK() *ConfigPatchOK {
	return &ConfigPatchOK{}
}

/*ConfigPatchOK handles this case with default header values.

Successfully modified the configuration
*/
type ConfigPatchOK struct {
	Payload ConfigPatchOKBody
}

func (o *ConfigPatchOK) Error() string {
	return fmt.Sprintf("[PATCH /config][%d] configPatchOK  %+v", 200, o.Payload)
}

func (o *ConfigPatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigPatchDefault creates a ConfigPatchDefault with default headers values
func NewConfigPatchDefault(code int) *ConfigPatchDefault {
	return &ConfigPatchDefault{
		_statusCode: code,
	}
}

/*ConfigPatchDefault handles this case with default header values.

Unexpected error
*/
type ConfigPatchDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the config patch default response
func (o *ConfigPatchDefault) Code() int {
	return o._statusCode
}

func (o *ConfigPatchDefault) Error() string {
	return fmt.Sprintf("[PATCH /config][%d] configPatch default  %+v", o._statusCode, o.Payload)
}

func (o *ConfigPatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ConfigPatchOKBody config patch o k body
swagger:model ConfigPatchOKBody
*/

type ConfigPatchOKBody interface{}
