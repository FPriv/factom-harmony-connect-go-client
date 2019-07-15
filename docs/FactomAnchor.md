# FactomAnchor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Indicates the blockchain network that contains this anchor. | 
**Status** | **string** | The confirmation status of the anchor. Either pending or confirmed. | 
**CreatedAt** | **string** | The time at which this entry was created. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: &#x60;YYYY-MM-DDThh:mm:ss.ssssssZ&#x60; This will be null if the chain is not at least at the &#x60;factom&#x60; immutability stage. | [optional] 
**EntrySerialized** | **string** | The raw data that makes up the entry. | [optional] 
**EntryHash** | **string** | The unique identitfier of the entry. | [optional] 
**Dblock** | [**DBlockLink**](DBlockLink.md) |  | [optional] 
**MerkleBranch** | [**[]MerkleNode**](MerkleNode.md) | The branch of the merkle tree that represents this anchor. Presented as an array of Merkle nodes. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


