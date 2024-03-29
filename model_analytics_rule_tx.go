/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

// The data of a particular Rule changing transaction
type AnalyticsRuleTx struct {
	// The public key of the Wallet used to create the Rule
	SenderPublicKey string `json:"sender_public_key,omitempty"`
	// The details of the Rules added in this transaction
	RulesToAdd []AnalyticsRuleDefinition `json:"rules_to_add,omitempty"`
	// The details of the Rules deleted in this transaction
	RulesToDelete []AnalyticsRuleDefinition `json:"rules_to_delete,omitempty"`
	// The nonce used to make this Rule transaction unique
	Nonce int32 `json:"nonce,omitempty"`
	// The signature authorizing the Rule changes, made by the Wallet that made the Rule changes
	Signature string `json:"signature,omitempty"`
}
