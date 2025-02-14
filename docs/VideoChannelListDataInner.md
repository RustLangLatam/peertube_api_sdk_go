# VideoChannelListDataInner

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
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DisplayName** | Pointer to **string** | editable name of the channel, displayed in its representations | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Support** | Pointer to **NullableString** | text shown by default on all videos of this channel, to tell the audience how to support it | [optional] 
**IsLocal** | Pointer to **bool** |  | [optional] [readonly] 
**Banners** | Pointer to [**[]ActorImage**](ActorImage.md) |  | [optional] 
**OwnerAccount** | Pointer to [**Account**](Account.md) |  | [optional] 

## Methods

### NewVideoChannelListDataInner

`func NewVideoChannelListDataInner() *VideoChannelListDataInner`

NewVideoChannelListDataInner instantiates a new VideoChannelListDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoChannelListDataInnerWithDefaults

`func NewVideoChannelListDataInnerWithDefaults() *VideoChannelListDataInner`

NewVideoChannelListDataInnerWithDefaults instantiates a new VideoChannelListDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoChannelListDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoChannelListDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoChannelListDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoChannelListDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *VideoChannelListDataInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VideoChannelListDataInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VideoChannelListDataInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *VideoChannelListDataInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *VideoChannelListDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoChannelListDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoChannelListDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoChannelListDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatars

`func (o *VideoChannelListDataInner) GetAvatars() []ActorImage`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *VideoChannelListDataInner) GetAvatarsOk() (*[]ActorImage, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *VideoChannelListDataInner) SetAvatars(v []ActorImage)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *VideoChannelListDataInner) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetHost

`func (o *VideoChannelListDataInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VideoChannelListDataInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VideoChannelListDataInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VideoChannelListDataInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHostRedundancyAllowed

`func (o *VideoChannelListDataInner) GetHostRedundancyAllowed() bool`

GetHostRedundancyAllowed returns the HostRedundancyAllowed field if non-nil, zero value otherwise.

### GetHostRedundancyAllowedOk

`func (o *VideoChannelListDataInner) GetHostRedundancyAllowedOk() (*bool, bool)`

GetHostRedundancyAllowedOk returns a tuple with the HostRedundancyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRedundancyAllowed

`func (o *VideoChannelListDataInner) SetHostRedundancyAllowed(v bool)`

SetHostRedundancyAllowed sets HostRedundancyAllowed field to given value.

### HasHostRedundancyAllowed

`func (o *VideoChannelListDataInner) HasHostRedundancyAllowed() bool`

HasHostRedundancyAllowed returns a boolean if a field has been set.

### GetFollowingCount

`func (o *VideoChannelListDataInner) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *VideoChannelListDataInner) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *VideoChannelListDataInner) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *VideoChannelListDataInner) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.

### GetFollowersCount

`func (o *VideoChannelListDataInner) GetFollowersCount() int32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *VideoChannelListDataInner) GetFollowersCountOk() (*int32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *VideoChannelListDataInner) SetFollowersCount(v int32)`

SetFollowersCount sets FollowersCount field to given value.

### HasFollowersCount

`func (o *VideoChannelListDataInner) HasFollowersCount() bool`

HasFollowersCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoChannelListDataInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoChannelListDataInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoChannelListDataInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoChannelListDataInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoChannelListDataInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoChannelListDataInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoChannelListDataInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoChannelListDataInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *VideoChannelListDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VideoChannelListDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VideoChannelListDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VideoChannelListDataInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *VideoChannelListDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoChannelListDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoChannelListDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoChannelListDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VideoChannelListDataInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VideoChannelListDataInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSupport

`func (o *VideoChannelListDataInner) GetSupport() string`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *VideoChannelListDataInner) GetSupportOk() (*string, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *VideoChannelListDataInner) SetSupport(v string)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *VideoChannelListDataInner) HasSupport() bool`

HasSupport returns a boolean if a field has been set.

### SetSupportNil

`func (o *VideoChannelListDataInner) SetSupportNil(b bool)`

 SetSupportNil sets the value for Support to be an explicit nil

### UnsetSupport
`func (o *VideoChannelListDataInner) UnsetSupport()`

UnsetSupport ensures that no value is present for Support, not even an explicit nil
### GetIsLocal

`func (o *VideoChannelListDataInner) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *VideoChannelListDataInner) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *VideoChannelListDataInner) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *VideoChannelListDataInner) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetBanners

`func (o *VideoChannelListDataInner) GetBanners() []ActorImage`

GetBanners returns the Banners field if non-nil, zero value otherwise.

### GetBannersOk

`func (o *VideoChannelListDataInner) GetBannersOk() (*[]ActorImage, bool)`

GetBannersOk returns a tuple with the Banners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanners

`func (o *VideoChannelListDataInner) SetBanners(v []ActorImage)`

SetBanners sets Banners field to given value.

### HasBanners

`func (o *VideoChannelListDataInner) HasBanners() bool`

HasBanners returns a boolean if a field has been set.

### GetOwnerAccount

`func (o *VideoChannelListDataInner) GetOwnerAccount() Account`

GetOwnerAccount returns the OwnerAccount field if non-nil, zero value otherwise.

### GetOwnerAccountOk

`func (o *VideoChannelListDataInner) GetOwnerAccountOk() (*Account, bool)`

GetOwnerAccountOk returns a tuple with the OwnerAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAccount

`func (o *VideoChannelListDataInner) SetOwnerAccount(v Account)`

SetOwnerAccount sets OwnerAccount field to given value.

### HasOwnerAccount

`func (o *VideoChannelListDataInner) HasOwnerAccount() bool`

HasOwnerAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


