# VerifyUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VerificationString** | **string** |  | 
**IsPendingEmail** | Pointer to **bool** |  | [optional] 

## Methods

### NewVerifyUserRequest

`func NewVerifyUserRequest(verificationString string, ) *VerifyUserRequest`

NewVerifyUserRequest instantiates a new VerifyUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyUserRequestWithDefaults

`func NewVerifyUserRequestWithDefaults() *VerifyUserRequest`

NewVerifyUserRequestWithDefaults instantiates a new VerifyUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerificationString

`func (o *VerifyUserRequest) GetVerificationString() string`

GetVerificationString returns the VerificationString field if non-nil, zero value otherwise.

### GetVerificationStringOk

`func (o *VerifyUserRequest) GetVerificationStringOk() (*string, bool)`

GetVerificationStringOk returns a tuple with the VerificationString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationString

`func (o *VerifyUserRequest) SetVerificationString(v string)`

SetVerificationString sets VerificationString field to given value.


### GetIsPendingEmail

`func (o *VerifyUserRequest) GetIsPendingEmail() bool`

GetIsPendingEmail returns the IsPendingEmail field if non-nil, zero value otherwise.

### GetIsPendingEmailOk

`func (o *VerifyUserRequest) GetIsPendingEmailOk() (*bool, bool)`

GetIsPendingEmailOk returns a tuple with the IsPendingEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPendingEmail

`func (o *VerifyUserRequest) SetIsPendingEmail(v bool)`

SetIsPendingEmail sets IsPendingEmail field to given value.

### HasIsPendingEmail

`func (o *VerifyUserRequest) HasIsPendingEmail() bool`

HasIsPendingEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


