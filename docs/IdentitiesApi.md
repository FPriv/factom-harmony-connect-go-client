# \IdentitiesApi

All URIs are relative to *https://ephemeral.api.factom.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdChainbyId**](IdentitiesApi.md#GetIdChainbyId) | **Get** /identities/{identity_chain_id} | Get Identity Chain Info
[**GetKeybyEntryHash**](IdentitiesApi.md#GetKeybyEntryHash) | **Get** /identities/{identity_chain_id}/keys/{key_string} | Get Key Info
[**GetKeysbyIdChainId**](IdentitiesApi.md#GetKeysbyIdChainId) | **Get** /identities/{identity_chain_id}/keys | Get Identity Chain&#39;s Keys
[**PostIdChain**](IdentitiesApi.md#PostIdChain) | **Post** /identities | Create Identity Chain
[**PostKeytoIdChainId**](IdentitiesApi.md#PostKeytoIdChainId) | **Post** /identities/{identity_chain_id}/keys | Replace Identity Key


# **GetIdChainbyId**
> IdentityChain GetIdChainbyId(ctx, identityChainId)
Get Identity Chain Info

Retrieve the details about a particular Identity Chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identityChainId** | **string**| The hex encoded string that points to the identity&#39;s chain | 

### Return type

[**IdentityChain**](IdentityChain.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeybyEntryHash**
> IdentityKey GetKeybyEntryHash(ctx, identityChainId, keyString)
Get Key Info

Retreive information about a specific public key for a given Identity, including the heights at which the key was activated and retired if applicable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identityChainId** | **string**| The hex encoded string that points to the identity&#39;s chain | 
  **keyString** | **string**| The public key string. Must be in base58 idpub format. | 

### Return type

[**IdentityKey**](IdentityKey.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeysbyIdChainId**
> KeyList GetKeysbyIdChainId(ctx, identityChainId, optional)
Get Identity Chain's Keys

Returns all of the keys, past and present, associated with a particular Identity. Results are paginated and ordered by activation height.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identityChainId** | **string**| The hex encoded string that points to the identity&#39;s chain | 
 **optional** | ***GetKeysbyIdChainIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetKeysbyIdChainIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32**| The number of items you would like back in each page. | 
 **offset** | **int32**| The offset parameter allows you to select which item you would like to start from when you get back a list from Connect. For example, if you&#39;ve already seen the first 15 items and you&#39;d like the next set, you would send an offset of 15. &#x60;offset&#x3D;0&#x60; starts from the first item of the set and is the default position. | 

### Return type

[**KeyList**](KeyList.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostIdChain**
> ChainShort PostIdChain(ctx, identityCreate)
Create Identity Chain

Creates a new identity chain. To create the chain, you'll need to include a unique name array and an array of public keys. If successful, returns information about the chain that was created.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identityCreate** | [**IdentityCreate**](IdentityCreate.md)|  | 

### Return type

[**ChainShort**](ChainShort.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostKeytoIdChainId**
> EntryShort PostKeytoIdChainId(ctx, identityChainId, identityKeyReplace)
Replace Identity Key

Retires an existing key from an identity and replaces it with a new one. To do this, a user must send the key to be replaced (`old_key`), the `new_key`, the signing key that authorizes the replacement and a signed message from the signing key. The signing key must be either equal to or senior to the key that is being replaced.  *Note: You may not reuse a key. If the `new_key` has been used by this Identity at any level, the replacement will fail.*

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **identityChainId** | **string**| The hex encoded string that points to the identity&#39;s chain | 
  **identityKeyReplace** | [**IdentityKeyReplace**](IdentityKeyReplace.md)|  | 

### Return type

[**EntryShort**](EntryShort.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

