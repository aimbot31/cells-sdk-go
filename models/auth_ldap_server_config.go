// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AuthLdapServerConfig auth ldap server config
// swagger:model authLdapServerConfig
type AuthLdapServerConfig struct {

	// bind d n
	BindDN string `json:"BindDN,omitempty"`

	// bind p w
	BindPW string `json:"BindPW,omitempty"`

	// config Id
	ConfigID string `json:"ConfigId,omitempty"`

	// connection
	Connection string `json:"Connection,omitempty"`

	// domain name
	DomainName string `json:"DomainName,omitempty"`

	// host
	Host string `json:"Host,omitempty"`

	// mapping rules
	MappingRules []*AuthLdapMapping `json:"MappingRules"`

	// member of mapping
	MemberOfMapping *AuthLdapMemberOfMapping `json:"MemberOfMapping,omitempty"`

	// page size
	PageSize int32 `json:"PageSize,omitempty"`

	// role prefix
	RolePrefix string `json:"RolePrefix,omitempty"`

	// root c a
	RootCA string `json:"RootCA,omitempty"`

	// To be converted to []byte
	RootCAData string `json:"RootCAData,omitempty"`

	// schedule
	Schedule string `json:"Schedule,omitempty"`

	// scheduler details
	SchedulerDetails string `json:"SchedulerDetails,omitempty"`

	// skip verify certificate
	SkipVerifyCertificate bool `json:"SkipVerifyCertificate,omitempty"`

	// user
	User *AuthLdapSearchFilter `json:"User,omitempty"`
}

// Validate validates this auth ldap server config
func (m *AuthLdapServerConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMappingRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberOfMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthLdapServerConfig) validateMappingRules(formats strfmt.Registry) error {

	if swag.IsZero(m.MappingRules) { // not required
		return nil
	}

	for i := 0; i < len(m.MappingRules); i++ {
		if swag.IsZero(m.MappingRules[i]) { // not required
			continue
		}

		if m.MappingRules[i] != nil {
			if err := m.MappingRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("MappingRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthLdapServerConfig) validateMemberOfMapping(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberOfMapping) { // not required
		return nil
	}

	if m.MemberOfMapping != nil {
		if err := m.MemberOfMapping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("MemberOfMapping")
			}
			return err
		}
	}

	return nil
}

func (m *AuthLdapServerConfig) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("User")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthLdapServerConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthLdapServerConfig) UnmarshalBinary(b []byte) error {
	var res AuthLdapServerConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}