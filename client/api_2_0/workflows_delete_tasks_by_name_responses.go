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

// WorkflowsDeleteTasksByNameReader is a Reader for the WorkflowsDeleteTasksByName structure.
type WorkflowsDeleteTasksByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *WorkflowsDeleteTasksByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWorkflowsDeleteTasksByNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewWorkflowsDeleteTasksByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewWorkflowsDeleteTasksByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewWorkflowsDeleteTasksByNameNoContent creates a WorkflowsDeleteTasksByNameNoContent with default headers values
func NewWorkflowsDeleteTasksByNameNoContent() *WorkflowsDeleteTasksByNameNoContent {
	return &WorkflowsDeleteTasksByNameNoContent{}
}

/*WorkflowsDeleteTasksByNameNoContent handles this case with default header values.

Successfully deleted the specified workflow task
*/
type WorkflowsDeleteTasksByNameNoContent struct {
	Payload WorkflowsDeleteTasksByNameNoContentBodyBody
}

func (o *WorkflowsDeleteTasksByNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workflows/tasks/{injectableName}][%d] workflowsDeleteTasksByNameNoContent  %+v", 204, o.Payload)
}

func (o *WorkflowsDeleteTasksByNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsDeleteTasksByNameNotFound creates a WorkflowsDeleteTasksByNameNotFound with default headers values
func NewWorkflowsDeleteTasksByNameNotFound() *WorkflowsDeleteTasksByNameNotFound {
	return &WorkflowsDeleteTasksByNameNotFound{}
}

/*WorkflowsDeleteTasksByNameNotFound handles this case with default header values.

The workflow task with the specified injectable name was not found
*/
type WorkflowsDeleteTasksByNameNotFound struct {
	Payload *models.Error
}

func (o *WorkflowsDeleteTasksByNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflows/tasks/{injectableName}][%d] workflowsDeleteTasksByNameNotFound  %+v", 404, o.Payload)
}

func (o *WorkflowsDeleteTasksByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsDeleteTasksByNameDefault creates a WorkflowsDeleteTasksByNameDefault with default headers values
func NewWorkflowsDeleteTasksByNameDefault(code int) *WorkflowsDeleteTasksByNameDefault {
	return &WorkflowsDeleteTasksByNameDefault{
		_statusCode: code,
	}
}

/*WorkflowsDeleteTasksByNameDefault handles this case with default header values.

Unexpected error
*/
type WorkflowsDeleteTasksByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the workflows delete tasks by name default response
func (o *WorkflowsDeleteTasksByNameDefault) Code() int {
	return o._statusCode
}

func (o *WorkflowsDeleteTasksByNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /workflows/tasks/{injectableName}][%d] workflowsDeleteTasksByName default  %+v", o._statusCode, o.Payload)
}

func (o *WorkflowsDeleteTasksByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WorkflowsDeleteTasksByNameNoContentBodyBody workflows delete tasks by name no content body body

swagger:model WorkflowsDeleteTasksByNameNoContentBodyBody
*/
type WorkflowsDeleteTasksByNameNoContentBodyBody interface{}