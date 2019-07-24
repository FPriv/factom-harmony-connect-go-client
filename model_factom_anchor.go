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

type FactomAnchor struct {
	// Indicates the blockchain network that contains this anchor.
	Network string `json:"network"`
	// The confirmation status of the anchor. Either pending or confirmed.
	Status string `json:"status"`
	// The time at which this entry was created. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: `YYYY-MM-DDThh:mm:ss.ssssssZ` This will be null if the chain is not at least at the `factom` immutability stage.
	CreatedAt string `json:"created_at,omitempty"`
	// The raw data that makes up the entry.
	EntrySerialized string `json:"entry_serialized,omitempty"`
	// The unique identitfier of the entry.
	EntryHash string `json:"entry_hash,omitempty"`
	Dblock DBlockLink `json:"dblock,omitempty"`
	// The branch of the merkle tree that represents this anchor. Presented as an array of Merkle nodes.
	MerkleBranch []MerkleNode `json:"merkle_branch,omitempty"`
}
