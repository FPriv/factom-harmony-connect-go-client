# CommitLong

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | The SHA256 Hash of this commit. | 
**TxId** | **string** | The entry credit transaction ID that resulted in this commit. | 
**Version** | **int32** | The type version of this transaction. | 
**CreatedAt** | **string** | The timestamp for when this commit was created. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: &#x60;YYYY-MM-DDThh:mm:ssZ&#x60; | 
**ChainId** | **string** | The ID of the chain that was altered by this commit. | 
**Weld** | **string** | The commit weld for this commit. This is the double hash (SHA256d) of the Entry Hash concatenated with the ChainID. | 
**EntryHash** | **string** | The unique identifier of the entry that was created by this commit. | 
**EntryCredits** | **int32** | The number of entry credits spent on this commit. | 
**Address** | **string** | The entry credit address that initiated this transaction. | 
**Signature** | **string** | The signature that authorized this commit. | 
**EcBlock** | [**EcBlockShort**](ECBlockShort.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


