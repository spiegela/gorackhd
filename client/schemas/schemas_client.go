// Code generated by go-swagger; DO NOT EDIT.

package schemas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new schemas API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schemas API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SchemasGet gets all schemas

Get a list of all schemas currently stored in the system.
*/
func (a *Client) SchemasGet(params *SchemasGetParams, authInfo runtime.ClientAuthInfoWriter) (*SchemasGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemasGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schemasGet",
		Method:             "GET",
		PathPattern:        "/schemas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SchemasGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemasGetOK), nil

}

/*
SchemasIDGet gets a schema

Get the specified schema.
*/
func (a *Client) SchemasIDGet(params *SchemasIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*SchemasIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSchemasIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "schemasIdGet",
		Method:             "GET",
		PathPattern:        "/schemas/{identifier}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SchemasIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SchemasIDGetOK), nil

}

/*
TaskSchemasGet gets all task schemas names

Get a list of all task schema names currently stored in the system.
*/
func (a *Client) TaskSchemasGet(params *TaskSchemasGetParams, authInfo runtime.ClientAuthInfoWriter) (*TaskSchemasGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTaskSchemasGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "taskSchemasGet",
		Method:             "GET",
		PathPattern:        "/schemas/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TaskSchemasGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TaskSchemasGetOK), nil

}

/*
TaskSchemasIDGet gets a task schema

Get the specified task schema.
*/
func (a *Client) TaskSchemasIDGet(params *TaskSchemasIDGetParams, authInfo runtime.ClientAuthInfoWriter) (*TaskSchemasIDGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTaskSchemasIDGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "taskSchemasIdGet",
		Method:             "GET",
		PathPattern:        "/schemas/tasks/{identifier}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TaskSchemasIDGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TaskSchemasIDGetOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
