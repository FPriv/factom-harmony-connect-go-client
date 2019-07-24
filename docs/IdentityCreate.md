# IdentityCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | **[]string** | A unique array of strings that together constitute the Identity&#39;s name. Each string should be in Base64 format.   *Note: It is best to avoid adding personally identifiable information to the blockchain.* | 
**Keys** | **[]string** | A list of public keys that will be used to verify this Indenty’s signatures. You may initiate an identity with as many keys as you like. The array of keys should be sent in order of priority with 0 being the master key. The keys should be sent as Base58 strings in IdPub format. | 
**CallbackUrl** | **string** | The URL where you would like to receive the callback from Connect. If this is not specified, callbacks will not activate. | [optional] 
**CallbackStages** | **[]string** | The stages that you would like to trigger a callback from Connect. This list can include any or all of the three stages: &#x60;replicated&#x60;, &#x60;factom&#x60;, and &#x60;anchored&#x60;. If callbacks are activated and this field is not sent, it will default to &#x60;factom&#x60; and &#x60;anchored&#x60; | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


