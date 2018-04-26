/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type IdmWorkspace struct {

	UUID string `json:"UUID,omitempty"`

	Label string `json:"Label,omitempty"`

	Description string `json:"Description,omitempty"`

	Slug string `json:"Slug,omitempty"`

	Scope *IdmWorkspaceScope `json:"Scope,omitempty"`

	LastUpdated int32 `json:"LastUpdated,omitempty"`

	Policies []ServiceResourcePolicy `json:"Policies,omitempty"`

	Attributes string `json:"Attributes,omitempty"`

	RootNodes []string `json:"RootNodes,omitempty"`

	PoliciesContextEditable bool `json:"PoliciesContextEditable,omitempty"`
}