// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TreeListNodesRequest tree list nodes request
// swagger:model treeListNodesRequest
type TreeListNodesRequest struct {

	// ancestors
	Ancestors bool `json:"Ancestors,omitempty"`

	// filter type
	FilterType TreeNodeType `json:"FilterType,omitempty"`

	// limit
	Limit int64 `json:"Limit,omitempty"`

	// node
	Node *TreeNode `json:"Node,omitempty"`

	// offset
	Offset int64 `json:"Offset,omitempty"`

	// recursive
	Recursive bool `json:"Recursive,omitempty"`

	// with commits
	WithCommits bool `json:"WithCommits,omitempty"`

	// with versions
	WithVersions bool `json:"WithVersions,omitempty"`
}

// Validate validates this tree list nodes request
func (m *TreeListNodesRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TreeListNodesRequest) validateFilterType(formats strfmt.Registry) error {

	if swag.IsZero(m.FilterType) { // not required
		return nil
	}

	if err := m.FilterType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("FilterType")
		}
		return err
	}

	return nil
}

func (m *TreeListNodesRequest) validateNode(formats strfmt.Registry) error {

	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TreeListNodesRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TreeListNodesRequest) UnmarshalBinary(b []byte) error {
	var res TreeListNodesRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
