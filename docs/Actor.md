# Actor

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

## Methods

### NewActor

`func NewActor() *Actor`

NewActor instantiates a new Actor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorWithDefaults

`func NewActorWithDefaults() *Actor`

NewActorWithDefaults instantiates a new Actor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Actor) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Actor) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Actor) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Actor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *Actor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Actor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Actor) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Actor) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *Actor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Actor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Actor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Actor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvatars

`func (o *Actor) GetAvatars() []ActorImage`

GetAvatars returns the Avatars field if non-nil, zero value otherwise.

### GetAvatarsOk

`func (o *Actor) GetAvatarsOk() (*[]ActorImage, bool)`

GetAvatarsOk returns a tuple with the Avatars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatars

`func (o *Actor) SetAvatars(v []ActorImage)`

SetAvatars sets Avatars field to given value.

### HasAvatars

`func (o *Actor) HasAvatars() bool`

HasAvatars returns a boolean if a field has been set.

### GetHost

`func (o *Actor) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Actor) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Actor) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *Actor) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHostRedundancyAllowed

`func (o *Actor) GetHostRedundancyAllowed() bool`

GetHostRedundancyAllowed returns the HostRedundancyAllowed field if non-nil, zero value otherwise.

### GetHostRedundancyAllowedOk

`func (o *Actor) GetHostRedundancyAllowedOk() (*bool, bool)`

GetHostRedundancyAllowedOk returns a tuple with the HostRedundancyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostRedundancyAllowed

`func (o *Actor) SetHostRedundancyAllowed(v bool)`

SetHostRedundancyAllowed sets HostRedundancyAllowed field to given value.

### HasHostRedundancyAllowed

`func (o *Actor) HasHostRedundancyAllowed() bool`

HasHostRedundancyAllowed returns a boolean if a field has been set.

### GetFollowingCount

`func (o *Actor) GetFollowingCount() int32`

GetFollowingCount returns the FollowingCount field if non-nil, zero value otherwise.

### GetFollowingCountOk

`func (o *Actor) GetFollowingCountOk() (*int32, bool)`

GetFollowingCountOk returns a tuple with the FollowingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowingCount

`func (o *Actor) SetFollowingCount(v int32)`

SetFollowingCount sets FollowingCount field to given value.

### HasFollowingCount

`func (o *Actor) HasFollowingCount() bool`

HasFollowingCount returns a boolean if a field has been set.

### GetFollowersCount

`func (o *Actor) GetFollowersCount() int32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *Actor) GetFollowersCountOk() (*int32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *Actor) SetFollowersCount(v int32)`

SetFollowersCount sets FollowersCount field to given value.

### HasFollowersCount

`func (o *Actor) HasFollowersCount() bool`

HasFollowersCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Actor) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Actor) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Actor) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Actor) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Actor) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Actor) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Actor) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Actor) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


