// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user meta service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user meta service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteUserMetaTags lists tags for a given namespace
*/
func (a *Client) DeleteUserMetaTags(params *DeleteUserMetaTagsParams) (*DeleteUserMetaTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserMetaTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteUserMetaTags",
		Method:             "DELETE",
		PathPattern:        "/user-meta/tags/{Namespace}/{Tags}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &DeleteUserMetaTagsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserMetaTagsOK), nil

}

/*
ListUserMetaNamespace lists defined meta namespaces
*/
func (a *Client) ListUserMetaNamespace(params *ListUserMetaNamespaceParams) (*ListUserMetaNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserMetaNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListUserMetaNamespace",
		Method:             "GET",
		PathPattern:        "/user-meta/namespace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListUserMetaNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUserMetaNamespaceOK), nil

}

/*
ListUserMetaTags lists tags for a given namespace
*/
func (a *Client) ListUserMetaTags(params *ListUserMetaTagsParams) (*ListUserMetaTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUserMetaTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListUserMetaTags",
		Method:             "GET",
		PathPattern:        "/user-meta/tags/{Namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &ListUserMetaTagsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListUserMetaTagsOK), nil

}

/*
PutUserMetaTag lists tags for a given namespace
*/
func (a *Client) PutUserMetaTag(params *PutUserMetaTagParams) (*PutUserMetaTagOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutUserMetaTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutUserMetaTag",
		Method:             "POST",
		PathPattern:        "/user-meta/tags/{Namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &PutUserMetaTagReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutUserMetaTagOK), nil

}

/*
SearchUserMeta searches a list of meta by node Id or by user id and by namespace
*/
func (a *Client) SearchUserMeta(params *SearchUserMetaParams) (*SearchUserMetaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchUserMetaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SearchUserMeta",
		Method:             "POST",
		PathPattern:        "/user-meta/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SearchUserMetaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchUserMetaOK), nil

}

/*
UpdateUserMeta updates delete user meta
*/
func (a *Client) UpdateUserMeta(params *UpdateUserMetaParams) (*UpdateUserMetaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserMetaParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateUserMeta",
		Method:             "PUT",
		PathPattern:        "/user-meta/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &UpdateUserMetaReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUserMetaOK), nil

}

/*
UpdateUserMetaNamespace admins update namespaces
*/
func (a *Client) UpdateUserMetaNamespace(params *UpdateUserMetaNamespaceParams) (*UpdateUserMetaNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserMetaNamespaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateUserMetaNamespace",
		Method:             "PUT",
		PathPattern:        "/user-meta/namespace",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &UpdateUserMetaNamespaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUserMetaNamespaceOK), nil

}

/*
UserBookmarks specials API for bookmarks will load user meta and the associated nodes and return as a node list
*/
func (a *Client) UserBookmarks(params *UserBookmarksParams) (*UserBookmarksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserBookmarksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UserBookmarks",
		Method:             "POST",
		PathPattern:        "/user-meta/bookmarks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &UserBookmarksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UserBookmarksOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
