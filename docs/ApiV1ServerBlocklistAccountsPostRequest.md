# ApiV1ServerBlocklistAccountsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | account to block, in the form &#x60;username@domain&#x60; | 

## Methods

### NewApiV1ServerBlocklistAccountsPostRequest

`func NewApiV1ServerBlocklistAccountsPostRequest(accountName string, ) *ApiV1ServerBlocklistAccountsPostRequest`

NewApiV1ServerBlocklistAccountsPostRequest instantiates a new ApiV1ServerBlocklistAccountsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1ServerBlocklistAccountsPostRequestWithDefaults

`func NewApiV1ServerBlocklistAccountsPostRequestWithDefaults() *ApiV1ServerBlocklistAccountsPostRequest`

NewApiV1ServerBlocklistAccountsPostRequestWithDefaults instantiates a new ApiV1ServerBlocklistAccountsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *ApiV1ServerBlocklistAccountsPostRequest) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ApiV1ServerBlocklistAccountsPostRequest) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ApiV1ServerBlocklistAccountsPostRequest) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


