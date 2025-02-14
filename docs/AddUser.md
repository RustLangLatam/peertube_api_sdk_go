# AddUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | immutable name of the user, used to find or mention its actor | 
**Password** | **string** |  | 
**Email** | **string** | The user email | 
**VideoQuota** | Pointer to **int32** | The user video quota in bytes | [optional] 
**VideoQuotaDaily** | Pointer to **int32** | The user daily video quota in bytes | [optional] 
**ChannelName** | Pointer to **string** | immutable name of the channel, used to interact with its actor | [optional] 
**Role** | [**UserRole**](UserRole.md) |  | 
**AdminFlags** | Pointer to [**UserAdminFlags**](UserAdminFlags.md) |  | [optional] 

## Methods

### NewAddUser

`func NewAddUser(username string, password string, email string, role UserRole, ) *AddUser`

NewAddUser instantiates a new AddUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserWithDefaults

`func NewAddUserWithDefaults() *AddUser`

NewAddUserWithDefaults instantiates a new AddUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *AddUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *AddUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *AddUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetVideoQuota

`func (o *AddUser) GetVideoQuota() int32`

GetVideoQuota returns the VideoQuota field if non-nil, zero value otherwise.

### GetVideoQuotaOk

`func (o *AddUser) GetVideoQuotaOk() (*int32, bool)`

GetVideoQuotaOk returns a tuple with the VideoQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuota

`func (o *AddUser) SetVideoQuota(v int32)`

SetVideoQuota sets VideoQuota field to given value.

### HasVideoQuota

`func (o *AddUser) HasVideoQuota() bool`

HasVideoQuota returns a boolean if a field has been set.

### GetVideoQuotaDaily

`func (o *AddUser) GetVideoQuotaDaily() int32`

GetVideoQuotaDaily returns the VideoQuotaDaily field if non-nil, zero value otherwise.

### GetVideoQuotaDailyOk

`func (o *AddUser) GetVideoQuotaDailyOk() (*int32, bool)`

GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuotaDaily

`func (o *AddUser) SetVideoQuotaDaily(v int32)`

SetVideoQuotaDaily sets VideoQuotaDaily field to given value.

### HasVideoQuotaDaily

`func (o *AddUser) HasVideoQuotaDaily() bool`

HasVideoQuotaDaily returns a boolean if a field has been set.

### GetChannelName

`func (o *AddUser) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *AddUser) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *AddUser) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *AddUser) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetRole

`func (o *AddUser) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddUser) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddUser) SetRole(v UserRole)`

SetRole sets Role field to given value.


### GetAdminFlags

`func (o *AddUser) GetAdminFlags() UserAdminFlags`

GetAdminFlags returns the AdminFlags field if non-nil, zero value otherwise.

### GetAdminFlagsOk

`func (o *AddUser) GetAdminFlagsOk() (*UserAdminFlags, bool)`

GetAdminFlagsOk returns a tuple with the AdminFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFlags

`func (o *AddUser) SetAdminFlags(v UserAdminFlags)`

SetAdminFlags sets AdminFlags field to given value.

### HasAdminFlags

`func (o *AddUser) HasAdminFlags() bool`

HasAdminFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


