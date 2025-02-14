# NotificationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**[]Notification**](Notification.md) |  | [optional] 

## Methods

### NewNotificationListResponse

`func NewNotificationListResponse() *NotificationListResponse`

NewNotificationListResponse instantiates a new NotificationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationListResponseWithDefaults

`func NewNotificationListResponseWithDefaults() *NotificationListResponse`

NewNotificationListResponseWithDefaults instantiates a new NotificationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *NotificationListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NotificationListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NotificationListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NotificationListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetData

`func (o *NotificationListResponse) GetData() []Notification`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NotificationListResponse) GetDataOk() (*[]Notification, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NotificationListResponse) SetData(v []Notification)`

SetData sets Data field to given value.

### HasData

`func (o *NotificationListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


