# NotificationActorFollow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Follower** | Pointer to [**ActorInfo**](ActorInfo.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Following** | Pointer to [**NotificationActorFollowFollowing**](NotificationActorFollowFollowing.md) |  | [optional] 

## Methods

### NewNotificationActorFollow

`func NewNotificationActorFollow() *NotificationActorFollow`

NewNotificationActorFollow instantiates a new NotificationActorFollow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationActorFollowWithDefaults

`func NewNotificationActorFollowWithDefaults() *NotificationActorFollow`

NewNotificationActorFollowWithDefaults instantiates a new NotificationActorFollow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationActorFollow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationActorFollow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationActorFollow) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationActorFollow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFollower

`func (o *NotificationActorFollow) GetFollower() ActorInfo`

GetFollower returns the Follower field if non-nil, zero value otherwise.

### GetFollowerOk

`func (o *NotificationActorFollow) GetFollowerOk() (*ActorInfo, bool)`

GetFollowerOk returns a tuple with the Follower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollower

`func (o *NotificationActorFollow) SetFollower(v ActorInfo)`

SetFollower sets Follower field to given value.

### HasFollower

`func (o *NotificationActorFollow) HasFollower() bool`

HasFollower returns a boolean if a field has been set.

### GetState

`func (o *NotificationActorFollow) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NotificationActorFollow) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NotificationActorFollow) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NotificationActorFollow) HasState() bool`

HasState returns a boolean if a field has been set.

### GetFollowing

`func (o *NotificationActorFollow) GetFollowing() NotificationActorFollowFollowing`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *NotificationActorFollow) GetFollowingOk() (*NotificationActorFollowFollowing, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *NotificationActorFollow) SetFollowing(v NotificationActorFollowFollowing)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *NotificationActorFollow) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


