/*
 * Pydio Cells Rest API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package api

// TimeRangeRequest contains the parameter to configure the query to  retrieve the number of audit events of this type for a given time range defined by last timestamp and a range type.
type LogTimeRangeRequest struct {

	MsgId string `json:"MsgId,omitempty"`

	TimeRangeType string `json:"TimeRangeType,omitempty"`

	RefTime int32 `json:"RefTime,omitempty"`
}