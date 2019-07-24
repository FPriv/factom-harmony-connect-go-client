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

type KeyList struct {
	// The keys returned in this set.
	Data []IdentityKey `json:"data"`
	// The index of the first key returned from the total set (Starting from 0).
	Offset int32 `json:"offset"`
	// The number of keys returned per page.
	Limit int32 `json:"limit"`
	// The total number of keys seen.
	Count int32 `json:"count"`
}
