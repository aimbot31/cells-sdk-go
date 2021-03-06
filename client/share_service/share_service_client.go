// Code generated by go-swagger; DO NOT EDIT.

package share_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new share service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for share service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCell deletes a share room
*/
func (a *Client) DeleteCell(params *DeleteCellParams) (*DeleteCellOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCellParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCell",
		Method:             "DELETE",
		PathPattern:        "/share/cell/{Uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &DeleteCellReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCellOK), nil

}

/*
DeleteShareLink deletes share link
*/
func (a *Client) DeleteShareLink(params *DeleteShareLinkParams) (*DeleteShareLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteShareLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteShareLink",
		Method:             "DELETE",
		PathPattern:        "/share/link/{Uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &DeleteShareLinkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteShareLinkOK), nil

}

/*
GetCell loads a share room
*/
func (a *Client) GetCell(params *GetCellParams) (*GetCellOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCellParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCell",
		Method:             "GET",
		PathPattern:        "/share/cell/{Uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &GetCellReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCellOK), nil

}

/*
GetShareLink loads a share link with all infos
*/
func (a *Client) GetShareLink(params *GetShareLinkParams) (*GetShareLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetShareLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetShareLink",
		Method:             "GET",
		PathPattern:        "/share/link/{Uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &GetShareLinkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetShareLinkOK), nil

}

/*
ListSharedResources lists shared resources for current user or all users
*/
func (a *Client) ListSharedResources(params *ListSharedResourcesParams) (*ListSharedResourcesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSharedResourcesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListSharedResources",
		Method:             "POST",
		PathPattern:        "/share/resources",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListSharedResourcesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSharedResourcesOK), nil

}

/*
PutCell puts or create a share room
*/
func (a *Client) PutCell(params *PutCellParams) (*PutCellOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCellParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCell",
		Method:             "PUT",
		PathPattern:        "/share/cell",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutCellReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutCellOK), nil

}

/*
PutShareLink puts or create a share room
*/
func (a *Client) PutShareLink(params *PutShareLinkParams) (*PutShareLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutShareLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutShareLink",
		Method:             "PUT",
		PathPattern:        "/share/link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutShareLinkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutShareLinkOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
