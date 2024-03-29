// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddTask(params *AddTaskParams, opts ...ClientOption) (*AddTaskCreated, error)

	DeleteTask(params *DeleteTaskParams, opts ...ClientOption) (*DeleteTaskNoContent, error)

	GetTaskByID(params *GetTaskByIDParams, opts ...ClientOption) (*GetTaskByIDOK, error)

	ListTasks(params *ListTasksParams, opts ...ClientOption) (*ListTasksOK, error)

	UpdateTask(params *UpdateTaskParams, opts ...ClientOption) (*UpdateTaskOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddTask add task API
*/
func (a *Client) AddTask(params *AddTaskParams, opts ...ClientOption) (*AddTaskCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addTask",
		Method:             "POST",
		PathPattern:        "/task",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddTaskCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteTask delete task API
*/
func (a *Client) DeleteTask(params *DeleteTaskParams, opts ...ClientOption) (*DeleteTaskNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteTask",
		Method:             "DELETE",
		PathPattern:        "/task/{ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteTaskNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetTaskByID get task by ID API
*/
func (a *Client) GetTaskByID(params *GetTaskByIDParams, opts ...ClientOption) (*GetTaskByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskByIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTaskByID",
		Method:             "GET",
		PathPattern:        "/task/{ID}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTaskByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTaskByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTaskByID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListTasks return tasks
*/
func (a *Client) ListTasks(params *ListTasksParams, opts ...ClientOption) (*ListTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTasks",
		Method:             "GET",
		PathPattern:        "/task",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateTask update task API
*/
func (a *Client) UpdateTask(params *UpdateTaskParams, opts ...ClientOption) (*UpdateTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateTask",
		Method:             "PUT",
		PathPattern:        "/task",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
