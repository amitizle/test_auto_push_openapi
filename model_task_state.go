/*
 * QEDIT - Asset Transfers
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.6.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sdk
// TaskState : A fine-grained status of the node's inner state. This field is subject to changes and should not be relied upon during development.
type TaskState string

// List of TaskState
const (
	PENDING_STATE TaskState = "pending_state"
	IN_PROGRESS_STATE TaskState = "in_progress_state"
	TX_GENERATED_STATE TaskState = "tx_generated_state"
	CONF_REQ_SENT_STATE TaskState = "conf_req_sent_state"
	CONF_REQ_RECEIVED_STATE TaskState = "conf_req_received_state"
	CONF_REQ_CANCELED_BY_SENDER_STATE TaskState = "conf_req_canceled_by_sender_state"
	CONF_REQ_CANCELED_BY_RECEIVER_STATE TaskState = "conf_req_canceled_by_receiver_state"
	TX_SUBMITTED_STATE TaskState = "tx_submitted_state"
	FAILURE_STATE TaskState = "failure_state"
	SUCCESS_STATE TaskState = "success_state"
)