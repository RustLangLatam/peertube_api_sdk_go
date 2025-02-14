# ServerConfigSignup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | Pointer to **bool** |  | [optional] 
**AllowedForCurrentIP** | Pointer to **bool** |  | [optional] 
**RequiresEmailVerification** | Pointer to **bool** |  | [optional] 

## Methods

### NewServerConfigSignup

`func NewServerConfigSignup() *ServerConfigSignup`

NewServerConfigSignup instantiates a new ServerConfigSignup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigSignupWithDefaults

`func NewServerConfigSignupWithDefaults() *ServerConfigSignup`

NewServerConfigSignupWithDefaults instantiates a new ServerConfigSignup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *ServerConfigSignup) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ServerConfigSignup) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ServerConfigSignup) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *ServerConfigSignup) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.

### GetAllowedForCurrentIP

`func (o *ServerConfigSignup) GetAllowedForCurrentIP() bool`

GetAllowedForCurrentIP returns the AllowedForCurrentIP field if non-nil, zero value otherwise.

### GetAllowedForCurrentIPOk

`func (o *ServerConfigSignup) GetAllowedForCurrentIPOk() (*bool, bool)`

GetAllowedForCurrentIPOk returns a tuple with the AllowedForCurrentIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedForCurrentIP

`func (o *ServerConfigSignup) SetAllowedForCurrentIP(v bool)`

SetAllowedForCurrentIP sets AllowedForCurrentIP field to given value.

### HasAllowedForCurrentIP

`func (o *ServerConfigSignup) HasAllowedForCurrentIP() bool`

HasAllowedForCurrentIP returns a boolean if a field has been set.

### GetRequiresEmailVerification

`func (o *ServerConfigSignup) GetRequiresEmailVerification() bool`

GetRequiresEmailVerification returns the RequiresEmailVerification field if non-nil, zero value otherwise.

### GetRequiresEmailVerificationOk

`func (o *ServerConfigSignup) GetRequiresEmailVerificationOk() (*bool, bool)`

GetRequiresEmailVerificationOk returns a tuple with the RequiresEmailVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresEmailVerification

`func (o *ServerConfigSignup) SetRequiresEmailVerification(v bool)`

SetRequiresEmailVerification sets RequiresEmailVerification field to given value.

### HasRequiresEmailVerification

`func (o *ServerConfigSignup) HasRequiresEmailVerification() bool`

HasRequiresEmailVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


