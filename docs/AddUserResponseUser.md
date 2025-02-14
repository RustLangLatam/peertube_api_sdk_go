# AddUserResponseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**ApiV1AbusesPost200ResponseAbuse**](ApiV1AbusesPost200ResponseAbuse.md) |  | [optional] 

## Methods

### NewAddUserResponseUser

`func NewAddUserResponseUser() *AddUserResponseUser`

NewAddUserResponseUser instantiates a new AddUserResponseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserResponseUserWithDefaults

`func NewAddUserResponseUserWithDefaults() *AddUserResponseUser`

NewAddUserResponseUserWithDefaults instantiates a new AddUserResponseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddUserResponseUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddUserResponseUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddUserResponseUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AddUserResponseUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *AddUserResponseUser) GetAccount() ApiV1AbusesPost200ResponseAbuse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddUserResponseUser) GetAccountOk() (*ApiV1AbusesPost200ResponseAbuse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddUserResponseUser) SetAccount(v ApiV1AbusesPost200ResponseAbuse)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddUserResponseUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


