/*
 * Harmony Connect
 *
 * An easy to use API that helps you access the Factom blockchain.
 *
 * API version: 1.0.19
 * Contact: harmony-support@factom.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package connectclient

// A single factoid transaction.
type Transaction struct {
	// The transaction ID for this transaction.
	TxId string `json:"tx_id,omitempty"`
	// The timestamp for this transaction. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: `YYYY-MM-DDThh:mm:ssZ`
	CreatedAt string `json:"created_at,omitempty"`
	// The number of factoids coming in from the sender.
	FctTotalInputs string `json:"fct_total_inputs,omitempty"`
	// The number of factoids going out to the receiver.
	FctTotalOutputs string `json:"fct_total_outputs,omitempty"`
	// The number of entry credits that have been purchased in this transaction.
	EcCreated string `json:"ec_created,omitempty"`
	// The price of an entry credit at the time of this transaction.
	EcRate string `json:"ec_rate,omitempty"`
	// The fee paid to the processor of this transaction.
	FctFee string `json:"fct_fee,omitempty"`
	// A link to the directory block that contains this transaction.
	Dblock string `json:"dblock,omitempty"`
	// A link to the factoid block that contains this transaction.
	Fblock string `json:"fblock,omitempty"`
	// All of the input transactions for this transaction.
	Inputs string `json:"inputs,omitempty"`
	// All of the output transactions for this transaction.
	Outputs string `json:"outputs,omitempty"`
}