# EthereumAnchor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | **string** | Indicates the blockchain network that contains this anchor. | 
**Status** | **string** | The confirmation status of the anchor. Either pending or confirmed. | 
**WindowStartHeight** | **int32** | The height of the first ethereum block that contains an anchor for this directory block. Valid anchors for the Directory block in question may also be contained in subsequent ethereum blocks. | [optional] 
**WindowMr** | **string** | Window Merkle Root. This is the Merkle root that was published in the ethereum transaction that anchors this directory block. The root represents every one of the 1000 Directory blocks that&#39;s included in this anchor. | [optional] 
**MerkleBranch** | [**[]MerkleNode**](MerkleNode.md) | The merkle branch that leads to the Directory block in question. Relates to the Window Merkle root. Presented as an array of Merkle nodes. | [optional] 
**ContractAddress** | **string** | The address of the contract that initiated the anchor transaction. | [optional] 
**TxId** | **string** | The ethereum transaction ID for the transaction that includes the anchor message. | [optional] 
**TxIndex** | **int32** | The index of the transaction within the block. | [optional] 
**BlockHash** | **string** | The hash of the ethereum block that contains the anchor transaction. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


