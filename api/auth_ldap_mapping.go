/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type AuthLdapMapping struct {

	LeftAttribute string `json:"LeftAttribute,omitempty"`

	RightAttribute string `json:"RightAttribute,omitempty"`

	RuleString string `json:"RuleString,omitempty"`

	RolePrefix string `json:"RolePrefix,omitempty"`
}