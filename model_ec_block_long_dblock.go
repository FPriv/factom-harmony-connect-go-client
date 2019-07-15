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

// The directory block that contains this entry credit block.
type EcBlockLongDblock struct {
	// The SHA256 hash of this directory block.
	Hash string `json:"hash,omitempty"`
	// An API link to this directory block.
	Href string `json:"href,omitempty"`
}