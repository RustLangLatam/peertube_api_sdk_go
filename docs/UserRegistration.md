# UserRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**State** | Pointer to [**UserRegistrationState**](UserRegistrationState.md) |  | [optional] 
**RegistrationReason** | Pointer to **string** |  | [optional] 
**ModerationResponse** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**EmailVerified** | Pointer to **bool** |  | [optional] 
**AccountDisplayName** | Pointer to **string** |  | [optional] 
**ChannelHandle** | Pointer to **string** |  | [optional] 
**ChannelDisplayName** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**NullableUserRegistrationUser**](UserRegistrationUser.md) |  | [optional] 

## Methods

### NewUserRegistration

`func NewUserRegistration() *UserRegistration`

NewUserRegistration instantiates a new UserRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRegistrationWithDefaults

`func NewUserRegistrationWithDefaults() *UserRegistration`

NewUserRegistrationWithDefaults instantiates a new UserRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserRegistration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRegistration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRegistration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserRegistration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *UserRegistration) GetState() UserRegistrationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserRegistration) GetStateOk() (*UserRegistrationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserRegistration) SetState(v UserRegistrationState)`

SetState sets State field to given value.

### HasState

`func (o *UserRegistration) HasState() bool`

HasState returns a boolean if a field has been set.

### GetRegistrationReason

`func (o *UserRegistration) GetRegistrationReason() string`

GetRegistrationReason returns the RegistrationReason field if non-nil, zero value otherwise.

### GetRegistrationReasonOk

`func (o *UserRegistration) GetRegistrationReasonOk() (*string, bool)`

GetRegistrationReasonOk returns a tuple with the RegistrationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationReason

`func (o *UserRegistration) SetRegistrationReason(v string)`

SetRegistrationReason sets RegistrationReason field to given value.

### HasRegistrationReason

`func (o *UserRegistration) HasRegistrationReason() bool`

HasRegistrationReason returns a boolean if a field has been set.

### GetModerationResponse

`func (o *UserRegistration) GetModerationResponse() string`

GetModerationResponse returns the ModerationResponse field if non-nil, zero value otherwise.

### GetModerationResponseOk

`func (o *UserRegistration) GetModerationResponseOk() (*string, bool)`

GetModerationResponseOk returns a tuple with the ModerationResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationResponse

`func (o *UserRegistration) SetModerationResponse(v string)`

SetModerationResponse sets ModerationResponse field to given value.

### HasModerationResponse

`func (o *UserRegistration) HasModerationResponse() bool`

HasModerationResponse returns a boolean if a field has been set.

### SetModerationResponseNil

`func (o *UserRegistration) SetModerationResponseNil(b bool)`

 SetModerationResponseNil sets the value for ModerationResponse to be an explicit nil

### UnsetModerationResponse
`func (o *UserRegistration) UnsetModerationResponse()`

UnsetModerationResponse ensures that no value is present for ModerationResponse, not even an explicit nil
### GetUsername

`func (o *UserRegistration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserRegistration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserRegistration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserRegistration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserRegistration) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserRegistration) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserRegistration) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserRegistration) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *UserRegistration) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UserRegistration) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UserRegistration) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *UserRegistration) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetAccountDisplayName

`func (o *UserRegistration) GetAccountDisplayName() string`

GetAccountDisplayName returns the AccountDisplayName field if non-nil, zero value otherwise.

### GetAccountDisplayNameOk

`func (o *UserRegistration) GetAccountDisplayNameOk() (*string, bool)`

GetAccountDisplayNameOk returns a tuple with the AccountDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountDisplayName

`func (o *UserRegistration) SetAccountDisplayName(v string)`

SetAccountDisplayName sets AccountDisplayName field to given value.

### HasAccountDisplayName

`func (o *UserRegistration) HasAccountDisplayName() bool`

HasAccountDisplayName returns a boolean if a field has been set.

### GetChannelHandle

`func (o *UserRegistration) GetChannelHandle() string`

GetChannelHandle returns the ChannelHandle field if non-nil, zero value otherwise.

### GetChannelHandleOk

`func (o *UserRegistration) GetChannelHandleOk() (*string, bool)`

GetChannelHandleOk returns a tuple with the ChannelHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelHandle

`func (o *UserRegistration) SetChannelHandle(v string)`

SetChannelHandle sets ChannelHandle field to given value.

### HasChannelHandle

`func (o *UserRegistration) HasChannelHandle() bool`

HasChannelHandle returns a boolean if a field has been set.

### GetChannelDisplayName

`func (o *UserRegistration) GetChannelDisplayName() string`

GetChannelDisplayName returns the ChannelDisplayName field if non-nil, zero value otherwise.

### GetChannelDisplayNameOk

`func (o *UserRegistration) GetChannelDisplayNameOk() (*string, bool)`

GetChannelDisplayNameOk returns a tuple with the ChannelDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDisplayName

`func (o *UserRegistration) SetChannelDisplayName(v string)`

SetChannelDisplayName sets ChannelDisplayName field to given value.

### HasChannelDisplayName

`func (o *UserRegistration) HasChannelDisplayName() bool`

HasChannelDisplayName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserRegistration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserRegistration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserRegistration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserRegistration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UserRegistration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserRegistration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserRegistration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UserRegistration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *UserRegistration) GetUser() UserRegistrationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserRegistration) GetUserOk() (*UserRegistrationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserRegistration) SetUser(v UserRegistrationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UserRegistration) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *UserRegistration) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *UserRegistration) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


