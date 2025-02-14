# VideoChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | immutable name of the actor, used to find or mention it | [optional] 
**Avatars** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 
**Host** | Pointer to **string** | server on which the actor is resident | [optional] 
**HostRedundancyAllowed** | Pointer to **bool** | whether this actor&#39;s host allows redundancy of its videos | [optional] 
**FollowingCount** | Pointer to **int32** | number of actors subscribed to by this actor, as seen by this instance | [optional] 
**FollowersCount** | Pointer to **int32** | number of followers of this actor, as seen by this instance | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**DisplayName** | Pointer to **string** | editable name of the channel, displayed in its representations | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Support** | Pointer to **NullableString** | text shown by default on all videos of this channel, to tell the audience how to support it | [optional] 
**IsLocal** | Pointer to **bool** |  | [optional] [readonly] 
**Banners** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 
**OwnerAccount** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewVideoChannel

`func NewVideoChannel() *VideoChannel`

NewVideoChannel instantiates a new VideoChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoChannelWithDefaults

`func NewVideoChannelWithDefaults() *VideoChannel`

NewVideoChannelWithDefaults instantiates a new VideoChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoChannel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoChannel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoChannel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *VideoChannel) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoChannel) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoChannel) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoChannel) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *VideoChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoChannel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoChannel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatars

`func (o *VideoChannel) GetAvatars() []ActorImage`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *VideoChannel) GetAvatarsOk() (*[]ActorImage, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *VideoChannel) SetAvatars(v []ActorImage)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *VideoChannel) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetHost

`func (o *VideoChannel) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VideoChannel) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VideoChannel) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VideoChannel) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHostRedundancyAllowed

`func (o *VideoChannel) GetHostRedundancyAllowed() bool`

GetHostRedundancyAllowed returns the HostRedundancyAllowed field if non-nil, zero value otherwise.

### GetHostRedundancyAllowedOk

`func (o *VideoChannel) GetHostRedundancyAllowedOk() (*bool, bool)`

GetHostRedundancyAllowedOk returns a tuple with the HostRedundancyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRedundancyAllowed

`func (o *VideoChannel) SetHostRedundancyAllowed(v bool)`

SetHostRedundancyAllowed sets HostRedundancyAllowed field to given value.

### HasHostRedundancyAllowed

`func (o *VideoChannel) HasHostRedundancyAllowed() bool`

HasHostRedundancyAllowed returns a boolean if a field has been set.

### GetFollowingCount

`func (o *VideoChannel) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *VideoChannel) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *VideoChannel) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *VideoChannel) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.

### GetFollowersCount

`func (o *VideoChannel) GetFollowersCount() int32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *VideoChannel) GetFollowersCountOk() (*int32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *VideoChannel) SetFollowersCount(v int32)`

SetFollowersCount sets FollowersCount field to given value.

### HasFollowersCount

`func (o *VideoChannel) HasFollowersCount() bool`

HasFollowersCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoChannel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoChannel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoChannel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoChannel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoChannel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoChannel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoChannel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoChannel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *VideoChannel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VideoChannel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VideoChannel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VideoChannel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *VideoChannel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoChannel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoChannel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoChannel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VideoChannel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VideoChannel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSupport

`func (o *VideoChannel) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *VideoChannel) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *VideoChannel) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *VideoChannel) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### SetSupportNil

`func (o *VideoChannel) SetSupportNil(b bool)`

 SetSupportNil sets the value for Support to be an explicit nil

### UnsetSupport
`func (o *VideoChannel) UnsetSupport()`

UnsetSupport ensures that no value is present for Support, not even an explicit nil
### GetIsLocal

`func (o *VideoChannel) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *VideoChannel) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *VideoChannel) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *VideoChannel) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetBanners

`func (o *VideoChannel) GetBanners() []ActorImage`

GetBanners returns the Banners field if non-nil, zero value otherwise.

### GetBannersOk

`func (o *VideoChannel) GetBannersOk() (*[]ActorImage, bool)`

GetBannersOk returns a tuple with the Banners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanners

`func (o *VideoChannel) SetBanners(v []ActorImage)`

SetBanners sets Banners field to given value.

### HasBanners

`func (o *VideoChannel) HasBanners() bool`

HasBanners returns a boolean if a field has been set.

### GetOwnerAccount

`func (o *VideoChannel) GetOwnerAccount() Account`

GetOwnerAccount returns the OwnerAccount field if non-nil, zero value otherwise.

### GetOwnerAccountOk

`func (o *VideoChannel) GetOwnerAccountOk() (*Account, bool)`

GetOwnerAccountOk returns a tuple with the OwnerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAccount

`func (o *VideoChannel) SetOwnerAccount(v Account)`

SetOwnerAccount sets OwnerAccount field to given value.

### HasOwnerAccount

`func (o *VideoChannel) HasOwnerAccount() bool`

HasOwnerAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


