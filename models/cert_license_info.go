// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CertLicenseInfo cert license info
// swagger:model certLicenseInfo
type CertLicenseInfo struct {

	// account name
	AccountName string `json:"AccountName,omitempty"`

	// expire time
	ExpireTime int32 `json:"ExpireTime,omitempty"`

	// features
	Features map[string]string `json:"Features,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// issue time
	IssueTime int32 `json:"IssueTime,omitempty"`

	// max peers
	MaxPeers int64 `json:"MaxPeers,omitempty"`

	// max users
	MaxUsers int64 `json:"MaxUsers,omitempty"`

	// server domain
	ServerDomain string `json:"ServerDomain,omitempty"`
}

// Validate validates this cert license info
func (m *CertLicenseInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertLicenseInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertLicenseInfo) UnmarshalBinary(b []byte) error {
	var res CertLicenseInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}