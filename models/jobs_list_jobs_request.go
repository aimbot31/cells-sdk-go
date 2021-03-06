// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// JobsListJobsRequest jobs list jobs request
// swagger:model jobsListJobsRequest
type JobsListJobsRequest struct {

	// events only
	EventsOnly bool `json:"EventsOnly,omitempty"`

	// load tasks
	LoadTasks JobsTaskStatus `json:"LoadTasks,omitempty"`

	// owner
	Owner string `json:"Owner,omitempty"`

	// timers only
	TimersOnly bool `json:"TimersOnly,omitempty"`
}

// Validate validates this jobs list jobs request
func (m *JobsListJobsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobsListJobsRequest) validateLoadTasks(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadTasks) { // not required
		return nil
	}

	if err := m.LoadTasks.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LoadTasks")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobsListJobsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobsListJobsRequest) UnmarshalBinary(b []byte) error {
	var res JobsListJobsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
