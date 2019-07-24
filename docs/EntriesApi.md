# \EntriesApi

All URIs are relative to *https://ephemeral.api.factom.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntriesByChainID**](EntriesApi.md#GetEntriesByChainID) | **Get** /chains/{chain_id}/entries | Get Chain&#39;s Entries
[**GetEntryByHash**](EntriesApi.md#GetEntryByHash) | **Get** /chains/{chain_id}/entries/{entry_hash} | Get Entry Info
[**GetFirstEntry**](EntriesApi.md#GetFirstEntry) | **Get** /chains/{chain_id}/entries/first | Get Chain&#39;s First Entry
[**GetLastEntry**](EntriesApi.md#GetLastEntry) | **Get** /chains/{chain_id}/entries/last | Get Chain&#39;s Last Entry
[**PostEntriesSearch**](EntriesApi.md#PostEntriesSearch) | **Post** /chains/{chain_id}/entries/search | Search Chain&#39;s Entries
[**PostEntryToChainID**](EntriesApi.md#PostEntryToChainID) | **Post** /chains/{chain_id}/entries | Create an Entry


# **GetEntriesByChainID**
> EntryList GetEntriesByChainID(ctx, chainId, optional)
Get Chain's Entries

List all entries contained on the specified chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chainId** | **string**| Chain identifier | 
 **optional** | ***GetEntriesByChainIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetEntriesByChainIDOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32**| The number of items you would like back in each page. | 
 **offset** | **int32**| The offset parameter allows you to select which item you would like to start from when you get back a list from Connect. For example, if you&#39;ve already seen the first 15 items and you&#39;d like the next set, you would send an offset of 15. &#x60;offset&#x3D;0&#x60; starts from the first item of the set and is the default position. | 
 **stages** | **string**| The immutability stages you want to restrict results to. You can choose any from &#x60;replicated&#x60;, &#x60;factom&#x60;, and &#x60;anchored&#x60;. If you would like to search among multiple stages, send them in a comma separated string. For example: &#x60;&#39;replicated,factom&#39;&#x60;. | 

### Return type

[**EntryList**](EntryList.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEntryByHash**
> Entry GetEntryByHash(ctx, chainId, entryHash)
Get Entry Info

Returns information about a specific entry on Connect. The requested entry must be specified using the Chain ID and Entry Hash.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chainId** | **string**| Chain identifier | 
  **entryHash** | **string**| The unique identitfier of the entry. | 

### Return type

[**Entry**](Entry.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFirstEntry**
> Entry GetFirstEntry(ctx, chainId)
Get Chain's First Entry

Retrieve the first entry that has been saved to this chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chainId** | **string**| Chain identifier | 

### Return type

[**Entry**](Entry.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLastEntry**
> Entry GetLastEntry(ctx, chainId)
Get Chain's Last Entry

Retrieve the last entry that has been saved to this chain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chainId** | **string**| Chain identifier | 

### Return type

[**Entry**](Entry.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEntriesSearch**
> EntrySearchResponse PostEntriesSearch(ctx, chainId, searchBody, optional)
Search Chain's Entries

Find all of the entries within the specified chain that have the requested `external_ids`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chainId** | **string**| Chain identifier | 
  **searchBody** | [**SearchBody**](SearchBody.md)|  | 
 **optional** | ***PostEntriesSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostEntriesSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32**| The number of items you would like back in each page. | 
 **offset** | **int32**| The offset parameter allows you to select which item you would like to start from when you get back a list from Connect. For example, if you&#39;ve already seen the first 15 items and you&#39;d like the next set, you would send an offset of 15. &#x60;offset&#x3D;0&#x60; starts from the first item of the set and is the default position. | 

### Return type

[**EntrySearchResponse**](EntrySearchResponse.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostEntryToChainID**
> EntryShort PostEntryToChainID(ctx, chainId, entryCreate)
Create an Entry

Create a new entry for the selected chain. Content and external id must be uploaded in Base64 format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chainId** | **string**| Chain identifier | 
  **entryCreate** | [**EntryCreate**](EntryCreate.md)|  | 

### Return type

[**EntryShort**](EntryShort.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

