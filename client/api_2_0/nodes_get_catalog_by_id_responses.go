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

// NodesGetCatalogByIDReader is a Reader for the NodesGetCatalogByID structure.
type NodesGetCatalogByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *NodesGetCatalogByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNodesGetCatalogByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNodesGetCatalogByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNodesGetCatalogByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewNodesGetCatalogByIDOK creates a NodesGetCatalogByIDOK with default headers values
func NewNodesGetCatalogByIDOK() *NodesGetCatalogByIDOK {
	return &NodesGetCatalogByIDOK{}
}

/*NodesGetCatalogByIDOK handles this case with default header values.

Successfully retrieved catalogs of specified node
*/
type NodesGetCatalogByIDOK struct {
	Payload NodesGetCatalogByIDOKBodyBody
}

func (o *NodesGetCatalogByIDOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs][%d] nodesGetCatalogByIdOK  %+v", 200, o.Payload)
}

func (o *NodesGetCatalogByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesGetCatalogByIDNotFound creates a NodesGetCatalogByIDNotFound with default headers values
func NewNodesGetCatalogByIDNotFound() *NodesGetCatalogByIDNotFound {
	return &NodesGetCatalogByIDNotFound{}
}

/*NodesGetCatalogByIDNotFound handles this case with default header values.

The specified node was not found
*/
type NodesGetCatalogByIDNotFound struct {
	Payload *models.Error
}

func (o *NodesGetCatalogByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs][%d] nodesGetCatalogByIdNotFound  %+v", 404, o.Payload)
}

func (o *NodesGetCatalogByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesGetCatalogByIDDefault creates a NodesGetCatalogByIDDefault with default headers values
func NewNodesGetCatalogByIDDefault(code int) *NodesGetCatalogByIDDefault {
	return &NodesGetCatalogByIDDefault{
		_statusCode: code,
	}
}

/*NodesGetCatalogByIDDefault handles this case with default header values.

Unexpected error
*/
type NodesGetCatalogByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the nodes get catalog by Id default response
func (o *NodesGetCatalogByIDDefault) Code() int {
	return o._statusCode
}

func (o *NodesGetCatalogByIDDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/catalogs][%d] nodesGetCatalogById default  %+v", o._statusCode, o.Payload)
}

func (o *NodesGetCatalogByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*NodesGetCatalogByIDOKBodyBody nodes get catalog by ID o k body body

swagger:model NodesGetCatalogByIDOKBodyBody
*/
type NodesGetCatalogByIDOKBodyBody interface{}