# NotificationVideoAbuse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Video** | Pointer to [**VideoInfo**](VideoInfo.md) |  | [optional] 

## Methods

### NewNotificationVideoAbuse

`func NewNotificationVideoAbuse() *NotificationVideoAbuse`

NewNotificationVideoAbuse instantiates a new NotificationVideoAbuse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationVideoAbuseWithDefaults

`func NewNotificationVideoAbuseWithDefaults() *NotificationVideoAbuse`

NewNotificationVideoAbuseWithDefaults instantiates a new NotificationVideoAbuse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationVideoAbuse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationVideoAbuse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationVideoAbuse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationVideoAbuse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVideo

`func (o *NotificationVideoAbuse) GetVideo() VideoInfo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *NotificationVideoAbuse) GetVideoOk() (*VideoInfo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *NotificationVideoAbuse) SetVideo(v VideoInfo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *NotificationVideoAbuse) HasVideo() bool`

HasVideo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


