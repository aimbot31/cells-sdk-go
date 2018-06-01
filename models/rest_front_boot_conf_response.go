// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RestFrontBootConfResponse rest front boot conf response
// swagger:model restFrontBootConfResponse
type RestFrontBootConfResponse struct {

	// Json data
	JSONData map[string]string `json:"JsonData,omitempty"`
}

// Validate validates this rest front boot conf response
func (m *RestFrontBootConfResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RestFrontBootConfResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestFrontBootConfResponse) UnmarshalBinary(b []byte) error {
	var res RestFrontBootConfResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}