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

// WorkflowsDeleteByInstanceIDReader is a Reader for the WorkflowsDeleteByInstanceID structure.
type WorkflowsDeleteByInstanceIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *WorkflowsDeleteByInstanceIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWorkflowsDeleteByInstanceIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewWorkflowsDeleteByInstanceIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewWorkflowsDeleteByInstanceIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewWorkflowsDeleteByInstanceIDNoContent creates a WorkflowsDeleteByInstanceIDNoContent with default headers values
func NewWorkflowsDeleteByInstanceIDNoContent() *WorkflowsDeleteByInstanceIDNoContent {
	return &WorkflowsDeleteByInstanceIDNoContent{}
}

/*WorkflowsDeleteByInstanceIDNoContent handles this case with default header values.

Successfully deleted the specified workflow
*/
type WorkflowsDeleteByInstanceIDNoContent struct {
	Payload WorkflowsDeleteByInstanceIDNoContentBodyBody
}

func (o *WorkflowsDeleteByInstanceIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{identifier}][%d] workflowsDeleteByInstanceIdNoContent  %+v", 204, o.Payload)
}

func (o *WorkflowsDeleteByInstanceIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsDeleteByInstanceIDNotFound creates a WorkflowsDeleteByInstanceIDNotFound with default headers values
func NewWorkflowsDeleteByInstanceIDNotFound() *WorkflowsDeleteByInstanceIDNotFound {
	return &WorkflowsDeleteByInstanceIDNotFound{}
}

/*WorkflowsDeleteByInstanceIDNotFound handles this case with default header values.

The specified workflow was not found
*/
type WorkflowsDeleteByInstanceIDNotFound struct {
	Payload *models.Error
}

func (o *WorkflowsDeleteByInstanceIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{identifier}][%d] workflowsDeleteByInstanceIdNotFound  %+v", 404, o.Payload)
}

func (o *WorkflowsDeleteByInstanceIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsDeleteByInstanceIDDefault creates a WorkflowsDeleteByInstanceIDDefault with default headers values
func NewWorkflowsDeleteByInstanceIDDefault(code int) *WorkflowsDeleteByInstanceIDDefault {
	return &WorkflowsDeleteByInstanceIDDefault{
		_statusCode: code,
	}
}

/*WorkflowsDeleteByInstanceIDDefault handles this case with default header values.

Unexpected error
*/
type WorkflowsDeleteByInstanceIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the workflows delete by instance Id default response
func (o *WorkflowsDeleteByInstanceIDDefault) Code() int {
	return o._statusCode
}

func (o *WorkflowsDeleteByInstanceIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /workflows/{identifier}][%d] workflowsDeleteByInstanceId default  %+v", o._statusCode, o.Payload)
}

func (o *WorkflowsDeleteByInstanceIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WorkflowsDeleteByInstanceIDNoContentBodyBody workflows delete by instance ID no content body body

swagger:model WorkflowsDeleteByInstanceIDNoContentBodyBody
*/
type WorkflowsDeleteByInstanceIDNoContentBodyBody interface{}