# RegisterUserChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | immutable name of the channel, used to interact with its actor | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 

## Methods

### NewRegisterUserChannel

`func NewRegisterUserChannel() *RegisterUserChannel`

NewRegisterUserChannel instantiates a new RegisterUserChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserChannelWithDefaults

`func NewRegisterUserChannelWithDefaults() *RegisterUserChannel`

NewRegisterUserChannelWithDefaults instantiates a new RegisterUserChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RegisterUserChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisterUserChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisterUserChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisterUserChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *RegisterUserChannel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RegisterUserChannel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RegisterUserChannel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RegisterUserChannel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


