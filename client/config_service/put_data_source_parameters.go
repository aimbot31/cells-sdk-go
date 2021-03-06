// Code generated by go-swagger; DO NOT EDIT.

package config_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// NewPutDataSourceParams creates a new PutDataSourceParams object
// with the default values initialized.
func NewPutDataSourceParams() *PutDataSourceParams {
	var ()
	return &PutDataSourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutDataSourceParamsWithTimeout creates a new PutDataSourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutDataSourceParamsWithTimeout(timeout time.Duration) *PutDataSourceParams {
	var ()
	return &PutDataSourceParams{

		timeout: timeout,
	}
}

// NewPutDataSourceParamsWithContext creates a new PutDataSourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutDataSourceParamsWithContext(ctx context.Context) *PutDataSourceParams {
	var ()
	return &PutDataSourceParams{

		Context: ctx,
	}
}

// NewPutDataSourceParamsWithHTTPClient creates a new PutDataSourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutDataSourceParamsWithHTTPClient(client *http.Client) *PutDataSourceParams {
	var ()
	return &PutDataSourceParams{
		HTTPClient: client,
	}
}

/*PutDataSourceParams contains all the parameters to send to the API endpoint
for the put data source operation typically these are written to a http.Request
*/
type PutDataSourceParams struct {

	/*Name*/
	Name string
	/*Body*/
	Body *models.ObjectDataSource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put data source params
func (o *PutDataSourceParams) WithTimeout(timeout time.Duration) *PutDataSourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put data source params
func (o *PutDataSourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put data source params
func (o *PutDataSourceParams) WithContext(ctx context.Context) *PutDataSourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put data source params
func (o *PutDataSourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put data source params
func (o *PutDataSourceParams) WithHTTPClient(client *http.Client) *PutDataSourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put data source params
func (o *PutDataSourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the put data source params
func (o *PutDataSourceParams) WithName(name string) *PutDataSourceParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the put data source params
func (o *PutDataSourceParams) SetName(name string) {
	o.Name = name
}

// WithBody adds the body to the put data source params
func (o *PutDataSourceParams) WithBody(body *models.ObjectDataSource) *PutDataSourceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put data source params
func (o *PutDataSourceParams) SetBody(body *models.ObjectDataSource) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutDataSourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Name
	if err := r.SetPathParam("Name", o.Name); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
