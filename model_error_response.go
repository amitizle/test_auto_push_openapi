/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

type ErrorResponse struct {
	// The error code returned from the server
	ErrorCode int32 `json:"error_code"`
	// The error message returned from the server
	Message string `json:"message,omitempty"`
}
