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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserParams creates a new GetUserParams object
// with the default values initialized.
func NewGetUserParams() *GetUserParams {
	var ()
	return &GetUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserParamsWithTimeout creates a new GetUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserParamsWithTimeout(timeout time.Duration) *GetUserParams {
	var ()
	return &GetUserParams{

		timeout: timeout,
	}
}

// NewGetUserParamsWithContext creates a new GetUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserParamsWithContext(ctx context.Context) *GetUserParams {
	var ()
	return &GetUserParams{

		Context: ctx,
	}
}

// NewGetUserParamsWithHTTPClient creates a new GetUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserParamsWithHTTPClient(client *http.Client) *GetUserParams {
	var ()
	return &GetUserParams{
		HTTPClient: client,
	}
}

/*GetUserParams contains all the parameters to send to the API endpoint
for the get user operation typically these are written to a http.Request
*/
type GetUserParams struct {

	/*GroupLabel*/
	GroupLabel *string
	/*GroupPath*/
	GroupPath *string
	/*IsGroup
	  Group specific data.

	*/
	IsGroup *bool
	/*Login*/
	Login string
	/*OldPassword*/
	OldPassword *string
	/*Password*/
	Password *string
	/*PoliciesContextEditable*/
	PoliciesContextEditable *bool
	/*UUID*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user params
func (o *GetUserParams) WithTimeout(timeout time.Duration) *GetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user params
func (o *GetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user params
func (o *GetUserParams) WithContext(ctx context.Context) *GetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user params
func (o *GetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) WithHTTPClient(client *http.Client) *GetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupLabel adds the groupLabel to the get user params
func (o *GetUserParams) WithGroupLabel(groupLabel *string) *GetUserParams {
	o.SetGroupLabel(groupLabel)
	return o
}

// SetGroupLabel adds the groupLabel to the get user params
func (o *GetUserParams) SetGroupLabel(groupLabel *string) {
	o.GroupLabel = groupLabel
}

// WithGroupPath adds the groupPath to the get user params
func (o *GetUserParams) WithGroupPath(groupPath *string) *GetUserParams {
	o.SetGroupPath(groupPath)
	return o
}

// SetGroupPath adds the groupPath to the get user params
func (o *GetUserParams) SetGroupPath(groupPath *string) {
	o.GroupPath = groupPath
}

// WithIsGroup adds the isGroup to the get user params
func (o *GetUserParams) WithIsGroup(isGroup *bool) *GetUserParams {
	o.SetIsGroup(isGroup)
	return o
}

// SetIsGroup adds the isGroup to the get user params
func (o *GetUserParams) SetIsGroup(isGroup *bool) {
	o.IsGroup = isGroup
}

// WithLogin adds the login to the get user params
func (o *GetUserParams) WithLogin(login string) *GetUserParams {
	o.SetLogin(login)
	return o
}

// SetLogin adds the login to the get user params
func (o *GetUserParams) SetLogin(login string) {
	o.Login = login
}

// WithOldPassword adds the oldPassword to the get user params
func (o *GetUserParams) WithOldPassword(oldPassword *string) *GetUserParams {
	o.SetOldPassword(oldPassword)
	return o
}

// SetOldPassword adds the oldPassword to the get user params
func (o *GetUserParams) SetOldPassword(oldPassword *string) {
	o.OldPassword = oldPassword
}

// WithPassword adds the password to the get user params
func (o *GetUserParams) WithPassword(password *string) *GetUserParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the get user params
func (o *GetUserParams) SetPassword(password *string) {
	o.Password = password
}

// WithPoliciesContextEditable adds the policiesContextEditable to the get user params
func (o *GetUserParams) WithPoliciesContextEditable(policiesContextEditable *bool) *GetUserParams {
	o.SetPoliciesContextEditable(policiesContextEditable)
	return o
}

// SetPoliciesContextEditable adds the policiesContextEditable to the get user params
func (o *GetUserParams) SetPoliciesContextEditable(policiesContextEditable *bool) {
	o.PoliciesContextEditable = policiesContextEditable
}

// WithUUID adds the uuid to the get user params
func (o *GetUserParams) WithUUID(uuid *string) *GetUserParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the get user params
func (o *GetUserParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.GroupLabel != nil {

		// query param GroupLabel
		var qrGroupLabel string
		if o.GroupLabel != nil {
			qrGroupLabel = *o.GroupLabel
		}
		qGroupLabel := qrGroupLabel
		if qGroupLabel != "" {
			if err := r.SetQueryParam("GroupLabel", qGroupLabel); err != nil {
				return err
			}
		}

	}

	if o.GroupPath != nil {

		// query param GroupPath
		var qrGroupPath string
		if o.GroupPath != nil {
			qrGroupPath = *o.GroupPath
		}
		qGroupPath := qrGroupPath
		if qGroupPath != "" {
			if err := r.SetQueryParam("GroupPath", qGroupPath); err != nil {
				return err
			}
		}

	}

	if o.IsGroup != nil {

		// query param IsGroup
		var qrIsGroup bool
		if o.IsGroup != nil {
			qrIsGroup = *o.IsGroup
		}
		qIsGroup := swag.FormatBool(qrIsGroup)
		if qIsGroup != "" {
			if err := r.SetQueryParam("IsGroup", qIsGroup); err != nil {
				return err
			}
		}

	}

	// path param Login
	if err := r.SetPathParam("Login", o.Login); err != nil {
		return err
	}

	if o.OldPassword != nil {

		// query param OldPassword
		var qrOldPassword string
		if o.OldPassword != nil {
			qrOldPassword = *o.OldPassword
		}
		qOldPassword := qrOldPassword
		if qOldPassword != "" {
			if err := r.SetQueryParam("OldPassword", qOldPassword); err != nil {
				return err
			}
		}

	}

	if o.Password != nil {

		// query param Password
		var qrPassword string
		if o.Password != nil {
			qrPassword = *o.Password
		}
		qPassword := qrPassword
		if qPassword != "" {
			if err := r.SetQueryParam("Password", qPassword); err != nil {
				return err
			}
		}

	}

	if o.PoliciesContextEditable != nil {

		// query param PoliciesContextEditable
		var qrPoliciesContextEditable bool
		if o.PoliciesContextEditable != nil {
			qrPoliciesContextEditable = *o.PoliciesContextEditable
		}
		qPoliciesContextEditable := swag.FormatBool(qrPoliciesContextEditable)
		if qPoliciesContextEditable != "" {
			if err := r.SetQueryParam("PoliciesContextEditable", qPoliciesContextEditable); err != nil {
				return err
			}
		}

	}

	if o.UUID != nil {

		// query param Uuid
		var qrUUID string
		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {
			if err := r.SetQueryParam("Uuid", qUUID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
