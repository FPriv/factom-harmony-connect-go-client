# \ProofsApi

All URIs are relative to *https://ephemeral.api.factom.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnchorbySearch**](ProofsApi.md#GetAnchorbySearch) | **Get** /anchors/{object_id} | Get Object&#39;s Anchors
[**GetReceiptsbyEntry**](ProofsApi.md#GetReceiptsbyEntry) | **Get** /receipts/{entry_hash} | Get Entry&#39;s Receipts
[**GetSearch**](ProofsApi.md#GetSearch) | **Get** /search | Search


# **GetAnchorbySearch**
> AnchorLong GetAnchorbySearch(ctx, objectId)
Get Object's Anchors

Retreive the blockchain anchors of an entry, chain, or block. Returns an array of anchors that may be of type ethereum, bitcoin, or factom. The valid identifiers for the objects that can be anchored are as follows:  * Entry - Entry Hash * Chain - Chain ID * Directory Block - Height or Key Merkle Root * Entry Block - Key Merkle Root * Factoid Block - Key Merkle Root

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| Object identifier.  Used to identify an entry, chain, or block that has been published on the Factom blockchain. These objects can be identified by their hash, ID, height, or key Merkle root. | 

### Return type

[**AnchorLong**](AnchorLong.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReceiptsbyEntry**
> ReceiptLong GetReceiptsbyEntry(ctx, entryHash)
Get Entry's Receipts

Retrieve a receipt providing cryptographically verifiable proof that information was recorded in the Factom blockchain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **entryHash** | **string**| The unique identitfier of the entry. | 

### Return type

[**ReceiptLong**](ReceiptLong.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSearch**
> InlineResponse200 GetSearch(ctx, optional)
Search

Search for something on the Factom blockchain. You may search for Directory Blocks (by height or key Merkle root), Entry Blocks (by key Merkle root), Factoid Blocks (by key Merkle root), Chains (by Chain ID), Entries (by Entry Hash), Addresses (by user address or address), and Transactions (by transaction ID).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSearchOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **string**| The term that you are searching for. You may search for Directory Blocks (by height or key Merkle root), Entry Blocks (by key Merkle root), Factoid Blocks (by key Merkle root), Chains (by Chain ID), Entries (by Entry Hash), Addresses (by user address or address), and Transactions (by transaction ID). | 
 **allowedStages** | **string**| When entered, filters chain and entry responses by immutability stage. Must be either &#x60;replicated&#x60;, &#x60;factom&#x60;, or &#x60;anchored&#x60;. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[AppId](../README.md#AppId), [AppKey](../README.md#AppKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

