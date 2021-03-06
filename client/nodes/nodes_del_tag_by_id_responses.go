// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NodesDelTagByIDReader is a Reader for the NodesDelTagByID structure.
type NodesDelTagByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesDelTagByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewNodesDelTagByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNodesDelTagByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNodesDelTagByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodesDelTagByIDNoContent creates a NodesDelTagByIDNoContent with default headers values
func NewNodesDelTagByIDNoContent() *NodesDelTagByIDNoContent {
	return &NodesDelTagByIDNoContent{}
}

/*NodesDelTagByIDNoContent handles this case with default header values.

Successfully deleted a specified tag from a node.
*/
type NodesDelTagByIDNoContent struct {
}

func (o *NodesDelTagByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}/tags/{tagName}][%d] nodesDelTagByIdNoContent ", 204)
}

func (o *NodesDelTagByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNodesDelTagByIDNotFound creates a NodesDelTagByIDNotFound with default headers values
func NewNodesDelTagByIDNotFound() *NodesDelTagByIDNotFound {
	return &NodesDelTagByIDNotFound{}
}

/*NodesDelTagByIDNotFound handles this case with default header values.

The specified node was not found.
*/
type NodesDelTagByIDNotFound struct {
	Payload *models.Error
}

func (o *NodesDelTagByIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}/tags/{tagName}][%d] nodesDelTagByIdNotFound  %+v", 404, o.Payload)
}

func (o *NodesDelTagByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesDelTagByIDDefault creates a NodesDelTagByIDDefault with default headers values
func NewNodesDelTagByIDDefault(code int) *NodesDelTagByIDDefault {
	return &NodesDelTagByIDDefault{
		_statusCode: code,
	}
}

/*NodesDelTagByIDDefault handles this case with default header values.

Unexpected error
*/
type NodesDelTagByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the nodes del tag by Id default response
func (o *NodesDelTagByIDDefault) Code() int {
	return o._statusCode
}

func (o *NodesDelTagByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{identifier}/tags/{tagName}][%d] nodesDelTagById default  %+v", o._statusCode, o.Payload)
}

func (o *NodesDelTagByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
