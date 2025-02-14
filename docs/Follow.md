# Follow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Follower** | Pointer to [**Actor**](Actor.md) |  | [optional] 
**Following** | Pointer to [**Actor**](Actor.md) |  | [optional] 
**Score** | Pointer to **float32** | score reflecting the reachability of the actor, with steps of &#x60;10&#x60; and a base score of &#x60;1000&#x60;. | [optional] 
**State** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewFollow

`func NewFollow() *Follow`

NewFollow instantiates a new Follow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowWithDefaults

`func NewFollowWithDefaults() *Follow`

NewFollowWithDefaults instantiates a new Follow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Follow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Follow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Follow) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Follow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFollower

`func (o *Follow) GetFollower() Actor`

GetFollower returns the Follower field if non-nil, zero value otherwise.

### GetFollowerOk

`func (o *Follow) GetFollowerOk() (*Actor, bool)`

GetFollowerOk returns a tuple with the Follower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollower

`func (o *Follow) SetFollower(v Actor)`

SetFollower sets Follower field to given value.

### HasFollower

`func (o *Follow) HasFollower() bool`

HasFollower returns a boolean if a field has been set.

### GetFollowing

`func (o *Follow) GetFollowing() Actor`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *Follow) GetFollowingOk() (*Actor, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *Follow) SetFollowing(v Actor)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *Follow) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetScore

`func (o *Follow) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Follow) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Follow) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Follow) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetState

`func (o *Follow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Follow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Follow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Follow) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Follow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Follow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Follow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Follow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Follow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Follow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Follow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Follow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


