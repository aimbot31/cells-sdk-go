// Code generated by go-swagger; DO NOT EDIT.

package scheduler_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new scheduler service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scheduler service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PutJob put job API
*/
func (a *Client) PutJob(params *PutJobParams) (*PutJobOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutJobParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutJob",
		Method:             "PUT",
		PathPattern:        "/scheduler/jobs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutJobReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutJobOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
