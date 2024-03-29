/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sdk

// The Rule used in the Issuance
type AnalyticsRule struct {
	// A `Namespace` describes what Asset IDs can be issued in an Issuance Rule. It is a string in the same format as `AssetId`. Additionally, if it ends with a wildcard character `*`, then the namespace covers all asset IDs with the namespace as a prefix. Without a final wildcard, the namespace covers exactly one asset ID. Example: The namespace `currencies.dollar` covers only this exact asset type, while `currencies.*` covers all asset types that start with `currencies.`. 
	Namespace string `json:"namespace"`
}
