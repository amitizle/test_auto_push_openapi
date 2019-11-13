/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

type TransferAssetRequest struct {
	// The ID of the Wallet to transfer from
	WalletId string `json:"wallet_id"`
	// The authorization password for the Wallet to transfer from
	Authorization string `json:"authorization"`
	// The Address of the recipient of the funds
	RecipientAddress string `json:"recipient_address"`
	// The ID of an Asset Type. It must be a string of length 0 to 40 characters. Allowed characters are lower- and uppercase letters, digits, dot (.), and hyphen (-). It must not end with an hyphen. Asset IDs are case-sensitive.
	AssetId string `json:"asset_id"`
	// The amount of assets to transfer
	Amount int32 `json:"amount"`
	// An app-customizable field to store additional private data relating to the transfer; the memo is shared between the sender and the receiver, but is not divulged to other parties
	Memo string `json:"memo"`
	// A user may request confirmation from the receiving party. If a public key of the approver is included in this optional field, the transaction will only become valid after the received signs it. The receiver will be able to decide whether to accept or reject the transfer by calling the /node/approve_task or the /node/reject_task respectively.
	RequireConfirmationFrom string `json:"require_confirmation_from,omitempty"`
}
