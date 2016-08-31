package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/emccode/gorackhd/models"
)

// GetNodesIdentifierWorkflowsReader is a Reader for the GetNodesIdentifierWorkflows structure.
type GetNodesIdentifierWorkflowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNodesIdentifierWorkflowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIdentifierWorkflowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNodesIdentifierWorkflowsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodesIdentifierWorkflowsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNodesIdentifierWorkflowsOK creates a GetNodesIdentifierWorkflowsOK with default headers values
func NewGetNodesIdentifierWorkflowsOK() *GetNodesIdentifierWorkflowsOK {
	return &GetNodesIdentifierWorkflowsOK{}
}

/*GetNodesIdentifierWorkflowsOK handles this case with default header values.

all workflows for specified node, empty object if none exist.

*/
type GetNodesIdentifierWorkflowsOK struct {
	Payload []*models.Graphobject
}

func (o *GetNodesIdentifierWorkflowsOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/workflows][%d] getNodesIdentifierWorkflowsOK  %+v", 200, o.Payload)
}

func (o *GetNodesIdentifierWorkflowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierWorkflowsNotFound creates a GetNodesIdentifierWorkflowsNotFound with default headers values
func NewGetNodesIdentifierWorkflowsNotFound() *GetNodesIdentifierWorkflowsNotFound {
	return &GetNodesIdentifierWorkflowsNotFound{}
}

/*GetNodesIdentifierWorkflowsNotFound handles this case with default header values.

The node with the identifier was not found

*/
type GetNodesIdentifierWorkflowsNotFound struct {
	Payload *models.Error
}

func (o *GetNodesIdentifierWorkflowsNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/workflows][%d] getNodesIdentifierWorkflowsNotFound  %+v", 404, o.Payload)
}

func (o *GetNodesIdentifierWorkflowsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodesIdentifierWorkflowsDefault creates a GetNodesIdentifierWorkflowsDefault with default headers values
func NewGetNodesIdentifierWorkflowsDefault(code int) *GetNodesIdentifierWorkflowsDefault {
	return &GetNodesIdentifierWorkflowsDefault{
		_statusCode: code,
	}
}

/*GetNodesIdentifierWorkflowsDefault handles this case with default header values.

Unexpected error
*/
type GetNodesIdentifierWorkflowsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes identifier workflows default response
func (o *GetNodesIdentifierWorkflowsDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesIdentifierWorkflowsDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/workflows][%d] GetNodesIdentifierWorkflows default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesIdentifierWorkflowsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
