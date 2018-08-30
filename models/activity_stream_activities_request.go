// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ActivityStreamActivitiesRequest activity stream activities request
// swagger:model activityStreamActivitiesRequest
type ActivityStreamActivitiesRequest struct {

	// as digest
	AsDigest bool `json:"AsDigest,omitempty"`

	// box name
	BoxName string `json:"BoxName,omitempty"`

	// context
	Context ActivityStreamContext `json:"Context,omitempty"`

	// context data
	ContextData string `json:"ContextData,omitempty"`

	// language
	Language string `json:"Language,omitempty"`

	// limit
	Limit int64 `json:"Limit,string,omitempty"`

	// offset
	Offset int64 `json:"Offset,string,omitempty"`

	// point of view
	PointOfView ActivitySummaryPointOfView `json:"PointOfView,omitempty"`

	// stream filter
	StreamFilter string `json:"StreamFilter,omitempty"`

	// unread count only
	UnreadCountOnly bool `json:"UnreadCountOnly,omitempty"`
}

// Validate validates this activity stream activities request
func (m *ActivityStreamActivitiesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePointOfView(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActivityStreamActivitiesRequest) validateContext(formats strfmt.Registry) error {

	if swag.IsZero(m.Context) { // not required
		return nil
	}

	if err := m.Context.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Context")
		}
		return err
	}

	return nil
}

func (m *ActivityStreamActivitiesRequest) validatePointOfView(formats strfmt.Registry) error {

	if swag.IsZero(m.PointOfView) { // not required
		return nil
	}

	if err := m.PointOfView.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PointOfView")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityStreamActivitiesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityStreamActivitiesRequest) UnmarshalBinary(b []byte) error {
	var res ActivityStreamActivitiesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
