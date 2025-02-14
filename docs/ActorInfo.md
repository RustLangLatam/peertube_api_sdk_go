# ActorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Avatars** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 

## Methods

### NewActorInfo

`func NewActorInfo() *ActorInfo`

NewActorInfo instantiates a new ActorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorInfoWithDefaults

`func NewActorInfoWithDefaults() *ActorInfo`

NewActorInfoWithDefaults instantiates a new ActorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActorInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActorInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActorInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActorInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActorInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActorInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ActorInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ActorInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ActorInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ActorInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHost

`func (o *ActorInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ActorInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ActorInfo) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ActorInfo) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAvatars

`func (o *ActorInfo) GetAvatars() []ActorImage`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *ActorInfo) GetAvatarsOk() (*[]ActorImage, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *ActorInfo) SetAvatars(v []ActorImage)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *ActorInfo) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


