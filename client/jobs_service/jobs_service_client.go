// Code generated by go-swagger; DO NOT EDIT.

package jobs_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new jobs service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for jobs service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
UserControlJob sends control commands to one or many jobs tasks
*/
func (a *Client) UserControlJob(params *UserControlJobParams) (*UserControlJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserControlJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserControlJob",
		Method:             "PUT",
		PathPattern:        "/jobs/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &UserControlJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserControlJobOK), nil

}

/*
UserCreateJob creates a predefined job to be run directly
*/
func (a *Client) UserCreateJob(params *UserCreateJobParams) (*UserCreateJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserCreateJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserCreateJob",
		Method:             "PUT",
		PathPattern:        "/jobs/user/{JobName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &UserCreateJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserCreateJobOK), nil

}

/*
UserDeleteTasks sends a control command to clean tasks on a given job
*/
func (a *Client) UserDeleteTasks(params *UserDeleteTasksParams) (*UserDeleteTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserDeleteTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserDeleteTasks",
		Method:             "POST",
		PathPattern:        "/jobs/tasks/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &UserDeleteTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserDeleteTasksOK), nil

}

/*
UserListJobs lists jobs associated with current user
*/
func (a *Client) UserListJobs(params *UserListJobsParams) (*UserListJobsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserListJobsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserListJobs",
		Method:             "POST",
		PathPattern:        "/jobs/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &UserListJobsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserListJobsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
