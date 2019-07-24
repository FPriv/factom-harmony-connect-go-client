# ReceiptLongData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The timestamp for this entry. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: &#x60;YYYY-MM-DDThh:mm:ssZ&#x60; | 
**EntrySerialized** | **string** | The raw data that makes up the entry. | 
**EntryHash** | **string** | The unique identitfier of the entry. | 
**MerkleBranch** | [**[]MerkleNode**](MerkleNode.md) | The branch of the merkle tree that represents this entry. Presented as an array of Merkle nodes. | 
**Eblock** | [**EBlockLink**](EBlockLink.md) |  | 
**Dblock** | [**DBlockLink**](DBlockLink.md) |  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


