// Code generated by go-swagger; DO NOT EDIT.

package user_service

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

// NewBindUserParams creates a new BindUserParams object
// with the default values initialized.
func NewBindUserParams() *BindUserParams {
	var ()
	return &BindUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBindUserParamsWithTimeout creates a new BindUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBindUserParamsWithTimeout(timeout time.Duration) *BindUserParams {
	var ()
	return &BindUserParams{

		timeout: timeout,
	}
}

// NewBindUserParamsWithContext creates a new BindUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewBindUserParamsWithContext(ctx context.Context) *BindUserParams {
	var ()
	return &BindUserParams{

		Context: ctx,
	}
}

// NewBindUserParamsWithHTTPClient creates a new BindUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBindUserParamsWithHTTPClient(client *http.Client) *BindUserParams {
	var ()
	return &BindUserParams{
		HTTPClient: client,
	}
}

/*BindUserParams contains all the parameters to send to the API endpoint
for the bind user operation typically these are written to a http.Request
*/
type BindUserParams struct {

	/*Login*/
	Login string
	/*Body*/
	Body *models.IdmUser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bind user params
func (o *BindUserParams) WithTimeout(timeout time.Duration) *BindUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bind user params
func (o *BindUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bind user params
func (o *BindUserParams) WithContext(ctx context.Context) *BindUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bind user params
func (o *BindUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bind user params
func (o *BindUserParams) WithHTTPClient(client *http.Client) *BindUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bind user params
func (o *BindUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLogin adds the login to the bind user params
func (o *BindUserParams) WithLogin(login string) *BindUserParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the bind user params
func (o *BindUserParams) SetLogin(login string) {
	o.Login = login
}

// WithBody adds the body to the bind user params
func (o *BindUserParams) WithBody(body *models.IdmUser) *BindUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the bind user params
func (o *BindUserParams) SetBody(body *models.IdmUser) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BindUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param Login
	if err := r.SetPathParam("Login", o.Login); err != nil {
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
