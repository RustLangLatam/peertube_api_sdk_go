# AccountSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Avatars** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 

## Methods

### NewAccountSummary

`func NewAccountSummary() *AccountSummary`

NewAccountSummary instantiates a new AccountSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSummaryWithDefaults

`func NewAccountSummaryWithDefaults() *AccountSummary`

NewAccountSummaryWithDefaults instantiates a new AccountSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountSummary) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountSummary) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountSummary) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AccountSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AccountSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *AccountSummary) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccountSummary) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccountSummary) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccountSummary) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetUrl

`func (o *AccountSummary) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AccountSummary) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AccountSummary) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AccountSummary) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetHost

`func (o *AccountSummary) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AccountSummary) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AccountSummary) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AccountSummary) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAvatars

`func (o *AccountSummary) GetAvatars() []ActorImage`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *AccountSummary) GetAvatarsOk() (*[]ActorImage, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *AccountSummary) SetAvatars(v []ActorImage)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *AccountSummary) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


