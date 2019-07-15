# IdentityKeyData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The public key string in Base58 idpub format. | 
**ActivatedHeight** | **int32** | The height at which this key became valid. | [optional] 
**RetiredHeight** | **int32** | The expiration height of this key. &#x60;null&#x60; if this key is currently active. | [optional] 
**Priority** | **int32** | The level of this key within the hierarchy. A lower number indicates a key that allows a holder to replace higher numbered keys. The master key is priority 0. | [optional] 
**EntryHash** | **string** | The unique identifier of the entry on the Identity chain where this key was activated. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


