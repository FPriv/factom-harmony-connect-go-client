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

// The chain that has been created or altered by this entry block.
type EBlockShortChain struct {
	// The ID for this chain on the Factom blockchain.
	ChainId string `json:"chain_id"`
	// An API link to retrieve all information about this chain.
	Href string `json:"href"`
}
