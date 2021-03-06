// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// FilesMd5GetReader is a Reader for the FilesMd5Get structure.
type FilesMd5GetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FilesMd5GetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFilesMd5GetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewFilesMd5GetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewFilesMd5GetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewFilesMd5GetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFilesMd5GetOK creates a FilesMd5GetOK with default headers values
func NewFilesMd5GetOK() *FilesMd5GetOK {
	return &FilesMd5GetOK{}
}

/*FilesMd5GetOK handles this case with default header values.

Successfully retrieved the md5sum of the specified file
*/
type FilesMd5GetOK struct {
	Payload FilesMd5GetOKBody
}

func (o *FilesMd5GetOK) Error() string {
	return fmt.Sprintf("[GET /files/{filename}/md5][%d] filesMd5GetOK  %+v", 200, o.Payload)
}

func (o *FilesMd5GetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesMd5GetNotFound creates a FilesMd5GetNotFound with default headers values
func NewFilesMd5GetNotFound() *FilesMd5GetNotFound {
	return &FilesMd5GetNotFound{}
}

/*FilesMd5GetNotFound handles this case with default header values.

File not found
*/
type FilesMd5GetNotFound struct {
	Payload *models.Error
}

func (o *FilesMd5GetNotFound) Error() string {
	return fmt.Sprintf("[GET /files/{filename}/md5][%d] filesMd5GetNotFound  %+v", 404, o.Payload)
}

func (o *FilesMd5GetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesMd5GetInternalServerError creates a FilesMd5GetInternalServerError with default headers values
func NewFilesMd5GetInternalServerError() *FilesMd5GetInternalServerError {
	return &FilesMd5GetInternalServerError{}
}

/*FilesMd5GetInternalServerError handles this case with default header values.

Failed to serve file request
*/
type FilesMd5GetInternalServerError struct {
	Payload *models.Error
}

func (o *FilesMd5GetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /files/{filename}/md5][%d] filesMd5GetInternalServerError  %+v", 500, o.Payload)
}

func (o *FilesMd5GetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFilesMd5GetDefault creates a FilesMd5GetDefault with default headers values
func NewFilesMd5GetDefault(code int) *FilesMd5GetDefault {
	return &FilesMd5GetDefault{
		_statusCode: code,
	}
}

/*FilesMd5GetDefault handles this case with default header values.

Unexpected error
*/
type FilesMd5GetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the files md5 get default response
func (o *FilesMd5GetDefault) Code() int {
	return o._statusCode
}

func (o *FilesMd5GetDefault) Error() string {
	return fmt.Sprintf("[GET /files/{filename}/md5][%d] filesMd5Get default  %+v", o._statusCode, o.Payload)
}

func (o *FilesMd5GetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*FilesMd5GetOKBody files md5 get o k body
swagger:model FilesMd5GetOKBody
*/

type FilesMd5GetOKBody interface{}
