/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

type GetPublicKeyResponse struct {
	// The ID of the Wallet the pubic key belongs to
	WalletId string `json:"wallet_id"`
	// The public key of the Wallet
	PublicKey string `json:"public_key"`
}
