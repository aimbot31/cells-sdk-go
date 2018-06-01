// Code generated by go-swagger; DO NOT EDIT.

package enterprise_log_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new enterprise log service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for enterprise log service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Audit auditables logs in Json or c s v format
*/
func (a *Client) Audit(params *AuditParams) (*AuditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuditParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Audit",
		Method:             "POST",
		PathPattern:        "/log/audit",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &AuditReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuditOK), nil

}

/*
AuditChartData retrieves aggregated audit logs to generate charts
*/
func (a *Client) AuditChartData(params *AuditChartDataParams) (*AuditChartDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuditChartDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AuditChartData",
		Method:             "POST",
		PathPattern:        "/log/audit/chartdata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &AuditChartDataReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuditChartDataOK), nil

}

/*
AuditExport auditables logs in Json or c s v format
*/
func (a *Client) AuditExport(params *AuditExportParams) (*AuditExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuditExportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AuditExport",
		Method:             "POST",
		PathPattern:        "/log/audit/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &AuditExportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuditExportOK), nil

}

/*
SyslogExport technicals logs in Json or c s v format
*/
func (a *Client) SyslogExport(params *SyslogExportParams) (*SyslogExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyslogExportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SyslogExport",
		Method:             "POST",
		PathPattern:        "/log/sys/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https", "wss"},
		Params:             params,
		Reader:             &SyslogExportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SyslogExportOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}