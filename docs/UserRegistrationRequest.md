# UserRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | immutable name of the user, used to find or mention its actor | 
**Password** | **string** |  | 
**Email** | **string** | email of the user, used for login or service communications | 
**DisplayName** | Pointer to **string** | editable name of the user, displayed in its representations | [optional] 
**Channel** | Pointer to [**RegisterUserChannel**](RegisterUserChannel.md) |  | [optional] 
**RegistrationReason** | **string** | reason for the user to register on the instance | 

## Methods

### NewUserRegistrationRequest

`func NewUserRegistrationRequest(username string, password string, email string, registrationReason string, ) *UserRegistrationRequest`

NewUserRegistrationRequest instantiates a new UserRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRegistrationRequestWithDefaults

`func NewUserRegistrationRequestWithDefaults() *UserRegistrationRequest`

NewUserRegistrationRequestWithDefaults instantiates a new UserRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserRegistrationRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserRegistrationRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserRegistrationRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *UserRegistrationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserRegistrationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserRegistrationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *UserRegistrationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRegistrationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRegistrationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *UserRegistrationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserRegistrationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserRegistrationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserRegistrationRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetChannel

`func (o *UserRegistrationRequest) GetChannel() RegisterUserChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UserRegistrationRequest) GetChannelOk() (*RegisterUserChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UserRegistrationRequest) SetChannel(v RegisterUserChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UserRegistrationRequest) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetRegistrationReason

`func (o *UserRegistrationRequest) GetRegistrationReason() string`

GetRegistrationReason returns the RegistrationReason field if non-nil, zero value otherwise.

### GetRegistrationReasonOk

`func (o *UserRegistrationRequest) GetRegistrationReasonOk() (*string, bool)`

GetRegistrationReasonOk returns a tuple with the RegistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationReason

`func (o *UserRegistrationRequest) SetRegistrationReason(v string)`

SetRegistrationReason sets RegistrationReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


