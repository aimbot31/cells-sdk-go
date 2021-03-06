// Code generated by go-swagger; DO NOT EDIT.

package user_meta_service

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

// NewPutUserMetaTagParams creates a new PutUserMetaTagParams object
// with the default values initialized.
func NewPutUserMetaTagParams() *PutUserMetaTagParams {
	var ()
	return &PutUserMetaTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutUserMetaTagParamsWithTimeout creates a new PutUserMetaTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutUserMetaTagParamsWithTimeout(timeout time.Duration) *PutUserMetaTagParams {
	var ()
	return &PutUserMetaTagParams{

		timeout: timeout,
	}
}

// NewPutUserMetaTagParamsWithContext creates a new PutUserMetaTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutUserMetaTagParamsWithContext(ctx context.Context) *PutUserMetaTagParams {
	var ()
	return &PutUserMetaTagParams{

		Context: ctx,
	}
}

// NewPutUserMetaTagParamsWithHTTPClient creates a new PutUserMetaTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutUserMetaTagParamsWithHTTPClient(client *http.Client) *PutUserMetaTagParams {
	var ()
	return &PutUserMetaTagParams{
		HTTPClient: client,
	}
}

/*PutUserMetaTagParams contains all the parameters to send to the API endpoint
for the put user meta tag operation typically these are written to a http.Request
*/
type PutUserMetaTagParams struct {

	/*Namespace*/
	Namespace string
	/*Body*/
	Body *models.RestPutUserMetaTagRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put user meta tag params
func (o *PutUserMetaTagParams) WithTimeout(timeout time.Duration) *PutUserMetaTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put user meta tag params
func (o *PutUserMetaTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put user meta tag params
func (o *PutUserMetaTagParams) WithContext(ctx context.Context) *PutUserMetaTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put user meta tag params
func (o *PutUserMetaTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put user meta tag params
func (o *PutUserMetaTagParams) WithHTTPClient(client *http.Client) *PutUserMetaTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put user meta tag params
func (o *PutUserMetaTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespace adds the namespace to the put user meta tag params
func (o *PutUserMetaTagParams) WithNamespace(namespace string) *PutUserMetaTagParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the put user meta tag params
func (o *PutUserMetaTagParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WithBody adds the body to the put user meta tag params
func (o *PutUserMetaTagParams) WithBody(body *models.RestPutUserMetaTagRequest) *PutUserMetaTagParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put user meta tag params
func (o *PutUserMetaTagParams) SetBody(body *models.RestPutUserMetaTagRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutUserMetaTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Namespace
	if err := r.SetPathParam("Namespace", o.Namespace); err != nil {
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
