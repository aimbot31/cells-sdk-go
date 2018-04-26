/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

type TreeSearchRequest struct {

	Query *TreeQuery `json:"Query,omitempty"`

	Size int32 `json:"Size,omitempty"`

	From int32 `json:"From,omitempty"`

	Details bool `json:"Details,omitempty"`

	Facet string `json:"Facet,omitempty"`
}