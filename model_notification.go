/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sdk

// A notification including some metadata and a content field whose type depends on the type of notification.
type Notification struct {
	// The unique identifier of the notification. Can be used for pagination purposes.
	Id int32 `json:"id"`
	// Timestamp the notification was created at, in RFC-3339 format.
	CreatedAt string `json:"created_at"`
	// The ID of the wallet the notification relates to. May be null in cases there is no such specific wallet.
	WalletId string `json:"wallet_id,omitempty"`
	// The ID of the task the notification relates to. May be empty in cases there is no such specific task.
	TaskId string `json:"task_id"`
	NotificationType NotificationType `json:"notification_type"`
	Content map[string]interface{} `json:"content,omitempty"`
}