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

type MerkleNode struct {
	// The top of this node of the Merkle tree.
	Top string `json:"top,omitempty"`
	// The right branch of this node of the Merkle tree.
	Right string `json:"right,omitempty"`
	// The left branch of this node of the Merkle tree.
	Left string `json:"left,omitempty"`
}
