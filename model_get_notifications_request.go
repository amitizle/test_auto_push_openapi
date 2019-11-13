/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package goqedit

// The request body determines which filters to apply to the notifications. All filters are optional, while the maximal number of results defaults to 100 and the order defaults to descending.
type GetNotificationsRequest struct {
	// Fetch notifications that relate to this wallet ID only
	WalletId string             `json:"wallet_id,omitempty"`
	Types    []NotificationType `json:"types,omitempty"`
	// Fetch notifications reported at this timestamp or later/earlier. Later notifications will be fetched if the order is ascending, and earlier ones if the order is descending. The timestamp is expected to be in RFC-3339 format.
	StartingWithTime string `json:"starting_with_time,omitempty"`
	// Fetch notifications after this given ID (not including the notification with the given ID). This is meant to facilitate pagination. Later notifications will be fetched if ascending order is selected, and earlier ones if descending order is selected. Warning - do not assume anything about the implementation of after_id; the values of the ID are intended to be copy-pasted from retrieved notifications for pagination.
	AfterId int32     `json:"after_id,omitempty"`
	Order   OrderEnum `json:"order,omitempty"`
	// The maximal number of results to fetch
	MaxResults int32 `json:"max_results,omitempty"`
}
