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

type ABlockEntryData struct {
	// The content of the admin block entry.
	Content string `json:"content"`
	// The timestamp for when this entry was created. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: `YYYY-MM-DDThh:mm:ssZ`
	CreatedAt string `json:"created_at"`
	Ablock ABlockShort `json:"ablock"`
}