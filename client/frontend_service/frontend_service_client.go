// Code generated by go-swagger; DO NOT EDIT.

package frontend_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new frontend service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for frontend service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
FrontBootConf adds some data to the initial set of parameters loaded by the frontend
*/
func (a *Client) FrontBootConf(params *FrontBootConfParams) (*FrontBootConfOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontBootConfParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontBootConf",
		Method:             "GET",
		PathPattern:        "/frontend/bootconf",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontBootConfReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontBootConfOK), nil

}

/*
FrontLog sends a log from front php to back
*/
func (a *Client) FrontLog(params *FrontLogParams) (*FrontLogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontLogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontLog",
		Method:             "PUT",
		PathPattern:        "/frontend/frontlogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontLogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontLogOK), nil

}

/*
FrontMessages serves list of i18n messages
*/
func (a *Client) FrontMessages(params *FrontMessagesParams) (*FrontMessagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontMessagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontMessages",
		Method:             "GET",
		PathPattern:        "/frontend/messages/{Lang}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontMessagesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontMessagesOK), nil

}

/*
FrontPlugins serves list of i18n messages
*/
func (a *Client) FrontPlugins(params *FrontPluginsParams) (*FrontPluginsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontPluginsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontPlugins",
		Method:             "GET",
		PathPattern:        "/frontend/plugins/{Lang}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontPluginsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontPluginsOK), nil

}

/*
FrontPutBinary uploads frontend binaries avatars logos bg images
*/
func (a *Client) FrontPutBinary(params *FrontPutBinaryParams) (*FrontPutBinaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontPutBinaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontPutBinary",
		Method:             "POST",
		PathPattern:        "/frontend/binaries/{BinaryType}/{Uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontPutBinaryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontPutBinaryOK), nil

}

/*
FrontServeBinary serves frontend binaries directly avatars logos bg images
*/
func (a *Client) FrontServeBinary(params *FrontServeBinaryParams) (*FrontServeBinaryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontServeBinaryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontServeBinary",
		Method:             "GET",
		PathPattern:        "/frontend/binaries/{BinaryType}/{Uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontServeBinaryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontServeBinaryOK), nil

}

/*
FrontSession handles j w t
*/
func (a *Client) FrontSession(params *FrontSessionParams) (*FrontSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontSession",
		Method:             "POST",
		PathPattern:        "/frontend/session",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontSessionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontSessionOK), nil

}

/*
FrontState sends XML state registry
*/
func (a *Client) FrontState(params *FrontStateParams) (*FrontStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFrontStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FrontState",
		Method:             "GET",
		PathPattern:        "/frontend/state",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &FrontStateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*FrontStateOK), nil

}

/*
SettingsMenu sends a tree of nodes to be used a menu in the settings panel
*/
func (a *Client) SettingsMenu(params *SettingsMenuParams) (*SettingsMenuOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSettingsMenuParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SettingsMenu",
		Method:             "GET",
		PathPattern:        "/frontend/settings-menu",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SettingsMenuReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SettingsMenuOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
