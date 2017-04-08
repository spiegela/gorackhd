package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// ListUsersReader is a Reader for the ListUsers structure.
type ListUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ListUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewListUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewListUsersOK creates a ListUsersOK with default headers values
func NewListUsersOK() *ListUsersOK {
	return &ListUsersOK{}
}

/*ListUsersOK handles this case with default header values.

Successfully retrieved the list of users
*/
type ListUsersOK struct {
	Payload []*models.GetUserObj
}

func (o *ListUsersOK) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersOK  %+v", 200, o.Payload)
}

func (o *ListUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsersUnauthorized creates a ListUsersUnauthorized with default headers values
func NewListUsersUnauthorized() *ListUsersUnauthorized {
	return &ListUsersUnauthorized{}
}

/*ListUsersUnauthorized handles this case with default header values.

Unauthorized
*/
type ListUsersUnauthorized struct {
	Payload ListUsersUnauthorizedBodyBody
}

func (o *ListUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *ListUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsersForbidden creates a ListUsersForbidden with default headers values
func NewListUsersForbidden() *ListUsersForbidden {
	return &ListUsersForbidden{}
}

/*ListUsersForbidden handles this case with default header values.

Forbidden
*/
type ListUsersForbidden struct {
	Payload ListUsersForbiddenBodyBody
}

func (o *ListUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsersForbidden  %+v", 403, o.Payload)
}

func (o *ListUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsersDefault creates a ListUsersDefault with default headers values
func NewListUsersDefault(code int) *ListUsersDefault {
	return &ListUsersDefault{
		_statusCode: code,
	}
}

/*ListUsersDefault handles this case with default header values.

Unexpected error
*/
type ListUsersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list users default response
func (o *ListUsersDefault) Code() int {
	return o._statusCode
}

func (o *ListUsersDefault) Error() string {
	return fmt.Sprintf("[GET /users][%d] listUsers default  %+v", o._statusCode, o.Payload)
}

func (o *ListUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ListUsersForbiddenBodyBody list users forbidden body body

swagger:model ListUsersForbiddenBodyBody
*/
type ListUsersForbiddenBodyBody interface{}

/*ListUsersUnauthorizedBodyBody list users unauthorized body body

swagger:model ListUsersUnauthorizedBodyBody
*/
type ListUsersUnauthorizedBodyBody interface{}
