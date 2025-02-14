# RegisterUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | immutable name of the user, used to find or mention its actor | 
**Password** | **string** |  | 
**Email** | **string** | email of the user, used for login or service communications | 
**DisplayName** | Pointer to **string** | editable name of the user, displayed in its representations | [optional] 
**Channel** | Pointer to [**RegisterUserChannel**](RegisterUserChannel.md) |  | [optional] 

## Methods

### NewRegisterUser

`func NewRegisterUser(username string, password string, email string, ) *RegisterUser`

NewRegisterUser instantiates a new RegisterUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserWithDefaults

`func NewRegisterUserWithDefaults() *RegisterUser`

NewRegisterUserWithDefaults instantiates a new RegisterUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *RegisterUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RegisterUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RegisterUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *RegisterUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RegisterUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RegisterUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *RegisterUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RegisterUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RegisterUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *RegisterUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *RegisterUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *RegisterUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *RegisterUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetChannel

`func (o *RegisterUser) GetChannel() RegisterUserChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *RegisterUser) GetChannelOk() (*RegisterUserChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *RegisterUser) SetChannel(v RegisterUserChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *RegisterUser) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


