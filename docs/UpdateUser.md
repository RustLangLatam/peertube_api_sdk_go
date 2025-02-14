# UpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The updated email of the user | [optional] 
**EmailVerified** | Pointer to **bool** | Set the email as verified | [optional] 
**VideoQuota** | Pointer to **int32** | The updated video quota of the user in bytes | [optional] 
**VideoQuotaDaily** | Pointer to **int32** | The updated daily video quota of the user in bytes | [optional] 
**PluginAuth** | Pointer to **NullableString** | The auth plugin to use to authenticate the user | [optional] 
**Role** | Pointer to [**UserRole**](UserRole.md) |  | [optional] 
**AdminFlags** | Pointer to [**UserAdminFlags**](UserAdminFlags.md) |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateUser

`func NewUpdateUser() *UpdateUser`

NewUpdateUser instantiates a new UpdateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserWithDefaults

`func NewUpdateUserWithDefaults() *UpdateUser`

NewUpdateUserWithDefaults instantiates a new UpdateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UpdateUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEmailVerified

`func (o *UpdateUser) GetEmailVerified() bool`

GetEmailVerified returns the EmailVerified field if non-nil, zero value otherwise.

### GetEmailVerifiedOk

`func (o *UpdateUser) GetEmailVerifiedOk() (*bool, bool)`

GetEmailVerifiedOk returns a tuple with the EmailVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailVerified

`func (o *UpdateUser) SetEmailVerified(v bool)`

SetEmailVerified sets EmailVerified field to given value.

### HasEmailVerified

`func (o *UpdateUser) HasEmailVerified() bool`

HasEmailVerified returns a boolean if a field has been set.

### GetVideoQuota

`func (o *UpdateUser) GetVideoQuota() int32`

GetVideoQuota returns the VideoQuota field if non-nil, zero value otherwise.

### GetVideoQuotaOk

`func (o *UpdateUser) GetVideoQuotaOk() (*int32, bool)`

GetVideoQuotaOk returns a tuple with the VideoQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuota

`func (o *UpdateUser) SetVideoQuota(v int32)`

SetVideoQuota sets VideoQuota field to given value.

### HasVideoQuota

`func (o *UpdateUser) HasVideoQuota() bool`

HasVideoQuota returns a boolean if a field has been set.

### GetVideoQuotaDaily

`func (o *UpdateUser) GetVideoQuotaDaily() int32`

GetVideoQuotaDaily returns the VideoQuotaDaily field if non-nil, zero value otherwise.

### GetVideoQuotaDailyOk

`func (o *UpdateUser) GetVideoQuotaDailyOk() (*int32, bool)`

GetVideoQuotaDailyOk returns a tuple with the VideoQuotaDaily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoQuotaDaily

`func (o *UpdateUser) SetVideoQuotaDaily(v int32)`

SetVideoQuotaDaily sets VideoQuotaDaily field to given value.

### HasVideoQuotaDaily

`func (o *UpdateUser) HasVideoQuotaDaily() bool`

HasVideoQuotaDaily returns a boolean if a field has been set.

### GetPluginAuth

`func (o *UpdateUser) GetPluginAuth() string`

GetPluginAuth returns the PluginAuth field if non-nil, zero value otherwise.

### GetPluginAuthOk

`func (o *UpdateUser) GetPluginAuthOk() (*string, bool)`

GetPluginAuthOk returns a tuple with the PluginAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginAuth

`func (o *UpdateUser) SetPluginAuth(v string)`

SetPluginAuth sets PluginAuth field to given value.

### HasPluginAuth

`func (o *UpdateUser) HasPluginAuth() bool`

HasPluginAuth returns a boolean if a field has been set.

### SetPluginAuthNil

`func (o *UpdateUser) SetPluginAuthNil(b bool)`

 SetPluginAuthNil sets the value for PluginAuth to be an explicit nil

### UnsetPluginAuth
`func (o *UpdateUser) UnsetPluginAuth()`

UnsetPluginAuth ensures that no value is present for PluginAuth, not even an explicit nil
### GetRole

`func (o *UpdateUser) GetRole() UserRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdateUser) GetRoleOk() (*UserRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdateUser) SetRole(v UserRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdateUser) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetAdminFlags

`func (o *UpdateUser) GetAdminFlags() UserAdminFlags`

GetAdminFlags returns the AdminFlags field if non-nil, zero value otherwise.

### GetAdminFlagsOk

`func (o *UpdateUser) GetAdminFlagsOk() (*UserAdminFlags, bool)`

GetAdminFlagsOk returns a tuple with the AdminFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminFlags

`func (o *UpdateUser) SetAdminFlags(v UserAdminFlags)`

SetAdminFlags sets AdminFlags field to given value.

### HasAdminFlags

`func (o *UpdateUser) HasAdminFlags() bool`

HasAdminFlags returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


