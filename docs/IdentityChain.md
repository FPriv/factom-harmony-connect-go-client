# IdentityChain

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | The Identity Chain’s schema version. This details the format of this digital identity. More information about Factom Identity Chain schemas can be seen [here](). | 
**Stage** | **string** | The immutability stage that this chain has reached. The identity can be considered in use once it reaches the &#x60;factom&#x60; stage. | 
**CreatedHeight** | **int32** | The block height at which this chain was written into the Factom blockchain. This is &#x60;null&#x60; if the chain has not reached the &#x60;factom&#x60; stage. | [optional] 
**ChainId** | **string** | The Chain ID for this identity chain. | 
**Name** | **[]string** | A unique array of strings that are associated with this identity. | 
**ActiveKeys** | [**[]IdentityKeyShort**](IdentityKeyShort.md) | Contains the currently active public keys for this identity. | 
**AllKeysHref** | **string** | An API link to retrieve the keys for this identity. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


