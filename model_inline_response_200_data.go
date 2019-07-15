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

// The search result. Includes an API link that allows you to retreive more information about the result. Also includes additional information based on the type of result reterned.
type InlineResponse200Data struct {
	// An API link to retrieve all information about this search result.
	Href string `json:"href,omitempty"`
}