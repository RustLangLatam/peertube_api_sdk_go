# RequestTwoFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | Pointer to **string** | Password of the currently authenticated user | [optional] 

## Methods

### NewRequestTwoFactorRequest

`func NewRequestTwoFactorRequest() *RequestTwoFactorRequest`

NewRequestTwoFactorRequest instantiates a new RequestTwoFactorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTwoFactorRequestWithDefaults

`func NewRequestTwoFactorRequestWithDefaults() *RequestTwoFactorRequest`

NewRequestTwoFactorRequestWithDefaults instantiates a new RequestTwoFactorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *RequestTwoFactorRequest) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *RequestTwoFactorRequest) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *RequestTwoFactorRequest) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *RequestTwoFactorRequest) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


