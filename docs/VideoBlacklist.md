# VideoBlacklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**VideoId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **int32** |  | [optional] 
**Views** | Pointer to **int32** |  | [optional] 
**Likes** | Pointer to **int32** |  | [optional] 
**Dislikes** | Pointer to **int32** |  | [optional] 
**Nsfw** | Pointer to **bool** |  | [optional] 

## Methods

### NewVideoBlacklist

`func NewVideoBlacklist() *VideoBlacklist`

NewVideoBlacklist instantiates a new VideoBlacklist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoBlacklistWithDefaults

`func NewVideoBlacklistWithDefaults() *VideoBlacklist`

NewVideoBlacklistWithDefaults instantiates a new VideoBlacklist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoBlacklist) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoBlacklist) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoBlacklist) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VideoBlacklist) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVideoId

`func (o *VideoBlacklist) GetVideoId() int32`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *VideoBlacklist) GetVideoIdOk() (*int32, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *VideoBlacklist) SetVideoId(v int32)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *VideoBlacklist) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VideoBlacklist) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VideoBlacklist) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VideoBlacklist) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VideoBlacklist) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VideoBlacklist) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VideoBlacklist) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VideoBlacklist) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VideoBlacklist) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *VideoBlacklist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VideoBlacklist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VideoBlacklist) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VideoBlacklist) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *VideoBlacklist) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VideoBlacklist) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VideoBlacklist) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VideoBlacklist) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *VideoBlacklist) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoBlacklist) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoBlacklist) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoBlacklist) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDuration

`func (o *VideoBlacklist) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoBlacklist) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoBlacklist) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoBlacklist) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetViews

`func (o *VideoBlacklist) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *VideoBlacklist) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *VideoBlacklist) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *VideoBlacklist) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetLikes

`func (o *VideoBlacklist) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *VideoBlacklist) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *VideoBlacklist) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *VideoBlacklist) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetDislikes

`func (o *VideoBlacklist) GetDislikes() int32`

GetDislikes returns the Dislikes field if non-nil, zero value otherwise.

### GetDislikesOk

`func (o *VideoBlacklist) GetDislikesOk() (*int32, bool)`

GetDislikesOk returns a tuple with the Dislikes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikes

`func (o *VideoBlacklist) SetDislikes(v int32)`

SetDislikes sets Dislikes field to given value.

### HasDislikes

`func (o *VideoBlacklist) HasDislikes() bool`

HasDislikes returns a boolean if a field has been set.

### GetNsfw

`func (o *VideoBlacklist) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *VideoBlacklist) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *VideoBlacklist) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *VideoBlacklist) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


