/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sdk

type AnalyticsTxMetadata struct {
	Type AnalyticsTxType `json:"type,omitempty"`
	// the QEDIT-generated hash of the transaction
	TxHash string `json:"tx_hash,omitempty"`
	// The height of the Block (inside the Blockchain) that the transaction is a part of
	BlockHeight int32 `json:"block_height,omitempty"`
	// the hash of the Block (inside the Blockchain) that the transaction is a part of
	BlockHash string `json:"block_hash,omitempty"`
	// UTC time of creation of the time the Block containing the transaction was created in RFC-3339 format
	Timestamp string `json:"timestamp,omitempty"`
	// The serial number within the Block of the transaction relative to other QEDIT transactions; indexing is 0-based
	IndexInBlock int32 `json:"index_in_block,omitempty"`
}