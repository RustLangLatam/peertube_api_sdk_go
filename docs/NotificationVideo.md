# NotificationVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | title of the video | [optional] 
**State** | Pointer to [**VideoStateConstant**](VideoStateConstant.md) | represents the internal state of the video processing within the PeerTube instance | [optional] 
**Channel** | Pointer to [**ActorInfo**](ActorInfo.md) |  | [optional] 

## Methods

### NewNotificationVideo

`func NewNotificationVideo() *NotificationVideo`

NewNotificationVideo instantiates a new NotificationVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationVideoWithDefaults

`func NewNotificationVideoWithDefaults() *NotificationVideo`

NewNotificationVideoWithDefaults instantiates a new NotificationVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationVideo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationVideo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationVideo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationVideo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *NotificationVideo) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *NotificationVideo) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *NotificationVideo) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *NotificationVideo) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *NotificationVideo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationVideo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationVideo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationVideo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *NotificationVideo) GetState() VideoStateConstant`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NotificationVideo) GetStateOk() (*VideoStateConstant, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NotificationVideo) SetState(v VideoStateConstant)`

SetState sets State field to given value.

### HasState

`func (o *NotificationVideo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChannel

`func (o *NotificationVideo) GetChannel() ActorInfo`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *NotificationVideo) GetChannelOk() (*ActorInfo, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *NotificationVideo) SetChannel(v ActorInfo)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *NotificationVideo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


