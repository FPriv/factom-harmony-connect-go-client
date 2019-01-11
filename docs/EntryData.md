# EntryData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryHash** | **string** | The SHA256 Hash of this entry. | [optional] 
**Chain** | [**EntryLinkChain**](EntryLink_chain.md) |  | [optional] 
**CreatedAt** | **string** | The time when this entry was created. Sent in [ISO 8601 Format](https://en.wikipedia.org/wiki/ISO_8601). For example: &#x60;YYYY-MM-DDThh:mm:ssZ&#x60; | [optional] 
**ExternalIds** | **[]string** | Tags that can be used to identify your entry. You can search for records that contain a particular external_id using Connect. External IDs are returned in Base64. | [optional] 
**Content** | **string** | This is the data that is stored by the entry. Content will be sent in Base64 format. | [optional] 
**Stage** | **string** | The level of immutability that this entry has reached. | [optional] 
**Eblock** | [**EntryDataEblock**](Entry_data_eblock.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


