/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

type AnalyticTransaction struct {
	Metadata AnalyticsTxMetadata `json:"metadata,omitempty"`
	// The detailed content of the transaction; format differs depending on the transaction type
	Content map[string]interface{} `json:"content,omitempty"`
}
