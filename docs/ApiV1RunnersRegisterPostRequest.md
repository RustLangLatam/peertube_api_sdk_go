# ApiV1RunnersRegisterPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegistrationToken** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewApiV1RunnersRegisterPostRequest

`func NewApiV1RunnersRegisterPostRequest(registrationToken string, name string, ) *ApiV1RunnersRegisterPostRequest`

NewApiV1RunnersRegisterPostRequest instantiates a new ApiV1RunnersRegisterPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiV1RunnersRegisterPostRequestWithDefaults

`func NewApiV1RunnersRegisterPostRequestWithDefaults() *ApiV1RunnersRegisterPostRequest`

NewApiV1RunnersRegisterPostRequestWithDefaults instantiates a new ApiV1RunnersRegisterPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegistrationToken

`func (o *ApiV1RunnersRegisterPostRequest) GetRegistrationToken() string`

GetRegistrationToken returns the RegistrationToken field if non-nil, zero value otherwise.

### GetRegistrationTokenOk

`func (o *ApiV1RunnersRegisterPostRequest) GetRegistrationTokenOk() (*string, bool)`

GetRegistrationTokenOk returns a tuple with the RegistrationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationToken

`func (o *ApiV1RunnersRegisterPostRequest) SetRegistrationToken(v string)`

SetRegistrationToken sets RegistrationToken field to given value.


### GetName

`func (o *ApiV1RunnersRegisterPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiV1RunnersRegisterPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiV1RunnersRegisterPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApiV1RunnersRegisterPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiV1RunnersRegisterPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiV1RunnersRegisterPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiV1RunnersRegisterPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


