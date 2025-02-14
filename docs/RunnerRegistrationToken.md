# RunnerRegistrationToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RegistrationToken** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**RegisteredRunnersCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewRunnerRegistrationToken

`func NewRunnerRegistrationToken() *RunnerRegistrationToken`

NewRunnerRegistrationToken instantiates a new RunnerRegistrationToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerRegistrationTokenWithDefaults

`func NewRunnerRegistrationTokenWithDefaults() *RunnerRegistrationToken`

NewRunnerRegistrationTokenWithDefaults instantiates a new RunnerRegistrationToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunnerRegistrationToken) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunnerRegistrationToken) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunnerRegistrationToken) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RunnerRegistrationToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRegistrationToken

`func (o *RunnerRegistrationToken) GetRegistrationToken() string`

GetRegistrationToken returns the RegistrationToken field if non-nil, zero value otherwise.

### GetRegistrationTokenOk

`func (o *RunnerRegistrationToken) GetRegistrationTokenOk() (*string, bool)`

GetRegistrationTokenOk returns a tuple with the RegistrationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationToken

`func (o *RunnerRegistrationToken) SetRegistrationToken(v string)`

SetRegistrationToken sets RegistrationToken field to given value.

### HasRegistrationToken

`func (o *RunnerRegistrationToken) HasRegistrationToken() bool`

HasRegistrationToken returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RunnerRegistrationToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RunnerRegistrationToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RunnerRegistrationToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RunnerRegistrationToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RunnerRegistrationToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RunnerRegistrationToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RunnerRegistrationToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RunnerRegistrationToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRegisteredRunnersCount

`func (o *RunnerRegistrationToken) GetRegisteredRunnersCount() int32`

GetRegisteredRunnersCount returns the RegisteredRunnersCount field if non-nil, zero value otherwise.

### GetRegisteredRunnersCountOk

`func (o *RunnerRegistrationToken) GetRegisteredRunnersCountOk() (*int32, bool)`

GetRegisteredRunnersCountOk returns a tuple with the RegisteredRunnersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredRunnersCount

`func (o *RunnerRegistrationToken) SetRegisteredRunnersCount(v int32)`

SetRegisteredRunnersCount sets RegisteredRunnersCount field to given value.

### HasRegisteredRunnersCount

`func (o *RunnerRegistrationToken) HasRegisteredRunnersCount() bool`

HasRegisteredRunnersCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


