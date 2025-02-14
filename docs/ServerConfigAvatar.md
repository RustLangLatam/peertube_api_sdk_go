# ServerConfigAvatar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | Pointer to [**ServerConfigAvatarFile**](ServerConfigAvatarFile.md) |  | [optional] 
**Extensions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServerConfigAvatar

`func NewServerConfigAvatar() *ServerConfigAvatar`

NewServerConfigAvatar instantiates a new ServerConfigAvatar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigAvatarWithDefaults

`func NewServerConfigAvatarWithDefaults() *ServerConfigAvatar`

NewServerConfigAvatarWithDefaults instantiates a new ServerConfigAvatar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *ServerConfigAvatar) GetFile() ServerConfigAvatarFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ServerConfigAvatar) GetFileOk() (*ServerConfigAvatarFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ServerConfigAvatar) SetFile(v ServerConfigAvatarFile)`

SetFile sets File field to given value.

### HasFile

`func (o *ServerConfigAvatar) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetExtensions

`func (o *ServerConfigAvatar) GetExtensions() []string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ServerConfigAvatar) GetExtensionsOk() (*[]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ServerConfigAvatar) SetExtensions(v []string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ServerConfigAvatar) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


