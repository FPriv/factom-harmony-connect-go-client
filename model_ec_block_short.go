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

// A single entry credit block from the factom blockchain.
type EcBlockShort struct {
	// The SHA256 Hash of this entry credit block.
	Hash string `json:"hash,omitempty"`
	// An API link to obtain the full set of data for this entry credit block.
	Href string `json:"href,omitempty"`
	Dblock EcBlockLongDblock `json:"dblock,omitempty"`
}
