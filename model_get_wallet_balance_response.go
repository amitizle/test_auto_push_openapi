/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sdk

type GetWalletBalanceResponse struct {
	// The ID of the Wallet
	WalletId string `json:"wallet_id"`
	// The balances of the various Asset Types held by the Wallet
	Assets []BalanceForAsset `json:"assets"`
}