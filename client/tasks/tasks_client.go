package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetBootstrap gets tasks bootstrap js

Used internally by the system - get tasks bootstrap.js
*/
func (a *Client) GetBootstrap(params *GetBootstrapParams, authInfo runtime.ClientAuthInfoWriter) (*GetBootstrapOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBootstrapParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getBootstrap",
		Method:             "GET",
		PathPattern:        "/tasks/bootstrap.js",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetBootstrapReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetBootstrapOK), nil
}

/*
GetTasksByID gets the specified task

Get the specified task.
*/
func (a *Client) GetTasksByID(params *GetTasksByIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetTasksByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTasksByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTasksById",
		Method:             "GET",
		PathPattern:        "/tasks/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTasksByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTasksByIDOK), nil
}

/*
PostTaskByID posts a task

Start the specified task
*/
func (a *Client) PostTaskByID(params *PostTaskByIDParams, authInfo runtime.ClientAuthInfoWriter) (*PostTaskByIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTaskByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postTaskById",
		Method:             "POST",
		PathPattern:        "/tasks/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostTaskByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTaskByIDCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
