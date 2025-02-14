# ServerConfigCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**ServerConfigCustomInstance**](ServerConfigCustomInstance.md) |  | [optional] 
**Theme** | Pointer to [**ServerConfigCustomTheme**](ServerConfigCustomTheme.md) |  | [optional] 
**Services** | Pointer to [**ServerConfigCustomServices**](ServerConfigCustomServices.md) |  | [optional] 
**Cache** | Pointer to [**ServerConfigCustomCache**](ServerConfigCustomCache.md) |  | [optional] 
**Signup** | Pointer to [**ServerConfigCustomSignup**](ServerConfigCustomSignup.md) |  | [optional] 
**Admin** | Pointer to [**ServerConfigCustomAdmin**](ServerConfigCustomAdmin.md) |  | [optional] 
**ContactForm** | Pointer to [**ServerConfigEmail**](ServerConfigEmail.md) |  | [optional] 
**User** | Pointer to [**ServerConfigCustomUser**](ServerConfigCustomUser.md) |  | [optional] 
**Transcoding** | Pointer to [**ServerConfigCustomTranscoding**](ServerConfigCustomTranscoding.md) |  | [optional] 
**Import** | Pointer to [**ServerConfigCustomImport**](ServerConfigCustomImport.md) |  | [optional] 
**AutoBlacklist** | Pointer to [**ServerConfigAutoBlacklist**](ServerConfigAutoBlacklist.md) |  | [optional] 
**Followers** | Pointer to [**ServerConfigCustomFollowers**](ServerConfigCustomFollowers.md) |  | [optional] 

## Methods

### NewServerConfigCustom

`func NewServerConfigCustom() *ServerConfigCustom`

NewServerConfigCustom instantiates a new ServerConfigCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigCustomWithDefaults

`func NewServerConfigCustomWithDefaults() *ServerConfigCustom`

NewServerConfigCustomWithDefaults instantiates a new ServerConfigCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *ServerConfigCustom) GetInstance() ServerConfigCustomInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ServerConfigCustom) GetInstanceOk() (*ServerConfigCustomInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ServerConfigCustom) SetInstance(v ServerConfigCustomInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ServerConfigCustom) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetTheme

`func (o *ServerConfigCustom) GetTheme() ServerConfigCustomTheme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *ServerConfigCustom) GetThemeOk() (*ServerConfigCustomTheme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *ServerConfigCustom) SetTheme(v ServerConfigCustomTheme)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *ServerConfigCustom) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetServices

`func (o *ServerConfigCustom) GetServices() ServerConfigCustomServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ServerConfigCustom) GetServicesOk() (*ServerConfigCustomServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ServerConfigCustom) SetServices(v ServerConfigCustomServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *ServerConfigCustom) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetCache

`func (o *ServerConfigCustom) GetCache() ServerConfigCustomCache`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *ServerConfigCustom) GetCacheOk() (*ServerConfigCustomCache, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *ServerConfigCustom) SetCache(v ServerConfigCustomCache)`

SetCache sets Cache field to given value.

### HasCache

`func (o *ServerConfigCustom) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetSignup

`func (o *ServerConfigCustom) GetSignup() ServerConfigCustomSignup`

GetSignup returns the Signup field if non-nil, zero value otherwise.

### GetSignupOk

`func (o *ServerConfigCustom) GetSignupOk() (*ServerConfigCustomSignup, bool)`

GetSignupOk returns a tuple with the Signup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignup

`func (o *ServerConfigCustom) SetSignup(v ServerConfigCustomSignup)`

SetSignup sets Signup field to given value.

### HasSignup

`func (o *ServerConfigCustom) HasSignup() bool`

HasSignup returns a boolean if a field has been set.

### GetAdmin

`func (o *ServerConfigCustom) GetAdmin() ServerConfigCustomAdmin`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *ServerConfigCustom) GetAdminOk() (*ServerConfigCustomAdmin, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *ServerConfigCustom) SetAdmin(v ServerConfigCustomAdmin)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *ServerConfigCustom) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetContactForm

`func (o *ServerConfigCustom) GetContactForm() ServerConfigEmail`

GetContactForm returns the ContactForm field if non-nil, zero value otherwise.

### GetContactFormOk

`func (o *ServerConfigCustom) GetContactFormOk() (*ServerConfigEmail, bool)`

GetContactFormOk returns a tuple with the ContactForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactForm

`func (o *ServerConfigCustom) SetContactForm(v ServerConfigEmail)`

SetContactForm sets ContactForm field to given value.

### HasContactForm

`func (o *ServerConfigCustom) HasContactForm() bool`

HasContactForm returns a boolean if a field has been set.

### GetUser

`func (o *ServerConfigCustom) GetUser() ServerConfigCustomUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ServerConfigCustom) GetUserOk() (*ServerConfigCustomUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ServerConfigCustom) SetUser(v ServerConfigCustomUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ServerConfigCustom) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetTranscoding

`func (o *ServerConfigCustom) GetTranscoding() ServerConfigCustomTranscoding`

GetTranscoding returns the Transcoding field if non-nil, zero value otherwise.

### GetTranscodingOk

`func (o *ServerConfigCustom) GetTranscodingOk() (*ServerConfigCustomTranscoding, bool)`

GetTranscodingOk returns a tuple with the Transcoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscoding

`func (o *ServerConfigCustom) SetTranscoding(v ServerConfigCustomTranscoding)`

SetTranscoding sets Transcoding field to given value.

### HasTranscoding

`func (o *ServerConfigCustom) HasTranscoding() bool`

HasTranscoding returns a boolean if a field has been set.

### GetImport

`func (o *ServerConfigCustom) GetImport() ServerConfigCustomImport`

GetImport returns the Import field if non-nil, zero value otherwise.

### GetImportOk

`func (o *ServerConfigCustom) GetImportOk() (*ServerConfigCustomImport, bool)`

GetImportOk returns a tuple with the Import field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImport

`func (o *ServerConfigCustom) SetImport(v ServerConfigCustomImport)`

SetImport sets Import field to given value.

### HasImport

`func (o *ServerConfigCustom) HasImport() bool`

HasImport returns a boolean if a field has been set.

### GetAutoBlacklist

`func (o *ServerConfigCustom) GetAutoBlacklist() ServerConfigAutoBlacklist`

GetAutoBlacklist returns the AutoBlacklist field if non-nil, zero value otherwise.

### GetAutoBlacklistOk

`func (o *ServerConfigCustom) GetAutoBlacklistOk() (*ServerConfigAutoBlacklist, bool)`

GetAutoBlacklistOk returns a tuple with the AutoBlacklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBlacklist

`func (o *ServerConfigCustom) SetAutoBlacklist(v ServerConfigAutoBlacklist)`

SetAutoBlacklist sets AutoBlacklist field to given value.

### HasAutoBlacklist

`func (o *ServerConfigCustom) HasAutoBlacklist() bool`

HasAutoBlacklist returns a boolean if a field has been set.

### GetFollowers

`func (o *ServerConfigCustom) GetFollowers() ServerConfigCustomFollowers`

GetFollowers returns the Followers field if non-nil, zero value otherwise.

### GetFollowersOk

`func (o *ServerConfigCustom) GetFollowersOk() (*ServerConfigCustomFollowers, bool)`

GetFollowersOk returns a tuple with the Followers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowers

`func (o *ServerConfigCustom) SetFollowers(v ServerConfigCustomFollowers)`

SetFollowers sets Followers field to given value.

### HasFollowers

`func (o *ServerConfigCustom) HasFollowers() bool`

HasFollowers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


